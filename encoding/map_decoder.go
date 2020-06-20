/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/map_decoder.go                                  |
|                                                          |
| LastModified: Jun 20, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"reflect"
	"strconv"
	"unsafe"

	"github.com/modern-go/reflect2"
)

// mapDecoder is the implementation of ValueDecoder for map[K]V.
type mapDecoder struct {
	t           *reflect2.UnsafeMapType
	kt          reflect.Type
	kt2         reflect2.Type
	vt          reflect.Type
	vt2         reflect2.Type
	decodeKey   DecodeHandler
	decodeValue DecodeHandler
}

func (valdec mapDecoder) convertKey(i int, p unsafe.Pointer) {
	switch valdec.kt.Kind() {
	case reflect.Int:
		*(*int)(p) = i
	case reflect.Int8:
		*(*int8)(p) = int8(i)
	case reflect.Int16:
		*(*int16)(p) = int16(i)
	case reflect.Int32:
		*(*int32)(p) = int32(i)
	case reflect.Int64:
		*(*int64)(p) = int64(i)
	case reflect.Uint:
		*(*uint)(p) = uint(i)
	case reflect.Uint8:
		*(*uint8)(p) = uint8(i)
	case reflect.Uint16:
		*(*uint16)(p) = uint16(i)
	case reflect.Uint32:
		*(*uint32)(p) = uint32(i)
	case reflect.Uint64:
		*(*uint64)(p) = uint64(i)
	case reflect.Uintptr:
		*(*uintptr)(p) = uintptr(i)
	case reflect.Float32:
		*(*float32)(p) = float32(i)
	case reflect.Float64:
		*(*float64)(p) = float64(i)
	case reflect.Complex64:
		*(*complex64)(p) = complex(float32(i), 0)
	case reflect.Complex128:
		*(*complex128)(p) = complex(float64(i), 0)
	case reflect.Interface:
		*(*interface{})(p) = i
	case reflect.String:
		*(*string)(p) = strconv.Itoa(i)
	}
}

func (valdec mapDecoder) canDecodeListAsMap() bool {
	switch valdec.kt.Kind() {
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32,
		reflect.Float64,
		reflect.Complex64,
		reflect.Complex128,
		reflect.Interface,
		reflect.String:
		return true
	}
	return false
}

func (valdec mapDecoder) decodeListAsMap(dec *Decoder, p interface{}, tag byte) {
	if !valdec.canDecodeListAsMap() {
		dec.decodeError(valdec.t.Type1(), tag)
		return
	}
	mp := reflect2.PtrOf(p)
	count := dec.ReadInt()
	valdec.t.UnsafeSet(mp, valdec.t.UnsafeMakeMap(count))
	dec.AddReference(p)
	kp := valdec.kt2.UnsafeNew()
	vp := valdec.vt2.UnsafeNew()
	for i := 0; i < count; i++ {
		valdec.convertKey(i, kp)
		valdec.decodeValue(dec, valdec.vt, vp)
		valdec.t.UnsafeSetIndex(mp, kp, vp)
	}
	dec.Skip()
}

func (valdec mapDecoder) decodeMap(dec *Decoder, p interface{}) {
	mp := reflect2.PtrOf(p)
	count := dec.ReadInt()
	valdec.t.UnsafeSet(mp, valdec.t.UnsafeMakeMap(count))
	dec.AddReference(p)
	kp := valdec.kt2.UnsafeNew()
	vp := valdec.vt2.UnsafeNew()
	for i := 0; i < count; i++ {
		valdec.decodeKey(dec, valdec.kt, kp)
		valdec.decodeValue(dec, valdec.vt, vp)
		valdec.t.UnsafeSetIndex(mp, kp, vp)
	}
	dec.Skip()
}

func (valdec mapDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch tag {
	case TagNull:
		mp := reflect2.PtrOf(p)
		if !valdec.t.UnsafeIsNil(mp) {
			*(*unsafe.Pointer)(mp) = nil
		}
	case TagEmpty:
		valdec.t.UnsafeSet(reflect2.PtrOf(p), valdec.t.UnsafeMakeMap(0))
	case TagList:
		valdec.decodeListAsMap(dec, p, tag)
	case TagMap:
		valdec.decodeMap(dec, p)
	default:
		dec.decodeError(valdec.t.Type1(), tag)
	}
}

func (valdec mapDecoder) Type() reflect.Type {
	return valdec.t.Type1()
}

// MapDecoder returns a ValueDecoder for map[K]V.
func MapDecoder(
	t reflect.Type,
	decodeKey DecodeHandler,
	decodeValue DecodeHandler,
) ValueDecoder {
	at := reflect2.Type2(t).(*reflect2.UnsafeMapType)
	kt := t.Key()
	vt := t.Elem()
	return mapDecoder{
		at,
		kt,
		reflect2.Type2(kt),
		vt,
		reflect2.Type2(vt),
		decodeKey,
		decodeValue,
	}
}

func getMapDecoder(t reflect.Type) ValueDecoder {
	return MapDecoder(t, getDecodeHandler(t.Key()), getDecodeHandler(t.Elem()))
}

func fastDecodeIntMap(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[int]bool:
		ibmdec.Decode(dec, p, tag)
	case *map[int]int:
		iimdec.Decode(dec, p, tag)
	case *map[int]int8:
		ii8mdec.Decode(dec, p, tag)
	case *map[int]int16:
		ii16mdec.Decode(dec, p, tag)
	case *map[int]int32:
		ii32mdec.Decode(dec, p, tag)
	case *map[int]int64:
		ii64mdec.Decode(dec, p, tag)
	case *map[int]uint:
		iumdec.Decode(dec, p, tag)
	case *map[int]uint8:
		iu8mdec.Decode(dec, p, tag)
	case *map[int]uint16:
		iu16mdec.Decode(dec, p, tag)
	case *map[int]uint32:
		iu32mdec.Decode(dec, p, tag)
	case *map[int]uint64:
		iu64mdec.Decode(dec, p, tag)
	case *map[int]float32:
		if32mdec.Decode(dec, p, tag)
	case *map[int]float64:
		if64mdec.Decode(dec, p, tag)
	case *map[int]string:
		ismdec.Decode(dec, p, tag)
	case *map[int]interface{}:
		iifmdec.Decode(dec, p, tag)
	default:
		return false
	}
	return true
}

func fastDecodeInt8Map(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[int8]bool:
		i8bmdec.Decode(dec, p, tag)
	case *map[int8]int:
		i8imdec.Decode(dec, p, tag)
	case *map[int8]int8:
		i8i8mdec.Decode(dec, p, tag)
	case *map[int8]int16:
		i8i16mdec.Decode(dec, p, tag)
	case *map[int8]int32:
		i8i32mdec.Decode(dec, p, tag)
	case *map[int8]int64:
		i8i64mdec.Decode(dec, p, tag)
	case *map[int8]uint:
		i8umdec.Decode(dec, p, tag)
	case *map[int8]uint8:
		i8u8mdec.Decode(dec, p, tag)
	case *map[int8]uint16:
		i8u16mdec.Decode(dec, p, tag)
	case *map[int8]uint32:
		i8u32mdec.Decode(dec, p, tag)
	case *map[int8]uint64:
		i8u64mdec.Decode(dec, p, tag)
	case *map[int8]float32:
		i8f32mdec.Decode(dec, p, tag)
	case *map[int8]float64:
		i8f64mdec.Decode(dec, p, tag)
	case *map[int8]string:
		i8smdec.Decode(dec, p, tag)
	case *map[int8]interface{}:
		i8ifmdec.Decode(dec, p, tag)
	default:
		return false
	}
	return true
}

func fastDecodeInt16Map(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[int16]bool:
		i16bmdec.Decode(dec, p, tag)
	case *map[int16]int:
		i16imdec.Decode(dec, p, tag)
	case *map[int16]int8:
		i16i8mdec.Decode(dec, p, tag)
	case *map[int16]int16:
		i16i16mdec.Decode(dec, p, tag)
	case *map[int16]int32:
		i16i32mdec.Decode(dec, p, tag)
	case *map[int16]int64:
		i16i64mdec.Decode(dec, p, tag)
	case *map[int16]uint:
		i16umdec.Decode(dec, p, tag)
	case *map[int16]uint8:
		i16u8mdec.Decode(dec, p, tag)
	case *map[int16]uint16:
		i16u16mdec.Decode(dec, p, tag)
	case *map[int16]uint32:
		i16u32mdec.Decode(dec, p, tag)
	case *map[int16]uint64:
		i16u64mdec.Decode(dec, p, tag)
	case *map[int16]float32:
		i16f32mdec.Decode(dec, p, tag)
	case *map[int16]float64:
		i16f64mdec.Decode(dec, p, tag)
	case *map[int16]string:
		i16smdec.Decode(dec, p, tag)
	case *map[int16]interface{}:
		i16ifmdec.Decode(dec, p, tag)
	default:
		return false
	}
	return true
}

