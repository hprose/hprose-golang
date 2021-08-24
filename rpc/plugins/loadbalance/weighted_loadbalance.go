/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/loadbalance/weighted_loadbalance.go          |
|                                                          |
| LastModified: Aug 24, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package loadbalance

import (
	"net/url"
)

// WeightedLoadBalance plugin for hprose.
type WeightedLoadBalance struct {
	URLs    []*url.URL
	Weights int64Slice
}

// MakeWeightedLoadBalance returns a WeightedLoadBalance instance.
func MakeWeightedLoadBalance(uris map[string]int) (lb WeightedLoadBalance) {
	n := len(uris)
	lb.URLs = make([]*url.URL, n)
	lb.Weights = make([]int64, n)
	i := 0
	for key, value := range uris {
		lb.URLs[i], _ = url.Parse(key)
		lb.Weights[i] = int64(value)
		if value <= 0 {
			panic("loadbalance: urls weight must be great than 0")
		}
		i++
	}
	return
}
