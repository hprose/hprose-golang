/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/plugins/oneway/oneway.go                             |
|                                                          |
| LastModified: Mar 24, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package oneway

import (
	"context"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

// Oneway plugin for hprose.
type Oneway struct{}

// Handler for Oneway.
func (f Oneway) Handler(ctx context.Context, name string, args []interface{}, next core.NextInvokeHandler) (result []interface{}, err error) {
	if c, ok := core.FromContext(ctx); ok && c.Items().GetBool("oneway") {
		go func() {
			_, _ = next(ctx, name, args)
		}()
		return
	}
	return next(ctx, name, args)
}
