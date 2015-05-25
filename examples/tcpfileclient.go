package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hprose/hprose-go"
)

type Stub struct {
	WriteFile func(filename string, data []byte, perm os.FileMode) error
}

func main() {
	client := hprose.NewClient("tcp4://127.0.0.1:4321/")
	var stub *Stub
	client.UseService(&stub)
	data, err := ioutil.ReadFile("hello.txt")
	if err == nil {
		err = stub.WriteFile("hello2.txt", data, 0777)
		if err == nil {
			fmt.Println("SUCCESS")
		}
	}
}
