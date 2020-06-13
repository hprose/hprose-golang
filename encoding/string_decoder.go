/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/string_decoder.go                               |
|                                                          |
| LastModified: Jun 13, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"reflect"

	"github.com/modern-go/reflect2"
)

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
				return buf[:off], false
			}
			data = append(data, buf[:off]...)
			return
		}
		safe = true
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
	return append(([]byte)(nil), data...)
}

// ReadStringAsBytes reads string as bytes
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
	data := dec.readStringAsSafeBytes(utf16Length)
	if data == nil {
		return
	}
	return unsafeString(data)
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

func (dec *Decoder) decodeString(t reflect.Type, tag byte) string {
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
		return unsafeString(dec.Until(TagSemicolon))
	case TagUTF8Char:
		return dec.readSafeString(1)
	case TagString:
		return dec.ReadString()
	default:
		dec.decodeError(t, tag)
	}
	return ""
}

// stringDecoder is the implementation of ValueDecoder for string.
type stringDecoder struct {
	t reflect.Type
}

func (valdec stringDecoder) decode(dec *Decoder, pv *string, tag byte) {
	if s := dec.decodeString(valdec.t, tag); dec.Error == nil {
		*pv = s
	}
}

func (valdec stringDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (*string)(reflect2.PtrOf(p)), tag)
}

func (valdec stringDecoder) Type() reflect.Type {
	return valdec.t
}

// stringPtrDecoder is the implementation of ValueDecoder for *string.
type stringPtrDecoder struct {
	t reflect.Type
}

func (valdec stringPtrDecoder) decode(dec *Decoder, pv **string, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if s := dec.decodeString(valdec.t, tag); dec.Error == nil {
		*pv = &s
	}
}

func (valdec stringPtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (**string)(reflect2.PtrOf(p)), tag)
}

func (valdec stringPtrDecoder) Type() reflect.Type {
	return valdec.t
}

var (
	sdec  = stringDecoder{reflect.TypeOf("")}
	psdec = stringPtrDecoder{reflect.TypeOf((*string)(nil))}
)

func init() {
	RegisterValueDecoder(sdec)
	RegisterValueDecoder(psdec)
}
