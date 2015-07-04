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
 * hprose/http_service.go                                 *
 *                                                        *
 * hprose http service for Go.                            *
 *                                                        *
 * LastModified: Jul 4, 2015                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// HttpContext is the hprose http context
type HttpContext struct {
	*BaseContext
	Response http.ResponseWriter
	Request  *http.Request
}

// HttpService is the hprose http service
type HttpService struct {
	*BaseService
	P3PEnabled                   bool
	GetEnabled                   bool
	CrossDomainEnabled           bool
	accessControlAllowOrigins    map[string]bool
	lastModified                 string
	etag                         string
	crossDomainXmlFile           string
	crossDomainXmlContent        []byte
	clientAccessPolicyXmlFile    string
	clientAccessPolicyXmlContent []byte
}

type sendHeaderEvent interface {
	OnSendHeader(context Context)
}

type sendHeader2Event interface {
	OnSendHeader(context *HttpContext)
}

type httpArgsFixer struct{}

func (httpArgsFixer) FixArgs(args []reflect.Value, lastParamType reflect.Type, context Context) []reflect.Value {
	if c, ok := context.(*HttpContext); ok {
		if lastParamType.String() == "*hprose.HttpContext" {
			return append(args, reflect.ValueOf(c))
		} else if lastParamType.String() == "*http.Request" {
			return append(args, reflect.ValueOf(c.Request))
		}
	}
	return fixArgs(args, lastParamType, context)
}

// NewHttpService is the constructor of HttpService
func NewHttpService() (service *HttpService) {
	t := time.Now().UTC()
	rand.Seed(t.UnixNano())
	service = new(HttpService)
	service.BaseService = NewBaseService()
	service.P3PEnabled = true
	service.GetEnabled = true
	service.CrossDomainEnabled = true
	service.accessControlAllowOrigins = make(map[string]bool)
	service.lastModified = t.Format(time.RFC1123)
	service.etag = `"` + strconv.FormatInt(rand.Int63(), 16) + `"`
	service.argsfixer = httpArgsFixer{}
	return
}

func (service *HttpService) crossDomainXmlHandler(response http.ResponseWriter, request *http.Request) bool {
	if strings.ToLower(request.URL.Path) == "/crossdomain.xml" {
		if request.Header.Get("if-modified-since") == service.lastModified &&
			request.Header.Get("if-none-match") == service.etag {
			response.WriteHeader(304)
		} else {
			response.Header().Set("Last-Modified", service.lastModified)
			response.Header().Set("Etag", service.etag)
			response.Header().Set("Content-Type", "text/xml")
			response.Header().Set("Content-Length", strconv.Itoa(len(service.crossDomainXmlContent)))
			response.Write(service.crossDomainXmlContent)
		}
		return true
	}
	return false
}

func (service *HttpService) clientAccessPolicyXmlHandler(response http.ResponseWriter, request *http.Request) bool {
	if strings.ToLower(request.URL.Path) == "/clientaccesspolicy.xml" {
		if request.Header.Get("if-modified-since") == service.lastModified &&
			request.Header.Get("if-none-match") == service.etag {
			response.WriteHeader(304)
		} else {
			response.Header().Set("Last-Modified", service.lastModified)
			response.Header().Set("Etag", service.etag)
			response.Header().Set("Content-Type", "text/xml")
			response.Header().Set("Content-Length", strconv.Itoa(len(service.clientAccessPolicyXmlContent)))
			response.Write(service.clientAccessPolicyXmlContent)
		}
		return true
	}
	return false
}

func (service *HttpService) sendHeader(context *HttpContext) {
	if service.ServiceEvent != nil {
		if event, ok := service.ServiceEvent.(sendHeaderEvent); ok {
			event.OnSendHeader(context)
		} else if event, ok := service.ServiceEvent.(sendHeader2Event); ok {
			event.OnSendHeader(context)
		}
	}
	context.Response.Header().Set("Content-Type", "text/plain")
	if service.P3PEnabled {
		context.Response.Header().Set("P3P", `CP="CAO DSP COR CUR ADM DEV TAI PSA PSD IVAi IVDi `+
			`CONi TELo OTPi OUR DELi SAMi OTRi UNRi PUBi IND PHY ONL `+
			`UNI PUR FIN COM NAV INT DEM CNT STA POL HEA PRE GOV"`)
	}
	if service.CrossDomainEnabled {
		origin := context.Request.Header.Get("origin")
		if origin != "" && origin != "null" {
			if len(service.accessControlAllowOrigins) == 0 || service.accessControlAllowOrigins[origin] {
				context.Response.Header().Set("Access-Control-Allow-Origin", origin)
				context.Response.Header().Set("Access-Control-Allow-Credentials", "true")
			}
		} else {
			context.Response.Header().Set("Access-Control-Allow-Origin", "*")
		}
	}
}

