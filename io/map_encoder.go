/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/map_encoder.go                                        |
|                                                          |
| LastModified: Feb 18, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

// mapEncoder is the implementation of ValueEncoder for *map.
type mapEncoder struct{}

var mapenc mapEncoder

func (valenc mapEncoder) Encode(enc *Encoder, v interface{}) {
	if ok := enc.WriteReference(v); !ok {
		valenc.Write(enc, v)
	}
}

func (mapEncoder) Write(enc *Encoder, v interface{}) {
	enc.setReference(v)
	enc.writeMap(reflect.ValueOf(v).Elem().Interface())
}

// WriteMap to encoder.
func (enc *Encoder) WriteMap(v interface{}) {
	enc.AddReferenceCount(1)
	enc.writeMap(v)
}

func (enc *Encoder) writeMap(v interface{}) {
	count := reflect.ValueOf(v).Len()
	if count == 0 {
		enc.buf = append(enc.buf, TagMap, TagOpenbrace, TagClosebrace)
		return
	}
	enc.WriteMapHead(count)
	enc.writeMapBody(v)
	enc.WriteFoot()
}

//nolint
func (enc *Encoder) writeStringMapBody(v interface{}) bool {
	switch v := v.(type) {
	case map[string]interface{}:
		enc.writeStringInterfaceMapBody(v)
	case map[string]string:
		enc.writeStringStringMapBody(v)
	case map[string]int:
		enc.writeStringIntMapBody(v)
	case map[string]int8:
		enc.writeStringInt8MapBody(v)
	case map[string]int16:
		enc.writeStringInt16MapBody(v)
	case map[string]int32:
		enc.writeStringInt32MapBody(v)
	case map[string]int64:
		enc.writeStringInt64MapBody(v)
	case map[string]uint:
		enc.writeStringUintMapBody(v)
	case map[string]uint8:
		enc.writeStringUint8MapBody(v)
	case map[string]uint16:
		enc.writeStringUint16MapBody(v)
	case map[string]uint32:
		enc.writeStringUint32MapBody(v)
	case map[string]uint64:
		enc.writeStringUint64MapBody(v)
	case map[string]bool:
		enc.writeStringBoolMapBody(v)
	case map[string]float32:
		enc.writeStringFloat32MapBody(v)
	case map[string]float64:
		enc.writeStringFloat64MapBody(v)
	default:
		return false
	}
	return true
}

//nolint
func (enc *Encoder) writeInterfaceMapBody(v interface{}) bool {
	switch v := v.(type) {
	case map[interface{}]interface{}:
		enc.writeInterfaceInterfaceMapBody(v)
	case map[interface{}]string:
		enc.writeInterfaceStringMapBody(v)
	case map[interface{}]int:
		enc.writeInterfaceIntMapBody(v)
	case map[interface{}]int8:
		enc.writeInterfaceInt8MapBody(v)
	case map[interface{}]int16:
		enc.writeInterfaceInt16MapBody(v)
	case map[interface{}]int32:
		enc.writeInterfaceInt32MapBody(v)
	case map[interface{}]int64:
		enc.writeInterfaceInt64MapBody(v)
	case map[interface{}]uint:
		enc.writeInterfaceUintMapBody(v)
	case map[interface{}]uint8:
		enc.writeInterfaceUint8MapBody(v)
	case map[interface{}]uint16:
		enc.writeInterfaceUint16MapBody(v)
	case map[interface{}]uint32:
		enc.writeInterfaceUint32MapBody(v)
	case map[interface{}]uint64:
		enc.writeInterfaceUint64MapBody(v)
	case map[interface{}]bool:
		enc.writeInterfaceBoolMapBody(v)
	case map[interface{}]float32:
		enc.writeInterfaceFloat32MapBody(v)
	case map[interface{}]float64:
		enc.writeInterfaceFloat64MapBody(v)
	default:
		return false
	}
	return true
}

//nolint
func (enc *Encoder) writeIntMapBody(v interface{}) bool {
	switch v := v.(type) {
	case map[int]interface{}:
		enc.writeIntInterfaceMapBody(v)
	case map[int]string:
		enc.writeIntStringMapBody(v)
	case map[int]int:
		enc.writeIntIntMapBody(v)
	case map[int]int8:
		enc.writeIntInt8MapBody(v)
	case map[int]int16:
		enc.writeIntInt16MapBody(v)
	case map[int]int32:
		enc.writeIntInt32MapBody(v)
	case map[int]int64:
		enc.writeIntInt64MapBody(v)
	case map[int]uint:
		enc.writeIntUintMapBody(v)
	case map[int]uint8:
		enc.writeIntUint8MapBody(v)
	case map[int]uint16:
		enc.writeIntUint16MapBody(v)
	case map[int]uint32:
		enc.writeIntUint32MapBody(v)
	case map[int]uint64:
		enc.writeIntUint64MapBody(v)
	case map[int]bool:
		enc.writeIntBoolMapBody(v)
	case map[int]float32:
		enc.writeIntFloat32MapBody(v)
	case map[int]float64:
		enc.writeIntFloat64MapBody(v)
	default:
		return false
	}
	return true
}

//nolint
func (enc *Encoder) writeInt8MapBody(v interface{}) bool {
	switch v := v.(type) {
	case map[int8]interface{}:
		enc.writeInt8InterfaceMapBody(v)
	case map[int8]string:
		enc.writeInt8StringMapBody(v)
	case map[int8]int:
		enc.writeInt8IntMapBody(v)
	case map[int8]int8:
		enc.writeInt8Int8MapBody(v)
	case map[int8]int16:
		enc.writeInt8Int16MapBody(v)
	case map[int8]int32:
		enc.writeInt8Int32MapBody(v)
	case map[int8]int64:
		enc.writeInt8Int64MapBody(v)
	case map[int8]uint:
		enc.writeInt8UintMapBody(v)
	case map[int8]uint8:
		enc.writeInt8Uint8MapBody(v)
	case map[int8]uint16:
		enc.writeInt8Uint16MapBody(v)
	case map[int8]uint32:
		enc.writeInt8Uint32MapBody(v)
	case map[int8]uint64:
		enc.writeInt8Uint64MapBody(v)
	case map[int8]bool:
		enc.writeInt8BoolMapBody(v)
	case map[int8]float32:
		enc.writeInt8Float32MapBody(v)
	case map[int8]float64:
		enc.writeInt8Float64MapBody(v)
	default:
		return false
	}
	return true
}

//nolint
func (enc *Encoder) writeInt16MapBody(v interface{}) bool {
	switch v := v.(type) {
	case map[int16]interface{}:
		enc.writeInt16InterfaceMapBody(v)
	case map[int16]string:
		enc.writeInt16StringMapBody(v)
	case map[int16]int:
		enc.writeInt16IntMapBody(v)
	case map[int16]int8:
		enc.writeInt16Int8MapBody(v)
	case map[int16]int16:
		enc.writeInt16Int16MapBody(v)
	case map[int16]int32:
		enc.writeInt16Int32MapBody(v)
	case map[int16]int64:
		enc.writeInt16Int64MapBody(v)
	case map[int16]uint:
		enc.writeInt16UintMapBody(v)
	case map[int16]uint8:
		enc.writeInt16Uint8MapBody(v)
	case map[int16]uint16:
		enc.writeInt16Uint16MapBody(v)
	case map[int16]uint32:
		enc.writeInt16Uint32MapBody(v)
	case map[int16]uint64:
		enc.writeInt16Uint64MapBody(v)
	case map[int16]bool:
		enc.writeInt16BoolMapBody(v)
	case map[int16]float32:
		enc.writeInt16Float32MapBody(v)
	case map[int16]float64:
		enc.writeInt16Float64MapBody(v)
	default:
		return false
	}
	return true
}

