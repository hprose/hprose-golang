/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/plugin_manager.go                               |
|                                                          |
| LastModified: Feb 18, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import "sync"

// NextPluginHandler must be one of NextInvokeHandler or NextIOHandler.
type NextPluginHandler interface{}

// PluginHandler must be one of InvokeHandler or IOHandler.
type PluginHandler interface{}

func separatePluginHandlers(handlers []PluginHandler) (invokeHandlers []PluginHandler, ioHandler []PluginHandler) {
	for _, handler := range handlers {
		switch handler.(type) {
		case InvokeHandler:
			invokeHandlers = append(invokeHandlers, handler)
		case IOHandler:
			ioHandler = append(ioHandler, handler)
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
	handlers       []PluginHandler
	rwlock         sync.RWMutex
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
	pm.rwlock.RLock()
	n := len(pm.handlers)
	for i := n - 1; i >= 0; i-- {
		next = pm.getNextHandler(pm.handlers[i], next)
	}
	pm.rwlock.RUnlock()
	pm.handler = next
}

func (pm *pluginManager) Handler() NextPluginHandler {
	return pm.handler
}

func (pm *pluginManager) Use(handler ...PluginHandler) {
	pm.rwlock.Lock()
	pm.handlers = append(pm.handlers, handler...)
	pm.rwlock.Unlock()
	pm.rebuildHandler()
}

func (pm *pluginManager) Unuse(handler ...PluginHandler) {
	rebuild := false
	pm.rwlock.Lock()
	var handlers []PluginHandler
	for _, h := range pm.handlers {
		for _, h2 := range handler {
			if h == h2 {
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
	pm.rwlock.Unlock()
	if rebuild {
		pm.rebuildHandler()
	}
}
