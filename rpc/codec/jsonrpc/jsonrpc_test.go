/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/codec/jsonrpc/common.go                              |
|                                                          |
| LastModified: May 10, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package jsonrpc_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/hprose/hprose-golang/v3/rpc"
	"github.com/hprose/hprose-golang/v3/rpc/codec/jsonrpc"
	"github.com/hprose/hprose-golang/v3/rpc/core"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/log"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

type testService int

func (testService) Hello(name string) string {
	return "hello " + name
}

func (testService) Sum(n ...int) (result int) {
	for _, x := range n {
		result += x
	}
	return
}

func (testService) Swap(a, b int) (int, int) {
	return b, a
}

func (testService) Divide(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("divide by zero")
	}
	return a / b, a % b, nil
}

func (testService) Panic(v interface{}) {
	panic(v)
}

func TestJSONRPC(t *testing.T) {
	service := rpc.NewService()
	service.Codec = jsonrpc.NewServiceCodec(nil)
	service.AddInstanceMethods(new(testService))
	server := &http.Server{Addr: ":8000"}
	err := service.Bind(server)
	assert.NoError(t, err)
	go server.ListenAndServe()

	time.Sleep(time.Millisecond * 5)

	client := rpc.NewClient("http://127.0.0.1:8000/")
	client.Use(log.Plugin)
	var proxy struct {
		Hello      func(name string) (string, error)
		Sum        func(n ...int) (int, error)
		Swap       func(a, b int) (int, int, error)
		Divide     func(a, b int) (int, int, error)
		Sub        func(a, b int) (int, error)
		SwapString func(a, b string) (string, string, error) `name:"swap"`
		Panic      func(v interface{}) error
	}
	client.UseService(&proxy)
	{
		result, err := proxy.Hello("world")
		assert.Equal(t, "hello world", result)
		assert.NoError(t, err)
	}

	client.Codec = jsonrpc.NewClientCodec(nil)
	{
		result, err := proxy.Hello("world")
		assert.Equal(t, "hello world", result)
		assert.NoError(t, err)
	}
	{
		result, err := proxy.Sum(1, 2, 3, 4, 5)
		assert.Equal(t, 15, result)
		assert.NoError(t, err)
	}
	{
		a, b, err := proxy.Swap(1, 2)
		assert.Equal(t, a, 2)
		assert.Equal(t, b, 1)
		assert.NoError(t, err)
	}
	{
		q, r, err := proxy.Divide(3, 2)
		assert.Equal(t, q, 1)
		assert.Equal(t, r, 1)
		assert.NoError(t, err)
	}
	{
		q, r, err := proxy.Divide(3, 0)
		assert.Equal(t, q, 0)
		assert.Equal(t, r, 0)
		assert.EqualError(t, err, "divide by zero")
	}
	{
		result, err := proxy.Sub(3, 2)
		assert.Equal(t, result, 0)
		assert.EqualError(t, err, "hprose/rpc/codec/jsonrpc: Method not found")
	}
	{
		a, b, err := proxy.SwapString("hello", "world")
		assert.Equal(t, a, "")
		assert.Equal(t, b, "")
		assert.EqualError(t, err, "hprose/rpc/codec/jsonrpc: Invalid params")
	}
	{
		err := proxy.Panic("test panic")
		assert.EqualError(t, err, "test panic")
	}
	server.Close()
}

func TestHeaders(t *testing.T) {
	service := rpc.NewService()
	service.Codec = jsonrpc.NewServiceCodec(jsoniter.ConfigDefault)
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	service.Use(func(ctx context.Context, name string, args []interface{}, next core.NextInvokeHandler) (result []interface{}, err error) {
		serviceContext := core.GetServiceContext(ctx)
		ping := serviceContext.RequestHeaders().GetBool("ping")
		assert.True(t, ping)
		serviceContext.ResponseHeaders().Set("pong", true)
		return next(ctx, name, args)
	})
	server := &http.Server{Addr: ":8000"}
	err := service.Bind(server)
	assert.NoError(t, err)
	go server.ListenAndServe()

	time.Sleep(time.Millisecond * 5)

	client := rpc.NewClient("http://127.0.0.1:8000/")
	client.Codec = jsonrpc.NewClientCodec(jsoniter.ConfigDefault)
	client.Use(log.Plugin)
	var proxy struct {
		Hello func(ctx context.Context, name string) string `header:"ping"`
	}
	client.UseService(&proxy)
	clientContext := core.NewClientContext()
	ctx := core.WithContext(context.Background(), clientContext)
	result := proxy.Hello(ctx, "world")
	assert.Equal(t, `hello world`, result)
	assert.True(t, clientContext.ResponseHeaders().GetBool("pong"))
	server.Close()
}
