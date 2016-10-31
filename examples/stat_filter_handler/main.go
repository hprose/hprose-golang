package main

import (
	"fmt"

	"time"

	"sync"

	"github.com/hprose/hprose-golang/rpc"
)

type cacheFilter struct {
	Cache map[string][]byte
	sync.RWMutex
}

func (cf *cacheFilter) handler(
	request []byte,
	context rpc.Context,
	next rpc.NextFilterHandler) (response []byte, err error) {
	if context.GetBool("cache") {
		var ok bool
		cf.RLock()
		if response, ok = cf.Cache[string(request)]; ok {
			cf.RUnlock()
			return response, nil
		}
		cf.RUnlock()
		response, err = next(request, context)
		if err != nil {
			return
		}
		cf.Lock()
		cf.Cache[string(request)] = response
		cf.Unlock()
		return response, nil
	}
	return next(request, context)
}

type sizeFilter struct {
	Message string
}

func (sf sizeFilter) handler(
	request []byte,
	context rpc.Context,
	next rpc.NextFilterHandler) (response []byte, err error) {
	fmt.Printf("%v request size is %d\r\n", sf.Message, len(request))
	response, err = next(request, context)
	fmt.Printf("%v response size is %d\r\n", sf.Message, len(response))
	return
}

type statFilter struct {
	Message string
}

func (sf statFilter) handler(
	request []byte,
	context rpc.Context,
	next rpc.NextFilterHandler) (response []byte, err error) {
	start := time.Now()
	response, err = next(request, context)
	end := time.Now()
	fmt.Printf("%v takes %d ms.\r\n", sf.Message, end.Sub(start)/time.Millisecond)
	return
}

// TestService is ...
type TestService struct {
	Test func([]int) ([]int, error) `userdata:"{\"cache\":true}"`
}

func main() {
	server := rpc.NewTCPServer("")
	server.AddFunction("test", func(data []int) []int {
		return data
	}).
		AddBeforeFilterHandler(
			statFilter{"Server: BeforeFilter"}.handler,
			sizeFilter{"Server: Compressed"}.handler,
		).
		AddFilter(CompressFilter{}).
		AddAfterFilterHandler(
			statFilter{"Server: AfterFilter"}.handler,
			sizeFilter{"Server: Non Compressed"}.handler,
		)
	server.Handle()
	client := rpc.NewClient(server.URI())
	client.AddFilter(CompressFilter{}).
		AddBeforeFilterHandler(
			(&cacheFilter{Cache: make(map[string][]byte)}).handler,
			statFilter{"Client: BeforeFilter"}.handler,
			sizeFilter{"Client: Compressed"}.handler,
		).
		AddAfterFilterHandler(
			statFilter{"Client: AfterFilter"}.handler,
			sizeFilter{"Client: Non Compressed"}.handler,
		)
	var testService *TestService
	client.UseService(&testService)
	args := make([]int, 100000)
	for i := range args {
		args[i] = i
	}
	result, err := testService.Test(args)
	fmt.Println(len(result), err)
	fmt.Println(len(result), err)
	client.Close()
	server.Close()
}
