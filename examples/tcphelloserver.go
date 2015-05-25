package main

import (
	"os"

	"github.com/hprose/hprose-go"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	server := hprose.NewTcpServer("tcp4://0.0.0.0:4321/")
	server.AddFunction("hello", hello)
	server.Start()
	b := make([]byte, 1)
	os.Stdin.Read(b)
	server.Stop()
}
