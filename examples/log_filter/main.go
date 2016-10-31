package main

import "github.com/hprose/hprose-golang/rpc"

func hello(name string) string {
	return "Hello " + name + "!"
}

// HelloService is ...
type HelloService struct {
	Hello func(string) (string, error)
}

func main() {
	server := rpc.NewTCPServer("")
	server.AddFunction("hello", hello).AddFilter(LogFilter{"Server"})
	server.Handle()
	client := rpc.NewClient(server.URI())
	client.AddFilter(LogFilter{"Client"})
	var helloService *HelloService
	client.UseService(&helloService)
	helloService.Hello("World")
	client.Close()
	server.Close()
}
