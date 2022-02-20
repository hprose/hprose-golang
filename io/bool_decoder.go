/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/bool_decoder.go                                       |
|                                                          |
| LastModified: Feb 20, 2022                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"reflect"
	"strconv"

	"github.com/modern-go/reflect2"
)

func (dec *Decoder) stringToBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		dec.Error = err
	}
	return b
}

func (dec *Decoder) decodeBool(t reflect.Type, tag byte, p *bool) {
	if i := intDigits[tag]; i != invalidDigit {
		*p = i > 0
		return
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
		*p = false
	case TagTrue, TagNaN:
		*p = true
	case TagInteger, TagLong, TagDouble:
		bytes := dec.UnsafeUntil(TagSemicolon)
		switch len(bytes) {
		case 0:
			*p = false
		case 1:
			*p = bytes[0] != '0'
		default:
			*p = true
		}
	case TagInfinity:
		dec.Skip()
		*p = true
	case TagUTF8Char:
		*p = dec.stringToBool(dec.readUnsafeString(1))
	case TagString:
		if dec.IsSimple() {
			*p = dec.stringToBool(dec.ReadUnsafeString())
		} else {
			*p = dec.stringToBool(dec.ReadString())
		}
	default:
		dec.defaultDecode(t, p, tag)
	}
}

func (dec *Decoder) decodeBoolPtr(t reflect.Type, tag byte, p **bool) {
	if tag == TagNull {
		*p = nil
		return
	}
	var b bool
	dec.decodeBool(t, tag, &b)
	*p = &b
}

// boolDecoder is the implementation of ValueDecoder for bool.
type boolDecoder struct {
	t reflect.Type
}

func (valdec boolDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeBool(valdec.t, tag, (*bool)(reflect2.PtrOf(p)))
}

// boolPtrDecoder is the implementation of ValueDecoder for *bool.
type boolPtrDecoder struct {
	t reflect.Type
}

func (valdec boolPtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeBoolPtr(valdec.t, tag, (**bool)(reflect2.PtrOf(p)))
}
