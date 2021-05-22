/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/udp/transport.go                                     |
|                                                          |
| LastModified: May 22, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package udp

import (
	"context"
	"errors"
	"net"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

type conn struct {
	net.Conn
	requests chan data
	results  map[int]chan data
	lock     sync.Mutex
	counter  int32
	onClose  func(net.Conn)
	once     sync.Once
}

// ErrClosed represents a error.
var ErrClosed = errors.New("network closed")

func dial(ctx context.Context) (net.Conn, error) {
	u := core.GetClientContext(ctx).URL
	var d net.Dialer
	switch u.Scheme {
	case "udp", "udp4", "udp6":
		address := u.Host
		if u.Port() == "" {
			address += ":8412"
		}
		return d.DialContext(ctx, "udp", address)
	}
	return nil, core.UnsupportedProtocolError{Scheme: u.Scheme}
}

func newConn(ctx context.Context, onConnect func(net.Conn) net.Conn, onClose func(net.Conn)) (*conn, error) {
	c, err := dial(ctx)
	if err != nil {
		return nil, err
	}
	return &conn{
		Conn:     onConnect(c),
		requests: make(chan data),
		onClose:  onClose,
		results:  make(map[int]chan data),
	}, nil
}

func (c *conn) store(index int, resultChan chan data) {
	c.lock.Lock()
	c.results[index] = resultChan
	c.lock.Unlock()
}

func (c *conn) delete(index int) {
	c.lock.Lock()
	delete(c.results, index)
	c.lock.Unlock()
}

func (c *conn) loadAndDelete(index int) (resultChan chan data, loaded bool) {
	c.lock.Lock()
	if resultChan, loaded = c.results[index]; loaded {
		delete(c.results, index)
	}
	c.lock.Unlock()
	return
}

func (c *conn) rangeAndClean(f func(index int, resultChan chan data)) {
	c.lock.Lock()
	for len(c.results) > 0 {
		results := c.results
		c.results = make(map[int]chan data)
		c.lock.Unlock()
		for index, resultChan := range results {
			f(index, resultChan)
		}
		runtime.Gosched()
		c.lock.Lock()
	}
	c.lock.Unlock()
}

func (c *conn) Transport(ctx context.Context, request []byte) (response []byte, err error) {
	index := int(atomic.AddInt32(&c.counter, 1) & 0x7fff)
	resultChan := make(chan data, 1)
	c.store(index, resultChan)
	select {
	case <-ctx.Done():
		c.delete(index)
		return nil, ctx.Err()
	case c.requests <- data{
		Index: index,
		Body:  request,
	}:
	case res := <-resultChan:
		return res.Body, res.Error
	}
	select {
	case <-ctx.Done():
		c.delete(index)
		return nil, ctx.Err()
	case res := <-resultChan:
		return res.Body, res.Error
	}
}

func (c *conn) Exit(onExit func(), err error) {
	onExit()
	if e := recover(); e != nil {
		err = core.NewPanicError(e)
	}
	if err != nil {
		c.Close(err)
	}
}

func (c *conn) send(request data) (err error) {
	var buffer [65507]byte
	header := makeHeader(len(request.Body), request.Index)
	copy(buffer[:], header[:])
	copy(buffer[8:], request.Body)
	_, err = c.Write(buffer[:8+len(request.Body)])
	return
}

func (c *conn) Send(ctx context.Context, onExit func()) {
	var err error
	defer func() {
		c.Exit(onExit, err)
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case request := <-c.requests:
			if err = c.send(request); err != nil {
				return
			}
		}
	}
}

func (c *conn) receive() (err error) {
	var buffer [65507]byte
	var n int
	switch n, err = c.Read(buffer[:]); {
	case err != nil:
	case n < 8:
		err = core.InvalidResponseError{}
	default:
		switch length, index, ok := parseHeader(buffer[:8]); {
		case length == 0 && index == -1 && !ok:
			err = core.InvalidResponseError{}
		default:
			body := make([]byte, length)
			copy(body, buffer[8:])
			if !ok {
				if string(body) == core.RequestEntityTooLarge {
					err = core.ErrRequestEntityTooLarge
				} else {
					err = core.InvalidResponseError{Response: body}
				}
			} else if resultChan, loaded := c.loadAndDelete(index); loaded {
				resultChan <- data{
					Index: index,
					Body:  body,
				}
			}
		}
	}
	return
}

func (c *conn) Receive(ctx context.Context, onExit func()) {
	var err error
	defer func() {
		c.Exit(onExit, err)
	}()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if err = c.receive(); err != nil {
				return
			}
		}
	}
}

func (c *conn) Close(err error) {
	c.once.Do(func() {
		c.onClose(c.Conn)
		_ = c.Conn.Close()
	})
	c.rangeAndClean(func(index int, resultChan chan data) {
		resultChan <- data{
			Index: index,
			Error: err,
		}
	})
}

type Transport struct {
	OnConnect func(net.Conn) net.Conn
	OnClose   func(net.Conn)
	conns     map[string]*conn
	lock      sync.RWMutex
}

func (trans *Transport) getConn(ctx context.Context) (conn *conn, err error) {
	u := core.GetClientContext(ctx).URL
	key := u.String()
	trans.lock.RLock()
	if conn = trans.conns[key]; conn != nil {
		trans.lock.RUnlock()
		return
	}
	trans.lock.RUnlock()
	trans.lock.Lock()
	defer trans.lock.Unlock()
	if conn = trans.conns[key]; conn != nil {
		return
	}
	conn, err = newConn(ctx, trans.onConnect, trans.onClose)
	if err != nil {
		return
	}
	trans.conns[key] = conn
	ctx, cancel := context.WithCancel(context.Background())
	onExit := func() {
		trans.lock.Lock()
		if trans.conns[key] == conn {
			delete(trans.conns, key)
			cancel()
		}
		trans.lock.Unlock()
	}
	go conn.Send(ctx, onExit)
	go conn.Receive(ctx, onExit)
	return
}

func (trans *Transport) onConnect(conn net.Conn) net.Conn {
	if trans.OnConnect != nil {
		return trans.OnConnect(conn)
	}
	return conn
}

func (trans *Transport) onClose(conn net.Conn) {
	if trans.OnClose != nil {
		trans.OnClose(conn)
	}
}

func (trans *Transport) Transport(ctx context.Context, request []byte) ([]byte, error) {
	conn, err := trans.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return conn.Transport(ctx, request)
}

func (trans *Transport) Abort() {
	trans.lock.Lock()
	conns := trans.conns
	trans.conns = make(map[string]*conn)
	trans.lock.Unlock()
	for _, conn := range conns {
		conn.Close(ErrClosed)
	}
}

type transportFactory struct {
	schemes []string
}

func (factory transportFactory) Schemes() []string {
	return factory.schemes
}

func (factory transportFactory) New() core.Transport {
	transport := &Transport{
		conns: make(map[string]*conn),
	}
	return transport
}

func RegisterTransport() {
	core.RegisterTransport("udp", transportFactory{[]string{"udp", "udp4", "udp6"}})
}
