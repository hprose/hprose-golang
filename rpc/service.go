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
	AddFunction(name string, function interface{}, options ...Option) Service
	AddFunctions(names []string, functions []interface{}, options ...Option) Service
	AddMethod(name string, obj interface{}, alias string, options ...Option) Service
	AddMethods(names []string, obj interface{}, aliases []string, options ...Option) Service
	AddInstanceMethods(obj interface{}, options ...Option) Service
	AddAllMethods(obj interface{}, options ...Option) Service
	AddMissingMethod(method MissingMethod, options ...Option) Service
	AddNetRPCMethods(rcvr interface{}, options ...Option) Service
	Remove(name string) Service
	FirstFilter() Filter
	FilterByIndex(index int) Filter
	SetFilters(filters ...Filter) Service
	AddFilters(filters ...Filter) Service
	RemoveFilterByIndex(index int) Service
	RemoveFilters(filter ...Filter) Service
	AddInvokeHandlers(handlers ...InvokeHandler) Service
	AddBeforeFilterHandlers(handlers ...FilterHandler) Service
	AddAfterFilterHandlers(handlers ...FilterHandler) Service
	SetUserData(userdata map[string]interface{}) Service
	Publish(topic string, timeout time.Duration, heartbeat time.Duration) Service
	Clients
}
