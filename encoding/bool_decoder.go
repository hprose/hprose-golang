/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/bool_decoder.go                                 |
|                                                          |
| LastModified: Jun 13, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

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

func (dec *Decoder) decodeBool(t reflect.Type, tag byte) bool {
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
		dec.decodeError(t, tag)
	}
	return false
}

// boolDecoder is the implementation of ValueDecoder for bool.
type boolDecoder struct {
	t reflect.Type
}

func (valdec boolDecoder) decode(dec *Decoder, pv *bool, tag byte) {
	if b := dec.decodeBool(valdec.t, tag); dec.Error == nil {
		*pv = b
	}
}

func (valdec boolDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (*bool)(reflect2.PtrOf(p)), tag)
}

func (valdec boolDecoder) Type() reflect.Type {
	return valdec.t
}

// boolPtrDecoder is the implementation of ValueDecoder for *bool.
type boolPtrDecoder struct {
	t reflect.Type
}

func (valdec boolPtrDecoder) decode(dec *Decoder, pv **bool, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if b := dec.decodeBool(valdec.t, tag); dec.Error == nil {
		*pv = &b
	}
}

func (valdec boolPtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (**bool)(reflect2.PtrOf(p)), tag)
}

func (valdec boolPtrDecoder) Type() reflect.Type {
	return valdec.t
}

var (
	bdec  = boolDecoder{reflect.TypeOf(false)}
	pbdec = boolPtrDecoder{reflect.TypeOf((*bool)(nil))}
)

func init() {
	RegisterValueDecoder(bdec)
	RegisterValueDecoder(pbdec)
}
