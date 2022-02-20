/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/decode_handler.go                                     |
|                                                          |
| LastModified: Feb 20, 2022                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"math/big"
	"reflect"
	"time"
	"unsafe"

	"github.com/modern-go/reflect2"
)

// DecodeHandler is an decode handler.
type DecodeHandler func(dec *Decoder, t reflect.Type, p unsafe.Pointer)

func invalidDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeError(t, dec.NextByte())
}

func boolDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeBool(t, dec.NextByte(), (*bool)(p))
}

func intDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeInt(t, dec.NextByte(), (*int)(p))
}

func int8Decode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeInt8(t, dec.NextByte(), (*int8)(p))
}

func int16Decode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeInt16(t, dec.NextByte(), (*int16)(p))
}

func int32Decode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeInt32(t, dec.NextByte(), (*int32)(p))
}

func int64Decode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeInt64(t, dec.NextByte(), (*int64)(p))
}

func uintDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeUint(t, dec.NextByte(), (*uint)(p))
}

func uint8Decode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeUint8(t, dec.NextByte(), (*uint8)(p))
}

func uint16Decode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeUint16(t, dec.NextByte(), (*uint16)(p))
}

func uint32Decode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeUint32(t, dec.NextByte(), (*uint32)(p))
}

func uint64Decode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeUint64(t, dec.NextByte(), (*uint64)(p))
}

func uintptrDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeUintptr(t, dec.NextByte(), (*uintptr)(p))
}

func float32Decode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeFloat32(t, dec.NextByte(), (*float32)(p))
}

func float64Decode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeFloat64(t, dec.NextByte(), (*float64)(p))
}

func complex64Decode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeComplex64(t, dec.NextByte(), (*complex64)(p))
}

func complex128Decode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeComplex128(t, dec.NextByte(), (*complex128)(p))
}

func interfaceDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	*(*interface{})(p) = dec.decodeInterface(dec.NextByte())
}

func stringDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	*(*string)(p) = dec.decodeString(t, dec.NextByte())
}

func bytesDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeBytes(t, dec.NextByte(), (*[]byte)(p))
}

func timeDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	*(*time.Time)(p) = dec.decodeTime(t, dec.NextByte())
}

func bigIntDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	*(**big.Int)(p) = dec.decodeBigInt(t, dec.NextByte())
}

func bigFloatDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	*(**big.Float)(p) = dec.decodeBigFloat(t, dec.NextByte())
}

func bigRatDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	*(**big.Rat)(p) = dec.decodeBigRat(t, dec.NextByte())
}

func boolPtrDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeBoolPtr(t, dec.NextByte(), (**bool)(p))
}

func intPtrDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeIntPtr(t, dec.NextByte(), (**int)(p))
}

func int8PtrDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeInt8Ptr(t, dec.NextByte(), (**int8)(p))
}

func int16PtrDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeInt16Ptr(t, dec.NextByte(), (**int16)(p))
}

func int32PtrDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeInt32Ptr(t, dec.NextByte(), (**int32)(p))
}

func int64PtrDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeInt64Ptr(t, dec.NextByte(), (**int64)(p))
}

func uintPtrDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeUintPtr(t, dec.NextByte(), (**uint)(p))
}

func uint8PtrDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeUint8Ptr(t, dec.NextByte(), (**uint8)(p))
}

func uint16PtrDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeUint16Ptr(t, dec.NextByte(), (**uint16)(p))
}

func uint32PtrDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeUint32(t, dec.NextByte(), (*uint32)(p))
}

func uint64PtrDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeUint64Ptr(t, dec.NextByte(), (**uint64)(p))
}

func uintptrPtrDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeUintptrPtr(t, dec.NextByte(), (**uintptr)(p))
}

func float32PtrDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeFloat32Ptr(t, dec.NextByte(), (**float32)(p))
}

func float64PtrDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeFloat64Ptr(t, dec.NextByte(), (**float64)(p))
}