func fastDecodeInt32Map(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[int32]bool:
		i32bmdec.Decode(dec, p, tag)
	case *map[int32]int:
		i32imdec.Decode(dec, p, tag)
	case *map[int32]int8:
		i32i8mdec.Decode(dec, p, tag)
	case *map[int32]int16:
		i32i16mdec.Decode(dec, p, tag)
	case *map[int32]int32:
		i32i32mdec.Decode(dec, p, tag)
	case *map[int32]int64:
		i32i64mdec.Decode(dec, p, tag)
	case *map[int32]uint:
		i32umdec.Decode(dec, p, tag)
	case *map[int32]uint8:
		i32u8mdec.Decode(dec, p, tag)
	case *map[int32]uint16:
		i32u16mdec.Decode(dec, p, tag)
	case *map[int32]uint32:
		i32u32mdec.Decode(dec, p, tag)
	case *map[int32]uint64:
		i32u64mdec.Decode(dec, p, tag)
	case *map[int32]float32:
		i32f32mdec.Decode(dec, p, tag)
	case *map[int32]float64:
		i32f64mdec.Decode(dec, p, tag)
	case *map[int32]string:
		i32smdec.Decode(dec, p, tag)
	case *map[int32]interface{}:
		i32ifmdec.Decode(dec, p, tag)
	default:
		return false
	}
	return true
}

func fastDecodeInt64Map(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[int64]bool:
		i64bmdec.Decode(dec, p, tag)
	case *map[int64]int:
		i64imdec.Decode(dec, p, tag)
	case *map[int64]int8:
		i64i8mdec.Decode(dec, p, tag)
	case *map[int64]int16:
		i64i16mdec.Decode(dec, p, tag)
	case *map[int64]int32:
		i64i32mdec.Decode(dec, p, tag)
	case *map[int64]int64:
		i64i64mdec.Decode(dec, p, tag)
	case *map[int64]uint:
		i64umdec.Decode(dec, p, tag)
	case *map[int64]uint8:
		i64u8mdec.Decode(dec, p, tag)
	case *map[int64]uint16:
		i64u16mdec.Decode(dec, p, tag)
	case *map[int64]uint32:
		i64u32mdec.Decode(dec, p, tag)
	case *map[int64]uint64:
		i64u64mdec.Decode(dec, p, tag)
	case *map[int64]float32:
		i64f32mdec.Decode(dec, p, tag)
	case *map[int64]float64:
		i64f64mdec.Decode(dec, p, tag)
	case *map[int64]string:
		i64smdec.Decode(dec, p, tag)
	case *map[int64]interface{}:
		i64ifmdec.Decode(dec, p, tag)
	default:
		return false
	}
	return true
}

func fastDecodeUintMap(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[uint]bool:
		ubmdec.Decode(dec, p, tag)
	case *map[uint]int:
		uimdec.Decode(dec, p, tag)
	case *map[uint]int8:
		ui8mdec.Decode(dec, p, tag)
	case *map[uint]int16:
		ui16mdec.Decode(dec, p, tag)
	case *map[uint]int32:
		ui32mdec.Decode(dec, p, tag)
	case *map[uint]int64:
		ui64mdec.Decode(dec, p, tag)
	case *map[uint]uint:
		uumdec.Decode(dec, p, tag)
	case *map[uint]uint8:
		uu8mdec.Decode(dec, p, tag)
	case *map[uint]uint16:
		uu16mdec.Decode(dec, p, tag)
	case *map[uint]uint32:
		uu32mdec.Decode(dec, p, tag)
	case *map[uint]uint64:
		uu64mdec.Decode(dec, p, tag)
	case *map[uint]float32:
		uf32mdec.Decode(dec, p, tag)
	case *map[uint]float64:
		uf64mdec.Decode(dec, p, tag)
	case *map[uint]string:
		usmdec.Decode(dec, p, tag)
	case *map[uint]interface{}:
		uifmdec.Decode(dec, p, tag)
	default:
		return false
	}
	return true
}

func fastDecodeUint8Map(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[uint8]bool:
		u8bmdec.Decode(dec, p, tag)
	case *map[uint8]int:
		u8imdec.Decode(dec, p, tag)
	case *map[uint8]int8:
		u8i8mdec.Decode(dec, p, tag)
	case *map[uint8]int16:
		u8i16mdec.Decode(dec, p, tag)
	case *map[uint8]int32:
		u8i32mdec.Decode(dec, p, tag)
	case *map[uint8]int64:
		u8i64mdec.Decode(dec, p, tag)
	case *map[uint8]uint:
		u8umdec.Decode(dec, p, tag)
	case *map[uint8]uint8:
		u8u8mdec.Decode(dec, p, tag)
	case *map[uint8]uint16:
		u8u16mdec.Decode(dec, p, tag)
	case *map[uint8]uint32:
		u8u32mdec.Decode(dec, p, tag)
	case *map[uint8]uint64:
		u8u64mdec.Decode(dec, p, tag)
	case *map[uint8]float32:
		u8f32mdec.Decode(dec, p, tag)
	case *map[uint8]float64:
		u8f64mdec.Decode(dec, p, tag)
	case *map[uint8]string:
		u8smdec.Decode(dec, p, tag)
	case *map[uint8]interface{}:
		u8ifmdec.Decode(dec, p, tag)
	default:
		return false
	}
	return true
}

func fastDecodeUint16Map(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[uint16]bool:
		u16bmdec.Decode(dec, p, tag)
	case *map[uint16]int:
		u16imdec.Decode(dec, p, tag)
	case *map[uint16]int8:
		u16i8mdec.Decode(dec, p, tag)
	case *map[uint16]int16:
		u16i16mdec.Decode(dec, p, tag)
	case *map[uint16]int32:
		u16i32mdec.Decode(dec, p, tag)
	case *map[uint16]int64:
		u16i64mdec.Decode(dec, p, tag)
	case *map[uint16]uint:
		u16umdec.Decode(dec, p, tag)
	case *map[uint16]uint8:
		u16u8mdec.Decode(dec, p, tag)
	case *map[uint16]uint16:
		u16u16mdec.Decode(dec, p, tag)
	case *map[uint16]uint32:
		u16u32mdec.Decode(dec, p, tag)
	case *map[uint16]uint64:
		u16u64mdec.Decode(dec, p, tag)
	case *map[uint16]float32:
		u16f32mdec.Decode(dec, p, tag)
	case *map[uint16]float64:
		u16f64mdec.Decode(dec, p, tag)
	case *map[uint16]string:
		u16smdec.Decode(dec, p, tag)
	case *map[uint16]interface{}:
		u16ifmdec.Decode(dec, p, tag)
	default:
		return false
	}
	return true
}

func fastDecodeUint32Map(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[uint32]bool:
		u32bmdec.Decode(dec, p, tag)
	case *map[uint32]int:
		u32imdec.Decode(dec, p, tag)
	case *map[uint32]int8:
		u32i8mdec.Decode(dec, p, tag)
	case *map[uint32]int16:
		u32i16mdec.Decode(dec, p, tag)
	case *map[uint32]int32:
		u32i32mdec.Decode(dec, p, tag)
	case *map[uint32]int64:
		u32i64mdec.Decode(dec, p, tag)
	case *map[uint32]uint:
		u32umdec.Decode(dec, p, tag)
	case *map[uint32]uint8:
		u32u8mdec.Decode(dec, p, tag)
	case *map[uint32]uint16:
		u32u16mdec.Decode(dec, p, tag)
	case *map[uint32]uint32:
		u32u32mdec.Decode(dec, p, tag)
	case *map[uint32]uint64:
		u32u64mdec.Decode(dec, p, tag)
	case *map[uint32]float32:
		u32f32mdec.Decode(dec, p, tag)
	case *map[uint32]float64:
		u32f64mdec.Decode(dec, p, tag)
	case *map[uint32]string:
		u32smdec.Decode(dec, p, tag)
	case *map[uint32]interface{}:
		u32ifmdec.Decode(dec, p, tag)
	default:
		return false
	}
	return true
}

