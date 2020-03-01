/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/encoder/encoder.go                           |
|                                                          |
| LastModified: Mar 1, 2020                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoder

import (
	"math/big"
	"reflect"
	"unsafe"

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

type interfaceStruct struct {
	typ unsafe.Pointer
	ptr unsafe.Pointer
}

func interfacePointer(p *interface{}) *interfaceStruct {
	return (*interfaceStruct)(unsafe.Pointer(p))
}

type encoderRefer struct {
	ref  map[interface{}]uint64
	last uint64
}

func newEncoderRefer() *encoderRefer {
	return &encoderRefer{
		ref:  make(map[interface{}]uint64),
		last: 0,
	}
}

func (r *encoderRefer) AddCount(count int) {
	r.last += uint64(count)
}

func (r *encoderRefer) Set(v interface{}) {
	r.ref[v] = r.last
	r.last++
}

func (r *encoderRefer) Write(writer io.Writer, v interface{}) (ok bool, err error) {
	var i uint64
	if i, ok = r.ref[v]; ok {
		if err = writer.WriteByte(io.TagRef); err == nil {
			if err = writeUint64(writer, i); err == nil {
				err = writer.WriteByte(io.TagSemicolon)
			}
		}
	}
	return
}
func (r *encoderRefer) Reset() {
	r.ref = map[interface{}]uint64{}
	r.last = 0
}

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

func (enc *Encoder) marshal(v interface{}, marshal func(m Marshaler, v interface{}) error) error {
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
		return marshal(stringMarshaler, v)
	case reflect.Array:
		return WriteArray(enc, v)
	case reflect.Struct:
		if m := GetMarshaler(reflect.PtrTo(t)); m != nil {
			return marshal(m, reflect.NewAt(t, interfacePointer(&v).ptr).Interface())
		}
	}
	if reflect.ValueOf(v).IsNil() {
		return WriteNil(enc.Writer)
	}
	switch kind {
	case reflect.Slice:
		return WriteSlice(enc, v)
	case reflect.Map:
		// return mapMarshaler
	case reflect.Ptr:
		return marshal(ptrMarshaler, v)
	case reflect.Interface:
		// return interfaceMarshaler
	}
	return &UnsupportedTypeError{Type: reflect.TypeOf(v)}
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
	if enc.refer != nil {
		return enc.refer.Write(enc.Writer, v)
	}
	return false, nil
}

// SetReference of v
func (enc *Encoder) SetReference(v interface{}) {
	if enc.refer != nil {
		enc.refer.Set(v)
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
