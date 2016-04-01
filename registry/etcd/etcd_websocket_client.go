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

package etcd

import (
	"github.com/hprose/hprose-go"
)


func NewWebSocketClientWithEtcd(serversRepo *EtcdServersRepo) (client *hprose.WebSocketClient) {
	client = hprose.CreateWebSocketClient()
	client.ServersRepo = &hprose.ServersRepo{Domain: serversRepo.Domain,
		ServersMap: serversRepo.ServersMap,
		PrimaryServer: serversRepo.PrimaryServer,
	}
	client.PrimaryServerManager = serversRepo
	client.SetUri(serversRepo.PrimaryServer.ServerUrl)
	return
}

func newWebSocketClientWithEtcd(serversRepo *EtcdServersRepo) hprose.Client {
	return NewWebSocketClientWithEtcd(serversRepo)
}
