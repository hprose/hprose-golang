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
	enc.writeSlice(reflect.ValueOf(v).Elem().Interface())
}

// WriteSlice to encoder
func (enc *Encoder) WriteSlice(v interface{}) {
	enc.AddReferenceCount(1)
	enc.writeSlice(v)
}

var emptySlice = []byte{TagList, TagOpenbrace, TagClosebrace}

func (enc *Encoder) writeSlice(v interface{}) {
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
	enc.writeSliceBody(v)
	enc.WriteFoot()
}

func (enc *Encoder) writeSliceBody(v interface{}) {
	switch v := v.(type) {
	case []uint16:
		enc.writeUint16SliceBody(v)
	case []uint32:
		enc.writeUint32SliceBody(v)
	case []uint64:
		enc.writeUint64SliceBody(v)
	case []uint:
		enc.writeUintSliceBody(v)
	case []int8:
		enc.writeInt8SliceBody(v)
	case []int16:
		enc.writeInt16SliceBody(v)
	case []int32:
		enc.writeInt32SliceBody(v)
	case []int64:
		enc.writeInt64SliceBody(v)
	case []int:
		enc.writeIntSliceBody(v)
	case []bool:
		enc.writeBoolSliceBody(v)
	case []float32:
		enc.writeFloat32SliceBody(v)
	case []float64:
		enc.writeFloat64SliceBody(v)
	case []complex64:
		enc.writeComplex64SliceBody(v)
	case []complex128:
		enc.writeComplex128SliceBody(v)
	case []string:
		enc.writeStringSliceBody(v)
	case []interface{}:
		enc.writeInterfaceSliceBody(v)
	case [][]uint16:
		enc.write2dUint16SliceBody(v)
	case [][]uint32:
		enc.write2dUint32SliceBody(v)
	case [][]uint64:
		enc.write2dUint64SliceBody(v)
	case [][]uint:
		enc.write2dUintSliceBody(v)
	case [][]int8:
		enc.write2dInt8SliceBody(v)
	case [][]int16:
		enc.write2dInt16SliceBody(v)
	case [][]int32:
		enc.write2dInt32SliceBody(v)
	case [][]int64:
		enc.write2dInt64SliceBody(v)
	case [][]int:
		enc.write2dIntSliceBody(v)
	case [][]bool:
		enc.write2dBoolSliceBody(v)
	case [][]float32:
		enc.write2dFloat32SliceBody(v)
	case [][]float64:
		enc.write2dFloat64SliceBody(v)
	case [][]complex64:
		enc.write2dComplex64SliceBody(v)
	case [][]complex128:
		enc.write2dComplex128SliceBody(v)
	case [][]string:
		enc.write2dStringSliceBody(v)
	case [][]interface{}:
		enc.write2dInterfaceSliceBody(v)
	case [][]byte:
		enc.writeBytesSliceBody(v)
	default:
		enc.writeOtherSliceBody(v)
	}
}

func (enc *Encoder) writeInt8SliceBody(slice []int8) {
	n := len(slice)
	for i := 0; i < n; i++ {
		enc.WriteInt8(slice[i])
	}
}

func (enc *Encoder) writeInt16SliceBody(slice []int16) {
	n := len(slice)
	for i := 0; i < n; i++ {
		enc.WriteInt16(slice[i])
	}
}

func (enc *Encoder) writeInt32SliceBody(slice []int32) {
	n := len(slice)
	for i := 0; i < n; i++ {
		enc.WriteInt32(slice[i])
	}
}

func (enc *Encoder) writeInt64SliceBody(slice []int64) {
	n := len(slice)
	for i := 0; i < n; i++ {
		enc.WriteInt64(slice[i])
	}
}

func (enc *Encoder) writeIntSliceBody(slice []int) {
	n := len(slice)
	for i := 0; i < n; i++ {
		enc.WriteInt(slice[i])
	}
}

func (enc *Encoder) writeUint16SliceBody(slice []uint16) {
	n := len(slice)
	for i := 0; i < n; i++ {
		enc.WriteUint16(slice[i])
	}
}

func (enc *Encoder) writeUint32SliceBody(slice []uint32) {
	n := len(slice)
	for i := 0; i < n; i++ {
		enc.WriteUint32(slice[i])
	}
}