//nolint
func (enc *Encoder) writeInt32MapBody(v interface{}) bool {
	switch v := v.(type) {
	case map[int32]interface{}:
		enc.writeInt32InterfaceMapBody(v)
	case map[int32]string:
		enc.writeInt32StringMapBody(v)
	case map[int32]int:
		enc.writeInt32IntMapBody(v)
	case map[int32]int8:
		enc.writeInt32Int8MapBody(v)
	case map[int32]int16:
		enc.writeInt32Int16MapBody(v)
	case map[int32]int32:
		enc.writeInt32Int32MapBody(v)
	case map[int32]int64:
		enc.writeInt32Int64MapBody(v)
	case map[int32]uint:
		enc.writeInt32UintMapBody(v)
	case map[int32]uint8:
		enc.writeInt32Uint8MapBody(v)
	case map[int32]uint16:
		enc.writeInt32Uint16MapBody(v)
	case map[int32]uint32:
		enc.writeInt32Uint32MapBody(v)
	case map[int32]uint64:
		enc.writeInt32Uint64MapBody(v)
	case map[int32]bool:
		enc.writeInt32BoolMapBody(v)
	case map[int32]float32:
		enc.writeInt32Float32MapBody(v)
	case map[int32]float64:
		enc.writeInt32Float64MapBody(v)
	default:
		return false
	}
	return true
}

//nolint
func (enc *Encoder) writeInt64MapBody(v interface{}) bool {
	switch v := v.(type) {
	case map[int64]interface{}:
		enc.writeInt64InterfaceMapBody(v)
	case map[int64]string:
		enc.writeInt64StringMapBody(v)
	case map[int64]int:
		enc.writeInt64IntMapBody(v)
	case map[int64]int8:
		enc.writeInt64Int8MapBody(v)
	case map[int64]int16:
		enc.writeInt64Int16MapBody(v)
	case map[int64]int32:
		enc.writeInt64Int32MapBody(v)
	case map[int64]int64:
		enc.writeInt64Int64MapBody(v)
	case map[int64]uint:
		enc.writeInt64UintMapBody(v)
	case map[int64]uint8:
		enc.writeInt64Uint8MapBody(v)
	case map[int64]uint16:
		enc.writeInt64Uint16MapBody(v)
	case map[int64]uint32:
		enc.writeInt64Uint32MapBody(v)
	case map[int64]uint64:
		enc.writeInt64Uint64MapBody(v)
	case map[int64]bool:
		enc.writeInt64BoolMapBody(v)
	case map[int64]float32:
		enc.writeInt64Float32MapBody(v)
	case map[int64]float64:
		enc.writeInt64Float64MapBody(v)
	default:
		return false
	}
	return true
}

//nolint
func (enc *Encoder) writeUintMapBody(v interface{}) bool {
	switch v := v.(type) {
	case map[uint]interface{}:
		enc.writeUintInterfaceMapBody(v)
	case map[uint]string:
		enc.writeUintStringMapBody(v)
	case map[uint]int:
		enc.writeUintIntMapBody(v)
	case map[uint]int8:
		enc.writeUintInt8MapBody(v)
	case map[uint]int16:
		enc.writeUintInt16MapBody(v)
	case map[uint]int32:
		enc.writeUintInt32MapBody(v)
	case map[uint]int64:
		enc.writeUintInt64MapBody(v)
	case map[uint]uint:
		enc.writeUintUintMapBody(v)
	case map[uint]uint8:
		enc.writeUintUint8MapBody(v)
	case map[uint]uint16:
		enc.writeUintUint16MapBody(v)
	case map[uint]uint32:
		enc.writeUintUint32MapBody(v)
	case map[uint]uint64:
		enc.writeUintUint64MapBody(v)
	case map[uint]bool:
		enc.writeUintBoolMapBody(v)
	case map[uint]float32:
		enc.writeUintFloat32MapBody(v)
	case map[uint]float64:
		enc.writeUintFloat64MapBody(v)
	default:
		return false
	}
	return true
}

//nolint
func (enc *Encoder) writeUint8MapBody(v interface{}) bool {
	switch v := v.(type) {
	case map[uint8]interface{}:
		enc.writeUint8InterfaceMapBody(v)
	case map[uint8]string:
		enc.writeUint8StringMapBody(v)
	case map[uint8]int:
		enc.writeUint8IntMapBody(v)
	case map[uint8]int8:
		enc.writeUint8Int8MapBody(v)
	case map[uint8]int16:
		enc.writeUint8Int16MapBody(v)
	case map[uint8]int32:
		enc.writeUint8Int32MapBody(v)
	case map[uint8]int64:
		enc.writeUint8Int64MapBody(v)
	case map[uint8]uint:
		enc.writeUint8UintMapBody(v)
	case map[uint8]uint8:
		enc.writeUint8Uint8MapBody(v)
	case map[uint8]uint16:
		enc.writeUint8Uint16MapBody(v)
	case map[uint8]uint32:
		enc.writeUint8Uint32MapBody(v)
	case map[uint8]uint64:
		enc.writeUint8Uint64MapBody(v)
	case map[uint8]bool:
		enc.writeUint8BoolMapBody(v)
	case map[uint8]float32:
		enc.writeUint8Float32MapBody(v)
	case map[uint8]float64:
		enc.writeUint8Float64MapBody(v)
	default:
		return false
	}
	return true
}

//nolint
func (enc *Encoder) writeUint16MapBody(v interface{}) bool {
	switch v := v.(type) {
	case map[uint16]interface{}:
		enc.writeUint16InterfaceMapBody(v)
	case map[uint16]string:
		enc.writeUint16StringMapBody(v)
	case map[uint16]int:
		enc.writeUint16IntMapBody(v)
	case map[uint16]int8:
		enc.writeUint16Int8MapBody(v)
	case map[uint16]int16:
		enc.writeUint16Int16MapBody(v)
	case map[uint16]int32:
		enc.writeUint16Int32MapBody(v)
	case map[uint16]int64:
		enc.writeUint16Int64MapBody(v)
	case map[uint16]uint:
		enc.writeUint16UintMapBody(v)
	case map[uint16]uint8:
		enc.writeUint16Uint8MapBody(v)
	case map[uint16]uint16:
		enc.writeUint16Uint16MapBody(v)
	case map[uint16]uint32:
		enc.writeUint16Uint32MapBody(v)
	case map[uint16]uint64:
		enc.writeUint16Uint64MapBody(v)
	case map[uint16]bool:
		enc.writeUint16BoolMapBody(v)
	case map[uint16]float32:
		enc.writeUint16Float32MapBody(v)
	case map[uint16]float64:
		enc.writeUint16Float64MapBody(v)
	default:
		return false
	}
	return true
}

//nolint
func (enc *Encoder) writeUint32MapBody(v interface{}) bool {
	switch v := v.(type) {
	case map[uint32]interface{}:
		enc.writeUint32InterfaceMapBody(v)
	case map[uint32]string:
		enc.writeUint32StringMapBody(v)
	case map[uint32]int:
		enc.writeUint32IntMapBody(v)
	case map[uint32]int8:
		enc.writeUint32Int8MapBody(v)
	case map[uint32]int16:
		enc.writeUint32Int16MapBody(v)
	case map[uint32]int32:
		enc.writeUint32Int32MapBody(v)
	case map[uint32]int64:
		enc.writeUint32Int64MapBody(v)
	case map[uint32]uint:
		enc.writeUint32UintMapBody(v)
	case map[uint32]uint8:
		enc.writeUint32Uint8MapBody(v)
	case map[uint32]uint16:
		enc.writeUint32Uint16MapBody(v)
	case map[uint32]uint32:
		enc.writeUint32Uint32MapBody(v)
	case map[uint32]uint64:
		enc.writeUint32Uint64MapBody(v)
	case map[uint32]bool:
		enc.writeUint32BoolMapBody(v)
	case map[uint32]float32:
		enc.writeUint32Float32MapBody(v)
	case map[uint32]float64:
		enc.writeUint32Float64MapBody(v)
	default:
		return false
	}
	return true
}

