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
 * hprose/stream_common.go                                *
 *                                                        *
 * hprose stream common for Go.                           *
 *                                                        *
 * LastModified: Jan 28, 2015                             *
 * Authors: Ma Bingyao <andot@hprose.com>                 *
 *          Ore_Ash <nanohugh@gmail.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"io"
)

func sendDataOverStream(w io.Writer, data []byte) (err error) {
	n := len(data)
	var len int
	switch {
	case n > 1020 && n <= 1400:
		len = 2048
	case n > 508:
		len = 1024
	default:
		len = 512
	}
	buf := make([]byte, len)
	buf[0] = byte((n >> 24) & 0xff)
	buf[1] = byte((n >> 16) & 0xff)
	buf[2] = byte((n >> 8) & 0xff)
	buf[3] = byte(n & 0xff)
	p := len - 4
	if n <= p {
		copy(buf[4:], data)
		_, err = w.Write(buf[:n+4])
	} else {
		copy(buf[4:], data[:p])
		_, err = w.Write(buf)
		if err != nil {
			return err
		}
		_, err = w.Write(data[p:])
	}
	return err
}

func receiveDataOverStream(r io.Reader) ([]byte, error) {
	var buf [512]byte
	n, err := io.ReadAtLeast(r, buf[:], 4)
	if err != nil {
		return nil, err
	}
	length := (int(buf[0])<<24 | int(buf[1])<<16 | int(buf[2])<<8 | int(buf[3]))
	size := length - (n - 4)
	if length <= 508 {
		if size > 0 {
			_, err = io.ReadAtLeast(r, buf[n:], size)
		}
		return buf[4 : length+4], err
	}
	data := make([]byte, length)
	copy(data, buf[4:n])
	_, err = io.ReadAtLeast(r, data[n-4:], size)
	return data, err
}
