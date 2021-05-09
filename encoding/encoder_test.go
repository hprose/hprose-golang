/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/encoder_test.go                                 |
|                                                          |
| LastModified: May 9, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding_test

import (
	"errors"
	"math"
	"math/big"
	"reflect"
	"strings"
	"testing"
	"time"

	. "github.com/hprose/hprose-golang/v3/encoding"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestWriteInt(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.WriteInt(-1)
	enc.WriteInt(0)
	enc.WriteInt(1)
	enc.WriteInt(123)
	enc.WriteInt(math.MinInt64)
	enc.WriteInt(-math.MaxInt64)
	enc.WriteInt(math.MaxInt64)
	assert.NoError(t, enc.Flush())
	assert.Equal(t, "i-1;01i123;l-9223372036854775808;l-9223372036854775807;l9223372036854775807;", sb.String())
}

func TestWriteUint(t *testing.T) {
	enc := new(Encoder)
	enc.WriteUint(0)
	enc.WriteUint(1)
	enc.WriteUint(123)
	enc.WriteUint(math.MaxUint64)
	enc.WriteUint(math.MaxUint32)
	enc.WriteUint(math.MaxInt32)
	assert.Equal(t, "01i123;l18446744073709551615;l4294967295;i2147483647;", enc.String())
}

func TestWriteInt32(t *testing.T) {
	enc := new(Encoder)
	enc.WriteInt32(-1)
	enc.WriteInt32(0)
	enc.WriteInt32(1)
	enc.WriteInt32(123)
	enc.WriteInt32(math.MinInt32)
	enc.WriteInt32(math.MaxInt32)
	assert.Equal(t, "i-1;01i123;i-2147483648;i2147483647;", enc.String())
}

func TestWriteUint32(t *testing.T) {
	enc := new(Encoder)
	enc.WriteUint32(0)
	enc.WriteUint32(1)
	enc.WriteUint32(123)
	enc.WriteUint32(math.MaxUint32)
	enc.WriteUint32(math.MaxInt32)
	assert.Equal(t, "01i123;l4294967295;i2147483647;", enc.String())
}

func TestWriteUint16(t *testing.T) {
	enc := new(Encoder)
	enc.WriteUint16(0)
	enc.WriteUint16(1)
	enc.WriteUint16(123)
	enc.WriteUint16(math.MaxUint16)
	enc.WriteUint16(math.MaxInt16)
	assert.Equal(t, "01i123;i65535;i32767;", enc.String())
}

func TestWriteInt16(t *testing.T) {
	enc := new(Encoder)
	enc.WriteInt16(0)
	enc.WriteInt16(1)
	enc.WriteInt16(123)
	enc.WriteInt16(math.MinInt16)
	enc.WriteInt16(math.MaxInt16)
	assert.Equal(t, "01i123;i-32768;i32767;", enc.String())
}

func TestWriteUint8(t *testing.T) {
	enc := new(Encoder)
	enc.WriteUint8(0)
	enc.WriteUint8(1)
	enc.WriteUint8(123)
	enc.WriteUint8(math.MaxUint8)
	enc.WriteUint8(math.MaxInt8)
	assert.Equal(t, "01i123;i255;i127;", enc.String())
}

func TestWriteInt8(t *testing.T) {
	enc := new(Encoder)
	enc.WriteInt8(0)
	enc.WriteInt8(1)
	enc.WriteInt8(123)
	enc.WriteInt8(math.MinInt8)
	enc.WriteInt8(math.MaxInt8)
	assert.Equal(t, "01i123;i-128;i127;", enc.String())
}

func TestWriteBool(t *testing.T) {
	enc := new(Encoder)
	enc.WriteBool(true)
	enc.WriteBool(false)
	assert.Equal(t, "tf", enc.String())
}

func TestWriteFloat(t *testing.T) {
	enc := new(Encoder)
	enc.WriteFloat32(math.E)
	enc.WriteFloat32(math.Pi)
	enc.WriteFloat64(math.E)
	enc.WriteFloat64(math.Pi)
	enc.WriteFloat64(math.Log(1))
	enc.WriteFloat64(math.Log(0))
	enc.WriteFloat64(-math.Log(0))
	enc.WriteFloat64(math.Log(-1))
	assert.Equal(t, "d2.7182817;d3.1415927;d2.718281828459045;d3.141592653589793;d0;I-I+N", enc.String())
}

func TestWriteBigInt(t *testing.T) {
	enc := new(Encoder)
	enc.WriteBigInt(big.NewInt(math.MaxInt64))
	assert.Equal(t, "l9223372036854775807;", enc.String())
}

func TestWriteBigFloat(t *testing.T) {
	enc := new(Encoder)
	enc.WriteBigFloat(big.NewFloat(math.MaxFloat64))
	assert.Equal(t, "d1.7976931348623157e+308;", enc.String())
}

func TestEncoderEncode(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb).Simple(false)
	i := 0
	i8 := int8(1)
	i16 := int16(2)
	i32 := int32(3)
	i64 := int64(4)
	u := uint(5)
	u8 := uint8(6)
	u16 := uint16(7)
	u32 := uint32(8)
	u64 := uint64(9)
	uptr := uintptr(10)
	b := true
	f32 := float32(math.Pi)
	f64 := float64(math.Pi)
	e := ""
	c := "我"
	s := "Hello"
	c64 := complex(float32(1), float32(2))
	c128 := complex(float64(3), float64(4))
	r64 := complex(float32(5), float32(0))
	r128 := complex(float64(6), float64(0))
	bi := big.NewInt(0)
	bf := big.NewFloat(1)
	br := big.NewRat(2, 3)
	bri := big.NewRat(4, 1)
	assert.NoError(t, enc.Encode(nil))
	assert.NoError(t, enc.Encode(i))
	assert.NoError(t, enc.Encode(i8))
	assert.NoError(t, enc.Encode(i16))
	assert.NoError(t, enc.Encode(i32))
	assert.NoError(t, enc.Encode(i64))
	assert.NoError(t, enc.Encode(u))
	assert.NoError(t, enc.Encode(u8))
	assert.NoError(t, enc.Encode(u16))
	assert.NoError(t, enc.Encode(u32))
	assert.NoError(t, enc.Encode(u64))
	assert.NoError(t, enc.Encode(uptr))
	assert.NoError(t, enc.Encode(b))
	assert.NoError(t, enc.Encode(f32))
	assert.NoError(t, enc.Encode(f64))
	assert.NoError(t, enc.Encode(e))
	assert.NoError(t, enc.Encode(c))
	assert.NoError(t, enc.Encode(s))
	assert.NoError(t, enc.Encode(c64))
	assert.NoError(t, enc.Encode(c128))
	assert.NoError(t, enc.Encode(r64))
	assert.NoError(t, enc.Encode(r128))
	assert.NoError(t, enc.Encode(*bi))
	assert.NoError(t, enc.Encode(*bf))
	assert.NoError(t, enc.Encode(*br))
	assert.NoError(t, enc.Encode(*bri))
	assert.NoError(t, enc.Encode((*int)(nil)))
	assert.NoError(t, enc.Encode(&i))
	assert.NoError(t, enc.Encode(&i8))
	assert.NoError(t, enc.Encode(&i16))
	assert.NoError(t, enc.Encode(&i32))
	assert.NoError(t, enc.Encode(&i64))
	assert.NoError(t, enc.Encode(&u))
	assert.NoError(t, enc.Encode(&u8))
	assert.NoError(t, enc.Encode(&u16))
	assert.NoError(t, enc.Encode(&u32))
	assert.NoError(t, enc.Encode(&u64))
	assert.NoError(t, enc.Encode(&uptr))
	assert.NoError(t, enc.Encode(&b))
	assert.NoError(t, enc.Encode(&f32))
	assert.NoError(t, enc.Encode(&f64))
	assert.NoError(t, enc.Encode(&e))
	assert.NoError(t, enc.Encode(&c))
	assert.NoError(t, enc.Encode(&s))
	assert.NoError(t, enc.Encode(&c64))
	assert.NoError(t, enc.Encode(&c128))
	assert.NoError(t, enc.Encode(&r64))
	assert.NoError(t, enc.Encode(&r128))
	assert.NoError(t, enc.Encode(bi))
	assert.NoError(t, enc.Encode(bf))
	assert.NoError(t, enc.Encode(br))
	assert.NoError(t, enc.Encode(bri))
	assert.Equal(t, `n0123l4;5678l9;l10;td3.1415927;d3.141592653589793;`+
		`eu我s5"Hello"a2{d1;d2;}a2{d3;d4;}d5;d6;l0;d1;s3"2/3"l4;`+
		`n0123l4;5678l9;l10;td3.1415927;d3.141592653589793;`+
		`eu我r0;a2{d1;d2;}a2{d3;d4;}d5;d6;l0;d1;s3"2/3"l4;`, sb.String())
}

