package main

import (
	"fmt"

	"github.com/hprose/hprose-golang/rpc"
)

func swap(a, b *int) {
	*b, *a = *a, *b
}

// RO is ...
type RO struct {
	Swap func(a, b *int) error `byref:"true"`
}

func main() {
	server := rpc.NewTCPServer("")
	server.AddFunction("swap", swap)
	server.Handle()
	client := rpc.NewClient(server.URI())
	var ro *RO
	client.UseService(&ro)
	a := 1
	b := 2
	ro.Swap(&a, &b)
	fmt.Println(a, b)
	client.Close()
	server.Close()
}
