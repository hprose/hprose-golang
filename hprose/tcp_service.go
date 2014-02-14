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
 * LastModified: Feb 3, 2014                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"bufio"
	"crypto/tls"
	"net"
	"net/url"
	"time"
)

type TcpService struct {
	*BaseService
}

func NewTcpService() *TcpService {
	return &TcpService{NewBaseService()}
}

func (service *TcpService) ServeTCP(conn net.Conn) {
	istream := bufio.NewReader(conn)
	ostream := conn
	go func() {
		for {
			service.Handle(istream, ostream)
			if service.IOError != nil {
				service.IOError = nil
				conn.Close()
				break
			}
		}
	}()
}

type TcpServer struct {
	*TcpService
	URL string
	*net.TCPListener
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
	var u *url.URL
	var err error
	if u, err = url.Parse(uri); err != nil {
		panic(err.Error())
	}
	var addr *net.TCPAddr
	if addr, err = net.ResolveTCPAddr(u.Scheme, u.Host); err != nil {
		panic(err.Error())
	}
	var listener *net.TCPListener
	if listener, err = net.ListenTCP(u.Scheme, addr); err != nil {
		panic(err.Error())
	}
	return &TcpServer{
		TcpService:  NewTcpService(),
		URL:         u.Scheme + "://" + listener.Addr().String(),
		TCPListener: listener,
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

func (server *TcpServer) Start() (err error) {
	for {
		var conn *net.TCPConn
		if conn, err = server.TCPListener.AcceptTCP(); err != nil {
			return err
		}
		if server.keepAlive != nil {
			if err := conn.SetKeepAlive(server.keepAlive.(bool)); err != nil {
				return err
			}
		}
		if server.keepAlivePeriod != nil {
			if err := conn.SetKeepAlivePeriod(server.keepAlivePeriod.(time.Duration)); err != nil {
				return err
			}
		}
		if server.linger != nil {
			if err := conn.SetLinger(server.linger.(int)); err != nil {
				return err
			}
		}
		if server.noDelay != nil {
			if err := conn.SetNoDelay(server.noDelay.(bool)); err != nil {
				return err
			}
		}
		if server.readBuffer != nil {
			if err := conn.SetReadBuffer(server.readBuffer.(int)); err != nil {
				return err
			}
		}
		if server.writerBuffer != nil {
			if err := conn.SetWriteBuffer(server.writerBuffer.(int)); err != nil {
				return err
			}
		}
		if server.deadline != nil {
			if err := conn.SetDeadline(server.deadline.(time.Time)); err != nil {
				return err
			}
		}
		if server.readDeadline != nil {
			if err := conn.SetReadDeadline(server.readDeadline.(time.Time)); err != nil {
				return err
			}
		}
		if server.writerDeadline != nil {
			if err := conn.SetWriteDeadline(server.writerDeadline.(time.Time)); err != nil {
				return err
			}
		}
		if server.config != nil {
			server.ServeTCP(tls.Client(conn, server.config))
		} else {
			server.ServeTCP(conn)
		}
	}
}
