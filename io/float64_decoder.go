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
 * io/float64_decoder.go                                  *
 *                                                        *
 * hprose float64 decoder for Go.                         *
 *                                                        *
 * LastModified: Oct 15, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import (
	"errors"
	"math"
	"reflect"
	"strconv"
	"time"

	"github.com/hprose/hprose-golang/util"
)

func readLongAsFloat64(r *Reader) float64 {
	return r.readLongAsFloat64()
}

func stringToFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}

func readInfinityAsFloat64(r *Reader) float64 {
	return r.readInf()
}

func readUTF8CharAsFloat64(r *Reader) float64 {
	return stringToFloat64(util.ByteString(r.readUTF8Slice(1)))
}

func readStringAsFloat64(r *Reader) float64 {
	return stringToFloat64(r.ReadStringWithoutTag())
}

func timeToFloat64(t time.Time) float64 {
	return float64(t.Unix()) + float64(t.Nanosecond())/1000000000
}

func readDateTimeAsFloat64(r *Reader) float64 {
	return timeToFloat64(r.ReadDateTimeWithoutTag())
}

func readTimeAsFloat64(r *Reader) float64 {
	return timeToFloat64(r.ReadTimeWithoutTag())
}

func readRefAsFloat64(r *Reader) float64 {
	ref := r.readRef()
	if str, ok := ref.(string); ok {
		return stringToFloat64(str)
	}
	if t, ok := ref.(*time.Time); ok {
		return timeToFloat64(*t)
	}
	panic(errors.New("value of type " +
		reflect.TypeOf(ref).String() +
		" cannot be converted to type float64"))
}

var float64Decoders = [256]func(r *Reader) float64{
	'0':         func(r *Reader) float64 { return 0 },
	'1':         func(r *Reader) float64 { return 1 },
	'2':         func(r *Reader) float64 { return 2 },
	'3':         func(r *Reader) float64 { return 3 },
	'4':         func(r *Reader) float64 { return 4 },
	'5':         func(r *Reader) float64 { return 5 },
	'6':         func(r *Reader) float64 { return 6 },
	'7':         func(r *Reader) float64 { return 7 },
	'8':         func(r *Reader) float64 { return 8 },
	'9':         func(r *Reader) float64 { return 9 },
	TagNull:     func(r *Reader) float64 { return 0 },
	TagEmpty:    func(r *Reader) float64 { return 0 },
	TagFalse:    func(r *Reader) float64 { return 0 },
	TagTrue:     func(r *Reader) float64 { return 1 },
	TagNaN:      func(r *Reader) float64 { return math.NaN() },
	TagInfinity: readInfinityAsFloat64,
	TagInteger:  readLongAsFloat64,
	TagLong:     readLongAsFloat64,
	TagDouble:   func(r *Reader) float64 { return r.readFloat64() },
	TagUTF8Char: readUTF8CharAsFloat64,
	TagString:   readStringAsFloat64,
	TagDate:     readDateTimeAsFloat64,
	TagTime:     readTimeAsFloat64,
	TagRef:      readRefAsFloat64,
}

func float64Decoder(r *Reader, v reflect.Value, tag byte) {
	decoder := float64Decoders[tag]
	if decoder != nil {
		v.SetFloat(decoder(r))
		return
	}
	castError(tag, "float64")
}
