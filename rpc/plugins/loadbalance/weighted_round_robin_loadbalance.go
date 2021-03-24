/*-----------------------------------------------------------*\
|                                                             |
|                          hprose                             |
|                                                             |
| Official WebSite: https://hprose.com                        |
|                                                             |
| rpc/plugins/loadbalance/weighted_round_robin_loadbalance.go |
|                                                             |
| LastModified: Mar 24, 2021                                  |
| Author: Ma Bingyao <andot@hprose.com>                       |
|                                                             |
\*___________________________________________________________*/

package loadbalance

import (
	"context"
	"sync"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

// WeightedRoundRobinLoadBalance plugin for hprose.
type WeightedRoundRobinLoadBalance struct {
	WeightedLoadBalance
	lock          sync.Mutex
	maxWeight     int
	gcdWeight     int
	index         int
	currentWeight int
}

// NewWeightedRoundRobinLoadBalance returns a WeightedRoundRobinLoadBalance instance.
func NewWeightedRoundRobinLoadBalance(uris map[string]int) *WeightedRoundRobinLoadBalance {
	lb := &WeightedRoundRobinLoadBalance{
		WeightedLoadBalance: MakeWeightedLoadBalance(uris),
		index:               -1,
		currentWeight:       0,
	}
	lb.maxWeight = lb.Weights.Max()
	lb.gcdWeight = lb.Weights.GCD()
	return lb
}

func (lb *WeightedRoundRobinLoadBalance) prepare(ctx context.Context) {
	n := len(lb.URLs)
	lb.lock.Lock()
	defer lb.lock.Unlock()
	for {
		lb.index = (lb.index + 1) % n
		if lb.index == 0 {
			lb.currentWeight -= lb.gcdWeight
			if lb.currentWeight <= 0 {
				lb.currentWeight = lb.maxWeight
			}
			if lb.Weights[lb.index] >= lb.currentWeight {
				core.GetClientContext(ctx).URL = lb.URLs[lb.index]
				break
			}
		}
	}
}

// Handler for WeightedRoundRobinLoadBalance.
func (lb *WeightedRoundRobinLoadBalance) Handler(ctx context.Context, request []byte, next core.NextIOHandler) (response []byte, err error) {
	lb.prepare(ctx)
	return next(ctx, request)
}
