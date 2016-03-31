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
 * Author: Henry Hu <henry.pf.hu@gmail.com                *
 *                                                        *
\**********************************************************/

package hprose

import (
	"strings"
	"net/url"
)

var clientWithEtcdFactories = make(map[string]func(*EtcdServersRepo) Client)

func init() {
	RegisterClientWithEtcdFactory("tcp", newTcpClientWithEtcd)
	RegisterClientWithEtcdFactory("tcp4", newTcpClientWithEtcd)
	RegisterClientWithEtcdFactory("tcp6", newTcpClientWithEtcd)

	RegisterClientWithEtcdFactory("http", newHTTPClientWithEtcd)
	RegisterClientWithEtcdFactory("https", newHTTPClientWithEtcd)

	RegisterClientWithEtcdFactory("ws", newWebSocketClientWithEtcd)
	RegisterClientWithEtcdFactory("wss", newWebSocketClientWithEtcd)
}


// RegisterClientFactory register client factory
func RegisterClientWithEtcdFactory(scheme string, newClient func(*EtcdServersRepo) Client) {
	clientWithEtcdFactories[strings.ToLower(scheme)] = newClient
}


func NewClientWithEtcd(domain string, etcEndpoints []string) Client {
	serversRepo := NewEtcdServersRepo(domain, etcEndpoints)
	serversRepo.Update()

	uri := serversRepo.PrimaryServer.ServerUrl
	if u, err := url.Parse(uri); err == nil {
		if newClient, ok := clientWithEtcdFactories[u.Scheme]; ok {
			return newClient(serversRepo)
		}
		panic("The " + u.Scheme + "client isn't implemented.")
	} else {
		panic("The uri can't be parsed.")
	}
}