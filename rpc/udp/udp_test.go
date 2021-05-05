/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/udp/udp_test.go                                      |
|                                                          |
| LastModified: May 5, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package udp_test

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/hprose/hprose-golang/v3/rpc/core"
	_ "github.com/hprose/hprose-golang/v3/rpc/http"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/circuitbreaker"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/cluster"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/forward"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/limiter"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/loadbalance"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/log"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/oneway"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/timeout"
	_ "github.com/hprose/hprose-golang/v3/rpc/socket"
	_ "github.com/hprose/hprose-golang/v3/rpc/udp"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8412")
	assert.NoError(t, err)
	server, err := net.ListenUDP("udp", addr)
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient("udp://127.0.0.1/")
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

func TestClientTimeout(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(d time.Duration) {
		time.Sleep(d)
	}, "wait")
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8412")
	assert.NoError(t, err)
	server, err := net.ListenUDP("udp", addr)
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient("udp://127.0.0.1/")
	client.Use(log.Plugin)
	client.Timeout = time.Millisecond
	var proxy struct {
		Wait func(d time.Duration) error
	}
	client.UseService(&proxy)
	err = proxy.Wait(time.Millisecond * 30)
	assert.True(t, core.IsTimeoutError(err))
	server.Close()
}

func TestServiceTimeout(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(d time.Duration) {
		time.Sleep(d)
	}, "wait")
	service.Use(timeout.New(5 * time.Millisecond))
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8412")
	assert.NoError(t, err)
	server, err := net.ListenUDP("udp", addr)
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient("udp://127.0.0.1/")
	var proxy struct {
		Wait func(d time.Duration) error
	}
	client.UseService(&proxy)
	client.Use(log.IOHandler, log.InvokeHandler)
	err = proxy.Wait(time.Millisecond)
	assert.False(t, core.IsTimeoutError(err))
	err = proxy.Wait(time.Millisecond * 30)
	assert.True(t, core.IsTimeoutError(err))
	server.Close()
}

func TestMissingMethod(t *testing.T) {
	service := core.NewService()
	service.AddMissingMethod(func(name string, args []interface{}) (result []interface{}, err error) {
		data, err := json.Marshal(args)
		if err != nil {
			return nil, err
		}
		return []interface{}{name + string(data)}, nil
	})
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8412")
	assert.NoError(t, err)
	server, err := net.ListenUDP("udp", addr)
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient("udp://127.0.0.1/")
	client.Use(log.IOHandler, log.InvokeHandler)
	var proxy struct {
		Hello func(name string) string
	}
	client.UseService(&proxy)
	result := proxy.Hello("world")
	assert.Equal(t, `Hello["world"]`, result)
	server.Close()
}

func TestMissingMethod2(t *testing.T) {
	service := core.NewService()
	service.AddMissingMethod(func(ctx context.Context, name string, args []interface{}) (result []interface{}, err error) {
		serviceContext := core.GetServiceContext(ctx)
		data, err := json.Marshal(args)
		if err != nil {
			return nil, err
		}
		return []interface{}{name + string(data) + serviceContext.LocalAddr.String()}, nil
	})
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8412")
	assert.NoError(t, err)
	server, err := net.ListenUDP("udp", addr)
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient("udp://127.0.0.1/")
	client.Use(log.Plugin)
	var proxy struct {
		Hello func(name string) string
	}
	client.UseService(&proxy)
	result := proxy.Hello("world")
	assert.Equal(t, `Hello["world"]127.0.0.1:8412`, result)
	server.Close()
}

func TestHeaders(t *testing.T) {
	service := core.NewService()
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
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8412")
	assert.NoError(t, err)
	server, err := net.ListenUDP("udp", addr)
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient("udp://127.0.0.1/")
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

func TestMaxRequestLength(t *testing.T) {
	service := core.NewService()
	service.MaxRequestLength = 10
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8412")
	assert.NoError(t, err)
	server, err := net.ListenUDP("udp", addr)
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient("udp://127.0.0.1/")
	client.Use(log.Plugin)
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.UseService(&proxy)
	_, err = proxy.Hello("world")
	if assert.Error(t, err) {
		assert.Equal(t, core.ErrRequestEntityTooLarge, err)
	}
	server.Close()
}

func TestCircuitBreaker(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8412")
	assert.NoError(t, err)
	server, err := net.ListenUDP("udp", addr)
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient("udp://127.0.0.1/")
	client.Use(circuitbreaker.New(
		circuitbreaker.WithThreshold(3),
		circuitbreaker.WithRecoverTime(time.Millisecond*10),
	))
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.UseService(&proxy)
	result, err := proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}
	server.Close()
	time.Sleep(time.Millisecond)

	for i := 0; i < 4; i++ {
		_, err = proxy.Hello("world")
		assert.Error(t, err)
	}
	_, err = proxy.Hello("world")
	if assert.Error(t, err) {
		assert.Equal(t, "service breaked", err.Error())
	}
	server, err = net.ListenUDP("udp", addr)
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	_, err = proxy.Hello("world")
	if assert.Error(t, err) {
		assert.Equal(t, "service breaked", err.Error())
	}
	time.Sleep(time.Millisecond * 10)
	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}
	server.Close()
}

