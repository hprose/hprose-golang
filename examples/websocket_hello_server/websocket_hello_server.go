package main

import (
	"net/http"

	"github.com/hprose/hprose-golang/rpc"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	service := rpc.NewWebSocketService()
	service.AddFunction("hello", hello)
	http.ListenAndServe(":8080", service)
}
