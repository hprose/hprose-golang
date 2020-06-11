/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/array_decoder_test.go                           |
|                                                          |
| LastModified: Jun 11, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"bytes"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestDecodeIntArray(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
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
	dec.Decode(&array)
	assert.Equal(t, [5]int{}, array) // ""
	dec.Decode(&array)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to [5]int`) // 1
}

func TestDecodeCustomIntArray(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
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
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to [5]encoding.Int`) // 1
}

func BenchmarkDecodeIntArray(b *testing.B) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
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
