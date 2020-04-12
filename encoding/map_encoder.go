/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/map_encoder.go                                  |
|                                                          |
| LastModified: Apr 12, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

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
	enc.SetPtrReference(v)
	writeMap(enc, reflect.ValueOf(v).Elem().Interface())
}

// WriteMap to encoder
func WriteMap(enc *Encoder, v interface{}) {
	enc.AddReferenceCount(1)
	writeMap(enc, v)
}

func writeMap(enc *Encoder, v interface{}) {
	count := reflect.ValueOf(v).Len()
	if count == 0 {
		enc.buf = append(enc.buf, TagMap, TagOpenbrace, TagClosebrace)
		return
	}
	enc.WriteHead(count, TagMap)
	writeMapBody(enc, v)
	enc.WriteFoot()
}

func writeMapBody(enc *Encoder, v interface{}) {
	switch v := v.(type) {
	case map[string]string:
		writeStringStringMapBody(enc, v)
	case map[string]int:
		writeStringIntMapBody(enc, v)
	case map[string]int8:
		writeStringInt8MapBody(enc, v)
	case map[string]int16:
		writeStringInt16MapBody(enc, v)
	case map[string]int32:
		writeStringInt32MapBody(enc, v)
	case map[string]int64:
		writeStringInt64MapBody(enc, v)
	case map[string]uint:
		writeStringUintMapBody(enc, v)
	case map[string]uint8:
		writeStringUint8MapBody(enc, v)
	case map[string]uint16:
		writeStringUint16MapBody(enc, v)
	case map[string]uint32:
		writeStringUint32MapBody(enc, v)
	case map[string]uint64:
		writeStringUint64MapBody(enc, v)
	case map[string]bool:
		writeStringBoolMapBody(enc, v)
	case map[string]float32:
		writeStringFloat32MapBody(enc, v)
	case map[string]float64:
		writeStringFloat64MapBody(enc, v)
	case map[string]interface{}:
		writeStringInterfaceMapBody(enc, v)
	case map[int]string:
		writeIntStringMapBody(enc, v)
	case map[int]int:
		writeIntIntMapBody(enc, v)
	case map[int]int8:
		writeIntInt8MapBody(enc, v)
	case map[int]int16:
		writeIntInt16MapBody(enc, v)
	case map[int]int32:
		writeIntInt32MapBody(enc, v)
	case map[int]int64:
		writeIntInt64MapBody(enc, v)
	case map[int]uint:
		writeIntUintMapBody(enc, v)
	case map[int]uint8:
		writeIntUint8MapBody(enc, v)
	case map[int]uint16:
		writeIntUint16MapBody(enc, v)
	case map[int]uint32:
		writeIntUint32MapBody(enc, v)
	case map[int]uint64:
		writeIntUint64MapBody(enc, v)
	case map[int]bool:
		writeIntBoolMapBody(enc, v)
	case map[int]float32:
		writeIntFloat32MapBody(enc, v)
	case map[int]float64:
		writeIntFloat64MapBody(enc, v)
	case map[int]interface{}:
		writeIntInterfaceMapBody(enc, v)
	case map[int8]string:
		writeInt8StringMapBody(enc, v)
	case map[int8]int:
		writeInt8IntMapBody(enc, v)
	case map[int8]int8:
		writeInt8Int8MapBody(enc, v)
	case map[int8]int16:
		writeInt8Int16MapBody(enc, v)
	case map[int8]int32:
		writeInt8Int32MapBody(enc, v)
	case map[int8]int64:
		writeInt8Int64MapBody(enc, v)
	case map[int8]uint:
		writeInt8UintMapBody(enc, v)
	case map[int8]uint8:
		writeInt8Uint8MapBody(enc, v)
	case map[int8]uint16:
		writeInt8Uint16MapBody(enc, v)
	case map[int8]uint32:
		writeInt8Uint32MapBody(enc, v)
	case map[int8]uint64:
		writeInt8Uint64MapBody(enc, v)
	case map[int8]bool:
		writeInt8BoolMapBody(enc, v)
	case map[int8]float32:
		writeInt8Float32MapBody(enc, v)
	case map[int8]float64:
		writeInt8Float64MapBody(enc, v)
	case map[int8]interface{}:
		writeInt8InterfaceMapBody(enc, v)
	case map[int16]string:
		writeInt16StringMapBody(enc, v)
	case map[int16]int:
		writeInt16IntMapBody(enc, v)
	case map[int16]int8:
		writeInt16Int8MapBody(enc, v)
	case map[int16]int16:
		writeInt16Int16MapBody(enc, v)
	case map[int16]int32:
		writeInt16Int32MapBody(enc, v)
	case map[int16]int64:
		writeInt16Int64MapBody(enc, v)
	case map[int16]uint:
		writeInt16UintMapBody(enc, v)
	case map[int16]uint8:
		writeInt16Uint8MapBody(enc, v)
	case map[int16]uint16:
		writeInt16Uint16MapBody(enc, v)
	case map[int16]uint32:
		writeInt16Uint32MapBody(enc, v)
	case map[int16]uint64:
		writeInt16Uint64MapBody(enc, v)
	case map[int16]bool:
		writeInt16BoolMapBody(enc, v)
	case map[int16]float32:
		writeInt16Float32MapBody(enc, v)
	case map[int16]float64:
		writeInt16Float64MapBody(enc, v)
	case map[int16]interface{}:
		writeInt16InterfaceMapBody(enc, v)
	case map[int32]string:
		writeInt32StringMapBody(enc, v)
	case map[int32]int:
		writeInt32IntMapBody(enc, v)
	case map[int32]int8:
		writeInt32Int8MapBody(enc, v)
	case map[int32]int16:
		writeInt32Int16MapBody(enc, v)
	case map[int32]int32:
		writeInt32Int32MapBody(enc, v)
	case map[int32]int64:
		writeInt32Int64MapBody(enc, v)
	case map[int32]uint:
		writeInt32UintMapBody(enc, v)
	case map[int32]uint8:
		writeInt32Uint8MapBody(enc, v)
	case map[int32]uint16:
		writeInt32Uint16MapBody(enc, v)
	case map[int32]uint32:
		writeInt32Uint32MapBody(enc, v)
	case map[int32]uint64:
		writeInt32Uint64MapBody(enc, v)
	case map[int32]bool:
		writeInt32BoolMapBody(enc, v)
	case map[int32]float32:
		writeInt32Float32MapBody(enc, v)
	case map[int32]float64:
		writeInt32Float64MapBody(enc, v)
	case map[int32]interface{}:
		writeInt32InterfaceMapBody(enc, v)
	case map[int64]string:
		writeInt64StringMapBody(enc, v)
	case map[int64]int:
		writeInt64IntMapBody(enc, v)
	case map[int64]int8:
		writeInt64Int8MapBody(enc, v)
	case map[int64]int16:
		writeInt64Int16MapBody(enc, v)
	case map[int64]int32:
		writeInt64Int32MapBody(enc, v)
	case map[int64]int64:
		writeInt64Int64MapBody(enc, v)
	case map[int64]uint:
		writeInt64UintMapBody(enc, v)
	case map[int64]uint8:
		writeInt64Uint8MapBody(enc, v)
	case map[int64]uint16:
		writeInt64Uint16MapBody(enc, v)
	case map[int64]uint32:
		writeInt64Uint32MapBody(enc, v)
	case map[int64]uint64:
		writeInt64Uint64MapBody(enc, v)
	case map[int64]bool:
		writeInt64BoolMapBody(enc, v)
	case map[int64]float32:
		writeInt64Float32MapBody(enc, v)
	case map[int64]float64:
		writeInt64Float64MapBody(enc, v)
	case map[int64]interface{}:
		writeInt64InterfaceMapBody(enc, v)
	case map[uint]string:
		writeUintStringMapBody(enc, v)
	case map[uint]int:
		writeUintIntMapBody(enc, v)
	case map[uint]int8:
		writeUintInt8MapBody(enc, v)
	case map[uint]int16:
		writeUintInt16MapBody(enc, v)
	case map[uint]int32:
		writeUintInt32MapBody(enc, v)
	case map[uint]int64:
		writeUintInt64MapBody(enc, v)
	case map[uint]uint:
		writeUintUintMapBody(enc, v)
	case map[uint]uint8:
		writeUintUint8MapBody(enc, v)
	case map[uint]uint16:
		writeUintUint16MapBody(enc, v)
	case map[uint]uint32:
		writeUintUint32MapBody(enc, v)
	case map[uint]uint64:
		writeUintUint64MapBody(enc, v)
	case map[uint]bool:
		writeUintBoolMapBody(enc, v)
	case map[uint]float32:
		writeUintFloat32MapBody(enc, v)
	case map[uint]float64:
		writeUintFloat64MapBody(enc, v)
	case map[uint]interface{}:
		writeUintInterfaceMapBody(enc, v)
	case map[uint8]string:
		writeUint8StringMapBody(enc, v)
	case map[uint8]int:
		writeUint8IntMapBody(enc, v)
	case map[uint8]int8:
		writeUint8Int8MapBody(enc, v)
	case map[uint8]int16:
		writeUint8Int16MapBody(enc, v)
	case map[uint8]int32:
		writeUint8Int32MapBody(enc, v)
	case map[uint8]int64:
		writeUint8Int64MapBody(enc, v)
	case map[uint8]uint:
		writeUint8UintMapBody(enc, v)
	case map[uint8]uint8:
		writeUint8Uint8MapBody(enc, v)
	case map[uint8]uint16:
		writeUint8Uint16MapBody(enc, v)
	case map[uint8]uint32:
		writeUint8Uint32MapBody(enc, v)
	case map[uint8]uint64:
		writeUint8Uint64MapBody(enc, v)
	case map[uint8]bool:
		writeUint8BoolMapBody(enc, v)
	case map[uint8]float32:
		writeUint8Float32MapBody(enc, v)
	case map[uint8]float64:
		writeUint8Float64MapBody(enc, v)
	case map[uint8]interface{}:
		writeUint8InterfaceMapBody(enc, v)
	case map[uint16]string:
		writeUint16StringMapBody(enc, v)
	case map[uint16]int:
		writeUint16IntMapBody(enc, v)
	case map[uint16]int8:
		writeUint16Int8MapBody(enc, v)
	case map[uint16]int16:
		writeUint16Int16MapBody(enc, v)
	case map[uint16]int32:
		writeUint16Int32MapBody(enc, v)
	case map[uint16]int64:
		writeUint16Int64MapBody(enc, v)
	case map[uint16]uint:
		writeUint16UintMapBody(enc, v)
	case map[uint16]uint8:
		writeUint16Uint8MapBody(enc, v)
	case map[uint16]uint16:
		writeUint16Uint16MapBody(enc, v)
	case map[uint16]uint32:
		writeUint16Uint32MapBody(enc, v)
	case map[uint16]uint64:
		writeUint16Uint64MapBody(enc, v)
	case map[uint16]bool:
		writeUint16BoolMapBody(enc, v)
	case map[uint16]float32:
		writeUint16Float32MapBody(enc, v)
	case map[uint16]float64:
		writeUint16Float64MapBody(enc, v)
	case map[uint16]interface{}:
		writeUint16InterfaceMapBody(enc, v)
	case map[uint32]string:
		writeUint32StringMapBody(enc, v)
	case map[uint32]int:
		writeUint32IntMapBody(enc, v)
	case map[uint32]int8:
		writeUint32Int8MapBody(enc, v)
	case map[uint32]int16:
		writeUint32Int16MapBody(enc, v)
	case map[uint32]int32:
		writeUint32Int32MapBody(enc, v)
	case map[uint32]int64:
		writeUint32Int64MapBody(enc, v)
	case map[uint32]uint:
		writeUint32UintMapBody(enc, v)
	case map[uint32]uint8:
		writeUint32Uint8MapBody(enc, v)
	case map[uint32]uint16:
		writeUint32Uint16MapBody(enc, v)
	case map[uint32]uint32:
		writeUint32Uint32MapBody(enc, v)
	case map[uint32]uint64:
		writeUint32Uint64MapBody(enc, v)
	case map[uint32]bool:
		writeUint32BoolMapBody(enc, v)
	case map[uint32]float32:
		writeUint32Float32MapBody(enc, v)
	case map[uint32]float64:
		writeUint32Float64MapBody(enc, v)
	case map[uint32]interface{}:
		writeUint32InterfaceMapBody(enc, v)
	case map[uint64]string:
		writeUint64StringMapBody(enc, v)
	case map[uint64]int:
		writeUint64IntMapBody(enc, v)
	case map[uint64]int8:
		writeUint64Int8MapBody(enc, v)
	case map[uint64]int16:
		writeUint64Int16MapBody(enc, v)
	case map[uint64]int32:
		writeUint64Int32MapBody(enc, v)
	case map[uint64]int64:
		writeUint64Int64MapBody(enc, v)
	case map[uint64]uint:
		writeUint64UintMapBody(enc, v)
	case map[uint64]uint8:
		writeUint64Uint8MapBody(enc, v)
	case map[uint64]uint16:
		writeUint64Uint16MapBody(enc, v)
	case map[uint64]uint32:
		writeUint64Uint32MapBody(enc, v)
	case map[uint64]uint64:
		writeUint64Uint64MapBody(enc, v)
	case map[uint64]bool:
		writeUint64BoolMapBody(enc, v)
	case map[uint64]float32:
		writeUint64Float32MapBody(enc, v)
	case map[uint64]float64:
		writeUint64Float64MapBody(enc, v)
	case map[uint64]interface{}:
		writeUint64InterfaceMapBody(enc, v)
	case map[float32]string:
		writeFloat32StringMapBody(enc, v)
	case map[float32]int:
		writeFloat32IntMapBody(enc, v)
	case map[float32]int8:
		writeFloat32Int8MapBody(enc, v)
	case map[float32]int16:
		writeFloat32Int16MapBody(enc, v)
	case map[float32]int32:
		writeFloat32Int32MapBody(enc, v)
	case map[float32]int64:
		writeFloat32Int64MapBody(enc, v)
	case map[float32]uint:
		writeFloat32UintMapBody(enc, v)
	case map[float32]uint8:
		writeFloat32Uint8MapBody(enc, v)
	case map[float32]uint16:
		writeFloat32Uint16MapBody(enc, v)
	case map[float32]uint32:
		writeFloat32Uint32MapBody(enc, v)
	case map[float32]uint64:
		writeFloat32Uint64MapBody(enc, v)
	case map[float32]bool:
		writeFloat32BoolMapBody(enc, v)
	case map[float32]float32:
		writeFloat32Float32MapBody(enc, v)
	case map[float32]float64:
		writeFloat32Float64MapBody(enc, v)
	case map[float32]interface{}:
		writeFloat32InterfaceMapBody(enc, v)
	case map[float64]string:
		writeFloat64StringMapBody(enc, v)
	case map[float64]int:
		writeFloat64IntMapBody(enc, v)
	case map[float64]int8:
		writeFloat64Int8MapBody(enc, v)
	case map[float64]int16:
		writeFloat64Int16MapBody(enc, v)
	case map[float64]int32:
		writeFloat64Int32MapBody(enc, v)
	case map[float64]int64:
		writeFloat64Int64MapBody(enc, v)
	case map[float64]uint:
		writeFloat64UintMapBody(enc, v)
	case map[float64]uint8:
		writeFloat64Uint8MapBody(enc, v)
	case map[float64]uint16:
		writeFloat64Uint16MapBody(enc, v)
	case map[float64]uint32:
		writeFloat64Uint32MapBody(enc, v)
	case map[float64]uint64:
		writeFloat64Uint64MapBody(enc, v)
	case map[float64]bool:
		writeFloat64BoolMapBody(enc, v)
	case map[float64]float32:
		writeFloat64Float32MapBody(enc, v)
	case map[float64]float64:
		writeFloat64Float64MapBody(enc, v)
	case map[float64]interface{}:
		writeFloat64InterfaceMapBody(enc, v)
	case map[interface{}]string:
		writeInterfaceStringMapBody(enc, v)
	case map[interface{}]int:
		writeInterfaceIntMapBody(enc, v)
	case map[interface{}]int8:
		writeInterfaceInt8MapBody(enc, v)
	case map[interface{}]int16:
		writeInterfaceInt16MapBody(enc, v)
	case map[interface{}]int32:
		writeInterfaceInt32MapBody(enc, v)
	case map[interface{}]int64:
		writeInterfaceInt64MapBody(enc, v)
	case map[interface{}]uint:
		writeInterfaceUintMapBody(enc, v)
	case map[interface{}]uint8:
		writeInterfaceUint8MapBody(enc, v)
	case map[interface{}]uint16:
		writeInterfaceUint16MapBody(enc, v)
	case map[interface{}]uint32:
		writeInterfaceUint32MapBody(enc, v)
	case map[interface{}]uint64:
		writeInterfaceUint64MapBody(enc, v)
	case map[interface{}]bool:
		writeInterfaceBoolMapBody(enc, v)
	case map[interface{}]float32:
		writeInterfaceFloat32MapBody(enc, v)
	case map[interface{}]float64:
		writeInterfaceFloat64MapBody(enc, v)
	case map[interface{}]interface{}:
		writeInterfaceInterfaceMapBody(enc, v)
	default:
		writeOtherMapBody(enc, v)
	}
}

