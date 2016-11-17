package main

import (
	"github.com/hprose/hprose-golang/rpc"
	"github.com/labstack/echo"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	service := rpc.NewHTTPService()
	service.AddFunction("hello", hello)
	e := echo.New()
	e.Any("/hello", echo.WrapHandler(service))
	e.Start(":8080")
}
