# Hprose for Golang Clustering support.

[![Join the chat at https://gitter.im/hprose/hprose-go](https://img.shields.io/badge/GITTER-join%20chat-green.svg)](https://gitter.im/hprose/hprose-go?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![GoDoc](https://godoc.org/github.com/hprose/hprose-go?status.svg&style=flat)](https://godoc.org/github.com/hprose/hprose-go)
[![Build Status](https://drone.io/github.com/hprose/hprose-go/status.png)](https://drone.io/github.com/hprose/hprose-go/latest)



## Introduction

*Hprose-Go* cluster support distributed service registration and discovery. Hprose can provide high available and load balance feature based on the etcd (distributed key value store) .


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
go run etcd_tcphelloserver.go
```

Start the testing client: hello tcp client
```sh
go run etcd_tcphelloclient.go
```
