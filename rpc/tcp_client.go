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
 * rpc/tcp_client.go                                      *
 *                                                        *
 * hprose tcp client for Go.                              *
 *                                                        *
 * LastModified: Nov 1, 2016                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import (
	"crypto/tls"
	"net"
	"net/url"
	"time"
)

// TCPClient is hprose tcp client
type TCPClient struct {
	SocketClient
	Linger          int
	NoDelay         bool
	KeepAlive       bool
	KeepAlivePeriod time.Duration
}

// NewTCPClient is the constructor of TCPClient
func NewTCPClient(uri ...string) (client *TCPClient) {
	client = new(TCPClient)
	client.initSocketClient()
	client.Linger = -1
	client.NoDelay = true
	client.KeepAlive = true
	client.setCreateConn(client.createTCPConn)
	client.SetURIList(uri)
	return
}

func newTCPClient(uri ...string) Client {
	return NewTCPClient(uri...)
}

// SetURIList set a list of server addresses
func (client *TCPClient) SetURIList(uriList []string) {
	CheckAddresses(uriList, tcpSchemes)
	client.BaseClient.SetURIList(uriList)
}

func (client *TCPClient) createTCPConn() net.Conn {
	u, err := url.Parse(client.uri)
	ifErrorPanic(err)
	tcpaddr, err := net.ResolveTCPAddr(u.Scheme, u.Host)
	ifErrorPanic(err)
	conn, err := net.DialTCP(u.Scheme, nil, tcpaddr)
	ifErrorPanic(err)
	ifErrorPanic(conn.SetLinger(client.Linger))
	ifErrorPanic(conn.SetNoDelay(client.NoDelay))
	ifErrorPanic(conn.SetKeepAlive(client.KeepAlive))
	if client.KeepAlivePeriod > 0 {
		ifErrorPanic(conn.SetKeepAlivePeriod(client.KeepAlivePeriod))
	}
	if client.ReadBuffer > 0 {
		ifErrorPanic(conn.SetReadBuffer(client.ReadBuffer))
	}
	if client.WriteBuffer > 0 {
		ifErrorPanic(conn.SetWriteBuffer(client.WriteBuffer))
	}
	if client.TLSConfig != nil {
		return tls.Client(conn, client.TLSConfig)
	}
	return conn
}
