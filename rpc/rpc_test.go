/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/rpc_test.go                                          |
|                                                          |
| LastModified: May 12, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package rpc_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"net/http"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/hprose/hprose-golang/v3/encoding"
	"github.com/hprose/hprose-golang/v3/rpc"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/log"
	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name     string
	Birthday time.Time
}

type Student struct {
	ID       int
	Name     string
	Birthday time.Time
	Grade    int
	Class    int
}

type StudentService struct {
	students []Student
	lock     sync.RWMutex
}

func (ss *StudentService) Add(s ...Student) {
	ss.lock.Lock()
	defer ss.lock.Unlock()
	ss.students = append(ss.students, s...)
}

func (ss *StudentService) Get(id int) *Student {
	ss.lock.RLock()
	defer ss.lock.RUnlock()
	for _, s := range ss.students {
		if s.ID == id {
			return &s
		}
	}
	return nil
}

func (ss *StudentService) Delete(id int) {
	ss.lock.Lock()
	defer ss.lock.Unlock()
	for i := len(ss.students) - 1; i >= 0; i-- {
		if ss.students[i].ID == id {
			ss.students = append(ss.students[:i], ss.students[i+1:]...)
		}
	}
}

func TestAddInstanceMethods(t *testing.T) {
	service := rpc.NewService()
	service.Codec = rpc.NewServiceCodec(rpc.WithDebug(true))
	service.AddInstanceMethods(&StudentService{})
	server, err := net.Listen("tcp", "127.0.0.1:8412")
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := rpc.NewClient("tcp://127.0.0.1/")
	client.Use(log.Plugin)
	var proxy struct {
		Add        func(s ...Student) error
		Get        func(id int) (*Student, error)
		GetStudent func(id int) Student `name:"get"`
		GetPerson  func(id int) Person  `name:"get"`
		Delete     func(id int) error
	}
	client.UseService(&proxy)
	s1 := Student{
		ID:       1,
		Name:     "张三",
		Birthday: time.Date(2008, 11, 23, 0, 0, 0, 0, time.Local),
		Grade:    6,
		Class:    1,
	}
	s2 := Student{
		ID:       2,
		Name:     "李四",
		Birthday: time.Date(2013, 12, 11, 0, 0, 0, 0, time.Local),
		Grade:    1,
		Class:    2,
	}
	err = proxy.Add(s1, s2)
	assert.NoError(t, err)
	var student *Student
	student, err = proxy.Get(1)
	assert.Equal(t, s1, *student)
	assert.NoError(t, err)
	student, err = proxy.Get(2)
	assert.Equal(t, s2, *student)
	assert.NoError(t, err)
	err = proxy.Delete(2)
	assert.NoError(t, err)
	student, err = proxy.Get(2)
	assert.Nil(t, student)
	assert.NoError(t, err)
	s := proxy.GetStudent(1)
	assert.Equal(t, s1, s)
	p := proxy.GetPerson(1)
	assert.Equal(t, "张三", p.Name)
	assert.Equal(t, time.Date(2008, 11, 23, 0, 0, 0, 0, time.Local), p.Birthday)
	server.Close()
}

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func TestAddNetRPCMethods(t *testing.T) {
	service := rpc.NewService()
	service.Codec = rpc.NewServiceCodec(rpc.WithDebug(true))
	service.AddNetRPCMethods(new(Arith), "Arith")
	server, err := net.Listen("tcp", "127.0.0.1:8412")
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := rpc.NewClient("tcp://127.0.0.1/")
	client.Use(log.Plugin)
	var proxy struct {
		Multiply func(args Args) (int, error)
		Divide   func(args Args) (Quotient, error)
	}
	client.UseService(&proxy, "Arith")
	{
		result, err := proxy.Multiply(Args{3, 2})
		assert.Equal(t, 6, result)
		assert.NoError(t, err)
	}
	{
		result, err := proxy.Divide(Args{3, 2})
		assert.Equal(t, Quotient{1, 1}, result)
		assert.NoError(t, err)
	}
	{
		_, err := proxy.Divide(Args{3, 0})
		assert.EqualError(t, err, "divide by zero")
	}
	server.Close()
}

type innerService int

func (s innerService) Sum(n ...int) (sum int) {
	for _, i := range n {
		sum += i
	}
	return
}

