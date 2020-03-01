/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/encoder/array_marshaler.go                   |
|                                                          |
| LastModified: Mar 1, 2020                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoder

import (
	"reflect"
	"unsafe"

	"github.com/hprose/hprose-golang/v3/io"
)

// ArrayMarshaler is the implementation of Marshaler for *array.
type ArrayMarshaler struct{}

var arrayMarshaler ArrayMarshaler

// Encode writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as reference
func (m ArrayMarshaler) Encode(enc *Encoder, v interface{}) (err error) {
	var ok bool
	if ok, err = enc.WriteReference(v); !ok && err == nil {
		err = m.Write(enc, v)
	}
	return
}

// Write writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as value
func (m ArrayMarshaler) Write(enc *Encoder, v interface{}) (err error) {
	enc.SetReference(v)
	return writeArray(enc, reflect.ValueOf(v).Elem().Interface())
}

// WriteArray to encoder
func WriteArray(enc *Encoder, v interface{}) (err error) {
	enc.AddReferenceCount(1)
	return writeArray(enc, v)
}

func getType(v interface{}) unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&v))
}

var bytesType = getType(([]byte)(nil))

func makeSlice(ptr *interface{}, count int) unsafe.Pointer {
	return unsafe.Pointer(&reflect.SliceHeader{
		Data: (uintptr)((*interfaceStruct)(unsafe.Pointer(ptr)).ptr),
		Len:  count,
		Cap:  count,
	})
}

func writeArray(enc *Encoder, array interface{}) (err error) {
	t := reflect.TypeOf(array)
	st := reflect.SliceOf(t.Elem())
	sliceType := (*interfaceStruct)(unsafe.Pointer(&st)).ptr
	count := t.Len()
	writer := enc.Writer
	if sliceType == bytesType {
		return writeBytes(writer, *(*[]byte)(makeSlice(&array, count)))
	}
	if count == 0 {
		_, err = writer.Write(emptySlice)
		return
	}
	if err = WriteHead(writer, count, io.TagList); err == nil {
		var slice interface{}
		sliceStruct := (*interfaceStruct)(unsafe.Pointer(&slice))
		sliceStruct.typ = uintptr(sliceType)
		sliceStruct.ptr = makeSlice(&array, count)
		if err = writeSliceBody(enc, slice); err == nil {
			err = WriteFoot(writer)
		}
	}
	return
}
