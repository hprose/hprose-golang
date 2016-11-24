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
 * rpc/http_service.go                                    *
 *                                                        *
 * hprose http service for Go.                            *
 *                                                        *
 * LastModified: Nov 24, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import (
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"sync"

	"github.com/hprose/hprose-golang/util"
)

// HTTPContext is the hprose http context
type HTTPContext struct {
	BaseServiceContext
	Response http.ResponseWriter
	Request  *http.Request
}

// InitHTTPContext initializes HTTPContext
func (context *HTTPContext) InitHTTPContext(
	service Service,
	response http.ResponseWriter,
	request *http.Request) {
	context.InitServiceContext(service)
	context.Response = response
	context.Request = request
}

// HTTPService is the hprose http service
type HTTPService struct {
	BaseHTTPService
	contextPool sync.Pool
}

type sendHeaderEvent interface {
	OnSendHeader(context *HTTPContext)
}

type sendHeaderEvent2 interface {
	OnSendHeader(context *HTTPContext) error
}

func httpFixArguments(args []reflect.Value, context ServiceContext) {
	i := len(args) - 1
	switch args[i].Type() {
	case httpContextType:
		if c, ok := context.(*HTTPContext); ok {
			args[i] = reflect.ValueOf(c)
		}
	case httpRequestType:
		if c, ok := context.(*HTTPContext); ok {
			args[i] = reflect.ValueOf(c.Request)
		}
	default:
		DefaultFixArguments(args, context)
	}
}

// NewHTTPService is the constructor of HTTPService
func NewHTTPService() (service *HTTPService) {
	service = new(HTTPService)
	service.InitHTTPService()
	return
}

// InitHTTPService initializes HTTPService
func (service *HTTPService) InitHTTPService() {
	service.InitBaseHTTPService()
	service.contextPool = sync.Pool{
		New: func() interface{} { return new(HTTPContext) },
	}
	service.FixArguments = httpFixArguments
}

func (service *HTTPService) acquireContext() (context *HTTPContext) {
	return service.contextPool.Get().(*HTTPContext)
}

func (service *HTTPService) releaseContext(context *HTTPContext) {
	service.contextPool.Put(context)
}

func (service *HTTPService) xmlFileHandler(
	response http.ResponseWriter, request *http.Request,
	path string, context []byte) bool {
	if context == nil || strings.ToLower(request.URL.Path) != path {
		return false
	}
	if request.Header.Get("if-modified-since") == service.LastModified &&
		request.Header.Get("if-none-match") == service.Etag {
		response.WriteHeader(304)
	} else {
		contentLength := len(context)
		header := response.Header()
		header.Set("Last-Modified", service.LastModified)
		header.Set("Etag", service.Etag)
		header.Set("Content-Type", "text/xml")
		header.Set("Content-Length", util.Itoa(contentLength))
		response.Write(context)
	}
	return true
}

func (service *HTTPService) crossDomainXMLHandler(
	response http.ResponseWriter, request *http.Request) bool {
	path := "/crossdomain.xml"
	context := service.crossDomainXMLContent
	return service.xmlFileHandler(response, request, path, context)
}

func (service *HTTPService) clientAccessPolicyXMLHandler(
	response http.ResponseWriter, request *http.Request) bool {
	path := "/clientaccesspolicy.xml"
	context := service.clientAccessPolicyXMLContent
	return service.xmlFileHandler(response, request, path, context)
}

func (service *HTTPService) fireSendHeaderEvent(
	context *HTTPContext) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = NewPanicError(e)
		}
	}()
	switch event := service.Event.(type) {
	case sendHeaderEvent:
		event.OnSendHeader(context)
	case sendHeaderEvent2:
		err = event.OnSendHeader(context)
	}
	return err
}

func (service *HTTPService) sendHeader(context *HTTPContext) (err error) {
	if err = service.fireSendHeaderEvent(context); err != nil {
		return err
	}
	header := context.Response.Header()
	header.Set("Content-Type", "text/plain")
	if service.P3P {
		header.Set("P3P",
			`CP="CAO DSP COR CUR ADM DEV TAI PSA PSD IVAi IVDi `+
				`CONi TELo OTPi OUR DELi SAMi OTRi UNRi PUBi IND PHY ONL `+
				`UNI PUR FIN COM NAV INT DEM CNT STA POL HEA PRE GOV"`)
	}
	if service.CrossDomain {
		origin := context.Request.Header.Get("origin")
		if origin != "" && origin != "null" {
			if len(service.AccessControlAllowOrigins) == 0 ||
				service.AccessControlAllowOrigins[origin] {
				header.Set("Access-Control-Allow-Origin", origin)
				header.Set("Access-Control-Allow-Credentials", "true")
			}
		} else {
			header.Set("Access-Control-Allow-Origin", "*")
		}
	}
	return nil
}

func readAllFromHTTPRequest(request *http.Request) ([]byte, error) {
	if request.ContentLength > 0 {
		data := make([]byte, request.ContentLength)
		_, err := io.ReadFull(request.Body, data)
		return data, err
	}
	if request.ContentLength < 0 {
		return ioutil.ReadAll(request.Body)
	}
	return nil, nil
}

// ServeHTTP is the hprose http handler method
func (service *HTTPService) ServeHTTP(
	response http.ResponseWriter, request *http.Request) {
	if service.clientAccessPolicyXMLHandler(response, request) ||
		service.crossDomainXMLHandler(response, request) {
		return
	}
	context := service.acquireContext()
	context.InitHTTPContext(service, response, request)
	var resp []byte
	err := service.sendHeader(context)
	if err == nil {
		switch request.Method {
		case "GET":
			if service.GET {
				resp = service.DoFunctionList(context)
			} else {
				response.WriteHeader(403)
			}
		case "POST":
			var req []byte
			if req, err = readAllFromHTTPRequest(request); err == nil {
				resp = service.Handle(req, context)
			}
		}
	}
	if err != nil {
		resp = service.EndError(err, context)
	}
	service.releaseContext(context)
	response.Header().Set("Content-Length", util.Itoa(len(resp)))
	response.Write(resp)
}