type outerService struct {
	innerService
	Inner innerService
	Sub   func(x int, n ...int) int
}

func newOuterService() *outerService {
	return &outerService{
		Sub: func(x int, n ...int) int {
			for _, i := range n {
				x -= i
			}
			return x
		},
	}
}

func TestAddAllMethods(t *testing.T) {
	service := rpc.NewService()
	service.Codec = rpc.NewServiceCodec(rpc.WithDebug(true))
	service.AddInstanceMethods(newOuterService(), "s1")
	service.AddAllMethods(newOuterService(), "s2")
	server, err := net.Listen("tcp", "127.0.0.1:8412")
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := rpc.NewClient("tcp://127.0.0.1/")
	client.Use(log.Plugin)
	var proxy1, proxy2 struct {
		Sum   func(n ...int) (int, error)
		Sub   func(x int, n ...int) (int, error)
		Inner struct {
			Sum func(n ...int) (int, error)
		}
	}
	client.UseService(&proxy1, "s1")
	client.UseService(&proxy2, "s2")
	{
		result, err := proxy1.Sum(1, 2, 3, 4, 5)
		assert.Equal(t, 15, result)
		assert.NoError(t, err)
	}
	{
		result, err := proxy2.Sum(1, 2, 3, 4, 5)
		assert.Equal(t, 15, result)
		assert.NoError(t, err)
	}
	{
		result, err := proxy1.Inner.Sum(1, 2, 3, 4, 5)
		assert.Equal(t, 0, result)
		assert.Error(t, err)
	}
	{
		result, err := proxy2.Inner.Sum(1, 2, 3, 4, 5)
		assert.Equal(t, 15, result)
		assert.NoError(t, err)
	}
	{
		result, err := proxy1.Sub(15, 1, 2, 3, 4)
		assert.Equal(t, 5, result)
		assert.NoError(t, err)
	}
	{
		result, err := proxy2.Sub(15, 1, 2, 3, 4)
		assert.Equal(t, 5, result)
		assert.NoError(t, err)
	}
	server.Close()
}

func TestAddMethods(t *testing.T) {
	service := rpc.NewService()
	service.Codec = rpc.NewServiceCodec(rpc.WithDebug(true))
	service.AddMethods([]string{"Sum", "Sub"}, newOuterService(), "s1")
	service.AddMethod("Sum", newOuterService(), "add")
	server, err := net.Listen("tcp", "127.0.0.1:8412")
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := rpc.NewClient("tcp://127.0.0.1/")
	client.Use(log.Plugin)
	var proxy1, proxy2 struct {
		Sum func(n ...int) (int, error)
		Sub func(x int, n ...int) (int, error)
		Add func(n ...int) (int, error)
	}
	client.UseService(&proxy1, "s1")
	client.UseService(&proxy2)
	{
		result, err := proxy1.Sum(1, 2, 3, 4, 5)
		assert.Equal(t, 15, result)
		assert.NoError(t, err)
	}
	{
		result, err := proxy2.Sum(1, 2, 3, 4, 5)
		assert.Equal(t, 0, result)
		assert.Error(t, err)
	}
	{
		result, err := proxy1.Add(1, 2, 3, 4, 5)
		assert.Equal(t, 0, result)
		assert.Error(t, err)
	}
	{
		result, err := proxy2.Add(1, 2, 3, 4, 5)
		assert.Equal(t, 15, result)
		assert.NoError(t, err)
	}
	{
		result, err := proxy1.Sub(15, 1, 2, 3, 4)
		assert.Equal(t, 5, result)
		assert.NoError(t, err)
	}
	{
		result, err := proxy2.Sub(15, 1, 2, 3, 4)
		assert.Equal(t, 0, result)
		assert.Error(t, err)
	}
	server.Close()
}

