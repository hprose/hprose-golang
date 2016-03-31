package main

import (
	"fmt"

	"github.com/hprose/hprose-go"
	"time"
	"runtime"
)

type Stub struct {
	Hello func(string) (string, error)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	client := hprose.NewClient("tcp4://127.0.0.1:4321/") //Used for Peer to Peer model...
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
