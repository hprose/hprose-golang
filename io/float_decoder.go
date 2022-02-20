/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/float_decoder.go                                      |
|                                                          |
| LastModified: Feb 20, 2022                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"math"
	"reflect"
	"strconv"

	"github.com/modern-go/reflect2"
)

func (dec *Decoder) stringToFloat32(s string) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		dec.Error = err
	}
	return float32(f)
}

func (dec *Decoder) stringToFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		dec.Error = err
	}
	return f
}

func (dec *Decoder) readInf() float64 {
	if dec.NextByte() == TagNeg {
		return math.Inf(-1)
	}
	return math.Inf(1)
}

func (dec *Decoder) decodeFloat32(t reflect.Type, tag byte, p *float32) {
	if i := intDigits[tag]; i != invalidDigit {
		*p = float32(i)
		return
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
		*p = 0
	case TagTrue:
		*p = 1
	case TagInteger:
		*p = float32(dec.ReadInt())
	case TagLong, TagDouble:
		*p = dec.ReadFloat32()
	case TagNaN:
		*p = float32(math.NaN())
	case TagInfinity:
		*p = float32(dec.readInf())
	case TagUTF8Char:
		*p = dec.stringToFloat32(dec.readUnsafeString(1))
	case TagString:
		if dec.IsSimple() {
			*p = dec.stringToFloat32(dec.ReadUnsafeString())
		} else {
			*p = dec.stringToFloat32(dec.ReadString())
		}
	default:
		dec.defaultDecode(t, p, tag)
	}
}

func (dec *Decoder) decodeFloat32Ptr(t reflect.Type, tag byte, p **float32) {
	if tag == TagNull {
		*p = nil
		return
	}
	var f float32
	dec.decodeFloat32(t, tag, &f)
	*p = &f
}

func (dec *Decoder) decodeFloat64(t reflect.Type, tag byte, p *float64) {
	if i := intDigits[tag]; i != invalidDigit {
		*p = float64(i)
		return
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
		*p = 0
	case TagTrue:
		*p = 1
	case TagInteger:
		*p = float64(dec.ReadInt())
	case TagLong, TagDouble:
		*p = dec.ReadFloat64()
	case TagNaN:
		*p = math.NaN()
	case TagInfinity:
		*p = dec.readInf()
	case TagUTF8Char:
		*p = dec.stringToFloat64(dec.readUnsafeString(1))
	case TagString:
		if dec.IsSimple() {
			*p = dec.stringToFloat64(dec.ReadUnsafeString())
		} else {
			*p = dec.stringToFloat64(dec.ReadString())
		}
	default:
		dec.defaultDecode(t, p, tag)
	}
}

func (dec *Decoder) decodeFloat64Ptr(t reflect.Type, tag byte, p **float64) {
	if tag == TagNull {
		*p = nil
		return
	}
	var f float64
	dec.decodeFloat64(t, tag, &f)
	*p = &f
}

// float32Decoder is the implementation of ValueDecoder for float32.
type float32Decoder struct {
	t reflect.Type
}

func (valdec float32Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeFloat32(valdec.t, tag, (*float32)(reflect2.PtrOf(p)))
}

// float32PtrDecoder is the implementation of ValueDecoder for *float32.
type float32PtrDecoder struct {
	t reflect.Type
}

func (valdec float32PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeFloat32Ptr(valdec.t, tag, (**float32)(reflect2.PtrOf(p)))
}

// float64Decoder is the implementation of ValueDecoder for float64.
type float64Decoder struct {
	t reflect.Type
}

func (valdec float64Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeFloat64(valdec.t, tag, (*float64)(reflect2.PtrOf(p)))
}

// float64PtrDecoder is the implementation of ValueDecoder for *float64.
type float64PtrDecoder struct {
	t reflect.Type
}

func (valdec float64PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeFloat64Ptr(valdec.t, tag, (**float64)(reflect2.PtrOf(p)))
}
