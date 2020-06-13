/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/slice_decoder.go                                |
|                                                          |
| LastModified: Jun 7, 2020                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"math/big"
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

// sliceDecoder is the implementation of ValueDecoder for []T.
type sliceDecoder struct {
	t        *reflect2.UnsafeSliceType
	et       reflect.Type
	empty    unsafe.Pointer
	readElem func(dec *Decoder, et reflect.Type, ep unsafe.Pointer)
}

func (valdec sliceDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch tag {
	case TagNull:
		valdec.t.UnsafeSetNil(reflect2.PtrOf(p))
	case TagEmpty:
		setSliceHeader(reflect2.PtrOf(p), valdec.empty, 0)
	case TagList:
		count := dec.ReadInt()
		slice := reflect2.PtrOf(p)
		valdec.t.UnsafeGrow(slice, count)
		dec.AddReference(p)
		for i := 0; i < count; i++ {
			valdec.readElem(dec, valdec.et, valdec.t.UnsafeGetIndex(slice, i))
		}
		dec.Skip()
	default:
		dec.decodeError(valdec.t.Type1(), tag)
	}
}

func (valdec sliceDecoder) Type() reflect.Type {
	return valdec.t.Type1()
}

// SliceDecoder returns a ValueDecoder for []T.
func SliceDecoder(t reflect.Type, readElem func(dec *Decoder, et reflect.Type, ep unsafe.Pointer)) ValueDecoder {
	return sliceDecoder{
		reflect2.Type2(t).(*reflect2.UnsafeSliceType),
		t.Elem(),
		reflect2.Type2(reflect.ArrayOf(0, t.Elem())).UnsafeNew(),
		readElem,
	}
}

func boolSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*bool)(ep) = dec.decodeBool(et, dec.NextByte())
	})
}

func intSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*int)(ep) = dec.decodeInt(et, dec.NextByte())
	})
}

func int8SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*int8)(ep) = dec.decodeInt8(et, dec.NextByte())
	})
}

func int16SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*int16)(ep) = dec.decodeInt16(et, dec.NextByte())
	})
}

func int32SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*int32)(ep) = dec.decodeInt32(et, dec.NextByte())
	})
}

func int64SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*int64)(ep) = dec.decodeInt64(et, dec.NextByte())
	})
}

func uintSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*uint)(ep) = dec.decodeUint(et, dec.NextByte())
	})
}

func uint8SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*uint8)(ep) = dec.decodeUint8(et, dec.NextByte())
	})
}

func uint16SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*uint16)(ep) = dec.decodeUint16(et, dec.NextByte())
	})
}

func uint32SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*uint32)(ep) = dec.decodeUint32(et, dec.NextByte())
	})
}

func uint64SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*uint64)(ep) = dec.decodeUint64(et, dec.NextByte())
	})
}

func uintptrSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*uintptr)(ep) = dec.decodeUintptr(et, dec.NextByte())
	})
}

func float32SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*float32)(ep) = dec.decodeFloat32(et, dec.NextByte())
	})
}

func float64SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*float64)(ep) = dec.decodeFloat64(et, dec.NextByte())
	})
}

func complex64SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*complex64)(ep) = dec.decodeComplex64(et, dec.NextByte())
	})
}

func complex128SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*complex128)(ep) = dec.decodeComplex128(et, dec.NextByte())
	})
}

func interfaceSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*interface{})(ep) = dec.decodeInterface(dec.NextByte())
	})
}

func stringSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(*string)(ep) = dec.decodeString(et, dec.NextByte())
	})
}

func bigIntSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(**big.Int)(ep) = dec.decodeBigInt(et, dec.NextByte())
	})
}

func bigFloatSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(**big.Float)(ep) = dec.decodeBigFloat(et, dec.NextByte())
	})
}

func bigRatSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		*(**big.Rat)(ep) = dec.decodeBigRat(et, dec.NextByte())
	})
}

func otherSliceDecoder(t reflect.Type) ValueDecoder {
	valdec := getValueDecoder(t.Elem())
	et2 := reflect2.Type2(t.Elem())
	return SliceDecoder(t, func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		valdec.Decode(dec, et2.UnsafeIndirect(ep), dec.NextByte())
	})
}

