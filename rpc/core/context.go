/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/context.go                                      |
|                                                          |
| LastModified: Feb 21, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"context"
)

type contextKeyT string

var contextKey = contextKeyT("github.com/hprose/hprose-golang/rpc/core.Context")

// Context for RPC.
type Context interface {
	Items() Dict
	HasRequestHeaders() bool
	RequestHeaders() Dict
	HasResponseHeaders() bool
	ResponseHeaders() Dict
	Clone() Context
}

type rpcContext struct {
	items           Dict
	requestHeaders  Dict
	responseHeaders Dict
}

// NewContext returns a core.Context.
func NewContext() Context {
	return &rpcContext{}
}

func (c *rpcContext) Items() Dict {
	if c.items == nil {
		c.items = NewDict(nil)
	}
	return c.items
}

func (c *rpcContext) HasRequestHeaders() bool {
	return c.requestHeaders != nil && !c.requestHeaders.Empty()
}

func (c *rpcContext) RequestHeaders() Dict {
	if c.requestHeaders == nil {
		c.requestHeaders = NewDict(nil)
	}
	return c.requestHeaders
}

func (c *rpcContext) HasResponseHeaders() bool {
	return c.responseHeaders != nil && !c.responseHeaders.Empty()
}

func (c *rpcContext) ResponseHeaders() Dict {
	if c.responseHeaders == nil {
		c.responseHeaders = NewDict(nil)
	}
	return c.responseHeaders
}

func (c *rpcContext) Clone() Context {
	clone := &rpcContext{}
	if c.items != nil {
		clone.items = NewDict(nil)
		c.items.CopyTo(clone.items)
	}
	if c.requestHeaders != nil {
		clone.requestHeaders = NewDict(nil)
		c.requestHeaders.CopyTo(clone.requestHeaders)
	}
	if c.responseHeaders != nil {
		clone.responseHeaders = NewDict(nil)
		c.responseHeaders.CopyTo(clone.responseHeaders)
	}
	return clone
}

// WithContext returns a copy of the parent context and associates it with a core.Context.
func WithContext(ctx context.Context, rpcContext Context) context.Context {
	return context.WithValue(ctx, contextKey, rpcContext)
}

// FromContext returns the core.Context bound to the context.
func FromContext(ctx context.Context) (Context, bool) {
	c, ok := ctx.Value(contextKey).(Context)
	return c, ok
}
