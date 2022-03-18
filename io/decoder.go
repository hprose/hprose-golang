/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/decoder.go                                            |
|                                                          |
| LastModified: Mar 18, 2022                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"bytes"
	"io"
	"math"
	"math/big"
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/modern-go/reflect2"
)

// LongType represents the default type for decode long integer.
type LongType int8

const (
	// LongTypeInt represents the default type is int.
	LongTypeInt LongType = iota
	// LongTypeUint represents the default type is uint.
	LongTypeUint
	// LongTypeInt64 represents the default type is int64.
	LongTypeInt64
	// LongTypeUint64 represents the default type is uint64.
	LongTypeUint64
	// LongTypeBigInt represents the default type is *big.Int.
	LongTypeBigInt
)

// RealType represents the default type for decode real number.
type RealType int8

const (
	// RealTypeFloat64 represents the default type is float64.
	RealTypeFloat64 RealType = iota
	// RealTypeFloat32 represents the default type is float32.
	RealTypeFloat32
	// RealTypeBigFloat represents the default type is *big.Float.
	RealTypeBigFloat
)

// MapType represents the default type for decode map.
type MapType int8

const (
	// MapTypeIIMap represents the default type is map[interface{}]interface{}.
	MapTypeIIMap MapType = iota
	// MapTypeSIMap represents the default type is map[string]interface{}.
	MapTypeSIMap
)

const defaultBufferSize = 256

// Decoder is a io.Reader like object, with hprose specific read functions.
// Error is not returned as return value, but stored as Error member on this decoder instance.
type Decoder struct {
	reader io.Reader
	buf    []byte
	head   int
	tail   int
	simple bool
	refer  decoderRefer
	ref    []structInfo
	Error  error
	LongType
	RealType
	MapType
}

// NewDecoder creates an Decoder instance from byte array.
func NewDecoder(input []byte) *Decoder {
	return &Decoder{
		reader: nil,
		buf:    input,
		simple: true,
		head:   0,
		tail:   len(input),
	}
}

// NewDecoderFromReader creates an Decoder instance from io.Reader.
func NewDecoderFromReader(reader io.Reader, bufSize ...int) *Decoder {
	size := defaultBufferSize
	if len(bufSize) > 0 && bufSize[0] >= defaultBufferSize {
		size = bufSize[0]
	}
	return &Decoder{
		reader: reader,
		buf:    make([]byte, size),
		simple: true,
		head:   0,
		tail:   0,
	}
}

func (dec *Decoder) fastDecode(p interface{}, tag byte) bool {
	switch pv := p.(type) {
	case *bool:
		dec.decodeBool(boolType, tag, pv)
	case *int:
		dec.decodeInt(intType, tag, pv)
	case *int8:
		dec.decodeInt8(int8Type, tag, pv)
	case *int16:
		dec.decodeInt16(int16Type, tag, pv)
	case *int32:
		dec.decodeInt32(int32Type, tag, pv)
	case *int64:
		dec.decodeInt64(int64Type, tag, pv)
	case *uint:
		dec.decodeUint(uintType, tag, pv)
	case *uint8:
		dec.decodeUint8(uint8Type, tag, pv)
	case *uint16:
		dec.decodeUint16(uint16Type, tag, pv)
	case *uint32:
		dec.decodeUint32(uint32Type, tag, pv)
	case *uint64:
		dec.decodeUint64(uint64Type, tag, pv)
	case *uintptr:
		dec.decodeUintptr(uintptrType, tag, pv)
	case *float32:
		dec.decodeFloat32(float32Type, tag, pv)
	case *float64:
		dec.decodeFloat64(float64Type, tag, pv)
	case *complex64:
		dec.decodeComplex64(complex64Type, tag, pv)
	case *complex128:
		dec.decodeComplex128(complex128Type, tag, pv)
	case *interface{}:
		dec.decodeInterface(tag, pv)
	case *[]byte:
		dec.decodeBytes(bytesType, tag, pv)
	case *string:
		dec.decodeString(stringType, tag, pv)
	case *time.Time:
		dec.decodeTime(timeType, tag, pv)
	case *uuid.UUID:
		dec.decodeUUID(uuidType, tag, pv)
	case *big.Int:
		dec.decodeBigIntValue(bigIntValueType, tag, pv)
	case *big.Float:
		dec.decodeBigFloatValue(bigFloatValueType, tag, pv)
	case *big.Rat:
		dec.decodeBigRatValue(bigRatValueType, tag, pv)
	default:
		return false
	}
	return true
}