func TestCircuitBreaker2(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8412")
	assert.NoError(t, err)
	server, err := net.ListenUDP("udp", addr)
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient("udp://127.0.0.1/")
	client.Use(circuitbreaker.New(
		circuitbreaker.WithThreshold(1),
		circuitbreaker.WithRecoverTime(time.Millisecond*10),
		circuitbreaker.WithMockService(func(ctx context.Context, name string, args []interface{}) (result []interface{}, err error) {
			return []interface{}{name + " breaked"}, nil
		}),
	))
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.UseService(&proxy)
	result, err := proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}
	server.Close()
	time.Sleep(time.Millisecond)

	_, err = proxy.Hello("world")
	assert.Error(t, err)
	_, err = proxy.Hello("world")
	assert.Error(t, err)
	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "Hello breaked", result)
	}

	server, err = net.ListenUDP("udp", addr)
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "Hello breaked", result)
	}
	time.Sleep(time.Millisecond * 10)
	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}
	server.Close()
}

func TestClusterFailover1(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")

	addr1, err := net.ResolveUDPAddr("udp", "127.0.0.1:8401")
	assert.NoError(t, err)
	server1, err := net.ListenUDP("udp", addr1)
	assert.NoError(t, err)
	err = service.Bind(server1)
	assert.NoError(t, err)

	addr2, err := net.ResolveUDPAddr("udp", "127.0.0.1:8402")
	assert.NoError(t, err)
	server2, err := net.ListenUDP("udp", addr2)
	assert.NoError(t, err)
	err = service.Bind(server2)
	assert.NoError(t, err)

	addr3, err := net.ResolveUDPAddr("udp", "127.0.0.1:8403")
	assert.NoError(t, err)
	server3, err := net.ListenUDP("udp", addr3)
	assert.NoError(t, err)
	err = service.Bind(server3)
	assert.NoError(t, err)

	addr4, err := net.ResolveUDPAddr("udp", "127.0.0.1:8404")
	assert.NoError(t, err)
	server4, err := net.ListenUDP("udp", addr4)
	assert.NoError(t, err)
	err = service.Bind(server4)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient(
		"udp://127.0.0.1:8401/",
		"udp://127.0.0.1:8402/",
		"udp://127.0.0.1:8403/",
		"udp://127.0.0.1:8404/",
	)
	client.Use(cluster.New(
		cluster.FailoverConfig(),
	)).Use(log.Plugin)
	var proxy struct {
		Hello func(name string) (string, error) `context:"idempotent,retry:3"`
	}
	client.UseService(&proxy)
	result, err := proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server1.Close()
	time.Sleep(time.Millisecond)

	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server2.Close()
	time.Sleep(time.Millisecond)

	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server3.Close()
	time.Sleep(time.Millisecond)

	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server4.Close()
	time.Sleep(time.Millisecond)

	client.UseService(&proxy)
	_, err = proxy.Hello("world")
	assert.Error(t, err)
}

func TestClusterFailover2(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr1, err := net.ResolveUDPAddr("udp", "127.0.0.1:8401")
	assert.NoError(t, err)
	server1, err := net.ListenUDP("udp", addr1)
	assert.NoError(t, err)
	err = service.Bind(server1)
	assert.NoError(t, err)

	addr2, err := net.ResolveUDPAddr("udp", "127.0.0.1:8402")
	assert.NoError(t, err)
	server2, err := net.ListenUDP("udp", addr2)
	assert.NoError(t, err)
	err = service.Bind(server2)
	assert.NoError(t, err)

	addr3, err := net.ResolveUDPAddr("udp", "127.0.0.1:8403")
	assert.NoError(t, err)
	server3, err := net.ListenUDP("udp", addr3)
	assert.NoError(t, err)
	err = service.Bind(server3)
	assert.NoError(t, err)

	addr4, err := net.ResolveUDPAddr("udp", "127.0.0.1:8404")
	assert.NoError(t, err)
	server4, err := net.ListenUDP("udp", addr4)
	assert.NoError(t, err)
	err = service.Bind(server4)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient(
		"udp://127.0.0.1:8401/",
		"udp://127.0.0.1:8402/",
		"udp://127.0.0.1:8403/",
		"udp://127.0.0.1:8404/",
	)
	client.Use(cluster.New(
		cluster.FailoverConfig(
			cluster.WithIdempotent(true),
			cluster.WithRetry(3),
		),
	)).Use(log.Plugin)
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.UseService(&proxy)
	result, err := proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server1.Close()
	time.Sleep(time.Millisecond)

	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server2.Close()
	time.Sleep(time.Millisecond)

	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server3.Close()
	time.Sleep(time.Millisecond)

	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server4.Close()
	time.Sleep(time.Millisecond)

	_, err = proxy.Hello("world")
	assert.Error(t, err)
}

