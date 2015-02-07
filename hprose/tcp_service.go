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
 * LastModified: Feb 7, 2015                              *
 * Authors: Ma Bingyao <andot@hprose.com>                 *
 *          Ore_Ash <nanohugh@gmail.com>                  *
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
	*StreamService
	keepAlive       interface{}
	keepAlivePeriod interface{}
	linger          interface{}
	noDelay         interface{}
	config          *tls.Config
}

type TcpContext StreamContext

type tcpArgsFixer struct{}

func (tcpArgsFixer) FixArgs(args []reflect.Value, lastParamType reflect.Type, context Context) []reflect.Value {
	if c, ok := context.(*TcpContext); ok {
		if lastParamType.String() == "*hprose.TcpContext" {
			return append(args, reflect.ValueOf(c))
		} else if lastParamType.String() == "*hprose.StreamContext" {
			return append(args, reflect.ValueOf((*StreamContext)(c)))
		} else if lastParamType.String() == "net.Conn" {
			return append(args, reflect.ValueOf(c.Conn))
		}
	}
	return fixArgs(args, lastParamType, context)
}

func NewTcpService() *TcpService {
	service := &TcpService{StreamService: newStreamService()}
	service.argsfixer = tcpArgsFixer{}
	return service
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
				data, err = receiveDataOverStream(conn)
			}
			if err == nil {
				data = service.Handle(data, &TcpContext{BaseContext: NewBaseContext(), Conn: conn})
				if service.writeTimeout != nil {
					err = conn.SetWriteDeadline(time.Now().Add(service.writeTimeout.(time.Duration)))
				}
				if err == nil {
					err = sendDataOverStream(conn, data)
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
