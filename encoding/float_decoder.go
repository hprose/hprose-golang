/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/float_decoder.go                                |
|                                                          |
| LastModified: Jun 12, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

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

func (dec *Decoder) decodeFloat32(t reflect.Type, tag byte) float32 {
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
		dec.decodeError(t, tag)
	}
	return 0
}

func (dec *Decoder) decodeFloat64(t reflect.Type, tag byte) float64 {
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
		dec.decodeError(t, tag)
	}
	return 0
}

// float32Decoder is the implementation of ValueDecoder for float32.
type float32Decoder struct {
	t reflect.Type
}

func (valdec float32Decoder) decode(dec *Decoder, pv *float32, tag byte) {
	if f := dec.decodeFloat32(valdec.t, tag); dec.Error == nil {
		*pv = f
	}
}

func (valdec float32Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (*float32)(reflect2.PtrOf(p)), tag)
}

func (valdec float32Decoder) Type() reflect.Type {
	return valdec.t
}

// float32PtrDecoder is the implementation of ValueDecoder for *float32.
type float32PtrDecoder struct {
	t reflect.Type
}

func (valdec float32PtrDecoder) decode(dec *Decoder, pv **float32, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if f := dec.decodeFloat32(valdec.t, tag); dec.Error == nil {
		*pv = &f
	}
}

func (valdec float32PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (**float32)(reflect2.PtrOf(p)), tag)
}

func (valdec float32PtrDecoder) Type() reflect.Type {
	return valdec.t
}

// float64Decoder is the implementation of ValueDecoder for float64.
type float64Decoder struct {
	t reflect.Type
}

func (valdec float64Decoder) decode(dec *Decoder, pv *float64, tag byte) {
	if f := dec.decodeFloat64(valdec.t, tag); dec.Error == nil {
		*pv = f
	}
}

func (valdec float64Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (*float64)(reflect2.PtrOf(p)), tag)
}

func (valdec float64Decoder) Type() reflect.Type {
	return valdec.t
}

// float64PtrDecoder is the implementation of ValueDecoder for *float64.
type float64PtrDecoder struct {
	t reflect.Type
}

func (valdec float64PtrDecoder) decode(dec *Decoder, pv **float64, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if f := dec.decodeFloat64(valdec.t, tag); dec.Error == nil {
		*pv = &f
	}
}

func (valdec float64PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (**float64)(reflect2.PtrOf(p)), tag)
}

func (valdec float64PtrDecoder) Type() reflect.Type {
	return valdec.t
}

var (
	f32dec  = float32Decoder{reflect.TypeOf((float32)(0))}
	f64dec  = float64Decoder{reflect.TypeOf((float64)(0))}
	pf32dec = float32PtrDecoder{reflect.TypeOf((*float32)(nil))}
	pf64dec = float64PtrDecoder{reflect.TypeOf((*float64)(nil))}
)

func init() {
	RegisterValueDecoder(f32dec)
	RegisterValueDecoder(f64dec)
	RegisterValueDecoder(pf32dec)
	RegisterValueDecoder(pf64dec)
}