func writeStringStringMapBody(enc *Encoder, m map[string]string) {
	for k, v := range m {
		EncodeString(enc, k)
		EncodeString(enc, v)
	}

}

func writeStringIntMapBody(enc *Encoder, m map[string]int) {
	for k, v := range m {
		EncodeString(enc, k)
		WriteInt(enc, v)
	}
}

func writeStringInt8MapBody(enc *Encoder, m map[string]int8) {
	for k, v := range m {
		EncodeString(enc, k)
		WriteInt8(enc, v)
	}
}

func writeStringInt16MapBody(enc *Encoder, m map[string]int16) {
	for k, v := range m {
		EncodeString(enc, k)
		WriteInt16(enc, v)
	}
}

func writeStringInt32MapBody(enc *Encoder, m map[string]int32) {
	for k, v := range m {
		EncodeString(enc, k)
		WriteInt32(enc, v)
	}
}

func writeStringInt64MapBody(enc *Encoder, m map[string]int64) {
	for k, v := range m {
		EncodeString(enc, k)
		WriteInt64(enc, v)
	}
}

func writeStringUintMapBody(enc *Encoder, m map[string]uint) {
	for k, v := range m {
		EncodeString(enc, k)
		WriteUint(enc, v)
	}
}

func writeStringUint8MapBody(enc *Encoder, m map[string]uint8) {
	for k, v := range m {
		EncodeString(enc, k)
		WriteUint8(enc, v)
	}
	return
}

