/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/complex_decoder.go                              |
|                                                          |
| LastModified: Jun 2, 2020                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"math"
	"reflect"

	"github.com/andot/complexconv"
)

// complex64Decoder is the implementation of ValueDecoder for complex64.
type complex64Decoder struct {
	descType reflect.Type
}

var c64dec = complex64Decoder{reflect.TypeOf((*complex64)(nil)).Elem()}

func (valdec complex64Decoder) decode(dec *Decoder, tag byte) complex64 {
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
		return dec.stringToComplex64(dec.ReadString())
	default:
		dec.decodeError(valdec.descType, tag)
	}
	return 0
}

func (valdec complex64Decoder) decodeValue(dec *Decoder, pv *complex64, tag byte) {
	if c := valdec.decode(dec, tag); dec.Error == nil {
		*pv = c
	}
}

func (valdec complex64Decoder) decodePtr(dec *Decoder, pv **complex64, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if c := valdec.decode(dec, tag); dec.Error == nil {
		*pv = &c
	}
}

func (valdec complex64Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch pv := p.(type) {
	case *complex64:
		valdec.decodeValue(dec, pv, tag)
	case **complex64:
		valdec.decodePtr(dec, pv, tag)
	}
}

// complex128Decoder is the implementation of ValueDecoder for complex128.
type complex128Decoder struct {
	descType reflect.Type
}

var c128dec = complex128Decoder{reflect.TypeOf((*complex128)(nil)).Elem()}

func (valdec complex128Decoder) decode(dec *Decoder, tag byte) complex128 {
	if i := intDigits[tag]; i != invalidDigit {
		return complex(float64(i), 0)
	}
	switch tag {
	case TagEmpty, TagFalse:
		return 0
	case TagTrue:
		return 1
	case TagNaN:
		return complex(float64(math.NaN()), 0)
	case TagInteger:
		return complex(float64(dec.ReadInt32()), 0)
	case TagLong, TagDouble:
		return complex(dec.ReadFloat64(), 0)
	case TagInfinity:
		return complex(float64(dec.readInf()), 0)
	case TagUTF8Char:
		return dec.stringToComplex128(dec.readUnsafeString(1))
	case TagString:
		return dec.stringToComplex128(dec.ReadString())
	default:
		dec.decodeError(valdec.descType, tag)
	}
	return 0
}

func (valdec complex128Decoder) decodeValue(dec *Decoder, pv *complex128, tag byte) {
	if c := valdec.decode(dec, tag); dec.Error == nil {
		*pv = c
	}
}

func (valdec complex128Decoder) decodePtr(dec *Decoder, pv **complex128, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if c := valdec.decode(dec, tag); dec.Error == nil {
		*pv = &c
	}
}

func (valdec complex128Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch pv := p.(type) {
	case *complex128:
		valdec.decodeValue(dec, pv, tag)
	case **complex128:
		valdec.decodePtr(dec, pv, tag)
	}
}

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