// AddAccessControlAllowOrigin add access control allow origin
func (service *HttpService) AddAccessControlAllowOrigin(origin string) {
	service.accessControlAllowOrigins[origin] = true
}

// RemoveAccessControlAllowOrigin remove access control allow origin
func (service *HttpService) RemoveAccessControlAllowOrigin(origin string) {
	delete(service.accessControlAllowOrigins, origin)
}

// CrossDomainXmlFile return the cross domain xml file
func (service *HttpService) CrossDomainXmlFile() string {
	return service.crossDomainXmlFile
}

// CrossDomainXmlContent return the cross domain xml content
func (service *HttpService) CrossDomainXmlContent() []byte {
	return service.crossDomainXmlContent
}

// ClientAccessPolicyXmlFile return the client access policy xml file
func (service *HttpService) ClientAccessPolicyXmlFile() string {
	return service.clientAccessPolicyXmlFile
}

// ClientAccessPolicyXmlContent return the client access policy xml content
func (service *HttpService) ClientAccessPolicyXmlContent() []byte {
	return service.clientAccessPolicyXmlContent
}

// SetCrossDomainXmlFile set the cross domain xml file
func (service *HttpService) SetCrossDomainXmlFile(filename string) {
	service.crossDomainXmlFile = filename
	service.crossDomainXmlContent, _ = ioutil.ReadFile(filename)
}

// SetClientAccessPolicyXmlFile set the client access policy xml file
func (service *HttpService) SetClientAccessPolicyXmlFile(filename string) {
	service.clientAccessPolicyXmlFile = filename
	service.clientAccessPolicyXmlContent, _ = ioutil.ReadFile(filename)
}

// SetCrossDomainXmlContent set the cross domain xml content
func (service *HttpService) SetCrossDomainXmlContent(content []byte) {
	service.crossDomainXmlFile = ""
	service.crossDomainXmlContent = content
}

// SetClientAccessPolicyXmlContent set the client access policy xml content
func (service *HttpService) SetClientAccessPolicyXmlContent(content []byte) {
	service.clientAccessPolicyXmlFile = ""
	service.clientAccessPolicyXmlContent = content
}

func (service *HttpService) readAll(request *http.Request) (data []byte, err error) {
	if request.ContentLength > 0 {
		data = make([]byte, request.ContentLength)
		_, err = io.ReadFull(request.Body, data)
		return data, err
	}
	if request.ContentLength < 0 {
		return ioutil.ReadAll(request.Body)
	}
	return make([]byte, 0), nil
}

// Serve ...
func (service *HttpService) Serve(response http.ResponseWriter, request *http.Request, userData map[string]interface{}) {
	if service.clientAccessPolicyXmlContent != nil && service.clientAccessPolicyXmlHandler(response, request) {
		return
	}
	if service.crossDomainXmlContent != nil && service.crossDomainXmlHandler(response, request) {
		return
	}
	context := new(HttpContext)
	context.BaseContext = NewBaseContext()
	context.Response = response
	context.Request = request
	if userData != nil {
		for k, v := range userData {
			context.SetInterface(k, v)
		}
	}
	service.sendHeader(context)
	switch request.Method {
	case "GET":
		if service.GetEnabled {
			response.Write(service.doFunctionList(context))
		} else {
			response.WriteHeader(403)
		}
	case "POST":
		data, err := service.readAll(request)
		request.Body.Close()
		if err != nil {
			response.Write(service.sendError(err, context))
		}
		response.Write(service.Handle(data, context))
	}
}

// ServeHTTP ...
func (service *HttpService) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	service.Serve(response, request, nil)
}
