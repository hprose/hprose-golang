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
 * io/bool_decoder.go                                     *
 *                                                        *
 * hprose bool decoder for Go.                            *
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

	"github.com/hprose/hprose-golang/util"
)

func readBoolFalse(r *Reader) bool {
	return false
}

func readBoolTrue(r *Reader) bool {
	return true
}

func readNumberAsBool(r *Reader) bool {
	bytes := r.readUntil(TagSemicolon)
	if len(bytes) == 0 {
		return true
	}
	if len(bytes) == 1 {
		return bytes[0] != '0'
	}
	return true
}

func readInfinityAsBool(r *Reader) bool {
	r.readInf()
	return true
}

func readUTF8CharAsBool(r *Reader) bool {
	b, err := strconv.ParseBool(util.ByteString(r.readUTF8Slice(1)))
	if err != nil {
		panic(err)
	}
	return b
}

func readStringAsBool(r *Reader) bool {
	b, err := strconv.ParseBool(r.ReadStringWithoutTag())
	if err != nil {
		panic(err)
	}
	return b
}

func readRefAsBool(r *Reader) bool {
	ref := r.readRef()
	if str, ok := ref.(string); ok {
		b, err := strconv.ParseBool(str)
		if err != nil {
			panic(err)
		}
		return b
	}
	panic(errors.New("value of type " +
		reflect.TypeOf(ref).String() +
		" cannot be converted to type bool"))
}

var boolDecoders = [256]func(r *Reader) bool{
	'0':         readBoolFalse,
	'1':         readBoolTrue,
	'2':         readBoolTrue,
	'3':         readBoolTrue,
	'4':         readBoolTrue,
	'5':         readBoolTrue,
	'6':         readBoolTrue,
	'7':         readBoolTrue,
	'8':         readBoolTrue,
	'9':         readBoolTrue,
	TagInteger:  readNumberAsBool,
	TagLong:     readNumberAsBool,
	TagDouble:   readNumberAsBool,
	TagNull:     readBoolFalse,
	TagEmpty:    readBoolFalse,
	TagFalse:    readBoolFalse,
	TagTrue:     readBoolTrue,
	TagNaN:      readBoolTrue,
	TagInfinity: readInfinityAsBool,
	TagUTF8Char: readUTF8CharAsBool,
	TagString:   readStringAsBool,
	TagRef:      readRefAsBool,
}

func boolDecoder(r *Reader, v reflect.Value, tag byte) {
	decoder := boolDecoders[tag]
	if decoder != nil {
		v.SetBool(decoder(r))
		return
	}
	castError(tag, "bool")
}
