/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * rpc/socket_client.go                                   *
 *                                                        *
 * hprose socket client for Go.                           *
 *                                                        *
 * LastModified: Jan 7, 2017                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import (
	"crypto/tls"
	"net"
	"time"
)

type socketTransport interface {
	IdleTimeout() time.Duration
	SetIdleTimeout(timeout time.Duration)
	MaxPoolSize() int
	SetMaxPoolSize(size int)
	setCreateConn(createConn func() (net.Conn, error))
	sendAndReceive(data []byte, context *ClientContext) ([]byte, error)
	close()
}

// SocketClient is base struct for TCPClient and UnixClient
type SocketClient struct {
	BaseClient
	socketTransport
	ReadBuffer  int
	WriteBuffer int
	TLSConfig   *tls.Config
}

func (client *SocketClient) initSocketClient() {
	client.InitBaseClient()
	client.socketTransport = newHalfDuplexSocketTransport()
	client.ReadBuffer = 0
	client.WriteBuffer = 0
	client.TLSConfig = nil
	client.SendAndReceive = client.sendAndReceive
}

// TLSClientConfig returns the tls.Config in hprose client
func (client *SocketClient) TLSClientConfig() *tls.Config {
	return client.TLSConfig
}

// SetTLSClientConfig sets the tls.Config
func (client *SocketClient) SetTLSClientConfig(config *tls.Config) {
	client.TLSConfig = config
}

// Close the client
func (client *SocketClient) Close() {
	client.socketTransport.close()
}
