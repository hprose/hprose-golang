/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/map_decoder.go                                  |
|                                                          |
| LastModified: May 11, 2021                               |
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
	kt          reflect2.Type
	vt          reflect2.Type
	decodeKey   DecodeHandler
	decodeValue DecodeHandler
}

func (valdec mapDecoder) canDecodeListAsMap() bool {
	switch valdec.kt.Kind() {
	case reflect.String,
		reflect.Interface,
		reflect.Int,
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
		reflect.Complex128:
		return true
	}
	return false
}

func (valdec mapDecoder) convertKey(i int, p unsafe.Pointer) {
	switch valdec.kt.Kind() {
	case reflect.String:
		*(*string)(p) = strconv.Itoa(i)
	case reflect.Interface:
		*(*interface{})(p) = i
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
	}
}

func (valdec mapDecoder) canDecodeObjectAsMap() bool {
	ktKind := valdec.kt.Kind()
	vtKind := valdec.vt.Kind()
	if (ktKind == reflect.String || ktKind == reflect.Interface) &&
		(vtKind == reflect.Interface) {
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
	kp := valdec.kt.UnsafeNew()
	vp := valdec.vt.UnsafeNew()
	vt := valdec.vt.Type1()
	for i := 0; i < count; i++ {
		valdec.convertKey(i, kp)
		valdec.decodeValue(dec, vt, vp)
		valdec.t.UnsafeSetIndex(mp, kp, vp)
	}
	dec.Skip()
}

func (valdec mapDecoder) decodeMap(dec *Decoder, p interface{}) {
	mp := reflect2.PtrOf(p)
	count := dec.ReadInt()
	valdec.t.UnsafeSet(mp, valdec.t.UnsafeMakeMap(count))
	dec.AddReference(p)
	kp := valdec.kt.UnsafeNew()
	vp := valdec.vt.UnsafeNew()
	kt := valdec.kt.Type1()
	vt := valdec.vt.Type1()
	for i := 0; i < count; i++ {
		valdec.decodeKey(dec, kt, kp)
		valdec.decodeValue(dec, vt, vp)
		valdec.t.UnsafeSetIndex(mp, kp, vp)
	}
	dec.Skip()
}

func (valdec mapDecoder) decodeObjectAsMap(dec *Decoder, p interface{}, tag byte) {
	if !valdec.canDecodeObjectAsMap() {
		dec.decodeError(valdec.t.Type1(), tag)
		return
	}
	index := dec.ReadInt()
	structInfo := dec.getStructInfo(index)
	mp := reflect2.PtrOf(p)
	count := len(structInfo.names)
	valdec.t.UnsafeSet(mp, valdec.t.UnsafeMakeMap(count))
	dec.AddReference(p)
	if fields := structInfo.fields; fields != nil {
		for _, name := range structInfo.names {
			field := fields[name]
			vp := field.Type.UnsafeNew()
			field.Decode(dec, field.Type.Type1(), vp)
			v := field.Type.UnsafeIndirect(vp)
			valdec.t.UnsafeSetIndex(mp, reflect2.PtrOf(name), reflect2.PtrOf(&v))
		}
	} else {
		for _, name := range structInfo.names {
			v := dec.decodeInterface(dec.NextByte())
			valdec.t.UnsafeSetIndex(mp, reflect2.PtrOf(name), reflect2.PtrOf(&v))
		}
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
	case TagMap:
		valdec.decodeMap(dec, p)
	case TagEmpty:
		valdec.t.UnsafeSet(reflect2.PtrOf(p), valdec.t.UnsafeMakeMap(0))
	case TagList:
		valdec.decodeListAsMap(dec, p, tag)
	case TagClass:
		dec.ReadStruct()
		dec.Decode(p)
	case TagObject:
		valdec.decodeObjectAsMap(dec, p, tag)
	default:
		dec.decodeError(valdec.t.Type1(), tag)
	}
}

func (valdec mapDecoder) Type() reflect.Type {
	return valdec.t.Type1()
}

// makeMapDecoder returns a mapDecoder for map[K]V.
func makeMapDecoder(t reflect.Type) mapDecoder {
	mt := reflect2.Type2(t).(*reflect2.UnsafeMapType)
	kt := t.Key()
	vt := t.Elem()
	return mapDecoder{
		mt,
		reflect2.Type2(kt),
		reflect2.Type2(vt),
		GetDecodeHandler(kt),
		GetDecodeHandler(vt),
	}
}

func getMapDecoder(t reflect.Type) ValueDecoder {
	return makeMapDecoder(t)
}

//nolint
func fastDecodeIntMap(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[int]string:
		ismdec.Decode(dec, p, tag)
	case *map[int]interface{}:
		iifmdec.Decode(dec, p, tag)
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
	default:
		return false
	}
	return true
}

