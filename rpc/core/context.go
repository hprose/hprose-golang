/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/context.go                                      |
|                                                          |
| LastModified: Jan 25, 2021                               |
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
	Items
	HasRequestHeaders() bool
	RequestHeaders() Headers
	HasResponseHeaders() bool
	ResponseHeaders() Headers
	Clone() Context
}

type rpcContext struct {
	items           map[string]interface{}
	requestHeaders  Headers
	responseHeaders Headers
}

// NewContext returns a core.Context.
func NewContext() Context {
	return &rpcContext{}
}

func (c *rpcContext) Set(name string, value interface{}) {
	if c.items == nil {
		c.items = make(map[string]interface{})
	}
	c.items[name] = value
}

func (c *rpcContext) Get(name string) (value interface{}, ok bool) {
	if c.items == nil {
		return nil, false
	}
	value, ok = c.items[name]
	return
}

func (c *rpcContext) Del(name string) {
	if c.items != nil {
		delete(c.items, name)
	}
}

func (c *rpcContext) Range(f func(name string, value interface{}) bool) {
	if c.items != nil {
		for k, v := range c.items {
			if !f(k, v) {
				return
			}
		}
	}
}

func (c *rpcContext) HasRequestHeaders() bool {
	return c.requestHeaders != nil && len(c.requestHeaders.(headers)) > 0
}

func (c *rpcContext) RequestHeaders() Headers {
	if c.requestHeaders == nil {
		c.requestHeaders = NewHeaders()
	}
	return c.requestHeaders
}

func (c *rpcContext) HasResponseHeaders() bool {
	return c.responseHeaders != nil && len(c.requestHeaders.(headers)) > 0
}

func (c *rpcContext) ResponseHeaders() Headers {
	if c.responseHeaders == nil {
		c.responseHeaders = NewHeaders()
	}
	return c.responseHeaders
}

func (c *rpcContext) Clone() Context {
	clone := &rpcContext{}
	if c.items != nil {
		clone.items = make(map[string]interface{})
		for k, v := range c.items {
			clone.items[k] = v
		}
	}
	if c.requestHeaders != nil {
		clone.requestHeaders = c.requestHeaders.Clone()
	}
	if c.responseHeaders != nil {
		clone.responseHeaders = c.responseHeaders.Clone()
	}
	return clone
}

// WithContext returns a copy of the parent context and associates it with a core.Context.
func WithContext(ctx context.Context, rpcCtx Context) context.Context {
	return context.WithValue(ctx, contextKey, rpcCtx)
}

// FromContext returns the core.Context bound to the context, if any.
func FromContext(ctx context.Context) (rpcCtx Context, ok bool) {
	rpcCtx, ok = ctx.Value(contextKey).(Context)
	return
}
