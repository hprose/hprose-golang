/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/int_decoder.go                                  |
|                                                          |
| LastModified: Jun 15, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

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
		return int(dec.stringToInt64(dec.readUnsafeString(1), 0))
	case TagString:
		if dec.IsSimple() {
			return int(dec.stringToInt64(dec.ReadUnsafeString(), 0))
		}
		return int(dec.stringToInt64(dec.ReadString(), 0))
	default:
		dec.decodeError(t, tag)
	}
	return 0
}

func (dec *Decoder) decodeIntPtr(t reflect.Type, tag byte) *int {
	if tag == TagNull {
		return nil
	}
	i := dec.decodeInt(t, tag)
	return &i
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
		return int8(dec.stringToInt64(dec.readUnsafeString(1), 8))
	case TagString:
		if dec.IsSimple() {
			return int8(dec.stringToInt64(dec.ReadUnsafeString(), 8))
		}
		return int8(dec.stringToInt64(dec.ReadString(), 8))
	default:
		dec.decodeError(t, tag)
	}
	return 0
}

func (dec *Decoder) decodeInt8Ptr(t reflect.Type, tag byte) *int8 {
	if tag == TagNull {
		return nil
	}
	i := dec.decodeInt8(t, tag)
	return &i
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
		return int16(dec.stringToInt64(dec.readUnsafeString(1), 16))
	case TagString:
		if dec.IsSimple() {
			return int16(dec.stringToInt64(dec.ReadUnsafeString(), 16))
		}
		return int16(dec.stringToInt64(dec.ReadString(), 16))
	default:
		dec.decodeError(t, tag)
	}
	return 0
}

func (dec *Decoder) decodeInt16Ptr(t reflect.Type, tag byte) *int16 {
	if tag == TagNull {
		return nil
	}
	i := dec.decodeInt16(t, tag)
	return &i
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
		return int32(dec.stringToInt64(dec.readUnsafeString(1), 32))
	case TagString:
		if dec.IsSimple() {
			return int32(dec.stringToInt64(dec.ReadUnsafeString(), 32))
		}
		return int32(dec.stringToInt64(dec.ReadString(), 32))
	default:
		dec.decodeError(t, tag)
	}
	return 0
}

func (dec *Decoder) decodeInt32Ptr(t reflect.Type, tag byte) *int32 {
	if tag == TagNull {
		return nil
	}
	i := dec.decodeInt32(t, tag)
	return &i
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
		return dec.stringToInt64(dec.readUnsafeString(1), 64)
	case TagString:
		if dec.IsSimple() {
			return dec.stringToInt64(dec.ReadUnsafeString(), 64)
		}
		return dec.stringToInt64(dec.ReadString(), 64)
	default:
		dec.decodeError(t, tag)
	}
	return 0
}

func (dec *Decoder) decodeInt64Ptr(t reflect.Type, tag byte) *int64 {
	if tag == TagNull {
		return nil
	}
	i := dec.decodeInt64(t, tag)
	return &i
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
		return uint(dec.stringToUint64(dec.readUnsafeString(1), 0))
	case TagString:
		if dec.IsSimple() {
			return uint(dec.stringToUint64(dec.ReadUnsafeString(), 0))
		}
		return uint(dec.stringToUint64(dec.ReadString(), 0))
	default:
		dec.decodeError(t, tag)
	}
	return 0
}

func (dec *Decoder) decodeUintPtr(t reflect.Type, tag byte) *uint {
	if tag == TagNull {
		return nil
	}
	i := dec.decodeUint(t, tag)
	return &i
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
		return uint8(dec.stringToUint64(dec.readUnsafeString(1), 8))
	case TagString:
		if dec.IsSimple() {
			return uint8(dec.stringToUint64(dec.ReadUnsafeString(), 8))
		}
		return uint8(dec.stringToUint64(dec.ReadString(), 8))
	default:
		dec.decodeError(t, tag)
	}
	return 0
}

func (dec *Decoder) decodeUint8Ptr(t reflect.Type, tag byte) *uint8 {
	if tag == TagNull {
		return nil
	}
	i := dec.decodeUint8(t, tag)
	return &i
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
		return uint16(dec.stringToUint64(dec.readUnsafeString(1), 16))
	case TagString:
		if dec.IsSimple() {
			return uint16(dec.stringToUint64(dec.ReadUnsafeString(), 16))
		}
		return uint16(dec.stringToUint64(dec.ReadString(), 16))
	default:
		dec.decodeError(t, tag)
	}
	return 0
}

