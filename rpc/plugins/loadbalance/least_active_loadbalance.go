/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/loadbalance/least_active_loadbalance.go      |
|                                                          |
| LastModified: Aug 24, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package loadbalance

import (
	"context"
	"math/rand"
	"sync"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

// LeastActiveLoadBalance plugin for hprose.
type LeastActiveLoadBalance struct {
	actives int64Slice
	rwlock  sync.RWMutex
}

// NewLeastActiveLoadBalance returns a LeastActiveLoadBalance instance.
func NewLeastActiveLoadBalance() *LeastActiveLoadBalance {
	return &LeastActiveLoadBalance{}
}

// Handler for LeastActiveLoadBalance.
func (lb *LeastActiveLoadBalance) Handler(ctx context.Context, request []byte, next core.NextIOHandler) (response []byte, err error) {
	clientContext := core.GetClientContext(ctx)
	urls := clientContext.Client().URLs
	n := len(urls)
	leastActiveIndexes := make([]int, 0, n)

	lb.rwlock.RLock()
	condition := len(lb.actives) < n
	lb.rwlock.RUnlock()
	if condition {
		lb.rwlock.Lock()
		if len(lb.actives) < n {
			lb.actives = make([]int64, n)
		}
		lb.rwlock.Unlock()
	}

	lb.rwlock.RLock()
	var leastActive int64
	if len(lb.actives) > n {
		leastActive = lb.actives[:n].Min()
	} else {
		leastActive = lb.actives.Min()
	}
	for i := 0; i < n; i++ {
		if lb.actives[i] == leastActive {
			leastActiveIndexes = append(leastActiveIndexes, i)
		}
	}
	lb.rwlock.RUnlock()

	index := leastActiveIndexes[0]
	count := len(leastActiveIndexes)
	if count > 1 {
		index = leastActiveIndexes[rand.Intn(count)]
	}
	clientContext.URL = urls[index]

	lb.rwlock.Lock()
	lb.actives[index]++
	lb.rwlock.Unlock()

	defer func() {
		lb.rwlock.Lock()
		lb.actives[index]--
		lb.rwlock.Unlock()
	}()

	return next(ctx, request)
}