func TestAddFunction(t *testing.T) {
	service := rpc.NewService()
	service.Codec = rpc.NewServiceCodec(rpc.WithDebug(true))
	value := reflect.ValueOf(newOuterService()).Elem()
	sub := value.FieldByName("Sub")
	sum, _ := value.Type().MethodByName("Sum")
	service.AddFunction(sub, "Sub")
	service.AddFunction(&sub, "ptr_sub")
	service.AddFunction(sum, "Sum")
	service.AddFunction(&sum, "ptr_sum")
	assert.Equal(t, 5, len(service.Names()))
	assert.Equal(t, 5, len(service.Methods()))
	server, err := net.Listen("tcp", "127.0.0.1:8412")
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := rpc.NewClient("tcp://127.0.0.1/")
	client.Use(log.Plugin)
	var proxy1, proxy2 struct {
		Sum func(_ interface{}, n ...int) (int, error)
		Sub func(x int, n ...int) (int, error)
	}
	client.UseService(&proxy1)
	client.UseService(&proxy2, "ptr")
	{
		result, err := proxy1.Sum(newOuterService(), 1, 2, 3, 4, 5)
		assert.Equal(t, 15, result)
		assert.NoError(t, err)
	}
	{
		result, err := proxy2.Sum(newOuterService(), 1, 2, 3, 4, 5)
		assert.Equal(t, 15, result)
		assert.NoError(t, err)
	}
	{
		result, err := proxy1.Sub(15, 1, 2, 3, 4)
		assert.Equal(t, 5, result)
		assert.NoError(t, err)
	}
	{
		result, err := proxy2.Sub(15, 1, 2, 3, 4)
		assert.Equal(t, 5, result)
		assert.NoError(t, err)
	}
	server.Close()
}

func autoTypeConvert(a interface{}) (string, interface{}) {
	switch a := a.(type) {
	case *big.Int:
		return "auto convert to *big.Int", a.Int64() + 1
	case *big.Float:
		return "auto convert to *big.Float", a.String()
	case map[string]interface{}:
		return "auto convert to map[string]interface{}", a["test"]
	default:
		return "", nil
	}
}

func TestAutoTypeConvert(t *testing.T) {
	service := rpc.NewService()
	service.Codec = rpc.NewServiceCodec(
		rpc.WithDebug(true),
		rpc.WithSimple(true),
		rpc.WithLongType(encoding.LongTypeBigInt),
		rpc.WithRealType(encoding.RealTypeBigFloat),
		rpc.WithMapType(encoding.MapTypeSIMap),
	)
	service.AddFunction(autoTypeConvert)
	server, err := net.Listen("tcp", "127.0.0.1:8412")
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := rpc.NewClient("tcp://127.0.0.1/")
	client.Use(log.Plugin)
	var proxy struct {
		AutoTypeConvert func(a interface{}) (string, interface{})
	}
	client.Codec = rpc.NewClientCodec(
		rpc.WithSimple(true),
		rpc.WithLongType(encoding.LongTypeUint64),
		rpc.WithRealType(encoding.RealTypeFloat64),
		rpc.WithMapType(encoding.MapTypeIIMap),
	)
	client.UseService(&proxy)
	msg, result := proxy.AutoTypeConvert(int64(12345))
	assert.Equal(t, "auto convert to *big.Int", msg)
	assert.Equal(t, uint64(12346), result)
	msg, result = proxy.AutoTypeConvert(float64(12345))
	assert.Equal(t, "auto convert to *big.Float", msg)
	assert.Equal(t, "12345", result)
	msg, result = proxy.AutoTypeConvert(map[interface{}]interface{}{"test": "test"})
	assert.Equal(t, "auto convert to map[string]interface{}", msg)
	assert.Equal(t, "test", result)
	server.Close()
}

