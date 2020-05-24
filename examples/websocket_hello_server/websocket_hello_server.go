package main

import (
	"net/http"

	rpc "github.com/lanfengye2008/hprose-golang/rpc/websocket"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	service := rpc.NewWebSocketService()
	service.AddFunction("hello", hello)
	http.ListenAndServe(":8080", service)
}
