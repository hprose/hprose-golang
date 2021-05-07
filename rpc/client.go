/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/client.go                                            |
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
	mock.RegisterTransport()
	http.RegisterTransport()
	socket.RegisterTransport()
	udp.RegisterTransport()
	websocket.RegisterTransport()
}

// Client for RPC.
type Client struct {
	*core.Client
}

func (c Client) Mock() *mock.Transport {
	return c.GetTransport("mock").(*mock.Transport)
}

func (c Client) HTTP() *http.Transport {
	return c.GetTransport("http").(*http.Transport)
}

func (c Client) Socket() *socket.Transport {
	return c.GetTransport("socket").(*socket.Transport)
}

func (c Client) UDP() *udp.Transport {
	return c.GetTransport("udp").(*udp.Transport)
}

func (c Client) WebSocket() *websocket.Transport {
	return c.GetTransport("websocket").(*websocket.Transport)
}

// NewClient returns an instance of Client.
func NewClient(uri ...string) Client {
	return Client{Client: core.NewClient(uri...)}
}
