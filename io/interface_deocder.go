/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/interface_decoder.go                                  |
|                                                          |
| LastModified: Feb 7, 2024                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

func (dec *Decoder) decodeLongAsInterface(p *interface{}) {
	switch dec.LongType {
	case LongTypeInt:
		*p = dec.ReadInt()
	case LongTypeUint:
		*p = dec.ReadUint()
	case LongTypeInt64:
		*p = dec.ReadInt64()
	case LongTypeUint64:
		*p = dec.ReadUint64()
	default:
		*p = dec.ReadBigInt()
	}
}

func (dec *Decoder) decodeNaNAsInterface(p *interface{}) {
	switch dec.RealType {
	case RealTypeFloat32:
		*p = float32(math.NaN())
	case RealTypeFloat64:
		*p = math.NaN()
	default:
		dec.Error = DecodeError("hprose/io: can not parse NaN to *big.Float")
	}
}

func (dec *Decoder) decodeInfinityAsInterface(p *interface{}) {
	var f float64
	if dec.NextByte() == TagNeg {
		f = math.Inf(-1)
	} else {
		f = math.Inf(1)
	}
	switch dec.RealType {
	case RealTypeFloat32:
		*p = float32(f)
	case RealTypeFloat64:
		*p = f
	default:
		*p = big.NewFloat(f)
	}
}

func (dec *Decoder) decodeDoubleAsInterface(p *interface{}) {
	switch dec.RealType {
	case RealTypeFloat32:
		*p = dec.ReadFloat32()
	case RealTypeFloat64:
		*p = dec.ReadFloat64()
	default:
		*p = dec.ReadBigFloat()
	}
}

func (dec *Decoder) decodeListAsInterface(tag byte, p *interface{}) {
	var result []interface{}
	ifsdec.Decode(dec, &result, tag)
	n := len(result)
	if n == 0 || dec.ListType == ListTypeInterfaceSlice {
		*p = result
		return
	}
	var t reflect.Type
	for i := 0; i < n; i++ {
		rt := reflect.TypeOf(result[i])
		if isNil(result[i]) || rt.Kind() == reflect.Invalid || rt.Kind() == reflect.Interface {
			continue
		}
		if t == nil {
			t = rt
		}
		if rt != t {
			*p = result
			return
		}
	}
	if t == nil {
		*p = result
		return
	}
	st := reflect2.Type2(reflect.SliceOf(t)).(*reflect2.UnsafeSliceType)
	s := st.UnsafeMakeSlice(n, n)
	for i := 0; i < n; i++ {
		if isNil(result[i]) {
			continue
		}
		p := reflect2.PtrOf(result[i])
		if t.Kind() == reflect.Ptr || t.Kind() == reflect.Map {
			st.UnsafeSetIndex(s, i, (unsafe.Pointer)(&p))
		} else {
			st.UnsafeSetIndex(s, i, p)
		}
	}
	*p = st.UnsafeIndirect(s)
}

func (dec *Decoder) decodeMapAsInterface(tag byte, p *interface{}) {
	if dec.MapType == MapTypeIIMap {
		var result map[interface{}]interface{}
		ififmdec.Decode(dec, &result, tag)
		*p = result
	} else {
		var result map[string]interface{}
		sifmdec.Decode(dec, &result, tag)
		*p = result
	}
}

func (dec *Decoder) decodeInterface(tag byte, p *interface{}) {
	if i := intDigits[tag]; i != invalidDigit {
		*p = int(i)
		return
	}
	switch tag {
	case TagNull:
		*p = nil
	case TagEmpty:
		*p = ""
	case TagFalse:
		*p = false
	case TagTrue:
		*p = true
	case TagInteger:
		*p = dec.ReadInt()
	case TagLong:
		dec.decodeLongAsInterface(p)
	case TagNaN:
		dec.decodeNaNAsInterface(p)
	case TagInfinity:
		dec.decodeInfinityAsInterface(p)
	case TagDouble:
		dec.decodeDoubleAsInterface(p)
	case TagTime:
		*p = dec.ReadTime()
	case TagDate:
		*p = dec.ReadDateTime()
	case TagGUID:
		*p = dec.ReadUUID()
	case TagUTF8Char:
		*p = dec.readSafeString(1)
	case TagString:
		*p = dec.ReadString()
	case TagBytes:
		*p = dec.ReadBytes()
	case TagList:
		dec.decodeListAsInterface(tag, p)
	case TagMap:
		dec.decodeMapAsInterface(tag, p)
	case TagObject:
		*p = dec.ReadObject()
	case TagRef:
		dec.ReadReference(p)
		return
	case TagClass:
		dec.ReadStruct(interfaceType)
		dec.Decode(p)
	case TagError:
		var s string
		dec.decodeString(stringType, dec.NextByte(), &s)
		dec.Error = DecodeError(s)
	default:
		if dec.Error == nil {
			dec.Error = DecodeError(fmt.Sprintf("hprose/io: invalid tag '%s'(0x%x)", string(tag), tag))
		}
		*p = nil
	}
}

func (dec *Decoder) decodeInterfacePtr(tag byte, p **interface{}) {
	if tag == TagNull {
		*p = nil
		return
	}
	var i interface{}
	dec.decodeInterface(tag, &i)
	*p = &i
}

// interfaceDecoder is the implementation of ValueDecoder for interface{}.
type interfaceDecoder struct{}

func (valdec interfaceDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeInterface(tag, (*interface{})(reflect2.PtrOf(p)))
}

// interfacePtrDecoder is the implementation of ValueDecoder for *interface{}.
type interfacePtrDecoder struct{}

func (valdec interfacePtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeInterfacePtr(tag, (**interface{})(reflect2.PtrOf(p)))
}
