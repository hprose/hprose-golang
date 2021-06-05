/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/big_decoder.go                                        |
|                                                          |
| LastModified: Jun 5, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"math"
	"math/big"
	"reflect"

	"github.com/modern-go/reflect2"
)

var (
	bigIntZero   = big.NewInt(0)
	bigIntOne    = big.NewInt(1)
	bigFloatZero = big.NewFloat(0)
	bigFloatOne  = big.NewFloat(1)
	bigRatZero   = big.NewRat(0, 1)
	bigRatOne    = big.NewRat(1, 1)
)

func (dec *Decoder) stringToBigInt(s string, t reflect.Type) *big.Int {
	if bi, ok := new(big.Int).SetString(s, 10); ok {
		return bi
	}
	typeName := "*big.Int"
	if t != nil {
		typeName = t.String()
	}
	dec.decodeStringError(s, typeName)
	return nil
}

func (dec *Decoder) stringToBigFloat(s string, t reflect.Type) *big.Float {
	if bf, ok := new(big.Float).SetString(s); ok {
		return bf
	}
	typeName := "*big.Float"
	if t != nil {
		typeName = t.String()
	}
	dec.decodeStringError(s, typeName)
	return nil
}

func (dec *Decoder) stringToBigRat(s string, t reflect.Type) *big.Rat {
	if bf, ok := new(big.Rat).SetString(s); ok {
		return bf
	}
	dec.decodeStringError(s, t.String())
	return nil
}

func (dec *Decoder) readBigInt(t reflect.Type) *big.Int {
	return dec.stringToBigInt(unsafeString(dec.UnsafeUntil(TagSemicolon)), t)
}

func (dec *Decoder) readBigFloat(t reflect.Type) *big.Float {
	return dec.stringToBigFloat(unsafeString(dec.UnsafeUntil(TagSemicolon)), t)
}

// ReadBigInt reads *big.Int.
func (dec *Decoder) ReadBigInt() *big.Int {
	return dec.readBigInt(nil)
}

// ReadBigFloat reads *big.Float.
func (dec *Decoder) ReadBigFloat() *big.Float {
	return dec.readBigFloat(nil)
}

func (dec *Decoder) decodeBigInt(t reflect.Type, tag byte) (result *big.Int) {
	if i := intDigits[tag]; i != invalidDigit {
		return big.NewInt(int64(i))
	}
	switch tag {
	case TagNull:
		return nil
	case TagEmpty, TagFalse:
		return bigIntZero
	case TagTrue:
		return bigIntOne
	case TagInteger:
		return big.NewInt(dec.ReadInt64())
	case TagLong:
		return dec.readBigInt(t)
	case TagDouble:
		if bf := dec.readBigFloat(t); bf != nil {
			bi, _ := bf.Int(nil)
			return bi
		}
	case TagUTF8Char:
		return dec.stringToBigInt(dec.readUnsafeString(1), t)
	case TagString:
		if dec.IsSimple() {
			return dec.stringToBigInt(dec.ReadUnsafeString(), t)
		}
		return dec.stringToBigInt(dec.ReadString(), t)
	default:
		dec.defaultDecode(t, &result, tag)
	}
	return
}

func (dec *Decoder) decodeBigIntValue(t reflect.Type, tag byte) big.Int {
	if i := dec.decodeBigInt(t, tag); i != nil {
		return *i
	}
	return *bigIntZero
}

func (dec *Decoder) decodeBigFloat(t reflect.Type, tag byte) (result *big.Float) {
	if i := intDigits[tag]; i != invalidDigit {
		return big.NewFloat(float64(i))
	}
	switch tag {
	case TagNull:
		return nil
	case TagEmpty, TagFalse:
		return bigFloatZero
	case TagTrue:
		return bigFloatOne
	case TagInteger:
		return big.NewFloat(float64(dec.ReadInt64()))
	case TagLong, TagDouble:
		return dec.readBigFloat(t)
	case TagInfinity:
		if dec.NextByte() == TagNeg {
			return big.NewFloat(math.Inf(-1))
		}
		return big.NewFloat(math.Inf(1))
	case TagUTF8Char:
		return dec.stringToBigFloat(dec.readUnsafeString(1), t)
	case TagString:
		if dec.IsSimple() {
			return dec.stringToBigFloat(dec.ReadUnsafeString(), t)
		}
		return dec.stringToBigFloat(dec.ReadString(), t)
	default:
		dec.defaultDecode(t, &result, tag)
	}
	return
}

