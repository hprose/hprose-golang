/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/complex_decoder.go                                    |
|                                                          |
| LastModified: Feb 20, 2022                               |
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

func (dec *Decoder) decodeComplex64(t reflect.Type, tag byte, p *complex64) {
	if i := intDigits[tag]; i != invalidDigit {
		*p = complex(float32(i), 0)
		return
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
		*p = 0
	case TagTrue:
		*p = 1
	case TagNaN:
		*p = complex(float32(math.NaN()), 0)
	case TagInteger:
		*p = complex(float32(dec.ReadInt32()), 0)
	case TagLong, TagDouble:
		*p = complex(dec.ReadFloat32(), 0)
	case TagInfinity:
		*p = complex(float32(dec.readInf()), 0)
	case TagUTF8Char:
		*p = dec.stringToComplex64(dec.readUnsafeString(1))
	case TagString:
		if dec.IsSimple() {
			*p = dec.stringToComplex64(dec.ReadUnsafeString())
		} else {
			*p = dec.stringToComplex64(dec.ReadString())
		}
	default:
		dec.defaultDecode(t, p, tag)
	}
}

func (dec *Decoder) decodeComplex64Ptr(t reflect.Type, tag byte, p **complex64) {
	if tag == TagNull {
		*p = nil
		return
	}
	var c complex64
	dec.decodeComplex64(t, tag, &c)
	*p = &c
}

func (dec *Decoder) decodeComplex128(t reflect.Type, tag byte, p *complex128) {
	if i := intDigits[tag]; i != invalidDigit {
		*p = complex(float64(i), 0)
		return
	}
	switch tag {
	case TagEmpty, TagFalse:
		*p = 0
	case TagTrue:
		*p = 1
	case TagNaN:
		*p = complex(math.NaN(), 0)
	case TagInteger:
		*p = complex(float64(dec.ReadInt32()), 0)
	case TagLong, TagDouble:
		*p = complex(dec.ReadFloat64(), 0)
	case TagInfinity:
		*p = complex(dec.readInf(), 0)
	case TagUTF8Char:
		*p = dec.stringToComplex128(dec.readUnsafeString(1))
	case TagString:
		if dec.IsSimple() {
			*p = dec.stringToComplex128(dec.ReadUnsafeString())
		} else {
			*p = dec.stringToComplex128(dec.ReadString())
		}
	default:
		dec.defaultDecode(t, p, tag)
	}
}

func (dec *Decoder) decodeComplex128Ptr(t reflect.Type, tag byte, p **complex128) {
	if tag == TagNull {
		*p = nil
		return
	}
	var c complex128
	dec.decodeComplex128(t, tag, &c)
	*p = &c
}

// complex64Decoder is the implementation of ValueDecoder for complex64.
type complex64Decoder struct {
	t reflect.Type
}

func (valdec complex64Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeComplex64(valdec.t, tag, (*complex64)(reflect2.PtrOf(p)))
}

// complex64PtrDecoder is the implementation of ValueDecoder for *complex64.
type complex64PtrDecoder struct {
	t reflect.Type
}

func (valdec complex64PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeComplex64Ptr(valdec.t, tag, (**complex64)(reflect2.PtrOf(p)))
}

// complex128Decoder is the implementation of ValueDecoder for complex128.
type complex128Decoder struct {
	t reflect.Type
}

func (valdec complex128Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeComplex128(valdec.t, tag, (*complex128)(reflect2.PtrOf(p)))
}

// complex128PtrDecoder is the implementation of ValueDecoder for *complex128.
type complex128PtrDecoder struct {
	t reflect.Type
}

func (valdec complex128PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeComplex128Ptr(valdec.t, tag, (**complex128)(reflect2.PtrOf(p)))
}
