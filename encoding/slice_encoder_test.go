/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/slice_encoder_test.go                           |
|                                                          |
| LastModified: Mar 20, 2020                               |
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
	enc := NewEncoder(sb).Simple(false)
	hello := []byte("Hello")
	assert.NoError(t, enc.Encode([]byte("")))
	assert.NoError(t, enc.Encode(hello))
	assert.NoError(t, enc.Encode(hello))
	assert.NoError(t, enc.Encode(&hello))
	assert.NoError(t, enc.Encode(&hello))
	assert.NoError(t, enc.Encode([]byte("Pokémon")))
	assert.NoError(t, enc.Encode([]byte("中文")))
	assert.NoError(t, enc.Encode([]byte("🐱🐶")))
	assert.NoError(t, enc.Encode([]byte("👩‍👩‍👧‍👧")))
	assert.Equal(t, `b""b5"Hello"b5"Hello"b5"Hello"r3;b8"Pokémon"b6"中文"b8"🐱🐶"b25"👩‍👩‍👧‍👧"`, sb.String())
}

func TestEncodeUint16Slice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	slice := []uint16{1, 2, 3, 4, 5}
	var nilslice []uint16
	assert.NoError(t, enc.Encode(nilslice))
	assert.NoError(t, enc.Encode([]uint16{}))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.Equal(t, `na{}a5{12345}a5{12345}a5{12345}r3;`, sb.String())
}

func TestEncodeUint32Slice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	slice := []uint32{1, 2, 3, 4, 5}
	var nilslice []uint32
	assert.NoError(t, enc.Encode(nilslice))
	assert.NoError(t, enc.Encode([]uint32{}))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.Equal(t, `na{}a5{12345}a5{12345}a5{12345}r3;`, sb.String())
}

func TestEncodeUint64Slice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	slice := []uint64{1, 2, 3, 4, 5}
	var nilslice []uint64
	assert.NoError(t, enc.Encode(nilslice))
	assert.NoError(t, enc.Encode([]uint64{}))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.Equal(t, `na{}a5{12345}a5{12345}a5{12345}r3;`, sb.String())
}

func TestEncodeUintSlice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	slice := []uint{1, 2, 3, 4, 5}
	var nilslice []uint
	assert.NoError(t, enc.Encode(nilslice))
	assert.NoError(t, enc.Encode([]uint{}))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.Equal(t, `na{}a5{12345}a5{12345}a5{12345}r3;`, sb.String())
}

func TestEncodeInt8Slice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	slice := []int8{1, 2, 3, 4, 5}
	var nilslice []int8
	assert.NoError(t, enc.Encode(nilslice))
	assert.NoError(t, enc.Encode([]int8{}))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.Equal(t, `na{}a5{12345}a5{12345}a5{12345}r3;`, sb.String())
}

func TestEncodeInt16Slice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	slice := []int16{1, 2, 3, 4, 5}
	var nilslice []int16
	assert.NoError(t, enc.Encode(nilslice))
	assert.NoError(t, enc.Encode([]int16{}))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.Equal(t, `na{}a5{12345}a5{12345}a5{12345}r3;`, sb.String())
}

func TestEncodeInt32Slice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	slice := []int32{1, 2, 3, 4, 5}
	var nilslice []int32
	assert.NoError(t, enc.Encode(nilslice))
	assert.NoError(t, enc.Encode([]int32{}))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.Equal(t, `na{}a5{12345}a5{12345}a5{12345}r3;`, sb.String())
}

func TestEncodeInt64Slice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	slice := []int64{1, 2, 3, 4, 5}
	var nilslice []int64
	assert.NoError(t, enc.Encode(nilslice))
	assert.NoError(t, enc.Encode([]int64{}))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.Equal(t, `na{}a5{12345}a5{12345}a5{12345}r3;`, sb.String())
}

func TestEncodeIntSlice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	slice := []int{1, 2, 3, 4, 5}
	var nilslice []int
	assert.NoError(t, enc.Encode(nilslice))
	assert.NoError(t, enc.Encode([]int{}))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.Equal(t, `na{}a5{12345}a5{12345}a5{12345}r3;`, sb.String())
}

func TestEncodeFloat32Slice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	slice := []float32{1, 2, 3, 4, 5}
	var nilslice []float32
	assert.NoError(t, enc.Encode(nilslice))
	assert.NoError(t, enc.Encode([]float32{}))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.Equal(t, `na{}a5{d1;d2;d3;d4;d5;}a5{d1;d2;d3;d4;d5;}a5{d1;d2;d3;d4;d5;}r3;`, sb.String())
}

func TestEncodeFloat64Slice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	slice := []float64{1, 2, 3, 4, 5}
	var nilslice []float64
	assert.NoError(t, enc.Encode(nilslice))
	assert.NoError(t, enc.Encode([]float64{}))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.Equal(t, `na{}a5{d1;d2;d3;d4;d5;}a5{d1;d2;d3;d4;d5;}a5{d1;d2;d3;d4;d5;}r3;`, sb.String())
}

func TestEncodeBoolSlice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	slice := []bool{true, false}
	var nilslice []bool
	assert.NoError(t, enc.Encode(nilslice))
	assert.NoError(t, enc.Encode([]bool{}))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.Equal(t, `na{}a2{tf}a2{tf}a2{tf}r3;`, sb.String())
}

