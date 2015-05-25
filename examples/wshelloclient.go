package main

import (
	"fmt"
	"reflect"

	"github.com/hprose/hprose-go"
)

type A struct {
	S string `json:"str"`
}

type Stub struct {
	Hello         func(string) string
	GetEmptySlice func() interface{}
}

func main() {
	hprose.ClassManager.Register(reflect.TypeOf(A{}), "A", "json")
	client := hprose.NewClient("ws://127.0.0.1:8080/")
	var stub *Stub
	client.UseService(&stub)
	fmt.Println(stub.Hello("world"))
	stub.GetEmptySlice()
}
