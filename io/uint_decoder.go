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
 * io/uint_decoder.go                                     *
 *                                                        *
 * hprose uint decoder for Go.                            *
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

func readUint64(r *Reader) uint64 {
	return r.readUint64(TagSemicolon)
}

func readFloat64AsUint(r *Reader) uint64 {
	return uint64(r.readFloat64())
}

func stringToUint(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func readUTF8CharAsUint(r *Reader) uint64 {
	return stringToUint(util.ByteString(r.readUTF8Slice(1)))
}

func readStringAsUint(r *Reader) uint64 {
	return stringToUint(r.ReadStringWithoutTag())
}

func readDateTimeAsUint(r *Reader) uint64 {
	return uint64(r.ReadDateTimeWithoutTag().UnixNano())
}

func readTimeAsUint(r *Reader) uint64 {
	return uint64(r.ReadTimeWithoutTag().UnixNano())
}

func readRefAsUint(r *Reader) uint64 {
	ref := r.readRef()
	if str, ok := ref.(string); ok {
		return stringToUint(str)
	}
	if t, ok := ref.(*time.Time); ok {
		return uint64(t.UnixNano())
	}
	panic(errors.New("value of type " +
		reflect.TypeOf(ref).String() +
		" cannot be converted to type uint64"))
}

var uintDecoders = [256]func(r *Reader) uint64{
	'0':         func(r *Reader) uint64 { return 0 },
	'1':         func(r *Reader) uint64 { return 1 },
	'2':         func(r *Reader) uint64 { return 2 },
	'3':         func(r *Reader) uint64 { return 3 },
	'4':         func(r *Reader) uint64 { return 4 },
	'5':         func(r *Reader) uint64 { return 5 },
	'6':         func(r *Reader) uint64 { return 6 },
	'7':         func(r *Reader) uint64 { return 7 },
	'8':         func(r *Reader) uint64 { return 8 },
	'9':         func(r *Reader) uint64 { return 9 },
	TagNull:     func(r *Reader) uint64 { return 0 },
	TagEmpty:    func(r *Reader) uint64 { return 0 },
	TagFalse:    func(r *Reader) uint64 { return 0 },
	TagTrue:     func(r *Reader) uint64 { return 1 },
	TagInteger:  readUint64,
	TagLong:     readUint64,
	TagDouble:   readFloat64AsUint,
	TagUTF8Char: readUTF8CharAsUint,
	TagString:   readStringAsUint,
	TagDate:     readDateTimeAsUint,
	TagTime:     readTimeAsUint,
	TagRef:      readRefAsUint,
}

func uintDecoder(r *Reader, v reflect.Value, tag byte) {
	decoder := uintDecoders[tag]
	if decoder != nil {
		v.SetUint(decoder(r))
		return
	}
	castError(tag, "uint64")
}
