package main

import (
	"fmt"
	"sync/atomic"
	"time"

	rpc "github.com/hprose/hprose-golang/rpc/fasthttp"
)

// Stub is ...
type Stub struct {
	Hello      func(string) (string, error)
	AsyncHello func(func(string, error), string) `name:"hello"`
	Sum        func(...int) int
}

func main() {
	client := rpc.NewFastHTTPClient("http://127.0.0.1:8080/")
	client.SetMaxConcurrentRequests(128)
	var stub *Stub
	client.UseService(&stub)
	stub.AsyncHello(func(result string, err error) {
		fmt.Println(result, err)
	}, "async world")
	fmt.Println(stub.Hello("world"))
	fmt.Println(stub.Sum())
	fmt.Println(stub.Sum(1))
	fmt.Println(stub.Sum(1, 2))
	fmt.Println(stub.Sum(1, 2, 3, 4, 5, 6, 7))
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
