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
 * io/interface_decoder.go                                *
 *                                                        *
 * hprose interface decoder for Go.                       *
 *                                                        *
 * LastModified: Oct 25, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import (
	"errors"
	"math"
	"reflect"
)

var reflectDigits = [...]reflect.Value{
	reflect.ValueOf(0),
	reflect.ValueOf(1),
	reflect.ValueOf(2),
	reflect.ValueOf(3),
	reflect.ValueOf(4),
	reflect.ValueOf(5),
	reflect.ValueOf(6),
	reflect.ValueOf(7),
	reflect.ValueOf(8),
	reflect.ValueOf(9),
}

func readNilAsInterface(r *Reader, v reflect.Value) {
	if v.IsNil() {
		return
	}
	v.Set(reflect.Zero(v.Type()))
}

func readEmptyAsInterface(r *Reader, v reflect.Value) {
	v.Set(reflect.ValueOf(""))
}

func readFalseAsInterface(r *Reader, v reflect.Value) {
	v.Set(reflect.ValueOf(false))
}

func readTrueAsInterface(r *Reader, v reflect.Value) {
	v.Set(reflect.ValueOf(true))
}

func readNaNAsInterface(r *Reader, v reflect.Value) {
	v.Set(reflect.ValueOf(math.NaN()))
}

func readInfAsInterface(r *Reader, v reflect.Value) {
	v.Set(reflect.ValueOf(r.readInf()))
}

func readIntAsInterface(r *Reader, v reflect.Value) {
	v.Set(reflect.ValueOf(r.readInt()))
}

func readLongAsInterface(r *Reader, v reflect.Value) {
	v.Set(reflect.ValueOf(r.ReadBigIntWithoutTag()))
}

func readFloatAsInterface(r *Reader, v reflect.Value) {
	v.Set(reflect.ValueOf(r.readFloat64()))
}

func readUTF8CharAsInterface(r *Reader, v reflect.Value) {
	v.Set(reflect.ValueOf(readUTF8CharAsString(r)))
}

func readStringAsInterface(r *Reader, v reflect.Value) {
	v.Set(reflect.ValueOf(r.ReadStringWithoutTag()))
}

func readBytesAsInterface(r *Reader, v reflect.Value) {
	v.Set(reflect.ValueOf(r.ReadBytesWithoutTag()))
}

func readGUIDAsInterface(r *Reader, v reflect.Value) {
	v.Set(reflect.ValueOf(readGUIDAsString(r)))
}

func readDateTimeAsInterface(r *Reader, v reflect.Value) {
	v.Set(reflect.ValueOf(r.ReadDateTimeWithoutTag()))
}

func readTimeAsInterface(r *Reader, v reflect.Value) {
	v.Set(reflect.ValueOf(r.ReadTimeWithoutTag()))
}

func readListAsInterface(r *Reader, v reflect.Value) {
	var slice []interface{}
	sv := reflect.ValueOf(&slice).Elem()
	readListAsSlice(r, sv)
	v.Set(sv)
}

func readMapAsInterface(r *Reader, v reflect.Value) {
	var mv reflect.Value
	if r.JSONCompatible {
		var m map[string]interface{}
		mv = reflect.ValueOf(&m).Elem()
	} else {
		var m map[interface{}]interface{}
		mv = reflect.ValueOf(&m).Elem()
	}
	readMap(r, mv)
	v.Set(mv)
}

func readStructMeta(r *Reader, v reflect.Value) {
	structName := r.readString()
	structType := v.Type()
	if structType.Kind() != reflect.Struct {
		structType = GetStructType(structName)
	}

	count := r.ReadCount()
	fields := make([]*fieldCache, count)

	if structType == nil {
		for i := 0; i < count; i++ {
			fields[i] = &fieldCache{
				Alias: r.ReadString(),
				Type:  interfaceType,
			}
		}
	} else {
		structCache := getStructCache(structType)
		fieldMap := structCache.FieldMap
		for i := 0; i < count; i++ {
			fields[i] = fieldMap[r.ReadString()]
		}
	}
	r.structTypeRef = append(r.structTypeRef, structType)
	r.fieldsRef = append(r.fieldsRef, fields)
	r.readByte()
	r.ReadValue(v)
}

