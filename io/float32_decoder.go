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
 * io/float32_decoder.go                                  *
 *                                                        *
 * hprose float32 decoder for Go.                         *
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

func readLongAsFloat32(r *Reader) float32 {
	return float32(r.readLongAsFloat64())
}

func stringToFloat32(s string) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		panic(err)
	}
	return float32(f)
}

func readInfinityAsFloat32(r *Reader) float32 {
	return float32(r.readInf())
}

func readUTF8CharAsFloat32(r *Reader) float32 {
	return stringToFloat32(util.ByteString(r.readUTF8Slice(1)))
}

func readStringAsFloat32(r *Reader) float32 {
	return stringToFloat32(r.ReadStringWithoutTag())
}

func timeToFloat32(t time.Time) float32 {
	return float32(t.Unix()) + float32(t.Nanosecond())/1000000000
}

func readDateTimeAsFloat32(r *Reader) float32 {
	return timeToFloat32(r.ReadDateTimeWithoutTag())
}

func readTimeAsFloat32(r *Reader) float32 {
	return timeToFloat32(r.ReadTimeWithoutTag())
}

func readRefAsFloat32(r *Reader) float32 {
	ref := r.readRef()
	if str, ok := ref.(string); ok {
		return stringToFloat32(str)
	}
	if t, ok := ref.(*time.Time); ok {
		return timeToFloat32(*t)
	}
	panic(errors.New("value of type " +
		reflect.TypeOf(ref).String() +
		" cannot be converted to type float32"))
}

var float32Decoders = [256]func(r *Reader) float32{
	'0':         func(r *Reader) float32 { return 0 },
	'1':         func(r *Reader) float32 { return 1 },
	'2':         func(r *Reader) float32 { return 2 },
	'3':         func(r *Reader) float32 { return 3 },
	'4':         func(r *Reader) float32 { return 4 },
	'5':         func(r *Reader) float32 { return 5 },
	'6':         func(r *Reader) float32 { return 6 },
	'7':         func(r *Reader) float32 { return 7 },
	'8':         func(r *Reader) float32 { return 8 },
	'9':         func(r *Reader) float32 { return 9 },
	TagNull:     func(r *Reader) float32 { return 0 },
	TagEmpty:    func(r *Reader) float32 { return 0 },
	TagFalse:    func(r *Reader) float32 { return 0 },
	TagTrue:     func(r *Reader) float32 { return 1 },
	TagNaN:      func(r *Reader) float32 { return float32(math.NaN()) },
	TagInfinity: readInfinityAsFloat32,
	TagInteger:  readLongAsFloat32,
	TagLong:     readLongAsFloat32,
	TagDouble:   func(r *Reader) float32 { return r.readFloat32() },
	TagUTF8Char: readUTF8CharAsFloat32,
	TagString:   readStringAsFloat32,
	TagDate:     readDateTimeAsFloat32,
	TagTime:     readTimeAsFloat32,
	TagRef:      readRefAsFloat32,
}

func float32Decoder(r *Reader, v reflect.Value, tag byte) {
	decoder := float32Decoders[tag]
	if decoder != nil {
		v.SetFloat(float64(decoder(r)))
		return
	}
	castError(tag, "float32")
}
