/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/http/transport.go                                    |
|                                                          |
| LastModified: Apr 18, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package http

import (
	"bytes"
	"container/list"
	"context"
	"errors"
	"net"
	"net/http"
	"net/http/cookiejar"
	"sync"
	"time"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

type Transport struct {
	Header      http.Header
	HTTPClient  http.Client
	cancelFuncs *list.List
	lock        sync.Mutex
}

func (trans *Transport) Transport(ctx context.Context, request []byte) ([]byte, error) {
	clientContext := core.GetClientContext(ctx)
	timeoutContext, cancel := context.WithTimeout(ctx, clientContext.Timeout)
	trans.lock.Lock()
	cancelFunc := trans.cancelFuncs.PushBack(cancel)
	trans.lock.Unlock()
	defer func() {
		trans.lock.Lock()
		trans.cancelFuncs.Remove(cancelFunc)
		trans.lock.Unlock()
	}()
	req, err := newRequestWithContext(timeoutContext, "POST", clientContext.URL.String(), bytes.NewReader(request))
	if err != nil {
		return nil, err
	}
	if trans.Header != nil {
		addHeader(req.Header, trans.Header)
	}
	if header, ok := clientContext.Items().Get("httpRequestHeaders"); ok {
		if header, ok := header.(http.Header); ok {
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
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		clientContext.Items().Set("httpResponseHeaders", resp.Header)
		return readAll(resp.Body, resp.ContentLength)
	}
	return nil, errors.New(resp.Status)
}

func (trans *Transport) Abort() {
	trans.lock.Lock()
	defer trans.lock.Unlock()
	var next *list.Element
	for e := trans.cancelFuncs.Front(); e != nil; e = next {
		next = e.Next()
		if cancelFunc := trans.cancelFuncs.Remove(e); cancelFunc != nil {
			cancelFunc.(context.CancelFunc)()
		}
	}
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
	transport.cancelFuncs = list.New()
	return transport
}

func init() {
	core.RegisterTransport("http", transportFactory{[]string{"http", "https"}})
}
