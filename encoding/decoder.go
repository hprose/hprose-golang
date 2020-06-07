/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/decoder.go                                      |
|                                                          |
| LastModified: Jun 2, 2020                                |
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
)

// LongType represents the default type for decode long integer
type LongType int8

const (
	// LongTypeBigInt represents the default type is *big.Int
	LongTypeBigInt LongType = iota
	// LongTypeInt64 represents the default type is int64
	LongTypeInt64
	// LongTypeUint64 represents the default type is uint64
	LongTypeUint64
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

const defaultBufferSize = 8192

// Decoder is a io.Reader like object, with hprose specific read functions.
// Error is not returned as return value, but stored as Error member on this decoder instance.
type Decoder struct {
	reader io.Reader
	buf    []byte
	head   int
	tail   int
	refer  *decoderRefer
	ref    []reflect.Type
	Error  error
	LongType
	RealType
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
	case *int:
		intdec.decodeValue(dec, pv, tag)
	case *int8:
		int8dec.decodeValue(dec, pv, tag)
	case *int16:
		int16dec.decodeValue(dec, pv, tag)
	case *int32:
		int32dec.decodeValue(dec, pv, tag)
	case *int64:
		int64dec.decodeValue(dec, pv, tag)
	case *uint:
		uintdec.decodeValue(dec, pv, tag)
	case *uint8:
		uint8dec.decodeValue(dec, pv, tag)
	case *uint16:
		uint16dec.decodeValue(dec, pv, tag)
	case *uint32:
		uint32dec.decodeValue(dec, pv, tag)
	case *uint64:
		uint64dec.decodeValue(dec, pv, tag)
	case *uintptr:
		uptrdec.decodeValue(dec, pv, tag)
	case *float32:
		f32dec.decodeValue(dec, pv, tag)
	case *float64:
		f64dec.decodeValue(dec, pv, tag)
	case *bool:
		booldec.decodeValue(dec, pv, tag)
	case *string:
		strdec.decodeValue(dec, pv, tag)
	case *complex64:
		c64dec.decodeValue(dec, pv, tag)
	case *complex128:
		c128dec.decodeValue(dec, pv, tag)
	case *[]byte:
		bytesdec.decodeValue(dec, pv, tag)
	case *interface{}:
		ifacedec.decodeValue(dec, pv, tag)
	case *big.Int:
		bigintdec.decodeValue(dec, pv, tag)
	case *big.Float:
		bigfloatdec.decodeValue(dec, pv, tag)
	case *big.Rat:
		bigratdec.decodeValue(dec, pv, tag)
	case **int:
		intdec.decodePtr(dec, pv, tag)
	case **int8:
		int8dec.decodePtr(dec, pv, tag)
	case **int16:
		int16dec.decodePtr(dec, pv, tag)
	case **int32:
		int32dec.decodePtr(dec, pv, tag)
	case **int64:
		int64dec.decodePtr(dec, pv, tag)
	case **uint:
		uintdec.decodePtr(dec, pv, tag)
	case **uint8:
		uint8dec.decodePtr(dec, pv, tag)
	case **uint16:
		uint16dec.decodePtr(dec, pv, tag)
	case **uint32:
		uint32dec.decodePtr(dec, pv, tag)
	case **uint64:
		uint64dec.decodePtr(dec, pv, tag)
	case **uintptr:
		uptrdec.decodePtr(dec, pv, tag)
	case **float32:
		f32dec.decodePtr(dec, pv, tag)
	case **float64:
		f64dec.decodePtr(dec, pv, tag)
	case **bool:
		booldec.decodePtr(dec, pv, tag)
	case **string:
		strdec.decodePtr(dec, pv, tag)
	case **complex64:
		c64dec.decodePtr(dec, pv, tag)
	case **complex128:
		c128dec.decodePtr(dec, pv, tag)
	case **[]byte:
		bytesdec.decodePtr(dec, pv, tag)
	case **interface{}:
		ifacedec.decodePtr(dec, pv, tag)
	case **big.Int:
		bigintdec.decodePtr(dec, pv, tag)
	case **big.Float:
		bigfloatdec.decodePtr(dec, pv, tag)
	case **big.Rat:
		bigratdec.decodePtr(dec, pv, tag)
	default:
		return false
	}
	return true
}

func (dec *Decoder) fastDecodeSlice(p interface{}, tag byte) bool {
	switch p.(type) {
	case *[]int:
		intsdec.Decode(dec, p, tag)
	case *[]int8:
		int8sdec.Decode(dec, p, tag)
	case *[]int16:
		int16sdec.Decode(dec, p, tag)
	case *[]int32:
		int32sdec.Decode(dec, p, tag)
	case *[]int64:
		int64sdec.Decode(dec, p, tag)
	case *[]uint:
		uintsdec.Decode(dec, p, tag)
	case *[]uint16:
		uint16sdec.Decode(dec, p, tag)
	case *[]uint32:
		uint32sdec.Decode(dec, p, tag)
	case *[]uint64:
		uint64sdec.Decode(dec, p, tag)
	case *[]uintptr:
		uptrsdec.Decode(dec, p, tag)
	case *[]bool:
		boolsdec.Decode(dec, p, tag)
	case *[]float32:
		f32sdec.Decode(dec, p, tag)
	case *[]float64:
		f64sdec.Decode(dec, p, tag)
	case *[]complex64:
		c64sdec.Decode(dec, p, tag)
	case *[]complex128:
		c128sdec.Decode(dec, p, tag)
	case *[]string:
		strsdec.Decode(dec, p, tag)
	case *[]interface{}:
		ifacesdec.Decode(dec, p, tag)
	case *[]*big.Int:
		bigintsdec.Decode(dec, p, tag)
	case *[]*big.Float:
		bigfloatsdec.Decode(dec, p, tag)
	case *[]*big.Rat:
		bigratsdec.Decode(dec, p, tag)
	default:
		return false
	}
	return true
}