func complex64PtrDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeComplex64Ptr(t, dec.NextByte(), (**complex64)(p))
}

func complex128PtrDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	dec.decodeComplex128Ptr(t, dec.NextByte(), (**complex128)(p))
}

func interfacePtrDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	*(**interface{})(p) = dec.decodeInterfacePtr(dec.NextByte())
}

func stringPtrDecode(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
	*(**string)(p) = dec.decodeStringPtr(t, dec.NextByte())
}

func otherDecode(t reflect.Type) DecodeHandler {
	valdec := getValueDecoder(t)
	t2 := reflect2.Type2(t)
	return func(dec *Decoder, t reflect.Type, p unsafe.Pointer) {
		valdec.Decode(dec, t2.PackEFace(p), dec.NextByte())
	}
}

// GetDecodeHandler for specified type.
func GetDecodeHandler(t reflect.Type) DecodeHandler {
	if getRegisteredValueDecoder(t) == nil {
		kind := t.Kind()
		if decode := decodeHandlers[kind]; decode != nil {
			return decode
		}
		if kind == reflect.Ptr {
			if decode := decodePtrHandlers[t.Elem().Kind()]; decode != nil {
				return decode
			}
		}
	}
	return otherDecode(t)
}

var (
	decodeHandlers    []DecodeHandler
	decodePtrHandlers []DecodeHandler
)

func init() {
	decodeHandlers = []DecodeHandler{
		reflect.Invalid:       invalidDecode,
		reflect.Bool:          boolDecode,
		reflect.Int:           intDecode,
		reflect.Int8:          int8Decode,
		reflect.Int16:         int16Decode,
		reflect.Int32:         int32Decode,
		reflect.Int64:         int64Decode,
		reflect.Uint:          uintDecode,
		reflect.Uint8:         uint8Decode,
		reflect.Uint16:        uint16Decode,
		reflect.Uint32:        uint32Decode,
		reflect.Uint64:        uint64Decode,
		reflect.Uintptr:       uintptrDecode,
		reflect.Float32:       float32Decode,
		reflect.Float64:       float64Decode,
		reflect.Complex64:     complex64Decode,
		reflect.Complex128:    complex128Decode,
		reflect.Array:         nil,
		reflect.Chan:          invalidDecode,
		reflect.Func:          invalidDecode,
		reflect.Interface:     interfaceDecode,
		reflect.Map:           nil,
		reflect.Ptr:           nil,
		reflect.Slice:         nil,
		reflect.String:        stringDecode,
		reflect.Struct:        nil,
		reflect.UnsafePointer: invalidDecode,
	}
	decodePtrHandlers = []DecodeHandler{
		reflect.Invalid:       invalidDecode,
		reflect.Bool:          boolPtrDecode,
		reflect.Int:           intPtrDecode,
		reflect.Int8:          int8PtrDecode,
		reflect.Int16:         int16PtrDecode,
		reflect.Int32:         int32PtrDecode,
		reflect.Int64:         int64PtrDecode,
		reflect.Uint:          uintPtrDecode,
		reflect.Uint8:         uint8PtrDecode,
		reflect.Uint16:        uint16PtrDecode,
		reflect.Uint32:        uint32PtrDecode,
		reflect.Uint64:        uint64PtrDecode,
		reflect.Uintptr:       uintptrPtrDecode,
		reflect.Float32:       float32PtrDecode,
		reflect.Float64:       float64PtrDecode,
		reflect.Complex64:     complex64PtrDecode,
		reflect.Complex128:    complex128PtrDecode,
		reflect.Array:         nil,
		reflect.Chan:          invalidDecode,
		reflect.Func:          invalidDecode,
		reflect.Interface:     interfacePtrDecode,
		reflect.Map:           nil,
		reflect.Ptr:           nil,
		reflect.Slice:         nil,
		reflect.String:        stringPtrDecode,
		reflect.Struct:        nil,
		reflect.UnsafePointer: invalidDecode,
	}
}
