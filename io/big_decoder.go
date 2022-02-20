/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/big_decoder.go                                        |
|                                                          |
| LastModified: Feb 20, 2022                               |
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

func (dec *Decoder) decodeBigInt(t reflect.Type, tag byte, p **big.Int) {
	if i := intDigits[tag]; i != invalidDigit {
		*p = big.NewInt(int64(i))
		return
	}
	switch tag {
	case TagNull:
		*p = nil
	case TagEmpty, TagFalse:
		*p = bigIntZero
	case TagTrue:
		*p = bigIntOne
	case TagInteger:
		*p = big.NewInt(dec.ReadInt64())
	case TagLong:
		*p = dec.readBigInt(t)
	case TagDouble:
		if bf := dec.readBigFloat(t); bf != nil {
			*p, _ = bf.Int(nil)
		}
	case TagUTF8Char:
		*p = dec.stringToBigInt(dec.readUnsafeString(1), t)
	case TagString:
		if dec.IsSimple() {
			*p = dec.stringToBigInt(dec.ReadUnsafeString(), t)
		} else {
			*p = dec.stringToBigInt(dec.ReadString(), t)
		}
	default:
		dec.defaultDecode(t, p, tag)
	}
}

func (dec *Decoder) decodeBigIntValue(t reflect.Type, tag byte, p *big.Int) {
	var pp *big.Int
	dec.decodeBigInt(t, tag, &pp)
	if pp == nil {
		*p = *bigIntZero
	} else {
		*p = *pp
	}
}

func (dec *Decoder) decodeBigFloat(t reflect.Type, tag byte, p **big.Float) {
	if i := intDigits[tag]; i != invalidDigit {
		*p = big.NewFloat(float64(i))
		return
	}
	switch tag {
	case TagNull:
		*p = nil
	case TagEmpty, TagFalse:
		*p = bigFloatZero
	case TagTrue:
		*p = bigFloatOne
	case TagInteger:
		*p = big.NewFloat(float64(dec.ReadInt64()))
	case TagLong, TagDouble:
		*p = dec.readBigFloat(t)
	case TagInfinity:
		if dec.NextByte() == TagNeg {
			*p = big.NewFloat(math.Inf(-1))
		} else {
			*p = big.NewFloat(math.Inf(1))
		}
	case TagUTF8Char:
		*p = dec.stringToBigFloat(dec.readUnsafeString(1), t)
	case TagString:
		if dec.IsSimple() {
			*p = dec.stringToBigFloat(dec.ReadUnsafeString(), t)
		} else {
			*p = dec.stringToBigFloat(dec.ReadString(), t)
		}
	default:
		dec.defaultDecode(t, p, tag)
	}
}

func (dec *Decoder) decodeBigFloatValue(t reflect.Type, tag byte, p *big.Float) {
	var pp *big.Float
	dec.decodeBigFloat(t, tag, &pp)
	if pp == nil {
		*p = *bigFloatZero
	} else {
		*p = *pp
	}
}

func (dec *Decoder) decodeBigRat(t reflect.Type, tag byte, p **big.Rat) {
	if i := intDigits[tag]; i != invalidDigit {
		*p = big.NewRat(int64(i), 1)
		return
	}
	switch tag {
	case TagNull:
		*p = nil
	case TagEmpty, TagFalse:
		*p = bigRatZero
	case TagTrue:
		*p = bigRatOne
	case TagInteger:
		*p = big.NewRat(dec.ReadInt64(), 1)
	case TagLong:
		*p = new(big.Rat).SetInt(dec.readBigInt(t))
	case TagDouble:
		*p = new(big.Rat).SetFloat64(dec.ReadFloat64())
	case TagUTF8Char:
		*p = dec.stringToBigRat(dec.readUnsafeString(1), t)
	case TagString:
		if dec.IsSimple() {
			*p = dec.stringToBigRat(dec.ReadUnsafeString(), t)
		} else {
			*p = dec.stringToBigRat(dec.ReadString(), t)
		}
	default:
		dec.defaultDecode(t, p, tag)
	}
}

func (dec *Decoder) decodeBigRatValue(t reflect.Type, tag byte, p *big.Rat) {
	var pp *big.Rat
	dec.decodeBigRat(t, tag, &pp)
	if pp == nil {
		*p = *bigRatZero
	} else {
		*p = *pp
	}
}

// bigIntValueDecoder is the implementation of ValueDecoder for big.Int.
type bigIntValueDecoder struct{}

func (bigIntValueDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeBigIntValue(reflect.TypeOf(p).Elem(), tag, (*big.Int)(reflect2.PtrOf(p)))
}

// bigIntDecoder is the implementation of ValueDecoder for *big.Int.
type bigIntDecoder struct{}

func (bigIntDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeBigInt(reflect.TypeOf(p).Elem(), tag, (**big.Int)(reflect2.PtrOf(p)))
}

// bigFloatValueDecoder is the implementation of ValueDecoder for big.Float.
type bigFloatValueDecoder struct{}

func (bigFloatValueDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeBigFloatValue(reflect.TypeOf(p).Elem(), tag, (*big.Float)(reflect2.PtrOf(p)))
}

// bigFloatDecoder is the implementation of ValueDecoder for *big.Float.
type bigFloatDecoder struct{}

func (bigFloatDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeBigFloat(reflect.TypeOf(p).Elem(), tag, (**big.Float)(reflect2.PtrOf(p)))
}

// bigRatValueDecoder is the implementation of ValueDecoder for big.Rat.
type bigRatValueDecoder struct{}

func (bigRatValueDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeBigRatValue(reflect.TypeOf(p).Elem(), tag, (*big.Rat)(reflect2.PtrOf(p)))
}

// bigRatDecoder is the implementation of ValueDecoder for big.Rat/*big.Rat.
type bigRatDecoder struct{}

func (bigRatDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeBigRat(reflect.TypeOf(p).Elem(), tag, (**big.Rat)(reflect2.PtrOf(p)))
}

func init() {
	registerValueDecoder(bigIntType, bigIntDecoder{})
	registerValueDecoder(bigFloatType, bigFloatDecoder{})
	registerValueDecoder(bigRatType, bigRatDecoder{})
	registerValueDecoder(bigIntValueType, bigIntValueDecoder{})
	registerValueDecoder(bigFloatValueType, bigFloatValueDecoder{})
	registerValueDecoder(bigRatValueType, bigRatValueDecoder{})
}
