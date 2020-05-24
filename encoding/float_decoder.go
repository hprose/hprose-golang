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

func (valdec float32Decoder) decode(dec *Decoder, p interface{}, tag byte) float32 {
	if i := intDigits[tag]; i != invalidDigit {
		return float32(i)
	}
	switch tag {
	case TagEmpty, TagFalse:
		return 0
	case TagTrue:
		return 1
	case TagInteger:
		return float32(dec.ReadInt())
	case TagLong, TagDouble:
		return dec.ReadFloat32()
	case TagNaN:
		return float32(math.NaN())
	case TagInfinity:
		return float32(dec.readInf())
	case TagUTF8Char:
		return dec.stringToFloat32(dec.readUnsafeString(1))
	case TagString:
		return dec.stringToFloat32(dec.ReadUnsafeString())
	default:
		dec.decodeError(p, tag)
	}
	return 0
}

func (valdec float32Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if tag == TagNull {
		switch pv := p.(type) {
		case **float32:
			*pv = nil
		case *float32:
			*pv = 0
		}
		return
	}
	f := valdec.decode(dec, p, tag)
	if dec.Error != nil {
		return
	}
	switch pv := p.(type) {
	case **float32:
		*pv = &f
	case *float32:
		*pv = f
	}
}

// float64Decoder is the implementation of ValueDecoder for *float64.
type float64Decoder struct{}

var f64dec float64Decoder

func (valdec float64Decoder) decode(dec *Decoder, p interface{}, tag byte) float64 {
	if i := intDigits[tag]; i != invalidDigit {
		return float64(i)
	}
	switch tag {
	case TagEmpty, TagFalse:
		return 0
	case TagTrue:
		return 1
	case TagInteger:
		return float64(dec.ReadInt())
	case TagLong, TagDouble:
		return dec.ReadFloat64()
	case TagNaN:
		return math.NaN()
	case TagInfinity:
		return dec.readInf()
	case TagUTF8Char:
		return dec.stringToFloat64(dec.readUnsafeString(1))
	case TagString:
		return dec.stringToFloat64(dec.ReadUnsafeString())
	default:
		dec.decodeError(p, tag)
	}
	return 0
}

func (valdec float64Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if tag == TagNull {
		switch pv := p.(type) {
		case **float64:
			*pv = nil
		case *float64:
			*pv = 0
		}
		return
	}
	f := valdec.decode(dec, p, tag)
	if dec.Error != nil {
		return
	}
	switch pv := p.(type) {
	case **float64:
		*pv = &f
	case *float64:
		*pv = f
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

func (dec *Decoder) readInf() float64 {
	if dec.NextByte() == TagNeg {
		return math.Inf(-1)
	}
	return math.Inf(1)
}