func writeStringUint16MapBody(enc *Encoder, m map[string]uint16) {
	for k, v := range m {
		EncodeString(enc, k)
		WriteUint16(enc, v)
	}
	return
}

func writeStringUint32MapBody(enc *Encoder, m map[string]uint32) {
	for k, v := range m {
		EncodeString(enc, k)
		WriteUint32(enc, v)
	}
	return
}

func writeStringUint64MapBody(enc *Encoder, m map[string]uint64) {
	for k, v := range m {
		EncodeString(enc, k)
		WriteUint64(enc, v)
	}
	return
}

func writeStringBoolMapBody(enc *Encoder, m map[string]bool) {
	for k, v := range m {
		EncodeString(enc, k)
		enc.WriteBool(v)
	}
	return
}

func writeStringFloat32MapBody(enc *Encoder, m map[string]float32) {
	for k, v := range m {
		EncodeString(enc, k)
		enc.WriteFloat32(v)
	}
	return
}

func writeStringFloat64MapBody(enc *Encoder, m map[string]float64) {
	for k, v := range m {
		EncodeString(enc, k)
		enc.WriteFloat64(v)
	}
	return
}

func writeStringInterfaceMapBody(enc *Encoder, m map[string]interface{}) {
	for k, v := range m {
		EncodeString(enc, k)
		enc.encode(v)
	}
	return
}

func writeIntStringMapBody(enc *Encoder, m map[int]string) {
	for k, v := range m {
		WriteInt(enc, k)
		EncodeString(enc, v)
	}
	return
}

func writeIntIntMapBody(enc *Encoder, m map[int]int) {
	for k, v := range m {
		WriteInt(enc, k)
		WriteInt(enc, v)
	}
	return
}

func writeIntInt8MapBody(enc *Encoder, m map[int]int8) {
	for k, v := range m {
		WriteInt(enc, k)
		WriteInt8(enc, v)
	}
	return
}

func writeIntInt16MapBody(enc *Encoder, m map[int]int16) {
	for k, v := range m {
		WriteInt(enc, k)
		WriteInt16(enc, v)
	}
	return
}

func writeIntInt32MapBody(enc *Encoder, m map[int]int32) {
	for k, v := range m {
		WriteInt(enc, k)
		WriteInt32(enc, v)
	}
	return
}

func writeIntInt64MapBody(enc *Encoder, m map[int]int64) {
	for k, v := range m {
		WriteInt(enc, k)
		WriteInt64(enc, v)
	}
	return
}

func writeIntUintMapBody(enc *Encoder, m map[int]uint) {
	for k, v := range m {
		WriteInt(enc, k)
		WriteUint(enc, v)
	}
	return
}

func writeIntUint8MapBody(enc *Encoder, m map[int]uint8) {
	for k, v := range m {
		WriteInt(enc, k)
		WriteUint8(enc, v)
	}
	return
}

func writeIntUint16MapBody(enc *Encoder, m map[int]uint16) {
	for k, v := range m {
		WriteInt(enc, k)
		WriteUint16(enc, v)
	}
	return
}

