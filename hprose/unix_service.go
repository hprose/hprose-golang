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
 * LastModified: Feb 7, 2015                              *
 * Authors: Ma Bingyao <andot@hprose.com>                 *
 *          Ore_Ash <nanohugh@gmail.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"fmt"
	"net"
	"reflect"
	"runtime"
	"runtime/debug"
	"time"
)

type UnixService StreamService

type UnixContext StreamContext

type unixArgsFixer struct{}

func (unixArgsFixer) FixArgs(args []reflect.Value, lastParamType reflect.Type, context Context) []reflect.Value {
	if c, ok := context.(*UnixContext); ok {
		if lastParamType.String() == "*hprose.UnixContext" {
			return append(args, reflect.ValueOf(c))
		} else if lastParamType.String() == "*hprose.StreamContext" {
			return append(args, reflect.ValueOf((*StreamContext)(c)))
		} else if lastParamType.String() == "net.Conn" {
			return append(args, reflect.ValueOf(c.Conn))
		}
	}
	return fixArgs(args, lastParamType, context)
}

func NewUnixService() *UnixService {
	service := (*UnixService)(newStreamService())
	service.argsfixer = unixArgsFixer{}
	return service
}

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
	if service.timeout != nil {
		if err = conn.SetDeadline(time.Now().Add(service.timeout.(time.Duration))); err != nil {
			return err
		}
	}
	go func(conn net.Conn) {
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
				data = service.Handle(data, &UnixContext{BaseContext: NewBaseContext(), Conn: conn})
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

type UnixServer struct {
	*UnixService
	URL         string
	ThreadCount int
	listener    *net.UnixListener
}

func NewUnixServer(uri string) *UnixServer {
	if uri == "" {
		uri = "unix:/tmp/hprose.sock"
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	return &UnixServer{
		UnixService: NewUnixService(),
		URL:         uri,
		ThreadCount: runtime.NumCPU(),
		listener:    nil,
	}
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

func (server *UnixServer) Start() (err error) {
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
		for i := 0; i < server.ThreadCount; i++ {
			go server.start()
		}
	}
	return nil
}

func (server *UnixServer) Stop() {
	if server.listener != nil {
		listener := server.listener
		server.listener = nil
		listener.Close()
	}
}
