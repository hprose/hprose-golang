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

func (ptrEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	return writePtr(enc, v, func(valenc ValueEncoder, enc *Encoder, v interface{}) error {
		return valenc.Encode(enc, v)
	})
}

func (ptrEncoder) Write(enc *Encoder, v interface{}) (err error) {
	return writePtr(enc, v, func(valenc ValueEncoder, enc *Encoder, v interface{}) error {
		return valenc.Write(enc, v)
	})
}

func writePtr(enc *Encoder, v interface{}, encode func(m ValueEncoder, enc *Encoder, v interface{}) error) (err error) {
	switch v := v.(type) {
	case *int:
		return WriteInt(enc, *v)
	case *int8:
		return WriteInt8(enc, *v)
	case *int16:
		return WriteInt16(enc, *v)
	case *int32:
		return WriteInt32(enc, *v)
	case *int64:
		return WriteInt64(enc, *v)
	case *uint:
		return WriteUint(enc, *v)
	case *uint8:
		return WriteUint8(enc, *v)
	case *uint16:
		return WriteUint16(enc, *v)
	case *uint32:
		return WriteUint32(enc, *v)
	case *uint64:
		return WriteUint64(enc, *v)
	case *uintptr:
		return WriteUint64(enc, uint64(*v))
	case *bool:
		return WriteBool(enc, *v)
	case *float32:
		return WriteFloat32(enc, *v)
	case *float64:
		return WriteFloat64(enc, *v)
	case *complex64:
		return WriteComplex64(enc, *v)
	case *complex128:
		return WriteComplex128(enc, *v)
	case *big.Int:
		return WriteBigInt(enc, v)
	case *big.Float:
		return WriteBigFloat(enc, v)
	case *big.Rat:
		return WriteBigRat(enc, v)
	case *error:
		return WriteError(enc, *v)
	}
	e := reflect.ValueOf(v).Elem()
	kind := e.Kind()
	switch kind {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface:
		if e.IsNil() {
			return WriteNil(enc)
		}
	}
	et := e.Type()
	if valenc := getOtherEncoder(et); valenc != nil {
		return encode(valenc, enc, v)
	}
	switch kind {
	case reflect.Int:
		return WriteInt(enc, *(*int)(reflect2.PtrOf(v)))
	case reflect.Int8:
		return WriteInt8(enc, *(*int8)(reflect2.PtrOf(v)))
	case reflect.Int16:
		return WriteInt16(enc, *(*int16)(reflect2.PtrOf(v)))
	case reflect.Int32:
		return WriteInt32(enc, *(*int32)(reflect2.PtrOf(v)))
	case reflect.Int64:
		return WriteInt64(enc, *(*int64)(reflect2.PtrOf(v)))
	case reflect.Uint:
		return WriteUint(enc, *(*uint)(reflect2.PtrOf(v)))
	case reflect.Uint8:
		return WriteUint8(enc, *(*uint8)(reflect2.PtrOf(v)))
	case reflect.Uint16:
		return WriteUint16(enc, *(*uint16)(reflect2.PtrOf(v)))
	case reflect.Uint32:
		return WriteUint32(enc, *(*uint32)(reflect2.PtrOf(v)))
	case reflect.Uint64, reflect.Uintptr:
		return WriteUint64(enc, *(*uint64)(reflect2.PtrOf(v)))
	case reflect.Bool:
		return WriteBool(enc, *(*bool)(reflect2.PtrOf(v)))
	case reflect.Float32:
		return WriteFloat32(enc, *(*float32)(reflect2.PtrOf(v)))
	case reflect.Float64:
		return WriteFloat64(enc, *(*float64)(reflect2.PtrOf(v)))
	case reflect.Complex64:
		return WriteComplex64(enc, *(*complex64)(reflect2.PtrOf(v)))
	case reflect.Complex128:
		return WriteComplex128(enc, *(*complex128)(reflect2.PtrOf(v)))
	case reflect.String:
		return encode(strenc, enc, e.String())
	case reflect.Array:
		return encode(arrayenc, enc, v)
	case reflect.Struct:
		return encode(getStructEncoder(et), enc, v)
	case reflect.Slice:
		return encode(slcenc, enc, v)
	case reflect.Map:
		return encode(mapenc, enc, v)
	case reflect.Ptr:
		return encode(ptrenc, enc, e.Interface())
	case reflect.Interface:
		return encode(intfenc, enc, e.Interface())
	}
	return &UnsupportedTypeError{Type: reflect.TypeOf(v)}
}
