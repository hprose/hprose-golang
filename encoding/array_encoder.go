/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/array_encoder.go                                |
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

// arrayEncoder is the implementation of ValueEncoder for *array.
type arrayEncoder struct{}

var arrayenc arrayEncoder

func (valenc arrayEncoder) Encode(enc *Encoder, v interface{}) {
	if ok := enc.WriteReference(v); !ok {
		valenc.Write(enc, v)
	}
}

func (arrayEncoder) Write(enc *Encoder, v interface{}) {
	enc.setReference(v)
	enc.writeArray(reflect.ValueOf(v).Elem().Interface())
}

// WriteArray to encoder
func (enc *Encoder) WriteArray(v interface{}) {
	enc.AddReferenceCount(1)
	enc.writeArray(v)
}

func makeSlice(array interface{}, count int) unsafe.Pointer {
	return unsafe.Pointer(&reflect.SliceHeader{
		Data: (uintptr)(reflect2.PtrOf(array)),
		Len:  count,
		Cap:  count,
	})
}

func (enc *Encoder) writeArray(array interface{}) {
	t := reflect.TypeOf(array)
	sliceType := reflect.SliceOf(t.Elem())
	var slice interface{}
	sliceStruct := unpackEFace(&slice)
	sliceStruct.typ = (uintptr)(reflect2.PtrOf(sliceType))
	sliceStruct.ptr = makeSlice(array, t.Len())
	enc.writeSlice(slice)
}
