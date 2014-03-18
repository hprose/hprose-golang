package main

import (
	"github.com/hprose/hprose-go/hprose"
	"net/http"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	service := hprose.NewHttpService()
	service.AddFunction("hello", hello)
	http.ListenAndServe(":8080", service)
}
