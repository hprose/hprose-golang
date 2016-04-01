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
 * Author: Ma Bingyao <andot@hprose.com>                  *
 * Modified By: Henry Hu <henry.pf.hu@gmail.com>          *
 *                                                        *
\**********************************************************/

package main

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/hprose/hprose-go"
	"runtime"
	"github.com/hprose/hprose-go/registry/etcd"
)

func hello(name string, context *hprose.HttpContext) string {
	return "Hello " + name + "!  -  " + context.Request.RemoteAddr
}

type A struct {
	S string `json:"str"`
}

func getEmptySlice() interface{} {
	s := make([]A, 100)
	return s
}

type ServerEvent struct{}

func (e *ServerEvent) OnBeforeInvoke(name string, args []reflect.Value, byref bool, context hprose.Context) {
	fmt.Println("Before OK")
}

func (e *ServerEvent) OnAfterInvoke(name string, args []reflect.Value, byref bool, result []reflect.Value, context hprose.Context) {
	fmt.Println("After OK")
}
func (e *ServerEvent) OnSendError(err error, context hprose.Context) {
	fmt.Println(err)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	hprose.ClassManager.Register(reflect.TypeOf(A{}), "A", "json")
	service := hprose.NewHttpService()
//	service.ServiceEvent = &ServerEvent{}
	service.DebugEnabled = false
	service.AddFunction("hello", hello)
	service.AddFunction("getEmptySlice", getEmptySlice)

	domain := "http.hello.server"
	httpEndpoint := "http://"+etcd.GetLocalIP()+":8080/"
	etcdEndpoints :=[]string{"http://127.0.0.1:2379"}
	etcd.RegisterServer(domain,httpEndpoint,etcdEndpoints)

	http.ListenAndServe(":8080", service)
}
