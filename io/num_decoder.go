/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/int_decoder.go                                        |
|                                                          |
| LastModified: Mar 5, 2022                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"strconv"

	"github.com/hprose/hprose-golang/v3/internal/convert"
)

const invalidDigit = uint64(0xff)

var intDigits [256]uint64

func init() {
	for i := 0; i < len(intDigits); i++ {
		intDigits[i] = invalidDigit
	}
	for i := uint64('0'); i <= uint64('9'); i++ {
		intDigits[i] = i - uint64('0')
	}
}

// ReadInt8 reads int8.
func (dec *Decoder) ReadInt8() (value int8) {
	return int8(dec.ReadInt64())
}

// ReadUint8 reads uint8.
func (dec *Decoder) ReadUint8() (value uint8) {
	return uint8(dec.ReadUint64())
}

// ReadInt16 reads int16.
func (dec *Decoder) ReadInt16() (value int16) {
	return int16(dec.ReadInt64())
}

// ReadUint16 reads uint16.
func (dec *Decoder) ReadUint16() (value uint16) {
	return uint16(dec.ReadUint64())
}

// ReadInt32 reads int32.
func (dec *Decoder) ReadInt32() (value int32) {
	return int32(dec.ReadInt64())
}

// ReadUint32 reads uint32.
func (dec *Decoder) ReadUint32() (value uint32) {
	return uint32(dec.ReadUint64())
}

// ReadInt64 reads int64.
func (dec *Decoder) ReadInt64() (value int64) {
	c := dec.NextByte()
	if c == '-' {
		return -int64(dec.readUint64(dec.NextByte()))
	}
	return int64(dec.readUint64(c))
}

// ReadUint64 reads uint64.
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

// ReadInt reads int.
func (dec *Decoder) ReadInt() (value int) {
	return int(dec.ReadInt64())
}

// ReadUint reads uint.
func (dec *Decoder) ReadUint() (value uint) {
	return uint(dec.ReadUint64())
}

// ReadFloat32 reads float32.
func (dec *Decoder) ReadFloat32() (value float32) {
	f, err := strconv.ParseFloat(convert.ToUnsafeString(dec.UnsafeUntil(TagSemicolon)), 32)
	if dec.Error == nil && err != nil {
		dec.Error = err
	}
	return float32(f)
}

// ReadFloat64 reads float64.
func (dec *Decoder) ReadFloat64() (value float64) {
	f, err := strconv.ParseFloat(convert.ToUnsafeString(dec.UnsafeUntil(TagSemicolon)), 64)
	if dec.Error == nil && err != nil {
		dec.Error = err
	}
	return f
}
