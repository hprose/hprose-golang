/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/mock/mock_test.go                                    |
|                                                          |
| LastModified: Feb 21, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package mock

import (
	"testing"
	"time"

	"github.com/hprose/hprose-golang/v3/rpc/core"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/log"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/timeout"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	server := Server{"testHelloWorld"}
	err := service.Bind(server)
	assert.NoError(t, err)
	client := core.NewClient("mock://testHelloWorld")
	client.Use(log.IOHandler)
	var proxy struct {
		Hello func(name string) string
	}
	client.UseService(&proxy)
	result := proxy.Hello("world")
	assert.Equal(t, "hello world", result)
	server.Close()
}

func TestClientTimeout(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(d time.Duration) {
		time.Sleep(d)
	}, "wait")
	server := Server{"testClientTimeout"}
	err := service.Bind(server)
	assert.NoError(t, err)
	client := core.NewClient("mock://testClientTimeout")
	client.Use(log.IOHandler)
	client.Timeout = time.Millisecond
	var proxy struct {
		Wait func(d time.Duration) error
	}
	client.UseService(&proxy)
	err = proxy.Wait(time.Second * 30)
	assert.True(t, core.IsTimeoutError(err))
	server.Close()
}

func TestServiceTimeout(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(d time.Duration) {
		time.Sleep(d)
	}, "wait")
	service.Use(timeout.GetHandler(time.Millisecond), log.IOHandler)
	server := Server{"testServiceTimeout"}
	err := service.Bind(server)
	assert.NoError(t, err)
	client := core.NewClient("mock://testServiceTimeout")
	var proxy struct {
		Wait func(d time.Duration) error
	}
	client.UseService(&proxy)
	err = proxy.Wait(time.Second * 30)
	assert.True(t, core.IsTimeoutError(err))
	server.Close()
}
