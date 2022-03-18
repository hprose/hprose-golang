/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/struct_manager.go                                     |
|                                                          |
| LastModified: Mar 18, 2022                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/modern-go/reflect2"
)

var defaultTags = []string{"hprose", "json"}

// FieldAccessor .
type FieldAccessor struct {
	Type   reflect2.Type
	Alias  string
	Field  reflect2.StructField
	Encode EncodeHandler
	Decode DecodeHandler
}

func stripOptions(tag string) string {
	i := strings.Index(tag, ",")
	if i < 0 {
		return tag
	}
	return tag[:i]
}

func _fieldAlias(tag reflect.StructTag, tagname string) string {
	return strings.Trim(stripOptions(tag.Get(tagname)), " ")
}

func fieldAlias(tag reflect.StructTag, name string, tags []string) string {
	for _, tagname := range defaultTags {
		if tagname != "" {
			if alias := _fieldAlias(tag, tagname); alias != "" {
				return alias
			}
		}
	}
	for _, tagname := range tags {
		if tagname != "" {
			if alias := _fieldAlias(tag, tagname); alias != "" {
				return alias
			}
		}
	}
	if name[0] >= 'A' && name[0] <= 'Z' {
		name = string(name[0]-'A'+'a') + name[1:]
	}
	return name
}

func _getFields(t reflect2.StructType, tags []string, mapping map[string]struct{}, fields []FieldAccessor) []FieldAccessor {
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
				fields = _getFields(ft.(reflect2.StructType), tags, mapping, fields)
				continue
			}
		}

		if f.PkgPath() != "" {
			continue
		}

		name := fieldAlias(f.Tag(), f.Name(), tags)
		if name == "-" {
			continue
		}
		if _, ok := mapping[name]; ok {
			panic(fmt.Sprintf("hprose/io: ambiguous fields with the same name or alias: %s", name))
		}

		var field FieldAccessor
		field.Type = ft
		field.Alias = name
		field.Field = f
		typ := ft.Type1()
		if field.Encode = GetEncodeHandler(typ); field.Encode == nil {
			continue
		}
		if field.Decode = GetDecodeHandler(typ); field.Decode == nil {
			continue
		}

		mapping[name] = struct{}{}
		fields = append(fields, field)
	}
	return fields
}

func getFields(t reflect.Type, tag ...string) []FieldAccessor {
	return _getFields(reflect2.Type2(t).(reflect2.StructType), tag, map[string]struct{}{}, nil)
}

var structFieldMapCache sync.Map

func getFieldMap(t reflect.Type, tag ...string) map[string]FieldAccessor {
	if fieldMap, ok := structFieldMapCache.Load(t); ok {
		return fieldMap.(map[string]FieldAccessor)
	}
	fields := getFields(t, tag...)
	fieldMap := make(map[string]FieldAccessor, len(fields))
	for _, field := range fields {
		fieldMap[field.Alias] = field
	}
	structFieldMapCache.Store(t, fieldMap)
	return fieldMap
}

type structInfo struct {
	name   string
	names  []string
	t      *reflect2.UnsafeStructType
	fields map[string]FieldAccessor
}

func makeStructInfo(name string, names []string, t reflect.Type) (info structInfo) {
	info.name = name
	info.names = names
	typ := GetStructType(name)
	if typ == nil {
		for t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		if t.Kind() == reflect.Struct {
			typ = t
		}
	}
	if typ != nil {
		info.t = reflect2.Type2(typ).(*reflect2.UnsafeStructType)
		info.fields = getFieldMap(typ)
	}
	return
}

// ReadStruct reads struct type.
func (dec *Decoder) ReadStruct(t reflect.Type) {
	name := dec.ReadSafeString()
	count := dec.ReadInt()
	names := make([]string, count)
	for i := 0; i < count; i++ {
		dec.decodeString(stringType, dec.NextByte(), &names[i])
	}
	dec.Skip()
	dec.ref = append(dec.ref, makeStructInfo(name, names, t))
}

func (dec *Decoder) getStructInfo(index int) structInfo {
	return dec.ref[index]
}

var structTypeMap sync.Map

// Register the type of the proto with tag.
func Register(proto interface{}, tag ...string) {
	RegisterName("", proto, tag...)
}

// RegisterName the type of the proto with name & tag.
func RegisterName(alias string, proto interface{}, tag ...string) {
	t := reflect.TypeOf(proto)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic(fmt.Sprintf("hprose/io: invalid type: %s", t.String()))
	}
	name := t.Name()
	if alias != "" {
		name = alias
	}
	structTypeMap.Store(name, t)
	if name == "" {
		newAnonymousStructEncoder(t, tag...)
		newAnonymousStructDecoder(t, tag...)
	} else {
		newNamedStructEncoder(t, name, tag...)
		newNamedStructDecoder(t, tag...)
	}
}

// GetStructType by alias.
func GetStructType(alias string) reflect.Type {
	if t, ok := structTypeMap.Load(alias); ok {
		return t.(reflect.Type)
	}
	return nil
}
