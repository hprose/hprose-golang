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

// Service for RPC.
type (
	Service       = core.Service
	HandlerGetter = core.HandlerGetter
)

var NewService = core.NewService

func HTTPHandler(h HandlerGetter) *http.Handler {
	return &h.GetHandler("websocket").(*websocket.Handler).Handler
}

func SocketHandler(h HandlerGetter) *socket.Handler {
	return h.GetHandler("socket").(*socket.Handler)
}

func UDPHandler(h HandlerGetter) *udp.Handler {
	return h.GetHandler("udp").(*udp.Handler)
}

func WebSocketHandler(h HandlerGetter) *websocket.Handler {
	return h.GetHandler("websocket").(*websocket.Handler)
}