var (
	bsdec    sliceDecoder
	isdec    sliceDecoder
	i8sdec   sliceDecoder
	i16sdec  sliceDecoder
	i32sdec  sliceDecoder
	i64sdec  sliceDecoder
	usdec    sliceDecoder
	u16sdec  sliceDecoder
	u32sdec  sliceDecoder
	u64sdec  sliceDecoder
	upsdec   sliceDecoder
	f32sdec  sliceDecoder
	f64sdec  sliceDecoder
	c64sdec  sliceDecoder
	c128sdec sliceDecoder
	ifsdec   sliceDecoder
	ssdec    sliceDecoder
	bisdec   sliceDecoder
	bfsdec   sliceDecoder
	brsdec   sliceDecoder
)

func init() {
	bsdec = boolSliceDecoder(reflect.TypeOf(([]bool)(nil))).(sliceDecoder)
	isdec = intSliceDecoder(reflect.TypeOf(([]int)(nil))).(sliceDecoder)
	i8sdec = int8SliceDecoder(reflect.TypeOf(([]int8)(nil))).(sliceDecoder)
	i16sdec = int16SliceDecoder(reflect.TypeOf(([]int16)(nil))).(sliceDecoder)
	i32sdec = int32SliceDecoder(reflect.TypeOf(([]int32)(nil))).(sliceDecoder)
	i64sdec = int64SliceDecoder(reflect.TypeOf(([]int64)(nil))).(sliceDecoder)
	usdec = uintSliceDecoder(reflect.TypeOf(([]uint)(nil))).(sliceDecoder)
	u16sdec = uint16SliceDecoder(reflect.TypeOf(([]uint16)(nil))).(sliceDecoder)
	u32sdec = uint32SliceDecoder(reflect.TypeOf(([]uint32)(nil))).(sliceDecoder)
	u64sdec = uint64SliceDecoder(reflect.TypeOf(([]uint64)(nil))).(sliceDecoder)
	upsdec = uintSliceDecoder(reflect.TypeOf(([]uintptr)(nil))).(sliceDecoder)
	f32sdec = float32SliceDecoder(reflect.TypeOf(([]float32)(nil))).(sliceDecoder)
	f64sdec = float64SliceDecoder(reflect.TypeOf(([]float64)(nil))).(sliceDecoder)
	c64sdec = complex64SliceDecoder(reflect.TypeOf(([]complex64)(nil))).(sliceDecoder)
	c128sdec = complex128SliceDecoder(reflect.TypeOf(([]complex128)(nil))).(sliceDecoder)
	ifsdec = interfaceSliceDecoder(reflect.TypeOf(([]interface{})(nil))).(sliceDecoder)
	ssdec = stringSliceDecoder(reflect.TypeOf(([]string)(nil))).(sliceDecoder)
	bisdec = bigIntSliceDecoder(reflect.TypeOf(([]*big.Int)(nil))).(sliceDecoder)
	bfsdec = bigFloatSliceDecoder(reflect.TypeOf(([]*big.Float)(nil))).(sliceDecoder)
	brsdec = bigRatSliceDecoder(reflect.TypeOf(([]*big.Rat)(nil))).(sliceDecoder)

	RegisterValueDecoder(bsdec)
	RegisterValueDecoder(isdec)
	RegisterValueDecoder(i8sdec)
	RegisterValueDecoder(i16sdec)
	RegisterValueDecoder(i32sdec)
	RegisterValueDecoder(i64sdec)
	RegisterValueDecoder(usdec)
	RegisterValueDecoder(u16sdec)
	RegisterValueDecoder(u32sdec)
	RegisterValueDecoder(u64sdec)
	RegisterValueDecoder(upsdec)
	RegisterValueDecoder(f32sdec)
	RegisterValueDecoder(f64sdec)
	RegisterValueDecoder(c64sdec)
	RegisterValueDecoder(c128sdec)
	RegisterValueDecoder(ifsdec)
	RegisterValueDecoder(ssdec)
	RegisterValueDecoder(bisdec)
	RegisterValueDecoder(bfsdec)
	RegisterValueDecoder(brsdec)
}
