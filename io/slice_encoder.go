/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/slice_encoder.go                                      |
|                                                          |
| LastModified: Feb 18, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

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
	enc.setReference(v)
	enc.writeSlice(reflect.ValueOf(v).Elem().Interface())
}

// WriteSlice to encoder.
func (enc *Encoder) WriteSlice(v interface{}) {
	enc.AddReferenceCount(1)
	enc.writeSlice(v)
}

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
	enc.WriteListHead(count)
	enc.writeSliceBody(v, count)
	enc.WriteFoot()
}

func (enc *Encoder) write2dSliceBody(v interface{}, n int) {
	switch v := v.(type) {
	case [][]uint16:
		enc.write2dUint16SliceBody(v, n)
	case [][]uint32:
		enc.write2dUint32SliceBody(v, n)
	case [][]uint64:
		enc.write2dUint64SliceBody(v, n)
	case [][]uint:
		enc.write2dUintSliceBody(v, n)
	case [][]int8:
		enc.write2dInt8SliceBody(v, n)
	case [][]int16:
		enc.write2dInt16SliceBody(v, n)
	case [][]int32:
		enc.write2dInt32SliceBody(v, n)
	case [][]int64:
		enc.write2dInt64SliceBody(v, n)
	case [][]int:
		enc.write2dIntSliceBody(v, n)
	case [][]bool:
		enc.write2dBoolSliceBody(v, n)
	case [][]float32:
		enc.write2dFloat32SliceBody(v, n)
	case [][]float64:
		enc.write2dFloat64SliceBody(v, n)
	case [][]complex64:
		enc.write2dComplex64SliceBody(v, n)
	case [][]complex128:
		enc.write2dComplex128SliceBody(v, n)
	case [][]string:
		enc.write2dStringSliceBody(v, n)
	case [][]interface{}:
		enc.write2dInterfaceSliceBody(v, n)
	case [][]byte:
		enc.writeBytesSliceBody(v, n)
	default:
		enc.writeOtherSliceBody(v, n)
	}
}

func (enc *Encoder) writeSliceBody(v interface{}, n int) {
	switch v := v.(type) {
	case []uint16:
		enc.writeUint16SliceBody(v, n)
	case []uint32:
		enc.writeUint32SliceBody(v, n)
	case []uint64:
		enc.writeUint64SliceBody(v, n)
	case []uint:
		enc.writeUintSliceBody(v, n)
	case []int8:
		enc.writeInt8SliceBody(v, n)
	case []int16:
		enc.writeInt16SliceBody(v, n)
	case []int32:
		enc.writeInt32SliceBody(v, n)
	case []int64:
		enc.writeInt64SliceBody(v, n)
	case []int:
		enc.writeIntSliceBody(v, n)
	case []bool:
		enc.writeBoolSliceBody(v, n)
	case []float32:
		enc.writeFloat32SliceBody(v, n)
	case []float64:
		enc.writeFloat64SliceBody(v, n)
	case []complex64:
		enc.writeComplex64SliceBody(v, n)
	case []complex128:
		enc.writeComplex128SliceBody(v, n)
	case []string:
		enc.writeStringSliceBody(v, n)
	case []interface{}:
		enc.writeInterfaceSliceBody(v, n)
	default:
		enc.write2dSliceBody(v, n)
	}
}

func (enc *Encoder) writeInt8SliceBody(slice []int8, n int) {
	for i := 0; i < n; i++ {
		enc.WriteInt8(slice[i])
	}
}

func (enc *Encoder) writeInt16SliceBody(slice []int16, n int) {
	for i := 0; i < n; i++ {
		enc.WriteInt16(slice[i])
	}
}

func (enc *Encoder) writeInt32SliceBody(slice []int32, n int) {
	for i := 0; i < n; i++ {
		enc.WriteInt32(slice[i])
	}
}

func (enc *Encoder) writeInt64SliceBody(slice []int64, n int) {
	for i := 0; i < n; i++ {
		enc.WriteInt64(slice[i])
	}
}

func (enc *Encoder) writeIntSliceBody(slice []int, n int) {
	for i := 0; i < n; i++ {
		enc.WriteInt(slice[i])
	}
}

func (enc *Encoder) writeUint16SliceBody(slice []uint16, n int) {
	for i := 0; i < n; i++ {
		enc.WriteUint16(slice[i])
	}
}

func (enc *Encoder) writeUint32SliceBody(slice []uint32, n int) {
	for i := 0; i < n; i++ {
		enc.WriteUint32(slice[i])
	}
}

func (enc *Encoder) writeUint64SliceBody(slice []uint64, n int) {
	for i := 0; i < n; i++ {
		enc.WriteUint64(slice[i])
	}
}

