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

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

// RandomLoadBalance plugin for hprose.
type RandomLoadBalance struct {
}

// NewRandomLoadBalance returns a RandomLoadBalance instance.
func NewRandomLoadBalance() *RandomLoadBalance {
	return &RandomLoadBalance{}
}

// Handler for RandomLoadBalance.
func (lb *RandomLoadBalance) Handler(ctx context.Context, request []byte, next core.NextIOHandler) (response []byte, err error) {
	clientContext := core.GetClientContext(ctx)
	urls := clientContext.Client().URLs
	clientContext.URL = urls[rand.Intn(len(urls))]
	return next(ctx, request)
}
