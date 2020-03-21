/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/encoder_refer.go                             |
|                                                          |
| LastModified: Mar 21, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"github.com/hprose/hprose-golang/v3/io"
)

type encoderRefer struct {
	ref  map[interface{}]uint64
	sref map[string]uint64
	last uint64
}

func newEncoderRefer() *encoderRefer {
	return &encoderRefer{
		ref:  make(map[interface{}]uint64),
		sref: make(map[string]uint64),
		last: 0,
	}
}

func (r *encoderRefer) AddCount(count int) {
	r.last += uint64(count)
}

func (r *encoderRefer) Set(p interface{}) {
	r.ref[p] = r.last
	r.last++
}

func (r *encoderRefer) SetString(s string) {
	r.sref[s] = r.last
	r.last++
}

func (r *encoderRefer) Write(writer io.BytesWriter, p interface{}) (ok bool, err error) {
	var i uint64
	if i, ok = r.ref[p]; ok {
		if err = writer.WriteByte(io.TagRef); err == nil {
			if err = writeUint64(writer, i); err == nil {
				err = writer.WriteByte(io.TagSemicolon)
			}
		}
	}
	return
}

func (r *encoderRefer) WriteString(writer io.BytesWriter, s string) (ok bool, err error) {
	var i uint64
	if i, ok = r.sref[s]; ok {
		if err = writer.WriteByte(io.TagRef); err == nil {
			if err = writeUint64(writer, i); err == nil {
				err = writer.WriteByte(io.TagSemicolon)
			}
		}
	}
	return
}

func (r *encoderRefer) Reset() {
	r.ref = make(map[interface{}]uint64)
	r.sref = make(map[string]uint64)
	r.last = 0
}
