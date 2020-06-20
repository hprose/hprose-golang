/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/slice_decoder.go                                |
|                                                          |
| LastModified: Jun 20, 2020                               |
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

func getSliceDecoder(t reflect.Type) ValueDecoder {
	et := t.Elem()
	if et.Kind() == reflect.Uint8 {
		return bytesDecoder{t}
	}
	return SliceDecoder(t, getDecodeHandler(et))
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
	bsdec = SliceDecoder(reflect.TypeOf(([]bool)(nil)), boolDecode).(sliceDecoder)
	isdec = SliceDecoder(reflect.TypeOf(([]int)(nil)), intDecode).(sliceDecoder)
	i8sdec = SliceDecoder(reflect.TypeOf(([]int8)(nil)), int8Decode).(sliceDecoder)
	i16sdec = SliceDecoder(reflect.TypeOf(([]int16)(nil)), int16Decode).(sliceDecoder)
	i32sdec = SliceDecoder(reflect.TypeOf(([]int32)(nil)), int32Decode).(sliceDecoder)
	i64sdec = SliceDecoder(reflect.TypeOf(([]int64)(nil)), int64Decode).(sliceDecoder)
	usdec = SliceDecoder(reflect.TypeOf(([]uint)(nil)), uintDecode).(sliceDecoder)
	u16sdec = SliceDecoder(reflect.TypeOf(([]uint16)(nil)), uint16Decode).(sliceDecoder)
	u32sdec = SliceDecoder(reflect.TypeOf(([]uint32)(nil)), uint32Decode).(sliceDecoder)
	u64sdec = SliceDecoder(reflect.TypeOf(([]uint64)(nil)), uint64Decode).(sliceDecoder)
	upsdec = SliceDecoder(reflect.TypeOf(([]uintptr)(nil)), uintDecode).(sliceDecoder)
	f32sdec = SliceDecoder(reflect.TypeOf(([]float32)(nil)), float32Decode).(sliceDecoder)
	f64sdec = SliceDecoder(reflect.TypeOf(([]float64)(nil)), float64Decode).(sliceDecoder)
	c64sdec = SliceDecoder(reflect.TypeOf(([]complex64)(nil)), complex64Decode).(sliceDecoder)
	c128sdec = SliceDecoder(reflect.TypeOf(([]complex128)(nil)), complex128Decode).(sliceDecoder)
	ifsdec = SliceDecoder(reflect.TypeOf(([]interface{})(nil)), interfaceDecode).(sliceDecoder)
	u8ssdec = SliceDecoder(reflect.TypeOf(([][]byte)(nil)), bytesDecode).(sliceDecoder)
	ssdec = SliceDecoder(reflect.TypeOf(([]string)(nil)), stringDecode).(sliceDecoder)
	bisdec = SliceDecoder(reflect.TypeOf(([]*big.Int)(nil)), bigIntDecode).(sliceDecoder)
	bfsdec = SliceDecoder(reflect.TypeOf(([]*big.Float)(nil)), bigFloatDecode).(sliceDecoder)
	brsdec = SliceDecoder(reflect.TypeOf(([]*big.Rat)(nil)), bigRatDecode).(sliceDecoder)

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
