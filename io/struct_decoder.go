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
 * io/struct_decoder.go                                   *
 *                                                        *
 * hprose struct decoder for Go.                          *
 *                                                        *
 * LastModified: Oct 25, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import (
	"container/list"
	"errors"
	"math/big"
	"reflect"
	"strconv"
	"time"
	"unsafe"

	"github.com/hprose/hprose-golang/util"
)

func readDigitAsStruct(r *Reader, v reflect.Value, tag byte) {
	typ := (*reflectValue)(unsafe.Pointer(&v)).typ
	switch typ {
	case bigIntType:
		v.Set(reflect.ValueOf(*big.NewInt(int64(tag - '0'))))
	case bigRatType:
		v.Set(reflect.ValueOf(*big.NewRat(int64(tag-'0'), 1)))
	case bigFloatType:
		v.Set(reflect.ValueOf(*big.NewFloat(float64(tag - '0'))))
	case timeType:
		v.Set(reflect.ValueOf(time.Unix(int64(tag-'0'), 0)))
	default:
		castError(tag, v.Type().String())
	}
}

func readIntAsStruct(r *Reader, v reflect.Value, tag byte) {
	i := r.readInt64(TagSemicolon)
	typ := (*reflectValue)(unsafe.Pointer(&v)).typ
	switch typ {
	case bigIntType:
		v.Set(reflect.ValueOf(*big.NewInt(i)))
	case bigRatType:
		v.Set(reflect.ValueOf(*big.NewRat(i, 1)))
	case bigFloatType:
		v.Set(reflect.ValueOf(*big.NewFloat(float64(i))))
	case timeType:
		v.Set(reflect.ValueOf(time.Unix(i, 0)))
	default:
		castError(tag, v.Type().String())
	}
}

func readLongAsStruct(r *Reader, v reflect.Value, tag byte) {
	i := util.ByteString(r.readUntil(TagSemicolon))
	typ := (*reflectValue)(unsafe.Pointer(&v)).typ
	switch typ {
	case bigIntType:
		if bi, ok := new(big.Int).SetString(i, 10); ok {
			v.Set(reflect.ValueOf(*bi))
		}
	case bigRatType:
		if br, ok := new(big.Rat).SetString(i); ok {
			v.Set(reflect.ValueOf(*br))
		}
	case bigFloatType:
		if bf, _, err := new(big.Float).Parse(i, 10); err == nil {
			v.Set(reflect.ValueOf(*bf))
		} else {
			panic(err)
		}
	case timeType:
		if unix, err := strconv.ParseInt(i, 10, 64); err == nil {
			v.Set(reflect.ValueOf(time.Unix(unix, 0)))
		} else {
			panic(err)
		}
	default:
		castError(tag, v.Type().String())
	}
}

func readDoubleAsStruct(r *Reader, v reflect.Value, tag byte) {
	f := util.ByteString(r.readUntil(TagSemicolon))
	typ := (*reflectValue)(unsafe.Pointer(&v)).typ
	switch typ {
	case bigFloatType:
		if bf, _, err := new(big.Float).Parse(f, 10); err == nil {
			v.Set(reflect.ValueOf(*bf))
		} else {
			panic(err)
		}
	case timeType:
		if unix, err := strconv.ParseFloat(f, 10); err == nil {
			sec := int64(unix)
			nsec := int64((unix - float64(sec)) * 1000000000)
			v.Set(reflect.ValueOf(time.Unix(sec, nsec)))
		} else {
			panic(err)
		}
	default:
		castError(tag, v.Type().String())
	}
}

const timeStringFormat = "2006-01-02 15:04:05.999999999 -0700 MST"

func readStringAsStruct(str string, v reflect.Value, tag byte) {
	typ := (*reflectValue)(unsafe.Pointer(&v)).typ
	switch typ {
	case bigIntType:
		if bi, ok := new(big.Int).SetString(str, 10); ok {
			v.Set(reflect.ValueOf(*bi))
		}
	case bigRatType:
		if br, ok := new(big.Rat).SetString(str); ok {
			v.Set(reflect.ValueOf(*br))
		}
	case bigFloatType:
		if bf, _, err := new(big.Float).Parse(str, 10); err == nil {
			v.Set(reflect.ValueOf(*bf))
		} else {
			panic(err)
		}
	case timeType:
		if t, err := time.Parse(timeStringFormat, str); err == nil {
			v.Set(reflect.ValueOf(t))
		} else {
			panic(err)
		}
	default:
		castError(tag, v.Type().String())
	}
}

