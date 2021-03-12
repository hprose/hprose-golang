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

func (lb *RoundRobinLoadBalance) getIndex(n int64) int64 {
	if n > 1 {
		if i := atomic.AddInt64(&lb.index, 1); i < n {
			return i
		}
		atomic.StoreInt64(&lb.index, 0)
	}
	return 0
}

// Handler for RoundRobinLoadBalance.
func (lb *RoundRobinLoadBalance) Handler(ctx context.Context, request []byte, next core.NextIOHandler) (response []byte, err error) {
	clientContext := core.GetClientContext(ctx)
	urls := clientContext.Client().URLs
	clientContext.URL = urls[lb.getIndex(int64(len(urls)))]
	return next(ctx, request)
}