func writeIntUint32MapBody(enc *Encoder, m map[int]uint32) {
	for k, v := range m {
		WriteInt(enc, k)
		WriteUint32(enc, v)
	}
	return
}

func writeIntUint64MapBody(enc *Encoder, m map[int]uint64) {
	for k, v := range m {
		WriteInt(enc, k)
		WriteUint64(enc, v)
	}
	return
}

func writeIntBoolMapBody(enc *Encoder, m map[int]bool) {
	for k, v := range m {
		WriteInt(enc, k)
		enc.WriteBool(v)
	}
	return
}

func writeIntFloat32MapBody(enc *Encoder, m map[int]float32) {
	for k, v := range m {
		WriteInt(enc, k)
		enc.WriteFloat32(v)
	}
	return
}

func writeIntFloat64MapBody(enc *Encoder, m map[int]float64) {
	for k, v := range m {
		WriteInt(enc, k)
		enc.WriteFloat64(v)
	}
	return
}

func writeIntInterfaceMapBody(enc *Encoder, m map[int]interface{}) {
	for k, v := range m {
		WriteInt(enc, k)
		enc.encode(v)
	}
	return
}

func writeInt8StringMapBody(enc *Encoder, m map[int8]string) {
	for k, v := range m {
		WriteInt8(enc, k)
		EncodeString(enc, v)
	}
	return
}

func writeInt8IntMapBody(enc *Encoder, m map[int8]int) {
	for k, v := range m {
		WriteInt8(enc, k)
		WriteInt(enc, v)
	}
	return
}

func writeInt8Int8MapBody(enc *Encoder, m map[int8]int8) {
	for k, v := range m {
		WriteInt8(enc, k)
		WriteInt8(enc, v)
	}
	return
}

func writeInt8Int16MapBody(enc *Encoder, m map[int8]int16) {
	for k, v := range m {
		WriteInt8(enc, k)
		WriteInt16(enc, v)
	}
	return
}

func writeInt8Int32MapBody(enc *Encoder, m map[int8]int32) {
	for k, v := range m {
		WriteInt8(enc, k)
		WriteInt32(enc, v)
	}
	return
}

func writeInt8Int64MapBody(enc *Encoder, m map[int8]int64) {
	for k, v := range m {
		WriteInt8(enc, k)
		WriteInt64(enc, v)
	}
	return
}

func writeInt8UintMapBody(enc *Encoder, m map[int8]uint) {
	for k, v := range m {
		WriteInt8(enc, k)
		WriteUint(enc, v)
	}
	return
}

func writeInt8Uint8MapBody(enc *Encoder, m map[int8]uint8) {
	for k, v := range m {
		WriteInt8(enc, k)
		WriteUint8(enc, v)
	}
	return
}

func writeInt8Uint16MapBody(enc *Encoder, m map[int8]uint16) {
	for k, v := range m {
		WriteInt8(enc, k)
		WriteUint16(enc, v)
	}
	return
}

func writeInt8Uint32MapBody(enc *Encoder, m map[int8]uint32) {
	for k, v := range m {
		WriteInt8(enc, k)
		WriteUint32(enc, v)
	}
	return
}

func writeInt8Uint64MapBody(enc *Encoder, m map[int8]uint64) {
	for k, v := range m {
		WriteInt8(enc, k)
		WriteUint64(enc, v)
	}
	return
}

func writeInt8BoolMapBody(enc *Encoder, m map[int8]bool) {
	for k, v := range m {
		WriteInt8(enc, k)
		enc.WriteBool(v)
	}
	return
}

func writeInt8Float32MapBody(enc *Encoder, m map[int8]float32) {
	for k, v := range m {
		WriteInt8(enc, k)
		enc.WriteFloat32(v)
	}
	return
}

func writeInt8Float64MapBody(enc *Encoder, m map[int8]float64) {
	for k, v := range m {
		WriteInt8(enc, k)
		enc.WriteFloat64(v)
	}
	return
}

func writeInt8InterfaceMapBody(enc *Encoder, m map[int8]interface{}) {
	for k, v := range m {
		WriteInt8(enc, k)
		enc.encode(v)
	}
	return
}

func writeInt16StringMapBody(enc *Encoder, m map[int16]string) {
	for k, v := range m {
		WriteInt16(enc, k)
		EncodeString(enc, v)
	}
	return
}

func writeInt16IntMapBody(enc *Encoder, m map[int16]int) {
	for k, v := range m {
		WriteInt16(enc, k)
		WriteInt(enc, v)
	}
	return
}

func writeInt16Int8MapBody(enc *Encoder, m map[int16]int8) {
	for k, v := range m {
		WriteInt16(enc, k)
		WriteInt8(enc, v)
	}
	return
}

func writeInt16Int16MapBody(enc *Encoder, m map[int16]int16) {
	for k, v := range m {
		WriteInt16(enc, k)
		WriteInt16(enc, v)
	}
	return
}

func writeInt16Int32MapBody(enc *Encoder, m map[int16]int32) {
	for k, v := range m {
		WriteInt16(enc, k)
		WriteInt32(enc, v)
	}
	return
}

func writeInt16Int64MapBody(enc *Encoder, m map[int16]int64) {
	for k, v := range m {
		WriteInt16(enc, k)
		WriteInt64(enc, v)
	}
	return
}

func writeInt16UintMapBody(enc *Encoder, m map[int16]uint) {
	for k, v := range m {
		WriteInt16(enc, k)
		WriteUint(enc, v)
	}
	return
}

func writeInt16Uint8MapBody(enc *Encoder, m map[int16]uint8) {
	for k, v := range m {
		WriteInt16(enc, k)
		WriteUint8(enc, v)
	}
	return
}

func writeInt16Uint16MapBody(enc *Encoder, m map[int16]uint16) {
	for k, v := range m {
		WriteInt16(enc, k)
		WriteUint16(enc, v)
	}
	return
}

func writeInt16Uint32MapBody(enc *Encoder, m map[int16]uint32) {
	for k, v := range m {
		WriteInt16(enc, k)
		WriteUint32(enc, v)
	}
	return
}

func writeInt16Uint64MapBody(enc *Encoder, m map[int16]uint64) {
	for k, v := range m {
		WriteInt16(enc, k)
		WriteUint64(enc, v)
	}
	return
}

func writeInt16BoolMapBody(enc *Encoder, m map[int16]bool) {
	for k, v := range m {
		WriteInt16(enc, k)
		enc.WriteBool(v)
	}
	return
}

func writeInt16Float32MapBody(enc *Encoder, m map[int16]float32) {
	for k, v := range m {
		WriteInt16(enc, k)
		enc.WriteFloat32(v)
	}
	return
}

func writeInt16Float64MapBody(enc *Encoder, m map[int16]float64) {
	for k, v := range m {
		WriteInt16(enc, k)
		enc.WriteFloat64(v)
	}
	return
}

func writeInt16InterfaceMapBody(enc *Encoder, m map[int16]interface{}) {
	for k, v := range m {
		WriteInt16(enc, k)
		enc.encode(v)
	}
	return
}

func writeInt32StringMapBody(enc *Encoder, m map[int32]string) {
	for k, v := range m {
		WriteInt32(enc, k)
		EncodeString(enc, v)
	}
	return
}

func writeInt32IntMapBody(enc *Encoder, m map[int32]int) {
	for k, v := range m {
		WriteInt32(enc, k)
		WriteInt(enc, v)
	}
	return
}

