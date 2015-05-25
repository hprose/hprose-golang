package main

import (
	"fmt"

	"github.com/hprose/hprose-go"
)

type Stub struct {
	Hello func(string) (string, error)
}

func main() {
	client := hprose.NewClient("unix:/tmp/my.sock")
	var stub *Stub
	client.UseService(&stub)
	fmt.Println(stub.Hello("world"))
}
