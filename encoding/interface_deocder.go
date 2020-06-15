/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/interface_decoder.go                            |
|                                                          |
| LastModified: Jun 15, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"fmt"
	"math"
	"math/big"
	"reflect"

	"github.com/modern-go/reflect2"
)

func (dec *Decoder) decodeInterface(t reflect.Type, tag byte) interface{} {
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
		case RealTypeFloat64:
			return math.NaN()
		default:
			dec.Error = DecodeError("hprose/encoding: can not parse NaN to *big.Float")
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
	case TagBytes:
		return dec.ReadBytes()
	}
	if dec.Error == nil {
		dec.Error = DecodeError(fmt.Sprintf("hprose/encoding: invalid tag '%s'(0x%x)", string(tag), tag))
	}
	return nil
}

func (dec *Decoder) decodeInterfacePtr(t reflect.Type, tag byte) *interface{} {
	if tag == TagNull {
		return nil
	}
	i := dec.decodeInterface(t, tag)
	return &i
}

// interfaceDecoder is the implementation of ValueDecoder for interface{}.
type interfaceDecoder struct {
	t reflect.Type
}

func (valdec interfaceDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*interface{})(reflect2.PtrOf(p)) = dec.decodeInterface(valdec.t, tag)
}

func (valdec interfaceDecoder) Type() reflect.Type {
	return valdec.t
}

// interfacePtrDecoder is the implementation of ValueDecoder for *interface{}.
type interfacePtrDecoder struct {
	t reflect.Type
}

func (valdec interfacePtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**interface{})(reflect2.PtrOf(p)) = dec.decodeInterfacePtr(valdec.t, tag)
}

func (valdec interfacePtrDecoder) Type() reflect.Type {
	return valdec.t
}