//nolint
func fastDecodeInt8Map(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[int8]string:
		i8smdec.Decode(dec, p, tag)
	case *map[int8]interface{}:
		i8ifmdec.Decode(dec, p, tag)
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
	default:
		return false
	}
	return true
}

//nolint
func fastDecodeInt16Map(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[int16]string:
		i16smdec.Decode(dec, p, tag)
	case *map[int16]interface{}:
		i16ifmdec.Decode(dec, p, tag)
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
	default:
		return false
	}
	return true
}

//nolint
func fastDecodeInt32Map(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[int32]string:
		i32smdec.Decode(dec, p, tag)
	case *map[int32]interface{}:
		i32ifmdec.Decode(dec, p, tag)
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
	default:
		return false
	}
	return true
}

//nolint
func fastDecodeInt64Map(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[int64]string:
		i64smdec.Decode(dec, p, tag)
	case *map[int64]interface{}:
		i64ifmdec.Decode(dec, p, tag)
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
	default:
		return false
	}
	return true
}

//nolint
func fastDecodeUintMap(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[uint]string:
		usmdec.Decode(dec, p, tag)
	case *map[uint]interface{}:
		uifmdec.Decode(dec, p, tag)
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
	default:
		return false
	}
	return true
}

//nolint
func fastDecodeUint8Map(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[uint8]string:
		u8smdec.Decode(dec, p, tag)
	case *map[uint8]interface{}:
		u8ifmdec.Decode(dec, p, tag)
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
	default:
		return false
	}
	return true
}

//nolint
func fastDecodeUint16Map(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[uint16]string:
		u16smdec.Decode(dec, p, tag)
	case *map[uint16]interface{}:
		u16ifmdec.Decode(dec, p, tag)
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
	default:
		return false
	}
	return true
}

//nolint
func fastDecodeUint32Map(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[uint32]string:
		u32smdec.Decode(dec, p, tag)
	case *map[uint32]interface{}:
		u32ifmdec.Decode(dec, p, tag)
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
	default:
		return false
	}
	return true
}

//nolint
func fastDecodeUint64Map(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[uint64]string:
		u64smdec.Decode(dec, p, tag)
	case *map[uint64]interface{}:
		u64ifmdec.Decode(dec, p, tag)
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
	default:
		return false
	}
	return true
}

//nolint
func fastDecodeFloat32Map(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[float32]string:
		f32smdec.Decode(dec, p, tag)
	case *map[float32]interface{}:
		f32ifmdec.Decode(dec, p, tag)
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
	default:
		return false
	}
	return true
}

//nolint
func fastDecodeFloat64Map(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[float64]string:
		f64smdec.Decode(dec, p, tag)
	case *map[float64]interface{}:
		f64ifmdec.Decode(dec, p, tag)
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
	default:
		return false
	}
	return true
}

//nolint
func fastDecodeStringMap(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[string]string:
		ssmdec.Decode(dec, p, tag)
	case *map[string]interface{}:
		sifmdec.Decode(dec, p, tag)
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
	default:
		return false
	}
	return true
}

//nolint
func fastDecodeInterfaceMap(dec *Decoder, p interface{}, tag byte) bool {
	switch p.(type) {
	case *map[interface{}]string:
		ifsmdec.Decode(dec, p, tag)
	case *map[interface{}]interface{}:
		ififmdec.Decode(dec, p, tag)
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
	default:
		return false
	}
	return true
}

