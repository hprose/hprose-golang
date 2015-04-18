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
 * LastModified: Apr 18, 2015                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"reflect"

	"golang.org/x/net/websocket"
)

type WebSocketContext struct {
	*BaseContext
	WebSocket *websocket.Conn
}

type WebSocketService struct {
	*BaseService
	Handler websocket.Handler
	Server  websocket.Server
}

type wsArgsFixer struct{}

func (wsArgsFixer) FixArgs(args []reflect.Value, lastParamType reflect.Type, context Context) []reflect.Value {
	if c, ok := context.(*WebSocketContext); ok {
		if lastParamType.String() == "*hprose.WebSocketContext" {
			return append(args, reflect.ValueOf(c))
		} else if lastParamType.String() == "*websocket.Conn" {
			return append(args, reflect.ValueOf(c.WebSocket))
		}
	}
	return fixArgs(args, lastParamType, context)
}

func NewWebSocketService() *WebSocketService {
	service := &WebSocketService{
		BaseService: NewBaseService(),
	}
	service.argsfixer = wsArgsFixer{}
	service.Handler = websocket.Handler(service.ServeWebSocket)
	service.Server = websocket.Server{websocket.Config{}, nil, service.ServeWebSocket}
	return service
}

func (service *WebSocketService) ServeWebSocket(ws *websocket.Conn) {
	defer ws.Close()
	for {
		context := &WebSocketContext{BaseContext: NewBaseContext(), WebSocket: ws}
		var data []byte
		if e := websocket.Message.Receive(ws, &data); e != nil {
			service.fireErrorEvent(e, context)
			break
		}
		go func(ws *websocket.Conn, data []byte, context *WebSocketContext) {
			id := data[0:4]
			data = append(id, service.Handle(data[4:], context)...)
			if e := websocket.Message.Send(ws, data); e != nil {
				service.fireErrorEvent(e, context)
				ws.Close()
			}
		}(ws, data, context)
	}
}
