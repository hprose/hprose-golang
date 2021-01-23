/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/decoder.go                                      |
|                                                          |
| LastModified: Jan 24, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"bytes"
	"io"
	"math"
	"math/big"
	"reflect"
	"time"

	"github.com/google/uuid"
)

// LongType represents the default type for decode long integer
type LongType int8

const (
	// LongTypeInt64 represents the default type is int64
	LongTypeInt64 LongType = iota
	// LongTypeUint64 represents the default type is uint64
	LongTypeUint64
	// LongTypeBigInt represents the default type is *big.Int
	LongTypeBigInt
)

// RealType represents the default type for decode real number
type RealType int8

const (
	// RealTypeFloat64 represents the default type is float64
	RealTypeFloat64 RealType = iota
	// RealTypeFloat32 represents the default type is float32
	RealTypeFloat32
	// RealTypeBigFloat represents the default type is *big.Float
	RealTypeBigFloat
)

// MapType represents the default type for decode map
type MapType int8

const (
	// MapTypeIIMap represents the default type is map[interface{}]interface{}
	MapTypeIIMap MapType = iota
	// MapTypeSIMap represents the default type is map[string]interface{}
	MapTypeSIMap
)

const defaultBufferSize = 8192

// Decoder is a io.Reader like object, with hprose specific read functions.
// Error is not returned as return value, but stored as Error member on this decoder instance.
type Decoder struct {
	reader io.Reader
	buf    []byte
	head   int
	tail   int
	refer  *decoderRefer
	ref    []structInfo
	Error  error
	LongType
	RealType
	MapType
}

// NewDecoder creates an Decoder instance from byte array
func NewDecoder(input []byte) *Decoder {
	return &Decoder{
		reader: nil,
		buf:    input,
		head:   0,
		tail:   len(input),
	}
}

// NewDecoderFromReader creates an Decoder instance from io.Reader
func NewDecoderFromReader(reader io.Reader, bufSize int) *Decoder {
	if bufSize < 32 {
		bufSize = defaultBufferSize
	}
	return &Decoder{
		reader: reader,
		buf:    make([]byte, bufSize),
		head:   0,
		tail:   0,
	}
}

func (dec *Decoder) fastDecode(p interface{}, tag byte) bool {
	switch pv := p.(type) {
	case *bool:
		*pv = dec.decodeBool(boolType, tag)
	case *int:
		*pv = dec.decodeInt(intType, tag)
	case *int8:
		*pv = dec.decodeInt8(int8Type, tag)
	case *int16:
		*pv = dec.decodeInt16(int16Type, tag)
	case *int32:
		*pv = dec.decodeInt32(int32Type, tag)
	case *int64:
		*pv = dec.decodeInt64(int64Type, tag)
	case *uint:
		*pv = dec.decodeUint(uintType, tag)
	case *uint8:
		*pv = dec.decodeUint8(uint8Type, tag)
	case *uint16:
		*pv = dec.decodeUint16(uint16Type, tag)
	case *uint32:
		*pv = dec.decodeUint32(uint32Type, tag)
	case *uint64:
		*pv = dec.decodeUint64(uint64Type, tag)
	case *uintptr:
		*pv = dec.decodeUintptr(uintptrType, tag)
	case *float32:
		*pv = dec.decodeFloat32(float32Type, tag)
	case *float64:
		*pv = dec.decodeFloat64(float64Type, tag)
	case *complex64:
		*pv = dec.decodeComplex64(complex64Type, tag)
	case *complex128:
		*pv = dec.decodeComplex128(complex128Type, tag)
	case *interface{}:
		*pv = dec.decodeInterface(interfaceType, tag)
	case *[]byte:
		*pv = dec.decodeBytes(bytesType, tag)
	case *string:
		*pv = dec.decodeString(stringType, tag)
	case *time.Time:
		*pv = dec.decodeTime(timeType, tag)
	case *uuid.UUID:
		*pv = dec.decodeUUID(uuidType, tag)
	case *big.Int:
		*pv = dec.decodeBigIntValue(bigIntValueType, tag)
	case *big.Float:
		*pv = dec.decodeBigFloatValue(bigFloatValueType, tag)
	case *big.Rat:
		*pv = dec.decodeBigRatValue(bigRatValueType, tag)
	default:
		return false
	}
	return true
}

