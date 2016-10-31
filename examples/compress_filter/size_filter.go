package main

import (
	"fmt"

	"github.com/hprose/hprose-golang/rpc"
)

// SizeFilter ...
type SizeFilter struct {
	Message string
}

// InputFilter ...
func (sf SizeFilter) InputFilter(data []byte, context rpc.Context) []byte {
	fmt.Printf("%v input size: %d\r\n", sf.Message, len(data))
	return data
}

// OutputFilter ...
func (sf SizeFilter) OutputFilter(data []byte, context rpc.Context) []byte {
	fmt.Printf("%v output size: %d\r\n", sf.Message, len(data))
	return data
}
