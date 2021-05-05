/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/udp/handler.go                                       |
|                                                          |
| LastModified: May 5, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package udp

import (
	"context"
	"net"
	"reflect"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

type Handler struct {
	Service *core.Service
	OnClose func(net.Conn)
	OnError func(error)
}

// BindContext to the http server.
func (h *Handler) BindContext(ctx context.Context, server core.Server) {
	go h.Serve(ctx, server.(*net.UDPConn))
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

func (h *Handler) sendResponse(ctx context.Context, queue chan data, index int, body []byte, err error, addr *net.UDPAddr) {
	select {
	case <-ctx.Done():
	case queue <- data{
		Index: index,
		Body:  body,
		Error: err,
		Addr:  addr,
	}:
	}
}

func (h *Handler) getServiceContext(conn *net.UDPConn, addr *net.UDPAddr) *core.ServiceContext {
	serviceContext := core.NewServiceContext(h.Service)
	serviceContext.Items().Set("conn", conn)
	serviceContext.LocalAddr = conn.LocalAddr()
	serviceContext.RemoteAddr = addr
	serviceContext.Handler = h
	return serviceContext
}

func (h *Handler) run(ctx context.Context, queue chan data, index int, body []byte, addr *net.UDPAddr) {
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = core.NewPanicError(e)
		}
		h.sendResponse(ctx, queue, index, body, err, addr)
	}()
	body, err = h.Service.Handle(ctx, body)
}

func (h *Handler) catch(ctx context.Context, errChan chan error) {
	if e := recover(); e != nil {
		h.reportError(ctx, errChan, core.NewPanicError(e))
	}
}

func (h *Handler) receive(ctx context.Context, conn *net.UDPConn, queue chan data, errChan chan error) {
	defer h.catch(ctx, errChan)
	var buffer [65507]byte
	for {
		select {
		case <-ctx.Done():
			return
		default:
			switch n, addr, err := conn.ReadFromUDP(buffer[:]); {
			case err != nil:
				if err, ok := err.(*net.OpError); ok && err.Addr == nil {
					h.reportError(ctx, errChan, err)
					return
				}
				h.onError(err)
			case n < 8:
				h.onError(core.InvalidRequestError{})
			default:
				switch length, index, ok := parseHeader(buffer[:8]); {
				case length == 0 && index == -1 && !ok:
					h.onError(core.InvalidRequestError{})
				case length > h.Service.MaxRequestLength:
					h.sendResponse(ctx, queue, index, nil, core.ErrRequestEntityTooLarge, addr)
				default:
					body := make([]byte, length)
					copy(body, buffer[8:])
					go h.run(core.WithContext(ctx, h.getServiceContext(conn, addr)), queue, index, body, addr)
				}
			}
		}
	}
}

func (h *Handler) send(ctx context.Context, conn *net.UDPConn, queue chan data, errChan chan error) {
	defer h.catch(ctx, errChan)
	var buffer [65507]byte
	for {
		select {
		case <-ctx.Done():
			return
		case response := <-queue:
			index, body, e, addr := response.Index, response.Body, response.Error, response.Addr
			if e != nil {
				index |= 0x8000
				if e == core.ErrRequestEntityTooLarge {
					body = []byte(core.RequestEntityTooLarge)
				} else {
					body = []byte(e.Error())
				}
				h.onError(e)
			}
			header := makeHeader(len(body), index)
			copy(buffer[:], header[:])
			copy(buffer[8:], body)
			if _, err := conn.WriteToUDP(buffer[:8+len(body)], addr); err != nil {
				if err, ok := err.(*net.OpError); ok && err.Addr == nil {
					h.reportError(ctx, errChan, err)
					return
				}
				h.onError(err)
			}
		}
	}
}

func (h *Handler) Serve(ctx context.Context, conn *net.UDPConn) {
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

func init() {
	core.RegisterHandler("udp", handlerFactory{
		[]reflect.Type{
			reflect.TypeOf((*net.UDPConn)(nil)),
		},
	})
}
