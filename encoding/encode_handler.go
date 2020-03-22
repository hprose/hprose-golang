/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/struct_encoder.go                               |
|                                                          |
| LastModified: Mar 22, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"reflect"

	"github.com/modern-go/reflect2"
)

// EncodeHandler is an encode handler
type EncodeHandler func(enc *Encoder, v interface{})

// GetEncodeHandler for specified type
func GetEncodeHandler(t reflect.Type) (handler EncodeHandler) {
	if handler = getOtherEncodeHandler(t); handler == nil {
		switch t.Kind() {
		case reflect.Int:
			handler = intEncode
		case reflect.Int8:
			handler = int8Encode
		case reflect.Int16:
			handler = int16Encode
		case reflect.Int32:
			handler = int32Encode
		case reflect.Int64:
			handler = int64Encode
		case reflect.Uint:
			handler = uintEncode
		case reflect.Uint8:
			handler = uint8Encode
		case reflect.Uint16:
			handler = uint16Encode
		case reflect.Uint32:
			handler = uint32Encode
		case reflect.Uint64, reflect.Uintptr:
			handler = uint64Encode
		case reflect.Bool:
			handler = boolEncode
		case reflect.Float32:
			handler = float32Encode
		case reflect.Float64:
			handler = float64Encode
		case reflect.Complex64:
			handler = complex64Encode
		case reflect.Complex128:
			handler = complex128Encode
		case reflect.Array:
			handler = arrayEncode
		case reflect.Interface:
			handler = interfaceEncode
		case reflect.Map:
			handler = mapEncode
		case reflect.Ptr:
			handler = getPtrEncodeHandler(t.Elem())
		case reflect.Slice:
			handler = sliceEncode
		case reflect.String:
			handler = stringEncode
		case reflect.Struct:
			handler = getStructEncodeHandler(t)
		}
	}
	return
}

func boolEncode(enc *Encoder, v interface{}) {
	WriteBool(enc, *(*bool)(reflect2.PtrOf(v)))
}

func intEncode(enc *Encoder, v interface{}) {
	WriteInt(enc, *(*int)(reflect2.PtrOf(v)))
}

func int8Encode(enc *Encoder, v interface{}) {
	WriteInt8(enc, *(*int8)(reflect2.PtrOf(v)))
}

func int16Encode(enc *Encoder, v interface{}) {
	WriteInt16(enc, *(*int16)(reflect2.PtrOf(v)))
}

func int32Encode(enc *Encoder, v interface{}) {
	WriteInt32(enc, *(*int32)(reflect2.PtrOf(v)))
}

func int64Encode(enc *Encoder, v interface{}) {
	WriteInt64(enc, *(*int64)(reflect2.PtrOf(v)))
}

func uintEncode(enc *Encoder, v interface{}) {
	WriteUint(enc, *(*uint)(reflect2.PtrOf(v)))
}

func uint8Encode(enc *Encoder, v interface{}) {
	WriteUint8(enc, *(*uint8)(reflect2.PtrOf(v)))
}

func uint16Encode(enc *Encoder, v interface{}) {
	WriteUint16(enc, *(*uint16)(reflect2.PtrOf(v)))
}

func uint32Encode(enc *Encoder, v interface{}) {
	WriteUint32(enc, *(*uint32)(reflect2.PtrOf(v)))
}

func uint64Encode(enc *Encoder, v interface{}) {
	WriteUint64(enc, *(*uint64)(reflect2.PtrOf(v)))
}

func float32Encode(enc *Encoder, v interface{}) {
	WriteFloat32(enc, *(*float32)(reflect2.PtrOf(v)))
}

func float64Encode(enc *Encoder, v interface{}) {
	WriteFloat64(enc, *(*float64)(reflect2.PtrOf(v)))
}

func complex64Encode(enc *Encoder, v interface{}) {
	WriteComplex64(enc, *(*complex64)(reflect2.PtrOf(v)))
}

func complex128Encode(enc *Encoder, v interface{}) {
	WriteComplex128(enc, *(*complex128)(reflect2.PtrOf(v)))
}

