/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/value_encode.go                              |
|                                                          |
| LastModified: Feb 22, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"math"
	"math/big"
	"strconv"
	"time"

	"github.com/hprose/hprose-golang/v3/io"
)

const (
	digits = "0123456789"
	digit2 = "" +
		"0001020304050607080910111213141516171819" +
		"2021222324252627282930313233343536373839" +
		"4041424344454647484950515253545556575859" +
		"6061626364656667686970717273747576777879" +
		"8081828384858687888990919293949596979899"
	digit3 = "" +
		"000001002003004005006007008009010011012013014015016017018019" +
		"020021022023024025026027028029030031032033034035036037038039" +
		"040041042043044045046047048049050051052053054055056057058059" +
		"060061062063064065066067068069070071072073074075076077078079" +
		"080081082083084085086087088089090091092093094095096097098099" +
		"100101102103104105106107108109110111112113114115116117118119" +
		"120121122123124125126127128129130131132133134135136137138139" +
		"140141142143144145146147148149150151152153154155156157158159" +
		"160161162163164165166167168169170171172173174175176177178179" +
		"180181182183184185186187188189190191192193194195196197198199" +
		"200201202203204205206207208209210211212213214215216217218219" +
		"220221222223224225226227228229230231232233234235236237238239" +
		"240241242243244245246247248249250251252253254255256257258259" +
		"260261262263264265266267268269270271272273274275276277278279" +
		"280281282283284285286287288289290291292293294295296297298299" +
		"300301302303304305306307308309310311312313314315316317318319" +
		"320321322323324325326327328329330331332333334335336337338339" +
		"340341342343344345346347348349350351352353354355356357358359" +
		"360361362363364365366367368369370371372373374375376377378379" +
		"380381382383384385386387388389390391392393394395396397398399" +
		"400401402403404405406407408409410411412413414415416417418419" +
		"420421422423424425426427428429430431432433434435436437438439" +
		"440441442443444445446447448449450451452453454455456457458459" +
		"460461462463464465466467468469470471472473474475476477478479" +
		"480481482483484485486487488489490491492493494495496497498499" +
		"500501502503504505506507508509510511512513514515516517518519" +
		"520521522523524525526527528529530531532533534535536537538539" +
		"540541542543544545546547548549550551552553554555556557558559" +
		"560561562563564565566567568569570571572573574575576577578579" +
		"580581582583584585586587588589590591592593594595596597598599" +
		"600601602603604605606607608609610611612613614615616617618619" +
		"620621622623624625626627628629630631632633634635636637638639" +
		"640641642643644645646647648649650651652653654655656657658659" +
		"660661662663664665666667668669670671672673674675676677678679" +
		"680681682683684685686687688689690691692693694695696697698699" +
		"700701702703704705706707708709710711712713714715716717718719" +
		"720721722723724725726727728729730731732733734735736737738739" +
		"740741742743744745746747748749750751752753754755756757758759" +
		"760761762763764765766767768769770771772773774775776777778779" +
		"780781782783784785786787788789790791792793794795796797798799" +
		"800801802803804805806807808809810811812813814815816817818819" +
		"820821822823824825826827828829830831832833834835836837838839" +
		"840841842843844845846847848849850851852853854855856857858859" +
		"860861862863864865866867868869870871872873874875876877878879" +
		"880881882883884885886887888889890891892893894895896897898899" +
		"900901902903904905906907908909910911912913914915916917918919" +
		"920921922923924925926927928929930931932933934935936937938939" +
		"940941942943944945946947948949950951952953954955956957958959" +
		"960961962963964965966967968969970971972973974975976977978979" +
		"980981982983984985986987988989990991992993994995996997998999"
)

var minInt64Buf = []byte("-9223372036854775808")

func toBytes(i uint64, buf []byte) (off int) {
	off = len(buf)
	var q, p uint64
	for i >= 100 {
		q = i / 1000
		p = (i - (q * 1000)) * 3
		i = q
		off -= 3
		buf[off] = digit3[p]
		buf[off+1] = digit3[p+1]
		buf[off+2] = digit3[p+2]
	}
	if i >= 10 {
		q = i / 100
		p = (i - (q * 100)) * 2
		i = q
		off -= 2
		buf[off] = digit2[p]
		buf[off+1] = digit2[p+1]
	}
	if i > 0 {
		off--
		buf[off] = digits[i]
	}
	return
}

func writeInt64(writer io.Writer, i int64) (err error) {
	if i >= 0 {
		return writeUint64(writer, uint64(i))
	}
	if i == math.MinInt64 {
		_, err = writer.Write(minInt64Buf)
		return err
	}
	var u uint64 = uint64(-i)
	var buf [20]byte
	off := toBytes(u, buf[:]) - 1
	buf[off] = '-'
	_, err = writer.Write(buf[off:])
	return
}

