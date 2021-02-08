/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/client_context.go                               |
|                                                          |
| LastModified: Feb 8, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"net/url"
	"reflect"
	"time"
)

// ClientContext for RPC.
type ClientContext interface {
	Context
	Init(client Client, returnType ...reflect.Type)
	Client() Client
	URL() url.URL
	ReturnType() []reflect.Type
	Timeout() time.Duration
	SetURL(url url.URL)
	SetReturnType(returnType []reflect.Type)
	SetTimeout(timeout time.Duration)
}

type clientContext struct {
	Context
	client     Client
	url        url.URL
	returnType []reflect.Type
	timeout    time.Duration
}

// NewClientContext returns a core.ClientContext.
func NewClientContext() ClientContext {
	return &clientContext{
		Context: NewContext(),
	}
}

func (c *clientContext) Init(client Client, returnType ...reflect.Type) {
	c.client = client
	if urls := c.client.URLs(); len(urls) > 0 {
		c.url = urls[0]
	}
	if c.returnType == nil {
		c.returnType = returnType
	}
	if c.timeout == 0 {
		c.timeout = client.Timeout()
	}
	if !client.RequestHeaders().Empty() {
		client.RequestHeaders().CopyTo(c.RequestHeaders())
	}
}

func (c *clientContext) Client() Client {
	return c.client
}

func (c *clientContext) URL() url.URL {
	return c.url
}

func (c *clientContext) ReturnType() []reflect.Type {
	return c.returnType
}

func (c *clientContext) Timeout() time.Duration {
	return c.timeout
}

func (c *clientContext) SetURL(url url.URL) {
	c.url = url
}

func (c *clientContext) SetReturnType(returnType []reflect.Type) {
	c.returnType = returnType
}

func (c *clientContext) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}

func (c *clientContext) Clone() Context {
	return &clientContext{
		c.Context.Clone(),
		c.client,
		c.url,
		c.returnType,
		c.timeout,
	}
}
