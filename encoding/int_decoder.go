/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/int_decoder.go                                  |
|                                                          |
| LastModified: May 1, 2020                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"strconv"
)

// intDecoder is the implementation of ValueDecoder for int.
type intDecoder struct{}

var intdec intDecoder

func (valdec intDecoder) decode(dec *Decoder, p interface{}, tag byte) int {
	if i := intDigits[tag]; i != invalidDigit {
		return int(i)
	}
	switch tag {
	case TagEmpty, TagFalse:
		return 0
	case TagTrue:
		return 1
	case TagInteger, TagLong:
		return dec.ReadInt()
	case TagDouble:
		return int(dec.ReadFloat64())
	case TagUTF8Char:
		return int(dec.stringToInt64(dec.readUnsafeString(1)))
	case TagString:
		return int(dec.stringToInt64(dec.ReadUnsafeString()))
	default:
		dec.decodeError(p, tag)
	}
	return 0
}

func (valdec intDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if tag == TagNull {
		switch pv := p.(type) {
		case **int:
			*pv = nil
		case *int:
			*pv = 0
		}
		return
	}
	i := valdec.decode(dec, p, tag)
	if dec.Error != nil {
		return
	}
	switch pv := p.(type) {
	case **int:
		*pv = &i
	case *int:
		*pv = i
	}
}

// int8Decoder is the implementation of ValueDecoder for int8.
type int8Decoder struct{}

var int8dec int8Decoder

func (valdec int8Decoder) decode(dec *Decoder, p interface{}, tag byte) int8 {
	if i := intDigits[tag]; i != invalidDigit {
		return int8(i)
	}
	switch tag {
	case TagEmpty, TagFalse:
		return 0
	case TagTrue:
		return 1
	case TagInteger, TagLong:
		return dec.ReadInt8()
	case TagDouble:
		return int8(dec.ReadFloat64())
	case TagUTF8Char:
		return int8(dec.stringToInt64(dec.readUnsafeString(1)))
	case TagString:
		return int8(dec.stringToInt64(dec.ReadUnsafeString()))
	default:
		dec.decodeError(p, tag)
	}
	return 0
}

func (valdec int8Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if tag == TagNull {
		switch pv := p.(type) {
		case **int8:
			*pv = nil
		case *int8:
			*pv = 0
		}
		return
	}
	i := valdec.decode(dec, p, tag)
	if dec.Error != nil {
		return
	}
	switch pv := p.(type) {
	case **int8:
		*pv = &i
	case *int8:
		*pv = i
	}
}

// int16Decoder is the implementation of ValueDecoder for int16.
type int16Decoder struct{}

var int16dec int16Decoder

func (valdec int16Decoder) decode(dec *Decoder, p interface{}, tag byte) int16 {
	if i := intDigits[tag]; i != invalidDigit {
		return int16(i)
	}
	switch tag {
	case TagEmpty, TagFalse:
		return 0
	case TagTrue:
		return 1
	case TagInteger, TagLong:
		return dec.ReadInt16()
	case TagDouble:
		return int16(dec.ReadFloat64())
	case TagUTF8Char:
		return int16(dec.stringToInt64(dec.readUnsafeString(1)))
	case TagString:
		return int16(dec.stringToInt64(dec.ReadUnsafeString()))
	default:
		dec.decodeError(p, tag)
	}
	return 0
}

func (valdec int16Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if tag == TagNull {
		switch pv := p.(type) {
		case **int16:
			*pv = nil
		case *int16:
			*pv = 0
		}
		return
	}
	i := valdec.decode(dec, p, tag)
	if dec.Error != nil {
		return
	}
	switch pv := p.(type) {
	case **int16:
		*pv = &i
	case *int16:
		*pv = i
	}
}

// int32Decoder is the implementation of ValueDecoder for int32.
type int32Decoder struct{}

var int32dec int32Decoder

