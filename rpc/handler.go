/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * rpc/handler.go                                         *
 *                                                        *
 * hprose handler manager for Go.                         *
 *                                                        *
 * LastModified: May 22, 2017                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import "reflect"

// NextInvokeHandler is the next invoke handler function
type NextInvokeHandler func(
	name string,
	args []reflect.Value,
	context Context) (results []reflect.Value, err error)

// InvokeHandler is the invoke handler function
type InvokeHandler func(
	name string,
	args []reflect.Value,
	context Context,
	next NextInvokeHandler) (results []reflect.Value, err error)

// NextFilterHandler is the next filter handler function
type NextFilterHandler func(
	request []byte,
	context Context) (response []byte, err error)

// FilterHandler is the filter handler function
type FilterHandler func(
	request []byte,
	context Context,
	next NextFilterHandler) (response []byte, err error)

type addFilterHandler struct {
	Use func(handler ...FilterHandler)
}

// handlerManager is the hprose handler manager
type handlerManager struct {
	BeforeFilter               addFilterHandler
	AfterFilter                addFilterHandler
	invokeHandlers             []InvokeHandler
	beforeFilterHandlers       []FilterHandler
	afterFilterHandlers        []FilterHandler
	defaultInvokeHandler       NextInvokeHandler
	defaultBeforeFilterHandler NextFilterHandler
	defaultAfterFilterHandler  NextFilterHandler
	invokeHandler              NextInvokeHandler
	beforeFilterHandler        NextFilterHandler
	afterFilterHandler         NextFilterHandler
	override                   struct {
		invokeHandler       NextInvokeHandler
		beforeFilterHandler NextFilterHandler
		afterFilterHandler  NextFilterHandler
	}
}

func (hm *handlerManager) initHandlerManager() {
	hm.BeforeFilter.Use = hm.AddBeforeFilterHandler
	hm.AfterFilter.Use = hm.AddAfterFilterHandler
	hm.defaultInvokeHandler = func(
		name string,
		args []reflect.Value,
		context Context) (results []reflect.Value, err error) {
		defer func() {
			if e := recover(); e != nil {
				err = NewPanicError(e)
			}
		}()
		return hm.override.invokeHandler(name, args, context)
	}
	hm.defaultBeforeFilterHandler = func(
		request []byte,
		context Context) (response []byte, err error) {
		defer func() {
			if e := recover(); e != nil {
				err = NewPanicError(e)
			}
		}()
		response, err = hm.override.beforeFilterHandler(request, context)
		return
	}
	hm.defaultAfterFilterHandler = func(
		request []byte,
		context Context) (response []byte, err error) {
		defer func() {
			if e := recover(); e != nil {
				err = NewPanicError(e)
			}
		}()
		response, err = hm.override.afterFilterHandler(request, context)
		return
	}
	hm.invokeHandler = hm.defaultInvokeHandler
	hm.beforeFilterHandler = hm.defaultBeforeFilterHandler
	hm.afterFilterHandler = hm.defaultAfterFilterHandler
}

func getNextInvokeHandler(
	next NextInvokeHandler, handler InvokeHandler) NextInvokeHandler {
	return func(name string,
		args []reflect.Value,
		context Context) (results []reflect.Value, err error) {
		defer func() {
			if e := recover(); e != nil {
				err = NewPanicError(e)
			}
		}()
		results, err = handler(name, args, context, next)
		return
	}
}

func getNextFilterHandler(
	next NextFilterHandler, handler FilterHandler) NextFilterHandler {
	return func(request []byte, context Context) (response []byte, err error) {
		defer func() {
			if e := recover(); e != nil {
				err = NewPanicError(e)
			}
		}()
		response, err = handler(request, context, next)
		return
	}
}

// AddInvokeHandler add the invoke handler
func (hm *handlerManager) AddInvokeHandler(handler ...InvokeHandler) {
	if len(handler) == 0 {
		return
	}
	hm.invokeHandlers = append(hm.invokeHandlers, handler...)
	next := hm.defaultInvokeHandler
	for i := len(hm.invokeHandlers) - 1; i >= 0; i-- {
		next = getNextInvokeHandler(next, hm.invokeHandlers[i])
	}
	hm.invokeHandler = next
}

// AddBeforeFilterHandler add the filter handler before filters
func (hm *handlerManager) AddBeforeFilterHandler(handler ...FilterHandler) {
	if len(handler) == 0 {
		return
	}
	hm.beforeFilterHandlers = append(hm.beforeFilterHandlers, handler...)
	next := hm.defaultBeforeFilterHandler
	for i := len(hm.beforeFilterHandlers) - 1; i >= 0; i-- {
		next = getNextFilterHandler(next, hm.beforeFilterHandlers[i])
	}
	hm.beforeFilterHandler = next
}

// AddAfterFilterHandler add the filter handler after filters
func (hm *handlerManager) AddAfterFilterHandler(handler ...FilterHandler) {
	if len(handler) == 0 {
		return
	}
	hm.afterFilterHandlers = append(hm.afterFilterHandlers, handler...)
	next := hm.defaultAfterFilterHandler
	for i := len(hm.afterFilterHandlers) - 1; i >= 0; i-- {
		next = getNextFilterHandler(next, hm.afterFilterHandlers[i])
	}
	hm.afterFilterHandler = next
}

// Use is a method alias of AddInvokeHandler
func (hm *handlerManager) Use(handler ...InvokeHandler) {
	hm.AddInvokeHandler(handler...)
}