//nolint
func (enc *Encoder) writeUint64MapBody(v interface{}) bool {
	switch v := v.(type) {
	case map[uint64]interface{}:
		enc.writeUint64InterfaceMapBody(v)
	case map[uint64]string:
		enc.writeUint64StringMapBody(v)
	case map[uint64]int:
		enc.writeUint64IntMapBody(v)
	case map[uint64]int8:
		enc.writeUint64Int8MapBody(v)
	case map[uint64]int16:
		enc.writeUint64Int16MapBody(v)
	case map[uint64]int32:
		enc.writeUint64Int32MapBody(v)
	case map[uint64]int64:
		enc.writeUint64Int64MapBody(v)
	case map[uint64]uint:
		enc.writeUint64UintMapBody(v)
	case map[uint64]uint8:
		enc.writeUint64Uint8MapBody(v)
	case map[uint64]uint16:
		enc.writeUint64Uint16MapBody(v)
	case map[uint64]uint32:
		enc.writeUint64Uint32MapBody(v)
	case map[uint64]uint64:
		enc.writeUint64Uint64MapBody(v)
	case map[uint64]bool:
		enc.writeUint64BoolMapBody(v)
	case map[uint64]float32:
		enc.writeUint64Float32MapBody(v)
	case map[uint64]float64:
		enc.writeUint64Float64MapBody(v)
	default:
		return false
	}
	return true
}

//nolint
func (enc *Encoder) writeFloat32MapBody(v interface{}) bool {
	switch v := v.(type) {
	case map[float32]interface{}:
		enc.writeFloat32InterfaceMapBody(v)
	case map[float32]string:
		enc.writeFloat32StringMapBody(v)
	case map[float32]int:
		enc.writeFloat32IntMapBody(v)
	case map[float32]int8:
		enc.writeFloat32Int8MapBody(v)
	case map[float32]int16:
		enc.writeFloat32Int16MapBody(v)
	case map[float32]int32:
		enc.writeFloat32Int32MapBody(v)
	case map[float32]int64:
		enc.writeFloat32Int64MapBody(v)
	case map[float32]uint:
		enc.writeFloat32UintMapBody(v)
	case map[float32]uint8:
		enc.writeFloat32Uint8MapBody(v)
	case map[float32]uint16:
		enc.writeFloat32Uint16MapBody(v)
	case map[float32]uint32:
		enc.writeFloat32Uint32MapBody(v)
	case map[float32]uint64:
		enc.writeFloat32Uint64MapBody(v)
	case map[float32]bool:
		enc.writeFloat32BoolMapBody(v)
	case map[float32]float32:
		enc.writeFloat32Float32MapBody(v)
	case map[float32]float64:
		enc.writeFloat32Float64MapBody(v)
	default:
		return false
	}
	return true
}

//nolint
func (enc *Encoder) writeFloat64MapBody(v interface{}) bool {
	switch v := v.(type) {
	case map[float64]interface{}:
		enc.writeFloat64InterfaceMapBody(v)
	case map[float64]string:
		enc.writeFloat64StringMapBody(v)
	case map[float64]int:
		enc.writeFloat64IntMapBody(v)
	case map[float64]int8:
		enc.writeFloat64Int8MapBody(v)
	case map[float64]int16:
		enc.writeFloat64Int16MapBody(v)
	case map[float64]int32:
		enc.writeFloat64Int32MapBody(v)
	case map[float64]int64:
		enc.writeFloat64Int64MapBody(v)
	case map[float64]uint:
		enc.writeFloat64UintMapBody(v)
	case map[float64]uint8:
		enc.writeFloat64Uint8MapBody(v)
	case map[float64]uint16:
		enc.writeFloat64Uint16MapBody(v)
	case map[float64]uint32:
		enc.writeFloat64Uint32MapBody(v)
	case map[float64]uint64:
		enc.writeFloat64Uint64MapBody(v)
	case map[float64]bool:
		enc.writeFloat64BoolMapBody(v)
	case map[float64]float32:
		enc.writeFloat64Float32MapBody(v)
	case map[float64]float64:
		enc.writeFloat64Float64MapBody(v)
	default:
		return false
	}
	return true
}

func (enc *Encoder) fastWriteMapBody(v interface{}) bool {
	switch reflect.TypeOf(v).Key().Kind() {
	case reflect.String:
		return enc.writeStringMapBody(v)
	case reflect.Interface:
		return enc.writeInterfaceMapBody(v)
	case reflect.Int:
		return enc.writeIntMapBody(v)
	case reflect.Int8:
		return enc.writeInt8MapBody(v)
	case reflect.Int16:
		return enc.writeInt16MapBody(v)
	case reflect.Int32:
		return enc.writeInt32MapBody(v)
	case reflect.Int64:
		return enc.writeInt64MapBody(v)
	case reflect.Uint:
		return enc.writeUintMapBody(v)
	case reflect.Uint8:
		return enc.writeUint8MapBody(v)
	case reflect.Uint16:
		return enc.writeUint16MapBody(v)
	case reflect.Uint32:
		return enc.writeUint32MapBody(v)
	case reflect.Uint64:
		return enc.writeUint64MapBody(v)
	case reflect.Float32:
		return enc.writeFloat32MapBody(v)
	case reflect.Float64:
		return enc.writeFloat64MapBody(v)
	}
	return false
}

func (enc *Encoder) writeMapBody(v interface{}) {
	if !enc.fastWriteMapBody(v) {
		enc.writeOtherMapBody(v)
	}
}

func (enc *Encoder) writeStringInterfaceMapBody(m map[string]interface{}) {
	for k, v := range m {
		enc.EncodeString(k)
		enc.encode(v)
	}
}

func (enc *Encoder) writeStringStringMapBody(m map[string]string) {
	for k, v := range m {
		enc.EncodeString(k)
		enc.EncodeString(v)
	}
}

func (enc *Encoder) writeStringIntMapBody(m map[string]int) {
	for k, v := range m {
		enc.EncodeString(k)
		enc.WriteInt(v)
	}
}

func (enc *Encoder) writeStringInt8MapBody(m map[string]int8) {
	for k, v := range m {
		enc.EncodeString(k)
		enc.WriteInt8(v)
	}
}

func (enc *Encoder) writeStringInt16MapBody(m map[string]int16) {
	for k, v := range m {
		enc.EncodeString(k)
		enc.WriteInt16(v)
	}
}

func (enc *Encoder) writeStringInt32MapBody(m map[string]int32) {
	for k, v := range m {
		enc.EncodeString(k)
		enc.WriteInt32(v)
	}
}

func (enc *Encoder) writeStringInt64MapBody(m map[string]int64) {
	for k, v := range m {
		enc.EncodeString(k)
		enc.WriteInt64(v)
	}
}

func (enc *Encoder) writeStringUintMapBody(m map[string]uint) {
	for k, v := range m {
		enc.EncodeString(k)
		enc.WriteUint(v)
	}
}

func (enc *Encoder) writeStringUint8MapBody(m map[string]uint8) {
	for k, v := range m {
		enc.EncodeString(k)
		enc.WriteUint8(v)
	}
}

func (enc *Encoder) writeStringUint16MapBody(m map[string]uint16) {
	for k, v := range m {
		enc.EncodeString(k)
		enc.WriteUint16(v)
	}
}

func (enc *Encoder) writeStringUint32MapBody(m map[string]uint32) {
	for k, v := range m {
		enc.EncodeString(k)
		enc.WriteUint32(v)
	}
}

