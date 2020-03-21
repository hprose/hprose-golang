/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/map_encoder.go                               |
|                                                          |
| LastModified: Mar 17, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"reflect"
	"unsafe"

	"github.com/hprose/hprose-golang/v3/io"
	"github.com/modern-go/reflect2"
)

// mapEncoder is the implementation of ValueEncoder for *map.
type mapEncoder struct{}

var mapenc mapEncoder

func (valenc mapEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	var ok bool
	if ok, err = enc.WriteReference(v); !ok && err == nil {
		err = valenc.Write(enc, v)
	}
	return
}

func (mapEncoder) Write(enc *Encoder, v interface{}) (err error) {
	enc.SetReference(v)
	return writeMap(enc, reflect.ValueOf(v).Elem().Interface())
}

// WriteMap to encoder
func WriteMap(enc *Encoder, v interface{}) (err error) {
	enc.AddReferenceCount(1)
	return writeMap(enc, v)
}

var emptyMap = []byte{io.TagMap, io.TagOpenbrace, io.TagClosebrace}

func writeMap(enc *Encoder, v interface{}) (err error) {
	writer := enc.Writer
	count := reflect.ValueOf(v).Len()
	if count == 0 {
		_, err = writer.Write(emptyMap)
		return
	}
	if err = WriteHead(writer, count, io.TagMap); err == nil {
		if err = writeMapBody(enc, v); err == nil {
			err = WriteFoot(writer)
		}
	}
	return
}

