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
 * hprose/unix_client.go                                   *
 *                                                        *
 * hprose unix client for Go.                              *
 *                                                        *
 * LastModified: Oct 15, 2014                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"net"
	"strings"
	"sync"
	"time"
)

type UnixClient struct {
	*BaseClient
	timeout      interface{}
	readBuffer   interface{}
	readTimeout  interface{}
	writeBuffer  interface{}
	writeTimeout interface{}
}

type UnixConnEntry struct {
	uri          string
	conn         net.Conn
	status       sockConnStatus
	lastUsedTime time.Time
}

func (connEntry *UnixConnEntry) Get() net.Conn {
	return connEntry.conn
}

func (connEntry *UnixConnEntry) Set(conn net.Conn) {
	if conn != nil {
		connEntry.conn = conn
	}
}

func (connEntry *UnixConnEntry) Close() {
	connEntry.status = closing
}

type UnixConnPool struct {
	sync.Mutex
	pool    []*UnixConnEntry
	timer   *time.Ticker
	timeout time.Duration
}

func (connPool *UnixConnPool) Timeout() time.Duration {
	return connPool.timeout
}

func (connPool *UnixConnPool) SetTimeout(d time.Duration) {
	if connPool.timer != nil {
		connPool.timer.Stop()
		connPool.timer = nil
	}
	connPool.timeout = d
	if d > 0 {
		connPool.timer = time.NewTicker(d)
		go connPool.closeTimeoutConns()
	}
}

func (connPool *UnixConnPool) closeTimeoutConns() {
	for t := range connPool.timer.C {
		connPool.Lock()
		defer connPool.Unlock()
		conns := make([]net.Conn, 0, len(connPool.pool))
		for _, entry := range connPool.pool {
			if entry.uri != "" &&
				entry.status == free &&
				entry.conn != nil &&
				t.After(entry.lastUsedTime.Add(connPool.timeout)) {
				conns = append(conns, entry.conn)
				entry.conn = nil
				entry.uri = ""
			}
		}
		go freeConns(conns)
	}
}

func (connPool *UnixConnPool) Get(uri string) *UnixConnEntry {
	connPool.Lock()
	defer connPool.Unlock()
	for _, entry := range connPool.pool {
		if entry.status == free {
			if entry.uri == uri {
				entry.status = using
				return entry
			} else if entry.uri == "" {
				entry.status = using
				entry.uri = uri
				return entry
			}
		}
	}
	entry := &UnixConnEntry{uri, nil, using, time.Now()}
	connPool.pool = append(connPool.pool, entry)
	return entry
}

func (connPool *UnixConnPool) Close(uri string) {
	connPool.Lock()
	defer connPool.Unlock()
	conns := make([]net.Conn, 0, len(connPool.pool))
	for _, entry := range connPool.pool {
		if entry.uri == uri {
			if entry.status == free {
				conns = append(conns, entry.conn)
				entry.conn = nil
				entry.uri = ""
			} else {
				entry.Close()
			}
		}
	}
	go freeConns(conns)
}

func (connPool *UnixConnPool) Free(entry *UnixConnEntry) {
	if entry.status == closing {
		if entry.conn != nil {
			go entry.conn.Close()
			entry.conn = nil
		}
		entry.uri = ""
	}
	entry.lastUsedTime = time.Now()
	entry.status = free
}

type UnixTransporter struct {
	connPool *UnixConnPool
	*UnixClient
}

func NewUnixClient(uri string) Client {
	trans := &UnixTransporter{connPool: &UnixConnPool{pool: make([]*UnixConnEntry, 0)}}
	client := &UnixClient{BaseClient: NewBaseClient(trans)}
	client.Client = client
	trans.UnixClient = client
	client.SetUri(uri)
	return client
}

func (client *UnixClient) SetUri(uri string) {
	scheme := strings.Split(uri, ":")[0]
	if scheme != "unix" && scheme != "unixpacket" {
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

func (client *UnixClient) SetReadBuffer(bytes int) {
	client.readBuffer = bytes
}

func (client *UnixClient) SetReadTimeout(d time.Duration) {
	client.readTimeout = d
}

func (client *UnixClient) SetWriteBuffer(bytes int) {
	client.writeBuffer = bytes
}

func (client *UnixClient) SetWriteTimeout(d time.Duration) {
	client.writeTimeout = d
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
	conn := connEntry.Get()
	if conn == nil {
		scheme, path := parseUnixUri(uri)
		var unixaddr *net.UnixAddr
		if unixaddr, err = net.ResolveUnixAddr(scheme, path); err != nil {
			return nil, err
		}
		if conn, err = net.DialUnix("unix", nil, unixaddr); err != nil {
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
		if t.timeout != nil {
			if err = conn.SetDeadline(time.Now().Add(t.timeout.(time.Duration))); err != nil {
				return nil, err
			}
		}
		connEntry.Set(conn)
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