func (enc *Encoder) writeStringUint64MapBody(m map[string]uint64) {
	for k, v := range m {
		enc.EncodeString(k)
		enc.WriteUint64(v)
	}
}

func (enc *Encoder) writeStringBoolMapBody(m map[string]bool) {
	for k, v := range m {
		enc.EncodeString(k)
		enc.WriteBool(v)
	}
}

func (enc *Encoder) writeStringFloat32MapBody(m map[string]float32) {
	for k, v := range m {
		enc.EncodeString(k)
		enc.WriteFloat32(v)
	}
}

func (enc *Encoder) writeStringFloat64MapBody(m map[string]float64) {
	for k, v := range m {
		enc.EncodeString(k)
		enc.WriteFloat64(v)
	}
}

func (enc *Encoder) writeInterfaceInterfaceMapBody(m map[interface{}]interface{}) {
	for k, v := range m {
		enc.encode(k)
		enc.encode(v)
	}
}

func (enc *Encoder) writeInterfaceStringMapBody(m map[interface{}]string) {
	for k, v := range m {
		enc.encode(k)
		enc.EncodeString(v)
	}
}

func (enc *Encoder) writeInterfaceIntMapBody(m map[interface{}]int) {
	for k, v := range m {
		enc.encode(k)
		enc.WriteInt(v)
	}
}

func (enc *Encoder) writeInterfaceInt8MapBody(m map[interface{}]int8) {
	for k, v := range m {
		enc.encode(k)
		enc.WriteInt8(v)
	}
}

func (enc *Encoder) writeInterfaceInt16MapBody(m map[interface{}]int16) {
	for k, v := range m {
		enc.encode(k)
		enc.WriteInt16(v)
	}
}

func (enc *Encoder) writeInterfaceInt32MapBody(m map[interface{}]int32) {
	for k, v := range m {
		enc.encode(k)
		enc.WriteInt32(v)
	}
}

func (enc *Encoder) writeInterfaceInt64MapBody(m map[interface{}]int64) {
	for k, v := range m {
		enc.encode(k)
		enc.WriteInt64(v)
	}
}

func (enc *Encoder) writeInterfaceUintMapBody(m map[interface{}]uint) {
	for k, v := range m {
		enc.encode(k)
		enc.WriteUint(v)
	}
}

func (enc *Encoder) writeInterfaceUint8MapBody(m map[interface{}]uint8) {
	for k, v := range m {
		enc.encode(k)
		enc.WriteUint8(v)
	}
}

func (enc *Encoder) writeInterfaceUint16MapBody(m map[interface{}]uint16) {
	for k, v := range m {
		enc.encode(k)
		enc.WriteUint16(v)
	}
}

func (enc *Encoder) writeInterfaceUint32MapBody(m map[interface{}]uint32) {
	for k, v := range m {
		enc.encode(k)
		enc.WriteUint32(v)
	}
}

func (enc *Encoder) writeInterfaceUint64MapBody(m map[interface{}]uint64) {
	for k, v := range m {
		enc.encode(k)
		enc.WriteUint64(v)
	}
}

func (enc *Encoder) writeInterfaceBoolMapBody(m map[interface{}]bool) {
	for k, v := range m {
		enc.encode(k)
		enc.WriteBool(v)
	}
}

func (enc *Encoder) writeInterfaceFloat32MapBody(m map[interface{}]float32) {
	for k, v := range m {
		enc.encode(k)
		enc.WriteFloat32(v)
	}
}

func (enc *Encoder) writeInterfaceFloat64MapBody(m map[interface{}]float64) {
	for k, v := range m {
		enc.encode(k)
		enc.WriteFloat64(v)
	}
}

func (enc *Encoder) writeIntInterfaceMapBody(m map[int]interface{}) {
	for k, v := range m {
		enc.WriteInt(k)
		enc.encode(v)
	}
}

func (enc *Encoder) writeIntStringMapBody(m map[int]string) {
	for k, v := range m {
		enc.WriteInt(k)
		enc.EncodeString(v)
	}
}

func (enc *Encoder) writeIntIntMapBody(m map[int]int) {
	for k, v := range m {
		enc.WriteInt(k)
		enc.WriteInt(v)
	}
}

func (enc *Encoder) writeIntInt8MapBody(m map[int]int8) {
	for k, v := range m {
		enc.WriteInt(k)
		enc.WriteInt8(v)
	}
}

func (enc *Encoder) writeIntInt16MapBody(m map[int]int16) {
	for k, v := range m {
		enc.WriteInt(k)
		enc.WriteInt16(v)
	}
}

func (enc *Encoder) writeIntInt32MapBody(m map[int]int32) {
	for k, v := range m {
		enc.WriteInt(k)
		enc.WriteInt32(v)
	}
}

func (enc *Encoder) writeIntInt64MapBody(m map[int]int64) {
	for k, v := range m {
		enc.WriteInt(k)
		enc.WriteInt64(v)
	}
}

func (enc *Encoder) writeIntUintMapBody(m map[int]uint) {
	for k, v := range m {
		enc.WriteInt(k)
		enc.WriteUint(v)
	}
}

func (enc *Encoder) writeIntUint8MapBody(m map[int]uint8) {
	for k, v := range m {
		enc.WriteInt(k)
		enc.WriteUint8(v)
	}
}

func (enc *Encoder) writeIntUint16MapBody(m map[int]uint16) {
	for k, v := range m {
		enc.WriteInt(k)
		enc.WriteUint16(v)
	}
}

func (enc *Encoder) writeIntUint32MapBody(m map[int]uint32) {
	for k, v := range m {
		enc.WriteInt(k)
		enc.WriteUint32(v)
	}
}

func (enc *Encoder) writeIntUint64MapBody(m map[int]uint64) {
	for k, v := range m {
		enc.WriteInt(k)
		enc.WriteUint64(v)
	}
}

func (enc *Encoder) writeIntBoolMapBody(m map[int]bool) {
	for k, v := range m {
		enc.WriteInt(k)
		enc.WriteBool(v)
	}
}

func (enc *Encoder) writeIntFloat32MapBody(m map[int]float32) {
	for k, v := range m {
		enc.WriteInt(k)
		enc.WriteFloat32(v)
	}
}

func (enc *Encoder) writeIntFloat64MapBody(m map[int]float64) {
	for k, v := range m {
		enc.WriteInt(k)
		enc.WriteFloat64(v)
	}
}

func (enc *Encoder) writeInt8InterfaceMapBody(m map[int8]interface{}) {
	for k, v := range m {
		enc.WriteInt8(k)
		enc.encode(v)
	}
}

func (enc *Encoder) writeInt8StringMapBody(m map[int8]string) {
	for k, v := range m {
		enc.WriteInt8(k)
		enc.EncodeString(v)
	}
}

func (enc *Encoder) writeInt8IntMapBody(m map[int8]int) {
	for k, v := range m {
		enc.WriteInt8(k)
		enc.WriteInt(v)
	}
}

func (enc *Encoder) writeInt8Int8MapBody(m map[int8]int8) {
	for k, v := range m {
		enc.WriteInt8(k)
		enc.WriteInt8(v)
	}
}

func (enc *Encoder) writeInt8Int16MapBody(m map[int8]int16) {
	for k, v := range m {
		enc.WriteInt8(k)
		enc.WriteInt16(v)
	}
}

func (enc *Encoder) writeInt8Int32MapBody(m map[int8]int32) {
	for k, v := range m {
		enc.WriteInt8(k)
		enc.WriteInt32(v)
	}
}

func (enc *Encoder) writeInt8Int64MapBody(m map[int8]int64) {
	for k, v := range m {
		enc.WriteInt8(k)
		enc.WriteInt64(v)
	}
}

