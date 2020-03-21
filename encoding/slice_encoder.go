/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/slice_encoder.go                                |
|                                                          |
| LastModified: Mar 17, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"reflect"

	"github.com/modern-go/reflect2"
)

// sliceEncoder is the implementation of ValueEncoder for *slice.
type sliceEncoder struct{}

var slcenc sliceEncoder

func (valenc sliceEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	var ok bool
	if ok, err = enc.WriteReference(v); !ok && err == nil {
		err = valenc.Write(enc, v)
	}
	return
}

func (sliceEncoder) Write(enc *Encoder, v interface{}) (err error) {
	enc.SetReference(v)
	return writeSlice(enc, reflect.ValueOf(v).Elem().Interface())
}

// WriteSlice to encoder
func WriteSlice(enc *Encoder, v interface{}) (err error) {
	enc.AddReferenceCount(1)
	return writeSlice(enc, v)
}

var emptySlice = []byte{TagList, TagOpenbrace, TagClosebrace}

func writeSlice(enc *Encoder, v interface{}) (err error) {
	if bytes, ok := v.([]byte); ok {
		return writeBytes(enc.writer, bytes)
	}
	count := (*reflect.SliceHeader)(reflect2.PtrOf(v)).Len
	if count == 0 {
		_, err = enc.writer.Write(emptySlice)
		return
	}
	if err = WriteHead(enc, count, TagList); err == nil {
		if err = writeSliceBody(enc, v); err == nil {
			err = WriteFoot(enc)
		}
	}
	return
}

func writeSliceBody(enc *Encoder, v interface{}) error {
	switch v := v.(type) {
	case []uint16:
		return writeUint16SliceBody(enc, v)
	case []uint32:
		return writeUint32SliceBody(enc, v)
	case []uint64:
		return writeUint64SliceBody(enc, v)
	case []uint:
		return writeUintSliceBody(enc, v)
	case []int8:
		return writeInt8SliceBody(enc, v)
	case []int16:
		return writeInt16SliceBody(enc, v)
	case []int32:
		return writeInt32SliceBody(enc, v)
	case []int64:
		return writeInt64SliceBody(enc, v)
	case []int:
		return writeIntSliceBody(enc, v)
	case []bool:
		return writeBoolSliceBody(enc, v)
	case []float32:
		return writeFloat32SliceBody(enc, v)
	case []float64:
		return writeFloat64SliceBody(enc, v)
	case []complex64:
		return writeComplex64SliceBody(enc, v)
	case []complex128:
		return writeComplex128SliceBody(enc, v)
	case []string:
		return writeStringSliceBody(enc, v)
	case []interface{}:
		return writeInterfaceSliceBody(enc, v)
	case [][]uint16:
		return write2dUint16SliceBody(enc, v)
	case [][]uint32:
		return write2dUint32SliceBody(enc, v)
	case [][]uint64:
		return write2dUint64SliceBody(enc, v)
	case [][]uint:
		return write2dUintSliceBody(enc, v)
	case [][]int8:
		return write2dInt8SliceBody(enc, v)
	case [][]int16:
		return write2dInt16SliceBody(enc, v)
	case [][]int32:
		return write2dInt32SliceBody(enc, v)
	case [][]int64:
		return write2dInt64SliceBody(enc, v)
	case [][]int:
		return write2dIntSliceBody(enc, v)
	case [][]bool:
		return write2dBoolSliceBody(enc, v)
	case [][]float32:
		return write2dFloat32SliceBody(enc, v)
	case [][]float64:
		return write2dFloat64SliceBody(enc, v)
	case [][]complex64:
		return write2dComplex64SliceBody(enc, v)
	case [][]complex128:
		return write2dComplex128SliceBody(enc, v)
	case [][]string:
		return write2dStringSliceBody(enc, v)
	case [][]interface{}:
		return write2dInterfaceSliceBody(enc, v)
	case [][]byte:
		return writeBytesSliceBody(enc, v)
	default:
		return writeOtherSliceBody(enc, v)
	}
}

func writeInt8SliceBody(enc *Encoder, slice []int8) (err error) {
	n := len(slice)
	for i := 0; i < n && err == nil; i++ {
		err = WriteInt8(enc, slice[i])
	}
	return
}

func writeInt16SliceBody(enc *Encoder, slice []int16) (err error) {
	n := len(slice)
	for i := 0; i < n && err == nil; i++ {
		err = WriteInt16(enc, slice[i])
	}
	return
}

func writeInt32SliceBody(enc *Encoder, slice []int32) (err error) {
	n := len(slice)
	for i := 0; i < n && err == nil; i++ {
		err = WriteInt32(enc, slice[i])
	}
	return
}

func writeInt64SliceBody(enc *Encoder, slice []int64) (err error) {
	n := len(slice)
	for i := 0; i < n && err == nil; i++ {
		err = WriteInt64(enc, slice[i])
	}
	return
}

