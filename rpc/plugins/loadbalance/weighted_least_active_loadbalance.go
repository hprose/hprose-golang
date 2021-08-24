/*------------------------------------------------------------*\
|                                                              |
|                          hprose                              |
|                                                              |
| Official WebSite: https://hprose.com                         |
|                                                              |
| rpc/plugins/loadbalance/weighted_least_active_loadbalance.go |
|                                                              |
| LastModified: Aug 24, 2021                                   |
| Author: Ma Bingyao <andot@hprose.com>                        |
|                                                              |
\*____________________________________________________________*/

package loadbalance

import (
	"context"
	"fmt"
	"math/rand"
	"sync"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

// WeightedLeastActiveLoadBalance plugin for hprose.
type WeightedLeastActiveLoadBalance struct {
	WeightedLoadBalance
	actives          int64Slice
	effectiveWeights int64Slice
	rwlock           sync.RWMutex
}

// NewWeightedLeastActiveLoadBalance returns a WeightedLeastActiveLoadBalance instance.
func NewWeightedLeastActiveLoadBalance(uris map[string]int) *WeightedLeastActiveLoadBalance {
	lb := &WeightedLeastActiveLoadBalance{
		WeightedLoadBalance: MakeWeightedLoadBalance(uris),
	}
	n := len(uris)
	lb.actives = make([]int64, n)
	lb.effectiveWeights = make([]int64, n)
	copy(lb.effectiveWeights, lb.Weights)
	return lb
}

func (lb *WeightedLeastActiveLoadBalance) getIndex() int {
	n := len(lb.URLs)
	leastActiveIndexes := make([]int, 0, n)

	lb.rwlock.RLock()
	leastActive := lb.actives.Min()
	var totalWeight int64
	for i := 0; i < n; i++ {
		if lb.actives[i] == leastActive {
			leastActiveIndexes = append(leastActiveIndexes, i)
			totalWeight += lb.effectiveWeights[i]
		}
	}
	lb.rwlock.RUnlock()

	index := leastActiveIndexes[0]
	count := len(leastActiveIndexes)
	if count <= 1 {
		return index
	}
	if totalWeight <= 0 {
		return leastActiveIndexes[rand.Intn(count)]
	}
	currentWeight := rand.Int63n(totalWeight)
	lb.rwlock.RLock()
	for i := 0; i < count; i++ {
		currentWeight -= lb.effectiveWeights[leastActiveIndexes[i]]
		if currentWeight < 0 {
			index = leastActiveIndexes[i]
			break
		}
	}
	lb.rwlock.RUnlock()
	return index
}

// Handler for WeightedLeastActiveLoadBalance.
func (lb *WeightedLeastActiveLoadBalance) Handler(ctx context.Context, request []byte, next core.NextIOHandler) (response []byte, err error) {
	index := lb.getIndex()
	core.GetClientContext(ctx).URL = lb.URLs[index]
	fmt.Println(lb.URLs[index])
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
		} else if lb.effectiveWeights[index] > 0 {
			lb.effectiveWeights[index]--
		}
		lb.rwlock.Unlock()
	}()
	return next(ctx, request)
}
