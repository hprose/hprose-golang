package main

import (
	"fmt"
	"reflect"

	"github.com/hprose/hprose-go"
)

type User struct {
	Name string `json:"n"`
	Age  int    `json:"a"`
	HaHa string `json:"-"`
}

type Stub struct {
	Hello   func(string) string
	GetUser func(name string, age int) *User
}

func main() {
	hprose.ClassManager.Register(reflect.TypeOf(User{}), "User", "json")
	client := hprose.NewClient("http://127.0.0.1:8080/")
	client.SetFilter(hprose.NewJSONRPCClientFilter("2.0"))
	var stub *Stub
	client.UseService(&stub)
	user := stub.GetUser("Jerry", 16)
	fmt.Println(user.Name)
	fmt.Println(user.Age)
	fmt.Println(user.HaHa)
}
