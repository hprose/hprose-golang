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
 * hprose/http_client.go                                  *
 *                                                        *
 * hprose http client for Go.                             *
 *                                                        *
 * LastModified: May 24, 2015                             *
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

// DisableGlobalCookie is a flag to disable global cookie
var DisableGlobalCookie = false

// HttpClient is hprose http client
type HttpClient struct {
	*BaseClient
}

type httpTransporter struct {
	*http.Client
	*http.Header
}

// NewHttpClient is the constructor of HttpClient
func NewHttpClient(uri string) (client *HttpClient) {
	client = new(HttpClient)
	client.BaseClient = NewBaseClient(newHttpTransporter())
	client.Client = client
	client.SetUri(uri)
	client.SetKeepAlive(true)
	return
}

func newHttpClient(uri string) Client {
	return NewHttpClient(uri)
}

// Close the client
func (client *HttpClient) Close() {}

// SetUri set the uri of hprose client
func (client *HttpClient) SetUri(uri string) {
	if u, err := url.Parse(uri); err == nil {
		if u.Scheme != "http" && u.Scheme != "https" {
			panic("This client desn't support " + u.Scheme + " scheme.")
		}
		if u.Scheme == "https" {
			client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		}
	}
	client.BaseClient.SetUri(uri)
}

// Http return the http.Client in hprose client
func (client *HttpClient) Http() *http.Client {
	return client.Transporter.(*httpTransporter).Client
}

// Header return the http.Header in hprose client
func (client *HttpClient) Header() *http.Header {
	return client.Transporter.(*httpTransporter).Header
}

func (client *HttpClient) transport() *http.Transport {
	return client.Http().Transport.(*http.Transport)
}

// TLSClientConfig return the tls.Config in hprose client
func (client *HttpClient) TLSClientConfig() *tls.Config {
	return client.transport().TLSClientConfig
}

// SetTLSClientConfig set the tls.Config
func (client *HttpClient) SetTLSClientConfig(config *tls.Config) {
	client.transport().TLSClientConfig = config
}

// KeepAlive return the keepalive status of hprose client
func (client *HttpClient) KeepAlive() bool {
	return !client.transport().DisableKeepAlives
}

// SetKeepAlive set the keepalive status of hprose client
func (client *HttpClient) SetKeepAlive(enable bool) {
	client.transport().DisableKeepAlives = !enable
}

// Compression return the compression status of hprose client
func (client *HttpClient) Compression() bool {
	return !client.transport().DisableCompression
}

// SetCompression set the compression status of hprose client
func (client *HttpClient) SetCompression(enable bool) {
	client.transport().DisableCompression = !enable
}

// MaxIdleConnsPerHost return the max idle connections per host of hprose client
func (client *HttpClient) MaxIdleConnsPerHost() int {
	return client.transport().MaxIdleConnsPerHost
}

// SetMaxIdleConnsPerHost set the max idle connections per host of hprose client
func (client *HttpClient) SetMaxIdleConnsPerHost(value int) {
	client.transport().MaxIdleConnsPerHost = value
}

func newHttpTransporter() (trans *httpTransporter) {
	tr := new(http.Transport)
	tr.DisableCompression = true
	tr.DisableKeepAlives = false
	tr.MaxIdleConnsPerHost = 4
	jar := cookieJar
	if DisableGlobalCookie {
		jar, _ = cookiejar.New(nil)
	}
	client := new(http.Client)
	client.Jar = jar
	client.Transport = tr
	trans = new(httpTransporter)
	trans.Client = client
	trans.Header = new(http.Header)
	return
}

func (h *httpTransporter) readAll(response *http.Response) (data []byte, err error) {
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

// SendAndReceive send and receive the data
func (h *httpTransporter) SendAndReceive(uri string, data []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", uri, NewBytesReader(data))
	if err != nil {
		return nil, err
	}
	for key, values := range *h.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
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
