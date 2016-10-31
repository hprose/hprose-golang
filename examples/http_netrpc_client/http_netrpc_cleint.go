package main

import (
	"fmt"
	"log"

	"github.com/hprose/hprose-golang/rpc"
)

// Args ...
type Args struct {
	A, B int
}

// Quotient ...
type Quotient struct {
	Quo, Rem int
}

// Stub ...
type Stub struct {
	// Synchronous call
	Multiply func(args *Args) int
	// Asynchronous call
	Divide func(func(*Quotient, error), *Args)
}

func main() {
	client := rpc.NewClient("http://127.0.0.1:8080")
	var stub *Stub
	client.UseService(&stub)
	fmt.Println(stub.Multiply(&Args{8, 7}))
	done := make(chan struct{})
	stub.Divide(func(result *Quotient, err error) {
		if err != nil {
			log.Fatal("arith error:", err)
		} else {
			fmt.Println(result.Quo, result.Rem)
		}
		done <- struct{}{}
	}, &Args{8, 7})
	<-done
}
