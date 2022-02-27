/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoder_refer.go                                      |
|                                                          |
| LastModified: Feb 27, 2022                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

type encoderRefer struct {
	ref  map[interface{}]uint64
	sref map[string]uint64
	last uint64
}

func (r *encoderRefer) AddCount(count int) {
	r.last += uint64(count)
}

func (r *encoderRefer) Set(p interface{}) {
	if r.ref == nil {
		r.ref = make(map[interface{}]uint64)
	}
	r.ref[p] = r.last
	r.last++
}

func (r *encoderRefer) SetString(s string) {
	if r.sref == nil {
		r.sref = make(map[string]uint64)
	}
	r.sref[s] = r.last
	r.last++
}

func (r *encoderRefer) Write(enc *Encoder, p interface{}) (ok bool) {
	if r.ref == nil {
		return false
	}
	var i uint64
	if i, ok = r.ref[p]; ok {
		enc.buf = append(enc.buf, TagRef)
		enc.buf = AppendUint64(enc.buf, i)
		enc.buf = append(enc.buf, TagSemicolon)
	}
	return
}

func (r *encoderRefer) WriteString(enc *Encoder, s string) (ok bool) {
	if r.sref == nil {
		return false
	}
	var i uint64
	if i, ok = r.sref[s]; ok {
		enc.buf = append(enc.buf, TagRef)
		enc.buf = AppendUint64(enc.buf, i)
		enc.buf = append(enc.buf, TagSemicolon)
	}
	return
}

func (r *encoderRefer) Reset() {
	for k := range r.ref {
		delete(r.ref, k)
	}
	for k := range r.sref {
		delete(r.sref, k)
	}
	r.last = 0
}
