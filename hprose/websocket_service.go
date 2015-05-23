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
 * hprose/websocket_service.go                            *
 *                                                        *
 * hprose websocket service for Go.                       *
 *                                                        *
 * LastModified: May 22, 2015                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

// WebSocketContext is the hprose websocket context
type WebSocketContext struct {
	*HttpContext
	WebSocket *websocket.Conn
}

// WebSocketService is the hprose websocket service
type WebSocketService struct {
	*HttpService
	*websocket.Upgrader
}

type wsArgsFixer struct {
	httpArgsFixer
}

func (fixer wsArgsFixer) FixArgs(args []reflect.Value, lastParamType reflect.Type, context Context) []reflect.Value {
	if c, ok := context.(*WebSocketContext); ok {
		if lastParamType.String() == "*hprose.WebSocketContext" {
			return append(args, reflect.ValueOf(c))
		} else if lastParamType.String() == "*websocket.Conn" {
			return append(args, reflect.ValueOf(c.WebSocket))
		} else if lastParamType.String() == "*hprose.HttpContext" {
			return append(args, reflect.ValueOf(c.HttpContext))
		} else if lastParamType.String() == "*http.Request" {
			return append(args, reflect.ValueOf(c.Request))
		}
	}
	return fixer.httpArgsFixer.FixArgs(args, lastParamType, context)
}

// NewWebSocketService is the constructor of WebSocketService
func NewWebSocketService() *WebSocketService {
	service := &WebSocketService{HttpService: NewHttpService()}
	service.argsfixer = wsArgsFixer{}
	service.Upgrader = &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("origin")
			if origin != "" && origin != "null" {
				if len(service.accessControlAllowOrigins) == 0 || service.accessControlAllowOrigins[origin] {
					return true
				}
				return false
			}
			return true
		},
	}
	return service
}

// ServeHTTP ...
func (service *WebSocketService) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && strings.ToLower(request.Header.Get("connection")) != "upgrade" || request.Method == "POST" {
		service.HttpService.ServeHTTP(response, request)
		return
	}
	conn, err := service.Upgrade(response, request, nil)
	if err != nil {
		context := &HttpContext{BaseContext: NewBaseContext(), Response: response, Request: request}
		service.fireErrorEvent(err, context)
		return
	}
	defer conn.Close()
	mutex := new(sync.Mutex)
	for {
		context := &WebSocketContext{HttpContext: &HttpContext{BaseContext: NewBaseContext(), Response: response, Request: request}, WebSocket: conn}
		_, data, err := conn.ReadMessage()
		if err != nil {
			if err != io.EOF {
				service.fireErrorEvent(err, context)
			}
			break
		}
		go func(conn *websocket.Conn, data []byte, context *WebSocketContext) {
			id := data[0:4]
			data = append(id, service.Handle(data[4:], context)...)
			err := func() error {
				mutex.Lock()
				defer mutex.Unlock()
				return conn.WriteMessage(websocket.BinaryMessage, data)
			}()
			if err != nil {
				service.fireErrorEvent(err, context)
				conn.Close()
			}
		}(conn, data, context)
	}
}
