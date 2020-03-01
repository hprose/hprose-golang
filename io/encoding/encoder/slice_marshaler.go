/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/encoder/slice_marshaler.go                   |
|                                                          |
| LastModified: Mar 1, 2020                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoder

import (
	"reflect"

	"github.com/hprose/hprose-golang/v3/io"
	"github.com/modern-go/reflect2"
)

// SliceMarshaler is the implementation of Marshaler for *slice.
type SliceMarshaler struct{}

var sliceMarshaler SliceMarshaler

// Encode writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as reference
func (m SliceMarshaler) Encode(enc *Encoder, v interface{}) (err error) {
	var ok bool
	if ok, err = enc.WriteReference(v); !ok && err == nil {
		err = m.Write(enc, v)
	}
	return
}

// Write writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as value
func (m SliceMarshaler) Write(enc *Encoder, v interface{}) (err error) {
	enc.SetReference(v)
	return writeSlice(enc, reflect.ValueOf(v).Elem().Interface())
}

// WriteSlice to encoder
func WriteSlice(enc *Encoder, v interface{}) (err error) {
	enc.AddReferenceCount(1)
	return writeSlice(enc, v)
}

var emptySlice = []byte{io.TagList, io.TagOpenbrace, io.TagClosebrace}

func writeSlice(enc *Encoder, v interface{}) (err error) {
	writer := enc.Writer
	if bytes, ok := v.([]byte); ok {
		return writeBytes(writer, bytes)
	}
	count := (*reflect.SliceHeader)(interfacePointer(&v).ptr).Len
	if count == 0 {
		_, err = writer.Write(emptySlice)
		return
	}
	if err = WriteHead(writer, count, io.TagList); err == nil {
		if err = writeSliceBody(enc, v); err == nil {
			err = WriteFoot(writer)
		}
	}
	return
}

func writeSliceBody(enc *Encoder, v interface{}) error {
	switch v := v.(type) {
	case []uint16:
		return writeUint16SliceBody(enc.Writer, v)
	case []uint32:
		return writeUint32SliceBody(enc.Writer, v)
	case []uint64:
		return writeUint64SliceBody(enc.Writer, v)
	case []uint:
		return writeUintSliceBody(enc.Writer, v)
	case []int8:
		return writeInt8SliceBody(enc.Writer, v)
	case []int16:
		return writeInt16SliceBody(enc.Writer, v)
	case []int32:
		return writeInt32SliceBody(enc.Writer, v)
	case []int64:
		return writeInt64SliceBody(enc.Writer, v)
	case []int:
		return writeIntSliceBody(enc.Writer, v)
	case []bool:
		return writeBoolSliceBody(enc.Writer, v)
	case []float32:
		return writeFloat32SliceBody(enc.Writer, v)
	case []float64:
		return writeFloat64SliceBody(enc.Writer, v)
	case []complex64:
		return writeComplex64SliceBody(enc, v)
	case []complex128:
		return writeComplex128SliceBody(enc, v)
	case []string:
		return writeStringSliceBody(enc, v)
	default:
		return writeOtherSliceBody(enc, v)
	}
}

func writeInt8SliceBody(writer io.Writer, slice []int8) (err error) {
	for _, e := range slice {
		err = WriteInt8(writer, e)
		if err != nil {
			return
		}
	}
	return
}

func writeInt16SliceBody(writer io.Writer, slice []int16) (err error) {
	for _, e := range slice {
		err = WriteInt16(writer, e)
		if err != nil {
			return
		}
	}
	return
}

func writeInt32SliceBody(writer io.Writer, slice []int32) (err error) {
	for _, e := range slice {
		err = WriteInt32(writer, e)
		if err != nil {
			return
		}
	}
	return
}

func writeInt64SliceBody(writer io.Writer, slice []int64) (err error) {
	for _, e := range slice {
		err = WriteInt64(writer, e)
		if err != nil {
			return
		}
	}
	return
}

func writeIntSliceBody(writer io.Writer, slice []int) (err error) {
	for _, e := range slice {
		err = WriteInt(writer, e)
		if err != nil {
			return
		}
	}
	return
}

func writeUint16SliceBody(writer io.Writer, slice []uint16) (err error) {
	for _, e := range slice {
		err = WriteUint16(writer, e)
		if err != nil {
			return
		}
	}
	return
}

func writeUint32SliceBody(writer io.Writer, slice []uint32) (err error) {
	for _, e := range slice {
		err = WriteUint32(writer, e)
		if err != nil {
			return
		}
	}
	return
}

func writeUint64SliceBody(writer io.Writer, slice []uint64) (err error) {
	for _, e := range slice {
		err = WriteUint64(writer, e)
		if err != nil {
			return
		}
	}
	return
}

func writeUintSliceBody(writer io.Writer, slice []uint) (err error) {
	for _, e := range slice {
		err = WriteUint(writer, e)
		if err != nil {
			return
		}
	}
	return
}

func writeBoolSliceBody(writer io.Writer, slice []bool) (err error) {
	for _, e := range slice {
		err = WriteBool(writer, e)
		if err != nil {
			return
		}
	}
	return
}

func writeFloat32SliceBody(writer io.Writer, slice []float32) (err error) {
	for _, e := range slice {
		err = WriteFloat32(writer, e)
		if err != nil {
			return
		}
	}
	return
}

func writeFloat64SliceBody(writer io.Writer, slice []float64) (err error) {
	for _, e := range slice {
		err = WriteFloat64(writer, e)
		if err != nil {
			return
		}
	}
	return
}

func writeComplex64SliceBody(enc *Encoder, slice []complex64) (err error) {
	for _, e := range slice {
		err = WriteComplex64(enc, e)
		if err != nil {
			return
		}
	}
	return
}

func writeComplex128SliceBody(enc *Encoder, slice []complex128) (err error) {
	for _, e := range slice {
		err = WriteComplex128(enc, e)
		if err != nil {
			return
		}
	}
	return
}

func writeStringSliceBody(enc *Encoder, slice []string) (err error) {
	for _, e := range slice {
		err = stringMarshaler.encode(enc, e)
		if err != nil {
			return
		}
	}
	return
}

func writeOtherSliceBody(enc *Encoder, slice interface{}) (err error) {
	t := reflect2.TypeOf(slice).(*reflect2.UnsafeSliceType)
	et := t.Elem()
	ptr := reflect2.PtrOf(slice)
	n := t.UnsafeLengthOf(ptr)
	for i := 0; i < n; i++ {
		if err = enc.Encode(et.UnsafeIndirect(t.UnsafeGetIndex(ptr, i))); err != nil {
			return
		}
	}
	return
}
