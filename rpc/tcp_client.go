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
 * LastModified: Jan 7, 2017                              *
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

func (client *TCPClient) createTCPConn() (net.Conn, error) {
	u, err := url.Parse(client.uri)
	if err != nil {
		return nil, err
	}
	tcpaddr, err := net.ResolveTCPAddr(u.Scheme, u.Host)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP(u.Scheme, nil, tcpaddr)
	if err != nil {
		return nil, err
	}
	err = conn.SetLinger(client.Linger)
	if err != nil {
		return nil, err
	}
	err = conn.SetNoDelay(client.NoDelay)
	if err != nil {
		return nil, err
	}
	err = conn.SetKeepAlive(client.KeepAlive)
	if err != nil {
		return nil, err
	}
	if client.KeepAlivePeriod > 0 {
		err = conn.SetKeepAlivePeriod(client.KeepAlivePeriod)
		if err != nil {
			return nil, err
		}
	}
	if client.ReadBuffer > 0 {
		err = conn.SetReadBuffer(client.ReadBuffer)
		if err != nil {
			return nil, err
		}
	}
	if client.WriteBuffer > 0 {
		err = conn.SetWriteBuffer(client.WriteBuffer)
		if err != nil {
			return nil, err
		}
	}
	if client.TLSConfig != nil {
		return tls.Client(conn, client.TLSConfig), nil
	}
	return conn, nil
}
