/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/map_decoder.go                                        |
|                                                          |
| LastModified: Jun 5, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

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
	case TagObject:
		valdec.decodeObjectAsMap(dec, p, tag)
	default:
		dec.defaultDecode(valdec.t.Type1(), p, tag)
	}
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

var (
	ififmdec mapDecoder
	sifmdec  mapDecoder
)

func init() {
	sifmdec = makeMapDecoder(reflect.TypeOf((map[string]interface{})(nil)))
	ififmdec = makeMapDecoder(reflect.TypeOf((map[interface{}]interface{})(nil)))

	registerValueDecoder(sifmdec.t.Type1(), sifmdec)
	registerValueDecoder(ififmdec.t.Type1(), ififmdec)
}