func (dec *Decoder) fastDecodeMap(t reflect.Type, p interface{}, tag byte) bool {
	switch t.Key().Kind() {
	case reflect.String:
		return fastDecodeStringMap(dec, p, tag)
	case reflect.Interface:
		return fastDecodeInterfaceMap(dec, p, tag)
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
	}
	return false
}

var (
	ismdec   mapDecoder
	iifmdec  mapDecoder
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

	i8smdec   mapDecoder
	i8ifmdec  mapDecoder
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

	i16smdec   mapDecoder
	i16ifmdec  mapDecoder
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

	i32smdec   mapDecoder
	i32ifmdec  mapDecoder
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

	i64smdec   mapDecoder
	i64ifmdec  mapDecoder
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

	usmdec   mapDecoder
	uifmdec  mapDecoder
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

	u8smdec   mapDecoder
	u8ifmdec  mapDecoder
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

	u16smdec   mapDecoder
	u16ifmdec  mapDecoder
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

	u32smdec   mapDecoder
	u32ifmdec  mapDecoder
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

	u64smdec   mapDecoder
	u64ifmdec  mapDecoder
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

	f32smdec   mapDecoder
	f32ifmdec  mapDecoder
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

	f64smdec   mapDecoder
	f64ifmdec  mapDecoder
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

	ssmdec   mapDecoder
	sifmdec  mapDecoder
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

	ifsmdec   mapDecoder
	ififmdec  mapDecoder
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
)

