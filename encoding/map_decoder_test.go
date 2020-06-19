/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/map_decoder_test.go                             |
|                                                          |
| LastModified: Jun 19, 2020                               |
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

func TestDecodeIntIntMap(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{1, 2, 3, 4, 5})
	enc.Encode([]float32{1, 2, 3, 4, 5})
	enc.Encode([]string{"1", "2", "3", "4", "5"})
	enc.Encode(map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var m map[int]int
	expected := map[int]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int]int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int]int`) // 1
}

func TestDecodeIntInt8Map(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{1, 2, 3, 4, 5})
	enc.Encode([]float32{1, 2, 3, 4, 5})
	enc.Encode([]string{"1", "2", "3", "4", "5"})
	enc.Encode(map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var m map[int]int8
	expected := map[int]int8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int]int8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int]int8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int]int8`) // 1
}

func BenchmarkDecodeIntIntMap(b *testing.B) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode(map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32})
	enc.Encode(nil)
	enc.Encode("")
	data := ([]byte)(sb.String())
	dec := &Decoder{}
	var m map[int]int
	for i := 0; i < b.N; i++ {
		dec.ResetBytes(data)
		dec.Decode(&m)
		dec.Decode(&m)
		dec.Decode(&m)
	}
}

func BenchmarkJsonDecodeIntIntMap(b *testing.B) {
	sb := new(strings.Builder)
	enc := jsoniter.NewEncoder(sb)
	enc.Encode(map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32})
	enc.Encode(nil)
	enc.Encode("")
	data := ([]byte)(sb.String())
	var m map[int]int
	for i := 0; i < b.N; i++ {
		dec := jsoniter.NewDecoder(bytes.NewReader(data))
		dec.Decode(&m)
		dec.Decode(&m)
		dec.Decode(&m)
	}
}
