/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/slice_decoder_test.go                           |
|                                                          |
| LastModified: Jun 11, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"bytes"
	"math/big"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestDecodeIntSlice(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
	enc.Encode([]int{1, 2, 3, 4, 5})
	enc.Encode([]float32{1, 2, 3, 4, 5})
	enc.Encode([]string{"1", "2", "3", "4", "5"})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var slice []int
	dec.Decode(&slice)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, slice) // []int{1, 2, 3, 4, 5}
	dec.Decode(&slice)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, slice) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&slice)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, slice) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&slice)
	assert.Nil(t, slice) // nil
	dec.Decode(&slice)
	assert.Equal(t, []int{}, slice) // ""
	dec.Decode(&slice)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to []int`) // 1
}

func TestDecodeCustomIntSlice(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
	enc.Encode([]int{1, 2, 3, 4, 5})
	enc.Encode([]float32{1, 2, 3, 4, 5})
	enc.Encode([]string{"1", "2", "3", "4", "5"})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	type Int int
	var slice []Int
	dec.Decode(&slice)
	assert.Equal(t, []Int{1, 2, 3, 4, 5}, slice) // []int{1, 2, 3, 4, 5}
	dec.Decode(&slice)
	assert.Equal(t, []Int{1, 2, 3, 4, 5}, slice) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&slice)
	assert.Equal(t, []Int{1, 2, 3, 4, 5}, slice) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&slice)
	assert.Nil(t, slice) // nil
	dec.Decode(&slice)
	assert.Equal(t, []Int{}, slice) // ""
	dec.Decode(&slice)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to []encoding.Int`) // 1
}

func BenchmarkDecodeIntSlice(b *testing.B) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
	enc.Encode([]int{1, 2, 3, 4, 5})
	enc.Encode([]float32{1, 2, 3, 4, 5})
	enc.Encode([]string{"1", "2", "3", "4", "5"})
	enc.Encode(nil)
	enc.Encode("")
	data := ([]byte)(sb.String())
	dec := &Decoder{}
	var slice []int
	for i := 0; i < b.N; i++ {
		dec.ResetBytes(data)
		dec.Decode(&slice)
		dec.Decode(&slice)
		dec.Decode(&slice)
		dec.Decode(&slice)
		dec.Decode(&slice)
	}
}

func BenchmarkDecodeCustomIntSlice(b *testing.B) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
	enc.Encode([]int{1, 2, 3, 4, 5})
	enc.Encode([]float32{1, 2, 3, 4, 5})
	enc.Encode([]string{"1", "2", "3", "4", "5"})
	enc.Encode(nil)
	enc.Encode("")
	data := ([]byte)(sb.String())
	dec := &Decoder{}
	type Int int
	var slice []Int
	for i := 0; i < b.N; i++ {
		dec.ResetBytes(data)
		dec.Decode(&slice)
		dec.Decode(&slice)
		dec.Decode(&slice)
		dec.Decode(&slice)
		dec.Decode(&slice)
	}
}

func BenchmarkJsonDecodeIntSlice(b *testing.B) {
	sb := new(strings.Builder)
	enc := jsoniter.NewEncoder(sb)
	enc.Encode([]int{1, 2, 3, 4, 5})
	enc.Encode([]float32{1, 2, 3, 4, 5})
	enc.Encode([]string{"1", "2", "3", "4", "5"})
	enc.Encode(nil)
	enc.Encode("")
	data := ([]byte)(sb.String())
	var slice []int
	for i := 0; i < b.N; i++ {
		dec := jsoniter.NewDecoder(bytes.NewReader(data))
		dec.Decode(&slice)
		dec.Decode(&slice)
		dec.Decode(&slice)
		dec.Decode(&slice)
		dec.Decode(&slice)
	}
}

func BenchmarkDecodeInt64Slice(b *testing.B) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
	enc.Encode([]int{1, 2, 3, 4, 5})
	enc.Encode([]float32{1, 2, 3, 4, 5})
	enc.Encode([]string{"1", "2", "3", "4", "5"})
	enc.Encode(nil)
	enc.Encode("")
	data := ([]byte)(sb.String())
	dec := &Decoder{}
	var slice []int64
	for i := 0; i < b.N; i++ {
		dec.ResetBytes(data)
		dec.Decode(&slice)
		dec.Decode(&slice)
		dec.Decode(&slice)
		dec.Decode(&slice)
		dec.Decode(&slice)
	}
}

func BenchmarkJsonDecodeInt64Slice(b *testing.B) {
	sb := new(strings.Builder)
	enc := jsoniter.NewEncoder(sb)
	enc.Encode([]int{1, 2, 3, 4, 5})
	enc.Encode([]float32{1, 2, 3, 4, 5})
	enc.Encode([]string{"1", "2", "3", "4", "5"})
	enc.Encode(nil)
	enc.Encode("")
	data := ([]byte)(sb.String())
	var slice []int64
	for i := 0; i < b.N; i++ {
		dec := jsoniter.NewDecoder(bytes.NewReader(data))
		dec.Decode(&slice)
		dec.Decode(&slice)
		dec.Decode(&slice)
		dec.Decode(&slice)
		dec.Decode(&slice)
	}
}
func BenchmarkDecodeBigIntSlice(b *testing.B) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
	enc.Encode([]int{1, 2, 3, 4, 5})
	enc.Encode([]float32{1, 2, 3, 4, 5})
	enc.Encode([]string{"1", "2", "3", "4", "5"})
	enc.Encode(nil)
	enc.Encode("")
	data := ([]byte)(sb.String())
	dec := &Decoder{}
	var slice []*big.Int
	for i := 0; i < b.N; i++ {
		dec.ResetBytes(data)
		dec.Decode(&slice)
		dec.Decode(&slice)
		dec.Decode(&slice)
		dec.Decode(&slice)
		dec.Decode(&slice)
	}
}

func BenchmarkJsonDecodeBigIntSlice(b *testing.B) {
	sb := new(strings.Builder)
	enc := jsoniter.NewEncoder(sb)
	enc.Encode([]int{1, 2, 3, 4, 5})
	enc.Encode([]float32{1, 2, 3, 4, 5})
	enc.Encode([]string{"1", "2", "3", "4", "5"})
	enc.Encode(nil)
	enc.Encode("")
	data := ([]byte)(sb.String())
	var slice []*big.Int
	for i := 0; i < b.N; i++ {
		dec := jsoniter.NewDecoder(bytes.NewReader(data))
		dec.Decode(&slice)
		dec.Decode(&slice)
		dec.Decode(&slice)
		dec.Decode(&slice)
		dec.Decode(&slice)
	}
}
