/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/struct_encoder.go                               |
|                                                          |
| LastModified: Apr 12, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/modern-go/reflect2"
)

type field struct {
	typ    reflect2.Type
	field  reflect2.StructField
	encode EncodeHandler
}

// structEncoder is the implementation of ValueEncoder for named struct/*struct.
type structEncoder struct {
	fields   []field
	metadata []byte
}

func (valenc *structEncoder) Encode(enc *Encoder, v interface{}) {
	enc.EncodeReference(valenc, v)
}

func (valenc *structEncoder) Write(enc *Encoder, v interface{}) {
	fields := valenc.fields
	n := len(fields)
	t := reflect.TypeOf(v)
	st := t
	if t.Kind() == reflect.Ptr {
		st = t.Elem()
	} else if n == 1 {
		v = toPtr(t, v)
	}
	var r = enc.WriteStruct(st, func() {
		enc.AddReferenceCount(n)
		enc.buf = append(enc.buf, valenc.metadata...)
	})
	if t.Kind() == reflect.Ptr {
		enc.SetPtrReference(v)
	} else {
		enc.AddReferenceCount(1)
	}
	p := reflect2.PtrOf(v)
	WriteObjectHead(enc, r)
	for i := 0; i < n; i++ {
		fields[i].encode(enc, fields[i].typ.UnsafeIndirect(fields[i].field.UnsafeGet(p)))
	}
	WriteFoot(enc)
}

func toPtr(t reflect.Type, v interface{}) interface{} {
	pv := reflect.New(t)
	pv.Elem().Set(reflect.ValueOf(v))
	return pv.Interface()
}

func appendName(buf []byte, s string, message string) []byte {
	length := utf16Length(s)
	if length < 0 {
		panic(fmt.Sprintf("hprose/encoding: invalid UTF-8 in %s", message))
	}
	return appendBinary(buf, reflect2.UnsafeCastString(s), length)
}

func stripOptions(tag string) string {
	i := strings.Index(tag, ",")
	if i < 0 {
		return tag
	}
	return string(tag[:i])
}

func _fieldAlias(tag reflect.StructTag, tagname string) string {
	return strings.Trim(stripOptions(tag.Get(tagname)), " ")
}

func fieldAlias(tag reflect.StructTag, name string, tagnames []string) string {
	alias := _fieldAlias(tag, defaultTagName)
	if alias != "" {
		return alias
	}
	for _, tagname := range tagnames {
		if tagname != "" && tagname != defaultTagName {
			alias = _fieldAlias(tag, tagname)
			if alias != "" {
				return alias
			}
		}
	}
	if name[0] >= 'A' && name[0] <= 'Z' {
		name = string(name[0]-'A'+'a') + name[1:]
	}
	return name
}

func getFields(t reflect2.StructType, tagnames []string, mapping map[string]bool, names []string, fields []field) ([]string, []field) {
	n := t.NumField()
	for i := 0; i < n; i++ {
		f := t.Field(i)
		ft := f.Type()
		kind := ft.Kind()

		switch kind {
		case reflect.Func, reflect.Chan, reflect.UnsafePointer:
			continue
		case reflect.Struct:
			if f.Anonymous() {
				names, fields = getFields(ft.(reflect2.StructType), tagnames, mapping, names, fields)
				continue
			}
		}

		if f.PkgPath() != "" {
			continue
		}

		name := fieldAlias(f.Tag(), f.Name(), tagnames)
		if name == "-" {
			continue
		}
		if mapping[name] {
			panic(fmt.Sprintf("hprose/encoding: ambiguous fields with the same name or alias: %s", name))
		}

		var field field
		field.typ = ft
		field.field = f
		typ := ft.Type1()
		if field.encode = GetEncodeHandler(typ); field.encode == nil {
			continue
		}

		mapping[name] = true
		names = append(names, name)
		fields = append(fields, field)
	}
	return names, fields
}

func newStructEncoder(t reflect.Type, name string, tagnames []string) ValueEncoder {
	encoder := &structEncoder{}
	registerValueEncoder(t, encoder)
	var names []string
	names, encoder.fields = getFields(reflect2.Type2(t).(reflect2.StructType), tagnames, map[string]bool{}, nil, nil)
	n := len(names)
	var buf []byte
	buf = append(buf, TagClass)
	buf = appendName(buf, name, "struct name")
	if n > 0 {
		buf = AppendUint64(buf, uint64(n))
	}
	buf = append(buf, TagOpenbrace)
	for i := 0; i < n; i++ {
		buf = append(buf, TagString)
		buf = appendName(buf, names[i], "struct field name or alias")
	}
	buf = append(buf, TagClosebrace)
	encoder.metadata = buf
	return encoder
}

// anonymousStructEncoder is the implementation of ValueEncoder for anonymous struct/*struct.
type anonymousStructEncoder struct {
	fields []field
	names  []string
}

func newAnonymousStructEncoder(t reflect.Type) ValueEncoder {
	encoder := &anonymousStructEncoder{}
	registerValueEncoder(t, encoder)
	encoder.names, encoder.fields = getFields(reflect2.Type2(t).(reflect2.StructType), []string{"json"}, map[string]bool{}, nil, nil)
	return encoder
}

func (valenc *anonymousStructEncoder) Encode(enc *Encoder, v interface{}) {
	enc.EncodeReference(valenc, v)
}

func (valenc *anonymousStructEncoder) Write(enc *Encoder, v interface{}) {
	enc.SetReference(v)
	names, fields := valenc.names, valenc.fields
	n := len(names)
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
	WriteHead(enc, n, TagMap)
	for i := 0; i < n; i++ {
		EncodeString(enc, names[i])
		fields[i].encode(enc, fields[i].typ.UnsafeIndirect(fields[i].field.UnsafeGet(p)))
	}
	WriteFoot(enc)
}
