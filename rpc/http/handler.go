/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/http/handler.go                                      |
|                                                          |
| LastModified: Mar 7, 2022                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package http

import (
	"bytes"
	"context"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/hprose/hprose-golang/v3/internal/convert"
	"github.com/hprose/hprose-golang/v3/rpc/core"
	"github.com/valyala/fasthttp"
)

type Handler struct {
	Service                      *core.Service
	OnError                      interface{}
	P3P                          bool
	GET                          bool
	CrossDomain                  bool
	DisableHTTPHeader            bool
	Header                       http.Header
	AccessControlAllowOrigins    map[string]bool
	LastModified                 string
	Etag                         string
	crossDomainXMLFile           string
	crossDomainXMLContent        []byte
	clientAccessPolicyXMLFile    string
	clientAccessPolicyXMLContent []byte
}

// common implementation.

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
	switch s := server.(type) {
	case *http.Server:
		s.Handler = h
		s.BaseContext = func(l net.Listener) context.Context {
			return ctx
		}
	case *fasthttp.Server:
		s.Handler = h.ServeFastHTTP
	}
}

// net/http implementation.

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
		h.onError(response, request, err)
	}
	if err = request.Body.Close(); err != nil {
		h.onError(response, request, err)
	}
	serviceContext := h.getServiceContext(response, request)
	ctx := core.WithContext(request.Context(), serviceContext)
	result, err := h.Service.Handle(ctx, data)
	if err != nil {
		h.onError(response, request, err)
	}
	response.Header().Set("Content-Length", strconv.Itoa(len(result)))
	h.sendHeader(serviceContext, response, request)
	_, err = response.Write(result)
	if err != nil {
		h.onError(response, request, err)
	}
}

func (h *Handler) onError(response http.ResponseWriter, request *http.Request, err error) {
	if h.OnError != nil {
		if onError, ok := h.OnError.(func(http.ResponseWriter, *http.Request, error)); ok {
			onError(response, request, err)
		}
	}
}

