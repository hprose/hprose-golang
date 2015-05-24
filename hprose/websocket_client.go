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
 * LastModified: May 24, 2015                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"crypto/tls"
	"errors"
	"net/http"
	"net/url"
	"sync"

	"github.com/gorilla/websocket"
)

var errClosed = errors.New("connection was closed")

// WebSocketClient is hprose websocket client
type WebSocketClient struct {
	*BaseClient
	keepAlive bool
}

// WebSocketTransporter is hprose websocket transporter
type WebSocketTransporter struct {
	dialer                *websocket.Dialer
	conn                  *websocket.Conn
	header                *http.Header
	mutex                 sync.Mutex
	maxConcurrentRequests int
	id                    chan uint32
	sendIDs               chan uint32
	sendMsgs              map[uint32][]byte
	recvMsgs              map[uint32](chan []byte)
}

// NewWebSocketClient is the constructor of WebSocketClient
func NewWebSocketClient(uri string) Client {
	client := new(WebSocketClient)
	transporter := new(WebSocketTransporter)
	transporter.dialer = new(websocket.Dialer)
	transporter.header = new(http.Header)
	client.BaseClient = NewBaseClient(transporter)
	client.Client = client
	client.SetUri(uri)
	client.SetKeepAlive(true)
	return client
}

// Close the client
func (client *WebSocketClient) Close() {
	trans := client.Transporter.(*WebSocketTransporter)
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

func (client *WebSocketClient) trans() *WebSocketTransporter {
	return client.Transporter.(*WebSocketTransporter)
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

// SendAndReceive send and receive the data
func (trans *WebSocketTransporter) SendAndReceive(uri string, data []byte) ([]byte, error) {
	err := func() (err error) {
		trans.mutex.Lock()
		defer trans.mutex.Unlock()
		if trans.conn == nil {
			trans.conn, _, err = trans.dialer.Dial(uri, *trans.header)
			if err != nil {
				return err
			}
			trans.id = make(chan uint32)
			go func() {
				defer close(trans.id)
				var i uint32
				for {
					if trans.conn == nil {
						break
					}
					trans.id <- i
					i++
				}
			}()
			trans.sendIDs = make(chan uint32, trans.maxConcurrentRequests)
			trans.sendMsgs = make(map[uint32][]byte, trans.maxConcurrentRequests)
			go func() {
				defer close(trans.sendIDs)
				for {
					if trans.conn == nil {
						break
					}
					id := <-trans.sendIDs
					msg := trans.sendMsgs[id]
					delete(trans.sendMsgs, id)
					trans.conn.WriteMessage(websocket.BinaryMessage, msg)
				}
			}()
			trans.recvMsgs = make(map[uint32](chan []byte), trans.maxConcurrentRequests)
			go func() {
				for {
					if trans.conn == nil {
						break
					}
					msgType, data, err := trans.conn.ReadMessage()
					if err != nil {
						break
					}
					if msgType == websocket.BinaryMessage {
						id := (uint32(data[0])<<24 |
							uint32(data[1])<<16 |
							uint32(data[2])<<8 |
							uint32(data[3]))
						recvMsg := trans.recvMsgs[id]
						delete(trans.recvMsgs, id)
						recvMsg <- data[4:]
					}
				}
				for _, recvMsg := range trans.recvMsgs {
					close(recvMsg)
				}
			}()
		}
		return nil
	}()
	if err != nil {
		return nil, err
	}
	id := <-trans.id
	msg := make([]byte, len(data)+4)
	msg[0] = byte((id >> 24) & 0xff)
	msg[1] = byte((id >> 16) & 0xff)
	msg[2] = byte((id >> 8) & 0xff)
	msg[3] = byte(id & 0xff)
	copy(msg[4:], data)
	recvMsg := make(chan []byte)
	trans.recvMsgs[id] = recvMsg
	trans.sendMsgs[id] = msg
	trans.sendIDs <- id
	if result, ok := <-recvMsg; ok {
		return result, nil
	}
	return nil, errClosed
}
