package main

import (
	"net/http"
	"reflect"

	"github.com/hprose/hprose-go"
)

func hello(name string, context *hprose.HttpContext) string {
	return "Hello " + name + "!  -  " + context.Request.RemoteAddr
}

type User struct {
	Name string `json:"n"`
	Age  int    `json:"a"`
	HaHa string `json:"-"`
}

func getUser(name string, age int) *User {
	return &User{Name: name, Age: age, HaHa: "Don't Serialize Me!"}
}

func main() {
	hprose.ClassManager.Register(reflect.TypeOf(User{}), "User", "json")
	service := hprose.NewHttpService()
	service.AddFunction("hello", hello)
	service.AddFunction("getUser", getUser)
	service.SetFilter(hprose.JSONRPCServiceFilter{})
	http.ListenAndServe(":8080", service)
}