func stringEncode(enc *Encoder, v interface{}) {
	EncodeString(enc, *(*string)(reflect2.PtrOf(v)))
}

func arrayEncode(enc *Encoder, v interface{}) {
	WriteArray(enc, v)
}

func mapEncode(enc *Encoder, v interface{}) {
	if reflect.ValueOf(v).IsNil() {
		WriteNil(enc)
	} else {
		WriteMap(enc, v)
	}
}

func sliceEncode(enc *Encoder, v interface{}) {
	if reflect.ValueOf(v).IsNil() {
		WriteNil(enc)
	} else {
		WriteSlice(enc, v)
	}
}

func interfaceEncode(enc *Encoder, v interface{}) {
	enc.encode(v)
}

func boolPtrEncode(enc *Encoder, v interface{}) {
	p := (*bool)(reflect2.PtrOf(v))
	if p == nil {
		WriteNil(enc)
	} else {
		WriteBool(enc, *p)
	}
}

func intPtrEncode(enc *Encoder, v interface{}) {
	p := (*int)(reflect2.PtrOf(v))
	if p == nil {
		WriteNil(enc)
	} else {
		WriteInt(enc, *p)
	}
}

func int8PtrEncode(enc *Encoder, v interface{}) {
	p := (*int8)(reflect2.PtrOf(v))
	if p == nil {
		WriteNil(enc)
	} else {
		WriteInt8(enc, *p)
	}
}

func int16PtrEncode(enc *Encoder, v interface{}) {
	p := (*int16)(reflect2.PtrOf(v))
	if p == nil {
		WriteNil(enc)
	} else {
		WriteInt16(enc, *p)
	}
}

func int32PtrEncode(enc *Encoder, v interface{}) {
	p := (*int32)(reflect2.PtrOf(v))
	if p == nil {
		WriteNil(enc)
	} else {
		WriteInt32(enc, *p)
	}
}

func int64PtrEncode(enc *Encoder, v interface{}) {
	p := (*int64)(reflect2.PtrOf(v))
	if p == nil {
		WriteNil(enc)
	} else {
		WriteInt64(enc, *p)
	}
}

func uintPtrEncode(enc *Encoder, v interface{}) {
	p := (*uint)(reflect2.PtrOf(v))
	if p == nil {
		WriteNil(enc)
	} else {
		WriteUint(enc, *p)
	}
}

func uint8PtrEncode(enc *Encoder, v interface{}) {
	p := (*uint8)(reflect2.PtrOf(v))
	if p == nil {
		WriteNil(enc)
	} else {
		WriteUint8(enc, *p)
	}
}

func uint16PtrEncode(enc *Encoder, v interface{}) {
	p := (*uint16)(reflect2.PtrOf(v))
	if p == nil {
		WriteNil(enc)
	} else {
		WriteUint16(enc, *p)
	}
}

func uint32PtrEncode(enc *Encoder, v interface{}) {
	p := (*uint32)(reflect2.PtrOf(v))
	if p == nil {
		WriteNil(enc)
	} else {
		WriteUint32(enc, *p)
	}
}

func uint64PtrEncode(enc *Encoder, v interface{}) {
	p := (*uint64)(reflect2.PtrOf(v))
	if p == nil {
		WriteNil(enc)
	} else {
		WriteUint64(enc, *p)
	}
}

func float32PtrEncode(enc *Encoder, v interface{}) {
	p := (*float32)(reflect2.PtrOf(v))
	if p == nil {
		WriteNil(enc)
	} else {
		WriteFloat32(enc, *p)
	}
}

func float64PtrEncode(enc *Encoder, v interface{}) {
	p := (*float64)(reflect2.PtrOf(v))
	if p == nil {
		WriteNil(enc)
	} else {
		WriteFloat64(enc, *p)
	}
}

func complex64PtrEncode(enc *Encoder, v interface{}) {
	p := (*complex64)(reflect2.PtrOf(v))
	if p == nil {
		WriteNil(enc)
	} else {
		WriteComplex64(enc, *p)
	}
}