func fastDecodeUint64Map(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[uint64]bool:
		u64bmdec.Decode(dec, p, tag)
	case *map[uint64]int:
		u64imdec.Decode(dec, p, tag)
	case *map[uint64]int8:
		u64i8mdec.Decode(dec, p, tag)
	case *map[uint64]int16:
		u64i16mdec.Decode(dec, p, tag)
	case *map[uint64]int32:
		u64i32mdec.Decode(dec, p, tag)
	case *map[uint64]int64:
		u64i64mdec.Decode(dec, p, tag)
	case *map[uint64]uint:
		u64umdec.Decode(dec, p, tag)
	case *map[uint64]uint8:
		u64u8mdec.Decode(dec, p, tag)
	case *map[uint64]uint16:
		u64u16mdec.Decode(dec, p, tag)
	case *map[uint64]uint32:
		u64u32mdec.Decode(dec, p, tag)
	case *map[uint64]uint64:
		u64u64mdec.Decode(dec, p, tag)
	case *map[uint64]float32:
		u64f32mdec.Decode(dec, p, tag)
	case *map[uint64]float64:
		u64f64mdec.Decode(dec, p, tag)
	case *map[uint64]string:
		u64smdec.Decode(dec, p, tag)
	case *map[uint64]interface{}:
		u64ifmdec.Decode(dec, p, tag)
	default:
		return false
	}
	return true
}

func fastDecodeFloat32Map(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[float32]bool:
		f32bmdec.Decode(dec, p, tag)
	case *map[float32]int:
		f32imdec.Decode(dec, p, tag)
	case *map[float32]int8:
		f32i8mdec.Decode(dec, p, tag)
	case *map[float32]int16:
		f32i16mdec.Decode(dec, p, tag)
	case *map[float32]int32:
		f32i32mdec.Decode(dec, p, tag)
	case *map[float32]int64:
		f32i64mdec.Decode(dec, p, tag)
	case *map[float32]uint:
		f32umdec.Decode(dec, p, tag)
	case *map[float32]uint8:
		f32u8mdec.Decode(dec, p, tag)
	case *map[float32]uint16:
		f32u16mdec.Decode(dec, p, tag)
	case *map[float32]uint32:
		f32u32mdec.Decode(dec, p, tag)
	case *map[float32]uint64:
		f32u64mdec.Decode(dec, p, tag)
	case *map[float32]float32:
		f32f32mdec.Decode(dec, p, tag)
	case *map[float32]float64:
		f32f64mdec.Decode(dec, p, tag)
	case *map[float32]string:
		f32smdec.Decode(dec, p, tag)
	case *map[float32]interface{}:
		f32ifmdec.Decode(dec, p, tag)
	default:
		return false
	}
	return true
}

func fastDecodeFloat64Map(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[float64]bool:
		f64bmdec.Decode(dec, p, tag)
	case *map[float64]int:
		f64imdec.Decode(dec, p, tag)
	case *map[float64]int8:
		f64i8mdec.Decode(dec, p, tag)
	case *map[float64]int16:
		f64i16mdec.Decode(dec, p, tag)
	case *map[float64]int32:
		f64i32mdec.Decode(dec, p, tag)
	case *map[float64]int64:
		f64i64mdec.Decode(dec, p, tag)
	case *map[float64]uint:
		f64umdec.Decode(dec, p, tag)
	case *map[float64]uint8:
		f64u8mdec.Decode(dec, p, tag)
	case *map[float64]uint16:
		f64u16mdec.Decode(dec, p, tag)
	case *map[float64]uint32:
		f64u32mdec.Decode(dec, p, tag)
	case *map[float64]uint64:
		f64u64mdec.Decode(dec, p, tag)
	case *map[float64]float32:
		f64f32mdec.Decode(dec, p, tag)
	case *map[float64]float64:
		f64f64mdec.Decode(dec, p, tag)
	case *map[float64]string:
		f64smdec.Decode(dec, p, tag)
	case *map[float64]interface{}:
		f64ifmdec.Decode(dec, p, tag)
	default:
		return false
	}
	return true
}

func fastDecodeStringMap(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[string]bool:
		sbmdec.Decode(dec, p, tag)
	case *map[string]int:
		simdec.Decode(dec, p, tag)
	case *map[string]int8:
		si8mdec.Decode(dec, p, tag)
	case *map[string]int16:
		si16mdec.Decode(dec, p, tag)
	case *map[string]int32:
		si32mdec.Decode(dec, p, tag)
	case *map[string]int64:
		si64mdec.Decode(dec, p, tag)
	case *map[string]uint:
		sumdec.Decode(dec, p, tag)
	case *map[string]uint8:
		su8mdec.Decode(dec, p, tag)
	case *map[string]uint16:
		su16mdec.Decode(dec, p, tag)
	case *map[string]uint32:
		su32mdec.Decode(dec, p, tag)
	case *map[string]uint64:
		su64mdec.Decode(dec, p, tag)
	case *map[string]float32:
		sf32mdec.Decode(dec, p, tag)
	case *map[string]float64:
		sf64mdec.Decode(dec, p, tag)
	case *map[string]string:
		ssmdec.Decode(dec, p, tag)
	case *map[string]interface{}:
		sifmdec.Decode(dec, p, tag)
	default:
		return false
	}
	return true
}

func fastDecodeInterfaceMap(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[interface{}]bool:
		ifbmdec.Decode(dec, p, tag)
	case *map[interface{}]int:
		ifimdec.Decode(dec, p, tag)
	case *map[interface{}]int8:
		ifi8mdec.Decode(dec, p, tag)
	case *map[interface{}]int16:
		ifi16mdec.Decode(dec, p, tag)
	case *map[interface{}]int32:
		ifi32mdec.Decode(dec, p, tag)
	case *map[interface{}]int64:
		ifi64mdec.Decode(dec, p, tag)
	case *map[interface{}]uint:
		ifumdec.Decode(dec, p, tag)
	case *map[interface{}]uint8:
		ifu8mdec.Decode(dec, p, tag)
	case *map[interface{}]uint16:
		ifu16mdec.Decode(dec, p, tag)
	case *map[interface{}]uint32:
		ifu32mdec.Decode(dec, p, tag)
	case *map[interface{}]uint64:
		ifu64mdec.Decode(dec, p, tag)
	case *map[interface{}]float32:
		iff32mdec.Decode(dec, p, tag)
	case *map[interface{}]float64:
		iff64mdec.Decode(dec, p, tag)
	case *map[interface{}]string:
		ifsmdec.Decode(dec, p, tag)
	case *map[interface{}]interface{}:
		ififmdec.Decode(dec, p, tag)
	default:
		return false
	}
	return true
}

func (dec *Decoder) fastDecodeMap(t reflect.Type, p interface{}, tag byte) bool {
	switch t.Key().Kind() {
	case reflect.Int:
		return fastDecodeIntMap(dec, p, tag)
	case reflect.Int8:
		return fastDecodeInt8Map(dec, p, tag)
	case reflect.Int16:
		return fastDecodeInt16Map(dec, p, tag)
	case reflect.Int32:
		return fastDecodeInt32Map(dec, p, tag)
	case reflect.Int64:
		return fastDecodeInt64Map(dec, p, tag)
	case reflect.Uint:
		return fastDecodeUintMap(dec, p, tag)
	case reflect.Uint8:
		return fastDecodeUint8Map(dec, p, tag)
	case reflect.Uint16:
		return fastDecodeUint16Map(dec, p, tag)
	case reflect.Uint32:
		return fastDecodeUint32Map(dec, p, tag)
	case reflect.Uint64:
		return fastDecodeUint64Map(dec, p, tag)
	case reflect.Float32:
		return fastDecodeFloat32Map(dec, p, tag)
	case reflect.Float64:
		return fastDecodeFloat64Map(dec, p, tag)
	case reflect.Interface:
		return fastDecodeInterfaceMap(dec, p, tag)
	case reflect.String:
		return fastDecodeStringMap(dec, p, tag)
	}
	return false
}

