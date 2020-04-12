/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/slice_encoder.go                                |
|                                                          |
| LastModified: Apr 12, 2020                               |
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

func (valenc sliceEncoder) Encode(enc *Encoder, v interface{}) {
	if ok := enc.WriteReference(v); !ok {
		valenc.Write(enc, v)
	}
}

func (sliceEncoder) Write(enc *Encoder, v interface{}) {
	enc.SetPtrReference(v)
	writeSlice(enc, reflect.ValueOf(v).Elem().Interface())
}

// WriteSlice to encoder
func WriteSlice(enc *Encoder, v interface{}) {
	enc.AddReferenceCount(1)
	writeSlice(enc, v)
}

var emptySlice = []byte{TagList, TagOpenbrace, TagClosebrace}

func writeSlice(enc *Encoder, v interface{}) {
	if bytes, ok := v.([]byte); ok {
		enc.buf = appendBytes(enc.buf, bytes)
		return
	}
	count := (*reflect.SliceHeader)(reflect2.PtrOf(v)).Len
	if count == 0 {
		enc.buf = append(enc.buf, TagList, TagOpenbrace, TagClosebrace)
		return
	}
	enc.WriteHead(count, TagList)
	writeSliceBody(enc, v)
	enc.WriteFoot()
}

func writeSliceBody(enc *Encoder, v interface{}) {
	switch v := v.(type) {
	case []uint16:
		writeUint16SliceBody(enc, v)
	case []uint32:
		writeUint32SliceBody(enc, v)
	case []uint64:
		writeUint64SliceBody(enc, v)
	case []uint:
		writeUintSliceBody(enc, v)
	case []int8:
		writeInt8SliceBody(enc, v)
	case []int16:
		writeInt16SliceBody(enc, v)
	case []int32:
		writeInt32SliceBody(enc, v)
	case []int64:
		writeInt64SliceBody(enc, v)
	case []int:
		writeIntSliceBody(enc, v)
	case []bool:
		writeBoolSliceBody(enc, v)
	case []float32:
		writeFloat32SliceBody(enc, v)
	case []float64:
		writeFloat64SliceBody(enc, v)
	case []complex64:
		writeComplex64SliceBody(enc, v)
	case []complex128:
		writeComplex128SliceBody(enc, v)
	case []string:
		writeStringSliceBody(enc, v)
	case []interface{}:
		writeInterfaceSliceBody(enc, v)
	case [][]uint16:
		write2dUint16SliceBody(enc, v)
	case [][]uint32:
		write2dUint32SliceBody(enc, v)
	case [][]uint64:
		write2dUint64SliceBody(enc, v)
	case [][]uint:
		write2dUintSliceBody(enc, v)
	case [][]int8:
		write2dInt8SliceBody(enc, v)
	case [][]int16:
		write2dInt16SliceBody(enc, v)
	case [][]int32:
		write2dInt32SliceBody(enc, v)
	case [][]int64:
		write2dInt64SliceBody(enc, v)
	case [][]int:
		write2dIntSliceBody(enc, v)
	case [][]bool:
		write2dBoolSliceBody(enc, v)
	case [][]float32:
		write2dFloat32SliceBody(enc, v)
	case [][]float64:
		write2dFloat64SliceBody(enc, v)
	case [][]complex64:
		write2dComplex64SliceBody(enc, v)
	case [][]complex128:
		write2dComplex128SliceBody(enc, v)
	case [][]string:
		write2dStringSliceBody(enc, v)
	case [][]interface{}:
		write2dInterfaceSliceBody(enc, v)
	case [][]byte:
		writeBytesSliceBody(enc, v)
	default:
		writeOtherSliceBody(enc, v)
	}
}

func writeInt8SliceBody(enc *Encoder, slice []int8) {
	n := len(slice)
	for i := 0; i < n; i++ {
		WriteInt8(enc, slice[i])
	}
}

func writeInt16SliceBody(enc *Encoder, slice []int16) {
	n := len(slice)
	for i := 0; i < n; i++ {
		WriteInt16(enc, slice[i])
	}
}

func writeInt32SliceBody(enc *Encoder, slice []int32) {
	n := len(slice)
	for i := 0; i < n; i++ {
		WriteInt32(enc, slice[i])
	}
}

func writeInt64SliceBody(enc *Encoder, slice []int64) {
	n := len(slice)
	for i := 0; i < n; i++ {
		WriteInt64(enc, slice[i])
	}
}

func writeIntSliceBody(enc *Encoder, slice []int) {
	n := len(slice)
	for i := 0; i < n; i++ {
		WriteInt(enc, slice[i])
	}
}

func writeUint16SliceBody(enc *Encoder, slice []uint16) {
	n := len(slice)
	for i := 0; i < n; i++ {
		WriteUint16(enc, slice[i])
	}
}

func writeUint32SliceBody(enc *Encoder, slice []uint32) {
	n := len(slice)
	for i := 0; i < n; i++ {
		WriteUint32(enc, slice[i])
	}
}

func writeUint64SliceBody(enc *Encoder, slice []uint64) {
	n := len(slice)
	for i := 0; i < n; i++ {
		WriteUint64(enc, slice[i])
	}
}

