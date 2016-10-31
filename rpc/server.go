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
 * rpc/server.go                                          *
 *                                                        *
 * hprose server for Go.                                  *
 *                                                        *
 * LastModified: Oct 5, 2016                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import (
	"os"
	"os/signal"
	"syscall"
)

// Server interface
type Server interface {
	Service
	URI() string
	Handle() (err error)
	Close()
	Start() (err error)
	Restart()
	Stop()
}

type starter struct {
	server Server
	c      chan os.Signal
}

// Start the hprose server
func (starter *starter) Start() (err error) {
	for {
		if err = starter.server.Handle(); err != nil {
			return err
		}
		starter.c = make(chan os.Signal, 1)
		signal.Notify(starter.c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		s := <-starter.c
		starter.server.Close()
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL:
			signal.Stop(starter.c)
			return
		}
	}
}

// Restart the hprose server
func (starter *starter) Restart() {
	starter.c <- syscall.SIGHUP
}

// Stop the hprose server
func (starter *starter) Stop() {
	starter.c <- syscall.SIGQUIT
}
