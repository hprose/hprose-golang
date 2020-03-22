/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/ptr_encoder.go                                  |
|                                                          |
| LastModified: Mar 22, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"math/big"
	"reflect"

	"github.com/modern-go/reflect2"
)

// ptrEncoder is the implementation of ValueEncoder for ptr.
type ptrEncoder struct{}

var ptrenc ptrEncoder

func (ptrEncoder) Encode(enc *Encoder, v interface{}) {
	writePtr(enc, v, func(valenc ValueEncoder, enc *Encoder, v interface{}) {
		valenc.Encode(enc, v)
	})
}

func (ptrEncoder) Write(enc *Encoder, v interface{}) {
	writePtr(enc, v, func(valenc ValueEncoder, enc *Encoder, v interface{}) {
		valenc.Write(enc, v)
	})
}

func fastWritePtr(enc *Encoder, v interface{}) (ok bool) {
	ok = true
	switch v := v.(type) {
	case *int:
		WriteInt(enc, *v)
	case *int8:
		WriteInt8(enc, *v)
	case *int16:
		WriteInt16(enc, *v)
	case *int32:
		WriteInt32(enc, *v)
	case *int64:
		WriteInt64(enc, *v)
	case *uint:
		WriteUint(enc, *v)
	case *uint8:
		WriteUint8(enc, *v)
	case *uint16:
		WriteUint16(enc, *v)
	case *uint32:
		WriteUint32(enc, *v)
	case *uint64:
		WriteUint64(enc, *v)
	case *uintptr:
		WriteUint64(enc, uint64(*v))
	case *bool:
		WriteBool(enc, *v)
	case *float32:
		WriteFloat32(enc, *v)
	case *float64:
		WriteFloat64(enc, *v)
	case *complex64:
		WriteComplex64(enc, *v)
	case *complex128:
		WriteComplex128(enc, *v)
	case *big.Int:
		WriteBigInt(enc, v)
	case *big.Float:
		WriteBigFloat(enc, v)
	case *big.Rat:
		WriteBigRat(enc, v)
	case *error:
		WriteError(enc, *v)
	default:
		ok = false
	}
	return
}
func writePtr(enc *Encoder, v interface{}, encode func(m ValueEncoder, enc *Encoder, v interface{})) {
	if fastWritePtr(enc, v) {
		return
	}
	e := reflect.ValueOf(v).Elem()
	kind := e.Kind()
	switch kind {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface:
		if e.IsNil() {
			WriteNil(enc)
			return
		}
	}
	et := e.Type()
	if valenc := getOtherEncoder(et); valenc != nil {
		encode(valenc, enc, v)
		return
	}
	switch kind {
	case reflect.Int:
		WriteInt(enc, *(*int)(reflect2.PtrOf(v)))
	case reflect.Int8:
		WriteInt8(enc, *(*int8)(reflect2.PtrOf(v)))
	case reflect.Int16:
		WriteInt16(enc, *(*int16)(reflect2.PtrOf(v)))
	case reflect.Int32:
		WriteInt32(enc, *(*int32)(reflect2.PtrOf(v)))
	case reflect.Int64:
		WriteInt64(enc, *(*int64)(reflect2.PtrOf(v)))
	case reflect.Uint:
		WriteUint(enc, *(*uint)(reflect2.PtrOf(v)))
	case reflect.Uint8:
		WriteUint8(enc, *(*uint8)(reflect2.PtrOf(v)))
	case reflect.Uint16:
		WriteUint16(enc, *(*uint16)(reflect2.PtrOf(v)))
	case reflect.Uint32:
		WriteUint32(enc, *(*uint32)(reflect2.PtrOf(v)))
	case reflect.Uint64, reflect.Uintptr:
		WriteUint64(enc, *(*uint64)(reflect2.PtrOf(v)))
	case reflect.Bool:
		WriteBool(enc, *(*bool)(reflect2.PtrOf(v)))
	case reflect.Float32:
		WriteFloat32(enc, *(*float32)(reflect2.PtrOf(v)))
	case reflect.Float64:
		WriteFloat64(enc, *(*float64)(reflect2.PtrOf(v)))
	case reflect.Complex64:
		WriteComplex64(enc, *(*complex64)(reflect2.PtrOf(v)))
	case reflect.Complex128:
		WriteComplex128(enc, *(*complex128)(reflect2.PtrOf(v)))
	case reflect.String:
		encode(strenc, enc, e.String())
	case reflect.Array:
		encode(arrayenc, enc, v)
	case reflect.Struct:
		encode(getStructEncoder(et), enc, v)
	case reflect.Slice:
		encode(slcenc, enc, v)
	case reflect.Map:
		encode(mapenc, enc, v)
	case reflect.Ptr:
		encode(ptrenc, enc, e.Interface())
	case reflect.Interface:
		encode(intfenc, enc, e.Interface())
	default:
		WriteNil(enc)
		enc.Error = &UnsupportedTypeError{Type: reflect.TypeOf(v)}
	}
}