func writeUint64(writer io.Writer, i uint64) (err error) {
	if (i >= 0) && (i <= 9) {
		return writer.WriteByte(digits[i])
	}
	var buf [20]byte
	off := toBytes(i, buf[:])
	_, err = writer.Write(buf[off:])
	return
}

// WriteInt64 to writer
func WriteInt64(writer io.Writer, i int64) (err error) {
	if (i >= 0) && (i <= 9) {
		return writer.WriteByte(digits[i])
	}
	var tag = io.TagInteger
	if (i < math.MinInt32) || (i > math.MaxInt32) {
		tag = io.TagLong
	}
	if err = writer.WriteByte(tag); err == nil {
		if err = writeInt64(writer, i); err == nil {
			err = writer.WriteByte(io.TagSemicolon)
		}
	}
	return
}

// WriteUint64 to writer
func WriteUint64(writer io.Writer, i uint64) (err error) {
	if (i >= 0) && (i <= 9) {
		return writer.WriteByte(digits[i])
	}
	var tag = io.TagInteger
	if i > math.MaxInt32 {
		tag = io.TagLong
	}
	if err = writer.WriteByte(tag); err == nil {
		if err = writeUint64(writer, i); err == nil {
			err = writer.WriteByte(io.TagSemicolon)
		}
	}
	return
}

// WriteInt32 to writer
func WriteInt32(writer io.Writer, i int32) (err error) {
	if (i >= 0) && (i <= 9) {
		return writer.WriteByte(digits[i])
	}
	if err = writer.WriteByte(io.TagInteger); err == nil {
		if err = writeInt64(writer, int64(i)); err == nil {
			err = writer.WriteByte(io.TagSemicolon)
		}
	}
	return
}

// WriteUint32 to writer
func WriteUint32(writer io.Writer, i uint32) (err error) {
	return WriteUint64(writer, uint64(i))
}

// WriteInt16 to writer
func WriteInt16(writer io.Writer, i int16) (err error) {
	return WriteInt32(writer, int32(i))
}

// WriteUint16 to writer
func WriteUint16(writer io.Writer, i uint16) (err error) {
	if (i >= 0) && (i <= 9) {
		return writer.WriteByte(digits[i])
	}
	if err = writer.WriteByte(io.TagInteger); err == nil {
		if err = writeUint64(writer, uint64(i)); err == nil {
			err = writer.WriteByte(io.TagSemicolon)
		}
	}
	return
}

// WriteInt8 to writer
func WriteInt8(writer io.Writer, i int8) (err error) {
	return WriteInt32(writer, int32(i))
}

// WriteUint8 to writer
func WriteUint8(writer io.Writer, i uint8) (err error) {
	return WriteUint16(writer, uint16(i))
}

// WriteInt to writer
func WriteInt(writer io.Writer, i int) (err error) {
	return WriteInt64(writer, int64(i))
}

// WriteUint to writer
func WriteUint(writer io.Writer, i uint) (err error) {
	return WriteUint64(writer, uint64(i))
}

// WriteNil to writer
func WriteNil(writer io.Writer) (err error) {
	return writer.WriteByte(io.TagNull)
}

// WriteBool to writer
func WriteBool(writer io.Writer, b bool) (err error) {
	if b {
		return writer.WriteByte(io.TagTrue)
	}
	return writer.WriteByte(io.TagFalse)
}

func writeFloat(writer io.Writer, f float64, bitSize int) (err error) {
	if f != f {
		return writer.WriteByte(io.TagNaN)
	}
	if f > math.MaxFloat64 {
		if err = writer.WriteByte(io.TagInfinity); err == nil {
			err = writer.WriteByte(io.TagPos)
		}
		return
	}
	if f < -math.MaxFloat64 {
		if err = writer.WriteByte(io.TagInfinity); err == nil {
			err = writer.WriteByte(io.TagNeg)
		}
		return
	}
	if err = writer.WriteByte(io.TagDouble); err == nil {
		var buf [24]byte
		if _, err = writer.Write(strconv.AppendFloat(buf[:0], f, 'g', -1, bitSize)); err == nil {
			err = writer.WriteByte(io.TagSemicolon)
		}
	}
	return
}

// WriteFloat32 to writer
func WriteFloat32(writer io.Writer, f float32) error {
	return writeFloat(writer, float64(f), 32)
}

// WriteFloat64 to writer
func WriteFloat64(writer io.Writer, f float64) error {
	return writeFloat(writer, float64(f), 64)
}

// WriteBigInt to writer
func WriteBigInt(writer io.Writer, i *big.Int) (err error) {
	if err = writer.WriteByte(io.TagLong); err == nil {
		if _, err = writer.Write(io.StringToBytes(i.String())); err == nil {
			err = writer.WriteByte(io.TagSemicolon)
		}
	}
	return
}

