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
 * rpc/websocket_service.go                               *
 *                                                        *
 * hprose websocket service for Go.                       *
 *                                                        *
 * LastModified: Oct 9, 2016                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import (
	"net/http"
	"reflect"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/hprose/hprose-golang/util"
)

// WebSocketContext is the hprose websocket context
type WebSocketContext struct {
	HTTPContext
	WebSocket *websocket.Conn
}

// WebSocketService is the hprose websocket service
type WebSocketService struct {
	HTTPService
	websocket.Upgrader
	contextPool sync.Pool
}

func websocketFixArguments(args []reflect.Value, context ServiceContext) {
	i := len(args) - 1
	switch args[i].Type() {
	case websocketContextType:
		if c, ok := context.(*WebSocketContext); ok {
			args[i] = reflect.ValueOf(c)
		}
	case websocketConnType:
		if c, ok := context.(*WebSocketContext); ok {
			args[i] = reflect.ValueOf(c.WebSocket)
		}
	case httpContextType:
		if c, ok := context.(*WebSocketContext); ok {
			args[i] = reflect.ValueOf(&c.HTTPContext)
		}
	case httpRequestType:
		if c, ok := context.(*WebSocketContext); ok {
			args[i] = reflect.ValueOf(c.Request)
		}
	default:
		defaultFixArguments(args, context)
	}
}

// NewWebSocketService is the constructor of WebSocketService
func NewWebSocketService() (service *WebSocketService) {
	service = new(WebSocketService)
	service.initBaseHTTPService()
	service.contextPool = sync.Pool{
		New: func() interface{} { return new(WebSocketContext) },
	}
	service.FixArguments = websocketFixArguments
	service.CheckOrigin = func(request *http.Request) bool {
		origin := request.Header.Get("origin")
		if origin != "" && origin != "null" {
			if len(service.accessControlAllowOrigins) == 0 ||
				service.accessControlAllowOrigins[origin] {
				return true
			}
			return false
		}
		return true
	}
	return
}

func (service *WebSocketService) acquireContext() (context *WebSocketContext) {
	return service.contextPool.Get().(*WebSocketContext)
}

func (service *WebSocketService) releaseContext(context *WebSocketContext) {
	service.contextPool.Put(context)
}

// ServeHTTP is the hprose http handler method
func (service *WebSocketService) ServeHTTP(
	response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && strings.ToLower(request.Header.Get("connection")) != "upgrade" || request.Method == "POST" {
		service.HTTPService.ServeHTTP(response, request)
		return
	}
	conn, err := service.Upgrade(response, request, nil)
	if err != nil {
		context := service.HTTPService.acquireContext()
		context.initHTTPContext(service, response, request)
		resp := service.endError(err, context)
		service.HTTPService.releaseContext(context)
		response.Header().Set("Content-Length", util.Itoa(len(resp)))
		response.Write(resp)
		return
	}
	defer conn.Close()

	mutex := new(sync.Mutex)
	for {
		msgType, data, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if msgType == websocket.BinaryMessage {
			go service.handle(data, mutex, response, request, conn)
		}
	}
}

func (service *WebSocketService) handle(
	data []byte,
	mutex *sync.Mutex,
	response http.ResponseWriter,
	request *http.Request,
	conn *websocket.Conn) {
	context := service.acquireContext()
	context.initHTTPContext(service, response, request)
	context.WebSocket = conn
	id := data[0:4]
	data = service.Handle(data[4:], context)
	mutex.Lock()
	writer, err := context.WebSocket.NextWriter(websocket.BinaryMessage)
	if err == nil {
		_, err = writer.Write(id)
	}
	if err == nil {
		_, err = writer.Write(data)
	}
	if err == nil {
		err = writer.Close()
	}
	mutex.Unlock()
	if err != nil {
		fireErrorEvent(service.Event, err, context)
	}
	service.releaseContext(context)
}
