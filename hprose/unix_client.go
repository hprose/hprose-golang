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
 * LastModified: May 28, 2015                             *
 * Authors: Ma Bingyao <andot@hprose.com>                 *
 *          Ore_Ash <nanohugh@gmail.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"crypto/tls"
	"net"
	"strings"
	"time"
)

// UnixClient is hprose unix client
type UnixClient struct {
	*StreamClient
	tlsConfig *tls.Config
}

type unixTransporter struct {
	ConnPool
	*UnixClient
}

var globalUnixConnPool = NewStreamConnPool(64)

// NewUnixClient is the constructor of UnixClient
func NewUnixClient(uri string) (client *UnixClient) {
	trans := new(unixTransporter)
	trans.ConnPool = globalUnixConnPool
	client = new(UnixClient)
	client.StreamClient = newStreamClient(trans)
	client.Client = client
	trans.UnixClient = client
	client.SetUri(uri)
	return client
}

func newUnixClient(uri string) Client {
	return NewUnixClient(uri)
}

// SetConnPool can set separate StreamConnPool for the client
func (client *UnixClient) SetConnPool(connPool ConnPool) {
	client.Transporter.(*unixTransporter).ConnPool = connPool
}

func parseUnixUri(uri string) (scheme, path string) {
	t := strings.SplitN(uri, ":", 2)
	return t[0], t[1]
}

// SetUri set the uri of hprose client
func (client *UnixClient) SetUri(uri string) {
	scheme, _ := parseUnixUri(uri)
	if scheme != "unix" {
		panic("This client desn't support " + scheme + " scheme.")
	}
	client.Close()
	client.BaseClient.SetUri(uri)
}

// Close the client
func (client *UnixClient) Close() {
	uri := client.Uri()
	if uri != "" {
		client.Transporter.(*unixTransporter).ConnPool.Close(uri)
	}
}

// SetKeepAlive do nothing on unix client
func (client *UnixClient) SetKeepAlive(keepalive bool) {
}

// Timeout return the timeout of the connection in client pool
func (client *UnixClient) Timeout() time.Duration {
	return client.Transporter.(*unixTransporter).ConnPool.Timeout()
}

// SetTimeout for connection in client pool
func (client *UnixClient) SetTimeout(d time.Duration) {
	client.timeout = d
	client.Transporter.(*unixTransporter).ConnPool.SetTimeout(d)
}

// TLSClientConfig returns the Config structure used to configure a TLS client
func (client *UnixClient) TLSClientConfig() *tls.Config {
	return client.tlsConfig
}

// SetTLSClientConfig sets the Config structure used to configure a TLS client
func (client *UnixClient) SetTLSClientConfig(config *tls.Config) {
	client.tlsConfig = config
}

// SendAndReceive send and receive the data
func (t *unixTransporter) SendAndReceive(uri string, odata []byte) (idata []byte, err error) {
	connEntry := t.ConnPool.Get(uri)
	defer func() {
		if err != nil {
			connEntry.Close()
			t.ConnPool.Free(connEntry)
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
		if t.tlsConfig != nil {
			conn = tls.Client(conn, t.tlsConfig)
		}
		connEntry.Set(conn)
	}
	if t.timeout != nil {
		if err = conn.SetDeadline(time.Now().Add(t.timeout.(time.Duration))); err != nil {
			err = nil
			connEntry.Close()
			t.ConnPool.Free(connEntry)
			connEntry = t.ConnPool.Get(uri)
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
	t.ConnPool.Free(connEntry)
	return idata, nil
}