func (dec *Decoder) decodeUint16Ptr(t reflect.Type, tag byte) *uint16 {
	if tag == TagNull {
		return nil
	}
	i := dec.decodeUint16(t, tag)
	return &i
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
		return uint32(dec.stringToUint64(dec.readUnsafeString(1), 32))
	case TagString:
		if dec.IsSimple() {
			return uint32(dec.stringToUint64(dec.ReadUnsafeString(), 32))
		}
		return uint32(dec.stringToUint64(dec.ReadString(), 32))
	default:
		dec.decodeError(t, tag)
	}
	return 0
}

func (dec *Decoder) decodeUint32Ptr(t reflect.Type, tag byte) *uint32 {
	if tag == TagNull {
		return nil
	}
	i := dec.decodeUint32(t, tag)
	return &i
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
		return dec.stringToUint64(dec.readUnsafeString(1), 64)
	case TagString:
		if dec.IsSimple() {
			return dec.stringToUint64(dec.ReadUnsafeString(), 64)
		}
		return dec.stringToUint64(dec.ReadString(), 64)
	default:
		dec.decodeError(t, tag)
	}
	return 0
}

func (dec *Decoder) decodeUint64Ptr(t reflect.Type, tag byte) *uint64 {
	if tag == TagNull {
		return nil
	}
	i := dec.decodeUint64(t, tag)
	return &i
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
		return uintptr(dec.stringToUint64(dec.readUnsafeString(1), 64))
	case TagString:
		if dec.IsSimple() {
			return uintptr(dec.stringToUint64(dec.ReadUnsafeString(), 64))
		}
		return uintptr(dec.stringToUint64(dec.ReadString(), 64))
	default:
		dec.decodeError(t, tag)
	}
	return 0
}

func (dec *Decoder) decodeUintptrPtr(t reflect.Type, tag byte) *uintptr {
	if tag == TagNull {
		return nil
	}
	i := dec.decodeUintptr(t, tag)
	return &i
}

// intDecoder is the implementation of ValueDecoder for int.
type intDecoder struct {
	t reflect.Type
}

func (valdec intDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*int)(reflect2.PtrOf(p)) = dec.decodeInt(valdec.t, tag)
}

func (valdec intDecoder) Type() reflect.Type {
	return valdec.t
}

// intPtrDecoder is the implementation of ValueDecoder for *int.
type intPtrDecoder struct {
	t reflect.Type
}

func (valdec intPtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**int)(reflect2.PtrOf(p)) = dec.decodeIntPtr(valdec.t, tag)
}

func (valdec intPtrDecoder) Type() reflect.Type {
	return valdec.t
}

// int8Decoder is the implementation of ValueDecoder for int8.
type int8Decoder struct {
	t reflect.Type
}

func (valdec int8Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*int8)(reflect2.PtrOf(p)) = dec.decodeInt8(valdec.t, tag)
}

func (valdec int8Decoder) Type() reflect.Type {
	return valdec.t
}

// int8PtrDecoder is the implementation of ValueDecoder for *int8.
type int8PtrDecoder struct {
	t reflect.Type
}

func (valdec int8PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**int8)(reflect2.PtrOf(p)) = dec.decodeInt8Ptr(valdec.t, tag)
}

func (valdec int8PtrDecoder) Type() reflect.Type {
	return valdec.t
}

// int16Decoder is the implementation of ValueDecoder for int16.
type int16Decoder struct {
	t reflect.Type
}

func (valdec int16Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*int16)(reflect2.PtrOf(p)) = dec.decodeInt16(valdec.t, tag)
}

func (valdec int16Decoder) Type() reflect.Type {
	return valdec.t
}

// int16PtrDecoder is the implementation of ValueDecoder for *int16.
type int16PtrDecoder struct {
	t reflect.Type
}

func (valdec int16PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*int16)(reflect2.PtrOf(p)) = dec.decodeInt16(valdec.t, tag)
}

func (valdec int16PtrDecoder) Type() reflect.Type {
	return valdec.t
}

// int32Decoder is the implementation of ValueDecoder for int32.
type int32Decoder struct {
	t reflect.Type
}

func (valdec int32Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*int32)(reflect2.PtrOf(p)) = dec.decodeInt32(valdec.t, tag)
}

func (valdec int32Decoder) Type() reflect.Type {
	return valdec.t
}

// int32PtrDecoder is the implementation of ValueDecoder for *int32.
type int32PtrDecoder struct {
	t reflect.Type
}

func (valdec int32PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**int32)(reflect2.PtrOf(p)) = dec.decodeInt32Ptr(valdec.t, tag)
}

func (valdec int32PtrDecoder) Type() reflect.Type {
	return valdec.t
}

// int64Decoder is the implementation of ValueDecoder for int64.
type int64Decoder struct {
	t reflect.Type
}

func (valdec int64Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*int64)(reflect2.PtrOf(p)) = dec.decodeInt64(valdec.t, tag)
}

func (valdec int64Decoder) Type() reflect.Type {
	return valdec.t
}

