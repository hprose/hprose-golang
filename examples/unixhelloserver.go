package main

import (
	"os"

	"github.com/hprose/hprose-go/hprose"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	server := hprose.NewUnixServer("unix:/tmp/my.sock")
	server.AddFunction("hello", hello)
	server.Start()
	b := make([]byte, 1)
	os.Stdin.Read(b)
	server.Stop()
}
