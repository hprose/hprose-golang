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
 * hprose/tcp_common.go                                   *
 *                                                        *
 * hprose tcp common for Go.                              *
 *                                                        *
 * LastModified: Feb 25, 2014                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"io"
	"net"
)

func sendDataOverTcp(conn net.Conn, data []byte) (err error) {
	n := len(data)
	var buflen int
	if n > 1024 {
		buflen = 2048
	} else if n > 512 {
		buflen = 1024
	} else {
		buflen = 512
	}
	buf := make([]byte, buflen)
	buf[0] = byte((n >> 24) & 0xff)
	buf[1] = byte((n >> 16) & 0xff)
	buf[2] = byte((n >> 8) & 0xff)
	buf[3] = byte(n & 0xff)
	if n <= buflen-4 {
		copy(buf[4:], data)
		_, err = conn.Write(buf[:n+4])
	} else {
		copy(buf[4:], data[:buflen-4])
		_, err = conn.Write(buf)
		if err != nil {
			return err
		}
		_, err = conn.Write(data[buflen-4:])
	}
	return err
}

func receiveDataOverTcp(conn net.Conn) ([]byte, error) {
	var buf [2048]byte
	n, err := io.ReadAtLeast(conn, buf[:], 4)
	if err != nil {
		return nil, err
	}
	length := (int(buf[0])<<24 | int(buf[1])<<16 | int(buf[2])<<8 | int(buf[3]))
	size := length - (n - 4)
	if length <= 2044 {
		if size > 0 {
			_, err = io.ReadAtLeast(conn, buf[n:], size)
		}
		return buf[4 : length+4], err
	}
	data := make([]byte, length)
	copy(data, buf[4:n])
	_, err = io.ReadAtLeast(conn, data[n-4:], size)
	return data, err
}