func (enc *Encoder) writeUintSliceBody(slice []uint, n int) {
	for i := 0; i < n; i++ {
		enc.WriteUint(slice[i])
	}
}

func (enc *Encoder) writeBoolSliceBody(slice []bool, n int) {
	for i := 0; i < n; i++ {
		enc.WriteBool(slice[i])
	}
}

func (enc *Encoder) writeFloat32SliceBody(slice []float32, n int) {
	for i := 0; i < n; i++ {
		enc.WriteFloat32(slice[i])
	}
}

func (enc *Encoder) writeFloat64SliceBody(slice []float64, n int) {
	for i := 0; i < n; i++ {
		enc.WriteFloat64(slice[i])
	}
}

func (enc *Encoder) writeComplex64SliceBody(slice []complex64, n int) {
	for i := 0; i < n; i++ {
		enc.WriteComplex64(slice[i])
	}
}

func (enc *Encoder) writeComplex128SliceBody(slice []complex128, n int) {
	for i := 0; i < n; i++ {
		enc.WriteComplex128(slice[i])
	}
}

func (enc *Encoder) writeStringSliceBody(slice []string, n int) {
	for i := 0; i < n; i++ {
		enc.EncodeString(slice[i])
	}
}

func (enc *Encoder) writeInterfaceSliceBody(slice []interface{}, n int) {
	for i := 0; i < n; i++ {
		enc.encode(slice[i])
	}
}

func (enc *Encoder) write2dInt8SliceBody(slice [][]int8, n int) {
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteListHead(m)
		enc.writeInt8SliceBody(slice[i], m)
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dInt16SliceBody(slice [][]int16, n int) {
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteListHead(m)
		enc.writeInt16SliceBody(slice[i], m)
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dInt32SliceBody(slice [][]int32, n int) {
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteListHead(m)
		enc.writeInt32SliceBody(slice[i], m)
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dInt64SliceBody(slice [][]int64, n int) {
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteListHead(m)
		enc.writeInt64SliceBody(slice[i], m)
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dIntSliceBody(slice [][]int, n int) {
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteListHead(m)
		enc.writeIntSliceBody(slice[i], m)
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dUint16SliceBody(slice [][]uint16, n int) {
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteListHead(m)
		enc.writeUint16SliceBody(slice[i], m)
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dUint32SliceBody(slice [][]uint32, n int) {
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteListHead(m)
		enc.writeUint32SliceBody(slice[i], m)
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dUint64SliceBody(slice [][]uint64, n int) {
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteListHead(m)
		enc.writeUint64SliceBody(slice[i], m)
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dUintSliceBody(slice [][]uint, n int) {
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteListHead(m)
		enc.writeUintSliceBody(slice[i], m)
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dBoolSliceBody(slice [][]bool, n int) {
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteListHead(m)
		enc.writeBoolSliceBody(slice[i], m)
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dFloat32SliceBody(slice [][]float32, n int) {
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteListHead(m)
		enc.writeFloat32SliceBody(slice[i], m)
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dFloat64SliceBody(slice [][]float64, n int) {
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteListHead(m)
		enc.writeFloat64SliceBody(slice[i], m)
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dComplex64SliceBody(slice [][]complex64, n int) {
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteListHead(m)
		enc.writeComplex64SliceBody(slice[i], m)
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dComplex128SliceBody(slice [][]complex128, n int) {
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.WriteListHead(m)
		enc.writeComplex128SliceBody(slice[i], m)
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dStringSliceBody(slice [][]string, n int) {
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.AddReferenceCount(1)
		enc.WriteListHead(m)
		enc.writeStringSliceBody(slice[i], m)
		enc.WriteFoot()
	}
}

func (enc *Encoder) write2dInterfaceSliceBody(slice [][]interface{}, n int) {
	for i := 0; i < n; i++ {
		m := len(slice[i])
		enc.AddReferenceCount(1)
		enc.WriteListHead(m)
		enc.writeInterfaceSliceBody(slice[i], m)
		enc.WriteFoot()
	}
}

func (enc *Encoder) writeBytesSliceBody(slice [][]byte, n int) {
	enc.AddReferenceCount(n)
	for i := 0; i < n; i++ {
		enc.buf = appendBytes(enc.buf, slice[i])
	}
}

func (enc *Encoder) writeOtherSliceBody(slice interface{}, n int) {
	t := reflect2.TypeOf(slice).(*reflect2.UnsafeSliceType)
	et := t.Elem()
	ptr := reflect2.PtrOf(slice)
	for i := 0; i < n; i++ {
		enc.encode(et.UnsafeIndirect(t.UnsafeGetIndex(ptr, i)))
	}
}
