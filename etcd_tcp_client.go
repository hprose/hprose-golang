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

func NewTcpClientWithEtcd(serversRepo *EtcdServersRepo) (client *TcpClient) {
	client = createTcpClient()
	client.ServersRepo = &ServersRepo{Domain: serversRepo.Domain,
		ServersMap: serversRepo.ServersMap,
		PrimaryServer: serversRepo.PrimaryServer,
	}
	client.PrimaryServerManager = serversRepo
	client.SetUri(serversRepo.PrimaryServer.ServerUrl)
	return
}

func newTcpClientWithEtcd(serversRepo *EtcdServersRepo) Client {
	return NewTcpClientWithEtcd(serversRepo)
}
