/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/float_decoder.go                                |
|                                                          |
| LastModified: Apr 25, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"math"
	"strconv"
)

// float32Decoder is the implementation of ValueDecoder for float32.
type float32Decoder struct{}

var f32dec float32Decoder

func (float32Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if pv, ok := p.(*float32); ok {
		if i := intDigits[tag]; i != invalidDigit {
			*pv = float32(i)
			return
		}
		switch tag {
		case TagNull, TagEmpty, TagFalse:
			*pv = 0
		case TagTrue:
			*pv = 1
		case TagInteger:
			*pv = float32(dec.ReadInt())
		case TagLong, TagDouble:
			*pv = dec.ReadFloat32()
		case TagNaN:
			*pv = float32(math.NaN())
		case TagInfinity:
			if dec.NextByte() == TagNeg {
				*pv = float32(math.Inf(-1))
			} else {
				*pv = float32(math.Inf(1))
			}
		case TagUTF8Char:
			*pv = dec.stringToFloat32(dec.readUnsafeString(1))
		case TagString:
			*pv = dec.stringToFloat32(dec.ReadUnsafeString())
		default:
			dec.decodeError(p, tag)
		}
	}
}

// float64Decoder is the implementation of ValueDecoder for *float64.
type float64Decoder struct{}

var f64dec float64Decoder

func (float64Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if pv, ok := p.(*float64); ok {
		if i := intDigits[tag]; i != invalidDigit {
			*pv = float64(i)
			return
		}
		switch tag {
		case TagNull, TagEmpty, TagFalse:
			*pv = 0
		case TagTrue:
			*pv = 1
		case TagInteger:
			*pv = float64(dec.ReadInt())
		case TagLong, TagDouble:
			*pv = dec.ReadFloat64()
		case TagNaN:
			*pv = math.NaN()
		case TagInfinity:
			if dec.NextByte() == TagNeg {
				*pv = math.Inf(-1)
			} else {
				*pv = math.Inf(1)
			}
		case TagUTF8Char:
			*pv = dec.stringToFloat64(dec.readUnsafeString(1))
		case TagString:
			*pv = dec.stringToFloat64(dec.ReadUnsafeString())
		default:
			dec.decodeError(p, tag)
		}
	}
}

func (dec *Decoder) stringToFloat32(s string) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		dec.Error = err
	}
	return float32(f)
}

func (dec *Decoder) stringToFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		dec.Error = err
	}
	return f
}
