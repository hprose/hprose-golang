/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/encoder.go                                      |
|                                                          |
| LastModified: Mar 22, 2020                               |
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
	addr   *Encoder // of receiver, to detect copies by value
	buf    []byte
	off    int
	refer  *encoderRefer
	ref    map[reflect.Type]int
	last   int
	Writer io.Writer
	Error  error
}

// NewEncoder create an encoder object
func NewEncoder(w io.Writer, simple bool) (encoder *Encoder) {
	encoder = &Encoder{Writer: w}
	if !simple {
		encoder.refer = &encoderRefer{}
	}
	return
}

// noescape hides a pointer from escape analysis.  noescape is
// the identity function but escape analysis doesn't think the
// output depends on the input. noescape is inlined and currently
// compiles down to zero instructions.
// USE CAREFULLY!
// This was copied from the runtime; see issues 23382 and 7921.
//go:nosplit
//go:nocheckptr
func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

func (enc *Encoder) copyCheck() {
	if enc.addr == nil {
		// This hack works around a failing of Go's escape analysis
		// that was causing enc to escape and be heap allocated.
		// See issue 23382.
		// TODO: once issue 7921 is fixed, this should be reverted to
		// just "enc.addr = enc".
		enc.addr = (*Encoder)(noescape(unsafe.Pointer(enc)))
	} else if enc.addr != enc {
		panic("hprose: illegal use of non-zero Encoder copied by value")
	}
}

func (enc *Encoder) fastWriteValue(v interface{}) (ok bool) {
	ok = true
	switch v := v.(type) {
	case nil:
		WriteNil(enc)
	case int:
		WriteInt(enc, v)
	case int8:
		WriteInt8(enc, v)
	case int16:
		WriteInt16(enc, v)
	case int32:
		WriteInt32(enc, v)
	case int64:
		WriteInt64(enc, v)
	case uint:
		WriteUint(enc, v)
	case uint8:
		WriteUint8(enc, v)
	case uint16:
		WriteUint16(enc, v)
	case uint32:
		WriteUint32(enc, v)
	case uint64:
		WriteUint64(enc, v)
	case uintptr:
		WriteUint64(enc, uint64(v))
	case bool:
		WriteBool(enc, v)
	case float32:
		WriteFloat32(enc, v)
	case float64:
		WriteFloat64(enc, v)
	case complex64:
		WriteComplex64(enc, v)
	case complex128:
		WriteComplex128(enc, v)
	case big.Int:
		WriteBigInt(enc, &v)
	case big.Float:
		WriteBigFloat(enc, &v)
	case big.Rat:
		WriteBigRat(enc, &v)
	case error:
		WriteError(enc, v)
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
	case reflect.Int8:
		WriteInt8(enc, *(*int8)(reflect2.PtrOf(v)))
	case reflect.Int16:
		WriteInt16(enc, *(*int16)(reflect2.PtrOf(v)))
	case reflect.Int32:
		WriteInt32(enc, *(*int32)(reflect2.PtrOf(v)))
	case reflect.Int64:
		WriteInt64(enc, *(*int64)(reflect2.PtrOf(v)))
	case reflect.Uint:
		WriteUint(enc, *(*uint)(reflect2.PtrOf(v)))
	case reflect.Uint8:
		WriteUint8(enc, *(*uint8)(reflect2.PtrOf(v)))
	case reflect.Uint16:
		WriteUint16(enc, *(*uint16)(reflect2.PtrOf(v)))
	case reflect.Uint32:
		WriteUint32(enc, *(*uint32)(reflect2.PtrOf(v)))
	case reflect.Uint64, reflect.Uintptr:
		WriteUint64(enc, *(*uint64)(reflect2.PtrOf(v)))
	case reflect.Bool:
		WriteBool(enc, *(*bool)(reflect2.PtrOf(v)))
	case reflect.Float32:
		WriteFloat32(enc, *(*float32)(reflect2.PtrOf(v)))
	case reflect.Float64:
		WriteFloat64(enc, *(*float64)(reflect2.PtrOf(v)))
	case reflect.Complex64:
		WriteComplex64(enc, *(*complex64)(reflect2.PtrOf(v)))
	case reflect.Complex128:
		WriteComplex128(enc, *(*complex128)(reflect2.PtrOf(v)))
	case reflect.String:
		encode(strenc, v)
	case reflect.Array:
		WriteArray(enc, v)
	case reflect.Struct:
		getStructEncoder(t).Write(enc, v)
	case reflect.Slice:
		WriteSlice(enc, v)
	case reflect.Map:
		WriteMap(enc, v)
	case reflect.Ptr:
		encode(ptrenc, v)
	default:
		enc.Error = &UnsupportedTypeError{Type: reflect.TypeOf(v)}
		WriteNil(enc)
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

// Flush writes the encoding data from buf to Writer
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

// Encode writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as reference
func (enc *Encoder) Encode(v interface{}) (err error) {
	enc.copyCheck()
	enc.encode(v)
	return enc.Flush()
}

// Write writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as value
func (enc *Encoder) Write(v interface{}) (err error) {
	enc.copyCheck()
	enc.write(v)
	return enc.Flush()
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
