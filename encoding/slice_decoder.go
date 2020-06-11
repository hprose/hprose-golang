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

func (dec *Decoder) readBoolSlice(et reflect.Type) []bool {
	count := dec.ReadInt()
	slice := make([]bool, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = dec.decodeBool(et, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readIntSlice(et reflect.Type) []int {
	count := dec.ReadInt()
	slice := make([]int, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = dec.decodeInt(et, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readInt8Slice(et reflect.Type) []int8 {
	count := dec.ReadInt()
	slice := make([]int8, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = dec.decodeInt8(et, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readInt16Slice(et reflect.Type) []int16 {
	count := dec.ReadInt()
	slice := make([]int16, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = dec.decodeInt16(et, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readInt32Slice(et reflect.Type) []int32 {
	count := dec.ReadInt()
	slice := make([]int32, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = dec.decodeInt32(et, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readInt64Slice(et reflect.Type) []int64 {
	count := dec.ReadInt()
	slice := make([]int64, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = dec.decodeInt64(et, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readUintSlice(et reflect.Type) []uint {
	count := dec.ReadInt()
	slice := make([]uint, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = dec.decodeUint(et, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readUint16Slice(et reflect.Type) []uint16 {
	count := dec.ReadInt()
	slice := make([]uint16, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = dec.decodeUint16(et, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readUint32Slice(et reflect.Type) []uint32 {
	count := dec.ReadInt()
	slice := make([]uint32, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = dec.decodeUint32(et, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readUint64Slice(et reflect.Type) []uint64 {
	count := dec.ReadInt()
	slice := make([]uint64, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = dec.decodeUint64(et, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readUintptrSlice(et reflect.Type) []uintptr {
	count := dec.ReadInt()
	slice := make([]uintptr, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = dec.decodeUintptr(et, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readFloat32Slice(et reflect.Type) []float32 {
	count := dec.ReadInt()
	slice := make([]float32, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = dec.decodeFloat32(et, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readFloat64Slice(et reflect.Type) []float64 {
	count := dec.ReadInt()
	slice := make([]float64, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = dec.decodeFloat64(et, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readComplex64Slice(et reflect.Type) []complex64 {
	count := dec.ReadInt()
	slice := make([]complex64, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = dec.decodeComplex64(et, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readComplex128Slice(et reflect.Type) []complex128 {
	count := dec.ReadInt()
	slice := make([]complex128, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = dec.decodeComplex128(et, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readInterfaceSlice() []interface{} {
	count := dec.ReadInt()
	slice := make([]interface{}, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = dec.decodeInterface(dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readStringSlice(et reflect.Type) []string {
	count := dec.ReadInt()
	slice := make([]string, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = dec.decodeString(et, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readBigIntSlice(et reflect.Type) []*big.Int {
	count := dec.ReadInt()
	slice := make([]*big.Int, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = dec.decodeBigInt(et, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readBigFloatSlice(et reflect.Type) []*big.Float {
	count := dec.ReadInt()
	slice := make([]*big.Float, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = dec.decodeBigFloat(et, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readBigRatSlice(et reflect.Type) []*big.Rat {
	count := dec.ReadInt()
	slice := make([]*big.Rat, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = dec.decodeBigRat(et, dec.NextByte())
	}
	dec.Skip()
	return slice
}

// sliceDecoder is the implementation of ValueDecoder for []T.
type sliceDecoder struct {
	t         *reflect2.UnsafeSliceType
	et        reflect.Type
	empty     unsafe.Pointer
	readSlice func(dec *Decoder, p interface{}, et reflect.Type)
}

func (valdec sliceDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch tag {
	case TagNull:
		valdec.t.UnsafeSetNil(reflect2.PtrOf(p))
	case TagEmpty:
		setSliceHeader(reflect2.PtrOf(p), valdec.empty, 0)
	case TagList:
		valdec.readSlice(dec, p, valdec.et)
	default:
		dec.decodeError(valdec.t.Type1(), tag)
	}
}

func (valdec sliceDecoder) Type() reflect.Type {
	return valdec.t.Type1()
}

// SliceDecoder returns a ValueDecoder for []T.
func SliceDecoder(t reflect.Type, readSlice func(dec *Decoder, p interface{}, et reflect.Type)) ValueDecoder {
	return sliceDecoder{
		reflect2.Type2(t).(*reflect2.UnsafeSliceType),
		t.Elem(),
		reflect2.Type2(reflect.ArrayOf(0, t.Elem())).UnsafeNew(),
		readSlice,
	}
}

func boolSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, p interface{}, et reflect.Type) {
		*(*[]bool)(reflect2.PtrOf(p)) = dec.readBoolSlice(et)
	})
}

func intSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, p interface{}, et reflect.Type) {
		*(*[]int)(reflect2.PtrOf(p)) = dec.readIntSlice(et)
	})
}

func int8SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, p interface{}, et reflect.Type) {
		*(*[]int8)(reflect2.PtrOf(p)) = dec.readInt8Slice(et)
	})
}

func int16SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, p interface{}, et reflect.Type) {
		*(*[]int16)(reflect2.PtrOf(p)) = dec.readInt16Slice(et)
	})
}

func int32SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, p interface{}, et reflect.Type) {
		*(*[]int32)(reflect2.PtrOf(p)) = dec.readInt32Slice(et)
	})
}

func int64SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, p interface{}, et reflect.Type) {
		*(*[]int64)(reflect2.PtrOf(p)) = dec.readInt64Slice(et)
	})
}

func uintSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, p interface{}, et reflect.Type) {
		*(*[]uint)(reflect2.PtrOf(p)) = dec.readUintSlice(et)
	})
}

func uint16SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, p interface{}, et reflect.Type) {
		*(*[]uint16)(reflect2.PtrOf(p)) = dec.readUint16Slice(et)
	})
}

func uint32SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, p interface{}, et reflect.Type) {
		*(*[]uint32)(reflect2.PtrOf(p)) = dec.readUint32Slice(et)
	})
}

func uint64SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, p interface{}, et reflect.Type) {
		*(*[]uint64)(reflect2.PtrOf(p)) = dec.readUint64Slice(et)
	})
}

func uintptrSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, p interface{}, et reflect.Type) {
		*(*[]uintptr)(reflect2.PtrOf(p)) = dec.readUintptrSlice(et)
	})
}

func float32SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, p interface{}, et reflect.Type) {
		*(*[]float32)(reflect2.PtrOf(p)) = dec.readFloat32Slice(et)
	})
}

func float64SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, p interface{}, et reflect.Type) {
		*(*[]float64)(reflect2.PtrOf(p)) = dec.readFloat64Slice(et)
	})
}

func complex64SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, p interface{}, et reflect.Type) {
		*(*[]complex64)(reflect2.PtrOf(p)) = dec.readComplex64Slice(et)
	})
}

func complex128SliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, p interface{}, et reflect.Type) {
		*(*[]complex128)(reflect2.PtrOf(p)) = dec.readComplex128Slice(et)
	})
}

func interfaceSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, p interface{}, et reflect.Type) {
		*(*[]interface{})(reflect2.PtrOf(p)) = dec.readInterfaceSlice()
	})
}

func stringSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, p interface{}, et reflect.Type) {
		*(*[]string)(reflect2.PtrOf(p)) = dec.readStringSlice(et)
	})
}

func bigIntSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, p interface{}, et reflect.Type) {
		*(*[]*big.Int)(reflect2.PtrOf(p)) = dec.readBigIntSlice(et)
	})
}

func bigFloatSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, p interface{}, et reflect.Type) {
		*(*[]*big.Float)(reflect2.PtrOf(p)) = dec.readBigFloatSlice(et)
	})
}

func bigRatSliceDecoder(t reflect.Type) ValueDecoder {
	return SliceDecoder(t, func(dec *Decoder, p interface{}, et reflect.Type) {
		*(*[]*big.Rat)(reflect2.PtrOf(p)) = dec.readBigRatSlice(et)
	})
}

func otherSliceDecoder(t reflect.Type) ValueDecoder {
	t2 := reflect2.Type2(t).(*reflect2.UnsafeSliceType)
	valdec := getValueDecoder(t.Elem())
	return SliceDecoder(t, func(dec *Decoder, p interface{}, et reflect.Type) {
		count := dec.ReadInt()
		slice := t2.MakeSlice(count, count)
		dec.AddReference(slice)
		for i := 0; i < count; i++ {
			valdec.Decode(dec, t2.GetIndex(slice, i), dec.NextByte())
		}
		dec.Skip()
		t2.Set(p, slice)
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
