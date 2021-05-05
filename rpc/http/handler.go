/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/http/handler.go                                      |
|                                                          |
| LastModified: May 5, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package http

import (
	"context"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

type Handler struct {
	Service                      *core.Service
	OnError                      func(error)
	P3P                          bool
	GET                          bool
	CrossDomain                  bool
	Header                       http.Header
	AccessControlAllowOrigins    map[string]bool
	LastModified                 string
	Etag                         string
	crossDomainXMLFile           string
	crossDomainXMLContent        []byte
	clientAccessPolicyXMLFile    string
	clientAccessPolicyXMLContent []byte
}

func (h *Handler) onError(err error) {
	if h.OnError != nil {
		h.OnError(err)
	}
}

// AddAccessControlAllowOrigin add access control allow origin.
func (h *Handler) AddAccessControlAllowOrigin(origins ...string) {
	for _, origin := range origins {
		h.AccessControlAllowOrigins[origin] = true
	}
}

// RemoveAccessControlAllowOrigin remove access control allow origin.
func (h *Handler) RemoveAccessControlAllowOrigin(origins ...string) {
	for _, origin := range origins {
		delete(h.AccessControlAllowOrigins, origin)
	}
}

// CrossDomainXMLFile return the cross domain xml file.
func (h *Handler) CrossDomainXMLFile() string {
	return h.crossDomainXMLFile
}

// CrossDomainXMLContent return the cross domain xml content.
func (h *Handler) CrossDomainXMLContent() []byte {
	return h.crossDomainXMLContent
}

// ClientAccessPolicyXMLFile return the client access policy xml file.
func (h *Handler) ClientAccessPolicyXMLFile() string {
	return h.clientAccessPolicyXMLFile
}

// ClientAccessPolicyXMLContent return the client access policy xml content.
func (h *Handler) ClientAccessPolicyXMLContent() []byte {
	return h.clientAccessPolicyXMLContent
}

// SetCrossDomainXMLFile set the cross domain xml file.
func (h *Handler) SetCrossDomainXMLFile(filename string) {
	h.crossDomainXMLFile = filename
	h.crossDomainXMLContent, _ = ioutil.ReadFile(filename)
}

// SetClientAccessPolicyXMLFile set the client access policy xml file.
func (h *Handler) SetClientAccessPolicyXMLFile(filename string) {
	h.clientAccessPolicyXMLFile = filename
	h.clientAccessPolicyXMLContent, _ = ioutil.ReadFile(filename)
}

// SetCrossDomainXMLContent set the cross domain xml content.
func (h *Handler) SetCrossDomainXMLContent(content []byte) {
	h.crossDomainXMLFile = ""
	h.crossDomainXMLContent = content
}

// SetClientAccessPolicyXMLContent set the client access policy xml content.
func (h *Handler) SetClientAccessPolicyXMLContent(content []byte) {
	h.clientAccessPolicyXMLFile = ""
	h.clientAccessPolicyXMLContent = content
}

// BindContext to the http server.
func (h *Handler) BindContext(ctx context.Context, server core.Server) {
	s := server.(*http.Server)
	s.Handler = h
	s.BaseContext = func(l net.Listener) context.Context {
		return ctx
	}
}

// ServeHTTP implements the http.Handler interface.
func (h *Handler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if request.ContentLength > int64(h.Service.MaxRequestLength) {
		response.WriteHeader(http.StatusRequestEntityTooLarge)
		return
	}
	if request.Method == "GET" {
		if h.clientAccessPolicyXMLHandler(response, request) ||
			h.crossDomainXMLHandler(response, request) {
			return
		}
		if !h.GET {
			response.WriteHeader(http.StatusForbidden)
			return
		}
	}
	data, err := readAll(request.Body, request.ContentLength)
	if err != nil {
		h.onError(err)
	}
	if err = request.Body.Close(); err != nil {
		h.onError(err)
	}
	serviceContext := h.getServiceContext(response, request)
	ctx := core.WithContext(request.Context(), serviceContext)
	result, err := h.Service.Handle(ctx, data)
	if err != nil {
		h.onError(err)
	}
	response.Header().Set("Content-Length", strconv.Itoa(len(result)))
	h.sendHeader(serviceContext, response, request)
	_, err = response.Write(result)
	if err != nil {
		h.onError(err)
	}
}

func (h *Handler) xmlFileHandler(response http.ResponseWriter, request *http.Request, path string, content []byte) bool {
	if content == nil || strings.ToLower(request.URL.Path) != path {
		return false
	}
	if request.Header.Get("if-modified-since") == h.LastModified &&
		request.Header.Get("if-none-match") == h.Etag {
		response.WriteHeader(304)
	} else {
		contentLength := len(content)
		header := response.Header()
		header.Set("Last-Modified", h.LastModified)
		header.Set("Etag", h.Etag)
		header.Set("Content-Type", "text/xml")
		header.Set("Content-Length", strconv.Itoa(contentLength))
		_, _ = response.Write(content)
	}
	return true
}

func (h *Handler) crossDomainXMLHandler(response http.ResponseWriter, request *http.Request) bool {
	return h.xmlFileHandler(response, request, "/crossdomain.xml", h.crossDomainXMLContent)
}

func (h *Handler) clientAccessPolicyXMLHandler(response http.ResponseWriter, request *http.Request) bool {
	return h.xmlFileHandler(response, request, "/clientaccesspolicy.xml", h.clientAccessPolicyXMLContent)
}

func (h *Handler) sendHeader(serviceContext *core.ServiceContext, response http.ResponseWriter, request *http.Request) {
	responseHeader := response.Header()
	responseHeader.Set("Content-Type", "text/plain")
	if h.P3P {
		responseHeader.Set("P3P",
			`CP="CAO DSP COR CUR ADM DEV TAI PSA PSD IVAi IVDi `+
				`CONi TELo OTPi OUR DELi SAMi OTRi UNRi PUBi IND PHY ONL `+
				`UNI PUR FIN COM NAV INT DEM CNT STA POL HEA PRE GOV"`)
	}
	if h.CrossDomain {
		origin := request.Header.Get("origin")
		if origin != "" && origin != "null" {
			if len(h.AccessControlAllowOrigins) == 0 ||
				h.AccessControlAllowOrigins[origin] {
				responseHeader.Set("Access-Control-Allow-Origin", origin)
				responseHeader.Set("Access-Control-Allow-Credentials", "true")
			}
		} else {
			responseHeader.Set("Access-Control-Allow-Origin", "*")
		}
	}
	if h.Header != nil {
		addHeader(responseHeader, h.Header)
	}
	if header, ok := serviceContext.Items().Get("httpResponseHeaders"); ok {
		if header, ok := header.(http.Header); ok {
			addHeader(responseHeader, header)
		}
	}
	if code := serviceContext.Items().GetInt("httpStatusCode"); code != 0 {
		response.WriteHeader(code)
	}
}

func (h *Handler) getServiceContext(response http.ResponseWriter, request *http.Request) *core.ServiceContext {
	serviceContext := core.NewServiceContext(h.Service)
	serviceContext.Items().Set("request", request)
	serviceContext.Items().Set("response", response)
	serviceContext.Items().Set("httpRequestHeaders", request.Header)
	serviceContext.LocalAddr, _ = net.ResolveTCPAddr("tcp", request.Host)
	serviceContext.RemoteAddr, _ = net.ResolveTCPAddr("tcp", request.RemoteAddr)
	serviceContext.Handler = h
	return serviceContext
}

type handlerFactory struct {
	serverTypes []reflect.Type
}

func (factory handlerFactory) ServerTypes() []reflect.Type {
	return factory.serverTypes
}

func (factory handlerFactory) New(service *core.Service) core.Handler {
	return &Handler{
		Service:                   service,
		P3P:                       true,
		GET:                       true,
		CrossDomain:               true,
		AccessControlAllowOrigins: make(map[string]bool),
		LastModified:              time.Now().UTC().Format(time.RFC1123),
		Etag:                      `"` + strconv.FormatInt(rand.Int63(), 16) + `"`,
	}
}

func init() {
	core.RegisterHandler("http", handlerFactory{
		[]reflect.Type{
			reflect.TypeOf((*http.Server)(nil)),
		},
	})
}
