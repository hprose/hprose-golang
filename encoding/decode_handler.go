/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/decode_handler.go                               |
|                                                          |
| LastModified: Jun 14, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"math/big"
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

// DecodeHandler is an decode handler
type DecodeHandler func(dec *Decoder, t reflect.Type, p unsafe.Pointer)

func boolDecode(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
	*(*bool)(ep) = dec.decodeBool(et, dec.NextByte())
}

func intDecode(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
	*(*int)(ep) = dec.decodeInt(et, dec.NextByte())
}

func int8Decode(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
	*(*int8)(ep) = dec.decodeInt8(et, dec.NextByte())
}

func int16Decode(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
	*(*int16)(ep) = dec.decodeInt16(et, dec.NextByte())
}

func int32Decode(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
	*(*int32)(ep) = dec.decodeInt32(et, dec.NextByte())
}

func int64Decode(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
	*(*int64)(ep) = dec.decodeInt64(et, dec.NextByte())
}

func uintDecode(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
	*(*uint)(ep) = dec.decodeUint(et, dec.NextByte())
}

func uint8Decode(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
	*(*uint8)(ep) = dec.decodeUint8(et, dec.NextByte())
}

func uint16Decode(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
	*(*uint16)(ep) = dec.decodeUint16(et, dec.NextByte())
}

func uint32Decode(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
	*(*uint32)(ep) = dec.decodeUint32(et, dec.NextByte())
}

func uint64Decode(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
	*(*uint64)(ep) = dec.decodeUint64(et, dec.NextByte())
}

func uintptrDecode(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
	*(*uintptr)(ep) = dec.decodeUintptr(et, dec.NextByte())
}

func float32Decode(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
	*(*float32)(ep) = dec.decodeFloat32(et, dec.NextByte())
}

func float64Decode(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
	*(*float64)(ep) = dec.decodeFloat64(et, dec.NextByte())
}

func complex64Decode(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
	*(*complex64)(ep) = dec.decodeComplex64(et, dec.NextByte())
}

func complex128Decode(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
	*(*complex128)(ep) = dec.decodeComplex128(et, dec.NextByte())
}

func interfaceDecode(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
	*(*interface{})(ep) = dec.decodeInterface(dec.NextByte())
}

func stringDecode(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
	*(*string)(ep) = dec.decodeString(et, dec.NextByte())
}

func bytesDecode(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
	*(*[]byte)(ep) = dec.decodeBytes(et, dec.NextByte())
}

func bigIntDecode(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
	*(**big.Int)(ep) = dec.decodeBigInt(et, dec.NextByte())
}

func bigFloatDecode(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
	*(**big.Float)(ep) = dec.decodeBigFloat(et, dec.NextByte())
}

func bigRatDecode(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
	*(**big.Rat)(ep) = dec.decodeBigRat(et, dec.NextByte())
}

func otherDecode(t reflect.Type) DecodeHandler {
	valdec := getValueDecoder(t.Elem())
	et2 := reflect2.Type2(t.Elem())
	return func(dec *Decoder, et reflect.Type, ep unsafe.Pointer) {
		valdec.Decode(dec, et2.UnsafeIndirect(ep), dec.NextByte())
	}
}
