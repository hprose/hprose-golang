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
 * hprose/unix_service.go                                 *
 *                                                        *
 * hprose unix service for Go.                            *
 *                                                        *
 * LastModified: Jul 3, 2015                              *
 * Authors: Ma Bingyao <andot@hprose.com>                 *
 *          Ore_Ash <nanohugh@gmail.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"reflect"
	"runtime/debug"
)

// UnixService is the hprose unix service
type UnixService StreamService

// UnixContext is the hprose unix context
type UnixContext StreamContext

type unixArgsFixer struct{}

func (unixArgsFixer) FixArgs(args []reflect.Value, lastParamType reflect.Type, context Context) []reflect.Value {
	if c, ok := context.(*StreamContext); ok {
		if lastParamType.String() == "*hprose.StreamContext" {
			return append(args, reflect.ValueOf(c))
		} else if lastParamType.String() == "*hprose.UnixContext" {
			return append(args, reflect.ValueOf((*UnixContext)(c)))
		} else if lastParamType.String() == "net.Conn" {
			return append(args, reflect.ValueOf(c.Conn))
		}
	}
	return fixArgs(args, lastParamType, context)
}

// NewUnixService is the constructor of UnixService
func NewUnixService() *UnixService {
	service := (*UnixService)(newStreamService())
	service.argsfixer = unixArgsFixer{}
	return service
}

// ServeUnix ...
func (service *UnixService) ServeUnix(conn *net.UnixConn) (err error) {
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
	return ((*StreamService)(service)).Serve(conn)
}

// UnixServer is a hprose unix server
type UnixServer struct {
	*UnixService
	URL      string
	listener *net.UnixListener
	signal   chan os.Signal
}

// NewUnixServer is a constructor for UnixServer
func NewUnixServer(uri string) (server *UnixServer) {
	if uri == "" {
		uri = "unix:/tmp/hprose.sock"
	}
	server = new(UnixServer)
	server.UnixService = NewUnixService()
	server.URL = uri
	return
}

func (server *UnixServer) handle() (err error) {
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
	var conn *net.UnixConn
	if conn, err = server.listener.AcceptUnix(); err != nil {
		return err
	}
	return server.ServeUnix(conn)
}

func (server *UnixServer) start() {
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

// Handle the hprose unix server
func (server *UnixServer) Handle() (err error) {
	if server.listener == nil {
		scheme, path := parseUnixUri(server.URL)
		var addr *net.UnixAddr
		if addr, err = net.ResolveUnixAddr(scheme, path); err != nil {
			return err
		}
		if server.listener, err = net.ListenUnix(scheme, addr); err != nil {
			return err
		}
		server.URL = scheme + ":" + server.listener.Addr().String()
		go server.start()
	}
	return nil
}

// Start the hprose unix server
func (server *UnixServer) Start() (err error) {
	if server.listener == nil {
		server.Handle()
		server.signal = make(chan os.Signal, 1)
		signal.Notify(server.signal, os.Interrupt, os.Kill)
		<-server.signal
		server.Stop()
	}
	return nil
}

// Stop the hprose unix server
func (server *UnixServer) Stop() {
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
