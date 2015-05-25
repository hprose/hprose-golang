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
 * LastModified: May 25, 2015                             *
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

// StreamClient is base struct for TcpClient and UnixClient
type StreamClient struct {
	*BaseClient
	timeout      interface{}
	readBuffer   interface{}
	readTimeout  interface{}
	writeBuffer  interface{}
	writeTimeout interface{}
}

func newStreamClient(trans Transporter) (client *StreamClient) {
	client = new(StreamClient)
	client.BaseClient = NewBaseClient(trans)
	return
}

type streamConnStatus int

const (
	free = streamConnStatus(iota)
	using
	closing
)

// StreamConnEntry is the connection entry in connection pool
type StreamConnEntry struct {
	uri          string
	conn         net.Conn
	status       streamConnStatus
	lastUsedTime time.Time
}

// NewStreamConnEntry is the constructor for StreamConnEntry
func NewStreamConnEntry(uri string) (entry *StreamConnEntry) {
	entry = new(StreamConnEntry)
	entry.uri = uri
	entry.status = using
	entry.lastUsedTime = time.Now()
	return
}

// Get the connection
func (connEntry *StreamConnEntry) Get() net.Conn {
	return connEntry.conn
}

// Set the connection
func (connEntry *StreamConnEntry) Set(conn net.Conn) {
	if conn != nil {
		connEntry.conn = conn
	}
}

// Close the connection
func (connEntry *StreamConnEntry) Close() {
	connEntry.status = closing
}

// StreamConnPool is the connection pool
type StreamConnPool struct {
	sync.Mutex
	pool    []*StreamConnEntry
	timer   *time.Ticker
	timeout time.Duration
}

// NewStreamConnPool is the constructor for StreamConnPool
func NewStreamConnPool(num int) (pool *StreamConnPool) {
	pool = new(StreamConnPool)
	pool.pool = make([]*StreamConnEntry, 0, num)
	return
}

func freeConns(conns []net.Conn) {
	for _, conn := range conns {
		conn.Close()
	}
}

// Timeout return the timeout of the connection in pool
func (connPool *StreamConnPool) Timeout() time.Duration {
	return connPool.timeout
}

// SetTimeout for connection in pool
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
		func() {
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
		}()
	}
}

// Get the StreamConnEntry in StreamConnPool
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
	entry := NewStreamConnEntry(uri)
	connPool.pool = append(connPool.pool, entry)
	return entry
}

// Close the specify uri connections in StreamConnPool
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

// Free the entry to pool
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

// SetReadBuffer sets the size of the operating system's receive buffer associated with the connection.
func (client *StreamClient) SetReadBuffer(bytes int) {
	client.readBuffer = bytes
}

// SetReadTimeout for stream client
func (client *StreamClient) SetReadTimeout(d time.Duration) {
	client.readTimeout = d
}

// SetWriteBuffer sets the size of the operating system's transmit buffer associated with the connection.
func (client *StreamClient) SetWriteBuffer(bytes int) {
	client.writeBuffer = bytes
}

// SetWriteTimeout for stream client
func (client *StreamClient) SetWriteTimeout(d time.Duration) {
	client.writeTimeout = d
}
