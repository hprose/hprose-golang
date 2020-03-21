/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/big_encoder.go                                  |
|                                                          |
| LastModified: Mar 21, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"math/big"

	"github.com/modern-go/reflect2"
)

// bigIntEncoder is the implementation of ValueEncoder for big.Int/*bit.Int.
type bigIntEncoder struct{}

func (valenc bigIntEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	return valenc.Write(enc, v)
}

func (bigIntEncoder) Write(enc *Encoder, v interface{}) (err error) {
	return WriteBigInt(enc, (*big.Int)(reflect2.PtrOf(v)))
}

// bigFloatEncoder is the implementation of ValueEncoder for big.Float/*bit.Float.
type bigFloatEncoder struct{}

func (valenc bigFloatEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	return valenc.Write(enc, v)
}

func (bigFloatEncoder) Write(enc *Encoder, v interface{}) (err error) {
	return WriteBigFloat(enc, (*big.Float)(reflect2.PtrOf(v)))
}

// bigRatEncoder is the implementation of ValueEncoder for big.Rat/*bit.Rat.
type bigRatEncoder struct{}

func (valenc bigRatEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	return valenc.Write(enc, v)
}

func (bigRatEncoder) Write(enc *Encoder, v interface{}) (err error) {
	return WriteBigRat(enc, (*big.Rat)(reflect2.PtrOf(v)))
}

func init() {
	RegisterValueEncoder((*big.Int)(nil), bigIntEncoder{})
	RegisterValueEncoder((*big.Float)(nil), bigFloatEncoder{})
	RegisterValueEncoder((*big.Rat)(nil), bigRatEncoder{})
}