func (dec *Decoder) fastDecodePtr(p interface{}, tag byte) bool {
	switch pv := p.(type) {
	case **bool:
		*pv = dec.decodeBoolPtr(boolPtrType, tag)
	case **int:
		*pv = dec.decodeIntPtr(intPtrType, tag)
	case **int8:
		*pv = dec.decodeInt8Ptr(int8PtrType, tag)
	case **int16:
		*pv = dec.decodeInt16Ptr(int16PtrType, tag)
	case **int32:
		*pv = dec.decodeInt32Ptr(int32PtrType, tag)
	case **int64:
		*pv = dec.decodeInt64Ptr(int64PtrType, tag)
	case **uint:
		*pv = dec.decodeUintPtr(uintPtrType, tag)
	case **uint8:
		*pv = dec.decodeUint8Ptr(uint8PtrType, tag)
	case **uint16:
		*pv = dec.decodeUint16Ptr(uint16PtrType, tag)
	case **uint32:
		*pv = dec.decodeUint32Ptr(uint32PtrType, tag)
	case **uint64:
		*pv = dec.decodeUint64Ptr(uint64PtrType, tag)
	case **uintptr:
		*pv = dec.decodeUintptrPtr(uintptrPtrType, tag)
	case **float32:
		*pv = dec.decodeFloat32Ptr(float32PtrType, tag)
	case **float64:
		*pv = dec.decodeFloat64Ptr(float64PtrType, tag)
	case **complex64:
		*pv = dec.decodeComplex64Ptr(complex64PtrType, tag)
	case **complex128:
		*pv = dec.decodeComplex128Ptr(complex128PtrType, tag)
	case **interface{}:
		*pv = dec.decodeInterfacePtr(interfacePtrType, tag)
	case **[]byte:
		*pv = dec.decodeBytesPtr(bytesPtrType, tag)
	case **string:
		*pv = dec.decodeStringPtr(stringPtrType, tag)
	case **time.Time:
		*pv = dec.decodeTimePtr(timePtrType, tag)
	case **uuid.UUID:
		*pv = dec.decodeUUIDPtr(uuidPtrType, tag)
	case **big.Int:
		*pv = dec.decodeBigInt(bigIntType, tag)
	case **big.Float:
		*pv = dec.decodeBigFloat(bigFloatType, tag)
	case **big.Rat:
		*pv = dec.decodeBigRat(bigRatType, tag)
	default:
		return false
	}
	return true
}

func (dec *Decoder) decode(p interface{}, tag byte) {
	switch tag {
	case TagRef:

	case TagClass:
		dec.ReadStruct()
		dec.Decode(p)
		return
	case TagError:
		dec.Error = DecodeError(dec.decodeString(stringType, dec.NextByte()))
		return
	}
	if dec.fastDecode(p, tag) {
		return
	}
	t := reflect.TypeOf(p).Elem()
	switch t.Kind() {
	case reflect.Map:
		if dec.fastDecodeMap(t, p, tag) {
			return
		}
	case reflect.Ptr:
		if dec.fastDecodePtr(p, tag) {
			return
		}
	case reflect.Slice:
		if dec.fastDecodeSlice(p, tag) {
			return
		}
	}
	if valdec := GetValueDecoder(t); valdec != nil {
		valdec.Decode(dec, p, tag)
	}
}

// Decode a data from the Decoder
func (dec *Decoder) Decode(p interface{}) {
	dec.decode(p, dec.NextByte())
}

// Reset the value reference and struct type reference
func (dec *Decoder) Reset() *Decoder {
	if !dec.IsSimple() {
		dec.refer.Reset()
	}
	dec.ref = dec.ref[:0]
	return dec
}

// Simple resets the decoder to simple mode or not
func (dec *Decoder) Simple(simple bool) *Decoder {
	if simple {
		dec.refer = nil
	} else {
		dec.refer = &decoderRefer{}
	}
	dec.ref = dec.ref[:0]
	return dec
}

// IsSimple returns the decoder is in simple mode or not
func (dec *Decoder) IsSimple() bool {
	return nil == dec.refer
}

// AddReference adds o to the reference
func (dec *Decoder) AddReference(o interface{}) {
	if !dec.IsSimple() {
		dec.refer.Add(o)
	}
}

// SetReference sets o to the reference at index i
func (dec *Decoder) SetReference(i int, o interface{}) {
	if !dec.IsSimple() {
		dec.refer.Set(i, o)
	}
}

// LastReferenceIndex returns the last index of the reference
func (dec *Decoder) LastReferenceIndex() int {
	if !dec.IsSimple() {
		dec.refer.Last()
	}
	return -1
}

// ResetReader reuse decoder instance by specifying another reader
func (dec *Decoder) ResetReader(reader io.Reader) *Decoder {
	dec.reader = reader
	dec.head = 0
	dec.tail = 0
	return dec
}

// ResetBytes reuse decoder instance by specifying another byte array as input
func (dec *Decoder) ResetBytes(input []byte) *Decoder {
	dec.reader = nil
	dec.buf = input
	dec.head = 0
	dec.tail = len(input)
	return dec
}

// NextByte reads and returns the next byte from the dec. If no byte is available, it returns 0.
func (dec *Decoder) NextByte() (b byte) {
	if (dec.head == dec.tail) && !dec.loadMore() {
		return 0
	}
	b = dec.buf[dec.head]
	dec.head++
	return b
}

