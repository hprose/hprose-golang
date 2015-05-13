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
 * LastModified: May 13, 2015                             *
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

var DisableGlobalCookie = false

type HttpClient struct {
	*BaseClient
}

type HttpTransporter struct {
	*http.Client
	Header *http.Header
}

func NewHttpClient(uri string) Client {
	client := &HttpClient{NewBaseClient(newHttpTransporter())}
	client.Client = client
	client.SetUri(uri)
	client.SetKeepAlive(true)
	return client
}

func (client *HttpClient) Close() {}

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

func (client *HttpClient) Http() *http.Client {
	return client.Transporter.(*HttpTransporter).Client
}

func (client *HttpClient) Header() *http.Header {
	return client.Transporter.(*HttpTransporter).Header
}

func (client *HttpClient) transport() *http.Transport {
	return client.Http().Transport.(*http.Transport)
}

func (client *HttpClient) TLSClientConfig() *tls.Config {
	return client.transport().TLSClientConfig
}

func (client *HttpClient) SetTLSClientConfig(config *tls.Config) {
	client.transport().TLSClientConfig = config
}

func (client *HttpClient) KeepAlive() bool {
	return !client.transport().DisableKeepAlives
}

func (client *HttpClient) SetKeepAlive(enable bool) {
	client.transport().DisableKeepAlives = !enable
}

func (client *HttpClient) Compression() bool {
	return !client.transport().DisableCompression
}

func (client *HttpClient) SetCompression(enable bool) {
	client.transport().DisableCompression = !enable
}

func (client *HttpClient) MaxIdleConnsPerHost() int {
	return client.transport().MaxIdleConnsPerHost
}

func (client *HttpClient) SetMaxIdleConnsPerHost(value int) {
	client.transport().MaxIdleConnsPerHost = value
}

func newHttpTransporter() *HttpTransporter {
	tr := &http.Transport{
		DisableCompression:  true,
		DisableKeepAlives:   false,
		MaxIdleConnsPerHost: 4}
	jar := cookieJar
	if DisableGlobalCookie {
		jar, _ = cookiejar.New(nil)
	}
	return &HttpTransporter{&http.Client{Jar: jar, Transport: tr}, &http.Header{}}
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
