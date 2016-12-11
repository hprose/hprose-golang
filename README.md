<p align="center"><img src="http://hprose.com/banner.@2x.png" alt="Hprose" title="Hprose" width="650" height="200" /></p>

# Hprose 2.0 for Golang

[![Join the chat at https://gitter.im/hprose/hprose-golang](https://img.shields.io/badge/GITTER-join%20chat-green.svg)](https://gitter.im/hprose/hprose-golang?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Build Status](https://travis-ci.org/hprose/hprose-golang.svg?branch=master)](https://travis-ci.org/hprose/hprose-golang)
[![GoDoc](https://godoc.org/github.com/hprose/hprose-golang?status.svg&style=flat)](https://godoc.org/github.com/hprose/hprose-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/hprose/hprose-golang)](https://goreportcard.com/report/github.com/hprose/hprose-golang)
[![codebeat badge](https://img.shields.io/badge/codebeat-A-398b39.svg)](https://codebeat.co/projects/github-com-hprose-hprose-golang)
[![Coverage Status](https://coveralls.io/repos/github/hprose/hprose-golang/badge.svg?branch=master)](https://coveralls.io/github/hprose/hprose-golang?branch=master)
[![License](https://img.shields.io/github/license/hprose/hprose-golang.svg)](http://opensource.org/licenses/MIT)


[Hprose 2.0 for Golang 中文文档](https://github.com/hprose/hprose-golang/wiki) 

>---
- **[Introduction](#introduction)**
- **[Installation](#installation)**
- **[Usage](#usage)**
    - **[Http Server](#http-server)**
        - [Based on net/http](#based-on-nethttp)
        - [Based on fasthttp](#based-on-fasthttp)
        - [Based on gin](#based-on-gin)
        - [Based on echo](#based-on-echo)
        - [Based on beego](#based-on-beego)
        - [Based on iris](#based-on-iris)
    - **[Http Client](#http-client)**
        - [Synchronous Invoking](#synchronous-invoking)
        - [Asynchronous Invoking](#asynchronous-invoking)
        - [Passing by reference parameters](#passing-by-reference-parameters)
    - **[Custom Struct](#custom-struct)**
		- [Field Alias of Custom Struct](#field-alias-of-custom-struct)
    - **[Hprose Proxy](#hprose-proxy)**
    - **[Simple Mode](#simple-mode)**
    - **[WebSocket Server & Client](#websocket-server--client)**
        - [WebSocket Server](#websocket-server)
        - [WebSocket Client](#websocket-client)
    - **[Unix Socket Server & Client](#unix-socket-server--client)**
        - [Unix Socket Server](#unix-socket-server)
        - [Unix Socket Client](#unix-socket-client)
- **[Benchmark](#benchmark)**

>---

## Introduction

*Hprose* is a High Performance Remote Object Service Engine.

It is a modern, lightweight, cross-language, cross-platform, object-oriented, high performance, remote dynamic communication middleware. It is not only easy to use, but powerful. You just need a little time to learn, then you can use it to easily construct cross language cross platform distributed application system.

*Hprose* supports many programming languages, for example:

* AAuto Quicker
* ActionScript
* ASP
* C++
* Dart
* Delphi/Free Pascal
* dotNET(C#, Visual Basic...)
* Golang
* Java
* JavaScript
* Node.js
* Objective-C
* Perl
* PHP
* Python
* Ruby
* ...

Through *Hprose*, You can conveniently and efficiently intercommunicate between those programming languages.

This project is the implementation of Hprose 2.0 for Golang.

## Installation

```sh
go get -u -v github.com/hprose/hprose-golang
```

## Usage

### Http Server

#### Based on net/http

```go
package main

import (
	"net/http"

	"github.com/hprose/hprose-golang/rpc"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	service := rpc.NewHTTPService()
	service.AddFunction("hello", hello)
	http.ListenAndServe(":8080", service)
}
```

#### Based on fasthttp

```go
package main

import (
	rpc "github.com/hprose/hprose-golang/rpc/fasthttp"
	"github.com/valyala/fasthttp"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	service := rpc.NewFastHTTPService()
	service.AddFunction("hello", hello)
	fasthttp.ListenAndServe(":8080", service.ServeFastHTTP)
}
```

#### Based on gin

```go
package main

import (
	"github.com/hprose/hprose-golang/rpc"
	"gopkg.in/gin-gonic/gin.v1"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	service := rpc.NewHTTPService()
	service.AddFunction("hello", hello)
	router := gin.Default()
	router.Any("/path", func(c *gin.Context) {
		service.ServeHTTP(c.Writer, c.Request)
	})
	router.Run(":8080")
}
```

#### Based on echo

```go
package main

import (
	"github.com/hprose/hprose-golang/rpc"
	"github.com/labstack/echo"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	service := rpc.NewHTTPService()
	service.AddFunction("hello", hello)
	e := echo.New()
	e.Any("/path", echo.WrapHandler(service))
	e.Start(":8080")
}
```

#### Based on beego

```go
package main

import (
	"github.com/astaxie/beego"
	"github.com/hprose/hprose-golang/rpc"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	service := rpc.NewHTTPService()
	service.AddFunction("hello", hello)
	beego.Handler("/path", service)
	beego.Run()
}
```

#### Based on iris

```go
package main

import (
	rpc "github.com/hprose/hprose-golang/rpc/fasthttp"
	"github.com/kataras/iris"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	service := rpc.NewFastHTTPService()
	service.AddFunction("hello", hello)
	iris.Any("/path", func(c *iris.Context) {
		service.ServeFastHTTP(c.RequestCtx)
	})
	iris.Listen(":8080")
}
```

### Http Client

#### Synchronous Invoking

```go
package main

import (
	"fmt"

	"github.com/hprose/hprose-golang/rpc"
)

type HelloService struct {
    Hello func(string) (string, error)
    Hello2 func(string) string `name:"hello"`
}

func main() {
	client := rpc.NewHTTPClient("http://127.0.0.1:8080/")
	var helloService *HelloService
	client.UseService(&helloService)
	fmt.Println(helloService.Hello("world"))
	fmt.Println(helloService.Hello2("world"))
}
```

Golang does not support function/method overload, but some other languages support. So hprose provides "Function/Method Alias" to invoke overloaded methods in other languages. You can also use it to invoke the same function/method with different names.

You just need define multiple func fields that correspond to the same remote method by the same `name` tag.

If an error (must be the last out parameter) returned by server-side function/method, or it panics in the server-side, the client will receive it. If the client func field defines an error out parameter (must be the last one), you can get the server-side error or panic from it. If the client func field has not defined an error out parameter, the client call will panic when receive the server-side error or panic.

#### Asynchronous Invoking

```go
package main

import (
	"fmt"
	"time"

	"github.com/hprose/hprose-golang/rpc"
)

type HelloService struct {
	Hello func(func(string, error), string)
	Hello2 func(func(string), string) `name:"hello"`
}

func main() {
	client := rpc.NewHTTPClient("http://127.0.0.1:8080/")
	var helloService *HelloService
	client.UseService(&helloService)
	helloService.Hello(func(result string, err error) {
		fmt.Println(result, err)
	}, "async world")
	helloService.Hello2(func(result string) {
		fmt.Println(result)
	}, "async world")
    time.Sleep(time.Second)
}
```

If the first in parameter is a callback function, you can invoke the remote method asynchronously.

The callback in parameters defines like the out parameters in synchronous invoking method. but if you omit the last error parameter, the asynchronous Invoking will NOT panic, the error will be ignore, too.

#### Passing by reference parameters

Hprose supports passing by reference parameters. The parameters must be pointer types and define func field with tag `byref:"true"`. For example:

```go
package main

import (
	"fmt"

	"github.com/hprose/hprose-golang/rpc"
)

func swap(a, b *int) {
	*b, *a = *a, *b
}

type SwapService struct {
	Swap func(a, b *int) error `byref:"true"`
}

func main() {
	server := rpc.NewTCPServer("")
	server.AddFunction("swap", swap)
	server.Handle()
	client := rpc.NewClient(server.URI())
	var swapService *SwapService
	client.UseService(&swapService)
	a := 1
	b := 2
	swapService.Swap(&a, &b)
	fmt.Println(a, b)
	client.Close()
	server.Close()
}
```

You will find that hprose also supports TCP server and client in this example.

### Custom Struct

You can transfer custom struct objects between hprose client and hprose server directly. Using the `Register` method to register your custom struct is the the only thing you need to do.

For example:

```go
package main

import (
    "fmt"
    "github.com/hprose/hprose-golang/io"
    "github.com/hprose/hprose-golang/rpc"
)

type TestUser struct {
    Name     string
    Sex      int
    Birthday time.Time
    Age      int
    Married  bool
}

type RemoteObject struct {
    GetUserList         func() []TestUser
}

func main() {
    io.Register(TestUser{}, "User")
    client := rpc.NewClient("http://www.hprose.com/example/")
    var ro *RemoteObject
    client.UseService(&ro)
    fmt.Println(ro.GetUserList())
}
```

The first argument of `Register` is an object or pointer of your custom struct. The second argument is the alias of your custom struct.

The real name of your custom struct can be different between the client and the server, as long as they registered the same alias.

The server of this example was written in PHP. In fact, You can use custom struct with golang server too.

#### Field Alias of Custom Struct

The first letter of the field name will be lowercased automatically when it is serialized by hprose. So we don't need to define a tag to implement this feature like JSON serialization when we interact with the other languages.

But it doesn't mean that hprose can't support to define field alias by tag. In fact, it can not only, and it can be compatible with the field alias definition in JSON serialization way. For example:

```go
type User struct {
	Name                         string `json:"n"`
	Age                          int    `json:"a"`
	ThisFieldWillNotBeSerialized string `json:"-"`
}

io.Register(User{}, "User", "json")
```

The struct above is defined for JSON serialization. But when we called `Register` by passing the third argument `"json"`, we can use the fields aliases defined in `json` tags for hprose serialization. If the field alias is `"-"`, it will be not serialized.

You can change the `json` tag to be anything else in the struct definition, such as `hprose`, as long as it is the same with the value of the `Register` third argument.

### Hprose Proxy

Hprose supports publishing a special method: MissingMethod. All methods not explicitly published will be redirected to the method. You can use it to implement an hprose proxy. And hprose provides an ResultMode options to improve performance of the proxy server. You can use it like this:

```go
package main

import (
	"net/http"
	"reflect"

	"github.com/hprose/hprose-golang/rpc"
)

type HproseProxy struct {
	client   rpc.Client
	settings rpc.InvokeSettings
}

func newHproseProxy() *HproseProxy {
	proxy := new(HproseProxy)
	proxy.client = rpc.NewClient("http://www.hprose.com/example/")
	proxy.settings = rpc.InvokeSettings{
		Mode:        rpc.Raw,
		ResultTypes: []reflect.Type{reflect.TypeOf(([]byte)(nil))},
	}
	return proxy
}

func (proxy *HproseProxy) Proxy(
	name string, args []reflect.Value, context rpc.Context) ([]reflect.Value, error) {
	return proxy.client.Invoke(name, args, &proxy.settings)
}

func main() {
	service := rpc.NewHTTPService()
	service.AddMissingMethod(newHproseProxy().Proxy, rpc.Options{Mode: rpc.Raw})
	http.ListenAndServe(":8080", service)
}
```

You can also define func field with tag `mode` in client, and the return value must be `[]byte`. The server result mode option is setting by `Options` parameter.

The ResultMode have 4 values:

* Normal
* Serialized
* Raw
* RawWithEndTag

The `Normal` result mode is the default value.

In `Serialized` result mode, the returned value is an hprose serialized data in []byte, but the arguments and exception will be parsed to the normal value.

In `Raw` result mode, all the reply will be returned directly to the result in []byte, but the result data doesn't have the hprose `TagEnd`.

The `RawWithEndTag` is similar to the `Raw` result mode, but it has the hprose `TagEnd`.

With the ResultMode option, you can store, cache and forward the result in the original format.

### Simple Mode

By default, the data between the hprose client and server can be passed with internal references. if your data have no internal references, you can open the simple mode to improve performance.

```go
package main

import (
	"fmt"

	"github.com/hprose/hprose-golang/rpc"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

type HelloService struct {
	Hello func(string) (string, error) `simple:"true"`
}

func main() {
	server := rpc.NewTCPServer("")
	server.AddFunction("hello", hello, rpc.Options{Simple: true})
	server.Handle()
	client := rpc.NewClient(server.URI())
	var helloService *HelloService
	client.UseService(&helloService)
	fmt.Println(helloService.Hello("World"))
	client.Close()
	server.Close()
}
```

### WebSocket Server & Client

#### WebSocket Server

```go
package main

import (
	"net/http"
	"runtime"

	rpc "github.com/hprose/hprose-golang/rpc/websocket"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	service := rpc.NewWebSocketService()
	service.AddFunction("hello", hello)
	http.ListenAndServe(":8080", service)
}
```

#### WebSocket Client

```go
package main

import (
	"fmt"

	rpc "github.com/hprose/hprose-golang/rpc/websocket"
)

type HelloService struct {
    Hello func(string) (string, error)
}

func main() {
	client := rpc.NewWebSocketClient("ws://127.0.0.1:8080/")
	var helloService *HelloService
	client.UseService(&helloService)
	fmt.Println(helloService.Hello("world"))
}
```

### Unix Socket Server & Client

#### Unix Socket Server

```go
package main

import (
	"net/http"
	"runtime"

	"github.com/hprose/hprose-golang/rpc"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	server := rpc.NewUnixServer("unix:///tmp/hprose.sock")
	server.AddFunction("hello", hello)
	server.Start()
}
```

#### Unix Socket Client

```go
package main

import (
	"fmt"

	"github.com/hprose/hprose-golang/rpc"
)

type HelloService struct {
    Hello func(string) (string, error)
}

func main() {
	client := rpc.NewUnixClient("unix:///tmp/hprose.sock")
	var helloService *HelloService
	client.UseService(&helloService)
	fmt.Println(helloService.Hello("world"))
}
```

## Benchmark

Hprose is faster than golang RPC, you can run benchmark like this:

```
go test --bench=".*" github.com/hprose/hprose-golang/examples/bench
```

* go1.7.1 darwin/amd64
* macOS Sierra
* iMac (Retina 5K, 27-inch, Late 2015)
* CPU 4GHz Intel Core i7
* Memory 32 GB 1867 MHz DDR3

```
BenchmarkParallelHprose2-8       	  200000	     11230 ns/op
BenchmarkParallelHprose2Unix-8   	  300000	      5234 ns/op
BenchmarkParallelGobRPC-8        	  100000	     16675 ns/op
BenchmarkParallelGobRPCUnix-8    	  200000	      6798 ns/op
BenchmarkParallelJSONRPC-8       	  100000	     17261 ns/op
BenchmarkParallelJSONRPCUnix-8   	  200000	      7917 ns/op
```

```
BenchmarkHprose2-8               	   50000	     34287 ns/op
BenchmarkHprose2Unix-8           	  200000	     11470 ns/op
BenchmarkGobRPC-8                	   30000	     45576 ns/op
BenchmarkGobRPCUnix-8            	  100000	     24216 ns/op
BenchmarkJSONRPC-8               	   30000	     51298 ns/op
BenchmarkJSONRPCUnix-8           	   50000	     27408 ns/op
```
