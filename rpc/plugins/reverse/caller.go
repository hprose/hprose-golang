/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/reverse/caller.go                            |
|                                                          |
| LastModified: May 19, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package reverse

import (
	"context"
	"errors"
	"reflect"
	"sync"
	"sync/atomic"
	"time"

	"github.com/hprose/hprose-golang/v3/io"
	"github.com/hprose/hprose-golang/v3/rpc/core"
	"github.com/modern-go/reflect2"
	cmap "github.com/orcaman/concurrent-map"
)

type call [3]interface{}

func newCall(index int, name string, args []interface{}) (c call) {
	c[0] = index
	c[1] = name
	c[2] = args
	return
}

func (c call) Value() (index int, name string, args []interface{}) {
	return c[0].(int), c[1].(string), c[2].([]interface{})
}

type callCache struct {
	c []call
	sync.Mutex
}

func (cc *callCache) Append(c call) {
	cc.Lock()
	defer cc.Unlock()
	cc.c = append(cc.c, c)
}

func (cc *callCache) Delete(index int) {
	cc.Lock()
	defer cc.Unlock()
	for i := 0; i < len(cc.c); i++ {
		if cc.c[i][0].(int) == index {
			cc.c = append(cc.c[:i], cc.c[i+1:]...)
			return
		}
	}
}

func (cc *callCache) Take() (calls []call) {
	cc.Lock()
	defer cc.Unlock()
	calls = cc.c
	cc.c = nil
	return
}

type returnValue [3]interface{}

func newReturnValue(index int, result interface{}, err string) (r returnValue) {
	r[0] = index
	r[1] = result
	r[2] = err
	return
}

func (r returnValue) Index() int {
	return r[0].(int)
}

func (r returnValue) Value(returnType []reflect.Type) ([]interface{}, error) {
	err := r[2].(string)
	if err != "" {
		return nil, errors.New(err)
	}
	n := len(returnType)
	switch n {
	case 0:
		return nil, nil
	case 1:
		if result, err := io.Convert(r[1], returnType[0]); err != nil {
			return nil, err
		} else {
			return []interface{}{result}, nil
		}
	default:
		results := make([]interface{}, n)
		values := r[1].([]interface{})
		count := len(values)
		for i := 0; i < n && i < count; i++ {
			if result, err := io.Convert(values[i], returnType[i]); err != nil {
				return nil, err
			} else {
				results[i] = result
			}
		}
		for i := count; i < n; i++ {
			t := reflect2.Type2(returnType[i])
			results[i] = t.Indirect(t.New())
		}
		return results, nil
	}
}

type resultMap struct {
	results map[int]chan returnValue
	sync.Mutex
}

func newResultMap() *resultMap {
	return &resultMap{
		results: make(map[int]chan returnValue),
	}
}

func (m *resultMap) GetAndDelete(index int) chan returnValue {
	m.Lock()
	defer m.Unlock()
	if result, ok := m.results[index]; ok {
		delete(m.results, index)
		return result
	}
	return nil
}

func (m *resultMap) Delete(index int) {
	m.Lock()
	defer m.Unlock()
	delete(m.results, index)
}

func (m *resultMap) Set(index int, result chan returnValue) {
	m.Lock()
	defer m.Unlock()
	m.results[index] = result
}

var (
	emptyArgs = make([]interface{}, 0)
	emptyCall = make([]call, 0)
)

type Caller struct {
	*core.Service
	IdleTimeout time.Duration
	Timeout     time.Duration
	HeartBeat   time.Duration
	calls       cmap.ConcurrentMap
	results     cmap.ConcurrentMap
	responders  cmap.ConcurrentMap
	onlines     cmap.ConcurrentMap
	counter     int32
}

func NewCaller(service *core.Service) *Caller {
	caller := &Caller{
		Service:     service,
		IdleTimeout: time.Minute * 2,
		Timeout:     time.Second * 30,
		HeartBeat:   time.Second * 3,
		calls:       cmap.New(),
		results:     cmap.New(),
		responders:  cmap.New(),
		onlines:     cmap.New(),
	}
	service.Use(caller.handler).
		AddFunction(caller.close, "!!").
		AddFunction(caller.begin, "!").
		AddFunction(caller.end, "=")
	return caller
}

func (c *Caller) ID(ctx context.Context) (id string) {
	if id = core.GetServiceContext(ctx).RequestHeaders().GetString("id"); id == "" {
		panic("client unique id not found")
	}
	return
}

func (c *Caller) send(id string, responder chan []call) bool {
	if calls, ok := c.calls.Get(id); ok {
		calls := calls.(*callCache).Take()
		if len(calls) == 0 {
			return false
		}
		responder <- calls
		return true
	}
	return false
}

