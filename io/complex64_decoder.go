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
 * io/complex64_decoder.go                                *
 *                                                        *
 * hprose complex64 decoder for Go.                       *
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
	"strconv"

	"github.com/hprose/hprose-golang/util"
)

func readLongAsComplex64(r *Reader) complex64 {
	return complex(float32(r.readLongAsFloat64()), 0)
}

func stringToComplex64(s string) complex64 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		panic(err)
	}
	return complex(float32(f), 0)
}

func readInfinityAsComplex64(r *Reader) complex64 {
	return complex(float32(r.readInf()), 0)
}

func readUTF8CharAsComplex64(r *Reader) complex64 {
	return stringToComplex64(util.ByteString(r.readUTF8Slice(1)))
}

func readStringAsComplex64(r *Reader) complex64 {
	return stringToComplex64(r.ReadStringWithoutTag())
}

func readListAsComplex64(r *Reader) complex64 {
	var floatPair [2]float32
	readListAsArray(r, reflect.ValueOf(&floatPair).Elem())
	return complex(floatPair[0], floatPair[1])
}

func readRefAsComplex64(r *Reader) complex64 {
	ref := r.readRef()
	if str, ok := ref.(string); ok {
		return stringToComplex64(str)
	}
	if v, ok := ref.(reflect.Value); ok {
		if floatPair, ok := v.Interface().([2]float32); ok {
			return complex(floatPair[0], floatPair[1])
		}
	}
	panic(errors.New("value of type " +
		reflect.TypeOf(ref).String() +
		" cannot be converted to type complex64"))
}

var complex64Decoders = [256]func(r *Reader) complex64{
	'0':         func(r *Reader) complex64 { return 0 },
	'1':         func(r *Reader) complex64 { return 1 },
	'2':         func(r *Reader) complex64 { return 2 },
	'3':         func(r *Reader) complex64 { return 3 },
	'4':         func(r *Reader) complex64 { return 4 },
	'5':         func(r *Reader) complex64 { return 5 },
	'6':         func(r *Reader) complex64 { return 6 },
	'7':         func(r *Reader) complex64 { return 7 },
	'8':         func(r *Reader) complex64 { return 8 },
	'9':         func(r *Reader) complex64 { return 9 },
	TagNull:     func(r *Reader) complex64 { return 0 },
	TagEmpty:    func(r *Reader) complex64 { return 0 },
	TagFalse:    func(r *Reader) complex64 { return 0 },
	TagTrue:     func(r *Reader) complex64 { return 1 },
	TagNaN:      func(r *Reader) complex64 { return complex(float32(math.NaN()), 0) },
	TagInfinity: readInfinityAsComplex64,
	TagInteger:  readLongAsComplex64,
	TagLong:     readLongAsComplex64,
	TagDouble:   func(r *Reader) complex64 { return complex(r.readFloat32(), 0) },
	TagUTF8Char: readUTF8CharAsComplex64,
	TagString:   readStringAsComplex64,
	TagList:     readListAsComplex64,
	TagRef:      readRefAsComplex64,
}

func complex64Decoder(r *Reader, v reflect.Value, tag byte) {
	decoder := complex64Decoders[tag]
	if decoder != nil {
		v.SetComplex(complex128(decoder(r)))
		return
	}
	castError(tag, "complex64")
}
