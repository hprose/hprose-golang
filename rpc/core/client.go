/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/client.go                                       |
|                                                          |
| LastModified: Feb 8, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"net/url"
	"time"
)

// Client for RPC.
type Client interface {
	URLs() []url.URL
	Timeout() time.Duration
	RequestHeaders() Dict
}
