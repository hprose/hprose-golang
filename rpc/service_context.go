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
 * rpc/service_context.go                                 *
 *                                                        *
 * hprose service context for Go.                         *
 *                                                        *
 * LastModified: Nov 1, 2016                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

// ServiceContext is the hprose service context
type ServiceContext interface {
	Context
	Service() Service
	Clients() Clients
	Method() *Method
	IsMissingMethod() bool
	ByRef() bool
	setMethod(method *Method)
	setIsMissingMethod(value bool)
	setByRef(value bool)
}

// BaseServiceContext is the base service context
type BaseServiceContext struct {
	BaseContext
	method          *Method
	service         Service
	isMissingMethod bool
	byRef           bool
}

// InitServiceContext initializes BaseServiceContext
func (context *BaseServiceContext) InitServiceContext(service Service) {
	context.InitBaseContext()
	context.service = service
	context.method = nil
	context.isMissingMethod = false
	context.byRef = false
}

// Method returns the method of current invoking
func (context *BaseServiceContext) Method() *Method {
	return context.method
}

// Service returns the Service interface
func (context *BaseServiceContext) Service() Service {
	return context.service
}

// Clients returns the Clients interface
func (context *BaseServiceContext) Clients() Clients {
	return context.service
}

// IsMissingMethod returns whether the current invoking is missing method
func (context *BaseServiceContext) IsMissingMethod() bool {
	return context.isMissingMethod
}

// ByRef returns whether the current invoking is parameter passing by reference.
func (context *BaseServiceContext) ByRef() bool {
	return context.byRef
}

func (context *BaseServiceContext) setMethod(method *Method) {
	context.method = method
}

func (context *BaseServiceContext) setIsMissingMethod(value bool) {
	context.isMissingMethod = value
}

func (context *BaseServiceContext) setByRef(value bool) {
	context.byRef = value
}
