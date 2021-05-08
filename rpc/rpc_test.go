/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/rpc_test.go                                          |
|                                                          |
| LastModified: May 7, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package rpc_test

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/hprose/hprose-golang/v3/rpc"
	"github.com/hprose/hprose-golang/v3/rpc/core"
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
	service.Codec = core.NewServiceCodec(core.WithDebug(true))
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
	service.Codec = core.NewServiceCodec(core.WithDebug(true))
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
	service.Codec = core.NewServiceCodec(core.WithDebug(true))
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
	service.Codec = core.NewServiceCodec(core.WithDebug(true))
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
	service.Codec = core.NewServiceCodec(core.WithDebug(true))
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
	service.Codec = core.NewServiceCodec(core.WithDebug(true))
	service.HTTP().OnError = func(err error) {
		fmt.Println(err)
	}
	service.HTTP().AddAccessControlAllowOrigin("www.google.com", "www.baidu.com", "hprose.com")
	service.HTTP().RemoveAccessControlAllowOrigin("www.baidu.com")
	assert.True(t, service.HTTP().AccessControlAllowOrigins["www.google.com"])
	assert.True(t, service.HTTP().AccessControlAllowOrigins["hprose.com"])
	assert.False(t, service.HTTP().AccessControlAllowOrigins["www.baidu.com"])
	service.HTTP().SetCrossDomainXMLContent([]byte(crossDomainXMLContent))
	service.HTTP().SetClientAccessPolicyXMLContent([]byte(clientAccessPolicyXMLContent))
	assert.Equal(t, crossDomainXMLContent, string(service.HTTP().CrossDomainXMLContent()))
	assert.Equal(t, clientAccessPolicyXMLContent, string(service.HTTP().ClientAccessPolicyXMLContent()))
	service.HTTP().SetCrossDomainXMLFile("")
	assert.Equal(t, "", string(service.HTTP().CrossDomainXMLFile()))
	assert.Equal(t, "", string(service.HTTP().CrossDomainXMLContent()))
	service.HTTP().SetClientAccessPolicyXMLFile("")
	assert.Equal(t, "", string(service.HTTP().ClientAccessPolicyXMLFile()))
	assert.Equal(t, "", string(service.HTTP().ClientAccessPolicyXMLContent()))
	service.HTTP().SetCrossDomainXMLContent([]byte(crossDomainXMLContent))
	service.HTTP().SetClientAccessPolicyXMLContent([]byte(clientAccessPolicyXMLContent))
	service.AddFunction(func(ctx context.Context, name string) string {
		serviceContext := core.GetServiceContext(ctx)
		header, _ := serviceContext.Items().Get("httpRequestHeaders")
		return header.(http.Header).Get("test") + ":hello " + name
	}, "hello")
	assert.True(t, service.Get("hello").PassContext())
	serverMux := http.NewServeMux()
	serverMux.Handle("/test1", service.HTTP())
	serverMux.Handle("/test2", service.HTTP())
	serverMux.Handle("/crossdomain.xml", service.HTTP())
	serverMux.Handle("/clientaccesspolicy.xml", service.HTTP())
	server := &http.Server{Addr: ":8000", Handler: serverMux}
	go server.ListenAndServe()

	time.Sleep(time.Millisecond * 5)

	client := rpc.NewClient("http://127.0.0.1:8000/test1", "http://127.0.0.1:8000/test2")
	client.ShuffleURLs()
	client.Use(log.Plugin)
	client.HTTP().Header = http.Header{
		"test":   []string{"test"},
		"Origin": []string{"hprose.com"},
	}
	resp, err := client.HTTP().HTTPClient.Get("http://127.0.0.1:8000/crossdomain.xml")
	assert.NoError(t, err)
	content, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	assert.Equal(t, crossDomainXMLContent, string(content))
	assert.NoError(t, err)
	resp, err = client.HTTP().HTTPClient.Get("http://127.0.0.1:8000/clientaccesspolicy.xml")
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
	resp, err = client.HTTP().HTTPClient.Do(req)
	assert.Equal(t, http.StatusNotModified, resp.StatusCode)
	resp.Body.Close()
	assert.NoError(t, err)
	resp, err = client.HTTP().HTTPClient.Get("http://127.0.0.1:8000/test1")
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp.Body.Close()
	assert.NoError(t, err)
	service.HTTP().GET = false
	resp, err = client.HTTP().HTTPClient.Get("http://127.0.0.1:8000/test1")
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
