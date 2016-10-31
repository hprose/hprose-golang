package main

import (
	"fmt"

	"github.com/hprose/hprose-golang/rpc"
)

// TestService is ...
type TestService struct {
	Test func([]int) ([]int, error)
}

func main() {
	server := rpc.NewTCPServer("")
	server.AddFunction("test", func(data []int) []int {
		return data
	}).
		AddFilter(
			SizeFilter{"Non compressed data on server"},
			CompressFilter{},
			SizeFilter{"Compressed data on server"},
		)
	server.Debug = true
	server.Handle()
	client := rpc.NewClient(server.URI())
	client.AddFilter(
		SizeFilter{"Non compressed data on client"},
		CompressFilter{},
		SizeFilter{"Compressed data on client"},
	)
	var testService *TestService
	client.UseService(&testService)
	args := make([]int, 100000)
	for i := range args {
		args[i] = i
	}
	result, err := testService.Test(args)
	fmt.Println(len(result), err)
	client.Close()
	server.Close()
}
