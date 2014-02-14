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
 * LastModified: Feb 2, 2014                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
)

var cookieJar, _ = cookiejar.New(nil)

type HttpClient struct {
	*BaseClient
}

type HttpTransporter struct {
	*http.Client
	keepAlive        bool
	keepAliveTimeout int
}

type HttpContext struct {
	uri  string
	body io.ReadCloser
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

func (client *HttpClient) SetTLSClientConfig(config *tls.Config) bool {
	transport, ok := client.Http().Transport.(*http.Transport)
	if ok {
		transport.TLSClientConfig = config
	}
	return ok
}

func (client *HttpClient) KeepAlive() bool {
	if transport, ok := client.Http().Transport.(*http.Transport); ok {
		return !transport.DisableKeepAlives
	}
	return client.Transporter.(*HttpTransporter).keepAlive
}

func (client *HttpClient) SetKeepAlive(enable bool) {
	if transport, ok := client.Http().Transport.(*http.Transport); ok {
		transport.DisableKeepAlives = !enable
		client.Transporter.(*HttpTransporter).keepAlive = enable
	}
}

func (client *HttpClient) KeepAliveTimeout() int {
	return client.Transporter.(*HttpTransporter).keepAliveTimeout
}

func (client *HttpClient) SetKeepAliveTimeout(timeout int) {
	client.Transporter.(*HttpTransporter).keepAliveTimeout = timeout
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
	return &HttpTransporter{&http.Client{Jar: cookieJar}, true, 300}
}

func (h *HttpTransporter) GetInvokeContext(uri string) (interface{}, error) {
	return &HttpContext{uri: uri}, nil
}

func (h *HttpTransporter) SendData(context interface{}, data []byte, success bool) error {
	if success {
		context := context.(*HttpContext)
		req, err := http.NewRequest("POST", context.uri, bytes.NewReader(data))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/hprose")
		if h.keepAlive {
			req.Header.Set("Connection", "keep-alive")
			req.Header.Set("Keep-Alive", strconv.Itoa(h.keepAliveTimeout))
		}
		resp, err := h.Do(req)
		if err != nil {
			return err
		}
		context.body = resp.Body
	}
	return nil
}

func (h *HttpTransporter) GetInputStream(context interface{}) (BufReader, error) {
	return bufio.NewReader(context.(*HttpContext).body), nil
}

func (h *HttpTransporter) EndInvoke(context interface{}, success bool) error {
	return context.(*HttpContext).body.Close()
}
