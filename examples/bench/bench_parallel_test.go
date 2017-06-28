package bench

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"runtime"
	"testing"

	hproserpc "github.com/hprose/hprose-golang/rpc"
)

// BenchmarkParallelHprose2 is ...
func BenchmarkParallelHprose2(b *testing.B) {
	b.StopTimer()
	server := hproserpc.NewTCPServer("")
	server.AddFunction("hello", hello)
	server.Handle()
	client := hproserpc.NewClient(server.URI())
	var ro *RO
	client.UseService(&ro)
	defer server.Close()
	defer client.Close()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ro.Hello("World")
		}
	})
	b.StopTimer()
}

// BenchmarkParallelHprose2Unix is ...
func BenchmarkParallelHprose2Unix(b *testing.B) {
	b.StopTimer()
	server := hproserpc.NewUnixServer("")
	server.AddFunction("hello", hello)
	server.Handle()
	client := hproserpc.NewClient(server.URI())
	var ro *RO
	client.UseService(&ro)
	defer server.Close()
	defer client.Close()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ro.Hello("World")
		}
	})
	b.StopTimer()
}

// BenchmarkParallelGobRPC is ...
func BenchmarkParallelGobRPC(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	b.StopTimer()
	server := rpc.NewServer()
	server.Register(new(Hello))
	listener, _ := net.Listen("tcp4", "")
	defer listener.Close()
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			go server.ServeConn(conn)
		}
	}()
	client, _ := rpc.Dial("tcp4", listener.Addr().String())
	defer client.Close()
	var args = &Args{"World"}
	var reply string
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			client.Call("Hello.Hello", &args, &reply)
		}
	})
	b.StopTimer()
}

// BenchmarkParallelGobRPCUnix is ...
func BenchmarkParallelGobRPCUnix(b *testing.B) {
	b.StopTimer()
	server := rpc.NewServer()
	server.Register(new(Hello))
	listener, _ := net.Listen("unix", "/tmp/gobrpc.sock")
	defer listener.Close()
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			go server.ServeConn(conn)
		}
	}()
	client, _ := rpc.Dial("unix", "/tmp/gobrpc.sock")
	defer client.Close()
	var args = &Args{"World"}
	var reply string
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			client.Call("Hello.Hello", &args, &reply)
		}
	})
	b.StopTimer()
}

// BenchmarkParallelJSONRPC is ...
func BenchmarkParallelJSONRPC(b *testing.B) {
	b.StopTimer()
	server := rpc.NewServer()
	server.Register(new(Hello))
	listener, _ := net.Listen("tcp4", "")
	defer listener.Close()
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}()
	client, _ := jsonrpc.Dial("tcp4", listener.Addr().String())
	defer client.Close()
	var args = &Args{"World"}
	var reply string
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			client.Call("Hello.Hello", &args, &reply)
		}
	})
	b.StopTimer()
}

// BenchmarkParallelJSONRPCUnix is ...
func BenchmarkParallelJSONRPCUnix(b *testing.B) {
	b.StopTimer()
	server := rpc.NewServer()
	server.Register(new(Hello))
	listener, _ := net.Listen("unix", "/tmp/jsonrpc.sock")
	defer listener.Close()
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}()
	client, _ := jsonrpc.Dial("unix", "/tmp/jsonrpc.sock")
	defer client.Close()
	var args = &Args{"World"}
	var reply string
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			client.Call("Hello.Hello", &args, &reply)
		}
	})
	b.StopTimer()
}
