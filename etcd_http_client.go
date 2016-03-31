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

func NewHTTPClientWithEtcd(serversRepo *EtcdServersRepo) (client *HttpClient) {
	client = createHttpClient()
	client.ServersRepo = &ServersRepo{Domain: serversRepo.Domain,
		ServersMap: serversRepo.ServersMap,
		PrimaryServer: serversRepo.PrimaryServer,
	}
	client.PrimaryServerManager = serversRepo
	client.SetUri(serversRepo.PrimaryServer.ServerUrl)
	return
}

func newHTTPClientWithEtcd(serversRepo *EtcdServersRepo) Client {
	return NewHTTPClientWithEtcd(serversRepo)
}
