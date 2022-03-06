/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/client.go                                            |
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
	mock.RegisterTransport()
	http.RegisterTransport()
	socket.RegisterTransport()
	udp.RegisterTransport()
	websocket.RegisterTransport()
}

type (
	// Client for RPC.
	Client = core.Client
)

// NewClient returns an instance of Client.
var NewClient = core.NewClient

// HTTPTransport returns http.Transport of Client.
func HTTPTransport(client *Client) *http.Transport {
	return client.GetTransport("http").(*http.Transport)
}

// SocketTransport returns socket.Transport of Client.
func SocketTransport(client *Client) *socket.Transport {
	return client.GetTransport("socket").(*socket.Transport)
}

// UDPTransport returns udp.Transport of Client.
func UDPTransport(client *Client) *udp.Transport {
	return client.GetTransport("udp").(*udp.Transport)
}

// WebSocketTransport returns websocket.Transport of Client.
func WebSocketTransport(client *Client) *websocket.Transport {
	return client.GetTransport("websocket").(*websocket.Transport)
}
