/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/slice_encoder_test.go                        |
|                                                          |
| LastModified: Mar 17, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"math/big"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestEncodeInterfaceSlice(t *testing.T) {
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
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&slice); err != nil {
		t.Error(err)
	}
	if sb.String() != `na{}a3{1s5"hello"t}a3{1r2;t}a3{1r2;t}r4;` {
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

func TestEncode2dSlice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, false)
	if err := enc.Encode([][]int{
		[]int{1, 2, 3}, []int{4, 5, 6},
	}); err != nil {
		t.Error(err)
	}
	assert.Equal(t, `a2{a3{123}a3{456}}`, sb.String())
	enc.Reset()
	sb.Reset()

	if err := enc.Encode([][]int8{
		[]int8{1, 2, 3}, []int8{4, 5, 6},
	}); err != nil {
		t.Error(err)
	}
	assert.Equal(t, `a2{a3{123}a3{456}}`, sb.String())
	enc.Reset()
	sb.Reset()

	if err := enc.Encode([][]int16{
		[]int16{1, 2, 3}, []int16{4, 5, 6},
	}); err != nil {
		t.Error(err)
	}
	assert.Equal(t, `a2{a3{123}a3{456}}`, sb.String())
	enc.Reset()
	sb.Reset()

	if err := enc.Encode([][]int32{
		[]int32{1, 2, 3}, []int32{4, 5, 6},
	}); err != nil {
		t.Error(err)
	}
	assert.Equal(t, `a2{a3{123}a3{456}}`, sb.String())
	enc.Reset()
	sb.Reset()

	if err := enc.Encode([][]int64{
		[]int64{1, 2, 3}, []int64{4, 5, 6},
	}); err != nil {
		t.Error(err)
	}
	assert.Equal(t, `a2{a3{123}a3{456}}`, sb.String())
	enc.Reset()
	sb.Reset()

	if err := enc.Encode([][]uint{
		[]uint{1, 2, 3}, []uint{4, 5, 6},
	}); err != nil {
		t.Error(err)
	}
	assert.Equal(t, `a2{a3{123}a3{456}}`, sb.String())
	enc.Reset()
	sb.Reset()

	if err := enc.Encode([][]uint8{
		[]uint8{1, 2, 3}, []uint8{4, 5, 6},
	}); err != nil {
		t.Error(err)
	}
	assert.Equal(t, "a2{b3\"\x01\x02\x03\"b3\"\x04\x05\x06\"}", sb.String())
	enc.Reset()
	sb.Reset()

	if err := enc.Encode([][]uint16{
		[]uint16{1, 2, 3}, []uint16{4, 5, 6},
	}); err != nil {
		t.Error(err)
	}
	assert.Equal(t, `a2{a3{123}a3{456}}`, sb.String())
	enc.Reset()
	sb.Reset()

	if err := enc.Encode([][]uint32{
		[]uint32{1, 2, 3}, []uint32{4, 5, 6},
	}); err != nil {
		t.Error(err)
	}
	assert.Equal(t, `a2{a3{123}a3{456}}`, sb.String())
	enc.Reset()
	sb.Reset()

	if err := enc.Encode([][]uint64{
		[]uint64{1, 2, 3}, []uint64{4, 5, 6},
	}); err != nil {
		t.Error(err)
	}
	assert.Equal(t, `a2{a3{123}a3{456}}`, sb.String())
	enc.Reset()
	sb.Reset()

	if err := enc.Encode([][]bool{
		[]bool{true, false, true}, []bool{false, true, false},
	}); err != nil {
		t.Error(err)
	}
	assert.Equal(t, `a2{a3{tft}a3{ftf}}`, sb.String())
	enc.Reset()
	sb.Reset()

	if err := enc.Encode([][]float32{
		[]float32{1, 2, 3}, []float32{4, 5, 6},
	}); err != nil {
		t.Error(err)
	}
	assert.Equal(t, `a2{a3{d1;d2;d3;}a3{d4;d5;d6;}}`, sb.String())
	enc.Reset()
	sb.Reset()

	if err := enc.Encode([][]float64{
		[]float64{1, 2, 3}, []float64{4, 5, 6},
	}); err != nil {
		t.Error(err)
	}
	assert.Equal(t, `a2{a3{d1;d2;d3;}a3{d4;d5;d6;}}`, sb.String())
	enc.Reset()
	sb.Reset()

	if err := enc.Encode([][]complex64{
		[]complex64{complex(1, 2), complex(3, 4)}, []complex64{5, 6},
	}); err != nil {
		t.Error(err)
	}
	assert.Equal(t, `a2{a2{a2{d1;d2;}a2{d3;d4;}}a2{d5;d6;}}`, sb.String())
	enc.Reset()
	sb.Reset()

	if err := enc.Encode([][]complex128{
		[]complex128{complex(1, 2), complex(3, 4)}, []complex128{5, 6},
	}); err != nil {
		t.Error(err)
	}
	assert.Equal(t, `a2{a2{a2{d1;d2;}a2{d3;d4;}}a2{d5;d6;}}`, sb.String())
	enc.Reset()
	sb.Reset()

	if err := enc.Encode([][]string{
		[]string{"hello", "world"}, []string{"hello", "🐱🐶"},
	}); err != nil {
		t.Error(err)
	}
	assert.Equal(t, `a2{a2{s5"hello"s5"world"}a2{r2;s4"🐱🐶"}}`, sb.String())
	enc.Reset()
	sb.Reset()

	if err := enc.Encode([][]interface{}{
		[]interface{}{1, 2.0}, []interface{}{"🐱🐶"},
	}); err != nil {
		t.Error(err)
	}
	assert.Equal(t, `a2{a2{1d2;}a1{s4"🐱🐶"}}`, sb.String())
	enc.Reset()
	sb.Reset()
}
