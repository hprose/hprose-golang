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
 * rpc/service_event.go                                   *
 *                                                        *
 * hprose service event for Go.                           *
 *                                                        *
 * LastModified: Sep 11, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import "reflect"

// ServiceEvent is the service event
type ServiceEvent interface{}

type beforeInvokeEvent interface {
	OnBeforeInvoke(name string, args []reflect.Value, byref bool, context Context)
}

type beforeInvokeEvent2 interface {
	OnBeforeInvoke(name string, args []reflect.Value, byref bool, context Context) error
}

type afterInvokeEvent interface {
	OnAfterInvoke(name string, args []reflect.Value, byref bool, result []reflect.Value, context Context)
}

type afterInvokeEvent2 interface {
	OnAfterInvoke(name string, args []reflect.Value, byref bool, result []reflect.Value, context Context) error
}

type sendErrorEvent interface {
	OnSendError(err error, context Context)
}

type sendErrorEvent2 interface {
	OnSendError(err error, context Context) error
}

type subscribeEvent interface {
	OnSubscribe(topic string, id string, service Service)
}

type unsubscribeEvent interface {
	OnUnsubscribe(topic string, id string, service Service)
}
