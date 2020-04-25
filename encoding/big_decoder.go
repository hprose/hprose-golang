/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/big_decoder.go                                  |
|                                                          |
| LastModified: Apr 25, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"math"
	"math/big"
)

var (
	bigIntZero   = big.NewInt(0)
	bigIntOne    = big.NewInt(1)
	bigFloatZero = big.NewFloat(0)
	bigFloatOne  = big.NewFloat(1)
)

func (dec *Decoder) strToBigInt(s string) *big.Int {
	if bi, ok := new(big.Int).SetString(s, 10); ok {
		return bi
	}
	dec.decodeStringError(s, "big.Int")
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
	dec.decodeStringError(s, "big.Float")
	return nil
}

// ReadBigFloat read *big.Float
func (dec *Decoder) ReadBigFloat() *big.Float {
	return dec.strToBigFloat(unsafeString(dec.Until(TagSemicolon)))
}

// bigIntDecoder is the implementation of ValueDecoder for big.Int/*big.Int.
type bigIntDecoder struct{}

func (valdec bigIntDecoder) decode(dec *Decoder, tag byte) *big.Int {
	if i := intDigits[tag]; i != invalidDigit {
		return big.NewInt(int64(i))
	}
	switch tag {
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
	}
	return nil
}

func (valdec bigIntDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if tag == TagNull {
		switch pv := p.(type) {
		case **big.Int:
			*pv = nil
		case *big.Int:
			*pv = *bigIntZero
		}
		return
	}
	bi := valdec.decode(dec, tag)
	if bi == nil {
		dec.decodeError(p, tag)
		return
	}
	switch pv := p.(type) {
	case **big.Int:
		*pv = bi
	case *big.Int:
		*pv = *bi
	}
}

// bigFloatDecoder is the implementation of ValueDecoder for big.Float/*big.Float.
type bigFloatDecoder struct{}

func (valdec bigFloatDecoder) decode(dec *Decoder, tag byte) *big.Float {
	if i := intDigits[tag]; i != invalidDigit {
		return big.NewFloat(float64(i))
	}
	switch tag {
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
	}
	return nil
}

func (valdec bigFloatDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if tag == TagNull {
		switch pv := p.(type) {
		case **big.Float:
			*pv = nil
		case *big.Float:
			*pv = *bigFloatZero
		}
		return
	}
	bf := valdec.decode(dec, tag)
	if bf == nil {
		dec.decodeError(p, tag)
		return
	}
	switch pv := p.(type) {
	case **big.Float:
		*pv = bf
	case *big.Float:
		*pv = *bf
	}
}

func init() {
	RegisterValueDecoder((*big.Int)(nil), bigIntDecoder{})
	RegisterValueDecoder((*big.Float)(nil), bigFloatDecoder{})
	// RegisterValueDecoder((*big.Rat)(nil), bigRatDecoder{})
}