func TestClusterFailtry(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8412")
	assert.NoError(t, err)
	server, err := net.ListenUDP("udp", addr)
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient("udp://127.0.0.1/")
	client.Use(cluster.New(
		cluster.FailtryConfig(
			cluster.WithIdempotent(true),
			cluster.WithRetry(3),
		),
	)).Use(log.Plugin)
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.UseService(&proxy)
	result, err := proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server.Close()
	time.Sleep(time.Millisecond)

	go func() {
		time.Sleep(time.Second)
		server, _ = net.ListenUDP("udp", addr)
		_ = service.Bind(server)
	}()

	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server.Close()
	time.Sleep(time.Millisecond)

	_, err = proxy.Hello("world")
	assert.Error(t, err)
}

func TestClusterFailfast(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8412")
	assert.NoError(t, err)
	server, err := net.ListenUDP("udp", addr)
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient("udp://127.0.0.1/")
	client.Use(cluster.New(
		cluster.FailfastConfig(
			func(c context.Context) {
				fmt.Println("TestClusterFailfast ok")
			},
		),
	)).Use(log.Plugin)
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.UseService(&proxy)
	result, err := proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server.Close()
	time.Sleep(time.Millisecond)

	_, err = proxy.Hello("world")
	assert.Error(t, err)
}

func TestClusterSuccess(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8412")
	assert.NoError(t, err)
	server, err := net.ListenUDP("udp", addr)
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient("udp://127.0.0.1/")
	client.Use(cluster.New(
		cluster.Config{
			OnSuccess: func(ctx context.Context) {
				fmt.Println("TestClusterSuccess ok")
			},
		},
	)).Use(log.Plugin)
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.UseService(&proxy)
	result, err := proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}
	server.Close()
}

func TestClusterForking(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr1, err := net.ResolveUDPAddr("udp", "127.0.0.1:8401")
	assert.NoError(t, err)
	server1, err := net.ListenUDP("udp", addr1)
	assert.NoError(t, err)
	err = service.Bind(server1)
	assert.NoError(t, err)

	addr2, err := net.ResolveUDPAddr("udp", "127.0.0.1:8402")
	assert.NoError(t, err)
	server2, err := net.ListenUDP("udp", addr2)
	assert.NoError(t, err)
	err = service.Bind(server2)
	assert.NoError(t, err)

	addr3, err := net.ResolveUDPAddr("udp", "127.0.0.1:8403")
	assert.NoError(t, err)
	server3, err := net.ListenUDP("udp", addr3)
	assert.NoError(t, err)
	err = service.Bind(server3)
	assert.NoError(t, err)

	addr4, err := net.ResolveUDPAddr("udp", "127.0.0.1:8404")
	assert.NoError(t, err)
	server4, err := net.ListenUDP("udp", addr4)
	assert.NoError(t, err)
	err = service.Bind(server4)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient(
		"udp://127.0.0.1:8401/",
		"udp://127.0.0.1:8402/",
		"udp://127.0.0.1:8403/",
		"udp://127.0.0.1:8404/",
	)
	client.Use(cluster.Forking).Use(log.Plugin)
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.UseService(&proxy)
	result, err := proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server1.Close()
	time.Sleep(time.Millisecond)

	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server2.Close()
	time.Sleep(time.Millisecond)

	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server3.Close()
	time.Sleep(time.Millisecond)

	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server4.Close()
	time.Sleep(time.Millisecond)

	_, err = proxy.Hello("world")
	assert.Error(t, err)
}

