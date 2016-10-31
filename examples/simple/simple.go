package main

import (
	"fmt"

	"github.com/hprose/hprose-golang/rpc"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

// Stub is ...
type Stub struct {
	Hello func(string) (string, error) `simple:"true" idempotent:"true" retry:"30"`
}

func main() {
	server := rpc.NewTCPServer("")
	server.AddFunction("hello", hello)
	server.Handle()
	client := rpc.NewClient(server.URI())
	var stub *Stub
	client.UseService(&stub)
	fmt.Println(stub.Hello("World"))
	client.Close()
	server.Close()
}