func (dec *Decoder) decode(p interface{}, tag byte) {
	if dec.fastDecode(p, tag) {
		return
	}
	if dec.fastDecodeSlice(p, tag) {
		return
	}

	typ := reflect.TypeOf(p).Elem()

	switch typ.Kind() {
	case reflect.Slice:
		valdec := getOtherDecoder(typ)

		valdec.Decode(dec, p, tag)
		return
	}

	if valdec := GetValueDecoder(p); valdec != nil {
		valdec.Decode(dec, p, tag)
	}
}

// Decode a data from the Decoder
func (dec *Decoder) Decode(p interface{}) {
	dec.decode(p, dec.NextByte())
}

// Reset the value reference and struct type reference
func (dec *Decoder) Reset() {
	if dec.refer != nil {
		dec.refer.Reset()
	}
	dec.ref = dec.ref[:0]
}

// Simple resets the encoder to simple mode or not
func (dec *Decoder) Simple(simple bool) {
	if simple {
		dec.refer = nil
	} else {
		dec.refer = &decoderRefer{}
	}
	dec.ref = dec.ref[:0]
}

// AddReference adds o to the reference
func (dec *Decoder) AddReference(o interface{}) {
	if dec.refer != nil {
		dec.refer.Add(o)
	}
}

// SetReference sets o to the reference at index i
func (dec *Decoder) SetReference(i int, o interface{}) {
	if dec.refer != nil {
		dec.refer.Set(i, o)
	}
}

// LastReferenceIndex returns the last index of the reference
func (dec *Decoder) LastReferenceIndex() int {
	if dec.refer != nil {
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

// Next returns a slice containing the next n bytes from the buffer,
// advancing the buffer as if the bytes had been returned by Read.
// If there are fewer than n bytes in the buffer, Next returns the entire buffer.
// The slice is only valid until the next call to a read method.
func (dec *Decoder) Next(n int) (data []byte) {
	if (dec.head == dec.tail) && !dec.loadMore() {
		return
	}
	remain := dec.tail - dec.head
	if remain >= n {
		data = dec.buf[dec.head : dec.head+n]
		dec.head += n
		return
	}
	data = make([]byte, 0, n)
	data = append(data, dec.buf[dec.head:dec.tail]...)
	n -= remain
	for {
		if !dec.loadMore() {
			return
		}
		if dec.tail >= n {
			dec.head = n
			return append(data, dec.buf[0:n]...)
		}
		data = append(data, dec.buf[:dec.tail]...)
		n -= dec.tail
	}
}

// Remains reads and returns all bytes data in this iter that has not been read.
func (dec *Decoder) Remains() (data []byte) {
	if (dec.head == dec.tail) && !dec.loadMore() {
		return
	}
	for {
		data = append(data, dec.buf[dec.head:dec.tail]...)
		if !dec.loadMore() {
			return data
		}
	}
}

// Until reads until the first occurrence of delim in the input,
// returning a slice containing the data up to and not including the delimiter.
func (dec *Decoder) Until(delim byte) (data []byte) {
	if (dec.head == dec.tail) && !dec.loadMore() {
		return
	}
	if i := bytes.IndexByte(dec.buf[dec.head:dec.tail], delim); i >= 0 {
		data = dec.buf[dec.head : dec.head+i]
		dec.head += i + 1
		return
	}
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

func (dec *Decoder) decodeStringError(s string, typename string) {
	if dec.Error == nil {
		dec.Error = DecodeError(`hprose/encoding: can not parse "` + s + `" to ` + typename)
	}
}

func (dec *Decoder) decodeError(destType reflect.Type, tag byte) {
	if dec.Error == nil {
		var iface interface{}
		dec.decode(&iface, tag)
		if dec.Error != nil {
			return
		}
		if v, ok := iface.(float64); ok {
			switch {
			case math.IsNaN(v):
				dec.Error = DecodeError("hprose/encoding: can not parse NaN to " + destType.String())
				return
			case math.IsInf(v, 1):
				dec.Error = DecodeError("hprose/encoding: can not parse +Inf to " + destType.String())
				return
			case math.IsInf(v, -1):
				dec.Error = DecodeError("hprose/encoding: can not parse -Inf to " + destType.String())
				return
			}
		}
		dec.Error = &CastError{
			Source:      reflect.TypeOf(iface),
			Destination: destType,
		}
	}
}
