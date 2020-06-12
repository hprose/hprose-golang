/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/int_decoder.go                                  |
|                                                          |
| LastModified: Jun 12, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"reflect"
	"strconv"

	"github.com/modern-go/reflect2"
)

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

func (dec *Decoder) decodeInt(t reflect.Type, tag byte) int {
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
		if dec.IsSimple() {
			return int(dec.stringToInt64(dec.ReadUnsafeString()))
		}
		return int(dec.stringToInt64(dec.ReadString()))
	default:
		dec.decodeError(t, tag)
	}
	return 0
}

func (dec *Decoder) decodeInt8(t reflect.Type, tag byte) int8 {
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
		if dec.IsSimple() {
			return int8(dec.stringToInt64(dec.ReadUnsafeString()))
		}
		return int8(dec.stringToInt64(dec.ReadString()))
	default:
		dec.decodeError(t, tag)
	}
	return 0
}

func (dec *Decoder) decodeInt16(t reflect.Type, tag byte) int16 {
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
		if dec.IsSimple() {
			return int16(dec.stringToInt64(dec.ReadUnsafeString()))
		}
		return int16(dec.stringToInt64(dec.ReadString()))
	default:
		dec.decodeError(t, tag)
	}
	return 0
}

func (dec *Decoder) decodeInt32(t reflect.Type, tag byte) int32 {
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
		if dec.IsSimple() {
			return int32(dec.stringToInt64(dec.ReadUnsafeString()))
		}
		return int32(dec.stringToInt64(dec.ReadString()))
	default:
		dec.decodeError(t, tag)
	}
	return 0
}

func (dec *Decoder) decodeInt64(t reflect.Type, tag byte) int64 {
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
		if dec.IsSimple() {
			return dec.stringToInt64(dec.ReadUnsafeString())
		}
		return dec.stringToInt64(dec.ReadString())
	default:
		dec.decodeError(t, tag)
	}
	return 0
}

func (dec *Decoder) decodeUint(t reflect.Type, tag byte) uint {
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
		if dec.IsSimple() {
			return uint(dec.stringToUint64(dec.ReadUnsafeString()))
		}
		return uint(dec.stringToUint64(dec.ReadString()))
	default:
		dec.decodeError(t, tag)
	}
	return 0
}

func (dec *Decoder) decodeUint8(t reflect.Type, tag byte) uint8 {
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
		if dec.IsSimple() {
			return uint8(dec.stringToUint64(dec.ReadUnsafeString()))
		}
		return uint8(dec.stringToUint64(dec.ReadString()))
	default:
		dec.decodeError(t, tag)
	}
	return 0
}

func (dec *Decoder) decodeUint16(t reflect.Type, tag byte) uint16 {
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
		if dec.IsSimple() {
			return uint16(dec.stringToUint64(dec.ReadUnsafeString()))
		}
		return uint16(dec.stringToUint64(dec.ReadString()))
	default:
		dec.decodeError(t, tag)
	}
	return 0
}

func (dec *Decoder) decodeUint32(t reflect.Type, tag byte) uint32 {
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
		if dec.IsSimple() {
			return uint32(dec.stringToUint64(dec.ReadUnsafeString()))
		}
		return uint32(dec.stringToUint64(dec.ReadString()))
	default:
		dec.decodeError(t, tag)
	}
	return 0
}

func (dec *Decoder) decodeUint64(t reflect.Type, tag byte) uint64 {
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
		if dec.IsSimple() {
			return dec.stringToUint64(dec.ReadUnsafeString())
		}
		return dec.stringToUint64(dec.ReadString())
	default:
		dec.decodeError(t, tag)
	}
	return 0
}

func (dec *Decoder) decodeUintptr(t reflect.Type, tag byte) uintptr {
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
		if dec.IsSimple() {
			return uintptr(dec.stringToUint64(dec.ReadUnsafeString()))
		}
		return uintptr(dec.stringToUint64(dec.ReadString()))
	default:
		dec.decodeError(t, tag)
	}
	return 0
}

// intDecoder is the implementation of ValueDecoder for int.
type intDecoder struct {
	t reflect.Type
}

