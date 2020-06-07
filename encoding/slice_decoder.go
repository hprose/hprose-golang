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

	"github.com/modern-go/reflect2"
)

// SliceDecoder is the implementation of ValueDecoder for []T.
type SliceDecoder struct {
	destType        reflect.Type
	nilSlice        func(p interface{})
	emptySlice      func(p interface{})
	readListAsSlice func(dec *Decoder, p interface{})
}

// Decode a slice from Decoder
func (valdec SliceDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch tag {
	case TagNull:
		valdec.nilSlice(p)
	case TagEmpty:
		valdec.emptySlice(p)
	case TagList:
		valdec.readListAsSlice(dec, p)
	default:
		dec.decodeError(valdec.destType, tag)
	}
}

func (dec *Decoder) readListAsIntSlice() []int {
	count := dec.ReadInt()
	slice := make([]int, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = intdec.decode(dec, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readListAsInt8Slice() []int8 {
	count := dec.ReadInt()
	slice := make([]int8, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = int8dec.decode(dec, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readListAsInt16Slice() []int16 {
	count := dec.ReadInt()
	slice := make([]int16, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = int16dec.decode(dec, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readListAsInt32Slice() []int32 {
	count := dec.ReadInt()
	slice := make([]int32, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = int32dec.decode(dec, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readListAsInt64Slice() []int64 {
	count := dec.ReadInt()
	slice := make([]int64, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = int64dec.decode(dec, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readListAsUintSlice() []uint {
	count := dec.ReadInt()
	slice := make([]uint, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = uintdec.decode(dec, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readListAsUint16Slice() []uint16 {
	count := dec.ReadInt()
	slice := make([]uint16, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = uint16dec.decode(dec, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readListAsUint32Slice() []uint32 {
	count := dec.ReadInt()
	slice := make([]uint32, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = uint32dec.decode(dec, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readListAsUint64Slice() []uint64 {
	count := dec.ReadInt()
	slice := make([]uint64, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = uint64dec.decode(dec, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readListAsUintptrSlice() []uintptr {
	count := dec.ReadInt()
	slice := make([]uintptr, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = uptrdec.decode(dec, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readListAsBoolSlice() []bool {
	count := dec.ReadInt()
	slice := make([]bool, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = booldec.decode(dec, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readListAsFloat32Slice() []float32 {
	count := dec.ReadInt()
	slice := make([]float32, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = f32dec.decode(dec, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readListAsFloat64Slice() []float64 {
	count := dec.ReadInt()
	slice := make([]float64, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = f64dec.decode(dec, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readListAsComplex64Slice() []complex64 {
	count := dec.ReadInt()
	slice := make([]complex64, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = c64dec.decode(dec, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readListAsComplex128Slice() []complex128 {
	count := dec.ReadInt()
	slice := make([]complex128, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = c128dec.decode(dec, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readListAsStringSlice() []string {
	count := dec.ReadInt()
	slice := make([]string, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = strdec.decode(dec, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readListAsInterfaceSlice() []interface{} {
	count := dec.ReadInt()
	slice := make([]interface{}, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = ifacedec.decode(dec, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readListAsBigIntSlice() []*big.Int {
	count := dec.ReadInt()
	slice := make([]*big.Int, count, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = bigintdec.decode(dec, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readListAsBigFloatSlice() []*big.Float {
	count := dec.ReadInt()
	slice := make([]*big.Float, count, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = bigfloatdec.decode(dec, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readListAsBigRatSlice() []*big.Rat {
	count := dec.ReadInt()
	slice := make([]*big.Rat, count, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = bigratdec.decode(dec, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) readListAsSlice(t *reflect2.UnsafeSliceType) interface{} {
	count := dec.ReadInt()
	slice := t.MakeSlice(count, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		dec.Decode(t.GetIndex(slice, i))
	}
	dec.Skip()
	return slice
}

var (
	intsdec      SliceDecoder
	int8sdec     SliceDecoder
	int16sdec    SliceDecoder
	int32sdec    SliceDecoder
	int64sdec    SliceDecoder
	uintsdec     SliceDecoder
	uint16sdec   SliceDecoder
	uint32sdec   SliceDecoder
	uint64sdec   SliceDecoder
	uptrsdec     SliceDecoder
	boolsdec     SliceDecoder
	f32sdec      SliceDecoder
	f64sdec      SliceDecoder
	c64sdec      SliceDecoder
	c128sdec     SliceDecoder
	strsdec      SliceDecoder
	ifacesdec    SliceDecoder
	bigintsdec   SliceDecoder
	bigfloatsdec SliceDecoder
	bigratsdec   SliceDecoder
)

func init() {
	intsdec = SliceDecoder{
		reflect.TypeOf(([]int)(nil)),
		func(p interface{}) { *(p.(*[]int)) = nil },
		func(p interface{}) { *(p.(*[]int)) = []int{} },
		func(dec *Decoder, p interface{}) { *(p.(*[]int)) = dec.readListAsIntSlice() },
	}
	int8sdec = SliceDecoder{
		reflect.TypeOf(([]int8)(nil)),
		func(p interface{}) { *(p.(*[]int8)) = nil },
		func(p interface{}) { *(p.(*[]int8)) = []int8{} },
		func(dec *Decoder, p interface{}) { *(p.(*[]int8)) = dec.readListAsInt8Slice() },
	}
	int16sdec = SliceDecoder{
		reflect.TypeOf(([]int16)(nil)),
		func(p interface{}) { *(p.(*[]int16)) = nil },
		func(p interface{}) { *(p.(*[]int16)) = []int16{} },
		func(dec *Decoder, p interface{}) { *(p.(*[]int16)) = dec.readListAsInt16Slice() },
	}
	int32sdec = SliceDecoder{
		reflect.TypeOf(([]int32)(nil)),
		func(p interface{}) { *(p.(*[]int32)) = nil },
		func(p interface{}) { *(p.(*[]int32)) = []int32{} },
		func(dec *Decoder, p interface{}) { *(p.(*[]int32)) = dec.readListAsInt32Slice() },
	}
	int64sdec = SliceDecoder{
		reflect.TypeOf(([]int64)(nil)),
		func(p interface{}) { *(p.(*[]int64)) = nil },
		func(p interface{}) { *(p.(*[]int64)) = []int64{} },
		func(dec *Decoder, p interface{}) { *(p.(*[]int64)) = dec.readListAsInt64Slice() },
	}
	uintsdec = SliceDecoder{
		reflect.TypeOf(([]uint)(nil)),
		func(p interface{}) { *(p.(*[]uint)) = nil },
		func(p interface{}) { *(p.(*[]uint)) = []uint{} },
		func(dec *Decoder, p interface{}) { *(p.(*[]uint)) = dec.readListAsUintSlice() },
	}
	uint16sdec = SliceDecoder{
		reflect.TypeOf(([]uint16)(nil)),
		func(p interface{}) { *(p.(*[]uint16)) = nil },
		func(p interface{}) { *(p.(*[]uint16)) = []uint16{} },
		func(dec *Decoder, p interface{}) { *(p.(*[]uint16)) = dec.readListAsUint16Slice() },
	}
	uint32sdec = SliceDecoder{
		reflect.TypeOf(([]uint32)(nil)),
		func(p interface{}) { *(p.(*[]uint32)) = nil },
		func(p interface{}) { *(p.(*[]uint32)) = []uint32{} },
		func(dec *Decoder, p interface{}) { *(p.(*[]uint32)) = dec.readListAsUint32Slice() },
	}
	uint64sdec = SliceDecoder{
		reflect.TypeOf(([]uint64)(nil)),
		func(p interface{}) { *(p.(*[]uint64)) = nil },
		func(p interface{}) { *(p.(*[]uint64)) = []uint64{} },
		func(dec *Decoder, p interface{}) { *(p.(*[]uint64)) = dec.readListAsUint64Slice() },
	}
	uptrsdec = SliceDecoder{
		reflect.TypeOf(([]uintptr)(nil)),
		func(p interface{}) { *(p.(*[]uintptr)) = nil },
		func(p interface{}) { *(p.(*[]uintptr)) = []uintptr{} },
		func(dec *Decoder, p interface{}) { *(p.(*[]uintptr)) = dec.readListAsUintptrSlice() },
	}
	boolsdec = SliceDecoder{
		reflect.TypeOf(([]bool)(nil)),
		func(p interface{}) { *(p.(*[]bool)) = nil },
		func(p interface{}) { *(p.(*[]bool)) = []bool{} },
		func(dec *Decoder, p interface{}) { *(p.(*[]bool)) = dec.readListAsBoolSlice() },
	}
	f32sdec = SliceDecoder{
		reflect.TypeOf(([]float32)(nil)),
		func(p interface{}) { *(p.(*[]float32)) = nil },
		func(p interface{}) { *(p.(*[]float32)) = []float32{} },
		func(dec *Decoder, p interface{}) { *(p.(*[]float32)) = dec.readListAsFloat32Slice() },
	}
	f64sdec = SliceDecoder{
		reflect.TypeOf(([]float64)(nil)),
		func(p interface{}) { *(p.(*[]float64)) = nil },
		func(p interface{}) { *(p.(*[]float64)) = []float64{} },
		func(dec *Decoder, p interface{}) { *(p.(*[]float64)) = dec.readListAsFloat64Slice() },
	}
	c64sdec = SliceDecoder{
		reflect.TypeOf(([]complex64)(nil)),
		func(p interface{}) { *(p.(*[]complex64)) = nil },
		func(p interface{}) { *(p.(*[]complex64)) = []complex64{} },
		func(dec *Decoder, p interface{}) { *(p.(*[]complex64)) = dec.readListAsComplex64Slice() },
	}
	c128sdec = SliceDecoder{
		reflect.TypeOf(([]complex128)(nil)),
		func(p interface{}) { *(p.(*[]complex128)) = nil },
		func(p interface{}) { *(p.(*[]complex128)) = []complex128{} },
		func(dec *Decoder, p interface{}) { *(p.(*[]complex128)) = dec.readListAsComplex128Slice() },
	}
	strsdec = SliceDecoder{
		reflect.TypeOf(([]string)(nil)),
		func(p interface{}) { *(p.(*[]string)) = nil },
		func(p interface{}) { *(p.(*[]string)) = []string{} },
		func(dec *Decoder, p interface{}) { *(p.(*[]string)) = dec.readListAsStringSlice() },
	}
	ifacesdec = SliceDecoder{
		reflect.TypeOf(([]interface{})(nil)),
		func(p interface{}) { *(p.(*[]interface{})) = nil },
		func(p interface{}) { *(p.(*[]interface{})) = []interface{}{} },
		func(dec *Decoder, p interface{}) { *(p.(*[]interface{})) = dec.readListAsInterfaceSlice() },
	}
	bigintsdec = SliceDecoder{
		reflect.TypeOf(([]*big.Int)(nil)),
		func(p interface{}) { *(p.(*[]*big.Int)) = nil },
		func(p interface{}) { *(p.(*[]*big.Int)) = []*big.Int{} },
		func(dec *Decoder, p interface{}) { *(p.(*[]*big.Int)) = dec.readListAsBigIntSlice() },
	}
	bigfloatsdec = SliceDecoder{
		reflect.TypeOf(([]*big.Float)(nil)),
		func(p interface{}) { *(p.(*[]*big.Float)) = nil },
		func(p interface{}) { *(p.(*[]*big.Float)) = []*big.Float{} },
		func(dec *Decoder, p interface{}) { *(p.(*[]*big.Float)) = dec.readListAsBigFloatSlice() },
	}
	bigratsdec = SliceDecoder{
		reflect.TypeOf(([]*big.Rat)(nil)),
		func(p interface{}) { *(p.(*[]*big.Rat)) = nil },
		func(p interface{}) { *(p.(*[]*big.Rat)) = []*big.Rat{} },
		func(dec *Decoder, p interface{}) { *(p.(*[]*big.Rat)) = dec.readListAsBigRatSlice() },
	}
}
