/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/struct_encoder.go                            |
|                                                          |
| LastModified: Mar 21, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"reflect"

	"github.com/modern-go/reflect2"
)

// EncodeHandler is an encode handler
type EncodeHandler func(enc *Encoder, v interface{}) error

// GetEncodeHandler for specified type
func GetEncodeHandler(t reflect.Type) EncodeHandler {
	if f := getOtherEncodeHandler(t); f != nil {
		return f
	}
	switch t.Kind() {
	case reflect.Int:
		return intEncode
	case reflect.Int8:
		return int8Encode
	case reflect.Int16:
		return int16Encode
	case reflect.Int32:
		return int32Encode
	case reflect.Int64:
		return int64Encode
	case reflect.Uint:
		return uintEncode
	case reflect.Uint8:
		return uint8Encode
	case reflect.Uint16:
		return uint16Encode
	case reflect.Uint32:
		return uint32Encode
	case reflect.Uint64, reflect.Uintptr:
		return uint64Encode
	case reflect.Bool:
		return boolEncode
	case reflect.Float32:
		return float32Encode
	case reflect.Float64:
		return float64Encode
	case reflect.Complex64:
		return complex64Encode
	case reflect.Complex128:
		return complex128Encode
	case reflect.Array:
		return arrayEncode
	case reflect.Interface:
		return interfaceEncode
	case reflect.Map:
		return mapEncode
	case reflect.Ptr:
		return getPtrEncodeHandler(t.Elem())
	case reflect.Slice:
		return sliceEncode
	case reflect.String:
		return stringEncode
	case reflect.Struct:
		return getStructEncodeHandler(t)
	}
	return nil
}

func boolEncode(enc *Encoder, v interface{}) error {
	return WriteBool(enc.Writer, *(*bool)(reflect2.PtrOf(v)))
}

func intEncode(enc *Encoder, v interface{}) error {
	return WriteInt(enc.Writer, *(*int)(reflect2.PtrOf(v)))
}

func int8Encode(enc *Encoder, v interface{}) error {
	return WriteInt8(enc.Writer, *(*int8)(reflect2.PtrOf(v)))
}

func int16Encode(enc *Encoder, v interface{}) error {
	return WriteInt16(enc.Writer, *(*int16)(reflect2.PtrOf(v)))
}

func int32Encode(enc *Encoder, v interface{}) error {
	return WriteInt32(enc.Writer, *(*int32)(reflect2.PtrOf(v)))
}

func int64Encode(enc *Encoder, v interface{}) error {
	return WriteInt64(enc.Writer, *(*int64)(reflect2.PtrOf(v)))
}

func uintEncode(enc *Encoder, v interface{}) error {
	return WriteUint(enc.Writer, *(*uint)(reflect2.PtrOf(v)))
}

func uint8Encode(enc *Encoder, v interface{}) error {
	return WriteUint8(enc.Writer, *(*uint8)(reflect2.PtrOf(v)))
}

func uint16Encode(enc *Encoder, v interface{}) error {
	return WriteUint16(enc.Writer, *(*uint16)(reflect2.PtrOf(v)))
}

func uint32Encode(enc *Encoder, v interface{}) error {
	return WriteUint32(enc.Writer, *(*uint32)(reflect2.PtrOf(v)))
}

func uint64Encode(enc *Encoder, v interface{}) error {
	return WriteUint64(enc.Writer, *(*uint64)(reflect2.PtrOf(v)))
}

func float32Encode(enc *Encoder, v interface{}) error {
	return WriteFloat32(enc.Writer, *(*float32)(reflect2.PtrOf(v)))
}

func float64Encode(enc *Encoder, v interface{}) error {
	return WriteFloat64(enc.Writer, *(*float64)(reflect2.PtrOf(v)))
}

func complex64Encode(enc *Encoder, v interface{}) error {
	return WriteComplex64(enc, *(*complex64)(reflect2.PtrOf(v)))
}

func complex128Encode(enc *Encoder, v interface{}) error {
	return WriteComplex128(enc, *(*complex128)(reflect2.PtrOf(v)))
}

func stringEncode(enc *Encoder, v interface{}) error {
	return EncodeString(enc, *(*string)(reflect2.PtrOf(v)))
}

func arrayEncode(enc *Encoder, v interface{}) error {
	return WriteArray(enc, v)
}

func mapEncode(enc *Encoder, v interface{}) error {
	if reflect.ValueOf(v).IsNil() {
		return WriteNil(enc.Writer)
	}
	return WriteMap(enc, v)
}

func sliceEncode(enc *Encoder, v interface{}) error {
	if reflect.ValueOf(v).IsNil() {
		return WriteNil(enc.Writer)
	}
	return WriteSlice(enc, v)
}

func interfaceEncode(enc *Encoder, v interface{}) error {
	return enc.Encode(v)
}

func boolPtrEncode(enc *Encoder, v interface{}) error {
	p := (*bool)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteBool(enc.Writer, *p)
}

func intPtrEncode(enc *Encoder, v interface{}) error {
	p := (*int)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteInt(enc.Writer, *p)
}

func int8PtrEncode(enc *Encoder, v interface{}) error {
	p := (*int8)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteInt8(enc.Writer, *p)
}

