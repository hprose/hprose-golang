/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/io_manager.go                                   |
|                                                          |
| LastModified: Feb 18, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import "context"

// NextIOHandler for RPC.
type NextIOHandler func(ctx context.Context, request []byte) (response []byte, err error)

// IOHandler for RPC.
type IOHandler func(ctx context.Context, request []byte, next NextIOHandler) (response []byte, err error)

// NewIOManager returns an IO PluginManager.
func NewIOManager(handler NextIOHandler) PluginManager {
	return newPluginManager(handler, func(handler PluginHandler, next NextPluginHandler) NextPluginHandler {
		h := handler.(IOHandler)
		n := next.(NextIOHandler)
		return NextIOHandler(func(ctx context.Context, request []byte) (response []byte, err error) {
			return h(ctx, request, n)
		})
	})
}