func (valdec intDecoder) decode(dec *Decoder, pv *int, tag byte) {
	if i := dec.decodeInt(valdec.t, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec intDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (*int)(reflect2.PtrOf(p)), tag)
}

func (valdec intDecoder) Type() reflect.Type {
	return valdec.t
}

// intPtrDecoder is the implementation of ValueDecoder for *int.
type intPtrDecoder struct {
	t reflect.Type
}

func (valdec intPtrDecoder) decode(dec *Decoder, pv **int, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if i := dec.decodeInt(valdec.t, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec intPtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (**int)(reflect2.PtrOf(p)), tag)
}

func (valdec intPtrDecoder) Type() reflect.Type {
	return valdec.t
}

// int8Decoder is the implementation of ValueDecoder for int8.
type int8Decoder struct {
	t reflect.Type
}

func (valdec int8Decoder) decode(dec *Decoder, pv *int8, tag byte) {
	if i := dec.decodeInt8(valdec.t, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec int8Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (*int8)(reflect2.PtrOf(p)), tag)
}

func (valdec int8Decoder) Type() reflect.Type {
	return valdec.t
}

// int8PtrDecoder is the implementation of ValueDecoder for *int8.
type int8PtrDecoder struct {
	t reflect.Type
}

func (valdec int8PtrDecoder) decode(dec *Decoder, pv **int8, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if i := dec.decodeInt8(valdec.t, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec int8PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (**int8)(reflect2.PtrOf(p)), tag)
}

func (valdec int8PtrDecoder) Type() reflect.Type {
	return valdec.t
}

// int16Decoder is the implementation of ValueDecoder for int16.
type int16Decoder struct {
	t reflect.Type
}

func (valdec int16Decoder) decode(dec *Decoder, pv *int16, tag byte) {
	if i := dec.decodeInt16(valdec.t, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec int16Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (*int16)(reflect2.PtrOf(p)), tag)
}

func (valdec int16Decoder) Type() reflect.Type {
	return valdec.t
}

// int16PtrDecoder is the implementation of ValueDecoder for *int16.
type int16PtrDecoder struct {
	t reflect.Type
}

func (valdec int16PtrDecoder) decode(dec *Decoder, pv **int16, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if i := dec.decodeInt16(valdec.t, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec int16PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (**int16)(reflect2.PtrOf(p)), tag)
}

func (valdec int16PtrDecoder) Type() reflect.Type {
	return valdec.t
}

// int32Decoder is the implementation of ValueDecoder for int32.
type int32Decoder struct {
	t reflect.Type
}

func (valdec int32Decoder) decode(dec *Decoder, pv *int32, tag byte) {
	if i := dec.decodeInt32(valdec.t, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec int32Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (*int32)(reflect2.PtrOf(p)), tag)
}

func (valdec int32Decoder) Type() reflect.Type {
	return valdec.t
}

// int32PtrDecoder is the implementation of ValueDecoder for *int32.
type int32PtrDecoder struct {
	t reflect.Type
}

func (valdec int32PtrDecoder) decode(dec *Decoder, pv **int32, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if i := dec.decodeInt32(valdec.t, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec int32PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (**int32)(reflect2.PtrOf(p)), tag)
}

func (valdec int32PtrDecoder) Type() reflect.Type {
	return valdec.t
}

// int64Decoder is the implementation of ValueDecoder for int64.
type int64Decoder struct {
	t reflect.Type
}

func (valdec int64Decoder) decode(dec *Decoder, pv *int64, tag byte) {
	if i := dec.decodeInt64(valdec.t, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec int64Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (*int64)(reflect2.PtrOf(p)), tag)
}

func (valdec int64Decoder) Type() reflect.Type {
	return valdec.t
}

// int64PtrDecoder is the implementation of ValueDecoder for *int64.
type int64PtrDecoder struct {
	t reflect.Type
}

func (valdec int64PtrDecoder) decode(dec *Decoder, pv **int64, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if i := dec.decodeInt64(valdec.t, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec int64PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (**int64)(reflect2.PtrOf(p)), tag)
}

func (valdec int64PtrDecoder) Type() reflect.Type {
	return valdec.t
}

// uintDecoder is the implementation of ValueDecoder for uint.
type uintDecoder struct {
	t reflect.Type
}

