/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/ptr_encoder.go                                  |
|                                                          |
| LastModified: Mar 21, 2020                               |
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

func writePtr(enc *Encoder, v interface{}, encode func(m ValueEncoder, enc *Encoder, v interface{})) {
	switch v := v.(type) {
	case *int:
		WriteInt(enc, *v)
		return
	case *int8:
		WriteInt8(enc, *v)
		return
	case *int16:
		WriteInt16(enc, *v)
		return
	case *int32:
		WriteInt32(enc, *v)
		return
	case *int64:
		WriteInt64(enc, *v)
		return
	case *uint:
		WriteUint(enc, *v)
		return
	case *uint8:
		WriteUint8(enc, *v)
		return
	case *uint16:
		WriteUint16(enc, *v)
		return
	case *uint32:
		WriteUint32(enc, *v)
		return
	case *uint64:
		WriteUint64(enc, *v)
		return
	case *uintptr:
		WriteUint64(enc, uint64(*v))
		return
	case *bool:
		WriteBool(enc, *v)
		return
	case *float32:
		WriteFloat32(enc, *v)
		return
	case *float64:
		WriteFloat64(enc, *v)
		return
	case *complex64:
		WriteComplex64(enc, *v)
		return
	case *complex128:
		WriteComplex128(enc, *v)
		return
	case *big.Int:
		WriteBigInt(enc, v)
		return
	case *big.Float:
		WriteBigFloat(enc, v)
		return
	case *big.Rat:
		WriteBigRat(enc, v)
		return
	case *error:
		WriteError(enc, *v)
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
		return
	case reflect.Int8:
		WriteInt8(enc, *(*int8)(reflect2.PtrOf(v)))
		return
	case reflect.Int16:
		WriteInt16(enc, *(*int16)(reflect2.PtrOf(v)))
		return
	case reflect.Int32:
		WriteInt32(enc, *(*int32)(reflect2.PtrOf(v)))
		return
	case reflect.Int64:
		WriteInt64(enc, *(*int64)(reflect2.PtrOf(v)))
		return
	case reflect.Uint:
		WriteUint(enc, *(*uint)(reflect2.PtrOf(v)))
		return
	case reflect.Uint8:
		WriteUint8(enc, *(*uint8)(reflect2.PtrOf(v)))
		return
	case reflect.Uint16:
		WriteUint16(enc, *(*uint16)(reflect2.PtrOf(v)))
		return
	case reflect.Uint32:
		WriteUint32(enc, *(*uint32)(reflect2.PtrOf(v)))
		return
	case reflect.Uint64, reflect.Uintptr:
		WriteUint64(enc, *(*uint64)(reflect2.PtrOf(v)))
		return
	case reflect.Bool:
		WriteBool(enc, *(*bool)(reflect2.PtrOf(v)))
		return
	case reflect.Float32:
		WriteFloat32(enc, *(*float32)(reflect2.PtrOf(v)))
		return
	case reflect.Float64:
		WriteFloat64(enc, *(*float64)(reflect2.PtrOf(v)))
		return
	case reflect.Complex64:
		WriteComplex64(enc, *(*complex64)(reflect2.PtrOf(v)))
		return
	case reflect.Complex128:
		WriteComplex128(enc, *(*complex128)(reflect2.PtrOf(v)))
		return
	case reflect.String:
		encode(strenc, enc, e.String())
		return
	case reflect.Array:
		encode(arrayenc, enc, v)
		return
	case reflect.Struct:
		encode(getStructEncoder(et), enc, v)
		return
	case reflect.Slice:
		encode(slcenc, enc, v)
		return
	case reflect.Map:
		encode(mapenc, enc, v)
		return
	case reflect.Ptr:
		encode(ptrenc, enc, e.Interface())
		return
	case reflect.Interface:
		encode(intfenc, enc, e.Interface())
		return
	}
	panic(&UnsupportedTypeError{Type: reflect.TypeOf(v)})
}
