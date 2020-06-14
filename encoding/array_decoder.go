/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/array_decoder.go                                |
|                                                          |
| LastModified: Jun 14, 2020                               |
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
	et         reflect.Type
	empty      unsafe.Pointer
	st         *reflect2.UnsafeSliceType
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
		slice := reflect2.PtrOf(sliceHeader{reflect2.PtrOf(p), length, length})
		dec.AddReference(p)
		n := length
		if n > count {
			n = count
		}
		for i := 0; i < n; i++ {
			valdec.decodeElem(dec, valdec.et, valdec.st.UnsafeGetIndex(slice, i))
		}
		switch {
		case n < length:
			for i := n; i < length; i++ {
				valdec.st.UnsafeSetIndex(slice, i, valdec.emptyElem)
			}
		case n < count:
			temp := reflect2.Type2(valdec.et).UnsafeNew()
			for i := n; i < count; i++ {
				valdec.decodeElem(dec, valdec.et, temp)
			}
		}
		dec.Skip()
	default:
		dec.decodeError(valdec.at.Type1(), tag)
	}
}

func (valdec arrayDecoder) Type() reflect.Type {
	return valdec.at.Type1()
}

// ArrayDecoder returns a ValueDecoder for [N]T.
func ArrayDecoder(t reflect.Type, decodeElem DecodeHandler) ValueDecoder {
	at := reflect2.Type2(t).(*reflect2.UnsafeArrayType)
	et := t.Elem()
	return arrayDecoder{
		at,
		et,
		at.UnsafeNew(),
		reflect2.Type2(reflect.SliceOf(et)).(*reflect2.UnsafeSliceType),
		reflect2.Type2(et).UnsafeNew(),
		decodeElem,
	}
}

func boolArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, boolDecode)
}

func intArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, intDecode)
}

func int8ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, int8Decode)
}

func int16ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, int16Decode)
}

func int32ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, int32Decode)
}

func int64ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, int64Decode)
}

func uintArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, uintDecode)
}

func uint8ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, uint8Decode)
}

func uint16ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, uint16Decode)
}

func uint32ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, uint32Decode)
}

func uint64ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, uint64Decode)
}

func uintptrArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, uintptrDecode)
}

func float32ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, float32Decode)
}

func float64ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, float64Decode)
}

func complex64ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, complex64Decode)
}

func complex128ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, complex128Decode)
}

func interfaceArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, interfaceDecode)
}

func stringArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, stringDecode)
}

func otherArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, otherDecode(t))
}

type byteArrayDecoder struct {
	arrayDecoder
}

func (valdec byteArrayDecoder) copy(p interface{}, data []byte) {
	count := len(data)
	length := valdec.at.Len()
	slice := *(*[]byte)(unsafe.Pointer(&sliceHeader{reflect2.PtrOf(p), length, length}))
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

// ByteArrayDecoder returns a ValueDecoder for [N]byte.
func ByteArrayDecoder(t reflect.Type) ValueDecoder {
	return byteArrayDecoder{uint8ArrayDecoder(t).(arrayDecoder)}
}
