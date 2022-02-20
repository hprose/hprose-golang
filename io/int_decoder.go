/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/int_decoder.go                                        |
|                                                          |
| LastModified: Feb 20, 2022                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"reflect"
	"strconv"

	"github.com/modern-go/reflect2"
)

func (dec *Decoder) stringToInt64(s string, bitSize int) int64 {
	i, err := strconv.ParseInt(s, 10, bitSize)
	if err != nil {
		dec.Error = err
	}
	return i
}

func (dec *Decoder) stringToUint64(s string, bitSize int) uint64 {
	i, err := strconv.ParseUint(s, 10, bitSize)
	if err != nil {
		dec.Error = err
	}
	return i
}

func (dec *Decoder) decodeInt(t reflect.Type, tag byte, p *int) {
	if i := intDigits[tag]; i != invalidDigit {
		*p = int(i)
		return
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
		*p = 0
	case TagTrue:
		*p = 1
	case TagInteger, TagLong:
		*p = dec.ReadInt()
	case TagDouble:
		*p = int(dec.ReadFloat64())
	case TagUTF8Char:
		*p = int(dec.stringToInt64(dec.readUnsafeString(1), 0))
	case TagString:
		if dec.IsSimple() {
			*p = int(dec.stringToInt64(dec.ReadUnsafeString(), 0))
		} else {
			*p = int(dec.stringToInt64(dec.ReadString(), 0))
		}
	default:
		dec.defaultDecode(t, p, tag)
	}
}

func (dec *Decoder) decodeIntPtr(t reflect.Type, tag byte, p **int) {
	if tag == TagNull {
		*p = nil
		return
	}
	var i int
	dec.decodeInt(t, tag, &i)
	*p = &i
}

func (dec *Decoder) decodeInt8(t reflect.Type, tag byte, p *int8) {
	if i := intDigits[tag]; i != invalidDigit {
		*p = int8(i)
		return
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
		*p = 0
	case TagTrue:
		*p = 1
	case TagInteger, TagLong:
		*p = dec.ReadInt8()
	case TagDouble:
		*p = int8(dec.ReadFloat64())
	case TagUTF8Char:
		*p = int8(dec.stringToInt64(dec.readUnsafeString(1), 8))
	case TagString:
		if dec.IsSimple() {
			*p = int8(dec.stringToInt64(dec.ReadUnsafeString(), 8))
		} else {
			*p = int8(dec.stringToInt64(dec.ReadString(), 8))
		}
	default:
		dec.defaultDecode(t, p, tag)
	}
}

func (dec *Decoder) decodeInt8Ptr(t reflect.Type, tag byte, p **int8) {
	if tag == TagNull {
		*p = nil
		return
	}
	var i int8
	dec.decodeInt8(t, tag, &i)
	*p = &i
}

func (dec *Decoder) decodeInt16(t reflect.Type, tag byte, p *int16) {
	if i := intDigits[tag]; i != invalidDigit {
		*p = int16(i)
		return
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
		*p = 0
	case TagTrue:
		*p = 1
	case TagInteger, TagLong:
		*p = dec.ReadInt16()
	case TagDouble:
		*p = int16(dec.ReadFloat64())
	case TagUTF8Char:
		*p = int16(dec.stringToInt64(dec.readUnsafeString(1), 16))
	case TagString:
		if dec.IsSimple() {
			*p = int16(dec.stringToInt64(dec.ReadUnsafeString(), 16))
		} else {
			*p = int16(dec.stringToInt64(dec.ReadString(), 16))
		}
	default:
		dec.defaultDecode(t, p, tag)
	}
}

func (dec *Decoder) decodeInt16Ptr(t reflect.Type, tag byte, p **int16) {
	if tag == TagNull {
		*p = nil
		return
	}
	var i int16
	dec.decodeInt16(t, tag, &i)
	*p = &i
}

func (dec *Decoder) decodeInt32(t reflect.Type, tag byte, p *int32) {
	if i := intDigits[tag]; i != invalidDigit {
		*p = int32(i)
		return
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
		*p = 0
	case TagTrue:
		*p = 1
	case TagInteger, TagLong:
		*p = dec.ReadInt32()
	case TagDouble:
		*p = int32(dec.ReadFloat64())
	case TagUTF8Char:
		*p = int32(dec.stringToInt64(dec.readUnsafeString(1), 32))
	case TagString:
		if dec.IsSimple() {
			*p = int32(dec.stringToInt64(dec.ReadUnsafeString(), 32))
		} else {
			*p = int32(dec.stringToInt64(dec.ReadString(), 32))
		}
	default:
		dec.defaultDecode(t, p, tag)
	}
}

