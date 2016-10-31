package main

import (
	"fmt"
	"reflect"

	"github.com/hprose/hprose-golang/rpc"
)

func logInvokeHandler(
	name string,
	args []reflect.Value,
	context rpc.Context,
	next rpc.NextInvokeHandler) (results []reflect.Value, err error) {
	fmt.Printf("%s(%v) = ", name, args)
	results, err = next(name, args, context)
	fmt.Printf("%v %v\r\n", results, err)
	return
}

func hello(name string) string {
	return "Hello " + name + "!"
}

// HelloService is ...
type HelloService struct {
	Hello func(string) (string, error)
	Hi    func(string) error
}

func main() {
	server := rpc.NewTCPServer("")
	server.AddFunction("hello", hello)
	server.Handle()
	client := rpc.NewClient(server.URI())
	client.AddInvokeHandler(logInvokeHandler)
	var helloService *HelloService
	client.UseService(&helloService)
	helloService.Hello("World")
	helloService.Hi("World")
	client.Close()
	server.Close()
}
