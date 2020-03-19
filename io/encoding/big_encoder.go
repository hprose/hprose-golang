/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/big_encoder.go                               |
|                                                          |
| LastModified: Mar 19, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"math/big"
	"reflect"
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
	switch v := v.(type) {
	case big.Int:
		return WriteBigInt(enc.Writer, &v)
	case *big.Int:
		return WriteBigInt(enc.Writer, v)
	}
	return &UnsupportedTypeError{Type: reflect.TypeOf(v)}
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
	switch v := v.(type) {
	case big.Float:
		return WriteBigFloat(enc.Writer, &v)
	case *big.Float:
		return WriteBigFloat(enc.Writer, v)
	}
	return &UnsupportedTypeError{Type: reflect.TypeOf(v)}
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
	switch v := v.(type) {
	case big.Rat:
		return WriteBigRat(enc, &v)
	case *big.Rat:
		return WriteBigRat(enc, v)
	}
	return &UnsupportedTypeError{Type: reflect.TypeOf(v)}
}

func init() {
	RegisterEncoder((*big.Int)(nil), BigIntEncoder{})
	RegisterEncoder((*big.Float)(nil), BigFloatEncoder{})
	RegisterEncoder((*big.Rat)(nil), BigRatEncoder{})
}
