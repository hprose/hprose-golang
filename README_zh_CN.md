# Hprose for Golang

[![Build Status](https://drone.io/github.com/hprose/hprose-go/status.png)](https://drone.io/github.com/hprose/hprose-go/latest)

> ---
- **[简介](#简介)**
- **[安装](#安装)**
- **[使用](#使用)**
	- **[Http 服务器](#http-服务器)**
	- **[Http 客户端](#http-客户端)**
		- [同步调用](#同步调用)
		- [同步异常处理](#同步异常处理)
		- [异步调用](#异步调用)
		- [异步异常处理](#异步异常处理)
		- [函数方法别名](#函数方法别名)
		- [引用参数传递](#引用参数传递)
	- **[Hprose 代理](#hprose-代理)**
		- [更好的代理](#更好的代理)
	- **[简单模式](#简单模式)**
	- **[缺失的方法](#缺失的方法)**
	- **[TCP 服务器和客户端](#tcp-服务器和客户端)**
	- **[服务事件](#服务事件)**

> ---

## 简介

*Hprose* 是高性能远程对象服务引擎（High Performance Remote Object Service Engine）的缩写。

它是一个先进的轻量级的跨语言跨平台面向对象的高性能远程动态通讯中间件。它不仅简单易用，而且功能强大。你只需要稍许的时间去学习，就能用它轻松构建跨语言跨平台的分布式应用系统了。

*Hprose* 支持众多编程语言，例如：

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

通过 *Hprose*，你就可以在这些语言之间方便高效的实现互通了。

本项目是 Hprose 的 Golang 语言版本实现。

## 安装

```sh
go get github.com/hprose/hprose-go/hprose
```

## 使用

### Http 服务器

Hprose for Golang 使用起来很简单，你可以像这样来创建一个 Hprose 的 http 服务:

```go
package main

import (
	"errors"
	"github.com/hprose/hprose-go/hprose"
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

你可以发布多值返回函数和方法，多值返回结果会自动转换为一个数组类型的结果。

### Http 客户端

#### 同步调用

然后你可以创建一个 Hprose 的 http 客户端来调用它了，就像这样：

```go
package main

import (
	"fmt"
	"github.com/hprose/hprose-go/hprose"
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

#### 同步异常处理

客户端接口通过 struct 的函数字段的方式来定义，这些函数接口不需要完全跟服务器端的接口一致，例如：

```go
package main

import (
	"fmt"
	"github.com/hprose/hprose-go/hprose"
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

如果服务器端返回一个错误（必须是通过最后一个输出参数），或者是服务器端产生了 panic（在其他的语言中就是抛出异常），客户端将会收到它。如果客户端函数接口中包含有一个错误输出参数（也必须是最后一个），你可以通过它来得到服务器端的错误或 panic（异常）。如果客户端没有定义错误输出参数，那么客户端在收到服务器端错误或 panic（异常）之后，将会在客户端产生 panic。

#### 异步调用

Hprose for golang 支持 golang 风格的异步调用。它不需要回调函数，但是需要定义通道型的输出参数。例如：

```go
package main

import (
	"fmt"
	"github.com/hprose/hprose-go/hprose"
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

#### 异步异常处理

当使用异步调用时，你需要定义一个 `<-chan error` 型的输出参数（也必须是最后一个）来接收服务器端的错误和 panic（或其它语言中的异常）。如果你省略了该参数，客户端也会忽略异常，就像从来没发生过一样。

例如：

```go
package main

import (
    "fmt"
    "github.com/hprose/hprose-go/hprose"
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

你将会得到结果 `0`，并且不会知道发生了什么。

#### 函数方法别名

Golang 本身不支持函数/方法的重载，但是其它一些语言支持。所以 Hprose 提供了 “函数/方法 别名” 来调用其它语言中的重载方法。你也可以使用它来通过不同的名字调用同一个函数或方法。

例如：

```go
package main

import (
    "fmt"
    "github.com/hprose/hprose-go/hprose"
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

远程方法或函数的真实名字在函数字段的 tag 中指定就可以了。

#### 引用参数传递

Hprose 还支持引用参数传递。在进行引用参数传递时，参数必须是指针类型（因为非指针类型没法被修改）。开启该选项也是通过在函数字段的 tag 中指定的。例如：

```go
package main

import (
    "fmt"
    "github.com/hprose/hprose-go/hprose"
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

这个例子中的服务器是用PHP编写的。实际上，你可以使用任何 Hprose 支持的语言来编写服务器，对于客户端调用上没有区别。

### Hprose 代理

你可以通过 Hprose 服务器和客户端来为 Hprose 创建代理服务器。所有的发送到 Hprose 代理服务器上的请求都将被转发到后端的 hprose 服务器上。例如：

```go
package main

import (
    "github.com/hprose/hprose-go/hprose"
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

不管是否定义了错误输出参数，异常都会被自动转发。

#### 更好的代理

Hprose 提供了结果模式选项来改进代理服务器的性能。你可以像这样来使用它：

```go
package main

import (
    "github.com/hprose/hprose-go/hprose"
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

客户端结果模式选项在函数字段的 tag 中设置，客户端接口的返回值必须是 `[]byte` 类型。服务器端的结果模式选项在 `AddMethods` 方法的参数中设置（其它几个 AddXXX 方法同样可以设置这个参数）。

结果模式包含有四个值：
* Normal
* Serialized
* Raw
* RawWithEndTag

`Normal` 是默认值。

在 `Serialized` 结果模式下，返回值是一个hprose序列化的数据，以 `[]byte` 类型返回，（即对返回结果不做解析）。但是参数和异常将被解析为正常值。

在 `Raw` 结果模式下，所有的应答信息都将直接以 `[]byte` 类型返回。但结果数据中不包含 Hprose 终结符。

`RawWithEndTag` 与 `Raw` 模式类似，但是它包含 Hprose 终结符。

通过结果模式选项，你可以以原始格式来存储，缓存和转发结果数据。

如果你愿意的话，你还可以将 Hprose 结合 memcache 之类的服务来实现更高效率的 Hprose 代理服务器。

> 注：客户端和服务器端的结果模式是相互独立的，你不需要同时在服务器端和客户端开启结果模式，只在一边设置为结果模式仍然可以正常通讯，不会有任何影响。

### 简单模式

在默认情况下，在客户端和服务器端传输的数据是可以包含有内部引用的复杂数据，这样可以解决用 json 格式无法传递的循环引用数据的问题，同时引用对于复杂数据来说可以起到有效的压缩效果。但如果你的数据没有包含内部引用，那么你可以开启简单模式来进一步改善性能。

你可以像这样在服务器端开启简单模式：

```go
package main

import (
    "github.com/hprose/hprose-go/hprose"
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

选项参数 `true` 表示打开简单模式开关，这样结果将以简单模式返回给客户端。

在客户端可以这样打开简单模式：

```go
package main

import (
    "fmt"
    "github.com/hprose/hprose-go/hprose"
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

在开启简单模式的情况下，参数将以简单模式传递给服务器。

> 注：客户端和服务器端的简单模式也是相互独立的，你不需要同时在服务器端和客户端开启简单模式，只在一边设置为简单模式仍然可以正常通讯，不会有任何影响。

### 缺失的方法

Hprose 支持发布一个特殊的方法：MissingMethod。所有对没有显式发布的方法的调用都将被重定向到它上面。例如：

```go
package main

import (
    "github.com/hprose/hprose-go/hprose"
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

如果你想返回一个错误给客户端，请使用panic。因为在该方法中返回的 error 类型数据不会被特殊处理。

简单模式和结果模式也可以在它上面使用，因此通过它你可以构建更通用的 Hprose 代理服务器。

对客户端来说，调用 `AddMissingMethod` 发布的方法跟调用普通方法没有任何区别。例如：

```go
package main

import (
    "fmt"
    "github.com/hprose/hprose-go/hprose"
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

结果为：

```
3
-1
2
0
0 The method 'Power' is not implemented.
```

### TCP 服务器和客户端

Hprose for Golang 已经支持 TCP 的服务器和客户端。它跟 HTTP 版本的服务器和客户端在使用上一样简单。

你可以使用 `NewTcpService` 或 `NewTcpServer`，来创建 Hprose 的 TCP 服务器。

使用 `NewTcpService`，你需要调用它的 `ServeTCP` 方法传入 TCP 连接。

使用 `NewTcpServer` 比 `NewTcpService` 则要简单的多。例如：

```go
    ...
    server := hprose.NewTcpServer("tcp://127.0.0.1:1234/")
    server.AddFunction("hello", hello)
    server.Start()
    ...
```

创建 Hprose 的 TCP 客户端跟 HTTP 客户端是一样的方式：

```go
    ...
    client := hprose.NewClient("tcp://127.0.0.1:1234/")
    ...
```

你也可以指定 `tcp4://` 方案来使用 ipv4 或 `tcp6://` 方案来使用 ipv6。

### 服务事件

Hprose 定义了一个 `ServiceEvent` 接口。

```go
type ServiceEvent interface {
    OnBeforeInvoke(name string, args []reflect.Value, byref bool, context interface{})
    OnAfterInvoke(name string, args []reflect.Value, byref bool, result []reflect.Value, context interface{})
    OnSendError(err error, context interface{})
}
```

如果你想针对服务器的一些行为做日志的话，你可以实现这个接口，例如：

```go
package main

import (
    "fmt"
    "github.com/hprose/hprose-go/hprose"
    "net/http"
    "reflect"
)

func hello(name string) string {
    return "Hello " + name + "!"
}

type myServiceEvent struct{}

func (myServiceEvent) OnBeforeInvoke(name string, args []reflect.Value, byref bool, context interface{}) {
    fmt.Println(name, args, byref)
}

func (myServiceEvent) OnAfterInvoke(name string, args []reflect.Value, byref bool, result []reflect.Value, context interface{}) {
    fmt.Println(name, args, byref, result)
}

func (myServiceEvent) OnSendError(err error, context interface{}) {
    fmt.Println(err)
}

func main() {
    service := hprose.NewHttpService()
    service.ServiceEvent = myServiceEvent{}
    service.AddFunction("hello", hello)
    http.ListenAndServe(":8080", service)
}
```

`TcpService` 和 `TcpServer` 同样包含这个接口字段。

另外，针对 Hprose HTTP 服务器，你还可以单独实现 `HttpServiceEvent` 接口，这个接口多了一个针对 Http 头的事件。

```go
type HttpServiceEvent interface {
	ServiceEvent
	OnSendHeader(response http.ResponseWriter, request *http.Request)
}
```

它的实现同样是赋值给 `service.ServiceEvent` 字段就可以了。
