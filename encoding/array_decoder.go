/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/array_decoder.go                                |
|                                                          |
| LastModified: May 14, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

// arrayDecoder is the implementation of ValueDecoder for [N]T.
type arrayDecoder struct {
	at         *reflect2.UnsafeArrayType
	et         reflect2.Type
	empty      unsafe.Pointer
	emptyElem  unsafe.Pointer
	decodeElem DecodeHandler
}

func (valdec arrayDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch tag {
	case TagNull, TagEmpty:
		valdec.at.UnsafeSet(reflect2.PtrOf(p), valdec.empty)
	case TagList:
		length := valdec.at.Len()
		count := dec.ReadInt()
		array := reflect2.PtrOf(p)
		dec.AddReference(p)
		n := length
		if n > count {
			n = count
		}
		et := valdec.et.Type1()
		for i := 0; i < n; i++ {
			valdec.decodeElem(dec, et, valdec.at.UnsafeGetIndex(array, i))
		}
		switch {
		case n < length:
			for i := n; i < length; i++ {
				valdec.at.UnsafeSetIndex(array, i, valdec.emptyElem)
			}
		case n < count:
			temp := valdec.et.UnsafeNew()
			for i := n; i < count; i++ {
				valdec.decodeElem(dec, et, temp)
			}
		}
		dec.Skip()
	default:
		dec.defaultDecode(valdec.at.Type1(), p, tag)
	}
}

func (valdec arrayDecoder) Type() reflect.Type {
	return valdec.at.Type1()
}

// makeArrayDecoder returns a arrayDecoder for [N]T.
func makeArrayDecoder(t reflect.Type, decodeElem DecodeHandler) arrayDecoder {
	at := reflect2.Type2(t).(*reflect2.UnsafeArrayType)
	et := reflect2.Type2(t.Elem())
	return arrayDecoder{
		at,
		et,
		at.UnsafeNew(),
		et.UnsafeNew(),
		decodeElem,
	}
}

type byteArrayDecoder struct {
	arrayDecoder
}

func (valdec byteArrayDecoder) copy(p interface{}, data []byte) {
	count := len(data)
	length := valdec.at.Len()
	slice := *(*[]byte)(unsafeToSlice(p, length))
	copy(slice, data)
	if length > count {
		for i := count; i < length; i++ {
			slice[i] = 0
		}
	}
}

func (valdec byteArrayDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch tag {
	case TagBytes:
		data := dec.UnsafeNext(dec.ReadInt())
		dec.Skip()
		valdec.copy(p, data)
		dec.AddReference(p)
	case TagUTF8Char:
		data, _ := dec.readStringAsBytes(1)
		valdec.copy(p, data)
	case TagString:
		if dec.IsSimple() {
			data, _ := dec.readStringAsBytes(dec.ReadInt())
			dec.Skip()
			valdec.copy(p, data)
		} else {
			valdec.copy(p, reflect2.UnsafeCastString(dec.ReadString()))
		}
	default:
		valdec.arrayDecoder.Decode(dec, p, tag)
	}
}

// makeByteArrayDecoder returns a ValueDecoder for [N]byte.
func makeByteArrayDecoder(t reflect.Type) byteArrayDecoder {
	return byteArrayDecoder{makeArrayDecoder(t, uint8Decode)}
}

func getArrayDecoder(t reflect.Type) ValueDecoder {
	et := t.Elem()
	if et.Kind() == reflect.Uint8 {
		return makeByteArrayDecoder(t)
	}
	return makeArrayDecoder(t, GetDecodeHandler(et))
}