func writeInt32Int8MapBody(enc *Encoder, m map[int32]int8) {
	for k, v := range m {
		WriteInt32(enc, k)
		WriteInt8(enc, v)
	}
	return
}

func writeInt32Int16MapBody(enc *Encoder, m map[int32]int16) {
	for k, v := range m {
		WriteInt32(enc, k)
		WriteInt16(enc, v)
	}
	return
}

func writeInt32Int32MapBody(enc *Encoder, m map[int32]int32) {
	for k, v := range m {
		WriteInt32(enc, k)
		WriteInt32(enc, v)
	}
	return
}

func writeInt32Int64MapBody(enc *Encoder, m map[int32]int64) {
	for k, v := range m {
		WriteInt32(enc, k)
		WriteInt64(enc, v)
	}
	return
}

func writeInt32UintMapBody(enc *Encoder, m map[int32]uint) {
	for k, v := range m {
		WriteInt32(enc, k)
		WriteUint(enc, v)
	}
	return
}

func writeInt32Uint8MapBody(enc *Encoder, m map[int32]uint8) {
	for k, v := range m {
		WriteInt32(enc, k)
		WriteUint8(enc, v)
	}
	return
}

func writeInt32Uint16MapBody(enc *Encoder, m map[int32]uint16) {
	for k, v := range m {
		WriteInt32(enc, k)
		WriteUint16(enc, v)
	}
	return
}

func writeInt32Uint32MapBody(enc *Encoder, m map[int32]uint32) {
	for k, v := range m {
		WriteInt32(enc, k)
		WriteUint32(enc, v)
	}
	return
}

func writeInt32Uint64MapBody(enc *Encoder, m map[int32]uint64) {
	for k, v := range m {
		WriteInt32(enc, k)
		WriteUint64(enc, v)
	}
	return
}

func writeInt32BoolMapBody(enc *Encoder, m map[int32]bool) {
	for k, v := range m {
		WriteInt32(enc, k)
		enc.WriteBool(v)
	}
	return
}

func writeInt32Float32MapBody(enc *Encoder, m map[int32]float32) {
	for k, v := range m {
		WriteInt32(enc, k)
		enc.WriteFloat32(v)
	}
	return
}

func writeInt32Float64MapBody(enc *Encoder, m map[int32]float64) {
	for k, v := range m {
		WriteInt32(enc, k)
		enc.WriteFloat64(v)
	}
	return
}

func writeInt32InterfaceMapBody(enc *Encoder, m map[int32]interface{}) {
	for k, v := range m {
		WriteInt32(enc, k)
		enc.encode(v)
	}
	return
}

func writeInt64StringMapBody(enc *Encoder, m map[int64]string) {
	for k, v := range m {
		WriteInt64(enc, k)
		EncodeString(enc, v)
	}
	return
}

func writeInt64IntMapBody(enc *Encoder, m map[int64]int) {
	for k, v := range m {
		WriteInt64(enc, k)
		WriteInt(enc, v)
	}
	return
}

func writeInt64Int8MapBody(enc *Encoder, m map[int64]int8) {
	for k, v := range m {
		WriteInt64(enc, k)
		WriteInt8(enc, v)
	}
	return
}

func writeInt64Int16MapBody(enc *Encoder, m map[int64]int16) {
	for k, v := range m {
		WriteInt64(enc, k)
		WriteInt16(enc, v)
	}
	return
}

func writeInt64Int32MapBody(enc *Encoder, m map[int64]int32) {
	for k, v := range m {
		WriteInt64(enc, k)
		WriteInt32(enc, v)
	}
	return
}

func writeInt64Int64MapBody(enc *Encoder, m map[int64]int64) {
	for k, v := range m {
		WriteInt64(enc, k)
		WriteInt64(enc, v)
	}
	return
}

func writeInt64UintMapBody(enc *Encoder, m map[int64]uint) {
	for k, v := range m {
		WriteInt64(enc, k)
		WriteUint(enc, v)
	}
	return
}

func writeInt64Uint8MapBody(enc *Encoder, m map[int64]uint8) {
	for k, v := range m {
		WriteInt64(enc, k)
		WriteUint8(enc, v)
	}
	return
}

func writeInt64Uint16MapBody(enc *Encoder, m map[int64]uint16) {
	for k, v := range m {
		WriteInt64(enc, k)
		WriteUint16(enc, v)
	}
	return
}

func writeInt64Uint32MapBody(enc *Encoder, m map[int64]uint32) {
	for k, v := range m {
		WriteInt64(enc, k)
		WriteUint32(enc, v)
	}
	return
}

func writeInt64Uint64MapBody(enc *Encoder, m map[int64]uint64) {
	for k, v := range m {
		WriteInt64(enc, k)
		WriteUint64(enc, v)
	}
	return
}

func writeInt64BoolMapBody(enc *Encoder, m map[int64]bool) {
	for k, v := range m {
		WriteInt64(enc, k)
		enc.WriteBool(v)
	}
	return
}

func writeInt64Float32MapBody(enc *Encoder, m map[int64]float32) {
	for k, v := range m {
		WriteInt64(enc, k)
		enc.WriteFloat32(v)
	}
	return
}

func writeInt64Float64MapBody(enc *Encoder, m map[int64]float64) {
	for k, v := range m {
		WriteInt64(enc, k)
		enc.WriteFloat64(v)
	}
	return
}

func writeInt64InterfaceMapBody(enc *Encoder, m map[int64]interface{}) {
	for k, v := range m {
		WriteInt64(enc, k)
		enc.encode(v)
	}
	return
}

func writeUintStringMapBody(enc *Encoder, m map[uint]string) {
	for k, v := range m {
		WriteUint(enc, k)
		EncodeString(enc, v)
	}
	return
}

func writeUintIntMapBody(enc *Encoder, m map[uint]int) {
	for k, v := range m {
		WriteUint(enc, k)
		WriteInt(enc, v)
	}
	return
}

func writeUintInt8MapBody(enc *Encoder, m map[uint]int8) {
	for k, v := range m {
		WriteUint(enc, k)
		WriteInt8(enc, v)
	}
	return
}

func writeUintInt16MapBody(enc *Encoder, m map[uint]int16) {
	for k, v := range m {
		WriteUint(enc, k)
		WriteInt16(enc, v)
	}
	return
}

func writeUintInt32MapBody(enc *Encoder, m map[uint]int32) {
	for k, v := range m {
		WriteUint(enc, k)
		WriteInt32(enc, v)
	}
	return
}

func writeUintInt64MapBody(enc *Encoder, m map[uint]int64) {
	for k, v := range m {
		WriteUint(enc, k)
		WriteInt64(enc, v)
	}
	return
}

func writeUintUintMapBody(enc *Encoder, m map[uint]uint) {
	for k, v := range m {
		WriteUint(enc, k)
		WriteUint(enc, v)
	}
	return
}

func writeUintUint8MapBody(enc *Encoder, m map[uint]uint8) {
	for k, v := range m {
		WriteUint(enc, k)
		WriteUint8(enc, v)
	}
	return
}

func writeUintUint16MapBody(enc *Encoder, m map[uint]uint16) {
	for k, v := range m {
		WriteUint(enc, k)
		WriteUint16(enc, v)
	}
	return
}

func writeUintUint32MapBody(enc *Encoder, m map[uint]uint32) {
	for k, v := range m {
		WriteUint(enc, k)
		WriteUint32(enc, v)
	}
	return
}

