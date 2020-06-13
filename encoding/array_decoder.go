/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/array_decoder.go                                |
|                                                          |
| LastModified: Jun 13, 2020                               |
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
	at        *reflect2.UnsafeArrayType
	et        reflect.Type
	empty     unsafe.Pointer
	st        *reflect2.UnsafeSliceType
	emptyElem unsafe.Pointer
	tempElem  unsafe.Pointer
	readElem  func(dec *Decoder, et reflect.Type, ep unsafe.Pointer)
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
			valdec.readElem(dec, valdec.et, valdec.st.UnsafeGetIndex(slice, i))
		}
		switch {
		case n < length:
			for i := n; i < length; i++ {
				valdec.st.UnsafeSetIndex(slice, i, valdec.emptyElem)
			}
		case n < count:
			for i := n; i < count; i++ {
				valdec.readElem(dec, valdec.et, valdec.tempElem)
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
func ArrayDecoder(t reflect.Type, readElem func(dec *Decoder, et reflect.Type, ep unsafe.Pointer)) ValueDecoder {
	at := reflect2.Type2(t).(*reflect2.UnsafeArrayType)
	et := t.Elem()
	return arrayDecoder{
		at,
		et,
		at.UnsafeNew(),
		reflect2.Type2(reflect.SliceOf(et)).(*reflect2.UnsafeSliceType),
		reflect2.Type2(et).UnsafeNew(),
		reflect2.Type2(et).UnsafeNew(),
		readElem,
	}
}

func boolArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*bool)(ep) = dec.decodeBool(et, dec.NextByte())
	})
}

func intArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*int)(ep) = dec.decodeInt(et, dec.NextByte())
	})
}

func int8ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*int8)(ep) = dec.decodeInt8(et, dec.NextByte())
	})
}

func int16ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*int16)(ep) = dec.decodeInt16(et, dec.NextByte())
	})
}

func int32ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*int32)(ep) = dec.decodeInt32(et, dec.NextByte())
	})
}

func int64ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*int64)(ep) = dec.decodeInt64(et, dec.NextByte())
	})
}

func uintArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*uint)(ep) = dec.decodeUint(et, dec.NextByte())
	})
}

func uint8ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*uint8)(ep) = dec.decodeUint8(et, dec.NextByte())
	})
}

func uint16ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*uint16)(ep) = dec.decodeUint16(et, dec.NextByte())
	})
}

func uint32ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*uint32)(ep) = dec.decodeUint32(et, dec.NextByte())
	})
}

func uint64ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*uint64)(ep) = dec.decodeUint64(et, dec.NextByte())
	})
}

func uintptrArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*uintptr)(ep) = dec.decodeUintptr(et, dec.NextByte())
	})
}

func float32ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*float32)(ep) = dec.decodeFloat32(et, dec.NextByte())
	})
}

func float64ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*float64)(ep) = dec.decodeFloat64(et, dec.NextByte())
	})
}

func complex64ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*complex64)(ep) = dec.decodeComplex64(et, dec.NextByte())
	})
}

func complex128ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*complex128)(ep) = dec.decodeComplex128(et, dec.NextByte())
	})
}

func interfaceArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*interface{})(ep) = dec.decodeInterface(dec.NextByte())
	})
}

func stringArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*string)(ep) = dec.decodeString(et, dec.NextByte())
	})
}

func otherArrayDecoder(t reflect.Type) ValueDecoder {
	valdec := getValueDecoder(t.Elem())
	et2 := reflect2.Type2(t.Elem())
	return ArrayDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		valdec.Decode(dec, et2.UnsafeIndirect(ep), dec.NextByte())
	})
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
