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
 * hprose/tcp_client.go                                   *
 *                                                        *
 * hprose tcp client for Go.                              *
 *                                                        *
 * LastModified: May 27, 2015                             *
 * Authors: Ma Bingyao <andot@hprose.com>                 *
 *          Ore_Ash <nanohugh@gmail.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"crypto/tls"
	"net"
	"net/url"
	"time"
)

// TcpClient is hprose tcp client
type TcpClient struct {
	*StreamClient
	keepAlive       interface{}
	keepAlivePeriod interface{}
	linger          interface{}
	noDelay         interface{}
	tlsConfig       *tls.Config
}

type tcpTransporter struct {
	ConnPool
	*TcpClient
}

var globalTcpConnPool = NewStreamConnPool(64)

// NewTcpClient is the constructor of TcpClient
func NewTcpClient(uri string) (client *TcpClient) {
	trans := new(tcpTransporter)
	trans.ConnPool = globalTcpConnPool
	client = new(TcpClient)
	client.StreamClient = newStreamClient(trans)
	client.Client = client
	trans.TcpClient = client
	client.SetUri(uri)
	return
}

func newTcpClient(uri string) Client {
	return NewTcpClient(uri)
}

// SetConnPool can set separate StreamConnPool for the client
func (client *TcpClient) SetConnPool(connPool ConnPool) {
	client.Transporter.(*tcpTransporter).ConnPool = connPool
}

// SetUri set the uri of hprose client
func (client *TcpClient) SetUri(uri string) {
	if u, err := url.Parse(uri); err == nil {
		if u.Scheme != "tcp" && u.Scheme != "tcp4" && u.Scheme != "tcp6" {
			panic("This client desn't support " + u.Scheme + " scheme.")
		}
	}
	client.Close()
	client.BaseClient.SetUri(uri)
}

// Close the client
func (client *TcpClient) Close() {
	uri := client.Uri()
	if uri != "" {
		client.Transporter.(*tcpTransporter).ConnPool.Close(uri)
	}
}

// Timeout return the timeout of the connection in client pool
func (client *TcpClient) Timeout() time.Duration {
	return client.Transporter.(*tcpTransporter).ConnPool.Timeout()
}

// SetTimeout for connection in client pool
func (client *TcpClient) SetTimeout(d time.Duration) {
	client.timeout = d
	client.Transporter.(*tcpTransporter).ConnPool.SetTimeout(d)
}

// SetKeepAlive sets whether the operating system should send keepalive messages on the connection.
func (client *TcpClient) SetKeepAlive(keepalive bool) {
	client.keepAlive = keepalive
}

// SetKeepAlivePeriod sets period between keep alives.
func (client *TcpClient) SetKeepAlivePeriod(d time.Duration) {
	client.keepAlivePeriod = d
}

// SetLinger sets the behavior of Close on a connection which still has data waiting to be sent or to be acknowledged.
//
// If sec < 0 (the default), the operating system finishes sending the data in the background.
//
// If sec == 0, the operating system discards any unsent or unacknowledged data.
//
// If sec > 0, the data is sent in the background as with sec < 0. On some operating systems after sec seconds have elapsed any remaining unsent data may be discarded.
func (client *TcpClient) SetLinger(sec int) {
	client.linger = sec
}

// SetNoDelay controls whether the operating system should delay packet transmission in hopes of sending fewer packets (Nagle's algorithm). The default is true (no delay), meaning that data is sent as soon as possible after a Write.
func (client *TcpClient) SetNoDelay(noDelay bool) {
	client.noDelay = noDelay
}

// TLSClientConfig returns the Config structure used to configure a TLS client
func (client *TcpClient) TLSClientConfig() *tls.Config {
	return client.tlsConfig
}

// SetTLSClientConfig sets the Config structure used to configure a TLS client
func (client *TcpClient) SetTLSClientConfig(config *tls.Config) {
	client.tlsConfig = config
}

// SendAndReceive send and receive the data
func (t *tcpTransporter) SendAndReceive(uri string, odata []byte) (idata []byte, err error) {
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
		var u *url.URL
		if u, err = url.Parse(uri); err != nil {
			return nil, err
		}
		var tcpaddr *net.TCPAddr
		if tcpaddr, err = net.ResolveTCPAddr(u.Scheme, u.Host); err != nil {
			return nil, err
		}
		if conn, err = net.DialTCP("tcp", nil, tcpaddr); err != nil {
			return nil, err
		}
		if t.keepAlive != nil {
			if err = conn.(*net.TCPConn).SetKeepAlive(t.keepAlive.(bool)); err != nil {
				return nil, err
			}
		}
		if t.keepAlivePeriod != nil {
			if kap, ok := conn.(iKeepAlivePeriod); ok {
				if err = kap.SetKeepAlivePeriod(t.keepAlivePeriod.(time.Duration)); err != nil {
					return nil, err
				}
			}
		}
		if t.linger != nil {
			if err = conn.(*net.TCPConn).SetLinger(t.linger.(int)); err != nil {
				return nil, err
			}
		}
		if t.noDelay != nil {
			if err = conn.(*net.TCPConn).SetNoDelay(t.noDelay.(bool)); err != nil {
				return nil, err
			}
		}
		if t.readBuffer != nil {
			if err = conn.(*net.TCPConn).SetReadBuffer(t.readBuffer.(int)); err != nil {
				return nil, err
			}
		}
		if t.writeBuffer != nil {
			if err = conn.(*net.TCPConn).SetWriteBuffer(t.writeBuffer.(int)); err != nil {
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
