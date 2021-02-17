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
	"context"
	"net"
)

// ServiceContext for RPC.
type ServiceContext struct {
	Context
	Method     Method
	LocalAddr  net.Addr
	RemoteAddr net.Addr
	service    *Service
}

// NewServiceContext returns a core.ServiceContext.
func NewServiceContext(service *Service) *ServiceContext {
	return &ServiceContext{
		Context: NewContext(),
		service: service,
	}
}

// Service returns the Service reference.
func (c *ServiceContext) Service() *Service {
	return c.service
}

// Clone returns a copy of this ServiceContext.
func (c *ServiceContext) Clone() Context {
	return &ServiceContext{
		c.Context.Clone(),
		c.Method,
		c.LocalAddr,
		c.RemoteAddr,
		c.service,
	}
}

// GetServiceContext returns the *core.ServiceContext bound to the context.
func GetServiceContext(ctx context.Context) *ServiceContext {
	return GetContext(ctx).(*ServiceContext)
}
