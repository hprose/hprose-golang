package main

import (
	"net/http"

	"github.com/hprose/hprose-golang/rpc"
)

// Foo ...
type Foo int

// MethodA1 ...
func (Foo) MethodA1() {}

// MethodA2 ...
func (*Foo) MethodA2() {}

// Bar ...
type Bar struct {
	Foo
	FuncB func()
}

// MethodB1 ...
func (Bar) MethodB1() {}

// MethodB2 ...
func (*Bar) MethodB2() {}

// Foobar ...
type Foobar struct {
	FooField Foo
	BarField *Bar
}

// MethodC1 ...
func (Foobar) MethodC1() {}

// MethodC2 ...
func (*Foobar) MethodC2() {}

func main() {
	service := rpc.NewHTTPService()
	foobar := &Foobar{}
	foobar.BarField = new(Bar)
	foobar.BarField.FuncB = func() {}
	service.AddAllMethods(foobar)
	http.ListenAndServe(":8080", service)
}