var (
	ibmdec   mapDecoder
	iimdec   mapDecoder
	ii8mdec  mapDecoder
	ii16mdec mapDecoder
	ii32mdec mapDecoder
	ii64mdec mapDecoder
	iumdec   mapDecoder
	iu8mdec  mapDecoder
	iu16mdec mapDecoder
	iu32mdec mapDecoder
	iu64mdec mapDecoder
	if32mdec mapDecoder
	if64mdec mapDecoder
	ismdec   mapDecoder
	iifmdec  mapDecoder

	i8bmdec   mapDecoder
	i8imdec   mapDecoder
	i8i8mdec  mapDecoder
	i8i16mdec mapDecoder
	i8i32mdec mapDecoder
	i8i64mdec mapDecoder
	i8umdec   mapDecoder
	i8u8mdec  mapDecoder
	i8u16mdec mapDecoder
	i8u32mdec mapDecoder
	i8u64mdec mapDecoder
	i8f32mdec mapDecoder
	i8f64mdec mapDecoder
	i8smdec   mapDecoder
	i8ifmdec  mapDecoder

	i16bmdec   mapDecoder
	i16imdec   mapDecoder
	i16i8mdec  mapDecoder
	i16i16mdec mapDecoder
	i16i32mdec mapDecoder
	i16i64mdec mapDecoder
	i16umdec   mapDecoder
	i16u8mdec  mapDecoder
	i16u16mdec mapDecoder
	i16u32mdec mapDecoder
	i16u64mdec mapDecoder
	i16f32mdec mapDecoder
	i16f64mdec mapDecoder
	i16smdec   mapDecoder
	i16ifmdec  mapDecoder

	i32bmdec   mapDecoder
	i32imdec   mapDecoder
	i32i8mdec  mapDecoder
	i32i16mdec mapDecoder
	i32i32mdec mapDecoder
	i32i64mdec mapDecoder
	i32umdec   mapDecoder
	i32u8mdec  mapDecoder
	i32u16mdec mapDecoder
	i32u32mdec mapDecoder
	i32u64mdec mapDecoder
	i32f32mdec mapDecoder
	i32f64mdec mapDecoder
	i32smdec   mapDecoder
	i32ifmdec  mapDecoder

	i64bmdec   mapDecoder
	i64imdec   mapDecoder
	i64i8mdec  mapDecoder
	i64i16mdec mapDecoder
	i64i32mdec mapDecoder
	i64i64mdec mapDecoder
	i64umdec   mapDecoder
	i64u8mdec  mapDecoder
	i64u16mdec mapDecoder
	i64u32mdec mapDecoder
	i64u64mdec mapDecoder
	i64f32mdec mapDecoder
	i64f64mdec mapDecoder
	i64smdec   mapDecoder
	i64ifmdec  mapDecoder

	ubmdec   mapDecoder
	uimdec   mapDecoder
	ui8mdec  mapDecoder
	ui16mdec mapDecoder
	ui32mdec mapDecoder
	ui64mdec mapDecoder
	uumdec   mapDecoder
	uu8mdec  mapDecoder
	uu16mdec mapDecoder
	uu32mdec mapDecoder
	uu64mdec mapDecoder
	uf32mdec mapDecoder
	uf64mdec mapDecoder
	usmdec   mapDecoder
	uifmdec  mapDecoder

	u8bmdec   mapDecoder
	u8imdec   mapDecoder
	u8i8mdec  mapDecoder
	u8i16mdec mapDecoder
	u8i32mdec mapDecoder
	u8i64mdec mapDecoder
	u8umdec   mapDecoder
	u8u8mdec  mapDecoder
	u8u16mdec mapDecoder
	u8u32mdec mapDecoder
	u8u64mdec mapDecoder
	u8f32mdec mapDecoder
	u8f64mdec mapDecoder
	u8smdec   mapDecoder
	u8ifmdec  mapDecoder

	u16bmdec   mapDecoder
	u16imdec   mapDecoder
	u16i8mdec  mapDecoder
	u16i16mdec mapDecoder
	u16i32mdec mapDecoder
	u16i64mdec mapDecoder
	u16umdec   mapDecoder
	u16u8mdec  mapDecoder
	u16u16mdec mapDecoder
	u16u32mdec mapDecoder
	u16u64mdec mapDecoder
	u16f32mdec mapDecoder
	u16f64mdec mapDecoder
	u16smdec   mapDecoder
	u16ifmdec  mapDecoder

	u32bmdec   mapDecoder
	u32imdec   mapDecoder
	u32i8mdec  mapDecoder
	u32i16mdec mapDecoder
	u32i32mdec mapDecoder
	u32i64mdec mapDecoder
	u32umdec   mapDecoder
	u32u8mdec  mapDecoder
	u32u16mdec mapDecoder
	u32u32mdec mapDecoder
	u32u64mdec mapDecoder
	u32f32mdec mapDecoder
	u32f64mdec mapDecoder
	u32smdec   mapDecoder
	u32ifmdec  mapDecoder

	u64bmdec   mapDecoder
	u64imdec   mapDecoder
	u64i8mdec  mapDecoder
	u64i16mdec mapDecoder
	u64i32mdec mapDecoder
	u64i64mdec mapDecoder
	u64umdec   mapDecoder
	u64u8mdec  mapDecoder
	u64u16mdec mapDecoder
	u64u32mdec mapDecoder
	u64u64mdec mapDecoder
	u64f32mdec mapDecoder
	u64f64mdec mapDecoder
	u64smdec   mapDecoder
	u64ifmdec  mapDecoder

	f32bmdec   mapDecoder
	f32imdec   mapDecoder
	f32i8mdec  mapDecoder
	f32i16mdec mapDecoder
	f32i32mdec mapDecoder
	f32i64mdec mapDecoder
	f32umdec   mapDecoder
	f32u8mdec  mapDecoder
	f32u16mdec mapDecoder
	f32u32mdec mapDecoder
	f32u64mdec mapDecoder
	f32f32mdec mapDecoder
	f32f64mdec mapDecoder
	f32smdec   mapDecoder
	f32ifmdec  mapDecoder

	f64bmdec   mapDecoder
	f64imdec   mapDecoder
	f64i8mdec  mapDecoder
	f64i16mdec mapDecoder
	f64i32mdec mapDecoder
	f64i64mdec mapDecoder
	f64umdec   mapDecoder
	f64u8mdec  mapDecoder
	f64u16mdec mapDecoder
	f64u32mdec mapDecoder
	f64u64mdec mapDecoder
	f64f32mdec mapDecoder
	f64f64mdec mapDecoder
	f64smdec   mapDecoder
	f64ifmdec  mapDecoder

	sbmdec   mapDecoder
	simdec   mapDecoder
	si8mdec  mapDecoder
	si16mdec mapDecoder
	si32mdec mapDecoder
	si64mdec mapDecoder
	sumdec   mapDecoder
	su8mdec  mapDecoder
	su16mdec mapDecoder
	su32mdec mapDecoder
	su64mdec mapDecoder
	sf32mdec mapDecoder
	sf64mdec mapDecoder
	ssmdec   mapDecoder
	sifmdec  mapDecoder

	ifbmdec   mapDecoder
	ifimdec   mapDecoder
	ifi8mdec  mapDecoder
	ifi16mdec mapDecoder
	ifi32mdec mapDecoder
	ifi64mdec mapDecoder
	ifumdec   mapDecoder
	ifu8mdec  mapDecoder
	ifu16mdec mapDecoder
	ifu32mdec mapDecoder
	ifu64mdec mapDecoder
	iff32mdec mapDecoder
	iff64mdec mapDecoder
	ifsmdec   mapDecoder
	ififmdec  mapDecoder
)

