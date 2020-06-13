/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/int_decoder.go                                  |
|                                                          |
| LastModified: Jun 13, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"strconv"
)

const invalidDigit = uint64(0xff)

var intDigits []uint64

func init() {
	intDigits = make([]uint64, 256)
	for i := 0; i < len(intDigits); i++ {
		intDigits[i] = invalidDigit
	}
	for i := uint64('0'); i <= uint64('9'); i++ {
		intDigits[i] = i - uint64('0')
	}
}

// ReadInt8 reads int8
func (dec *Decoder) ReadInt8() (value int8) {
	return int8(dec.ReadInt64())
}

// ReadUint8 reads uint8
func (dec *Decoder) ReadUint8() (value uint8) {
	return uint8(dec.ReadUint64())
}

// ReadInt16 reads int16
func (dec *Decoder) ReadInt16() (value int16) {
	return int16(dec.ReadInt64())
}

// ReadUint16 reads uint16
func (dec *Decoder) ReadUint16() (value uint16) {
	return uint16(dec.ReadUint64())
}

// ReadInt32 reads int32
func (dec *Decoder) ReadInt32() (value int32) {
	return int32(dec.ReadInt64())
}

// ReadUint32 reads uint32
func (dec *Decoder) ReadUint32() (value uint32) {
	return uint32(dec.ReadUint64())
}

// ReadInt64 reads int64
func (dec *Decoder) ReadInt64() (value int64) {
	c := dec.NextByte()
	if c == '-' {
		return -int64(dec.readUint64(dec.NextByte()))
	}
	return int64(dec.readUint64(c))
}

// ReadUint64 reads uint64
func (dec *Decoder) ReadUint64() (value uint64) {
	c := dec.NextByte()
	if c == '-' {
		return uint64(-int64(dec.readUint64(dec.NextByte())))
	}
	return dec.readUint64(c)
}

func (dec *Decoder) readUint64(c byte) (value uint64) {
	i := intDigits[c]
	if i == invalidDigit {
		return
	}
	value = i
	if dec.tail-dec.head > 9 {
		p := dec.head
		i1 := intDigits[dec.buf[p]]
		p++
		if i1 == invalidDigit {
			dec.head = p
			return value
		}
		i2 := intDigits[dec.buf[p]]
		p++
		if i2 == invalidDigit {
			dec.head = p
			return value*10 + i1
		}
		i3 := intDigits[dec.buf[p]]
		p++
		if i3 == invalidDigit {
			dec.head = p
			return value*100 + i1*10 + i2
		}
		i4 := intDigits[dec.buf[p]]
		p++
		if i4 == invalidDigit {
			dec.head = p
			return value*1000 + i1*100 + i2*10 + i3
		}
		i5 := intDigits[dec.buf[p]]
		p++
		if i5 == invalidDigit {
			dec.head = p
			return value*10000 + i1*1000 + i2*100 + i3*10 + i4
		}
		i6 := intDigits[dec.buf[p]]
		p++
		if i6 == invalidDigit {
			dec.head = p
			return value*100000 + i1*10000 + i2*1000 + i3*100 + i4*10 + i5
		}
		i7 := intDigits[dec.buf[p]]
		p++
		if i7 == invalidDigit {
			dec.head = p
			return value*1000000 + i1*100000 + i2*10000 + i3*1000 + i4*100 + i5*10 + i6
		}
		i8 := intDigits[dec.buf[p]]
		p++
		if i8 == invalidDigit {
			dec.head = p
			return value*10000000 + i1*1000000 + i2*100000 + i3*10000 + i4*1000 + i5*100 + i6*10 + i7
		}
		i9 := intDigits[dec.buf[p]]
		p++
		if i9 == invalidDigit {
			dec.head = p
			return value*100000000 + i1*10000000 + i2*1000000 + i3*100000 + i4*10000 + i5*1000 + i6*100 + i7*10 + i8
		}
		value = value*1000000000 + i1*100000000 + i2*10000000 + i3*1000000 + i4*100000 + i5*10000 + i6*1000 + i7*100 + i8*10 + i9
		i10 := intDigits[dec.buf[p]]
		if i10 == invalidDigit {
			dec.head = p + 1
			return
		}
		dec.head = p
	}
	for {
		for p := dec.head; p < dec.tail; p++ {
			i = intDigits[dec.buf[p]]
			if i == invalidDigit {
				dec.head = p + 1
				return
			}
			value = value*10 + i
		}
		if !dec.loadMore() {
			return
		}
	}
}

// ReadInt reads int
func (dec *Decoder) ReadInt() (value int) {
	return int(dec.ReadInt64())
}

// ReadUint reads uint
func (dec *Decoder) ReadUint() (value uint) {
	return uint(dec.ReadUint64())
}

// ReadFloat32 reads float32
func (dec *Decoder) ReadFloat32() (value float32) {
	f, err := strconv.ParseFloat(unsafeString(dec.UnsafeUntil(TagSemicolon)), 32)
	if dec.Error == nil && err != nil {
		dec.Error = err
	}
	return float32(f)
}

// ReadFloat64 reads float64
func (dec *Decoder) ReadFloat64() (value float64) {
	f, err := strconv.ParseFloat(unsafeString(dec.UnsafeUntil(TagSemicolon)), 64)
	if dec.Error == nil && err != nil {
		dec.Error = err
	}
	return f
}