func TestClusterBroadcast(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr1, err := net.ResolveUDPAddr("udp", "127.0.0.1:8401")
	assert.NoError(t, err)
	server1, err := net.ListenUDP("udp", addr1)
	assert.NoError(t, err)
	err = service.Bind(server1)
	assert.NoError(t, err)

	addr2, err := net.ResolveUDPAddr("udp", "127.0.0.1:8402")
	assert.NoError(t, err)
	server2, err := net.ListenUDP("udp", addr2)
	assert.NoError(t, err)
	err = service.Bind(server2)
	assert.NoError(t, err)

	addr3, err := net.ResolveUDPAddr("udp", "127.0.0.1:8403")
	assert.NoError(t, err)
	server3, err := net.ListenUDP("udp", addr3)
	assert.NoError(t, err)
	err = service.Bind(server3)
	assert.NoError(t, err)

	addr4, err := net.ResolveUDPAddr("udp", "127.0.0.1:8404")
	assert.NoError(t, err)
	server4, err := net.ListenUDP("udp", addr4)
	assert.NoError(t, err)
	err = service.Bind(server4)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient(
		"udp://127.0.0.1:8401/",
		"udp://127.0.0.1:8402/",
		"udp://127.0.0.1:8403/",
		"udp://127.0.0.1:8404/",
	)
	client.Use(cluster.Broadcast).Use(log.Plugin)
	clientContext := core.NewClientContext()
	clientContext.ReturnType = append(clientContext.ReturnType, reflect.TypeOf(""))
	result, err := client.InvokeContext(core.WithContext(context.Background(), clientContext), "hello", []interface{}{"world"})
	if assert.NoError(t, err) {
		assert.Equal(t, []interface{}{
			[]interface{}{"hello world"},
			[]interface{}{"hello world"},
			[]interface{}{"hello world"},
			[]interface{}{"hello world"},
		}, result)
	}

	server1.Close()
	time.Sleep(time.Millisecond)

	result, err = client.InvokeContext(core.WithContext(context.Background(), clientContext), "hello", []interface{}{"world"})
	assert.Error(t, err)
	assert.Equal(t, []interface{}{
		[]interface{}(nil),
		[]interface{}{"hello world"},
		[]interface{}{"hello world"},
		[]interface{}{"hello world"},
	}, result)

	server2.Close()
	time.Sleep(time.Millisecond)

	result, err = client.InvokeContext(core.WithContext(context.Background(), clientContext), "hello", []interface{}{"world"})
	assert.Error(t, err)
	assert.Equal(t, []interface{}{
		[]interface{}(nil),
		[]interface{}(nil),
		[]interface{}{"hello world"},
		[]interface{}{"hello world"},
	}, result)

	server3.Close()
	time.Sleep(time.Millisecond)

	result, err = client.InvokeContext(core.WithContext(context.Background(), clientContext), "hello", []interface{}{"world"})
	assert.Error(t, err)
	assert.Equal(t, []interface{}{
		[]interface{}(nil),
		[]interface{}(nil),
		[]interface{}(nil),
		[]interface{}{"hello world"},
	}, result)

	server4.Close()
	time.Sleep(time.Millisecond)

	result, err = client.InvokeContext(core.WithContext(context.Background(), clientContext), "hello", []interface{}{"world"})
	assert.Error(t, err)
	assert.Equal(t, []interface{}{
		[]interface{}(nil),
		[]interface{}(nil),
		[]interface{}(nil),
		[]interface{}(nil),
	}, result)
}

func TestForward(t *testing.T) {
	service1 := core.NewService()
	service1.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr1, err := net.ResolveUDPAddr("udp", "127.0.0.1:8401")
	assert.NoError(t, err)
	server1, err := net.ListenUDP("udp", addr1)
	assert.NoError(t, err)
	err = service1.Bind(server1)
	assert.NoError(t, err)

	fw := forward.New("udp://127.0.0.1:8401/")
	fw.Use(log.Plugin)
	service2 := core.NewService()
	service2.AddMissingMethod(fw.Forward)
	// service2.AddMissingMethod(func(ctx context.Context, name string, args []interface{}) (result []interface{}, err error) {
	// 	return
	// })
	// service2.Use(fw.InvokeHandler)
	addr2, err := net.ResolveUDPAddr("udp", "127.0.0.1:8402")
	assert.NoError(t, err)
	server2, err := net.ListenUDP("udp", addr2)
	assert.NoError(t, err)
	err = service2.Bind(server2)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient("udp://127.0.0.1:8402/")
	client.Use(log.Plugin)
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.UseService(&proxy)
	result, err := proxy.Hello("invoke forward")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello invoke forward", result)
	}

	service2.Remove("*")
	service2.Use(fw.IOHandler)

	result, err = proxy.Hello("io forward")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello io forward", result)
	}

	fw.Unuse(log.Plugin)
	result, err = proxy.Hello("forward")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello forward", result)
	}

	server1.Close()
	server2.Close()
}

func TestConcurrentLimiter(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8412")
	assert.NoError(t, err)
	server, err := net.ListenUDP("udp", addr)
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)
	time.Sleep(time.Millisecond * 5)

	client := core.NewClient("udp://127.0.0.1/")
	var proxy struct {
		Hello func(name string) (string, error)
	}
	cl := limiter.NewConcurrentLimiter(3, time.Nanosecond)
	client.Use(cl)
	client.UseService(&proxy)
	assert.Equal(t, 3, cl.MaxConcurrentRequests())
	assert.Equal(t, time.Nanosecond, cl.Timeout())
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(i int) {
			defer wg.Done()
			assert.GreaterOrEqual(t, cl.ConcurrentRequests(), 0)
			result, err := proxy.Hello(fmt.Sprintf("world %d", i))
			assert.LessOrEqual(t, cl.ConcurrentRequests(), 3)
			if err == nil {
				assert.Equal(t, fmt.Sprintf("hello world %d", i), result)
			} else {
				assert.Equal(t, core.ErrTimeout, err)
			}
		}(i)
	}
	wg.Wait()
	server.Close()
}

