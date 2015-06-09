package hprosebench

import (
	//"fmt"
	"github.com/hprose/hprose-go"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

// BenchmarkHprose is ...
func BenchmarkHprose(b *testing.B) {
	b.StopTimer()
	server := hprose.NewTcpServer("")
	server.AddFunction("hello", hello)
	server.Handle()
	client := hprose.NewClient(server.URL)
	defer server.Stop()
	var args = []interface{}{"World"}
	var result string
	// client.Invoke("Hello", args, nil, &result)
	// fmt.Println(result)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		client.Invoke("Hello", args, nil, &result)
	}
}

// RO is Reomote object
type RO struct {
	Hello func(string) (string, error)
}

// BenchmarkHprose2 is ...
func BenchmarkHprose2(b *testing.B) {
	b.StopTimer()
	server := hprose.NewTcpServer("")
	server.AddFunction("hello", hello)
	server.Handle()
	client := hprose.NewClient(server.URL)
	var ro *RO
	client.UseService(&ro)
	defer server.Stop()
	// result, _ := ro.Hello("World")
	// fmt.Println(result)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ro.Hello("World")
	}
}

type Args struct {
	Name string
}

type Hello int

func (this *Hello) Hello(args *Args, result *string) error {
	*result = "Hello " + args.Name + "!"
	return nil
}

// BenchmarkGobRPC is ...
func BenchmarkGobRPC(b *testing.B) {
	b.StopTimer()
	server := rpc.NewServer()
	server.Register(new(Hello))
	listener, _ := net.Listen("tcp", "")
	go server.Accept(listener)
	client, _ := rpc.Dial("tcp", listener.Addr().String())
	defer client.Close()
	var args = &Args{"World"}
	var reply string
	// client.Call("Hello.Hello", &args, &reply)
	// fmt.Println(reply)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		client.Call("Hello.Hello", &args, &reply)
	}
}

// BenchmarkJSONRPC is ...
func BenchmarkJSONRPC(b *testing.B) {
	b.StopTimer()
	server := rpc.NewServer()
	server.Register(new(Hello))
	listener, _ := net.Listen("tcp", "")
	go func() {
		for {
			conn, _ := listener.Accept()
			server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}()
	client, _ := jsonrpc.Dial("tcp", listener.Addr().String())
	defer client.Close()
	var args = &Args{"World"}
	var reply string
	// client.Call("Hello.Hello", &args, &reply)
	// fmt.Println(reply)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		client.Call("Hello.Hello", &args, &reply)
	}
}