func (valdec int32Decoder) decode(dec *Decoder, p interface{}, tag byte) int32 {
	if i := intDigits[tag]; i != invalidDigit {
		return int32(i)
	}
	switch tag {
	case TagEmpty, TagFalse:
		return 0
	case TagTrue:
		return 1
	case TagInteger, TagLong:
		return dec.ReadInt32()
	case TagDouble:
		return int32(dec.ReadFloat64())
	case TagUTF8Char:
		return int32(dec.stringToInt64(dec.readUnsafeString(1)))
	case TagString:
		return int32(dec.stringToInt64(dec.ReadUnsafeString()))
	default:
		dec.decodeError(p, tag)
	}
	return 0
}

func (valdec int32Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if tag == TagNull {
		switch pv := p.(type) {
		case **int32:
			*pv = nil
		case *int32:
			*pv = 0
		}
		return
	}
	i := valdec.decode(dec, p, tag)
	if dec.Error != nil {
		return
	}
	switch pv := p.(type) {
	case **int32:
		*pv = &i
	case *int32:
		*pv = i
	}
}

// int64Decoder is the implementation of ValueDecoder for int64.
type int64Decoder struct{}

var int64dec int64Decoder

func (valdec int64Decoder) decode(dec *Decoder, p interface{}, tag byte) int64 {
	if i := intDigits[tag]; i != invalidDigit {
		return int64(i)
	}
	switch tag {
	case TagEmpty, TagFalse:
		return 0
	case TagTrue:
		return 1
	case TagInteger, TagLong:
		return dec.ReadInt64()
	case TagDouble:
		return int64(dec.ReadFloat64())
	case TagUTF8Char:
		return dec.stringToInt64(dec.readUnsafeString(1))
	case TagString:
		return dec.stringToInt64(dec.ReadUnsafeString())
	default:
		dec.decodeError(p, tag)
	}
	return 0
}

func (valdec int64Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if tag == TagNull {
		switch pv := p.(type) {
		case **int64:
			*pv = nil
		case *int64:
			*pv = 0
		}
		return
	}
	i := valdec.decode(dec, p, tag)
	if dec.Error != nil {
		return
	}
	switch pv := p.(type) {
	case **int64:
		*pv = &i
	case *int64:
		*pv = i
	}
}

// uintDecoder is the implementation of ValueDecoder for uint.
type uintDecoder struct{}

var uintdec uintDecoder

func (valdec uintDecoder) decode(dec *Decoder, p interface{}, tag byte) uint {
	if i := intDigits[tag]; i != invalidDigit {
		return uint(i)
	}
	switch tag {
	case TagEmpty, TagFalse:
		return 0
	case TagTrue:
		return 1
	case TagInteger, TagLong:
		return dec.ReadUint()
	case TagDouble:
		return uint(dec.ReadFloat64())
	case TagUTF8Char:
		return uint(dec.stringToUint64(dec.readUnsafeString(1)))
	case TagString:
		return uint(dec.stringToUint64(dec.ReadUnsafeString()))
	default:
		dec.decodeError(p, tag)
	}
	return 0
}

func (valdec uintDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if tag == TagNull {
		switch pv := p.(type) {
		case **uint:
			*pv = nil
		case *uint:
			*pv = 0
		}
		return
	}
	i := valdec.decode(dec, p, tag)
	if dec.Error != nil {
		return
	}
	switch pv := p.(type) {
	case **uint:
		*pv = &i
	case *uint:
		*pv = i
	}
}

// uint8Decoder is the implementation of ValueDecoder for uint8.
type uint8Decoder struct{}

var uint8dec uint8Decoder

func (valdec uint8Decoder) decode(dec *Decoder, p interface{}, tag byte) uint8 {
	if i := intDigits[tag]; i != invalidDigit {
		return uint8(i)
	}
	switch tag {
	case TagEmpty, TagFalse:
		return 0
	case TagTrue:
		return 1
	case TagInteger, TagLong:
		return dec.ReadUint8()
	case TagDouble:
		return uint8(dec.ReadFloat64())
	case TagUTF8Char:
		return uint8(dec.stringToUint64(dec.readUnsafeString(1)))
	case TagString:
		return uint8(dec.stringToUint64(dec.ReadUnsafeString()))
	default:
		dec.decodeError(p, tag)
	}
	return 0
}

