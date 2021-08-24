/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/websocket/handler.go                                 |
|                                                          |
| LastModified: Aug 24, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package websocket

import (
	"context"
	"math/rand"
	"net"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/hprose/hprose-golang/v3/rpc/core"
	rpchttp "github.com/hprose/hprose-golang/v3/rpc/http"
)

type Handler struct {
	rpchttp.Handler
	OnAccept func(*websocket.Conn) *websocket.Conn
	OnClose  func(*websocket.Conn)
	OnError  func(*websocket.Conn, error)
}

func (h *Handler) onAccept(conn *websocket.Conn) *websocket.Conn {
	if h.OnAccept != nil {
		return h.OnAccept(conn)
	}
	return conn
}

func (h *Handler) onClose(conn *websocket.Conn) {
	if h.OnClose != nil {
		h.OnClose(conn)
	}
}

func (h *Handler) onError(conn *websocket.Conn, err error) {
	if h.OnError != nil {
		h.OnError(conn, err)
	}
}

// BindContext to the websocket server.
func (h *Handler) BindContext(ctx context.Context, server core.Server) {
	s := server.(*http.Server)
	s.Handler = h
	s.BaseContext = func(l net.Listener) context.Context {
		return ctx
	}
}

// ServeHTTP implements the http.Handler interface.
func (h *Handler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if !websocket.IsWebSocketUpgrade(request) {
		h.Handler.ServeHTTP(response, request)
		return
	}
	upgrader := websocket.Upgrader{
		Subprotocols: []string{"hprose"},
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("origin")
			if origin != "" && origin != "null" {
				if len(h.AccessControlAllowOrigins) == 0 ||
					h.AccessControlAllowOrigins[origin] {
					return true
				}
				return false
			}
			return true
		},
	}
	conn, err := upgrader.Upgrade(response, request, nil)
	if err != nil {
		h.onError(conn, err)
		return
	}
	h.Serve(request.Context(), conn)
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

func (h *Handler) getServiceContext(conn *websocket.Conn) *core.ServiceContext {
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

func (h *Handler) receive(ctx context.Context, conn *websocket.Conn, queue chan data, errChan chan error) {
	defer h.catch(ctx, errChan)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			messageType, data, err := conn.ReadMessage()
			if err != nil {
				h.reportError(ctx, errChan, err)
				return
			}
			switch messageType {
			case websocket.CloseMessage:
				h.reportError(ctx, errChan, core.ErrClosed)
				return
			case websocket.BinaryMessage:
			default:
				continue
			}
			index, ok := parseHeader(data[:4])
			if !ok {
				h.reportError(ctx, errChan, core.InvalidRequestError{})
				return
			}
			body := data[4:]
			if len(body) > h.Service.MaxRequestLength {
				h.sendResponse(ctx, queue, index, nil, core.ErrRequestEntityTooLarge)
				return
			}
			go h.run(core.WithContext(ctx, h.getServiceContext(conn)), queue, index, body)
		}
	}
}

func (h *Handler) send(ctx context.Context, conn *websocket.Conn, queue chan data, errChan chan error) {
	defer h.catch(ctx, errChan)
	for {
		select {
		case <-ctx.Done():
			return
		case response := <-queue:
			index, body, e := response.Index, response.Body, response.Error
			if e != nil {
				index |= int(0x80000000)
				if e == core.ErrRequestEntityTooLarge {
					body = []byte(core.RequestEntityTooLarge)
				} else {
					body = []byte(e.Error())
				}
			}
			header := makeHeader(index)
			writer, err := conn.NextWriter(websocket.BinaryMessage)
			if err == nil {
				_, err = writer.Write(header[:])
				if err == nil {
					_, err = writer.Write(body)
					if err == nil {
						err = writer.Close()
					}
				}
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

func (h *Handler) Serve(ctx context.Context, conn *websocket.Conn) {
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
		Handler: rpchttp.Handler{
			Service:                   service,
			P3P:                       true,
			GET:                       true,
			CrossDomain:               true,
			AccessControlAllowOrigins: make(map[string]bool),
			LastModified:              time.Now().UTC().Format(time.RFC1123),
			Etag:                      `"` + strconv.FormatInt(rand.Int63(), 16) + `"`,
		},
	}
}

func RegisterHandler() {
	core.RegisterHandler("websocket", handlerFactory{
		[]reflect.Type{
			reflect.TypeOf((*http.Server)(nil)),
		},
	})
}
