package main

import (
	rpc "github.com/hprose/hprose-golang/rpc"
	"github.com/kataras/iris"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	service := rpc.NewHTTPService()
	service.AddFunction("hello", hello)
	iris.Any("/hello", func(c *iris.Context) {
		service.ServeHTTP(c.ResponseWriter, c.Request)
	})
	iris.Listen(":8080")
}