func TestHTTP(t *testing.T) {
	crossDomainXMLContent := `<?xml version="1.0"?>
	<!DOCTYPE cross-domain-policy SYSTEM "http://www.adobe.com/xml/dtds/cross-domain-policy.dtd">
	<cross-domain-policy>
		<site-control permitted-cross-domain-policies="master-only"/>
		<allow-access-from domain="*.hprose.com"/>
	</cross-domain-policy>`
	clientAccessPolicyXMLContent := `<?xml version="1.0" encoding="utf-8" ?>
	<access-policy>
	  <cross-domain-access>
		<policy>
		  <allow-from http-request-headers="*">
			<domain uri="*"/>
		  </allow-from>
		  <grant-to>
			<resource path="/" include-subpaths="true"/>
		  </grant-to>
		</policy>
	  </cross-domain-access>
	</access-policy>`
	service := rpc.NewService()
	service.Codec = rpc.NewServiceCodec(rpc.WithDebug(true))
	httpHandler := rpc.HTTPHandler(service)
	httpHandler.OnError = func(response http.ResponseWriter, request *http.Request, err error) {
		fmt.Println(err)
	}
	httpHandler.AddAccessControlAllowOrigin("www.google.com", "www.baidu.com", "hprose.com")
	httpHandler.RemoveAccessControlAllowOrigin("www.baidu.com")
	assert.True(t, httpHandler.AccessControlAllowOrigins["www.google.com"])
	assert.True(t, httpHandler.AccessControlAllowOrigins["hprose.com"])
	assert.False(t, httpHandler.AccessControlAllowOrigins["www.baidu.com"])
	httpHandler.SetCrossDomainXMLContent([]byte(crossDomainXMLContent))
	httpHandler.SetClientAccessPolicyXMLContent([]byte(clientAccessPolicyXMLContent))
	assert.Equal(t, crossDomainXMLContent, string(httpHandler.CrossDomainXMLContent()))
	assert.Equal(t, clientAccessPolicyXMLContent, string(httpHandler.ClientAccessPolicyXMLContent()))
	httpHandler.SetCrossDomainXMLFile("")
	assert.Equal(t, "", string(httpHandler.CrossDomainXMLFile()))
	assert.Equal(t, "", string(httpHandler.CrossDomainXMLContent()))
	httpHandler.SetClientAccessPolicyXMLFile("")
	assert.Equal(t, "", string(httpHandler.ClientAccessPolicyXMLFile()))
	assert.Equal(t, "", string(httpHandler.ClientAccessPolicyXMLContent()))
	httpHandler.SetCrossDomainXMLContent([]byte(crossDomainXMLContent))
	httpHandler.SetClientAccessPolicyXMLContent([]byte(clientAccessPolicyXMLContent))
	service.AddFunction(func(ctx context.Context, name string) string {
		serviceContext := rpc.GetServiceContext(ctx)
		header, _ := serviceContext.Items().Get("httpRequestHeaders")
		return header.(http.Header).Get("test") + ":hello " + name
	}, "hello")
	assert.True(t, service.Get("hello").PassContext())
	serverMux := http.NewServeMux()
	serverMux.Handle("/test1", httpHandler)
	serverMux.Handle("/test2", httpHandler)
	serverMux.Handle("/crossdomain.xml", httpHandler)
	serverMux.Handle("/clientaccesspolicy.xml", httpHandler)
	server := &http.Server{Addr: ":8000", Handler: serverMux}
	go server.ListenAndServe()

	time.Sleep(time.Millisecond * 5)

	client := rpc.NewClient("http://127.0.0.1:8000/test1", "http://127.0.0.1:8000/test2")
	client.ShuffleURLs()
	client.Use(log.Plugin)
	httpTransport := rpc.HTTPTransport(client)
	httpTransport.Header = http.Header{
		"test":   []string{"test"},
		"Origin": []string{"hprose.com"},
	}
	resp, err := httpTransport.HTTPClient.Get("http://127.0.0.1:8000/crossdomain.xml")
	assert.NoError(t, err)
	content, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	assert.Equal(t, crossDomainXMLContent, string(content))
	assert.NoError(t, err)
	resp, err = httpTransport.HTTPClient.Get("http://127.0.0.1:8000/clientaccesspolicy.xml")
	lastModified := resp.Header.Get("Last-Modified")
	etag := resp.Header.Get("Etag")
	assert.NoError(t, err)
	content, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	assert.Equal(t, clientAccessPolicyXMLContent, string(content))
	assert.NoError(t, err)
	req, _ := http.NewRequest("GET", "http://127.0.0.1:8000/clientaccesspolicy.xml", nil)
	req.Header.Set("if-modified-since", lastModified)
	req.Header.Set("if-none-match", etag)
	resp, err = httpTransport.HTTPClient.Do(req)
	assert.Equal(t, http.StatusNotModified, resp.StatusCode)
	resp.Body.Close()
	assert.NoError(t, err)
	resp, err = httpTransport.HTTPClient.Get("http://127.0.0.1:8000/test1")
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp.Body.Close()
	assert.NoError(t, err)
	httpHandler.GET = false
	resp, err = httpTransport.HTTPClient.Get("http://127.0.0.1:8000/test1")
	assert.Equal(t, http.StatusForbidden, resp.StatusCode)
	resp.Body.Close()
	assert.NoError(t, err)
	{
		result, err := client.Invoke("hello", []interface{}{"world"})
		assert.Equal(t, "test:hello world", result[0])
		assert.NoError(t, err)
	}
	var proxy struct {
		Hello func(name string) (string, error)
	}
	client.Timeout = 0
	client.UseService(&proxy)
	result, err := proxy.Hello("world")
	assert.Equal(t, "test:hello world", result)
	assert.NoError(t, err)
	client.SetURI("http://127.0.0.1:8000/")
	_, err = proxy.Hello("world")
	assert.Equal(t, errors.New("404 Not Found"), err)
	server.Close()
}

