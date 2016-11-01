package main

import (
	rpc "github.com/hprose/hprose-golang/rpc/fasthttp"
	"github.com/valyala/fasthttp"
)

// Service ...
type Service struct{}

// Hello ...
func (Service) Hello(name string) string {
	return "Hello " + name + "!"
}

// Sum ...
func (Service) Sum(a ...int) int {
	sum := 0
	for _, i := range a {
		sum += i
	}
	return sum
}

func main() {
	service := rpc.NewFastHTTPService()
	service.Debug = true
	service.AddInstanceMethods(Service{})
	fasthttp.ListenAndServe(":8080", service.ServeFastHTTP)
}
