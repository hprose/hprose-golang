/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/push/message.go                              |
|                                                          |
| LastModified: May 16, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package push

import (
	"sync"

	"github.com/hprose/hprose-golang/v3/encoding"
)

type Message struct {
	Data interface{} `json:"data"`
	From string      `json:"from"`
}

type MessageCache struct {
	m []Message
	l sync.Mutex
}

func (m *MessageCache) Append(message Message) {
	m.l.Lock()
	defer m.l.Unlock()
	m.m = append(m.m, message)
}

func (m *MessageCache) Take() (result []Message) {
	m.l.Lock()
	defer m.l.Unlock()
	result = m.m
	m.m = nil
	return
}

func init() {
	encoding.RegisterAlias((*Message)(nil), "@")
}
