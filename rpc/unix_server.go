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
 * rpc/unix_server.go                                     *
 *                                                        *
 * hprose unix server for Go.                             *
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

// UnixServer is a hprose unix server
type UnixServer struct {
	UnixService
	starter
	uri      string
	listener *net.UnixListener
}

// NewUnixServer is the constructor for UnixServer
func NewUnixServer(uri string) (server *UnixServer) {
	if uri == "" {
		uri = "unix:/tmp/hprose.sock"
	}
	server = new(UnixServer)
	server.initSocketService()
	server.starter.server = server
	server.uri = uri
	return
}

// URI return the real address of this server
func (server *UnixServer) URI() string {
	if server.listener == nil {
		panic(ErrServerIsNotStarted)
	}
	return "unix:" + server.listener.Addr().String()
}

// Handle the hprose unix server
func (server *UnixServer) Handle() (err error) {
	if server.listener != nil {
		return ErrServerIsAlreadyStarted
	}
	u, err := url.Parse(server.uri)
	if err != nil {
		return err
	}
	addr, err := net.ResolveUnixAddr(u.Scheme, u.Path)
	if err != nil {
		return err
	}
	if server.listener, err = net.ListenUnix(u.Scheme, addr); err != nil {
		return err
	}
	go server.ServeUnix(server.listener)
	return nil
}

// Close the hprose unix server
func (server *UnixServer) Close() {
	if server.listener != nil {
		listener := server.listener
		server.listener = nil
		listener.Close()
	}
}