func (enc *Encoder) writeInt8UintMapBody(m map[int8]uint) {
	for k, v := range m {
		enc.WriteInt8(k)
		enc.WriteUint(v)
	}
}

func (enc *Encoder) writeInt8Uint8MapBody(m map[int8]uint8) {
	for k, v := range m {
		enc.WriteInt8(k)
		enc.WriteUint8(v)
	}
}

func (enc *Encoder) writeInt8Uint16MapBody(m map[int8]uint16) {
	for k, v := range m {
		enc.WriteInt8(k)
		enc.WriteUint16(v)
	}
}

func (enc *Encoder) writeInt8Uint32MapBody(m map[int8]uint32) {
	for k, v := range m {
		enc.WriteInt8(k)
		enc.WriteUint32(v)
	}
}

func (enc *Encoder) writeInt8Uint64MapBody(m map[int8]uint64) {
	for k, v := range m {
		enc.WriteInt8(k)
		enc.WriteUint64(v)
	}
}

func (enc *Encoder) writeInt8BoolMapBody(m map[int8]bool) {
	for k, v := range m {
		enc.WriteInt8(k)
		enc.WriteBool(v)
	}
}

func (enc *Encoder) writeInt8Float32MapBody(m map[int8]float32) {
	for k, v := range m {
		enc.WriteInt8(k)
		enc.WriteFloat32(v)
	}
}

func (enc *Encoder) writeInt8Float64MapBody(m map[int8]float64) {
	for k, v := range m {
		enc.WriteInt8(k)
		enc.WriteFloat64(v)
	}
}

func (enc *Encoder) writeInt16InterfaceMapBody(m map[int16]interface{}) {
	for k, v := range m {
		enc.WriteInt16(k)
		enc.encode(v)
	}
}

func (enc *Encoder) writeInt16StringMapBody(m map[int16]string) {
	for k, v := range m {
		enc.WriteInt16(k)
		enc.EncodeString(v)
	}
}

func (enc *Encoder) writeInt16IntMapBody(m map[int16]int) {
	for k, v := range m {
		enc.WriteInt16(k)
		enc.WriteInt(v)
	}
}

func (enc *Encoder) writeInt16Int8MapBody(m map[int16]int8) {
	for k, v := range m {
		enc.WriteInt16(k)
		enc.WriteInt8(v)
	}
}

func (enc *Encoder) writeInt16Int16MapBody(m map[int16]int16) {
	for k, v := range m {
		enc.WriteInt16(k)
		enc.WriteInt16(v)
	}
}

func (enc *Encoder) writeInt16Int32MapBody(m map[int16]int32) {
	for k, v := range m {
		enc.WriteInt16(k)
		enc.WriteInt32(v)
	}
}

func (enc *Encoder) writeInt16Int64MapBody(m map[int16]int64) {
	for k, v := range m {
		enc.WriteInt16(k)
		enc.WriteInt64(v)
	}
}

func (enc *Encoder) writeInt16UintMapBody(m map[int16]uint) {
	for k, v := range m {
		enc.WriteInt16(k)
		enc.WriteUint(v)
	}
}

func (enc *Encoder) writeInt16Uint8MapBody(m map[int16]uint8) {
	for k, v := range m {
		enc.WriteInt16(k)
		enc.WriteUint8(v)
	}
}

func (enc *Encoder) writeInt16Uint16MapBody(m map[int16]uint16) {
	for k, v := range m {
		enc.WriteInt16(k)
		enc.WriteUint16(v)
	}
}

func (enc *Encoder) writeInt16Uint32MapBody(m map[int16]uint32) {
	for k, v := range m {
		enc.WriteInt16(k)
		enc.WriteUint32(v)
	}
}

func (enc *Encoder) writeInt16Uint64MapBody(m map[int16]uint64) {
	for k, v := range m {
		enc.WriteInt16(k)
		enc.WriteUint64(v)
	}
}

func (enc *Encoder) writeInt16BoolMapBody(m map[int16]bool) {
	for k, v := range m {
		enc.WriteInt16(k)
		enc.WriteBool(v)
	}
}

func (enc *Encoder) writeInt16Float32MapBody(m map[int16]float32) {
	for k, v := range m {
		enc.WriteInt16(k)
		enc.WriteFloat32(v)
	}
}

func (enc *Encoder) writeInt16Float64MapBody(m map[int16]float64) {
	for k, v := range m {
		enc.WriteInt16(k)
		enc.WriteFloat64(v)
	}
}

func (enc *Encoder) writeInt32InterfaceMapBody(m map[int32]interface{}) {
	for k, v := range m {
		enc.WriteInt32(k)
		enc.encode(v)
	}
}

func (enc *Encoder) writeInt32StringMapBody(m map[int32]string) {
	for k, v := range m {
		enc.WriteInt32(k)
		enc.EncodeString(v)
	}
}

func (enc *Encoder) writeInt32IntMapBody(m map[int32]int) {
	for k, v := range m {
		enc.WriteInt32(k)
		enc.WriteInt(v)
	}
}

func (enc *Encoder) writeInt32Int8MapBody(m map[int32]int8) {
	for k, v := range m {
		enc.WriteInt32(k)
		enc.WriteInt8(v)
	}
}

func (enc *Encoder) writeInt32Int16MapBody(m map[int32]int16) {
	for k, v := range m {
		enc.WriteInt32(k)
		enc.WriteInt16(v)
	}
}

func (enc *Encoder) writeInt32Int32MapBody(m map[int32]int32) {
	for k, v := range m {
		enc.WriteInt32(k)
		enc.WriteInt32(v)
	}
}

func (enc *Encoder) writeInt32Int64MapBody(m map[int32]int64) {
	for k, v := range m {
		enc.WriteInt32(k)
		enc.WriteInt64(v)
	}
}

func (enc *Encoder) writeInt32UintMapBody(m map[int32]uint) {
	for k, v := range m {
		enc.WriteInt32(k)
		enc.WriteUint(v)
	}
}

func (enc *Encoder) writeInt32Uint8MapBody(m map[int32]uint8) {
	for k, v := range m {
		enc.WriteInt32(k)
		enc.WriteUint8(v)
	}
}

func (enc *Encoder) writeInt32Uint16MapBody(m map[int32]uint16) {
	for k, v := range m {
		enc.WriteInt32(k)
		enc.WriteUint16(v)
	}
}

func (enc *Encoder) writeInt32Uint32MapBody(m map[int32]uint32) {
	for k, v := range m {
		enc.WriteInt32(k)
		enc.WriteUint32(v)
	}
}

func (enc *Encoder) writeInt32Uint64MapBody(m map[int32]uint64) {
	for k, v := range m {
		enc.WriteInt32(k)
		enc.WriteUint64(v)
	}
}

func (enc *Encoder) writeInt32BoolMapBody(m map[int32]bool) {
	for k, v := range m {
		enc.WriteInt32(k)
		enc.WriteBool(v)
	}
}

func (enc *Encoder) writeInt32Float32MapBody(m map[int32]float32) {
	for k, v := range m {
		enc.WriteInt32(k)
		enc.WriteFloat32(v)
	}
}

func (enc *Encoder) writeInt32Float64MapBody(m map[int32]float64) {
	for k, v := range m {
		enc.WriteInt32(k)
		enc.WriteFloat64(v)
	}
}

func (enc *Encoder) writeInt64InterfaceMapBody(m map[int64]interface{}) {
	for k, v := range m {
		enc.WriteInt64(k)
		enc.encode(v)
	}
}

func (enc *Encoder) writeInt64StringMapBody(m map[int64]string) {
	for k, v := range m {
		enc.WriteInt64(k)
		enc.EncodeString(v)
	}
}

func (enc *Encoder) writeInt64IntMapBody(m map[int64]int) {
	for k, v := range m {
		enc.WriteInt64(k)
		enc.WriteInt(v)
	}
}

