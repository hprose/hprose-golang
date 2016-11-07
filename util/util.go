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
 * util/util.go                                           *
 *                                                        *
 * some utility functions for Go.                         *
 *                                                        *
 * LastModified: Nov 7, 2016                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

/*
Package util defines some utility functions for Golang.
*/
package util

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

const digits = "0123456789"

const digit2 = "" +
	"0001020304050607080910111213141516171819" +
	"2021222324252627282930313233343536373839" +
	"4041424344454647484950515253545556575859" +
	"6061626364656667686970717273747576777879" +
	"8081828384858687888990919293949596979899"

const digit3 = "" +
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

var minInt64Buf = [...]byte{
	'-', '9', '2', '2', '3', '3', '7', '2', '0', '3',
	'6', '8', '5', '4', '7', '7', '5', '8', '0', '8'}

// GetIntBytes returns the []byte representation of i in base 10.
// buf length must be greater than or equal to 20
func GetIntBytes(buf []byte, i int64) []byte {
	if i == 0 {
		return []byte{'0'}
	}
	if i == math.MinInt64 {
		return minInt64Buf[:]
	}
	var sign byte
	if i < 0 {
		sign = '-'
		i = -i
	}
	off := len(buf)
	var q, p int64
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
	if sign == '-' {
		off--
		buf[off] = sign
	}
	return buf[off:]
}

// GetUintBytes returns the []byte representation of u in base 10.
// buf length must be greater than or equal to 20
func GetUintBytes(buf []byte, u uint64) []byte {
	if u == 0 {
		return []byte{'0'}
	}
	off := len(buf)
	var q, p uint64
	for u >= 100 {
		q = u / 1000
		p = (u - (q * 1000)) * 3
		u = q
		off -= 3
		buf[off] = digit3[p]
		buf[off+1] = digit3[p+1]
		buf[off+2] = digit3[p+2]
	}
	if u >= 10 {
		q = u / 100
		p = (u - (q * 100)) * 2
		u = q
		off -= 2
		buf[off] = digit2[p]
		buf[off+1] = digit2[p+1]
	}
	if u > 0 {
		off--
		buf[off] = digits[u]
	}
	return buf[off:]
}

// GetDateBytes returns the []byte representation of year, month and day.
// The format of []byte returned is 20060102
// buf length must be greater than or equal to 8
func GetDateBytes(buf []byte, year int, month int, day int) []byte {
	q := year / 100
	p := q << 1
	buf[0] = digit2[p]
	buf[1] = digit2[p+1]
	p = (year - q*100) << 1
	buf[2] = digit2[p]
	buf[3] = digit2[p+1]
	p = month << 1
	buf[4] = digit2[p]
	buf[5] = digit2[p+1]
	p = day << 1
	buf[6] = digit2[p]
	buf[7] = digit2[p+1]
	return buf[:8]
}

// GetTimeBytes returns the []byte representation of hour, min and sec.
// The format of []byte returned is 150405
// buf length must be greater than or equal to 6
func GetTimeBytes(buf []byte, hour int, min int, sec int) []byte {
	p := hour << 1
	buf[0] = digit2[p]
	buf[1] = digit2[p+1]
	p = min << 1
	buf[2] = digit2[p]
	buf[3] = digit2[p+1]
	p = sec << 1
	buf[4] = digit2[p]
	buf[5] = digit2[p+1]
	return buf[:6]
}

// GetNsecBytes returns the []byte representation of nsec.
// The format of []byte returned is 123, 123456 or 123456789
// buf length must be greater than or equal to 9
func GetNsecBytes(buf []byte, nsec int) []byte {
	q := nsec / 1000000
	p := q * 3
	nsec = nsec - q*1000000
	buf[0] = digit3[p]
	buf[1] = digit3[p+1]
	buf[2] = digit3[p+2]
	if nsec == 0 {
		return buf[:3]
	}
	q = nsec / 1000
	p = q * 3
	nsec = nsec - q*1000
	buf[3] = digit3[p]
	buf[4] = digit3[p+1]
	buf[5] = digit3[p+2]
	if nsec == 0 {
		return buf[:6]
	}
	p = nsec * 3
	buf[6] = digit3[p]
	buf[7] = digit3[p+1]
	buf[8] = digit3[p+2]
	return buf[:9]
}

// UTF16Length return the UTF16 length of str.
// str must be an UTF8 encode string, otherwise return -1.
func UTF16Length(str string) (n int) {
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
			if a&8 == 8 {
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

// ByteString converts []byte to string without memory allocation by block magic
func ByteString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringByte converts string to string without memory allocation by block magic
func StringByte(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// Itoa returns the string representation of i
func Itoa(i int) string {
	var buf [20]byte
	return ByteString(GetIntBytes(buf[:], int64(i)))
}

// Min returns the min one of a, b
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// Max returns the max one of a, b
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// UUIDv4 returns a version 4 UUID string
func UUIDv4() (uid string) {
	u := make([]byte, 16)
	rand.Read(u)
	u[6] = (u[6] & 0x0f) | 0x40
	u[8] = (u[8] & 0x3f) | 0x80
	uid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// ToUint32 translates 4 byte to uint32
func ToUint32(b []byte) uint32 {
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}

// FromUint32 translates uint32 to 4 byte
func FromUint32(b []byte, i uint32) {
	b[0] = byte(i >> 24)
	b[1] = byte(i >> 16)
	b[2] = byte(i >> 8)
	b[3] = byte(i)
}

// LocalProxy make a local object to a proxy struct
func LocalProxy(proxy, local interface{}) error {
	dstValue := reflect.ValueOf(proxy)
	srcValue := reflect.ValueOf(local)
	if dstValue.Kind() != reflect.Ptr {
		return errors.New("proxy must be a pointer")
	}
	dstValue = dstValue.Elem()
	t := dstValue.Type()
	et := t
	if et.Kind() == reflect.Ptr {
		et = et.Elem()
	}
	if et.Kind() != reflect.Struct {
		return errors.New("proxy must be a struct pointer or pointer to a struct pointer")
	}
	ptr := reflect.New(et)
	obj := ptr.Elem()
	count := obj.NumField()
	for i := 0; i < count; i++ {
		f := obj.Field(i)
		ft := f.Type()
		sf := et.Field(i)
		if f.CanSet() && ft.Kind() == reflect.Func {
			f.Set(srcValue.MethodByName(sf.Name))
		}
	}
	if t.Kind() == reflect.Ptr {
		dstValue.Set(ptr)
	} else {
		dstValue.Set(obj)
	}
	return nil
}
