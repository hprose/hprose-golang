/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/websocket/transport.go                               |
|                                                          |
| LastModified: May 5, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package websocket

import (
	"context"
	"net/http"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/gorilla/websocket"
	"github.com/hprose/hprose-golang/v3/rpc/core"
)

type conn struct {
	*websocket.Conn
	requests chan data
	results  map[int]chan data
	lock     sync.Mutex
	counter  int32
	onClose  func(*websocket.Conn)
	once     sync.Once
}

func dial(ctx context.Context) (*websocket.Conn, error) {
	u := core.GetClientContext(ctx).URL
	var d websocket.Dialer
	switch u.Scheme {
	case "ws", "wss":
		header := http.Header{"Sec-WebSocket-Protocol": []string{"hprose"}}
		conn, response, err := d.DialContext(ctx, u.String(), header)
		if response != nil {
			response.Body.Close()
		}
		return conn, err
	}
	return nil, core.UnsupportedProtocolError{Scheme: u.Scheme}
}

func newConn(ctx context.Context, onConnect func(*websocket.Conn) *websocket.Conn, onClose func(*websocket.Conn)) (*conn, error) {
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
	index := int(atomic.AddInt32(&c.counter, 1) & 0x7fffffff)
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

func (c *conn) Send(onExit func()) {
	var err error
	defer func() {
		c.Exit(onExit, err)
	}()
	for request := range c.requests {
		header := makeHeader(request.Index)
		writer, err := c.NextWriter(websocket.BinaryMessage)
		if err == nil {
			_, err = writer.Write(header[:])
			if err == nil {
				_, err = writer.Write(request.Body)
				if err == nil {
					err = writer.Close()
				}
			}
		}
		if err != nil {
			return
		}
	}
}

func (c *conn) Receive(onExit func()) {
	var (
		messageType int
		body        []byte
		err         error
	)
	defer func() {
		c.Exit(onExit, err)
	}()
	for {
		messageType, body, err = c.ReadMessage()
		if err != nil {
			return
		}
		switch messageType {
		case websocket.CloseMessage:
			err = core.ErrClosed
			return
		case websocket.BinaryMessage:
		default:
			continue
		}
		index, ok := parseHeader(body[:4])
		body = body[4:]
		if !ok {
			if string(body) == core.RequestEntityTooLarge {
				err = core.ErrRequestEntityTooLarge
			} else {
				err = core.InvalidResponseError{Response: body}
			}
			return
		}
		if resultChan, loaded := c.loadAndDelete(index); loaded {
			resultChan <- data{
				Index: index,
				Body:  body,
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
	OnConnect func(*websocket.Conn) *websocket.Conn
	OnClose   func(*websocket.Conn)
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
	onExit := func() {
		trans.lock.Lock()
		if trans.conns[key] == conn {
			delete(trans.conns, key)
		}
		trans.lock.Unlock()
	}
	go conn.Send(onExit)
	go conn.Receive(onExit)
	return
}

func (trans *Transport) onConnect(conn *websocket.Conn) *websocket.Conn {
	if trans.OnConnect != nil {
		return trans.OnConnect(conn)
	}
	return conn
}

func (trans *Transport) onClose(conn *websocket.Conn) {
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
		conn.Close(core.ErrClosed)
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
	core.RegisterTransport("websocket", transportFactory{[]string{"ws", "wss"}})
}
