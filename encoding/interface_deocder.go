/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/interface_decoder.go                            |
|                                                          |
| LastModified: Jun 2, 2020                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"fmt"
	"math"
	"math/big"
)

// interfaceDecoder is the implementation of ValueDecoder for interface{}.
type interfaceDecoder struct{}

var ifacedec interfaceDecoder

func (valdec interfaceDecoder) decode(dec *Decoder, tag byte) interface{} {
	if i := intDigits[tag]; i != invalidDigit {
		return int(i)
	}
	switch tag {
	case TagNull:
		return nil
	case TagEmpty:
		return ""
	case TagFalse:
		return false
	case TagTrue:
		return true
	case TagInteger:
		return dec.ReadInt()
	case TagLong:
		switch dec.LongType {
		case LongTypeInt64:
			return dec.ReadInt64()
		case LongTypeUint64:
			return dec.ReadUint64()
		default:
			return dec.ReadBigInt()
		}
	case TagNaN:
		switch dec.RealType {
		case RealTypeFloat32:
			return float32(math.NaN())
		default:
			return math.NaN()
		}
	case TagInfinity:
		var f float64
		if dec.NextByte() == TagNeg {
			f = math.Inf(-1)
		} else {
			f = math.Inf(1)
		}
		switch dec.RealType {
		case RealTypeFloat32:
			return float32(f)
		case RealTypeFloat64:
			return f
		default:
			return big.NewFloat(f)
		}
	case TagDouble:
		switch dec.RealType {
		case RealTypeFloat32:
			return dec.ReadFloat32()
		case RealTypeFloat64:
			return dec.ReadFloat64()
		default:
			return dec.ReadBigFloat()
		}
	case TagUTF8Char:
		return dec.readSafeString(1)
	case TagString:
		return dec.ReadString()
	}
	if dec.Error == nil {
		dec.Error = DecodeError(fmt.Sprintf("hprose/encoding: invalid tag '%s'(0x%x)", string(tag), tag))
	}
	return nil
}

func (valdec interfaceDecoder) decodeValue(dec *Decoder, pv *interface{}, tag byte) {
	if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec interfaceDecoder) decodePtr(dec *Decoder, pv **interface{}, tag byte) {
	if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec interfaceDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch pv := p.(type) {
	case *interface{}:
		valdec.decodeValue(dec, pv, tag)
	case **interface{}:
		valdec.decodePtr(dec, pv, tag)
	}
}
