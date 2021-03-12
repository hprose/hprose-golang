/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/loadbalance/weighted_random_loadbalance.go   |
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

// WeightedRandomLoadBalance plugin for hprose.
type WeightedRandomLoadBalance struct {
	WeightedLoadBalance
	random           *rand.Rand
	effectiveWeights []int
	rwlock           sync.RWMutex
}

// NewWeightedRandomLoadBalance returns a WeightedRandomLoadBalance instance.
func NewWeightedRandomLoadBalance(uris map[string]int) *WeightedRandomLoadBalance {
	lb := &WeightedRandomLoadBalance{
		WeightedLoadBalance: MakeWeightedLoadBalance(uris),
		random:              rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
	}
	lb.effectiveWeights = make([]int, len(uris))
	copy(lb.effectiveWeights, lb.Weights)
	return lb
}

// Handler for WeightedRandomLoadBalance.
func (lb *WeightedRandomLoadBalance) Handler(ctx context.Context, request []byte, next core.NextIOHandler) (response []byte, err error) {
	n := len(lb.URLs)
	index := n - 1
	lb.rwlock.RLock()
	totalWeight := sum(lb.effectiveWeights)
	if totalWeight > 0 {
		currentWeight := lb.random.Intn(totalWeight)
		for i := 0; i < n; i++ {
			currentWeight -= lb.effectiveWeights[i]
			if currentWeight < 0 {
				index = i
				break
			}
		}
	} else {
		index = lb.random.Intn(n)
	}
	lb.rwlock.RUnlock()
	core.GetClientContext(ctx).URL = lb.URLs[index]
	defer func() {
		if e := recover(); e != nil {
			err = core.NewPanicError(e)
		}
		if err == nil {
			return
		}
		lb.rwlock.RLock()
		condition := lb.effectiveWeights[index] > 0
		lb.rwlock.RUnlock()
		if condition {
			lb.rwlock.Lock()
			if lb.effectiveWeights[index] > 0 {
				lb.effectiveWeights[index]--
			}
			lb.rwlock.Unlock()
		}
	}()
	if response, err = next(ctx, request); err != nil {
		return
	}
	lb.rwlock.RLock()
	condition := lb.effectiveWeights[index] < lb.Weights[index]
	lb.rwlock.RUnlock()
	if condition {
		lb.rwlock.Lock()
		if lb.effectiveWeights[index] < lb.Weights[index] {
			lb.effectiveWeights[index]++
		}
		lb.rwlock.Unlock()
	}
	return
}

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
