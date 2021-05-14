/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/struct_decoder.go                               |
|                                                          |
| LastModified: May 14, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"reflect"
	"sync"
	"unsafe"

	"github.com/modern-go/reflect2"
)

func (dec *Decoder) readObjectAsMap(structInfo structInfo) map[string]interface{} {
	m := make(map[string]interface{}, len(structInfo.names))
	t := reflect2.TypeOf(m).(*reflect2.UnsafeMapType)
	dec.AddReference(m)
	ptr := reflect2.PtrOf(&m)
	for _, name := range structInfo.names {
		v := dec.decodeInterface(dec.NextByte())
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
			dec.decodeInterface(dec.NextByte())
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
	lock   sync.RWMutex
}

func (valdec *structDecoder) decodeField(dec *Decoder, ptr unsafe.Pointer, name string) {
	if field, ok := valdec.fields[name]; ok {
		field.Decode(dec, field.Type.Type1(), field.Field.UnsafeGet(ptr))
	} else {
		dec.decodeInterface(dec.NextByte())
	}
}

func (valdec *structDecoder) decodeObject(dec *Decoder, p interface{}) {
	index := dec.ReadInt()
	structInfo := dec.getStructInfo(index)
	dec.AddReference(p)
	ptr := reflect2.PtrOf(p)
	valdec.lock.RLock()
	defer valdec.lock.RUnlock()
	for _, name := range structInfo.names {
		valdec.decodeField(dec, ptr, name)
	}
	dec.Skip()
}

func (valdec *structDecoder) decodeMapAsObject(dec *Decoder, p interface{}) {
	ptr := reflect2.PtrOf(p)
	valdec.lock.RLock()
	defer valdec.lock.RUnlock()
	count := dec.ReadInt()
	dec.AddReference(p)
	for i := 0; i < count; i++ {
		valdec.decodeField(dec, ptr, dec.decodeString(stringType, dec.NextByte()))
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
		dec.defaultDecode(valdec.Type(), p, tag)
	}
}

func (valdec *structDecoder) Type() reflect.Type {
	return valdec.t.Type1()
}

// newStructDecoder returns a ValueDecoder for struct T.
func newStructDecoder(t reflect.Type) *structDecoder {
	decoder := &structDecoder{t: reflect2.Type2(t).(*reflect2.UnsafeStructType)}
	decoder.lock.Lock()
	defer decoder.lock.Unlock()
	RegisterValueDecoder(decoder)
	decoder.fields = getFieldMap(t)
	return decoder
}

func getStructDecoder(t reflect.Type) ValueDecoder {
	return newStructDecoder(t)
}
