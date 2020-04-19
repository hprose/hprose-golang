/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/int_decoder.go                                  |
|                                                          |
| LastModified: Apr 19, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"strconv"
)

// intDecoder is the implementation of ValueDecoder for *int.
type intDecoder struct{}

var intdec intDecoder

func (intDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if pv, ok := p.(*int); ok {
		if i := intDigits[tag]; i != invalidDigit {
			*pv = int(i)
			return
		}
		switch tag {
		case TagNull, TagEmpty, TagFalse:
			*pv = 0
		case TagTrue:
			*pv = 1
		case TagInteger, TagLong:
			*pv = dec.ReadInt()
		case TagDouble:
			*pv = int(dec.ReadFloat64())
		case TagUTF8Char:
			*pv = int(dec.stringToInt64(dec.readUnsafeString(1)))
		case TagString:
			*pv = int(dec.stringToInt64(dec.ReadUnsafeString()))
		default:
			dec.castError(p)
		}
	}
}

// int8Decoder is the implementation of ValueDecoder for *int8.
type int8Decoder struct{}

var int8dec int8Decoder

func (int8Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if pv, ok := p.(*int8); ok {
		if i := intDigits[tag]; i != invalidDigit {
			*pv = int8(i)
			return
		}
		switch tag {
		case TagNull, TagEmpty, TagFalse:
			*pv = 0
		case TagTrue:
			*pv = 1
		case TagInteger, TagLong:
			*pv = dec.ReadInt8()
		case TagDouble:
			*pv = int8(dec.ReadFloat64())
		case TagUTF8Char:
			*pv = int8(dec.stringToInt64(dec.readUnsafeString(1)))
		case TagString:
			*pv = int8(dec.stringToInt64(dec.ReadUnsafeString()))
		default:
			dec.castError(p)
		}
	}
}

// int16Decoder is the implementation of ValueDecoder for *int16.
type int16Decoder struct{}

var int16dec int16Decoder

func (int16Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if pv, ok := p.(*int16); ok {
		if i := intDigits[tag]; i != invalidDigit {
			*pv = int16(i)
			return
		}
		switch tag {
		case TagNull, TagEmpty, TagFalse:
			*pv = 0
		case TagTrue:
			*pv = 1
		case TagInteger, TagLong:
			*pv = dec.ReadInt16()
		case TagDouble:
			*pv = int16(dec.ReadFloat64())
		case TagUTF8Char:
			*pv = int16(dec.stringToInt64(dec.readUnsafeString(1)))
		case TagString:
			*pv = int16(dec.stringToInt64(dec.ReadUnsafeString()))
		default:
			dec.castError(p)
		}
	}
}

// int32Decoder is the implementation of ValueDecoder for *int32.
type int32Decoder struct{}

var int32dec int32Decoder

func (int32Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if pv, ok := p.(*int32); ok {
		if i := intDigits[tag]; i != invalidDigit {
			*pv = int32(i)
			return
		}
		switch tag {
		case TagNull, TagEmpty, TagFalse:
			*pv = 0
		case TagTrue:
			*pv = 1
		case TagInteger, TagLong:
			*pv = dec.ReadInt32()
		case TagDouble:
			*pv = int32(dec.ReadFloat64())
		case TagUTF8Char:
			*pv = int32(dec.stringToInt64(dec.readUnsafeString(1)))
		case TagString:
			*pv = int32(dec.stringToInt64(dec.ReadUnsafeString()))
		default:
			dec.castError(p)
		}
	}
}

// int64Decoder is the implementation of ValueDecoder for *int64.
type int64Decoder struct{}

var int64dec int64Decoder

func (int64Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if pv, ok := p.(*int64); ok {
		if i := intDigits[tag]; i != invalidDigit {
			*pv = int64(i)
			return
		}
		switch tag {
		case TagNull, TagEmpty, TagFalse:
			*pv = 0
		case TagTrue:
			*pv = 1
		case TagInteger, TagLong:
			*pv = dec.ReadInt64()
		case TagDouble:
			*pv = int64(dec.ReadFloat64())
		case TagUTF8Char:
			*pv = dec.stringToInt64(dec.readUnsafeString(1))
		case TagString:
			*pv = dec.stringToInt64(dec.ReadUnsafeString())
		default:
			dec.castError(p)
		}
	}
}

