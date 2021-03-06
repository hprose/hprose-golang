/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/mock/mock_test.go                                    |
|                                                          |
| LastModified: Mar 6, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package mock

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/hprose/hprose-golang/v3/rpc/core"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/circuitbreaker"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/cluster"
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
	server := Server{"testClientTimeout"}
	err := service.Bind(server)
	assert.NoError(t, err)
	client := core.NewClient("mock://testClientTimeout")
	client.Use(log.Plugin)
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
	service.Use(timeout.GetHandler(time.Millisecond))
	server := Server{"testServiceTimeout"}
	err := service.Bind(server)
	assert.NoError(t, err)
	client := core.NewClient("mock://testServiceTimeout")
	var proxy struct {
		Wait func(d time.Duration) error
	}
	client.UseService(&proxy)
	client.Use(log.IOHandler, log.InvokeHandler)
	err = proxy.Wait(time.Second * 30)
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
	server := Server{"testMissingMethod"}
	err := service.Bind(server)
	assert.NoError(t, err)
	client := core.NewClient("mock://testMissingMethod")
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
		return []interface{}{name + string(data) + serviceContext.RemoteAddr.String()}, nil
	})
	server := Server{"testMissingMethod2"}
	err := service.Bind(server)
	assert.NoError(t, err)
	client := core.NewClient("mock://testMissingMethod2")
	client.Use(log.Plugin)
	var proxy struct {
		Hello func(name string) string
	}
	client.UseService(&proxy)
	result := proxy.Hello("world")
	assert.Equal(t, `Hello["world"]testMissingMethod2`, result)
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
	server := Server{"testHeaders"}
	err := service.Bind(server)
	assert.NoError(t, err)
	client := core.NewClient("mock://testHeaders")
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
	server := Server{"testMaxRequestLength"}
	err := service.Bind(server)
	assert.NoError(t, err)
	client := core.NewClient("mock://testMaxRequestLength")
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
	server := Server{"testCircuitBreaker"}
	err := service.Bind(server)
	assert.NoError(t, err)
	client := core.NewClient("mock://testCircuitBreaker")
	client.Use(circuitbreaker.New(circuitbreaker.WithThreshold(3)))
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.UseService(&proxy)
	result, err := proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}
	server.Close()
	for i := 0; i < 4; i++ {
		_, err = proxy.Hello("world")
		if assert.Error(t, err) {
			assert.Equal(t, "hprose/rpc/core: server is stoped", err.Error())
		}
	}
	_, err = proxy.Hello("world")
	if assert.Error(t, err) {
		assert.Equal(t, "service breaked", err.Error())
	}
	_ = service.Bind(server)
	_, err = proxy.Hello("world")
	if assert.Error(t, err) {
		assert.Equal(t, "service breaked", err.Error())
	}
	server.Close()
}

func TestCircuitBreaker2(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	server := Server{"testCircuitBreaker2"}
	err := service.Bind(server)
	assert.NoError(t, err)
	client := core.NewClient("mock://testCircuitBreaker2")
	client.Use(circuitbreaker.New(
		circuitbreaker.WithThreshold(1),
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
	_, err = proxy.Hello("world")
	if assert.Error(t, err) {
		assert.Equal(t, "hprose/rpc/core: server is stoped", err.Error())
	}
	_, err = proxy.Hello("world")
	if assert.Error(t, err) {
		assert.Equal(t, "hprose/rpc/core: server is stoped", err.Error())
	}
	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "Hello breaked", result)
	}
	_ = service.Bind(server)
	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "Hello breaked", result)
	}
	server.Close()
}

func TestClusterFailover1(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	server1 := Server{"testClusterFailover1"}
	err := service.Bind(server1)
	assert.NoError(t, err)

	server2 := Server{"testClusterFailover2"}
	err = service.Bind(server2)
	assert.NoError(t, err)

	server3 := Server{"testClusterFailover3"}
	err = service.Bind(server3)
	assert.NoError(t, err)

	server4 := Server{"testClusterFailover4"}
	err = service.Bind(server4)
	assert.NoError(t, err)

	client := core.NewClient(
		"mock://testClusterFailover1",
		"mock://testClusterFailover2",
		"mock://testClusterFailover3",
		"mock://testClusterFailover4",
	)
	client.Use(cluster.New(
		cluster.FailoverConfig(),
	).Handler).Use(log.Plugin)
	var proxy struct {
		Hello func(name string) (string, error) `context:"idempotent,retry:3"`
	}
	client.UseService(&proxy)
	result, err := proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server1.Close()

	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server2.Close()

	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server3.Close()

	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server4.Close()

	client.UseService(&proxy)
	_, err = proxy.Hello("world")
	assert.Error(t, err)
}

func TestClusterFailover2(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	server1 := Server{"testClusterFailover1"}
	err := service.Bind(server1)
	assert.NoError(t, err)

	server2 := Server{"testClusterFailover2"}
	err = service.Bind(server2)
	assert.NoError(t, err)

	server3 := Server{"testClusterFailover3"}
	err = service.Bind(server3)
	assert.NoError(t, err)

	server4 := Server{"testClusterFailover4"}
	err = service.Bind(server4)
	assert.NoError(t, err)

	client := core.NewClient(
		"mock://testClusterFailover1",
		"mock://testClusterFailover2",
		"mock://testClusterFailover3",
		"mock://testClusterFailover4",
	)
	client.Use(cluster.New(
		cluster.FailoverConfig(
			cluster.WithIdempotent(true),
			cluster.WithRetry(3),
		),
	).Handler).Use(log.Plugin)
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.UseService(&proxy)
	result, err := proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server1.Close()

	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server2.Close()

	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server3.Close()

	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server4.Close()

	_, err = proxy.Hello("world")
	assert.Error(t, err)
}

