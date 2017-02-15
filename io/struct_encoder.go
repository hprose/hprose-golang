/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * hprose/struct_encoder.go                               *
 *                                                        *
 * hprose struct encoder for Go.                          *
 *                                                        *
 * LastModified: Feb 15, 2017                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import (
	"reflect"
	"strings"
	"sync"
	"unsafe"

	"github.com/hprose/hprose-golang/util"
)

type fieldCache struct {
	Name  string
	Alias string
	Index []int
	Type  reflect.Type
	Kind  reflect.Kind
}

type structCache struct {
	Alias    string
	Tag      string
	Fields   []*fieldCache
	FieldMap map[string]*fieldCache
	Data     []byte
}

var structTypeCache = map[uintptr]*structCache{}
var structTypeCacheLocker = sync.RWMutex{}

var structTypes = map[string]reflect.Type{}
var structTypesLocker = sync.RWMutex{}

func getFieldAlias(f *reflect.StructField, tag string) (alias string) {
	fname := f.Name
	if fname != "" && 'A' <= fname[0] && fname[0] <= 'Z' {
		if tag != "" && f.Tag != "" {
			alias = strings.SplitN(f.Tag.Get(tag), ",", 2)[0]
			alias = strings.TrimSpace(strings.SplitN(alias, ">", 2)[0])
			if alias == "-" {
				return ""
			}
		}
		if alias == "" {
			alias = string(fname[0]-'A'+'a') + fname[1:]
		}
	}
	return alias
}

func getSubFields(t reflect.Type, tag string, index []int) []*fieldCache {
	subFields := getFields(t, tag)
	for _, subField := range subFields {
		subField.Index = append(index, subField.Index...)
	}
	return subFields
}

func getFields(t reflect.Type, tag string) []*fieldCache {
	n := t.NumField()
	fields := make([]*fieldCache, 0, n)
	for i := 0; i < n; i++ {
		f := t.Field(i)
		ft := f.Type
		fkind := ft.Kind()
		if fkind == reflect.Chan ||
			fkind == reflect.Func ||
			fkind == reflect.UnsafePointer {
			continue
		}
		if f.Anonymous {
			if fkind == reflect.Struct {
				subFields := getSubFields(ft, tag, f.Index)
				fields = append(fields, subFields...)
				continue
			}
		}
		alias := getFieldAlias(&f, tag)
		if alias == "" {
			continue
		}
		field := fieldCache{}
		field.Name = f.Name
		field.Alias = alias
		field.Type = ft
		field.Kind = fkind
		field.Index = f.Index
		fields = append(fields, &field)
	}
	return fields
}

func initStructCacheData(cache *structCache) {
	w := &ByteWriter{}
	fields := cache.Fields
	count := len(fields)
	cache.FieldMap = make(map[string]*fieldCache, count)
	w.writeByte(TagClass)
	var buf [20]byte
	w.write(util.GetIntBytes(buf[:], int64(util.UTF16Length(cache.Alias))))
	w.writeByte(TagQuote)
	w.writeString(cache.Alias)
	w.writeByte(TagQuote)
	if count > 0 {
		w.write(util.GetIntBytes(buf[:], int64(count)))
	}
	w.writeByte(TagOpenbrace)
	for _, field := range fields {
		cache.FieldMap[field.Alias] = field
		w.writeByte(TagString)
		w.write(util.GetIntBytes(buf[:], int64(util.UTF16Length(field.Alias))))
		w.writeByte(TagQuote)
		w.writeString(field.Alias)
		w.writeByte(TagQuote)
	}
	w.writeByte(TagClosebrace)
	cache.Data = w.Bytes()
}

func getStructCache(structType reflect.Type) *structCache {
	typ := (*emptyInterface)(unsafe.Pointer(&structType)).ptr
	structTypeCacheLocker.RLock()
	if cache, ok := structTypeCache[typ]; ok {
		structTypeCacheLocker.RUnlock()
		return cache
	}
	structTypeCacheLocker.RUnlock()
	structTypeCacheLocker.Lock()
	cache, ok := structTypeCache[typ]
	if !ok {
		cache = &structCache{}
		cache.Alias = structType.Name()
		cache.Fields = getFields(structType, "")
		initStructCacheData(cache)
		structTypeCache[typ] = cache
		structTypesLocker.Lock()
		structTypes[cache.Alias] = structType
		structTypesLocker.Unlock()
	}
	structTypeCacheLocker.Unlock()
	return cache
}

// Register the type of the proto with alias & tag.
func Register(proto interface{}, alias string, tag ...string) {
	structType := reflect.TypeOf(proto)
	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}
	if structType.Kind() != reflect.Struct {
		panic("invalid type: " + structType.String())
	}
	structTypesLocker.Lock()
	structTypes[alias] = structType
	structTypesLocker.Unlock()

	structTypeCacheLocker.Lock()
	cache := &structCache{Alias: alias}
	if len(tag) == 1 {
		cache.Tag = tag[0]
	}
	cache.Fields = getFields(structType, cache.Tag)
	initStructCacheData(cache)
	structTypeCache[(*emptyInterface)(unsafe.Pointer(&structType)).ptr] = cache
	structTypeCacheLocker.Unlock()
}

// GetStructType by alias.
func GetStructType(alias string) (structType reflect.Type) {
	structTypesLocker.RLock()
	structType = structTypes[alias]
	structTypesLocker.RUnlock()
	return structType
}

// GetAlias of structType
func GetAlias(structType reflect.Type) string {
	return getStructCache(structType).Alias
}

// GetTag by structType.
func GetTag(structType reflect.Type) string {
	return getStructCache(structType).Tag
}
