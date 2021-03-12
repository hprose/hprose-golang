/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/loadbalance/random_loadbalance.go            |
|                                                          |
| LastModified: Mar 12, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package loadbalance

import (
	"context"
	"math/rand"
	"time"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

// RandomLoadBalance plugin for hprose.
type RandomLoadBalance struct {
	random *rand.Rand
}

// NewRandomLoadBalance returns a RandomLoadBalance instance.
func NewRandomLoadBalance() *RandomLoadBalance {
	return &RandomLoadBalance{
		random: rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
	}
}

// Handler for RandomLoadBalance.
func (lb *RandomLoadBalance) Handler(ctx context.Context, request []byte, next core.NextIOHandler) (response []byte, err error) {
	clientContext := core.GetClientContext(ctx)
	urls := clientContext.Client().URLs
	clientContext.URL = urls[lb.random.Intn(len(urls))]
	return next(ctx, request)
}
