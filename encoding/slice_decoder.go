/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/slice_decoder.go                                |
|                                                          |
| LastModified: Jun 14, 2020                               |
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
	t          *reflect2.UnsafeSliceType
	et         reflect.Type
	empty      unsafe.Pointer
	decodeElem DecodeHandler
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
			valdec.decodeElem(dec, valdec.et, valdec.t.UnsafeGetIndex(slice, i))
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
func SliceDecoder(t reflect.Type, decodeElem DecodeHandler) ValueDecoder {
	return sliceDecoder{
		reflect2.Type2(t).(*reflect2.UnsafeSliceType),
		t.Elem(),
		reflect2.Type2(reflect.ArrayOf(0, t.Elem())).UnsafeNew(),
		decodeElem,
	}
}

func boolSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, boolDecode)
}

func intSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, intDecode)
}

func int8SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, int8Decode)
}

func int16SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, int16Decode)
}

func int32SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, int32Decode)
}

func int64SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, int64Decode)
}

func uintSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, uintDecode)
}

func uint8SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, uint8Decode)
}

func uint16SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, uint16Decode)
}

func uint32SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, uint32Decode)
}

func uint64SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, uint64Decode)
}

func uintptrSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, uintptrDecode)
}

func float32SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, float32Decode)
}

func float64SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, float64Decode)
}

func complex64SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, complex64Decode)
}

func complex128SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, complex128Decode)
}

func interfaceSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, interfaceDecode)
}

func bytesSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, bytesDecode)
}

func stringSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, stringDecode)
}

func bigIntSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, bigIntDecode)
}

func bigFloatSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, bigFloatDecode)
}

func bigRatSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, bigRatDecode)
}

func otherSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, otherDecode(t))
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
	u8ssdec  sliceDecoder
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
	u8ssdec = bytesSliceDecoder(reflect.TypeOf(([][]byte)(nil))).(sliceDecoder)
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
	RegisterValueDecoder(u8ssdec)
	RegisterValueDecoder(ssdec)
	RegisterValueDecoder(bisdec)
	RegisterValueDecoder(bfsdec)
	RegisterValueDecoder(brsdec)
}
