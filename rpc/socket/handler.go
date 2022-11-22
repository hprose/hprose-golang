/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/socket/handler.go                                    |
|                                                          |
| LastModified: Nov 22, 2022                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package socket

import (
	"context"
	"crypto/tls"
	"io"
	"math"
	"net"
	"reflect"
	"time"

	"github.com/hprose/hprose-golang/v3/internal/convert"
	"github.com/hprose/hprose-golang/v3/rpc/core"
)

type Handler struct {
	Service  *core.Service
	Pool     core.WorkerPool
	OnAccept func(net.Conn) net.Conn
	OnClose  func(net.Conn)
	OnError  func(net.Conn, error)
}

// BindContext to the http server.
func (h *Handler) BindContext(ctx context.Context, server core.Server) {
	go h.bind(ctx, server.(net.Listener))
}

func (h *Handler) bind(ctx context.Context, listener net.Listener) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	var tempDelay time.Duration // how long to sleep on accept failure
	for {
		conn, err := listener.Accept()
		if err != nil {
			tempDelay = nextTempDelay(err, h.onError, tempDelay)
			if tempDelay > 0 {
				continue
			}
			return
		}
		tempDelay = 0
		go h.Serve(ctx, conn)
	}
}

func (h *Handler) onAccept(conn net.Conn) net.Conn {
	if h.OnAccept != nil {
		return h.OnAccept(conn)
	}
	return conn
}

func (h *Handler) onClose(conn net.Conn) {
	if h.OnClose != nil {
		h.OnClose(conn)
	}
}

func (h *Handler) onError(conn net.Conn, err error) {
	if h.OnError != nil {
		h.OnError(conn, err)
	}
}

func (h *Handler) reportError(ctx context.Context, errChan chan error, err error) {
	select {
	case <-ctx.Done():
	case errChan <- err:
	default:
	}
}

func (h *Handler) sendResponse(ctx context.Context, queue chan data, index int, body []byte, err error) {
	select {
	case <-ctx.Done():
	case queue <- data{
		Index: index,
		Body:  body,
		Error: err,
	}:
	}
}

func (h *Handler) getServiceContext(conn net.Conn) *core.ServiceContext {
	serviceContext := core.NewServiceContext(h.Service)
	serviceContext.Items().Set("conn", conn)
	serviceContext.LocalAddr = conn.LocalAddr()
	serviceContext.RemoteAddr = conn.RemoteAddr()
	serviceContext.Handler = h
	return serviceContext
}

func (h *Handler) run(ctx context.Context, queue chan data, index int, body []byte) {
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = core.NewPanicError(e)
		}
		h.sendResponse(ctx, queue, index, body, err)
	}()
	body, err = h.Service.Handle(ctx, body)
}

func (h *Handler) task(ctx context.Context, queue chan data, index int, body []byte) func() {
	return func() {
		h.run(ctx, queue, index, body)
	}
}

func (h *Handler) catch(ctx context.Context, errChan chan error) {
	if e := recover(); e != nil {
		h.reportError(ctx, errChan, core.NewPanicError(e))
	}
}

func (h *Handler) receive(ctx context.Context, conn net.Conn, queue chan data, errChan chan error) {
	defer h.catch(ctx, errChan)
	var header [12]byte
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if _, err := io.ReadAtLeast(conn, header[:], 12); err != nil {
				h.reportError(ctx, errChan, err)
				return
			}
			length, index, ok := parseHeader(header)
			if length == 0 && index == -1 && !ok {
				h.reportError(ctx, errChan, core.InvalidRequestError{})
				return
			}
			if length > h.Service.MaxRequestLength {
				h.sendResponse(ctx, queue, index, nil, core.ErrRequestEntityTooLarge)
				return
			}
			body := make([]byte, length)
			if _, err := io.ReadAtLeast(conn, body, length); err != nil {
				h.reportError(ctx, errChan, err)
				return
			}
			if h.Pool != nil {
				h.Pool.Submit(h.task(core.WithContext(ctx, h.getServiceContext(conn)), queue, index, body))
			} else {
				go h.run(core.WithContext(ctx, h.getServiceContext(conn)), queue, index, body)
			}
		}
	}
}

func (h *Handler) send(ctx context.Context, conn net.Conn, queue chan data, errChan chan error) {
	defer h.catch(ctx, errChan)
	for {
		select {
		case <-ctx.Done():
			return
		case response := <-queue:
			index, body, e := response.Index, response.Body, response.Error
			if e != nil {
				index |= math.MinInt32
				if e == core.ErrRequestEntityTooLarge {
					body = convert.ToUnsafeBytes(core.RequestEntityTooLarge)
				} else {
					body = convert.ToUnsafeBytes(e.Error())
				}
			}
			header := makeHeader(len(body), index)
			_, err := conn.Write(header[:])
			if err == nil {
				_, err = conn.Write(body)
			}
			if err != nil {
				h.reportError(ctx, errChan, err)
				return
			}
			if e != nil {
				h.reportError(ctx, errChan, e)
				return
			}
		}
	}
}

func (h *Handler) Serve(ctx context.Context, conn net.Conn) {
	if conn = h.onAccept(conn); conn == nil {
		return
	}
	ctx, cancel := context.WithCancel(ctx)
	var err error
	defer func() {
		cancel()
		if e := recover(); e != nil {
			err = core.NewPanicError(e)
		}
		if err != nil {
			h.onError(conn, err)
		}
		h.onClose(conn)
		conn.Close()
	}()
	queue := make(chan data)
	errChan := make(chan error, 1)
	go h.receive(ctx, conn, queue, errChan)
	go h.send(ctx, conn, queue, errChan)
	select {
	case <-ctx.Done():
		err = ctx.Err()
	case err = <-errChan:
	}
}

type handlerFactory struct {
	serverTypes []reflect.Type
}

func (factory handlerFactory) ServerTypes() []reflect.Type {
	return factory.serverTypes
}

func (factory handlerFactory) New(service *core.Service) core.Handler {
	return &Handler{
		Service: service,
	}
}

func RegisterHandler() {
	core.RegisterHandler("socket", handlerFactory{
		[]reflect.Type{
			reflect.TypeOf((*net.TCPListener)(nil)),
			reflect.TypeOf((*net.UnixListener)(nil)),
			reflect.TypeOf(tls.NewListener(nil, nil)),
		},
	})
}