func TestEncoderWrite(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb).Simple(false)
	i := 0
	i8 := int8(1)
	i16 := int16(2)
	i32 := int32(3)
	i64 := int64(4)
	u := uint(5)
	u8 := uint8(6)
	u16 := uint16(7)
	u32 := uint32(8)
	u64 := uint64(9)
	b := true
	f32 := float32(math.Pi)
	f64 := float64(math.Pi)
	e := ""
	c := "我"
	s := "Hello"
	c64 := complex(float32(1), float32(2))
	c128 := complex(float64(3), float64(4))
	r64 := complex(float32(5), float32(0))
	r128 := complex(float64(6), float64(0))
	bi := big.NewInt(0)
	bf := big.NewFloat(1)
	br := big.NewRat(2, 3)
	bri := big.NewRat(4, 1)
	assert.NoError(t, enc.Write(nil))
	assert.NoError(t, enc.Write(i))
	assert.NoError(t, enc.Write(i8))
	assert.NoError(t, enc.Write(i16))
	assert.NoError(t, enc.Write(i32))
	assert.NoError(t, enc.Write(i64))
	assert.NoError(t, enc.Write(u))
	assert.NoError(t, enc.Write(u8))
	assert.NoError(t, enc.Write(u16))
	assert.NoError(t, enc.Write(u32))
	assert.NoError(t, enc.Write(u64))
	assert.NoError(t, enc.Write(b))
	assert.NoError(t, enc.Write(f32))
	assert.NoError(t, enc.Write(f64))
	assert.NoError(t, enc.Write(e))
	assert.NoError(t, enc.Write(c))
	assert.NoError(t, enc.Write(s))
	assert.NoError(t, enc.Write(c64))
	assert.NoError(t, enc.Write(c128))
	assert.NoError(t, enc.Write(r64))
	assert.NoError(t, enc.Write(r128))
	assert.NoError(t, enc.Write(*bi))
	assert.NoError(t, enc.Write(*bf))
	assert.NoError(t, enc.Write(*br))
	assert.NoError(t, enc.Write(*bri))
	assert.NoError(t, enc.Write((*int)(nil)))
	assert.NoError(t, enc.Write(&i))
	assert.NoError(t, enc.Write(&i8))
	assert.NoError(t, enc.Write(&i16))
	assert.NoError(t, enc.Write(&i32))
	assert.NoError(t, enc.Write(&i64))
	assert.NoError(t, enc.Write(&u))
	assert.NoError(t, enc.Write(&u8))
	assert.NoError(t, enc.Write(&u16))
	assert.NoError(t, enc.Write(&u32))
	assert.NoError(t, enc.Write(&u64))
	assert.NoError(t, enc.Write(&b))
	assert.NoError(t, enc.Write(&f32))
	assert.NoError(t, enc.Write(&f64))
	assert.NoError(t, enc.Write(&e))
	assert.NoError(t, enc.Write(&c))
	assert.NoError(t, enc.Write(&s))
	assert.NoError(t, enc.Write(&c64))
	assert.NoError(t, enc.Write(&c128))
	assert.NoError(t, enc.Write(&r64))
	assert.NoError(t, enc.Write(&r128))
	assert.NoError(t, enc.Write(bi))
	assert.NoError(t, enc.Write(bf))
	assert.NoError(t, enc.Write(br))
	assert.NoError(t, enc.Write(bri))
	assert.Equal(t, `n0123l4;5678l9;td3.1415927;d3.141592653589793;`+
		`s""s1"我"s5"Hello"a2{d1;d2;}a2{d3;d4;}d5;d6;l0;d1;s3"2/3"l4;`+
		`n0123l4;5678l9;td3.1415927;d3.141592653589793;`+
		`s""s1"我"s5"Hello"a2{d1;d2;}a2{d3;d4;}d5;d6;l0;d1;s3"2/3"l4;`, sb.String())
}