func (h *Handler) xmlFileHandler(response http.ResponseWriter, request *http.Request, path string, content []byte) bool {
	if content == nil || strings.ToLower(request.URL.Path) != path {
		return false
	}
	if request.Header.Get("if-modified-since") == h.LastModified &&
		request.Header.Get("if-none-match") == h.Etag {
		response.WriteHeader(http.StatusNotModified)
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
	if !h.DisableHTTPHeader {
		if h.Header != nil {
			addHeader(responseHeader, h.Header)
		}
		if header, ok := serviceContext.Items().GetInterface("httpResponseHeaders").(http.Header); ok {
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
	if !h.DisableHTTPHeader {
		serviceContext.Items().Set("httpRequestHeaders", request.Header)
	}
	serviceContext.LocalAddr, _ = net.ResolveTCPAddr("tcp", request.Host)
	serviceContext.RemoteAddr, _ = net.ResolveTCPAddr("tcp", request.RemoteAddr)
	serviceContext.Handler = h
	return serviceContext
}

// fasthttp implementation.

func addResponseHeader(dest *fasthttp.ResponseHeader, src http.Header) {
	for key, values := range src {
		for _, value := range values {
			dest.Add(key, value)
		}
	}
}

func getRequestHeader(src *fasthttp.RequestHeader) (dest http.Header) {
	dest = make(http.Header)
	src.VisitAll(func(key, value []byte) {
		dest.Add(string(key), string(value))
	})
	return
}

// ServeFastHTTP implements the fasthttp.RequestHandler.
func (h *Handler) ServeFastHTTP(ctx *fasthttp.RequestCtx) {
	if ctx.Request.Header.ContentLength() > h.Service.MaxRequestLength {
		ctx.SetStatusCode(fasthttp.StatusRequestEntityTooLarge)
		return
	}
	if convert.ToUnsafeString(ctx.Request.Header.Method()) == "GET" {
		if h.clientAccessPolicyXMLFastHTTPHandler(ctx) ||
			h.crossDomainXMLFastHTTPHandler(ctx) {
			return
		}
		if !h.GET {
			ctx.SetStatusCode(fasthttp.StatusForbidden)
			return
		}
	}
	serviceContext := h.getFastHTTPServiceContext(ctx)
	result, err := h.Service.Handle(core.WithContext(context.Background(), serviceContext), ctx.Request.Body())
	if err != nil {
		h.onFastHTTPError(ctx, err)
	}
	ctx.Response.Header.SetContentLength(len(result))
	h.sendFastHTTPHeader(serviceContext, ctx)
	_, err = ctx.Write(result)
	if err != nil {
		h.onFastHTTPError(ctx, err)
	}
}

func (h *Handler) onFastHTTPError(ctx *fasthttp.RequestCtx, err error) {
	if h.OnError != nil {
		if onError, ok := h.OnError.(func(ctx *fasthttp.RequestCtx, err error)); ok {
			onError(ctx, err)
		}
	}
}

func (h *Handler) xmlFileFastHTTPHandler(ctx *fasthttp.RequestCtx, path string, content []byte) bool {
	if content == nil || convert.ToUnsafeString(bytes.ToLower(ctx.Path())) != path {
		return false
	}

	if convert.ToUnsafeString(ctx.Request.Header.Peek("if-modified-since")) == h.LastModified &&
		convert.ToUnsafeString(ctx.Request.Header.Peek("if-none-match")) == h.Etag {
		ctx.SetStatusCode(fasthttp.StatusNotModified)
	} else {
		contentLength := len(content)
		ctx.Response.Header.Set("Last-Modified", h.LastModified)
		ctx.Response.Header.Set("Etag", h.Etag)
		ctx.Response.Header.SetContentType("text/xml")
		ctx.Response.Header.SetContentLength(contentLength)
		_, _ = ctx.Write(content)
	}
	return true
}

func (h *Handler) crossDomainXMLFastHTTPHandler(ctx *fasthttp.RequestCtx) bool {
	return h.xmlFileFastHTTPHandler(ctx, "/crossdomain.xml", h.crossDomainXMLContent)
}

func (h *Handler) clientAccessPolicyXMLFastHTTPHandler(ctx *fasthttp.RequestCtx) bool {
	return h.xmlFileFastHTTPHandler(ctx, "/clientaccesspolicy.xml", h.clientAccessPolicyXMLContent)
}

func (h *Handler) sendFastHTTPHeader(serviceContext *core.ServiceContext, ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "text/plain")
	if h.P3P {
		ctx.Response.Header.Set("P3P",
			`CP="CAO DSP COR CUR ADM DEV TAI PSA PSD IVAi IVDi `+
				`CONi TELo OTPi OUR DELi SAMi OTRi UNRi PUBi IND PHY ONL `+
				`UNI PUR FIN COM NAV INT DEM CNT STA POL HEA PRE GOV"`)
	}
	if h.CrossDomain {
		origin := convert.ToUnsafeString(ctx.Request.Header.Peek("origin"))
		if origin != "" && origin != "null" {
			if len(h.AccessControlAllowOrigins) == 0 ||
				h.AccessControlAllowOrigins[origin] {
				ctx.Response.Header.Set("Access-Control-Allow-Origin", origin)
				ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
			}
		} else {
			ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		}
	}
	if !h.DisableHTTPHeader {
		if h.Header != nil {
			addResponseHeader(&ctx.Response.Header, h.Header)
		}
		if header, ok := serviceContext.Items().GetInterface("httpResponseHeaders").(http.Header); ok {
			addResponseHeader(&ctx.Response.Header, header)
		}
	}
	if code := serviceContext.Items().GetInt("httpStatusCode"); code != 0 {
		ctx.SetStatusCode(code)
	}
}

func (h *Handler) getFastHTTPServiceContext(ctx *fasthttp.RequestCtx) *core.ServiceContext {
	serviceContext := core.NewServiceContext(h.Service)
	serviceContext.Items().Set("requestCtx", ctx)
	if !h.DisableHTTPHeader {
		serviceContext.Items().Set("httpRequestHeaders", getRequestHeader(&ctx.Request.Header))
	}
	serviceContext.LocalAddr, _ = net.ResolveTCPAddr("tcp", convert.ToUnsafeString(ctx.Host()))
	serviceContext.RemoteAddr = ctx.RemoteAddr()
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

func RegisterHandler() {
	core.RegisterHandler("http", handlerFactory{
		[]reflect.Type{
			reflect.TypeOf((*http.Server)(nil)),
			reflect.TypeOf((*fasthttp.Server)(nil)),
		},
	})
}