func (enc *Encoder) writeUint64SliceBody(slice []uint64) {
	n := len(slice)
	for i := 0; i < n; i++ {
		enc.WriteUint64(slice[i])
	}
}

func (enc *Encoder) writeUintSliceBody(slice []uint) {
	n := len(slice)
	for i := 0; i < n; i++ {
		enc.WriteUint(slice[i])
	}
}

func (enc *Encoder) writeBoolSliceBody(slice []bool) {
	n := len(slice)
	for i := 0; i < n; i++ {
		enc.WriteBool(slice[i])
	}
}

func (enc *Encoder) writeFloat32SliceBody(slice []float32) {
	n := len(slice)
	for i := 0; i < n; i++ {
		enc.WriteFloat32(slice[i])
	}
}

func (enc *Encoder) writeFloat64SliceBody(slice []float64) {
	n := len(slice)
	for i := 0; i < n; i++ {
		enc.WriteFloat64(slice[i])
	}
}

func (enc *Encoder) writeComplex64SliceBody(slice []complex64) {
	n := len(slice)
	for i := 0; i < n; i++ {
		enc.WriteComplex64(slice[i])
	}
}

func (enc *Encoder) writeComplex128SliceBody(slice []complex128) {
	n := len(slice)
	for i := 0; i < n; i++ {
		enc.WriteComplex128(slice[i])
	}
}

func (enc *Encoder) writeStringSliceBody(slice []string) {
	n := len(slice)
	for i := 0; i < n; i++ {
		enc.EncodeString(slice[i])
	}
}

func (enc *Encoder) writeInterfaceSliceBody(slice []interface{}) {
	n := len(slice)
	for i := 0; i < n; i++ {
		enc.encode(slice[i])
	}
}

func (enc *Encoder) write2dInt8SliceBody(slice [][]int8) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			enc.WriteInt8(slice[i][j])
		}
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dInt16SliceBody(slice [][]int16) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			enc.WriteInt16(slice[i][j])
		}
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dInt32SliceBody(slice [][]int32) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			enc.WriteInt32(slice[i][j])
		}
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dInt64SliceBody(slice [][]int64) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			enc.WriteInt64(slice[i][j])
		}
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dIntSliceBody(slice [][]int) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			enc.WriteInt(slice[i][j])
		}
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dUint16SliceBody(slice [][]uint16) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			enc.WriteUint16(slice[i][j])
		}
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dUint32SliceBody(slice [][]uint32) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			enc.WriteUint32(slice[i][j])
		}
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dUint64SliceBody(slice [][]uint64) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			enc.WriteUint64(slice[i][j])
		}
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dUintSliceBody(slice [][]uint) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			enc.WriteUint(slice[i][j])
		}
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dBoolSliceBody(slice [][]bool) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			enc.WriteBool(slice[i][j])
		}
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dFloat32SliceBody(slice [][]float32) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			enc.WriteFloat32(slice[i][j])
		}
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dFloat64SliceBody(slice [][]float64) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			enc.WriteFloat64(slice[i][j])
		}
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dComplex64SliceBody(slice [][]complex64) {
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

func (enc *Encoder) write2dComplex128SliceBody(slice [][]complex128) {
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

func (enc *Encoder) write2dStringSliceBody(slice [][]string) {
	n := len(slice)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.AddReferenceCount(1)
		enc.WriteHead(m, TagList)
		for j := 0; j < m; j++ {
			enc.EncodeString(slice[i][j])
		}
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dInterfaceSliceBody(slice [][]interface{}) {
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

func (enc *Encoder) writeBytesSliceBody(slice [][]byte) {
	n := len(slice)
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		enc.buf = appendBytes(enc.buf, slice[i])
	}
}

func (enc *Encoder) writeOtherSliceBody(slice interface{}) {
	t := reflect2.TypeOf(slice).(*reflect2.UnsafeSliceType)
	et := t.Elem()
	ptr := reflect2.PtrOf(slice)
	n := t.UnsafeLengthOf(ptr)
	for i := 0; i < n; i++ {
		enc.encode(et.UnsafeIndirect(t.UnsafeGetIndex(ptr, i)))
	}
}
