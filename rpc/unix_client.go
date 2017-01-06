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
 * rpc/unix_client.go                                     *
 *                                                        *
 * hprose unx client for Go.                              *
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
)

// UnixClient is hprose unix client
type UnixClient struct {
	SocketClient
}

// NewUnixClient is the constructor of UnixClient
func NewUnixClient(uri ...string) (client *UnixClient) {
	client = new(UnixClient)
	client.initSocketClient()
	client.setCreateConn(client.createUnixConn)
	client.SetURIList(uri)
	return
}

func newUnixClient(uri ...string) Client {
	return NewUnixClient(uri...)
}

// SetURIList set a list of server addresses
func (client *UnixClient) SetURIList(uriList []string) {
	CheckAddresses(uriList, unixSchemes)
	client.BaseClient.SetURIList(uriList)
}

func (client *UnixClient) createUnixConn() (net.Conn, error) {
	u, err := url.Parse(client.uri)
	if err != nil {
		return nil, err
	}
	unixaddr, err := net.ResolveUnixAddr(u.Scheme, u.Path)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialUnix(u.Scheme, nil, unixaddr)
	if err != nil {
		return nil, err
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