func writeIntSliceBody(enc *Encoder, slice []int) (err error) {
	n := len(slice)
	for i := 0; i < n && err == nil; i++ {
		err = WriteInt(enc, slice[i])
	}
	return
}

func writeUint16SliceBody(enc *Encoder, slice []uint16) (err error) {
	n := len(slice)
	for i := 0; i < n && err == nil; i++ {
		err = WriteUint16(enc, slice[i])
	}
	return
}

func writeUint32SliceBody(enc *Encoder, slice []uint32) (err error) {
	n := len(slice)
	for i := 0; i < n && err == nil; i++ {
		err = WriteUint32(enc, slice[i])
	}
	return
}

func writeUint64SliceBody(enc *Encoder, slice []uint64) (err error) {
	n := len(slice)
	for i := 0; i < n && err == nil; i++ {
		err = WriteUint64(enc, slice[i])
	}
	return
}

func writeUintSliceBody(enc *Encoder, slice []uint) (err error) {
	n := len(slice)
	for i := 0; i < n && err == nil; i++ {
		err = WriteUint(enc, slice[i])
	}
	return
}

func writeBoolSliceBody(enc *Encoder, slice []bool) (err error) {
	n := len(slice)
	for i := 0; i < n && err == nil; i++ {
		err = WriteBool(enc, slice[i])
	}
	return
}

func writeFloat32SliceBody(enc *Encoder, slice []float32) (err error) {
	n := len(slice)
	for i := 0; i < n && err == nil; i++ {
		err = WriteFloat32(enc, slice[i])
	}
	return
}

func writeFloat64SliceBody(enc *Encoder, slice []float64) (err error) {
	n := len(slice)
	for i := 0; i < n && err == nil; i++ {
		err = WriteFloat64(enc, slice[i])
	}
	return
}

func writeComplex64SliceBody(enc *Encoder, slice []complex64) (err error) {
	n := len(slice)
	for i := 0; i < n && err == nil; i++ {
		err = WriteComplex64(enc, slice[i])
	}
	return
}

func writeComplex128SliceBody(enc *Encoder, slice []complex128) (err error) {
	n := len(slice)
	for i := 0; i < n && err == nil; i++ {
		err = WriteComplex128(enc, slice[i])
	}
	return
}

func writeStringSliceBody(enc *Encoder, slice []string) (err error) {
	n := len(slice)
	for i := 0; i < n && err == nil; i++ {
		err = EncodeString(enc, slice[i])
	}
	return
}

func writeInterfaceSliceBody(enc *Encoder, slice []interface{}) (err error) {
	n := len(slice)
	for i := 0; i < n && err == nil; i++ {
		err = enc.Encode(slice[i])
	}
	return
}

func write2dInt8SliceBody(enc *Encoder, slice [][]int8) (err error) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n && err == nil; i++ {
		m := len(slice[i])
		err = WriteHead(enc, m, TagList)
		for j := 0; j < m && err == nil; j++ {
			err = WriteInt8(enc, slice[i][j])
		}
		if err == nil {
			err = WriteFoot(enc)
		}
	}
	return
}

func write2dInt16SliceBody(enc *Encoder, slice [][]int16) (err error) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n && err == nil; i++ {
		m := len(slice[i])
		err = WriteHead(enc, m, TagList)
		for j := 0; j < m && err == nil; j++ {
			err = WriteInt16(enc, slice[i][j])
		}
		if err == nil {
			err = WriteFoot(enc)
		}
	}
	return
}

func write2dInt32SliceBody(enc *Encoder, slice [][]int32) (err error) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n && err == nil; i++ {
		m := len(slice[i])
		err = WriteHead(enc, m, TagList)
		for j := 0; j < m && err == nil; j++ {
			err = WriteInt32(enc, slice[i][j])
		}
		if err == nil {
			err = WriteFoot(enc)
		}
	}
	return
}

func write2dInt64SliceBody(enc *Encoder, slice [][]int64) (err error) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n && err == nil; i++ {
		m := len(slice[i])
		err = WriteHead(enc, m, TagList)
		for j := 0; j < m && err == nil; j++ {
			err = WriteInt64(enc, slice[i][j])
		}
		if err == nil {
			err = WriteFoot(enc)
		}
	}
	return
}

func write2dIntSliceBody(enc *Encoder, slice [][]int) (err error) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n && err == nil; i++ {
		m := len(slice[i])
		err = WriteHead(enc, m, TagList)
		for j := 0; j < m && err == nil; j++ {
			err = WriteInt(enc, slice[i][j])
		}
		if err == nil {
			err = WriteFoot(enc)
		}
	}
	return
}

func write2dUint16SliceBody(enc *Encoder, slice [][]uint16) (err error) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n && err == nil; i++ {
		m := len(slice[i])
		err = WriteHead(enc, m, TagList)
		for j := 0; j < m && err == nil; j++ {
			err = WriteUint16(enc, slice[i][j])
		}
		if err == nil {
			err = WriteFoot(enc)
		}
	}
	return
}