func writeMapBody(enc *Encoder, v interface{}) error {
	switch v := v.(type) {
	case map[string]string:
		return writeStringStringMapBody(enc, v)
	case map[string]int:
		return writeStringIntMapBody(enc, v)
	case map[string]int8:
		return writeStringInt8MapBody(enc, v)
	case map[string]int16:
		return writeStringInt16MapBody(enc, v)
	case map[string]int32:
		return writeStringInt32MapBody(enc, v)
	case map[string]int64:
		return writeStringInt64MapBody(enc, v)
	case map[string]uint:
		return writeStringUintMapBody(enc, v)
	case map[string]uint8:
		return writeStringUint8MapBody(enc, v)
	case map[string]uint16:
		return writeStringUint16MapBody(enc, v)
	case map[string]uint32:
		return writeStringUint32MapBody(enc, v)
	case map[string]uint64:
		return writeStringUint64MapBody(enc, v)
	case map[string]bool:
		return writeStringBoolMapBody(enc, v)
	case map[string]float32:
		return writeStringFloat32MapBody(enc, v)
	case map[string]float64:
		return writeStringFloat64MapBody(enc, v)
	case map[string]interface{}:
		return writeStringInterfaceMapBody(enc, v)
	case map[int]string:
		return writeIntStringMapBody(enc, v)
	case map[int]int:
		return writeIntIntMapBody(enc, v)
	case map[int]int8:
		return writeIntInt8MapBody(enc, v)
	case map[int]int16:
		return writeIntInt16MapBody(enc, v)
	case map[int]int32:
		return writeIntInt32MapBody(enc, v)
	case map[int]int64:
		return writeIntInt64MapBody(enc, v)
	case map[int]uint:
		return writeIntUintMapBody(enc, v)
	case map[int]uint8:
		return writeIntUint8MapBody(enc, v)
	case map[int]uint16:
		return writeIntUint16MapBody(enc, v)
	case map[int]uint32:
		return writeIntUint32MapBody(enc, v)
	case map[int]uint64:
		return writeIntUint64MapBody(enc, v)
	case map[int]bool:
		return writeIntBoolMapBody(enc, v)
	case map[int]float32:
		return writeIntFloat32MapBody(enc, v)
	case map[int]float64:
		return writeIntFloat64MapBody(enc, v)
	case map[int]interface{}:
		return writeIntInterfaceMapBody(enc, v)
	case map[int8]string:
		return writeInt8StringMapBody(enc, v)
	case map[int8]int:
		return writeInt8IntMapBody(enc, v)
	case map[int8]int8:
		return writeInt8Int8MapBody(enc, v)
	case map[int8]int16:
		return writeInt8Int16MapBody(enc, v)
	case map[int8]int32:
		return writeInt8Int32MapBody(enc, v)
	case map[int8]int64:
		return writeInt8Int64MapBody(enc, v)
	case map[int8]uint:
		return writeInt8UintMapBody(enc, v)
	case map[int8]uint8:
		return writeInt8Uint8MapBody(enc, v)
	case map[int8]uint16:
		return writeInt8Uint16MapBody(enc, v)
	case map[int8]uint32:
		return writeInt8Uint32MapBody(enc, v)
	case map[int8]uint64:
		return writeInt8Uint64MapBody(enc, v)
	case map[int8]bool:
		return writeInt8BoolMapBody(enc, v)
	case map[int8]float32:
		return writeInt8Float32MapBody(enc, v)
	case map[int8]float64:
		return writeInt8Float64MapBody(enc, v)
	case map[int8]interface{}:
		return writeInt8InterfaceMapBody(enc, v)
	case map[int16]string:
		return writeInt16StringMapBody(enc, v)
	case map[int16]int:
		return writeInt16IntMapBody(enc, v)
	case map[int16]int8:
		return writeInt16Int8MapBody(enc, v)
	case map[int16]int16:
		return writeInt16Int16MapBody(enc, v)
	case map[int16]int32:
		return writeInt16Int32MapBody(enc, v)
	case map[int16]int64:
		return writeInt16Int64MapBody(enc, v)
	case map[int16]uint:
		return writeInt16UintMapBody(enc, v)
	case map[int16]uint8:
		return writeInt16Uint8MapBody(enc, v)
	case map[int16]uint16:
		return writeInt16Uint16MapBody(enc, v)
	case map[int16]uint32:
		return writeInt16Uint32MapBody(enc, v)
	case map[int16]uint64:
		return writeInt16Uint64MapBody(enc, v)
	case map[int16]bool:
		return writeInt16BoolMapBody(enc, v)
	case map[int16]float32:
		return writeInt16Float32MapBody(enc, v)
	case map[int16]float64:
		return writeInt16Float64MapBody(enc, v)
	case map[int16]interface{}:
		return writeInt16InterfaceMapBody(enc, v)
	case map[int32]string:
		return writeInt32StringMapBody(enc, v)
	case map[int32]int:
		return writeInt32IntMapBody(enc, v)
	case map[int32]int8:
		return writeInt32Int8MapBody(enc, v)
	case map[int32]int16:
		return writeInt32Int16MapBody(enc, v)
	case map[int32]int32:
		return writeInt32Int32MapBody(enc, v)
	case map[int32]int64:
		return writeInt32Int64MapBody(enc, v)
	case map[int32]uint:
		return writeInt32UintMapBody(enc, v)
	case map[int32]uint8:
		return writeInt32Uint8MapBody(enc, v)
	case map[int32]uint16:
		return writeInt32Uint16MapBody(enc, v)
	case map[int32]uint32:
		return writeInt32Uint32MapBody(enc, v)
	case map[int32]uint64:
		return writeInt32Uint64MapBody(enc, v)
	case map[int32]bool:
		return writeInt32BoolMapBody(enc, v)
	case map[int32]float32:
		return writeInt32Float32MapBody(enc, v)
	case map[int32]float64:
		return writeInt32Float64MapBody(enc, v)
	case map[int32]interface{}:
		return writeInt32InterfaceMapBody(enc, v)
	case map[int64]string:
		return writeInt64StringMapBody(enc, v)
	case map[int64]int:
		return writeInt64IntMapBody(enc, v)
	case map[int64]int8:
		return writeInt64Int8MapBody(enc, v)
	case map[int64]int16:
		return writeInt64Int16MapBody(enc, v)
	case map[int64]int32:
		return writeInt64Int32MapBody(enc, v)
	case map[int64]int64:
		return writeInt64Int64MapBody(enc, v)
	case map[int64]uint:
		return writeInt64UintMapBody(enc, v)
	case map[int64]uint8:
		return writeInt64Uint8MapBody(enc, v)
	case map[int64]uint16:
		return writeInt64Uint16MapBody(enc, v)
	case map[int64]uint32:
		return writeInt64Uint32MapBody(enc, v)
	case map[int64]uint64:
		return writeInt64Uint64MapBody(enc, v)
	case map[int64]bool:
		return writeInt64BoolMapBody(enc, v)
	case map[int64]float32:
		return writeInt64Float32MapBody(enc, v)
	case map[int64]float64:
		return writeInt64Float64MapBody(enc, v)
	case map[int64]interface{}:
		return writeInt64InterfaceMapBody(enc, v)
	case map[uint]string:
		return writeUintStringMapBody(enc, v)
	case map[uint]int:
		return writeUintIntMapBody(enc, v)
	case map[uint]int8:
		return writeUintInt8MapBody(enc, v)
	case map[uint]int16:
		return writeUintInt16MapBody(enc, v)
	case map[uint]int32:
		return writeUintInt32MapBody(enc, v)
	case map[uint]int64:
		return writeUintInt64MapBody(enc, v)
	case map[uint]uint:
		return writeUintUintMapBody(enc, v)
	case map[uint]uint8:
		return writeUintUint8MapBody(enc, v)
	case map[uint]uint16:
		return writeUintUint16MapBody(enc, v)
	case map[uint]uint32:
		return writeUintUint32MapBody(enc, v)
	case map[uint]uint64:
		return writeUintUint64MapBody(enc, v)
	case map[uint]bool:
		return writeUintBoolMapBody(enc, v)
	case map[uint]float32:
		return writeUintFloat32MapBody(enc, v)
	case map[uint]float64:
		return writeUintFloat64MapBody(enc, v)
	case map[uint]interface{}:
		return writeUintInterfaceMapBody(enc, v)
	case map[uint8]string:
		return writeUint8StringMapBody(enc, v)
	case map[uint8]int:
		return writeUint8IntMapBody(enc, v)
	case map[uint8]int8:
		return writeUint8Int8MapBody(enc, v)
	case map[uint8]int16:
		return writeUint8Int16MapBody(enc, v)
	case map[uint8]int32:
		return writeUint8Int32MapBody(enc, v)
	case map[uint8]int64:
		return writeUint8Int64MapBody(enc, v)
	case map[uint8]uint:
		return writeUint8UintMapBody(enc, v)
	case map[uint8]uint8:
		return writeUint8Uint8MapBody(enc, v)
	case map[uint8]uint16:
		return writeUint8Uint16MapBody(enc, v)
	case map[uint8]uint32:
		return writeUint8Uint32MapBody(enc, v)
	case map[uint8]uint64:
		return writeUint8Uint64MapBody(enc, v)
	case map[uint8]bool:
		return writeUint8BoolMapBody(enc, v)
	case map[uint8]float32:
		return writeUint8Float32MapBody(enc, v)
	case map[uint8]float64:
		return writeUint8Float64MapBody(enc, v)
	case map[uint8]interface{}:
		return writeUint8InterfaceMapBody(enc, v)
	case map[uint16]string:
		return writeUint16StringMapBody(enc, v)
	case map[uint16]int:
		return writeUint16IntMapBody(enc, v)
	case map[uint16]int8:
		return writeUint16Int8MapBody(enc, v)
	case map[uint16]int16:
		return writeUint16Int16MapBody(enc, v)
	case map[uint16]int32:
		return writeUint16Int32MapBody(enc, v)
	case map[uint16]int64:
		return writeUint16Int64MapBody(enc, v)
	case map[uint16]uint:
		return writeUint16UintMapBody(enc, v)
	case map[uint16]uint8:
		return writeUint16Uint8MapBody(enc, v)
	case map[uint16]uint16:
		return writeUint16Uint16MapBody(enc, v)
	case map[uint16]uint32:
		return writeUint16Uint32MapBody(enc, v)
	case map[uint16]uint64:
		return writeUint16Uint64MapBody(enc, v)
	case map[uint16]bool:
		return writeUint16BoolMapBody(enc, v)
	case map[uint16]float32:
		return writeUint16Float32MapBody(enc, v)
	case map[uint16]float64:
		return writeUint16Float64MapBody(enc, v)
	case map[uint16]interface{}:
		return writeUint16InterfaceMapBody(enc, v)
	case map[uint32]string:
		return writeUint32StringMapBody(enc, v)
	case map[uint32]int:
		return writeUint32IntMapBody(enc, v)
	case map[uint32]int8:
		return writeUint32Int8MapBody(enc, v)
	case map[uint32]int16:
		return writeUint32Int16MapBody(enc, v)
	case map[uint32]int32:
		return writeUint32Int32MapBody(enc, v)
	case map[uint32]int64:
		return writeUint32Int64MapBody(enc, v)
	case map[uint32]uint:
		return writeUint32UintMapBody(enc, v)
	case map[uint32]uint8:
		return writeUint32Uint8MapBody(enc, v)
	case map[uint32]uint16:
		return writeUint32Uint16MapBody(enc, v)
	case map[uint32]uint32:
		return writeUint32Uint32MapBody(enc, v)
	case map[uint32]uint64:
		return writeUint32Uint64MapBody(enc, v)
	case map[uint32]bool:
		return writeUint32BoolMapBody(enc, v)
	case map[uint32]float32:
		return writeUint32Float32MapBody(enc, v)
	case map[uint32]float64:
		return writeUint32Float64MapBody(enc, v)
	case map[uint32]interface{}:
		return writeUint32InterfaceMapBody(enc, v)
	case map[uint64]string:
		return writeUint64StringMapBody(enc, v)
	case map[uint64]int:
		return writeUint64IntMapBody(enc, v)
	case map[uint64]int8:
		return writeUint64Int8MapBody(enc, v)
	case map[uint64]int16:
		return writeUint64Int16MapBody(enc, v)
	case map[uint64]int32:
		return writeUint64Int32MapBody(enc, v)
	case map[uint64]int64:
		return writeUint64Int64MapBody(enc, v)
	case map[uint64]uint:
		return writeUint64UintMapBody(enc, v)
	case map[uint64]uint8:
		return writeUint64Uint8MapBody(enc, v)
	case map[uint64]uint16:
		return writeUint64Uint16MapBody(enc, v)
	case map[uint64]uint32:
		return writeUint64Uint32MapBody(enc, v)
	case map[uint64]uint64:
		return writeUint64Uint64MapBody(enc, v)
	case map[uint64]bool:
		return writeUint64BoolMapBody(enc, v)
	case map[uint64]float32:
		return writeUint64Float32MapBody(enc, v)
	case map[uint64]float64:
		return writeUint64Float64MapBody(enc, v)
	case map[uint64]interface{}:
		return writeUint64InterfaceMapBody(enc, v)
	case map[float32]string:
		return writeFloat32StringMapBody(enc, v)
	case map[float32]int:
		return writeFloat32IntMapBody(enc, v)
	case map[float32]int8:
		return writeFloat32Int8MapBody(enc, v)
	case map[float32]int16:
		return writeFloat32Int16MapBody(enc, v)
	case map[float32]int32:
		return writeFloat32Int32MapBody(enc, v)
	case map[float32]int64:
		return writeFloat32Int64MapBody(enc, v)
	case map[float32]uint:
		return writeFloat32UintMapBody(enc, v)
	case map[float32]uint8:
		return writeFloat32Uint8MapBody(enc, v)
	case map[float32]uint16:
		return writeFloat32Uint16MapBody(enc, v)
	case map[float32]uint32:
		return writeFloat32Uint32MapBody(enc, v)
	case map[float32]uint64:
		return writeFloat32Uint64MapBody(enc, v)
	case map[float32]bool:
		return writeFloat32BoolMapBody(enc, v)
	case map[float32]float32:
		return writeFloat32Float32MapBody(enc, v)
	case map[float32]float64:
		return writeFloat32Float64MapBody(enc, v)
	case map[float32]interface{}:
		return writeFloat32InterfaceMapBody(enc, v)
	case map[float64]string:
		return writeFloat64StringMapBody(enc, v)
	case map[float64]int:
		return writeFloat64IntMapBody(enc, v)
	case map[float64]int8:
		return writeFloat64Int8MapBody(enc, v)
	case map[float64]int16:
		return writeFloat64Int16MapBody(enc, v)
	case map[float64]int32:
		return writeFloat64Int32MapBody(enc, v)
	case map[float64]int64:
		return writeFloat64Int64MapBody(enc, v)
	case map[float64]uint:
		return writeFloat64UintMapBody(enc, v)
	case map[float64]uint8:
		return writeFloat64Uint8MapBody(enc, v)
	case map[float64]uint16:
		return writeFloat64Uint16MapBody(enc, v)
	case map[float64]uint32:
		return writeFloat64Uint32MapBody(enc, v)
	case map[float64]uint64:
		return writeFloat64Uint64MapBody(enc, v)
	case map[float64]bool:
		return writeFloat64BoolMapBody(enc, v)
	case map[float64]float32:
		return writeFloat64Float32MapBody(enc, v)
	case map[float64]float64:
		return writeFloat64Float64MapBody(enc, v)
	case map[float64]interface{}:
		return writeFloat64InterfaceMapBody(enc, v)
	case map[interface{}]string:
		return writeInterfaceStringMapBody(enc, v)
	case map[interface{}]int:
		return writeInterfaceIntMapBody(enc, v)
	case map[interface{}]int8:
		return writeInterfaceInt8MapBody(enc, v)
	case map[interface{}]int16:
		return writeInterfaceInt16MapBody(enc, v)
	case map[interface{}]int32:
		return writeInterfaceInt32MapBody(enc, v)
	case map[interface{}]int64:
		return writeInterfaceInt64MapBody(enc, v)
	case map[interface{}]uint:
		return writeInterfaceUintMapBody(enc, v)
	case map[interface{}]uint8:
		return writeInterfaceUint8MapBody(enc, v)
	case map[interface{}]uint16:
		return writeInterfaceUint16MapBody(enc, v)
	case map[interface{}]uint32:
		return writeInterfaceUint32MapBody(enc, v)
	case map[interface{}]uint64:
		return writeInterfaceUint64MapBody(enc, v)
	case map[interface{}]bool:
		return writeInterfaceBoolMapBody(enc, v)
	case map[interface{}]float32:
		return writeInterfaceFloat32MapBody(enc, v)
	case map[interface{}]float64:
		return writeInterfaceFloat64MapBody(enc, v)
	case map[interface{}]interface{}:
		return writeInterfaceInterfaceMapBody(enc, v)
	default:
		return writeOtherMapBody(enc, v)
	}
}

