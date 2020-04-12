/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/time_encoder.go                                 |
|                                                          |
| LastModified: Apr 12, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"time"

	"github.com/modern-go/reflect2"
)

// timeEncoder is the implementation of ValueEncoder for time.Time/*time.Time.
type timeEncoder struct{}

func (valenc timeEncoder) Encode(enc *Encoder, v interface{}) {
	enc.EncodeReference(valenc, v)
}

func (timeEncoder) Write(enc *Encoder, v interface{}) {
	enc.SetReference(v)
	writeTime(enc, *(*time.Time)(reflect2.PtrOf(v)))
}

func writeDatePart(enc *Encoder, year int, month int, day int) {
	enc.buf = append(enc.buf, TagDate)
	q := year / 100
	p := q << 1
	enc.buf = append(enc.buf, digit2[p:p+2]...)
	p = (year - q*100) << 1
	enc.buf = append(enc.buf, digit2[p:p+2]...)
	p = month << 1
	enc.buf = append(enc.buf, digit2[p:p+2]...)
	p = day << 1
	enc.buf = append(enc.buf, digit2[p:p+2]...)
}

func writeTimePart(enc *Encoder, hour int, min int, sec int, nsec int) {
	enc.buf = append(enc.buf, TagTime)
	p := hour << 1
	enc.buf = append(enc.buf, digit2[p:p+2]...)
	p = min << 1
	enc.buf = append(enc.buf, digit2[p:p+2]...)
	p = sec << 1
	enc.buf = append(enc.buf, digit2[p:p+2]...)
	if nsec == 0 {
		return
	}
	enc.buf = append(enc.buf, TagPoint)
	q := nsec / 1000000
	p = q * 3
	nsec = nsec - q*1000000
	enc.buf = append(enc.buf, digit3[p:p+3]...)
	if nsec == 0 {
		return
	}
	q = nsec / 1000
	p = q * 3
	nsec = nsec - q*1000
	enc.buf = append(enc.buf, digit3[p:p+3]...)
	if nsec == 0 {
		return
	}
	p = nsec * 3
	enc.buf = append(enc.buf, digit3[p:p+3]...)
	return
}

func writeTime(enc *Encoder, t time.Time) {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	nsec := t.Nanosecond()
	if (hour == 0) && (min == 0) && (sec == 0) && (nsec == 0) {
		writeDatePart(enc, year, int(month), day)
	} else if (year == 1970) && (month == 1) && (day == 1) {
		writeTimePart(enc, hour, min, sec, nsec)
	} else {
		writeDatePart(enc, year, int(month), day)
		writeTimePart(enc, hour, min, sec, nsec)
	}
	loc := TagSemicolon
	if t.Location() == time.UTC {
		loc = TagUTC
	}
	enc.buf = append(enc.buf, loc)
}

func init() {
	RegisterValueEncoder((*time.Time)(nil), timeEncoder{})
}
