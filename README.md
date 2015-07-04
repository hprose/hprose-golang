# Hprose for Golang

[![Join the chat at https://gitter.im/hprose/hprose-go](https://img.shields.io/badge/GITTER-join%20chat-green.svg)](https://gitter.im/hprose/hprose-go?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![GoDoc](https://godoc.org/github.com/hprose/hprose-go?status.svg&style=flat)](https://godoc.org/github.com/hprose/hprose-go)
[![Build Status](https://drone.io/github.com/hprose/hprose-go/status.png)](https://drone.io/github.com/hprose/hprose-go/latest)

>---
- **[Introduction](#introduction)**
- **[Installation](#installation)**
- **[Usage](#usage)**
    - **[Http Server](#http-server)**
    - **[Http Client](#http-client)**
        - [Synchronous Invoking](#synchronous-invoking)
        - [Synchronous Exception Handling](#synchronous-exception-handling)
        - [Asynchronous Invoking](#asynchronous-invoking)
        - [Asynchronous Exception Handling](#asynchronous-exception-handling)
        - [Function Alias](#function-alias)
        - [Passing by reference parameters](#passing-by-reference-parameters)
    - **[Custom Struct](#custom-struct)**
		- [Field Alias of Custom Struct](#field-alias-of-custom-struct)
    - **[Hprose Proxy](#hprose-proxy)**
        - [Better Proxy](#better-proxy)
    - **[Simple Mode](#simple-mode)**
    - **[Missing Method](#missing-method)**
    - **[TCP Server and Client](#tcp-server-and-client)**
    - **[Unix Server and Client](#unix-server-and-client)**
	- **[WebSocket Server and Client](#websocket-server-and-client)**
    - **[Service Event](#service-event)**
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

This project is the implementation of Hprose for Golang.

## Installation

```sh
go get github.com/hprose/hprose-go
go install github.com/hprose/hprose-go
```

## Usage

### Http Server

Hprose for Golang is very easy to use. You can create a hprose http server like this:

```go
package main

import (
	"errors"
	"github.com/hprose/hprose-go"
	"net/http"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

type myService struct{}

func (myService) Swap(a int, b int) (int, int) {
	return b, a
}

func (myService) Sum(args ...int) (int, error) {
	if len(args) < 2 {
		return 0, errors.New("Requires at least two parameters")
	}
	a := args[0]
	for i := 1; i < len(args); i++ {
		a += args[i]
	}
	return a, nil
}

func main() {
	service := hprose.NewHttpService()
	service.AddFunction("hello", hello)
	service.AddMethods(myService{})
	http.ListenAndServe(":8080", service)
}
```

You can publish multi-valued functions/methods, the multi-valued result will be automatically converted to an array result.

### Http Client

#### Synchronous Invoking

Then you can create a hprose http client to invoke it like this:

```go
package main

import (
	"fmt"
	"github.com/hprose/hprose-go"
)

type clientStub struct {
	Hello func(string) string
	Swap  func(int, int) (int, int)
	Sum   func(...int) (int, error)
}

func main() {
	client := hprose.NewClient("http://127.0.0.1:8080/")
	var ro *clientStub
	client.UseService(&ro)
	fmt.Println(ro.Hello("World"))
	fmt.Println(ro.Swap(1, 2))
	fmt.Println(ro.Sum(1, 2, 3, 4, 5))
	fmt.Println(ro.Sum(1))
}
```

#### Synchronous Exception Handling

Client stubs do not have exactly the same with the server-side interfaces. For example:

```go
package main

import (
	"fmt"
	"github.com/hprose/hprose-go"
)

type clientStub struct {
	Sum   func(...int) int
}

func main() {
	client := hprose.NewClient("http://127.0.0.1:8080/")
	var ro *clientStub
	client.UseService(&ro)
	fmt.Println(ro.Sum(1, 2, 3, 4, 5))
	fmt.Println(ro.Sum(1))
}
```

If an error (must be the last out parameter) returned by server-side function/method, or it panics in the server-side, the client will receive it. If the client stub has an error out parameter (also must be the last one), you can get the server-side error or panic from it. If the client stub have not define an error out parameter, the client stub will panic when receive the server-side error or panic.

#### Asynchronous Invoking

Hprose for golang supports golang style asynchronous invoke. It does not require a callback function, but need to define the channel out parameters. for example:

```go
package main

import (
	"fmt"
	"github.com/hprose/hprose-go"
)

type clientStub struct {
	Sum func(...int) (<-chan int, <-chan error)
}

func main() {
	client := hprose.NewClient("http://127.0.0.1:8080/")
	var ro *clientStub
	client.UseService(&ro)
	sum, err := ro.Sum(1, 2, 3, 4, 5)
	fmt.Println(<-sum, <-err)
	sum, err = ro.Sum(1)
	fmt.Println(<-sum, <-err)
}
```

#### Asynchronous Exception Handling

When using asynchronous invoking, you need to define a `<-chan error` out parameter (also the last one) to receive the server-side error or panic (or exception in other languages). If you omit this parameter, the client will ignore the exception, like never happened.

For example:

```go
package main

import (
    "fmt"
    "github.com/hprose/hprose-go"
)

type clientStub struct {
    Sum func(...int) (<-chan int)
}

func main() {
    client := hprose.NewClient("http://127.0.0.1:8080/")
    var ro *clientStub
    client.UseService(&ro)
    fmt.Println(<-ro.Sum(1))
}
```
You will get the result `0`, but do not know what happened.

#### Function Alias

Golang does not support function/method overload, but some other languages support. So hprose provides "Function/Method Alias" to invoke overloaded methods in other languages. You can also use it to invoke the same function/method with different names.

For example:

```go
package main

import (
    "fmt"
    "github.com/hprose/hprose-go"
)

type clientStub struct {
    Hello      func(string) string
    AsyncHello func(string) <-chan string `name:"hello"`
}

func main() {
    client := hprose.NewClient("http://127.0.0.1:8080/")
    var ro *clientStub
    client.UseService(&ro)
    fmt.Println(ro.Hello("Synchronous Invoking"))
    fmt.Println(<-ro.AsyncHello("Asynchronous Invoking"))
}
```

The real remote function/method name is specified in the function field tag.

#### Passing by reference parameters

Hprose supports passing by reference parameters. The parameters must be pointer types. Open this option also in the function field tag. For example:

```go
package main

import (
    "fmt"
    "github.com/hprose/hprose-go"
)

type clientStub struct {
    Swap func(*map[string]string) `name:"swapKeyAndValue" byref:"true"`
}

func main() {
    client := hprose.NewClient("http://hprose.com/example/")
    var ro *clientStub
    client.UseService(&ro)
    m := map[string]string{
        "Jan": "January",
        "Feb": "February",
        "Mar": "March",
        "Apr": "April",
        "May": "May",
        "Jun": "June",
        "Jul": "July",
        "Aug": "August",
        "Sep": "September",
        "Oct": "October",
        "Nov": "November",
        "Dec": "December",
    }
    fmt.Println(m)
    ro.Swap(&m)
    fmt.Println(m)
}
```

The server of this example was written in PHP. In fact, You can use any language which hprose supported to write the server.

### Custom Struct

You can transfer custom struct objects between hprose client and hprose server directly.

Using the `ClassManager.Register` method to register your custom struct is the the only thing you need to do. For example:

```go
package main

import (
    "fmt"
    "github.com/hprose/hprose-go"
)

type TestUser struct {
    Name     string
    Sex      int
    Birthday time.Time
    Age      int
    Married  bool
}

type remoteObject struct {
    GetUserList         func() []TestUser
}

func main() {
    ClassManager.Register(reflect.TypeOf(TestUser{}), "User")
    client := NewClient("http://www.hprose.com/example/")
    var ro *remoteObject
    client.UseService(&ro)

    fmt.Println(ro.GetUserList())
}
```

The first argument of `ClassManager.Register` is the type of your custom struct. The second argument is the alias of your custom struct.

The real name of your custom struct can be different between the client and the server, as long as they registered the same alias.

The server of this example was written in PHP. In fact, You can use custom struct with go server too.

#### Field Alias of Custom Struct

The first letter of the field name will be lowercased automatically when it is serialized. So we don't need to define a tag to implement this feature like Json serialization when we interact with the other languages.

But it doesn't mean that hprose can't support to define field alias by tag. In fact, it can not only, and it can be compatible with the field alias definition in Json serialization way. For example:

```go
...
type User struct {
	Name string `json:"n"`
	Age  int    `json:"a"`
	OOXX string `json:"-"`
}
...
hprose.ClassManager.Register(reflect.TypeOf(User{}), "User", "json")
...
```

The struct above is defined for Json serialization. But when we called `ClassManager.Register` by passing the third argument `"json"`, we can use the fields aliases defined in `json` tags for hprose serialization. If the field alias is `"-"`, it will be not serialized.

You can change the `json` tag to be anything else in the struct definition, such as `hprose`, as long as it is the same with the value of the `ClassManager.Register` third argument.

### Hprose Proxy

You can use hprose server and client to create a hprose proxy server. All requests sent to the hprose proxy server will be forwarded to the backend hprose server. For example:

```go
package main

import (
    "github.com/hprose/hprose-go"
    "net/http"
)

type proxyStub struct {
    Hello func(string) (string, error)
    Swap  func(int, int) (int, int)
    Sum   func(...int) (int)
}

func main() {
    client := hprose.NewClient("http://127.0.0.1:8080/")
    var ro *proxyStub
    client.UseService(&ro)
    service := hprose.NewHttpService()
    service.AddMethods(ro)
    http.ListenAndServe(":8181", service)
}
```
Whether the definition of the error out parameter does not matter, the exception will be automatically forwarded.

#### Better Proxy

Hprose provides an ResultMode options to improve performance of the proxy server. You can use it like this:

```go
package main

import (
    "github.com/hprose/hprose-go"
    "net/http"
)

type proxyStub struct {
    Hello func(string) []byte   `result:"raw"`
    Swap  func(int, int) []byte `result:"raw"`
    Sum   func(...int) []byte   `result:"raw"`
}

func main() {
    client := hprose.NewClient("http://127.0.0.1:8080/")
    var ro *proxyStub
    client.UseService(&ro)
    service := hprose.NewHttpService()
    service.AddMethods(ro, hprose.Raw)
    http.ListenAndServe(":8181", service)
}
```

The client result mode option is setting in the func field tag, and the return value must be `[]byte`. The server result mode option is setting by `AddMethods` parameter.

The ResultMode have 4 values:
* Normal
* Serialized
* Raw
* RawWithEndTag

The `Normal` result mode is the default value.

In `Serialized` result mode, the returned value is a hprose serialized data in []byte, but the arguments and exception will be parsed to the normal value.

In `Raw` result mode, all the reply will be returned directly to the result in []byte, but the result data doesn't have the hprose end tag.

The `RawWithEndTag` is similar to the `Raw` result mode, but it has the hprose end tag.

With the ResultMode option, you can store, cache and forward the result in the original format.

### Simple Mode

By default, the data between the hprose client and server can be passed with internal references. if your data have no internal references, you can open the simple mode to improve performance.

You can open simple mode in server like this:

```go
package main

import (
    "github.com/hprose/hprose-go"
    "net/http"
)

func hello(name string) string {
    return "Hello " + name + "!"
}

func main() {
    service := hprose.NewHttpService()
    service.AddFunction("hello", hello, true)
    http.ListenAndServe(":8080", service)
}
```

The option parameter `true` is the simple mode switch. The result will be transmitted to the client in simple mode when it is on.

To open the client simple mode is like this:

```go
package main

import (
    "fmt"
    "github.com/hprose/hprose-go"
)

type clientStub struct {
    Hello func(string) string       `simple:"true"`
    Swap  func(int, int) (int, int) `simple:"true"`
    Sum   func(...int) (int, error)
}

func main() {
    client := hprose.NewClient("http://127.0.0.1:8181/")
    var ro *clientStub
    client.UseService(&ro)
    fmt.Println(ro.Hello("World"))
    fmt.Println(ro.Swap(1, 2))
    fmt.Println(ro.Sum(1, 2, 3, 4, 5))
    fmt.Println(ro.Sum(1))
}
```

The arguments will be transmitted to the server in simple mode when it is on.

### Missing Method

Hprose supports publishing a special method: MissingMethod. All methods not explicitly published will be redirected to the method. For example:

```go
package main

import (
    "github.com/hprose/hprose-go"
    "net/http"
    "reflect"
    "strings"
)

func hello(name string) string {
    return "Hello " + name + "!"
}

func missing(name string, args []reflect.Value) (result []reflect.Value) {
    result = make([]reflect.Value, 1)
    switch strings.ToLower(name) {
    case "add":
        result[0] = reflect.ValueOf(args[0].Interface().(int) + args[1].Interface().(int))
    case "sub":
        result[0] = reflect.ValueOf(args[0].Interface().(int) - args[1].Interface().(int))
    case "mul":
        result[0] = reflect.ValueOf(args[0].Interface().(int) * args[1].Interface().(int))
    case "div":
        result[0] = reflect.ValueOf(args[0].Interface().(int) / args[1].Interface().(int))
    default:
        panic("The method '" + name + "' is not implemented.")
    }
    return
}

func main() {
    service := hprose.NewHttpService()
    service.AddFunction("hello", hello, true)
    service.AddMissingMethod(missing, true)
    http.ListenAndServe(":8080", service)
}
```

If you want return an error to the client, please use panic. The error type return value can't be processed in the method.

The simple mode and the result mode options can also be used with it.

Invoking the missing method makes no difference with the normal method. For example:

```go
package main

import (
    "fmt"
    "github.com/hprose/hprose-go"
)

type clientStub struct {
    Add   func(int, int) int
    Sub   func(int, int) int
    Mul   func(int, int) int
    Div   func(int, int) int
    Power func(int, int) (int, error)
}

func main() {
    client := hprose.NewClient("http://127.0.0.1:8080/")
    var ro *clientStub
    client.UseService(&ro)
    fmt.Println(ro.Add(1, 2))
    fmt.Println(ro.Sub(1, 2))
    fmt.Println(ro.Mul(1, 2))
    fmt.Println(ro.Div(1, 2))
    fmt.Println(ro.Power(1, 2))
}
```

The result is:

```
3
-1
2
0
0 The method 'Power' is not implemented.
```
### TCP Server and Client

Hprose for Golang supports TCP Server and Client. It is very easy to use like the HTTP Server and Client.

To create a hprose TCP server, you can use `NewTcpService` or `NewTcpServer`.

To use `NewTcpService`, you need call the ServeTCP method and passing the TCP Connection to it.

using `NewTcpServer` is easier than `NewTcpService`. For example:

```go
    ...
    server := hprose.NewTcpServer("tcp://127.0.0.1:1234/")
    server.AddFunction("hello", hello)
    server.Start()
    ...
```

To create a hprose TCP client is the same as HTTP client:

```go
    ...
    client := hprose.NewClient("tcp://127.0.0.1:1234/")
    ...
```

You can also specify `tcp4://` scheme to using ipv4 or `tcp6://` scheme to using ipv6.

### Unix Server and Client

Hprose for Golang supports Unix Socket Server and Client. It is very easy to use like the Tcp Server and Client.

To create a hprose Unix server, you can use `NewUnixService` or `NewUnixServer`.

To use `NewUnixService`, you need call the ServeUnix method and passing the Unix Connection to it.

using `NewUnixServer` is easier than `NewUserService`. For example:

```go
    ...
    server := hprose.NewUnixServer("unix:/tmp/my.sock")
    server.AddFunction("hello", hello)
    server.Start()
    ...
```

To create a hprose Unix client is the same as TCP client:

```go
    ...
    client := hprose.NewClient("unix:/tmp/my.sock")
    ...
```

### WebSocket Server and Client

Hprose for Golang supports WebSocket Socket Server and Client. It is very easy to use like the HTTP Server and Client.

To create a hprose WebSocket service, you can use `NewWebSocketService`. For example:

```go
    ...
	service := hprose.NewWebSocketService()
    service.AddFunction("hello", hello, true)
    http.ListenAndServe(":8080", service)
    ...
```

Hprose WebSocket service is also HTTP service. Not only hprose WebSocket client can invoke it, but also hprose HTTP client.

To create a hprose WebSocket client is the same as HTTP client:

```go
    ...
    client := hprose.NewClient("ws://127.0.0.1:8080/")
    ...
```

### Service Event

Hprose defines a `ServiceEvent` interface.

```go
type ServiceEvent interface {}
```

It is a empty interface, but you can add the following event methods to your implementation:

```go
    OnBeforeInvoke(name string, args []reflect.Value, byref bool, context hprose.Context)
	OnBeforeInvoke(name string, args []reflect.Value, byref bool, context hprose.Context) error
    OnAfterInvoke(name string, args []reflect.Value, byref bool, result []reflect.Value, context hprose.Context)
	OnAfterInvoke(name string, args []reflect.Value, byref bool, result []reflect.Value, context hprose.Context) error
    OnSendError(err error, context hprose.Context)
	OnSendError(err error, context hprose.Context) error
```

`OnBeforeInvoke`, `OnAfterInvoke` and `OnSendError` all have two kinds of definitions. You just need and only can implement one of them.

For example, if you want to log some thing about the service, you can do it like this:

```go
package main

import (
    "fmt"
    "github.com/hprose/hprose-go"
    "net/http"
    "reflect"
)

func hello(name string) string {
    return "Hello " + name + "!"
}

type myServiceEvent struct{}

func (myServiceEvent) OnBeforeInvoke(name string, args []reflect.Value, byref bool, context hprose.Context) {
    fmt.Println(name, args, byref)
}

func (myServiceEvent) OnAfterInvoke(name string, args []reflect.Value, byref bool, result []reflect.Value, context hprose.Context) {
    fmt.Println(name, args, byref, result)
}

func (myServiceEvent) OnSendError(err error, context hprose.Context) {
    fmt.Println(err)
}

func main() {
    service := hprose.NewHttpService()
    service.ServiceEvent = myServiceEvent{}
    service.AddFunction("hello", hello)
    http.ListenAndServe(":8080", service)
}
```

The `TcpService` and `TcpServer` also have this interface field.

For hprose HTTP Service, there are two more `OnSendHeader` event methods:

```go
	OnSendHeader(context Context)
    OnSendHeader(context *hprose.HttpContext)
```

You also just need to implement only one of them.

## Benchmark

Hprose is faster than golang RPC, you can run benchmark like this:

```
go test --bench=".*" github.com/hprose/hprose-go/bench
```

Here is the result with Go 1.4 on an Intel i7-2600:

benchmark | iter | time/iter
:------|------:|------:|
BenchmarkHprose| 30000|46696 ns/op
BenchmarkHprose2| 30000|48215 ns/op
BenchmarkGobRPC| 20000|66818 ns/op
BenchmarkJSONRPC| 10000|104709 ns/op
