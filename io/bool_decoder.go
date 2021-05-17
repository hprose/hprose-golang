/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/bool_decoder.go                                       |
|                                                          |
| LastModified: May 14, 2021                               |
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

func (dec *Decoder) decodeBool(t reflect.Type, tag byte) (result bool) {
	if i := intDigits[tag]; i != invalidDigit {
		return i > 0
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
		return false
	case TagTrue, TagNaN:
		return true
	case TagInteger, TagLong, TagDouble:
		bytes := dec.UnsafeUntil(TagSemicolon)
		if len(bytes) == 0 {
			return false
		}
		if len(bytes) == 1 {
			return bytes[0] != '0'
		}
		return true
	case TagInfinity:
		dec.Skip()
		return true
	case TagUTF8Char:
		return dec.stringToBool(dec.readUnsafeString(1))
	case TagString:
		if dec.IsSimple() {
			return dec.stringToBool(dec.ReadUnsafeString())
		}
		return dec.stringToBool(dec.ReadString())
	default:
		dec.defaultDecode(t, &result, tag)
	}
	return
}

func (dec *Decoder) decodeBoolPtr(t reflect.Type, tag byte) *bool {
	if tag == TagNull {
		return nil
	}
	b := dec.decodeBool(t, tag)
	return &b
}

// boolDecoder is the implementation of ValueDecoder for bool.
type boolDecoder struct {
	t reflect.Type
}

func (valdec boolDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*bool)(reflect2.PtrOf(p)) = dec.decodeBool(valdec.t, tag)
}

func (valdec boolDecoder) Type() reflect.Type {
	return valdec.t
}

// boolPtrDecoder is the implementation of ValueDecoder for *bool.
type boolPtrDecoder struct {
	t reflect.Type
}

func (valdec boolPtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**bool)(reflect2.PtrOf(p)) = dec.decodeBoolPtr(valdec.t, tag)
}

func (valdec boolPtrDecoder) Type() reflect.Type {
	return valdec.t
}
