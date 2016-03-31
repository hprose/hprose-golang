package main

import (
	"github.com/hprose/hprose-go"
	"runtime"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	server := hprose.NewTcpServer("tcp4://localhost:4321/")
	server.AddFunction("hello", hello)
	server.SetKeepAlive(true)
	server.Start()
}
