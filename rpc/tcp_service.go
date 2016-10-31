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
 * rpc/tcp_service.go                                     *
 *                                                        *
 * hprose tcp service for Go.                             *
 *                                                        *
 * LastModified: Oct 5, 2016                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import (
	"crypto/tls"
	"net"
	"time"
)

// TCPService is the hprose tcp service
type TCPService struct {
	SocketService
	Linger          int
	NoDelay         bool
	KeepAlive       bool
	KeepAlivePeriod time.Duration
}

// NewTCPService is the constructor of TCPService
func NewTCPService() (service *TCPService) {
	service = new(TCPService)
	service.initTCPService()
	return service
}

func (service *TCPService) initTCPService() {
	service.initSocketService()
	service.Linger = -1
	service.NoDelay = true
	service.KeepAlive = true
	service.KeepAlivePeriod = 0
}

// ServeTCPConn runs on a single tcp connection. ServeTCPConn blocks, serving
// the connection until the client hangs up. The caller typically invokes
// ServeTCPConn in a go statement.
func (service *TCPService) ServeTCPConn(conn *net.TCPConn) {
	conn.SetLinger(service.Linger)
	conn.SetNoDelay(service.NoDelay)
	conn.SetKeepAlive(service.KeepAlive)
	if service.KeepAlivePeriod > 0 {
		conn.SetKeepAlivePeriod(service.KeepAlivePeriod)
	}
	if service.TLSConfig != nil {
		tlsConn := tls.Server(conn, service.TLSConfig)
		tlsConn.Handshake()
		service.serveConn(tlsConn)
	} else {
		service.serveConn(conn)
	}
}

// ServeConn runs on a single net connection. ServeConn blocks, serving the
// connection until the client hangs up. The caller typically invokes
// ServeConn in a go statement.
func (service *TCPService) ServeConn(conn net.Conn) {
	service.ServeTCPConn(conn.(*net.TCPConn))
}

// ServeTCP runs on the TCPListener. ServeTCP blocks, serving the listener
// until the server is stop. The caller typically invokes ServeTCP in a go
// statement.
func (service *TCPService) ServeTCP(listener *net.TCPListener) {
	var tempDelay time.Duration // how long to sleep on accept failure
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			tempDelay = nextTempDelay(err, service.Event, tempDelay)
			if tempDelay > 0 {
				continue
			}
			return
		}
		tempDelay = 0
		go service.ServeTCPConn(conn)
	}
}

// Serve runs on the Listener. Serve blocks, serving the listener
// until the server is stop. The caller typically invokes Serve in a go
// statement.
func (service *TCPService) Serve(listener net.Listener) {
	service.ServeTCP(listener.(*net.TCPListener))
}
