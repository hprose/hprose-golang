package main

import (
	"fmt"

	"github.com/hprose/hprose-golang/rpc"
	"github.com/hprose/hprose-golang/rpc/filter/jsonrpc"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

// HelloService is ...
type HelloService struct {
	Hello1 func(string) (string, error) `name:"hello" userdata:"{\"jsonrpc\":true}"`
	Hello2 func(string) (string, error) `name:"hello"`
}

func main() {
	server := rpc.NewTCPServer("")
	server.AddFunction("hello", hello).AddFilter(
		jsonrpc.ServiceFilter{},
		LogFilter{"Server"},
	)
	server.Handle()
	client := rpc.NewClient(server.URI()).AddFilter(
		jsonrpc.NewClientFilter("2.0"),
		LogFilter{"Client"},
	)
	var helloService *HelloService
	client.UseService(&helloService)
	fmt.Println(helloService.Hello1("JSONRPC"))
	fmt.Println(helloService.Hello2("Hprose"))
	fmt.Println(helloService.Hello1("World"))
	client.Close()
	server.Close()
}