// int64PtrDecoder is the implementation of ValueDecoder for *int64.
type int64PtrDecoder struct {
	t reflect.Type
}

func (valdec int64PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**int64)(reflect2.PtrOf(p)) = dec.decodeInt64Ptr(valdec.t, tag)
}

func (valdec int64PtrDecoder) Type() reflect.Type {
	return valdec.t
}

// uintDecoder is the implementation of ValueDecoder for uint.
type uintDecoder struct {
	t reflect.Type
}

func (valdec uintDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*uint)(reflect2.PtrOf(p)) = dec.decodeUint(valdec.t, tag)
}

func (valdec uintDecoder) Type() reflect.Type {
	return valdec.t
}

// uintPtrDecoder is the implementation of ValueDecoder for *uint.
type uintPtrDecoder struct {
	t reflect.Type
}

func (valdec uintPtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**uint)(reflect2.PtrOf(p)) = dec.decodeUintPtr(valdec.t, tag)
}

func (valdec uintPtrDecoder) Type() reflect.Type {
	return valdec.t
}

// uint8Decoder is the implementation of ValueDecoder for uint8.
type uint8Decoder struct {
	t reflect.Type
}

func (valdec uint8Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*uint8)(reflect2.PtrOf(p)) = dec.decodeUint8(valdec.t, tag)
}

func (valdec uint8Decoder) Type() reflect.Type {
	return valdec.t
}

// uint8PtrDecoder is the implementation of ValueDecoder for *uint8.
type uint8PtrDecoder struct {
	t reflect.Type
}

func (valdec uint8PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**uint8)(reflect2.PtrOf(p)) = dec.decodeUint8Ptr(valdec.t, tag)
}

func (valdec uint8PtrDecoder) Type() reflect.Type {
	return valdec.t
}

// uint16Decoder is the implementation of ValueDecoder for uint16.
type uint16Decoder struct {
	t reflect.Type
}

func (valdec uint16Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*uint16)(reflect2.PtrOf(p)) = dec.decodeUint16(valdec.t, tag)
}

func (valdec uint16Decoder) Type() reflect.Type {
	return valdec.t
}

// uint16PtrDecoder is the implementation of ValueDecoder for *uint16.
type uint16PtrDecoder struct {
	t reflect.Type
}

func (valdec uint16PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**uint16)(reflect2.PtrOf(p)) = dec.decodeUint16Ptr(valdec.t, tag)
}

func (valdec uint16PtrDecoder) Type() reflect.Type {
	return valdec.t
}

// uint32Decoder is the implementation of ValueDecoder for uint32.
type uint32Decoder struct {
	t reflect.Type
}

func (valdec uint32Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*uint32)(reflect2.PtrOf(p)) = dec.decodeUint32(valdec.t, tag)
}

func (valdec uint32Decoder) Type() reflect.Type {
	return valdec.t
}

// uint32PtrDecoder is the implementation of ValueDecoder for *uint32.
type uint32PtrDecoder struct {
	t reflect.Type
}

func (valdec uint32PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**uint32)(reflect2.PtrOf(p)) = dec.decodeUint32Ptr(valdec.t, tag)
}

func (valdec uint32PtrDecoder) Type() reflect.Type {
	return valdec.t
}

// uint64Decoder is the implementation of ValueDecoder for uint64.
type uint64Decoder struct {
	t reflect.Type
}

func (valdec uint64Decoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*uint64)(reflect2.PtrOf(p)) = dec.decodeUint64(valdec.t, tag)
}

func (valdec uint64Decoder) Type() reflect.Type {
	return valdec.t
}

// uint64PtrDecoder is the implementation of ValueDecoder for *uint64.
type uint64PtrDecoder struct {
	t reflect.Type
}

func (valdec uint64PtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**uint64)(reflect2.PtrOf(p)) = dec.decodeUint64Ptr(valdec.t, tag)
}

func (valdec uint64PtrDecoder) Type() reflect.Type {
	return valdec.t
}

// uintptrDecoder is the implementation of ValueDecoder for uintptr.
type uintptrDecoder struct {
	t reflect.Type
}

func (valdec uintptrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*uintptr)(reflect2.PtrOf(p)) = dec.decodeUintptr(valdec.t, tag)
}

func (valdec uintptrDecoder) Type() reflect.Type {
	return valdec.t
}

// uintptrPtrDecoder is the implementation of ValueDecoder for *uintptr.
type uintptrPtrDecoder struct {
	t reflect.Type
}

func (valdec uintptrPtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**uintptr)(reflect2.PtrOf(p)) = dec.decodeUintptrPtr(valdec.t, tag)
}

func (valdec uintptrPtrDecoder) Type() reflect.Type {
	return valdec.t
}