func readStringStruct(r *Reader, v reflect.Value, tag byte) {
	readStringAsStruct(r.ReadStringWithoutTag(), v, tag)
}

func readTimeAsStruct(t time.Time, v reflect.Value, tag byte) {
	typ := (*reflectValue)(unsafe.Pointer(&v)).typ
	switch typ {
	case bigIntType:
		v.Set(reflect.ValueOf(*new(big.Int).SetInt64(t.Unix())))
	case bigRatType:
		v.Set(reflect.ValueOf(*new(big.Rat).SetInt64(t.Unix())))
	case bigFloatType:
		ft := float64(t.Unix()) + float64(t.Nanosecond())/1000000000
		v.Set(reflect.ValueOf(*new(big.Float).SetFloat64(ft)))
	case timeType:
		v.Set(reflect.ValueOf(t))
	default:
		castError(tag, v.Type().String())
	}
}

func readDateTimeStruct(r *Reader, v reflect.Value, tag byte) {
	readTimeAsStruct(r.ReadDateTimeWithoutTag(), v, tag)
}

func readTimeStruct(r *Reader, v reflect.Value, tag byte) {
	readTimeAsStruct(r.ReadTimeWithoutTag(), v, tag)
}

func readListAsStruct(r *Reader, v reflect.Value, tag byte) {
	typ := (*reflectValue)(unsafe.Pointer(&v)).typ
	if typ != listType {
		castError(tag, v.Type().String())
	}
	lst := list.New()
	l := r.ReadCount()
	if !r.Simple {
		setReaderRef(r, v)
	}
	for i := 0; i < l; i++ {
		var e interface{}
		r.Unserialize(&e)
		lst.PushBack(e)
	}
	r.readByte()
	v.Set(reflect.ValueOf(*lst))
}

func readMapAsStruct(r *Reader, v reflect.Value, tag byte) {
	structCache := getStructCache(v.Type())
	fieldMap := structCache.FieldMap
	l := r.ReadCount()
	if !r.Simple {
		setReaderRef(r, v)
	}
	for i := 0; i < l; i++ {
		key := r.ReadString()
		if field, ok := fieldMap[key]; ok {
			f := v.FieldByIndex(field.Index)
			r.ReadValue(f)
		} else {
			var x interface{}
			r.Unserialize(&x)
		}
	}
	r.readByte()
}

func readRefAsStruct(r *Reader, v reflect.Value, tag byte) {
	ref := r.readRef()
	if str, ok := ref.(string); ok {
		readStringAsStruct(str, v, tag)
		return
	}
	if t, ok := ref.(*time.Time); ok {
		readTimeAsStruct(*t, v, tag)
		return
	}
	if r, ok := ref.(reflect.Value); ok {
		v.Set(r)
		return
	}
	panic(errors.New("value of type " +
		reflect.TypeOf(ref).String() +
		" cannot be converted to type map"))
}

var structDecoders = [256]func(r *Reader, v reflect.Value, tag byte){
	TagNull:    func(r *Reader, v reflect.Value, tag byte) { nilDecoder(r, v) },
	'0':        readDigitAsStruct,
	'1':        readDigitAsStruct,
	'2':        readDigitAsStruct,
	'3':        readDigitAsStruct,
	'4':        readDigitAsStruct,
	'5':        readDigitAsStruct,
	'6':        readDigitAsStruct,
	'7':        readDigitAsStruct,
	'8':        readDigitAsStruct,
	'9':        readDigitAsStruct,
	TagInteger: readIntAsStruct,
	TagLong:    readLongAsStruct,
	TagDouble:  readDoubleAsStruct,
	TagString:  readStringStruct,
	TagDate:    readDateTimeStruct,
	TagTime:    readTimeStruct,
	TagList:    readListAsStruct,
	TagMap:     readMapAsStruct,
	TagClass:   func(r *Reader, v reflect.Value, tag byte) { readStructMeta(r, v) },
	TagObject:  func(r *Reader, v reflect.Value, tag byte) { readStructData(r, v) },
	TagRef:     readRefAsStruct,
}

func structDecoder(r *Reader, v reflect.Value, tag byte) {
	if (*reflectValue)(unsafe.Pointer(&v)).typ == reflectValueType {
		rv := reflect.New(interfaceType).Elem()
		interfaceDecoder(r, rv, tag)
		v.Set(reflect.ValueOf(rv))
		return
	}
	decoder := structDecoders[tag]
	if decoder != nil {
		decoder(r, v, tag)
		return
	}
	castError(tag, v.Type().String())
}
