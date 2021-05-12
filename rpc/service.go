/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/service.go                                           |
|                                                          |
| LastModified: May 12, 2021                               |
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
type Service = core.Service

var NewService = core.NewService

func HTTPHandler(s *Service) *http.Handler {
	return &s.GetHandler("websocket").(*websocket.Handler).Handler
}

func SocketHandler(s *Service) *socket.Handler {
	return s.GetHandler("socket").(*socket.Handler)
}

func UDPHandler(s *Service) *udp.Handler {
	return s.GetHandler("udp").(*udp.Handler)
}

func WebSocketHandler(s *Service) *websocket.Handler {
	return s.GetHandler("websocket").(*websocket.Handler)
}