// WriteBigFloat to writer
func WriteBigFloat(writer io.Writer, f *big.Float) (err error) {
	if err = writer.WriteByte(io.TagDouble); err == nil {
		var buf [32]byte
		if _, err = writer.Write(f.Append(buf[:0], 'g', -1)); err == nil {
			err = writer.WriteByte(io.TagSemicolon)
		}
	}
	return
}

func utf16Length(str string) (n int) {
	length := len(str)
	n = length
	p := 0
	for p < length {
		a := str[p]
		switch a >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			p++
		case 12, 13:
			p += 2
			n--
		case 14:
			p += 3
			n -= 2
		case 15:
			if (a & 8) == 8 {
				return -1
			}
			p += 4
			n -= 2
		default:
			return -1
		}
	}
	return n
}

func writeBytes(writer io.Writer, bytes []byte, length int) (err error) {
	if length > 0 {
		if err = writeUint64(writer, uint64(length)); err != nil {
			return
		}
	}
	if err = writer.WriteByte(io.TagQuote); err == nil {
		if _, err = writer.Write(bytes); err == nil {
			err = writer.WriteByte(io.TagQuote)
		}
	}
	return
}

func writeString(writer io.Writer, s string) (err error) {
	return writeBytes(writer, io.StringToBytes(s), utf16Length(s))
}

// WriteBytes to writer
func WriteBytes(writer io.Writer, bytes []byte) (err error) {
	if err = writer.WriteByte(io.TagBytes); err == nil {
		err = writeBytes(writer, bytes, len(bytes))
	}
	return
}

// WriteString to writer
func WriteString(writer io.Writer, s string) (err error) {
	length := utf16Length(s)
	bytes := io.StringToBytes(s)
	if length < 0 {
		return WriteBytes(writer, bytes)
	}
	if err = writer.WriteByte(io.TagString); err == nil {
		return writeBytes(writer, bytes, length)
	}
	return
}

func writeDate(writer io.Writer, year int, month int, day int) (err error) {
	var buf [9]byte
	buf[0] = io.TagDate
	q := year / 100
	p := q << 1
	buf[1] = digit2[p]
	buf[2] = digit2[p+1]
	p = (year - q*100) << 1
	buf[3] = digit2[p]
	buf[4] = digit2[p+1]
	p = month << 1
	buf[5] = digit2[p]
	buf[6] = digit2[p+1]
	p = day << 1
	buf[7] = digit2[p]
	buf[8] = digit2[p+1]
	_, err = writer.Write(buf[:])
	return
}

func writeTime(writer io.Writer, hour int, min int, sec int, nsec int) (err error) {
	var buf [17]byte
	buf[0] = io.TagTime
	p := hour << 1
	buf[1] = digit2[p]
	buf[2] = digit2[p+1]
	p = min << 1
	buf[3] = digit2[p]
	buf[4] = digit2[p+1]
	p = sec << 1
	buf[5] = digit2[p]
	buf[6] = digit2[p+1]
	if nsec == 0 {
		_, err = writer.Write(buf[:7])
		return
	}
	buf[7] = io.TagPoint
	q := nsec / 1000000
	p = q * 3
	nsec = nsec - q*1000000
	buf[8] = digit3[p]
	buf[9] = digit3[p+1]
	buf[10] = digit3[p+2]
	if nsec == 0 {
		_, err = writer.Write(buf[:11])
		return
	}
	q = nsec / 1000
	p = q * 3
	nsec = nsec - q*1000
	buf[11] = digit3[p]
	buf[12] = digit3[p+1]
	buf[13] = digit3[p+2]
	if nsec == 0 {
		_, err = writer.Write(buf[:14])
		return
	}
	p = nsec * 3
	buf[14] = digit3[p]
	buf[15] = digit3[p+1]
	buf[16] = digit3[p+2]
	_, err = writer.Write(buf[:17])
	return
}

// WriteTime to writer
func WriteTime(writer io.Writer, t time.Time) (err error) {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	nsec := t.Nanosecond()
	if (hour == 0) && (min == 0) && (sec == 0) && (nsec == 0) {
		err = writeDate(writer, year, int(month), day)
	} else if (year == 1970) && (month == 1) && (day == 1) {
		err = writeTime(writer, hour, min, sec, nsec)
	} else if err = writeDate(writer, year, int(month), day); err == nil {
		err = writeTime(writer, hour, min, sec, nsec)
	}
	if err == nil {
		loc := io.TagSemicolon
		if t.Location() == time.UTC {
			loc = io.TagUTC
		}
		err = writer.WriteByte(loc)
	}
	return
}
