/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/mock/agent.go                                        |
|                                                          |
| LastModified: Mar 26, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package mock

import (
	"sync"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

type agent struct {
	handlers map[string]func(address string, request []byte) (response []byte, err error)
	rwlock   sync.RWMutex
}

func (a *agent) Register(address string, handler func(address string, request []byte) (response []byte, err error)) {
	a.rwlock.Lock()
	a.handlers[address] = handler
	a.rwlock.Unlock()
}

func (a *agent) Cancel(address string) {
	a.rwlock.Lock()
	delete(a.handlers, address)
	a.rwlock.Unlock()
}

func (a *agent) Handler(address string, request []byte) ([]byte, error) {
	a.rwlock.RLock()
	defer a.rwlock.RUnlock()
	if handler, ok := a.handlers[address]; ok {
		return handler(address, request)
	}
	return nil, core.ErrServerIsStoped
}

// Agent for mock.
var Agent = &agent{handlers: make(map[string]func(address string, request []byte) (response []byte, err error))}