func TestEncodeError(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb)
	e := errors.New("test error")
	assert.NoError(t, enc.Encode(e))
	assert.NoError(t, enc.Encode(&e))
	assert.NoError(t, enc.Encode(&e))
	assert.Equal(t, `Es10"test error"Es10"test error"Es10"test error"`, sb.String())
}

func TestEncodeString(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb)
	assert.NoError(t, enc.Encode(""))
	assert.NoError(t, enc.Encode("Hello"))
	assert.NoError(t, enc.Encode("Pokémon"))
	assert.NoError(t, enc.Encode("中文"))
	assert.NoError(t, enc.Encode("🐱🐶"))
	assert.NoError(t, enc.Encode("👩‍👩‍👧‍👧"))
	assert.Equal(t, `es5"Hello"s7"Pokémon"s2"中文"s4"🐱🐶"s11"👩‍👩‍👧‍👧"`, sb.String())
}

func TestWriteBadString(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb)
	assert.NoError(t, enc.Write("\xfe\xfe"))
	assert.NoError(t, enc.Write("\xf0\xfe"))
	assert.NoError(t, enc.Write("\xf0\x80"))
	assert.Equal(t, "b2\"\xfe\xfe\"b2\"\xf0\xfe\"b2\"\xf0\x80\"", sb.String())
}

