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
	empty     unsafe.Pointer
	st        *reflect2.UnsafeSliceType
	emptyElem unsafe.Pointer
	tempElem  unsafe.Pointer
	readElem  func(dec *Decoder, ep unsafe.Pointer)
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
			valdec.readElem(dec, valdec.st.UnsafeGetIndex(slice, i))
		}
		switch {
		case n < length:
			for i := n; i < length; i++ {
				valdec.st.UnsafeSetIndex(slice, i, valdec.emptyElem)
			}
		case n < count:
			for i := n; i < count; i++ {
				valdec.readElem(dec, valdec.tempElem)
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
func ArrayDecoder(t reflect.Type, readElem func(dec *Decoder, ep unsafe.Pointer)) ValueDecoder {
	at := reflect2.Type2(t).(*reflect2.UnsafeArrayType)
	et := t.Elem()
	return arrayDecoder{
		at,
		at.UnsafeNew(),
		reflect2.Type2(reflect.SliceOf(et)).(*reflect2.UnsafeSliceType),
		reflect2.Type2(et).UnsafeNew(),
		reflect2.Type2(et).UnsafeNew(),
		readElem,
	}
}

func boolArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, ep unsafe.Pointer) {
		bdec.decode(dec, (*bool)(ep), dec.NextByte())
	})
}

func intArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, ep unsafe.Pointer) {
		idec.decode(dec, (*int)(ep), dec.NextByte())
	})
}

func int8ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, ep unsafe.Pointer) {
		i8dec.decode(dec, (*int8)(ep), dec.NextByte())
	})
}

func int16ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, ep unsafe.Pointer) {
		i16dec.decode(dec, (*int16)(ep), dec.NextByte())
	})
}

func int32ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, ep unsafe.Pointer) {
		i32dec.decode(dec, (*int32)(ep), dec.NextByte())
	})
}

func int64ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, ep unsafe.Pointer) {
		i64dec.decode(dec, (*int64)(ep), dec.NextByte())
	})
}

func uintArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, ep unsafe.Pointer) {
		udec.decode(dec, (*uint)(ep), dec.NextByte())
	})
}

func uint8ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, ep unsafe.Pointer) {
		u8dec.decode(dec, (*uint8)(ep), dec.NextByte())
	})
}

func uint16ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, ep unsafe.Pointer) {
		u16dec.decode(dec, (*uint16)(ep), dec.NextByte())
	})
}

func uint32ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, ep unsafe.Pointer) {
		u32dec.decode(dec, (*uint32)(ep), dec.NextByte())
	})
}

func uint64ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, ep unsafe.Pointer) {
		u64dec.decode(dec, (*uint64)(ep), dec.NextByte())
	})
}

func uintptrArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, ep unsafe.Pointer) {
		updec.decode(dec, (*uintptr)(ep), dec.NextByte())
	})
}

func float32ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, ep unsafe.Pointer) {
		f32dec.decode(dec, (*float32)(ep), dec.NextByte())
	})
}

func float64ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, ep unsafe.Pointer) {
		f64dec.decode(dec, (*float64)(ep), dec.NextByte())
	})
}

func complex64ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, ep unsafe.Pointer) {
		c64dec.decode(dec, (*complex64)(ep), dec.NextByte())
	})
}

func complex128ArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, ep unsafe.Pointer) {
		c128dec.decode(dec, (*complex128)(ep), dec.NextByte())
	})
}

func interfaceArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, ep unsafe.Pointer) {
		ifdec.decode(dec, (*interface{})(ep), dec.NextByte())
	})
}

func stringArrayDecoder(t reflect.Type) ValueDecoder {
	return ArrayDecoder(t, func(dec *Decoder, ep unsafe.Pointer) {
		sdec.decode(dec, (*string)(ep), dec.NextByte())
	})
}

func otherArrayDecoder(t reflect.Type) ValueDecoder {
	valdec := getValueDecoder(t.Elem())
	et := reflect2.Type2(t.Elem())
	return ArrayDecoder(t, func(dec *Decoder, ep unsafe.Pointer) {
		valdec.Decode(dec, et.UnsafeIndirect(ep), dec.NextByte())
	})
}
