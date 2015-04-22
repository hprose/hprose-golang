package main

import (
	"io/ioutil"
	"os"

	"github.com/hprose/hprose-go/hprose"
)

func main() {
	server := hprose.NewTcpServer("tcp4://0.0.0.0:4321/")
	server.AddFunction("writeFile", ioutil.WriteFile)
	server.Start()
	b := make([]byte, 1)
	os.Stdin.Read(b)
	server.Stop()
}