func readStructData(r *Reader, v reflect.Value) {
	index := r.ReadCount()
	if v.Kind() == reflect.Interface {
		typ := r.structTypeRef[index]
		if typ == nil {
			x := map[string]interface{}{}
			v2 := reflect.ValueOf(x)
			readStructAsMapByIndex(r, v2, index)
			v.Set(v2)
			return
		}
		if !reflect.PtrTo(typ).Implements(v.Type()) {
			panic(errors.New("*" + typ.String() + " does not implements " + v.Type().String() + " interface"))
		} else {
			ptr := reflect.New(typ)
			v.Set(ptr)
			v = ptr.Elem()
		}
	}
	fields := r.fieldsRef[index]
	count := len(fields)
	if !r.Simple {
		setReaderRef(r, v)
	}
	for i := 0; i < count; i++ {
		if field := fields[i]; field != nil {
			f := v.FieldByIndex(field.Index)
			r.ReadValue(f)
		} else {
			var x interface{}
			r.Unserialize(&x)
		}
	}
	r.readByte()
}

func readRefAsInterface(r *Reader, v reflect.Value) {
	iv := reflect.ValueOf(r.readRef())
	t := v.Type()
	it := iv.Type()
	if it.AssignableTo(t) {
		v.Set(iv)
	} else if it.ConvertibleTo(t) {
		v.Set(iv.Convert(t))
	} else {
		panic(errors.New(it.String() +
			" cannot be converted to type" +
			t.String()))
	}
}

var interfaceDecoders = [256]func(r *Reader, v reflect.Value){
	'0':         func(r *Reader, v reflect.Value) { v.Set(reflectDigits[0]) },
	'1':         func(r *Reader, v reflect.Value) { v.Set(reflectDigits[1]) },
	'2':         func(r *Reader, v reflect.Value) { v.Set(reflectDigits[2]) },
	'3':         func(r *Reader, v reflect.Value) { v.Set(reflectDigits[3]) },
	'4':         func(r *Reader, v reflect.Value) { v.Set(reflectDigits[4]) },
	'5':         func(r *Reader, v reflect.Value) { v.Set(reflectDigits[5]) },
	'6':         func(r *Reader, v reflect.Value) { v.Set(reflectDigits[6]) },
	'7':         func(r *Reader, v reflect.Value) { v.Set(reflectDigits[7]) },
	'8':         func(r *Reader, v reflect.Value) { v.Set(reflectDigits[8]) },
	'9':         func(r *Reader, v reflect.Value) { v.Set(reflectDigits[9]) },
	TagNull:     readNilAsInterface,
	TagEmpty:    readEmptyAsInterface,
	TagFalse:    readFalseAsInterface,
	TagTrue:     readTrueAsInterface,
	TagNaN:      readNaNAsInterface,
	TagInfinity: readInfAsInterface,
	TagInteger:  readIntAsInterface,
	TagLong:     readLongAsInterface,
	TagDouble:   readFloatAsInterface,
	TagUTF8Char: readUTF8CharAsInterface,
	TagString:   readStringAsInterface,
	TagBytes:    readBytesAsInterface,
	TagGUID:     readGUIDAsInterface,
	TagDate:     readDateTimeAsInterface,
	TagTime:     readTimeAsInterface,
	TagList:     readListAsInterface,
	TagMap:      readMapAsInterface,
	TagClass:    readStructMeta,
	TagObject:   readStructData,
	TagRef:      readRefAsInterface,
}

func interfaceDecoder(r *Reader, v reflect.Value, tag byte) {
	decoder := interfaceDecoders[tag]
	if decoder != nil {
		decoder(r, v)
		return
	}
	castError(tag, "interface{}")
}
