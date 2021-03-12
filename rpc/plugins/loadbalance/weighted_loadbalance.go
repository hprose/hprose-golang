/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/loadbalance/weighted_loadbalance.go          |
|                                                          |
| LastModified: Mar 12, 2021                               |
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
	Weights []int
}

// MakeWeightedLoadBalance returns a WeightedLoadBalance instance.
func MakeWeightedLoadBalance(uris map[string]int) (lb WeightedLoadBalance) {
	n := len(uris)
	lb.URLs = make([]*url.URL, n)
	lb.Weights = make([]int, n)
	i := 0
	for key, value := range uris {
		lb.URLs[i], _ = url.Parse(key)
		lb.Weights[i] = value
		if value <= 0 {
			panic("loadbalance: urls weight must be great than 0")
		}
		i++
	}
	return
}