func writeUintUint64MapBody(enc *Encoder, m map[uint]uint64) {
	for k, v := range m {
		WriteUint(enc, k)
		WriteUint64(enc, v)
	}
	return
}

func writeUintBoolMapBody(enc *Encoder, m map[uint]bool) {
	for k, v := range m {
		WriteUint(enc, k)
		enc.WriteBool(v)
	}
	return
}

func writeUintFloat32MapBody(enc *Encoder, m map[uint]float32) {
	for k, v := range m {
		WriteUint(enc, k)
		enc.WriteFloat32(v)
	}
	return
}

func writeUintFloat64MapBody(enc *Encoder, m map[uint]float64) {
	for k, v := range m {
		WriteUint(enc, k)
		enc.WriteFloat64(v)
	}
	return
}

func writeUintInterfaceMapBody(enc *Encoder, m map[uint]interface{}) {
	for k, v := range m {
		WriteUint(enc, k)
		enc.encode(v)
	}
	return
}
func writeUint8StringMapBody(enc *Encoder, m map[uint8]string) {
	for k, v := range m {
		WriteUint8(enc, k)
		EncodeString(enc, v)
	}
	return
}

func writeUint8IntMapBody(enc *Encoder, m map[uint8]int) {
	for k, v := range m {
		WriteUint8(enc, k)
		WriteInt(enc, v)
	}
	return
}

func writeUint8Int8MapBody(enc *Encoder, m map[uint8]int8) {
	for k, v := range m {
		WriteUint8(enc, k)
		WriteInt8(enc, v)
	}
	return
}

func writeUint8Int16MapBody(enc *Encoder, m map[uint8]int16) {
	for k, v := range m {
		WriteUint8(enc, k)
		WriteInt16(enc, v)
	}
	return
}

func writeUint8Int32MapBody(enc *Encoder, m map[uint8]int32) {
	for k, v := range m {
		WriteUint8(enc, k)
		WriteInt32(enc, v)
	}
	return
}

func writeUint8Int64MapBody(enc *Encoder, m map[uint8]int64) {
	for k, v := range m {
		WriteUint8(enc, k)
		WriteInt64(enc, v)
	}
	return
}

func writeUint8UintMapBody(enc *Encoder, m map[uint8]uint) {
	for k, v := range m {
		WriteUint8(enc, k)
		WriteUint(enc, v)
	}
	return
}

func writeUint8Uint8MapBody(enc *Encoder, m map[uint8]uint8) {
	for k, v := range m {
		WriteUint8(enc, k)
		WriteUint8(enc, v)
	}
	return
}

func writeUint8Uint16MapBody(enc *Encoder, m map[uint8]uint16) {
	for k, v := range m {
		WriteUint8(enc, k)
		WriteUint16(enc, v)
	}
	return
}

func writeUint8Uint32MapBody(enc *Encoder, m map[uint8]uint32) {
	for k, v := range m {
		WriteUint8(enc, k)
		WriteUint32(enc, v)
	}
	return
}

func writeUint8Uint64MapBody(enc *Encoder, m map[uint8]uint64) {
	for k, v := range m {
		WriteUint8(enc, k)
		WriteUint64(enc, v)
	}
	return
}

func writeUint8BoolMapBody(enc *Encoder, m map[uint8]bool) {
	for k, v := range m {
		WriteUint8(enc, k)
		enc.WriteBool(v)
	}
	return
}

func writeUint8Float32MapBody(enc *Encoder, m map[uint8]float32) {
	for k, v := range m {
		WriteUint8(enc, k)
		enc.WriteFloat32(v)
	}
	return
}

func writeUint8Float64MapBody(enc *Encoder, m map[uint8]float64) {
	for k, v := range m {
		WriteUint8(enc, k)
		enc.WriteFloat64(v)
	}
	return
}

func writeUint8InterfaceMapBody(enc *Encoder, m map[uint8]interface{}) {
	for k, v := range m {
		WriteUint8(enc, k)
		enc.encode(v)
	}
	return
}

func writeUint16StringMapBody(enc *Encoder, m map[uint16]string) {
	for k, v := range m {
		WriteUint16(enc, k)
		EncodeString(enc, v)
	}
	return
}

func writeUint16IntMapBody(enc *Encoder, m map[uint16]int) {
	for k, v := range m {
		WriteUint16(enc, k)
		WriteInt(enc, v)
	}
	return
}

func writeUint16Int8MapBody(enc *Encoder, m map[uint16]int8) {
	for k, v := range m {
		WriteUint16(enc, k)
		WriteInt8(enc, v)
	}
	return
}

func writeUint16Int16MapBody(enc *Encoder, m map[uint16]int16) {
	for k, v := range m {
		WriteUint16(enc, k)
		WriteInt16(enc, v)
	}
	return
}

func writeUint16Int32MapBody(enc *Encoder, m map[uint16]int32) {
	for k, v := range m {
		WriteUint16(enc, k)
		WriteInt32(enc, v)
	}
	return
}

func writeUint16Int64MapBody(enc *Encoder, m map[uint16]int64) {
	for k, v := range m {
		WriteUint16(enc, k)
		WriteInt64(enc, v)
	}
	return
}

func writeUint16UintMapBody(enc *Encoder, m map[uint16]uint) {
	for k, v := range m {
		WriteUint16(enc, k)
		WriteUint(enc, v)
	}
	return
}

func writeUint16Uint8MapBody(enc *Encoder, m map[uint16]uint8) {
	for k, v := range m {
		WriteUint16(enc, k)
		WriteUint8(enc, v)
	}
	return
}

func writeUint16Uint16MapBody(enc *Encoder, m map[uint16]uint16) {
	for k, v := range m {
		WriteUint16(enc, k)
		WriteUint16(enc, v)
	}
	return
}

func writeUint16Uint32MapBody(enc *Encoder, m map[uint16]uint32) {
	for k, v := range m {
		WriteUint16(enc, k)
		WriteUint32(enc, v)
	}
	return
}

func writeUint16Uint64MapBody(enc *Encoder, m map[uint16]uint64) {
	for k, v := range m {
		WriteUint16(enc, k)
		WriteUint64(enc, v)
	}
	return
}

func writeUint16BoolMapBody(enc *Encoder, m map[uint16]bool) {
	for k, v := range m {
		WriteUint16(enc, k)
		enc.WriteBool(v)
	}
	return
}

func writeUint16Float32MapBody(enc *Encoder, m map[uint16]float32) {
	for k, v := range m {
		WriteUint16(enc, k)
		enc.WriteFloat32(v)
	}
	return
}

func writeUint16Float64MapBody(enc *Encoder, m map[uint16]float64) {
	for k, v := range m {
		WriteUint16(enc, k)
		enc.WriteFloat64(v)
	}
	return
}

func writeUint16InterfaceMapBody(enc *Encoder, m map[uint16]interface{}) {
	for k, v := range m {
		WriteUint16(enc, k)
		enc.encode(v)
	}
	return
}

func writeUint32StringMapBody(enc *Encoder, m map[uint32]string) {
	for k, v := range m {
		WriteUint32(enc, k)
		EncodeString(enc, v)
	}
	return
}

func writeUint32IntMapBody(enc *Encoder, m map[uint32]int) {
	for k, v := range m {
		WriteUint32(enc, k)
		WriteInt(enc, v)
	}
	return
}

