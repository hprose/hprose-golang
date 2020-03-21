/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/encoder.go                                      |
|                                                          |
| LastModified: Mar 21, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"bufio"
	"io"
	"math/big"
	"reflect"

	"github.com/modern-go/reflect2"
)

// An Encoder writes hprose data to an output stream
type Encoder struct {
	writer BytesWriter
	refer  *encoderRefer
	ref    map[reflect.Type]int
	last   int
}

// NewEncoder create an encoder object
func NewEncoder(w io.Writer, simple bool) (encoder *Encoder) {
	encoder = &Encoder{
		ref:  make(map[reflect.Type]int),
		last: 0,
	}
	if writer, ok := w.(BytesWriter); ok {
		encoder.writer = writer
	} else {
		encoder.writer = bufio.NewWriter(w)
	}
	if !simple {
		encoder.refer = newEncoderRefer()
	}
	return
}

func (enc *Encoder) writeValue(v interface{}, encode func(m ValueEncoder, v interface{}) error) error {
	switch v := v.(type) {
	case nil:
		return WriteNil(enc)
	case int:
		return WriteInt(enc, v)
	case int8:
		return WriteInt8(enc, v)
	case int16:
		return WriteInt16(enc, v)
	case int32:
		return WriteInt32(enc, v)
	case int64:
		return WriteInt64(enc, v)
	case uint:
		return WriteUint(enc, v)
	case uint8:
		return WriteUint8(enc, v)
	case uint16:
		return WriteUint16(enc, v)
	case uint32:
		return WriteUint32(enc, v)
	case uint64:
		return WriteUint64(enc, v)
	case uintptr:
		return WriteUint64(enc, uint64(v))
	case bool:
		return WriteBool(enc, v)
	case float32:
		return WriteFloat32(enc, v)
	case float64:
		return WriteFloat64(enc, v)
	case complex64:
		return WriteComplex64(enc, v)
	case complex128:
		return WriteComplex128(enc, v)
	case big.Int:
		return WriteBigInt(enc, &v)
	case big.Float:
		return WriteBigFloat(enc, &v)
	case big.Rat:
		return WriteBigRat(enc, &v)
	case error:
		return WriteError(enc, v)
	}
	t := reflect.TypeOf(v)
	kind := t.Kind()
	switch kind {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface:
		if reflect.ValueOf(v).IsNil() {
			return WriteNil(enc)
		}
	}
	if valenc := getOtherEncoder(t); valenc != nil {
		return encode(valenc, v)
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
		return encode(strenc, v)
	case reflect.Array:
		return WriteArray(enc, v)
	case reflect.Struct:
		return getStructEncoder(t).Write(enc, v)
	case reflect.Slice:
		return WriteSlice(enc, v)
	case reflect.Map:
		return WriteMap(enc, v)
	case reflect.Ptr:
		return encode(ptrenc, v)
	}
	return &UnsupportedTypeError{Type: reflect.TypeOf(v)}
}

// Writer returns returns the output destination for the encoder.
func (enc *Encoder) Writer() io.Writer {
	return enc.writer
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
		return enc.refer.Write(enc, v)
	}
	return false, nil
}

// WriteStringReference of v to stream
func (enc *Encoder) WriteStringReference(s string) (bool, error) {
	if enc.refer != nil {
		return enc.refer.WriteString(enc, s)
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
func (enc *Encoder) WriteStructType(t reflect.Type, action func() error) (r int, err error) {
	if r, ok := enc.ref[t]; ok {
		return r, nil
	}
	if err = action(); err == nil {
		r = enc.last
		enc.last++
		enc.ref[t] = r
	}
	return
}

// Reset the value reference and struct type reference
func (enc *Encoder) Reset() {
	if enc.refer != nil {
		enc.refer.Reset()
	}
	enc.ref = make(map[reflect.Type]int)
	enc.last = 0
}
