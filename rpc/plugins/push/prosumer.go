/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/push/prosumer.go                             |
|                                                          |
| LastModified: May 16, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package push

import (
	"reflect"
	"sync"
	"sync/atomic"
	"time"

	"github.com/hprose/hprose-golang/v3/io"
	"github.com/hprose/hprose-golang/v3/rpc/core"
)

type Callback interface{}

type Prosumer struct {
	client        *core.Client
	proxy         prosumer
	callbacks     sync.Map // map[string]func(Message)
	RetryInterval time.Duration
	OnError       func(error)
	OnSubscribe   func(topic string)
	OnUnsubscribe func(topic string)
	loop          int32
}

type prosumer struct {
	message     func() (map[string][]Message, error)                                        `name:"<"`
	subscribe   func(topic string) (bool, error)                                            `name:"+"`
	unsubscribe func(topic string) (bool, error)                                            `name:"-"`
	unicast     func(data interface{}, topic string, id string) (bool, error)               `name:">"`
	multicast   func(data interface{}, topic string, ids []string) (map[string]bool, error) `name:">?"`
	broadcast   func(data interface{}, topic string) (map[string]bool, error)               `name:">*"`
	exists      func(topic string, id string) (bool, error)                                 `name:"?"`
	idList      func(topic string) ([]string, error)                                        `name:"|"`
}

func NewProsumer(client *core.Client, id ...string) *Prosumer {
	p := &Prosumer{
		client:        client,
		RetryInterval: time.Second,
	}
	if len(id) > 0 && id[0] != "" {
		p.SetID(id[0])
	}
	p.client.UseService(&p.proxy)
	return p
}

func (p *Prosumer) onError(err error) {
	if p.OnError != nil {
		p.OnError(err)
	}
}

func (p *Prosumer) onSubscribe(topic string) {
	if p.OnSubscribe != nil {
		p.OnSubscribe(topic)
	}
}

func (p *Prosumer) onUnsubscribe(topic string) {
	if p.OnUnsubscribe != nil {
		p.OnUnsubscribe(topic)
	}
}

func (p *Prosumer) Client() *core.Client {
	return p.client
}

func (p *Prosumer) ID() (id string) {
	if id = p.client.RequestHeaders().GetString("id"); id == "" {
		panic("client unique id not found")
	}
	return
}

func (p *Prosumer) SetID(id string) {
	p.client.RequestHeaders().Set("id", id)
}

func (p *Prosumer) dispatch(topics map[string][]Message) {
	for topic, messages := range topics {
		if callback, ok := p.callbacks.Load(topic); ok {
			if messages == nil {
				p.callbacks.Delete(topic)
				if p.OnUnsubscribe != nil {
					p.OnUnsubscribe(topic)
				}
			} else {
				for _, message := range messages {
					p.call(callback, message)
				}
			}
		}
	}
}

func (p *Prosumer) call(callback Callback, message Message) {
	switch callback := callback.(type) {
	case func(Message):
		callback(message)
	case func(data interface{}, from string):
		callback(message.Data, message.From)
	case func(data interface{}):
		callback(message.Data)
	default:
		v := reflect.ValueOf(callback)
		t := v.Type()
		if n := t.NumIn(); n >= 1 {
			data, err := io.Convert(message.Data, t.In(0))
			if err != nil {
				p.onError(err)
				return
			}
			switch n {
			case 1:
				v.Call([]reflect.Value{reflect.ValueOf(data)})
			case 2:
				v.Call([]reflect.Value{reflect.ValueOf(data), reflect.ValueOf(message.From)})
			default:
				panic("invalid callback: " + t.String())
			}
		}
	}
}

func (p *Prosumer) message() {
	if atomic.LoadInt32(&p.loop) == 1 {
		return
	}
	for {
		atomic.StoreInt32(&p.loop, 1)
		topics, err := p.proxy.message()
		if err != nil {
			if err != core.ErrTimeout {
				if p.RetryInterval != 0 {
					<-time.After(p.RetryInterval)
				}
				p.onError(err)
				p.callbacks.Range(func(key, value interface{}) bool {
					p.proxy.subscribe(key.(string))
					return true
				})
			}
			continue
		}
		if topics == nil {
			atomic.StoreInt32(&p.loop, 0)
			return
		}
		go p.dispatch(topics)
	}
}

func (p *Prosumer) Subscribe(topic string, callback Callback) (result bool, err error) {
	if p.ID() != "" {
		p.callbacks.Store(topic, callback)
		result, err = p.proxy.subscribe(topic)
		go p.message()
		p.onSubscribe(topic)
	}
	return
}

func (p *Prosumer) Unsubscribe(topic string) (result bool, err error) {
	if p.ID() != "" {
		result, err = p.proxy.unsubscribe(topic)
		p.callbacks.Delete(topic)
		p.onUnsubscribe(topic)
	}
	return
}

func (p *Prosumer) Unicast(data interface{}, topic string, id string) (bool, error) {
	return p.proxy.unicast(data, topic, id)
}

func (p *Prosumer) Multicast(data interface{}, topic string, ids []string) (map[string]bool, error) {
	return p.proxy.multicast(data, topic, ids)
}

func (p *Prosumer) Broadcast(data interface{}, topic string) (map[string]bool, error) {
	return p.proxy.broadcast(data, topic)
}

func (p *Prosumer) Push(data interface{}, topic string, id ...string) (map[string]bool, error) {
	switch len(id) {
	case 0:
		return p.Broadcast(data, topic)
	case 1:
		result, err := p.Unicast(data, topic, id[0])
		return map[string]bool{
			id[0]: result,
		}, err
	default:
		return p.Multicast(data, topic, id)
	}
}

func (p *Prosumer) Exists(topic string, id string) (bool, error) {
	if id == "" {
		id = p.ID()
	}
	return p.proxy.exists(topic, id)
}

func (p *Prosumer) IdList(topic string) ([]string, error) {
	return p.proxy.idList(topic)
}
