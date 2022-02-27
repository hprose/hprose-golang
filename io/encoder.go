/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoder.go                                            |
|                                                          |
| LastModified: Feb 27, 2022                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"io"
	"math/big"
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

// An Encoder writes hprose data to an output stream.
type Encoder struct {
	addr   *Encoder // of receiver, to detect copies by value.
	buf    []byte
	off    int
	simple bool
	refer  encoderRefer
	ref    map[reflect.Type]int
	last   int
	Writer io.Writer
	Error  error
}

// NewEncoder create an encoder object.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		Writer: w,
		simple: true,
	}
}

func (enc *Encoder) copyCheck() {
	if enc.addr == nil {
		// This hack works around a failing of Go's escape analysis
		// that was causing enc to escape and be heap allocated.
		// See issue 23382.
		// TODO: once issue 7921 is fixed, this should be reverted to
		// just "enc.addr = enc".
		enc.addr = (*Encoder)(reflect2.NoEscape(unsafe.Pointer(enc)))
	} else if enc.addr != enc {
		panic("hprose/io: illegal use of non-zero Encoder copied by value")
	}
}

func (enc *Encoder) fastWriteValue(v interface{}) (ok bool) {
	ok = true
	switch v := v.(type) {
	case nil:
		enc.WriteNil()
	case int:
		enc.WriteInt(v)
	case int8:
		enc.WriteInt8(v)
	case int16:
		enc.WriteInt16(v)
	case int32:
		enc.WriteInt32(v)
	case int64:
		enc.WriteInt64(v)
	case uint:
		enc.WriteUint(v)
	case uint8:
		enc.WriteUint8(v)
	case uint16:
		enc.WriteUint16(v)
	case uint32:
		enc.WriteUint32(v)
	case uint64:
		enc.WriteUint64(v)
	case uintptr:
		enc.WriteUint64(uint64(v))
	case bool:
		enc.WriteBool(v)
	case float32:
		enc.WriteFloat32(v)
	case float64:
		enc.WriteFloat64(v)
	case complex64:
		enc.WriteComplex64(v)
	case complex128:
		enc.WriteComplex128(v)
	case big.Int:
		enc.WriteBigInt(&v)
	case big.Float:
		enc.WriteBigFloat(&v)
	case big.Rat:
		enc.WriteBigRat(&v)
	case error:
		enc.WriteError(v)
	default:
		ok = false
	}
	return
}

func (enc *Encoder) writeValue(v interface{}, encode func(m ValueEncoder, v interface{})) {
	if enc.fastWriteValue(v) {
		return
	}
	t := reflect.TypeOf(v)
	kind := t.Kind()
	switch kind {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface:
		if reflect.ValueOf(v).IsNil() {
			enc.WriteNil()
			return
		}
	}
	if valenc := getOtherEncoder(t); valenc != nil {
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
		encode(strenc, v)
	case reflect.Array:
		enc.WriteArray(v)
	case reflect.Struct:
		getStructEncoder(t).Write(enc, v)
	case reflect.Slice:
		enc.WriteSlice(v)
	case reflect.Map:
		enc.WriteMap(v)
	case reflect.Ptr:
		encode(ptrenc, v)
	default:
		enc.Error = UnsupportedTypeError{reflect.TypeOf(v)}
		enc.WriteNil()
	}
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

// Flush writes the io data from buf to Writer.
func (enc *Encoder) Flush() (err error) {
	if enc.Error != nil {
		return enc.Error
	}
	if enc.Writer != nil && enc.off < len(enc.buf) {
		_, err = enc.Writer.Write(enc.buf[enc.off:])
		enc.off = len(enc.buf)
	}
	return
}

// Encode writes the hprose io of v to stream.
// If v is already written to stream, it will writes it as reference.
func (enc *Encoder) Encode(v interface{}) (err error) {
	enc.copyCheck()
	enc.encode(v)
	return enc.Flush()
}

// Write writes the hprose io of v to stream.
// If v is already written to stream, it will writes it as value.
func (enc *Encoder) Write(v interface{}) (err error) {
	enc.copyCheck()
	enc.write(v)
	return enc.Flush()
}

// Buffer returns the accumulated bytes.
func (enc *Encoder) Buffer() []byte {
	return enc.buf
}

// Bytes returns a copy of the accumulated bytes.
func (enc *Encoder) Bytes() []byte {
	bytes := make([]byte, len(enc.buf))
	copy(bytes, enc.buf)
	return bytes
}

// UnsafeString returns the accumulated string.
func (enc *Encoder) UnsafeString() string {
	return *(*string)(unsafe.Pointer(&enc.buf))
}

// String returns a copy of the accumulated string.
func (enc *Encoder) String() string {
	return string(enc.buf)
}

// WriteReference of v to stream.
func (enc *Encoder) WriteReference(v interface{}) bool {
	if !enc.IsSimple() {
		return enc.refer.Write(enc, v)
	}
	return false
}

// WriteStringReference of v to stream.
func (enc *Encoder) WriteStringReference(s string) bool {
	if !enc.IsSimple() {
		return enc.refer.WriteString(enc, s)
	}
	return false
}

// SetReference of v.
func (enc *Encoder) SetReference(v interface{}) {
	if !enc.IsSimple() {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Ptr:
			enc.refer.Set(v)
		default:
			enc.refer.AddCount(1)
		}
	}
}