func writeUintSliceBody(enc *Encoder, slice []uint) {
	n := len(slice)
	for i := 0; i < n; i++ {
		WriteUint(enc, slice[i])
	}
}

func writeBoolSliceBody(enc *Encoder, slice []bool) {
	n := len(slice)
	for i := 0; i < n; i++ {
		WriteBool(enc, slice[i])
	}
}

func writeFloat32SliceBody(enc *Encoder, slice []float32) {
	n := len(slice)
	for i := 0; i < n; i++ {
		WriteFloat32(enc, slice[i])
	}
}

func writeFloat64SliceBody(enc *Encoder, slice []float64) {
	n := len(slice)
	for i := 0; i < n; i++ {
		WriteFloat64(enc, slice[i])
	}
}

func writeComplex64SliceBody(enc *Encoder, slice []complex64) {
	n := len(slice)
	for i := 0; i < n; i++ {
		enc.WriteComplex64(slice[i])
	}
}

func writeComplex128SliceBody(enc *Encoder, slice []complex128) {
	n := len(slice)
	for i := 0; i < n; i++ {
		enc.WriteComplex128(slice[i])
	}
}

func writeStringSliceBody(enc *Encoder, slice []string) {
	n := len(slice)
	for i := 0; i < n; i++ {
		EncodeString(enc, slice[i])
	}
}

func writeInterfaceSliceBody(enc *Encoder, slice []interface{}) {
	n := len(slice)
	for i := 0; i < n; i++ {
		enc.encode(slice[i])
	}
}

func write2dInt8SliceBody(enc *Encoder, slice [][]int8) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			WriteInt8(enc, slice[i][j])
		}
		enc.WriteFoot()
	}
}

func write2dInt16SliceBody(enc *Encoder, slice [][]int16) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			WriteInt16(enc, slice[i][j])
		}
		enc.WriteFoot()
	}
}

func write2dInt32SliceBody(enc *Encoder, slice [][]int32) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			WriteInt32(enc, slice[i][j])
		}
		enc.WriteFoot()
	}
}

func write2dInt64SliceBody(enc *Encoder, slice [][]int64) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			WriteInt64(enc, slice[i][j])
		}
		enc.WriteFoot()
	}
}

func write2dIntSliceBody(enc *Encoder, slice [][]int) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			WriteInt(enc, slice[i][j])
		}
		enc.WriteFoot()
	}
}

func write2dUint16SliceBody(enc *Encoder, slice [][]uint16) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			WriteUint16(enc, slice[i][j])
		}
		enc.WriteFoot()
	}
}

func write2dUint32SliceBody(enc *Encoder, slice [][]uint32) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			WriteUint32(enc, slice[i][j])
		}
		enc.WriteFoot()
	}
}

func write2dUint64SliceBody(enc *Encoder, slice [][]uint64) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			WriteUint64(enc, slice[i][j])
		}
		enc.WriteFoot()
	}
}

func write2dUintSliceBody(enc *Encoder, slice [][]uint) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			WriteUint(enc, slice[i][j])
		}
		enc.WriteFoot()
	}
}

func write2dBoolSliceBody(enc *Encoder, slice [][]bool) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			WriteBool(enc, slice[i][j])
		}
		enc.WriteFoot()
	}
}

func write2dFloat32SliceBody(enc *Encoder, slice [][]float32) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			WriteFloat32(enc, slice[i][j])
		}
		enc.WriteFoot()
	}
}

func write2dFloat64SliceBody(enc *Encoder, slice [][]float64) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			WriteFloat64(enc, slice[i][j])
		}
		enc.WriteFoot()
	}
}

func write2dComplex64SliceBody(enc *Encoder, slice [][]complex64) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			enc.WriteComplex64(slice[i][j])
		}
		enc.WriteFoot()
	}
}

func write2dComplex128SliceBody(enc *Encoder, slice [][]complex128) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			enc.WriteComplex128(slice[i][j])
		}
		enc.WriteFoot()
	}
}

func write2dStringSliceBody(enc *Encoder, slice [][]string) {
	n := len(slice)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.AddReferenceCount(1)
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			EncodeString(enc, slice[i][j])
		}
		enc.WriteFoot()
	}
}

func write2dInterfaceSliceBody(enc *Encoder, slice [][]interface{}) {
	n := len(slice)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.AddReferenceCount(1)
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			enc.encode(slice[i][j])
		}
		enc.WriteFoot()
	}
}

func writeBytesSliceBody(enc *Encoder, slice [][]byte) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		enc.buf = appendBytes(enc.buf, slice[i])
	}
}

func writeOtherSliceBody(enc *Encoder, slice interface{}) {
	t := reflect2.TypeOf(slice).(*reflect2.UnsafeSliceType)
	et := t.Elem()
	ptr := reflect2.PtrOf(slice)
	n := t.UnsafeLengthOf(ptr)
	for i := 0; i < n; i++ {
		enc.encode(et.UnsafeIndirect(t.UnsafeGetIndex(ptr, i)))
	}
}