func TestConcurrentLimiterWithoutTimeout(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8412")
	assert.NoError(t, err)
	server, err := net.ListenUDP("udp", addr)
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient("udp://127.0.0.1/")
	var proxy struct {
		Hello func(name string) (string, error)
	}
	cl := limiter.NewConcurrentLimiter(3)
	client.Use(cl)
	client.UseService(&proxy)
	assert.Equal(t, 3, cl.MaxConcurrentRequests())
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			assert.GreaterOrEqual(t, cl.ConcurrentRequests(), 0)
			result, err := proxy.Hello(fmt.Sprintf("world %d", i))
			assert.LessOrEqual(t, cl.ConcurrentRequests(), 3)
			if assert.NoError(t, err) {
				assert.Equal(t, fmt.Sprintf("hello world %d", i), result)
			}
		}(i)
	}
	wg.Wait()
	server.Close()
}

func TestRateLimiterIOHandler(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8412")
	assert.NoError(t, err)
	server, err := net.ListenUDP("udp", addr)
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient("udp://127.0.0.1/")
	var proxy struct {
		Hello func(name string) (string, error)
	}
	rl := limiter.NewRateLimiter(10000, limiter.WithTimeout(time.Millisecond*250))
	client.Use(rl.IOHandler)
	client.UseService(&proxy)
	assert.Equal(t, int64(10000), rl.PermitsPerSecond())
	assert.Equal(t, math.Inf(0), rl.MaxPermits())
	assert.Equal(t, time.Millisecond*250, rl.Timeout())
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			result, err := proxy.Hello(fmt.Sprintf("world %d", i))
			if err == nil {
				assert.Equal(t, fmt.Sprintf("hello world %d", i), result)
			} else {
				assert.Equal(t, core.ErrTimeout, err)
			}
		}(i)
	}
	wg.Wait()
	server.Close()
}

func TestRateLimiterInvokeHandler(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8412")
	assert.NoError(t, err)
	server, err := net.ListenUDP("udp", addr)
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient("udp://127.0.0.1/")
	var proxy struct {
		Hello func(name string) (string, error)
	}
	rl := limiter.NewRateLimiter(1000, limiter.WithMaxPermits(1), limiter.WithTimeout(time.Millisecond*80))
	client.Use(rl.InvokeHandler)
	client.UseService(&proxy)
	assert.Equal(t, int64(1000), rl.PermitsPerSecond())
	assert.Equal(t, float64(1), rl.MaxPermits())
	assert.Equal(t, time.Millisecond*80, rl.Timeout())
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			result, err := proxy.Hello(fmt.Sprintf("world %d", i))
			if err == nil {
				assert.Equal(t, fmt.Sprintf("hello world %d", i), result)
			} else {
				assert.Equal(t, core.ErrTimeout, err)
			}
		}(i)
	}
	wg.Wait()
	server.Close()
}

func TestRandomLoadBalance(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr1, err := net.ResolveUDPAddr("udp", "127.0.0.1:8401")
	assert.NoError(t, err)
	server1, err := net.ListenUDP("udp", addr1)
	assert.NoError(t, err)
	err = service.Bind(server1)
	assert.NoError(t, err)

	addr2, err := net.ResolveUDPAddr("udp", "127.0.0.1:8402")
	assert.NoError(t, err)
	server2, err := net.ListenUDP("udp", addr2)
	assert.NoError(t, err)
	err = service.Bind(server2)
	assert.NoError(t, err)

	addr3, err := net.ResolveUDPAddr("udp", "127.0.0.1:8403")
	assert.NoError(t, err)
	server3, err := net.ListenUDP("udp", addr3)
	assert.NoError(t, err)
	err = service.Bind(server3)
	assert.NoError(t, err)

	addr4, err := net.ResolveUDPAddr("udp", "127.0.0.1:8404")
	assert.NoError(t, err)
	server4, err := net.ListenUDP("udp", addr4)
	assert.NoError(t, err)
	err = service.Bind(server4)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient(
		"udp://127.0.0.1:8401/",
		"udp://127.0.0.1:8402/",
		"udp://127.0.0.1:8403/",
		"udp://127.0.0.1:8404/",
	)
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.Use(loadbalance.NewRandomLoadBalance())
	client.UseService(&proxy)
	var wg sync.WaitGroup
	wg.Add(100)
	var rwlock sync.RWMutex
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			rwlock.RLock()
			result, err := proxy.Hello(fmt.Sprintf("world %d", i))
			rwlock.RUnlock()
			if err == nil {
				assert.Equal(t, fmt.Sprintf("hello world %d", i), result)
			} else {
				assert.Equal(t, core.ErrTimeout, err)
			}
			if i == 50 {
				rwlock.Lock()
				client.URLs = client.URLs[:3]
				rwlock.Unlock()
			}
		}(i)
	}
	wg.Wait()
	server1.Close()
	server2.Close()
	server3.Close()
	server4.Close()
}

