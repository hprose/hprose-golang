/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/http/fasthttp/cookie_manager.go                      |
|                                                          |
| LastModified: Mar 6, 2022                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package fasthttp

import (
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/hprose/hprose-golang/v3/internal/convert"
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
	resp.Header.VisitAllCookie(func(key, value []byte) {
		k := convert.ToUnsafeString(key)
		cookie := fasthttp.AcquireCookie()
		cookie.SetKeyBytes(key)
		resp.Header.Cookie(cookie)
		if c := cm.store[k]; c != nil {
			delete(cm.store, k)
			fasthttp.ReleaseCookie(c)
		}
		cm.store[k] = cookie
	})
	cm.locker.Unlock()
}

func (cm *cookieManager) loadCookie(req *fasthttp.Request, url *url.URL) {
	cm.locker.Lock()
	for k, v := range cm.store {
		if !strings.Contains(url.Host, convert.ToUnsafeString(v.Domain())) {
			continue
		}
		if strings.Index(url.Path, convert.ToUnsafeString(v.Path())) != 0 {
			continue
		}
		if (url.Scheme == "https" && v.Secure()) || !v.Secure() {
			req.Header.SetCookie(k, convert.ToUnsafeString(v.Value()))
		}
		if v.Expire().After(time.Now()) {
			delete(cm.store, k)
			fasthttp.ReleaseCookie(v)
		}
	}
	cm.locker.Unlock()
}

var globalCookieManager = newCookieManager()
