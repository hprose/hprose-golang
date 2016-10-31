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
 * io/string_decoder.go                                   *
 *                                                        *
 * hprose string decoder for Go.                          *
 *                                                        *
 * LastModified: Oct 15, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import (
	"errors"
	"reflect"
	"time"
)

func readNumberAsString(r *Reader) string {
	return string(r.readUntil(TagSemicolon))
}

func readInfAsString(r *Reader) string {
	if r.readByte() == TagNeg {
		return "-Inf"
	}
	return "+Inf"
}

func readUTF8CharAsString(r *Reader) string {
	return string(r.readUTF8Slice(1))
}

func readBytesAsString(r *Reader) (str string) {
	l := r.readLength()
	str = string(r.Next(l))
	r.readByte()
	if !r.Simple {
		setReaderRef(r, str)
	}
	return
}

func readGUIDAsString(r *Reader) (str string) {
	r.readByte()
	str = string(r.Next(36))
	r.readByte()
	if !r.Simple {
		setReaderRef(r, str)
	}
	return
}

func readDateTimeAsString(r *Reader) string {
	return r.ReadDateTimeWithoutTag().String()
}

func readTimeAsString(r *Reader) string {
	return r.ReadTimeWithoutTag().String()
}

func readRefAsString(r *Reader) string {
	ref := r.readRef()
	if str, ok := ref.(string); ok {
		return str
	}
	if t, ok := ref.(*time.Time); ok {
		return t.String()
	}
	panic(errors.New("value of type " +
		reflect.TypeOf(ref).String() +
		" cannot be converted to type string"))
}

var stringDecoders = [256]func(r *Reader) string{
	'0':         func(r *Reader) string { return "0" },
	'1':         func(r *Reader) string { return "1" },
	'2':         func(r *Reader) string { return "2" },
	'3':         func(r *Reader) string { return "3" },
	'4':         func(r *Reader) string { return "4" },
	'5':         func(r *Reader) string { return "5" },
	'6':         func(r *Reader) string { return "6" },
	'7':         func(r *Reader) string { return "7" },
	'8':         func(r *Reader) string { return "8" },
	'9':         func(r *Reader) string { return "9" },
	TagNull:     func(r *Reader) string { return "" },
	TagEmpty:    func(r *Reader) string { return "" },
	TagFalse:    func(r *Reader) string { return "false" },
	TagTrue:     func(r *Reader) string { return "true" },
	TagNaN:      func(r *Reader) string { return "NaN" },
	TagInfinity: readInfAsString,
	TagInteger:  readNumberAsString,
	TagLong:     readNumberAsString,
	TagDouble:   readNumberAsString,
	TagUTF8Char: readUTF8CharAsString,
	TagString:   func(r *Reader) string { return r.ReadStringWithoutTag() },
	TagBytes:    readBytesAsString,
	TagGUID:     readGUIDAsString,
	TagDate:     readDateTimeAsString,
	TagTime:     readTimeAsString,
	TagRef:      readRefAsString,
}

func stringDecoder(r *Reader, v reflect.Value, tag byte) {
	decoder := stringDecoders[tag]
	if decoder != nil {
		v.SetString(decoder(r))
		return
	}
	castError(tag, "string")
}
