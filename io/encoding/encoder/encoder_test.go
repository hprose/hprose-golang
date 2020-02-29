/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/encoder/encoder_test.go                      |
|                                                          |
| LastModified: Feb 25, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoder

import (
	"math"
	"math/big"
	"strings"
	"testing"
)

func TestEncoderEncode(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, false)
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
	if err := enc.Encode(nil); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(i); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(i8); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(i16); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(i32); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(i64); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(u); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(u8); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(u16); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(u32); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(u64); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(b); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(f32); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(f64); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(e); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(c); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(s); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(c64); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(c128); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(r64); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(r128); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(*bi); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(*bf); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(*br); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(*bri); err != nil {
		t.Error(err)
	}
	if err := enc.Encode((*int)(nil)); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&i); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&i8); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&i16); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&i32); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&i64); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&u); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&u8); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&u16); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&u32); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&u64); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&b); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&f32); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&f64); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&e); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&c); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&s); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&c64); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&c128); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&r64); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&r128); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(bi); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(bf); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(br); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(bri); err != nil {
		t.Error(err)
	}
	if sb.String() != ``+
		`n0123456789td3.1415927;d3.141592653589793;`+
		`eu我s5"Hello"a2{d1;d2;}a2{d3;d4;}d5;d6;l0;d1;s3"2/3"l4;`+
		`n0123456789td3.1415927;d3.141592653589793;`+
		`eu我r0;a2{d1;d2;}a2{d3;d4;}d5;d6;l0;d1;s3"2/3"l4;` {
		t.Error(sb)
	}
}

func TestEncoderWrite(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, false)
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
	if err := enc.Write(nil); err != nil {
		t.Error(err)
	}
	if err := enc.Write(i); err != nil {
		t.Error(err)
	}
	if err := enc.Write(i8); err != nil {
		t.Error(err)
	}
	if err := enc.Write(i16); err != nil {
		t.Error(err)
	}
	if err := enc.Write(i32); err != nil {
		t.Error(err)
	}
	if err := enc.Write(i64); err != nil {
		t.Error(err)
	}
	if err := enc.Write(u); err != nil {
		t.Error(err)
	}
	if err := enc.Write(u8); err != nil {
		t.Error(err)
	}
	if err := enc.Write(u16); err != nil {
		t.Error(err)
	}
	if err := enc.Write(u32); err != nil {
		t.Error(err)
	}
	if err := enc.Write(u64); err != nil {
		t.Error(err)
	}
	if err := enc.Write(b); err != nil {
		t.Error(err)
	}
	if err := enc.Write(f32); err != nil {
		t.Error(err)
	}
	if err := enc.Write(f64); err != nil {
		t.Error(err)
	}
	if err := enc.Write(e); err != nil {
		t.Error(err)
	}
	if err := enc.Write(c); err != nil {
		t.Error(err)
	}
	if err := enc.Write(s); err != nil {
		t.Error(err)
	}
	if err := enc.Write(c64); err != nil {
		t.Error(err)
	}
	if err := enc.Write(c128); err != nil {
		t.Error(err)
	}
	if err := enc.Write(r64); err != nil {
		t.Error(err)
	}
	if err := enc.Write(r128); err != nil {
		t.Error(err)
	}
	if err := enc.Write(*bi); err != nil {
		t.Error(err)
	}
	if err := enc.Write(*bf); err != nil {
		t.Error(err)
	}
	if err := enc.Write(*br); err != nil {
		t.Error(err)
	}
	if err := enc.Write(*bri); err != nil {
		t.Error(err)
	}
	if err := enc.Write((*int)(nil)); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&i); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&i8); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&i16); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&i32); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&i64); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&u); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&u8); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&u16); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&u32); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&u64); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&b); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&f32); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&f64); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&e); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&c); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&s); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&c64); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&c128); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&r64); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&r128); err != nil {
		t.Error(err)
	}
	if err := enc.Write(bi); err != nil {
		t.Error(err)
	}
	if err := enc.Write(bf); err != nil {
		t.Error(err)
	}
	if err := enc.Write(br); err != nil {
		t.Error(err)
	}
	if err := enc.Write(bri); err != nil {
		t.Error(err)
	}
	if sb.String() != ``+
		`n0123456789td3.1415927;d3.141592653589793;`+
		`s""s1"我"s5"Hello"a2{d1;d2;}a2{d3;d4;}d5;d6;l0;d1;s3"2/3"l4;`+
		`n0123456789td3.1415927;d3.141592653589793;`+
		`s""s1"我"s5"Hello"a2{d1;d2;}a2{d3;d4;}d5;d6;l0;d1;s3"2/3"l4;` {
		t.Error(sb)
	}
}

