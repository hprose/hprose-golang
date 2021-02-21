/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/mock/handler.go                                      |
|                                                          |
| LastModified: Feb 21, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package mock

import (
	"context"
	"net/url"
	"reflect"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

// Server for mock.
type Server struct {
	Address string
}

// Close the mock server.
func (server Server) Close() {
	Agent.Cancel(server.Address)
}

// Handler for mock.
type Handler struct {
	Service *core.Service
}

// Bind to the mock server.
func (h Handler) Bind(server core.Server) {
	Agent.Register(server.(Server).Address, h.Handler)
}

// Handler for mock.
func (h Handler) Handler(address string, request []byte) (response []byte, err error) {
	if len(request) > h.Service.MaxRequestLength {
		return nil, core.ErrRequestEntityTooLarge
	}
	serviceContext := core.NewServiceContext(h.Service)
	ctx := core.WithContext(context.Background(), serviceContext)
	url, err := url.Parse("mock://" + address)
	if err != nil {
		return nil, err
	}
	addr := NewAddr(url)
	serviceContext.LocalAddr = addr
	serviceContext.RemoteAddr = addr
	serviceContext.Handler = h
	return h.Service.Handle(ctx, request)
}

type handlerFactory struct {
	serverTypes []reflect.Type
}

func (factory handlerFactory) ServerTypes() []reflect.Type {
	return factory.serverTypes
}

func (factory handlerFactory) New(service *core.Service) core.Handler {
	return Handler{service}
}

func init() {
	core.RegisterHandler("mock", handlerFactory{
		[]reflect.Type{
			reflect.TypeOf((*Server)(nil)),
			reflect.TypeOf((*Server)(nil)).Elem(),
		},
	})
}
