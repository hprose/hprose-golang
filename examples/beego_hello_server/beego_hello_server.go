package main

import (
	"github.com/astaxie/beego"
	"github.com/hprose/hprose-golang/rpc"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	service := rpc.NewHTTPService()
	service.AddFunction("hello", hello)
	beego.Handler("/hello", service)
	beego.Run()
}
