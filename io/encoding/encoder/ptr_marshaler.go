/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/encoder/ptr_marshaler.go                     |
|                                                          |
| LastModified: Mar 1, 2020                                |
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

func (PtrMarshaler) marshal(enc *Encoder, v interface{}, marshal func(m Marshaler, enc *Encoder, v interface{}) error) (err error) {
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
		return marshal(stringMarshaler, enc, *(v.(*string)))
	case reflect.Array:
		return marshal(arrayMarshaler, enc, v)
	case reflect.Struct:
		if marshaler := GetMarshaler(reflect.TypeOf(v)); marshaler != nil {
			return marshal(marshaler, enc, v)
		}
	}
	if e.IsNil() {
		return WriteNil(enc.Writer)
	}
	switch kind {
	case reflect.Slice:
		return marshal(sliceMarshaler, enc, v)
	case reflect.Map:
		// return mapMarshaler
	case reflect.Ptr:
		return marshal(ptrMarshaler, enc, e.Interface())
	case reflect.Interface:
		return marshal(interfaceMarshaler, enc, e.Interface())
	}
	return &UnsupportedTypeError{Type: reflect.TypeOf(v)}
}

func (PtrMarshaler) encode(m Marshaler, enc *Encoder, v interface{}) (err error) {
	return m.Encode(enc, v)
}

func (PtrMarshaler) write(m Marshaler, enc *Encoder, v interface{}) (err error) {
	return m.Write(enc, v)
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
