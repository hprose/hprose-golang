# Hprose for Golang Etcd Service registry protocol.

[![Join the chat at https://gitter.im/hprose/hprose-go](https://img.shields.io/badge/GITTER-join%20chat-green.svg)](https://gitter.im/hprose/hprose-go?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![GoDoc](https://godoc.org/github.com/hprose/hprose-go?status.svg&style=flat)](https://godoc.org/github.com/hprose/hprose-go)
[![Build Status](https://drone.io/github.com/hprose/hprose-go/status.png)](https://drone.io/github.com/hprose/hprose-go/latest)


## Introduction

*Hprose-Go* clustering service servers will registry itself to etcd servers (distributed key value store) and keep the heartbeat.

*Hprose-Go* clustering service clients will watch service domain and its' server list which they are interest in. Client will be always updated with the available servers.

## Service Server registry
During service server boot up, hprose server will depend on the elements: 

- Services domain name, which stand for services collection for particular business domain. Service developer need to define it. And service client/invoker will rely on this domain name.
- Server UUID, hprose etcd register module will automatically generate it.

Then etcd register wil use *key*: 
```url
"hprose-service/<services domain>/<server uuid>"
```

and *value*:
```json
{"UUID":"<server uuid>","Domain":"<services domain>","ServerUrl":"<service server listen url>","CPU":<CPU cores>}
```
to register itself to etcd cluster.


## Client Service discovery
Etcd client will retrieve all server list and watch any update belong to registered *domain key*:
```url
"hprose-service/<services domain>/<server uuid>"
```

Etcd client will use the retrieved server list to generate the primary server, and will use this chosen primary server to invoke behind service.

If hprose client invocation meet any error, Etcd client will use updated server list to re-generate service primary server as the target.
  

## etcd registry sample:
##### list the service domains
```sh
#etcdctl ls /hprose-service/

/hprose-service/http.hello.server
/hprose-service/tcp.hello.server
/hprose-service/ws.hello.server
```
> #### Service domains:
> http.hello.server
> tcp.hello.server
> ws.hello.server


##### list the server list of service domain 'tcp.hello.server'
```sh
#etcdctl ls /hprose-service/tcp.hello.server
/hprose-service/tcp.hello.server/fe091d2b-1f58-4eee-a5af-e38ea3099686
```

#### retrieve the service server's definition
```sh
#etcdctl get /hprose-service/tcp.hello.server/fe091d2b-1f58-4eee-a5af-e38ea3099686
{"UUID":"fe091d2b-1f58-4eee-a5af-e38ea3099686","Domain":"tcp.hello.server","ServerUrl":"tcp4://169.254.128.35:4321/","CPU":8}
```


## High lights of etcd register and discovery implementation

#### Service Server Register:
1. Register the server item to etcd cluster with TTL 10 seconds.
2. Keep the heartbeat to update server item with etcd cluster in 3 seconds.

#### Service client discovery:
1. During client initialization, need to retrieve all servers list for the given service domain.
2. Then used etcd watching feature to listen any servers update belong to the given service domain.
 