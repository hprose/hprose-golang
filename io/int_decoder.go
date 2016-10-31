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
 * io/int_decoder.go                                      *
 *                                                        *
 * hprose int decoder for Go.                             *
 *                                                        *
 * LastModified: Oct 15, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import (
	"errors"
	"reflect"
	"strconv"
	"time"

	"github.com/hprose/hprose-golang/util"
)

func readInt64(r *Reader) int64 {
	return r.readInt64(TagSemicolon)
}

func readFloat64AsInt(r *Reader) int64 {
	return int64(r.readFloat64())
}

func stringToInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func readUTF8CharAsInt(r *Reader) int64 {
	return stringToInt(util.ByteString(r.readUTF8Slice(1)))
}

func readStringAsInt(r *Reader) int64 {
	return stringToInt(r.ReadStringWithoutTag())
}

func readDateTimeAsInt(r *Reader) int64 {
	return r.ReadDateTimeWithoutTag().UnixNano()
}

func readTimeAsInt(r *Reader) int64 {
	return r.ReadTimeWithoutTag().UnixNano()
}

func readRefAsInt(r *Reader) int64 {
	ref := r.readRef()
	if str, ok := ref.(string); ok {
		return stringToInt(str)
	}
	if t, ok := ref.(*time.Time); ok {
		return t.UnixNano()
	}
	panic(errors.New("value of type " +
		reflect.TypeOf(ref).String() +
		" cannot be converted to type int64"))
}

var intDecoders = [256]func(r *Reader) int64{
	'0':         func(r *Reader) int64 { return 0 },
	'1':         func(r *Reader) int64 { return 1 },
	'2':         func(r *Reader) int64 { return 2 },
	'3':         func(r *Reader) int64 { return 3 },
	'4':         func(r *Reader) int64 { return 4 },
	'5':         func(r *Reader) int64 { return 5 },
	'6':         func(r *Reader) int64 { return 6 },
	'7':         func(r *Reader) int64 { return 7 },
	'8':         func(r *Reader) int64 { return 8 },
	'9':         func(r *Reader) int64 { return 9 },
	TagNull:     func(r *Reader) int64 { return 0 },
	TagEmpty:    func(r *Reader) int64 { return 0 },
	TagFalse:    func(r *Reader) int64 { return 0 },
	TagTrue:     func(r *Reader) int64 { return 1 },
	TagInteger:  readInt64,
	TagLong:     readInt64,
	TagDouble:   readFloat64AsInt,
	TagUTF8Char: readUTF8CharAsInt,
	TagString:   readStringAsInt,
	TagDate:     readDateTimeAsInt,
	TagTime:     readTimeAsInt,
	TagRef:      readRefAsInt,
}

func intDecoder(r *Reader, v reflect.Value, tag byte) {
	decoder := intDecoders[tag]
	if decoder != nil {
		v.SetInt(decoder(r))
		return
	}
	castError(tag, "int64")
}
