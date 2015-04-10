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
 * hprose/unix_client.go                                  *
 *                                                        *
 * hprose unix client for Go.                             *
 *                                                        *
 * LastModified: Apr 10, 2015                             *
 * Author: Ore_Ash <nanohugh@gmail.com>                   *
 *                                                        *
\**********************************************************/

package hprose

import (
	"net"
	"strings"
	"time"
)

type UnixClient struct {
	*StreamClient
}

type UnixTransporter struct {
	connPool *StreamConnPool
	*UnixClient
}

func NewUnixClient(uri string) Client {
	trans := &UnixTransporter{connPool: &StreamConnPool{pool: make([]*StreamConnEntry, 0)}}
	client := &UnixClient{StreamClient: newStreamClient(trans)}
	client.Client = client
	trans.UnixClient = client
	client.SetUri(uri)
	return client
}

func (client *UnixClient) SetUri(uri string) {
	scheme, _ := parseUnixUri(uri)
	if scheme != "unix" {
		panic("This client desn't support " + scheme + " scheme.")
	}
	client.Close()
	client.BaseClient.SetUri(uri)
}

func (client *UnixClient) Close() {
	uri := client.Uri()
	if uri == "" {
		client.Transporter.(*UnixTransporter).connPool.Close(uri)
	}
}

func (client *UnixClient) Timeout() time.Duration {
	return client.Transporter.(*UnixTransporter).connPool.Timeout()
}

func (client *UnixClient) SetTimeout(d time.Duration) {
	client.timeout = d
	client.Transporter.(*UnixTransporter).connPool.SetTimeout(d)
}

func parseUnixUri(uri string) (scheme, path string) {
	t := strings.Split(uri, ":")
	return t[0], t[1]
}

func (t *UnixTransporter) SendAndReceive(uri string, odata []byte) (idata []byte, err error) {
	connEntry := t.connPool.Get(uri)
	defer func() {
		if err != nil {
			connEntry.Close()
			t.connPool.Free(connEntry)
		}
	}()
begin:
	conn := connEntry.Get()
	if conn == nil {
		scheme, path := parseUnixUri(uri)
		var unixaddr *net.UnixAddr
		if unixaddr, err = net.ResolveUnixAddr(scheme, path); err != nil {
			return nil, err
		}
		if conn, err = net.DialUnix(scheme, nil, unixaddr); err != nil {
			return nil, err
		}
		if t.readBuffer != nil {
			if err = conn.(*net.UnixConn).SetReadBuffer(t.readBuffer.(int)); err != nil {
				return nil, err
			}
		}
		if t.writeBuffer != nil {
			if err = conn.(*net.UnixConn).SetWriteBuffer(t.writeBuffer.(int)); err != nil {
				return nil, err
			}
		}
		connEntry.Set(conn)
	}
	if t.timeout != nil {
		if err = conn.SetDeadline(time.Now().Add(t.timeout.(time.Duration))); err != nil {
			err = nil;
			connEntry.Close()
			t.connPool.Free(connEntry)
			connEntry = t.connPool.Get(uri)
			goto begin
		}
	}
	if t.writeTimeout != nil {
		if err = conn.SetWriteDeadline(time.Now().Add(t.writeTimeout.(time.Duration))); err != nil {
			return nil, err
		}
	}
	if err = sendDataOverStream(conn, odata); err != nil {
		return nil, err
	}
	if t.readTimeout != nil {
		if err = conn.SetReadDeadline(time.Now().Add(t.readTimeout.(time.Duration))); err != nil {
			return nil, err
		}
	}
	if idata, err = receiveDataOverStream(conn); err != nil {
		return nil, err
	}
	t.connPool.Free(connEntry)
	return idata, nil
}
