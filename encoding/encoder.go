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
	"io"
	"math/big"
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

// An Encoder writes hprose data to an output stream
type Encoder struct {
	Writer io.Writer
	buf    []byte
	off    int
	refer  *encoderRefer
	ref    map[reflect.Type]int
	last   int
}

// NewEncoder create an encoder object
func NewEncoder(w io.Writer, simple bool) (encoder *Encoder) {
	encoder = &Encoder{Writer: w}
	if !simple {
		encoder.refer = &encoderRefer{}
	}
	return
}

func (enc *Encoder) writeValue(v interface{}, encode func(m ValueEncoder, v interface{})) {
	switch v := v.(type) {
	case nil:
		WriteNil(enc)
		return
	case int:
		WriteInt(enc, v)
		return
	case int8:
		WriteInt8(enc, v)
		return
	case int16:
		WriteInt16(enc, v)
		return
	case int32:
		WriteInt32(enc, v)
		return
	case int64:
		WriteInt64(enc, v)
		return
	case uint:
		WriteUint(enc, v)
		return
	case uint8:
		WriteUint8(enc, v)
		return
	case uint16:
		WriteUint16(enc, v)
		return
	case uint32:
		WriteUint32(enc, v)
		return
	case uint64:
		WriteUint64(enc, v)
		return
	case uintptr:
		WriteUint64(enc, uint64(v))
		return
	case bool:
		WriteBool(enc, v)
		return
	case float32:
		WriteFloat32(enc, v)
		return
	case float64:
		WriteFloat64(enc, v)
		return
	case complex64:
		WriteComplex64(enc, v)
		return
	case complex128:
		WriteComplex128(enc, v)
		return
	case big.Int:
		WriteBigInt(enc, &v)
		return
	case big.Float:
		WriteBigFloat(enc, &v)
		return
	case big.Rat:
		WriteBigRat(enc, &v)
		return
	case error:
		WriteError(enc, v)
		return
	}
	t := reflect.TypeOf(v)
	kind := t.Kind()
	switch kind {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface:
		if reflect.ValueOf(v).IsNil() {
			WriteNil(enc)
			return
		}
	}
	if valenc := getOtherEncoder(t); valenc != nil {
		encode(valenc, v)
		return
	}
	switch kind {
	case reflect.Int:
		WriteInt(enc, *(*int)(reflect2.PtrOf(v)))
		return
	case reflect.Int8:
		WriteInt8(enc, *(*int8)(reflect2.PtrOf(v)))
		return
	case reflect.Int16:
		WriteInt16(enc, *(*int16)(reflect2.PtrOf(v)))
		return
	case reflect.Int32:
		WriteInt32(enc, *(*int32)(reflect2.PtrOf(v)))
		return
	case reflect.Int64:
		WriteInt64(enc, *(*int64)(reflect2.PtrOf(v)))
		return
	case reflect.Uint:
		WriteUint(enc, *(*uint)(reflect2.PtrOf(v)))
		return
	case reflect.Uint8:
		WriteUint8(enc, *(*uint8)(reflect2.PtrOf(v)))
		return
	case reflect.Uint16:
		WriteUint16(enc, *(*uint16)(reflect2.PtrOf(v)))
		return
	case reflect.Uint32:
		WriteUint32(enc, *(*uint32)(reflect2.PtrOf(v)))
		return
	case reflect.Uint64, reflect.Uintptr:
		WriteUint64(enc, *(*uint64)(reflect2.PtrOf(v)))
		return
	case reflect.Bool:
		WriteBool(enc, *(*bool)(reflect2.PtrOf(v)))
		return
	case reflect.Float32:
		WriteFloat32(enc, *(*float32)(reflect2.PtrOf(v)))
		return
	case reflect.Float64:
		WriteFloat64(enc, *(*float64)(reflect2.PtrOf(v)))
		return
	case reflect.Complex64:
		WriteComplex64(enc, *(*complex64)(reflect2.PtrOf(v)))
		return
	case reflect.Complex128:
		WriteComplex128(enc, *(*complex128)(reflect2.PtrOf(v)))
		return
	case reflect.String:
		encode(strenc, v)
		return
	case reflect.Array:
		WriteArray(enc, v)
		return
	case reflect.Struct:
		getStructEncoder(t).Write(enc, v)
		return
	case reflect.Slice:
		WriteSlice(enc, v)
		return
	case reflect.Map:
		WriteMap(enc, v)
		return
	case reflect.Ptr:
		encode(ptrenc, v)
		return
	}
	panic(&UnsupportedTypeError{Type: reflect.TypeOf(v)})
}

func (enc *Encoder) encode(v interface{}) {
	enc.writeValue(v, func(valenc ValueEncoder, v interface{}) {
		valenc.Encode(enc, v)
	})
}

func (enc *Encoder) write(v interface{}) {
	enc.writeValue(v, func(valenc ValueEncoder, v interface{}) {
		valenc.Write(enc, v)
	})
}

func (enc *Encoder) flush() (err error) {
	if enc.Writer != nil {
		_, err = enc.Writer.Write(enc.buf[enc.off:])
		enc.off = len(enc.buf)
	}
	return
}

// Encode writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as reference
func (enc *Encoder) Encode(v interface{}) (err error) {
	enc.encode(v)
	return enc.flush()
}

// Write writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as value
func (enc *Encoder) Write(v interface{}) (err error) {
	enc.write(v)
	return enc.flush()
}

// Bytes returns the accumulated bytes.
func (enc *Encoder) Bytes() []byte {
	return enc.buf
}

// String returns the accumulated string.
func (enc *Encoder) String() string {
	return *(*string)(unsafe.Pointer(&enc.buf))
}

// WriteReference of v to stream
func (enc *Encoder) WriteReference(v interface{}) bool {
	if enc.refer != nil {
		return enc.refer.Write(enc, v)
	}
	return false
}

// WriteStringReference of v to stream
func (enc *Encoder) WriteStringReference(s string) bool {
	if enc.refer != nil {
		return enc.refer.WriteString(enc, s)
	}
	return false
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

// WriteStruct of t to stream with action
func (enc *Encoder) WriteStruct(t reflect.Type, action func()) (r int) {
	if enc.ref == nil {
		enc.ref = make(map[reflect.Type]int)
	}
	if r, ok := enc.ref[t]; ok {
		return r
	}
	action()
	r = enc.last
	enc.last++
	enc.ref[t] = r
	return
}

// Reset the value reference and struct type reference
func (enc *Encoder) Reset() {
	if enc.refer != nil {
		enc.refer.Reset()
	}
	enc.ref = nil
	enc.last = 0
}
