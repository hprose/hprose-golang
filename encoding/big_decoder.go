/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/big_decoder.go                                  |
|                                                          |
| LastModified: Jun 7, 2020                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"math"
	"math/big"
	"reflect"
)

var (
	bigIntZero   = big.NewInt(0)
	bigIntOne    = big.NewInt(1)
	bigFloatZero = big.NewFloat(0)
	bigFloatOne  = big.NewFloat(1)
	bigRatZero   = big.NewRat(0, 1)
	bigRatOne    = big.NewRat(1, 1)
)

func (dec *Decoder) strToBigInt(s string) *big.Int {
	if bi, ok := new(big.Int).SetString(s, 10); ok {
		return bi
	}
	dec.decodeStringError(s, "*big.Int")
	return nil
}

// ReadBigInt read *big.Int
func (dec *Decoder) ReadBigInt() *big.Int {
	return dec.strToBigInt(unsafeString(dec.Until(TagSemicolon)))
}

func (dec *Decoder) strToBigFloat(s string) *big.Float {
	if bf, ok := new(big.Float).SetString(s); ok {
		return bf
	}
	dec.decodeStringError(s, "*big.Float")
	return nil
}

// ReadBigFloat read *big.Float
func (dec *Decoder) ReadBigFloat() *big.Float {
	return dec.strToBigFloat(unsafeString(dec.Until(TagSemicolon)))
}

func (dec *Decoder) strToBigRat(s string) *big.Rat {
	if bf, ok := new(big.Rat).SetString(s); ok {
		return bf
	}
	dec.decodeStringError(s, "*big.Rat")
	return nil
}

// bigIntDecoder is the implementation of ValueDecoder for big.Int/*big.Int.
type bigIntDecoder struct {
	destType reflect.Type
}

func (valdec bigIntDecoder) decode(dec *Decoder, tag byte) *big.Int {
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
		return dec.ReadBigInt()
	case TagDouble:
		if bf := dec.strToBigFloat(unsafeString(dec.Until(TagSemicolon))); bf != nil {
			bi, _ := bf.Int(nil)
			return bi
		}
	case TagUTF8Char:
		return dec.strToBigInt(dec.readUnsafeString(1))
	case TagString:
		return dec.strToBigInt(dec.ReadString())
	default:
		dec.decodeError(valdec.destType, tag)
	}
	return nil
}

func (valdec bigIntDecoder) decodeValue(dec *Decoder, pv *big.Int, tag byte) {
	if tag == TagNull {
		*pv = *bigIntZero
	} else if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = *i
	}
}

func (valdec bigIntDecoder) decodePtr(dec *Decoder, pv **big.Int, tag byte) {
	if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec bigIntDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch pv := p.(type) {
	case *big.Int:
		valdec.decodeValue(dec, pv, tag)
	case **big.Int:
		valdec.decodePtr(dec, pv, tag)
	}
}

// bigFloatDecoder is the implementation of ValueDecoder for big.Float/*big.Float.
type bigFloatDecoder struct {
	destType reflect.Type
}

func (valdec bigFloatDecoder) decode(dec *Decoder, tag byte) *big.Float {
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
		return dec.ReadBigFloat()
	case TagInfinity:
		if dec.NextByte() == TagNeg {
			return big.NewFloat(math.Inf(-1))
		}
		return big.NewFloat(math.Inf(1))
	case TagUTF8Char:
		return dec.strToBigFloat(dec.readUnsafeString(1))
	case TagString:
		return dec.strToBigFloat(dec.ReadString())
	default:
		dec.decodeError(valdec.destType, tag)
	}
	return nil
}

func (valdec bigFloatDecoder) decodeValue(dec *Decoder, pv *big.Float, tag byte) {
	if tag == TagNull {
		*pv = *bigFloatZero
	} else if f := valdec.decode(dec, tag); dec.Error == nil {
		*pv = *f
	}
}

func (valdec bigFloatDecoder) decodePtr(dec *Decoder, pv **big.Float, tag byte) {
	if f := valdec.decode(dec, tag); dec.Error == nil {
		*pv = f
	}
}

func (valdec bigFloatDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch pv := p.(type) {
	case *big.Float:
		valdec.decodeValue(dec, pv, tag)
	case **big.Float:
		valdec.decodePtr(dec, pv, tag)
	}
}

// bigRatDecoder is the implementation of ValueDecoder for big.Rat/*big.Rat.
type bigRatDecoder struct {
	destType reflect.Type
}

func (valdec bigRatDecoder) decode(dec *Decoder, tag byte) *big.Rat {
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
		return new(big.Rat).SetInt(dec.ReadBigInt())
	case TagDouble:
		return new(big.Rat).SetFloat64(dec.ReadFloat64())
	case TagUTF8Char:
		return dec.strToBigRat(dec.readUnsafeString(1))
	case TagString:
		return dec.strToBigRat(dec.ReadString())
	default:
		dec.decodeError(valdec.destType, tag)
	}
	return nil
}

func (valdec bigRatDecoder) decodeValue(dec *Decoder, pv *big.Rat, tag byte) {
	if tag == TagNull {
		*pv = *bigRatZero
	} else if r := valdec.decode(dec, tag); dec.Error == nil {
		*pv = *r
	}
}

func (valdec bigRatDecoder) decodePtr(dec *Decoder, pv **big.Rat, tag byte) {
	if r := valdec.decode(dec, tag); dec.Error == nil {
		*pv = r
	}
}

func (valdec bigRatDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch pv := p.(type) {
	case *big.Rat:
		valdec.decodeValue(dec, pv, tag)
	case **big.Rat:
		valdec.decodePtr(dec, pv, tag)
	}
}

var bigintdec = bigIntDecoder{reflect.TypeOf((*big.Int)(nil))}
var bigfloatdec = bigFloatDecoder{reflect.TypeOf((*big.Float)(nil))}
var bigratdec = bigRatDecoder{reflect.TypeOf((*big.Rat)(nil))}

func init() {
	RegisterValueDecoder((*big.Int)(nil), bigintdec)
	RegisterValueDecoder((*big.Float)(nil), bigfloatdec)
	RegisterValueDecoder((*big.Rat)(nil), bigratdec)
}
