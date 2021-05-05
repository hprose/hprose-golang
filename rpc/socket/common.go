/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/socket/common.go                                     |
|                                                          |
| LastModified: May 5, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package socket

import (
	"errors"
	"hash/crc32"
	"net"
	"time"
)

// ErrClosed represents a error.
var ErrClosed = errors.New("connection closed")

type data struct {
	Index int
	Body  []byte
	Error error
}

func makeHeader(length int, index int) (header [12]byte) {
	header[4] = byte((length >> 24 & 0xff) | 0x80)
	header[5] = byte(length >> 16 & 0xff)
	header[6] = byte(length >> 8 & 0xff)
	header[7] = byte(length & 0xff)
	header[8] = byte(index >> 24 & 0xff)
	header[9] = byte(index >> 16 & 0xff)
	header[10] = byte(index >> 8 & 0xff)
	header[11] = byte(index & 0xff)
	crc := crc32.ChecksumIEEE(header[4:])
	header[0] = byte(crc >> 24 & 0xff)
	header[1] = byte(crc >> 16 & 0xff)
	header[2] = byte(crc >> 8 & 0xff)
	header[3] = byte(crc & 0xff)
	return
}

func parseHeader(header [12]byte) (length int, index int, ok bool) {
	crc := uint32(header[0])<<24 | uint32(header[1])<<16 | uint32(header[2])<<8 | uint32(header[3])
	if crc32.ChecksumIEEE(header[4:]) != crc {
		index = -1
		return
	}
	length = int(header[4]&0x7F)<<24 | int(header[5])<<16 | int(header[6])<<8 | int(header[7])
	index = int(header[8])<<24 | int(header[9])<<16 | int(header[10])<<8 | int(header[11])
	if ok = (header[8]&0x80 == 0); !ok {
		index &= 0x7fffffff
	}
	return
}

func nextTempDelay(err error, onError func(error), tempDelay time.Duration) time.Duration {
	if ne, ok := err.(net.Error); ok && ne.Temporary() {
		if tempDelay == 0 {
			tempDelay = 5 * time.Millisecond
		} else {
			tempDelay *= 2
		}
		if max := 1 * time.Second; tempDelay > max {
			tempDelay = max
		}
		onError(err)
		time.Sleep(tempDelay)
		return tempDelay
	}
	return 0
}
