/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/socket/handler.go                                    |
|                                                          |
| LastModified: Apr 29, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package socket

import (
	"context"
	"io"
	"net"
	"reflect"
	"time"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

type Handler struct {
	Service  *core.Service
	OnAccept func(net.Conn)
	OnClose  func(net.Conn)
	OnError  func(error)
}

// Bind to the http server.
func (h *Handler) Bind(server core.Server) {
	go h.bind(server.(net.Listener))
}

func (h *Handler) bind(listener net.Listener) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var tempDelay time.Duration // how long to sleep on accept failure
	for {
		conn, err := listener.Accept()
		if err != nil {
			tempDelay = nextTempDelay(err, h.onError, tempDelay)
			if tempDelay > 0 {
				continue
			}
			break
		}
		tempDelay = 0
		go h.Serve(ctx, conn)
	}
}

func (h *Handler) onAccept(conn net.Conn) {
	if h.OnAccept != nil {
		h.OnAccept(conn)
	}
}

func (h *Handler) onClose(conn net.Conn) {
	if h.OnClose != nil {
		h.OnClose(conn)
	}
}

func (h *Handler) onError(err error) {
	if h.OnError != nil {
		h.OnError(err)
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

func (h *Handler) catch(ctx context.Context, errChan chan error) {
	if e := recover(); e != nil {
		h.reportError(ctx, errChan, core.NewPanicError(e))
	}
}

func (h *Handler) receive(ctx context.Context, conn net.Conn, queue chan data, errChan chan error) {
	defer h.catch(ctx, errChan)
	var header [12]byte
	for {
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
		go h.run(core.WithContext(ctx, h.getServiceContext(conn)), queue, index, body)
	}
}

func (h *Handler) send(ctx context.Context, conn net.Conn, queue chan data, errChan chan error) {
	defer h.catch(ctx, errChan)
	for {
		select {
		case <-ctx.Done():
			return
		case response := <-queue:
			index, body, err := response.Index, response.Body, response.Error
			if err != nil {
				index |= 0x80000000
				if err == core.ErrRequestEntityTooLarge {
					body = []byte(requestEntityTooLarge)
				} else {
					body = []byte(err.Error())
				}
			}
			header := makeHeader(len(body), index)
			if _, err := conn.Write(header[:]); err != nil {
				h.reportError(ctx, errChan, err)
				return
			}
			if _, err := conn.Write(body); err != nil {
				h.reportError(ctx, errChan, err)
				return
			}
			if err != nil {
				h.reportError(ctx, errChan, err)
				return
			}
		}
	}
}

func (h *Handler) Serve(ctx context.Context, conn net.Conn) {
	ctx, cancel := context.WithCancel(ctx)
	var err error
	defer func() {
		cancel()
		if e := recover(); e != nil {
			err = core.NewPanicError(e)
		}
		if err != nil {
			h.onError(err)
		}
		h.onClose(conn)
		conn.Close()
	}()
	h.onAccept(conn)
	queue := make(chan data)
	errChan := make(chan error)
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

func init() {
	core.RegisterHandler("socket", handlerFactory{
		[]reflect.Type{
			reflect.TypeOf((*net.TCPListener)(nil)),
			reflect.TypeOf((*net.UnixListener)(nil)),
		},
	})
}
