/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/int_decoder.go                                  |
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

// intDecoder is the implementation of ValueDecoder for int.
type intDecoder struct {
	descType reflect.Type
}

var intdec = intDecoder{reflect.TypeOf((*int)(nil)).Elem()}

func (valdec intDecoder) decode(dec *Decoder, tag byte) int {
	if i := intDigits[tag]; i != invalidDigit {
		return int(i)
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
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
		return int(dec.stringToInt64(dec.ReadString()))
	default:
		dec.decodeError(valdec.descType, tag)
	}
	return 0
}

func (valdec intDecoder) decodeValue(dec *Decoder, pv *int, tag byte) {
	if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec intDecoder) decodePtr(dec *Decoder, pv **int, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec intDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch pv := p.(type) {
	case *int:
		valdec.decodeValue(dec, pv, tag)
	case **int:
		valdec.decodePtr(dec, pv, tag)
	}
}

// int8Decoder is the implementation of ValueDecoder for int8.
type int8Decoder struct {
	descType reflect.Type
}

var int8dec = int8Decoder{reflect.TypeOf((*int8)(nil)).Elem()}

func (valdec int8Decoder) decode(dec *Decoder, tag byte) int8 {
	if i := intDigits[tag]; i != invalidDigit {
		return int8(i)
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
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
		return int8(dec.stringToInt64(dec.ReadString()))
	default:
		dec.decodeError(valdec.descType, tag)
	}
	return 0
}

func (valdec int8Decoder) decodeValue(dec *Decoder, pv *int8, tag byte) {
	if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec int8Decoder) decodePtr(dec *Decoder, pv **int8, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec int8Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch pv := p.(type) {
	case *int8:
		valdec.decodeValue(dec, pv, tag)
	case **int8:
		valdec.decodePtr(dec, pv, tag)
	}
}

// int16Decoder is the implementation of ValueDecoder for int16.
type int16Decoder struct {
	descType reflect.Type
}

var int16dec = int16Decoder{reflect.TypeOf((*int16)(nil)).Elem()}

func (valdec int16Decoder) decode(dec *Decoder, tag byte) int16 {
	if i := intDigits[tag]; i != invalidDigit {
		return int16(i)
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
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
		return int16(dec.stringToInt64(dec.ReadString()))
	default:
		dec.decodeError(valdec.descType, tag)
	}
	return 0
}

func (valdec int16Decoder) decodeValue(dec *Decoder, pv *int16, tag byte) {
	if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec int16Decoder) decodePtr(dec *Decoder, pv **int16, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec int16Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch pv := p.(type) {
	case *int16:
		valdec.decodeValue(dec, pv, tag)
	case **int16:
		valdec.decodePtr(dec, pv, tag)
	}
}

// int32Decoder is the implementation of ValueDecoder for int32.
type int32Decoder struct {
	descType reflect.Type
}

var int32dec = int32Decoder{reflect.TypeOf((*int32)(nil)).Elem()}

func (valdec int32Decoder) decode(dec *Decoder, tag byte) int32 {
	if i := intDigits[tag]; i != invalidDigit {
		return int32(i)
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
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
		return int32(dec.stringToInt64(dec.ReadString()))
	default:
		dec.decodeError(valdec.descType, tag)
	}
	return 0
}

func (valdec int32Decoder) decodeValue(dec *Decoder, pv *int32, tag byte) {
	if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec int32Decoder) decodePtr(dec *Decoder, pv **int32, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec int32Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch pv := p.(type) {
	case *int32:
		valdec.decodeValue(dec, pv, tag)
	case **int32:
		valdec.decodePtr(dec, pv, tag)
	}
}

// int64Decoder is the implementation of ValueDecoder for int64.
type int64Decoder struct {
	descType reflect.Type
}

var int64dec = int64Decoder{reflect.TypeOf((*int64)(nil)).Elem()}

func (valdec int64Decoder) decode(dec *Decoder, tag byte) int64 {
	if i := intDigits[tag]; i != invalidDigit {
		return int64(i)
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
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
		return dec.stringToInt64(dec.ReadString())
	default:
		dec.decodeError(valdec.descType, tag)
	}
	return 0
}

func (valdec int64Decoder) decodeValue(dec *Decoder, pv *int64, tag byte) {
	if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec int64Decoder) decodePtr(dec *Decoder, pv **int64, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec int64Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch pv := p.(type) {
	case *int64:
		valdec.decodeValue(dec, pv, tag)
	case **int64:
		valdec.decodePtr(dec, pv, tag)
	}
}

// uintDecoder is the implementation of ValueDecoder for uint.
type uintDecoder struct {
	descType reflect.Type
}

var uintdec = uintDecoder{reflect.TypeOf((*uint)(nil)).Elem()}

func (valdec uintDecoder) decode(dec *Decoder, tag byte) uint {
	if i := intDigits[tag]; i != invalidDigit {
		return uint(i)
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
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
		return uint(dec.stringToUint64(dec.ReadString()))
	default:
		dec.decodeError(valdec.descType, tag)
	}
	return 0
}

func (valdec uintDecoder) decodeValue(dec *Decoder, pv *uint, tag byte) {
	if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec uintDecoder) decodePtr(dec *Decoder, pv **uint, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec uintDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch pv := p.(type) {
	case *uint:
		valdec.decodeValue(dec, pv, tag)
	case **uint:
		valdec.decodePtr(dec, pv, tag)
	}
}

