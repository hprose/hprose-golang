/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/big_decoder.go                                  |
|                                                          |
| LastModified: Jun 26, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

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

func (dec *Decoder) strToBigInt(s string, t reflect.Type) *big.Int {
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

func (dec *Decoder) strToBigFloat(s string, t reflect.Type) *big.Float {
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

func (dec *Decoder) strToBigRat(s string, t reflect.Type) *big.Rat {
	if bf, ok := new(big.Rat).SetString(s); ok {
		return bf
	}
	dec.decodeStringError(s, t.String())
	return nil
}

func (dec *Decoder) readBigInt(t reflect.Type) *big.Int {
	return dec.strToBigInt(unsafeString(dec.UnsafeUntil(TagSemicolon)), t)
}

func (dec *Decoder) readBigFloat(t reflect.Type) *big.Float {
	return dec.strToBigFloat(unsafeString(dec.UnsafeUntil(TagSemicolon)), t)
}

// ReadBigInt reads *big.Int
func (dec *Decoder) ReadBigInt() *big.Int {
	return dec.readBigInt(nil)
}

// ReadBigFloat reads *big.Float
func (dec *Decoder) ReadBigFloat() *big.Float {
	return dec.readBigFloat(nil)
}

func (dec *Decoder) decodeBigInt(t reflect.Type, tag byte) *big.Int {
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
		return dec.strToBigInt(dec.readUnsafeString(1), t)
	case TagString:
		if dec.IsSimple() {
			return dec.strToBigInt(dec.ReadUnsafeString(), t)
		}
		return dec.strToBigInt(dec.ReadString(), t)
	default:
		dec.decodeError(t, tag)
	}
	return nil
}

func (dec *Decoder) decodeBigIntValue(t reflect.Type, tag byte) big.Int {
	if i := dec.decodeBigInt(t, tag); i != nil {
		return *i
	}
	return *bigIntZero
}

func (dec *Decoder) decodeBigFloat(t reflect.Type, tag byte) *big.Float {
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
		return dec.strToBigFloat(dec.readUnsafeString(1), t)
	case TagString:
		if dec.IsSimple() {
			return dec.strToBigFloat(dec.ReadUnsafeString(), t)
		}
		return dec.strToBigFloat(dec.ReadString(), t)
	default:
		dec.decodeError(t, tag)
	}
	return nil
}

func (dec *Decoder) decodeBigFloatValue(t reflect.Type, tag byte) big.Float {
	if f := dec.decodeBigFloat(t, tag); f != nil {
		return *f
	}
	return *bigFloatZero
}

func (dec *Decoder) decodeBigRat(t reflect.Type, tag byte) *big.Rat {
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
		return dec.strToBigRat(dec.readUnsafeString(1), t)
	case TagString:
		if dec.IsSimple() {
			return dec.strToBigRat(dec.ReadUnsafeString(), t)
		}
		return dec.strToBigRat(dec.ReadString(), t)
	default:
		dec.decodeError(t, tag)
	}
	return nil
}

func (dec *Decoder) decodeBigRatValue(t reflect.Type, tag byte) big.Rat {
	if r := dec.decodeBigRat(t, tag); r != nil {
		return *r
	}
	return *bigRatZero
}

// bigIntValueDecoder is the implementation of ValueDecoder for big.Int.
type bigIntValueDecoder struct {
	t reflect.Type
}

func (valdec bigIntValueDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*big.Int)(reflect2.PtrOf(p)) = dec.decodeBigIntValue(valdec.t, tag)
}

func (valdec bigIntValueDecoder) Type() reflect.Type {
	return valdec.t
}

// bigIntDecoder is the implementation of ValueDecoder for *big.Int.
type bigIntDecoder struct {
	t reflect.Type
}

func (valdec bigIntDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**big.Int)(reflect2.PtrOf(p)) = dec.decodeBigInt(valdec.t, tag)
}

func (valdec bigIntDecoder) Type() reflect.Type {
	return valdec.t
}

// bigFloatValueDecoder is the implementation of ValueDecoder for big.Float.
type bigFloatValueDecoder struct {
	t reflect.Type
}

func (valdec bigFloatValueDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*big.Float)(reflect2.PtrOf(p)) = dec.decodeBigFloatValue(valdec.t, tag)
}

func (valdec bigFloatValueDecoder) Type() reflect.Type {
	return valdec.t
}

// bigFloatDecoder is the implementation of ValueDecoder for *big.Float.
type bigFloatDecoder struct {
	t reflect.Type
}

func (valdec bigFloatDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**big.Float)(reflect2.PtrOf(p)) = dec.decodeBigFloat(valdec.t, tag)
}

func (valdec bigFloatDecoder) Type() reflect.Type {
	return valdec.t
}

// bigRatValueDecoder is the implementation of ValueDecoder for big.Rat.
type bigRatValueDecoder struct {
	t reflect.Type
}

func (valdec bigRatValueDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*big.Rat)(reflect2.PtrOf(p)) = dec.decodeBigRatValue(valdec.t, tag)
}

func (valdec bigRatValueDecoder) Type() reflect.Type {
	return valdec.t
}

// bigRatDecoder is the implementation of ValueDecoder for big.Rat/*big.Rat.
type bigRatDecoder struct {
	t reflect.Type
}

func (valdec bigRatDecoder) decode(dec *Decoder, pv **big.Rat, tag byte) {
	if r := dec.decodeBigRat(valdec.t, tag); dec.Error == nil {
		*pv = r
	}
}

func (valdec bigRatDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**big.Rat)(reflect2.PtrOf(p)) = dec.decodeBigRat(valdec.t, tag)
}

func (valdec bigRatDecoder) Type() reflect.Type {
	return valdec.t
}

func init() {
	RegisterValueDecoder(bigIntDecoder{bigIntType})
	RegisterValueDecoder(bigFloatDecoder{bigFloatType})
	RegisterValueDecoder(bigRatDecoder{bigRatType})
	RegisterValueDecoder(bigIntValueDecoder{bigIntValueType})
	RegisterValueDecoder(bigFloatValueDecoder{bigFloatValueType})
	RegisterValueDecoder(bigRatValueDecoder{bigRatValueType})
}
