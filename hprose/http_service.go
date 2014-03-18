/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.net/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * hprose/http_service.go                                 *
 *                                                        *
 * hprose http service for Go.                            *
 *                                                        *
 * LastModified: Mar 18, 2014                             *
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

type HttpServiceEvent interface {
	ServiceEvent
	OnSendHeader(response http.ResponseWriter, request *http.Request)
}

type HttpService struct {
	*BaseService
	P3PEnabled                   bool
	GetEnabled                   bool
	CrossDomainEnabled           bool
	lastModified                 string
	etag                         string
	crossDomainXmlFile           string
	crossDomainXmlContent        []byte
	clientAccessPolicyXmlFile    string
	clientAccessPolicyXmlContent []byte
}

type httpArgsFixer struct{}

func (httpArgsFixer) FixArgs(args []reflect.Value, lastParamType reflect.Type, context interface{}) []reflect.Value {
	if request, ok := context.(*http.Request); ok && lastParamType.String() == "*http.Request" {
		return append(args, reflect.ValueOf(request))
	}
	return fixArgs(args, lastParamType, context)
}

func NewHttpService() *HttpService {
	t := time.Now().UTC()
	rand.Seed(t.UnixNano())
	service := &HttpService{
		BaseService:        NewBaseService(),
		P3PEnabled:         true,
		GetEnabled:         true,
		CrossDomainEnabled: true,
		lastModified:       t.Format(time.RFC1123),
		etag:               `"` + strconv.FormatInt(rand.Int63(), 16) + `"`,
	}
	service.ArgsFixer = httpArgsFixer{}
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

func (service *HttpService) sendHeader(response http.ResponseWriter, request *http.Request) {
	if service.ServiceEvent != nil {
		if event, ok := service.ServiceEvent.(HttpServiceEvent); ok {
			event.OnSendHeader(response, request)
		}
	}
	response.Header().Set("Content-Type", "text/plain")
	if service.P3PEnabled {
		response.Header().Set("P3P", `CP="CAO DSP COR CUR ADM DEV TAI PSA PSD IVAi IVDi `+
			`CONi TELo OTPi OUR DELi SAMi OTRi UNRi PUBi IND PHY ONL `+
			`UNI PUR FIN COM NAV INT DEM CNT STA POL HEA PRE GOV"`)
	}
	if service.CrossDomainEnabled {
		origin := request.Header.Get("origin")
		if origin != "" && origin != "null" {
			response.Header().Set("Access-Control-Allow-Origin", origin)
			response.Header().Set("Access-Control-Allow-Credentials", "true")
		} else {
			response.Header().Set("Access-Control-Allow-Origin", "*")
		}
	}
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

func (service *HttpService) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if service.clientAccessPolicyXmlContent != nil && service.clientAccessPolicyXmlHandler(response, request) {
		return
	}
	if service.crossDomainXmlContent != nil && service.crossDomainXmlHandler(response, request) {
		return
	}
	service.sendHeader(response, request)
	switch request.Method {
	case "GET":
		if service.GetEnabled {
			response.Write(service.doFunctionList(request))
		} else {
			response.WriteHeader(403)
		}
	case "POST":
		data, err := service.readAll(request)
		request.Body.Close()
		if err != nil {
			response.Write(service.sendError(err, request))
		}
		response.Write(service.Handle(data, request))
	}
}
