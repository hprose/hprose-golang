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
 * LastModified: May 25, 2015                             *
 * Authors: Ma Bingyao <andot@hprose.com>                 *
 *          Ore_Ash <nanohugh@gmail.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"net"
	"time"
)

// StreamService is the base service for TcpService and UnixService
type StreamService struct {
	*BaseService
	timeout      interface{}
	readTimeout  interface{}
	readBuffer   interface{}
	writeTimeout interface{}
	writeBuffer  interface{}
}

// StreamContext is the hprose stream context for service
type StreamContext struct {
	*BaseContext
	net.Conn
}

func newStreamService() (service *StreamService) {
	service = new(StreamService)
	service.BaseService = NewBaseService()
	return
}

// SetTimeout for stream service
func (service *StreamService) SetTimeout(d time.Duration) {
	service.timeout = d
}

// SetReadTimeout for stream service
func (service *StreamService) SetReadTimeout(d time.Duration) {
	service.readTimeout = d
}

// SetReadBuffer for stream service
func (service *StreamService) SetReadBuffer(bytes int) {
	service.readBuffer = bytes
}

// SetWriteTimeout for stream service
func (service *StreamService) SetWriteTimeout(d time.Duration) {
	service.writeTimeout = d
}

// SetWriteBuffer for stream service
func (service *StreamService) SetWriteBuffer(bytes int) {
	service.writeBuffer = bytes
}