// uint8Decoder is the implementation of ValueDecoder for uint8.
type uint8Decoder struct {
	descType reflect.Type
}

var uint8dec = uint8Decoder{reflect.TypeOf((*uint8)(nil)).Elem()}

func (valdec uint8Decoder) decode(dec *Decoder, tag byte) uint8 {
	if i := intDigits[tag]; i != invalidDigit {
		return uint8(i)
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
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
		return uint8(dec.stringToUint64(dec.ReadString()))
	default:
		dec.decodeError(valdec.descType, tag)
	}
	return 0
}

func (valdec uint8Decoder) decodeValue(dec *Decoder, pv *uint8, tag byte) {
	if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec uint8Decoder) decodePtr(dec *Decoder, pv **uint8, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec uint8Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch pv := p.(type) {
	case *uint8:
		valdec.decodeValue(dec, pv, tag)
	case **uint8:
		valdec.decodePtr(dec, pv, tag)
	}
}

// uint16Decoder is the implementation of ValueDecoder for uint16.
type uint16Decoder struct {
	descType reflect.Type
}

var uint16dec = uint16Decoder{reflect.TypeOf((*uint16)(nil)).Elem()}

func (valdec uint16Decoder) decode(dec *Decoder, tag byte) uint16 {
	if i := intDigits[tag]; i != invalidDigit {
		return uint16(i)
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
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
		return uint16(dec.stringToUint64(dec.ReadString()))
	default:
		dec.decodeError(valdec.descType, tag)
	}
	return 0
}

func (valdec uint16Decoder) decodeValue(dec *Decoder, pv *uint16, tag byte) {
	if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec uint16Decoder) decodePtr(dec *Decoder, pv **uint16, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec uint16Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch pv := p.(type) {
	case *uint16:
		valdec.decodeValue(dec, pv, tag)
	case **uint16:
		valdec.decodePtr(dec, pv, tag)
	}
}

// uint32Decoder is the implementation of ValueDecoder for uint32.
type uint32Decoder struct {
	descType reflect.Type
}

var uint32dec = uint32Decoder{reflect.TypeOf((*uint32)(nil)).Elem()}

func (valdec uint32Decoder) decode(dec *Decoder, tag byte) uint32 {
	if i := intDigits[tag]; i != invalidDigit {
		return uint32(i)
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
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
		return uint32(dec.stringToUint64(dec.ReadString()))
	default:
		dec.decodeError(valdec.descType, tag)
	}
	return 0
}

func (valdec uint32Decoder) decodeValue(dec *Decoder, pv *uint32, tag byte) {
	if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec uint32Decoder) decodePtr(dec *Decoder, pv **uint32, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec uint32Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch pv := p.(type) {
	case *uint32:
		valdec.decodeValue(dec, pv, tag)
	case **uint32:
		valdec.decodePtr(dec, pv, tag)
	}
}

// uint64Decoder is the implementation of ValueDecoder for uint64.
type uint64Decoder struct {
	descType reflect.Type
}

var uint64dec = uint64Decoder{reflect.TypeOf((*uint64)(nil)).Elem()}

func (valdec uint64Decoder) decode(dec *Decoder, tag byte) uint64 {
	if i := intDigits[tag]; i != invalidDigit {
		return i
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
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
		return dec.stringToUint64(dec.ReadString())
	default:
		dec.decodeError(valdec.descType, tag)
	}
	return 0
}

func (valdec uint64Decoder) decodeValue(dec *Decoder, pv *uint64, tag byte) {
	if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec uint64Decoder) decodePtr(dec *Decoder, pv **uint64, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec uint64Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch pv := p.(type) {
	case *uint64:
		valdec.decodeValue(dec, pv, tag)
	case **uint64:
		valdec.decodePtr(dec, pv, tag)
	}
}

// uintptrDecoder is the implementation of ValueDecoder for uintptr.
type uintptrDecoder struct {
	descType reflect.Type
}

var uptrdec = uintptrDecoder{reflect.TypeOf((*uintptr)(nil)).Elem()}

func (valdec uintptrDecoder) decode(dec *Decoder, tag byte) uintptr {
	if i := intDigits[tag]; i != invalidDigit {
		return uintptr(i)
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
		return 0
	case TagTrue:
		return 1
	case TagInteger, TagLong:
		return uintptr(dec.ReadUint64())
	case TagDouble:
		return uintptr(dec.ReadFloat64())
	case TagUTF8Char:
		return uintptr(dec.stringToUint64(dec.readUnsafeString(1)))
	case TagString:
		return uintptr(dec.stringToUint64(dec.ReadString()))
	default:
		dec.decodeError(valdec.descType, tag)
	}
	return 0
}

func (valdec uintptrDecoder) decodeValue(dec *Decoder, pv *uintptr, tag byte) {
	if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec uintptrDecoder) decodePtr(dec *Decoder, pv **uintptr, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if i := valdec.decode(dec, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec uintptrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch pv := p.(type) {
	case *uintptr:
		valdec.decodeValue(dec, pv, tag)
	case **uintptr:
		valdec.decodePtr(dec, pv, tag)
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
