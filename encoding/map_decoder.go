/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/map_decoder.go                                  |
|                                                          |
| LastModified: Jun 19, 2020                               |
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
}