func (enc *Encoder) writeInt64Int8MapBody(m map[int64]int8) {
	for k, v := range m {
		enc.WriteInt64(k)
		enc.WriteInt8(v)
	}
}

func (enc *Encoder) writeInt64Int16MapBody(m map[int64]int16) {
	for k, v := range m {
		enc.WriteInt64(k)
		enc.WriteInt16(v)
	}
}

func (enc *Encoder) writeInt64Int32MapBody(m map[int64]int32) {
	for k, v := range m {
		enc.WriteInt64(k)
		enc.WriteInt32(v)
	}
}

func (enc *Encoder) writeInt64Int64MapBody(m map[int64]int64) {
	for k, v := range m {
		enc.WriteInt64(k)
		enc.WriteInt64(v)
	}
}

func (enc *Encoder) writeInt64UintMapBody(m map[int64]uint) {
	for k, v := range m {
		enc.WriteInt64(k)
		enc.WriteUint(v)
	}
}

func (enc *Encoder) writeInt64Uint8MapBody(m map[int64]uint8) {
	for k, v := range m {
		enc.WriteInt64(k)
		enc.WriteUint8(v)
	}
}

func (enc *Encoder) writeInt64Uint16MapBody(m map[int64]uint16) {
	for k, v := range m {
		enc.WriteInt64(k)
		enc.WriteUint16(v)
	}
}

func (enc *Encoder) writeInt64Uint32MapBody(m map[int64]uint32) {
	for k, v := range m {
		enc.WriteInt64(k)
		enc.WriteUint32(v)
	}
}

func (enc *Encoder) writeInt64Uint64MapBody(m map[int64]uint64) {
	for k, v := range m {
		enc.WriteInt64(k)
		enc.WriteUint64(v)
	}
}

func (enc *Encoder) writeInt64BoolMapBody(m map[int64]bool) {
	for k, v := range m {
		enc.WriteInt64(k)
		enc.WriteBool(v)
	}
}

func (enc *Encoder) writeInt64Float32MapBody(m map[int64]float32) {
	for k, v := range m {
		enc.WriteInt64(k)
		enc.WriteFloat32(v)
	}
}

func (enc *Encoder) writeInt64Float64MapBody(m map[int64]float64) {
	for k, v := range m {
		enc.WriteInt64(k)
		enc.WriteFloat64(v)
	}
}

func (enc *Encoder) writeUintInterfaceMapBody(m map[uint]interface{}) {
	for k, v := range m {
		enc.WriteUint(k)
		enc.encode(v)
	}
}

func (enc *Encoder) writeUintStringMapBody(m map[uint]string) {
	for k, v := range m {
		enc.WriteUint(k)
		enc.EncodeString(v)
	}
}

func (enc *Encoder) writeUintIntMapBody(m map[uint]int) {
	for k, v := range m {
		enc.WriteUint(k)
		enc.WriteInt(v)
	}
}

func (enc *Encoder) writeUintInt8MapBody(m map[uint]int8) {
	for k, v := range m {
		enc.WriteUint(k)
		enc.WriteInt8(v)
	}
}

func (enc *Encoder) writeUintInt16MapBody(m map[uint]int16) {
	for k, v := range m {
		enc.WriteUint(k)
		enc.WriteInt16(v)
	}
}

func (enc *Encoder) writeUintInt32MapBody(m map[uint]int32) {
	for k, v := range m {
		enc.WriteUint(k)
		enc.WriteInt32(v)
	}
}

func (enc *Encoder) writeUintInt64MapBody(m map[uint]int64) {
	for k, v := range m {
		enc.WriteUint(k)
		enc.WriteInt64(v)
	}
}

func (enc *Encoder) writeUintUintMapBody(m map[uint]uint) {
	for k, v := range m {
		enc.WriteUint(k)
		enc.WriteUint(v)
	}
}

func (enc *Encoder) writeUintUint8MapBody(m map[uint]uint8) {
	for k, v := range m {
		enc.WriteUint(k)
		enc.WriteUint8(v)
	}
}

func (enc *Encoder) writeUintUint16MapBody(m map[uint]uint16) {
	for k, v := range m {
		enc.WriteUint(k)
		enc.WriteUint16(v)
	}
}

func (enc *Encoder) writeUintUint32MapBody(m map[uint]uint32) {
	for k, v := range m {
		enc.WriteUint(k)
		enc.WriteUint32(v)
	}
}

func (enc *Encoder) writeUintUint64MapBody(m map[uint]uint64) {
	for k, v := range m {
		enc.WriteUint(k)
		enc.WriteUint64(v)
	}
}

func (enc *Encoder) writeUintBoolMapBody(m map[uint]bool) {
	for k, v := range m {
		enc.WriteUint(k)
		enc.WriteBool(v)
	}
}

func (enc *Encoder) writeUintFloat32MapBody(m map[uint]float32) {
	for k, v := range m {
		enc.WriteUint(k)
		enc.WriteFloat32(v)
	}
}

func (enc *Encoder) writeUintFloat64MapBody(m map[uint]float64) {
	for k, v := range m {
		enc.WriteUint(k)
		enc.WriteFloat64(v)
	}
}

func (enc *Encoder) writeUint8InterfaceMapBody(m map[uint8]interface{}) {
	for k, v := range m {
		enc.WriteUint8(k)
		enc.encode(v)
	}
}

func (enc *Encoder) writeUint8StringMapBody(m map[uint8]string) {
	for k, v := range m {
		enc.WriteUint8(k)
		enc.EncodeString(v)
	}
}

func (enc *Encoder) writeUint8IntMapBody(m map[uint8]int) {
	for k, v := range m {
		enc.WriteUint8(k)
		enc.WriteInt(v)
	}
}

func (enc *Encoder) writeUint8Int8MapBody(m map[uint8]int8) {
	for k, v := range m {
		enc.WriteUint8(k)
		enc.WriteInt8(v)
	}
}

func (enc *Encoder) writeUint8Int16MapBody(m map[uint8]int16) {
	for k, v := range m {
		enc.WriteUint8(k)
		enc.WriteInt16(v)
	}
}

func (enc *Encoder) writeUint8Int32MapBody(m map[uint8]int32) {
	for k, v := range m {
		enc.WriteUint8(k)
		enc.WriteInt32(v)
	}
}

func (enc *Encoder) writeUint8Int64MapBody(m map[uint8]int64) {
	for k, v := range m {
		enc.WriteUint8(k)
		enc.WriteInt64(v)
	}
}

func (enc *Encoder) writeUint8UintMapBody(m map[uint8]uint) {
	for k, v := range m {
		enc.WriteUint8(k)
		enc.WriteUint(v)
	}
}

func (enc *Encoder) writeUint8Uint8MapBody(m map[uint8]uint8) {
	for k, v := range m {
		enc.WriteUint8(k)
		enc.WriteUint8(v)
	}
}

func (enc *Encoder) writeUint8Uint16MapBody(m map[uint8]uint16) {
	for k, v := range m {
		enc.WriteUint8(k)
		enc.WriteUint16(v)
	}
}

func (enc *Encoder) writeUint8Uint32MapBody(m map[uint8]uint32) {
	for k, v := range m {
		enc.WriteUint8(k)
		enc.WriteUint32(v)
	}
}

func (enc *Encoder) writeUint8Uint64MapBody(m map[uint8]uint64) {
	for k, v := range m {
		enc.WriteUint8(k)
		enc.WriteUint64(v)
	}
}

func (enc *Encoder) writeUint8BoolMapBody(m map[uint8]bool) {
	for k, v := range m {
		enc.WriteUint8(k)
		enc.WriteBool(v)
	}
}

