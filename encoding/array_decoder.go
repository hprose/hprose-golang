/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/array_decoder.go                                |
|                                                          |
| LastModified: Jun 11, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

func (dec *Decoder) readIntArray(p interface{}, et reflect.Type, length int) {
	var slice []int
	setSliceHeader(unsafe.Pointer(&slice), reflect2.PtrOf(p), length)
	count := dec.ReadInt()
	dec.AddReference(slice)
	n := length
	if n > count {
		n = count
	}
	for i := 0; i < n; i++ {
		slice[i] = dec.decodeInt(et, dec.NextByte())
	}
	switch {
	case n < length:
		for i := n; i < length; i++ {
			slice[i] = 0
		}
	case n < count:
		for i := n; i < count; i++ {
			dec.decodeInt(et, dec.NextByte())
		}
	}
	dec.Skip()
}

// arrayDecoder is the implementation of ValueDecoder for [N]T.
type arrayDecoder struct {
	t         *reflect2.UnsafeArrayType
	et        reflect.Type
	empty     unsafe.Pointer
	readArray func(dec *Decoder, p interface{}, et reflect.Type, length int)
}

func (valdec arrayDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch tag {
	case TagNull, TagEmpty:
		valdec.t.UnsafeSet(reflect2.PtrOf(p), valdec.empty)
	case TagList:
		valdec.readArray(dec, p, valdec.et, valdec.t.Len())
	default:
		dec.decodeError(valdec.t.Type1(), tag)
	}
}

func (valdec arrayDecoder) Type() reflect.Type {
	return valdec.t.Type1()
}

// ArrayDecoder returns a ValueDecoder for [N]T.
func ArrayDecoder(t reflect.Type, readArray func(dec *Decoder, p interface{}, et reflect.Type, length int)) ValueDecoder {
	t2 := reflect2.Type2(t).(*reflect2.UnsafeArrayType)
	return arrayDecoder{
		t2,
		t.Elem(),
		t2.UnsafeNew(),
		readArray,
	}
}

func intArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, p interface{}, et reflect.Type, length int) {
		dec.readIntArray(p, et, length)
	})
}
