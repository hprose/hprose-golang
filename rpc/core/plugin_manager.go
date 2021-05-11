/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/plugin_manager.go                               |
|                                                          |
| LastModified: May 11, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"context"
	"reflect"
	"sync"
)

// NextPluginHandler must be one of NextInvokeHandler or NextIOHandler.
type NextPluginHandler interface{}

// PluginHandler must be one of InvokeHandler or IOHandler.
type PluginHandler interface{}

type plugin interface {
	IOHandler(ctx context.Context, request []byte, next NextIOHandler) (response []byte, err error)
	InvokeHandler(ctx context.Context, name string, args []interface{}, next NextInvokeHandler) (result []interface{}, err error)
}

type invokePlugin interface {
	Handler(ctx context.Context, name string, args []interface{}, next NextInvokeHandler) (result []interface{}, err error)
}

type ioPlugin interface {
	Handler(ctx context.Context, request []byte, next NextIOHandler) (response []byte, err error)
}

func separatePluginHandlers(handlers []PluginHandler) (invokeHandlers []PluginHandler, ioHandlers []PluginHandler) {
	for _, handler := range handlers {
		switch handler := handler.(type) {
		case InvokeHandler:
			invokeHandlers = append(invokeHandlers, handler)
		case func(ctx context.Context, name string, args []interface{}, next NextInvokeHandler) (result []interface{}, err error):
			invokeHandlers = append(invokeHandlers, InvokeHandler(handler))
		case IOHandler:
			ioHandlers = append(ioHandlers, handler)
		case func(ctx context.Context, request []byte, next NextIOHandler) (response []byte, err error):
			ioHandlers = append(ioHandlers, IOHandler(handler))
		case plugin:
			invokeHandlers = append(invokeHandlers, InvokeHandler(handler.InvokeHandler))
			ioHandlers = append(ioHandlers, IOHandler(handler.IOHandler))
		case invokePlugin:
			invokeHandlers = append(invokeHandlers, InvokeHandler(handler.Handler))
		case ioPlugin:
			ioHandlers = append(ioHandlers, IOHandler(handler.Handler))
		default:
			panic("invalid plugin handler")
		}
	}
	return
}

// PluginManager for RPC.
type PluginManager interface {
	Handler() NextPluginHandler
	Use(handler ...PluginHandler)
	Unuse(handler ...PluginHandler)
}

type pluginManager struct {
	sync.RWMutex
	handlers       []PluginHandler
	defaultHandler NextPluginHandler
	handler        NextPluginHandler
	getNextHandler func(handler PluginHandler, next NextPluginHandler) NextPluginHandler
}

func newPluginManager(handler NextPluginHandler, getNextHandler func(handler PluginHandler, next NextPluginHandler) NextPluginHandler) PluginManager {
	return &pluginManager{
		handler:        handler,
		defaultHandler: handler,
		getNextHandler: getNextHandler,
	}
}

func (pm *pluginManager) rebuildHandler() {
	next := pm.defaultHandler
	n := len(pm.handlers)
	for i := n - 1; i >= 0; i-- {
		next = pm.getNextHandler(pm.handlers[i], next)
	}
	pm.handler = next
}

func (pm *pluginManager) Handler() NextPluginHandler {
	pm.RLock()
	defer pm.RUnlock()
	return pm.handler
}

func (pm *pluginManager) Use(handler ...PluginHandler) {
	pm.Lock()
	defer pm.Unlock()
	pm.handlers = append(pm.handlers, handler...)
	pm.rebuildHandler()
}

func (pm *pluginManager) Unuse(handler ...PluginHandler) {
	pm.Lock()
	defer pm.Unlock()
	rebuild := false
	var handlers []PluginHandler
	for _, h := range pm.handlers {
		hp := reflect.ValueOf(h).Pointer()
		for _, h2 := range handler {
			if hp == reflect.ValueOf(h2).Pointer() {
				h = nil
				rebuild = true
				break
			}
		}
		if h != nil {
			handlers = append(handlers, h)
		}
	}
	pm.handlers = handlers
	if rebuild {
		pm.rebuildHandler()
	}
}
