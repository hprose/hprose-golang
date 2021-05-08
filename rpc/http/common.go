/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/http/common.go                                       |
|                                                          |
| LastModified: May 8, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package http

import (
	"io"
	"io/ioutil"
	"net/http"
)

func readAll(body io.Reader, length int64) ([]byte, error) {
	if length > 0 {
		data := make([]byte, length)
		_, err := io.ReadFull(body, data)
		return data, err
	}
	if body != nil {
		return ioutil.ReadAll(body)
	}
	return nil, nil
}

func addHeader(dest http.Header, src http.Header) {
	for key, values := range src {
		for _, value := range values {
			dest.Add(key, value)
		}
	}
}
