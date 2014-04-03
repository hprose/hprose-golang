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
 * hprose/http_client.go                                  *
 *                                                        *
 * hprose http client for Go.                             *
 *                                                        *
 * LastModified: Apr 3, 2014                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

var cookieJar, _ = cookiejar.New(nil)

type HttpClient struct {
	*BaseClient
}

type HttpTransporter struct {
	*http.Client
}

func NewHttpClient(uri string) Client {
	client := &HttpClient{NewBaseClient(newHttpTransporter())}
	client.SetUri(uri)
	client.SetKeepAlive(true)
	return client
}

func (client *HttpClient) SetUri(uri string) {
	if u, err := url.Parse(uri); err == nil {
		if u.Scheme != "http" && u.Scheme != "https" {
			panic("This client desn't support " + u.Scheme + " scheme.")
		}
	}
	client.BaseClient.SetUri(uri)
}

func (client *HttpClient) TLSClientConfig() *tls.Config {
	if transport, ok := client.Http().Transport.(*http.Transport); ok {
		return transport.TLSClientConfig
	}
	return nil
}

func (client *HttpClient) SetTLSClientConfig(config *tls.Config) {
	if transport, ok := client.Http().Transport.(*http.Transport); ok {
		transport.TLSClientConfig = config
	}
}

func (client *HttpClient) KeepAlive() bool {
	if transport, ok := client.Http().Transport.(*http.Transport); ok {
		return !transport.DisableKeepAlives
	}
	return true
}

func (client *HttpClient) SetKeepAlive(enable bool) {
	if transport, ok := client.Http().Transport.(*http.Transport); ok {
		transport.DisableKeepAlives = !enable
	}
}

func (client *HttpClient) Compression() bool {
	if transport, ok := client.Http().Transport.(*http.Transport); ok {
		return !transport.DisableCompression
	}
	return false
}

func (client *HttpClient) SetCompression(enable bool) {
	if transport, ok := client.Http().Transport.(*http.Transport); ok {
		transport.DisableCompression = !enable
	}
}

func (client *HttpClient) MaxIdleConnsPerHost() int {
	if transport, ok := client.Http().Transport.(*http.Transport); ok {
		return transport.MaxIdleConnsPerHost
	}
	return http.DefaultMaxIdleConnsPerHost
}

func (client *HttpClient) SetMaxIdleConnsPerHost(value int) bool {
	transport, ok := client.Http().Transport.(*http.Transport)
	if ok {
		transport.MaxIdleConnsPerHost = value
	}
	return ok
}

func (client *HttpClient) Http() *http.Client {
	return client.Transporter.(*HttpTransporter).Client
}

func newHttpTransporter() *HttpTransporter {
	return &HttpTransporter{&http.Client{Jar: cookieJar}}
}

func (h *HttpTransporter) readAll(response *http.Response) (data []byte, err error) {
	if response.ContentLength > 0 {
		data = make([]byte, response.ContentLength)
		_, err = io.ReadFull(response.Body, data)
		return data, err
	}
	if response.ContentLength < 0 {
		return ioutil.ReadAll(response.Body)
	}
	return make([]byte, 0), nil
}

func (h *HttpTransporter) SendAndReceive(uri string, data []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", uri, NewBytesReader(data))
	if err != nil {
		return nil, err
	}
	req.ContentLength = int64(len(data))
	req.Header.Set("Content-Type", "application/hprose")
	resp, err := h.Do(req)
	if err != nil {
		return nil, err
	}
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, resp.Body.Close()
}
