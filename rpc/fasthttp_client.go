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
 * rpc/http_client.go                                     *
 *                                                        *
 * hprose http client for Go.                             *
 *                                                        *
 * LastModified: Oct 24, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import (
	"crypto/tls"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/hprose/hprose-golang/util"
	"github.com/valyala/fasthttp"
)

type cookieManager struct {
	store  map[string]*fasthttp.Cookie
	locker sync.Mutex
}

func newCookieManager() (cm *cookieManager) {
	cm = &cookieManager{store: make(map[string]*fasthttp.Cookie)}
	return
}

func (cm *cookieManager) saveCookie(resp *fasthttp.Response) {
	cm.locker.Lock()
	defer cm.locker.Unlock()
	resp.Header.VisitAllCookie(func(key, value []byte) {
		k := util.ByteString(key)
		cookie := fasthttp.AcquireCookie()
		cookie.SetKeyBytes(key)
		resp.Header.Cookie(cookie)
		if c := cm.store[k]; c != nil {
			delete(cm.store, k)
			fasthttp.ReleaseCookie(c)
		}
		cm.store[k] = cookie
	})
}

func (cm *cookieManager) loadCookie(req *fasthttp.Request, url *url.URL) {
	cm.locker.Lock()
	for k, v := range cm.store {
		if strings.Index(url.Host, util.ByteString(v.Domain())) < 0 {
			continue
		}
		if strings.Index(url.Path, util.ByteString(v.Path())) != 0 {
			continue
		}
		if (url.Scheme == "https" && v.Secure()) || !v.Secure() {
			req.Header.SetCookie(k, util.ByteString(v.Value()))
		}
		if v.Expire().After(time.Now()) {
			delete(cm.store, k)
			fasthttp.ReleaseCookie(v)
		}
	}
	cm.locker.Unlock()
}

var globalCookieManager = newCookieManager()

// FastHTTPClient is hprose fasthttp client
type FastHTTPClient struct {
	baseClient
	limiter
	*cookieManager
	fasthttp.Client
	Header      fasthttp.RequestHeader
	compression bool
	keepAlive   bool
}

// NewFastHTTPClient is the constructor of FastHTTPClient
func NewFastHTTPClient(uri ...string) (client *FastHTTPClient) {
	client = new(FastHTTPClient)
	client.initBaseClient()
	client.initLimiter()
	client.cookieManager = globalCookieManager
	if DisableGlobalCookie {
		client.cookieManager = newCookieManager()
	}
	client.compression = false
	client.keepAlive = true
	client.SetURIList(uri)
	client.SendAndReceive = client.sendAndReceive
	return
}

func newFastHTTPClient(uri ...string) Client {
	return NewFastHTTPClient(uri...)
}

// SetURIList set a list of server addresses
func (client *FastHTTPClient) SetURIList(uriList []string) {
	if checkAddresses(uriList, httpSchemes) == "https" {
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	client.baseClient.SetURIList(uriList)
}

// TLSClientConfig return the tls.Config in hprose client
func (client *FastHTTPClient) TLSClientConfig() *tls.Config {
	return client.TLSConfig
}

// SetTLSClientConfig set the tls.Config
func (client *FastHTTPClient) SetTLSClientConfig(config *tls.Config) {
	client.TLSConfig = config
}

// KeepAlive return the keepalive status of hprose client
func (client *FastHTTPClient) KeepAlive() bool {
	return client.keepAlive
}

// SetKeepAlive set the keepalive status of hprose client
func (client *FastHTTPClient) SetKeepAlive(enable bool) {
	client.keepAlive = enable
}

// Compression return the compression status of hprose client
func (client *FastHTTPClient) Compression() bool {
	return client.compression
}

// SetCompression set the compression status of hprose client
func (client *FastHTTPClient) SetCompression(enable bool) {
	client.compression = enable
}

func (client *FastHTTPClient) sendAndReceive(
	data []byte, context *ClientContext) ([]byte, error) {
	client.cond.L.Lock()
	client.limit()
	client.cond.L.Unlock()
	req := fasthttp.AcquireRequest()
	client.Header.CopyTo(&req.Header)
	req.Header.SetMethod("POST")
	client.loadCookie(req, client.url)
	req.SetRequestURI(client.uri)
	req.SetBody(data)
	req.Header.SetContentLength(len(data))
	req.Header.SetContentType("application/hprose")
	if client.keepAlive {
		req.Header.Set("Connection", "keep-alive")
	} else {
		req.Header.Set("Connection", "close")
	}
	if client.compression {
		req.Header.Set("Content-Encoding", "gzip")
	}
	resp := fasthttp.AcquireResponse()
	err := client.Client.DoTimeout(req, resp, context.Timeout)
	if err != nil {
		data = nil
	} else {
		data = resp.Body()
		client.saveCookie(resp)
	}
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(resp)
	client.cond.L.Lock()
	client.unlimit()
	client.cond.L.Unlock()
	return data, err
}
