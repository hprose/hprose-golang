package main

import (
	"github.com/hprose/hprose-golang/rpc"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	service := rpc.NewHTTPService()
	service.AddFunction("hello", hello)
	e := echo.New()
	e.Any("/hello", standard.WrapHandler(service))
	e.Run(standard.New(":8080"))
}
