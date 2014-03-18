package main

import (
	"fmt"
	"github.com/hprose/hprose-go/hprose"
)

type Stub struct {
	Hello func(string) string
}

func main() {
	client := hprose.NewClient("http://127.0.0.1:8080/")
	var stub *Stub
	client.UseService(&stub)
	fmt.Println(stub.Hello("world"))
}