func (enc *Encoder) writeUint8Float32MapBody(m map[uint8]float32) {
	for k, v := range m {
		enc.WriteUint8(k)
		enc.WriteFloat32(v)
	}
}

func (enc *Encoder) writeUint8Float64MapBody(m map[uint8]float64) {
	for k, v := range m {
		enc.WriteUint8(k)
		enc.WriteFloat64(v)
	}
}

func (enc *Encoder) writeUint16InterfaceMapBody(m map[uint16]interface{}) {
	for k, v := range m {
		enc.WriteUint16(k)
		enc.encode(v)
	}
}

func (enc *Encoder) writeUint16StringMapBody(m map[uint16]string) {
	for k, v := range m {
		enc.WriteUint16(k)
		enc.EncodeString(v)
	}
}

func (enc *Encoder) writeUint16IntMapBody(m map[uint16]int) {
	for k, v := range m {
		enc.WriteUint16(k)
		enc.WriteInt(v)
	}
}

func (enc *Encoder) writeUint16Int8MapBody(m map[uint16]int8) {
	for k, v := range m {
		enc.WriteUint16(k)
		enc.WriteInt8(v)
	}
}

func (enc *Encoder) writeUint16Int16MapBody(m map[uint16]int16) {
	for k, v := range m {
		enc.WriteUint16(k)
		enc.WriteInt16(v)
	}
}

func (enc *Encoder) writeUint16Int32MapBody(m map[uint16]int32) {
	for k, v := range m {
		enc.WriteUint16(k)
		enc.WriteInt32(v)
	}
}

func (enc *Encoder) writeUint16Int64MapBody(m map[uint16]int64) {
	for k, v := range m {
		enc.WriteUint16(k)
		enc.WriteInt64(v)
	}
}

func (enc *Encoder) writeUint16UintMapBody(m map[uint16]uint) {
	for k, v := range m {
		enc.WriteUint16(k)
		enc.WriteUint(v)
	}
}

func (enc *Encoder) writeUint16Uint8MapBody(m map[uint16]uint8) {
	for k, v := range m {
		enc.WriteUint16(k)
		enc.WriteUint8(v)
	}
}

func (enc *Encoder) writeUint16Uint16MapBody(m map[uint16]uint16) {
	for k, v := range m {
		enc.WriteUint16(k)
		enc.WriteUint16(v)
	}
}

func (enc *Encoder) writeUint16Uint32MapBody(m map[uint16]uint32) {
	for k, v := range m {
		enc.WriteUint16(k)
		enc.WriteUint32(v)
	}
}

func (enc *Encoder) writeUint16Uint64MapBody(m map[uint16]uint64) {
	for k, v := range m {
		enc.WriteUint16(k)
		enc.WriteUint64(v)
	}
}

func (enc *Encoder) writeUint16BoolMapBody(m map[uint16]bool) {
	for k, v := range m {
		enc.WriteUint16(k)
		enc.WriteBool(v)
	}
}

func (enc *Encoder) writeUint16Float32MapBody(m map[uint16]float32) {
	for k, v := range m {
		enc.WriteUint16(k)
		enc.WriteFloat32(v)
	}
}

func (enc *Encoder) writeUint16Float64MapBody(m map[uint16]float64) {
	for k, v := range m {
		enc.WriteUint16(k)
		enc.WriteFloat64(v)
	}
}

func (enc *Encoder) writeUint32InterfaceMapBody(m map[uint32]interface{}) {
	for k, v := range m {
		enc.WriteUint32(k)
		enc.encode(v)
	}
}

func (enc *Encoder) writeUint32StringMapBody(m map[uint32]string) {
	for k, v := range m {
		enc.WriteUint32(k)
		enc.EncodeString(v)
	}
}

func (enc *Encoder) writeUint32IntMapBody(m map[uint32]int) {
	for k, v := range m {
		enc.WriteUint32(k)
		enc.WriteInt(v)
	}
}

func (enc *Encoder) writeUint32Int8MapBody(m map[uint32]int8) {
	for k, v := range m {
		enc.WriteUint32(k)
		enc.WriteInt8(v)
	}
}

func (enc *Encoder) writeUint32Int16MapBody(m map[uint32]int16) {
	for k, v := range m {
		enc.WriteUint32(k)
		enc.WriteInt16(v)
	}
}

func (enc *Encoder) writeUint32Int32MapBody(m map[uint32]int32) {
	for k, v := range m {
		enc.WriteUint32(k)
		enc.WriteInt32(v)
	}
}

func (enc *Encoder) writeUint32Int64MapBody(m map[uint32]int64) {
	for k, v := range m {
		enc.WriteUint32(k)
		enc.WriteInt64(v)
	}
}

func (enc *Encoder) writeUint32UintMapBody(m map[uint32]uint) {
	for k, v := range m {
		enc.WriteUint32(k)
		enc.WriteUint(v)
	}
}

func (enc *Encoder) writeUint32Uint8MapBody(m map[uint32]uint8) {
	for k, v := range m {
		enc.WriteUint32(k)
		enc.WriteUint8(v)
	}
}

func (enc *Encoder) writeUint32Uint16MapBody(m map[uint32]uint16) {
	for k, v := range m {
		enc.WriteUint32(k)
		enc.WriteUint16(v)
	}
}

func (enc *Encoder) writeUint32Uint32MapBody(m map[uint32]uint32) {
	for k, v := range m {
		enc.WriteUint32(k)
		enc.WriteUint32(v)
	}
}

func (enc *Encoder) writeUint32Uint64MapBody(m map[uint32]uint64) {
	for k, v := range m {
		enc.WriteUint32(k)
		enc.WriteUint64(v)
	}
}

func (enc *Encoder) writeUint32BoolMapBody(m map[uint32]bool) {
	for k, v := range m {
		enc.WriteUint32(k)
		enc.WriteBool(v)
	}
}

func (enc *Encoder) writeUint32Float32MapBody(m map[uint32]float32) {
	for k, v := range m {
		enc.WriteUint32(k)
		enc.WriteFloat32(v)
	}
}

func (enc *Encoder) writeUint32Float64MapBody(m map[uint32]float64) {
	for k, v := range m {
		enc.WriteUint32(k)
		enc.WriteFloat64(v)
	}
}

func (enc *Encoder) writeUint64InterfaceMapBody(m map[uint64]interface{}) {
	for k, v := range m {
		enc.WriteUint64(k)
		enc.encode(v)
	}
}

func (enc *Encoder) writeUint64StringMapBody(m map[uint64]string) {
	for k, v := range m {
		enc.WriteUint64(k)
		enc.EncodeString(v)
	}
}

func (enc *Encoder) writeUint64IntMapBody(m map[uint64]int) {
	for k, v := range m {
		enc.WriteUint64(k)
		enc.WriteInt(v)
	}
}

func (enc *Encoder) writeUint64Int8MapBody(m map[uint64]int8) {
	for k, v := range m {
		enc.WriteUint64(k)
		enc.WriteInt8(v)
	}
}

func (enc *Encoder) writeUint64Int16MapBody(m map[uint64]int16) {
	for k, v := range m {
		enc.WriteUint64(k)
		enc.WriteInt16(v)
	}
}

func (enc *Encoder) writeUint64Int32MapBody(m map[uint64]int32) {
	for k, v := range m {
		enc.WriteUint64(k)
		enc.WriteInt32(v)
	}
}

func (enc *Encoder) writeUint64Int64MapBody(m map[uint64]int64) {
	for k, v := range m {
		enc.WriteUint64(k)
		enc.WriteInt64(v)
	}
}

func (enc *Encoder) writeUint64UintMapBody(m map[uint64]uint) {
	for k, v := range m {
		enc.WriteUint64(k)
		enc.WriteUint(v)
	}
}