func (enc *Encoder) setReference(v interface{}) {
	if !enc.IsSimple() {
		enc.refer.Set(v)
	}
}

// SetStringReference of v.
func (enc *Encoder) SetStringReference(s string) {
	if !enc.IsSimple() {
		enc.refer.SetString(s)
	}
}

// AddReferenceCount n.
func (enc *Encoder) AddReferenceCount(n int) {
	if !enc.IsSimple() {
		enc.refer.AddCount(n)
	}
}

// WriteStructType of t to stream with action.
func (enc *Encoder) WriteStructType(t reflect.Type, action func()) (r int) {
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

// Reset the value reference and struct type reference.
func (enc *Encoder) Reset() *Encoder {
	if !enc.IsSimple() {
		enc.refer.Reset()
	}
	for k := range enc.ref {
		delete(enc.ref, k)
	}
	enc.last = 0
	return enc
}

// ResetBuffer of the Encoder.
func (enc *Encoder) ResetBuffer() *Encoder {
	enc.buf = enc.buf[:0]
	return enc
}

// Simple resets the encoder to simple mode or not.
func (enc *Encoder) Simple(simple bool) *Encoder {
	enc.simple = simple
	enc.Reset()
	return enc
}

// IsSimple returns the encoder is in simple mode or not.
func (enc *Encoder) IsSimple() bool {
	return enc.simple
}

// WriteNil to encoder.
func (enc *Encoder) WriteNil() {
	enc.buf = append(enc.buf, TagNull)
}

func (enc *Encoder) writeHead(n int, tag byte) {
	enc.buf = append(enc.buf, tag)
	if n > 0 {
		enc.buf = AppendUint64(enc.buf, uint64(n))
	}
	enc.buf = append(enc.buf, TagOpenbrace)
}

// WriteListHead to encoder, n is the count of elements in list.
func (enc *Encoder) WriteListHead(n int) {
	enc.writeHead(n, TagList)
}

// WriteMapHead to encoder, n is the count of elements in map.
func (enc *Encoder) WriteMapHead(n int) {
	enc.writeHead(n, TagMap)
}

// WriteObjectHead to encoder, r is the reference number of struct.
func (enc *Encoder) WriteObjectHead(r int) {
	enc.buf = append(enc.buf, TagObject)
	enc.buf = AppendUint64(enc.buf, uint64(r))
	enc.buf = append(enc.buf, TagOpenbrace)
}

// WriteFoot of list or map to encoder.
func (enc *Encoder) WriteFoot() {
	enc.buf = append(enc.buf, TagClosebrace)
}

// EncodeReference to encoder.
func (enc *Encoder) EncodeReference(valenc ValueEncoder, v interface{}) {
	if reflect2.IsNil(v) {
		enc.WriteNil()
	} else if ok := enc.WriteReference(v); !ok {
		valenc.Write(enc, v)
	}
}

// WriteTag to encoder.
func (enc *Encoder) WriteTag(tag byte) {
	enc.buf = append(enc.buf, tag)
}
