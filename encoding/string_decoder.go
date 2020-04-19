/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/string_decoder.go                               |
|                                                          |
| LastModified: Apr 18, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

// stringDecoder is the implementation of ValueDecoder for *string.
type stringDecoder struct{}

var strdec stringDecoder

func (stringDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if pv, ok := p.(*string); ok {
		switch tag {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			*pv = string(tag)
		case TagNull, TagEmpty:
			*pv = ""
		case TagTrue:
			*pv = "true"
		case TagFalse:
			*pv = "false"
		case TagNaN:
			*pv = "NaN"
		case TagInfinity:
			if dec.NextByte() == TagNeg {
				*pv = "-Inf"
			} else {
				*pv = "+Inf"
			}
		case TagInteger, TagLong, TagDouble:
			*pv = string(dec.Until(TagSemicolon))
		case TagUTF8Char:
			*pv = dec.readSafeString(1)
		case TagString:
			*pv = dec.ReadString()
		default:
			dec.castError(p)
		}
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
				dec.Error = ErrInvalidUTF8
				return
			}
			off += 4
			utf16Length--
		default:
			dec.Error = ErrInvalidUTF8
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
					dec.Error = ErrInvalidUTF8
					return
				}
				off += 4
				utf16Length--
			default:
				dec.Error = ErrInvalidUTF8
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
				dec.Error = ErrInvalidUTF8
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

// ReadUnsafeString read unsafe string
func (dec *Decoder) ReadUnsafeString() (s string) {
	s = dec.readUnsafeString(dec.ReadInt())
	dec.Skip()
	return
}

// ReadSafeString read safe string
func (dec *Decoder) ReadSafeString() (s string) {
	s = dec.readSafeString(dec.ReadInt())
	dec.Skip()
	return
}

// ReadString read safe string and add reference
func (dec *Decoder) ReadString() (s string) {
	s = dec.ReadSafeString()
	dec.AddReference(s)
	return
}
