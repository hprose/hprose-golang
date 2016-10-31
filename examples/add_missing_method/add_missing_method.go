package main

import (
	"net/http"
	"reflect"

	"github.com/hprose/hprose-golang/rpc"
)

// HproseProxy ...
type HproseProxy struct {
	client rpc.Client
}

func newHproseProxy() *HproseProxy {
	proxy := new(HproseProxy)
	proxy.client = rpc.NewClient("http://www.hprose.com/example/")
	return proxy
}

// MissingMethod ...
func (proxy *HproseProxy) MissingMethod(
	name string, args []reflect.Value, context rpc.Context) ([]reflect.Value, error) {
	return proxy.client.Invoke(name, args, &rpc.InvokeSettings{
		Mode:        rpc.Raw,
		ResultTypes: []reflect.Type{reflect.TypeOf(([]byte)(nil))},
	})
}

func main() {
	service := rpc.NewHTTPService()
	service.AddMissingMethod(newHproseProxy().MissingMethod, rpc.Options{Mode: rpc.Raw})
	http.ListenAndServe(":8080", service)
}
