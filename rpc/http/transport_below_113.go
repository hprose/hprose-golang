// +build !go1.13

/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/http/transport.go                                    |
|                                                          |
| LastModified: Mar 25, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package http

import (
	"context"
	"io"
	"net/http"
)

func newRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, uri, body)
	req = req.WithContext(ctx)
	return req, err
}