func TestEncodeStringSlice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	slice := []string{"hello", "world"}
	var nilslice []string
	assert.NoError(t, enc.Encode(nilslice))
	assert.NoError(t, enc.Encode([]string{}))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.Equal(t, `na{}a2{s5"hello"s5"world"}a2{r2;r3;}a2{r2;r3;}r5;`, sb.String())
}

func TestEncodeComplex64Slice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	slice := []complex64{complex(1, 2), complex(3, 4)}
	var nilslice []complex64
	assert.NoError(t, enc.Encode(nilslice))
	assert.NoError(t, enc.Encode([]complex64{}))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.Equal(t, `na{}a2{a2{d1;d2;}a2{d3;d4;}}a2{a2{d1;d2;}a2{d3;d4;}}a2{a2{d1;d2;}a2{d3;d4;}}r7;`, sb.String())
}

func TestEncodeComplex128Slice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	slice := []complex128{complex(1, 2), complex(3, 4)}
	var nilslice []complex128
	assert.NoError(t, enc.Encode(nilslice))
	assert.NoError(t, enc.Encode([]complex128{}))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.Equal(t, `na{}a2{a2{d1;d2;}a2{d3;d4;}}a2{a2{d1;d2;}a2{d3;d4;}}a2{a2{d1;d2;}a2{d3;d4;}}r7;`, sb.String())
}

func TestEncodeInterfaceSlice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	slice := []interface{}{1, "hello", true}
	var nilslice []interface{}
	assert.NoError(t, enc.Encode(nilslice))
	assert.NoError(t, enc.Encode([]interface{}{}))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.Equal(t, `na{}a3{1s5"hello"t}a3{1r2;t}a3{1r2;t}r4;`, sb.String())
}

func TestEncodeBigIntSlice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	slice := []*big.Int{big.NewInt(1), big.NewInt(2)}
	var nilslice []*big.Int
	assert.NoError(t, enc.Encode(nilslice))
	assert.NoError(t, enc.Encode([]*big.Int{}))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.NoError(t, enc.Encode(&slice))
	assert.Equal(t, `na{}a2{l1;l2;}a2{l1;l2;}a2{l1;l2;}r3;`, sb.String())
}

func TestEncode2dSlice(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	assert.NoError(t, enc.Encode([][]int{
		{1, 2, 3}, {4, 5, 6},
	}))
	assert.Equal(t, `a2{a3{123}a3{456}}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode([][]int8{
		{1, 2, 3}, {4, 5, 6},
	}))
	assert.Equal(t, `a2{a3{123}a3{456}}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode([][]int16{
		{1, 2, 3}, {4, 5, 6},
	}))
	assert.Equal(t, `a2{a3{123}a3{456}}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode([][]int32{
		{1, 2, 3}, {4, 5, 6},
	}))
	assert.Equal(t, `a2{a3{123}a3{456}}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode([][]int64{
		{1, 2, 3}, {4, 5, 6},
	}))
	assert.Equal(t, `a2{a3{123}a3{456}}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode([][]uint{
		{1, 2, 3}, {4, 5, 6},
	}))
	assert.Equal(t, `a2{a3{123}a3{456}}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode([][]uint8{
		{1, 2, 3}, {4, 5, 6},
	}))
	assert.Equal(t, "a2{b3\"\x01\x02\x03\"b3\"\x04\x05\x06\"}", sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode([][]uint16{
		{1, 2, 3}, {4, 5, 6},
	}))
	assert.Equal(t, `a2{a3{123}a3{456}}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode([][]uint32{
		{1, 2, 3}, {4, 5, 6},
	}))
	assert.Equal(t, `a2{a3{123}a3{456}}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode([][]uint64{
		{1, 2, 3}, {4, 5, 6},
	}))
	assert.Equal(t, `a2{a3{123}a3{456}}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode([][]bool{
		{true, false, true}, {false, true, false},
	}))
	assert.Equal(t, `a2{a3{tft}a3{ftf}}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode([][]float32{
		{1, 2, 3}, {4, 5, 6},
	}))
	assert.Equal(t, `a2{a3{d1;d2;d3;}a3{d4;d5;d6;}}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode([][]float64{
		{1, 2, 3}, {4, 5, 6},
	}))
	assert.Equal(t, `a2{a3{d1;d2;d3;}a3{d4;d5;d6;}}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode([][]complex64{
		{complex(1, 2), complex(3, 4)}, {5, 6},
	}))
	assert.Equal(t, `a2{a2{a2{d1;d2;}a2{d3;d4;}}a2{d5;d6;}}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode([][]complex128{
		{complex(1, 2), complex(3, 4)}, {5, 6},
	}))
	assert.Equal(t, `a2{a2{a2{d1;d2;}a2{d3;d4;}}a2{d5;d6;}}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode([][]string{
		{"hello", "world"}, {"hello", "🐱🐶"},
	}))
	assert.Equal(t, `a2{a2{s5"hello"s5"world"}a2{r2;s4"🐱🐶"}}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode([][]interface{}{
		{1, 2.0}, {"🐱🐶"},
	}))
	assert.Equal(t, `a2{a2{1d2;}a1{s4"🐱🐶"}}`, sb.String())
	enc.Reset()
	sb.Reset()
}
