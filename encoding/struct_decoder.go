/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/struct_decoder.go                               |
|                                                          |
| LastModified: Jun 25, 2020                               |
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
		_ = dec.decodeInterface(interfaceType, dec.NextByte())
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
		dec.decodeError(valdec.Type(), tag)
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
