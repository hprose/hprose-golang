/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/loadbalance/least_active_loadbalance.go      |
|                                                          |
| LastModified: Mar 12, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package loadbalance

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

// LeastActiveLoadBalance plugin for hprose.
type LeastActiveLoadBalance struct {
	random  *rand.Rand
	actives []int
	rwlock  sync.RWMutex
}

// NewLeastActiveLoadBalance returns a LeastActiveLoadBalance instance.
func NewLeastActiveLoadBalance() *LeastActiveLoadBalance {
	return &LeastActiveLoadBalance{
		random: rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
	}
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
			lb.actives = make([]int, n)
		}
		lb.rwlock.Unlock()
	}

	lb.rwlock.RLock()
	var leastActive int
	if len(lb.actives) > n {
		leastActive = min(lb.actives[:n])
	} else {
		leastActive = min(lb.actives)
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
		index = leastActiveIndexes[lb.random.Intn(count)]
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

func min(nums []int) int {
	minVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if minVal > nums[i] {
			minVal = nums[i]
		}
	}
	return minVal
}
