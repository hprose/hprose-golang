/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/mock/agent.go                                        |
|                                                          |
| LastModified: Feb 21, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package mock

import "github.com/hprose/hprose-golang/v3/rpc/core"

type agent struct {
	handlers map[string]func(address string, request []byte) (response []byte, err error)
}

func (a *agent) Register(address string, handler func(address string, request []byte) (response []byte, err error)) {
	a.handlers[address] = handler
}

func (a *agent) Cancel(address string) {
	delete(a.handlers, address)
}

func (a *agent) Handler(address string, request []byte) ([]byte, error) {
	if handler, ok := a.handlers[address]; ok {
		return handler(address, request)
	}
	return nil, core.ErrServerIsStoped
}

// Agent for mock.
var Agent = &agent{make(map[string]func(address string, request []byte) (response []byte, err error))}
