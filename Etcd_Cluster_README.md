# Hprose for Golang Clustering support.

[![Join the chat at https://gitter.im/hprose/hprose-go](https://img.shields.io/badge/GITTER-join%20chat-green.svg)](https://gitter.im/hprose/hprose-go?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![GoDoc](https://godoc.org/github.com/hprose/hprose-go?status.svg&style=flat)](https://godoc.org/github.com/hprose/hprose-go)
[![Build Status](https://drone.io/github.com/hprose/hprose-go/status.png)](https://drone.io/github.com/hprose/hprose-go/latest)


## Introduction
 
*Hprose-Go* cluster support distributed service registration and discovery. Hprose-Go can provide high available and load balance feature based on the etcd (distributed key value store).

## Installation: Hprose-Go

```sh
go get github.com/hprose/hprose-go
go install github.com/hprose/hprose-go
```

## Install and Start Etcd
https://coreos.com/etcd/

Download the last etcd binary code based on your os from https://github.com/coreos/etcd/releases

Run the etcd command to start etcd standalone testing instance.
```sh
./etcd
```

## Test the SOA feature through sample code
Start the sample: hello tcp server
```sh
cd examples
go run etcd_tcphelloserver.go
```

Start the testing client: hello tcp client
```sh
cd examples
go run etcd_tcphelloclient.go
```

## etcd clustering sample code explanation
- hello tcp server sample code:

```go
func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

    //Services domain name, this present as the service collector for given services.
	domain := "tcp.hello.server"
    //The server endpoint which provide the services...
	tcpEndpoint := "tcp4://"+hprose.GetLocalIP()+":4321/"
    //Service registry cluster etcd servers' endpoints
	etcdEndpoints :=[]string{"http://127.0.0.1:2379"}

    /** 
     * Used to register the service domain "tcp.hello.server" to the etcd registry
     * And keep services heartbeat with registry.
     * If service domain server being shutdown, this server will be deleted 
     * from registry and clients listened to this domain will be updated.
     * Services client will keep last updated server list.    
     */
	hprose.EtcdRegisterServer(domain,tcpEndpoint,etcdEndpoints)

	server := hprose.NewTcpServer(tcpEndpoint)
	server.AddFunction("hello", hello)
	server.SetKeepAlive(true)
	server.Start()
}
```

- hello tcp client sample code:

```go
type Stub struct {
	Hello func(string) (string, error)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

    //Domain name which include the services we want to invoke
	domain := "tcp.hello.server"

    /**
     *Service registry cluster etcd servers' endpoints
     *We do not need to know where service server is.
     *etcd cluster will provide us last updated server list.
     */
	etcdEndpoints :=[]string{"http://127.0.0.1:2379"}

    /**
     * NewClientWithEtcd method will generate hprose client 
     * which will inject the etcd discovery logic during service invocation.
     * And discovery logic will keep the last updated service list for the service invoked.
     * Logic will provide services load balance and high avaialbe clustering features.
     */
	client := hprose.NewClientWithEtcd(domain,etcdEndpoints) //Used for Clustering model...

	var stub *Stub
	client.UseService(&stub)
	client.SetKeepAlive(true)
    ......
}

```
