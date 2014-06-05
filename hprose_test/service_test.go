/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * hprose/service_test.go                                 *
 *                                                        *
 * hprose Service Test for Go.                            *
 *                                                        *
 * LastModified: Mar 15, 2014                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose_test

import (
	"errors"
	"fmt"
	"github.com/hprose/hprose-go/hprose"
	"net"
	"net/http/httptest"
	"testing"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

type testServe int

func (*testServe) Swap(a int, b int) (int, int) {
	return b, a
}

func (*testServe) Sum(args ...int) (int, error) {
	if len(args) < 2 {
		return 0, errors.New("Requires at least two parameters")
	}
	a := args[0]
	for i := 1; i < len(args); i++ {
		a += args[i]
	}
	return a, nil
}

func (*testServe) PanicTest() {
	panic("I'm crazy")
}

func getIP(conn net.Conn) string {
	return conn.RemoteAddr().String()
}

type testRemoteObject2 struct {
	Hello     func(string) (string, error)
	Swap      func(int, int) (int, int, error)
	Sum       func(...int) (int, error)
	GetIP     func() string
	PanicTest func() error
}

func TestHttpService(t *testing.T) {
	service := hprose.NewHttpService()
	service.AddFunction("hello", hello)
	service.AddMethods(new(testServe))
	server := httptest.NewServer(service)
	defer server.Close()
	client := hprose.NewClient(server.URL)
	var ro *testRemoteObject2
	client.UseService(&ro)
	if s, err := ro.Hello("World"); err != nil {
		t.Error(err.Error())
	} else {
		fmt.Println(s)
	}
	if a, b, err := ro.Swap(1, 2); err != nil {
		t.Error(err.Error())
	} else {
		fmt.Println(a, b)
	}
	if sum, err := ro.Sum(1); err != nil {
		fmt.Println(err.Error())
	} else {
		t.Error(sum)
	}
	if sum, err := ro.Sum(1, 2, 3, 4, 5); err != nil {
		t.Error(err.Error())
	} else {
		fmt.Println(sum)
	}
	if err := ro.PanicTest(); err != nil {
		fmt.Println(err.Error())
	} else {
		t.Error("missing panic")
	}
}

func TestTcpService(t *testing.T) {
	server := hprose.NewTcpServer("")
	server.AddFunction("hello", hello)
	server.AddMethods(new(testServe))
	server.AddFunction("getIP", getIP)
	server.Start()
	defer server.Stop()
	client := hprose.NewClient(server.URL)
	var ro *testRemoteObject2
	client.UseService(&ro)
	if s, err := ro.Hello("World"); err != nil {
		t.Error(err.Error())
	} else {
		fmt.Println(s)
	}
	if a, b, err := ro.Swap(1, 2); err != nil {
		t.Error(err.Error())
	} else {
		fmt.Println(a, b)
	}
	if sum, err := ro.Sum(1); err != nil {
		fmt.Println(err.Error())
	} else {
		t.Error(sum)
	}
	if sum, err := ro.Sum(1, 2, 3, 4, 5); err != nil {
		t.Error(err.Error())
	} else {
		fmt.Println(sum)
	}
	fmt.Println(ro.GetIP())
	if err := ro.PanicTest(); err != nil {
		fmt.Println(err.Error())
	} else {
		t.Error("missing panic")
	}
}

type testRemoteObject3 struct {
	Hello     func(string) (<-chan string, <-chan error)
	Swap      func(int, int) (<-chan int, <-chan int, <-chan error)
	Sum       func(...int) (<-chan int, <-chan error)
	PanicTest func() <-chan error
}

func TestTcpServiceAsync(t *testing.T) {
	server := hprose.NewTcpServer("")
	server.AddFunction("hello", hello)
	server.AddMethods(new(testServe))
	server.Start()
	defer server.Stop()
	client := hprose.NewClient(server.URL)
	var ro *testRemoteObject3
	client.UseService(&ro)
	s, err1 := ro.Hello("World")
	a, b, err2 := ro.Swap(1, 2)
	sum1, err3 := ro.Sum(1)
	if err := <-err1; err != nil {
		t.Error(err.Error())
	} else {
		fmt.Println(<-s)
	}
	if err := <-err2; err != nil {
		t.Error(err.Error())
	} else {
		fmt.Println(<-a, <-b)
	}
	if err := <-err3; err != nil {
		fmt.Println(err.Error())
	} else {
		t.Error(<-sum1)
	}
	sum2, err4 := ro.Sum(1, 2, 3, 4, 5)
	err5 := ro.PanicTest()
	if err := <-err4; err != nil {
		t.Error(err.Error())
	} else {
		fmt.Println(<-sum2)
	}
	if err := <-err5; err != nil {
		fmt.Println(err.Error())
	} else {
		t.Error("missing panic")
	}
}
