/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/websocket/websocket_test.go                          |
|                                                          |
| LastModified: May 5, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package websocket_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/hprose/hprose-golang/v3/rpc/core"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/log"
	_ "github.com/hprose/hprose-golang/v3/rpc/websocket"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	server := &http.Server{Addr: ":8000"}
	err := service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient("ws://127.0.0.1:8000/")
	client.Use(log.Plugin)
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.UseService(&proxy)
	result, err := proxy.Hello("world")
	assert.Equal(t, "hello world", result)
	assert.NoError(t, err)
	server.Close()
}
