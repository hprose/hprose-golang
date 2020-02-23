/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/encoder/ptr_marshaler.go                     |
|                                                          |
| LastModified: Feb 24, 2020                               |
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

func (m PtrMarshaler) marshal(enc *Encoder, v interface{}, marshal func(enc *Encoder, v interface{}) error) (err error) {
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
	default:
		e := reflect.TypeOf(v).Elem()
		if e.Kind() == reflect.Struct {
			marshaler := getValueMarshaler(e)
			if marshaler != nil {
				return marshaler(enc, v)
			}
		}
		return marshal(enc, v)
	}
}

func (m PtrMarshaler) encode(enc *Encoder, v interface{}) (err error) {
	var ok bool
	if ok, err = enc.WriteReference(v); !ok && err == nil {
		err = m.write(enc, v)
	}
	return
}

func (m PtrMarshaler) write(enc *Encoder, v interface{}) (err error) {
	enc.SetReference(v)
	return enc.Encode(reflect.ValueOf(v).Elem().Interface())
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
