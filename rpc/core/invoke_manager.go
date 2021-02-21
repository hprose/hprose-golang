/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/invoke_manager.go                               |
|                                                          |
| LastModified: Feb 18, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import "context"

// NextInvokeHandler for RPC.
type NextInvokeHandler func(ctx context.Context, name string, args []interface{}) (result []interface{}, err error)

// InvokeHandler for RPC.
type InvokeHandler func(ctx context.Context, name string, args []interface{}, next NextInvokeHandler) (result []interface{}, err error)

// NewInvokeManager returns an Invoke PluginManager.
func NewInvokeManager(handler NextInvokeHandler) PluginManager {
	return newPluginManager(handler, func(handler PluginHandler, next NextPluginHandler) NextPluginHandler {
		h := handler.(InvokeHandler)
		n := next.(NextInvokeHandler)
		return NextInvokeHandler(func(ctx context.Context, name string, args []interface{}) (result []interface{}, err error) {
			return h(ctx, name, args, n)
		})
	})
}