func init() {
	ismdec = makeMapDecoder(reflect.TypeOf((map[int]string)(nil)))
	iifmdec = makeMapDecoder(reflect.TypeOf((map[int]interface{})(nil)))
	ibmdec = makeMapDecoder(reflect.TypeOf((map[int]bool)(nil)))
	iimdec = makeMapDecoder(reflect.TypeOf((map[int]int)(nil)))
	ii8mdec = makeMapDecoder(reflect.TypeOf((map[int]int8)(nil)))
	ii16mdec = makeMapDecoder(reflect.TypeOf((map[int]int16)(nil)))
	ii32mdec = makeMapDecoder(reflect.TypeOf((map[int]int32)(nil)))
	ii64mdec = makeMapDecoder(reflect.TypeOf((map[int]int64)(nil)))
	iumdec = makeMapDecoder(reflect.TypeOf((map[int]uint)(nil)))
	iu8mdec = makeMapDecoder(reflect.TypeOf((map[int]uint8)(nil)))
	iu16mdec = makeMapDecoder(reflect.TypeOf((map[int]uint16)(nil)))
	iu32mdec = makeMapDecoder(reflect.TypeOf((map[int]uint32)(nil)))
	iu64mdec = makeMapDecoder(reflect.TypeOf((map[int]uint64)(nil)))
	if32mdec = makeMapDecoder(reflect.TypeOf((map[int]float32)(nil)))
	if64mdec = makeMapDecoder(reflect.TypeOf((map[int]float64)(nil)))

	i8smdec = makeMapDecoder(reflect.TypeOf((map[int8]string)(nil)))
	i8ifmdec = makeMapDecoder(reflect.TypeOf((map[int8]interface{})(nil)))
	i8bmdec = makeMapDecoder(reflect.TypeOf((map[int8]bool)(nil)))
	i8imdec = makeMapDecoder(reflect.TypeOf((map[int8]int)(nil)))
	i8i8mdec = makeMapDecoder(reflect.TypeOf((map[int8]int8)(nil)))
	i8i16mdec = makeMapDecoder(reflect.TypeOf((map[int8]int16)(nil)))
	i8i32mdec = makeMapDecoder(reflect.TypeOf((map[int8]int32)(nil)))
	i8i64mdec = makeMapDecoder(reflect.TypeOf((map[int8]int64)(nil)))
	i8umdec = makeMapDecoder(reflect.TypeOf((map[int8]uint)(nil)))
	i8u8mdec = makeMapDecoder(reflect.TypeOf((map[int8]uint8)(nil)))
	i8u16mdec = makeMapDecoder(reflect.TypeOf((map[int8]uint16)(nil)))
	i8u32mdec = makeMapDecoder(reflect.TypeOf((map[int8]uint32)(nil)))
	i8u64mdec = makeMapDecoder(reflect.TypeOf((map[int8]uint64)(nil)))
	i8f32mdec = makeMapDecoder(reflect.TypeOf((map[int8]float32)(nil)))
	i8f64mdec = makeMapDecoder(reflect.TypeOf((map[int8]float64)(nil)))

	i16smdec = makeMapDecoder(reflect.TypeOf((map[int16]string)(nil)))
	i16ifmdec = makeMapDecoder(reflect.TypeOf((map[int16]interface{})(nil)))
	i16bmdec = makeMapDecoder(reflect.TypeOf((map[int16]bool)(nil)))
	i16imdec = makeMapDecoder(reflect.TypeOf((map[int16]int)(nil)))
	i16i8mdec = makeMapDecoder(reflect.TypeOf((map[int16]int8)(nil)))
	i16i16mdec = makeMapDecoder(reflect.TypeOf((map[int16]int16)(nil)))
	i16i32mdec = makeMapDecoder(reflect.TypeOf((map[int16]int32)(nil)))
	i16i64mdec = makeMapDecoder(reflect.TypeOf((map[int16]int64)(nil)))
	i16umdec = makeMapDecoder(reflect.TypeOf((map[int16]uint)(nil)))
	i16u8mdec = makeMapDecoder(reflect.TypeOf((map[int16]uint8)(nil)))
	i16u16mdec = makeMapDecoder(reflect.TypeOf((map[int16]uint16)(nil)))
	i16u32mdec = makeMapDecoder(reflect.TypeOf((map[int16]uint32)(nil)))
	i16u64mdec = makeMapDecoder(reflect.TypeOf((map[int16]uint64)(nil)))
	i16f32mdec = makeMapDecoder(reflect.TypeOf((map[int16]float32)(nil)))
	i16f64mdec = makeMapDecoder(reflect.TypeOf((map[int16]float64)(nil)))

	i32smdec = makeMapDecoder(reflect.TypeOf((map[int32]string)(nil)))
	i32ifmdec = makeMapDecoder(reflect.TypeOf((map[int32]interface{})(nil)))
	i32bmdec = makeMapDecoder(reflect.TypeOf((map[int32]bool)(nil)))
	i32imdec = makeMapDecoder(reflect.TypeOf((map[int32]int)(nil)))
	i32i8mdec = makeMapDecoder(reflect.TypeOf((map[int32]int8)(nil)))
	i32i16mdec = makeMapDecoder(reflect.TypeOf((map[int32]int16)(nil)))
	i32i32mdec = makeMapDecoder(reflect.TypeOf((map[int32]int32)(nil)))
	i32i64mdec = makeMapDecoder(reflect.TypeOf((map[int32]int64)(nil)))
	i32umdec = makeMapDecoder(reflect.TypeOf((map[int32]uint)(nil)))
	i32u8mdec = makeMapDecoder(reflect.TypeOf((map[int32]uint8)(nil)))
	i32u16mdec = makeMapDecoder(reflect.TypeOf((map[int32]uint16)(nil)))
	i32u32mdec = makeMapDecoder(reflect.TypeOf((map[int32]uint32)(nil)))
	i32u64mdec = makeMapDecoder(reflect.TypeOf((map[int32]uint64)(nil)))
	i32f32mdec = makeMapDecoder(reflect.TypeOf((map[int32]float32)(nil)))
	i32f64mdec = makeMapDecoder(reflect.TypeOf((map[int32]float64)(nil)))

	i64smdec = makeMapDecoder(reflect.TypeOf((map[int64]string)(nil)))
	i64ifmdec = makeMapDecoder(reflect.TypeOf((map[int64]interface{})(nil)))
	i64bmdec = makeMapDecoder(reflect.TypeOf((map[int64]bool)(nil)))
	i64imdec = makeMapDecoder(reflect.TypeOf((map[int64]int)(nil)))
	i64i8mdec = makeMapDecoder(reflect.TypeOf((map[int64]int8)(nil)))
	i64i16mdec = makeMapDecoder(reflect.TypeOf((map[int64]int16)(nil)))
	i64i32mdec = makeMapDecoder(reflect.TypeOf((map[int64]int32)(nil)))
	i64i64mdec = makeMapDecoder(reflect.TypeOf((map[int64]int64)(nil)))
	i64umdec = makeMapDecoder(reflect.TypeOf((map[int64]uint)(nil)))
	i64u8mdec = makeMapDecoder(reflect.TypeOf((map[int64]uint8)(nil)))
	i64u16mdec = makeMapDecoder(reflect.TypeOf((map[int64]uint16)(nil)))
	i64u32mdec = makeMapDecoder(reflect.TypeOf((map[int64]uint32)(nil)))
	i64u64mdec = makeMapDecoder(reflect.TypeOf((map[int64]uint64)(nil)))
	i64f32mdec = makeMapDecoder(reflect.TypeOf((map[int64]float32)(nil)))
	i64f64mdec = makeMapDecoder(reflect.TypeOf((map[int64]float64)(nil)))

	usmdec = makeMapDecoder(reflect.TypeOf((map[uint]string)(nil)))
	uifmdec = makeMapDecoder(reflect.TypeOf((map[uint]interface{})(nil)))
	ubmdec = makeMapDecoder(reflect.TypeOf((map[uint]bool)(nil)))
	uimdec = makeMapDecoder(reflect.TypeOf((map[uint]int)(nil)))
	ui8mdec = makeMapDecoder(reflect.TypeOf((map[uint]int8)(nil)))
	ui16mdec = makeMapDecoder(reflect.TypeOf((map[uint]int16)(nil)))
	ui32mdec = makeMapDecoder(reflect.TypeOf((map[uint]int32)(nil)))
	ui64mdec = makeMapDecoder(reflect.TypeOf((map[uint]int64)(nil)))
	uumdec = makeMapDecoder(reflect.TypeOf((map[uint]uint)(nil)))
	uu8mdec = makeMapDecoder(reflect.TypeOf((map[uint]uint8)(nil)))
	uu16mdec = makeMapDecoder(reflect.TypeOf((map[uint]uint16)(nil)))
	uu32mdec = makeMapDecoder(reflect.TypeOf((map[uint]uint32)(nil)))
	uu64mdec = makeMapDecoder(reflect.TypeOf((map[uint]uint64)(nil)))
	uf32mdec = makeMapDecoder(reflect.TypeOf((map[uint]float32)(nil)))
	uf64mdec = makeMapDecoder(reflect.TypeOf((map[uint]float64)(nil)))

	u8smdec = makeMapDecoder(reflect.TypeOf((map[uint8]string)(nil)))
	u8ifmdec = makeMapDecoder(reflect.TypeOf((map[uint8]interface{})(nil)))
	u8bmdec = makeMapDecoder(reflect.TypeOf((map[uint8]bool)(nil)))
	u8imdec = makeMapDecoder(reflect.TypeOf((map[uint8]int)(nil)))
	u8i8mdec = makeMapDecoder(reflect.TypeOf((map[uint8]int8)(nil)))
	u8i16mdec = makeMapDecoder(reflect.TypeOf((map[uint8]int16)(nil)))
	u8i32mdec = makeMapDecoder(reflect.TypeOf((map[uint8]int32)(nil)))
	u8i64mdec = makeMapDecoder(reflect.TypeOf((map[uint8]int64)(nil)))
	u8umdec = makeMapDecoder(reflect.TypeOf((map[uint8]uint)(nil)))
	u8u8mdec = makeMapDecoder(reflect.TypeOf((map[uint8]uint8)(nil)))
	u8u16mdec = makeMapDecoder(reflect.TypeOf((map[uint8]uint16)(nil)))
	u8u32mdec = makeMapDecoder(reflect.TypeOf((map[uint8]uint32)(nil)))
	u8u64mdec = makeMapDecoder(reflect.TypeOf((map[uint8]uint64)(nil)))
	u8f32mdec = makeMapDecoder(reflect.TypeOf((map[uint8]float32)(nil)))
	u8f64mdec = makeMapDecoder(reflect.TypeOf((map[uint8]float64)(nil)))

	u16smdec = makeMapDecoder(reflect.TypeOf((map[uint16]string)(nil)))
	u16ifmdec = makeMapDecoder(reflect.TypeOf((map[uint16]interface{})(nil)))
	u16bmdec = makeMapDecoder(reflect.TypeOf((map[uint16]bool)(nil)))
	u16imdec = makeMapDecoder(reflect.TypeOf((map[uint16]int)(nil)))
	u16i8mdec = makeMapDecoder(reflect.TypeOf((map[uint16]int8)(nil)))
	u16i16mdec = makeMapDecoder(reflect.TypeOf((map[uint16]int16)(nil)))
	u16i32mdec = makeMapDecoder(reflect.TypeOf((map[uint16]int32)(nil)))
	u16i64mdec = makeMapDecoder(reflect.TypeOf((map[uint16]int64)(nil)))
	u16umdec = makeMapDecoder(reflect.TypeOf((map[uint16]uint)(nil)))
	u16u8mdec = makeMapDecoder(reflect.TypeOf((map[uint16]uint8)(nil)))
	u16u16mdec = makeMapDecoder(reflect.TypeOf((map[uint16]uint16)(nil)))
	u16u32mdec = makeMapDecoder(reflect.TypeOf((map[uint16]uint32)(nil)))
	u16u64mdec = makeMapDecoder(reflect.TypeOf((map[uint16]uint64)(nil)))
	u16f32mdec = makeMapDecoder(reflect.TypeOf((map[uint16]float32)(nil)))
	u16f64mdec = makeMapDecoder(reflect.TypeOf((map[uint16]float64)(nil)))

	u32smdec = makeMapDecoder(reflect.TypeOf((map[uint32]string)(nil)))
	u32ifmdec = makeMapDecoder(reflect.TypeOf((map[uint32]interface{})(nil)))
	u32bmdec = makeMapDecoder(reflect.TypeOf((map[uint32]bool)(nil)))
	u32imdec = makeMapDecoder(reflect.TypeOf((map[uint32]int)(nil)))
	u32i8mdec = makeMapDecoder(reflect.TypeOf((map[uint32]int8)(nil)))
	u32i16mdec = makeMapDecoder(reflect.TypeOf((map[uint32]int16)(nil)))
	u32i32mdec = makeMapDecoder(reflect.TypeOf((map[uint32]int32)(nil)))
	u32i64mdec = makeMapDecoder(reflect.TypeOf((map[uint32]int64)(nil)))
	u32umdec = makeMapDecoder(reflect.TypeOf((map[uint32]uint)(nil)))
	u32u8mdec = makeMapDecoder(reflect.TypeOf((map[uint32]uint8)(nil)))
	u32u16mdec = makeMapDecoder(reflect.TypeOf((map[uint32]uint16)(nil)))
	u32u32mdec = makeMapDecoder(reflect.TypeOf((map[uint32]uint32)(nil)))
	u32u64mdec = makeMapDecoder(reflect.TypeOf((map[uint32]uint64)(nil)))
	u32f32mdec = makeMapDecoder(reflect.TypeOf((map[uint32]float32)(nil)))
	u32f64mdec = makeMapDecoder(reflect.TypeOf((map[uint32]float64)(nil)))

	u64smdec = makeMapDecoder(reflect.TypeOf((map[uint64]string)(nil)))
	u64ifmdec = makeMapDecoder(reflect.TypeOf((map[uint64]interface{})(nil)))
	u64bmdec = makeMapDecoder(reflect.TypeOf((map[uint64]bool)(nil)))
	u64imdec = makeMapDecoder(reflect.TypeOf((map[uint64]int)(nil)))
	u64i8mdec = makeMapDecoder(reflect.TypeOf((map[uint64]int8)(nil)))
	u64i16mdec = makeMapDecoder(reflect.TypeOf((map[uint64]int16)(nil)))
	u64i32mdec = makeMapDecoder(reflect.TypeOf((map[uint64]int32)(nil)))
	u64i64mdec = makeMapDecoder(reflect.TypeOf((map[uint64]int64)(nil)))
	u64umdec = makeMapDecoder(reflect.TypeOf((map[uint64]uint)(nil)))
	u64u8mdec = makeMapDecoder(reflect.TypeOf((map[uint64]uint8)(nil)))
	u64u16mdec = makeMapDecoder(reflect.TypeOf((map[uint64]uint16)(nil)))
	u64u32mdec = makeMapDecoder(reflect.TypeOf((map[uint64]uint32)(nil)))
	u64u64mdec = makeMapDecoder(reflect.TypeOf((map[uint64]uint64)(nil)))
	u64f32mdec = makeMapDecoder(reflect.TypeOf((map[uint64]float32)(nil)))
	u64f64mdec = makeMapDecoder(reflect.TypeOf((map[uint64]float64)(nil)))

	f32smdec = makeMapDecoder(reflect.TypeOf((map[float32]string)(nil)))
	f32ifmdec = makeMapDecoder(reflect.TypeOf((map[float32]interface{})(nil)))
	f32bmdec = makeMapDecoder(reflect.TypeOf((map[float32]bool)(nil)))
	f32imdec = makeMapDecoder(reflect.TypeOf((map[float32]int)(nil)))
	f32i8mdec = makeMapDecoder(reflect.TypeOf((map[float32]int8)(nil)))
	f32i16mdec = makeMapDecoder(reflect.TypeOf((map[float32]int16)(nil)))
	f32i32mdec = makeMapDecoder(reflect.TypeOf((map[float32]int32)(nil)))
	f32i64mdec = makeMapDecoder(reflect.TypeOf((map[float32]int64)(nil)))
	f32umdec = makeMapDecoder(reflect.TypeOf((map[float32]uint)(nil)))
	f32u8mdec = makeMapDecoder(reflect.TypeOf((map[float32]uint8)(nil)))
	f32u16mdec = makeMapDecoder(reflect.TypeOf((map[float32]uint16)(nil)))
	f32u32mdec = makeMapDecoder(reflect.TypeOf((map[float32]uint32)(nil)))
	f32u64mdec = makeMapDecoder(reflect.TypeOf((map[float32]uint64)(nil)))
	f32f32mdec = makeMapDecoder(reflect.TypeOf((map[float32]float32)(nil)))
	f32f64mdec = makeMapDecoder(reflect.TypeOf((map[float32]float64)(nil)))

	f64smdec = makeMapDecoder(reflect.TypeOf((map[float64]string)(nil)))
	f64ifmdec = makeMapDecoder(reflect.TypeOf((map[float64]interface{})(nil)))
	f64bmdec = makeMapDecoder(reflect.TypeOf((map[float64]bool)(nil)))
	f64imdec = makeMapDecoder(reflect.TypeOf((map[float64]int)(nil)))
	f64i8mdec = makeMapDecoder(reflect.TypeOf((map[float64]int8)(nil)))
	f64i16mdec = makeMapDecoder(reflect.TypeOf((map[float64]int16)(nil)))
	f64i32mdec = makeMapDecoder(reflect.TypeOf((map[float64]int32)(nil)))
	f64i64mdec = makeMapDecoder(reflect.TypeOf((map[float64]int64)(nil)))
	f64umdec = makeMapDecoder(reflect.TypeOf((map[float64]uint)(nil)))
	f64u8mdec = makeMapDecoder(reflect.TypeOf((map[float64]uint8)(nil)))
	f64u16mdec = makeMapDecoder(reflect.TypeOf((map[float64]uint16)(nil)))
	f64u32mdec = makeMapDecoder(reflect.TypeOf((map[float64]uint32)(nil)))
	f64u64mdec = makeMapDecoder(reflect.TypeOf((map[float64]uint64)(nil)))
	f64f32mdec = makeMapDecoder(reflect.TypeOf((map[float64]float32)(nil)))
	f64f64mdec = makeMapDecoder(reflect.TypeOf((map[float64]float64)(nil)))

	ssmdec = makeMapDecoder(reflect.TypeOf((map[string]string)(nil)))
	sifmdec = makeMapDecoder(reflect.TypeOf((map[string]interface{})(nil)))
	sbmdec = makeMapDecoder(reflect.TypeOf((map[string]bool)(nil)))
	simdec = makeMapDecoder(reflect.TypeOf((map[string]int)(nil)))
	si8mdec = makeMapDecoder(reflect.TypeOf((map[string]int8)(nil)))
	si16mdec = makeMapDecoder(reflect.TypeOf((map[string]int16)(nil)))
	si32mdec = makeMapDecoder(reflect.TypeOf((map[string]int32)(nil)))
	si64mdec = makeMapDecoder(reflect.TypeOf((map[string]int64)(nil)))
	sumdec = makeMapDecoder(reflect.TypeOf((map[string]uint)(nil)))
	su8mdec = makeMapDecoder(reflect.TypeOf((map[string]uint8)(nil)))
	su16mdec = makeMapDecoder(reflect.TypeOf((map[string]uint16)(nil)))
	su32mdec = makeMapDecoder(reflect.TypeOf((map[string]uint32)(nil)))
	su64mdec = makeMapDecoder(reflect.TypeOf((map[string]uint64)(nil)))
	sf32mdec = makeMapDecoder(reflect.TypeOf((map[string]float32)(nil)))
	sf64mdec = makeMapDecoder(reflect.TypeOf((map[string]float64)(nil)))

	ifsmdec = makeMapDecoder(reflect.TypeOf((map[interface{}]string)(nil)))
	ififmdec = makeMapDecoder(reflect.TypeOf((map[interface{}]interface{})(nil)))
	ifbmdec = makeMapDecoder(reflect.TypeOf((map[interface{}]bool)(nil)))
	ifimdec = makeMapDecoder(reflect.TypeOf((map[interface{}]int)(nil)))
	ifi8mdec = makeMapDecoder(reflect.TypeOf((map[interface{}]int8)(nil)))
	ifi16mdec = makeMapDecoder(reflect.TypeOf((map[interface{}]int16)(nil)))
	ifi32mdec = makeMapDecoder(reflect.TypeOf((map[interface{}]int32)(nil)))
	ifi64mdec = makeMapDecoder(reflect.TypeOf((map[interface{}]int64)(nil)))
	ifumdec = makeMapDecoder(reflect.TypeOf((map[interface{}]uint)(nil)))
	ifu8mdec = makeMapDecoder(reflect.TypeOf((map[interface{}]uint8)(nil)))
	ifu16mdec = makeMapDecoder(reflect.TypeOf((map[interface{}]uint16)(nil)))
	ifu32mdec = makeMapDecoder(reflect.TypeOf((map[interface{}]uint32)(nil)))
	ifu64mdec = makeMapDecoder(reflect.TypeOf((map[interface{}]uint64)(nil)))
	iff32mdec = makeMapDecoder(reflect.TypeOf((map[interface{}]float32)(nil)))
	iff64mdec = makeMapDecoder(reflect.TypeOf((map[interface{}]float64)(nil)))

	RegisterValueDecoder(ismdec)
	RegisterValueDecoder(iifmdec)
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

	RegisterValueDecoder(i8smdec)
	RegisterValueDecoder(i8ifmdec)
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

	RegisterValueDecoder(i16smdec)
	RegisterValueDecoder(i16ifmdec)
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

	RegisterValueDecoder(i32smdec)
	RegisterValueDecoder(i32ifmdec)
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

	RegisterValueDecoder(i64smdec)
	RegisterValueDecoder(i64ifmdec)
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

	RegisterValueDecoder(usmdec)
	RegisterValueDecoder(uifmdec)
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

	RegisterValueDecoder(u8smdec)
	RegisterValueDecoder(u8ifmdec)
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

	RegisterValueDecoder(u16smdec)
	RegisterValueDecoder(u16ifmdec)
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

	RegisterValueDecoder(u32smdec)
	RegisterValueDecoder(u32ifmdec)
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

	RegisterValueDecoder(u64smdec)
	RegisterValueDecoder(u64ifmdec)
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

	RegisterValueDecoder(f32smdec)
	RegisterValueDecoder(f32ifmdec)
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

	RegisterValueDecoder(f64smdec)
	RegisterValueDecoder(f64ifmdec)
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

	RegisterValueDecoder(ssmdec)
	RegisterValueDecoder(sifmdec)
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

	RegisterValueDecoder(ifsmdec)
	RegisterValueDecoder(ififmdec)
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
}
