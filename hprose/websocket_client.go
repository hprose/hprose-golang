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
 * hprose/websocket_client.go                             *
 *                                                        *
 * hprose websocket client for Go.                        *
 *                                                        *
 * LastModified: May 25, 2015                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"sync"

	"github.com/gorilla/websocket"
)

// WebSocketClient is hprose websocket client
type WebSocketClient struct {
	*BaseClient
	keepAlive bool
}

type receiveMessage struct {
	data []byte
	err  error
}

type webSocketTransporter struct {
	dialer                *websocket.Dialer
	conn                  *websocket.Conn
	header                *http.Header
	mutex                 sync.Mutex
	maxConcurrentRequests int
	id                    chan uint32
	sendIDs               chan uint32
	sendMsgs              map[uint32][]byte
	recvMsgs              map[uint32](chan receiveMessage)
}

// NewWebSocketClient is the constructor of WebSocketClient
func NewWebSocketClient(uri string) (client *WebSocketClient) {
	client = new(WebSocketClient)
	transporter := new(webSocketTransporter)
	transporter.dialer = new(websocket.Dialer)
	transporter.header = new(http.Header)
	transporter.maxConcurrentRequests = 10
	client.BaseClient = NewBaseClient(transporter)
	client.Client = client
	client.SetUri(uri)
	client.SetKeepAlive(true)
	return
}

func newWebSocketClient(uri string) Client {
	return NewWebSocketClient(uri)
}

// Close the client
func (client *WebSocketClient) Close() {
	trans := client.Transporter.(*webSocketTransporter)
	if trans.conn != nil {
		trans.conn.Close()
		trans.conn = nil
	}
}

// SetUri set the uri of hprose client
func (client *WebSocketClient) SetUri(uri string) {
	if u, err := url.Parse(uri); err == nil {
		if u.Scheme != "ws" && u.Scheme != "wss" {
			panic("This client desn't support " + u.Scheme + " scheme.")
		}
		if u.Scheme == "wss" {
			config := new(tls.Config)
			config.InsecureSkipVerify = true
			client.SetTLSClientConfig(config)
		}
	}
	client.BaseClient.SetUri(uri)
}

func (client *WebSocketClient) trans() *webSocketTransporter {
	return client.Transporter.(*webSocketTransporter)
}

func (client *WebSocketClient) dialer() *websocket.Dialer {
	return client.trans().dialer
}

// Header returns the http.Header in hprose client
func (client *WebSocketClient) Header() *http.Header {
	return client.trans().header
}

// TLSClientConfig returns the tls.Config in hprose client
func (client *WebSocketClient) TLSClientConfig() *tls.Config {
	return client.dialer().TLSClientConfig
}

// SetTLSClientConfig sets the tls.Config
func (client *WebSocketClient) SetTLSClientConfig(config *tls.Config) {
	client.dialer().TLSClientConfig = config
}

// KeepAlive returns the keepalive status of hprose client
func (client *WebSocketClient) KeepAlive() bool {
	return client.keepAlive
}

// SetKeepAlive sets the keepalive status of hprose client
func (client *WebSocketClient) SetKeepAlive(enable bool) {
	client.keepAlive = enable
}

// MaxConcurrentRequests returns the max concurrent requests of hprose client
func (client *WebSocketClient) MaxConcurrentRequests() int {
	return client.trans().maxConcurrentRequests
}

// SetMaxConcurrentRequests sets the max concurrent requests of hprose client
func (client *WebSocketClient) SetMaxConcurrentRequests(value int) {
	client.trans().maxConcurrentRequests = value
}

func (trans *webSocketTransporter) idGen() {
	defer func() {
		close(trans.id)
		trans.id = nil
	}()
	var i uint32
	for {
		if trans.conn == nil {
			break
		}
		trans.id <- i
		i++
	}
}

func (trans *webSocketTransporter) sendLoop() {
	defer func() {
		close(trans.sendIDs)
		trans.sendIDs = nil
		trans.sendMsgs = nil
	}()
	for {
		if trans.conn == nil {
			break
		}
		id := <-trans.sendIDs
		msg := trans.sendMsgs[id]
		delete(trans.sendMsgs, id)
		err := trans.conn.WriteMessage(websocket.BinaryMessage, msg)
		if err != nil {
			trans.conn.Close()
			trans.conn = nil
			recvMsg := trans.recvMsgs[id]
			delete(trans.recvMsgs, id)
			recvMsg <- receiveMessage{nil, err}
			close(recvMsg)
		}
	}
}

func (trans *webSocketTransporter) recvLoop() {
	var msgType int
	var data []byte
	var err error
	defer func() {
		for _, recvMsg := range trans.recvMsgs {
			recvMsg <- receiveMessage{nil, err}
			close(recvMsg)
		}
		trans.recvMsgs = nil
	}()
	for {
		if trans.conn == nil {
			break
		}
		msgType, data, err = trans.conn.ReadMessage()
		if err != nil {
			trans.conn.Close()
			trans.conn = nil
			break
		}
		if msgType == websocket.BinaryMessage {
			id := (uint32(data[0])<<24 |
				uint32(data[1])<<16 |
				uint32(data[2])<<8 |
				uint32(data[3]))
			recvMsg := trans.recvMsgs[id]
			delete(trans.recvMsgs, id)
			recvMsg <- receiveMessage{data[4:], nil}
			close(recvMsg)
		}
	}
}

func (trans *webSocketTransporter) getConn(uri string) (err error) {
	trans.mutex.Lock()
	defer trans.mutex.Unlock()
	if trans.conn == nil {
		trans.conn, _, err = trans.dialer.Dial(uri, *trans.header)
		if err != nil {
			return err
		}
		trans.id = make(chan uint32)
		trans.sendIDs = make(chan uint32, trans.maxConcurrentRequests)
		trans.sendMsgs = make(map[uint32][]byte, trans.maxConcurrentRequests)
		trans.recvMsgs = make(map[uint32](chan receiveMessage), trans.maxConcurrentRequests)
		go trans.idGen()
		go trans.sendLoop()
		go trans.recvLoop()
	}
	return nil
}

// SendAndReceive send and receive the data
func (trans *webSocketTransporter) SendAndReceive(uri string, data []byte) ([]byte, error) {
	if err := trans.getConn(uri); err != nil {
		return nil, err
	}
	id := <-trans.id
	msg := make([]byte, len(data)+4)
	msg[0] = byte((id >> 24) & 0xff)
	msg[1] = byte((id >> 16) & 0xff)
	msg[2] = byte((id >> 8) & 0xff)
	msg[3] = byte(id & 0xff)
	copy(msg[4:], data)
	recvMsg := make(chan receiveMessage)
	trans.recvMsgs[id] = recvMsg
	trans.sendMsgs[id] = msg
	trans.sendIDs <- id
	result := <-recvMsg
	return result.data, result.err
}