func TestRoundRobinLoadBalance(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr1, err := net.ResolveUDPAddr("udp", "127.0.0.1:8401")
	assert.NoError(t, err)
	server1, err := net.ListenUDP("udp", addr1)
	assert.NoError(t, err)
	err = service.Bind(server1)
	assert.NoError(t, err)

	addr2, err := net.ResolveUDPAddr("udp", "127.0.0.1:8402")
	assert.NoError(t, err)
	server2, err := net.ListenUDP("udp", addr2)
	assert.NoError(t, err)
	err = service.Bind(server2)
	assert.NoError(t, err)

	addr3, err := net.ResolveUDPAddr("udp", "127.0.0.1:8403")
	assert.NoError(t, err)
	server3, err := net.ListenUDP("udp", addr3)
	assert.NoError(t, err)
	err = service.Bind(server3)
	assert.NoError(t, err)

	addr4, err := net.ResolveUDPAddr("udp", "127.0.0.1:8404")
	assert.NoError(t, err)
	server4, err := net.ListenUDP("udp", addr4)
	assert.NoError(t, err)
	err = service.Bind(server4)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient(
		"udp://127.0.0.1:8401/",
		"udp://127.0.0.1:8402/",
		"udp://127.0.0.1:8403/",
		"udp://127.0.0.1:8404/",
	)
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.Use(loadbalance.NewRoundRobinLoadBalance())
	client.UseService(&proxy)
	var wg sync.WaitGroup
	wg.Add(100)
	var rwlock sync.RWMutex
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			rwlock.RLock()
			result, err := proxy.Hello(fmt.Sprintf("world %d", i))
			rwlock.RUnlock()
			if err == nil {
				assert.Equal(t, fmt.Sprintf("hello world %d", i), result)
			} else {
				assert.Equal(t, core.ErrTimeout, err)
			}
			if i == 50 {
				rwlock.Lock()
				client.URLs = client.URLs[:3]
				rwlock.Unlock()
			}
		}(i)
	}
	wg.Wait()
	server1.Close()
	server2.Close()
	server3.Close()
	server4.Close()
}

func TestLeastActiveLoadBalance(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr1, err := net.ResolveUDPAddr("udp", "127.0.0.1:8401")
	assert.NoError(t, err)
	server1, err := net.ListenUDP("udp", addr1)
	assert.NoError(t, err)
	err = service.Bind(server1)
	assert.NoError(t, err)

	addr2, err := net.ResolveUDPAddr("udp", "127.0.0.1:8402")
	assert.NoError(t, err)
	server2, err := net.ListenUDP("udp", addr2)
	assert.NoError(t, err)
	err = service.Bind(server2)
	assert.NoError(t, err)

	addr3, err := net.ResolveUDPAddr("udp", "127.0.0.1:8403")
	assert.NoError(t, err)
	server3, err := net.ListenUDP("udp", addr3)
	assert.NoError(t, err)
	err = service.Bind(server3)
	assert.NoError(t, err)

	addr4, err := net.ResolveUDPAddr("udp", "127.0.0.1:8404")
	assert.NoError(t, err)
	server4, err := net.ListenUDP("udp", addr4)
	assert.NoError(t, err)
	err = service.Bind(server4)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient(
		"udp://127.0.0.1:8401/",
		"udp://127.0.0.1:8402/",
		"udp://127.0.0.1:8403/",
		"udp://127.0.0.1:8404/",
	)
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.Use(loadbalance.NewLeastActiveLoadBalance())
	client.UseService(&proxy)
	var wg sync.WaitGroup
	wg.Add(100)
	var rwlock sync.RWMutex
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			rwlock.RLock()
			result, err := proxy.Hello(fmt.Sprintf("world %d", i))
			rwlock.RUnlock()
			if err == nil {
				assert.Equal(t, fmt.Sprintf("hello world %d", i), result)
			} else {
				assert.Equal(t, core.ErrTimeout, err)
			}
			if i == 50 {
				rwlock.Lock()
				client.URLs = client.URLs[:3]
				rwlock.Unlock()
			}
		}(i)
	}
	wg.Wait()
	server1.Close()
	server2.Close()
	server3.Close()
	server4.Close()
}