func init() {
	ibmdec = getMapDecoder(reflect.TypeOf((map[int]bool)(nil))).(mapDecoder)
	iimdec = getMapDecoder(reflect.TypeOf((map[int]int)(nil))).(mapDecoder)
	ii8mdec = getMapDecoder(reflect.TypeOf((map[int]int8)(nil))).(mapDecoder)
	ii16mdec = getMapDecoder(reflect.TypeOf((map[int]int16)(nil))).(mapDecoder)
	ii32mdec = getMapDecoder(reflect.TypeOf((map[int]int32)(nil))).(mapDecoder)
	ii64mdec = getMapDecoder(reflect.TypeOf((map[int]int64)(nil))).(mapDecoder)
	iumdec = getMapDecoder(reflect.TypeOf((map[int]uint)(nil))).(mapDecoder)
	iu8mdec = getMapDecoder(reflect.TypeOf((map[int]uint8)(nil))).(mapDecoder)
	iu16mdec = getMapDecoder(reflect.TypeOf((map[int]uint16)(nil))).(mapDecoder)
	iu32mdec = getMapDecoder(reflect.TypeOf((map[int]uint32)(nil))).(mapDecoder)
	iu64mdec = getMapDecoder(reflect.TypeOf((map[int]uint64)(nil))).(mapDecoder)
	if32mdec = getMapDecoder(reflect.TypeOf((map[int]float32)(nil))).(mapDecoder)
	if64mdec = getMapDecoder(reflect.TypeOf((map[int]float64)(nil))).(mapDecoder)
	ismdec = getMapDecoder(reflect.TypeOf((map[int]string)(nil))).(mapDecoder)
	iifmdec = getMapDecoder(reflect.TypeOf((map[int]interface{})(nil))).(mapDecoder)

	i8bmdec = getMapDecoder(reflect.TypeOf((map[int8]bool)(nil))).(mapDecoder)
	i8imdec = getMapDecoder(reflect.TypeOf((map[int8]int)(nil))).(mapDecoder)
	i8i8mdec = getMapDecoder(reflect.TypeOf((map[int8]int8)(nil))).(mapDecoder)
	i8i16mdec = getMapDecoder(reflect.TypeOf((map[int8]int16)(nil))).(mapDecoder)
	i8i32mdec = getMapDecoder(reflect.TypeOf((map[int8]int32)(nil))).(mapDecoder)
	i8i64mdec = getMapDecoder(reflect.TypeOf((map[int8]int64)(nil))).(mapDecoder)
	i8umdec = getMapDecoder(reflect.TypeOf((map[int8]uint)(nil))).(mapDecoder)
	i8u8mdec = getMapDecoder(reflect.TypeOf((map[int8]uint8)(nil))).(mapDecoder)
	i8u16mdec = getMapDecoder(reflect.TypeOf((map[int8]uint16)(nil))).(mapDecoder)
	i8u32mdec = getMapDecoder(reflect.TypeOf((map[int8]uint32)(nil))).(mapDecoder)
	i8u64mdec = getMapDecoder(reflect.TypeOf((map[int8]uint64)(nil))).(mapDecoder)
	i8f32mdec = getMapDecoder(reflect.TypeOf((map[int8]float32)(nil))).(mapDecoder)
	i8f64mdec = getMapDecoder(reflect.TypeOf((map[int8]float64)(nil))).(mapDecoder)
	i8smdec = getMapDecoder(reflect.TypeOf((map[int8]string)(nil))).(mapDecoder)
	i8ifmdec = getMapDecoder(reflect.TypeOf((map[int8]interface{})(nil))).(mapDecoder)

	i16bmdec = getMapDecoder(reflect.TypeOf((map[int16]bool)(nil))).(mapDecoder)
	i16imdec = getMapDecoder(reflect.TypeOf((map[int16]int)(nil))).(mapDecoder)
	i16i8mdec = getMapDecoder(reflect.TypeOf((map[int16]int8)(nil))).(mapDecoder)
	i16i16mdec = getMapDecoder(reflect.TypeOf((map[int16]int16)(nil))).(mapDecoder)
	i16i32mdec = getMapDecoder(reflect.TypeOf((map[int16]int32)(nil))).(mapDecoder)
	i16i64mdec = getMapDecoder(reflect.TypeOf((map[int16]int64)(nil))).(mapDecoder)
	i16umdec = getMapDecoder(reflect.TypeOf((map[int16]uint)(nil))).(mapDecoder)
	i16u8mdec = getMapDecoder(reflect.TypeOf((map[int16]uint8)(nil))).(mapDecoder)
	i16u16mdec = getMapDecoder(reflect.TypeOf((map[int16]uint16)(nil))).(mapDecoder)
	i16u32mdec = getMapDecoder(reflect.TypeOf((map[int16]uint32)(nil))).(mapDecoder)
	i16u64mdec = getMapDecoder(reflect.TypeOf((map[int16]uint64)(nil))).(mapDecoder)
	i16f32mdec = getMapDecoder(reflect.TypeOf((map[int16]float32)(nil))).(mapDecoder)
	i16f64mdec = getMapDecoder(reflect.TypeOf((map[int16]float64)(nil))).(mapDecoder)
	i16smdec = getMapDecoder(reflect.TypeOf((map[int16]string)(nil))).(mapDecoder)
	i16ifmdec = getMapDecoder(reflect.TypeOf((map[int16]interface{})(nil))).(mapDecoder)

	i32bmdec = getMapDecoder(reflect.TypeOf((map[int32]bool)(nil))).(mapDecoder)
	i32imdec = getMapDecoder(reflect.TypeOf((map[int32]int)(nil))).(mapDecoder)
	i32i8mdec = getMapDecoder(reflect.TypeOf((map[int32]int8)(nil))).(mapDecoder)
	i32i16mdec = getMapDecoder(reflect.TypeOf((map[int32]int16)(nil))).(mapDecoder)
	i32i32mdec = getMapDecoder(reflect.TypeOf((map[int32]int32)(nil))).(mapDecoder)
	i32i64mdec = getMapDecoder(reflect.TypeOf((map[int32]int64)(nil))).(mapDecoder)
	i32umdec = getMapDecoder(reflect.TypeOf((map[int32]uint)(nil))).(mapDecoder)
	i32u8mdec = getMapDecoder(reflect.TypeOf((map[int32]uint8)(nil))).(mapDecoder)
	i32u16mdec = getMapDecoder(reflect.TypeOf((map[int32]uint16)(nil))).(mapDecoder)
	i32u32mdec = getMapDecoder(reflect.TypeOf((map[int32]uint32)(nil))).(mapDecoder)
	i32u64mdec = getMapDecoder(reflect.TypeOf((map[int32]uint64)(nil))).(mapDecoder)
	i32f32mdec = getMapDecoder(reflect.TypeOf((map[int32]float32)(nil))).(mapDecoder)
	i32f64mdec = getMapDecoder(reflect.TypeOf((map[int32]float64)(nil))).(mapDecoder)
	i32smdec = getMapDecoder(reflect.TypeOf((map[int32]string)(nil))).(mapDecoder)
	i32ifmdec = getMapDecoder(reflect.TypeOf((map[int32]interface{})(nil))).(mapDecoder)

	i64bmdec = getMapDecoder(reflect.TypeOf((map[int64]bool)(nil))).(mapDecoder)
	i64imdec = getMapDecoder(reflect.TypeOf((map[int64]int)(nil))).(mapDecoder)
	i64i8mdec = getMapDecoder(reflect.TypeOf((map[int64]int8)(nil))).(mapDecoder)
	i64i16mdec = getMapDecoder(reflect.TypeOf((map[int64]int16)(nil))).(mapDecoder)
	i64i32mdec = getMapDecoder(reflect.TypeOf((map[int64]int32)(nil))).(mapDecoder)
	i64i64mdec = getMapDecoder(reflect.TypeOf((map[int64]int64)(nil))).(mapDecoder)
	i64umdec = getMapDecoder(reflect.TypeOf((map[int64]uint)(nil))).(mapDecoder)
	i64u8mdec = getMapDecoder(reflect.TypeOf((map[int64]uint8)(nil))).(mapDecoder)
	i64u16mdec = getMapDecoder(reflect.TypeOf((map[int64]uint16)(nil))).(mapDecoder)
	i64u32mdec = getMapDecoder(reflect.TypeOf((map[int64]uint32)(nil))).(mapDecoder)
	i64u64mdec = getMapDecoder(reflect.TypeOf((map[int64]uint64)(nil))).(mapDecoder)
	i64f32mdec = getMapDecoder(reflect.TypeOf((map[int64]float32)(nil))).(mapDecoder)
	i64f64mdec = getMapDecoder(reflect.TypeOf((map[int64]float64)(nil))).(mapDecoder)
	i64smdec = getMapDecoder(reflect.TypeOf((map[int64]string)(nil))).(mapDecoder)
	i64ifmdec = getMapDecoder(reflect.TypeOf((map[int64]interface{})(nil))).(mapDecoder)

	ubmdec = getMapDecoder(reflect.TypeOf((map[uint]bool)(nil))).(mapDecoder)
	uimdec = getMapDecoder(reflect.TypeOf((map[uint]int)(nil))).(mapDecoder)
	ui8mdec = getMapDecoder(reflect.TypeOf((map[uint]int8)(nil))).(mapDecoder)
	ui16mdec = getMapDecoder(reflect.TypeOf((map[uint]int16)(nil))).(mapDecoder)
	ui32mdec = getMapDecoder(reflect.TypeOf((map[uint]int32)(nil))).(mapDecoder)
	ui64mdec = getMapDecoder(reflect.TypeOf((map[uint]int64)(nil))).(mapDecoder)
	uumdec = getMapDecoder(reflect.TypeOf((map[uint]uint)(nil))).(mapDecoder)
	uu8mdec = getMapDecoder(reflect.TypeOf((map[uint]uint8)(nil))).(mapDecoder)
	uu16mdec = getMapDecoder(reflect.TypeOf((map[uint]uint16)(nil))).(mapDecoder)
	uu32mdec = getMapDecoder(reflect.TypeOf((map[uint]uint32)(nil))).(mapDecoder)
	uu64mdec = getMapDecoder(reflect.TypeOf((map[uint]uint64)(nil))).(mapDecoder)
	uf32mdec = getMapDecoder(reflect.TypeOf((map[uint]float32)(nil))).(mapDecoder)
	uf64mdec = getMapDecoder(reflect.TypeOf((map[uint]float64)(nil))).(mapDecoder)
	usmdec = getMapDecoder(reflect.TypeOf((map[uint]string)(nil))).(mapDecoder)
	uifmdec = getMapDecoder(reflect.TypeOf((map[uint]interface{})(nil))).(mapDecoder)

	u8bmdec = getMapDecoder(reflect.TypeOf((map[uint8]bool)(nil))).(mapDecoder)
	u8imdec = getMapDecoder(reflect.TypeOf((map[uint8]int)(nil))).(mapDecoder)
	u8i8mdec = getMapDecoder(reflect.TypeOf((map[uint8]int8)(nil))).(mapDecoder)
	u8i16mdec = getMapDecoder(reflect.TypeOf((map[uint8]int16)(nil))).(mapDecoder)
	u8i32mdec = getMapDecoder(reflect.TypeOf((map[uint8]int32)(nil))).(mapDecoder)
	u8i64mdec = getMapDecoder(reflect.TypeOf((map[uint8]int64)(nil))).(mapDecoder)
	u8umdec = getMapDecoder(reflect.TypeOf((map[uint8]uint)(nil))).(mapDecoder)
	u8u8mdec = getMapDecoder(reflect.TypeOf((map[uint8]uint8)(nil))).(mapDecoder)
	u8u16mdec = getMapDecoder(reflect.TypeOf((map[uint8]uint16)(nil))).(mapDecoder)
	u8u32mdec = getMapDecoder(reflect.TypeOf((map[uint8]uint32)(nil))).(mapDecoder)
	u8u64mdec = getMapDecoder(reflect.TypeOf((map[uint8]uint64)(nil))).(mapDecoder)
	u8f32mdec = getMapDecoder(reflect.TypeOf((map[uint8]float32)(nil))).(mapDecoder)
	u8f64mdec = getMapDecoder(reflect.TypeOf((map[uint8]float64)(nil))).(mapDecoder)
	u8smdec = getMapDecoder(reflect.TypeOf((map[uint8]string)(nil))).(mapDecoder)
	u8ifmdec = getMapDecoder(reflect.TypeOf((map[uint8]interface{})(nil))).(mapDecoder)

	u16bmdec = getMapDecoder(reflect.TypeOf((map[uint16]bool)(nil))).(mapDecoder)
	u16imdec = getMapDecoder(reflect.TypeOf((map[uint16]int)(nil))).(mapDecoder)
	u16i8mdec = getMapDecoder(reflect.TypeOf((map[uint16]int8)(nil))).(mapDecoder)
	u16i16mdec = getMapDecoder(reflect.TypeOf((map[uint16]int16)(nil))).(mapDecoder)
	u16i32mdec = getMapDecoder(reflect.TypeOf((map[uint16]int32)(nil))).(mapDecoder)
	u16i64mdec = getMapDecoder(reflect.TypeOf((map[uint16]int64)(nil))).(mapDecoder)
	u16umdec = getMapDecoder(reflect.TypeOf((map[uint16]uint)(nil))).(mapDecoder)
	u16u8mdec = getMapDecoder(reflect.TypeOf((map[uint16]uint8)(nil))).(mapDecoder)
	u16u16mdec = getMapDecoder(reflect.TypeOf((map[uint16]uint16)(nil))).(mapDecoder)
	u16u32mdec = getMapDecoder(reflect.TypeOf((map[uint16]uint32)(nil))).(mapDecoder)
	u16u64mdec = getMapDecoder(reflect.TypeOf((map[uint16]uint64)(nil))).(mapDecoder)
	u16f32mdec = getMapDecoder(reflect.TypeOf((map[uint16]float32)(nil))).(mapDecoder)
	u16f64mdec = getMapDecoder(reflect.TypeOf((map[uint16]float64)(nil))).(mapDecoder)
	u16smdec = getMapDecoder(reflect.TypeOf((map[uint16]string)(nil))).(mapDecoder)
	u16ifmdec = getMapDecoder(reflect.TypeOf((map[uint16]interface{})(nil))).(mapDecoder)

	u32bmdec = getMapDecoder(reflect.TypeOf((map[uint32]bool)(nil))).(mapDecoder)
	u32imdec = getMapDecoder(reflect.TypeOf((map[uint32]int)(nil))).(mapDecoder)
	u32i8mdec = getMapDecoder(reflect.TypeOf((map[uint32]int8)(nil))).(mapDecoder)
	u32i16mdec = getMapDecoder(reflect.TypeOf((map[uint32]int16)(nil))).(mapDecoder)
	u32i32mdec = getMapDecoder(reflect.TypeOf((map[uint32]int32)(nil))).(mapDecoder)
	u32i64mdec = getMapDecoder(reflect.TypeOf((map[uint32]int64)(nil))).(mapDecoder)
	u32umdec = getMapDecoder(reflect.TypeOf((map[uint32]uint)(nil))).(mapDecoder)
	u32u8mdec = getMapDecoder(reflect.TypeOf((map[uint32]uint8)(nil))).(mapDecoder)
	u32u16mdec = getMapDecoder(reflect.TypeOf((map[uint32]uint16)(nil))).(mapDecoder)
	u32u32mdec = getMapDecoder(reflect.TypeOf((map[uint32]uint32)(nil))).(mapDecoder)
	u32u64mdec = getMapDecoder(reflect.TypeOf((map[uint32]uint64)(nil))).(mapDecoder)
	u32f32mdec = getMapDecoder(reflect.TypeOf((map[uint32]float32)(nil))).(mapDecoder)
	u32f64mdec = getMapDecoder(reflect.TypeOf((map[uint32]float64)(nil))).(mapDecoder)
	u32smdec = getMapDecoder(reflect.TypeOf((map[uint32]string)(nil))).(mapDecoder)
	u32ifmdec = getMapDecoder(reflect.TypeOf((map[uint32]interface{})(nil))).(mapDecoder)

	u64bmdec = getMapDecoder(reflect.TypeOf((map[uint64]bool)(nil))).(mapDecoder)
	u64imdec = getMapDecoder(reflect.TypeOf((map[uint64]int)(nil))).(mapDecoder)
	u64i8mdec = getMapDecoder(reflect.TypeOf((map[uint64]int8)(nil))).(mapDecoder)
	u64i16mdec = getMapDecoder(reflect.TypeOf((map[uint64]int16)(nil))).(mapDecoder)
	u64i32mdec = getMapDecoder(reflect.TypeOf((map[uint64]int32)(nil))).(mapDecoder)
	u64i64mdec = getMapDecoder(reflect.TypeOf((map[uint64]int64)(nil))).(mapDecoder)
	u64umdec = getMapDecoder(reflect.TypeOf((map[uint64]uint)(nil))).(mapDecoder)
	u64u8mdec = getMapDecoder(reflect.TypeOf((map[uint64]uint8)(nil))).(mapDecoder)
	u64u16mdec = getMapDecoder(reflect.TypeOf((map[uint64]uint16)(nil))).(mapDecoder)
	u64u32mdec = getMapDecoder(reflect.TypeOf((map[uint64]uint32)(nil))).(mapDecoder)
	u64u64mdec = getMapDecoder(reflect.TypeOf((map[uint64]uint64)(nil))).(mapDecoder)
	u64f32mdec = getMapDecoder(reflect.TypeOf((map[uint64]float32)(nil))).(mapDecoder)
	u64f64mdec = getMapDecoder(reflect.TypeOf((map[uint64]float64)(nil))).(mapDecoder)
	u64smdec = getMapDecoder(reflect.TypeOf((map[uint64]string)(nil))).(mapDecoder)
	u64ifmdec = getMapDecoder(reflect.TypeOf((map[uint64]interface{})(nil))).(mapDecoder)

	f32bmdec = getMapDecoder(reflect.TypeOf((map[float32]bool)(nil))).(mapDecoder)
	f32imdec = getMapDecoder(reflect.TypeOf((map[float32]int)(nil))).(mapDecoder)
	f32i8mdec = getMapDecoder(reflect.TypeOf((map[float32]int8)(nil))).(mapDecoder)
	f32i16mdec = getMapDecoder(reflect.TypeOf((map[float32]int16)(nil))).(mapDecoder)
	f32i32mdec = getMapDecoder(reflect.TypeOf((map[float32]int32)(nil))).(mapDecoder)
	f32i64mdec = getMapDecoder(reflect.TypeOf((map[float32]int64)(nil))).(mapDecoder)
	f32umdec = getMapDecoder(reflect.TypeOf((map[float32]uint)(nil))).(mapDecoder)
	f32u8mdec = getMapDecoder(reflect.TypeOf((map[float32]uint8)(nil))).(mapDecoder)
	f32u16mdec = getMapDecoder(reflect.TypeOf((map[float32]uint16)(nil))).(mapDecoder)
	f32u32mdec = getMapDecoder(reflect.TypeOf((map[float32]uint32)(nil))).(mapDecoder)
	f32u64mdec = getMapDecoder(reflect.TypeOf((map[float32]uint64)(nil))).(mapDecoder)
	f32f32mdec = getMapDecoder(reflect.TypeOf((map[float32]float32)(nil))).(mapDecoder)
	f32f64mdec = getMapDecoder(reflect.TypeOf((map[float32]float64)(nil))).(mapDecoder)
	f32smdec = getMapDecoder(reflect.TypeOf((map[float32]string)(nil))).(mapDecoder)
	f32ifmdec = getMapDecoder(reflect.TypeOf((map[float32]interface{})(nil))).(mapDecoder)

	f64bmdec = getMapDecoder(reflect.TypeOf((map[float64]bool)(nil))).(mapDecoder)
	f64imdec = getMapDecoder(reflect.TypeOf((map[float64]int)(nil))).(mapDecoder)
	f64i8mdec = getMapDecoder(reflect.TypeOf((map[float64]int8)(nil))).(mapDecoder)
	f64i16mdec = getMapDecoder(reflect.TypeOf((map[float64]int16)(nil))).(mapDecoder)
	f64i32mdec = getMapDecoder(reflect.TypeOf((map[float64]int32)(nil))).(mapDecoder)
	f64i64mdec = getMapDecoder(reflect.TypeOf((map[float64]int64)(nil))).(mapDecoder)
	f64umdec = getMapDecoder(reflect.TypeOf((map[float64]uint)(nil))).(mapDecoder)
	f64u8mdec = getMapDecoder(reflect.TypeOf((map[float64]uint8)(nil))).(mapDecoder)
	f64u16mdec = getMapDecoder(reflect.TypeOf((map[float64]uint16)(nil))).(mapDecoder)
	f64u32mdec = getMapDecoder(reflect.TypeOf((map[float64]uint32)(nil))).(mapDecoder)
	f64u64mdec = getMapDecoder(reflect.TypeOf((map[float64]uint64)(nil))).(mapDecoder)
	f64f32mdec = getMapDecoder(reflect.TypeOf((map[float64]float32)(nil))).(mapDecoder)
	f64f64mdec = getMapDecoder(reflect.TypeOf((map[float64]float64)(nil))).(mapDecoder)
	f64smdec = getMapDecoder(reflect.TypeOf((map[float64]string)(nil))).(mapDecoder)
	f64ifmdec = getMapDecoder(reflect.TypeOf((map[float64]interface{})(nil))).(mapDecoder)

	sbmdec = getMapDecoder(reflect.TypeOf((map[string]bool)(nil))).(mapDecoder)
	simdec = getMapDecoder(reflect.TypeOf((map[string]int)(nil))).(mapDecoder)
	si8mdec = getMapDecoder(reflect.TypeOf((map[string]int8)(nil))).(mapDecoder)
	si16mdec = getMapDecoder(reflect.TypeOf((map[string]int16)(nil))).(mapDecoder)
	si32mdec = getMapDecoder(reflect.TypeOf((map[string]int32)(nil))).(mapDecoder)
	si64mdec = getMapDecoder(reflect.TypeOf((map[string]int64)(nil))).(mapDecoder)
	sumdec = getMapDecoder(reflect.TypeOf((map[string]uint)(nil))).(mapDecoder)
	su8mdec = getMapDecoder(reflect.TypeOf((map[string]uint8)(nil))).(mapDecoder)
	su16mdec = getMapDecoder(reflect.TypeOf((map[string]uint16)(nil))).(mapDecoder)
	su32mdec = getMapDecoder(reflect.TypeOf((map[string]uint32)(nil))).(mapDecoder)
	su64mdec = getMapDecoder(reflect.TypeOf((map[string]uint64)(nil))).(mapDecoder)
	sf32mdec = getMapDecoder(reflect.TypeOf((map[string]float32)(nil))).(mapDecoder)
	sf64mdec = getMapDecoder(reflect.TypeOf((map[string]float64)(nil))).(mapDecoder)
	ssmdec = getMapDecoder(reflect.TypeOf((map[string]string)(nil))).(mapDecoder)
	sifmdec = getMapDecoder(reflect.TypeOf((map[string]interface{})(nil))).(mapDecoder)

	ifbmdec = getMapDecoder(reflect.TypeOf((map[interface{}]bool)(nil))).(mapDecoder)
	ifimdec = getMapDecoder(reflect.TypeOf((map[interface{}]int)(nil))).(mapDecoder)
	ifi8mdec = getMapDecoder(reflect.TypeOf((map[interface{}]int8)(nil))).(mapDecoder)
	ifi16mdec = getMapDecoder(reflect.TypeOf((map[interface{}]int16)(nil))).(mapDecoder)
	ifi32mdec = getMapDecoder(reflect.TypeOf((map[interface{}]int32)(nil))).(mapDecoder)
	ifi64mdec = getMapDecoder(reflect.TypeOf((map[interface{}]int64)(nil))).(mapDecoder)
	ifumdec = getMapDecoder(reflect.TypeOf((map[interface{}]uint)(nil))).(mapDecoder)
	ifu8mdec = getMapDecoder(reflect.TypeOf((map[interface{}]uint8)(nil))).(mapDecoder)
	ifu16mdec = getMapDecoder(reflect.TypeOf((map[interface{}]uint16)(nil))).(mapDecoder)
	ifu32mdec = getMapDecoder(reflect.TypeOf((map[interface{}]uint32)(nil))).(mapDecoder)
	ifu64mdec = getMapDecoder(reflect.TypeOf((map[interface{}]uint64)(nil))).(mapDecoder)
	iff32mdec = getMapDecoder(reflect.TypeOf((map[interface{}]float32)(nil))).(mapDecoder)
	iff64mdec = getMapDecoder(reflect.TypeOf((map[interface{}]float64)(nil))).(mapDecoder)
	ifsmdec = getMapDecoder(reflect.TypeOf((map[interface{}]string)(nil))).(mapDecoder)
	ififmdec = getMapDecoder(reflect.TypeOf((map[interface{}]interface{})(nil))).(mapDecoder)

	RegisterValueDecoder(ibmdec)
	RegisterValueDecoder(iimdec)
	RegisterValueDecoder(ii8mdec)
	RegisterValueDecoder(ii16mdec)
	RegisterValueDecoder(ii32mdec)
	RegisterValueDecoder(ii64mdec)
	RegisterValueDecoder(iumdec)
	RegisterValueDecoder(iu8mdec)
	RegisterValueDecoder(iu16mdec)
	RegisterValueDecoder(iu32mdec)
	RegisterValueDecoder(iu64mdec)
	RegisterValueDecoder(if32mdec)
	RegisterValueDecoder(if64mdec)
	RegisterValueDecoder(ismdec)
	RegisterValueDecoder(iifmdec)

	RegisterValueDecoder(i8bmdec)
	RegisterValueDecoder(i8imdec)
	RegisterValueDecoder(i8i8mdec)
	RegisterValueDecoder(i8i16mdec)
	RegisterValueDecoder(i8i32mdec)
	RegisterValueDecoder(i8i64mdec)
	RegisterValueDecoder(i8umdec)
	RegisterValueDecoder(i8u8mdec)
	RegisterValueDecoder(i8u16mdec)
	RegisterValueDecoder(i8u32mdec)
	RegisterValueDecoder(i8u64mdec)
	RegisterValueDecoder(i8f32mdec)
	RegisterValueDecoder(i8f64mdec)
	RegisterValueDecoder(i8smdec)
	RegisterValueDecoder(i8ifmdec)

	RegisterValueDecoder(i16bmdec)
	RegisterValueDecoder(i16imdec)
	RegisterValueDecoder(i16i8mdec)
	RegisterValueDecoder(i16i16mdec)
	RegisterValueDecoder(i16i32mdec)
	RegisterValueDecoder(i16i64mdec)
	RegisterValueDecoder(i16umdec)
	RegisterValueDecoder(i16u8mdec)
	RegisterValueDecoder(i16u16mdec)
	RegisterValueDecoder(i16u32mdec)
	RegisterValueDecoder(i16u64mdec)
	RegisterValueDecoder(i16f32mdec)
	RegisterValueDecoder(i16f64mdec)
	RegisterValueDecoder(i16smdec)
	RegisterValueDecoder(i16ifmdec)

	RegisterValueDecoder(i32bmdec)
	RegisterValueDecoder(i32imdec)
	RegisterValueDecoder(i32i8mdec)
	RegisterValueDecoder(i32i16mdec)
	RegisterValueDecoder(i32i32mdec)
	RegisterValueDecoder(i32i64mdec)
	RegisterValueDecoder(i32umdec)
	RegisterValueDecoder(i32u8mdec)
	RegisterValueDecoder(i32u16mdec)
	RegisterValueDecoder(i32u32mdec)
	RegisterValueDecoder(i32u64mdec)
	RegisterValueDecoder(i32f32mdec)
	RegisterValueDecoder(i32f64mdec)
	RegisterValueDecoder(i32smdec)
	RegisterValueDecoder(i32ifmdec)

	RegisterValueDecoder(i64bmdec)
	RegisterValueDecoder(i64imdec)
	RegisterValueDecoder(i64i8mdec)
	RegisterValueDecoder(i64i16mdec)
	RegisterValueDecoder(i64i32mdec)
	RegisterValueDecoder(i64i64mdec)
	RegisterValueDecoder(i64umdec)
	RegisterValueDecoder(i64u8mdec)
	RegisterValueDecoder(i64u16mdec)
	RegisterValueDecoder(i64u32mdec)
	RegisterValueDecoder(i64u64mdec)
	RegisterValueDecoder(i64f32mdec)
	RegisterValueDecoder(i64f64mdec)
	RegisterValueDecoder(i64smdec)
	RegisterValueDecoder(i64ifmdec)

	RegisterValueDecoder(ubmdec)
	RegisterValueDecoder(uimdec)
	RegisterValueDecoder(ui8mdec)
	RegisterValueDecoder(ui16mdec)
	RegisterValueDecoder(ui32mdec)
	RegisterValueDecoder(ui64mdec)
	RegisterValueDecoder(uumdec)
	RegisterValueDecoder(uu8mdec)
	RegisterValueDecoder(uu16mdec)
	RegisterValueDecoder(uu32mdec)
	RegisterValueDecoder(uu64mdec)
	RegisterValueDecoder(uf32mdec)
	RegisterValueDecoder(uf64mdec)
	RegisterValueDecoder(usmdec)
	RegisterValueDecoder(uifmdec)

	RegisterValueDecoder(u8bmdec)
	RegisterValueDecoder(u8imdec)
	RegisterValueDecoder(u8i8mdec)
	RegisterValueDecoder(u8i16mdec)
	RegisterValueDecoder(u8i32mdec)
	RegisterValueDecoder(u8i64mdec)
	RegisterValueDecoder(u8umdec)
	RegisterValueDecoder(u8u8mdec)
	RegisterValueDecoder(u8u16mdec)
	RegisterValueDecoder(u8u32mdec)
	RegisterValueDecoder(u8u64mdec)
	RegisterValueDecoder(u8f32mdec)
	RegisterValueDecoder(u8f64mdec)
	RegisterValueDecoder(u8smdec)
	RegisterValueDecoder(u8ifmdec)

	RegisterValueDecoder(u16bmdec)
	RegisterValueDecoder(u16imdec)
	RegisterValueDecoder(u16i8mdec)
	RegisterValueDecoder(u16i16mdec)
	RegisterValueDecoder(u16i32mdec)
	RegisterValueDecoder(u16i64mdec)
	RegisterValueDecoder(u16umdec)
	RegisterValueDecoder(u16u8mdec)
	RegisterValueDecoder(u16u16mdec)
	RegisterValueDecoder(u16u32mdec)
	RegisterValueDecoder(u16u64mdec)
	RegisterValueDecoder(u16f32mdec)
	RegisterValueDecoder(u16f64mdec)
	RegisterValueDecoder(u16smdec)
	RegisterValueDecoder(u16ifmdec)

	RegisterValueDecoder(u32bmdec)
	RegisterValueDecoder(u32imdec)
	RegisterValueDecoder(u32i8mdec)
	RegisterValueDecoder(u32i16mdec)
	RegisterValueDecoder(u32i32mdec)
	RegisterValueDecoder(u32i64mdec)
	RegisterValueDecoder(u32umdec)
	RegisterValueDecoder(u32u8mdec)
	RegisterValueDecoder(u32u16mdec)
	RegisterValueDecoder(u32u32mdec)
	RegisterValueDecoder(u32u64mdec)
	RegisterValueDecoder(u32f32mdec)
	RegisterValueDecoder(u32f64mdec)
	RegisterValueDecoder(u32smdec)
	RegisterValueDecoder(u32ifmdec)

	RegisterValueDecoder(u64bmdec)
	RegisterValueDecoder(u64imdec)
	RegisterValueDecoder(u64i8mdec)
	RegisterValueDecoder(u64i16mdec)
	RegisterValueDecoder(u64i32mdec)
	RegisterValueDecoder(u64i64mdec)
	RegisterValueDecoder(u64umdec)
	RegisterValueDecoder(u64u8mdec)
	RegisterValueDecoder(u64u16mdec)
	RegisterValueDecoder(u64u32mdec)
	RegisterValueDecoder(u64u64mdec)
	RegisterValueDecoder(u64f32mdec)
	RegisterValueDecoder(u64f64mdec)
	RegisterValueDecoder(u64smdec)
	RegisterValueDecoder(u64ifmdec)

	RegisterValueDecoder(f32bmdec)
	RegisterValueDecoder(f32imdec)
	RegisterValueDecoder(f32i8mdec)
	RegisterValueDecoder(f32i16mdec)
	RegisterValueDecoder(f32i32mdec)
	RegisterValueDecoder(f32i64mdec)
	RegisterValueDecoder(f32umdec)
	RegisterValueDecoder(f32u8mdec)
	RegisterValueDecoder(f32u16mdec)
	RegisterValueDecoder(f32u32mdec)
	RegisterValueDecoder(f32u64mdec)
	RegisterValueDecoder(f32f32mdec)
	RegisterValueDecoder(f32f64mdec)
	RegisterValueDecoder(f32smdec)
	RegisterValueDecoder(f32ifmdec)

	RegisterValueDecoder(f64bmdec)
	RegisterValueDecoder(f64imdec)
	RegisterValueDecoder(f64i8mdec)
	RegisterValueDecoder(f64i16mdec)
	RegisterValueDecoder(f64i32mdec)
	RegisterValueDecoder(f64i64mdec)
	RegisterValueDecoder(f64umdec)
	RegisterValueDecoder(f64u8mdec)
	RegisterValueDecoder(f64u16mdec)
	RegisterValueDecoder(f64u32mdec)
	RegisterValueDecoder(f64u64mdec)
	RegisterValueDecoder(f64f32mdec)
	RegisterValueDecoder(f64f64mdec)
	RegisterValueDecoder(f64smdec)
	RegisterValueDecoder(f64ifmdec)

	RegisterValueDecoder(sbmdec)
	RegisterValueDecoder(simdec)
	RegisterValueDecoder(si8mdec)
	RegisterValueDecoder(si16mdec)
	RegisterValueDecoder(si32mdec)
	RegisterValueDecoder(si64mdec)
	RegisterValueDecoder(sumdec)
	RegisterValueDecoder(su8mdec)
	RegisterValueDecoder(su16mdec)
	RegisterValueDecoder(su32mdec)
	RegisterValueDecoder(su64mdec)
	RegisterValueDecoder(sf32mdec)
	RegisterValueDecoder(sf64mdec)
	RegisterValueDecoder(ssmdec)
	RegisterValueDecoder(sifmdec)

	RegisterValueDecoder(ifbmdec)
	RegisterValueDecoder(ifimdec)
	RegisterValueDecoder(ifi8mdec)
	RegisterValueDecoder(ifi16mdec)
	RegisterValueDecoder(ifi32mdec)
	RegisterValueDecoder(ifi64mdec)
	RegisterValueDecoder(ifumdec)
	RegisterValueDecoder(ifu8mdec)
	RegisterValueDecoder(ifu16mdec)
	RegisterValueDecoder(ifu32mdec)
	RegisterValueDecoder(ifu64mdec)
	RegisterValueDecoder(iff32mdec)
	RegisterValueDecoder(iff64mdec)
	RegisterValueDecoder(ifsmdec)
	RegisterValueDecoder(ififmdec)
}
