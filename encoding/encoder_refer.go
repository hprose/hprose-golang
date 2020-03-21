/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/encoder_refer.go                                |
|                                                          |
| LastModified: Mar 21, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

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

func (r *encoderRefer) Write(enc *Encoder, p interface{}) (ok bool, err error) {
	writer := enc.writer
	var i uint64
	if i, ok = r.ref[p]; ok {
		if err = writer.WriteByte(TagRef); err == nil {
			if err = writeUint64(writer, i); err == nil {
				err = writer.WriteByte(TagSemicolon)
			}
		}
	}
	return
}

func (r *encoderRefer) WriteString(enc *Encoder, s string) (ok bool, err error) {
	writer := enc.writer
	var i uint64
	if i, ok = r.sref[s]; ok {
		if err = writer.WriteByte(TagRef); err == nil {
			if err = writeUint64(writer, i); err == nil {
				err = writer.WriteByte(TagSemicolon)
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
