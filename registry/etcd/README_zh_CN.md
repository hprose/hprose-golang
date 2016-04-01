# Hprose for Golang 集群高可用支持.

[![Join the chat at https://gitter.im/hprose/hprose-go](https://img.shields.io/badge/GITTER-join%20chat-green.svg)](https://gitter.im/hprose/hprose-go?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![GoDoc](https://godoc.org/github.com/hprose/hprose-go?status.svg&style=flat)](https://godoc.org/github.com/hprose/hprose-go)
[![Build Status](https://drone.io/github.com/hprose/hprose-go/status.png)](https://drone.io/github.com/hprose/hprose-go/latest)


## 简介
 
*Hprose-Go* 集群功能支持分布式服务注册和分布式服务自动发现。 Hprose-Go 基于etcd(分布式服务注册引擎)能提供高可用性和自动负载均衡的能力。

## Hprose-Go安装

```sh
go get github.com/hprose/hprose-go
go install github.com/hprose/hprose-go
```

## 安装和启动Etcd
https://coreos.com/etcd/

可以从https://github.com/coreos/etcd/releases 选择下载对应操作系统最新的二进制版本

直接运行etcd命令启动独立的测试运行实例。
```sh
./etcd
```

##通过样例测试SOA特性
启动样例服务器: hello tcp server
```sh
cd registry/etcd/examples
go run etcd_tcphelloserver.go
```

启动样例客户端: hello tcp client
```sh
cd registry/etcd/examples
go run etcd_tcphelloclient.go
```

## etcd 集群功能样例核心代码解读
- hello tcp 服务器样例代码:

```go
func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

    //服务器服务域domain名称, domain代表是服务器上提供的各种服务集合。
	domain := "tcp.hello.server"
    //提供Service服务器的服务地址端口。
	tcpEndpoint := "tcp4://"+hprose.GetLocalIP()+":4321/"
    //提供服务注册的etcd集群服务器服务地址列表。
	etcdEndpoints :=[]string{"http://127.0.0.1:2379"}

    /** 
     * 以下方法用来把服务域domain名称为"tcp.hello.server"服务集合所在的
     * 本机服务器注册到etcd注册引擎中去，并保持与注册引擎的心跳更新。
     * 如果此服务器被停止，此服务器将会被自动从注册引擎中删除。
     * 与此同时，监听此服务域domain的客户端将会自动剔除此失效的服务器
     * 服务客户端会维持一份最新有效的服务器列表。    
     */
	etcd.RegisterServer(domain,tcpEndpoint,etcdEndpoints)

	server := hprose.NewTcpServer(tcpEndpoint)
	server.AddFunction("hello", hello)
	server.SetKeepAlive(true)
	server.Start()
}
```

- hello tcp 客户端样例代码:

```go
type Stub struct {
	Hello func(string) (string, error)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

    //Domain名称包含了我们所希望访问的服务。
	domain := "tcp.hello.server"

    /**
     * 服务注册引擎集群的服务器地址列表。
     * 客户端无需知道服务具体由哪个服务器提供的，
     * etcd注册引擎集群将会自动提供给客户端最新对应的服务器列表。
     */
	etcdEndpoints :=[]string{"http://127.0.0.1:2379"}

    /**
     * etcd.NewClient 方法将会自动产生hprose代理客户端，
     * 此客户端中将会在服务调用过程中自动注入etcd 服务自动发现逻辑。
     * 此服务发现逻辑将会在客户端服务调用过程中自动更新对应的服务器列表。
     * 此服务发现逻辑可以提供服务请求负载均衡和服务调用的高可用集群特性。
     */
	client := etcd.NewClient(domain,etcdEndpoints) 

	var stub *Stub
	client.UseService(&stub)
	client.SetKeepAlive(true)
    ......
}

```
