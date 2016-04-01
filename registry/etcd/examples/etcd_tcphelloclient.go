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

	"time"
	"runtime"
	"github.com/hprose/hprose-go/registry/etcd"
)

type Stub struct {
	Hello func(string) (string, error)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	domain := "tcp.hello.server"
	etcdEndpoints :=[]string{"http://127.0.0.1:2379"}
	client := etcd.NewClient(domain,etcdEndpoints) //Used for Clustering model...

	var stub *Stub
	client.UseService(&stub)
	client.SetKeepAlive(true)

	runFlags := make(chan int, 1)

	startTime := time.Now()
	for i := 1; i < 20; i ++ {
		go func( c chan int) {
			for i := 1; i < 100000; i++ {
				stub.Hello("world")
			}
			c <- 1
		}(runFlags)
	}

	for i := 1; i < 20; i ++ {
		<- runFlags
	}

	endTime := time.Now()
	fmt.Println("Time used: ", endTime.Sub(startTime).Seconds())

	result, error := stub.Hello("world")
	if error == nil {
		fmt.Println(result)
	} else {
		fmt.Println(error)
	}
}