func (c *Caller) response(id string) {
	if responder, ok := c.responders.Pop(id); ok {
		responder := responder.(chan []call)
		if !c.send(id, responder) {
			if !c.responders.SetIfAbsent(id, responder) {
				responder <- nil
			}
		}
	}
}

func (c *Caller) stop(ctx context.Context) string {
	id := c.ID(ctx)
	if responder, ok := c.responders.Pop(id); ok {
		responder.(chan []call) <- nil
	}
	return id
}

func (c *Caller) close(ctx context.Context) {
	id := c.stop(ctx)
	c.onlines.Remove(id)
}

func (c *Caller) begin(ctx context.Context) []call {
	id := c.stop(ctx)
	online := make(chan bool, 1)
	c.onlines.Upsert(id, online, func(exist bool, valueInMap, newValue interface{}) interface{} {
		if exist {
			if online, ok := valueInMap.(chan bool); ok {
				online <- true
			}
		}
		return newValue
	})
	if c.HeartBeat > 0 {
		defer func() {
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), c.HeartBeat)
				defer cancel()
				select {
				case <-ctx.Done():
					c.onlines.Remove(id)
				case <-online:
				}
			}()
		}()
	}
	responder := make(chan []call, 1)
	if !c.send(id, responder) {
		c.responders.Upsert(id, responder, func(exist bool, valueInMap interface{}, newValue interface{}) interface{} {
			if exist {
				valueInMap.(chan []call) <- nil
			}
			return newValue
		})
		if c.IdleTimeout > 0 {
			ctx, cancel := context.WithTimeout(ctx, c.IdleTimeout)
			defer cancel()
			select {
			case <-ctx.Done():
				responder <- emptyCall
			case result := <-responder:
				return result
			}
		}
	}
	return <-responder
}

func (c *Caller) end(ctx context.Context, results []returnValue) {
	id := c.ID(ctx)
	for _, rv := range results {
		if r, ok := c.results.Get(id); ok {
			if value := r.(*resultMap).GetAndDelete(rv.Index()); value != nil {
				value <- rv
			}
		}
	}
}

func (c *Caller) Invoke(id string, name string, args []interface{}, returnType ...reflect.Type) ([]interface{}, error) {
	return c.InvokeContext(context.Background(), id, name, args, returnType...)
}

func (c *Caller) InvokeContext(ctx context.Context, id string, name string, args []interface{}, returnType ...reflect.Type) ([]interface{}, error) {
	if args == nil {
		args = emptyArgs
	}
	index := int(atomic.AddInt32(&c.counter, 1) & 0x7fffffff)
	var calls *callCache
	if cc, ok := c.calls.Get(id); ok {
		calls = cc.(*callCache)
	} else {
		calls = new(callCache)
		if !c.calls.SetIfAbsent(id, calls) {
			cc, _ := c.calls.Get(id)
			calls = cc.(*callCache)
		}
	}
	calls.Append(newCall(index, name, args))
	var results *resultMap
	if rm, ok := c.results.Get(id); ok {
		results = rm.(*resultMap)
	} else {
		results = newResultMap()
		if !c.results.SetIfAbsent(id, results) {
			rm, _ := c.results.Get(id)
			results = rm.(*resultMap)
		}
	}
	result := make(chan returnValue, 1)
	results.Set(index, result)
	c.response(id)
	if c.Timeout > 0 {
		ctx, cancel := context.WithTimeout(ctx, c.Timeout)
		defer cancel()
		select {
		case <-ctx.Done():
			calls.Delete(index)
			results.Delete(index)
			return nil, core.ErrTimeout
		case result := <-result:
			return result.Value(returnType)
		}
	}
	return (<-result).Value(returnType)
}

func (c *Caller) UseService(remoteService interface{}, id string, namespace ...string) {
	ns := ""
	if len(namespace) > 0 {
		ns = namespace[0]
	}
	core.Proxy.Build(remoteService, invocation{caller: c, id: id, namespace: ns}.Invoke)
}

func (c *Caller) Exists(id string) bool {
	return c.onlines.Has(id)
}

func (c *Caller) IdList() []string {
	return c.onlines.Keys()
}

func (c *Caller) handler(ctx context.Context, name string, args []interface{}, next core.NextInvokeHandler) (result []interface{}, err error) {
	core.GetServiceContext(ctx).Items().Set("caller", c)
	return next(ctx, name, args)
}

func UseService(ctx context.Context, remoteService interface{}, namespace ...string) *Caller {
	caller := core.GetServiceContext(ctx).Items().GetInterface("caller").(*Caller)
	caller.UseService(remoteService, caller.ID(ctx), namespace...)
	return caller
}