func writeUint32Int8MapBody(enc *Encoder, m map[uint32]int8) {
	for k, v := range m {
		WriteUint32(enc, k)
		WriteInt8(enc, v)
	}
	return
}

func writeUint32Int16MapBody(enc *Encoder, m map[uint32]int16) {
	for k, v := range m {
		WriteUint32(enc, k)
		WriteInt16(enc, v)
	}
	return
}

func writeUint32Int32MapBody(enc *Encoder, m map[uint32]int32) {
	for k, v := range m {
		WriteUint32(enc, k)
		WriteInt32(enc, v)
	}
	return
}

func writeUint32Int64MapBody(enc *Encoder, m map[uint32]int64) {
	for k, v := range m {
		WriteUint32(enc, k)
		WriteInt64(enc, v)
	}
	return
}

func writeUint32UintMapBody(enc *Encoder, m map[uint32]uint) {
	for k, v := range m {
		WriteUint32(enc, k)
		WriteUint(enc, v)
	}
	return
}

func writeUint32Uint8MapBody(enc *Encoder, m map[uint32]uint8) {
	for k, v := range m {
		WriteUint32(enc, k)
		WriteUint8(enc, v)
	}
	return
}

func writeUint32Uint16MapBody(enc *Encoder, m map[uint32]uint16) {
	for k, v := range m {
		WriteUint32(enc, k)
		WriteUint16(enc, v)
	}
	return
}

func writeUint32Uint32MapBody(enc *Encoder, m map[uint32]uint32) {
	for k, v := range m {
		WriteUint32(enc, k)
		WriteUint32(enc, v)
	}
	return
}

func writeUint32Uint64MapBody(enc *Encoder, m map[uint32]uint64) {
	for k, v := range m {
		WriteUint32(enc, k)
		WriteUint64(enc, v)
	}
	return
}

func writeUint32BoolMapBody(enc *Encoder, m map[uint32]bool) {
	for k, v := range m {
		WriteUint32(enc, k)
		enc.WriteBool(v)
	}
	return
}

func writeUint32Float32MapBody(enc *Encoder, m map[uint32]float32) {
	for k, v := range m {
		WriteUint32(enc, k)
		enc.WriteFloat32(v)
	}
	return
}

func writeUint32Float64MapBody(enc *Encoder, m map[uint32]float64) {
	for k, v := range m {
		WriteUint32(enc, k)
		enc.WriteFloat64(v)
	}
	return
}

func writeUint32InterfaceMapBody(enc *Encoder, m map[uint32]interface{}) {
	for k, v := range m {
		WriteUint32(enc, k)
		enc.encode(v)
	}
	return
}

func writeUint64StringMapBody(enc *Encoder, m map[uint64]string) {
	for k, v := range m {
		WriteUint64(enc, k)
		EncodeString(enc, v)
	}
	return
}

func writeUint64IntMapBody(enc *Encoder, m map[uint64]int) {
	for k, v := range m {
		WriteUint64(enc, k)
		WriteInt(enc, v)
	}
	return
}

func writeUint64Int8MapBody(enc *Encoder, m map[uint64]int8) {
	for k, v := range m {
		WriteUint64(enc, k)
		WriteInt8(enc, v)
	}
	return
}

func writeUint64Int16MapBody(enc *Encoder, m map[uint64]int16) {
	for k, v := range m {
		WriteUint64(enc, k)
		WriteInt16(enc, v)
	}
	return
}

func writeUint64Int32MapBody(enc *Encoder, m map[uint64]int32) {
	for k, v := range m {
		WriteUint64(enc, k)
		WriteInt32(enc, v)
	}
	return
}

func writeUint64Int64MapBody(enc *Encoder, m map[uint64]int64) {
	for k, v := range m {
		WriteUint64(enc, k)
		WriteInt64(enc, v)
	}
	return
}

func writeUint64UintMapBody(enc *Encoder, m map[uint64]uint) {
	for k, v := range m {
		WriteUint64(enc, k)
		WriteUint(enc, v)
	}
	return
}

func writeUint64Uint8MapBody(enc *Encoder, m map[uint64]uint8) {
	for k, v := range m {
		WriteUint64(enc, k)
		WriteUint8(enc, v)
	}
	return
}

func writeUint64Uint16MapBody(enc *Encoder, m map[uint64]uint16) {
	for k, v := range m {
		WriteUint64(enc, k)
		WriteUint16(enc, v)
	}
	return
}

func writeUint64Uint32MapBody(enc *Encoder, m map[uint64]uint32) {
	for k, v := range m {
		WriteUint64(enc, k)
		WriteUint32(enc, v)
	}
	return
}

func writeUint64Uint64MapBody(enc *Encoder, m map[uint64]uint64) {
	for k, v := range m {
		WriteUint64(enc, k)
		WriteUint64(enc, v)
	}
	return
}

func writeUint64BoolMapBody(enc *Encoder, m map[uint64]bool) {
	for k, v := range m {
		WriteUint64(enc, k)
		enc.WriteBool(v)
	}
	return
}

func writeUint64Float32MapBody(enc *Encoder, m map[uint64]float32) {
	for k, v := range m {
		WriteUint64(enc, k)
		enc.WriteFloat32(v)
	}
	return
}

func writeUint64Float64MapBody(enc *Encoder, m map[uint64]float64) {
	for k, v := range m {
		WriteUint64(enc, k)
		enc.WriteFloat64(v)
	}
	return
}

func writeUint64InterfaceMapBody(enc *Encoder, m map[uint64]interface{}) {
	for k, v := range m {
		WriteUint64(enc, k)
		enc.encode(v)
	}
	return
}

func writeFloat32StringMapBody(enc *Encoder, m map[float32]string) {
	for k, v := range m {
		enc.WriteFloat32(k)
		EncodeString(enc, v)
	}
	return
}

func writeFloat32IntMapBody(enc *Encoder, m map[float32]int) {
	for k, v := range m {
		enc.WriteFloat32(k)
		WriteInt(enc, v)
	}
	return
}

func writeFloat32Int8MapBody(enc *Encoder, m map[float32]int8) {
	for k, v := range m {
		enc.WriteFloat32(k)
		WriteInt8(enc, v)
	}
	return
}

func writeFloat32Int16MapBody(enc *Encoder, m map[float32]int16) {
	for k, v := range m {
		enc.WriteFloat32(k)
		WriteInt16(enc, v)
	}
	return
}

func writeFloat32Int32MapBody(enc *Encoder, m map[float32]int32) {
	for k, v := range m {
		enc.WriteFloat32(k)
		WriteInt32(enc, v)
	}
	return
}

func writeFloat32Int64MapBody(enc *Encoder, m map[float32]int64) {
	for k, v := range m {
		enc.WriteFloat32(k)
		WriteInt64(enc, v)
	}
	return
}

func writeFloat32UintMapBody(enc *Encoder, m map[float32]uint) {
	for k, v := range m {
		enc.WriteFloat32(k)
		WriteUint(enc, v)
	}
	return
}

func writeFloat32Uint8MapBody(enc *Encoder, m map[float32]uint8) {
	for k, v := range m {
		enc.WriteFloat32(k)
		WriteUint8(enc, v)
	}
	return
}

func writeFloat32Uint16MapBody(enc *Encoder, m map[float32]uint16) {
	for k, v := range m {
		enc.WriteFloat32(k)
		WriteUint16(enc, v)
	}
	return
}