func TestTCP(t *testing.T) {
	service := rpc.NewService()
	service.AddMissingMethod(func(ctx context.Context, name string, args []interface{}) (result []interface{}, err error) {
		serviceContext := rpc.GetServiceContext(ctx)
		data, err := json.Marshal(args)
		if err != nil {
			return nil, err
		}
		return []interface{}{name + string(data) + serviceContext.RemoteAddr.String()}, nil
	})
	method := service.Get("*")
	assert.Equal(t, reflect.Func, method.Func().Kind())
	assert.Equal(t, []reflect.Type{reflect.TypeOf(""), reflect.TypeOf([]interface{}{})}, method.Parameters())
	assert.True(t, method.ReturnError())
	assert.Nil(t, method.Options())
	socketHandler := rpc.SocketHandler(service)
	socketHandler.OnAccept = func(c net.Conn) net.Conn {
		fmt.Println(c.RemoteAddr().String() + "->" + c.LocalAddr().String() + " accepted")
		return c
	}
	socketHandler.OnClose = func(c net.Conn) {
		fmt.Println(c.RemoteAddr().String() + "->" + c.LocalAddr().String() + " closed on server")
	}
	socketHandler.OnError = func(c net.Conn, e error) {
		if c != nil {
			fmt.Println(c.RemoteAddr().String()+"->"+c.LocalAddr().String(), e)
		} else {
			fmt.Println(e)
		}
	}
	server, err := net.Listen("tcp", "127.0.0.1:8412")
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := rpc.NewClient("tcp://127.0.0.1/")
	socketTransport := rpc.SocketTransport(client)
	socketTransport.OnConnect = func(c net.Conn) net.Conn {
		fmt.Println(c.LocalAddr().String() + "->" + c.RemoteAddr().String() + " connected")
		return c
	}
	socketTransport.OnClose = func(c net.Conn) {
		fmt.Println(c.LocalAddr().String() + "->" + c.RemoteAddr().String() + " closed on client")
	}
	client.Use(log.Plugin)
	var proxy struct {
		Hello func(name string) string
	}
	client.UseService(&proxy)
	proxy.Hello("world")
	server.Close()
}

func TestUnix(t *testing.T) {
	service := rpc.NewService()
	service.AddMissingMethod(func(name string, args []interface{}) (result []interface{}, err error) {
		data, err := json.Marshal(args)
		if err != nil {
			return nil, err
		}
		return []interface{}{name + string(data)}, nil
	})
	method := service.Get("*")
	assert.Equal(t, reflect.Func, method.Func().Kind())
	assert.Equal(t, []reflect.Type{reflect.TypeOf(""), reflect.TypeOf([]interface{}{})}, method.Parameters())
	assert.True(t, method.ReturnError())
	assert.Nil(t, method.Options())
	socketHandler := rpc.SocketHandler(service)
	socketHandler.OnAccept = func(c net.Conn) net.Conn {
		fmt.Println(c.RemoteAddr().String() + "->" + c.LocalAddr().String() + " accepted")
		return c
	}
	socketHandler.OnClose = func(c net.Conn) {
		fmt.Println(c.RemoteAddr().String() + "->" + c.LocalAddr().String() + " closed on server")
	}
	socketHandler.OnError = func(c net.Conn, e error) {
		if c != nil {
			fmt.Println(c.RemoteAddr().String()+"->"+c.LocalAddr().String(), e)
		} else {
			fmt.Println(e)
		}
	}
	server, err := net.Listen("unix", "/tmp/hprose_test.sock")
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := rpc.NewClient("unix://1/tmp/hprose_test.sock")
	socketTransport := rpc.SocketTransport(client)
	socketTransport.OnConnect = func(c net.Conn) net.Conn {
		fmt.Println(c.LocalAddr().String() + "->" + c.RemoteAddr().String() + " connected")
		return c
	}
	socketTransport.OnClose = func(c net.Conn) {
		fmt.Println(c.LocalAddr().String() + "->" + c.RemoteAddr().String() + " closed on client")
	}
	client.Use(log.Plugin)
	var proxy struct {
		Hello func(name string) string
	}
	client.UseService(&proxy)
	proxy.Hello("world")
	server.Close()
}

