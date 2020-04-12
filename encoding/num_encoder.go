/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/num_encoder.go                                  |
|                                                          |
| LastModified: Apr 12, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"math"
	"strconv"
)

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
