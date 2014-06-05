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
 * hprose/tcp_service.go                                  *
 *                                                        *
 * hprose tcp service for Go.                             *
 *                                                        *
 * LastModified: Apr 14, 2014                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/url"
	"reflect"
	"runtime"
	"runtime/debug"
	"time"
)

type TcpService struct {
	*BaseService
	timeout         interface{}
	keepAlive       interface{}
	keepAlivePeriod interface{}
	linger          interface{}
	noDelay         interface{}
	readTimeout     interface{}
	readBuffer      interface{}
	writeTimeout    interface{}
	writeBuffer     interface{}
	config          *tls.Config
}

type tcpArgsFixer struct{}

func (tcpArgsFixer) FixArgs(args []reflect.Value, lastParamType reflect.Type, context interface{}) []reflect.Value {
	if conn, ok := context.(net.Conn); ok && lastParamType.String() == "net.Conn" {
		return append(args, reflect.ValueOf(conn))
	}
	return fixArgs(args, lastParamType, context)
}

func NewTcpService() *TcpService {
	service := &TcpService{BaseService: NewBaseService()}
	service.argsfixer = tcpArgsFixer{}
	return service
}

func (service *TcpService) SetTimeout(d time.Duration) {
	service.timeout = d
}

func (service *TcpService) SetKeepAlive(keepalive bool) {
	service.keepAlive = keepalive
}

func (service *TcpService) SetKeepAlivePeriod(d time.Duration) {
	service.keepAlivePeriod = d
}

func (service *TcpService) SetLinger(sec int) {
	service.linger = sec
}

func (service *TcpService) SetNoDelay(noDelay bool) {
	service.noDelay = noDelay
}

func (service *TcpService) SetReadTimeout(d time.Duration) {
	service.readTimeout = d
}

func (service *TcpService) SetReadBuffer(bytes int) {
	service.readBuffer = bytes
}

func (service *TcpService) SetWriteTimeout(d time.Duration) {
	service.writeTimeout = d
}

func (service *TcpService) SetWriteBuffer(bytes int) {
	service.writeBuffer = bytes
}

func (service *TcpService) SetTLSConfig(config *tls.Config) {
	service.config = config
}

func (service *TcpService) ServeTCP(conn *net.TCPConn) (err error) {
	if service.keepAlive != nil {
		if err = conn.SetKeepAlive(service.keepAlive.(bool)); err != nil {
			return err
		}
	}
	if service.keepAlivePeriod != nil {
		if kap, ok := (net.Conn(conn)).(iKeepAlivePeriod); ok {
			if err = kap.SetKeepAlivePeriod(service.keepAlivePeriod.(time.Duration)); err != nil {
				return err
			}
		}
	}
	if service.linger != nil {
		if err = conn.SetLinger(service.linger.(int)); err != nil {
			return err
		}
	}
	if service.noDelay != nil {
		if err = conn.SetNoDelay(service.noDelay.(bool)); err != nil {
			return err
		}
	}
	if service.readBuffer != nil {
		if err = conn.SetReadBuffer(service.readBuffer.(int)); err != nil {
			return err
		}
	}
	if service.writeBuffer != nil {
		if err = conn.SetWriteBuffer(service.writeBuffer.(int)); err != nil {
			return err
		}
	}
	if service.timeout != nil {
		if err = conn.SetDeadline(time.Now().Add(service.timeout.(time.Duration))); err != nil {
			return err
		}
	}
	go func(conn net.Conn) {
		if service.config != nil {
			tlsConn := tls.Server(conn, service.config)
			tlsConn.Handshake()
			conn = tlsConn
		}
		var data []byte
		var err error
		for {
			if service.readTimeout != nil {
				err = conn.SetReadDeadline(time.Now().Add(service.readTimeout.(time.Duration)))
			}
			if err == nil {
				data, err = receiveDataOverTcp(conn)
			}
			if err == nil {
				data = service.Handle(data, conn)
				if service.writeTimeout != nil {
					err = conn.SetWriteDeadline(time.Now().Add(service.writeTimeout.(time.Duration)))
				}
				if err == nil {
					err = sendDataOverTcp(conn, data)
				}
			}
			if err != nil {
				conn.Close()
				break
			}
		}
	}(conn)
	return nil
}

type TcpServer struct {
	*TcpService
	URL         string
	ThreadCount int
	listener    *net.TCPListener
}

func NewTcpServer(uri string) *TcpServer {
	if uri == "" {
		uri = "tcp://127.0.0.1:0"
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	return &TcpServer{
		TcpService:  NewTcpService(),
		URL:         uri,
		ThreadCount: runtime.NumCPU(),
		listener:    nil,
	}
}

func (server *TcpServer) handle() (err error) {
	defer func() {
		if e := recover(); e != nil && err == nil {
			if server.DebugEnabled {
				err = fmt.Errorf("%v\r\n%s", e, debug.Stack())
			} else {
				err = fmt.Errorf("%v", e)
			}
		}
	}()
	if server.listener == nil {
		return nil
	}
	var conn *net.TCPConn
	if conn, err = server.listener.AcceptTCP(); err != nil {
		return err
	}
	return server.ServeTCP(conn)
}

func (server *TcpServer) start() {
	for {
		if server.listener != nil {
			if err := server.handle(); err != nil {
				server.fireErrorEvent(err, nil)
			}
		} else {
			break
		}
	}
}

func (server *TcpServer) Start() (err error) {
	if server.listener == nil {
		var u *url.URL
		if u, err = url.Parse(server.URL); err != nil {
			return err
		}
		var addr *net.TCPAddr
		if addr, err = net.ResolveTCPAddr(u.Scheme, u.Host); err != nil {
			return err
		}
		if server.listener, err = net.ListenTCP(u.Scheme, addr); err != nil {
			return err
		}
		server.URL = u.Scheme + "://" + server.listener.Addr().String()
		for i := 0; i < server.ThreadCount; i++ {
			go server.start()
		}
	}
	return nil
}

func (server *TcpServer) Stop() {
	if server.listener != nil {
		listener := server.listener
		server.listener = nil
		listener.Close()
	}
}
