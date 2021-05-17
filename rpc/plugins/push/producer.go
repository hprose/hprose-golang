/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/push/producer.go                             |
|                                                          |
| LastModified: May 17, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package push

import "context"

type Producer interface {
	From() string
	Unicast(ctx context.Context, data interface{}, topic string, id string) bool
	Multicast(ctx context.Context, data interface{}, topic string, ids []string) map[string]bool
	Broadcast(ctx context.Context, data interface{}, topic string) map[string]bool
	Push(data interface{}, topic string, id ...string) map[string]bool
	Deny(ctx context.Context, id string, topic string)
	Exists(topic string, id string) bool
	IdList(topic string) []string
}

type producer struct {
	broker *Broker
	from   string
}

func (p producer) From() string {
	return p.from
}

func (p producer) Unicast(ctx context.Context, data interface{}, topic string, id string) bool {
	return p.broker.Unicast(ctx, data, topic, id, p.from)
}

func (p producer) Multicast(ctx context.Context, data interface{}, topic string, ids []string) map[string]bool {
	return p.broker.Multicast(ctx, data, topic, ids, p.from)
}

func (p producer) Broadcast(ctx context.Context, data interface{}, topic string) map[string]bool {
	return p.broker.Broadcast(ctx, data, topic, p.from)
}

func (p producer) Push(data interface{}, topic string, id ...string) map[string]bool {
	ctx := context.Background()
	switch len(id) {
	case 0:
		return p.broker.Broadcast(ctx, data, topic, p.from)
	case 1:
		return map[string]bool{
			id[0]: p.broker.Unicast(ctx, data, topic, id[0], p.from),
		}
	default:
		return p.broker.Multicast(ctx, data, topic, id, p.from)
	}
}

func (p producer) Deny(ctx context.Context, id string, topic string) {
	if id == "" {
		id = p.from
	}
	p.broker.Deny(ctx, id, topic)
}

func (p producer) Exists(topic string, id string) bool {
	if id == "" {
		id = p.from
	}
	return p.broker.Exists(topic, id)
}

func (p producer) IdList(topic string) []string {
	return p.broker.IdList(topic)
}
