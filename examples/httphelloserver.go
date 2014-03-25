package main

import (
	"github.com/hprose/hprose-go/hprose"
	"net/http"
)

func hello(name string, context *hprose.HttpContext) string {
	return "Hello " + name + "!  -  " + context.Request.RemoteAddr
}

func main() {
	service := hprose.NewHttpService()
	service.AddFunction("hello", hello)
	http.ListenAndServe(":8080", service)
}
