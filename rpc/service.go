/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/service.go                                           |
|                                                          |
| LastModified: Mar 6, 2022                                |
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

type (
	// Service for RPC.
	Service = core.Service
)

// NewService returns an instance of Service.
var NewService = core.NewService

// HTTPHandler returns http.Handler of Service.
func HTTPHandler(service *Service) *http.Handler {
	return &service.GetHandler("websocket").(*websocket.Handler).Handler
}

// SocketHandler returns socket.Handler of Service.
func SocketHandler(service *Service) *socket.Handler {
	return service.GetHandler("socket").(*socket.Handler)
}

// UDPHandler returns udp.Handler of Service.
func UDPHandler(service *Service) *udp.Handler {
	return service.GetHandler("udp").(*udp.Handler)
}

// WebSocketHandler returns websocket.Handler of Service.
func WebSocketHandler(service *Service) *websocket.Handler {
	return service.GetHandler("websocket").(*websocket.Handler)
}
