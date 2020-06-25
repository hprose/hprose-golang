/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/slice_decoder.go                                |
|                                                          |
| LastModified: Jun 25, 2020                               |
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

// makeSliceDecoder returns a ValueDecoder for []T.
func makeSliceDecoder(t reflect.Type, decodeElem DecodeHandler) sliceDecoder {
	return sliceDecoder{
		reflect2.Type2(t).(*reflect2.UnsafeSliceType),
		t.Elem(),
		reflect2.Type2(reflect.ArrayOf(0, t.Elem())).UnsafeNew(),
		decodeElem,
	}
}

func getSliceDecoder(t reflect.Type) ValueDecoder {
	et := t.Elem()
	if et.Kind() == reflect.Uint8 {
		return bytesDecoder{t}
	}
	return makeSliceDecoder(t, GetDecodeHandler(et))
}

func (dec *Decoder) fastDecodeSlice(p interface{}, tag byte) bool {
	switch p.(type) {
	case *[]bool:
		bsdec.Decode(dec, p, tag)
	case *[]int:
		isdec.Decode(dec, p, tag)
	case *[]int8:
		i8sdec.Decode(dec, p, tag)
	case *[]int16:
		i16sdec.Decode(dec, p, tag)
	case *[]int32:
		i32sdec.Decode(dec, p, tag)
	case *[]int64:
		i64sdec.Decode(dec, p, tag)
	case *[]uint:
		usdec.Decode(dec, p, tag)
	case *[]uint16:
		u16sdec.Decode(dec, p, tag)
	case *[]uint32:
		u32sdec.Decode(dec, p, tag)
	case *[]uint64:
		u64sdec.Decode(dec, p, tag)
	case *[]uintptr:
		upsdec.Decode(dec, p, tag)
	case *[]float32:
		f32sdec.Decode(dec, p, tag)
	case *[]float64:
		f64sdec.Decode(dec, p, tag)
	case *[]complex64:
		c64sdec.Decode(dec, p, tag)
	case *[]complex128:
		c128sdec.Decode(dec, p, tag)
	case *[]interface{}:
		ifsdec.Decode(dec, p, tag)
	case *[][]byte:
		u8ssdec.Decode(dec, p, tag)
	case *[]string:
		ssdec.Decode(dec, p, tag)
	case *[]*big.Int:
		bisdec.Decode(dec, p, tag)
	case *[]*big.Float:
		bfsdec.Decode(dec, p, tag)
	case *[]*big.Rat:
		brsdec.Decode(dec, p, tag)
	default:
		return false
	}
	return true
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
	bsdec = makeSliceDecoder(reflect.TypeOf(([]bool)(nil)), boolDecode)
	isdec = makeSliceDecoder(reflect.TypeOf(([]int)(nil)), intDecode)
	i8sdec = makeSliceDecoder(reflect.TypeOf(([]int8)(nil)), int8Decode)
	i16sdec = makeSliceDecoder(reflect.TypeOf(([]int16)(nil)), int16Decode)
	i32sdec = makeSliceDecoder(reflect.TypeOf(([]int32)(nil)), int32Decode)
	i64sdec = makeSliceDecoder(reflect.TypeOf(([]int64)(nil)), int64Decode)
	usdec = makeSliceDecoder(reflect.TypeOf(([]uint)(nil)), uintDecode)
	u16sdec = makeSliceDecoder(reflect.TypeOf(([]uint16)(nil)), uint16Decode)
	u32sdec = makeSliceDecoder(reflect.TypeOf(([]uint32)(nil)), uint32Decode)
	u64sdec = makeSliceDecoder(reflect.TypeOf(([]uint64)(nil)), uint64Decode)
	upsdec = makeSliceDecoder(reflect.TypeOf(([]uintptr)(nil)), uintDecode)
	f32sdec = makeSliceDecoder(reflect.TypeOf(([]float32)(nil)), float32Decode)
	f64sdec = makeSliceDecoder(reflect.TypeOf(([]float64)(nil)), float64Decode)
	c64sdec = makeSliceDecoder(reflect.TypeOf(([]complex64)(nil)), complex64Decode)
	c128sdec = makeSliceDecoder(reflect.TypeOf(([]complex128)(nil)), complex128Decode)
	ifsdec = makeSliceDecoder(reflect.TypeOf(([]interface{})(nil)), interfaceDecode)
	u8ssdec = makeSliceDecoder(reflect.TypeOf(([][]byte)(nil)), bytesDecode)
	ssdec = makeSliceDecoder(reflect.TypeOf(([]string)(nil)), stringDecode)
	bisdec = makeSliceDecoder(reflect.TypeOf(([]*big.Int)(nil)), bigIntDecode)
	bfsdec = makeSliceDecoder(reflect.TypeOf(([]*big.Float)(nil)), bigFloatDecode)
	brsdec = makeSliceDecoder(reflect.TypeOf(([]*big.Rat)(nil)), bigRatDecode)

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
