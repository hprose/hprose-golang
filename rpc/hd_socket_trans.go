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
 * rpc/hd_socket_trans.go                                 *
 *                                                        *
 * hprose half duplex socket transport for Go.            *
 *                                                        *
 * LastModified: Jan 7, 2017                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import (
	"net"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type halfDuplexConnEntry struct {
	conn  net.Conn
	timer *time.Timer
}

type halfDuplexSocketTransport struct {
	idleTimeout time.Duration
	connPool    chan *halfDuplexConnEntry
	connCount   int32
	nextid      uint32
	createConn  func() (net.Conn, error)
	cond        sync.Cond
}

func newHalfDuplexSocketTransport() (hd *halfDuplexSocketTransport) {
	hd = new(halfDuplexSocketTransport)
	hd.idleTimeout = 0
	hd.connPool = make(chan *halfDuplexConnEntry, runtime.NumCPU())
	hd.connCount = 0
	hd.nextid = 0
	hd.cond.L = &sync.Mutex{}
	return
}

func (hd *halfDuplexSocketTransport) setCreateConn(createConn func() (net.Conn, error)) {
	hd.createConn = createConn
}

// IdleTimeout returns the conn pool idle timeout of hprose socket client
func (hd *halfDuplexSocketTransport) IdleTimeout() time.Duration {
	return hd.idleTimeout
}

// SetIdleTimeout sets the conn pool idle timeout of hprose socket client
func (hd *halfDuplexSocketTransport) SetIdleTimeout(timeout time.Duration) {
	hd.idleTimeout = timeout
}

// MaxPoolSize returns the max conn pool size of hprose socket client
func (hd *halfDuplexSocketTransport) MaxPoolSize() int {
	return cap(hd.connPool)
}

// SetMaxPoolSize sets the max conn pool size of hprose socket client
func (hd *halfDuplexSocketTransport) SetMaxPoolSize(size int) {
	if size > 0 {
		pool := make(chan *halfDuplexConnEntry, size)
		for i := 0; i < len(hd.connPool); i++ {
			select {
			case pool <- <-hd.connPool:
			default:
			}
		}
		hd.connPool = pool
	}
}

func (hd *halfDuplexSocketTransport) getConn() *halfDuplexConnEntry {
	for hd.connPool != nil {
		select {
		case entry, ok := <-hd.connPool:
			if !ok {
				panic(ErrClientIsAlreadyClosed)
			}
			if entry.timer != nil {
				entry.timer.Stop()
			}
			if entry.conn != nil {
				return entry
			}
			continue
		default:
			return nil
		}
	}
	panic(ErrClientIsAlreadyClosed)
}

func (hd *halfDuplexSocketTransport) fetchConn() (*halfDuplexConnEntry, error) {
	hd.cond.L.Lock()
	for {
		entry := hd.getConn()
		if entry != nil && entry.conn != nil {
			hd.cond.L.Unlock()
			return entry, nil
		}
		if int(atomic.AddInt32(&hd.connCount, 1)) <= cap(hd.connPool) {
			hd.cond.L.Unlock()
			conn, err := hd.createConn()
			if err == nil {
				return &halfDuplexConnEntry{conn: conn}, nil
			}
			atomic.AddInt32(&hd.connCount, -1)
			return nil, err
		}
		atomic.AddInt32(&hd.connCount, -1)
		hd.cond.Wait()
	}
}

func (hd *halfDuplexSocketTransport) close() {
	if hd.connPool != nil {
		connPool := hd.connPool
		hd.connPool = nil
		hd.connCount = 0
		close(connPool)
		for entry := range connPool {
			if entry.timer != nil {
				entry.timer.Stop()
			}
			if entry.conn != nil {
				entry.conn.Close()
			}
		}
	}
}

func (hd *halfDuplexSocketTransport) closeConn(conn net.Conn) {
	conn.Close()
	atomic.AddInt32(&hd.connCount, -1)
}

func (hd *halfDuplexSocketTransport) sendAndReceive(
	data []byte, context *ClientContext) ([]byte, error) {
	entry, err := hd.fetchConn()
	if err != nil {
		hd.cond.Signal()
		return nil, err
	}
	conn := entry.conn
	err = conn.SetDeadline(time.Now().Add(context.Timeout))
	if err == nil {
		err = hdSendData(conn, data)
	}
	if err == nil {
		data, err = hdRecvData(conn, data)
	}
	if err == nil {
		err = conn.SetDeadline(time.Time{})
	}
	if err != nil {
		hd.closeConn(conn)
		hd.cond.Signal()
		return nil, err
	}
	if hd.idleTimeout > 0 {
		if entry.timer == nil {
			entry.timer = time.AfterFunc(hd.idleTimeout, func() {
				hd.closeConn(conn)
				entry.conn = nil
				entry.timer = nil
			})
		} else {
			entry.timer.Reset(hd.idleTimeout)
		}
	}
	hd.connPool <- entry
	hd.cond.Signal()
	return data, nil
}
