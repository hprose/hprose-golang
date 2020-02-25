/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/encoder/encoder.go                           |
|                                                          |
| LastModified: Feb 25, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoder

import (
	"reflect"

	"github.com/hprose/hprose-golang/v3/io"
)

// An UnsupportedTypeError is returned by Marshal when attempting
// to encode an unsupported value type.
type UnsupportedTypeError struct {
	Type reflect.Type
}

func (e *UnsupportedTypeError) Error() string {
	return "hprose: unsupported type: " + e.Type.String()
}

// An Encoder writes hprose data to an output stream
type Encoder struct {
	io.Writer
}

func (enc *Encoder) marshal(v interface{}, marshal func(m Marshaler, v interface{}) error) error {
	switch value := v.(type) {
	case nil:
		return WriteNil(enc.Writer)
	case int:
		return WriteInt(enc.Writer, value)
	case int8:
		return WriteInt8(enc.Writer, value)
	case int16:
		return WriteInt16(enc.Writer, value)
	case int32:
		return WriteInt32(enc.Writer, value)
	case int64:
		return WriteInt64(enc.Writer, value)
	case uint:
		return WriteUint(enc.Writer, value)
	case uint8:
		return WriteUint8(enc.Writer, value)
	case uint16:
		return WriteUint16(enc.Writer, value)
	case uint32:
		return WriteUint32(enc.Writer, value)
	case uint64:
		return WriteUint64(enc.Writer, value)
	case bool:
		return WriteBool(enc.Writer, value)
	case float32:
		return WriteFloat32(enc.Writer, value)
	case float64:
		return WriteFloat64(enc.Writer, value)
	case complex64:
		return WriteComplex64(enc, value)
	case complex128:
		return WriteComplex128(enc, value)
	default:
		e := reflect.TypeOf(v)
		if e.Kind() == reflect.Struct {
			marshaler := GetValueMarshaler(e)
			if marshaler != nil {
				return marshaler(enc, v)
			}
		}
		if m := getMarshaler(v); m != nil {
			return marshal(m, v)
		}
		return &UnsupportedTypeError{Type: reflect.TypeOf(v)}
	}
}

func (enc *Encoder) encode(m Marshaler, v interface{}) error {
	return m.Encode(enc, v)
}

func (enc *Encoder) write(m Marshaler, v interface{}) error {
	return m.Write(enc, v)
}

// Encode writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as reference
func (enc *Encoder) Encode(v interface{}) error {
	return enc.marshal(v, enc.encode)
}

// Write writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as value
func (enc *Encoder) Write(v interface{}) error {
	return enc.marshal(v, enc.write)
}

// WriteReference of v to stream
func (enc *Encoder) WriteReference(v interface{}) (bool, error) {
	return false, nil
}

// SetReference of v
func (enc *Encoder) SetReference(v interface{}) {
}

// AddReferenceCount n
func (enc *Encoder) AddReferenceCount(n int) {
}

// WriteStructType of t to stream with action
func (enc *Encoder) WriteStructType(t reflect.Type, action func()) (int, error) {
	return 0, nil
}

// Reset the value reference and struct type reference
func (enc *Encoder) Reset() {

}
