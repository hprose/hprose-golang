/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/string_decoder.go                               |
|                                                          |
| LastModified: Apr 25, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import "reflect"

// stringDecoder is the implementation of ValueDecoder for string.
type stringDecoder struct {
	destType reflect.Type
}

var strdec = stringDecoder{reflect.TypeOf((*string)(nil)).Elem()}

func (valdec stringDecoder) decode(dec *Decoder, tag byte) string {
	switch tag {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return string(tag)
	case TagNull, TagEmpty:
		return ""
	case TagTrue:
		return "true"
	case TagFalse:
		return "false"
	case TagNaN:
		return "NaN"
	case TagInfinity:
		if dec.NextByte() == TagNeg {
			return "-Inf"
		}
		return "+Inf"
	case TagInteger, TagLong, TagDouble:
		return string(dec.Until(TagSemicolon))
	case TagUTF8Char:
		return dec.readSafeString(1)
	case TagString:
		return dec.ReadString()
	default:
		dec.decodeError(valdec.destType, tag)
	}
	return ""
}

func (valdec stringDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if tag == TagNull {
		switch pv := p.(type) {
		case **string:
			*pv = nil
		case *string:
			*pv = ""
		}
		return
	}
	s := valdec.decode(dec, tag)
	if dec.Error != nil {
		return
	}
	switch pv := p.(type) {
	case **string:
		*pv = &s
	case *string:
		*pv = s
	}
}

func (dec *Decoder) fastReadStringAsBytes(utf16Length int) (data []byte) {
	buf := dec.buf[dec.head:dec.tail]
	off := 0
	for ; utf16Length > 0; utf16Length-- {
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
				return
			}
			off += 4
			utf16Length--
		default:
			if dec.Error == nil {
				dec.Error = ErrInvalidUTF8
			}
			return
		}
	}
	dec.head += off
	return buf[:off]
}

func (dec *Decoder) readStringAsBytes(utf16Length int) (data []byte) {
	if (utf16Length == 0) || (dec.head == dec.tail) && !dec.loadMore() {
		return
	}
	length := dec.tail - dec.head
	if length >= utf16Length*3 {
		return dec.fastReadStringAsBytes(utf16Length)
	}
	for {
		buf := dec.buf[dec.head:dec.tail]
		off := 0
		for ; utf16Length > 0 && off < length; utf16Length-- {
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
					return
				}
				off += 4
				utf16Length--
			default:
				if dec.Error == nil {
					dec.Error = ErrInvalidUTF8
				}
				return
			}
		}
		remains := length - off
		if remains > 0 {
			dec.head += off
			if data == nil {
				return buf[:off]
			}
			data = append(data, buf[:off]...)
			return
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

func (dec *Decoder) readUnsafeString(utf16Length int) (s string) {
	data := dec.readStringAsBytes(utf16Length)
	if data == nil {
		return
	}
	return unsafeString(data)
}

func (dec *Decoder) readSafeString(utf16Length int) (s string) {
	data := dec.readStringAsBytes(utf16Length)
	if data == nil {
		return
	}
	return string(data)
}

// ReadUnsafeString reads unsafe string
func (dec *Decoder) ReadUnsafeString() (s string) {
	s = dec.readUnsafeString(dec.ReadInt())
	dec.Skip()
	return
}

// ReadSafeString reads safe string
func (dec *Decoder) ReadSafeString() (s string) {
	s = dec.readSafeString(dec.ReadInt())
	dec.Skip()
	return
}

// ReadString reads safe string and add reference
func (dec *Decoder) ReadString() (s string) {
	s = dec.ReadSafeString()
	dec.AddReference(s)
	return
}