func writeStringStringMapBody(enc *Encoder, m map[string]string) (err error) {
	for k, v := range m {
		if err == nil {
			if err = EncodeString(enc, k); err == nil {
				err = EncodeString(enc, v)
			}
		}
	}
	return
}

func writeStringIntMapBody(enc *Encoder, m map[string]int) (err error) {
	for k, v := range m {
		if err == nil {
			if err = EncodeString(enc, k); err == nil {
				err = WriteInt(enc.Writer, v)
			}
		}
	}
	return
}

func writeStringInt8MapBody(enc *Encoder, m map[string]int8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = EncodeString(enc, k); err == nil {
				err = WriteInt8(enc.Writer, v)
			}
		}
	}
	return
}

func writeStringInt16MapBody(enc *Encoder, m map[string]int16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = EncodeString(enc, k); err == nil {
				err = WriteInt16(enc.Writer, v)
			}
		}
	}
	return
}

func writeStringInt32MapBody(enc *Encoder, m map[string]int32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = EncodeString(enc, k); err == nil {
				err = WriteInt32(enc.Writer, v)
			}
		}
	}
	return
}

func writeStringInt64MapBody(enc *Encoder, m map[string]int64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = EncodeString(enc, k); err == nil {
				err = WriteInt64(enc.Writer, v)
			}
		}
	}
	return
}

func writeStringUintMapBody(enc *Encoder, m map[string]uint) (err error) {
	for k, v := range m {
		if err == nil {
			if err = EncodeString(enc, k); err == nil {
				err = WriteUint(enc.Writer, v)
			}
		}
	}
	return
}

func writeStringUint8MapBody(enc *Encoder, m map[string]uint8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = EncodeString(enc, k); err == nil {
				err = WriteUint8(enc.Writer, v)
			}
		}
	}
	return
}

func writeStringUint16MapBody(enc *Encoder, m map[string]uint16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = EncodeString(enc, k); err == nil {
				err = WriteUint16(enc.Writer, v)
			}
		}
	}
	return
}

