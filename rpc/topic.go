/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * rpc/topic.go                                           *
 *                                                        *
 * hprose push topic for Go.                              *
 *                                                        *
 * LastModified: Sep 21, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import (
	"sync"
	"time"
)

type topic struct {
	sync.RWMutex
	messages  map[string]chan interface{}
	heartbeat time.Duration
}

func newTopic(heartbeat time.Duration) *topic {
	t := new(topic)
	t.messages = make(map[string]chan interface{})
	t.heartbeat = heartbeat
	return t
}

func (t *topic) get(id string) (message chan interface{}) {
	t.RLock()
	message = t.messages[id]
	t.RUnlock()
	return
}

func (t *topic) put(id string, message chan interface{}) {
	t.Lock()
	t.messages[id] = message
	t.Unlock()
}

func (t *topic) remove(id string) {
	t.Lock()
	delete(t.messages, id)
	t.Unlock()
}

func (t *topic) idlist() (result []string) {
	t.RLock()
	result = make([]string, len(t.messages))
	i := 0
	for id := range t.messages {
		result[i] = id
		i++
	}
	t.RUnlock()
	return
}

func (t *topic) exist(id string) (exist bool) {
	t.RLock()
	_, exist = t.messages[id]
	t.RUnlock()
	return
}
