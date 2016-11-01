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
 * rpc/fasthttp/fasthttp_service.go                       *
 *                                                        *
 * hprose fasthttp service for Go.                        *
 *                                                        *
 * LastModified: Nov 1, 2016                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package fasthttp

import (
	"reflect"
	"strings"
	"sync"

	"github.com/hprose/hprose-golang/rpc"
	"github.com/hprose/hprose-golang/util"
	"github.com/valyala/fasthttp"
)

var fasthttpContextType = reflect.TypeOf((*FastHTTPContext)(nil))
var fasthttpRequestCtxType = reflect.TypeOf((*fasthttp.RequestCtx)(nil))

// FastHTTPContext is the hprose fasthttp context
type FastHTTPContext struct {
	rpc.BaseServiceContext
	RequestCtx *fasthttp.RequestCtx
}

// InitFastHTTPContext initializes FastHTTPContext
func (context *FastHTTPContext) InitFastHTTPContext(
	service rpc.Service, ctx *fasthttp.RequestCtx) {
	context.InitServiceContext(service)
	context.RequestCtx = ctx
}

// FastHTTPService is the hprose fasthttp service
type FastHTTPService struct {
	rpc.BaseHTTPService
	contextPool sync.Pool
}

type fastSendHeaderEvent interface {
	OnSendHeader(context *FastHTTPContext)
}

type fastSendHeaderEvent2 interface {
	OnSendHeader(context *FastHTTPContext) error
}

func fasthttpFixArguments(args []reflect.Value, context rpc.ServiceContext) {
	i := len(args) - 1
	switch args[i].Type() {
	case fasthttpContextType:
		if c, ok := context.(*FastHTTPContext); ok {
			args[i] = reflect.ValueOf(c)
		}
	case fasthttpRequestCtxType:
		if c, ok := context.(*FastHTTPContext); ok {
			args[i] = reflect.ValueOf(c.RequestCtx)
		}
	default:
		rpc.DefaultFixArguments(args, context)
	}
}

// NewFastHTTPService is the constructor of FastHTTPService
func NewFastHTTPService() (service *FastHTTPService) {
	service = new(FastHTTPService)
	service.InitBaseHTTPService()
	service.contextPool = sync.Pool{
		New: func() interface{} { return new(FastHTTPContext) },
	}
	service.FixArguments = fasthttpFixArguments
	return
}

func (service *FastHTTPService) acquireContext() *FastHTTPContext {
	return service.contextPool.Get().(*FastHTTPContext)
}

func (service *FastHTTPService) releaseContext(context *FastHTTPContext) {
	service.contextPool.Put(context)
}

func (service *FastHTTPService) xmlFileHandler(
	ctx *fasthttp.RequestCtx, path string, context []byte) bool {
	requestPath := util.ByteString(ctx.Path())
	if context == nil || strings.ToLower(requestPath) != path {
		return false
	}
	ifModifiedSince := util.ByteString(ctx.Request.Header.Peek("if-modified-since"))
	ifNoneMatch := util.ByteString(ctx.Request.Header.Peek("if-none-match"))
	if ifModifiedSince == service.LastModified && ifNoneMatch == service.Etag {
		ctx.SetStatusCode(304)
	} else {
		contentLength := len(context)
		ctx.Response.Header.Set("Last-Modified", service.LastModified)
		ctx.Response.Header.Set("Etag", service.Etag)
		ctx.Response.Header.SetContentType("text/xml")
		ctx.Response.Header.Set("Content-Length", util.Itoa(contentLength))
		ctx.SetBody(context)
	}
	return true
}

func (service *FastHTTPService) crossDomainXMLHandler(
	ctx *fasthttp.RequestCtx) bool {
	path := "/crossdomain.xml"
	context := service.CrossDomainXMLContent()
	return service.xmlFileHandler(ctx, path, context)
}

func (service *FastHTTPService) clientAccessPolicyXMLHandler(
	ctx *fasthttp.RequestCtx) bool {
	path := "/clientaccesspolicy.xml"
	context := service.ClientAccessPolicyXMLContent()
	return service.xmlFileHandler(ctx, path, context)
}

func (service *FastHTTPService) fireSendHeaderEvent(
	context *FastHTTPContext) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = rpc.NewPanicError(e)
		}
	}()
	switch event := service.Event.(type) {
	case fastSendHeaderEvent:
		event.OnSendHeader(context)
	case fastSendHeaderEvent2:
		err = event.OnSendHeader(context)
	}
	return err
}

func (service *FastHTTPService) sendHeader(
	context *FastHTTPContext) (err error) {
	if err = service.fireSendHeaderEvent(context); err != nil {
		return err
	}
	ctx := context.RequestCtx
	ctx.Response.Header.Set("Content-Type", "text/plain")
	if service.P3P {
		ctx.Response.Header.Set("P3P",
			`CP="CAO DSP COR CUR ADM DEV TAI PSA PSD IVAi IVDi `+
				`CONi TELo OTPi OUR DELi SAMi OTRi UNRi PUBi IND PHY ONL `+
				`UNI PUR FIN COM NAV INT DEM CNT STA POL HEA PRE GOV"`)
	}
	if service.CrossDomain {
		origin := util.ByteString(ctx.Request.Header.Peek("origin"))
		if origin != "" && origin != "null" {
			if len(service.AccessControlAllowOrigins) == 0 ||
				service.AccessControlAllowOrigins[origin] {
				ctx.Response.Header.Set("Access-Control-Allow-Origin", origin)
				ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
			}
		} else {
			ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		}
	}
	return nil
}

// ServeFastHTTP is the hprose fasthttp handler method
func (service *FastHTTPService) ServeFastHTTP(ctx *fasthttp.RequestCtx) {
	if service.clientAccessPolicyXMLHandler(ctx) ||
		service.crossDomainXMLHandler(ctx) {
		return
	}
	context := service.acquireContext()
	context.InitFastHTTPContext(service, ctx)
	var resp []byte
	if err := service.sendHeader(context); err == nil {
		switch util.ByteString(ctx.Method()) {
		case "GET":
			if service.GET {
				resp = service.DoFunctionList(context)
			} else {
				ctx.SetStatusCode(403)
			}
		case "POST":
			resp = service.Handle(ctx.PostBody(), context)
		}
	} else {
		resp = service.EndError(err, context)
	}
	context.RequestCtx = nil
	service.releaseContext(context)
	ctx.Response.Header.Set("Content-Length", util.Itoa(len(resp)))
	ctx.SetBody(resp)
}