func (valdec uint8Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if tag == TagNull {
		switch pv := p.(type) {
		case **uint8:
			*pv = nil
		case *uint8:
			*pv = 0
		}
		return
	}
	i := valdec.decode(dec, p, tag)
	if dec.Error != nil {
		return
	}
	switch pv := p.(type) {
	case **uint8:
		*pv = &i
	case *uint8:
		*pv = i
	}
}

// uint16Decoder is the implementation of ValueDecoder for uint16.
type uint16Decoder struct{}

var uint16dec uint16Decoder

func (valdec uint16Decoder) decode(dec *Decoder, p interface{}, tag byte) uint16 {
	if i := intDigits[tag]; i != invalidDigit {
		return uint16(i)
	}
	switch tag {
	case TagEmpty, TagFalse:
		return 0
	case TagTrue:
		return 1
	case TagInteger, TagLong:
		return dec.ReadUint16()
	case TagDouble:
		return uint16(dec.ReadFloat64())
	case TagUTF8Char:
		return uint16(dec.stringToUint64(dec.readUnsafeString(1)))
	case TagString:
		return uint16(dec.stringToUint64(dec.ReadUnsafeString()))
	default:
		dec.decodeError(p, tag)
	}
	return 0
}

func (valdec uint16Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if tag == TagNull {
		switch pv := p.(type) {
		case **uint16:
			*pv = nil
		case *uint16:
			*pv = 0
		}
		return
	}
	i := valdec.decode(dec, p, tag)
	if dec.Error != nil {
		return
	}
	switch pv := p.(type) {
	case **uint16:
		*pv = &i
	case *uint16:
		*pv = i
	}
}

// uint32Decoder is the implementation of ValueDecoder for uint32.
type uint32Decoder struct{}

var uint32dec uint32Decoder

func (valdec uint32Decoder) decode(dec *Decoder, p interface{}, tag byte) uint32 {
	if i := intDigits[tag]; i != invalidDigit {
		return uint32(i)
	}
	switch tag {
	case TagEmpty, TagFalse:
		return 0
	case TagTrue:
		return 1
	case TagInteger, TagLong:
		return dec.ReadUint32()
	case TagDouble:
		return uint32(dec.ReadFloat64())
	case TagUTF8Char:
		return uint32(dec.stringToUint64(dec.readUnsafeString(1)))
	case TagString:
		return uint32(dec.stringToUint64(dec.ReadUnsafeString()))
	default:
		dec.decodeError(p, tag)
	}
	return 0
}

func (valdec uint32Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if tag == TagNull {
		switch pv := p.(type) {
		case **uint32:
			*pv = nil
		case *uint32:
			*pv = 0
		}
		return
	}
	i := valdec.decode(dec, p, tag)
	if dec.Error != nil {
		return
	}
	switch pv := p.(type) {
	case **uint32:
		*pv = &i
	case *uint32:
		*pv = i
	}
}

// uint64Decoder is the implementation of ValueDecoder for uint64.
type uint64Decoder struct{}

var uint64dec uint64Decoder

func (valdec uint64Decoder) decode(dec *Decoder, p interface{}, tag byte) uint64 {
	if i := intDigits[tag]; i != invalidDigit {
		return i
	}
	switch tag {
	case TagEmpty, TagFalse:
		return 0
	case TagTrue:
		return 1
	case TagInteger, TagLong:
		return dec.ReadUint64()
	case TagDouble:
		return uint64(dec.ReadFloat64())
	case TagUTF8Char:
		return dec.stringToUint64(dec.readUnsafeString(1))
	case TagString:
		return dec.stringToUint64(dec.ReadUnsafeString())
	default:
		dec.decodeError(p, tag)
	}
	return 0
}

func (valdec uint64Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if tag == TagNull {
		switch pv := p.(type) {
		case **uint64:
			*pv = nil
		case *uint64:
			*pv = 0
		}
		return
	}
	i := valdec.decode(dec, p, tag)
	if dec.Error != nil {
		return
	}
	switch pv := p.(type) {
	case **uint64:
		*pv = &i
	case *uint64:
		*pv = i
	}
}

func (dec *Decoder) stringToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		dec.Error = err
	}
	return i
}

func (dec *Decoder) stringToUint64(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		dec.Error = err
	}
	return i
}
