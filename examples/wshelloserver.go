package main

import (
	"net/http"

	"github.com/hprose/hprose-go/hprose"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	service := hprose.NewWebSocketService()
	service.AddFunction("hello", hello)
	http.ListenAndServe(":8080", service.Server)
}
