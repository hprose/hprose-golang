/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/encoder/map_marshaler.go                     |
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
	"github.com/modern-go/reflect2"
)

// MapMarshaler is the implementation of Marshaler for *map.
type MapMarshaler struct{}

var mapMarshaler MapMarshaler

// Encode writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as reference
func (m MapMarshaler) Encode(enc *Encoder, v interface{}) (err error) {
	var ok bool
	if ok, err = enc.WriteReference(v); !ok && err == nil {
		err = m.Write(enc, v)
	}
	return
}

// Write writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as value
func (m MapMarshaler) Write(enc *Encoder, v interface{}) (err error) {
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
	// case map[string]int:
	// 	return writeStringStringMapBody(enc.Writer, v)
	// case map[string]interface{}:
	// 	return writeStringInterfaceMapBody(enc.Writer, v)
	// case map[int]string:
	// 	return writeIntStringMapBody(enc.Writer, v)
	// case map[int]int:
	// 	return writeIntIntMapBody(enc.Writer, v)
	// case map[int]interface{}:
	// 	return writeIntInterfaceMapBody(enc.Writer, v)
	// case map[interface{}]string:
	// 	return writeInterfaceStringMapBody(enc.Writer, v)
	// case map[interface{}]int:
	// 	return writeInterfaceIntMapBody(enc.Writer, v)
	// case map[interface{}]interface{}:
	// 	return writeInterfaceInterfaceMapBody(enc.Writer, v)
	default:
		return writeOtherMapBody(enc, v)
	}
}

func writeStringStringMapBody(enc *Encoder, v interface{}) (err error) {
	m := v.(map[string]string)
	for k, v := range m {
		if err = stringMarshaler.encode(enc, k); err != nil {
			return
		}
		if err = stringMarshaler.encode(enc, v); err != nil {
			return
		}
	}
	return
}

func writeOtherMapBody(enc *Encoder, v interface{}) (err error) {
	mapType := reflect2.TypeOf(v).(*reflect2.UnsafeMapType)
	ptr := unsafe.Pointer(&(*interfaceStruct)(unsafe.Pointer(&v)).ptr)
	p := reflect.NewAt(reflect.TypeOf(v), ptr).Interface()
	iter := mapType.Iterate(p)
	var key, value interface{}
	kp := (*interfaceStruct)(unsafe.Pointer(&key))
	vp := (*interfaceStruct)(unsafe.Pointer(&value))
	kp.typ = mapType.Key().RType()
	vp.typ = mapType.Elem().RType()
	for iter.HasNext() {
		kp.ptr, vp.ptr = iter.UnsafeNext()
		if err = enc.Encode(key); err != nil {
			return
		}
		if err = enc.Encode(value); err != nil {
			return
		}
	}
	return
}
