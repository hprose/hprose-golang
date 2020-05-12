package main

import (
	"net/http"

	"github.com/lanfengye2008/hprose-golang/rpc"
)

func hello(name string, context *rpc.HTTPContext) string {
	return "Hello !" + context.Request.RemoteAddr
}

func main() {
	service := rpc.NewHTTPService()
	service.Debug = true
	service.AddFunction("hello", hello)
	http.ListenAndServe(":8080", service)
}
