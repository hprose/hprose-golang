/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/loadbalance/nginx_round_robin_loadbalance.go |
|                                                          |
| LastModified: Mar 24, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package loadbalance

import (
	"context"
	"math"
	"math/rand"
	"sync"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

// NginxRoundRobinLoadBalance plugin for hprose.
type NginxRoundRobinLoadBalance struct {
	WeightedLoadBalance
	effectiveWeights intSlice
	currentWeights   intSlice
	lock             sync.Mutex
}

// NewNginxRoundRobinLoadBalance returns a NginxRoundRobinLoadBalance instance.
func NewNginxRoundRobinLoadBalance(uris map[string]int) *NginxRoundRobinLoadBalance {
	lb := &NginxRoundRobinLoadBalance{
		WeightedLoadBalance: MakeWeightedLoadBalance(uris),
	}
	n := len(uris)
	lb.currentWeights = make([]int, n)
	lb.effectiveWeights = make([]int, n)
	copy(lb.effectiveWeights, lb.Weights)
	return lb
}

func (lb *NginxRoundRobinLoadBalance) getIndex() int {
	n := len(lb.URLs)
	lb.lock.Lock()
	defer lb.lock.Unlock()
	totalWeight := lb.effectiveWeights.Sum()
	if totalWeight > 0 {
		var index int
		currentWeight := math.MinInt64
		for i := 0; i < n; i++ {
			lb.currentWeights[i] += lb.effectiveWeights[i]
			weight := lb.currentWeights[i]
			if currentWeight < weight {
				currentWeight = weight
				index = i
			}
		}
		lb.currentWeights[index] = currentWeight - totalWeight
		return index
	}
	return rand.Intn(n)
}

// Handler for NginxRoundRobinLoadBalance.
func (lb *NginxRoundRobinLoadBalance) Handler(ctx context.Context, request []byte, next core.NextIOHandler) (response []byte, err error) {
	index := lb.getIndex()
	core.GetClientContext(ctx).URL = lb.URLs[index]
	defer func() {
		if e := recover(); e != nil {
			err = core.NewPanicError(e)
		}
		lb.lock.Lock()
		if err == nil {
			if lb.effectiveWeights[index] < lb.Weights[index] {
				lb.effectiveWeights[index]++
			}
		} else {
			if lb.effectiveWeights[index] > 0 {
				lb.effectiveWeights[index]--
			}
		}
		lb.lock.Unlock()
	}()
	return next(ctx, request)
}
