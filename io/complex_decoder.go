/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/complex_decoder.go                                    |
|                                                          |
| LastModified: May 14, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"math"
	"reflect"

	"github.com/andot/complexconv"
	"github.com/modern-go/reflect2"
)

func (dec *Decoder) stringToComplex64(s string) complex64 {
	c, err := complexconv.ParseComplex(s, 64)
	if err != nil {
		dec.Error = err
	}
	return complex64(c)
}

func (dec *Decoder) stringToComplex128(s string) complex128 {
	c, err := complexconv.ParseComplex(s, 128)
	if err != nil {
		dec.Error = err
	}
	return c
}

func (dec *Decoder) decodeComplex64(t reflect.Type, tag byte) (result complex64) {
	if i := intDigits[tag]; i != invalidDigit {
		return complex(float32(i), 0)
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
		return 0
	case TagTrue:
		return 1
	case TagNaN:
		return complex(float32(math.NaN()), 0)
	case TagInteger:
		return complex(float32(dec.ReadInt32()), 0)
	case TagLong, TagDouble:
		return complex(dec.ReadFloat32(), 0)
	case TagInfinity:
		return complex(float32(dec.readInf()), 0)
	case TagUTF8Char:
		return dec.stringToComplex64(dec.readUnsafeString(1))
	case TagString:
		if dec.IsSimple() {
			return dec.stringToComplex64(dec.ReadUnsafeString())
		}
		return dec.stringToComplex64(dec.ReadString())
	default:
		dec.defaultDecode(t, &result, tag)
	}
	return
}

func (dec *Decoder) decodeComplex64Ptr(t reflect.Type, tag byte) *complex64 {
	if tag == TagNull {
		return nil
	}
	c := dec.decodeComplex64(t, tag)
	return &c
}

func (dec *Decoder) decodeComplex128(t reflect.Type, tag byte) (result complex128) {
	if i := intDigits[tag]; i != invalidDigit {
		return complex(float64(i), 0)
	}
	switch tag {
	case TagEmpty, TagFalse:
		return 0
	case TagTrue:
		return 1
	case TagNaN:
		return complex(math.NaN(), 0)
	case TagInteger:
		return complex(float64(dec.ReadInt32()), 0)
	case TagLong, TagDouble:
		return complex(dec.ReadFloat64(), 0)
	case TagInfinity:
		return complex(dec.readInf(), 0)
	case TagUTF8Char:
		return dec.stringToComplex128(dec.readUnsafeString(1))
	case TagString:
		if dec.IsSimple() {
			return dec.stringToComplex128(dec.ReadUnsafeString())
		}
		return dec.stringToComplex128(dec.ReadString())
	default:
		dec.defaultDecode(t, &result, tag)
	}
	return
}

func (dec *Decoder) decodeComplex128Ptr(t reflect.Type, tag byte) *complex128 {
	if tag == TagNull {
		return nil
	}
	c := dec.decodeComplex128(t, tag)
	return &c
}

// complex64Decoder is the implementation of ValueDecoder for complex64.
type complex64Decoder struct {
	t reflect.Type
}

func (valdec complex64Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*complex64)(reflect2.PtrOf(p)) = dec.decodeComplex64(valdec.t, tag)
}

func (valdec complex64Decoder) Type() reflect.Type {
	return valdec.t
}

// complex64PtrDecoder is the implementation of ValueDecoder for *complex64.
type complex64PtrDecoder struct {
	t reflect.Type
}

func (valdec complex64PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**complex64)(reflect2.PtrOf(p)) = dec.decodeComplex64Ptr(valdec.t, tag)
}

func (valdec complex64PtrDecoder) Type() reflect.Type {
	return valdec.t
}

// complex128Decoder is the implementation of ValueDecoder for complex128.
type complex128Decoder struct {
	t reflect.Type
}

func (valdec complex128Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*complex128)(reflect2.PtrOf(p)) = dec.decodeComplex128(valdec.t, tag)
}

func (valdec complex128Decoder) Type() reflect.Type {
	return valdec.t
}

// complex128PtrDecoder is the implementation of ValueDecoder for *complex128.
type complex128PtrDecoder struct {
	t reflect.Type
}

func (valdec complex128PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**complex128)(reflect2.PtrOf(p)) = dec.decodeComplex128Ptr(valdec.t, tag)
}

func (valdec complex128PtrDecoder) Type() reflect.Type {
	return valdec.t
}