func TestReset(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	slice := []interface{}{1, "hello", true}
	var nilslice []interface{}
	assert.NoError(t, enc.Encode(nilslice))
	assert.NoError(t, enc.Encode([]interface{}{}))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(slice))
	enc.Reset()
	assert.NoError(t, enc.Encode(&slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.Equal(t, `na{}a3{1s5"hello"t}a3{1r2;t}a3{1s5"hello"t}r0;`, sb.String())
}

func TestEncodeByteArray(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	hello := [5]byte{'H', 'e', 'l', 'l', 'o'}
	assert.NoError(t, enc.Encode([0]byte{}))
	assert.NoError(t, enc.Encode(hello))
	assert.NoError(t, enc.Encode(hello))
	assert.NoError(t, enc.Encode(&hello))
	assert.NoError(t, enc.Encode(&hello))
	assert.Equal(t, `b""b5"Hello"b5"Hello"b5"Hello"r3;`, sb.String())
}

func TestEncodeByteArray2(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb)
	hello := [5]byte{'H', 'e', 'l', 'l', 'o'}
	assert.NoError(t, enc.Encode([0]byte{}))
	assert.NoError(t, enc.Encode(hello))
	assert.NoError(t, enc.Encode(hello))
	assert.NoError(t, enc.Encode(&hello))
	assert.NoError(t, enc.Encode(&hello))
	assert.Equal(t, `b""b5"Hello"b5"Hello"b5"Hello"b5"Hello"`, sb.String())
}

func TestEncodeBigIntArray(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	array := [2]*big.Int{big.NewInt(1), big.NewInt(2)}
	var emptyArray [0]*big.Int
	assert.NoError(t, enc.Encode(emptyArray))
	assert.NoError(t, enc.Encode([2]*big.Int{}))
	assert.NoError(t, enc.Encode(array))
	assert.NoError(t, enc.Encode(array))
	assert.NoError(t, enc.Encode(&array))
	assert.NoError(t, enc.Encode(&array))
	assert.Equal(t, `a{}a2{nn}a2{l1;l2;}a2{l1;l2;}a2{l1;l2;}r4;`, sb.String())
}

func TestEncodeNil(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	var x interface{}
	var xp interface{} = &x
	var i interface{} = (****int)(nil)
	var ip interface{} = &i
	assert.NoError(t, enc.Encode(nil))
	assert.NoError(t, enc.Encode(x))
	assert.NoError(t, enc.Encode(&x))
	assert.NoError(t, enc.Encode(&xp))
	assert.NoError(t, enc.Encode(i))
	assert.NoError(t, enc.Encode(&i))
	assert.NoError(t, enc.Encode(&ip))
	assert.Equal(t, `nnnnnnn`, sb.String())
}

