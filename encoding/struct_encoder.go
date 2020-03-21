/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/struct_encoder.go                               |
|                                                          |
| LastModified: Mar 21, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"bytes"
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

func (valenc *structEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	return EncodeReference(valenc, enc, v)
}

func (valenc *structEncoder) Write(enc *Encoder, v interface{}) (err error) {
	t := reflect.TypeOf(v)
	st := t
	if t.Kind() == reflect.Ptr {
		st = t.Elem()
	}
	writer := enc.Writer
	fields := valenc.fields
	n := len(fields)
	var r int
	r, err = enc.WriteStructType(st, func() (err error) {
		enc.AddReferenceCount(n)
		_, err = writer.Write(valenc.metadata)
		return
	})
	if err == nil {
		if t.Kind() == reflect.Ptr {
			enc.SetReference(v)
		} else {
			enc.AddReferenceCount(1)
		}
		p := reflect2.PtrOf(v)
		err = WriteObjectHead(writer, r)
		for i := 0; i < n && err == nil; i++ {
			err = fields[i].encode(enc, fields[i].typ.UnsafeIndirect(fields[i].field.UnsafeGet(p)))
		}
		if err == nil {
			err = WriteFoot(writer)
		}
	}
	return
}

func writeName(writer BytesWriter, s string) (err error) {
	length := utf16Length(s)
	if length < 0 {
		return ErrInvalidUTF8
	}
	return writeBinary(writer, reflect2.UnsafeCastString(s), length)
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
			panic(fmt.Sprintf("ambiguous fields with the same name or alias: %s", name))
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
	buffer := &bytes.Buffer{}
	buffer.WriteByte(TagClass)
	writeName(buffer, name)
	if n > 0 {
		writeUint64(buffer, uint64(n))
	}
	buffer.WriteByte(TagOpenbrace)
	for i := 0; i < n; i++ {
		buffer.WriteByte(TagString)
		writeName(buffer, names[i])
	}
	buffer.WriteByte(TagClosebrace)
	encoder.metadata = buffer.Bytes()
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

func (valenc *anonymousStructEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	return EncodeReference(valenc, enc, v)
}

func (valenc *anonymousStructEncoder) Write(enc *Encoder, v interface{}) (err error) {
	SetReference(enc, v)
	writer := enc.Writer
	names, fields := valenc.names, valenc.fields
	n := len(names)
	if n == 0 {
		_, err = writer.Write(emptyMap)
		return
	}
	p := reflect2.PtrOf(v)
	err = WriteHead(writer, n, TagMap)
	for i := 0; i < n && err == nil; i++ {
		EncodeString(enc, names[i])
		err = fields[i].encode(enc, fields[i].typ.UnsafeIndirect(fields[i].field.UnsafeGet(p)))
	}
	if err == nil {
		err = WriteFoot(writer)
	}
	return
}
