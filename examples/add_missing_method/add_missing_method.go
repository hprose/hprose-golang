package main

import (
	"net/http"
	"reflect"

	"github.com/hprose/hprose-golang/rpc"
)

// HproseProxy ...
type HproseProxy struct {
	client   rpc.Client
	settings rpc.InvokeSettings
}

func newHproseProxy() *HproseProxy {
	proxy := new(HproseProxy)
	proxy.client = rpc.NewClient("http://www.hprose.com/example/")
	proxy.settings = rpc.InvokeSettings{
		Mode:        rpc.Raw,
		ResultTypes: []reflect.Type{reflect.TypeOf(([]byte)(nil))},
	}
	return proxy
}

// Proxy ...
func (proxy *HproseProxy) Proxy(
	name string, args []reflect.Value, context rpc.Context) ([]reflect.Value, error) {
	return proxy.client.Invoke(name, args, &proxy.settings)
}

func main() {
	service := rpc.NewHTTPService()
	service.AddMissingMethod(newHproseProxy().Proxy, rpc.Options{Mode: rpc.Raw})
	http.ListenAndServe(":8080", service)
}