func TestWeightedRandomLoadBalance(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr1, err := net.ResolveUDPAddr("udp", "127.0.0.1:8401")
	assert.NoError(t, err)
	server1, err := net.ListenUDP("udp", addr1)
	assert.NoError(t, err)
	err = service.Bind(server1)
	assert.NoError(t, err)

	addr2, err := net.ResolveUDPAddr("udp", "127.0.0.1:8402")
	assert.NoError(t, err)
	server2, err := net.ListenUDP("udp", addr2)
	assert.NoError(t, err)
	err = service.Bind(server2)
	assert.NoError(t, err)

	addr3, err := net.ResolveUDPAddr("udp", "127.0.0.1:8403")
	assert.NoError(t, err)
	server3, err := net.ListenUDP("udp", addr3)
	assert.NoError(t, err)
	err = service.Bind(server3)
	assert.NoError(t, err)

	addr4, err := net.ResolveUDPAddr("udp", "127.0.0.1:8404")
	assert.NoError(t, err)
	server4, err := net.ListenUDP("udp", addr4)
	assert.NoError(t, err)
	err = service.Bind(server4)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient()
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.Use(loadbalance.NewWeightedRandomLoadBalance(map[string]int{
		"udp://127.0.0.1:8401/": 1,
		"udp://127.0.0.1:8402/": 2,
		"udp://127.0.0.1:8403/": 3,
		"udp://127.0.0.1:8404/": 4,
	}))
	client.UseService(&proxy)
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			result, err := proxy.Hello(fmt.Sprintf("world %d", i))
			if err == nil {
				assert.Equal(t, fmt.Sprintf("hello world %d", i), result)
			} else {
				assert.Equal(t, core.ErrTimeout, err)
			}
		}(i)
	}
	wg.Wait()
	server1.Close()
	server2.Close()
	server3.Close()
	server4.Close()
}

func TestWeightedRoundRobinLoadBalance(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr1, err := net.ResolveUDPAddr("udp", "127.0.0.1:8401")
	assert.NoError(t, err)
	server1, err := net.ListenUDP("udp", addr1)
	assert.NoError(t, err)
	err = service.Bind(server1)
	assert.NoError(t, err)

	addr2, err := net.ResolveUDPAddr("udp", "127.0.0.1:8402")
	assert.NoError(t, err)
	server2, err := net.ListenUDP("udp", addr2)
	assert.NoError(t, err)
	err = service.Bind(server2)
	assert.NoError(t, err)

	addr3, err := net.ResolveUDPAddr("udp", "127.0.0.1:8403")
	assert.NoError(t, err)
	server3, err := net.ListenUDP("udp", addr3)
	assert.NoError(t, err)
	err = service.Bind(server3)
	assert.NoError(t, err)

	addr4, err := net.ResolveUDPAddr("udp", "127.0.0.1:8404")
	assert.NoError(t, err)
	server4, err := net.ListenUDP("udp", addr4)
	assert.NoError(t, err)
	err = service.Bind(server4)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient()
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.Use(loadbalance.NewWeightedRoundRobinLoadBalance(map[string]int{
		"udp://127.0.0.1:8401/": 1,
		"udp://127.0.0.1:8402/": 2,
		"udp://127.0.0.1:8403/": 3,
		"udp://127.0.0.1:8404/": 4,
	}))
	client.UseService(&proxy)
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			result, err := proxy.Hello(fmt.Sprintf("world %d", i))
			if err == nil {
				assert.Equal(t, fmt.Sprintf("hello world %d", i), result)
			} else {
				assert.Equal(t, core.ErrTimeout, err)
			}
		}(i)
	}
	wg.Wait()
	server1.Close()
	server2.Close()
	server3.Close()
	server4.Close()
}

func TestNginxRoundRobinLoadBalance(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr1, err := net.ResolveUDPAddr("udp", "127.0.0.1:8401")
	assert.NoError(t, err)
	server1, err := net.ListenUDP("udp", addr1)
	assert.NoError(t, err)
	err = service.Bind(server1)
	assert.NoError(t, err)

	addr2, err := net.ResolveUDPAddr("udp", "127.0.0.1:8402")
	assert.NoError(t, err)
	server2, err := net.ListenUDP("udp", addr2)
	assert.NoError(t, err)
	err = service.Bind(server2)
	assert.NoError(t, err)

	addr3, err := net.ResolveUDPAddr("udp", "127.0.0.1:8403")
	assert.NoError(t, err)
	server3, err := net.ListenUDP("udp", addr3)
	assert.NoError(t, err)
	err = service.Bind(server3)
	assert.NoError(t, err)

	addr4, err := net.ResolveUDPAddr("udp", "127.0.0.1:8404")
	assert.NoError(t, err)
	server4, err := net.ListenUDP("udp", addr4)
	assert.NoError(t, err)
	err = service.Bind(server4)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient()
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.Use(loadbalance.NewNginxRoundRobinLoadBalance(map[string]int{
		"udp://127.0.0.1:8401/": 1,
		"udp://127.0.0.1:8402/": 2,
		"udp://127.0.0.1:8403/": 3,
		"udp://127.0.0.1:8404/": 4,
	}))
	client.UseService(&proxy)
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			result, err := proxy.Hello(fmt.Sprintf("world %d", i))
			if err == nil {
				assert.Equal(t, fmt.Sprintf("hello world %d", i), result)
			} else {
				assert.Equal(t, core.ErrTimeout, err)
			}
		}(i)
	}
	wg.Wait()
	server1.Close()
	server2.Close()
	server3.Close()
	server4.Close()
}