func TestWriteNil(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	var x interface{}
	var xp interface{} = &x
	var i interface{} = (****int)(nil)
	var ip interface{} = &i
	assert.NoError(t, enc.Write(nil))
	assert.NoError(t, enc.Write(x))
	assert.NoError(t, enc.Write(&x))
	assert.NoError(t, enc.Write(&xp))
	assert.NoError(t, enc.Write(i))
	assert.NoError(t, enc.Write(&i))
	assert.NoError(t, enc.Write(&ip))
	assert.Equal(t, `nnnnnnn`, sb.String())
}

func TestWriteDuration(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	d := time.Duration(1000)
	assert.NoError(t, enc.Write(d))
	assert.NoError(t, enc.Write(&d))
	assert.Equal(t, `l1000;l1000;`, sb.String())
}

func TestWriteTime(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb)
	assert.NoError(t, enc.Write(time.Date(2020, 2, 22, 0, 0, 0, 0, time.UTC)))
	assert.NoError(t, enc.Write(time.Date(1970, 1, 1, 12, 12, 12, 0, time.UTC)))
	assert.NoError(t, enc.Write(time.Date(1970, 1, 1, 12, 12, 12, 123456789, time.Local)))
	assert.NoError(t, enc.Write(time.Date(2020, 2, 22, 12, 12, 12, 123456000, time.Local)))
	assert.NoError(t, enc.Write(time.Date(2020, 2, 22, 12, 12, 12, 123000000, time.UTC)))
	assert.Equal(t, "D20200222ZT121212ZT121212.123456789;D20200222T121212.123456;D20200222T121212.123Z", sb.String())
}

func TestEncodeStringStringMap(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	m := map[string]string{
		"hello": "world",
	}
	var nilmap map[string]string
	assert.NoError(t, enc.Encode(nilmap))
	assert.NoError(t, enc.Encode(map[string]string{}))
	assert.NoError(t, enc.Encode(m))
	assert.NoError(t, enc.Encode(&m))
	assert.NoError(t, enc.Encode(m))
	assert.NoError(t, enc.Encode(&m))
	assert.Equal(t, `nm{}m1{s5"hello"s5"world"}m1{r2;r3;}m1{r2;r3;}r4;`, sb.String())
}

func TestEncodeIntBigIntMap(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	m := map[int]*big.Int{
		1: big.NewInt(1),
	}
	var nilmap map[int]*big.Int
	assert.NoError(t, enc.Encode(nilmap))
	assert.NoError(t, enc.Encode(map[int]*big.Int{}))
	assert.NoError(t, enc.Encode(m))
	assert.NoError(t, enc.Encode(&m))
	assert.NoError(t, enc.Encode(m))
	assert.NoError(t, enc.Encode(&m))
	assert.Equal(t, `nm{}m1{1l1;}m1{1l1;}m1{1l1;}r2;`, sb.String())
}

func TestEncodeStringInterfaceMap(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	m := map[string]interface{}{
		"hello": "world",
	}
	var nilmap map[string]interface{}
	assert.NoError(t, enc.Encode(nilmap))
	assert.NoError(t, enc.Encode(map[string]interface{}{}))
	assert.NoError(t, enc.Encode(m))
	assert.NoError(t, enc.Encode(&m))
	assert.NoError(t, enc.Encode(m))
	assert.NoError(t, enc.Encode(&m))
	assert.Equal(t, `nm{}m1{s5"hello"s5"world"}m1{r2;r3;}m1{r2;r3;}r4;`, sb.String())
}

