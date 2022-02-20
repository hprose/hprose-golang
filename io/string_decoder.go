/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/string_decoder.go                                     |
|                                                          |
| LastModified: Feb 20, 2022                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"reflect"

	"github.com/modern-go/reflect2"
)

func (dec *Decoder) checkUTF8String(buf []byte, off, utf16Length int) (int, int, bool) {
	b := buf[off]
	switch b >> 4 {
	case 0, 1, 2, 3, 4, 5, 6, 7:
		off++
	case 12, 13:
		off += 2
	case 14:
		off += 3
	case 15:
		if b&8 == 8 {
			if dec.Error == nil {
				dec.Error = ErrInvalidUTF8
			}
			return off, utf16Length, false
		}
		off += 4
		utf16Length--
	default:
		if dec.Error == nil {
			dec.Error = ErrInvalidUTF8
		}
		return off, utf16Length, false
	}
	return off, utf16Length, true
}

func (dec *Decoder) fastReadStringAsBytes(utf16Length int) (data []byte) {
	buf := dec.buf[dec.head:dec.tail]
	off := 0
	for ; utf16Length > 0; utf16Length-- {
		var ok bool
		if off, utf16Length, ok = dec.checkUTF8String(buf, off, utf16Length); !ok {
			return
		}
	}
	dec.head += off
	return buf[:off]
}

func (dec *Decoder) readStringAsBytes(utf16Length int) (data []byte, safe bool) {
	if (utf16Length == 0) || (dec.head == dec.tail) && !dec.loadMore() {
		return nil, true
	}
	length := dec.tail - dec.head
	if length >= utf16Length*3 {
		return dec.fastReadStringAsBytes(utf16Length), false
	}
	for {
		buf := dec.buf[dec.head:dec.tail]
		off := 0
		for ; utf16Length > 0 && off < length; utf16Length-- {
			var ok bool
			if off, utf16Length, ok = dec.checkUTF8String(buf, off, utf16Length); !ok {
				return
			}
		}
		remains := length - off
		if remains > 0 {
			dec.head += off
			if data == nil {
				return buf[:off], false
			}
			data = append(data, buf[:off]...)
			return
		}
		if !safe {
			safe = true
			data = make([]byte, 0, utf16Length*3)
		}
		data = append(data, buf...)
		if !dec.loadMore() {
			if remains < 0 {
				if dec.Error == nil {
					dec.Error = ErrInvalidUTF8
				}
			}
			return
		}
		data = append(data, dec.buf[dec.head:dec.head-remains]...)
		dec.head -= remains
		length = dec.tail - dec.head
	}
}

func (dec *Decoder) readStringAsSafeBytes(utf16Length int) []byte {
	data, safe := dec.readStringAsBytes(utf16Length)
	if safe {
		return data
	}
	result := make([]byte, len(data))
	copy(result, data)
	return result
}

// ReadStringAsBytes reads string as bytes.
func (dec *Decoder) ReadStringAsBytes() (data []byte) {
	data = dec.readStringAsSafeBytes(dec.ReadInt())
	dec.Skip()
	return
}

func (dec *Decoder) readUnsafeString(utf16Length int) (s string) {
	data, _ := dec.readStringAsBytes(utf16Length)
	if data == nil {
		return
	}
	return unsafeString(data)
}

func (dec *Decoder) readSafeString(utf16Length int) (s string) {
	data, safe := dec.readStringAsBytes(utf16Length)
	if data == nil {
		return
	}
	if safe {
		return unsafeString(data)
	}
	return string(data)
}

// ReadUnsafeString reads unsafe string.
func (dec *Decoder) ReadUnsafeString() (s string) {
	s = dec.readUnsafeString(dec.ReadInt())
	dec.Skip()
	return
}

// ReadSafeString reads safe string.
func (dec *Decoder) ReadSafeString() (s string) {
	s = dec.readSafeString(dec.ReadInt())
	dec.Skip()
	return
}

// ReadString reads safe string and add reference.
func (dec *Decoder) ReadString() (s string) {
	s = dec.ReadSafeString()
	if !dec.IsSimple() {
		dec.refer.Add(s)
	}
	return
}

func (dec *Decoder) decodeString(t reflect.Type, tag byte, p *string) {
	switch tag {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		*p = string(tag)
	case TagNull, TagEmpty:
		*p = ""
	case TagTrue:
		*p = "true"
	case TagFalse:
		*p = "false"
	case TagNaN:
		*p = "NaN"
	case TagInfinity:
		if dec.NextByte() == TagNeg {
			*p = "-Inf"
		} else {
			*p = "+Inf"
		}
	case TagInteger, TagLong, TagDouble:
		*p = unsafeString(dec.Until(TagSemicolon))
	case TagUTF8Char:
		*p = dec.readSafeString(1)
	case TagString:
		*p = dec.ReadString()
	case TagBytes:
		*p = unsafeString(dec.ReadBytes())
	case TagTime:
		*p = dec.ReadTime().String()
	case TagDate:
		*p = dec.ReadDateTime().String()
	case TagGUID:
		*p = dec.ReadUUID().String()
	default:
		dec.defaultDecode(t, p, tag)
	}
}

func (dec *Decoder) decodeStringPtr(t reflect.Type, tag byte, p **string) {
	if tag == TagNull {
		*p = nil
		return
	}
	var s string
	dec.decodeString(t, tag, &s)
	*p = &s
}

// stringDecoder is the implementation of ValueDecoder for string.
type stringDecoder struct {
	t reflect.Type
}

func (valdec stringDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeString(valdec.t, tag, (*string)(reflect2.PtrOf(p)))
}

// stringPtrDecoder is the implementation of ValueDecoder for *string.
type stringPtrDecoder struct {
	t reflect.Type
}

func (valdec stringPtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeStringPtr(valdec.t, tag, (**string)(reflect2.PtrOf(p)))
}
