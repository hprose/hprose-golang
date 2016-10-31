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
 * io/array_decoder.go                                    *
 *                                                        *
 * hprose array decoder for Go.                           *
 *                                                        *
 * LastModified: Oct 25, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import (
	"errors"
	"reflect"
	"unsafe"

	"github.com/hprose/hprose-golang/util"
)

func readBytesAsArray(r *Reader, v reflect.Value) {
	if !r.Simple {
		setReaderRef(r, v)
	}
	if v.Type().Elem().Kind() != reflect.Uint8 {
		panic(errors.New("cannot be converted []byte to " + v.Type().String()))
	}
	n := v.Len()
	sliceHeader := reflect.SliceHeader{
		Data: (*emptyInterface)(unsafe.Pointer(&v)).ptr,
		Len:  n,
		Cap:  n,
	}
	b := *(*[]byte)(unsafe.Pointer(&sliceHeader))
	l := r.readLength()
	min := util.Min(n, l)
	if _, err := r.Read(b[:min]); err != nil {
		panic(err)
	}
	if l > min {
		r.Next(l - min)
	}
	r.readByte()
}

func readListAsArray(r *Reader, v reflect.Value) {
	n := v.Len()
	l := r.ReadCount()
	if !r.Simple {
		setReaderRef(r, v)
	}
	min := util.Min(n, l)
	for i := 0; i < min; i++ {
		r.ReadValue(v.Index(i))
	}
	if min < l {
		x := reflect.New(v.Type().Elem()).Elem()
		for i := min; i < l; i++ {
			r.ReadValue(x)
		}
	}
	r.readByte()
}

func readRefAsArray(r *Reader, v reflect.Value) {
	ref := r.readRef()
	if b, ok := ref.([]byte); ok {
		reflect.Copy(v, reflect.ValueOf(b))
		return
	}
	if a, ok := ref.(reflect.Value); ok {
		reflect.Copy(v, a)
		return
	}
	panic(errors.New("value of type " +
		reflect.TypeOf(ref).String() +
		" cannot be converted to type array"))
}

var arrayDecoders = [256]func(r *Reader, v reflect.Value){
	TagNull:  nilDecoder,
	TagEmpty: nilDecoder,
	TagBytes: readBytesAsArray,
	TagList:  readListAsArray,
	TagRef:   readRefAsArray,
}

func arrayDecoder(r *Reader, v reflect.Value, tag byte) {
	decoder := arrayDecoders[tag]
	if decoder != nil {
		decoder(r, v)
		return
	}
	castError(tag, v.Type().String())
}
