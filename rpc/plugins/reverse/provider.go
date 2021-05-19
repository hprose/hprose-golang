/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/reverse/provider.go                          |
|                                                          |
| LastModified: May 19, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package reverse

import (
	"context"
	"reflect"
	"sync"
	"sync/atomic"
	"time"

	"github.com/hprose/hprose-golang/v3/io"
	"github.com/hprose/hprose-golang/v3/rpc/core"
)

type contextMissingMethod = func(ctx context.Context, name string, args []interface{}) (result []interface{}, err error)
type missingMethod = func(name string, args []interface{}) (result []interface{}, err error)

type ProviderContext struct {
	core.Context
	client *core.Client
	method core.Method
}

func NewProviderContext(client *core.Client, method core.Method) *ProviderContext {
	return &ProviderContext{
		client: client,
		method: method,
	}
}

func (c *ProviderContext) Client() *core.Client {
	return c.client
}

func (c *ProviderContext) Method() core.Method {
	return c.method
}

func (c *ProviderContext) Clone() core.Context {
	return &ProviderContext{
		c.Context.Clone(),
		c.client,
		c.method,
	}
}

// GetProviderContext returns the *reverse.ProviderContext bound to the context.
func GetProviderContext(ctx context.Context) *ProviderContext {
	if c, ok := core.FromContext(ctx); ok {
		return c.(*ProviderContext)
	}
	return nil
}

type Provider struct {
	client        *core.Client
	proxy         provider
	invokeManager core.PluginManager
	methodManager core.MethodManager
	closed        int32
	RetryInterval time.Duration
	OnError       func(error)
	Debug         bool
}

type provider struct {
	close func() error                      `name:"!!"`
	begin func() ([]call, error)            `name:"!"`
	end   func(results []returnValue) error `name:"="`
}

func NewProvider(client *core.Client, id ...string) *Provider {
	p := &Provider{
		client:        client,
		RetryInterval: time.Second,
		closed:        1,
	}
	if len(id) > 0 && id[0] != "" {
		p.SetID(id[0])
	}
	p.client.UseService(&p.proxy)
	p.invokeManager = core.NewInvokeManager(p.Execute)
	p.methodManager = core.NewMethodManager()
	p.AddFunction(p.methodManager.Names, "~")
	return p
}

func (p *Provider) onError(err error) {
	if p.OnError != nil {
		p.OnError(err)
	}
}

func (p *Provider) Client() *core.Client {
	return p.client
}

func (p *Provider) ID() (id string) {
	if id = p.client.RequestHeaders().GetString("id"); id == "" {
		panic("client unique id not found")
	}
	return
}

func (p *Provider) SetID(id string) {
	p.client.RequestHeaders().Set("id", id)
}

func (p *Provider) Execute(ctx context.Context, name string, args []interface{}) (result []interface{}, err error) {
	method := GetProviderContext(ctx).method
	if method.Missing() {
		if method.PassContext() {
			return method.(interface{}).(contextMissingMethod)(ctx, name, args)
		}
		return method.(interface{}).(missingMethod)(name, args)
	}
	n := len(args)
	var in []reflect.Value
	if method.PassContext() {
		in = make([]reflect.Value, n+1)
		in[0] = reflect.ValueOf(ctx)
		for i := 0; i < n; i++ {
			in[i+1] = reflect.ValueOf(args[i])
		}
	} else {
		in = make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			in[i] = reflect.ValueOf(args[i])
		}
	}
	f := method.Func()
	out := f.Call(in)
	n = len(out)
	if method.ReturnError() {
		if !out[n-1].IsNil() {
			err = out[n-1].Interface().(error)
		}
		out = out[:n-1]
		n--
	}
	for i := 0; i < n; i++ {
		result = append(result, out[i].Interface())
	}
	return
}

func (p *Provider) process(c call) (rv returnValue) {
	index, name, args := c.Value()
	defer func() {
		if e := recover(); e != nil {
			err := core.NewPanicError(e)
			if p.Debug {
				rv = newReturnValue(index, nil, err.String())
			} else {
				rv = newReturnValue(index, nil, err.Error())
			}
		}
	}()
	method := p.Get(name)
	if method == nil {
		return newReturnValue(index, nil, "Can't find this method "+name+"().")
	}
	if !method.Missing() {
		count := len(args)
		parameters := method.Parameters()
		paramTypes := make([]reflect.Type, count)
		if method.Func().Type().IsVariadic() {
			n := len(parameters)
			copy(paramTypes, parameters[:n-1])
			for i := n - 1; i < count; i++ {
				paramTypes[i] = parameters[n-1].Elem()
			}
		} else {
			copy(paramTypes, parameters)
		}
		for i, t := range paramTypes {
			if arg, err := io.Convert(args[i], t); err != nil {
				return newReturnValue(index, nil, err.Error())
			} else {
				args[i] = arg
			}
		}
	}
	ctx := core.WithContext(context.Background(), NewProviderContext(p.client, method))
	results, err := p.invokeManager.Handler().(core.NextInvokeHandler)(ctx, name, args)
	var result interface{}
	switch len(results) {
	case 0:
		result = nil
	case 1:
		result = results[0]
	default:
		result = results
	}
	if err != nil {
		return newReturnValue(index, result, err.Error())
	}
	return newReturnValue(index, result, "")
}

