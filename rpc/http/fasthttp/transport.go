/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/http/fasthttp/transport.go                           |
|                                                          |
| LastModified: Mar 7, 2022                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package fasthttp

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/hprose/hprose-golang/v3/internal/convert"
	"github.com/hprose/hprose-golang/v3/rpc/core"
	"github.com/hprose/hprose-golang/v3/rpc/http/cookie"
	"github.com/valyala/fasthttp"
)

type Transport struct {
	DisableHTTPHeader bool
	Header            http.Header
	FastHTTPClient    fasthttp.Client
	compression       bool
	keepAlive         bool
	*cookieManager
}

func addRequestHeader(dest *fasthttp.RequestHeader, src http.Header) {
	for key, values := range src {
		for _, value := range values {
			dest.Add(key, value)
		}
	}
}

func getResponseHeader(src *fasthttp.ResponseHeader) (dest http.Header) {
	dest = make(http.Header)
	src.VisitAll(func(key, value []byte) {
		dest.Add(string(key), string(value))
	})
	return
}

func (trans *Transport) Transport(ctx context.Context, request []byte) (response []byte, err error) {
	clientContext := core.GetClientContext(ctx)
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetMethod("POST")
	req.SetRequestURI(clientContext.URL.String())
	req.SetBody(request)
	if !trans.DisableHTTPHeader {
		if trans.Header != nil {
			addRequestHeader(&req.Header, trans.Header)
		}
		if header, ok := clientContext.Items().GetInterface("httpRequestHeaders").(http.Header); ok {
			addRequestHeader(&req.Header, header)
		}
	}
	if trans.keepAlive {
		req.Header.Set("Connection", "keep-alive")
	} else {
		req.Header.Set("Connection", "close")
	}
	if trans.compression {
		req.Header.Set("Accept-Encoding", "gzip")
	}
	if trans.cookieManager != nil {
		trans.loadCookie(req, clientContext.URL)
	}
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	if deadline, ok := ctx.Deadline(); ok {
		err = trans.FastHTTPClient.DoDeadline(req, resp, deadline)
	} else {
		err = trans.FastHTTPClient.Do(req, resp)
	}
	if err != nil {
		return nil, err
	}
	if trans.cookieManager != nil {
		trans.saveCookie(resp)
	}
	clientContext.Items().Set("httpStatusCode", resp.Header.StatusCode())
	clientContext.Items().Set("httpStatusText", string(resp.Header.StatusMessage()))
	switch resp.Header.StatusCode() {
	case fasthttp.StatusOK:
		if !trans.DisableHTTPHeader {
			clientContext.Items().Set("httpResponseHeaders", getResponseHeader(&resp.Header))
		}
		body := resp.Body()
		response := make([]byte, len(body))
		copy(response, body)
		return response, nil
	case fasthttp.StatusRequestEntityTooLarge:
		return nil, core.ErrRequestEntityTooLarge
	default:
		return nil, fmt.Errorf("%d %s", resp.Header.StatusCode(), convert.ToUnsafeString(resp.Header.StatusMessage()))
	}
}

func (trans *Transport) Abort() {
}

// CookieManagerOption returns the CookieManagerOption
func (trans *Transport) CookieManagerOption() cookie.CookieManagerOption {
	switch trans.cookieManager {
	case nil:
		return cookie.NoCookieManager
	case globalCookieManager:
		return cookie.GlobalCookieManager
	default:
		return cookie.ClientCookieManager
	}
}

// SetCookieManagerOption sets the CookieManagerOption
func (trans *Transport) SetCookieManagerOption(option cookie.CookieManagerOption) {
	switch option {
	case cookie.NoCookieManager:
		trans.cookieManager = nil
	case cookie.GlobalCookieManager:
		trans.cookieManager = globalCookieManager
	default:
		trans.cookieManager = newCookieManager()
	}
}

// TLSClientConfig returns the tls.Config
func (trans *Transport) TLSClientConfig() *tls.Config {
	return trans.FastHTTPClient.TLSConfig
}

// SetTLSClientConfig sets the tls.Config
func (trans *Transport) SetTLSClientConfig(config *tls.Config) {
	trans.FastHTTPClient.TLSConfig = config
}

// KeepAlive returns the keepalive status
func (trans *Transport) KeepAlive() bool {
	return trans.keepAlive
}

// SetKeepAlive sets the keepalive status
func (trans *Transport) SetKeepAlive(enable bool) {
	trans.keepAlive = enable
}

// Compression returns the compression status
func (trans *Transport) Compression() bool {
	return trans.compression
}

// SetCompression sets the compression status
func (trans *Transport) SetCompression(enable bool) {
	trans.compression = enable
}

type transportFactory struct {
	schemes []string
}

func (factory transportFactory) Schemes() []string {
	return factory.schemes
}

func (factory transportFactory) New() core.Transport {
	return &Transport{
		keepAlive:     true,
		cookieManager: globalCookieManager,
	}
}

func RegisterTransport() {
	core.RegisterTransport("fasthttp", transportFactory{[]string{"http", "https"}})
}
