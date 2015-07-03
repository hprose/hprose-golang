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
 * LastModified: Jul 3, 2015                              *
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
	"os"
	"os/signal"
	"reflect"
	"runtime/debug"
	"time"
)

// TcpService is the hprose tcp service
type TcpService struct {
	*StreamService
	keepAlive       interface{}
	keepAlivePeriod interface{}
	linger          interface{}
	noDelay         interface{}
	config          *tls.Config
}

// TcpContext is the hprose tcp context
type TcpContext StreamContext

type tcpArgsFixer struct{}

func (tcpArgsFixer) FixArgs(args []reflect.Value, lastParamType reflect.Type, context Context) []reflect.Value {
	if c, ok := context.(*StreamContext); ok {
		if lastParamType.String() == "*hprose.StreamContext" {
			return append(args, reflect.ValueOf(c))
		} else if lastParamType.String() == "*hprose.TcpContext" {
			return append(args, reflect.ValueOf((*TcpContext)(c)))
		} else if lastParamType.String() == "net.Conn" {
			return append(args, reflect.ValueOf(c.Conn))
		}
	}
	return fixArgs(args, lastParamType, context)
}

// NewTcpService is the constructor of TcpService
func NewTcpService() (service *TcpService) {
	service = new(TcpService)
	service.StreamService = newStreamService()
	service.argsfixer = tcpArgsFixer{}
	return
}

// SetKeepAlive sets whether the operating system should send keepalive messages on the connection.
func (service *TcpService) SetKeepAlive(keepalive bool) {
	service.keepAlive = keepalive
}

// SetKeepAlivePeriod sets period between keep alives.
func (service *TcpService) SetKeepAlivePeriod(d time.Duration) {
	service.keepAlivePeriod = d
}

// SetLinger sets the behavior of Close on a connection which still has data waiting to be sent or to be acknowledged.
//
// If sec < 0 (the default), the operating system finishes sending the data in the background.
//
// If sec == 0, the operating system discards any unsent or unacknowledged data.
//
// If sec > 0, the data is sent in the background as with sec < 0. On some operating systems after sec seconds have elapsed any remaining unsent data may be discarded.
func (service *TcpService) SetLinger(sec int) {
	service.linger = sec
}

// SetNoDelay controls whether the operating system should delay packet transmission in hopes of sending fewer packets (Nagle's algorithm). The default is true (no delay), meaning that data is sent as soon as possible after a Write.
func (service *TcpService) SetNoDelay(noDelay bool) {
	service.noDelay = noDelay
}

// SetTLSConfig sets the Config structure used to configure TLS service
func (service *TcpService) SetTLSConfig(config *tls.Config) {
	service.config = config
}

// ServeTCP ...
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
	if service.config != nil {
		tlsConn := tls.Server(conn, service.config)
		tlsConn.Handshake()
		return service.Serve(tlsConn)
	}
	return service.Serve(conn)
}

// TcpServer is a hprose tcp server
type TcpServer struct {
	*TcpService
	URL      string
	listener *net.TCPListener
	signal   chan os.Signal
}

// NewTcpServer is a constructor for TcpServer
func NewTcpServer(uri string) (server *TcpServer) {
	if uri == "" {
		uri = "tcp://127.0.0.1:0"
	}
	server = new(TcpServer)
	server.TcpService = NewTcpService()
	server.URL = uri
	return
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

// Handle the hprose tcp server
func (server *TcpServer) Handle() (err error) {
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
		go server.start()
	}
	return nil
}

// Start the hprose tcp server
func (server *TcpServer) Start() (err error) {
	if server.listener == nil {
		server.Handle()
		server.signal = make(chan os.Signal, 1)
		signal.Notify(server.signal, os.Interrupt, os.Kill)
		<-server.signal
		server.Stop()
	}
	return nil
}

// Stop the hprose tcp server
func (server *TcpServer) Stop() {
	if server.signal != nil {
		signal.Stop(server.signal)
		server.signal = nil
	}
	if server.listener != nil {
		listener := server.listener
		server.listener = nil
		listener.Close()
	}
}