func writeStringUint32MapBody(enc *Encoder, m map[string]uint32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = EncodeString(enc, k); err == nil {
				err = WriteUint32(enc.Writer, v)
			}
		}
	}
	return
}

func writeStringUint64MapBody(enc *Encoder, m map[string]uint64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = EncodeString(enc, k); err == nil {
				err = WriteUint64(enc.Writer, v)
			}
		}
	}
	return
}

func writeStringBoolMapBody(enc *Encoder, m map[string]bool) (err error) {
	for k, v := range m {
		if err == nil {
			if err = EncodeString(enc, k); err == nil {
				err = WriteBool(enc.Writer, v)
			}
		}
	}
	return
}

func writeStringFloat32MapBody(enc *Encoder, m map[string]float32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = EncodeString(enc, k); err == nil {
				err = WriteFloat32(enc.Writer, v)
			}
		}
	}
	return
}

func writeStringFloat64MapBody(enc *Encoder, m map[string]float64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = EncodeString(enc, k); err == nil {
				err = WriteFloat64(enc.Writer, v)
			}
		}
	}
	return
}

func writeStringInterfaceMapBody(enc *Encoder, m map[string]interface{}) (err error) {
	for k, v := range m {
		if err == nil {
			if err = EncodeString(enc, k); err == nil {
				err = enc.Encode(v)
			}
		}
	}
	return
}

func writeIntStringMapBody(enc *Encoder, m map[int]string) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt(enc.Writer, k); err == nil {
				err = EncodeString(enc, v)
			}
		}
	}
	return
}

func writeIntIntMapBody(enc *Encoder, m map[int]int) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt(enc.Writer, k); err == nil {
				err = WriteInt(enc.Writer, v)
			}
		}
	}
	return
}

func writeIntInt8MapBody(enc *Encoder, m map[int]int8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt(enc.Writer, k); err == nil {
				err = WriteInt8(enc.Writer, v)
			}
		}
	}
	return
}

func writeIntInt16MapBody(enc *Encoder, m map[int]int16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt(enc.Writer, k); err == nil {
				err = WriteInt16(enc.Writer, v)
			}
		}
	}
	return
}

func writeIntInt32MapBody(enc *Encoder, m map[int]int32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt(enc.Writer, k); err == nil {
				err = WriteInt32(enc.Writer, v)
			}
		}
	}
	return
}

func writeIntInt64MapBody(enc *Encoder, m map[int]int64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt(enc.Writer, k); err == nil {
				err = WriteInt64(enc.Writer, v)
			}
		}
	}
	return
}

func writeIntUintMapBody(enc *Encoder, m map[int]uint) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt(enc.Writer, k); err == nil {
				err = WriteUint(enc.Writer, v)
			}
		}
	}
	return
}

func writeIntUint8MapBody(enc *Encoder, m map[int]uint8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt(enc.Writer, k); err == nil {
				err = WriteUint8(enc.Writer, v)
			}
		}
	}
	return
}

func writeIntUint16MapBody(enc *Encoder, m map[int]uint16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt(enc.Writer, k); err == nil {
				err = WriteUint16(enc.Writer, v)
			}
		}
	}
	return
}

func writeIntUint32MapBody(enc *Encoder, m map[int]uint32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt(enc.Writer, k); err == nil {
				err = WriteUint32(enc.Writer, v)
			}
		}
	}
	return
}

func writeIntUint64MapBody(enc *Encoder, m map[int]uint64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt(enc.Writer, k); err == nil {
				err = WriteUint64(enc.Writer, v)
			}
		}
	}
	return
}

func writeIntBoolMapBody(enc *Encoder, m map[int]bool) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt(enc.Writer, k); err == nil {
				err = WriteBool(enc.Writer, v)
			}
		}
	}
	return
}

func writeIntFloat32MapBody(enc *Encoder, m map[int]float32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt(enc.Writer, k); err == nil {
				err = WriteFloat32(enc.Writer, v)
			}
		}
	}
	return
}

func writeIntFloat64MapBody(enc *Encoder, m map[int]float64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt(enc.Writer, k); err == nil {
				err = WriteFloat64(enc.Writer, v)
			}
		}
	}
	return
}

func writeIntInterfaceMapBody(enc *Encoder, m map[int]interface{}) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt(enc.Writer, k); err == nil {
				err = enc.Encode(v)
			}
		}
	}
	return
}

func writeInt8StringMapBody(enc *Encoder, m map[int8]string) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt8(enc.Writer, k); err == nil {
				err = EncodeString(enc, v)
			}
		}
	}
	return
}

