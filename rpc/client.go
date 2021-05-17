/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/client.go                                            |
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
	mock.RegisterTransport()
	http.RegisterTransport()
	socket.RegisterTransport()
	udp.RegisterTransport()
	websocket.RegisterTransport()
}

// Client for RPC.
type (
	Client          = core.Client
	TransportGetter = core.TransportGetter
)

var NewClient = core.NewClient

func HTTPTransport(t TransportGetter) *http.Transport {
	return t.GetTransport("http").(*http.Transport)
}

func SocketTransport(t TransportGetter) *socket.Transport {
	return t.GetTransport("socket").(*socket.Transport)
}

func UDPTransport(t TransportGetter) *udp.Transport {
	return t.GetTransport("udp").(*udp.Transport)
}

func WebSocketTransport(t TransportGetter) *websocket.Transport {
	return t.GetTransport("websocket").(*websocket.Transport)
}
