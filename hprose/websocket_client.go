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
 * LastModified: Mar 31, 2016                             *
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
}

type sendMessage struct {
	id   uint32
	data []byte
}

type recvMessage struct {
	data []byte
	err  error
}

type recvCommand struct {
	id   uint32
	recv chan recvMessage
	data []byte
	err  error
}

type webSocketTransporter struct {
	dialer                *websocket.Dialer
	conn                  *websocket.Conn
	header                *http.Header
	mutex                 sync.RWMutex
	maxConcurrentRequests int
	id                    chan uint32
	sendChan              chan sendMessage
	recvChan              chan recvCommand
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

// SetKeepAlive do nothing on WebSocketClient, it is always keepAlive
func (client *WebSocketClient) SetKeepAlive(enable bool) {
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
		i++
		trans.id <- i
		if i == 0xFFFFFFFF {
			i = 0
		}
	}
}

func (trans *webSocketTransporter) sendLoop() {
	defer func() {
		close(trans.sendChan)
		trans.sendChan = nil
	}()
	for {
		trans.mutex.RLock()
		if trans.conn == nil {
			trans.mutex.RUnlock()
			break
		}
		trans.mutex.RUnlock()
		send := <-trans.sendChan
		err := trans.conn.WriteMessage(websocket.BinaryMessage, send.data)
		if err != nil {
			trans.mutex.Lock()
			trans.conn.Close()
			trans.conn = nil
			trans.mutex.Unlock()
			trans.recvChan <- recvCommand{send.id, nil, nil, err}
		}
	}
}

func (trans *webSocketTransporter) resultLoop() {
	results := make(map[uint32](chan recvMessage))
	for r := range trans.recvChan {
		if r.recv != nil {
			results[r.id] = r.recv
		} else if r.data != nil {
			recv := results[r.id]
			delete(results, r.id)
			recv <- recvMessage{r.data, nil}
			close(recv)
		} else if r.err != nil {
			if r.id != 0 {
				recv := results[r.id]
				delete(results, r.id)
				recv <- recvMessage{nil, r.err}
				close(recv)
			} else {
				for _, recv := range results {
					recv <- recvMessage{nil, r.err}
					close(recv)
				}
				break
			}
		}
	}
	results = nil
	close(trans.recvChan)
	trans.recvChan = nil
}

func (trans *webSocketTransporter) recvLoop() {
	var msgType int
	var data []byte
	var err error
	defer func() {
		trans.recvChan <- recvCommand{0, nil, nil, err}
	}()
	for {
		trans.mutex.RLock()
		if trans.conn == nil {
			trans.mutex.RUnlock()
			break
		}
		msgType, data, err = trans.conn.ReadMessage()
		if err != nil {
			trans.mutex.Lock()
			trans.conn.Close()
			trans.conn = nil
			trans.mutex.Unlock()
			trans.mutex.RUnlock()
			break
		}
		trans.mutex.RUnlock()
		if msgType == websocket.BinaryMessage {
			id := (uint32(data[0])<<24 |
				uint32(data[1])<<16 |
				uint32(data[2])<<8 |
				uint32(data[3]))
			trans.recvChan <- recvCommand{id, nil, data[4:], nil}
		}
	}
}

func (trans *webSocketTransporter) getConn(uri string) (err error) {
	trans.mutex.RLock()
	defer trans.mutex.RUnlock()
	if trans.conn == nil {
		trans.mutex.Lock()
		trans.conn, _, err = trans.dialer.Dial(uri, *trans.header)
		trans.mutex.Unlock()
		if err != nil {
			return err
		}
		trans.id = make(chan uint32)
		trans.sendChan = make(chan sendMessage, trans.maxConcurrentRequests)
		trans.recvChan = make(chan recvCommand, trans.maxConcurrentRequests)
		go trans.idGen()
		go trans.resultLoop()
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
	buf := make([]byte, len(data)+4)
	buf[0] = byte((id >> 24) & 0xff)
	buf[1] = byte((id >> 16) & 0xff)
	buf[2] = byte((id >> 8) & 0xff)
	buf[3] = byte(id & 0xff)
	copy(buf[4:], data)
	recv := make(chan recvMessage)
	trans.recvChan <- recvCommand{id, recv, nil, nil}
	trans.sendChan <- sendMessage{id, buf}
	result := <-recv
	return result.data, result.err
}
