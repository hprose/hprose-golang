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
 * rpc/tcp_server.go                                      *
 *                                                        *
 * hprose tcp server for Go.                              *
 *                                                        *
 * LastModified: Nov 1, 2016                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import (
	"net"
	"net/url"
)

// TCPServer is a hprose tcp server
type TCPServer struct {
	TCPService
	starter
	uri      string
	listener *net.TCPListener
}

// NewTCPServer is the constructor for TCPServer
func NewTCPServer(uri string) (server *TCPServer) {
	if uri == "" {
		uri = "tcp://127.0.0.1:0"
	}
	server = new(TCPServer)
	server.initTCPService()
	server.starter.server = server
	server.uri = uri
	return
}

// URI return the real address of this server
func (server *TCPServer) URI() string {
	if server.listener == nil {
		panic(ErrServerIsNotStarted)
	}
	u, err := url.Parse(server.uri)
	if err != nil {
		panic(err)
	}
	return u.Scheme + "://" + server.listener.Addr().String()
}

// Handle the hprose tcp server
func (server *TCPServer) Handle() (err error) {
	if server.listener != nil {
		return ErrServerIsAlreadyStarted
	}
	u, err := url.Parse(server.uri)
	if err != nil {
		return err
	}
	addr, err := net.ResolveTCPAddr(u.Scheme, u.Host)
	if err != nil {
		return err
	}
	if server.listener, err = net.ListenTCP(u.Scheme, addr); err != nil {
		return err
	}
	go server.ServeTCP(server.listener)
	return nil
}

// Close the hprose tcp server
func (server *TCPServer) Close() {
	if server.listener != nil {
		listener := server.listener
		server.listener = nil
		listener.Close()
	}
}
