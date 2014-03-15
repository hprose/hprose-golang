/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.net/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * hprose/tcp_service.go                                  *
 *                                                        *
 * hprose tcp service for Go.                             *
 *                                                        *
 * LastModified: Mar 15, 2014                             *
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
	"time"
)

type TcpService struct {
	*BaseService
}

type tcpArgsFixed struct{}

func (tcpArgsFixed) FixArgs(args []reflect.Value, lastParamType reflect.Type, context interface{}) []reflect.Value {
	if conn, ok := context.(net.Conn); ok && lastParamType.String() == "net.Conn" {
		args = append(args, reflect.ValueOf(conn))
	}
	return args
}

func NewTcpService() *TcpService {
	service := &TcpService{NewBaseService()}
	service.ArgsFixer = tcpArgsFixed{}
	return service
}

func (service *TcpService) ServeTCP(conn net.Conn) {
	go func() {
		for {
			data, err := receiveDataOverTcp(conn)
			if err == nil {
				err = sendDataOverTcp(conn, service.Handle(data, conn))
			}
			if err != nil {
				conn.Close()
				break
			}
		}
	}()
}

type TcpServer struct {
	*TcpService
	URL             string
	ThreadCount     int
	listener        *net.TCPListener
	deadline        interface{}
	keepAlive       interface{}
	keepAlivePeriod interface{}
	linger          interface{}
	noDelay         interface{}
	readBuffer      interface{}
	readDeadline    interface{}
	writerBuffer    interface{}
	writerDeadline  interface{}
	config          *tls.Config
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

func (server *TcpServer) SetDeadline(t time.Time) {
	server.deadline = t
}

func (server *TcpServer) SetKeepAlive(keepalive bool) {
	server.keepAlive = keepalive
}

func (server *TcpServer) SetKeepAlivePeriod(d time.Duration) {
	server.keepAlivePeriod = d
}

func (server *TcpServer) SetLinger(sec int) {
	server.linger = sec
}

func (server *TcpServer) SetNoDelay(noDelay bool) {
	server.noDelay = noDelay
}

func (server *TcpServer) SetReadBuffer(bytes int) {
	server.readBuffer = bytes
}

func (server *TcpServer) SetReadDeadline(t time.Time) {
	server.readDeadline = t
}

func (server *TcpServer) SetWriteBuffer(bytes int) {
	server.writerBuffer = bytes
}

func (server *TcpServer) SetWriteDeadline(t time.Time) {
	server.writerDeadline = t
}

func (server *TcpServer) SetTLSConfig(config *tls.Config) {
	server.config = config
}

func (server *TcpServer) handle() (err error) {
	defer func() {
		if e := recover(); e != nil && err == nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	if server.listener == nil {
		return nil
	}
	var conn *net.TCPConn
	if conn, err = server.listener.AcceptTCP(); err != nil {
		return err
	}
	if server.keepAlive != nil {
		if err = conn.SetKeepAlive(server.keepAlive.(bool)); err != nil {
			return err
		}
	}
	if server.keepAlivePeriod != nil {
		if kap, ok := (net.Conn(conn)).(iKeepAlivePeriod); ok {
			if err = kap.SetKeepAlivePeriod(server.keepAlivePeriod.(time.Duration)); err != nil {
				return err
			}
		}
	}
	if server.linger != nil {
		if err = conn.SetLinger(server.linger.(int)); err != nil {
			return err
		}
	}
	if server.noDelay != nil {
		if err = conn.SetNoDelay(server.noDelay.(bool)); err != nil {
			return err
		}
	}
	if server.readBuffer != nil {
		if err = conn.SetReadBuffer(server.readBuffer.(int)); err != nil {
			return err
		}
	}
	if server.writerBuffer != nil {
		if err = conn.SetWriteBuffer(server.writerBuffer.(int)); err != nil {
			return err
		}
	}
	if server.deadline != nil {
		if err = conn.SetDeadline(server.deadline.(time.Time)); err != nil {
			return err
		}
	}
	if server.readDeadline != nil {
		if err = conn.SetReadDeadline(server.readDeadline.(time.Time)); err != nil {
			return err
		}
	}
	if server.writerDeadline != nil {
		if err = conn.SetWriteDeadline(server.writerDeadline.(time.Time)); err != nil {
			return err
		}
	}
	if server.config != nil {
		server.ServeTCP(tls.Client(conn, server.config))
	} else {
		server.ServeTCP(conn)
	}
	return nil
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
