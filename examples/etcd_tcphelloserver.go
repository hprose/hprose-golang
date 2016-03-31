/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 * Modified By: Henry Hu <henry.pf.hu@gmail.com>          *
 *                                                        *
\**********************************************************/

package main

import (
	"github.com/hprose/hprose-go"
	"runtime"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	domain := "tcp.hello.server"
	tcpEndpoint := "tcp4://"+hprose.GetLocalIP()+":4321/"
	etcdEndpoints :=[]string{"http://127.0.0.1:2379"}

	hprose.EtcdRegisterServer(domain,tcpEndpoint,etcdEndpoints)

	server := hprose.NewTcpServer(tcpEndpoint)
	server.AddFunction("hello", hello)
	server.SetKeepAlive(true)
	server.Start()
}
