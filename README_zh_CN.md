# Hprose for Golang

[![Join the chat at https://gitter.im/hprose/hprose-go](https://img.shields.io/badge/GITTER-join%20chat-green.svg)](https://gitter.im/hprose/hprose-go?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![GoDoc](https://godoc.org/github.com/hprose/hprose-go?status.svg&style=flat)](https://godoc.org/github.com/hprose/hprose-go)
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
	- **[自定义结构体](#自定义结构体)**
		- [自定义结构体的字段别名](#自定义结构体的字段别名)
	- **[Hprose 代理](#hprose-代理)**
		- [更好的代理](#更好的代理)
	- **[简单模式](#简单模式)**
	- **[缺失的方法](#缺失的方法)**
	- **[TCP 服务器和客户端](#tcp-服务器和客户端)**
	- **[Unix 服务器和客户端](#unix-服务器和客户端)**
	- **[WebSocket 服务器和客户端](#websocket-服务器和客户端)**
	- **[服务事件](#服务事件)**
- **[性能测试](#性能测试)**

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
go get github.com/hprose/hprose-go
go install github.com/hprose/hprose-go
```

## 使用

### Http 服务器

Hprose for Golang 使用起来很简单，你可以像这样来创建一个 Hprose 的 http 服务:

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

你可以发布多值返回函数和方法，多值返回结果会自动转换为一个数组类型的结果。

### Http 客户端

#### 同步调用

然后你可以创建一个 Hprose 的 http 客户端来调用它了，就像这样：

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

#### 同步异常处理

客户端接口通过 struct 的函数字段的方式来定义，这些函数接口不需要完全跟服务器端的接口一致，例如：

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

如果服务器端返回一个错误（必须是通过最后一个输出参数），或者是服务器端产生了 panic（在其他的语言中就是抛出异常），客户端将会收到它。如果客户端函数接口中包含有一个错误输出参数（也必须是最后一个），你可以通过它来得到服务器端的错误或 panic（异常）。如果客户端没有定义错误输出参数，那么客户端在收到服务器端错误或 panic（异常）之后，将会在客户端产生 panic。

#### 异步调用

Hprose for golang 支持 golang 风格的异步调用。它不需要回调函数，但是需要定义通道型的输出参数。例如：

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

#### 异步异常处理

当使用异步调用时，你需要定义一个 `<-chan error` 型的输出参数（也必须是最后一个）来接收服务器端的错误和 panic（或其它语言中的异常）。如果你省略了该参数，客户端也会忽略异常，就像从来没发生过一样。

例如：

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

你将会得到结果 `0`，并且不会知道发生了什么。

#### 函数方法别名

Golang 本身不支持函数/方法的重载，但是其它一些语言支持。所以 Hprose 提供了 “函数/方法 别名” 来调用其它语言中的重载方法。你也可以使用它来通过不同的名字调用同一个函数或方法。

例如：

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

远程方法或函数的真实名字在函数字段的 tag 中指定就可以了。

#### 引用参数传递

Hprose 还支持引用参数传递。在进行引用参数传递时，参数必须是指针类型（因为非指针类型没法被修改）。开启该选项也是通过在函数字段的 tag 中指定的。例如：

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

这个例子中的服务器是用 PHP 编写的。实际上，你可以使用任何 Hprose 支持的语言来编写服务器，对于客户端调用上没有区别。

### 自定义结构体

你可以在 Hprose 客户端和服务器之间直接传输自定义结构体的对象。

你唯一需要做的事情是使用 `ClassManager.Register` 方法来注册一下你的自定义结构体。例如：

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
    hprose.ClassManager.Register(reflect.TypeOf(TestUser{}), "User")
    client := hprose.NewClient("http://www.hprose.com/example/")
    var ro *remoteObject
    client.UseService(&ro)

    fmt.Println(ro.GetUserList())
}
```

`ClassManager.Register` 的第一个参数是你的自定义结构体的类型。第二个参数是自定义结构体的别名。

客户端和服务器两端的自定义结构体的真实名字可以不同，只要它们注册相同的别名就可以了。

这个例子中的服务器是用 PHP 编写的。实际上，你同样可以在 go 服务器端使用自定义结构体。

#### 自定义结构体的字段别名

字段名在序列化时，会自动小写开头字母，所以不需要像 Json 序列化那样为了跟其他语言交互需要通过定义 `json` 标签才能实现这个功能。

但是这并不代表 hprose 不支持通过定义标签的方式来定义字段别名。事实上，它不但可以，而且它可以兼容你为 Json 序列化而定义的字段别名。例如：

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

上面的结构体是一个为 Json 序列化而定义的结构体，只要在调用 `ClassManager.Register` 时，加上第三个参数 `json`，就可以使用 `json` 标签中定义的别名来做 hprose 序列化了，如果标签中定义的别名为 "-"，则自动忽略这个字段。

当然你也可以把结构体中的 `json` 标签换成别的什么东西，比如 `hprose`，只要跟 `ClassManager.Register` 的第三个参数值一致就可以了。

### Hprose 代理

你可以通过 Hprose 服务器和客户端来为 Hprose 创建代理服务器。所有的发送到 Hprose 代理服务器上的请求都将被转发到后端的 hprose 服务器上。例如：

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

不管是否定义了错误输出参数，异常都会被自动转发。

#### 更好的代理

Hprose 提供了结果模式选项来改进代理服务器的性能。你可以像这样来使用它：

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

选项参数 `true` 表示打开简单模式开关，这样结果将以简单模式返回给客户端。

在客户端可以这样打开简单模式：

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

在开启简单模式的情况下，参数将以简单模式传递给服务器。

> 注：客户端和服务器端的简单模式也是相互独立的，你不需要同时在服务器端和客户端开启简单模式，只在一边设置为简单模式仍然可以正常通讯，不会有任何影响。

### 缺失的方法

Hprose 支持发布一个特殊的方法：MissingMethod。所有对没有显式发布的方法的调用都将被重定向到它上面。例如：

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

如果你想返回一个错误给客户端，请使用panic。因为在该方法中返回的 error 类型数据不会被特殊处理。

简单模式和结果模式也可以在它上面使用，因此通过它你可以构建更通用的 Hprose 代理服务器。

对客户端来说，调用 `AddMissingMethod` 发布的方法跟调用普通方法没有任何区别。例如：

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

结果为：

```
3
-1
2
0
0 The method 'Power' is not implemented.
```

### TCP 服务器和客户端

Hprose for Golang 还支持 TCP 的服务器和客户端。它跟 HTTP 版本的服务器和客户端在使用上一样简单。

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

### Unix 服务器和客户端

Hprose for Golang 还支持 Unix 的服务器和客户端。它跟 TCP 版本的服务器和客户端在使用上一样简单。

你可以使用 `NewUnixService` 或 `NewUnixServer`，来创建 Hprose 的 Unix 服务器。

使用 `NewUnixService`，你需要调用它的 `ServeUnix` 方法传入 Unix 连接。

使用 `NewUnixServer` 比 `NewUnixService` 则要简单的多。例如：

```go
    ...
    server := hprose.NewUnixServer("unix:/tmp/my.sock")
    server.AddFunction("hello", hello)
    server.Start()
    ...
```

创建 Hprose 的 Unix 客户端跟 TCP 客户端是一样的方式：

```go
    ...
    client := hprose.NewClient("unix:/tmp/my.sock")
    ...
```

### WebSocket 服务器和客户端

Hprose for Golang 还支持 WebSocket 的服务器和客户端。它跟 HTTP 版本的服务器和客户端在使用上一样简单。

你可以使用 `NewWebSocketService` 来创建 Hprose 的 WebSocket 服务。例如：

```go
    ...
	service := hprose.NewWebSocketService()
    service.AddFunction("hello", hello, true)
    http.ListenAndServe(":8080", service)
    ...
```

Hprose 的 WebSocket 服务器同时也是 HTTP 服务器，客户端可以用 WebSocket 访问它，也可以用 HTTP 访问它。

创建 Hprose 的 WebSocket 客户端跟 HTTP 客户端是一样的方式：

```go
    ...
    client := hprose.NewClient("ws://127.0.0.1:8080/")
    ...
```

### 服务事件

Hprose 定义了一个 `ServiceEvent` 接口。

```go
type ServiceEvent interface {}
```

这是一个空接口，但是你可以在你的实现中添加下面一些事件方法：

```go
    OnBeforeInvoke(name string, args []reflect.Value, byref bool, context hprose.Context)
	OnBeforeInvoke(name string, args []reflect.Value, byref bool, context hprose.Context) error
    OnAfterInvoke(name string, args []reflect.Value, byref bool, result []reflect.Value, context hprose.Context)
	OnAfterInvoke(name string, args []reflect.Value, byref bool, result []reflect.Value, context hprose.Context) error
    OnSendError(err error, context hprose.Context)
	OnSendError(err error, context hprose.Context) error
```

`OnBeforeInvoke`, `OnAfterInvoke` 和 `OnSendError` 都具有两种形式的定义，你只需要且仅可以实现其中的一种。

例如，如果你想针对服务器的一些行为做日志的话，你可以这样做：

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

`TcpService` 和 `TcpServer` 同样包含这个接口字段。

另外，针对 Hprose HTTP 服务器，还增加了两个 `OnSendHeader` 事件方法：

```go
	OnSendHeader(context Context)
	OnSendHeader(context *hprose.HttpContext)
```

这两种形式仍然只需要实现一种即可。它的实现同样是赋值给 `service.ServiceEvent` 字段就可以了。

## 性能测试

Hprose 比 golang 内置的 RPC 要快，你可以像这样运行性能测试程序：

```
go test --bench=".*" github.com/hprose/hprose-go/bench
```

下面是一组在 Intel i7-2600 上使用 Go 1.4 的测试结果：

benchmark | iter | time/iter
:------|------:|------:|
BenchmarkHprose| 30000|46696 ns/op
BenchmarkHprose2| 30000|48215 ns/op
BenchmarkGobRPC| 20000|66818 ns/op
BenchmarkJSONRPC| 10000|104709 ns/op
