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
 * io/writer_test.go                                      *
 *                                                        *
 * hprose writer test for Go.                             *
 *                                                        *
 * LastModified: Oct 29, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import (
	"container/list"
	"math"
	"math/big"
	"math/rand"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestSerializeNil(t *testing.T) {
	w := NewWriter(true)
	w.Serialize(nil)
	if w.String() != "n" {
		t.Error(w.String())
	}
}

func BenchmarkSerializeNil(b *testing.B) {
	w := NewWriter(true)
	for i := 0; i < b.N; i++ {
		w.Serialize(nil)
	}
}

func BenchmarkWriteNil(b *testing.B) {
	w := NewWriter(true)
	for i := 0; i < b.N; i++ {
		w.WriteNil()
	}
}

func TestSerializeTrue(t *testing.T) {
	w := NewWriter(true)
	w.Serialize(true)
	if w.String() != "t" {
		t.Error(w.String())
	}
}

func BenchmarkSerializeTrue(b *testing.B) {
	w := NewWriter(true)
	for i := 0; i < b.N; i++ {
		w.Serialize(true)
	}
}

func BenchmarkWriteTrue(b *testing.B) {
	w := NewWriter(true)
	for i := 0; i < b.N; i++ {
		w.WriteBool(true)
	}
}

func TestSerializeFalse(t *testing.T) {
	w := NewWriter(true)
	w.Serialize(false)
	if w.String() != "f" {
		t.Error(w.String())
	}
}

func BenchmarkSerializeFalse(b *testing.B) {
	w := NewWriter(true)
	for i := 0; i < b.N; i++ {
		w.Serialize(false)
	}
}

func BenchmarkWriteFalse(b *testing.B) {
	w := NewWriter(true)
	for i := 0; i < b.N; i++ {
		w.WriteBool(false)
	}
}

func TestSerializeDigit(t *testing.T) {
	w := NewWriter(true)
	for i := 0; i <= 9; i++ {
		w.Clear()
		w.Serialize(i)
		if w.String() != strconv.Itoa(i) {
			t.Error(w.String())
		}
	}
}

func TestSerializeInt(t *testing.T) {
	w := NewWriter(true)
	for i := 0; i <= 100; i++ {
		w.Clear()
		x := rand.Intn(math.MaxInt32-10) + 10
		w.Serialize(x)
		if w.String() != "i"+strconv.Itoa(x)+";" {
			t.Error(w.String())
		}
	}
	for i := 0; i <= 100; i++ {
		w.Clear()
		x := rand.Intn(math.MaxInt64-math.MaxInt32-1) + math.MaxInt32 + 1
		w.Serialize(x)
		if w.String() != "l"+strconv.Itoa(x)+";" {
			t.Error(w.String())
		}
	}
}

func BenchmarkSerializeInt(b *testing.B) {
	w := NewWriter(true)
	for i := 0; i < b.N; i++ {
		w.Serialize(i)
	}
}

func BenchmarkWriteInt(b *testing.B) {
	w := NewWriter(true)
	for i := 0; i < b.N; i++ {
		w.WriteInt(int64(i))
	}
}

func TestSerializeInt8(t *testing.T) {
	w := NewWriter(true)
	for i := 0; i <= 9; i++ {
		w.Clear()
		w.Serialize(int8(i))
		if w.String() != strconv.Itoa(i) {
			t.Error(w.String())
		}
	}
	for i := 10; i <= 127; i++ {
		w.Clear()
		w.Serialize(int8(i))
		if w.String() != "i"+strconv.Itoa(i)+";" {
			t.Error(w.String())
		}
	}
	for i := -128; i < 0; i++ {
		w.Clear()
		w.Serialize(int8(i))
		if w.String() != "i"+strconv.Itoa(i)+";" {
			t.Error(w.String())
		}
	}
}

func TestSerializeInt16(t *testing.T) {
	w := NewWriter(true)
	w.Serialize(int16(math.MaxInt16))
	if w.String() != "i"+strconv.Itoa(math.MaxInt16)+";" {
		t.Error(w.String())
	}
}

func TestSerializeInt32(t *testing.T) {
	w := NewWriter(true)
	w.Serialize(int32(math.MaxInt32))
	if w.String() != "i"+strconv.Itoa(math.MaxInt32)+";" {
		t.Error(w.String())
	}
}

func BenchmarkSerializeInt32(b *testing.B) {
	w := NewWriter(true)
	for i := 0; i < b.N; i++ {
		w.Serialize(int32(i))
	}
}

func TestSerializeInt64(t *testing.T) {
	w := NewWriter(true)
	w.Serialize(int64(math.MaxInt32))
	if w.String() != "i"+strconv.Itoa(math.MaxInt32)+";" {
		t.Error(w.String())
	}
	w.Clear()
	w.Serialize(int64(math.MaxInt64))
	if w.String() != "l"+strconv.Itoa(math.MaxInt64)+";" {
		t.Error(w.String())
	}
}

func TestSerializeUint(t *testing.T) {
	w := NewWriter(true)
	for i := 0; i <= 100; i++ {
		w.Clear()
		x := rand.Intn(math.MaxInt32-10) + 10
		w.Serialize(uint(x))
		if w.String() != "i"+strconv.Itoa(x)+";" {
			t.Error(w.String())
		}
	}
	for i := 0; i <= 100; i++ {
		w.Clear()
		x := rand.Intn(math.MaxInt64-math.MaxInt32-1) + math.MaxInt32 + 1
		w.Serialize(uint(x))
		if w.String() != "l"+strconv.Itoa(x)+";" {
			t.Error(w.String())
		}
	}
}

func TestSerializeUint8(t *testing.T) {
	w := NewWriter(true)
	for i := 0; i <= 9; i++ {
		w.Clear()
		w.Serialize(uint8(i))
		if w.String() != strconv.Itoa(i) {
			t.Error(w.String())
		}
	}
	for i := 10; i <= 255; i++ {
		w.Clear()
		w.Serialize(uint8(i))
		if w.String() != "i"+strconv.Itoa(i)+";" {
			t.Error(w.String())
		}
	}
}

func TestSerializeUint16(t *testing.T) {
	w := NewWriter(true)
	w.Serialize(uint16(math.MaxUint16))
	if w.String() != "i"+strconv.Itoa(math.MaxUint16)+";" {
		t.Error(w.String())
	}
}

func TestSerializeUint32(t *testing.T) {
	w := NewWriter(true)
	w.Serialize(uint32(math.MaxUint32))
	if w.String() != "l"+strconv.Itoa(math.MaxUint32)+";" {
		t.Error(w.String())
	}
}

func TestSerializeUint64(t *testing.T) {
	w := NewWriter(true)
	w.Serialize(uint64(math.MaxUint32))
	if w.String() != "l"+strconv.Itoa(math.MaxUint32)+";" {
		t.Error(w.String())
	}
	w.Clear()
	w.Serialize(uint64(math.MaxUint64))
	if w.String() != "l"+strconv.FormatUint(math.MaxUint64, 10)+";" {
		t.Error(w.String())
	}
}

func BenchmarkSerializeUint64(b *testing.B) {
	w := NewWriter(true)
	for i := 0; i < b.N; i++ {
		w.Serialize(uint64(i))
	}
}

func BenchmarkWriteUint64(b *testing.B) {
	w := NewWriter(true)
	for i := 0; i < b.N; i++ {
		w.WriteUint(uint64(i))
	}
}

func TestSerializeUintptr(t *testing.T) {
	w := NewWriter(true)
	w.Serialize(uintptr(123))
	if w.String() != "i123;" {
		t.Error(w.String())
	}
}

func TestSerializeFloat32(t *testing.T) {
	w := NewWriter(true)
	testdata := map[float32]string{
		float32(math.NaN()):   "N",
		float32(math.Inf(1)):  "I+",
		float32(math.Inf(-1)): "I-",
		float32(3.14159):      "d3.14159;",
		math.MaxFloat32:       "d3.4028235e+38;",
	}
	for k, v := range testdata {
		w.Serialize(k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func BenchmarkSerializeFloat32(b *testing.B) {
	w := NewWriter(true)
	for i := 0; i < b.N; i++ {
		w.Serialize(float32(i))
	}
}

func BenchmarkWriteFloat32(b *testing.B) {
	w := NewWriter(true)
	for i := 0; i < b.N; i++ {
		w.WriteFloat(float64(i), 32)
	}
}

func TestSerializeFloat64(t *testing.T) {
	w := NewWriter(true)
	testdata := map[float64]string{
		math.NaN():       "N",
		math.Inf(1):      "I+",
		math.Inf(-1):     "I-",
		3.14159265358979: "d3.14159265358979;",
		math.MaxFloat64:  "d1.7976931348623157e+308;",
	}
	for k, v := range testdata {
		w.Serialize(k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func BenchmarkSerializeFloat64(b *testing.B) {
	w := NewWriter(true)
	for i := 0; i < b.N; i++ {
		w.Serialize(float64(i))
	}
}

func BenchmarkWriteFloat64(b *testing.B) {
	w := NewWriter(true)
	for i := 0; i < b.N; i++ {
		w.WriteFloat(float64(i), 64)
	}
}

func TestSerializeComplex64(t *testing.T) {
	w := NewWriter(true)
	w.Serialize(complex(float32(100), 0))
	if w.String() != "d100;" {
		t.Error(w.String())
	}
	w.Clear()
	w.Serialize(complex(0, float32(100)))
	if w.String() != "a2{d0;d100;}" {
		t.Error(w.String())
	}
}

func TestSerializeComplex128(t *testing.T) {
	w := NewWriter(true)
	w.Serialize(complex(100, 0))
	if w.String() != "d100;" {
		t.Error(w.String())
	}
	w.Clear()
	w.Serialize(complex(0, 100))
	if w.String() != "a2{d0;d100;}" {
		t.Error(w.String())
	}
}

func BenchmarkSerializeComplex128(b *testing.B) {
	w := NewWriter(true)
	for i := 0; i < b.N; i++ {
		w.Serialize(complex(float64(i), float64(i)))
	}
}

func BenchmarkWriteComplex128(b *testing.B) {
	w := NewWriter(true)
	for i := 0; i < b.N; i++ {
		w.WriteComplex128(complex(float64(i), float64(i)))
	}
}

func TestWriteTuple(t *testing.T) {
	w := NewWriter(true)
	w.WriteTuple()
	if w.String() != "a{}" {
		t.Error(w.String())
	}
	w.Clear()
	w.WriteTuple(1, 3.14, true)
	if w.String() != "a3{1d3.14;t}" {
		t.Error(w.String())
	}
}

func TestWriteBytes(t *testing.T) {
	w := NewWriter(true)
	testdata := map[*[]byte]string{
		&[]byte{'h', 'e', 'l', 'l', 'o'}: `b5"hello"`,
		&[]byte{}:                        `b""`,
	}
	for k, v := range testdata {
		w.WriteBytes(*k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestSerializeString(t *testing.T) {
	w := NewWriter(true)
	testdata := map[string]string{
		"":                            "e",
		"π":                           "uπ",
		"你":                           "u你",
		"你好":                          `s2"你好"`,
		"你好啊,hello!":                  `s10"你好啊,hello!"`,
		"🇨🇳":                          `s4"🇨🇳"`,
		string([]byte{128, 129, 130}): string([]byte{'b', '3', '"', 128, 129, 130, '"'}),
	}
	for k, v := range testdata {
		w.Serialize(k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestWriteString(t *testing.T) {
	w := NewWriter(true)
	testdata := map[string]string{
		"":                            "e",
		"π":                           "uπ",
		"你":                           "u你",
		"你好":                          `s2"你好"`,
		"你好啊,hello!":                  `s10"你好啊,hello!"`,
		"🇨🇳":                          `s4"🇨🇳"`,
		string([]byte{128, 129, 130}): string([]byte{'b', '3', '"', 128, 129, 130, '"'}),
	}
	for k, v := range testdata {
		w.WriteString(k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestSerializeStringPtr(t *testing.T) {
	w := NewWriter(true)
	testdata := map[string]string{
		"":                            "e",
		"π":                           "uπ",
		"你":                           "u你",
		"你好":                          `s2"你好"`,
		"你好啊,hello!":                  `s10"你好啊,hello!"`,
		"🇨🇳":                          `s4"🇨🇳"`,
		string([]byte{128, 129, 130}): string([]byte{'b', '3', '"', 128, 129, 130, '"'}),
	}
	for k, v := range testdata {
		w.Serialize(&k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestSerializeArray(t *testing.T) {
	w := NewWriter(true)
	testdata := map[interface{}]string{
		&[...]int{1, 2, 3}:                   "a3{123}",
		&[...]float64{1, 2, 3}:               "a3{d1;d2;d3;}",
		&[...]byte{'h', 'e', 'l', 'l', 'o'}:  `b5"hello"`,
		&[...]byte{}:                         `b""`,
		&[...]interface{}{1, 2.0, nil, true}: "a4{1d2;nt}",
		&[...]bool{true, false, true}:        "a3{tft}",
		&[...]int{}:                          "a{}",
		&[...]bool{}:                         "a{}",
		&[...]interface{}{}:                  "a{}",
	}
	for k, v := range testdata {
		w.Serialize(k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestSerializeSlice(t *testing.T) {
	w := NewWriter(true)
	testdata := map[interface{}]string{
		&[]int{1, 2, 3}:                   "a3{123}",
		&[]float64{1, 2, 3}:               "a3{d1;d2;d3;}",
		&[]byte{'h', 'e', 'l', 'l', 'o'}:  `b5"hello"`,
		&[]byte{}:                         `b""`,
		&[]interface{}{1, 2.0, nil, true}: "a4{1d2;nt}",
		&[]bool{true, false, true}:        "a3{tft}",
		&[]int{}:                          "a{}",
		&[]bool{}:                         "a{}",
		&[]interface{}{}:                  "a{}",
	}
	for k, v := range testdata {
		w.Serialize(k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestSerializeBoolSlice(t *testing.T) {
	w := NewWriter(true)
	testdata := map[*[]bool]string{
		&[]bool{true, false, true}: "a3{tft}",
		&[]bool{}:                  "a{}",
	}
	for k, v := range testdata {
		w.Serialize(*k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestSerializeIntSlice(t *testing.T) {
	w := NewWriter(true)
	testdata := map[*[]int]string{
		&[]int{1, 2, 3}: "a3{123}",
		&[]int{}:        "a{}",
	}
	for k, v := range testdata {
		w.Serialize(*k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestSerializeInt8Slice(t *testing.T) {
	w := NewWriter(true)
	testdata := map[*[]int8]string{
		&[]int8{1, 2, 3}: "a3{123}",
		&[]int8{}:        "a{}",
	}
	for k, v := range testdata {
		w.Serialize(*k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestSerializeInt16Slice(t *testing.T) {
	w := NewWriter(true)
	testdata := map[*[]int16]string{
		&[]int16{1, 2, 3}: "a3{123}",
		&[]int16{}:        "a{}",
	}
	for k, v := range testdata {
		w.Serialize(*k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestSerializeInt32Slice(t *testing.T) {
	w := NewWriter(true)
	testdata := map[*[]int32]string{
		&[]int32{1, 2, 3}: "a3{123}",
		&[]int32{}:        "a{}",
	}
	for k, v := range testdata {
		w.Serialize(*k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestSerializeInt64Slice(t *testing.T) {
	w := NewWriter(true)
	testdata := map[*[]int64]string{
		&[]int64{1, 2, 3}: "a3{123}",
		&[]int64{}:        "a{}",
	}
	for k, v := range testdata {
		w.Serialize(*k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestSerializeUintSlice(t *testing.T) {
	w := NewWriter(true)
	testdata := map[*[]uint]string{
		&[]uint{1, 2, 3}: "a3{123}",
		&[]uint{}:        "a{}",
	}
	for k, v := range testdata {
		w.Serialize(*k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestSerializeUint8Slice(t *testing.T) {
	w := NewWriter(true)
	testdata := map[*[]uint8]string{
		&[]uint8{1, 2, 3}: `b3"` + string([]byte{1, 2, 3}) + `"`,
		&[]uint8{}:        `b""`,
	}
	for k, v := range testdata {
		w.Serialize(*k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestSerializeUint16Slice(t *testing.T) {
	w := NewWriter(true)
	testdata := map[*[]uint16]string{
		&[]uint16{1, 2, 3}: "a3{123}",
		&[]uint16{}:        "a{}",
	}
	for k, v := range testdata {
		w.Serialize(*k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestSerializeUint32Slice(t *testing.T) {
	w := NewWriter(true)
	testdata := map[*[]uint32]string{
		&[]uint32{1, 2, 3}: "a3{123}",
		&[]uint32{}:        "a{}",
	}
	for k, v := range testdata {
		w.Serialize(*k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestSerializeUint64Slice(t *testing.T) {
	w := NewWriter(true)
	testdata := map[*[]uint64]string{
		&[]uint64{1, 2, 3}: "a3{123}",
		&[]uint64{}:        "a{}",
	}
	for k, v := range testdata {
		w.Serialize(*k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestSerializeUintptrSlice(t *testing.T) {
	w := NewWriter(true)
	testdata := map[*[]uintptr]string{
		&[]uintptr{1, 2, 3}: "a3{123}",
		&[]uintptr{}:        "a{}",
	}
	for k, v := range testdata {
		w.Serialize(*k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestSerializeFloat32Slice(t *testing.T) {
	w := NewWriter(true)
	testdata := map[*[]float32]string{
		&[]float32{1, 2, 3}: "a3{d1;d2;d3;}",
		&[]float32{}:        "a{}",
	}
	for k, v := range testdata {
		w.Serialize(*k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestSerializeFloat64Slice(t *testing.T) {
	w := NewWriter(true)
	testdata := map[*[]float64]string{
		&[]float64{1, 2, 3}: "a3{d1;d2;d3;}",
		&[]float64{}:        "a{}",
	}
	for k, v := range testdata {
		w.Serialize(*k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestSerializeComplex64Slice(t *testing.T) {
	w := NewWriter(true)
	testdata := map[*[]complex64]string{
		&[]complex64{complex(0, 0), complex(1, 0), complex(0, 1)}: "a3{d0;d1;a2{d0;d1;}}",
		&[]complex64{}: "a{}",
	}
	for k, v := range testdata {
		w.Serialize(*k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestSerializeComplex128Slice(t *testing.T) {
	w := NewWriter(true)
	testdata := map[*[]complex128]string{
		&[]complex128{complex(0, 0), complex(1, 0), complex(0, 1)}: "a3{d0;d1;a2{d0;d1;}}",
		&[]complex128{}:                                            "a{}",
	}
	for k, v := range testdata {
		w.Serialize(*k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestSerializeStringSlice(t *testing.T) {
	w := NewWriter(true)
	testdata := map[*[]string]string{
		&[]string{"", "π", "hello"}: `a3{euπs5"hello"}`,
		&[]string{}:                 "a{}",
	}
	for k, v := range testdata {
		w.Serialize(*k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestSerializeBytesSlice(t *testing.T) {
	w := NewWriter(true)
	testdata := map[*[][]byte]string{
		&[][]byte{[]byte(""), []byte("π"), []byte("hello")}: `a3{b""b2"π"b5"hello"}`,
		&[][]byte{}: "a{}",
	}
	for k, v := range testdata {
		w.Serialize(*k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func BenchmarkSerializeIntArray(b *testing.B) {
	w := NewWriter(true)
	array := [...]int{0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 1, 2, 3, 4, 0, 1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		w.Serialize(array)
	}
}

func BenchmarkSerializeIntSlice(b *testing.B) {
	w := NewWriter(true)
	slice := []int{0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 1, 2, 3, 4, 0, 1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		w.Serialize(slice)
	}
}

func BenchmarkSerializeBytes(b *testing.B) {
	w := NewWriter(true)
	slice := ([]byte)("你好,hello!")
	for i := 0; i < b.N; i++ {
		w.Serialize(slice)
	}
}

func BenchmarkWriteBytes(b *testing.B) {
	w := NewWriter(true)
	slice := ([]byte)("你好,hello!")
	for i := 0; i < b.N; i++ {
		w.WriteBytes(slice)
	}
}

func BenchmarkSerializeString(b *testing.B) {
	w := NewWriter(true)
	str := "你好,hello!"
	for i := 0; i < b.N; i++ {
		w.Serialize(str)
	}
}

func BenchmarkWriteString(b *testing.B) {
	w := NewWriter(true)
	str := "你好,hello!"
	for i := 0; i < b.N; i++ {
		w.WriteString(str)
	}
}

func TestSerializeBigInt(t *testing.T) {
	w := NewWriter(true)
	w.Serialize(big.NewInt(123))
	if w.String() != "l123;" {
		t.Error(w.String())
	}
	w.Clear()
	w.Serialize(*big.NewInt(123))
	if w.String() != "l123;" {
		t.Error(w.String())
	}
}

func BenchmarkWriteBigInt(b *testing.B) {
	w := NewWriter(true)
	x := big.NewInt(123)
	for i := 0; i < b.N; i++ {
		w.WriteBigInt(x)
	}
}

func BenchmarkSerializeBigInt(b *testing.B) {
	w := NewWriter(true)
	x := big.NewInt(123)
	for i := 0; i < b.N; i++ {
		w.Serialize(x)
	}
}

func TestSerializeBigRat(t *testing.T) {
	w := NewWriter(true)
	w.Serialize(big.NewRat(123, 1))
	if w.String() != "l123;" {
		t.Error(w.String())
	}
	w.Clear()
	w.Serialize(*big.NewRat(123, 2))
	if w.String() != `s5"123/2"` {
		t.Error(w.String())
	}
}

func BenchmarkWriteBigRat(b *testing.B) {
	w := NewWriter(true)
	x := big.NewRat(123, 2)
	for i := 0; i < b.N; i++ {
		w.WriteBigRat(x)
	}
}

func BenchmarkSerializeBigRat(b *testing.B) {
	w := NewWriter(true)
	x := big.NewRat(123, 2)
	for i := 0; i < b.N; i++ {
		w.Serialize(x)
	}
}

func TestSerializeBigFloat(t *testing.T) {
	w := NewWriter(true)
	w.Serialize(big.NewFloat(3.14159265358979))
	if w.String() != "d3.14159265358979;" {
		t.Error(w.String())
	}
	w.Clear()
	w.Serialize(*big.NewFloat(3.14159265358979))
	if w.String() != "d3.14159265358979;" {
		t.Error(w.String())
	}
}

func BenchmarkWriteBigFloat(b *testing.B) {
	w := NewWriter(true)
	x := big.NewFloat(3.14159265358979)
	for i := 0; i < b.N; i++ {
		w.WriteBigFloat(x)
	}
}

func BenchmarkSerializeBigFloat(b *testing.B) {
	w := NewWriter(true)
	x := big.NewFloat(3.14159265358979)
	for i := 0; i < b.N; i++ {
		w.Serialize(x)
	}
}

func TestWriteTime(t *testing.T) {
	w := NewWriter(true)
	testdata := map[time.Time]string{
		time.Date(1980, 12, 1, 0, 0, 0, 0, time.UTC):              "D19801201Z",
		time.Date(1970, 1, 1, 12, 34, 56, 0, time.UTC):            "T123456Z",
		time.Date(1970, 1, 1, 12, 34, 56, 789000000, time.UTC):    "T123456.789Z",
		time.Date(1970, 1, 1, 12, 34, 56, 789456000, time.UTC):    "T123456.789456Z",
		time.Date(1970, 1, 1, 12, 34, 56, 789456123, time.UTC):    "T123456.789456123Z",
		time.Date(1980, 12, 1, 12, 34, 56, 0, time.UTC):           "D19801201T123456Z",
		time.Date(1980, 12, 1, 12, 34, 56, 789000000, time.UTC):   "D19801201T123456.789Z",
		time.Date(1980, 12, 1, 12, 34, 56, 789456000, time.UTC):   "D19801201T123456.789456Z",
		time.Date(1980, 12, 1, 12, 34, 56, 789456123, time.UTC):   "D19801201T123456.789456123Z",
		time.Date(1980, 12, 1, 0, 0, 0, 0, time.Local):            "D19801201;",
		time.Date(1970, 1, 1, 12, 34, 56, 0, time.Local):          "T123456;",
		time.Date(1980, 12, 1, 12, 34, 56, 0, time.Local):         "D19801201T123456;",
		time.Date(1980, 12, 1, 12, 34, 56, 789456123, time.Local): "D19801201T123456.789456123;",
	}
	for k, v := range testdata {
		w.WriteTime(&k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func TestWriteTimeRef(t *testing.T) {
	d := time.Date(1980, 12, 1, 0, 0, 0, 0, time.UTC)
	w := NewWriter(false)
	w.WriteTime(&d)
	w.WriteTime(&d)
	if w.String() != "D19801201Zr0;" {
		t.Error(w.String())
	}
}

func TestSerializeTime(t *testing.T) {
	w := NewWriter(true)
	testdata := map[time.Time]string{
		time.Date(1980, 12, 1, 0, 0, 0, 0, time.UTC):              "D19801201Z",
		time.Date(1970, 1, 1, 12, 34, 56, 0, time.UTC):            "T123456Z",
		time.Date(1970, 1, 1, 12, 34, 56, 789000000, time.UTC):    "T123456.789Z",
		time.Date(1970, 1, 1, 12, 34, 56, 789456000, time.UTC):    "T123456.789456Z",
		time.Date(1970, 1, 1, 12, 34, 56, 789456123, time.UTC):    "T123456.789456123Z",
		time.Date(1980, 12, 1, 12, 34, 56, 0, time.UTC):           "D19801201T123456Z",
		time.Date(1980, 12, 1, 12, 34, 56, 789000000, time.UTC):   "D19801201T123456.789Z",
		time.Date(1980, 12, 1, 12, 34, 56, 789456000, time.UTC):   "D19801201T123456.789456Z",
		time.Date(1980, 12, 1, 12, 34, 56, 789456123, time.UTC):   "D19801201T123456.789456123Z",
		time.Date(1980, 12, 1, 0, 0, 0, 0, time.Local):            "D19801201;",
		time.Date(1970, 1, 1, 12, 34, 56, 0, time.Local):          "T123456;",
		time.Date(1980, 12, 1, 12, 34, 56, 0, time.Local):         "D19801201T123456;",
		time.Date(1980, 12, 1, 12, 34, 56, 789456123, time.Local): "D19801201T123456.789456123;",
	}
	for k, v := range testdata {
		w.Serialize(&k)
		if w.String() != v {
			t.Error(w.String())
		}
		w.Clear()
	}
}

func BenchmarkWriteTime(b *testing.B) {
	w := NewWriter(true)
	x := time.Date(1980, 12, 1, 12, 34, 56, 789456123, time.UTC)
	for i := 0; i < b.N; i++ {
		w.WriteTime(&x)
	}
}

func BenchmarkSerializeTime(b *testing.B) {
	w := NewWriter(true)
	x := time.Date(1980, 12, 1, 12, 34, 56, 789456123, time.UTC)
	for i := 0; i < b.N; i++ {
		w.Serialize(x)
	}
}

func TestSerializeList(t *testing.T) {
	w := NewWriter(false)
	lst := list.New()
	w.Serialize(lst)
	if w.String() != "a{}" {
		t.Error(w.String())
	}
	w.Clear()
	w.Reset()
	lst.PushBack(1)
	lst.PushBack("hello")
	lst.PushBack(nil)
	lst.PushBack(3.14159)
	w.Serialize(lst)
	w.Serialize(lst)
	if w.String() != `a4{1s5"hello"nd3.14159;}r0;` {
		t.Error(w.String())
	}
}

func TestWriteList(t *testing.T) {
	w := NewWriter(false)
	lst := list.New()
	w.WriteList(lst)
	if w.String() != "a{}" {
		t.Error(w.String())
	}
	w.Clear()
	w.Reset()
	lst.PushBack(1)
	lst.PushBack("hello")
	lst.PushBack(nil)
	lst.PushBack(3.14159)
	w.WriteList(lst)
	w.WriteList(lst)
	if w.String() != `a4{1s5"hello"nd3.14159;}r0;` {
		t.Error(w.String())
	}
}

func BenchmarkSerializeList(b *testing.B) {
	w := NewWriter(true)
	lst := list.New()
	lst.PushBack(1)
	lst.PushBack("hello")
	lst.PushBack(nil)
	lst.PushBack(3.14159)
	for i := 0; i < b.N; i++ {
		w.Serialize(lst)
	}
}

func BenchmarkWriteList(b *testing.B) {
	w := NewWriter(true)
	lst := list.New()
	lst.PushBack(1)
	lst.PushBack("hello")
	lst.PushBack(nil)
	lst.PushBack(3.14159)
	for i := 0; i < b.N; i++ {
		w.WriteList(lst)
	}
}

func TestWriterMap(t *testing.T) {
	w := NewWriter(true)
	m := make(map[string]interface{})
	w.Serialize(m)
	if w.String() != "m{}" {
		t.Error(w.String())
	}
	w.Clear()
	m["name"] = "Tom"
	m["age"] = 36
	m["male"] = true
	w.Serialize(m)
	s := w.String()
	s1 := `m3{s4"name"s3"Tom"s3"age"i36;s4"male"t}`
	s2 := `m3{s3"age"i36;s4"male"ts4"name"s3"Tom"}`
	s3 := `m3{s3"age"i36;s4"name"s3"Tom"s4"male"t}`
	s4 := `m3{s4"name"s3"Tom"s4"male"ts3"age"i36;}`
	s5 := `m3{s4"male"ts3"age"i36;s4"name"s3"Tom"}`
	s6 := `m3{s4"male"ts4"name"s3"Tom"s3"age"i36;}`
	if s != s1 && s != s2 && s != s3 && s != s4 && s != s5 && s != s6 {
		t.Error(w.String())
	}
}

func TestWriterMapRef(t *testing.T) {
	w := NewWriter(false)
	m := make(map[string]interface{})
	w.Serialize(m)
	if w.String() != "m{}" {
		t.Error(w.String())
	}
	w.Reset()
	w.Clear()
	m["name"] = "Tom"
	m["age"] = 36
	m["male"] = true
	w.Serialize(&m)
	w.Serialize(&m)
	s := w.String()
	s1 := `m3{s4"name"s3"Tom"s3"age"i36;s4"male"t}r0;`
	s2 := `m3{s3"age"i36;s4"male"ts4"name"s3"Tom"}r0;`
	s3 := `m3{s3"age"i36;s4"name"s3"Tom"s4"male"t}r0;`
	s4 := `m3{s4"name"s3"Tom"s4"male"ts3"age"i36;}r0;`
	s5 := `m3{s4"male"ts3"age"i36;s4"name"s3"Tom"}r0;`
	s6 := `m3{s4"male"ts4"name"s3"Tom"s3"age"i36;}r0;`
	if s != s1 && s != s2 && s != s3 && s != s4 && s != s5 && s != s6 {
		t.Error(w.String())
	}
}

func BenchmarkSerializeStringKeyMap(b *testing.B) {
	w := NewWriter(true)
	m := make(map[string]interface{})
	m["name"] = "Tom"
	m["age"] = 36
	m["male"] = true
	for i := 0; i < b.N; i++ {
		w.Serialize(m)
	}
}

func BenchmarkSerializeEmptyMap(b *testing.B) {
	w := NewWriter(true)
	m := make(map[string]interface{})
	for i := 0; i < b.N; i++ {
		w.Serialize(m)
	}
}

func BenchmarkSerializeInterfaceKeyMap(b *testing.B) {
	w := NewWriter(true)
	m := make(map[interface{}]interface{})
	m["name"] = "Tom"
	m["age"] = 36
	m["male"] = true
	for i := 0; i < b.N; i++ {
		w.Serialize(&m)
	}
}

func TestSerializeStruct(t *testing.T) {
	type TestStruct struct {
		ID int `hprose:"id"`
	}
	type TestStruct1 struct {
		TestStruct
		Name string
		Age  *int
	}
	type TestStruct2 struct {
		OOXX bool `hprose:"ooxx"`
		*TestStruct2
		TestStruct1
		Test     TestStruct
		birthday time.Time
	}
	st := TestStruct2{}
	st.TestStruct2 = &st
	st.ID = 100
	st.Name = "Tom"
	age := 18
	st.Age = &age
	st.OOXX = false
	st.Test.ID = 200
	Register((*TestStruct)(nil), "Test", "hprose")
	Register((*TestStruct1)(nil), "Test1", "hprose")
	Register((*TestStruct2)(nil), "Test2", "hprose")
	w := NewWriter(false)
	w.Serialize(st)
	s := `c5"Test2"6{s4"ooxx"s11"testStruct2"s2"id"s4"name"s3"age"s4"test"}o0{fo0{fr7;i100;s3"Tom"i18;c4"Test"1{s2"id"}o1{i200;}}i100;s3"Tom"i18;o1{i200;}}`
	if w.String() != s {
		t.Error(w.String())
	}
}

func BenchmarkSerializeStruct(b *testing.B) {
	type TestStruct struct {
		ID int `hprose:"id"`
	}
	type TestStruct1 struct {
		TestStruct
		Name string
		Age  *int
	}
	type TestStruct2 struct {
		OOXX bool `hprose:"ooxx"`
		*TestStruct2
		TestStruct1
		Test     TestStruct
		birthday time.Time
	}
	st := TestStruct2{}
	st.TestStruct2 = &st
	st.ID = 100
	st.Name = "Tom"
	age := 18
	st.Age = &age
	st.OOXX = false
	st.Test.ID = 200
	Register((*TestStruct)(nil), "Test", "hprose")
	Register((*TestStruct1)(nil), "Test1", "hprose")
	Register((*TestStruct2)(nil), "Test2", "hprose")
	w := NewWriter(false)
	for i := 0; i < b.N; i++ {
		w.Serialize(st)
	}
}

func TestSerializeStructPtr(t *testing.T) {
	type Quotient struct {
		Quo, Rem int
	}
	q := Quotient{10, 1}
	pq := &q
	w := NewWriter(false)
	w.Serialize(&pq)
	s := `c8"Quotient"2{s3"quo"s3"rem"}o0{i10;1}`
	if w.String() != s {
		t.Error(w.String())
	}
	w.Reset()
	w.Serialize(&pq)
	s = `c8"Quotient"2{s3"quo"s3"rem"}o0{i10;1}c8"Quotient"2{s3"quo"s3"rem"}o0{i10;1}`
	if w.String() != s {
		t.Error(w.String())
	}
}

func TestSerializeBigIntPtr(t *testing.T) {
	w := NewWriter(true)
	bi := big.NewInt(123)
	pbi := &bi
	w.Serialize(&pbi)
	if w.String() != "l123;" {
		t.Error(w.String())
	}
}

func TestSerializeValue(t *testing.T) {
	w := NewWriter(true)
	w.Serialize(reflect.ValueOf(123))
	if w.String() != "i123;" {
		t.Error(w.String())
	}
}

func TestWriteSlice(t *testing.T) {
	w := NewWriter(true)
	w.WriteSlice([]reflect.Value{})
	if w.String() != `a{}` {
		t.Error(w.String())
	}
	w.Clear()
	w.Reset()
	w.WriteSlice([]reflect.Value{reflect.ValueOf(123), reflect.ValueOf("Hello")})
	if w.String() != `a2{i123;s5"Hello"}` {
		t.Error(w.String())
	}
}

func TestWriteStringSlice(t *testing.T) {
	w := NewWriter(true)
	w.WriteStringSlice([]string{})
	if w.String() != `a{}` {
		t.Error(w.String())
	}
	w.Clear()
	w.Reset()
	w.WriteStringSlice([]string{"你好", "Hello"})
	if w.String() != `a2{s2"你好"s5"Hello"}` {
		t.Error(w.String())
	}
}

func TestWriteSliceArray(t *testing.T) {
	w := NewWriter(false)
	w.Serialize([0]*[]string{})
	if w.String() != `a{}` {
		t.Error(w.String())
	}
	w.Clear()
	w.Reset()
	s := []string{"你好", "Hello"}
	w.Serialize([2]*[]string{&s, &s})
	if w.String() != `a2{a2{s2"你好"s5"Hello"}r1;}` {
		t.Error(w.String())
	}
	w.Clear()
	w.Reset()
	w.Serialize([2][]string{s, s})
	if w.String() != `a2{a2{s2"你好"s5"Hello"}a2{s2"你好"s5"Hello"}}` {
		t.Error(w.String())
	}
}

func TestWriteSliceSlice(t *testing.T) {
	w := NewWriter(false)
	w.Serialize([]*[]string{})
	if w.String() != `a{}` {
		t.Error(w.String())
	}
	w.Clear()
	w.Reset()
	s := []string{"你好", "Hello"}
	w.Serialize([]*[]string{&s, &s})
	if w.String() != `a2{a2{s2"你好"s5"Hello"}r1;}` {
		t.Error(w.String())
	}
	w.Clear()
	w.Reset()
	w.Serialize([][]string{s, s})
	if w.String() != `a2{a2{s2"你好"s5"Hello"}a2{s2"你好"s5"Hello"}}` {
		t.Error(w.String())
	}
}

func TestWriteSliceSlice2(t *testing.T) {
	w := NewWriter(true)
	s := []string{"你好", "Hello"}
	w.Serialize([]*[]string{&s, &s})
	if w.String() != `a2{a2{s2"你好"s5"Hello"}a2{s2"你好"s5"Hello"}}` {
		t.Error(w.String())
	}
	w.Clear()
	w.Reset()
	w.Serialize([][]string{s, s})
	if w.String() != `a2{a2{s2"你好"s5"Hello"}a2{s2"你好"s5"Hello"}}` {
		t.Error(w.String())
	}
}

func TestWriteSliceMap(t *testing.T) {
	w := NewWriter(false)
	w.Serialize(map[string]*[]string{})
	if w.String() != `m{}` {
		t.Error(w.String())
	}
	w.Clear()
	w.Reset()
	s := []string{"你好", "Hello"}
	w.Serialize(map[string]*[]string{"e1": &s})
	if w.String() != `m1{s2"e1"a2{s2"你好"s5"Hello"}}` {
		t.Error(w.String())
	}
}

func TestSerializeInterface(t *testing.T) {
	w := NewWriter(false)
	var i interface{}
	w.Serialize(&i)
	i = 123
	w.Serialize(&i)
	if w.String() != `ni123;` {
		t.Error(w.String())
	}
}

func TestSerializeFunc(t *testing.T) {
	w := NewWriter(false)
	w.Serialize(func() {})
	if w.String() != `n` {
		t.Error(w.String())
	}
}