func TestEncodeCustomType(t *testing.T) {
	type IntType int
	type Int8Type int8
	type Int16Type int16
	type Int32Type int32
	type Int64Type int64
	type UintType uint
	type Uint8Type uint8
	type Uint16Type uint16
	type Uint32Type uint32
	type Uint64Type uint64
	type UintptrType uintptr
	type BoolType bool
	type Float32Type float32
	type Float64Type float64
	type Complex64Type complex64
	type Complex128Type complex128
	type StringType string
	type BigIntType big.Int

	RegisterValueEncoder((*BigIntType)(nil), GetValueEncoder((*big.Int)(nil)))

	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	i := IntType(0)
	i8 := Int8Type(1)
	i16 := Int16Type(2)
	i32 := Int32Type(3)
	i64 := Int64Type(4)
	u := UintType(5)
	u8 := Uint8Type(6)
	u16 := Uint16Type(7)
	u32 := Uint32Type(8)
	u64 := Uint64Type(9)
	uptr := UintptrType(10)
	b := BoolType(true)
	f32 := Float32Type(3.14159)
	f64 := Float64Type(2.17828)
	c64 := Complex64Type(complex(1, 2))
	c128 := Complex128Type(complex(3, 4))
	s := StringType("hello")
	bi := (*BigIntType)(big.NewInt(100))

	assert.NoError(t, enc.Encode(i))
	assert.NoError(t, enc.Encode(i8))
	assert.NoError(t, enc.Encode(i16))
	assert.NoError(t, enc.Encode(i32))
	assert.NoError(t, enc.Encode(i64))
	assert.NoError(t, enc.Encode(u))
	assert.NoError(t, enc.Encode(u8))
	assert.NoError(t, enc.Encode(u16))
	assert.NoError(t, enc.Encode(u32))
	assert.NoError(t, enc.Encode(u64))
	assert.NoError(t, enc.Encode(uptr))
	assert.NoError(t, enc.Encode(b))
	assert.NoError(t, enc.Encode(f32))
	assert.NoError(t, enc.Encode(f64))
	assert.NoError(t, enc.Encode(c64))
	assert.NoError(t, enc.Encode(c128))
	assert.NoError(t, enc.Encode(s))
	assert.NoError(t, enc.Encode(bi))
	assert.NoError(t, enc.Encode(&i))
	assert.NoError(t, enc.Encode(&i8))
	assert.NoError(t, enc.Encode(&i16))
	assert.NoError(t, enc.Encode(&i32))
	assert.NoError(t, enc.Encode(&i64))
	assert.NoError(t, enc.Encode(&u))
	assert.NoError(t, enc.Encode(&u8))
	assert.NoError(t, enc.Encode(&u16))
	assert.NoError(t, enc.Encode(&u32))
	assert.NoError(t, enc.Encode(&u64))
	assert.NoError(t, enc.Encode(&uptr))
	assert.NoError(t, enc.Encode(&b))
	assert.NoError(t, enc.Encode(&f32))
	assert.NoError(t, enc.Encode(&f64))
	assert.NoError(t, enc.Encode(&c64))
	assert.NoError(t, enc.Encode(&c128))
	assert.NoError(t, enc.Encode(&s))
	assert.NoError(t, enc.Encode(&bi))
	assert.Equal(t, `0123l4;5678l9;l10;td3.14159;d2.17828;a2{d1;d2;}a2{d3;d4;}`+
		`s5"hello"l100;0123l4;5678l9;l10;td3.14159;d2.17828;a2{d1;d2;}a2{d3;d4;}r2;l100;`, sb.String())
}

func TestUnsupportedTypeError(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	f := func() {}
	var ch chan int
	assert.EqualError(t, enc.Encode(f), (UnsupportedTypeError{reflect.TypeOf(f)}).Error())
	assert.EqualError(t, enc.Encode(ch), (UnsupportedTypeError{reflect.TypeOf(ch)}).Error())
	assert.EqualError(t, enc.Encode(&f), (UnsupportedTypeError{reflect.TypeOf(&f)}).Error())
	assert.EqualError(t, enc.Encode(&ch), (UnsupportedTypeError{reflect.TypeOf(&ch)}).Error())
}

func TestEncoderCopiedByValuePanic(t *testing.T) {
	assert.PanicsWithValue(t, "hprose/encoding: illegal use of non-zero Encoder copied by value", func() {
		sb := &strings.Builder{}
		enc := NewEncoder(sb).Simple(false)
		enc.Encode(1)
		enc2 := *enc
		enc2.Encode(2)
	})
}

func TestEncoderBytes(t *testing.T) {
	enc := new(Encoder)
	assert.Nil(t, enc.Bytes())
	enc.Encode(1)
	assert.Equal(t, []byte{'1'}, enc.Bytes())
}

func TestEncoderSimple(t *testing.T) {
	enc := new(Encoder)
	enc.Simple(false)
	enc.Encode("hello")
	enc.Encode("hello")
	assert.Equal(t, `s5"hello"r0;`, enc.String())
	enc.Simple(true)
	enc.Encode("hello")
	enc.Encode("hello")
	assert.Equal(t, `s5"hello"r0;s5"hello"s5"hello"`, enc.String())
}

