/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/ptr_encoder.go                               |
|                                                          |
| LastModified: Mar 15, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"math/big"
	"reflect"
)

// PtrEncoder is the implementation of ValueEncoder for ptr.
type PtrEncoder struct{}

var ptrEncoder PtrEncoder

// Encode writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as reference
func (PtrEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	return writePtr(enc, v, func(valenc ValueEncoder, enc *Encoder, v interface{}) error {
		return valenc.Encode(enc, v)
	})
}

// Write writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as value
func (PtrEncoder) Write(enc *Encoder, v interface{}) (err error) {
	return writePtr(enc, v, func(valenc ValueEncoder, enc *Encoder, v interface{}) error {
		return valenc.Write(enc, v)
	})
}

func writePtr(enc *Encoder, v interface{}, encode func(m ValueEncoder, enc *Encoder, v interface{}) error) (err error) {
	switch v := v.(type) {
	case *uint8:
		return WriteUint8(enc.Writer, *v)
	case *uint16:
		return WriteUint16(enc.Writer, *v)
	case *uint32:
		return WriteUint32(enc.Writer, *v)
	case *uint64:
		return WriteUint64(enc.Writer, *v)
	case *uint:
		return WriteUint(enc.Writer, *v)
	case *int8:
		return WriteInt8(enc.Writer, *v)
	case *int16:
		return WriteInt16(enc.Writer, *v)
	case *int32:
		return WriteInt32(enc.Writer, *v)
	case *int64:
		return WriteInt64(enc.Writer, *v)
	case *int:
		return WriteInt(enc.Writer, *v)
	case *uintptr:
		return WriteUint64(enc.Writer, uint64(*v))
	case *bool:
		return WriteBool(enc.Writer, *v)
	case *float32:
		return WriteFloat32(enc.Writer, *v)
	case *float64:
		return WriteFloat64(enc.Writer, *v)
	case *complex64:
		return WriteComplex64(enc, *v)
	case *complex128:
		return WriteComplex128(enc, *v)
	case *big.Int:
		return WriteBigInt(enc.Writer, v)
	case *big.Float:
		return WriteBigFloat(enc.Writer, v)
	case *big.Rat:
		return WriteBigRat(enc, v)
	}
	e := reflect.ValueOf(v).Elem()
	kind := e.Kind()
	switch kind {
	case reflect.String:
		return encode(stringEncoder, enc, *(v.(*string)))
	case reflect.Array:
		return encode(arrayEncoder, enc, v)
	case reflect.Struct:
		if valenc := GetEncoder(reflect.TypeOf(v)); valenc != nil {
			return encode(valenc, enc, v)
		}
	}
	if e.IsNil() {
		return WriteNil(enc.Writer)
	}
	switch kind {
	case reflect.Slice:
		return encode(sliceEncoder, enc, v)
	case reflect.Map:
		return encode(mapEncoder, enc, v)
	case reflect.Ptr:
		return encode(ptrEncoder, enc, e.Interface())
	case reflect.Interface:
		return encode(interfaceEncoder, enc, e.Interface())
	}
	return &UnsupportedTypeError{Type: reflect.TypeOf(v)}
}
