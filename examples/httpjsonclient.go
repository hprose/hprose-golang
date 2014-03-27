package main

import (
	"fmt"
	"github.com/hprose/hprose-go/hprose"
	"reflect"
)

type User struct {
	Name string `json:"n"`
	Age  int    `json:"a"`
	HaHa string `json:"-"`
}

type Stub struct {
	GetUser func() *User
}

func main() {
	hprose.ClassManager.Register(reflect.TypeOf(User{}), "User", "json")
	client := hprose.NewClient("http://127.0.0.1:8080/")
	var stub *Stub
	client.UseService(&stub)
	user := stub.GetUser()
	fmt.Println(user.Name)
	fmt.Println(user.Age)
	fmt.Println(user.HaHa)
}
