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

	"github.com/modern-go/reflect2"
)

// structDecoder is the implementation of ValueEncoder for named struct.
type structDecoder struct {
	t      *reflect2.UnsafeStructType
	fields map[string]FieldAccessor
	lock   sync.RWMutex
}

func (valdec *structDecoder) decode(dec *Decoder, p interface{}) {
	index := dec.ReadInt()
	structInfo := dec.getStructInfo(index)
	dec.AddReference(p)
	ptr := reflect2.PtrOf(p)
	valdec.lock.RLock()
	defer valdec.lock.RUnlock()
	for _, name := range structInfo.names {
		field := valdec.fields[name]
		field.Decode(dec, field.Type.Type1(), field.Field.UnsafeGet(ptr))
	}
	dec.Skip()
}

func (valdec *structDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch tag {
	case TagObject:
		valdec.decode(dec, p)
	case TagMap:
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
	RegisterValueDecoder(decoder)
	decoder.lock.Lock()
	defer decoder.lock.Unlock()
	decoder.fields = getFieldMap(t)
	return decoder
}

func getStructDecoder(t reflect.Type) ValueDecoder {
	return newStructDecoder(t)
}
