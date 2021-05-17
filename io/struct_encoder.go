/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/struct_encoder.go                                     |
|                                                          |
| LastModified: Jun 25, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/modern-go/reflect2"
)

// structEncoder is the implementation of ValueEncoder for named struct/*struct.
type structEncoder struct {
	fields   []FieldAccessor
	metadata []byte
	lock     sync.RWMutex
}

func (valenc *structEncoder) Encode(enc *Encoder, v interface{}) {
	enc.EncodeReference(valenc, v)
}

func (valenc *structEncoder) Write(enc *Encoder, v interface{}) {
	valenc.lock.RLock()
	defer valenc.lock.RUnlock()
	fields := valenc.fields
	n := len(fields)
	t := reflect.TypeOf(v)
	st := t
	if t.Kind() == reflect.Ptr {
		st = t.Elem()
	} else if n == 1 {
		v = toPtr(t, v)
	}
	var r = enc.WriteStructType(st, func() {
		enc.AddReferenceCount(n)
		enc.buf = append(enc.buf, valenc.metadata...)
	})
	enc.SetReference(v)
	p := reflect2.PtrOf(v)
	enc.WriteObjectHead(r)
	for i := 0; i < n; i++ {
		fields[i].Encode(enc, fields[i].Type.UnsafeIndirect(fields[i].Field.UnsafeGet(p)))
	}
	enc.WriteFoot()
}

func appendName(buf []byte, s string, message string) []byte {
	length := utf16Length(s)
	if length < 0 {
		panic(fmt.Sprintf("hprose/io: invalid UTF-8 in %s", message))
	}
	return appendBinary(buf, reflect2.UnsafeCastString(s), length)
}

func toPtr(t reflect.Type, v interface{}) interface{} {
	pv := reflect.New(t)
	pv.Elem().Set(reflect.ValueOf(v))
	return pv.Interface()
}

func newStructEncoder(t reflect.Type, name string, tag ...string) *structEncoder {
	encoder := &structEncoder{}
	encoder.lock.Lock()
	defer encoder.lock.Unlock()
	registerValueEncoder(t, encoder)
	fields := getFields(t, tag...)
	n := len(fields)
	var metadata []byte
	metadata = append(metadata, TagClass)
	metadata = appendName(metadata, name, "struct name")
	if n > 0 {
		metadata = AppendUint64(metadata, uint64(n))
	}
	metadata = append(metadata, TagOpenbrace)
	for i := 0; i < n; i++ {
		metadata = append(metadata, TagString)
		metadata = appendName(metadata, fields[i].Alias, "struct field name or alias")
	}
	metadata = append(metadata, TagClosebrace)
	encoder.fields = fields
	encoder.metadata = metadata
	return encoder
}

// anonymousStructEncoder is the implementation of ValueEncoder for anonymous struct/*struct.
type anonymousStructEncoder struct {
	fields []FieldAccessor
	lock   sync.RWMutex
}

func newAnonymousStructEncoder(t reflect.Type, tag ...string) *anonymousStructEncoder {
	encoder := &anonymousStructEncoder{}
	encoder.lock.Lock()
	defer encoder.lock.Unlock()
	registerValueEncoder(t, encoder)
	encoder.fields = getFields(t, tag...)
	return encoder
}

func (valenc *anonymousStructEncoder) Encode(enc *Encoder, v interface{}) {
	enc.EncodeReference(valenc, v)
}

func (valenc *anonymousStructEncoder) Write(enc *Encoder, v interface{}) {
	valenc.lock.RLock()
	defer valenc.lock.RUnlock()
	enc.SetReference(v)
	fields := valenc.fields
	n := len(fields)
	switch n {
	case 0:
		enc.buf = append(enc.buf, TagMap, TagOpenbrace, TagClosebrace)
		return
	case 1:
		if t := reflect.TypeOf(v); t.Kind() == reflect.Struct {
			v = toPtr(t, v)
		}
	}
	p := reflect2.PtrOf(v)
	enc.WriteMapHead(n)
	for i := 0; i < n; i++ {
		enc.EncodeString(fields[i].Alias)
		fields[i].Encode(enc, fields[i].Type.UnsafeIndirect(fields[i].Field.UnsafeGet(p)))
	}
	enc.WriteFoot()
}
