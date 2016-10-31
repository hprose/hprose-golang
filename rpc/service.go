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
 * rpc/service.go                                         *
 *                                                        *
 * hprose service for Go.                                 *
 *                                                        *
 * LastModified: Oct 27, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import "time"

// Service interface
type Service interface {
	AddFunction(name string, function interface{}, option ...Options) Service
	AddFunctions(names []string, functions []interface{}, option ...Options) Service
	AddMethod(name string, obj interface{}, alias string, option ...Options) Service
	AddMethods(names []string, obj interface{}, aliases []string, option ...Options) Service
	AddInstanceMethods(obj interface{}, option ...Options) Service
	AddAllMethods(obj interface{}, option ...Options) Service
	AddMissingMethod(method MissingMethod, option ...Options) Service
	AddNetRPCMethods(rcvr interface{}, option ...Options) Service
	Remove(name string) Service
	Filter() Filter
	FilterByIndex(index int) Filter
	SetFilter(filter ...Filter) Service
	AddFilter(filter ...Filter) Service
	RemoveFilterByIndex(index int) Service
	RemoveFilter(filter ...Filter) Service
	AddInvokeHandler(handler ...InvokeHandler) Service
	AddBeforeFilterHandler(handler ...FilterHandler) Service
	AddAfterFilterHandler(handler ...FilterHandler) Service
	SetUserData(userdata map[string]interface{}) Service
	Publish(topic string, timeout time.Duration, heartbeat time.Duration) Service
	Clients
}
