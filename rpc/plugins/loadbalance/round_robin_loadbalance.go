/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/loadbalance/round_robin_loadbalance.go       |
|                                                          |
| LastModified: Mar 12, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package loadbalance

import (
	"context"
	"sync/atomic"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

// RoundRobinLoadBalance plugin for hprose.
type RoundRobinLoadBalance struct {
	index int64
}

// NewRoundRobinLoadBalance returns a RoundRobinLoadBalance instance.
func NewRoundRobinLoadBalance() *RoundRobinLoadBalance {
	return &RoundRobinLoadBalance{index: -1}
}

// Handler for RoundRobinLoadBalance.
func (lb *RoundRobinLoadBalance) Handler(ctx context.Context, request []byte, next core.NextIOHandler) (response []byte, err error) {
	clientContext := core.GetClientContext(ctx)
	urls := clientContext.Client().URLs
	n := int64(len(urls))
	if n > 1 {
		if atomic.AddInt64(&lb.index, 1) >= n {
			atomic.StoreInt64(&lb.index, 0)
		}
	}
	clientContext.URL = urls[atomic.LoadInt64(&lb.index)]
	return next(ctx, request)
}
