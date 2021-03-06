/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/forward/forward.go                           |
|                                                          |
| LastModified: Mar 7, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package forward

import (
	"context"
	"time"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

// Forward plugin for hprose.
type Forward struct {
	Timeout time.Duration
	client  *core.Client
}

// New returns a Forward instance.
func New(uri ...string) *Forward {
	return &Forward{client: core.NewClient(uri...)}
}

// IOHandler for Forward.
func (f *Forward) IOHandler(ctx context.Context, request []byte, next core.NextIOHandler) (response []byte, err error) {
	clientContext := core.NewClientContext()
	clientContext.Timeout = f.Timeout
	clientContext.Init(f.client)
	return f.client.Request(core.WithContext(ctx, clientContext), request)
}

// InvokeHandler for Forward.
func (f *Forward) InvokeHandler(ctx context.Context, name string, args []interface{}, next core.NextInvokeHandler) (result []interface{}, err error) {
	serviceContext := core.GetServiceContext(ctx)
	clientContext := core.NewClientContext()
	clientContext.Timeout = f.Timeout
	if serviceContext.HasRequestHeaders() {
		serviceContext.RequestHeaders().CopyTo(clientContext.RequestHeaders())
	}
	result, err = f.client.InvokeContext(core.WithContext(ctx, clientContext), name, args)
	if clientContext.HasResponseHeaders() {
		clientContext.ResponseHeaders().CopyTo(serviceContext.ResponseHeaders())
	}
	return
}

// Use plugin handlers.
func (f *Forward) Use(handler ...core.PluginHandler) *Forward {
	f.client.Use(handler...)
	return f
}

// Unuse plugin handlers.
func (f *Forward) Unuse(handler ...core.PluginHandler) *Forward {
	f.client.Unuse(handler...)
	return f
}
