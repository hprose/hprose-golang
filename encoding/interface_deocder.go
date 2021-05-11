/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/interface_decoder.go                            |
|                                                          |
| LastModified: May 11, 2021                               |
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

func (dec *Decoder) decodeLongAsInterface() interface{} {
	switch dec.LongType {
	case LongTypeInt64:
		return dec.ReadInt64()
	case LongTypeUint64:
		return dec.ReadUint64()
	}
	return dec.ReadBigInt()
}

func (dec *Decoder) decodeNaNAsInterface() interface{} {
	switch dec.RealType {
	case RealTypeFloat32:
		return float32(math.NaN())
	case RealTypeFloat64:
		return math.NaN()
	}
	dec.Error = DecodeError("hprose/encoding: can not parse NaN to *big.Float")
	return nil
}

func (dec *Decoder) decodeInfinityAsInterface() interface{} {
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
	}
	return big.NewFloat(f)
}

func (dec *Decoder) decodeDoubleAsInterface() interface{} {
	switch dec.RealType {
	case RealTypeFloat32:
		return dec.ReadFloat32()
	case RealTypeFloat64:
		return dec.ReadFloat64()
	}
	return dec.ReadBigFloat()
}

func (dec *Decoder) decodeListAsInterface(tag byte) interface{} {
	var result []interface{}
	ifsdec.Decode(dec, &result, tag)
	return result
}

func (dec *Decoder) decodeMapAsInterface(tag byte) interface{} {
	if dec.MapType == MapTypeIIMap {
		var result map[interface{}]interface{}
		ififmdec.Decode(dec, &result, tag)
		return result
	}
	var result map[string]interface{}
	sifmdec.Decode(dec, &result, tag)
	return result
}

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
		return dec.decodeLongAsInterface()
	case TagNaN:
		return dec.decodeNaNAsInterface()
	case TagInfinity:
		return dec.decodeInfinityAsInterface()
	case TagDouble:
		return dec.decodeDoubleAsInterface()
	case TagTime:
		return dec.ReadTime()
	case TagDate:
		return dec.ReadDateTime()
	case TagGUID:
		return dec.ReadUUID()
	case TagUTF8Char:
		return dec.readSafeString(1)
	case TagString:
		return dec.ReadString()
	case TagBytes:
		return dec.ReadBytes()
	case TagList:
		return dec.decodeListAsInterface(tag)
	case TagMap:
		return dec.decodeMapAsInterface(tag)
	case TagClass:
		dec.ReadStruct()
		return dec.Read(interfaceType)
	case TagObject:
		return dec.ReadObject()
	}
	if dec.Error == nil {
		dec.Error = DecodeError(fmt.Sprintf("hprose/encoding: invalid tag '%s'(0x%x)", string(tag), tag))
	}
	return nil
}

func (dec *Decoder) decodeInterfacePtr(tag byte) *interface{} {
	if tag == TagNull {
		return nil
	}
	i := dec.decodeInterface(tag)
	return &i
}

// interfaceDecoder is the implementation of ValueDecoder for interface{}.
type interfaceDecoder struct{}

func (valdec interfaceDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*interface{})(reflect2.PtrOf(p)) = dec.decodeInterface(tag)
}

func (valdec interfaceDecoder) Type() reflect.Type {
	return interfaceType
}

// interfacePtrDecoder is the implementation of ValueDecoder for *interface{}.
type interfacePtrDecoder struct{}

func (valdec interfacePtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**interface{})(reflect2.PtrOf(p)) = dec.decodeInterfacePtr(tag)
}

func (valdec interfacePtrDecoder) Type() reflect.Type {
	return interfacePtrType
}
