/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/push/broker.go                               |
|                                                          |
| LastModified: May 17, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package push

import (
	"context"
	"sync"
	"time"

	"github.com/hprose/hprose-golang/v3/rpc/core"
	cmap "github.com/orcaman/concurrent-map"
)

type Broker struct {
	*core.Service
	messages              sync.Map           // map[string]map[string]messageCache
	responders            cmap.ConcurrentMap // map[string]chan map[string][]Message
	signals               cmap.ConcurrentMap // map[string]chan bool
	MessageQueueMaxLength int
	Timeout               time.Duration
	HeartBeat             time.Duration
	OnSubscribe           func(ctx context.Context, id string, topic string)
	OnUnsubscribe         func(ctx context.Context, id string, topic string, messages []Message)
}

func NewBroker(service *core.Service) *Broker {
	broker := &Broker{
		Service:               service,
		responders:            cmap.New(),
		signals:               cmap.New(),
		MessageQueueMaxLength: 10,
		Timeout:               time.Minute * 2,
		HeartBeat:             time.Second * 10,
	}
	service.Use(broker.handler).
		AddFunction(broker.subscribe, "+").
		AddFunction(broker.unsubscribe, "-").
		AddFunction(broker.message, "<").
		AddFunction(broker.Unicast, ">").
		AddFunction(broker.Multicast, ">?").
		AddFunction(broker.Broadcast, ">*").
		AddFunction(broker.Exists, "?").
		AddFunction(broker.IdList, "|")
	return broker
}

func (b *Broker) send(ctx context.Context, id string, responder chan map[string][]Message) bool {
	var topics *sync.Map
	if value, ok := b.messages.Load(id); ok {
		topics = value.(*sync.Map)
	}
	if topics == nil {
		responder <- nil
		return true
	}
	var size int
	result := make(map[string][]Message)
	topics.Range(func(key, value interface{}) bool {
		size++
		topic := key.(string)
		cache := value.(*MessageCache)
		if cache == nil {
			result[topic] = nil
			topics.Delete(topic)
		} else {
			messages := cache.Take()
			if len(messages) > 0 {
				result[topic] = messages
			}
		}
		return true
	})
	if size == 0 {
		responder <- nil
		return true
	}
	if len(result) == 0 {
		return false
	}
	responder <- result
	go b.doHeartBeat(ctx, id)
	return true
}

func (b *Broker) doHeartBeat(ctx context.Context, id string) {
	if b.HeartBeat <= 0 {
		return
	}
	signal := make(chan bool, 1)
	b.signals.Upsert(id, signal, func(exist bool, valueInMap interface{}, newValue interface{}) interface{} {
		if exist {
			close(valueInMap.(chan bool))
		}
		return newValue
	})
	ctx, cancel := context.WithTimeout(ctx, b.HeartBeat)
	defer cancel()
	select {
	case <-ctx.Done():
		if topics, ok := b.messages.Load(id); ok {
			topics := topics.(*sync.Map)
			topics.Range(func(key, value interface{}) bool {
				b.offline(ctx, topics, id, key.(string))
				return true
			})
		}
	case <-signal:
	}
}

func (b *Broker) ID(ctx context.Context) (id string) {
	if id = core.GetServiceContext(ctx).RequestHeaders().GetString("id"); id == "" {
		panic("client unique id not found")
	}
	return
}

func (b *Broker) subscribe(ctx context.Context, topic string) bool {
	id := b.ID(ctx)
	t, ok := b.messages.Load(id)
	if !ok {
		t, _ = b.messages.LoadOrStore(id, new(sync.Map))
	}
	topics := t.(*sync.Map)
	if _, ok := topics.Load(topic); ok {
		return false
	}
	_, loaded := topics.LoadOrStore(topic, new(MessageCache))
	if !loaded && b.OnSubscribe != nil {
		b.OnSubscribe(ctx, id, topic)

	}
	return !loaded
}

func (b *Broker) response(ctx context.Context, id string) {
	if responder, ok := b.responders.Pop(id); ok {
		responder := responder.(chan map[string][]Message)
		if !b.send(ctx, id, responder) {
			if !b.responders.SetIfAbsent(id, responder) {
				responder <- nil
			}
		}
	}
}

func (b *Broker) offline(ctx context.Context, topics *sync.Map, id string, topic string) bool {
	if messages, ok := topics.Load(topic); ok {
		topics.Delete(topic)
		if b.OnUnsubscribe != nil {
			b.OnUnsubscribe(ctx, id, topic, messages.(*MessageCache).Take())
		}
		b.response(ctx, id)
		return true
	}
	return false
}