// uintDecoder is the implementation of ValueDecoder for *uint.
type uintDecoder struct{}

var uintdec uintDecoder

func (uintDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if pv, ok := p.(*uint); ok {
		if i := intDigits[tag]; i != invalidDigit {
			*pv = uint(i)
			return
		}
		switch tag {
		case TagNull, TagEmpty, TagFalse:
			*pv = 0
		case TagTrue:
			*pv = 1
		case TagInteger, TagLong:
			*pv = dec.ReadUint()
		case TagDouble:
			*pv = uint(dec.ReadFloat64())
		case TagUTF8Char:
			*pv = uint(dec.stringToUint64(dec.readUnsafeString(1)))
		case TagString:
			*pv = uint(dec.stringToUint64(dec.ReadUnsafeString()))
		default:
			dec.castError(p)
		}
	}
}

// uint8Decoder is the implementation of ValueDecoder for *uint8.
type uint8Decoder struct{}

var uint8dec uint8Decoder

func (uint8Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if pv, ok := p.(*uint8); ok {
		if i := intDigits[tag]; i != invalidDigit {
			*pv = uint8(i)
			return
		}
		switch tag {
		case TagNull, TagEmpty, TagFalse:
			*pv = 0
		case TagTrue:
			*pv = 1
		case TagInteger, TagLong:
			*pv = dec.ReadUint8()
		case TagDouble:
			*pv = uint8(dec.ReadFloat64())
		case TagUTF8Char:
			*pv = uint8(dec.stringToUint64(dec.readUnsafeString(1)))
		case TagString:
			*pv = uint8(dec.stringToUint64(dec.ReadUnsafeString()))
		default:
			dec.castError(p)
		}
	}
}

// uint16Decoder is the implementation of ValueDecoder for *uint16.
type uint16Decoder struct{}

var uint16dec uint16Decoder

func (uint16Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if pv, ok := p.(*uint16); ok {
		if i := intDigits[tag]; i != invalidDigit {
			*pv = uint16(i)
			return
		}
		switch tag {
		case TagNull, TagEmpty, TagFalse:
			*pv = 0
		case TagTrue:
			*pv = 1
		case TagInteger, TagLong:
			*pv = dec.ReadUint16()
		case TagDouble:
			*pv = uint16(dec.ReadFloat64())
		case TagUTF8Char:
			*pv = uint16(dec.stringToUint64(dec.readUnsafeString(1)))
		case TagString:
			*pv = uint16(dec.stringToUint64(dec.ReadUnsafeString()))
		default:
			dec.castError(p)
		}
	}
}

// uint32Decoder is the implementation of ValueDecoder for *uint32.
type uint32Decoder struct{}

var uint32dec uint32Decoder

func (uint32Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if pv, ok := p.(*uint32); ok {
		if i := intDigits[tag]; i != invalidDigit {
			*pv = uint32(i)
			return
		}
		switch tag {
		case TagNull, TagEmpty, TagFalse:
			*pv = 0
		case TagTrue:
			*pv = 1
		case TagInteger, TagLong:
			*pv = dec.ReadUint32()
		case TagDouble:
			*pv = uint32(dec.ReadFloat64())
		case TagUTF8Char:
			*pv = uint32(dec.stringToUint64(dec.readUnsafeString(1)))
		case TagString:
			*pv = uint32(dec.stringToUint64(dec.ReadUnsafeString()))
		default:
			dec.castError(p)
		}
	}
}

// uint64Decoder is the implementation of ValueDecoder for *uint64.
type uint64Decoder struct{}

var uint64dec uint64Decoder

func (uint64Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if pv, ok := p.(*uint64); ok {
		if i := intDigits[tag]; i != invalidDigit {
			*pv = i
			return
		}
		switch tag {
		case TagNull, TagEmpty, TagFalse:
			*pv = 0
		case TagTrue:
			*pv = 1
		case TagInteger, TagLong:
			*pv = dec.ReadUint64()
		case TagDouble:
			*pv = uint64(dec.ReadFloat64())
		case TagUTF8Char:
			*pv = dec.stringToUint64(dec.readUnsafeString(1))
		case TagString:
			*pv = dec.stringToUint64(dec.ReadUnsafeString())
		default:
			dec.castError(p)
		}
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
