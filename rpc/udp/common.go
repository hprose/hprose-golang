/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/udp/common.go                                        |
|                                                          |
| LastModified: May 18, 2022                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package udp

import (
	"hash/crc32"
	"net"
)

type data struct {
	Index int
	Body  []byte
	Error error
	Addr  *net.UDPAddr
}

func makeHeader(length int, index int) (header [8]byte) {
	header[7] = byte(index & 0xff)
	header[6] = byte(index >> 8 & 0xff)
	header[5] = byte(length & 0xff)
	header[4] = byte(length >> 8 & 0xff)
	crc := crc32.ChecksumIEEE(header[4:])
	header[3] = byte(crc & 0xff)
	header[2] = byte(crc >> 8 & 0xff)
	header[1] = byte(crc >> 16 & 0xff)
	header[0] = byte(crc >> 24 & 0xff)
	return
}

func parseHeader(header []byte) (length int, index int, ok bool) {
	index = int(header[7]) | int(header[6])<<8
	length = int(header[5]) | int(header[4])<<8
	crc := uint32(header[3]) | uint32(header[2])<<8 | uint32(header[1])<<16 | uint32(header[0])<<24
	if crc32.ChecksumIEEE(header[4:]) != crc {
		index = -1
		length = 0
		return
	}
	if ok = (header[6]&0x80 == 0); !ok {
		index &= 0x7fff
	}
	return
}
