/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/service_context.go                              |
|                                                          |
| LastModified: Jan 24, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"net"
	"reflect"
)

// ServiceContext for RPC.
type ServiceContext interface {
	Context
	Service() Service
	Method() reflect.Method
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	SetMethod(method reflect.Method)
	SetLocalAddr(addr net.Addr)
	SetRemoteAddr(addr net.Addr)
}

type serviceContext struct {
	Context
	service    Service
	method     reflect.Method
	localAddr  net.Addr
	remoteAddr net.Addr
}

// NewServiceContext returns a core.ServiceContext.
func NewServiceContext(service Service) ServiceContext {
	return &serviceContext{
		Context: NewContext(),
		service: service,
	}
}

func (c *serviceContext) Service() Service {
	return c.service
}

func (c *serviceContext) Method() reflect.Method {
	return c.method
}

func (c *serviceContext) LocalAddr() net.Addr {
	return c.localAddr
}

func (c *serviceContext) RemoteAddr() net.Addr {
	return c.remoteAddr
}

func (c *serviceContext) SetMethod(method reflect.Method) {
	c.method = method
}

func (c *serviceContext) SetLocalAddr(addr net.Addr) {
	c.localAddr = addr
}

func (c *serviceContext) SetRemoteAddr(addr net.Addr) {
	c.remoteAddr = addr
}

func (c *serviceContext) Clone() Context {
	return &serviceContext{
		c.Context.Clone(),
		c.service,
		c.method,
		c.localAddr,
		c.remoteAddr,
	}
}