func TestWeightedLeastActiveLoadBalance(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr1, err := net.ResolveUDPAddr("udp", "127.0.0.1:8401")
	assert.NoError(t, err)
	server1, err := net.ListenUDP("udp", addr1)
	assert.NoError(t, err)
	err = service.Bind(server1)
	assert.NoError(t, err)

	addr2, err := net.ResolveUDPAddr("udp", "127.0.0.1:8402")
	assert.NoError(t, err)
	server2, err := net.ListenUDP("udp", addr2)
	assert.NoError(t, err)
	err = service.Bind(server2)
	assert.NoError(t, err)

	addr3, err := net.ResolveUDPAddr("udp", "127.0.0.1:8403")
	assert.NoError(t, err)
	server3, err := net.ListenUDP("udp", addr3)
	assert.NoError(t, err)
	err = service.Bind(server3)
	assert.NoError(t, err)

	addr4, err := net.ResolveUDPAddr("udp", "127.0.0.1:8404")
	assert.NoError(t, err)
	server4, err := net.ListenUDP("udp", addr4)
	assert.NoError(t, err)
	err = service.Bind(server4)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient()
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.Use(loadbalance.NewWeightedLeastActiveLoadBalance(map[string]int{
		"udp://127.0.0.1:8401/": 1,
		"udp://127.0.0.1:8402/": 2,
		"udp://127.0.0.1:8403/": 3,
		"udp://127.0.0.1:8404/": 4,
	}))
	client.UseService(&proxy)
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			result, err := proxy.Hello(fmt.Sprintf("world %d", i))
			if err == nil {
				assert.Equal(t, fmt.Sprintf("hello world %d", i), result)
			} else {
				assert.Equal(t, core.ErrTimeout, err)
			}
		}(i)
	}
	wg.Wait()
	server1.Close()
	server2.Close()
	server3.Close()
	server4.Close()
}

func TestOneway(t *testing.T) {
	service := core.NewService()
	service.Codec = core.NewServiceCodec(core.WithDebug(true))
	service.AddFunction(func() {
		time.Sleep(time.Millisecond * 50)
	}, "sleep")
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8412")
	assert.NoError(t, err)
	server, err := net.ListenUDP("udp", addr)
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient("udp://127.0.0.1/")
	client.Use(log.Plugin)
	var proxy struct {
		Sleep func() `context:"oneway"`
	}
	client.UseService(&proxy)
	start := time.Now()
	proxy.Sleep()
	duration := time.Since(start)
	assert.True(t, duration > time.Millisecond*40 && duration < time.Millisecond*60)
	client.Use(oneway.Oneway{})
	start = time.Now()
	proxy.Sleep()
	duration = time.Since(start)
	assert.True(t, duration < time.Millisecond*10)
	time.Sleep(time.Millisecond * 60)
	server.Close()
}

func TestClientAbort(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8412")
	assert.NoError(t, err)
	server, err := net.ListenUDP("udp", addr)
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient("udp://127.0.0.1/")
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.UseService(&proxy)
	client.Use(limiter.NewConcurrentLimiter(72))
	client.Use(limiter.NewRateLimiter(5000).InvokeHandler)
	n := int32(0)
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(i int) {
			defer wg.Done()
			result, err := proxy.Hello(fmt.Sprintf("world %d", i))
			if err == nil {
				atomic.AddInt32(&n, 1)
				assert.Equal(t, fmt.Sprintf("hello world %d", i), result)
			}
		}(i)
	}
	client.Abort()
	wg.Wait()
	assert.Greater(t, n, int32(0))
	server.Close()
}

func TestRobustness(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8412")
	assert.NoError(t, err)
	server, err := net.ListenUDP("udp", addr)
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	httpClient := core.NewClient("http://127.0.0.1:8412/")
	tcpClient := core.NewClient("tcp://127.0.0.1/")
	udpClient := core.NewClient("udp://127.0.0.1/")
	var proxy1, proxy2, proxy3 struct {
		Hello func(name string) (string, error)
	}
	httpClient.UseService(&proxy1)
	tcpClient.UseService(&proxy2)
	udpClient.UseService(&proxy3)
	for i := 0; i < 100; i++ {
		result, err := proxy1.Hello("world")
		assert.Equal(t, "", result)
		assert.Error(t, err)
		result, err = proxy2.Hello("world")
		assert.Equal(t, "", result)
		assert.Error(t, err)
		result, err = proxy3.Hello("world")
		assert.Equal(t, "hello world", result)
		assert.NoError(t, err)
	}
	server.Close()
}