// Skip the next byte from the dec.
func (dec *Decoder) Skip() {
	if (dec.head == dec.tail) && !dec.loadMore() {
		return
	}
	dec.head++
}

func (dec *Decoder) next(n int) (data []byte, safe bool) {
	if (dec.head == dec.tail) && !dec.loadMore() {
		return nil, true
	}
	remain := dec.tail - dec.head
	if remain >= n {
		data = dec.buf[dec.head : dec.head+n]
		dec.head += n
		return data, false
	}
	safe = true
	data = make([]byte, remain, n)
	copy(data, dec.buf[dec.head:dec.tail])
	n -= remain
	for {
		if !dec.loadMore() {
			return
		}
		if dec.tail >= n {
			dec.head = n
			data = append(data, dec.buf[0:n]...)
			return
		}
		data = append(data, dec.buf[:dec.tail]...)
		n -= dec.tail
	}
}

// UnsafeNext returns a slice containing the next n bytes from the buffer,
// advancing the buffer as if the bytes had been returned by Read.
// If there are fewer than n bytes in the buffer, Next returns the entire buffer.
// The returned slice is only valid until the next call to a read method.
func (dec *Decoder) UnsafeNext(n int) (data []byte) {
	data, _ = dec.next(n)
	return
}

// Next returns a slice containing the next n bytes from the buffer,
// advancing the buffer as if the bytes had been returned by Read.
// If there are fewer than n bytes in the buffer, Next returns the entire buffer.
// The returned slice is always valid.
func (dec *Decoder) Next(n int) []byte {
	data, safe := dec.next(n)
	if safe {
		return data
	}
	result := make([]byte, len(data))
	copy(result, data)
	return result
}

// Remains reads and returns all bytes data in this iter that has not been read.
func (dec *Decoder) Remains() (data []byte) {
	if (dec.head == dec.tail) && !dec.loadMore() {
		return
	}
	for {
		data = append(data, dec.buf[dec.head:dec.tail]...)
		if !dec.loadMore() {
			return
		}
	}
}

func (dec *Decoder) until(delim byte) (data []byte, safe bool) {
	if (dec.head == dec.tail) && !dec.loadMore() {
		return nil, true
	}
	if i := bytes.IndexByte(dec.buf[dec.head:dec.tail], delim); i >= 0 {
		data = dec.buf[dec.head : dec.head+i]
		dec.head += i + 1
		return data, false
	}
	safe = true
	for {
		data = append(data, dec.buf[dec.head:dec.tail]...)
		if !dec.loadMore() {
			return
		}
		if i := bytes.IndexByte(dec.buf[dec.head:dec.tail], delim); i >= 0 {
			data = append(data, dec.buf[dec.head:dec.head+i]...)
			dec.head += i + 1
			return
		}
	}
}

// UnsafeUntil reads until the first occurrence of delim in the input,
// returning a slice containing the data up to and not including the delimiter.
// The returned slice is only valid until the next call to a read method.
func (dec *Decoder) UnsafeUntil(delim byte) (data []byte) {
	data, _ = dec.until(delim)
	return
}

// Until reads until the first occurrence of delim in the input,
// returning a slice containing the data up to and not including the delimiter.
// The returned slice is always valid.
func (dec *Decoder) Until(delim byte) []byte {
	data, safe := dec.until(delim)
	if safe {
		return data
	}
	result := make([]byte, len(data))
	copy(result, data)
	return result
}

func (dec *Decoder) loadMore() bool {
	if dec.reader == nil {
		dec.head = dec.tail
		if dec.Error == nil {
			dec.Error = io.EOF
		}
		return false
	} else if dec.buf == nil {
		dec.buf = make([]byte, defaultBufferSize)
	}
	for {
		n, err := dec.reader.Read(dec.buf)
		dec.head = 0
		dec.tail = n
		if n > 0 {
			return true
		}
		if err != nil {
			if dec.Error == nil {
				dec.Error = err
			}
			return false
		}
	}
}

func (dec *Decoder) decodeStringError(s string, typeName string) {
	if dec.Error == nil {
		dec.Error = DecodeError(`hprose/encoding: can not parse "` + s + `" to ` + typeName)
	}
}

func (dec *Decoder) decodeError(t reflect.Type, tag byte) {
	if dec.Error == nil {
		var iface interface{}
		dec.decode(&iface, tag)
		if dec.Error != nil {
			return
		}
		if v, ok := iface.(float64); ok {
			switch {
			case math.IsNaN(v):
				dec.Error = DecodeError("hprose/encoding: can not parse NaN to " + t.String())
				return
			case math.IsInf(v, 1):
				dec.Error = DecodeError("hprose/encoding: can not parse +Inf to " + t.String())
				return
			case math.IsInf(v, -1):
				dec.Error = DecodeError("hprose/encoding: can not parse -Inf to " + t.String())
				return
			}
		}
		dec.Error = CastError{
			Source:      reflect.TypeOf(iface),
			Destination: t,
		}
	}
}
