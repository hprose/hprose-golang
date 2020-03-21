/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/time_encoder.go                                 |
|                                                          |
| LastModified: Mar 21, 2020                               |
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

func (valenc timeEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	return EncodeReference(valenc, enc, v)
}

func (timeEncoder) Write(enc *Encoder, v interface{}) (err error) {
	SetReference(enc, v)
	return writeTime(enc.writer, *(*time.Time)(reflect2.PtrOf(v)))
}

func writeDatePart(writer bytesWriter, year int, month int, day int) (err error) {
	var buf [9]byte
	buf[0] = TagDate
	q := year / 100
	p := q << 1
	copy(buf[1:3], digit2[p:p+2])
	p = (year - q*100) << 1
	copy(buf[3:5], digit2[p:p+2])
	p = month << 1
	copy(buf[5:7], digit2[p:p+2])
	p = day << 1
	copy(buf[7:9], digit2[p:p+2])
	_, err = writer.Write(buf[:])
	return
}

func writeTimePart(writer bytesWriter, hour int, min int, sec int, nsec int) (err error) {
	var buf [17]byte
	buf[0] = TagTime
	p := hour << 1
	copy(buf[1:3], digit2[p:p+2])
	p = min << 1
	copy(buf[3:5], digit2[p:p+2])
	p = sec << 1
	copy(buf[5:7], digit2[p:p+2])
	if nsec == 0 {
		_, err = writer.Write(buf[:7])
		return
	}
	buf[7] = TagPoint
	q := nsec / 1000000
	p = q * 3
	nsec = nsec - q*1000000
	copy(buf[8:11], digit3[p:p+3])
	if nsec == 0 {
		_, err = writer.Write(buf[:11])
		return
	}
	q = nsec / 1000
	p = q * 3
	nsec = nsec - q*1000
	copy(buf[11:14], digit3[p:p+3])
	if nsec == 0 {
		_, err = writer.Write(buf[:14])
		return
	}
	p = nsec * 3
	copy(buf[14:17], digit3[p:p+3])
	_, err = writer.Write(buf[:17])
	return
}

func writeTime(writer bytesWriter, t time.Time) (err error) {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	nsec := t.Nanosecond()
	if (hour == 0) && (min == 0) && (sec == 0) && (nsec == 0) {
		err = writeDatePart(writer, year, int(month), day)
	} else if (year == 1970) && (month == 1) && (day == 1) {
		err = writeTimePart(writer, hour, min, sec, nsec)
	} else if err = writeDatePart(writer, year, int(month), day); err == nil {
		err = writeTimePart(writer, hour, min, sec, nsec)
	}
	if err == nil {
		loc := TagSemicolon
		if t.Location() == time.UTC {
			loc = TagUTC
		}
		err = writer.WriteByte(loc)
	}
	return
}

func init() {
	RegisterValueEncoder((*time.Time)(nil), timeEncoder{})
}
