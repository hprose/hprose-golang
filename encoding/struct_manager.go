/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/struct_manager.go                               |
|                                                          |
| LastModified: Apr 27, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

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

func _getFields(t reflect2.StructType, tags []string, mapping map[string]bool, fields []FieldAccessor) []FieldAccessor {
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
		if mapping[name] {
			panic(fmt.Sprintf("hprose/encoding: ambiguous fields with the same name or alias: %s", name))
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

		mapping[name] = true
		fields = append(fields, field)
	}
	return fields
}

func getFields(t reflect.Type, tag ...string) []FieldAccessor {
	return _getFields(reflect2.Type2(t).(reflect2.StructType), tag, map[string]bool{}, nil)
}

var structFieldMapCache sync.Map

func getFieldMap(t reflect.Type) map[string]FieldAccessor {
	if fieldMap, ok := structFieldMapCache.Load(t); ok {
		return fieldMap.(map[string]FieldAccessor)
	}
	fields := getFields(t)
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

func makeStructInfo(name string, names []string) (info structInfo) {
	info.name = name
	info.names = names
	if t := GetStructType(name); t != nil {
		info.t = reflect2.Type2(t).(*reflect2.UnsafeStructType)
		info.fields = getFieldMap(t)
	}
	return
}

// ReadStruct reads struct type.
func (dec *Decoder) ReadStruct() {
	name := dec.ReadSafeString()
	count := dec.ReadInt()
	names := make([]string, count)
	for i := 0; i < count; i++ {
		names[i] = dec.decodeString(stringType, dec.NextByte())
	}
	dec.Skip()
	dec.ref = append(dec.ref, makeStructInfo(name, names))
}

func (dec *Decoder) getStructInfo(index int) structInfo {
	return dec.ref[index]
}

var structTypeMap sync.Map

// Register the type of the proto with tag.
func Register(proto interface{}, tag ...string) {
	RegisterAlias(proto, "", tag...)
}

// RegisterAlias the type of the proto with alias & tag.
func RegisterAlias(proto interface{}, alias string, tag ...string) {
	t := reflect.TypeOf(proto)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic(fmt.Sprintf("hprose/encoding: invalid type: %s", t.String()))
	}
	name := t.Name()
	if alias != "" {
		name = alias
	}
	structTypeMap.Store(name, t)
	if name == "" {
		newAnonymousStructEncoder(t, tag...)
	} else {
		newStructEncoder(t, name, tag...)
	}
	newStructDecoder(t)
}

// GetStructType by alias.
func GetStructType(alias string) reflect.Type {
	if t, ok := structTypeMap.Load(alias); ok {
		return t.(reflect.Type)
	}
	return nil
}