func (p *Provider) dispatch(calls []call) {
	n := len(calls)
	results := make([]returnValue, n)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			results[i] = p.process(calls[i])
			wg.Done()
		}(i)
	}
	wg.Wait()
	for {
		if err := p.proxy.end(results); err != nil {
			if !core.IsTimeoutError(err) {
				if p.RetryInterval != 0 {
					<-time.After(p.RetryInterval)
				}
				p.onError(err)
			}
			continue
		}
		return
	}
}

func (p *Provider) Listen() {
	if !atomic.CompareAndSwapInt32(&p.closed, 1, 0) {
		return
	}
	for atomic.LoadInt32(&p.closed) == 0 {
		calls, err := p.proxy.begin()
		if err != nil {
			if !core.IsTimeoutError(err) {
				if p.RetryInterval != 0 {
					<-time.After(p.RetryInterval)
				}
				p.onError(err)
			}
			continue
		}
		if calls == nil {
			return
		}
		go p.dispatch(calls)
	}
	atomic.StoreInt32(&p.closed, 1)
}

func (p *Provider) Close() error {
	if atomic.CompareAndSwapInt32(&p.closed, 0, 1) {
		return p.proxy.close()
	}
	return core.ErrClosed
}

// Use plugin handlers.
func (p *Provider) Use(handler ...core.PluginHandler) *Provider {
	invokeHandlers, _ := core.SeparatePluginHandlers(handler)
	if len(invokeHandlers) > 0 {
		p.invokeManager.Use(invokeHandlers...)
	}
	return p
}

// Unuse plugin handlers.
func (p *Provider) Unuse(handler ...core.PluginHandler) *Provider {
	invokeHandlers, _ := core.SeparatePluginHandlers(handler)
	if len(invokeHandlers) > 0 {
		p.invokeManager.Unuse(invokeHandlers...)
	}
	return p
}

// Get returns the published method by name.
func (p *Provider) Get(name string) core.Method {
	return p.methodManager.Get(name)
}

// Remove is used for unpublishing method by the specified name.
func (p *Provider) Remove(name string) *Provider {
	p.methodManager.Remove(name)
	return p
}

// Add is used for publishing the method.
func (p *Provider) Add(method core.Method) *Provider {
	p.methodManager.Add(method)
	return p
}

// AddFunction is used for publishing function f with alias.
func (p *Provider) AddFunction(f interface{}, alias ...string) *Provider {
	p.methodManager.AddFunction(f, alias...)
	return p
}

// AddMethod is used for publishing method named name on target with alias.
func (p *Provider) AddMethod(name string, target interface{}, alias ...string) *Provider {
	p.methodManager.AddMethod(name, target, alias...)
	return p
}

// AddMethods is used for publishing methods named names on target with namespace.
func (p *Provider) AddMethods(names []string, target interface{}, namespace ...string) *Provider {
	p.methodManager.AddMethods(names, target, namespace...)
	return p
}

// AddInstanceMethods is used for publishing all the public methods and func fields with namespace.
func (p *Provider) AddInstanceMethods(target interface{}, namespace ...string) *Provider {
	p.methodManager.AddInstanceMethods(target, namespace...)
	return p
}

// AddAllMethods will publish all methods and non-nil function fields on the
// obj self and on its anonymous or non-anonymous struct fields (or pointer to
// pointer ... to pointer struct fields). This is a recursive operation.
// So it's a pit, if you do not know what you are doing, do not step on.
func (p *Provider) AddAllMethods(target interface{}, namespace ...string) *Provider {
	p.methodManager.AddAllMethods(target, namespace...)
	return p
}

// AddMissingMethod is used for publishing a method,
// all methods not explicitly published will be redirected to this method.
func (p *Provider) AddMissingMethod(f interface{}) *Provider {
	p.methodManager.AddMissingMethod(f)
	return p
}

// AddNetRPCMethods is used for publishing methods defined for net/rpc.
func (p *Provider) AddNetRPCMethods(rcvr interface{}, namespace ...string) *Provider {
	p.methodManager.AddNetRPCMethods(rcvr, namespace...)
	return p
}
