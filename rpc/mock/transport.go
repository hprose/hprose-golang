/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/mock/transport.go                                    |
|                                                          |
| LastModified: May 7, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package mock

import (
	"context"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

type Transport struct{}

func (*Transport) Transport(ctx context.Context, request []byte) (response []byte, err error) {
	clientContext := core.GetClientContext(ctx)
	url := clientContext.URL
	ch := make(chan struct {
		response []byte
		err      error
	}, 1)
	go func() {
		response, err := Agent.Handler(ctx, url.Host, request)
		ch <- struct {
			response []byte
			err      error
		}{
			response: response,
			err:      err,
		}
		close(ch)
	}()
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case result := <-ch:
		return result.response, result.err
	}
}

func (*Transport) Abort() {}

type transportFactory struct {
	schemes []string
}

func (factory transportFactory) Schemes() []string {
	return factory.schemes
}

func (factory transportFactory) New() core.Transport {
	return &Transport{}
}

func RegisterTransport() {
	core.RegisterTransport("mock", transportFactory{[]string{"mock"}})
}
