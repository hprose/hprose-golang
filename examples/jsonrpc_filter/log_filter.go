package main

import (
	"fmt"

	"github.com/hprose/hprose-golang/rpc"
)

// LogFilter ...
type LogFilter struct {
	Prompt string
}

// InputFilter ...
func (lf LogFilter) InputFilter(data []byte, context rpc.Context) []byte {
	fmt.Printf("%v: %s\r\n", lf.Prompt, data)
	return data
}

// OutputFilter ...
func (lf LogFilter) OutputFilter(data []byte, context rpc.Context) []byte {
	fmt.Printf("%v: %s\r\n", lf.Prompt, data)
	return data
}
