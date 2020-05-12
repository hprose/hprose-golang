package main

import "github.com/lanfengye2008/hprose-golang/rpc"

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	server := rpc.NewTCPServer("tcp4://0.0.0.0:4321/")
	server.AddFunction("hello", hello)
	server.Start()
}
