/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/client.go                                       |
|                                                          |
| LastModified: Feb 18, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"context"
	"math/rand"
	"net/url"
	"reflect"
	"sync"
	"time"
)

// Transport is an interface used to represent client transport layer.
type Transport interface {
	Transport(context context.Context, request []byte) (response []byte, err error)
	Abort()
}

// TransportFactory is a constructor for Transport.
type TransportFactory interface {
	Schemes() []string
	New() Transport
}

var interfaceType = reflect.TypeOf((interface{})(nil))
var transportFactories sync.Map
var protocols sync.Map

// RegisterTransport for Client.
func RegisterTransport(name string, transportFactory TransportFactory) {
	transportFactories.Store(name, transportFactory)
	for _, scheme := range transportFactory.Schemes() {
		protocols.Store(scheme, name)
	}
}

// Client for RPC.
type Client struct {
	Codec          ClientCodec
	URLs           []*url.URL
	Timeout        time.Duration
	requestHeaders Dict
	invokeManager  PluginManager
	ioManager      PluginManager
	transports     map[string]Transport
}

// NewClient returns an instance of Client.
func NewClient(uri ...string) *Client {
	client := (&Client{
		Codec:          clientCodec{},
		Timeout:        time.Second * 30,
		requestHeaders: NewSafeDict(),
	})
	for _, u := range uri {
		if url, err := url.Parse(u); err != nil {
			client.URLs = append(client.URLs, url)
		}
	}
	transportFactories.Range(func(key, value interface{}) bool {
		transport := value.(TransportFactory).New()
		client.transports[key.(string)] = transport
		return true
	})
	client.invokeManager = NewInvokeManager(client.Call)
	client.ioManager = NewIOManager(client.Transport)
	return client
}

// ShuffleURLs sorts the URLs in random order.
func (c *Client) ShuffleURLs() *Client {
	src := c.URLs
	if n := len(src); n > 0 {
		dest := make([]*url.URL, n)
		rand.Seed(time.Now().UTC().UnixNano())
		perm := rand.Perm(n)
		for i, v := range perm {
			dest[v] = src[i]
		}
		c.URLs = dest
	}
	return c
}

// UseService build a remote service proxy object with namespace
func (c *Client) UseService(remoteService interface{}, namespace ...string) {

}

// RequestHeaders returns the global request headers.
func (c *Client) RequestHeaders() Dict {
	return c.requestHeaders
}

// InvokeContext the remote method with context.Context.
func (c *Client) InvokeContext(ctx context.Context, name string, args []interface{}) (result []interface{}, err error) {
	clientContext := GetClientContext(ctx)
	if clientContext == nil {
		clientContext = NewClientContext()
		ctx = WithContext(ctx, clientContext)
	}
	clientContext.Init(c, interfaceType)
	return c.invokeManager.Handler().(NextInvokeHandler)(ctx, name, args)
}

// Invoke the remote method.
func (c *Client) Invoke(name string, args []interface{}) (result []interface{}, err error) {
	return c.InvokeContext(context.Background(), name, args)
}

// Call the remote method.
func (c *Client) Call(ctx context.Context, name string, args []interface{}) (result []interface{}, err error) {
	var request, response []byte
	clientContext := GetClientContext(ctx)
	if request, err = c.Codec.Encode(name, args, clientContext); err == nil {
		if response, err = c.Request(ctx, request); err == nil {
			result, err = c.Codec.Decode(response, clientContext)
		}
	}
	return
}

// Request data to the server and returns the response data.
func (c *Client) Request(ctx context.Context, request []byte) (response []byte, err error) {
	return c.ioManager.Handler().(NextIOHandler)(ctx, request)
}

// Transport the request data to the server and returns the response data.
func (c *Client) Transport(ctx context.Context, request []byte) (response []byte, err error) {
	url := GetClientContext(ctx).URL
	if name, ok := protocols.Load(url.Scheme); ok {
		return c.transports[name.(string)].Transport(ctx, request)
	}
	return nil, UnsupportedProtocolError{url.Scheme}
}

// Abort the remote call.
func (c *Client) Abort() {
	var wg sync.WaitGroup
	for _, transport := range c.transports {
		wg.Add(1)
		go func(transport Transport) {
			defer wg.Done()
			transport.Abort()
		}(transport)
	}
	wg.Wait()
}

// Use plugin handlers.
func (c *Client) Use(handler ...PluginHandler) *Client {
	invokeHandlers, ioHandler := separatePluginHandlers(handler)
	if len(invokeHandlers) > 0 {
		c.invokeManager.Use(invokeHandlers...)
	}
	if len(ioHandler) > 0 {
		c.ioManager.Use(ioHandler...)
	}
	return c
}

// Unuse plugin handlers.
func (c *Client) Unuse(handler ...PluginHandler) *Client {
	invokeHandlers, ioHandler := separatePluginHandlers(handler)
	if len(invokeHandlers) > 0 {
		c.invokeManager.Unuse(invokeHandlers...)
	}
	if len(ioHandler) > 0 {
		c.ioManager.Unuse(ioHandler...)
	}
	return c
}
