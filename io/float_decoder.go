/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/float_decoder.go                                      |
|                                                          |
| LastModified: Jun 5, 2021                                |
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

func (dec *Decoder) decodeFloat32(t reflect.Type, tag byte) (result float32) {
	if i := intDigits[tag]; i != invalidDigit {
		return float32(i)
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
		return 0
	case TagTrue:
		return 1
	case TagInteger:
		return float32(dec.ReadInt())
	case TagLong, TagDouble:
		return dec.ReadFloat32()
	case TagNaN:
		return float32(math.NaN())
	case TagInfinity:
		return float32(dec.readInf())
	case TagUTF8Char:
		return dec.stringToFloat32(dec.readUnsafeString(1))
	case TagString:
		if dec.IsSimple() {
			return dec.stringToFloat32(dec.ReadUnsafeString())
		}
		return dec.stringToFloat32(dec.ReadString())
	default:
		dec.defaultDecode(t, &result, tag)
	}
	return
}

func (dec *Decoder) decodeFloat32Ptr(t reflect.Type, tag byte) *float32 {
	if tag == TagNull {
		return nil
	}
	f := dec.decodeFloat32(t, tag)
	return &f
}

func (dec *Decoder) decodeFloat64(t reflect.Type, tag byte) (result float64) {
	if i := intDigits[tag]; i != invalidDigit {
		return float64(i)
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
		return 0
	case TagTrue:
		return 1
	case TagInteger:
		return float64(dec.ReadInt())
	case TagLong, TagDouble:
		return dec.ReadFloat64()
	case TagNaN:
		return math.NaN()
	case TagInfinity:
		return dec.readInf()
	case TagUTF8Char:
		return dec.stringToFloat64(dec.readUnsafeString(1))
	case TagString:
		if dec.IsSimple() {
			return dec.stringToFloat64(dec.ReadUnsafeString())
		}
		return dec.stringToFloat64(dec.ReadString())
	default:
		dec.defaultDecode(t, &result, tag)
	}
	return
}

func (dec *Decoder) decodeFloat64Ptr(t reflect.Type, tag byte) *float64 {
	if tag == TagNull {
		return nil
	}
	f := dec.decodeFloat64(t, tag)
	return &f
}

// float32Decoder is the implementation of ValueDecoder for float32.
type float32Decoder struct {
	t reflect.Type
}

func (valdec float32Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*float32)(reflect2.PtrOf(p)) = dec.decodeFloat32(valdec.t, tag)
}

// float32PtrDecoder is the implementation of ValueDecoder for *float32.
type float32PtrDecoder struct {
	t reflect.Type
}

func (valdec float32PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**float32)(reflect2.PtrOf(p)) = dec.decodeFloat32Ptr(valdec.t, tag)
}

// float64Decoder is the implementation of ValueDecoder for float64.
type float64Decoder struct {
	t reflect.Type
}

func (valdec float64Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*float64)(reflect2.PtrOf(p)) = dec.decodeFloat64(valdec.t, tag)
}

// float64PtrDecoder is the implementation of ValueDecoder for *float64.
type float64PtrDecoder struct {
	t reflect.Type
}

func (valdec float64PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**float64)(reflect2.PtrOf(p)) = dec.decodeFloat64Ptr(valdec.t, tag)
}
