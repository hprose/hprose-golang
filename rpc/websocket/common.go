/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/websocket/common.go                                  |
|                                                          |
| LastModified: May 5, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package websocket

type data struct {
	Index int
	Body  []byte
	Error error
}

func makeHeader(index int) (header [4]byte) {
	header[0] = byte(index >> 24 & 0xff)
	header[1] = byte(index >> 16 & 0xff)
	header[2] = byte(index >> 8 & 0xff)
	header[3] = byte(index & 0xff)
	return
}

func parseHeader(header []byte) (index int, ok bool) {
	index = int(header[0])<<24 | int(header[1])<<16 | int(header[2])<<8 | int(header[3])
	if ok = (header[0]&0x80 == 0); !ok {
		index &= 0x7fffffff
	}
	return
}
