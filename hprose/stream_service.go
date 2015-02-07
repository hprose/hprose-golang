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
 * hprose/stream_service.go                               *
 *                                                        *
 * hprose stream service for Go.                          *
 *                                                        *
 * LastModified: Feb 7, 2015                              *
 * Authors: Ma Bingyao <andot@hprose.com>                 *
 *          Ore_Ash <nanohugh@gmail.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"net"
	"time"
)

type StreamService struct {
	*BaseService
	timeout      interface{}
	readTimeout  interface{}
	readBuffer   interface{}
	writeTimeout interface{}
	writeBuffer  interface{}
}

type StreamContext struct {
	*BaseContext
	net.Conn
}

func newStreamService() *StreamService {
	return &StreamService{BaseService: NewBaseService()}
}

func (service *StreamService) SetTimeout(d time.Duration) {
	service.timeout = d
}

func (service *StreamService) SetReadTimeout(d time.Duration) {
	service.readTimeout = d
}

func (service *StreamService) SetReadBuffer(bytes int) {
	service.readBuffer = bytes
}

func (service *StreamService) SetWriteTimeout(d time.Duration) {
	service.writeTimeout = d
}

func (service *StreamService) SetWriteBuffer(bytes int) {
	service.writeBuffer = bytes
}
