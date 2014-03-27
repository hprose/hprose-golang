package main

import (
	"github.com/hprose/hprose-go/hprose"
	"net/http"
	"reflect"
)

type User struct {
	Name string `json:"n"`
	Age  int    `json:"a"`
	HaHa string `json:"-"`
}

func getUser() *User {
	return &User{Name: "Tom", Age: 28, HaHa: "Don't Serialize Me!"}
}

func main() {
	hprose.ClassManager.Register(reflect.TypeOf(User{}), "User", "json")
	service := hprose.NewHttpService()
	service.AddFunction("getUser", getUser)
	http.ListenAndServe(":8080", service)
}