func TestWriteTag(t *testing.T) {
	enc := new(Encoder)
	enc.WriteTag(TagHeader)
	enc.Encode(map[string]string{"id": "12345"})
	assert.Equal(t, `Hm1{s2"id"s5"12345"}`, enc.String())
}

func BenchmarkHproseEncodeSlice(b *testing.B) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	slice := []int16{
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
	}
	for i := 0; i < b.N; i++ {
		enc.Encode(slice)
	}
}

func BenchmarkSampleHproseEncodeSlice(b *testing.B) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb)
	slice := []int16{
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
	}
	for i := 0; i < b.N; i++ {
		enc.Encode(slice)
	}
}

func BenchmarkJsonEncodeSlice(b *testing.B) {
	sb := &strings.Builder{}
	enc := jsoniter.NewEncoder(sb)
	slice := []int16{
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
	}
	for i := 0; i < b.N; i++ {
		enc.Encode(slice)
	}
}

func BenchmarkHproseEncodeArray(b *testing.B) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	array := [50]int16{
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
	}
	for i := 0; i < b.N; i++ {
		enc.Encode(array)
	}
}

func BenchmarkSampleHproseEncodeArray(b *testing.B) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb)
	array := [50]int16{
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
	}
	for i := 0; i < b.N; i++ {
		enc.Encode(array)
	}
}

func BenchmarkJsonEncodeArray(b *testing.B) {
	sb := &strings.Builder{}
	enc := jsoniter.NewEncoder(sb)
	array := [50]int16{
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
		1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
	}
	for i := 0; i < b.N; i++ {
		enc.Encode(array)
	}
}

func BenchmarkHproseEncodeMap(b *testing.B) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	m := map[int16]int16{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
		6: 6,
		7: 7,
		8: 8,
		9: 9,
		0: 0,
	}
	for i := 0; i < b.N; i++ {
		enc.Encode(m)
	}
}

func BenchmarkSampleHproseEncodeMap(b *testing.B) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb)
	m := map[int16]int16{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
		6: 6,
		7: 7,
		8: 8,
		9: 9,
		0: 0,
	}
	for i := 0; i < b.N; i++ {
		enc.Encode(m)
	}
}

func BenchmarkJsonEncodeMap(b *testing.B) {
	sb := &strings.Builder{}
	enc := jsoniter.NewEncoder(sb)
	m := map[int16]int16{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
		6: 6,
		7: 7,
		8: 8,
		9: 9,
		0: 0,
	}
	for i := 0; i < b.N; i++ {
		enc.Encode(m)
	}
}

func BenchmarkHproseEncodeStruct(b *testing.B) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	type TestStruct struct {
		Name     string
		Age      int
		Birthday time.Time
		Male     bool
	}
	ts := &TestStruct{
		Name:     "Tom",
		Age:      18,
		Birthday: time.Date(2002, 1, 2, 3, 4, 5, 6, time.Local),
		Male:     true,
	}
	for i := 0; i < b.N; i++ {
		enc.Encode(ts)
	}
}

func BenchmarkSampleHproseEncodeStruct(b *testing.B) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb)
	type TestStruct struct {
		Name     string
		Age      int
		Birthday time.Time
		Male     bool
	}
	ts := &TestStruct{
		Name:     "Tom",
		Age:      18,
		Birthday: time.Date(2002, 1, 2, 3, 4, 5, 6, time.Local),
		Male:     true,
	}
	for i := 0; i < b.N; i++ {
		enc.Encode(ts)
	}
}

func BenchmarkJsonEncodeStruct(b *testing.B) {
	sb := &strings.Builder{}
	enc := jsoniter.NewEncoder(sb)
	type TestStruct struct {
		Name     string
		Age      int
		Birthday time.Time
		Male     bool
	}
	ts := &TestStruct{
		Name:     "Tom",
		Age:      18,
		Birthday: time.Date(2002, 1, 2, 3, 4, 5, 6, time.Local),
		Male:     true,
	}
	for i := 0; i < b.N; i++ {
		enc.Encode(ts)
	}
}