func (valdec uintDecoder) decode(dec *Decoder, pv *uint, tag byte) {
	if i := dec.decodeUint(valdec.t, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec uintDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (*uint)(reflect2.PtrOf(p)), tag)
}

func (valdec uintDecoder) Type() reflect.Type {
	return valdec.t
}

// uintPtrDecoder is the implementation of ValueDecoder for *uint.
type uintPtrDecoder struct {
	t reflect.Type
}

func (valdec uintPtrDecoder) decode(dec *Decoder, pv **uint, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if i := dec.decodeUint(valdec.t, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec uintPtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (**uint)(reflect2.PtrOf(p)), tag)
}

func (valdec uintPtrDecoder) Type() reflect.Type {
	return valdec.t
}

// uint8Decoder is the implementation of ValueDecoder for uint8.
type uint8Decoder struct {
	t reflect.Type
}

func (valdec uint8Decoder) decode(dec *Decoder, pv *uint8, tag byte) {
	if i := dec.decodeUint8(valdec.t, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec uint8Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (*uint8)(reflect2.PtrOf(p)), tag)
}

func (valdec uint8Decoder) Type() reflect.Type {
	return valdec.t
}

// uint8PtrDecoder is the implementation of ValueDecoder for *uint8.
type uint8PtrDecoder struct {
	t reflect.Type
}

func (valdec uint8PtrDecoder) decode(dec *Decoder, pv **uint8, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if i := dec.decodeUint8(valdec.t, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec uint8PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (**uint8)(reflect2.PtrOf(p)), tag)
}

func (valdec uint8PtrDecoder) Type() reflect.Type {
	return valdec.t
}

// uint16Decoder is the implementation of ValueDecoder for uint16.
type uint16Decoder struct {
	t reflect.Type
}

func (valdec uint16Decoder) decode(dec *Decoder, pv *uint16, tag byte) {
	if i := dec.decodeUint16(valdec.t, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec uint16Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (*uint16)(reflect2.PtrOf(p)), tag)
}

func (valdec uint16Decoder) Type() reflect.Type {
	return valdec.t
}

// uint16PtrDecoder is the implementation of ValueDecoder for *uint16.
type uint16PtrDecoder struct {
	t reflect.Type
}

func (valdec uint16PtrDecoder) decode(dec *Decoder, pv **uint16, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if i := dec.decodeUint16(valdec.t, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec uint16PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (**uint16)(reflect2.PtrOf(p)), tag)
}

func (valdec uint16PtrDecoder) Type() reflect.Type {
	return valdec.t
}

// uint32Decoder is the implementation of ValueDecoder for uint32.
type uint32Decoder struct {
	t reflect.Type
}

func (valdec uint32Decoder) decode(dec *Decoder, pv *uint32, tag byte) {
	if i := dec.decodeUint32(valdec.t, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec uint32Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (*uint32)(reflect2.PtrOf(p)), tag)
}

func (valdec uint32Decoder) Type() reflect.Type {
	return valdec.t
}

// uint32PtrDecoder is the implementation of ValueDecoder for *uint32.
type uint32PtrDecoder struct {
	t reflect.Type
}

func (valdec uint32PtrDecoder) decode(dec *Decoder, pv **uint32, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if i := dec.decodeUint32(valdec.t, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec uint32PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (**uint32)(reflect2.PtrOf(p)), tag)
}

func (valdec uint32PtrDecoder) Type() reflect.Type {
	return valdec.t
}

// uint64Decoder is the implementation of ValueDecoder for uint64.
type uint64Decoder struct {
	t reflect.Type
}

func (valdec uint64Decoder) decode(dec *Decoder, pv *uint64, tag byte) {
	if i := dec.decodeUint64(valdec.t, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec uint64Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (*uint64)(reflect2.PtrOf(p)), tag)
}

func (valdec uint64Decoder) Type() reflect.Type {
	return valdec.t
}

// uint64PtrDecoder is the implementation of ValueDecoder for *uint64.
type uint64PtrDecoder struct {
	t reflect.Type
}

func (valdec uint64PtrDecoder) decode(dec *Decoder, pv **uint64, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if i := dec.decodeUint64(valdec.t, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec uint64PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (**uint64)(reflect2.PtrOf(p)), tag)
}