func complex128PtrEncode(enc *Encoder, v interface{}) {
	p := (*complex128)(reflect2.PtrOf(v))
	if p == nil {
		WriteNil(enc)
	} else {
		WriteComplex128(enc, *p)
	}
}

func stringPtrEncode(enc *Encoder, v interface{}) {
	p := (*string)(reflect2.PtrOf(v))
	if p == nil {
		WriteNil(enc)
	} else {
		EncodeString(enc, *p)
	}
}

func arrayPtrEncode(enc *Encoder, v interface{}) {
	if reflect2.IsNil(v) {
		WriteNil(enc)
	} else {
		arrayenc.Encode(enc, v)
	}
}

func mapPtrEncode(enc *Encoder, v interface{}) {
	if rv := reflect.ValueOf(v); rv.IsNil() || rv.Elem().IsNil() {
		WriteNil(enc)
	} else {
		mapenc.Encode(enc, v)
	}
}

func slicePtrEncode(enc *Encoder, v interface{}) {
	if rv := reflect.ValueOf(v); rv.IsNil() || rv.Elem().IsNil() {
		WriteNil(enc)
	} else {
		slcenc.Encode(enc, v)
	}
}

func interfacePtrEncode(enc *Encoder, v interface{}) {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		WriteNil(enc)
	} else {
		enc.encode(rv.Elem().Interface())
	}
}

func ptrEncode(enc *Encoder, v interface{}) {
	if reflect2.IsNil(v) {
		WriteNil(enc)
	} else {
		ptrenc.Encode(enc, v)
	}
}

func getStructEncodeHandler(t reflect.Type) EncodeHandler {
	return getStructEncoder(t).Write
}

func getStructPtrEncodeHandler(t reflect.Type) EncodeHandler {
	return getStructEncoder(t).Encode
}

func getOtherEncodeHandler(t reflect.Type) (handler EncodeHandler) {
	if encoder := getOtherEncoder(t); encoder != nil {
		handler = encoder.Write
	}
	return
}

func getOtherPtrEncodeHandler(t reflect.Type) (handler EncodeHandler) {
	if encoder := getOtherEncoder(t); encoder != nil {
		handler = encoder.Encode
	}
	return
}

func getPtrEncodeHandler(t reflect.Type) (handler EncodeHandler) {
	if handler = getOtherPtrEncodeHandler(t); handler == nil {
		switch t.Kind() {
		case reflect.Int:
			handler = intPtrEncode
		case reflect.Int8:
			handler = int8PtrEncode
		case reflect.Int16:
			handler = int16PtrEncode
		case reflect.Int32:
			handler = int32PtrEncode
		case reflect.Int64:
			handler = int64PtrEncode
		case reflect.Uint:
			handler = uintPtrEncode
		case reflect.Uint8:
			handler = uint8PtrEncode
		case reflect.Uint16:
			handler = uint16PtrEncode
		case reflect.Uint32:
			handler = uint32PtrEncode
		case reflect.Uint64, reflect.Uintptr:
			handler = uint64PtrEncode
		case reflect.Bool:
			handler = boolPtrEncode
		case reflect.Float32:
			handler = float32PtrEncode
		case reflect.Float64:
			handler = float64PtrEncode
		case reflect.Complex64:
			handler = complex64PtrEncode
		case reflect.Complex128:
			handler = complex128PtrEncode
		case reflect.Array:
			handler = arrayPtrEncode
		case reflect.Interface:
			handler = interfacePtrEncode
		case reflect.Map:
			handler = mapPtrEncode
		case reflect.Ptr:
			for t.Kind() == reflect.Ptr {
				t = t.Elem()
			}
			switch t.Kind() {
			case reflect.Func, reflect.Chan, reflect.UnsafePointer:
				handler = nil
			default:
				handler = ptrEncode
			}
		case reflect.Slice:
			handler = slicePtrEncode
		case reflect.String:
			handler = stringPtrEncode
		case reflect.Struct:
			handler = getStructPtrEncodeHandler(t)
		}
	}
	return handler
}
