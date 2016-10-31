package main

import (
	"fmt"

	"github.com/hprose/hprose-golang/rpc"
)

// HelloService ...
type HelloService struct {
	Hello func(string) (string, error)
}

func main() {
	client := rpc.NewClient("tcp://127.0.0.1:4321/")
	client.Subscribe("ip", "", nil, func(ip string) {
		fmt.Println(ip)
	})
	var helloService *HelloService
	client.UseService(&helloService)
	for i := 0; i < 10; i++ {
		fmt.Println(helloService.Hello("world"))
	}
}