func (enc *Encoder) writeUint64Uint8MapBody(m map[uint64]uint8) {
	for k, v := range m {
		enc.WriteUint64(k)
		enc.WriteUint8(v)
	}
}

func (enc *Encoder) writeUint64Uint16MapBody(m map[uint64]uint16) {
	for k, v := range m {
		enc.WriteUint64(k)
		enc.WriteUint16(v)
	}
}

func (enc *Encoder) writeUint64Uint32MapBody(m map[uint64]uint32) {
	for k, v := range m {
		enc.WriteUint64(k)
		enc.WriteUint32(v)
	}
}

func (enc *Encoder) writeUint64Uint64MapBody(m map[uint64]uint64) {
	for k, v := range m {
		enc.WriteUint64(k)
		enc.WriteUint64(v)
	}
}

func (enc *Encoder) writeUint64BoolMapBody(m map[uint64]bool) {
	for k, v := range m {
		enc.WriteUint64(k)
		enc.WriteBool(v)
	}
}

func (enc *Encoder) writeUint64Float32MapBody(m map[uint64]float32) {
	for k, v := range m {
		enc.WriteUint64(k)
		enc.WriteFloat32(v)
	}
}

func (enc *Encoder) writeUint64Float64MapBody(m map[uint64]float64) {
	for k, v := range m {
		enc.WriteUint64(k)
		enc.WriteFloat64(v)
	}
}

func (enc *Encoder) writeFloat32InterfaceMapBody(m map[float32]interface{}) {
	for k, v := range m {
		enc.WriteFloat32(k)
		enc.encode(v)
	}
}

func (enc *Encoder) writeFloat32StringMapBody(m map[float32]string) {
	for k, v := range m {
		enc.WriteFloat32(k)
		enc.EncodeString(v)
	}
}

func (enc *Encoder) writeFloat32IntMapBody(m map[float32]int) {
	for k, v := range m {
		enc.WriteFloat32(k)
		enc.WriteInt(v)
	}
}

func (enc *Encoder) writeFloat32Int8MapBody(m map[float32]int8) {
	for k, v := range m {
		enc.WriteFloat32(k)
		enc.WriteInt8(v)
	}
}

func (enc *Encoder) writeFloat32Int16MapBody(m map[float32]int16) {
	for k, v := range m {
		enc.WriteFloat32(k)
		enc.WriteInt16(v)
	}
}

func (enc *Encoder) writeFloat32Int32MapBody(m map[float32]int32) {
	for k, v := range m {
		enc.WriteFloat32(k)
		enc.WriteInt32(v)
	}
}

func (enc *Encoder) writeFloat32Int64MapBody(m map[float32]int64) {
	for k, v := range m {
		enc.WriteFloat32(k)
		enc.WriteInt64(v)
	}
}

func (enc *Encoder) writeFloat32UintMapBody(m map[float32]uint) {
	for k, v := range m {
		enc.WriteFloat32(k)
		enc.WriteUint(v)
	}
}

func (enc *Encoder) writeFloat32Uint8MapBody(m map[float32]uint8) {
	for k, v := range m {
		enc.WriteFloat32(k)
		enc.WriteUint8(v)
	}
}

func (enc *Encoder) writeFloat32Uint16MapBody(m map[float32]uint16) {
	for k, v := range m {
		enc.WriteFloat32(k)
		enc.WriteUint16(v)
	}
}

func (enc *Encoder) writeFloat32Uint32MapBody(m map[float32]uint32) {
	for k, v := range m {
		enc.WriteFloat32(k)
		enc.WriteUint32(v)
	}
}

func (enc *Encoder) writeFloat32Uint64MapBody(m map[float32]uint64) {
	for k, v := range m {
		enc.WriteFloat32(k)
		enc.WriteUint64(v)
	}
}

func (enc *Encoder) writeFloat32BoolMapBody(m map[float32]bool) {
	for k, v := range m {
		enc.WriteFloat32(k)
		enc.WriteBool(v)
	}
}

func (enc *Encoder) writeFloat32Float32MapBody(m map[float32]float32) {
	for k, v := range m {
		enc.WriteFloat32(k)
		enc.WriteFloat32(v)
	}
}

func (enc *Encoder) writeFloat32Float64MapBody(m map[float32]float64) {
	for k, v := range m {
		enc.WriteFloat32(k)
		enc.WriteFloat64(v)
	}
}

func (enc *Encoder) writeFloat64InterfaceMapBody(m map[float64]interface{}) {
	for k, v := range m {
		enc.WriteFloat64(k)
		enc.encode(v)
	}
}

func (enc *Encoder) writeFloat64StringMapBody(m map[float64]string) {
	for k, v := range m {
		enc.WriteFloat64(k)
		enc.EncodeString(v)
	}
}

func (enc *Encoder) writeFloat64IntMapBody(m map[float64]int) {
	for k, v := range m {
		enc.WriteFloat64(k)
		enc.WriteInt(v)
	}
}

func (enc *Encoder) writeFloat64Int8MapBody(m map[float64]int8) {
	for k, v := range m {
		enc.WriteFloat64(k)
		enc.WriteInt8(v)
	}
}

func (enc *Encoder) writeFloat64Int16MapBody(m map[float64]int16) {
	for k, v := range m {
		enc.WriteFloat64(k)
		enc.WriteInt16(v)
	}
}

func (enc *Encoder) writeFloat64Int32MapBody(m map[float64]int32) {
	for k, v := range m {
		enc.WriteFloat64(k)
		enc.WriteInt32(v)
	}
}

func (enc *Encoder) writeFloat64Int64MapBody(m map[float64]int64) {
	for k, v := range m {
		enc.WriteFloat64(k)
		enc.WriteInt64(v)
	}
}

func (enc *Encoder) writeFloat64UintMapBody(m map[float64]uint) {
	for k, v := range m {
		enc.WriteFloat64(k)
		enc.WriteUint(v)
	}
}

func (enc *Encoder) writeFloat64Uint8MapBody(m map[float64]uint8) {
	for k, v := range m {
		enc.WriteFloat64(k)
		enc.WriteUint8(v)
	}
}

func (enc *Encoder) writeFloat64Uint16MapBody(m map[float64]uint16) {
	for k, v := range m {
		enc.WriteFloat64(k)
		enc.WriteUint16(v)
	}
}

func (enc *Encoder) writeFloat64Uint32MapBody(m map[float64]uint32) {
	for k, v := range m {
		enc.WriteFloat64(k)
		enc.WriteUint32(v)
	}
}

func (enc *Encoder) writeFloat64Uint64MapBody(m map[float64]uint64) {
	for k, v := range m {
		enc.WriteFloat64(k)
		enc.WriteUint64(v)
	}
}

func (enc *Encoder) writeFloat64BoolMapBody(m map[float64]bool) {
	for k, v := range m {
		enc.WriteFloat64(k)
		enc.WriteBool(v)
	}
}

func (enc *Encoder) writeFloat64Float32MapBody(m map[float64]float32) {
	for k, v := range m {
		enc.WriteFloat64(k)
		enc.WriteFloat32(v)
	}
}

func (enc *Encoder) writeFloat64Float64MapBody(m map[float64]float64) {
	for k, v := range m {
		enc.WriteFloat64(k)
		enc.WriteFloat64(v)
	}
}

func (enc *Encoder) writeOtherMapBody(v interface{}) {
	mapType := reflect2.TypeOf(v).(*reflect2.UnsafeMapType)
	p := reflect2.PtrOf(v)
	iter := mapType.UnsafeIterate(unsafe.Pointer(&p))
	kt := mapType.Key()
	vt := mapType.Elem()
	for iter.HasNext() {
		kp, vp := iter.UnsafeNext()
		enc.encode(kt.UnsafeIndirect(kp))
		enc.encode(vt.UnsafeIndirect(vp))
	}
}
