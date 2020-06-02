/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/bool_decoder.go                                 |
|                                                          |
| LastModified: Jun 2, 2020                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"reflect"
	"strconv"
)

// boolDecoder is the implementation of ValueDecoder for bool.
type boolDecoder struct {
	descType reflect.Type
}

var booldec = boolDecoder{reflect.TypeOf((*bool)(nil)).Elem()}

func (valdec boolDecoder) decode(dec *Decoder, tag byte) bool {
	if i := intDigits[tag]; i != invalidDigit {
		return i > 0
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
		return false
	case TagTrue, TagNaN:
		return true
	case TagInteger, TagLong, TagDouble:
		bytes := dec.Until(TagSemicolon)
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
		return dec.stringToBool(dec.ReadString())
	default:
		dec.decodeError(valdec.descType, tag)
	}
	return false
}

func (valdec boolDecoder) decodeValue(dec *Decoder, pv *bool, tag byte) {
	if b := valdec.decode(dec, tag); dec.Error == nil {
		*pv = b
	}
}

func (valdec boolDecoder) decodePtr(dec *Decoder, pv **bool, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if b := valdec.decode(dec, tag); dec.Error == nil {
		*pv = &b
	}
}

func (valdec boolDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch pv := p.(type) {
	case *bool:
		valdec.decodeValue(dec, pv, tag)
	case **bool:
		valdec.decodePtr(dec, pv, tag)
	}
}

func (dec *Decoder) stringToBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		dec.Error = err
	}
	return b
}
