/*------------------------------------------------------------*\
|                                                              |
|                          hprose                              |
|                                                              |
| Official WebSite: https://hprose.com                         |
|                                                              |
| rpc/plugins/loadbalance/weighted_least_active_loadbalance.go |
|                                                              |
| LastModified: Mar 24, 2021                                   |
| Author: Ma Bingyao <andot@hprose.com>                        |
|                                                              |
\*____________________________________________________________*/

package loadbalance

import (
	"context"
	"math/rand"
	"sync"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

// WeightedLeastActiveLoadBalance plugin for hprose.
type WeightedLeastActiveLoadBalance struct {
	WeightedLoadBalance
	actives          intSlice
	effectiveWeights intSlice
	rwlock           sync.RWMutex
}

// NewWeightedLeastActiveLoadBalance returns a WeightedLeastActiveLoadBalance instance.
func NewWeightedLeastActiveLoadBalance(uris map[string]int) *WeightedLeastActiveLoadBalance {
	lb := &WeightedLeastActiveLoadBalance{
		WeightedLoadBalance: MakeWeightedLoadBalance(uris),
	}
	n := len(uris)
	lb.actives = make([]int, n)
	lb.effectiveWeights = make([]int, n)
	copy(lb.effectiveWeights, lb.Weights)
	return lb
}

// Handler for WeightedLeastActiveLoadBalance.
func (lb *WeightedLeastActiveLoadBalance) Handler(ctx context.Context, request []byte, next core.NextIOHandler) (response []byte, err error) {
	n := len(lb.URLs)
	leastActiveIndexes := make([]int, 0, n)

	lb.rwlock.RLock()
	leastActive := lb.actives.Min()
	totalWeight := 0
	for i := 0; i < n; i++ {
		if lb.actives[i] == leastActive {
			leastActiveIndexes = append(leastActiveIndexes, i)
			totalWeight += lb.effectiveWeights[i]
		}
	}
	lb.rwlock.RUnlock()

	index := leastActiveIndexes[0]
	count := len(leastActiveIndexes)
	if count > 1 {
		if totalWeight > 0 {
			currentWeight := rand.Intn(totalWeight)
			lb.rwlock.RLock()
			for i := 0; i < count; i++ {
				currentWeight -= lb.effectiveWeights[leastActiveIndexes[i]]
				if currentWeight < 0 {
					index = leastActiveIndexes[i]
					break
				}
			}
			lb.rwlock.RUnlock()
		} else {
			index = leastActiveIndexes[rand.Intn(count)]
		}
	}
	core.GetClientContext(ctx).URL = lb.URLs[index]
	lb.rwlock.Lock()
	lb.actives[index]++
	lb.rwlock.Unlock()

	defer func() {
		if e := recover(); e != nil {
			err = core.NewPanicError(e)
		}
		lb.rwlock.Lock()
		lb.actives[index]--
		if err == nil {
			if lb.effectiveWeights[index] < lb.Weights[index] {
				lb.effectiveWeights[index]++
			}
		} else {
			if lb.effectiveWeights[index] > 0 {
				lb.effectiveWeights[index]--
			}
		}
		lb.rwlock.Unlock()
	}()
	return next(ctx, request)
}
