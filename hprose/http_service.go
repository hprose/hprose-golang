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
 * LastModified: Dec 21, 2014                             *
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

type HttpContext struct {
	*BaseContext
	Response http.ResponseWriter
	Request  *http.Request
}

type HttpServiceEvent interface {
	ServiceEvent
	OnSendHeader(context *HttpContext)
}

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

func NewHttpService() *HttpService {
	t := time.Now().UTC()
	rand.Seed(t.UnixNano())
	service := &HttpService{
		BaseService:               NewBaseService(),
		P3PEnabled:                true,
		GetEnabled:                true,
		CrossDomainEnabled:        true,
		accessControlAllowOrigins: make(map[string]bool),
		lastModified:              t.Format(time.RFC1123),
		etag:                      `"` + strconv.FormatInt(rand.Int63(), 16) + `"`,
	}
	service.argsfixer = httpArgsFixer{}
	return service
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
		if event, ok := service.ServiceEvent.(HttpServiceEvent); ok {
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

func (service *HttpService) AddAccessControlAllowOrigin(origin string) {
	service.accessControlAllowOrigins[origin] = true
}

func (service *HttpService) RemoveAccessControlAllowOrigin(origin string) {
	delete(service.accessControlAllowOrigins, origin)
}

func (service *HttpService) CrossDomainXmlFile() string {
	return service.crossDomainXmlFile
}

func (service *HttpService) CrossDomainXmlContent() []byte {
	return service.crossDomainXmlContent
}

func (service *HttpService) ClientAccessPolicyXmlFile() string {
	return service.clientAccessPolicyXmlFile
}

func (service *HttpService) ClientAccessPolicyXmlContent() []byte {
	return service.clientAccessPolicyXmlContent
}

func (service *HttpService) SetCrossDomainXmlFile(filename string) {
	service.crossDomainXmlFile = filename
	service.crossDomainXmlContent, _ = ioutil.ReadFile(filename)
}

func (service *HttpService) SetClientAccessPolicyXmlFile(filename string) {
	service.clientAccessPolicyXmlFile = filename
	service.clientAccessPolicyXmlContent, _ = ioutil.ReadFile(filename)
}

func (service *HttpService) SetCrossDomainXmlContent(content []byte) {
	service.crossDomainXmlFile = ""
	service.crossDomainXmlContent = content
}

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

func (service *HttpService) Serve(response http.ResponseWriter, request *http.Request, userData map[string]interface{}) {
	if service.clientAccessPolicyXmlContent != nil && service.clientAccessPolicyXmlHandler(response, request) {
		return
	}
	if service.crossDomainXmlContent != nil && service.crossDomainXmlHandler(response, request) {
		return
	}
	context := &HttpContext{BaseContext: NewBaseContext(), Response: response, Request: request}
	if userData != nil {
		for k, v := range userData {
			context.SetInterface(k, v);
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

func (service *HttpService) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	service.Serve(response, request, nil)
}
