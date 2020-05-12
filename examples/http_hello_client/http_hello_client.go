package main

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/lanfengye2008/hprose-golang/rpc"
)

// Stub is ...
type Stub struct {
	Hello      func(string) (string, error)
	AsyncHello func(func(string, error), string) `name:"hello"`
}

func main() {
	client := rpc.NewHTTPClient("http://127.0.0.1:8080/")
	client.MaxIdleConnsPerHost = 128
	client.SetMaxConcurrentRequests(128)
	var stub *Stub
	client.UseService(&stub)
	stub.AsyncHello(func(result string, err error) {
		fmt.Println(result, err)
	}, "async world")
	fmt.Println(stub.Hello("world"))
	start := time.Now()
	var n int32 = 500000
	done := make(chan bool)
	for i := 0; i < 500000; i++ {
		stub.AsyncHello(func(result string, err error) {
			if atomic.AddInt32(&n, -1) == 0 {
				done <- true
			}
		}, "async world")
	}
	<-done
	stop := time.Now()
	fmt.Println((stop.UnixNano() - start.UnixNano()) / 1000000)
}