func (b *Broker) unsubscribe(ctx context.Context, topic string) bool {
	id := b.ID(ctx)
	if topics, ok := b.messages.Load(id); ok {
		return b.offline(ctx, topics.(*sync.Map), id, topic)
	}
	return false
}

func (b *Broker) message(ctx context.Context) map[string][]Message {
	id := b.ID(ctx)
	if responder, ok := b.responders.Pop(id); ok {
		responder.(chan map[string][]Message) <- nil
	}
	if signal, ok := b.signals.Pop(id); ok {
		close(signal.(chan bool))
	}
	responder := make(chan map[string][]Message, 1)
	if !b.send(ctx, id, responder) {
		b.responders.Upsert(id, responder, func(exist bool, valueInMap interface{}, newValue interface{}) interface{} {
			if exist {
				valueInMap.(chan map[string][]Message) <- nil
			}
			return newValue
		})
		if b.Timeout > 0 {
			ctx, cancel := context.WithTimeout(ctx, b.Timeout)
			defer cancel()
			select {
			case <-ctx.Done():
				go b.doHeartBeat(context.Background(), id)
				return map[string][]Message{}
			case result := <-responder:
				return result
			}
		}
	}
	return <-responder
}

func (b *Broker) Unicast(ctx context.Context, data interface{}, topic string, id string, from string) bool {
	if topics, ok := b.messages.Load(id); ok {
		if cache, ok := topics.(*sync.Map).Load(topic); ok && cache != nil {
			cache.(*MessageCache).Append(Message{Data: data, From: from})
			b.response(ctx, id)
			return true
		}
	}
	return false
}

func (b *Broker) Multicast(ctx context.Context, data interface{}, topic string, ids []string, from string) map[string]bool {
	result := make(map[string]bool)
	for _, id := range ids {
		result[id] = b.Unicast(ctx, data, topic, id, from)
	}
	return result
}

func (b *Broker) Broadcast(ctx context.Context, data interface{}, topic string, from string) map[string]bool {
	result := make(map[string]bool)
	b.messages.Range(func(key, value interface{}) bool {
		id := key.(string)
		topics := value.(*sync.Map)
		if cache, ok := topics.Load(topic); ok && cache != nil {
			cache.(*MessageCache).Append(Message{Data: data, From: from})
			b.response(ctx, id)
			result[id] = true
		}
		result[id] = false
		return true
	})
	return result
}

func (b *Broker) Push(data interface{}, topic string, id ...string) map[string]bool {
	ctx := context.Background()
	switch len(id) {
	case 0:
		return b.Broadcast(ctx, data, topic, "")
	case 1:
		return map[string]bool{
			id[0]: b.Unicast(ctx, data, topic, id[0], ""),
		}
	default:
		return b.Multicast(ctx, data, topic, id, "")
	}
}

func (b *Broker) Deny(ctx context.Context, id string, topic string) {
	if topics, ok := b.messages.Load(id); ok {
		topics := topics.(*sync.Map)
		if topic != "" {
			if cache, ok := topics.Load(topic); ok && cache != nil {
				topics.Store(topic, nil)
			}
		} else {
			topics.Range(func(key, _ interface{}) bool {
				topics.Store(key, nil)
				return false
			})
		}
		b.response(ctx, id)
	}
}

func (b *Broker) Exists(topic string, id string) bool {
	if topics, ok := b.messages.Load(id); ok {
		if cache, ok := topics.(*sync.Map).Load(topic); ok {
			return cache != nil
		}
	}
	return false
}

func (b *Broker) IdList(topic string) (idlist []string) {
	b.messages.Range(func(key, value interface{}) bool {
		id := key.(string)
		topics := value.(*sync.Map)
		if cache, ok := topics.Load(topic); ok && cache != nil {
			idlist = append(idlist, id)
		}
		return true
	})
	return
}

func (b *Broker) handler(ctx context.Context, name string, args []interface{}, next core.NextInvokeHandler) (result []interface{}, err error) {
	serviceContext := core.GetServiceContext(ctx)
	var from string
	if id := serviceContext.RequestHeaders().GetString("id"); id != "" {
		from = id
	}
	switch name {
	case ">", ">?", ">*":
		args = append(args, from)
	}
	serviceContext.Items().Set("producer", producer{b, from})
	return next(ctx, name, args)
}
