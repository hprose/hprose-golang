/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/array_decoder_test.go                           |
|                                                          |
| LastModified: Apr 27, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding_test

import (
	"bytes"
	"strings"
	"testing"

	. "github.com/hprose/hprose-golang/v3/encoding"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestDecodeIntArray(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{1, 2, 3, 4, 5, 6})
	enc.Encode([]float32{4, 3, 2, 1})
	enc.Encode([]string{"1", "2", "3", "4", "5"})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var array [5]int
	dec.Decode(&array)
	assert.Equal(t, [...]int{1, 2, 3, 4, 5}, array) // []int{1, 2, 3, 4, 5, 6}
	dec.Decode(&array)
	assert.Equal(t, [...]int{4, 3, 2, 1, 0}, array) // []float32{4, 3, 2, 1}
	dec.Decode(&array)
	assert.Equal(t, [...]int{1, 2, 3, 4, 5}, array) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&array)
	assert.Equal(t, [5]int{}, array) // nil
	array[3] = 3
	assert.Equal(t, [5]int{0, 0, 0, 3, 0}, array)
	dec.Decode(&array)
	assert.Equal(t, [5]int{}, array) // ""
	dec.Decode(&array)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to [5]int`) // 1
}

func TestDecodeIntPtrArray(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{1, 2, 3})
	dec := NewDecoder(([]byte)(sb.String()))
	var array [5]*int
	dec.Decode(&array)
	assert.Equal(t, 1, *array[0])
	assert.Equal(t, 2, *array[1])
	assert.Equal(t, 3, *array[2])
	assert.Equal(t, (*int)(nil), array[3])
	assert.Equal(t, (*int)(nil), array[4])
}

func TestDecodeCustomIntArray(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{1, 2, 3, 4, 5, 6})
	enc.Encode([]float32{4, 3, 2, 1})
	enc.Encode([]string{"1", "2", "3", "4", "5"})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	type Int int
	var array [5]Int
	dec.Decode(&array)
	assert.Equal(t, [...]Int{1, 2, 3, 4, 5}, array) // []int{1, 2, 3, 4, 5, 6}
	dec.Decode(&array)
	assert.Equal(t, [...]Int{4, 3, 2, 1, 0}, array) // []float32{4, 3, 2, 1}
	dec.Decode(&array)
	assert.Equal(t, [...]Int{1, 2, 3, 4, 5}, array) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&array)
	assert.Equal(t, [5]Int{}, array) // nil
	dec.Decode(&array)
	assert.Equal(t, [5]Int{}, array) // ""
	dec.Decode(&array)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to [5]encoding_test.Int`) // 1
}

func TestDecodeByteArray(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{1, 2, 3, 4, 5, 6})
	enc.Encode([]float32{4, 3, 2, 1})
	enc.Encode([]string{"1", "2", "3", "4", "5"})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode("123456789")
	enc.Encode("A")
	enc.Encode([]byte("OK"))
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var array [5]byte
	dec.Decode(&array)
	assert.Equal(t, [...]byte{1, 2, 3, 4, 5}, array) // []int{1, 2, 3, 4, 5, 6}
	dec.Decode(&array)
	assert.Equal(t, [...]byte{4, 3, 2, 1, 0}, array) // []float32{4, 3, 2, 1}
	dec.Decode(&array)
	assert.Equal(t, [...]byte{1, 2, 3, 4, 5}, array) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&array)
	assert.Equal(t, [5]byte{}, array) // nil
	dec.Decode(&array)
	assert.Equal(t, [5]byte{}, array) // ""
	dec.Decode(&array)
	assert.Equal(t, [5]byte{'1', '2', '3', '4', '5'}, array) // "123456789"
	dec.Decode(&array)
	assert.Equal(t, [5]byte{'A', 0, 0, 0, 0}, array) // "A"
	dec.Decode(&array)
	assert.Equal(t, [5]byte{'O', 'K', 0, 0, 0}, array) // []byte("OK")
	dec.Decode(&array)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to [5]uint8`) // 1
}

func TestDecodeInterfaceArray(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{1, 2, 3, 4, 5, 6})
	enc.Encode([]float32{4, 3, 2, 1})
	enc.Encode([]string{"1", "2", "3", "4", "5"})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var array [5]interface{}
	dec.Decode(&array)
	assert.Equal(t, [...]interface{}{1, 2, 3, 4, 5}, array) // []int{1, 2, 3, 4, 5, 6}
	dec.Decode(&array)
	assert.Equal(t, [...]interface{}{4.0, 3.0, 2.0, 1.0, nil}, array) // []float32{4, 3, 2, 1}
	dec.Decode(&array)
	assert.Equal(t, [...]interface{}{"1", "2", "3", "4", "5"}, array) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&array)
	assert.Equal(t, [5]interface{}{}, array) // nil
	dec.Decode(&array)
	assert.Equal(t, [5]interface{}{}, array) // ""
	dec.Decode(&array)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to [5]interface {}`) // 1
}

