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

func NewWebSocketClientWithEtcd(serversRepo *EtcdServersRepo) (client *WebSocketClient) {
	client = createWebSocketClient()
	client.ServersRepo = &ServersRepo{Domain: serversRepo.Domain,
		ServersMap: serversRepo.ServersMap,
		PrimaryServer: serversRepo.PrimaryServer,
	}
	client.PrimaryServerManager = serversRepo
	client.SetUri(serversRepo.PrimaryServer.ServerUrl)
	return
}

func newWebSocketClientWithEtcd(serversRepo *EtcdServersRepo) Client {
	return NewWebSocketClientWithEtcd(serversRepo)
}
