/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/mock/mock_test.go                                    |
|                                                          |
| LastModified: Feb 18, 2024                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package mock_test

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/hprose/hprose-golang/v3/io"
	"github.com/hprose/hprose-golang/v3/rpc/core"
	. "github.com/hprose/hprose-golang/v3/rpc/mock"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/circuitbreaker"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/cluster"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/forward"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/limiter"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/loadbalance"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/log"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/oneway"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/timeout"
	"github.com/stretchr/testify/assert"
)

func init() {
	RegisterHandler()
	RegisterTransport()
}

func TestHelloWorld(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	server := Server{Address: "testHelloWorld"}
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
	results, err := client.Invoke("hello", []interface{}{"world"})
	assert.Equal(t, "hello world", results[0])
	assert.NoError(t, err)
	server.Close()
}

func TestClientTimeout(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(d time.Duration) {
		time.Sleep(d)
	}, "wait")
	server := Server{Address: "testClientTimeout"}
	err := service.Bind(server)
	assert.NoError(t, err)
	client := core.NewClient("mock://testClientTimeout")
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
	server := Server{Address: "testServiceTimeout"}
	err := service.Bind(server)
	assert.NoError(t, err)
	client := core.NewClient("mock://testServiceTimeout")
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
	server := Server{Address: "testMissingMethod"}
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
	server := Server{Address: "testMissingMethod2"}
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
	server := Server{Address: "testHeaders"}
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
	server := Server{Address: "testMaxRequestLength"}
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
	server := Server{Address: "testCircuitBreaker"}
	err := service.Bind(server)
	assert.NoError(t, err)
	client := core.NewClient("mock://testCircuitBreaker")
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
		if assert.Error(t, err) {
			assert.Equal(t, "server is stoped", err.Error())
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
	server := Server{Address: "testCircuitBreaker2"}
	err := service.Bind(server)
	assert.NoError(t, err)
	client := core.NewClient("mock://testCircuitBreaker2")
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
	if assert.Error(t, err) {
		assert.Equal(t, "server is stoped", err.Error())
	}
	_, err = proxy.Hello("world")
	if assert.Error(t, err) {
		assert.Equal(t, "server is stoped", err.Error())
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
	server1 := Server{Address: "testClusterFailover1"}
	err := service.Bind(server1)
	assert.NoError(t, err)

	server2 := Server{Address: "testClusterFailover2"}
	err = service.Bind(server2)
	assert.NoError(t, err)

	server3 := Server{Address: "testClusterFailover3"}
	err = service.Bind(server3)
	assert.NoError(t, err)

	server4 := Server{Address: "testClusterFailover4"}
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
	server1 := Server{Address: "testClusterFailover1"}
	err := service.Bind(server1)
	assert.NoError(t, err)

	server2 := Server{Address: "testClusterFailover2"}
	err = service.Bind(server2)
	assert.NoError(t, err)

	server3 := Server{Address: "testClusterFailover3"}
	err = service.Bind(server3)
	assert.NoError(t, err)

	server4 := Server{Address: "testClusterFailover4"}
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
	server := Server{Address: "testClusterFailtry"}
	err := service.Bind(server)
	assert.NoError(t, err)

	client := core.NewClient("mock://testClusterFailtry")
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
		service.Bind(server)
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
	server := Server{Address: "testClusterFailfast"}
	err := service.Bind(server)
	assert.NoError(t, err)

	client := core.NewClient("mock://testClusterFailfast")
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

	go func() {
		time.Sleep(time.Second)
		service.Bind(server)
	}()

	_, err = proxy.Hello("world")
	assert.Error(t, err)

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
	server := Server{Address: "testClusterSuccess"}
	err := service.Bind(server)
	assert.NoError(t, err)

	client := core.NewClient("mock://testClusterSuccess")
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
}

func TestClusterForking(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	server1 := Server{Address: "testClusterForking1"}
	err := service.Bind(server1)
	assert.NoError(t, err)
	server2 := Server{Address: "testClusterForking2"}
	err = service.Bind(server2)
	assert.NoError(t, err)
	server3 := Server{Address: "testClusterForking3"}
	err = service.Bind(server3)
	assert.NoError(t, err)
	server4 := Server{Address: "testClusterForking4"}
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
	server1 := Server{Address: "testClusterBroadcast1"}
	err := service.Bind(server1)
	assert.NoError(t, err)
	server2 := Server{Address: "testClusterBroadcast2"}
	err = service.Bind(server2)
	assert.NoError(t, err)
	server3 := Server{Address: "testClusterBroadcast3"}
	err = service.Bind(server3)
	assert.NoError(t, err)
	server4 := Server{Address: "testClusterBroadcast4"}
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
	server1 := Server{Address: "testForward.RealServer"}
	err := service1.Bind(server1)
	assert.NoError(t, err)

	fw := forward.New("mock://testForward.RealServer")
	fw.Use(log.Plugin)
	service2 := core.NewService()
	service2.AddMissingMethod(func(ctx context.Context, name string, args []interface{}) (result []interface{}, err error) {
		return
	})
	service2.Use(fw.InvokeHandler)
	server2 := Server{Address: "testForward.ForwardServer"}
	err = service2.Bind(server2)
	assert.NoError(t, err)

	client := core.NewClient("mock://testForward.ForwardServer")
	client.Use(log.Plugin)
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.UseService(&proxy)
	result, err := proxy.Hello("invoke forward")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello invoke forward", result)
	}

	service2.Unuse(fw.InvokeHandler)
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
	server := Server{Address: "testConcurrentLimiter"}
	err := service.Bind(server)
	assert.NoError(t, err)
	client := core.NewClient("mock://testConcurrentLimiter")
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
	server := Server{Address: "testConcurrentLimiterWithoutTimeout"}
	err := service.Bind(server)
	assert.NoError(t, err)
	client := core.NewClient("mock://testConcurrentLimiterWithoutTimeout")
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
	server := Server{Address: "testRateLimiterIOHandler"}
	err := service.Bind(server)
	assert.NoError(t, err)
	client := core.NewClient("mock://testRateLimiterIOHandler")
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
	server := Server{Address: "testRateLimiterInvokeHandler"}
	err := service.Bind(server)
	assert.NoError(t, err)
	client := core.NewClient("mock://testRateLimiterInvokeHandler")
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
	server1 := Server{Address: "testRandomLoadBalance1"}
	err := service.Bind(server1)
	assert.NoError(t, err)
	server2 := Server{Address: "testRandomLoadBalance2"}
	err = service.Bind(server2)
	assert.NoError(t, err)
	server3 := Server{Address: "testRandomLoadBalance3"}
	err = service.Bind(server3)
	assert.NoError(t, err)
	server4 := Server{Address: "testRandomLoadBalance4"}
	err = service.Bind(server4)
	assert.NoError(t, err)
	client := core.NewClient(
		"mock://testRandomLoadBalance1",
		"mock://testRandomLoadBalance2",
		"mock://testRandomLoadBalance3",
		"mock://testRandomLoadBalance4",
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
	server1 := Server{Address: "testRoundRobinLoadBalance1"}
	err := service.Bind(server1)
	assert.NoError(t, err)
	server2 := Server{Address: "testRoundRobinLoadBalance2"}
	err = service.Bind(server2)
	assert.NoError(t, err)
	server3 := Server{Address: "testRoundRobinLoadBalance3"}
	err = service.Bind(server3)
	assert.NoError(t, err)
	server4 := Server{Address: "testRoundRobinLoadBalance4"}
	err = service.Bind(server4)
	assert.NoError(t, err)
	client := core.NewClient(
		"mock://testRoundRobinLoadBalance1",
		"mock://testRoundRobinLoadBalance2",
		"mock://testRoundRobinLoadBalance3",
		"mock://testRoundRobinLoadBalance4",
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
	server1 := Server{Address: "testLeastActiveLoadBalance1"}
	err := service.Bind(server1)
	assert.NoError(t, err)
	server2 := Server{Address: "testLeastActiveLoadBalance2"}
	err = service.Bind(server2)
	assert.NoError(t, err)
	server3 := Server{Address: "testLeastActiveLoadBalance3"}
	err = service.Bind(server3)
	assert.NoError(t, err)
	server4 := Server{Address: "testLeastActiveLoadBalance4"}
	err = service.Bind(server4)
	assert.NoError(t, err)
	client := core.NewClient(
		"mock://testLeastActiveLoadBalance1",
		"mock://testLeastActiveLoadBalance2",
		"mock://testLeastActiveLoadBalance3",
		"mock://testLeastActiveLoadBalance4",
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
	server1 := Server{Address: "testWeightedRandomLoadBalance1"}
	err := service.Bind(server1)
	assert.NoError(t, err)
	server2 := Server{Address: "testWeightedRandomLoadBalance2"}
	err = service.Bind(server2)
	assert.NoError(t, err)
	server3 := Server{Address: "testWeightedRandomLoadBalance3"}
	err = service.Bind(server3)
	assert.NoError(t, err)
	server4 := Server{Address: "testWeightedRandomLoadBalance4"}
	err = service.Bind(server4)
	assert.NoError(t, err)
	client := core.NewClient()
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.Use(loadbalance.NewWeightedRandomLoadBalance(map[string]int{
		"mock://testWeightedRandomLoadBalance1": 1,
		"mock://testWeightedRandomLoadBalance2": 2,
		"mock://testWeightedRandomLoadBalance3": 3,
		"mock://testWeightedRandomLoadBalance4": 4,
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
	server1 := Server{Address: "testWeightedRoundRobinLoadBalance1"}
	err := service.Bind(server1)
	assert.NoError(t, err)
	server2 := Server{Address: "testWeightedRoundRobinLoadBalance2"}
	err = service.Bind(server2)
	assert.NoError(t, err)
	server3 := Server{Address: "testWeightedRoundRobinLoadBalance3"}
	err = service.Bind(server3)
	assert.NoError(t, err)
	server4 := Server{Address: "testWeightedRoundRobinLoadBalance4"}
	err = service.Bind(server4)
	assert.NoError(t, err)
	client := core.NewClient()
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.Use(loadbalance.NewWeightedRoundRobinLoadBalance(map[string]int{
		"mock://testWeightedRoundRobinLoadBalance1": 1,
		"mock://testWeightedRoundRobinLoadBalance2": 2,
		"mock://testWeightedRoundRobinLoadBalance3": 3,
		"mock://testWeightedRoundRobinLoadBalance4": 4,
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
	server1 := Server{Address: "testNginxRoundRobinLoadBalance1"}
	err := service.Bind(server1)
	assert.NoError(t, err)
	server2 := Server{Address: "testNginxRoundRobinLoadBalance2"}
	err = service.Bind(server2)
	assert.NoError(t, err)
	server3 := Server{Address: "testNginxRoundRobinLoadBalance3"}
	err = service.Bind(server3)
	assert.NoError(t, err)
	server4 := Server{Address: "testNginxRoundRobinLoadBalance4"}
	err = service.Bind(server4)
	assert.NoError(t, err)
	client := core.NewClient()
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.Use(loadbalance.NewNginxRoundRobinLoadBalance(map[string]int{
		"mock://testNginxRoundRobinLoadBalance1": 1,
		"mock://testNginxRoundRobinLoadBalance2": 2,
		"mock://testNginxRoundRobinLoadBalance3": 3,
		"mock://testNginxRoundRobinLoadBalance4": 4,
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
	server1 := Server{Address: "testWeightedLeastActiveLoadBalance1"}
	err := service.Bind(server1)
	assert.NoError(t, err)
	server2 := Server{Address: "testWeightedLeastActiveLoadBalance2"}
	err = service.Bind(server2)
	assert.NoError(t, err)
	server3 := Server{Address: "testWeightedLeastActiveLoadBalance3"}
	err = service.Bind(server3)
	assert.NoError(t, err)
	server4 := Server{Address: "testWeightedLeastActiveLoadBalance4"}
	err = service.Bind(server4)
	assert.NoError(t, err)
	client := core.NewClient()
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.Use(loadbalance.NewWeightedLeastActiveLoadBalance(map[string]int{
		"mock://testWeightedLeastActiveLoadBalance1": 1,
		"mock://testWeightedLeastActiveLoadBalance2": 2,
		"mock://testWeightedLeastActiveLoadBalance3": 3,
		"mock://testWeightedLeastActiveLoadBalance4": 4,
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
		time.Sleep(time.Millisecond * 100)
	}, "sleep")
	server := Server{Address: "testOneway"}
	err := service.Bind(server)
	assert.NoError(t, err)
	client := core.NewClient("mock://testOneway")
	client.Use(log.Plugin)
	var proxy struct {
		Sleep func() `context:"oneway"`
	}
	client.UseService(&proxy)
	start := time.Now()
	proxy.Sleep()
	duration := time.Since(start)
	assert.True(t, duration > time.Millisecond*90 && duration < time.Millisecond*110)
	client.Use(oneway.Oneway{})
	start = time.Now()
	proxy.Sleep()
	duration = time.Since(start)
	assert.True(t, duration < time.Millisecond*10)
	server.Close()
}

func TestClientAbort(t *testing.T) {
	service := core.NewService()
	service.AddFunction(func(name string) string {
		return "hello " + name
	}, "hello")
	server := Server{Address: "testClientAbort"}
	err := service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := core.NewClient("mock://testClientAbort/")
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.UseService(&proxy)
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

func TestReturnStructSlice(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	io.RegisterName("User", (*User)(nil))
	service := core.NewService()
	service.AddFunction(func() []User {
		return []User{{"Tom", 18}, {"Jerry", 20}}
	}, "users")
	server := Server{Address: "testReturnStructSlice"}
	err := service.Bind(server)
	assert.NoError(t, err)
	client := core.NewClient("mock://testReturnStructSlice")
	client.Codec = core.NewClientCodec(
		core.WithStructType(io.StructTypeValue),
		core.WithListType(io.ListTypeSlice),
	)
	var proxy struct {
		Users func() ([]User, error)
	}
	client.UseService(&proxy)
	result, err := proxy.Users()
	if assert.NoError(t, err) {
		assert.Equal(t, []User{{"Tom", 18}, {"Jerry", 20}}, result)
	}
	results, err := client.Invoke("users", nil)
	if assert.NoError(t, err) {
		assert.Equal(t, []User{{"Tom", 18}, {"Jerry", 20}}, results[0])
	}
	server.Close()
}
