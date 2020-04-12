/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/encoder.go                                      |
|                                                          |
| LastModified: Apr 12, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"io"
	"math"
	"math/big"
	"reflect"
	"strconv"
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

func (enc *Encoder) copyCheck() {
	if enc.addr == nil {
		// This hack works around a failing of Go's escape analysis
		// that was causing enc to escape and be heap allocated.
		// See issue 23382.
		// TODO: once issue 7921 is fixed, this should be reverted to
		// just "enc.addr = enc".
		enc.addr = (*Encoder)(reflect2.NoEscape(unsafe.Pointer(enc)))
	} else if enc.addr != enc {
		panic("hprose/encoding: illegal use of non-zero Encoder copied by value")
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
		WriteSlice(enc, v)
	case reflect.Map:
		WriteMap(enc, v)
	case reflect.Ptr:
		encode(ptrenc, v)
	default:
		enc.Error = &UnsupportedTypeError{Type: reflect.TypeOf(v)}
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
		switch reflect.TypeOf(v).Kind() {
		case reflect.Ptr:
			enc.refer.Set(v)
		case reflect.String:
			enc.refer.SetString(v.(string))
		default:
			enc.refer.AddCount(1)
		}
	}
}

// SetPtrReference of v
func (enc *Encoder) SetPtrReference(v interface{}) {
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

// Simple resets the encoder to simple mode or not
func (enc *Encoder) Simple(simple bool) {
	if simple {
		enc.refer = nil
	} else {
		enc.refer = &encoderRefer{}
	}
	enc.ref = nil
	enc.last = 0
}

// WriteNil to encoder
func (enc *Encoder) WriteNil() {
	enc.buf = append(enc.buf, TagNull)
}

// WriteBool to encoder
func (enc *Encoder) WriteBool(b bool) {
	if b {
		enc.buf = append(enc.buf, TagTrue)
	} else {
		enc.buf = append(enc.buf, TagFalse)
	}
}

// WriteInt64 to encoder
func (enc *Encoder) WriteInt64(i int64) {
	if (i >= 0) && (i <= 9) {
		enc.buf = append(enc.buf, digits[i])
	} else {
		var tag = TagInteger
		if (i < math.MinInt32) || (i > math.MaxInt32) {
			tag = TagLong
		}
		enc.buf = append(enc.buf, tag)
		enc.buf = AppendInt64(enc.buf, i)
		enc.buf = append(enc.buf, TagSemicolon)
	}
}

// WriteUint64 to encoder
func (enc *Encoder) WriteUint64(i uint64) {
	if (i >= 0) && (i <= 9) {
		enc.buf = append(enc.buf, digits[i])
	} else {
		var tag = TagInteger
		if i > math.MaxInt32 {
			tag = TagLong
		}
		enc.buf = append(enc.buf, tag)
		enc.buf = AppendUint64(enc.buf, i)
		enc.buf = append(enc.buf, TagSemicolon)
	}
}

// WriteInt32 to encoder
func (enc *Encoder) WriteInt32(i int32) {
	if (i >= 0) && (i <= 9) {
		enc.buf = append(enc.buf, digits[i])
	} else {
		enc.buf = append(enc.buf, TagInteger)
		enc.buf = AppendInt64(enc.buf, int64(i))
		enc.buf = append(enc.buf, TagSemicolon)
	}
}

// WriteUint32 to encoder
func (enc *Encoder) WriteUint32(i uint32) {
	enc.WriteUint64(uint64(i))
}

// WriteInt16 to encoder
func (enc *Encoder) WriteInt16(i int16) {
	enc.WriteInt32(int32(i))
}

// WriteUint16 to encoder
func (enc *Encoder) WriteUint16(i uint16) {
	if (i >= 0) && (i <= 9) {
		enc.buf = append(enc.buf, digits[i])
		return
	}
	enc.buf = append(enc.buf, TagInteger)
	enc.buf = AppendUint64(enc.buf, uint64(i))
	enc.buf = append(enc.buf, TagSemicolon)
	return
}

// WriteInt8 to encoder
func (enc *Encoder) WriteInt8(i int8) {
	enc.WriteInt32(int32(i))
}

// WriteUint8 to encoder
func (enc *Encoder) WriteUint8(i uint8) {
	enc.WriteUint16(uint16(i))
}

// WriteInt to encoder
func (enc *Encoder) WriteInt(i int) {
	enc.WriteInt64(int64(i))
}

// WriteUint to encoder
func (enc *Encoder) WriteUint(i uint) {
	enc.WriteUint64(uint64(i))
}

func (enc *Encoder) writeFloat(f float64, bitSize int) {
	switch {
	case f != f:
		enc.buf = append(enc.buf, TagNaN)
	case f > math.MaxFloat64:
		enc.buf = append(enc.buf, TagInfinity, TagPos)
	case f < -math.MaxFloat64:
		enc.buf = append(enc.buf, TagInfinity, TagNeg)
	default:
		enc.buf = append(enc.buf, TagDouble)
		enc.buf = strconv.AppendFloat(enc.buf, f, 'g', -1, bitSize)
		enc.buf = append(enc.buf, TagSemicolon)
	}
}

// WriteFloat32 to encoder
func (enc *Encoder) WriteFloat32(f float32) {
	enc.writeFloat(float64(f), 32)
}

// WriteFloat64 to encoder
func (enc *Encoder) WriteFloat64(f float64) {
	enc.writeFloat(f, 64)
}

// WriteHead to encoder, n is the count of elements in list or map
func (enc *Encoder) WriteHead(n int, tag byte) {
	enc.buf = append(enc.buf, tag)
	if n > 0 {
		enc.buf = AppendUint64(enc.buf, uint64(n))
	}
	enc.buf = append(enc.buf, TagOpenbrace)
}

// WriteObjectHead to encoder, r is the reference number of struct
func (enc *Encoder) WriteObjectHead(r int) {
	enc.buf = append(enc.buf, TagObject)
	enc.buf = AppendUint64(enc.buf, uint64(r))
	enc.buf = append(enc.buf, TagOpenbrace)
}

// WriteFoot of list or map to encoder
func (enc *Encoder) WriteFoot() {
	enc.buf = append(enc.buf, TagClosebrace)
}

func (enc *Encoder) writeComplex(r float64, i float64, bitSize int) {
	if i == 0 {
		enc.writeFloat(r, bitSize)
	} else {
		enc.AddReferenceCount(1)
		enc.WriteHead(2, TagList)
		enc.writeFloat(r, bitSize)
		enc.writeFloat(i, bitSize)
		enc.WriteFoot()
	}
}

// WriteComplex64 to encoder
func (enc *Encoder) WriteComplex64(c complex64) {
	enc.writeComplex(float64(real(c)), float64(imag(c)), 32)
}

// WriteComplex128 to encoder
func (enc *Encoder) WriteComplex128(c complex128) {
	enc.writeComplex(real(c), imag(c), 64)
}

// WriteBigFloat to encoder
func (enc *Encoder) WriteBigFloat(f *big.Float) {
	enc.buf = append(enc.buf, TagDouble)
	enc.buf = f.Append(enc.buf, 'g', -1)
	enc.buf = append(enc.buf, TagSemicolon)
}

// WriteBigInt to encoder
func (enc *Encoder) WriteBigInt(i *big.Int) {
	enc.buf = append(enc.buf, TagLong)
	enc.buf = append(enc.buf, i.String()...)
	enc.buf = append(enc.buf, TagSemicolon)
}

// WriteBigRat to encoder
func (enc *Encoder) WriteBigRat(r *big.Rat) {
	if r.IsInt() {
		enc.WriteBigInt(r.Num())
	} else {
		enc.AddReferenceCount(1)
		s := r.String()
		enc.buf = appendString(enc.buf, s, len(s))
	}
}

// WriteError to encoder
func (enc *Encoder) WriteError(e error) {
	enc.AddReferenceCount(1)
	s := e.Error()
	enc.buf = append(enc.buf, TagError)
	enc.buf = appendString(enc.buf, s, utf16Length(s))
}

// EncodeReference to encoder
func (enc *Encoder) EncodeReference(valenc ValueEncoder, v interface{}) {
	if reflect2.IsNil(v) {
		enc.WriteNil()
	} else if ok := enc.WriteReference(v); !ok {
		valenc.Write(enc, v)
	}
}
