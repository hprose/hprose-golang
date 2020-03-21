/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/array_encoder.go                             |
|                                                          |
| LastModified: Mar 21, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

// arrayEncoder is the implementation of ValueEncoder for *array.
type arrayEncoder struct{}

var arrayenc arrayEncoder

func (valenc arrayEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	var ok bool
	if ok, err = enc.WriteReference(v); !ok && err == nil {
		err = valenc.Write(enc, v)
	}
	return
}

func (arrayEncoder) Write(enc *Encoder, v interface{}) (err error) {
	enc.SetReference(v)
	return writeArray(enc, reflect.ValueOf(v).Elem().Interface())
}

// WriteArray to encoder
func WriteArray(enc *Encoder, v interface{}) (err error) {
	enc.AddReferenceCount(1)
	return writeArray(enc, v)
}

func makeSlice(array interface{}, count int) unsafe.Pointer {
	return unsafe.Pointer(&reflect.SliceHeader{
		Data: (uintptr)(reflect2.PtrOf(array)),
		Len:  count,
		Cap:  count,
	})
}

func writeArray(enc *Encoder, array interface{}) (err error) {
	t := reflect.TypeOf(array)
	sliceType := reflect.SliceOf(t.Elem())
	var slice interface{}
	sliceStruct := unpackEFace(&slice)
	sliceStruct.typ = (uintptr)(reflect2.PtrOf(sliceType))
	sliceStruct.ptr = makeSlice(array, t.Len())
	return writeSlice(enc, slice)
}
