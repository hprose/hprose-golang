/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/client.go                                            |
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
	mock.RegisterTransport()
	http.RegisterTransport()
	socket.RegisterTransport()
	udp.RegisterTransport()
	websocket.RegisterTransport()
}

// Client for RPC.
type Client = core.Client

var NewClient = core.NewClient

func HTTPTransport(c *Client) *http.Transport {
	return c.GetTransport("http").(*http.Transport)
}

func SocketTransport(c *Client) *socket.Transport {
	return c.GetTransport("socket").(*socket.Transport)
}

func UDPTransport(c *Client) *udp.Transport {
	return c.GetTransport("udp").(*udp.Transport)
}

func WebSocketTransport(c *Client) *websocket.Transport {
	return c.GetTransport("websocket").(*websocket.Transport)
}
