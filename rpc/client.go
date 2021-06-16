/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/client.go                                            |
|                                                          |
| LastModified: Jun 16, 2021                               |
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

type (
	// Client for RPC.
	Client = core.Client
	// TransportGetter for Client.
	TransportGetter = core.TransportGetter
)

// NewClient returns an instance of Client.
var NewClient = core.NewClient

// HTTPTransport returns http.Transport of Client.
func HTTPTransport(t TransportGetter) *http.Transport {
	return t.GetTransport("http").(*http.Transport)
}

// SocketTransport returns socket.Transport of Client.
func SocketTransport(t TransportGetter) *socket.Transport {
	return t.GetTransport("socket").(*socket.Transport)
}

// UDPTransport returns udp.Transport of Client.
func UDPTransport(t TransportGetter) *udp.Transport {
	return t.GetTransport("udp").(*udp.Transport)
}

// WebSocketTransport returns websocket.Transport of Client.
func WebSocketTransport(t TransportGetter) *websocket.Transport {
	return t.GetTransport("websocket").(*websocket.Transport)
}
