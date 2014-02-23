/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.net/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * hprose/tcp_client.go                                   *
 *                                                        *
 * hprose tcp client for Go.                              *
 *                                                        *
 * LastModified: Feb 23, 2014                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"crypto/tls"
	"io"
	"net"
	"net/url"
	"sync"
	"time"
)

type TcpClient struct {
	*BaseClient
	deadline        interface{}
	keepAlive       interface{}
	keepAlivePeriod interface{}
	linger          interface{}
	noDelay         interface{}
	readBuffer      interface{}
	readDeadline    interface{}
	writerBuffer    interface{}
	writerDeadline  interface{}
	config          *tls.Config
}

type TcpConnEntry struct {
	uri   string
	conn  net.Conn
	free  bool
	valid bool
}

func (connEntry *TcpConnEntry) Get() net.Conn {
	return connEntry.conn
}

func (connEntry *TcpConnEntry) Set(conn net.Conn) {
	if conn != nil {
		connEntry.conn = conn
	}
}

func (connEntry *TcpConnEntry) Close() {
	connEntry.free = true
	connEntry.valid = false
}

type TcpConnPool struct {
	sync.Mutex
	pool []*TcpConnEntry
}

func (connPool *TcpConnPool) Get(uri string) *TcpConnEntry {
	connPool.Lock()
	defer connPool.Unlock()
	for _, entry := range connPool.pool {
		if entry.free && entry.valid {
			if entry.uri == uri {
				entry.free = false
				return entry
			} else if entry.uri == "" {
				entry.free = false
				entry.valid = false
				entry.uri = uri
				return entry
			}
		}
	}
	entry := &TcpConnEntry{uri, nil, false, false}
	connPool.pool = append(connPool.pool, entry)
	return entry
}

func (connPool *TcpConnPool) freeConns(conns []net.Conn) {
	for _, conn := range conns {
		conn.Close()
	}
}

func (connPool *TcpConnPool) Close(uri string) {
	connPool.Lock()
	defer connPool.Unlock()
	conns := make([]net.Conn, 0, len(connPool.pool))
	for _, entry := range connPool.pool {
		if entry.uri == uri {
			if entry.free && entry.valid {
				conns = append(conns, entry.conn)
				entry.conn = nil
				entry.uri = ""
			} else {
				entry.Close()
			}
		}
	}
	go connPool.freeConns(conns)
}

func (connPool *TcpConnPool) Free(entry *TcpConnEntry) {
	if entry.free && !entry.valid {
		if entry.conn != nil {
			entry.conn.Close()
			entry.conn = nil
		}
		entry.uri = ""
	}
	connPool.Lock()
	entry.free = true
	entry.valid = true
	connPool.Unlock()
}

type TcpTransporter struct {
	connPool *TcpConnPool
	*TcpClient
}

func NewTcpClient(uri string) Client {
	trans := &TcpTransporter{connPool: &TcpConnPool{pool: make([]*TcpConnEntry, 0)}}
	client := &TcpClient{BaseClient: NewBaseClient(trans)}
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

func (client *TcpClient) SetDeadline(t time.Time) {
	client.deadline = t
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

func (client *TcpClient) SetReadBuffer(bytes int) {
	client.readBuffer = bytes
}

func (client *TcpClient) SetReadDeadline(t time.Time) {
	client.readDeadline = t
}

func (client *TcpClient) SetWriteBuffer(bytes int) {
	client.writerBuffer = bytes
}

func (client *TcpClient) SetWriteDeadline(t time.Time) {
	client.writerDeadline = t
}

func (client *TcpClient) SetTLSConfig(config *tls.Config) {
	client.config = config
}

func (t *TcpTransporter) GetInvokeContext(uri string) (context interface{}, err error) {
	connEntry := t.connPool.Get(uri)
	defer func() {
		if err != nil {
			connEntry.Close()
			t.connPool.Free(connEntry)
		}
	}()
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
		if t.writerBuffer != nil {
			if err = conn.(*net.TCPConn).SetWriteBuffer(t.writerBuffer.(int)); err != nil {
				return nil, err
			}
		}
		if t.deadline != nil {
			if err = conn.SetDeadline(t.deadline.(time.Time)); err != nil {
				return nil, err
			}
		}
		if t.readDeadline != nil {
			if err = conn.SetReadDeadline(t.readDeadline.(time.Time)); err != nil {
				return nil, err
			}
		}
		if t.writerDeadline != nil {
			if err = conn.SetWriteDeadline(t.writerDeadline.(time.Time)); err != nil {
				return nil, err
			}
		}
		if t.config != nil {
			conn = tls.Client(conn, t.config)
		}
		connEntry.Set(conn)
	}
	return connEntry, err
}

func (t *TcpTransporter) SendData(context interface{}, data []byte, success bool) (err error) {
	connEntry := context.(*TcpConnEntry)
	if success {
		if err = writeContentLength(connEntry.conn, len(data)); err == nil {
			_, err = connEntry.conn.Write(data)
		}
		if err != nil {
			success = false
		}
	}
	if !success {
		connEntry.Close()
		t.connPool.Free(connEntry)
	}
	return err
}

func (t *TcpTransporter) GetInputStream(context interface{}) ([]byte, error) {
	connEntry := context.(*TcpConnEntry)
	conn := connEntry.conn
	n, err := readContentLength(conn)
	if err != nil {
		connEntry.Close()
		t.connPool.Free(connEntry)
		return nil, err
	}
	data := make([]byte, n)
	if _, err = io.ReadAtLeast(conn, data, n); err != nil {
		connEntry.Close()
		t.connPool.Free(connEntry)
	}
	return data, err
}

func (t *TcpTransporter) EndInvoke(context interface{}, success bool) error {
	connEntry := context.(*TcpConnEntry)
	if !success {
		connEntry.Close()
	}
	t.connPool.Free(connEntry)
	return nil
}
