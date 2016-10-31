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
 * LastModified: Oct 6, 2016                              *
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

type serviceContext struct {
	baseContext
	method          *Method
	service         Service
	isMissingMethod bool
	byRef           bool
}

func (context *serviceContext) initServiceContext(service Service) {
	context.initBaseContext()
	context.service = service
	context.method = nil
	context.isMissingMethod = false
	context.byRef = false
}

func (context *serviceContext) Method() *Method {
	return context.method
}

func (context *serviceContext) Service() Service {
	return context.service
}

func (context *serviceContext) Clients() Clients {
	return context.service
}

func (context *serviceContext) IsMissingMethod() bool {
	return context.isMissingMethod
}

func (context *serviceContext) ByRef() bool {
	return context.byRef
}

func (context *serviceContext) setMethod(method *Method) {
	context.method = method
}

func (context *serviceContext) setIsMissingMethod(value bool) {
	context.isMissingMethod = value
}

func (context *serviceContext) setByRef(value bool) {
	context.byRef = value
}