func (valdec uint64PtrDecoder) Type() reflect.Type {
	return valdec.t
}

// uintptrDecoder is the implementation of ValueDecoder for uintptr.
type uintptrDecoder struct {
	t reflect.Type
}

func (valdec uintptrDecoder) decode(dec *Decoder, pv *uintptr, tag byte) {
	if i := dec.decodeUintptr(valdec.t, tag); dec.Error == nil {
		*pv = i
	}
}

func (valdec uintptrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (*uintptr)(reflect2.PtrOf(p)), tag)
}

func (valdec uintptrDecoder) Type() reflect.Type {
	return valdec.t
}

// uintptrPtrDecoder is the implementation of ValueDecoder for *uintptr.
type uintptrPtrDecoder struct {
	t reflect.Type
}

func (valdec uintptrPtrDecoder) decode(dec *Decoder, pv **uintptr, tag byte) {
	if tag == TagNull {
		*pv = nil
	} else if i := dec.decodeUintptr(valdec.t, tag); dec.Error == nil {
		*pv = &i
	}
}

func (valdec uintptrPtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (**uintptr)(reflect2.PtrOf(p)), tag)
}

func (valdec uintptrPtrDecoder) Type() reflect.Type {
	return valdec.t
}

var (
	idec    = intDecoder{reflect.TypeOf((int)(0))}
	i8dec   = int8Decoder{reflect.TypeOf((int8)(0))}
	i16dec  = int16Decoder{reflect.TypeOf((int16)(0))}
	i32dec  = int32Decoder{reflect.TypeOf((int32)(0))}
	i64dec  = int64Decoder{reflect.TypeOf((int64)(0))}
	udec    = uintDecoder{reflect.TypeOf((uint)(0))}
	u8dec   = uint8Decoder{reflect.TypeOf((uint8)(0))}
	u16dec  = uint16Decoder{reflect.TypeOf((uint16)(0))}
	u32dec  = uint32Decoder{reflect.TypeOf((uint32)(0))}
	u64dec  = uint64Decoder{reflect.TypeOf((uint64)(0))}
	updec   = uintptrDecoder{reflect.TypeOf((uintptr)(0))}
	pidec   = intPtrDecoder{reflect.TypeOf((*int)(nil))}
	pi8dec  = int8PtrDecoder{reflect.TypeOf((*int8)(nil))}
	pi16dec = int16PtrDecoder{reflect.TypeOf((*int16)(nil))}
	pi32dec = int32PtrDecoder{reflect.TypeOf((*int32)(nil))}
	pi64dec = int64PtrDecoder{reflect.TypeOf((*int64)(nil))}
	pudec   = uintPtrDecoder{reflect.TypeOf((*uint)(nil))}
	pu8dec  = uint8PtrDecoder{reflect.TypeOf((*uint8)(nil))}
	pu16dec = uint16PtrDecoder{reflect.TypeOf((*uint16)(nil))}
	pu32dec = uint32PtrDecoder{reflect.TypeOf((*uint32)(nil))}
	pu64dec = uint64PtrDecoder{reflect.TypeOf((*uint64)(nil))}
	pupdec  = uintptrPtrDecoder{reflect.TypeOf((*uintptr)(nil))}
)

func init() {
	RegisterValueDecoder(idec)
	RegisterValueDecoder(i8dec)
	RegisterValueDecoder(i16dec)
	RegisterValueDecoder(i32dec)
	RegisterValueDecoder(i64dec)
	RegisterValueDecoder(udec)
	RegisterValueDecoder(u8dec)
	RegisterValueDecoder(u16dec)
	RegisterValueDecoder(u32dec)
	RegisterValueDecoder(u64dec)
	RegisterValueDecoder(updec)
	RegisterValueDecoder(pidec)
	RegisterValueDecoder(pi8dec)
	RegisterValueDecoder(pi16dec)
	RegisterValueDecoder(pi32dec)
	RegisterValueDecoder(pi64dec)
	RegisterValueDecoder(pudec)
	RegisterValueDecoder(pu8dec)
	RegisterValueDecoder(pu16dec)
	RegisterValueDecoder(pu32dec)
	RegisterValueDecoder(pu64dec)
	RegisterValueDecoder(pupdec)
}
