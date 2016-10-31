package main

import (
	"fmt"
	"time"

	"github.com/hprose/hprose-golang/rpc"
)

// StatFilter ...
type StatFilter struct {
	Prompt string
}

func (sf StatFilter) stat(context rpc.Context) {
	startTime := context.GetInt64("startTime")
	if startTime > 0 {
		fmt.Printf("%v takes %d ns.\r\n", sf.Prompt, time.Now().UnixNano()-startTime)
	} else {
		context.SetInt64("startTime", time.Now().UnixNano())
	}
}

// InputFilter ...
func (sf StatFilter) InputFilter(data []byte, context rpc.Context) []byte {
	sf.stat(context)
	return data
}

// OutputFilter ...
func (sf StatFilter) OutputFilter(data []byte, context rpc.Context) []byte {
	sf.stat(context)
	return data
}