func writeInt8IntMapBody(enc *Encoder, m map[int8]int) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt8(enc.Writer, k); err == nil {
				err = WriteInt(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt8Int8MapBody(enc *Encoder, m map[int8]int8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt8(enc.Writer, k); err == nil {
				err = WriteInt8(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt8Int16MapBody(enc *Encoder, m map[int8]int16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt8(enc.Writer, k); err == nil {
				err = WriteInt16(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt8Int32MapBody(enc *Encoder, m map[int8]int32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt8(enc.Writer, k); err == nil {
				err = WriteInt32(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt8Int64MapBody(enc *Encoder, m map[int8]int64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt8(enc.Writer, k); err == nil {
				err = WriteInt64(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt8UintMapBody(enc *Encoder, m map[int8]uint) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt8(enc.Writer, k); err == nil {
				err = WriteUint(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt8Uint8MapBody(enc *Encoder, m map[int8]uint8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt8(enc.Writer, k); err == nil {
				err = WriteUint8(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt8Uint16MapBody(enc *Encoder, m map[int8]uint16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt8(enc.Writer, k); err == nil {
				err = WriteUint16(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt8Uint32MapBody(enc *Encoder, m map[int8]uint32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt8(enc.Writer, k); err == nil {
				err = WriteUint32(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt8Uint64MapBody(enc *Encoder, m map[int8]uint64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt8(enc.Writer, k); err == nil {
				err = WriteUint64(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt8BoolMapBody(enc *Encoder, m map[int8]bool) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt8(enc.Writer, k); err == nil {
				err = WriteBool(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt8Float32MapBody(enc *Encoder, m map[int8]float32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt8(enc.Writer, k); err == nil {
				err = WriteFloat32(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt8Float64MapBody(enc *Encoder, m map[int8]float64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt8(enc.Writer, k); err == nil {
				err = WriteFloat64(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt8InterfaceMapBody(enc *Encoder, m map[int8]interface{}) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt8(enc.Writer, k); err == nil {
				err = enc.Encode(v)
			}
		}
	}
	return
}

func writeInt16StringMapBody(enc *Encoder, m map[int16]string) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt16(enc.Writer, k); err == nil {
				err = EncodeString(enc, v)
			}
		}
	}
	return
}

func writeInt16IntMapBody(enc *Encoder, m map[int16]int) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt16(enc.Writer, k); err == nil {
				err = WriteInt(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt16Int8MapBody(enc *Encoder, m map[int16]int8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt16(enc.Writer, k); err == nil {
				err = WriteInt8(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt16Int16MapBody(enc *Encoder, m map[int16]int16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt16(enc.Writer, k); err == nil {
				err = WriteInt16(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt16Int32MapBody(enc *Encoder, m map[int16]int32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt16(enc.Writer, k); err == nil {
				err = WriteInt32(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt16Int64MapBody(enc *Encoder, m map[int16]int64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt16(enc.Writer, k); err == nil {
				err = WriteInt64(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt16UintMapBody(enc *Encoder, m map[int16]uint) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt16(enc.Writer, k); err == nil {
				err = WriteUint(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt16Uint8MapBody(enc *Encoder, m map[int16]uint8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt16(enc.Writer, k); err == nil {
				err = WriteUint8(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt16Uint16MapBody(enc *Encoder, m map[int16]uint16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt16(enc.Writer, k); err == nil {
				err = WriteUint16(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt16Uint32MapBody(enc *Encoder, m map[int16]uint32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt16(enc.Writer, k); err == nil {
				err = WriteUint32(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt16Uint64MapBody(enc *Encoder, m map[int16]uint64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt16(enc.Writer, k); err == nil {
				err = WriteUint64(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt16BoolMapBody(enc *Encoder, m map[int16]bool) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt16(enc.Writer, k); err == nil {
				err = WriteBool(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt16Float32MapBody(enc *Encoder, m map[int16]float32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt16(enc.Writer, k); err == nil {
				err = WriteFloat32(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt16Float64MapBody(enc *Encoder, m map[int16]float64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt16(enc.Writer, k); err == nil {
				err = WriteFloat64(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt16InterfaceMapBody(enc *Encoder, m map[int16]interface{}) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt16(enc.Writer, k); err == nil {
				err = enc.Encode(v)
			}
		}
	}
	return
}

func writeInt32StringMapBody(enc *Encoder, m map[int32]string) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt32(enc.Writer, k); err == nil {
				err = EncodeString(enc, v)
			}
		}
	}
	return
}

func writeInt32IntMapBody(enc *Encoder, m map[int32]int) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt32(enc.Writer, k); err == nil {
				err = WriteInt(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt32Int8MapBody(enc *Encoder, m map[int32]int8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt32(enc.Writer, k); err == nil {
				err = WriteInt8(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt32Int16MapBody(enc *Encoder, m map[int32]int16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt32(enc.Writer, k); err == nil {
				err = WriteInt16(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt32Int32MapBody(enc *Encoder, m map[int32]int32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt32(enc.Writer, k); err == nil {
				err = WriteInt32(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt32Int64MapBody(enc *Encoder, m map[int32]int64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt32(enc.Writer, k); err == nil {
				err = WriteInt64(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt32UintMapBody(enc *Encoder, m map[int32]uint) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt32(enc.Writer, k); err == nil {
				err = WriteUint(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt32Uint8MapBody(enc *Encoder, m map[int32]uint8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt32(enc.Writer, k); err == nil {
				err = WriteUint8(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt32Uint16MapBody(enc *Encoder, m map[int32]uint16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt32(enc.Writer, k); err == nil {
				err = WriteUint16(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt32Uint32MapBody(enc *Encoder, m map[int32]uint32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt32(enc.Writer, k); err == nil {
				err = WriteUint32(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt32Uint64MapBody(enc *Encoder, m map[int32]uint64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt32(enc.Writer, k); err == nil {
				err = WriteUint64(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt32BoolMapBody(enc *Encoder, m map[int32]bool) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt32(enc.Writer, k); err == nil {
				err = WriteBool(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt32Float32MapBody(enc *Encoder, m map[int32]float32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt32(enc.Writer, k); err == nil {
				err = WriteFloat32(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt32Float64MapBody(enc *Encoder, m map[int32]float64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt32(enc.Writer, k); err == nil {
				err = WriteFloat64(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt32InterfaceMapBody(enc *Encoder, m map[int32]interface{}) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt32(enc.Writer, k); err == nil {
				err = enc.Encode(v)
			}
		}
	}
	return
}

func writeInt64StringMapBody(enc *Encoder, m map[int64]string) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt64(enc.Writer, k); err == nil {
				err = EncodeString(enc, v)
			}
		}
	}
	return
}

func writeInt64IntMapBody(enc *Encoder, m map[int64]int) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt64(enc.Writer, k); err == nil {
				err = WriteInt(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt64Int8MapBody(enc *Encoder, m map[int64]int8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt64(enc.Writer, k); err == nil {
				err = WriteInt8(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt64Int16MapBody(enc *Encoder, m map[int64]int16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt64(enc.Writer, k); err == nil {
				err = WriteInt16(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt64Int32MapBody(enc *Encoder, m map[int64]int32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt64(enc.Writer, k); err == nil {
				err = WriteInt32(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt64Int64MapBody(enc *Encoder, m map[int64]int64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt64(enc.Writer, k); err == nil {
				err = WriteInt64(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt64UintMapBody(enc *Encoder, m map[int64]uint) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt64(enc.Writer, k); err == nil {
				err = WriteUint(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt64Uint8MapBody(enc *Encoder, m map[int64]uint8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt64(enc.Writer, k); err == nil {
				err = WriteUint8(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt64Uint16MapBody(enc *Encoder, m map[int64]uint16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt64(enc.Writer, k); err == nil {
				err = WriteUint16(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt64Uint32MapBody(enc *Encoder, m map[int64]uint32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt64(enc.Writer, k); err == nil {
				err = WriteUint32(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt64Uint64MapBody(enc *Encoder, m map[int64]uint64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt64(enc.Writer, k); err == nil {
				err = WriteUint64(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt64BoolMapBody(enc *Encoder, m map[int64]bool) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt64(enc.Writer, k); err == nil {
				err = WriteBool(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt64Float32MapBody(enc *Encoder, m map[int64]float32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt64(enc.Writer, k); err == nil {
				err = WriteFloat32(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt64Float64MapBody(enc *Encoder, m map[int64]float64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt64(enc.Writer, k); err == nil {
				err = WriteFloat64(enc.Writer, v)
			}
		}
	}
	return
}

func writeInt64InterfaceMapBody(enc *Encoder, m map[int64]interface{}) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteInt64(enc.Writer, k); err == nil {
				err = enc.Encode(v)
			}
		}
	}
	return
}

func writeUintStringMapBody(enc *Encoder, m map[uint]string) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint(enc.Writer, k); err == nil {
				err = EncodeString(enc, v)
			}
		}
	}
	return
}

func writeUintIntMapBody(enc *Encoder, m map[uint]int) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint(enc.Writer, k); err == nil {
				err = WriteInt(enc.Writer, v)
			}
		}
	}
	return
}

func writeUintInt8MapBody(enc *Encoder, m map[uint]int8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint(enc.Writer, k); err == nil {
				err = WriteInt8(enc.Writer, v)
			}
		}
	}
	return
}

func writeUintInt16MapBody(enc *Encoder, m map[uint]int16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint(enc.Writer, k); err == nil {
				err = WriteInt16(enc.Writer, v)
			}
		}
	}
	return
}

func writeUintInt32MapBody(enc *Encoder, m map[uint]int32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint(enc.Writer, k); err == nil {
				err = WriteInt32(enc.Writer, v)
			}
		}
	}
	return
}

func writeUintInt64MapBody(enc *Encoder, m map[uint]int64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint(enc.Writer, k); err == nil {
				err = WriteInt64(enc.Writer, v)
			}
		}
	}
	return
}

func writeUintUintMapBody(enc *Encoder, m map[uint]uint) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint(enc.Writer, k); err == nil {
				err = WriteUint(enc.Writer, v)
			}
		}
	}
	return
}

func writeUintUint8MapBody(enc *Encoder, m map[uint]uint8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint(enc.Writer, k); err == nil {
				err = WriteUint8(enc.Writer, v)
			}
		}
	}
	return
}

func writeUintUint16MapBody(enc *Encoder, m map[uint]uint16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint(enc.Writer, k); err == nil {
				err = WriteUint16(enc.Writer, v)
			}
		}
	}
	return
}

func writeUintUint32MapBody(enc *Encoder, m map[uint]uint32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint(enc.Writer, k); err == nil {
				err = WriteUint32(enc.Writer, v)
			}
		}
	}
	return
}

func writeUintUint64MapBody(enc *Encoder, m map[uint]uint64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint(enc.Writer, k); err == nil {
				err = WriteUint64(enc.Writer, v)
			}
		}
	}
	return
}

func writeUintBoolMapBody(enc *Encoder, m map[uint]bool) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint(enc.Writer, k); err == nil {
				err = WriteBool(enc.Writer, v)
			}
		}
	}
	return
}

func writeUintFloat32MapBody(enc *Encoder, m map[uint]float32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint(enc.Writer, k); err == nil {
				err = WriteFloat32(enc.Writer, v)
			}
		}
	}
	return
}

func writeUintFloat64MapBody(enc *Encoder, m map[uint]float64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint(enc.Writer, k); err == nil {
				err = WriteFloat64(enc.Writer, v)
			}
		}
	}
	return
}

func writeUintInterfaceMapBody(enc *Encoder, m map[uint]interface{}) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint(enc.Writer, k); err == nil {
				err = enc.Encode(v)
			}
		}
	}
	return
}
func writeUint8StringMapBody(enc *Encoder, m map[uint8]string) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint8(enc.Writer, k); err == nil {
				err = EncodeString(enc, v)
			}
		}
	}
	return
}

func writeUint8IntMapBody(enc *Encoder, m map[uint8]int) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint8(enc.Writer, k); err == nil {
				err = WriteInt(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint8Int8MapBody(enc *Encoder, m map[uint8]int8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint8(enc.Writer, k); err == nil {
				err = WriteInt8(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint8Int16MapBody(enc *Encoder, m map[uint8]int16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint8(enc.Writer, k); err == nil {
				err = WriteInt16(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint8Int32MapBody(enc *Encoder, m map[uint8]int32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint8(enc.Writer, k); err == nil {
				err = WriteInt32(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint8Int64MapBody(enc *Encoder, m map[uint8]int64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint8(enc.Writer, k); err == nil {
				err = WriteInt64(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint8UintMapBody(enc *Encoder, m map[uint8]uint) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint8(enc.Writer, k); err == nil {
				err = WriteUint(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint8Uint8MapBody(enc *Encoder, m map[uint8]uint8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint8(enc.Writer, k); err == nil {
				err = WriteUint8(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint8Uint16MapBody(enc *Encoder, m map[uint8]uint16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint8(enc.Writer, k); err == nil {
				err = WriteUint16(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint8Uint32MapBody(enc *Encoder, m map[uint8]uint32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint8(enc.Writer, k); err == nil {
				err = WriteUint32(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint8Uint64MapBody(enc *Encoder, m map[uint8]uint64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint8(enc.Writer, k); err == nil {
				err = WriteUint64(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint8BoolMapBody(enc *Encoder, m map[uint8]bool) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint8(enc.Writer, k); err == nil {
				err = WriteBool(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint8Float32MapBody(enc *Encoder, m map[uint8]float32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint8(enc.Writer, k); err == nil {
				err = WriteFloat32(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint8Float64MapBody(enc *Encoder, m map[uint8]float64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint8(enc.Writer, k); err == nil {
				err = WriteFloat64(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint8InterfaceMapBody(enc *Encoder, m map[uint8]interface{}) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint8(enc.Writer, k); err == nil {
				err = enc.Encode(v)
			}
		}
	}
	return
}

func writeUint16StringMapBody(enc *Encoder, m map[uint16]string) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint16(enc.Writer, k); err == nil {
				err = EncodeString(enc, v)
			}
		}
	}
	return
}

func writeUint16IntMapBody(enc *Encoder, m map[uint16]int) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint16(enc.Writer, k); err == nil {
				err = WriteInt(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint16Int8MapBody(enc *Encoder, m map[uint16]int8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint16(enc.Writer, k); err == nil {
				err = WriteInt8(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint16Int16MapBody(enc *Encoder, m map[uint16]int16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint16(enc.Writer, k); err == nil {
				err = WriteInt16(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint16Int32MapBody(enc *Encoder, m map[uint16]int32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint16(enc.Writer, k); err == nil {
				err = WriteInt32(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint16Int64MapBody(enc *Encoder, m map[uint16]int64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint16(enc.Writer, k); err == nil {
				err = WriteInt64(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint16UintMapBody(enc *Encoder, m map[uint16]uint) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint16(enc.Writer, k); err == nil {
				err = WriteUint(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint16Uint8MapBody(enc *Encoder, m map[uint16]uint8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint16(enc.Writer, k); err == nil {
				err = WriteUint8(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint16Uint16MapBody(enc *Encoder, m map[uint16]uint16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint16(enc.Writer, k); err == nil {
				err = WriteUint16(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint16Uint32MapBody(enc *Encoder, m map[uint16]uint32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint16(enc.Writer, k); err == nil {
				err = WriteUint32(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint16Uint64MapBody(enc *Encoder, m map[uint16]uint64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint16(enc.Writer, k); err == nil {
				err = WriteUint64(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint16BoolMapBody(enc *Encoder, m map[uint16]bool) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint16(enc.Writer, k); err == nil {
				err = WriteBool(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint16Float32MapBody(enc *Encoder, m map[uint16]float32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint16(enc.Writer, k); err == nil {
				err = WriteFloat32(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint16Float64MapBody(enc *Encoder, m map[uint16]float64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint16(enc.Writer, k); err == nil {
				err = WriteFloat64(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint16InterfaceMapBody(enc *Encoder, m map[uint16]interface{}) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint16(enc.Writer, k); err == nil {
				err = enc.Encode(v)
			}
		}
	}
	return
}

func writeUint32StringMapBody(enc *Encoder, m map[uint32]string) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint32(enc.Writer, k); err == nil {
				err = EncodeString(enc, v)
			}
		}
	}
	return
}

func writeUint32IntMapBody(enc *Encoder, m map[uint32]int) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint32(enc.Writer, k); err == nil {
				err = WriteInt(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint32Int8MapBody(enc *Encoder, m map[uint32]int8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint32(enc.Writer, k); err == nil {
				err = WriteInt8(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint32Int16MapBody(enc *Encoder, m map[uint32]int16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint32(enc.Writer, k); err == nil {
				err = WriteInt16(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint32Int32MapBody(enc *Encoder, m map[uint32]int32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint32(enc.Writer, k); err == nil {
				err = WriteInt32(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint32Int64MapBody(enc *Encoder, m map[uint32]int64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint32(enc.Writer, k); err == nil {
				err = WriteInt64(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint32UintMapBody(enc *Encoder, m map[uint32]uint) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint32(enc.Writer, k); err == nil {
				err = WriteUint(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint32Uint8MapBody(enc *Encoder, m map[uint32]uint8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint32(enc.Writer, k); err == nil {
				err = WriteUint8(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint32Uint16MapBody(enc *Encoder, m map[uint32]uint16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint32(enc.Writer, k); err == nil {
				err = WriteUint16(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint32Uint32MapBody(enc *Encoder, m map[uint32]uint32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint32(enc.Writer, k); err == nil {
				err = WriteUint32(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint32Uint64MapBody(enc *Encoder, m map[uint32]uint64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint32(enc.Writer, k); err == nil {
				err = WriteUint64(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint32BoolMapBody(enc *Encoder, m map[uint32]bool) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint32(enc.Writer, k); err == nil {
				err = WriteBool(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint32Float32MapBody(enc *Encoder, m map[uint32]float32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint32(enc.Writer, k); err == nil {
				err = WriteFloat32(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint32Float64MapBody(enc *Encoder, m map[uint32]float64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint32(enc.Writer, k); err == nil {
				err = WriteFloat64(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint32InterfaceMapBody(enc *Encoder, m map[uint32]interface{}) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint32(enc.Writer, k); err == nil {
				err = enc.Encode(v)
			}
		}
	}
	return
}

func writeUint64StringMapBody(enc *Encoder, m map[uint64]string) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint64(enc.Writer, k); err == nil {
				err = EncodeString(enc, v)
			}
		}
	}
	return
}

func writeUint64IntMapBody(enc *Encoder, m map[uint64]int) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint64(enc.Writer, k); err == nil {
				err = WriteInt(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint64Int8MapBody(enc *Encoder, m map[uint64]int8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint64(enc.Writer, k); err == nil {
				err = WriteInt8(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint64Int16MapBody(enc *Encoder, m map[uint64]int16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint64(enc.Writer, k); err == nil {
				err = WriteInt16(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint64Int32MapBody(enc *Encoder, m map[uint64]int32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint64(enc.Writer, k); err == nil {
				err = WriteInt32(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint64Int64MapBody(enc *Encoder, m map[uint64]int64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint64(enc.Writer, k); err == nil {
				err = WriteInt64(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint64UintMapBody(enc *Encoder, m map[uint64]uint) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint64(enc.Writer, k); err == nil {
				err = WriteUint(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint64Uint8MapBody(enc *Encoder, m map[uint64]uint8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint64(enc.Writer, k); err == nil {
				err = WriteUint8(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint64Uint16MapBody(enc *Encoder, m map[uint64]uint16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint64(enc.Writer, k); err == nil {
				err = WriteUint16(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint64Uint32MapBody(enc *Encoder, m map[uint64]uint32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint64(enc.Writer, k); err == nil {
				err = WriteUint32(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint64Uint64MapBody(enc *Encoder, m map[uint64]uint64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint64(enc.Writer, k); err == nil {
				err = WriteUint64(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint64BoolMapBody(enc *Encoder, m map[uint64]bool) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint64(enc.Writer, k); err == nil {
				err = WriteBool(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint64Float32MapBody(enc *Encoder, m map[uint64]float32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint64(enc.Writer, k); err == nil {
				err = WriteFloat32(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint64Float64MapBody(enc *Encoder, m map[uint64]float64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint64(enc.Writer, k); err == nil {
				err = WriteFloat64(enc.Writer, v)
			}
		}
	}
	return
}

func writeUint64InterfaceMapBody(enc *Encoder, m map[uint64]interface{}) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteUint64(enc.Writer, k); err == nil {
				err = enc.Encode(v)
			}
		}
	}
	return
}

func writeFloat32StringMapBody(enc *Encoder, m map[float32]string) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat32(enc.Writer, k); err == nil {
				err = EncodeString(enc, v)
			}
		}
	}
	return
}

func writeFloat32IntMapBody(enc *Encoder, m map[float32]int) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat32(enc.Writer, k); err == nil {
				err = WriteInt(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat32Int8MapBody(enc *Encoder, m map[float32]int8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat32(enc.Writer, k); err == nil {
				err = WriteInt8(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat32Int16MapBody(enc *Encoder, m map[float32]int16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat32(enc.Writer, k); err == nil {
				err = WriteInt16(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat32Int32MapBody(enc *Encoder, m map[float32]int32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat32(enc.Writer, k); err == nil {
				err = WriteInt32(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat32Int64MapBody(enc *Encoder, m map[float32]int64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat32(enc.Writer, k); err == nil {
				err = WriteInt64(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat32UintMapBody(enc *Encoder, m map[float32]uint) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat32(enc.Writer, k); err == nil {
				err = WriteUint(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat32Uint8MapBody(enc *Encoder, m map[float32]uint8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat32(enc.Writer, k); err == nil {
				err = WriteUint8(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat32Uint16MapBody(enc *Encoder, m map[float32]uint16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat32(enc.Writer, k); err == nil {
				err = WriteUint16(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat32Uint32MapBody(enc *Encoder, m map[float32]uint32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat32(enc.Writer, k); err == nil {
				err = WriteUint32(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat32Uint64MapBody(enc *Encoder, m map[float32]uint64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat32(enc.Writer, k); err == nil {
				err = WriteUint64(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat32BoolMapBody(enc *Encoder, m map[float32]bool) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat32(enc.Writer, k); err == nil {
				err = WriteBool(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat32Float32MapBody(enc *Encoder, m map[float32]float32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat32(enc.Writer, k); err == nil {
				err = WriteFloat32(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat32Float64MapBody(enc *Encoder, m map[float32]float64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat32(enc.Writer, k); err == nil {
				err = WriteFloat64(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat32InterfaceMapBody(enc *Encoder, m map[float32]interface{}) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat32(enc.Writer, k); err == nil {
				err = enc.Encode(v)
			}
		}
	}
	return
}

func writeFloat64StringMapBody(enc *Encoder, m map[float64]string) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat64(enc.Writer, k); err == nil {
				err = EncodeString(enc, v)
			}
		}
	}
	return
}

func writeFloat64IntMapBody(enc *Encoder, m map[float64]int) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat64(enc.Writer, k); err == nil {
				err = WriteInt(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat64Int8MapBody(enc *Encoder, m map[float64]int8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat64(enc.Writer, k); err == nil {
				err = WriteInt8(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat64Int16MapBody(enc *Encoder, m map[float64]int16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat64(enc.Writer, k); err == nil {
				err = WriteInt16(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat64Int32MapBody(enc *Encoder, m map[float64]int32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat64(enc.Writer, k); err == nil {
				err = WriteInt32(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat64Int64MapBody(enc *Encoder, m map[float64]int64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat64(enc.Writer, k); err == nil {
				err = WriteInt64(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat64UintMapBody(enc *Encoder, m map[float64]uint) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat64(enc.Writer, k); err == nil {
				err = WriteUint(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat64Uint8MapBody(enc *Encoder, m map[float64]uint8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat64(enc.Writer, k); err == nil {
				err = WriteUint8(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat64Uint16MapBody(enc *Encoder, m map[float64]uint16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat64(enc.Writer, k); err == nil {
				err = WriteUint16(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat64Uint32MapBody(enc *Encoder, m map[float64]uint32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat64(enc.Writer, k); err == nil {
				err = WriteUint32(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat64Uint64MapBody(enc *Encoder, m map[float64]uint64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat64(enc.Writer, k); err == nil {
				err = WriteUint64(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat64BoolMapBody(enc *Encoder, m map[float64]bool) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat64(enc.Writer, k); err == nil {
				err = WriteBool(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat64Float32MapBody(enc *Encoder, m map[float64]float32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat64(enc.Writer, k); err == nil {
				err = WriteFloat32(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat64Float64MapBody(enc *Encoder, m map[float64]float64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat64(enc.Writer, k); err == nil {
				err = WriteFloat64(enc.Writer, v)
			}
		}
	}
	return
}

func writeFloat64InterfaceMapBody(enc *Encoder, m map[float64]interface{}) (err error) {
	for k, v := range m {
		if err == nil {
			if err = WriteFloat64(enc.Writer, k); err == nil {
				err = enc.Encode(v)
			}
		}
	}
	return
}

func writeInterfaceStringMapBody(enc *Encoder, m map[interface{}]string) (err error) {
	for k, v := range m {
		if err == nil {
			if err = enc.Encode(k); err == nil {
				err = EncodeString(enc, v)
			}
		}
	}
	return
}

func writeInterfaceIntMapBody(enc *Encoder, m map[interface{}]int) (err error) {
	for k, v := range m {
		if err == nil {
			if err = enc.Encode(k); err == nil {
				err = WriteInt(enc.Writer, v)
			}
		}
	}
	return
}

func writeInterfaceInt8MapBody(enc *Encoder, m map[interface{}]int8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = enc.Encode(k); err == nil {
				err = WriteInt8(enc.Writer, v)
			}
		}
	}
	return
}

func writeInterfaceInt16MapBody(enc *Encoder, m map[interface{}]int16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = enc.Encode(k); err == nil {
				err = WriteInt16(enc.Writer, v)
			}
		}
	}
	return
}

func writeInterfaceInt32MapBody(enc *Encoder, m map[interface{}]int32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = enc.Encode(k); err == nil {
				err = WriteInt32(enc.Writer, v)
			}
		}
	}
	return
}

func writeInterfaceInt64MapBody(enc *Encoder, m map[interface{}]int64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = enc.Encode(k); err == nil {
				err = WriteInt64(enc.Writer, v)
			}
		}
	}
	return
}

func writeInterfaceUintMapBody(enc *Encoder, m map[interface{}]uint) (err error) {
	for k, v := range m {
		if err == nil {
			if err = enc.Encode(k); err == nil {
				err = WriteUint(enc.Writer, v)
			}
		}
	}
	return
}

func writeInterfaceUint8MapBody(enc *Encoder, m map[interface{}]uint8) (err error) {
	for k, v := range m {
		if err == nil {
			if err = enc.Encode(k); err == nil {
				err = WriteUint8(enc.Writer, v)
			}
		}
	}
	return
}

func writeInterfaceUint16MapBody(enc *Encoder, m map[interface{}]uint16) (err error) {
	for k, v := range m {
		if err == nil {
			if err = enc.Encode(k); err == nil {
				err = WriteUint16(enc.Writer, v)
			}
		}
	}
	return
}

func writeInterfaceUint32MapBody(enc *Encoder, m map[interface{}]uint32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = enc.Encode(k); err == nil {
				err = WriteUint32(enc.Writer, v)
			}
		}
	}
	return
}

func writeInterfaceUint64MapBody(enc *Encoder, m map[interface{}]uint64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = enc.Encode(k); err == nil {
				err = WriteUint64(enc.Writer, v)
			}
		}
	}
	return
}

func writeInterfaceBoolMapBody(enc *Encoder, m map[interface{}]bool) (err error) {
	for k, v := range m {
		if err == nil {
			if err = enc.Encode(k); err == nil {
				err = WriteBool(enc.Writer, v)
			}
		}
	}
	return
}

func writeInterfaceFloat32MapBody(enc *Encoder, m map[interface{}]float32) (err error) {
	for k, v := range m {
		if err == nil {
			if err = enc.Encode(k); err == nil {
				err = WriteFloat32(enc.Writer, v)
			}
		}
	}
	return
}

func writeInterfaceFloat64MapBody(enc *Encoder, m map[interface{}]float64) (err error) {
	for k, v := range m {
		if err == nil {
			if err = enc.Encode(k); err == nil {
				err = WriteFloat64(enc.Writer, v)
			}
		}
	}
	return
}

func writeInterfaceInterfaceMapBody(enc *Encoder, m map[interface{}]interface{}) (err error) {
	for k, v := range m {
		if err == nil {
			if err = enc.Encode(k); err == nil {
				err = enc.Encode(v)
			}
		}
	}
	return
}

func writeOtherMapBody(enc *Encoder, v interface{}) (err error) {
	mapType := reflect2.TypeOf(v).(*reflect2.UnsafeMapType)
	p := reflect2.PtrOf(v)
	iter := mapType.UnsafeIterate(unsafe.Pointer(&p))
	kt := mapType.Key()
	vt := mapType.Elem()
	for iter.HasNext() && err == nil {
		kp, vp := iter.UnsafeNext()
		if err = enc.Encode(kt.UnsafeIndirect(kp)); err == nil {
			err = enc.Encode(vt.UnsafeIndirect(vp))
		}
	}
	return
}
