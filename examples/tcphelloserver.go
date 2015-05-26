package main

import "github.com/hprose/hprose-go"

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	server := hprose.NewTcpServer("tcp4://0.0.0.0:4321/")
	server.AddFunction("hello", hello)
	server.Start()
}
