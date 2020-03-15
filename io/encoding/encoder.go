/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/encoder.go                                   |
|                                                          |
| LastModified: Mar 15, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"math/big"
	"reflect"

	"github.com/hprose/hprose-golang/v3/io"
	"github.com/modern-go/reflect2"
)

// An Encoder writes hprose data to an output stream
type Encoder struct {
	io.Writer
	refer *encoderRefer
	ref   map[reflect.Type]uint64
	last  uint64
}

// NewEncoder create an encoder object
func NewEncoder(writer io.Writer, simple bool) (encoder *Encoder) {
	encoder = &Encoder{
		Writer: writer,
		ref:    make(map[reflect.Type]uint64),
		last:   0,
	}
	if !simple {
		encoder.refer = newEncoderRefer()
	}
	return
}

func (enc *Encoder) writeValue(v interface{}, encode func(m ValueEncoder, v interface{}) error) error {
	switch v := v.(type) {
	case nil:
		return WriteNil(enc.Writer)
	case uint8:
		return WriteUint8(enc.Writer, v)
	case uint16:
		return WriteUint16(enc.Writer, v)
	case uint32:
		return WriteUint32(enc.Writer, v)
	case uint64:
		return WriteUint64(enc.Writer, v)
	case uint:
		return WriteUint(enc.Writer, v)
	case int8:
		return WriteInt8(enc.Writer, v)
	case int16:
		return WriteInt16(enc.Writer, v)
	case int32:
		return WriteInt32(enc.Writer, v)
	case int64:
		return WriteInt64(enc.Writer, v)
	case int:
		return WriteInt(enc.Writer, v)
	case uintptr:
		return WriteUint64(enc.Writer, uint64(v))
	case bool:
		return WriteBool(enc.Writer, v)
	case float32:
		return WriteFloat32(enc.Writer, v)
	case float64:
		return WriteFloat64(enc.Writer, v)
	case complex64:
		return WriteComplex64(enc, v)
	case complex128:
		return WriteComplex128(enc, v)
	case big.Int:
		return WriteBigInt(enc.Writer, &v)
	case big.Float:
		return WriteBigFloat(enc.Writer, &v)
	case big.Rat:
		return WriteBigRat(enc, &v)
	}
	t := reflect.TypeOf(v)
	kind := t.Kind()
	switch kind {
	case reflect.String:
		return encode(stringEncoder, v)
	case reflect.Array:
		return WriteArray(enc, v)
	case reflect.Struct:
		if m := GetEncoder(reflect.PtrTo(t)); m != nil {
			return encode(m, reflect.NewAt(t, reflect2.PtrOf(v)).Interface())
		}
	}
	if reflect.ValueOf(v).IsNil() {
		return WriteNil(enc.Writer)
	}
	switch kind {
	case reflect.Slice:
		return WriteSlice(enc, v)
	case reflect.Map:
		return WriteMap(enc, v)
	case reflect.Ptr:
		return encode(ptrEncoder, v)
	}
	return &UnsupportedTypeError{Type: reflect.TypeOf(v)}
}

// Encode writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as reference
func (enc *Encoder) Encode(v interface{}) error {
	return enc.writeValue(v, func(valenc ValueEncoder, v interface{}) error {
		return valenc.Encode(enc, v)
	})
}

// Write writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as value
func (enc *Encoder) Write(v interface{}) error {
	return enc.writeValue(v, func(valenc ValueEncoder, v interface{}) error {
		return valenc.Write(enc, v)
	})
}

// WriteReference of v to stream
func (enc *Encoder) WriteReference(v interface{}) (bool, error) {
	if enc.refer != nil {
		return enc.refer.Write(enc.Writer, v)
	}
	return false, nil
}

// WriteStringReference of v to stream
func (enc *Encoder) WriteStringReference(s string) (bool, error) {
	if enc.refer != nil {
		return enc.refer.WriteString(enc.Writer, s)
	}
	return false, nil
}

// SetReference of v
func (enc *Encoder) SetReference(v interface{}) {
	if enc.refer != nil {
		enc.refer.Set(v)
	}
}

// SetStringReference of v
func (enc *Encoder) SetStringReference(s string) {
	if enc.refer != nil {
		enc.refer.SetString(s)
	}
}

// AddReferenceCount n
func (enc *Encoder) AddReferenceCount(n int) {
	if enc.refer != nil {
		enc.refer.AddCount(n)
	}
}

// WriteStructType of t to stream with action
func (enc *Encoder) WriteStructType(t reflect.Type, action func()) (int, error) {
	return 0, nil
}

// Reset the value reference and struct type reference
func (enc *Encoder) Reset() {
	if enc.refer != nil {
		enc.refer.Reset()
	}
	enc.ref = make(map[reflect.Type]uint64)
	enc.last = 0
}
