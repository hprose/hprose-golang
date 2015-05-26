package main

import (
	"io/ioutil"

	"github.com/hprose/hprose-go"
)

func main() {
	server := hprose.NewTcpServer("tcp4://0.0.0.0:4321/")
	server.AddFunction("writeFile", ioutil.WriteFile)
	server.Start()
}
