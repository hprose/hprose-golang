package main

import "github.com/hprose/hprose-golang/rpc"

func hello(name string, context *rpc.SocketContext) string {
	context.Clients().Push("ip", context.Conn.RemoteAddr().String())
	return "Hello " + name + "!"
}

func main() {
	server := rpc.NewTCPServer("tcp4://0.0.0.0:4321/")
	server.AddFunction("hello", hello)
	server.Publish("ip", 0, 0)
	server.Start()
}
