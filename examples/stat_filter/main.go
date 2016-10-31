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
	server.AddFunction("hello", hello).AddFilter(StatFilter{"Server"})
	server.Handle()
	client := rpc.NewClient(server.URI())
	client.AddFilter(StatFilter{"Client"})
	var helloService *HelloService
	client.UseService(&helloService)
	for i := 0; i < 3; i++ {
		helloService.Hello("World")
	}
	client.Close()
	server.Close()
}