func (dec *Decoder) decodeInt32Ptr(t reflect.Type, tag byte, p **int32) {
	if tag == TagNull {
		*p = nil
		return
	}
	var i int32
	dec.decodeInt32(t, tag, &i)
	*p = &i
}

func (dec *Decoder) decodeInt64(t reflect.Type, tag byte, p *int64) {
	if i := intDigits[tag]; i != invalidDigit {
		*p = int64(i)
		return
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
		*p = 0
	case TagTrue:
		*p = 1
	case TagInteger, TagLong:
		*p = dec.ReadInt64()
	case TagDouble:
		*p = int64(dec.ReadFloat64())
	case TagUTF8Char:
		*p = dec.stringToInt64(dec.readUnsafeString(1), 64)
	case TagString:
		if dec.IsSimple() {
			*p = dec.stringToInt64(dec.ReadUnsafeString(), 64)
		} else {
			*p = dec.stringToInt64(dec.ReadString(), 64)
		}
	default:
		dec.defaultDecode(t, p, tag)
	}
}

func (dec *Decoder) decodeInt64Ptr(t reflect.Type, tag byte, p **int64) {
	if tag == TagNull {
		*p = nil
		return
	}
	var i int64
	dec.decodeInt64(t, tag, &i)
	*p = &i
}

func (dec *Decoder) decodeUint(t reflect.Type, tag byte, p *uint) {
	if i := intDigits[tag]; i != invalidDigit {
		*p = uint(i)
		return
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
		*p = 0
	case TagTrue:
		*p = 1
	case TagInteger, TagLong:
		*p = dec.ReadUint()
	case TagDouble:
		*p = uint(dec.ReadFloat64())
	case TagUTF8Char:
		*p = uint(dec.stringToUint64(dec.readUnsafeString(1), 0))
	case TagString:
		if dec.IsSimple() {
			*p = uint(dec.stringToUint64(dec.ReadUnsafeString(), 0))
		} else {
			*p = uint(dec.stringToUint64(dec.ReadString(), 0))
		}
	default:
		dec.defaultDecode(t, p, tag)
	}
}

func (dec *Decoder) decodeUintPtr(t reflect.Type, tag byte, p **uint) {
	if tag == TagNull {
		*p = nil
		return
	}
	var i uint
	dec.decodeUint(t, tag, &i)
	*p = &i
}

func (dec *Decoder) decodeUint8(t reflect.Type, tag byte, p *uint8) {
	if i := intDigits[tag]; i != invalidDigit {
		*p = uint8(i)
		return
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
		*p = 0
	case TagTrue:
		*p = 1
	case TagInteger, TagLong:
		*p = dec.ReadUint8()
	case TagDouble:
		*p = uint8(dec.ReadFloat64())
	case TagUTF8Char:
		*p = uint8(dec.stringToUint64(dec.readUnsafeString(1), 8))
	case TagString:
		if dec.IsSimple() {
			*p = uint8(dec.stringToUint64(dec.ReadUnsafeString(), 8))
		} else {
			*p = uint8(dec.stringToUint64(dec.ReadString(), 8))
		}
	default:
		dec.defaultDecode(t, p, tag)
	}
}

func (dec *Decoder) decodeUint8Ptr(t reflect.Type, tag byte, p **uint8) {
	if tag == TagNull {
		*p = nil
		return
	}
	var i uint8
	dec.decodeUint8(t, tag, &i)
	*p = &i
}

func (dec *Decoder) decodeUint16(t reflect.Type, tag byte, p *uint16) {
	if i := intDigits[tag]; i != invalidDigit {
		*p = uint16(i)
		return
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
		*p = 0
	case TagTrue:
		*p = 1
	case TagInteger, TagLong:
		*p = dec.ReadUint16()
	case TagDouble:
		*p = uint16(dec.ReadFloat64())
	case TagUTF8Char:
		*p = uint16(dec.stringToUint64(dec.readUnsafeString(1), 16))
	case TagString:
		if dec.IsSimple() {
			*p = uint16(dec.stringToUint64(dec.ReadUnsafeString(), 16))
		} else {
			*p = uint16(dec.stringToUint64(dec.ReadString(), 16))
		}
	default:
		dec.defaultDecode(t, p, tag)
	}
}

