/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/http/transport.go                                    |
|                                                          |
| LastModified: Mar 7, 2022                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package http

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"net"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/hprose/hprose-golang/v3/rpc/core"
	"github.com/hprose/hprose-golang/v3/rpc/http/cookie"
)

type Transport struct {
	DisableHTTPHeader bool
	Header            http.Header
	HTTPClient        http.Client
}

func (trans *Transport) Transport(ctx context.Context, request []byte) ([]byte, error) {
	clientContext := core.GetClientContext(ctx)
	req, err := http.NewRequestWithContext(ctx, "POST", clientContext.URL.String(), bytes.NewReader(request))
	if err != nil {
		return nil, err
	}
	if !trans.DisableHTTPHeader {
		if trans.Header != nil {
			addHeader(req.Header, trans.Header)
		}
		if header, ok := clientContext.Items().GetInterface("httpRequestHeaders").(http.Header); ok {
			addHeader(req.Header, header)
		}
	}
	var resp *http.Response
	resp, err = trans.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	clientContext.Items().Set("httpStatusCode", resp.StatusCode)
	clientContext.Items().Set("httpStatusText", http.StatusText(resp.StatusCode))
	switch resp.StatusCode {
	case http.StatusOK:
		if !trans.DisableHTTPHeader {
			clientContext.Items().Set("httpResponseHeaders", resp.Header)
		}
		return readAll(resp.Body, resp.ContentLength)
	case http.StatusRequestEntityTooLarge:
		return nil, core.ErrRequestEntityTooLarge
	default:
		return nil, errors.New(resp.Status)
	}
}

func (trans *Transport) Abort() {
}

// CookieManagerOption returns the CookieManagerOption
func (trans *Transport) CookieManagerOption() cookie.CookieManagerOption {
	switch trans.HTTPClient.Jar {
	case nil:
		return cookie.NoCookieManager
	case globalCookieJar:
		return cookie.GlobalCookieManager
	default:
		return cookie.ClientCookieManager
	}
}

// SetCookieManagerOption sets the CookieManagerOption
func (trans *Transport) SetCookieManagerOption(option cookie.CookieManagerOption) {
	switch option {
	case cookie.NoCookieManager:
		trans.HTTPClient.Jar = nil
	case cookie.GlobalCookieManager:
		trans.HTTPClient.Jar = globalCookieJar
	default:
		trans.HTTPClient.Jar, _ = cookiejar.New(nil)
	}
}

// TLSClientConfig returns the tls.Config
func (trans *Transport) TLSClientConfig() *tls.Config {
	return trans.HTTPClient.Transport.(*http.Transport).TLSClientConfig
}

// SetTLSClientConfig sets the tls.Config
func (trans *Transport) SetTLSClientConfig(config *tls.Config) {
	trans.HTTPClient.Transport.(*http.Transport).TLSClientConfig = config
}

// KeepAlive returns the keepalive status
func (trans *Transport) KeepAlive() bool {
	return !trans.HTTPClient.Transport.(*http.Transport).DisableKeepAlives
}

// SetKeepAlive sets the keepalive status
func (trans *Transport) SetKeepAlive(enable bool) {
	trans.HTTPClient.Transport.(*http.Transport).DisableKeepAlives = !enable
}

// Compression returns the compression status
func (trans *Transport) Compression() bool {
	return !trans.HTTPClient.Transport.(*http.Transport).DisableCompression
}

// SetCompression sets the compression status
func (trans *Transport) SetCompression(enable bool) {
	trans.HTTPClient.Transport.(*http.Transport).DisableCompression = !enable
}

var globalCookieJar, _ = cookiejar.New(nil)

type transportFactory struct {
	schemes []string
}

func (factory transportFactory) Schemes() []string {
	return factory.schemes
}

func (factory transportFactory) New() core.Transport {
	transport := &Transport{}
	transport.HTTPClient.Transport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   time.Second,
			KeepAlive: time.Second * 30,
			DualStack: true,
		}).DialContext,
		MaxIdleConnsPerHost:   100,
		IdleConnTimeout:       time.Minute,
		TLSHandshakeTimeout:   time.Second,
		ExpectContinueTimeout: time.Millisecond * 500,
	}
	transport.HTTPClient.Jar = globalCookieJar
	return transport
}

func RegisterTransport() {
	core.RegisterTransport("http", transportFactory{[]string{"http", "https"}})
}
