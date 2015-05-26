package main

import "github.com/hprose/hprose-go"

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	server := hprose.NewUnixServer("unix:/tmp/my.sock")
	server.AddFunction("hello", hello)
	server.Start()
}
