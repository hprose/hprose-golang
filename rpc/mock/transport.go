/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/mock/transport.go                                    |
|                                                          |
| LastModified: Feb 21, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package mock

import (
	"context"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

type transport struct{}

func (transport) Transport(ctx context.Context, request []byte) (response []byte, err error) {
	clientContext := core.GetClientContext(ctx)
	url := clientContext.URL
	timeout := clientContext.Timeout
	if timeout <= 0 {
		return Agent.Handler(url.Host, request)
	}
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, timeout)
	defer cancel()
	ch := make(chan struct{}, 1)
	go func() {
		response, err = Agent.Handler(url.Host, request)
		ch <- struct{}{}
		close(ch)
	}()
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-ch:
		return response, err
	}
}

func (transport) Abort() {}

type transportFactory struct {
	schemes []string
}

func (factory transportFactory) Schemes() []string {
	return factory.schemes
}

func (factory transportFactory) New() core.Transport {
	return transport{}
}

func init() {
	core.RegisterTransport("mock", transportFactory{[]string{"mock"}})
}
