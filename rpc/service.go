/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/service.go                                           |
|                                                          |
| LastModified: May 7, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package rpc

import (
	"github.com/hprose/hprose-golang/v3/rpc/core"
	"github.com/hprose/hprose-golang/v3/rpc/http"
	"github.com/hprose/hprose-golang/v3/rpc/mock"
	"github.com/hprose/hprose-golang/v3/rpc/socket"
	"github.com/hprose/hprose-golang/v3/rpc/udp"
	"github.com/hprose/hprose-golang/v3/rpc/websocket"
)

func init() {
	mock.RegisterHandler()
	socket.RegisterHandler()
	udp.RegisterHandler()
	websocket.RegisterHandler()
}

// Service for RPC.
type Service struct {
	*core.Service
}

func (s Service) Mock() *mock.Handler {
	return s.GetHandler("mock").(*mock.Handler)
}

func (s Service) HTTP() *http.Handler {
	return &s.GetHandler("websocket").(*websocket.Handler).Handler
}

func (s Service) Socket() *socket.Handler {
	return s.GetHandler("socket").(*socket.Handler)
}

func (s Service) UDP() *udp.Handler {
	return s.GetHandler("udp").(*udp.Handler)
}

func (s Service) WebSocket() *websocket.Handler {
	return s.GetHandler("websocket").(*websocket.Handler)
}

// NewService returns an instance of Service.
func NewService() Service {
	return Service{Service: core.NewService()}
}