func (dec *Decoder) decodeUint16Ptr(t reflect.Type, tag byte, p **uint16) {
	if tag == TagNull {
		*p = nil
		return
	}
	var i uint16
	dec.decodeUint16(t, tag, &i)
	*p = &i
}

func (dec *Decoder) decodeUint32(t reflect.Type, tag byte, p *uint32) {
	if i := intDigits[tag]; i != invalidDigit {
		*p = uint32(i)
		return
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
		*p = 0
	case TagTrue:
		*p = 1
	case TagInteger, TagLong:
		*p = dec.ReadUint32()
	case TagDouble:
		*p = uint32(dec.ReadFloat64())
	case TagUTF8Char:
		*p = uint32(dec.stringToUint64(dec.readUnsafeString(1), 32))
	case TagString:
		if dec.IsSimple() {
			*p = uint32(dec.stringToUint64(dec.ReadUnsafeString(), 32))
		} else {
			*p = uint32(dec.stringToUint64(dec.ReadString(), 32))
		}
	default:
		dec.defaultDecode(t, p, tag)
	}
}

func (dec *Decoder) decodeUint32Ptr(t reflect.Type, tag byte, p **uint32) {
	if tag == TagNull {
		*p = nil
		return
	}
	var i uint32
	dec.decodeUint32(t, tag, &i)
	*p = &i
}

func (dec *Decoder) decodeUint64(t reflect.Type, tag byte, p *uint64) {
	if i := intDigits[tag]; i != invalidDigit {
		*p = i
		return
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
		*p = 0
	case TagTrue:
		*p = 1
	case TagInteger, TagLong:
		*p = dec.ReadUint64()
	case TagDouble:
		*p = uint64(dec.ReadFloat64())
	case TagUTF8Char:
		*p = dec.stringToUint64(dec.readUnsafeString(1), 64)
	case TagString:
		if dec.IsSimple() {
			*p = dec.stringToUint64(dec.ReadUnsafeString(), 64)
		} else {
			*p = dec.stringToUint64(dec.ReadString(), 64)
		}
	default:
		dec.defaultDecode(t, p, tag)
	}
}

func (dec *Decoder) decodeUint64Ptr(t reflect.Type, tag byte, p **uint64) {
	if tag == TagNull {
		*p = nil
		return
	}
	var i uint64
	dec.decodeUint64(t, tag, &i)
	*p = &i
}

func (dec *Decoder) decodeUintptr(t reflect.Type, tag byte, p *uintptr) {
	if i := intDigits[tag]; i != invalidDigit {
		*p = uintptr(i)
		return
	}
	switch tag {
	case TagNull, TagEmpty, TagFalse:
		*p = 0
	case TagTrue:
		*p = 1
	case TagInteger, TagLong:
		*p = uintptr(dec.ReadUint64())
	case TagDouble:
		*p = uintptr(dec.ReadFloat64())
	case TagUTF8Char:
		*p = uintptr(dec.stringToUint64(dec.readUnsafeString(1), 64))
	case TagString:
		if dec.IsSimple() {
			*p = uintptr(dec.stringToUint64(dec.ReadUnsafeString(), 64))
		} else {
			*p = uintptr(dec.stringToUint64(dec.ReadString(), 64))
		}
	default:
		dec.defaultDecode(t, p, tag)
	}
}

func (dec *Decoder) decodeUintptrPtr(t reflect.Type, tag byte, p **uintptr) {
	if tag == TagNull {
		*p = nil
		return
	}
	var i uintptr
	dec.decodeUintptr(t, tag, &i)
	*p = &i
}

// intDecoder is the implementation of ValueDecoder for int.
type intDecoder struct {
	t reflect.Type
}

func (valdec intDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeInt(valdec.t, tag, (*int)(reflect2.PtrOf(p)))
}

// intPtrDecoder is the implementation of ValueDecoder for *int.
type intPtrDecoder struct {
	t reflect.Type
}

func (valdec intPtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeIntPtr(valdec.t, tag, (**int)(reflect2.PtrOf(p)))
}

// int8Decoder is the implementation of ValueDecoder for int8.
type int8Decoder struct {
	t reflect.Type
}

func (valdec int8Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeInt8(valdec.t, tag, (*int8)(reflect2.PtrOf(p)))
}

// int8PtrDecoder is the implementation of ValueDecoder for *int8.
type int8PtrDecoder struct {
	t reflect.Type
}

func (valdec int8PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeInt8Ptr(valdec.t, tag, (**int8)(reflect2.PtrOf(p)))
}

// int16Decoder is the implementation of ValueDecoder for int16.
type int16Decoder struct {
	t reflect.Type
}

