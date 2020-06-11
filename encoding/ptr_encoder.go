/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/ptr_encoder.go                                  |
|                                                          |
| LastModified: Apr 12, 2020                               |
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
	enc.writePtr(v, func(valenc ValueEncoder, v interface{}) {
		valenc.Encode(enc, v)
	})
}

func (ptrEncoder) Write(enc *Encoder, v interface{}) {
	enc.writePtr(v, func(valenc ValueEncoder, v interface{}) {
		valenc.Write(enc, v)
	})
}

func (enc *Encoder) fastWritePtr(v interface{}) (ok bool) {
	ok = true
	switch v := v.(type) {
	case *int:
		enc.WriteInt(*v)
	case *int8:
		enc.WriteInt8(*v)
	case *int16:
		enc.WriteInt16(*v)
	case *int32:
		enc.WriteInt32(*v)
	case *int64:
		enc.WriteInt64(*v)
	case *uint:
		enc.WriteUint(*v)
	case *uint8:
		enc.WriteUint8(*v)
	case *uint16:
		enc.WriteUint16(*v)
	case *uint32:
		enc.WriteUint32(*v)
	case *uint64:
		enc.WriteUint64(*v)
	case *uintptr:
		enc.WriteUint64(uint64(*v))
	case *bool:
		enc.WriteBool(*v)
	case *float32:
		enc.WriteFloat32(*v)
	case *float64:
		enc.WriteFloat64(*v)
	case *complex64:
		enc.WriteComplex64(*v)
	case *complex128:
		enc.WriteComplex128(*v)
	case *big.Int:
		enc.WriteBigInt(v)
	case *big.Float:
		enc.WriteBigFloat(v)
	case *big.Rat:
		enc.WriteBigRat(v)
	case *error:
		enc.WriteError(*v)
	default:
		ok = false
	}
	return
}
func (enc *Encoder) writePtr(v interface{}, encode func(m ValueEncoder, v interface{})) {
	if enc.fastWritePtr(v) {
		return
	}
	e := reflect.ValueOf(v).Elem()
	kind := e.Kind()
	switch kind {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface:
		if e.IsNil() {
			enc.WriteNil()
			return
		}
	}
	et := e.Type()
	if valenc := getOtherEncoder(et); valenc != nil {
		encode(valenc, v)
		return
	}
	switch kind {
	case reflect.Int:
		enc.WriteInt(*(*int)(reflect2.PtrOf(v)))
	case reflect.Int8:
		enc.WriteInt8(*(*int8)(reflect2.PtrOf(v)))
	case reflect.Int16:
		enc.WriteInt16(*(*int16)(reflect2.PtrOf(v)))
	case reflect.Int32:
		enc.WriteInt32(*(*int32)(reflect2.PtrOf(v)))
	case reflect.Int64:
		enc.WriteInt64(*(*int64)(reflect2.PtrOf(v)))
	case reflect.Uint:
		enc.WriteUint(*(*uint)(reflect2.PtrOf(v)))
	case reflect.Uint8:
		enc.WriteUint8(*(*uint8)(reflect2.PtrOf(v)))
	case reflect.Uint16:
		enc.WriteUint16(*(*uint16)(reflect2.PtrOf(v)))
	case reflect.Uint32:
		enc.WriteUint32(*(*uint32)(reflect2.PtrOf(v)))
	case reflect.Uint64, reflect.Uintptr:
		enc.WriteUint64(*(*uint64)(reflect2.PtrOf(v)))
	case reflect.Bool:
		enc.WriteBool(*(*bool)(reflect2.PtrOf(v)))
	case reflect.Float32:
		enc.WriteFloat32(*(*float32)(reflect2.PtrOf(v)))
	case reflect.Float64:
		enc.WriteFloat64(*(*float64)(reflect2.PtrOf(v)))
	case reflect.Complex64:
		enc.WriteComplex64(*(*complex64)(reflect2.PtrOf(v)))
	case reflect.Complex128:
		enc.WriteComplex128(*(*complex128)(reflect2.PtrOf(v)))
	case reflect.String:
		encode(strenc, e.String())
	case reflect.Array:
		encode(arrayenc, v)
	case reflect.Struct:
		encode(getStructEncoder(et), v)
	case reflect.Slice:
		encode(slcenc, v)
	case reflect.Map:
		encode(mapenc, v)
	case reflect.Ptr:
		encode(ptrenc, e.Interface())
	case reflect.Interface:
		encode(intfenc, e.Interface())
	default:
		enc.WriteNil()
		enc.Error = UnsupportedTypeError{Type: reflect.TypeOf(v)}
	}
}
