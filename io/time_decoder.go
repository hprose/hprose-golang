/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/time_decoder.go                                       |
|                                                          |
| LastModified: May 14, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"reflect"
	"time"

	"github.com/modern-go/reflect2"
)

var timeFormat = []string{
	"2006-01-02 15:04:05",
	"2006-01-02 15:04:05.999999999",
	"2006-01-02 15:04:05Z07:00",
	"2006-01-02 15:04:05.999999999Z07:00",
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
	"2006-01-02",
	"02 Jan 06",
	"02-Jan-06",
	"02 Jan 2006",
	"15:04:05",
	"15:04:05.999999999",
	"15:04:05Z07:00",
	"15:04:05.999999999Z07:00",
}

func (dec *Decoder) stringToTime(value string) time.Time {
	for _, layout := range timeFormat {
		if t, e := time.Parse(layout, value); e == nil {
			return t
		}
	}
	dec.decodeStringError(value, "time.Time")
	return time.Unix(0, 0)
}

func (dec *Decoder) read2Digit() (n int) {
	i := intDigits[dec.NextByte()]
	i2 := intDigits[dec.NextByte()]
	return int(i*10 + i2)
}

func (dec *Decoder) read3Digit() (n int) {
	i := intDigits[dec.NextByte()]
	i2 := intDigits[dec.NextByte()]
	i3 := intDigits[dec.NextByte()]
	return int(i*100 + i2*10 + i3)
}

func (dec *Decoder) read4Digit() (n int) {
	i := intDigits[dec.NextByte()]
	i2 := intDigits[dec.NextByte()]
	i3 := intDigits[dec.NextByte()]
	i4 := intDigits[dec.NextByte()]
	return int(i*1000 + i2*100 + i3*10 + i4)
}

func (dec *Decoder) readNsec() (nsec int, tag byte) {
	nsec = dec.read3Digit()
	nsec *= 1000000
	tag = dec.NextByte()
	i := intDigits[tag]
	if i == invalidDigit {
		return
	}
	nsec += int(i * 100000)
	nsec += dec.read2Digit() * 1000
	tag = dec.NextByte()
	i = intDigits[tag]
	if i == invalidDigit {
		return
	}
	nsec += int(i * 100)
	nsec += dec.read2Digit()
	tag = dec.NextByte()
	return
}

// ReadTime reads time.Time and add reference.
func (dec *Decoder) ReadTime() (t time.Time) {
	hour := dec.read2Digit()
	min := dec.read2Digit()
	sec := dec.read2Digit()
	tag := dec.NextByte()
	var nsec int
	if tag == TagPoint {
		nsec, tag = dec.readNsec()
	}
	loc := time.Local
	if tag == TagUTC {
		loc = time.UTC
	}
	t = time.Date(1970, time.January, 1, hour, min, sec, nsec, loc)
	dec.AddReference(t)
	return
}

// ReadDateTime reads time.Time and add reference.
func (dec *Decoder) ReadDateTime() (t time.Time) {
	year := dec.read4Digit()
	month := dec.read2Digit()
	day := dec.read2Digit()
	tag := dec.NextByte()
	var hour, min, sec, nsec int
	if tag == TagTime {
		hour = dec.read2Digit()
		min = dec.read2Digit()
		sec = dec.read2Digit()
		tag = dec.NextByte()
		if tag == TagPoint {
			nsec, tag = dec.readNsec()
		}
	}
	loc := time.Local
	if tag == TagUTC {
		loc = time.UTC
	}
	t = time.Date(year, time.Month(month), day, hour, min, sec, nsec, loc)
	dec.AddReference(t)
	return
}

func (dec *Decoder) decodeTime(t reflect.Type, tag byte) (result time.Time) {
	if i := intDigits[tag]; i != invalidDigit {
		return time.Unix(0, int64(i))
	}
	switch tag {
	case TagEmpty, TagFalse:
		return time.Unix(0, 0)
	case TagTrue:
		return time.Unix(0, 1)
	case TagInteger, TagLong:
		return time.Unix(0, dec.ReadInt64())
	case TagDouble:
		return time.Unix(0, int64(dec.ReadFloat64()))
	case TagTime:
		return dec.ReadTime()
	case TagDate:
		return dec.ReadDateTime()
	case TagString:
		if dec.IsSimple() {
			return dec.stringToTime(dec.ReadUnsafeString())
		}
		return dec.stringToTime(dec.ReadString())
	default:
		dec.defaultDecode(t, &result, tag)
	}
	return
}

func (dec *Decoder) decodeTimePtr(t reflect.Type, tag byte) *time.Time {
	if tag == TagNull {
		return nil
	}
	time := dec.decodeTime(t, tag)
	return &time
}

// timeDecoder is the implementation of ValueDecoder for time.Time.
type timeDecoder struct{}

func (timeDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*time.Time)(reflect2.PtrOf(p)) = dec.decodeTime(timeType, tag)
}

func (timeDecoder) Type() reflect.Type {
	return timeType
}

func init() {
	RegisterValueDecoder(timeDecoder{})
}
