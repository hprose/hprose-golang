/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/mock/addr.go                                         |
|                                                          |
| LastModified: Feb 21, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package mock

import (
	"net"
	"net/url"
)

type addr url.URL

func (a addr) Network() string {
	return a.Scheme
}

func (a addr) String() string {
	return a.Host
}

// NewAddr for mock.
func NewAddr(u *url.URL) net.Addr {
	return (*addr)(u)
}
