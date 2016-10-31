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
 * rpc/types.go                                           *
 *                                                        *
 * reflect types for Go.                                  *
 *                                                        *
 * LastModified: Sep 16, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import (
	"net"
	"net/http"
	"reflect"

	"github.com/gorilla/websocket"
	"github.com/valyala/fasthttp"
)

var stringType = reflect.TypeOf("")
var errorType = reflect.TypeOf((*error)(nil)).Elem()
var interfaceType = reflect.TypeOf((*interface{})(nil)).Elem()
var contextType = reflect.TypeOf((*Context)(nil)).Elem()
var serviceContextType = reflect.TypeOf((*ServiceContext)(nil)).Elem()
var httpContextType = reflect.TypeOf((*HTTPContext)(nil))
var httpRequestType = reflect.TypeOf((*http.Request)(nil))
var fasthttpContextType = reflect.TypeOf((*FastHTTPContext)(nil))
var fasthttpRequestCtxType = reflect.TypeOf((*fasthttp.RequestCtx)(nil))
var socketContextType = reflect.TypeOf((*SocketContext)(nil))
var netConnType = reflect.TypeOf((*net.Conn)(nil)).Elem()
var websocketContextType = reflect.TypeOf((*WebSocketContext)(nil))
var websocketConnType = reflect.TypeOf((*websocket.Conn)(nil))
