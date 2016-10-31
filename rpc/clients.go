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
 * rpc/clients.go                                         *
 *                                                        *
 * hprose clients for Go.                                 *
 *                                                        *
 * LastModified: Sep 21, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

// Clients interface for server push
type Clients interface {
	IDList(topic string) []string
	Exist(topic string, id string) bool
	Push(topic string, result interface{}, id ...string)
	Broadcast(topic string, result interface{}, callback func([]string))
	Multicast(topic string, ids []string, result interface{}, callback func([]string))
	Unicast(topic string, id string, result interface{}, callback func(bool))
}