func int16PtrEncode(enc *Encoder, v interface{}) error {
	p := (*int16)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteInt16(enc.Writer, *p)
}

func int32PtrEncode(enc *Encoder, v interface{}) error {
	p := (*int32)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteInt32(enc.Writer, *p)
}

func int64PtrEncode(enc *Encoder, v interface{}) error {
	p := (*int64)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteInt64(enc.Writer, *p)
}

func uintPtrEncode(enc *Encoder, v interface{}) error {
	p := (*uint)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteUint(enc.Writer, *p)
}

func uint8PtrEncode(enc *Encoder, v interface{}) error {
	p := (*uint8)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteUint8(enc.Writer, *p)
}

func uint16PtrEncode(enc *Encoder, v interface{}) error {
	p := (*uint16)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteUint16(enc.Writer, *p)
}

func uint32PtrEncode(enc *Encoder, v interface{}) error {
	p := (*uint32)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteUint32(enc.Writer, *p)
}

func uint64PtrEncode(enc *Encoder, v interface{}) error {
	p := (*uint64)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteUint64(enc.Writer, *p)
}

func float32PtrEncode(enc *Encoder, v interface{}) error {
	p := (*float32)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteFloat32(enc.Writer, *p)
}

func float64PtrEncode(enc *Encoder, v interface{}) error {
	p := (*float64)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteFloat64(enc.Writer, *p)
}

func complex64PtrEncode(enc *Encoder, v interface{}) error {
	p := (*complex64)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteComplex64(enc, *p)
}

func complex128PtrEncode(enc *Encoder, v interface{}) error {
	p := (*complex128)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return WriteComplex128(enc, *p)
}

func stringPtrEncode(enc *Encoder, v interface{}) error {
	p := (*string)(reflect2.PtrOf(v))
	if p == nil {
		return WriteNil(enc.Writer)
	}
	return EncodeString(enc, *p)
}

func arrayPtrEncode(enc *Encoder, v interface{}) error {
	if reflect2.IsNil(v) {
		return WriteNil(enc.Writer)
	}
	return arrayenc.Encode(enc, v)
}

func mapPtrEncode(enc *Encoder, v interface{}) error {
	if rv := reflect.ValueOf(v); rv.IsNil() || rv.Elem().IsNil() {
		return WriteNil(enc.Writer)
	}
	return mapenc.Encode(enc, v)
}

func slicePtrEncode(enc *Encoder, v interface{}) error {
	if rv := reflect.ValueOf(v); rv.IsNil() || rv.Elem().IsNil() {
		return WriteNil(enc.Writer)
	}
	return slcenc.Encode(enc, v)
}

func interfacePtrEncode(enc *Encoder, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return WriteNil(enc.Writer)
	}
	return enc.Encode(rv.Elem().Interface())
}

func ptrEncode(enc *Encoder, v interface{}) error {
	if reflect2.IsNil(v) {
		return WriteNil(enc.Writer)
	}
	return ptrenc.Encode(enc, v)
}

func getStructEncodeHandler(t reflect.Type) EncodeHandler {
	return getStructEncoder(t).Write
}

func getStructPtrEncodeHandler(t reflect.Type) EncodeHandler {
	return getStructEncoder(t).Encode
}

func getOtherEncodeHandler(t reflect.Type) EncodeHandler {
	if encoder := getOtherEncoder(t); encoder != nil {
		return encoder.Write
	}
	return nil
}

func getOtherPtrEncodeHandler(t reflect.Type) EncodeHandler {
	if encoder := getOtherEncoder(t); encoder != nil {
		return encoder.Encode
	}
	return nil
}

func getPtrEncodeHandler(t reflect.Type) EncodeHandler {
	if f := getOtherPtrEncodeHandler(t); f != nil {
		return f
	}
	switch t.Kind() {
	case reflect.Int:
		return intPtrEncode
	case reflect.Int8:
		return int8PtrEncode
	case reflect.Int16:
		return int16PtrEncode
	case reflect.Int32:
		return int32PtrEncode
	case reflect.Int64:
		return int64PtrEncode
	case reflect.Uint:
		return uintPtrEncode
	case reflect.Uint8:
		return uint8PtrEncode
	case reflect.Uint16:
		return uint16PtrEncode
	case reflect.Uint32:
		return uint32PtrEncode
	case reflect.Uint64, reflect.Uintptr:
		return uint64PtrEncode
	case reflect.Bool:
		return boolPtrEncode
	case reflect.Float32:
		return float32PtrEncode
	case reflect.Float64:
		return float64PtrEncode
	case reflect.Complex64:
		return complex64PtrEncode
	case reflect.Complex128:
		return complex128PtrEncode
	case reflect.Array:
		return arrayPtrEncode
	case reflect.Interface:
		return interfacePtrEncode
	case reflect.Map:
		return mapPtrEncode
	case reflect.Ptr:
		for t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		switch t.Kind() {
		case reflect.Func, reflect.Chan, reflect.UnsafePointer:
			return nil
		}
		return ptrEncode
	case reflect.Slice:
		return slicePtrEncode
	case reflect.String:
		return stringPtrEncode
	case reflect.Struct:
		return getStructPtrEncodeHandler(t)
	}
	return nil
}
