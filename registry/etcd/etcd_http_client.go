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

func NewHTTPClientWithEtcd(serversRepo *EtcdServersRepo) (client *hprose.HttpClient) {
	client = hprose.CreateHttpClient()
	client.ServersRepo = &hprose.ServersRepo{Domain: serversRepo.Domain,
		ServersMap: serversRepo.ServersMap,
		PrimaryServer: serversRepo.PrimaryServer,
	}
	client.PrimaryServerManager = serversRepo
	client.SetUri(serversRepo.PrimaryServer.ServerUrl)
	return
}

func newHTTPClientWithEtcd(serversRepo *EtcdServersRepo) hprose.Client {
	return NewHTTPClientWithEtcd(serversRepo)
}
