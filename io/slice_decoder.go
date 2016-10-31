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
 * io/slice_decoder.go                                    *
 *                                                        *
 * hprose slice decoder for Go.                           *
 *                                                        *
 * LastModified: Oct 25, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import (
	"errors"
	"reflect"
)

func readBytesAsSlice(r *Reader, v reflect.Value) {
	if v.Type().Elem().Kind() != reflect.Uint8 {
		panic(errors.New("cannot be converted []byte to " + v.Type().String()))
	}
	b := v.Bytes()
	n := cap(b)
	l := r.readLength()
	if n >= l {
		b = b[:l]
		v.SetLen(l)
	} else {
		b = make([]byte, l)
		v.SetBytes(b)
	}
	if !r.Simple {
		setReaderRef(r, v)
	}
	if _, err := r.Read(b); err != nil {
		panic(err)
	}
	r.readByte()
}

func readListAsSlice(r *Reader, v reflect.Value) {
	n := v.Cap()
	l := r.ReadCount()
	if n >= l {
		v.SetLen(l)
	} else {
		v.Set(reflect.MakeSlice(v.Type(), l, l))
	}
	if !r.Simple {
		setReaderRef(r, v)
	}
	for i := 0; i < l; i++ {
		r.ReadValue(v.Index(i))
	}
	r.readByte()
}

func readRefAsSlice(r *Reader, v reflect.Value) {
	ref := r.readRef()
	if b, ok := ref.([]byte); ok {
		reflect.Copy(v, reflect.ValueOf(b))
		return
	}
	if s, ok := ref.(reflect.Value); ok {
		if s.Kind() == reflect.Slice {
			v.Set(s)
		} else {
			reflect.Copy(v, s)
		}
		return
	}
	panic(errors.New("value of type " +
		reflect.TypeOf(ref).String() +
		" cannot be converted to type slice"))
}

var sliceDecoders = [256]func(r *Reader, v reflect.Value){
	TagNull:  nilDecoder,
	TagEmpty: nilDecoder,
	TagBytes: readBytesAsSlice,
	TagList:  readListAsSlice,
	TagRef:   readRefAsSlice,
}

func sliceDecoder(r *Reader, v reflect.Value, tag byte) {
	decoder := sliceDecoders[tag]
	if decoder != nil {
		decoder(r, v)
		return
	}
	castError(tag, v.Type().String())
}
