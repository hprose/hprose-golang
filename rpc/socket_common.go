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
 * rpc/socket_common.go                                   *
 *                                                        *
 * hprose socket common for Go.                           *
 *                                                        *
 * LastModified: Jan 7, 2017                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import (
	"io"
	"net"
	"time"

	hio "github.com/hprose/hprose-golang/io"
	"github.com/hprose/hprose-golang/util"
)

type packet struct {
	fullDuplex bool
	id         [4]byte
	body       []byte
}

type socketResponse struct {
	data []byte
	err  error
}

func sendData(writer io.Writer, data packet) (err error) {
	n := len(data.body)
	i := 4
	var len int
	switch {
	case n > 1016 && n <= 1400:
		len = 2048
	case n > 504:
		len = 1024
	default:
		len = 512
	}
	buf := hio.AcquireBytes(len)
	if data.fullDuplex {
		util.FromUint32(buf, uint32(n)|0x80000000)
		buf[4] = data.id[0]
		buf[5] = data.id[1]
		buf[6] = data.id[2]
		buf[7] = data.id[3]
		i = 8
	} else {
		util.FromUint32(buf, uint32(n))
	}
	p := len - i
	if n <= p {
		copy(buf[i:], data.body)
		_, err = writer.Write(buf[:n+i])
		hio.ReleaseBytes(buf)
	} else {
		copy(buf[i:], data.body[:p])
		_, err = writer.Write(buf)
		hio.ReleaseBytes(buf)
		if err != nil {
			return err
		}
		_, err = writer.Write(data.body[p:])
	}
	return err
}

func recvData(reader io.Reader, data *packet) (err error) {
	header := data.id[:]
	if _, err = io.ReadAtLeast(reader, header, 4); err != nil {
		return
	}
	size := util.ToUint32(header)
	data.fullDuplex = (size&0x80000000 != 0)
	if data.fullDuplex {
		size &= 0x7FFFFFFF
		data.fullDuplex = true
		data.body = nil
		if _, err = io.ReadAtLeast(reader, data.id[:], 4); err != nil {
			return
		}
	}
	if cap(data.body) >= int(size) {
		data.body = data.body[:size]
	} else {
		data.body = make([]byte, size)
	}
	_, err = io.ReadAtLeast(reader, data.body, int(size))
	return
}

func hdSendData(writer io.Writer, data []byte) (err error) {
	n := len(data)
	const i = 4
	var len int
	switch {
	case n > 1020 && n <= 1400:
		len = 2048
	case n > 508:
		len = 1024
	default:
		len = 512
	}
	buf := hio.AcquireBytes(len)
	util.FromUint32(buf, uint32(n))
	p := len - i
	if n <= p {
		copy(buf[i:], data)
		_, err = writer.Write(buf[:n+i])
		hio.ReleaseBytes(buf)
	} else {
		copy(buf[i:], data[:p])
		_, err = writer.Write(buf)
		hio.ReleaseBytes(buf)
		if err != nil {
			return err
		}
		_, err = writer.Write(data[p:])
	}
	return err
}

func hdRecvData(reader io.Reader, buf []byte) (data []byte, err error) {
	var header [4]byte
	if _, err = io.ReadAtLeast(reader, header[:], 4); err != nil {
		return
	}
	size := util.ToUint32(header[:])
	if cap(buf) >= int(size) {
		data = buf[:size]
	} else {
		data = make([]byte, size)
	}
	_, err = io.ReadAtLeast(reader, data, int(size))
	return
}

func nextTempDelay(
	err error, event ServiceEvent, tempDelay time.Duration) time.Duration {
	if ne, ok := err.(net.Error); ok && ne.Temporary() {
		if tempDelay == 0 {
			tempDelay = 5 * time.Millisecond
		} else {
			tempDelay *= 2
		}
		if max := 1 * time.Second; tempDelay > max {
			tempDelay = max
		}
		FireErrorEvent(event, err, nil)
		time.Sleep(tempDelay)
		return tempDelay
	}
	return 0
}
