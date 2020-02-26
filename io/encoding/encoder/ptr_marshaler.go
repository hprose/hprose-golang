/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/encoder/ptr_marshaler.go                     |
|                                                          |
| LastModified: Feb 25, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoder

import (
	"math/big"
	"reflect"
)

// PtrMarshaler is the implementation of Marshaler for ptr.
type PtrMarshaler struct{}

var ptrMarshaler PtrMarshaler

func (m PtrMarshaler) marshal(enc *Encoder, v interface{}, marshal func(marshaler Marshaler, enc *Encoder, v interface{}) error) (err error) {
	if reflect.ValueOf(v).IsNil() {
		return WriteNil(enc.Writer)
	}
	switch v := v.(type) {
	case *int:
		return WriteInt(enc.Writer, *v)
	case *int8:
		return WriteInt8(enc.Writer, *v)
	case *int16:
		return WriteInt16(enc.Writer, *v)
	case *int32:
		return WriteInt32(enc.Writer, *v)
	case *int64:
		return WriteInt64(enc.Writer, *v)
	case *uint:
		return WriteUint(enc.Writer, *v)
	case *uint8:
		return WriteUint8(enc.Writer, *v)
	case *uint16:
		return WriteUint16(enc.Writer, *v)
	case *uint32:
		return WriteUint32(enc.Writer, *v)
	case *uint64:
		return WriteUint64(enc.Writer, *v)
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
	t := reflect.TypeOf(v)
	switch t.Elem().Kind() {
	case reflect.Struct:
		if marshaler := GetMarshaler(t); marshaler != nil {
			return marshal(marshaler, enc, v)
		}
	case reflect.String:
		return marshal(stringMarshaler, enc, *(v.(*string)))
	case reflect.Array:
	case reflect.Slice:
	case reflect.Map:
	case reflect.Ptr:
	case reflect.Interface:
	}
	return &UnsupportedTypeError{Type: reflect.TypeOf(v)}
}

func (m PtrMarshaler) encode(marshaler Marshaler, enc *Encoder, v interface{}) (err error) {
	return marshaler.Encode(enc, v)
}

func (m PtrMarshaler) write(marshaler Marshaler, enc *Encoder, v interface{}) (err error) {
	return marshaler.Write(enc, v)
}

// Encode writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as reference
func (m PtrMarshaler) Encode(enc *Encoder, v interface{}) (err error) {
	return m.marshal(enc, v, m.encode)
}

// Write writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as value
func (m PtrMarshaler) Write(enc *Encoder, v interface{}) (err error) {
	return m.marshal(enc, v, m.write)
}