func TestUDP(t *testing.T) {
	service := rpc.NewService()
	udpHandler := rpc.UDPHandler(service)
	udpHandler.OnClose = func(c net.Conn) {
		fmt.Println(c.LocalAddr().String() + " closed on server")
	}
	udpHandler.OnError = func(c net.Conn, e error) {
		if c != nil {
			fmt.Println(c.LocalAddr().String(), e)
		} else {
			fmt.Println(e)
		}
	}
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8412")
	assert.NoError(t, err)
	server, err := net.ListenUDP("udp", addr)
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 5)

	client := rpc.NewClient("udp://127.0.0.1/")
	udpTransport := rpc.UDPTransport(client)
	udpTransport.OnConnect = func(c net.Conn) net.Conn {
		fmt.Println(c.LocalAddr().String() + "->" + c.RemoteAddr().String() + " connected")
		return c
	}
	udpTransport.OnClose = func(c net.Conn) {
		fmt.Println(c.LocalAddr().String() + "->" + c.RemoteAddr().String() + " closed on client")
	}
	client.Use(log.Plugin)
	var proxy struct {
		Hello func(name string) string
	}
	client.UseService(&proxy)
	proxy.Hello("world")
	client.Abort()
	server.Close()
}

func TestWebSocket(t *testing.T) {
	service := rpc.NewService()
	service.AddMissingMethod(func(ctx context.Context, name string, args []interface{}) (result []interface{}, err error) {
		serviceContext := rpc.GetServiceContext(ctx)
		data, err := json.Marshal(args)
		if err != nil {
			return nil, err
		}
		return []interface{}{name + string(data) + serviceContext.RemoteAddr.String()}, nil
	})
	method := service.Get("*")
	assert.Equal(t, reflect.Func, method.Func().Kind())
	assert.Equal(t, []reflect.Type{reflect.TypeOf(""), reflect.TypeOf([]interface{}{})}, method.Parameters())
	assert.True(t, method.ReturnError())
	assert.Nil(t, method.Options())
	webSocketHandler := rpc.WebSocketHandler(service)
	webSocketHandler.OnAccept = func(c *websocket.Conn) *websocket.Conn {
		fmt.Println(c.RemoteAddr().String() + "->" + c.LocalAddr().String() + " accepted")
		return c
	}
	webSocketHandler.OnClose = func(c *websocket.Conn) {
		fmt.Println(c.RemoteAddr().String() + "->" + c.LocalAddr().String() + " closed on server")
	}
	webSocketHandler.OnError = func(c *websocket.Conn, e error) {
		if c != nil {
			fmt.Println(c.RemoteAddr().String()+"->"+c.LocalAddr().String(), e)
		} else {
			fmt.Println(e)
		}
	}
	server := &http.Server{Addr: ":8005"}
	err := service.Bind(server)
	assert.NoError(t, err)
	go server.ListenAndServe()

	time.Sleep(time.Millisecond * 5)

	client := rpc.NewClient("ws://127.0.0.1:8005/")
	webSocketTransport := rpc.WebSocketTransport(client)
	webSocketTransport.OnConnect = func(c *websocket.Conn) *websocket.Conn {
		fmt.Println(c.LocalAddr().String() + "->" + c.RemoteAddr().String() + " connected")
		return c
	}
	webSocketTransport.OnClose = func(c *websocket.Conn) {
		fmt.Println(c.LocalAddr().String() + "->" + c.RemoteAddr().String() + " closed on client")
	}
	client.Use(log.Plugin)
	var proxy struct {
		Hello func(name string) string
	}
	client.UseService(&proxy)
	proxy.Hello("world")
	client.Abort()

	client = rpc.NewClient("http://127.0.0.1:8005/")
	client.Use(log.Plugin)
	client.UseService(&proxy)
	proxy.Hello("world")

	server.Close()
}
