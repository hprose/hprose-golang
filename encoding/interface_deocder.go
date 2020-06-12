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
	"reflect"
)

func (dec *Decoder) decodeInterface(tag byte) interface{} {
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
	}
	if dec.Error == nil {
		dec.Error = DecodeError(fmt.Sprintf("hprose/encoding: invalid tag '%s'(0x%x)", string(tag), tag))
	}
	return nil
}

// interfaceDecoder is the implementation of ValueDecoder for interface{}.
type interfaceDecoder struct {
	t reflect.Type
}

func (valdec interfaceDecoder) decode(dec *Decoder, pv *interface{}, tag byte) {
	if i := dec.decodeInterface(tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec interfaceDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, p.(*interface{}), tag)
}

func (valdec interfaceDecoder) Type() reflect.Type {
	return valdec.t
}

// interfacePtrDecoder is the implementation of ValueDecoder for *interface{}.
type interfacePtrDecoder struct {
	t reflect.Type
}

func (valdec interfacePtrDecoder) decode(dec *Decoder, pv **interface{}, tag byte) {
	if i := dec.decodeInterface(tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec interfacePtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, p.(**interface{}), tag)
}

func (valdec interfacePtrDecoder) Type() reflect.Type {
	return valdec.t
}

var (
	ifdec  = interfaceDecoder{reflect.TypeOf(interface{}(nil))}
	pifdec = interfacePtrDecoder{reflect.TypeOf((*interface{})(nil))}
)

func init() {
	RegisterValueDecoder(ifdec)
	RegisterValueDecoder(pifdec)
}