func TestDecodeIntIntArray(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([][]int{{1, 2, 3}, {4, 5, 6}})
	enc.Encode([][]float32{{4, 3, 2}, {1}})
	enc.Encode([][]string{{"1", "2", "3"}, {"4", "5"}})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var array [2][3]int
	dec.Decode(&array)
	assert.Equal(t, [2][3]int{{1, 2, 3}, {4, 5, 6}}, array) // [][]int{{1, 2, 3}, {4, 5, 6}}
	dec.Decode(&array)
	assert.Equal(t, [2][3]int{{4, 3, 2}, {1, 0, 0}}, array) // [][]float32{{4, 3, 2}, {1}}
	dec.Decode(&array)
	assert.Equal(t, [2][3]int{{1, 2, 3}, {4, 5, 0}}, array) // [][]string{{"1", "2", "3"}, {"4", "5"}}
	dec.Decode(&array)
	assert.Equal(t, [2][3]int{}, array) // nil
	dec.Decode(&array)
	assert.Equal(t, [2][3]int{}, array) // ""
	dec.Decode(&array)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to [2][3]int`) // 1
}

func BenchmarkDecodeIntArray(b *testing.B) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{1, 2, 3, 4, 5})
	enc.Encode([]float32{1, 2, 3, 4, 5})
	enc.Encode([]string{"1", "2", "3", "4", "5"})
	enc.Encode(nil)
	enc.Encode("")
	data := ([]byte)(sb.String())
	dec := &Decoder{}
	var array [5]int
	for i := 0; i < b.N; i++ {
		dec.ResetBytes(data)
		dec.Decode(&array)
		dec.Decode(&array)
		dec.Decode(&array)
		dec.Decode(&array)
		dec.Decode(&array)
	}
}

func BenchmarkJsonDecodeIntArray(b *testing.B) {
	sb := new(strings.Builder)
	enc := jsoniter.NewEncoder(sb)
	enc.Encode([]int{1, 2, 3, 4, 5})
	enc.Encode([]float32{1, 2, 3, 4, 5})
	enc.Encode([]string{"1", "2", "3", "4", "5"})
	enc.Encode(nil)
	enc.Encode("")
	data := ([]byte)(sb.String())
	var array [5]int
	for i := 0; i < b.N; i++ {
		dec := jsoniter.NewDecoder(bytes.NewReader(data))
		dec.Decode(&array)
		dec.Decode(&array)
		dec.Decode(&array)
		dec.Decode(&array)
		dec.Decode(&array)
	}
}

func BenchmarkDecodeInterfaceArray(b *testing.B) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{1, 2, 3, 4, 5})
	enc.Encode([]float32{1, 2, 3, 4, 5})
	enc.Encode([]string{"1", "2", "3", "4", "5"})
	enc.Encode(nil)
	enc.Encode("")
	data := ([]byte)(sb.String())
	dec := &Decoder{}
	var array [5]interface{}
	for i := 0; i < b.N; i++ {
		dec.ResetBytes(data)
		dec.Decode(&array)
		dec.Decode(&array)
		dec.Decode(&array)
		dec.Decode(&array)
		dec.Decode(&array)
	}
}

func BenchmarkJsonDecodeInterfaceArray(b *testing.B) {
	sb := new(strings.Builder)
	enc := jsoniter.NewEncoder(sb)
	enc.Encode([]int{1, 2, 3, 4, 5})
	enc.Encode([]float32{1, 2, 3, 4, 5})
	enc.Encode([]string{"1", "2", "3", "4", "5"})
	enc.Encode(nil)
	enc.Encode("")
	data := ([]byte)(sb.String())
	var array [5]interface{}
	for i := 0; i < b.N; i++ {
		dec := jsoniter.NewDecoder(bytes.NewReader(data))
		dec.Decode(&array)
		dec.Decode(&array)
		dec.Decode(&array)
		dec.Decode(&array)
		dec.Decode(&array)
	}
}

func BenchmarkDecodeIntIntArray(b *testing.B) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([][]int{{1, 2, 3}, {4, 5, 6}})
	enc.Encode([][]float32{{4, 3, 2}, {1}})
	enc.Encode([][]string{{"1", "2", "3"}, {"4", "5"}})
	enc.Encode(nil)
	enc.Encode("")
	data := ([]byte)(sb.String())
	dec := &Decoder{}
	var array [2][3]int
	for i := 0; i < b.N; i++ {
		dec.ResetBytes(data)
		dec.Decode(&array)
		dec.Decode(&array)
		dec.Decode(&array)
		dec.Decode(&array)
		dec.Decode(&array)
	}
}

func BenchmarkJsonDecodeIntIntArray(b *testing.B) {
	sb := new(strings.Builder)
	enc := jsoniter.NewEncoder(sb)
	enc.Encode([][]int{{1, 2, 3}, {4, 5, 6}})
	enc.Encode([][]float32{{4, 3, 2}, {1}})
	enc.Encode([][]string{{"1", "2", "3"}, {"4", "5"}})
	enc.Encode(nil)
	enc.Encode("")
	data := ([]byte)(sb.String())
	var array [2][3]int
	for i := 0; i < b.N; i++ {
		dec := jsoniter.NewDecoder(bytes.NewReader(data))
		dec.Decode(&array)
		dec.Decode(&array)
		dec.Decode(&array)
		dec.Decode(&array)
		dec.Decode(&array)
	}
}
