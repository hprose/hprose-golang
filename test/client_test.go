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
 * hprose/client_test.go                                  *
 *                                                        *
 * hprose Client Test for Go.                             *
 *                                                        *
 * LastModified: Feb 24, 2014                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose_test

import (
	//"fmt"
	//. "../hprose"
	//"testing"
	"time"
)

type testUser struct {
	Name     string
	Sex      int
	Birthday time.Time
	Age      int
	Married  bool
}

type testRemoteObject struct {
	Hello               func(string) string
	HelloWithError      func(string) (string, error)               `name:"hello"`
	AsyncHello          func(string) <-chan string                 `name:"hello"`
	AsyncHelloWithError func(string) (<-chan string, <-chan error) `name:"hello"`
	Sum                 func(...int) int
	SwapKeyAndValue     func(*map[string]string) map[string]string `byref:"true"`
	SwapInt             func(int, int) (int, int)                  `name:"swap"`
	SwapFloat           func(float64, float64) (float64, float64)  `name:"swap"`
	Swap                func(interface{}, interface{}) (interface{}, interface{})
	GetUserList         func() []testUser
}

/*
func TestRemoteObject(t *testing.T) {
	//client := NewClient("http://127.0.0.1/")
	client := NewClient("http://www.hprose.com/example/")
	var ro *testRemoteObject
	client.UseService(&ro)

	// If an error occurs, it will panic
	fmt.Println(ro.Hello("World"))

	// If an error occurs, an error value will be returned
	if result, err := ro.HelloWithError("World"); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err.Error())
	}

	// If an error occurs, it will be ignored
	result := ro.AsyncHello("World")
	fmt.Println(<-result)

	// If an error occurs, an error chan will be returned
	result, err := ro.AsyncHelloWithError("World")
	if e := <-err; e == nil {
		fmt.Println(<-result)
	} else {
		fmt.Println(e.Error())
	}
	fmt.Println(ro.Sum(1, 2, 3, 4, 5))

	m := make(map[string]string)
	m["Jan"] = "January"
	m["Feb"] = "February"
	m["Mar"] = "March"
	m["Apr"] = "April"
	m["May"] = "May"
	m["Jun"] = "June"
	m["Jul"] = "July"
	m["Aug"] = "August"
	m["Sep"] = "September"
	m["Oct"] = "October"
	m["Nov"] = "November"
	m["Dec"] = "December"

	fmt.Println(m)
	mm := ro.SwapKeyAndValue(&m)
	fmt.Println(m)
	fmt.Println(mm)

	fmt.Println(ro.GetUserList())
	fmt.Println(ro.SwapInt(1, 2))
	fmt.Println(ro.SwapFloat(1.2, 3.4))
	fmt.Println(ro.Swap("Hello", "World"))
}

func TestClient(t *testing.T) {
	//client := NewClient("http://127.0.0.1/")
	client := NewClient("http://www.hprose.com/example/")
	var r1 string
	if err := <-client.Invoke("hello", []interface{}{"world"}, nil, &r1); err != nil {
		t.Error(err.Error())
	}
	fmt.Println(r1)

	var r2 chan int
	if err := <-client.Invoke("sum", []interface{}{1, 2, 3, 4, 5, 6, 7}, nil, &r2); err != nil {
		t.Error(err.Error())
	}
	fmt.Println(<-r2)

	var r3 chan map[string]string
	m := make(map[string]string)
	m["Jan"] = "January"
	m["Feb"] = "February"
	m["Mar"] = "March"
	m["Apr"] = "April"
	m["May"] = "May"
	m["Jun"] = "June"
	m["Jul"] = "July"
	m["Aug"] = "August"
	m["Sep"] = "September"
	m["Oct"] = "October"
	m["Nov"] = "November"
	m["Dec"] = "December"
	fmt.Println(m)
	if err := <-client.Invoke("swapKeyAndValue", []interface{}{&m}, &InvokeOptions{ByRef: true}, &r3); err != nil {
		t.Error(err.Error())
	}
	fmt.Println(<-r3)
	fmt.Println(m)

	//ClassManager.Register(reflect.TypeOf(testUser{}), "User")

	var r4 chan []testUser
	if err := <-client.Invoke("getUserList", []interface{}{}, nil, &r4); err != nil {
		t.Error(err.Error())
	}
	fmt.Println(<-r4)

	var r5 chan []byte
	if err := <-client.Invoke("hello", []interface{}{"马秉尧"}, &InvokeOptions{ResultMode: Serialized}, &r5); err != nil {
		t.Error(err.Error())
	}
	fmt.Println(string(<-r5))
	if err := <-client.Invoke("hello", []interface{}{"马秉尧"}, &InvokeOptions{ResultMode: Raw}, &r5); err != nil {
		t.Error(err.Error())
	}
	fmt.Println(string(<-r5))
	if err := <-client.Invoke("hello", []interface{}{"马秉尧"}, &InvokeOptions{ResultMode: RawWithEndTag}, &r5); err != nil {
		t.Error(err.Error())
	}
	fmt.Println(string(<-r5))
	if err := <-client.Invoke("swapKeyAndValue", []interface{}{&m}, &InvokeOptions{ByRef: true, ResultMode: RawWithEndTag}, &r5); err != nil {
		t.Error(err.Error())
	}
	fmt.Println(string(<-r5))
}
*/
