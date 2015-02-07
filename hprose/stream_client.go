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
 * hprose/stream_client.go                                *
 *                                                        *
 * hprose stream client for Go.                           *
 *                                                        *
 * LastModified: Jan 28, 2015                             *
 * Authors: Ma Bingyao <andot@hprose.com>                 *
 *          Ore_Ash <nanohugh@gmail.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"net"
	"sync"
	"time"
)

type StreamClient struct {
	*BaseClient
	timeout      interface{}
	readBuffer   interface{}
	readTimeout  interface{}
	writeBuffer  interface{}
	writeTimeout interface{}
}

func newStreamClient(trans Transporter) *StreamClient {
	return &StreamClient{
		BaseClient: NewBaseClient(trans),
	}
}

type streamConnStatus int

const (
	free = streamConnStatus(iota)
	using
	closing
)

type StreamConnEntry struct {
	uri          string
	conn         net.Conn
	status       streamConnStatus
	lastUsedTime time.Time
}

func (connEntry *StreamConnEntry) Get() net.Conn {
	return connEntry.conn
}

func (connEntry *StreamConnEntry) Set(conn net.Conn) {
	if conn != nil {
		connEntry.conn = conn
	}
}

func (connEntry *StreamConnEntry) Close() {
	connEntry.status = closing
}

type StreamConnPool struct {
	sync.Mutex
	pool    []*StreamConnEntry
	timer   *time.Ticker
	timeout time.Duration
}

func freeConns(conns []net.Conn) {
	for _, conn := range conns {
		conn.Close()
	}
}

func (connPool *StreamConnPool) Timeout() time.Duration {
	return connPool.timeout
}

func (connPool *StreamConnPool) SetTimeout(d time.Duration) {
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

func (connPool *StreamConnPool) closeTimeoutConns() {
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

func (connPool *StreamConnPool) Get(uri string) *StreamConnEntry {
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
	entry := &StreamConnEntry{uri, nil, using, time.Now()}
	connPool.pool = append(connPool.pool, entry)
	return entry
}

func (connPool *StreamConnPool) Close(uri string) {
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

func (connPool *StreamConnPool) Free(entry *StreamConnEntry) {
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

func (client *StreamClient) SetReadBuffer(bytes int) {
	client.readBuffer = bytes
}

func (client *StreamClient) SetReadTimeout(d time.Duration) {
	client.readTimeout = d
}

func (client *StreamClient) SetWriteBuffer(bytes int) {
	client.writeBuffer = bytes
}

func (client *StreamClient) SetWriteTimeout(d time.Duration) {
	client.writeTimeout = d
}
