/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/encoder_test.go                              |
|                                                          |
| LastModified: Mar 19, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"math"
	"math/big"
	"strings"
	"testing"
	"time"

	jsoniter "github.com/json-iterator/go"
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
	if err := enc.Encode(uptr); err != nil {
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
	if err := enc.Encode(&uptr); err != nil {
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
		`n0123456789i10;td3.1415927;d3.141592653589793;`+
		`eu我s5"Hello"a2{d1;d2;}a2{d3;d4;}d5;d6;l0;d1;s3"2/3"l4;`+
		`n0123456789i10;td3.1415927;d3.141592653589793;`+
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

func TestTeset(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	slice := []interface{}{1, "hello", true}
	var nilslice []interface{}
	if err := enc.Encode(nilslice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode([]interface{}{}); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(slice); err != nil {
		t.Error(err)
	}
	enc.Reset()
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if sb.String() != `na{}a3{1s5"hello"t}a3{1r2;t}a3{1s5"hello"t}r0;` {
		t.Error(sb)
	}
}

func TestEncodeByteArray(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	hello := [5]byte{'H', 'e', 'l', 'l', 'o'}
	if err := enc.Encode([0]byte{}); err != nil {
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
	if sb.String() != `b""b5"Hello"b5"Hello"b5"Hello"r3;` {
		t.Error(sb)
	}
}

func TestEncodeBigIntArray(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	array := [2]*big.Int{big.NewInt(1), big.NewInt(2)}
	var emptyArray [0]*big.Int
	if err := enc.Encode(emptyArray); err != nil {
		t.Error(err)
	}
	if err := enc.Encode([2]*big.Int{}); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(array); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(array); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&array); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&array); err != nil {
		t.Error(err)
	}
	if sb.String() != `a{}a2{nn}a2{l1;l2;}a2{l1;l2;}a2{l1;l2;}r4;` {
		t.Error(sb)
	}
}

func TestEncodeNil(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	var x interface{} = nil
	var xp interface{} = &x
	var i interface{} = (****int)(nil)
	var ip interface{} = &i
	if err := enc.Encode(nil); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(x); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&x); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&xp); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(i); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&i); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&ip); err != nil {
		t.Error(err)
	}
	if sb.String() != `nnnnnnn` {
		t.Error(sb)
	}
}

func TestWriteNil(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	var x interface{} = nil
	var xp interface{} = &x
	var i interface{} = (****int)(nil)
	var ip interface{} = &i
	if err := enc.Write(nil); err != nil {
		t.Error(err)
	}
	if err := enc.Write(x); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&x); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&xp); err != nil {
		t.Error(err)
	}
	if err := enc.Write(i); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&i); err != nil {
		t.Error(err)
	}
	if err := enc.Write(&ip); err != nil {
		t.Error(err)
	}
	if sb.String() != `nnnnnnn` {
		t.Error(sb)
	}
}

func TestWriteDuration(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	d := time.Duration(1000)
	dp := &d
	if err := enc.Write(d); err != nil {
		t.Error(err)
	}
	if err := enc.Write(dp); err != nil {
		t.Error(err)
	}
	if sb.String() != `i1000;i1000;` {
		t.Error(sb)
	}

}

func TestEncodeStringStringMap(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	m := map[string]string{
		"hello": "world",
	}
	var nilmap map[string]string
	if err := enc.Encode(nilmap); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(map[string]string{}); err != nil {
		t.Error(err)
	}
	if err := enc.Write(m); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&m); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(m); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&m); err != nil {
		t.Error(err)
	}
	if sb.String() != `nm{}m1{s5"hello"s5"world"}m1{r2;r3;}m1{r2;r3;}r4;` {
		t.Error(sb)
	}
}

func TestEncodeIntBigIntMap(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	m := map[int]*big.Int{
		1: big.NewInt(1),
	}
	var nilmap map[int]*big.Int
	if err := enc.Encode(nilmap); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(map[int]*big.Int{}); err != nil {
		t.Error(err)
	}
	if err := enc.Write(m); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&m); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(m); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&m); err != nil {
		t.Error(err)
	}
	if sb.String() != `nm{}m1{1l1;}m1{1l1;}m1{1l1;}r2;` {
		t.Error(sb)
	}
}

func TestEncodeStringInterfaceMap(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	m := map[string]interface{}{
		"hello": "world",
	}
	var nilmap map[string]interface{}
	if err := enc.Encode(nilmap); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(map[string]interface{}{}); err != nil {
		t.Error(err)
	}
	if err := enc.Write(m); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&m); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(m); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&m); err != nil {
		t.Error(err)
	}
	if sb.String() != `nm{}m1{s5"hello"s5"world"}m1{r2;r3;}m1{r2;r3;}r4;` {
		t.Error(sb)
	}
}

func BenchmarkEncodeSlice(b *testing.B) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
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

func BenchmarkEncodeArray(b *testing.B) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
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

func BenchmarkEncodeMap(b *testing.B) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
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

func BenchmarkEncodeStruct(b *testing.B) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	type TestStruct struct {
		Name     string
		Age      int
		Birthday time.Time
		Male     bool
	}
	ts := TestStruct{
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
	ts := TestStruct{
		Name:     "Tom",
		Age:      18,
		Birthday: time.Date(2002, 1, 2, 3, 4, 5, 6, time.Local),
		Male:     true,
	}
	for i := 0; i < b.N; i++ {
		enc.Encode(ts)
	}
}
