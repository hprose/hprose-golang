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
 * rpc/client_event.go                                    *
 *                                                        *
 * hprose client event for Go.                            *
 *                                                        *
 * LastModified: Sep 23, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

// ClientEvent is the client event
type ClientEvent interface{}

type onErrorEvent interface {
	OnError(name string, err error)
}

type onFailswitchEvent interface {
	OnFailswitch(Client)
}
