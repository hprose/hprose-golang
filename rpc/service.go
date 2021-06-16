/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/service.go                                           |
|                                                          |
| LastModified: May 17, 2021                               |
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
	// HandlerGetter for Service.
	HandlerGetter = core.HandlerGetter
)

// NewService returns an instance of Service.
var NewService = core.NewService

// HTTPHandler returns http.Handler of Service.
func HTTPHandler(h HandlerGetter) *http.Handler {
	return &h.GetHandler("websocket").(*websocket.Handler).Handler
}

// SocketHandler returns socket.Handler of Service.
func SocketHandler(h HandlerGetter) *socket.Handler {
	return h.GetHandler("socket").(*socket.Handler)
}

// UDPHandler returns udp.Handler of Service.
func UDPHandler(h HandlerGetter) *udp.Handler {
	return h.GetHandler("udp").(*udp.Handler)
}

// WebSocketHandler returns websocket.Handler of Service.
func WebSocketHandler(h HandlerGetter) *websocket.Handler {
	return h.GetHandler("websocket").(*websocket.Handler)
}