func writeFloat32Uint32MapBody(enc *Encoder, m map[float32]uint32) {
	for k, v := range m {
		enc.WriteFloat32(k)
		WriteUint32(enc, v)
	}
	return
}

func writeFloat32Uint64MapBody(enc *Encoder, m map[float32]uint64) {
	for k, v := range m {
		enc.WriteFloat32(k)
		WriteUint64(enc, v)
	}
	return
}

func writeFloat32BoolMapBody(enc *Encoder, m map[float32]bool) {
	for k, v := range m {
		enc.WriteFloat32(k)
		enc.WriteBool(v)
	}
	return
}

func writeFloat32Float32MapBody(enc *Encoder, m map[float32]float32) {
	for k, v := range m {
		enc.WriteFloat32(k)
		enc.WriteFloat32(v)
	}
	return
}

func writeFloat32Float64MapBody(enc *Encoder, m map[float32]float64) {
	for k, v := range m {
		enc.WriteFloat32(k)
		enc.WriteFloat64(v)
	}
	return
}

func writeFloat32InterfaceMapBody(enc *Encoder, m map[float32]interface{}) {
	for k, v := range m {
		enc.WriteFloat32(k)
		enc.encode(v)
	}
	return
}

func writeFloat64StringMapBody(enc *Encoder, m map[float64]string) {
	for k, v := range m {
		enc.WriteFloat64(k)
		EncodeString(enc, v)
	}
	return
}

func writeFloat64IntMapBody(enc *Encoder, m map[float64]int) {
	for k, v := range m {
		enc.WriteFloat64(k)
		WriteInt(enc, v)
	}
	return
}

func writeFloat64Int8MapBody(enc *Encoder, m map[float64]int8) {
	for k, v := range m {
		enc.WriteFloat64(k)
		WriteInt8(enc, v)
	}
	return
}

func writeFloat64Int16MapBody(enc *Encoder, m map[float64]int16) {
	for k, v := range m {
		enc.WriteFloat64(k)
		WriteInt16(enc, v)
	}
	return
}

func writeFloat64Int32MapBody(enc *Encoder, m map[float64]int32) {
	for k, v := range m {
		enc.WriteFloat64(k)
		WriteInt32(enc, v)
	}
	return
}

func writeFloat64Int64MapBody(enc *Encoder, m map[float64]int64) {
	for k, v := range m {
		enc.WriteFloat64(k)
		WriteInt64(enc, v)
	}
	return
}

func writeFloat64UintMapBody(enc *Encoder, m map[float64]uint) {
	for k, v := range m {
		enc.WriteFloat64(k)
		WriteUint(enc, v)
	}
	return
}

func writeFloat64Uint8MapBody(enc *Encoder, m map[float64]uint8) {
	for k, v := range m {
		enc.WriteFloat64(k)
		WriteUint8(enc, v)
	}
	return
}

func writeFloat64Uint16MapBody(enc *Encoder, m map[float64]uint16) {
	for k, v := range m {
		enc.WriteFloat64(k)
		WriteUint16(enc, v)
	}
	return
}

func writeFloat64Uint32MapBody(enc *Encoder, m map[float64]uint32) {
	for k, v := range m {
		enc.WriteFloat64(k)
		WriteUint32(enc, v)
	}
	return
}

func writeFloat64Uint64MapBody(enc *Encoder, m map[float64]uint64) {
	for k, v := range m {
		enc.WriteFloat64(k)
		WriteUint64(enc, v)
	}
	return
}

func writeFloat64BoolMapBody(enc *Encoder, m map[float64]bool) {
	for k, v := range m {
		enc.WriteFloat64(k)
		enc.WriteBool(v)
	}
	return
}

func writeFloat64Float32MapBody(enc *Encoder, m map[float64]float32) {
	for k, v := range m {
		enc.WriteFloat64(k)
		enc.WriteFloat32(v)
	}
	return
}

func writeFloat64Float64MapBody(enc *Encoder, m map[float64]float64) {
	for k, v := range m {
		enc.WriteFloat64(k)
		enc.WriteFloat64(v)
	}
	return
}

func writeFloat64InterfaceMapBody(enc *Encoder, m map[float64]interface{}) {
	for k, v := range m {
		enc.WriteFloat64(k)
		enc.encode(v)
	}
	return
}

func writeInterfaceStringMapBody(enc *Encoder, m map[interface{}]string) {
	for k, v := range m {
		enc.encode(k)
		EncodeString(enc, v)
	}
	return
}

func writeInterfaceIntMapBody(enc *Encoder, m map[interface{}]int) {
	for k, v := range m {
		enc.encode(k)
		WriteInt(enc, v)
	}
	return
}

func writeInterfaceInt8MapBody(enc *Encoder, m map[interface{}]int8) {
	for k, v := range m {
		enc.encode(k)
		WriteInt8(enc, v)
	}
	return
}

func writeInterfaceInt16MapBody(enc *Encoder, m map[interface{}]int16) {
	for k, v := range m {
		enc.encode(k)
		WriteInt16(enc, v)
	}
	return
}

func writeInterfaceInt32MapBody(enc *Encoder, m map[interface{}]int32) {
	for k, v := range m {
		enc.encode(k)
		WriteInt32(enc, v)
	}
	return
}

func writeInterfaceInt64MapBody(enc *Encoder, m map[interface{}]int64) {
	for k, v := range m {
		enc.encode(k)
		WriteInt64(enc, v)
	}
	return
}

func writeInterfaceUintMapBody(enc *Encoder, m map[interface{}]uint) {
	for k, v := range m {
		enc.encode(k)
		WriteUint(enc, v)
	}
	return
}

func writeInterfaceUint8MapBody(enc *Encoder, m map[interface{}]uint8) {
	for k, v := range m {
		enc.encode(k)
		WriteUint8(enc, v)
	}
	return
}

func writeInterfaceUint16MapBody(enc *Encoder, m map[interface{}]uint16) {
	for k, v := range m {
		enc.encode(k)
		WriteUint16(enc, v)
	}
	return
}

func writeInterfaceUint32MapBody(enc *Encoder, m map[interface{}]uint32) {
	for k, v := range m {
		enc.encode(k)
		WriteUint32(enc, v)
	}
	return
}

func writeInterfaceUint64MapBody(enc *Encoder, m map[interface{}]uint64) {
	for k, v := range m {
		enc.encode(k)
		WriteUint64(enc, v)
	}
	return
}

func writeInterfaceBoolMapBody(enc *Encoder, m map[interface{}]bool) {
	for k, v := range m {
		enc.encode(k)
		enc.WriteBool(v)
	}
	return
}

func writeInterfaceFloat32MapBody(enc *Encoder, m map[interface{}]float32) {
	for k, v := range m {
		enc.encode(k)
		enc.WriteFloat32(v)
	}
	return
}

func writeInterfaceFloat64MapBody(enc *Encoder, m map[interface{}]float64) {
	for k, v := range m {
		enc.encode(k)
		enc.WriteFloat64(v)
	}
	return
}

func writeInterfaceInterfaceMapBody(enc *Encoder, m map[interface{}]interface{}) {
	for k, v := range m {
		enc.encode(k)
		enc.encode(v)
	}
	return
}

func writeOtherMapBody(enc *Encoder, v interface{}) {
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
	return
}