func write2dUint32SliceBody(enc *Encoder, slice [][]uint32) (err error) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n && err == nil; i++ {
		m := len(slice[i])
		err = WriteHead(enc, m, TagList)
		for j := 0; j < m && err == nil; j++ {
			err = WriteUint32(enc, slice[i][j])
		}
		if err == nil {
			err = WriteFoot(enc)
		}
	}
	return
}

func write2dUint64SliceBody(enc *Encoder, slice [][]uint64) (err error) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n && err == nil; i++ {
		m := len(slice[i])
		err = WriteHead(enc, m, TagList)
		for j := 0; j < m && err == nil; j++ {
			err = WriteUint64(enc, slice[i][j])
		}
		if err == nil {
			err = WriteFoot(enc)
		}
	}
	return
}

func write2dUintSliceBody(enc *Encoder, slice [][]uint) (err error) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n && err == nil; i++ {
		m := len(slice[i])
		err = WriteHead(enc, m, TagList)
		for j := 0; j < m && err == nil; j++ {
			err = WriteUint(enc, slice[i][j])
		}
		if err == nil {
			err = WriteFoot(enc)
		}
	}
	return
}

func write2dBoolSliceBody(enc *Encoder, slice [][]bool) (err error) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n && err == nil; i++ {
		m := len(slice[i])
		err = WriteHead(enc, m, TagList)
		for j := 0; j < m && err == nil; j++ {
			err = WriteBool(enc, slice[i][j])
		}
		if err == nil {
			err = WriteFoot(enc)
		}
	}
	return
}

func write2dFloat32SliceBody(enc *Encoder, slice [][]float32) (err error) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n && err == nil; i++ {
		m := len(slice[i])
		err = WriteHead(enc, m, TagList)
		for j := 0; j < m && err == nil; j++ {
			err = WriteFloat32(enc, slice[i][j])
		}
		if err == nil {
			err = WriteFoot(enc)
		}
	}
	return
}

func write2dFloat64SliceBody(enc *Encoder, slice [][]float64) (err error) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n && err == nil; i++ {
		m := len(slice[i])
		err = WriteHead(enc, m, TagList)
		for j := 0; j < m && err == nil; j++ {
			err = WriteFloat64(enc, slice[i][j])
		}
		if err == nil {
			err = WriteFoot(enc)
		}
	}
	return
}

func write2dComplex64SliceBody(enc *Encoder, slice [][]complex64) (err error) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n && err == nil; i++ {
		m := len(slice[i])
		err = WriteHead(enc, m, TagList)
		for j := 0; j < m && err == nil; j++ {
			err = WriteComplex64(enc, slice[i][j])
		}
		if err == nil {
			err = WriteFoot(enc)
		}
	}
	return
}

func write2dComplex128SliceBody(enc *Encoder, slice [][]complex128) (err error) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n && err == nil; i++ {
		m := len(slice[i])
		err = WriteHead(enc, m, TagList)
		for j := 0; j < m && err == nil; j++ {
			err = WriteComplex128(enc, slice[i][j])
		}
		if err == nil {
			err = WriteFoot(enc)
		}
	}
	return
}

func write2dStringSliceBody(enc *Encoder, slice [][]string) (err error) {
	n := len(slice)
	for i := 0; i < n && err == nil; i++ {
		m := len(slice[i])
		enc.AddReferenceCount(1)
		err = WriteHead(enc, m, TagList)
		for j := 0; j < m && err == nil; j++ {
			err = EncodeString(enc, slice[i][j])
		}
		if err == nil {
			err = WriteFoot(enc)
		}
	}
	return
}

func write2dInterfaceSliceBody(enc *Encoder, slice [][]interface{}) (err error) {
	n := len(slice)
	for i := 0; i < n && err == nil; i++ {
		m := len(slice[i])
		enc.AddReferenceCount(1)
		err = WriteHead(enc, m, TagList)
		for j := 0; j < m && err == nil; j++ {
			err = enc.Encode(slice[i][j])
		}
		if err == nil {
			err = WriteFoot(enc)
		}
	}
	return
}

func writeBytesSliceBody(enc *Encoder, slice [][]byte) (err error) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n && err == nil; i++ {
		writeBytes(enc.writer, slice[i])
	}
	return
}

func writeOtherSliceBody(enc *Encoder, slice interface{}) (err error) {
	t := reflect2.TypeOf(slice).(*reflect2.UnsafeSliceType)
	et := t.Elem()
	ptr := reflect2.PtrOf(slice)
	n := t.UnsafeLengthOf(ptr)
	for i := 0; i < n && err == nil; i++ {
		err = enc.Encode(et.UnsafeIndirect(t.UnsafeGetIndex(ptr, i)))
	}
	return
}
