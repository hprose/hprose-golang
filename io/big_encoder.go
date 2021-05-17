/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/big_encoder.go                                        |
|                                                          |
| LastModified: Feb 18, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"math/big"

	"github.com/modern-go/reflect2"
)

// bigIntEncoder is the implementation of ValueEncoder for big.Int/*bit.Int.
type bigIntEncoder struct{}

func (valenc bigIntEncoder) Encode(enc *Encoder, v interface{}) {
	valenc.Write(enc, v)
}

func (bigIntEncoder) Write(enc *Encoder, v interface{}) {
	enc.WriteBigInt((*big.Int)(reflect2.PtrOf(v)))
}

// bigFloatEncoder is the implementation of ValueEncoder for big.Float/*bit.Float.
type bigFloatEncoder struct{}

func (valenc bigFloatEncoder) Encode(enc *Encoder, v interface{}) {
	valenc.Write(enc, v)
}

func (bigFloatEncoder) Write(enc *Encoder, v interface{}) {
	enc.WriteBigFloat((*big.Float)(reflect2.PtrOf(v)))
}

// bigRatEncoder is the implementation of ValueEncoder for big.Rat/*bit.Rat.
type bigRatEncoder struct{}

func (valenc bigRatEncoder) Encode(enc *Encoder, v interface{}) {
	valenc.Write(enc, v)
}

func (bigRatEncoder) Write(enc *Encoder, v interface{}) {
	enc.WriteBigRat((*big.Rat)(reflect2.PtrOf(v)))
}

// WriteBigFloat to encoder.
func (enc *Encoder) WriteBigFloat(f *big.Float) {
	enc.buf = append(enc.buf, TagDouble)
	enc.buf = f.Append(enc.buf, 'g', -1)
	enc.buf = append(enc.buf, TagSemicolon)
}

// WriteBigInt to encoder.
func (enc *Encoder) WriteBigInt(i *big.Int) {
	enc.buf = append(enc.buf, TagLong)
	enc.buf = append(enc.buf, i.String()...)
	enc.buf = append(enc.buf, TagSemicolon)
}

// WriteBigRat to encoder.
func (enc *Encoder) WriteBigRat(r *big.Rat) {
	if r.IsInt() {
		enc.WriteBigInt(r.Num())
	} else {
		enc.AddReferenceCount(1)
		s := r.String()
		enc.buf = appendString(enc.buf, s, len(s))
	}
}

func init() {
	RegisterValueEncoder((*big.Int)(nil), bigIntEncoder{})
	RegisterValueEncoder((*big.Float)(nil), bigFloatEncoder{})
	RegisterValueEncoder((*big.Rat)(nil), bigRatEncoder{})
}
