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
 * LastModified: Apr 10, 2015                             *
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

type TcpClient struct {
	*StreamClient
	keepAlive       interface{}
	keepAlivePeriod interface{}
	linger          interface{}
	noDelay         interface{}
	tlsConfig       *tls.Config
}

type TcpTransporter struct {
	connPool *StreamConnPool
	*TcpClient
}

func NewTcpClient(uri string) Client {
	trans := &TcpTransporter{connPool: &StreamConnPool{pool: make([]*StreamConnEntry, 0)}}
	client := &TcpClient{StreamClient: newStreamClient(trans)}
	client.Client = client
	trans.TcpClient = client
	client.SetUri(uri)
	return client
}

func (client *TcpClient) SetUri(uri string) {
	if u, err := url.Parse(uri); err == nil {
		if u.Scheme != "tcp" && u.Scheme != "tcp4" && u.Scheme != "tcp6" {
			panic("This client desn't support " + u.Scheme + " scheme.")
		}
	}
	client.Close()
	client.BaseClient.SetUri(uri)
}

func (client *TcpClient) Close() {
	uri := client.Uri()
	if uri == "" {
		client.Transporter.(*TcpTransporter).connPool.Close(uri)
	}
}

func (client *TcpClient) Timeout() time.Duration {
	return client.Transporter.(*TcpTransporter).connPool.Timeout()
}

func (client *TcpClient) SetTimeout(d time.Duration) {
	client.timeout = d
	client.Transporter.(*TcpTransporter).connPool.SetTimeout(d)
}

func (client *TcpClient) SetKeepAlive(keepalive bool) {
	client.keepAlive = keepalive
}

func (client *TcpClient) SetKeepAlivePeriod(d time.Duration) {
	client.keepAlivePeriod = d
}

func (client *TcpClient) SetLinger(sec int) {
	client.linger = sec
}

func (client *TcpClient) SetNoDelay(noDelay bool) {
	client.noDelay = noDelay
}

func (client *TcpClient) TLSClientConfig() *tls.Config {
	return client.tlsConfig
}

func (client *TcpClient) SetTLSClientConfig(config *tls.Config) {
	client.tlsConfig = config
}

func (t *TcpTransporter) SendAndReceive(uri string, odata []byte) (idata []byte, err error) {
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
