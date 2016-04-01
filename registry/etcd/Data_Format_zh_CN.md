# Hprose for Golang Etcd 服务注册数据格式协议.

[![Join the chat at https://gitter.im/hprose/hprose-go](https://img.shields.io/badge/GITTER-join%20chat-green.svg)](https://gitter.im/hprose/hprose-go?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![GoDoc](https://godoc.org/github.com/hprose/hprose-go?status.svg&style=flat)](https://godoc.org/github.com/hprose/hprose-go)
[![Build Status](https://drone.io/github.com/hprose/hprose-go/status.png)](https://drone.io/github.com/hprose/hprose-go/latest)


## 简介

*Hprose-Go* 提供服务的集群服务器将会自动把自己注册到etcd集群服务器（分布式K/V存储引擎）并保持心跳更新。

*Hprose-Go* 集群服务的客户端将会通过etcd引擎监听所关注服务域domain下级的集群服务器列表。客户端时刻保持最新可用的服务器列表

## 服务所在服务器的自动注册
在服务所在*Hprose-Go*服务器启动过程中，Hprose服务器将会依赖以下两个元素： 

- 所在服务器上所有服务域domain名称, 是特定服务领域的约定集合名称。此名称将会由服务开发人员和服务调用开发人员协商命名，服务调用客户端将会依赖此服务域domain名称来发现提供服务的服务器列表。
- 提供Service服务的服务器UUID, hprose etcd 注册模块将会自动产生服务器的唯一标示ID。

然后etcd服务注册模块将会使用**key**: 
```url
"hprose-service/<services domain>/<server uuid>"
```

和 **value**:
```json
{"UUID":"<server uuid>","Domain":"<services domain>","ServerUrl":"<service server listen url>","CPU":<CPU cores>}
```
来把所在服务器自动注册到etcd集群中去。


## 服务调用客户端的服务自动发现
服务调用客户端在初始化过程中，etcd 客户端模块将会自动获取所关注服务域domain下所挂接的所有服务器列表，并持续监听所关注的**domain key**:
```url
"hprose-service/<services domain>/<server uuid>"
```

客户端将会根据相应的算法（目前仅仅实现随机）从服务器列表中自动选择主服务器，并使用此挑选的服务来承担后续的客户端服务调用。

如果hprose服务客户端调用过程中遇到任何与后台服务器调用相关的异常和错误，etcd客户端服务模块将会重新更新服务器列表并选择对应的后台服务器作为主服务。
  

## etcd 注册引擎中存储 hprose服务的样例数据:
##### 查看所有hprose注册的服务service domains
```sh
#etcdctl ls /hprose-service/

/hprose-service/http.hello.server
/hprose-service/tcp.hello.server
/hprose-service/ws.hello.server
```
> #### 目前在etcd中注册了三个domain服务域:
> http.hello.server
> tcp.hello.server
> ws.hello.server


##### 查看服务域domain 'tcp.hello.server'下的所有服务器节点
```sh
#etcdctl ls /hprose-service/tcp.hello.server
/hprose-service/tcp.hello.server/fe091d2b-1f58-4eee-a5af-e38ea3099686
```

#### 获取某个服务器的详细定义
```sh
#etcdctl get /hprose-service/tcp.hello.server/fe091d2b-1f58-4eee-a5af-e38ea3099686
{"UUID":"fe091d2b-1f58-4eee-a5af-e38ea3099686","Domain":"tcp.hello.server","ServerUrl":"tcp4://169.254.128.35:4321/","CPU":8}
```


## etcd 服务器注册和发现实现的几个要点

#### 提供服务的服务器注册:
1. 在注册服务器条目到etcd 引擎时，需要设置TTL为10秒，即为此生命周期10秒。意味着10秒后，此服务条目将会失效并自动删除。
2. 必须持续与etcd服务器进行心跳交互，并重新更新服务器条目到etcd引擎中，心跳频率目前为3秒。

#### 服务请求客户端的服务发现:
1. 在客户单初始化过程中，需要主动从etcd引擎中获取所关注服务域(Service Domain)所对应的所有服务器列表。
2. 客户单并持续监听挂接在此服务域下的所有服务器的更新状态，并及时更新提供服务的主服务器
 