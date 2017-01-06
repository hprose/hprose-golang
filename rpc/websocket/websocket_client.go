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
 * rpc/websocket/websocket_client.go                      *
 *                                                        *
 * hprose websocket client for Go.                        *
 *                                                        *
 * LastModified: Jan 7, 2017                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package websocket

import (
	"crypto/tls"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/hprose/hprose-golang/util"
)

var websocketSchemes = []string{"ws", "wss"}

type reqeust struct {
	id   uint32
	data []byte
}

type socketResponse struct {
	data []byte
	err  error
}

// WebSocketClient is hprose websocket client
type WebSocketClient struct {
	rpc.BaseClient
	http.Header
	dialer    websocket.Dialer
	conn      *websocket.Conn
	nextid    uint32
	requests  chan reqeust
	responses map[uint32]chan socketResponse
	closed    bool
	limiter   rpc.Limiter
}

// NewWebSocketClient is the constructor of WebSocketClient
func NewWebSocketClient(uri ...string) (client *WebSocketClient) {
	client = new(WebSocketClient)
	client.InitBaseClient()
	client.limiter.InitLimiter()
	client.closed = false
	client.SetURIList(uri)
	client.SendAndReceive = client.sendAndReceive
	return
}

func newWebSocketClient(uri ...string) rpc.Client {
	return NewWebSocketClient(uri...)
}

// SetURIList sets a list of server addresses
func (client *WebSocketClient) SetURIList(uriList []string) {
	if rpc.CheckAddresses(uriList, websocketSchemes) == "wss" {
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	client.BaseClient.SetURIList(uriList)
}

func (client *WebSocketClient) close(err error) {
	client.limiter.L.Lock()
	if err != nil && client.responses != nil {
		for _, response := range client.responses {
			response <- socketResponse{nil, err}
		}
	}
	client.responses = nil
	if client.conn != nil {
		client.conn.Close()
		client.conn = nil
	}
	client.limiter.Reset()
	client.limiter.L.Unlock()
}

// Close the client
func (client *WebSocketClient) Close() {
	client.closed = true
	client.close(rpc.ErrClientIsAlreadyClosed)
}

// TLSClientConfig returns the tls.Config in hprose client
func (client *WebSocketClient) TLSClientConfig() *tls.Config {
	return client.dialer.TLSClientConfig
}

// SetTLSClientConfig sets the tls.Config
func (client *WebSocketClient) SetTLSClientConfig(config *tls.Config) {
	client.dialer.TLSClientConfig = config
}

// MaxConcurrentRequests returns max concurrent request count
func (client *WebSocketClient) MaxConcurrentRequests() int {
	return client.limiter.MaxConcurrentRequests
}

// SetMaxConcurrentRequests sets max concurrent request count
func (client *WebSocketClient) SetMaxConcurrentRequests(value int) {
	client.limiter.MaxConcurrentRequests = value
}

func (client *WebSocketClient) sendLoop() {
	conn := client.conn
	for request := range client.requests {
		err := conn.WriteMessage(websocket.BinaryMessage, request.data)
		if err != nil {
			client.close(err)
			break
		}
	}
	client.requests = nil
}

func (client *WebSocketClient) recvLoop() {
	conn := client.conn
	for {
		msgType, data, err := conn.ReadMessage()
		if err != nil {
			client.close(err)
			break
		}
		if msgType == websocket.BinaryMessage {
			id := util.ToUint32(data)
			client.limiter.L.Lock()
			response := client.responses[id]
			if response != nil {
				response <- socketResponse{data[4:], nil}
				delete(client.responses, id)
			}
			client.limiter.Unlimit()
			client.limiter.L.Unlock()
		}
	}
	close(client.requests)
}

func (client *WebSocketClient) getConn(uri string) (err error) {
	if client.conn == nil {
		client.conn, _, err = client.dialer.Dial(uri, client.Header)
		if err != nil {
			return err
		}
		count := client.limiter.MaxConcurrentRequests
		client.requests = make(chan reqeust, count)
		client.responses = make(map[uint32]chan socketResponse, count)
		go client.sendLoop()
		go client.recvLoop()
	}
	return nil
}

func (client *WebSocketClient) sendAndReceive(
	data []byte, context *rpc.ClientContext) ([]byte, error) {
	id := atomic.AddUint32(&client.nextid, 1)
	buf := make([]byte, len(data)+4)
	util.FromUint32(buf, id)
	copy(buf[4:], data)
	response := make(chan socketResponse)
	client.limiter.L.Lock()
	client.limiter.Limit()
	if client.closed {
		client.limiter.Unlimit()
		client.limiter.L.Unlock()
		return nil, rpc.ErrClientIsAlreadyClosed
	}
	if err := client.getConn(client.URI()); err != nil {
		client.limiter.Unlimit()
		client.limiter.L.Unlock()
		return nil, err
	}
	client.responses[id] = response
	client.limiter.L.Unlock()
	client.requests <- reqeust{id, buf}
	select {
	case resp := <-response:
		return resp.data, resp.err
	case <-time.After(context.Timeout):
		client.limiter.L.Lock()
		delete(client.responses, id)
		client.limiter.Unlimit()
		client.limiter.L.Unlock()
		return nil, rpc.ErrTimeout
	}
}

func init() {
	rpc.RegisterClientFactory("ws", newWebSocketClient)
	rpc.RegisterClientFactory("wss", newWebSocketClient)
}
