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
 * LastModified: Feb 23, 2014                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"io"
	"net"
)

func readContentLength(conn net.Conn) (n int, err error) {
	var buf [4]byte
	_, err = io.ReadAtLeast(conn, buf[:], 4)
	if err != nil {
		return 0, err
	}
	return (int(buf[0])<<24 | int(buf[1])<<16 | int(buf[2])<<8 | int(buf[3])), nil
}

func writeContentLength(conn net.Conn, n int) (err error) {
	var buf = [4]byte{
		byte((n >> 24) & 0xff),
		byte((n >> 16) & 0xff),
		byte((n >> 8) & 0xff),
		byte(n & 0xff),
	}
	_, err = conn.Write(buf[:])
	return err
}