func (dec *Decoder) fastDecodePtr(p interface{}, tag byte) bool {
	switch pv := p.(type) {
	case **bool:
		dec.decodeBoolPtr(boolPtrType, tag, pv)
	case **int:
		dec.decodeIntPtr(intPtrType, tag, pv)
	case **int8:
		dec.decodeInt8Ptr(int8PtrType, tag, pv)
	case **int16:
		dec.decodeInt16Ptr(int16PtrType, tag, pv)
	case **int32:
		dec.decodeInt32Ptr(int32PtrType, tag, pv)
	case **int64:
		dec.decodeInt64Ptr(int64PtrType, tag, pv)
	case **uint:
		dec.decodeUintPtr(uintPtrType, tag, pv)
	case **uint8:
		dec.decodeUint8Ptr(uint8PtrType, tag, pv)
	case **uint16:
		dec.decodeUint16Ptr(uint16PtrType, tag, pv)
	case **uint32:
		dec.decodeUint32Ptr(uint32PtrType, tag, pv)
	case **uint64:
		dec.decodeUint64Ptr(uint64PtrType, tag, pv)
	case **uintptr:
		dec.decodeUintptrPtr(uintptrPtrType, tag, pv)
	case **float32:
		dec.decodeFloat32Ptr(float32PtrType, tag, pv)
	case **float64:
		dec.decodeFloat64Ptr(float64PtrType, tag, pv)
	case **complex64:
		dec.decodeComplex64Ptr(complex64PtrType, tag, pv)
	case **complex128:
		dec.decodeComplex128Ptr(complex128PtrType, tag, pv)
	case **interface{}:
		dec.decodeInterfacePtr(tag, pv)
	case **[]byte:
		dec.decodeBytesPtr(bytesPtrType, tag, pv)
	case **string:
		dec.decodeStringPtr(stringPtrType, tag, pv)
	case **time.Time:
		dec.decodeTimePtr(timePtrType, tag, pv)
	case **uuid.UUID:
		dec.decodeUUIDPtr(uuidPtrType, tag, pv)
	case **big.Int:
		dec.decodeBigInt(bigIntType, tag, pv)
	case **big.Float:
		dec.decodeBigFloat(bigFloatType, tag, pv)
	case **big.Rat:
		dec.decodeBigRat(bigRatType, tag, pv)
	default:
		return false
	}
	return true
}

func (dec *Decoder) decode(p interface{}, tag byte) {
	if dec.fastDecode(p, tag) {
		return
	}
	t := reflect.TypeOf(p).Elem()
	if t.Kind() == reflect.Ptr && dec.fastDecodePtr(p, tag) {
		return
	}
	if valdec := getValueDecoder(t); valdec != nil {
		valdec.Decode(dec, p, tag)
	}
}

// Decode a data from the Decoder.
func (dec *Decoder) Decode(p interface{}, tag ...byte) {
	if len(tag) > 0 {
		dec.decode(p, tag[0])
	} else {
		dec.decode(p, dec.NextByte())
	}
}

// Read returns a data of the specified type from the Decoder.
func (dec *Decoder) Read(t reflect.Type, tag ...byte) (result interface{}) {
	if t == nil {
		dec.Decode(&result, tag...)
		return
	}
	t2 := reflect2.Type2(t)
	p := t2.New()
	if len(tag) > 0 {
		dec.decode(p, tag[0])
	} else {
		dec.decode(p, dec.NextByte())
	}
	return t2.Indirect(p)
}

// Reset the value reference and struct type reference.
func (dec *Decoder) Reset() *Decoder {
	if !dec.IsSimple() {
		dec.refer.Reset()
	}
	dec.ref = dec.ref[:0]
	return dec
}

// Simple resets the decoder to simple mode or not.
func (dec *Decoder) Simple(simple bool) *Decoder {
	dec.simple = simple
	dec.Reset()
	return dec
}

// IsSimple returns the decoder is in simple mode or not.
func (dec *Decoder) IsSimple() bool {
	return dec.simple
}

// AddReference adds o to the reference.
func (dec *Decoder) AddReference(o interface{}) {
	if !dec.IsSimple() {
		dec.refer.Add(o)
	}
}

// SetReference sets o to the reference at index i.
func (dec *Decoder) SetReference(i int, o interface{}) {
	if !dec.IsSimple() {
		dec.refer.Set(i, o)
	}
}

// LastReferenceIndex returns the last index of the reference.
func (dec *Decoder) LastReferenceIndex() int {
	if !dec.IsSimple() {
		dec.refer.Last()
	}
	return -1
}

// ReadReference to p.
func (dec *Decoder) ReadReference(p interface{}) {
	o := dec.refer.Read(dec.ReadInt())
	src := reflect.TypeOf(o)
	dest := reflect.TypeOf(p).Elem()
	if conv := GetConverter(src, dest); conv != nil {
		conv(dec, o, p)
	} else if dec.Error == nil {
		dec.Error = CastError{
			Source:      src,
			Destination: dest,
		}
	}
}

// ResetReader reuse decoder instance by specifying another reader.
func (dec *Decoder) ResetReader(reader io.Reader) *Decoder {
	dec.reader = reader
	dec.head = 0
	dec.tail = 0
	return dec
}

// ResetBytes reuse decoder instance by specifying another byte array as input.
func (dec *Decoder) ResetBytes(input []byte) *Decoder {
	dec.reader = nil
	dec.buf = input
	dec.head = 0
	dec.tail = len(input)
	return dec
}

// ResetBuffer of the Decoder.
func (dec *Decoder) ResetBuffer() *Decoder {
	if dec.reader == nil {
		dec.buf = nil
	} else {
		dec.reader = nil
	}
	dec.head = 0
	dec.tail = 0
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
	}
	if dec.buf == nil {
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
		dec.Error = DecodeError(`hprose/io: can not parse "` + s + `" to ` + typeName)
	}
}

func (dec *Decoder) defaultDecode(t reflect.Type, p interface{}, tag byte) {
	switch tag {
	case TagRef:
		dec.ReadReference(p)
		return
	case TagClass:
		dec.ReadStruct(t)
		dec.Decode(p)
		return
	case TagError:
		var s string
		dec.decodeString(stringType, dec.NextByte(), &s)
		dec.Error = DecodeError(s)
		return
	default:
		dec.decodeError(t, tag)
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
				dec.Error = DecodeError("hprose/io: can not parse NaN to " + t.String())
				return
			case math.IsInf(v, 1):
				dec.Error = DecodeError("hprose/io: can not parse +Inf to " + t.String())
				return
			case math.IsInf(v, -1):
				dec.Error = DecodeError("hprose/io: can not parse -Inf to " + t.String())
				return
			}
		}
		dec.Error = CastError{
			Source:      reflect.TypeOf(iface),
			Destination: t,
		}
	}
}
