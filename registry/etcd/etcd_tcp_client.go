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


func NewTcpClientWithEtcd(serversRepo *EtcdServersRepo) (client *hprose.TcpClient) {
	client = hprose.CreateTcpClient()
	client.ServersRepo = &hprose.ServersRepo{Domain: serversRepo.Domain,
		ServersMap: serversRepo.ServersMap,
		PrimaryServer: serversRepo.PrimaryServer,
	}
	client.PrimaryServerManager = serversRepo
	client.SetUri(serversRepo.PrimaryServer.ServerUrl)
	return
}

func newTcpClientWithEtcd(serversRepo *EtcdServersRepo) hprose.Client {
	return NewTcpClientWithEtcd(serversRepo)
}
