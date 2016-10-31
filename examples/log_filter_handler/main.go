package main

import (
	"fmt"

	"github.com/hprose/hprose-golang/rpc"
)

type logFilter struct {
	Prompt string
}

func (lf logFilter) handler(
	request []byte,
	context rpc.Context,
	next rpc.NextFilterHandler) (response []byte, err error) {
	fmt.Printf("%v: %s\r\n", lf.Prompt, request)
	response, err = next(request, context)
	fmt.Printf("%v: %s\r\n", lf.Prompt, response)
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
	server.AddBeforeFilterHandler(logFilter{"Server"}.handler)
	server.Handle()
	client := rpc.NewClient(server.URI())
	client.AddBeforeFilterHandler(logFilter{"Client"}.handler)
	var helloService *HelloService
	client.UseService(&helloService)
	helloService.Hello("World")
	helloService.Hi("World")
	client.Close()
	server.Close()
}
