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
	"reflect"

	"github.com/hprose/hprose-go"
	"time"
	"github.com/hprose/hprose-go/registry/etcd"
)

type A struct {
	S string `json:"str"`
}

type Stub struct {
	Hello         func(string) string
	GetEmptySlice func() interface{}
}

func main() {
	hprose.ClassManager.Register(reflect.TypeOf(A{}), "A", "json")

	domain := "ws.hello.server"
	etcdEndpoints :=[]string{"http://127.0.0.1:2379"}
	client := etcd.NewClient(domain,etcdEndpoints) //Used for Clustering model...

	//client := hprose.NewClient("ws://127.0.0.1:8080/")
	var stub *Stub
	client.UseService(&stub)

	startTime := time.Now()
	for i := 1; i < 500000; i++ {
		result := stub.Hello("world")
		if i%10000 == 0 {
			println("HttpRequest Result: ", result)
		}
	}
	endTime := time.Now()
	fmt.Println("Time used: ", endTime.Sub(startTime).Seconds())

	stub.GetEmptySlice()
}
