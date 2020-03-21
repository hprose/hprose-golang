/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/big_encoder.go                               |
|                                                          |
| LastModified: Mar 20, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"math/big"

	"github.com/modern-go/reflect2"
)

// BigIntEncoder is the implementation of ValueEncoder for big.Int/*bit.Int.
type BigIntEncoder struct{}

// Encode writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as reference
func (valenc BigIntEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	return valenc.Write(enc, v)
}

// Write writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as value
func (BigIntEncoder) Write(enc *Encoder, v interface{}) (err error) {
	return WriteBigInt(enc.Writer, (*big.Int)(reflect2.PtrOf(v)))
}

// BigFloatEncoder is the implementation of ValueEncoder for big.Float/*bit.Float.
type BigFloatEncoder struct{}

// Encode writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as reference
func (valenc BigFloatEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	return valenc.Write(enc, v)
}

// Write writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as value
func (BigFloatEncoder) Write(enc *Encoder, v interface{}) (err error) {
	return WriteBigFloat(enc.Writer, (*big.Float)(reflect2.PtrOf(v)))
}

// BigRatEncoder is the implementation of ValueEncoder for big.Rat/*bit.Rat.
type BigRatEncoder struct{}

// Encode writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as reference
func (valenc BigRatEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	return valenc.Write(enc, v)
}

// Write writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as value
func (BigRatEncoder) Write(enc *Encoder, v interface{}) (err error) {
	return WriteBigRat(enc, (*big.Rat)(reflect2.PtrOf(v)))
}

func init() {
	RegisterValueEncoder((*big.Int)(nil), BigIntEncoder{})
	RegisterValueEncoder((*big.Float)(nil), BigFloatEncoder{})
	RegisterValueEncoder((*big.Rat)(nil), BigRatEncoder{})
}
