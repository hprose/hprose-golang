/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/client_context.go                               |
|                                                          |
| LastModified: Feb 18, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"context"
	"net/url"
	"reflect"
	"time"
)

// ClientContext for RPC.
type ClientContext struct {
	Context
	URL        *url.URL
	ReturnType []reflect.Type
	Timeout    time.Duration
	client     *Client
}

// NewClientContext returns a core.ClientContext.
func NewClientContext() *ClientContext {
	return &ClientContext{
		Context: NewContext(),
	}
}

// Init this ClientContext.
func (c *ClientContext) Init(client *Client, returnType ...reflect.Type) {
	c.client = client
	if urls := client.URLs; len(urls) > 0 {
		c.URL = urls[0]
	}
	if c.ReturnType == nil {
		c.ReturnType = returnType
	}
	if c.Timeout == 0 {
		c.Timeout = client.Timeout
	}
	if !client.RequestHeaders().Empty() {
		client.RequestHeaders().CopyTo(c.RequestHeaders())
	}
}

// Client returns the Client reference.
func (c *ClientContext) Client() *Client {
	return c.client
}

// Clone returns a copy of this ClientContext.
func (c *ClientContext) Clone() Context {
	return &ClientContext{
		c.Context.Clone(),
		c.URL,
		c.ReturnType,
		c.Timeout,
		c.client,
	}
}

// GetClientContext returns the *core.ClientContext bound to the context.
func GetClientContext(ctx context.Context) *ClientContext {
	if c, ok := FromContext(ctx); ok {
		return c.(*ClientContext)
	}
	return nil
}