func TestWriteString(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	if err := enc.Write(""); err != nil {
		t.Error(err)
	}
	if err := enc.Write("Hello"); err != nil {
		t.Error(err)
	}
	if err := enc.Write("Pokémon"); err != nil {
		t.Error(err)
	}
	if err := enc.Write("中文"); err != nil {
		t.Error(err)
	}
	if err := enc.Write("🐱🐶"); err != nil {
		t.Error(err)
	}
	if err := enc.Write("👩‍👩‍👧‍👧"); err != nil {
		t.Error(err)
	}
	if sb.String() != `s""s5"Hello"s7"Pokémon"s2"中文"s4"🐱🐶"s11"👩‍👩‍👧‍👧"` {
		t.Error(sb)
	}
}

func TestWriteBadString(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	if err := enc.Write(string([]byte{254, 254})); err != nil {
		t.Error(err)
	}
	if sb.String() != `b2"`+string([]byte{254, 254})+`"` {
		t.Error(sb)
	}
}

func TestEncodeBytes(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	hello := []byte("Hello")
	if err := enc.Encode([]byte("")); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(hello); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(hello); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&hello); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&hello); err != nil {
		t.Error(err)
	}
	if err := enc.Encode([]byte("Pokémon")); err != nil {
		t.Error(err)
	}
	if err := enc.Encode([]byte("中文")); err != nil {
		t.Error(err)
	}
	if err := enc.Encode([]byte("🐱🐶")); err != nil {
		t.Error(err)
	}
	if err := enc.Encode([]byte("👩‍👩‍👧‍👧")); err != nil {
		t.Error(err)
	}
	if sb.String() != `b""b5"Hello"b5"Hello"b5"Hello"r3;b8"Pokémon"b6"中文"b8"🐱🐶"b25"👩‍👩‍👧‍👧"` {
		t.Error(sb)
	}
}

func TestEncodeUint16Slice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	slice := []uint16{1, 2, 3, 4, 5}
	var nilslice []uint16
	if err := enc.Encode(nilslice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode([]uint16{}); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if sb.String() != `na{}a5{12345}a5{12345}a5{12345}r3;` {
		t.Error(sb)
	}
}

func TestEncodeUint32Slice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	slice := []uint32{1, 2, 3, 4, 5}
	var nilslice []uint32
	if err := enc.Encode(nilslice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode([]uint32{}); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if sb.String() != `na{}a5{12345}a5{12345}a5{12345}r3;` {
		t.Error(sb)
	}
}

func TestEncodeUint64Slice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	slice := []uint64{1, 2, 3, 4, 5}
	var nilslice []uint64
	if err := enc.Encode(nilslice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode([]uint64{}); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if sb.String() != `na{}a5{12345}a5{12345}a5{12345}r3;` {
		t.Error(sb)
	}
}

func TestEncodeUintSlice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	slice := []uint{1, 2, 3, 4, 5}
	var nilslice []uint
	if err := enc.Encode(nilslice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode([]uint{}); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if sb.String() != `na{}a5{12345}a5{12345}a5{12345}r3;` {
		t.Error(sb)
	}
}

func TestEncodeInt8Slice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	slice := []int8{1, 2, 3, 4, 5}
	var nilslice []int8
	if err := enc.Encode(nilslice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode([]int8{}); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if sb.String() != `na{}a5{12345}a5{12345}a5{12345}r3;` {
		t.Error(sb)
	}
}