func (dec *Decoder) decodeBigFloatValue(t reflect.Type, tag byte) big.Float {
	if f := dec.decodeBigFloat(t, tag); f != nil {
		return *f
	}
	return *bigFloatZero
}

func (dec *Decoder) decodeBigRat(t reflect.Type, tag byte) (result *big.Rat) {
	if i := intDigits[tag]; i != invalidDigit {
		return big.NewRat(int64(i), 1)
	}
	switch tag {
	case TagNull:
		return nil
	case TagEmpty, TagFalse:
		return bigRatZero
	case TagTrue:
		return bigRatOne
	case TagInteger:
		return big.NewRat(dec.ReadInt64(), 1)
	case TagLong:
		return new(big.Rat).SetInt(dec.readBigInt(t))
	case TagDouble:
		return new(big.Rat).SetFloat64(dec.ReadFloat64())
	case TagUTF8Char:
		return dec.stringToBigRat(dec.readUnsafeString(1), t)
	case TagString:
		if dec.IsSimple() {
			return dec.stringToBigRat(dec.ReadUnsafeString(), t)
		}
		return dec.stringToBigRat(dec.ReadString(), t)
	default:
		dec.defaultDecode(t, &result, tag)
	}
	return
}

func (dec *Decoder) decodeBigRatValue(t reflect.Type, tag byte) big.Rat {
	if r := dec.decodeBigRat(t, tag); r != nil {
		return *r
	}
	return *bigRatZero
}

// bigIntValueDecoder is the implementation of ValueDecoder for big.Int.
type bigIntValueDecoder struct{}

func (bigIntValueDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*big.Int)(reflect2.PtrOf(p)) = dec.decodeBigIntValue(reflect.TypeOf(p).Elem(), tag)
}

// bigIntDecoder is the implementation of ValueDecoder for *big.Int.
type bigIntDecoder struct{}

func (bigIntDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**big.Int)(reflect2.PtrOf(p)) = dec.decodeBigInt(reflect.TypeOf(p).Elem(), tag)
}

// bigFloatValueDecoder is the implementation of ValueDecoder for big.Float.
type bigFloatValueDecoder struct{}

func (bigFloatValueDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*big.Float)(reflect2.PtrOf(p)) = dec.decodeBigFloatValue(reflect.TypeOf(p).Elem(), tag)
}

// bigFloatDecoder is the implementation of ValueDecoder for *big.Float.
type bigFloatDecoder struct{}

func (bigFloatDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**big.Float)(reflect2.PtrOf(p)) = dec.decodeBigFloat(reflect.TypeOf(p).Elem(), tag)
}

// bigRatValueDecoder is the implementation of ValueDecoder for big.Rat.
type bigRatValueDecoder struct{}

func (bigRatValueDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*big.Rat)(reflect2.PtrOf(p)) = dec.decodeBigRatValue(reflect.TypeOf(p).Elem(), tag)
}

// bigRatDecoder is the implementation of ValueDecoder for big.Rat/*big.Rat.
type bigRatDecoder struct{}

func (bigRatDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**big.Rat)(reflect2.PtrOf(p)) = dec.decodeBigRat(reflect.TypeOf(p).Elem(), tag)
}

func init() {
	registerValueDecoder(bigIntType, bigIntDecoder{})
	registerValueDecoder(bigFloatType, bigFloatDecoder{})
	registerValueDecoder(bigRatType, bigRatDecoder{})
	registerValueDecoder(bigIntValueType, bigIntValueDecoder{})
	registerValueDecoder(bigFloatValueType, bigFloatValueDecoder{})
	registerValueDecoder(bigRatValueType, bigRatValueDecoder{})
}