func TestClusterFailtry(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	server := Server{"testClusterFailtry"}
	err := service.Bind(server)
	assert.NoError(t, err)

	client := core.NewClient("mock://testClusterFailtry")
	client.Use(cluster.New(
		cluster.FailtryConfig(
			cluster.WithIdempotent(true),
			cluster.WithRetry(3),
		),
	).Handler).Use(log.Plugin)
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.UseService(&proxy)
	result, err := proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server.Close()

	go func() {
		time.Sleep(time.Second)
		service.Bind(server)
	}()

	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server.Close()

	_, err = proxy.Hello("world")
	assert.Error(t, err)
}

func TestClusterFailfast(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	server := Server{"testClusterFailfast"}
	err := service.Bind(server)
	assert.NoError(t, err)

	client := core.NewClient("mock://testClusterFailfast")
	client.Use(cluster.New(
		cluster.FailfastConfig(
			func(c context.Context) {
				fmt.Println("TestClusterFailfast ok")
			},
		),
	).Handler).Use(log.Plugin)
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.UseService(&proxy)
	result, err := proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server.Close()

	go func() {
		time.Sleep(time.Second)
		service.Bind(server)
	}()

	_, err = proxy.Hello("world")
	assert.Error(t, err)

	server.Close()

	_, err = proxy.Hello("world")
	assert.Error(t, err)
}

func TestClusterSuccess(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	server := Server{"testClusterSuccess"}
	err := service.Bind(server)
	assert.NoError(t, err)

	client := core.NewClient("mock://testClusterSuccess")
	client.Use(cluster.New(
		cluster.Config{
			OnSuccess: func(ctx context.Context) {
				fmt.Println("TestClusterSuccess ok")
			},
		},
	).Handler).Use(log.Plugin)
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.UseService(&proxy)
	result, err := proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}
}

func TestClusterForking(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	server1 := Server{"testClusterForking1"}
	err := service.Bind(server1)
	assert.NoError(t, err)
	server2 := Server{"testClusterForking2"}
	err = service.Bind(server2)
	assert.NoError(t, err)
	server3 := Server{"testClusterForking3"}
	err = service.Bind(server3)
	assert.NoError(t, err)
	server4 := Server{"testClusterForking4"}
	err = service.Bind(server4)
	assert.NoError(t, err)

	client := core.NewClient(
		"mock://testClusterForking1",
		"mock://testClusterForking2",
		"mock://testClusterForking3",
		"mock://testClusterForking4",
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

	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server2.Close()

	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server3.Close()

	result, err = proxy.Hello("world")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello world", result)
	}

	server4.Close()

	_, err = proxy.Hello("world")
	assert.Error(t, err)
}

func TestClusterBroadcast(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	server1 := Server{"testClusterBroadcast1"}
	err := service.Bind(server1)
	assert.NoError(t, err)
	server2 := Server{"testClusterBroadcast2"}
	err = service.Bind(server2)
	assert.NoError(t, err)
	server3 := Server{"testClusterBroadcast3"}
	err = service.Bind(server3)
	assert.NoError(t, err)
	server4 := Server{"testClusterBroadcast4"}
	err = service.Bind(server4)
	assert.NoError(t, err)

	client := core.NewClient(
		"mock://testClusterBroadcast1",
		"mock://testClusterBroadcast2",
		"mock://testClusterBroadcast3",
		"mock://testClusterBroadcast4",
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
	result, err = client.InvokeContext(core.WithContext(context.Background(), clientContext), "hello", []interface{}{"world"})
	assert.Error(t, err)
	assert.Equal(t, []interface{}{
		[]interface{}(nil),
		[]interface{}{"hello world"},
		[]interface{}{"hello world"},
		[]interface{}{"hello world"},
	}, result)
	server2.Close()
	result, err = client.InvokeContext(core.WithContext(context.Background(), clientContext), "hello", []interface{}{"world"})
	assert.Error(t, err)
	assert.Equal(t, []interface{}{
		[]interface{}(nil),
		[]interface{}(nil),
		[]interface{}{"hello world"},
		[]interface{}{"hello world"},
	}, result)
	server3.Close()
	result, err = client.InvokeContext(core.WithContext(context.Background(), clientContext), "hello", []interface{}{"world"})
	assert.Error(t, err)
	assert.Equal(t, []interface{}{
		[]interface{}(nil),
		[]interface{}(nil),
		[]interface{}(nil),
		[]interface{}{"hello world"},
	}, result)
	server4.Close()
	result, err = client.InvokeContext(core.WithContext(context.Background(), clientContext), "hello", []interface{}{"world"})
	assert.Error(t, err)
	assert.Equal(t, []interface{}{
		[]interface{}(nil),
		[]interface{}(nil),
		[]interface{}(nil),
		[]interface{}(nil),
	}, result)
}