func TestEncodeInt16Slice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	slice := []int16{1, 2, 3, 4, 5}
	var nilslice []int16
	if err := enc.Encode(nilslice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode([]int16{}); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if sb.String() != `na{}a5{12345}a5{12345}a5{12345}r3;` {
		t.Error(sb)
	}
}

func TestEncodeInt32Slice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	slice := []int32{1, 2, 3, 4, 5}
	var nilslice []int32
	if err := enc.Encode(nilslice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode([]int32{}); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if sb.String() != `na{}a5{12345}a5{12345}a5{12345}r3;` {
		t.Error(sb)
	}
}

func TestEncodeInt64Slice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	slice := []int64{1, 2, 3, 4, 5}
	var nilslice []int64
	if err := enc.Encode(nilslice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode([]int64{}); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if sb.String() != `na{}a5{12345}a5{12345}a5{12345}r3;` {
		t.Error(sb)
	}
}

func TestEncodeIntSlice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	slice := []int{1, 2, 3, 4, 5}
	var nilslice []int
	if err := enc.Encode(nilslice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode([]int{}); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if sb.String() != `na{}a5{12345}a5{12345}a5{12345}r3;` {
		t.Error(sb)
	}
}

func TestEncodeFloat32Slice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	slice := []float32{1, 2, 3, 4, 5}
	var nilslice []float32
	if err := enc.Encode(nilslice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode([]float32{}); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if sb.String() != `na{}a5{d1;d2;d3;d4;d5;}a5{d1;d2;d3;d4;d5;}a5{d1;d2;d3;d4;d5;}r3;` {
		t.Error(sb)
	}
}

func TestEncodeFloat64Slice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	slice := []float64{1, 2, 3, 4, 5}
	var nilslice []float64
	if err := enc.Encode(nilslice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode([]float64{}); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if sb.String() != `na{}a5{d1;d2;d3;d4;d5;}a5{d1;d2;d3;d4;d5;}a5{d1;d2;d3;d4;d5;}r3;` {
		t.Error(sb)
	}
}

func TestEncodeBoolSlice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	slice := []bool{true, false}
	var nilslice []bool
	if err := enc.Encode(nilslice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode([]bool{}); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if sb.String() != `na{}a2{tf}a2{tf}a2{tf}r3;` {
		t.Error(sb)
	}
}

func TestEncodeStringSlice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	slice := []string{"hello", "world"}
	var nilslice []string
	if err := enc.Encode(nilslice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode([]string{}); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if sb.String() != `na{}a2{s5"hello"s5"world"}a2{r2;r3;}a2{r2;r3;}r5;` {
		t.Error(sb)
	}
}

func TestEncodeComplex64Slice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	slice := []complex64{complex(1, 2), complex(3, 4)}
	var nilslice []complex64
	if err := enc.Encode(nilslice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode([]complex64{}); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if sb.String() != `na{}a2{a2{d1;d2;}a2{d3;d4;}}a2{a2{d1;d2;}a2{d3;d4;}}a2{a2{d1;d2;}a2{d3;d4;}}r7;` {
		t.Error(sb)
	}
}

func TestEncodeComplex128Slice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	slice := []complex128{complex(1, 2), complex(3, 4)}
	var nilslice []complex128
	if err := enc.Encode(nilslice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode([]complex128{}); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if sb.String() != `na{}a2{a2{d1;d2;}a2{d3;d4;}}a2{a2{d1;d2;}a2{d3;d4;}}a2{a2{d1;d2;}a2{d3;d4;}}r7;` {
		t.Error(sb)
	}
}

func TestEncodeBigIntSlice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	slice := []*big.Int{big.NewInt(1), big.NewInt(2)}
	var nilslice []*big.Int
	if err := enc.Encode(nilslice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode([]*big.Int{}); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if sb.String() != `na{}a2{l1;l2;}a2{l1;l2;}a2{l1;l2;}r3;` {
		t.Error(sb)
	}
}

func BenchmarkSlice1(b *testing.B) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	slice := []int16{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		enc.Encode(slice)
	}
}
