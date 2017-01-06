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
 * rpc/fasthttp/fasthttp_client.go                        *
 *                                                        *
 * hprose fasthttp client for Go.                         *
 *                                                        *
 * LastModified: Jan 7, 2017                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package fasthttp

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/hprose/hprose-golang/rpc"
	"github.com/hprose/hprose-golang/util"
	"github.com/valyala/fasthttp"
)

var httpSchemes = []string{"http", "https"}

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
	rpc.BaseClient
	*cookieManager
	fasthttp.Client
	Header      fasthttp.RequestHeader
	compression bool
	keepAlive   bool
	limiter     rpc.Limiter
}

// NewFastHTTPClient is the constructor of FastHTTPClient
func NewFastHTTPClient(uri ...string) (client *FastHTTPClient) {
	client = new(FastHTTPClient)
	client.InitBaseClient()
	client.limiter.InitLimiter()
	client.cookieManager = globalCookieManager
	if rpc.DisableGlobalCookie {
		client.cookieManager = newCookieManager()
	}
	client.compression = false
	client.keepAlive = true
	client.SetURIList(uri)
	client.SendAndReceive = client.sendAndReceive
	return
}

func newFastHTTPClient(uri ...string) rpc.Client {
	return NewFastHTTPClient(uri...)
}

// SetURIList sets a list of server addresses
func (client *FastHTTPClient) SetURIList(uriList []string) {
	if rpc.CheckAddresses(uriList, httpSchemes) == "https" {
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	client.BaseClient.SetURIList(uriList)
}

// TLSClientConfig returns the tls.Config in hprose client
func (client *FastHTTPClient) TLSClientConfig() *tls.Config {
	return client.TLSConfig
}

// SetTLSClientConfig sets the tls.Config
func (client *FastHTTPClient) SetTLSClientConfig(config *tls.Config) {
	client.TLSConfig = config
}

// MaxConcurrentRequests returns max concurrent request count
func (client *FastHTTPClient) MaxConcurrentRequests() int {
	return client.limiter.MaxConcurrentRequests
}

// SetMaxConcurrentRequests sets max concurrent request count
func (client *FastHTTPClient) SetMaxConcurrentRequests(value int) {
	client.limiter.MaxConcurrentRequests = value
}

// KeepAlive returns the keepalive status of hprose client
func (client *FastHTTPClient) KeepAlive() bool {
	return client.keepAlive
}

// SetKeepAlive sets the keepalive status of hprose client
func (client *FastHTTPClient) SetKeepAlive(enable bool) {
	client.keepAlive = enable
}

// Compression returns the compression status of hprose client
func (client *FastHTTPClient) Compression() bool {
	return client.compression
}

// SetCompression sets the compression status of hprose client
func (client *FastHTTPClient) SetCompression(enable bool) {
	client.compression = enable
}

func (client *FastHTTPClient) limit() {
	client.limiter.L.Lock()
	client.limiter.Limit()
	client.limiter.L.Unlock()
}

func (client *FastHTTPClient) unlimit() {
	client.limiter.L.Lock()
	client.limiter.Unlimit()
	client.limiter.L.Unlock()
}

func (client *FastHTTPClient) sendAndReceive(
	data []byte, context *rpc.ClientContext) ([]byte, error) {
	client.limit()
	defer client.unlimit()
	req := fasthttp.AcquireRequest()
	client.Header.CopyTo(&req.Header)
	header, ok := context.Get("httpHeader").(http.Header)
	if ok && header != nil {
		for key, values := range header {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}
	}
	req.Header.SetMethod("POST")
	client.loadCookie(req, client.URL())
	req.SetRequestURI(client.URI())
	req.SetBody(data)
	req.Header.SetContentLength(len(data))
	req.Header.SetContentType("application/hprose")
	if client.keepAlive {
		req.Header.Set("Connection", "keep-alive")
	} else {
		req.Header.Set("Connection", "close")
	}
	if client.compression {
		req.Header.Set("Accept-Encoding", "gzip")
	}
	resp := fasthttp.AcquireResponse()
	err := client.Client.DoTimeout(req, resp, context.Timeout)
	if err != nil {
		data = nil
	} else {
		data = resp.Body()
		client.saveCookie(resp)
	}
	header = make(http.Header)
	resp.Header.VisitAll(func(key, value []byte) {
		header.Add(util.ByteString(key), util.ByteString(value))
	})
	context.Set("httpHeader", header)
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(resp)
	return data, err
}

func init() {
	rpc.RegisterClientFactory("http", newFastHTTPClient)
	rpc.RegisterClientFactory("https", newFastHTTPClient)
}