func (valdec int16Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeInt16(valdec.t, tag, (*int16)(reflect2.PtrOf(p)))
}

// int16PtrDecoder is the implementation of ValueDecoder for *int16.
type int16PtrDecoder struct {
	t reflect.Type
}

func (valdec int16PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeInt16Ptr(valdec.t, tag, (**int16)(reflect2.PtrOf(p)))
}

// int32Decoder is the implementation of ValueDecoder for int32.
type int32Decoder struct {
	t reflect.Type
}

func (valdec int32Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeInt32(valdec.t, tag, (*int32)(reflect2.PtrOf(p)))
}

// int32PtrDecoder is the implementation of ValueDecoder for *int32.
type int32PtrDecoder struct {
	t reflect.Type
}

func (valdec int32PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeInt32Ptr(valdec.t, tag, (**int32)(reflect2.PtrOf(p)))
}

// int64Decoder is the implementation of ValueDecoder for int64.
type int64Decoder struct {
	t reflect.Type
}

func (valdec int64Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeInt64(valdec.t, tag, (*int64)(reflect2.PtrOf(p)))
}

// int64PtrDecoder is the implementation of ValueDecoder for *int64.
type int64PtrDecoder struct {
	t reflect.Type
}

func (valdec int64PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeInt64Ptr(valdec.t, tag, (**int64)(reflect2.PtrOf(p)))
}

// uintDecoder is the implementation of ValueDecoder for uint.
type uintDecoder struct {
	t reflect.Type
}

func (valdec uintDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeUint(valdec.t, tag, (*uint)(reflect2.PtrOf(p)))
}

// uintPtrDecoder is the implementation of ValueDecoder for *uint.
type uintPtrDecoder struct {
	t reflect.Type
}

func (valdec uintPtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeUintPtr(valdec.t, tag, (**uint)(reflect2.PtrOf(p)))
}

// uint8Decoder is the implementation of ValueDecoder for uint8.
type uint8Decoder struct {
	t reflect.Type
}

func (valdec uint8Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeUint8(valdec.t, tag, (*uint8)(reflect2.PtrOf(p)))
}

// uint8PtrDecoder is the implementation of ValueDecoder for *uint8.
type uint8PtrDecoder struct {
	t reflect.Type
}

func (valdec uint8PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeUint8Ptr(valdec.t, tag, (**uint8)(reflect2.PtrOf(p)))
}

// uint16Decoder is the implementation of ValueDecoder for uint16.
type uint16Decoder struct {
	t reflect.Type
}

func (valdec uint16Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeUint16(valdec.t, tag, (*uint16)(reflect2.PtrOf(p)))
}

// uint16PtrDecoder is the implementation of ValueDecoder for *uint16.
type uint16PtrDecoder struct {
	t reflect.Type
}

func (valdec uint16PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeUint16Ptr(valdec.t, tag, (**uint16)(reflect2.PtrOf(p)))
}

// uint32Decoder is the implementation of ValueDecoder for uint32.
type uint32Decoder struct {
	t reflect.Type
}

func (valdec uint32Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeUint32(valdec.t, tag, (*uint32)(reflect2.PtrOf(p)))
}

// uint32PtrDecoder is the implementation of ValueDecoder for *uint32.
type uint32PtrDecoder struct {
	t reflect.Type
}

func (valdec uint32PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeUint32Ptr(valdec.t, tag, (**uint32)(reflect2.PtrOf(p)))
}

// uint64Decoder is the implementation of ValueDecoder for uint64.
type uint64Decoder struct {
	t reflect.Type
}

func (valdec uint64Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeUint64(valdec.t, tag, (*uint64)(reflect2.PtrOf(p)))
}

// uint64PtrDecoder is the implementation of ValueDecoder for *uint64.
type uint64PtrDecoder struct {
	t reflect.Type
}

func (valdec uint64PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeUint64Ptr(valdec.t, tag, (**uint64)(reflect2.PtrOf(p)))
}

// uintptrDecoder is the implementation of ValueDecoder for uintptr.
type uintptrDecoder struct {
	t reflect.Type
}

func (valdec uintptrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeUintptr(valdec.t, tag, (*uintptr)(reflect2.PtrOf(p)))
}

// uintptrPtrDecoder is the implementation of ValueDecoder for *uintptr.
type uintptrPtrDecoder struct {
	t reflect.Type
}

func (valdec uintptrPtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeUintptrPtr(valdec.t, tag, (**uintptr)(reflect2.PtrOf(p)))
}
