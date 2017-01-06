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
 * LastModified: Jan 7, 2017                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"

	hio "github.com/hprose/hprose-golang/io"
)

var cookieJar, _ = cookiejar.New(nil)

// DisableGlobalCookie is a flag to disable global cookie
var DisableGlobalCookie = false

// HTTPClient is hprose http client
type HTTPClient struct {
	BaseClient
	http.Transport
	Header     http.Header
	httpClient http.Client
	limiter    Limiter
}

// NewHTTPClient is the constructor of HTTPClient
func NewHTTPClient(uri ...string) (client *HTTPClient) {
	client = new(HTTPClient)
	client.InitBaseClient()
	client.limiter.InitLimiter()
	client.httpClient.Transport = &client.Transport
	client.Header = make(http.Header)
	client.DisableCompression = true
	client.DisableKeepAlives = false
	client.MaxIdleConnsPerHost = 10
	client.httpClient.Jar = cookieJar
	if DisableGlobalCookie {
		client.httpClient.Jar, _ = cookiejar.New(nil)
	}
	client.SetURIList(uri)
	client.SendAndReceive = client.sendAndReceive
	return
}

func newHTTPClient(uri ...string) Client {
	return NewHTTPClient(uri...)
}

// SetURIList sets a list of server addresses
func (client *HTTPClient) SetURIList(uriList []string) {
	if CheckAddresses(uriList, httpSchemes) == "https" {
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	client.BaseClient.SetURIList(uriList)
}

// TLSClientConfig returns the tls.Config in hprose client
func (client *HTTPClient) TLSClientConfig() *tls.Config {
	return client.Transport.TLSClientConfig
}

// SetTLSClientConfig sets the tls.Config
func (client *HTTPClient) SetTLSClientConfig(config *tls.Config) {
	client.Transport.TLSClientConfig = config
}

// MaxConcurrentRequests returns max concurrent request count
func (client *HTTPClient) MaxConcurrentRequests() int {
	return client.limiter.MaxConcurrentRequests
}

// SetMaxConcurrentRequests sets max concurrent request count
func (client *HTTPClient) SetMaxConcurrentRequests(value int) {
	client.limiter.MaxConcurrentRequests = value
}

// KeepAlive returns the keepalive status of hprose client
func (client *HTTPClient) KeepAlive() bool {
	return !client.DisableKeepAlives
}

// SetKeepAlive sets the keepalive status of hprose client
func (client *HTTPClient) SetKeepAlive(enable bool) {
	client.DisableKeepAlives = !enable
}

// Compression returns the compression status of hprose client
func (client *HTTPClient) Compression() bool {
	return !client.DisableCompression
}

// SetCompression sets the compression status of hprose client
func (client *HTTPClient) SetCompression(enable bool) {
	client.DisableCompression = !enable
}

func (client *HTTPClient) readAll(
	response *http.Response) (data []byte, err error) {
	if response.ContentLength > 0 {
		data = make([]byte, response.ContentLength)
		_, err = io.ReadFull(response.Body, data)
		return data, err
	}
	if response.ContentLength < 0 {
		return ioutil.ReadAll(response.Body)
	}
	return nil, nil
}

func (client *HTTPClient) limit() {
	client.limiter.L.Lock()
	client.limiter.Limit()
	client.limiter.L.Unlock()
}

func (client *HTTPClient) unlimit() {
	client.limiter.L.Lock()
	client.limiter.Unlimit()
	client.limiter.L.Unlock()
}

func (client *HTTPClient) sendAndReceive(
	data []byte, context *ClientContext) ([]byte, error) {
	client.limit()
	defer client.unlimit()
	req, err := http.NewRequest("POST", client.uri, hio.NewByteReader(data))
	if err != nil {
		return nil, err
	}
	for key, values := range client.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	header, ok := context.Get("httpHeader").(http.Header)
	if ok && header != nil {
		for key, values := range header {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}
	}
	req.ContentLength = int64(len(data))
	req.Header.Set("Content-Type", "application/hprose")
	client.httpClient.Timeout = context.Timeout
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	context.Set("httpHeader", resp.Header)
	data, err = ioutil.ReadAll(resp.Body)
	if err == nil {
		err = resp.Body.Close()
	}
	return data, err
}
