package main

import (
	"fmt"
	"reflect"

	"github.com/hprose/hprose-go"
	"runtime"
	"time"
)

type A struct {
	S string `json:"str"`
}

type Stub struct {
	Hello         func(string) string
	GetEmptySlice func() interface{}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	hprose.ClassManager.Register(reflect.TypeOf(A{}), "A", "json")

	client := hprose.NewClient("http://127.0.0.1:8080/")
	var stub *Stub
	client.UseService(&stub)

	startTime := time.Now()
	for i := 1; i < 500000; i++ {
		result := stub.Hello("world")
		if i%10000 == 0 {
			println("HttpRequest Result: ", result)
		}
	}
	endTime := time.Now()
	fmt.Println("Time used: ", endTime.Sub(startTime).Seconds())

	fmt.Println(stub.Hello("world"))
	stub.GetEmptySlice()
}
