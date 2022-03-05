/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/struct_decoder.go                                     |
|                                                          |
| LastModified: Mar 5, 2022                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"reflect"
	"sync"
	"unsafe"

	"github.com/modern-go/reflect2"
)

func (dec *Decoder) readObjectAsMap(structInfo structInfo) map[string]interface{} {
	m := make(map[string]interface{}, len(structInfo.names))
	t := reflect2.TypeOf(m).(*reflect2.UnsafeMapType)
	if !dec.IsSimple() {
		dec.refer.Add(m)
	}
	ptr := reflect2.PtrOf(&m)
	for _, name := range structInfo.names {
		var v interface{}
		dec.decodeInterface(dec.NextByte(), &v)
		t.UnsafeSetIndex(ptr, reflect2.PtrOf(name), reflect2.PtrOf(&v))
	}
	dec.Skip()
	return m
}

func (dec *Decoder) readObject(structInfo structInfo) interface{} {
	obj := structInfo.t.New()
	dec.AddReference(obj)
	ptr := reflect2.PtrOf(obj)
	for _, name := range structInfo.names {
		if field, ok := structInfo.fields[name]; ok {
			field.Decode(dec, field.Type.Type1(), field.Field.UnsafeGet(ptr))
		} else {
			var v interface{}
			dec.decodeInterface(dec.NextByte(), &v)
		}
	}
	dec.Skip()
	return obj
}

// ReadObject reads object and add reference.
func (dec *Decoder) ReadObject() interface{} {
	index := dec.ReadInt()
	structInfo := dec.getStructInfo(index)
	if structInfo.fields == nil {
		return dec.readObjectAsMap(structInfo)
	}
	return dec.readObject(structInfo)
}

// structDecoder is the implementation of ValueEncoder for named struct.
type structDecoder struct {
	t      *reflect2.UnsafeStructType
	fields map[string]FieldAccessor
}

func (valdec *structDecoder) decodeField(dec *Decoder, ptr unsafe.Pointer, name string) {
	if field, ok := valdec.fields[name]; ok {
		field.Decode(dec, field.Type.Type1(), field.Field.UnsafeGet(ptr))
	} else {
		var v interface{}
		dec.decodeInterface(dec.NextByte(), &v)
	}
}

func (valdec *structDecoder) decodeObject(dec *Decoder, p interface{}) {
	index := dec.ReadInt()
	structInfo := dec.getStructInfo(index)
	dec.AddReference(p)
	ptr := reflect2.PtrOf(p)
	for _, name := range structInfo.names {
		valdec.decodeField(dec, ptr, name)
	}
	dec.Skip()
}

func (valdec *structDecoder) decodeMapAsObject(dec *Decoder, p interface{}) {
	ptr := reflect2.PtrOf(p)
	count := dec.ReadInt()
	dec.AddReference(p)
	for i := 0; i < count; i++ {
		var name string
		dec.decodeString(stringType, dec.NextByte(), &name)
		valdec.decodeField(dec, ptr, name)
	}
	dec.Skip()
}

func (valdec *structDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch tag {
	case TagObject:
		valdec.decodeObject(dec, p)
	case TagMap:
		valdec.decodeMapAsObject(dec, p)
	case TagEmpty:
		valdec.t.UnsafeSet(reflect2.PtrOf(p), valdec.t.UnsafeNew())
	default:
		dec.defaultDecode(valdec.t.Type1(), p, tag)
	}
}

var namedStructDecoderMap sync.Map

func registerNamedStructDecoder(t reflect.Type, valdec ValueDecoder) {
	namedStructDecoderMap.Store(t, valdec)
}

func getNamedStructDecoder(t reflect.Type) ValueDecoder {
	if valdec, ok := namedStructDecoderMap.Load(t); ok {
		return valdec.(ValueDecoder)
	}
	return nil
}

func newNamedStructDecoder(t reflect.Type, tag ...string) *structDecoder {
	t2 := reflect2.Type2(t).(*reflect2.UnsafeStructType)
	decoder := &structDecoder{t: t2}
	registerNamedStructDecoder(t, decoder)
	decoder.fields = getFieldMap(t, tag...)
	return decoder
}

func newAnonymousStructDecoder(t reflect.Type, tag ...string) *structDecoder {
	t2 := reflect2.Type2(t).(*reflect2.UnsafeStructType)
	decoder := &structDecoder{t: t2}
	decoder.fields = getFieldMap(t, tag...)
	return decoder
}

func getStructDecoder(t reflect.Type) ValueDecoder {
	if t.Name() == "" {
		return newAnonymousStructDecoder(t)
	}
	return newNamedStructDecoder(t)
}
