/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/map_decoder_test.go                             |
|                                                          |
| LastModified: Jun 27, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"bytes"
	"strings"
	"testing"
	"time"

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

func TestDecodeIntInt16Map(t *testing.T) {
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
	var m map[int]int16
	expected := map[int]int16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int]int16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int]int16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int]int16`) // 1
}

func TestDecodeIntInt32Map(t *testing.T) {
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
	var m map[int]int32
	expected := map[int]int32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int]int32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int]int32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int]int32`) // 1
}

func TestDecodeIntInt64Map(t *testing.T) {
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
	var m map[int]int64
	expected := map[int]int64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int]int64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int]int64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int]int64`) // 1
}

func TestDecodeIntUintMap(t *testing.T) {
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
	var m map[int]uint
	expected := map[int]uint{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int]uint{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int]uint{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int]uint`) // 1
}

func TestDecodeIntUint8Map(t *testing.T) {
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
	var m map[int]uint8
	expected := map[int]uint8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int]uint8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int]uint8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int]uint8`) // 1
}

func TestDecodeIntUint16Map(t *testing.T) {
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
	var m map[int]uint16
	expected := map[int]uint16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int]uint16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int]uint16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int]uint16`) // 1
}

func TestDecodeIntUint32Map(t *testing.T) {
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
	var m map[int]uint32
	expected := map[int]uint32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int]uint32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int]uint32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int]uint32`) // 1
}

func TestDecodeIntUint64Map(t *testing.T) {
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
	var m map[int]uint64
	expected := map[int]uint64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int]uint64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int]uint64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int]uint64`) // 1
}

func TestDecodeIntFloat32Map(t *testing.T) {
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
	var m map[int]float32
	expected := map[int]float32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int]float32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int]float32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int]float32`) // 1
}

func TestDecodeIntFloat64Map(t *testing.T) {
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
	var m map[int]float64
	expected := map[int]float64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int]float64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int]float64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int]float64`) // 1
}

func TestDecodeIntBoolMap(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{0, 1, 2, 3, 4})
	enc.Encode([]float32{0, 1, 2, 3, 4})
	enc.Encode(map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var m map[int]bool
	expected := map[int]bool{0: false, 1: true, 2: true, 3: true, 4: true}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int]bool{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int]bool`) // 1
}

func TestDecodeIntStringMap(t *testing.T) {
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
	var m map[int]string
	expected := map[int]string{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int]string{1: "1", 2: "4", 3: "8", 4: "16", 5: "32"}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int]string{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int]string`) // 1
}

func TestDecodeIntInterfaceMap(t *testing.T) {
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
	dec.RealType = RealTypeFloat32
	var m map[int]interface{}
	dec.Decode(&m)
	assert.Equal(t, map[int]interface{}{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[int]interface{}{0: float32(1), 1: float32(2), 2: float32(3), 3: float32(4), 4: float32(5)}, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[int]interface{}{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int]interface{}{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int]interface{}{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int]interface {}`) // 1
}

func TestDecodeIntCustomIntMap(t *testing.T) {
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
	type Int int
	var m map[int]Int
	expected := map[int]Int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int]Int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int]Int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int]encoding.Int`) // 1
}

func TestDecodeInt8IntMap(t *testing.T) {
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
	var m map[int8]int
	expected := map[int8]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int8]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int8]int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int8]int`) // 1
}

func TestDecodeInt8Int8Map(t *testing.T) {
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
	var m map[int8]int8
	expected := map[int8]int8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int8]int8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int8]int8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int8]int8`) // 1
}

func TestDecodeInt8Int16Map(t *testing.T) {
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
	var m map[int8]int16
	expected := map[int8]int16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int8]int16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int8]int16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int8]int16`) // 1
}

func TestDecodeInt8Int32Map(t *testing.T) {
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
	var m map[int8]int32
	expected := map[int8]int32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int8]int32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int8]int32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int8]int32`) // 1
}

func TestDecodeInt8Int64Map(t *testing.T) {
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
	var m map[int8]int64
	expected := map[int8]int64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int8]int64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int8]int64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int8]int64`) // 1
}

func TestDecodeInt8UintMap(t *testing.T) {
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
	var m map[int8]uint
	expected := map[int8]uint{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int8]uint{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int8]uint{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int8]uint`) // 1
}

func TestDecodeInt8Uint8Map(t *testing.T) {
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
	var m map[int8]uint8
	expected := map[int8]uint8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int8]uint8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int8]uint8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int8]uint8`) // 1
}

func TestDecodeInt8Uint16Map(t *testing.T) {
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
	var m map[int8]uint16
	expected := map[int8]uint16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int8]uint16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int8]uint16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int8]uint16`) // 1
}

func TestDecodeInt8Uint32Map(t *testing.T) {
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
	var m map[int8]uint32
	expected := map[int8]uint32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int8]uint32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int8]uint32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int8]uint32`) // 1
}

func TestDecodeInt8Uint64Map(t *testing.T) {
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
	var m map[int8]uint64
	expected := map[int8]uint64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int8]uint64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int8]uint64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int8]uint64`) // 1
}

func TestDecodeInt8Float32Map(t *testing.T) {
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
	var m map[int8]float32
	expected := map[int8]float32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int8]float32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int8]float32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int8]float32`) // 1
}

func TestDecodeInt8Float64Map(t *testing.T) {
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
	var m map[int8]float64
	expected := map[int8]float64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int8]float64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int8]float64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int8]float64`) // 1
}

func TestDecodeInt8BoolMap(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{0, 1, 2, 3, 4})
	enc.Encode([]float32{0, 1, 2, 3, 4})
	enc.Encode(map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var m map[int8]bool
	expected := map[int8]bool{0: false, 1: true, 2: true, 3: true, 4: true}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, map[int8]bool{1: true, 2: true, 3: true, 4: true, 5: true}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int8]bool{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int8]bool`) // 1
}

func TestDecodeInt8StringMap(t *testing.T) {
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
	var m map[int8]string
	expected := map[int8]string{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int8]string{1: "1", 2: "4", 3: "8", 4: "16", 5: "32"}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int8]string{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int8]string`) // 1
}

func TestDecodeInt8InterfaceMap(t *testing.T) {
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
	dec.RealType = RealTypeFloat32
	var m map[int8]interface{}
	dec.Decode(&m)
	assert.Equal(t, map[int8]interface{}{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[int8]interface{}{0: float32(1), 1: float32(2), 2: float32(3), 3: float32(4), 4: float32(5)}, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[int8]interface{}{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int8]interface{}{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int8]interface{}{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int8]interface {}`) // 1
}

func TestDecodeInt8CustomIntMap(t *testing.T) {
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
	type Int int
	var m map[int8]Int
	expected := map[int8]Int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int8]Int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int8]Int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int8]encoding.Int`) // 1
}

func TestDecodeInt16IntMap(t *testing.T) {
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
	var m map[int16]int
	expected := map[int16]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int16]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int16]int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int16]int`) // 1
}

func TestDecodeInt16Int8Map(t *testing.T) {
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
	var m map[int16]int8
	expected := map[int16]int8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int16]int8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int16]int8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int16]int8`) // 1
}

func TestDecodeInt16Int16Map(t *testing.T) {
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
	var m map[int16]int16
	expected := map[int16]int16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int16]int16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int16]int16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int16]int16`) // 1
}

func TestDecodeInt16Int32Map(t *testing.T) {
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
	var m map[int16]int32
	expected := map[int16]int32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int16]int32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int16]int32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int16]int32`) // 1
}

func TestDecodeInt16Int64Map(t *testing.T) {
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
	var m map[int16]int64
	expected := map[int16]int64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int16]int64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int16]int64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int16]int64`) // 1
}

func TestDecodeInt16UintMap(t *testing.T) {
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
	var m map[int16]uint
	expected := map[int16]uint{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int16]uint{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int16]uint{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int16]uint`) // 1
}

func TestDecodeInt16Uint8Map(t *testing.T) {
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
	var m map[int16]uint8
	expected := map[int16]uint8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int16]uint8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int16]uint8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int16]uint8`) // 1
}

func TestDecodeInt16Uint16Map(t *testing.T) {
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
	var m map[int16]uint16
	expected := map[int16]uint16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int16]uint16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int16]uint16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int16]uint16`) // 1
}

func TestDecodeInt16Uint32Map(t *testing.T) {
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
	var m map[int16]uint32
	expected := map[int16]uint32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int16]uint32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int16]uint32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int16]uint32`) // 1
}

func TestDecodeInt16Uint64Map(t *testing.T) {
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
	var m map[int16]uint64
	expected := map[int16]uint64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int16]uint64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int16]uint64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int16]uint64`) // 1
}

func TestDecodeInt16Float32Map(t *testing.T) {
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
	var m map[int16]float32
	expected := map[int16]float32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int16]float32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int16]float32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int16]float32`) // 1
}

func TestDecodeInt16Float64Map(t *testing.T) {
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
	var m map[int16]float64
	expected := map[int16]float64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int16]float64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int16]float64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int16]float64`) // 1
}

func TestDecodeInt16BoolMap(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{0, 1, 2, 3, 4})
	enc.Encode([]float32{0, 1, 2, 3, 4})
	enc.Encode(map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var m map[int16]bool
	expected := map[int16]bool{0: false, 1: true, 2: true, 3: true, 4: true}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, map[int16]bool{1: true, 2: true, 3: true, 4: true, 5: true}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int16]bool{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int16]bool`) // 1
}

func TestDecodeInt16StringMap(t *testing.T) {
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
	var m map[int16]string
	expected := map[int16]string{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int16]string{1: "1", 2: "4", 3: "8", 4: "16", 5: "32"}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int16]string{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int16]string`) // 1
}

func TestDecodeInt16InterfaceMap(t *testing.T) {
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
	dec.RealType = RealTypeFloat32
	var m map[int16]interface{}
	dec.Decode(&m)
	assert.Equal(t, map[int16]interface{}{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[int16]interface{}{0: float32(1), 1: float32(2), 2: float32(3), 3: float32(4), 4: float32(5)}, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[int16]interface{}{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int16]interface{}{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int16]interface{}{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int16]interface {}`) // 1
}

func TestDecodeInt16CustomIntMap(t *testing.T) {
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
	type Int int
	var m map[int16]Int
	expected := map[int16]Int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int16]Int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int16]Int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int16]encoding.Int`) // 1
}

func TestDecodeInt32IntMap(t *testing.T) {
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
	var m map[int32]int
	expected := map[int32]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int32]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int32]int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int32]int`) // 1
}

func TestDecodeInt32Int8Map(t *testing.T) {
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
	var m map[int32]int8
	expected := map[int32]int8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int32]int8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int32]int8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int32]int8`) // 1
}

func TestDecodeInt32Int16Map(t *testing.T) {
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
	var m map[int32]int16
	expected := map[int32]int16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int32]int16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int32]int16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int32]int16`) // 1
}

func TestDecodeInt32Int32Map(t *testing.T) {
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
	var m map[int32]int32
	expected := map[int32]int32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int32]int32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int32]int32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int32]int32`) // 1
}

func TestDecodeInt32Int64Map(t *testing.T) {
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
	var m map[int32]int64
	expected := map[int32]int64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int32]int64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int32]int64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int32]int64`) // 1
}

func TestDecodeInt32UintMap(t *testing.T) {
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
	var m map[int32]uint
	expected := map[int32]uint{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int32]uint{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int32]uint{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int32]uint`) // 1
}

func TestDecodeInt32Uint8Map(t *testing.T) {
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
	var m map[int32]uint8
	expected := map[int32]uint8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int32]uint8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int32]uint8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int32]uint8`) // 1
}

func TestDecodeInt32Uint16Map(t *testing.T) {
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
	var m map[int32]uint16
	expected := map[int32]uint16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int32]uint16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int32]uint16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int32]uint16`) // 1
}

func TestDecodeInt32Uint32Map(t *testing.T) {
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
	var m map[int32]uint32
	expected := map[int32]uint32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int32]uint32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int32]uint32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int32]uint32`) // 1
}

func TestDecodeInt32Uint64Map(t *testing.T) {
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
	var m map[int32]uint64
	expected := map[int32]uint64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int32]uint64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int32]uint64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int32]uint64`) // 1
}

func TestDecodeInt32Float32Map(t *testing.T) {
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
	var m map[int32]float32
	expected := map[int32]float32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int32]float32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int32]float32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int32]float32`) // 1
}

func TestDecodeInt32Float64Map(t *testing.T) {
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
	var m map[int32]float64
	expected := map[int32]float64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int32]float64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int32]float64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int32]float64`) // 1
}

func TestDecodeInt32BoolMap(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{0, 1, 2, 3, 4})
	enc.Encode([]float32{0, 1, 2, 3, 4})
	enc.Encode(map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var m map[int32]bool
	expected := map[int32]bool{0: false, 1: true, 2: true, 3: true, 4: true}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, map[int32]bool{1: true, 2: true, 3: true, 4: true, 5: true}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int32]bool{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int32]bool`) // 1
}

func TestDecodeInt32StringMap(t *testing.T) {
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
	var m map[int32]string
	expected := map[int32]string{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int32]string{1: "1", 2: "4", 3: "8", 4: "16", 5: "32"}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int32]string{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int32]string`) // 1
}

func TestDecodeInt32InterfaceMap(t *testing.T) {
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
	dec.RealType = RealTypeFloat32
	var m map[int32]interface{}
	dec.Decode(&m)
	assert.Equal(t, map[int32]interface{}{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[int32]interface{}{0: float32(1), 1: float32(2), 2: float32(3), 3: float32(4), 4: float32(5)}, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[int32]interface{}{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int32]interface{}{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int32]interface{}{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int32]interface {}`) // 1
}

func TestDecodeInt32CustomIntMap(t *testing.T) {
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
	type Int int
	var m map[int32]Int
	expected := map[int32]Int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int32]Int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int32]Int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int32]encoding.Int`) // 1
}

func TestDecodeInt64IntMap(t *testing.T) {
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
	var m map[int64]int
	expected := map[int64]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int64]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int64]int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int64]int`) // 1
}

func TestDecodeInt64Int8Map(t *testing.T) {
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
	var m map[int64]int8
	expected := map[int64]int8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int64]int8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int64]int8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int64]int8`) // 1
}

func TestDecodeInt64Int16Map(t *testing.T) {
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
	var m map[int64]int16
	expected := map[int64]int16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int64]int16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int64]int16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int64]int16`) // 1
}

func TestDecodeInt64Int32Map(t *testing.T) {
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
	var m map[int64]int32
	expected := map[int64]int32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int64]int32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int64]int32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int64]int32`) // 1
}

func TestDecodeInt64Int64Map(t *testing.T) {
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
	var m map[int64]int64
	expected := map[int64]int64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int64]int64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int64]int64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int64]int64`) // 1
}

func TestDecodeInt64UintMap(t *testing.T) {
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
	var m map[int64]uint
	expected := map[int64]uint{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int64]uint{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int64]uint{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int64]uint`) // 1
}

func TestDecodeInt64Uint8Map(t *testing.T) {
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
	var m map[int64]uint8
	expected := map[int64]uint8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int64]uint8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int64]uint8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int64]uint8`) // 1
}

func TestDecodeInt64Uint16Map(t *testing.T) {
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
	var m map[int64]uint16
	expected := map[int64]uint16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int64]uint16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int64]uint16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int64]uint16`) // 1
}

func TestDecodeInt64Uint32Map(t *testing.T) {
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
	var m map[int64]uint32
	expected := map[int64]uint32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int64]uint32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int64]uint32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int64]uint32`) // 1
}

func TestDecodeInt64Uint64Map(t *testing.T) {
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
	var m map[int64]uint64
	expected := map[int64]uint64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int64]uint64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int64]uint64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int64]uint64`) // 1
}

func TestDecodeInt64Float32Map(t *testing.T) {
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
	var m map[int64]float32
	expected := map[int64]float32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int64]float32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int64]float32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int64]float32`) // 1
}

func TestDecodeInt64Float64Map(t *testing.T) {
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
	var m map[int64]float64
	expected := map[int64]float64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int64]float64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int64]float64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int64]float64`) // 1
}

func TestDecodeInt64BoolMap(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{0, 1, 2, 3, 4})
	enc.Encode([]float32{0, 1, 2, 3, 4})
	enc.Encode(map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var m map[int64]bool
	expected := map[int64]bool{0: false, 1: true, 2: true, 3: true, 4: true}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, map[int64]bool{1: true, 2: true, 3: true, 4: true, 5: true}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int64]bool{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int64]bool`) // 1
}

func TestDecodeInt64StringMap(t *testing.T) {
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
	var m map[int64]string
	expected := map[int64]string{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int64]string{1: "1", 2: "4", 3: "8", 4: "16", 5: "32"}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int64]string{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int64]string`) // 1
}

func TestDecodeInt64InterfaceMap(t *testing.T) {
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
	dec.RealType = RealTypeFloat32
	var m map[int64]interface{}
	dec.Decode(&m)
	assert.Equal(t, map[int64]interface{}{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[int64]interface{}{0: float32(1), 1: float32(2), 2: float32(3), 3: float32(4), 4: float32(5)}, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[int64]interface{}{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int64]interface{}{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int64]interface{}{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int64]interface {}`) // 1
}

func TestDecodeInt64CustomIntMap(t *testing.T) {
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
	type Int int
	var m map[int64]Int
	expected := map[int64]Int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[int64]Int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[int64]Int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[int64]encoding.Int`) // 1
}

func TestDecodeUintIntMap(t *testing.T) {
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
	var m map[uint]int
	expected := map[uint]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint]int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint]int`) // 1
}

func TestDecodeUintInt8Map(t *testing.T) {
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
	var m map[uint]int8
	expected := map[uint]int8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint]int8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint]int8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint]int8`) // 1
}

func TestDecodeUintInt16Map(t *testing.T) {
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
	var m map[uint]int16
	expected := map[uint]int16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint]int16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint]int16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint]int16`) // 1
}

func TestDecodeUintInt32Map(t *testing.T) {
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
	var m map[uint]int32
	expected := map[uint]int32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint]int32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint]int32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint]int32`) // 1
}

func TestDecodeUintInt64Map(t *testing.T) {
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
	var m map[uint]int64
	expected := map[uint]int64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint]int64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint]int64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint]int64`) // 1
}

func TestDecodeUintUintMap(t *testing.T) {
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
	var m map[uint]uint
	expected := map[uint]uint{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint]uint{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint]uint{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint]uint`) // 1
}

func TestDecodeUintUint8Map(t *testing.T) {
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
	var m map[uint]uint8
	expected := map[uint]uint8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint]uint8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint]uint8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint]uint8`) // 1
}

func TestDecodeUintUint16Map(t *testing.T) {
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
	var m map[uint]uint16
	expected := map[uint]uint16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint]uint16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint]uint16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint]uint16`) // 1
}

func TestDecodeUintUint32Map(t *testing.T) {
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
	var m map[uint]uint32
	expected := map[uint]uint32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint]uint32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint]uint32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint]uint32`) // 1
}

func TestDecodeUintUint64Map(t *testing.T) {
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
	var m map[uint]uint64
	expected := map[uint]uint64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint]uint64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint]uint64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint]uint64`) // 1
}

func TestDecodeUintFloat32Map(t *testing.T) {
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
	var m map[uint]float32
	expected := map[uint]float32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint]float32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint]float32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint]float32`) // 1
}

func TestDecodeUintFloat64Map(t *testing.T) {
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
	var m map[uint]float64
	expected := map[uint]float64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint]float64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint]float64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint]float64`) // 1
}

func TestDecodeUintBoolMap(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{0, 1, 2, 3, 4})
	enc.Encode([]float32{0, 1, 2, 3, 4})
	enc.Encode(map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var m map[uint]bool
	expected := map[uint]bool{0: false, 1: true, 2: true, 3: true, 4: true}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, map[uint]bool{1: true, 2: true, 3: true, 4: true, 5: true}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint]bool{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint]bool`) // 1
}

func TestDecodeUintStringMap(t *testing.T) {
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
	var m map[uint]string
	expected := map[uint]string{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint]string{1: "1", 2: "4", 3: "8", 4: "16", 5: "32"}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint]string{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint]string`) // 1
}

func TestDecodeUintInterfaceMap(t *testing.T) {
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
	dec.RealType = RealTypeFloat32
	var m map[uint]interface{}
	dec.Decode(&m)
	assert.Equal(t, map[uint]interface{}{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[uint]interface{}{0: float32(1), 1: float32(2), 2: float32(3), 3: float32(4), 4: float32(5)}, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[uint]interface{}{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint]interface{}{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint]interface{}{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint]interface {}`) // 1
}

func TestDecodeUintCustomIntMap(t *testing.T) {
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
	type Int int
	var m map[uint]Int
	expected := map[uint]Int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint]Int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint]Int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint]encoding.Int`) // 1
}

func TestDecodeUint8IntMap(t *testing.T) {
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
	var m map[uint8]int
	expected := map[uint8]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint8]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint8]int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint8]int`) // 1
}

func TestDecodeUint8Int8Map(t *testing.T) {
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
	var m map[uint8]int8
	expected := map[uint8]int8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint8]int8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint8]int8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint8]int8`) // 1
}

func TestDecodeUint8Int16Map(t *testing.T) {
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
	var m map[uint8]int16
	expected := map[uint8]int16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint8]int16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint8]int16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint8]int16`) // 1
}

func TestDecodeUint8Int32Map(t *testing.T) {
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
	var m map[uint8]int32
	expected := map[uint8]int32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint8]int32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint8]int32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint8]int32`) // 1
}

func TestDecodeUint8Int64Map(t *testing.T) {
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
	var m map[uint8]int64
	expected := map[uint8]int64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint8]int64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint8]int64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint8]int64`) // 1
}

func TestDecodeUint8UintMap(t *testing.T) {
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
	var m map[uint8]uint
	expected := map[uint8]uint{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint8]uint{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint8]uint{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint8]uint`) // 1
}

func TestDecodeUint8Uint8Map(t *testing.T) {
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
	var m map[uint8]uint8
	expected := map[uint8]uint8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint8]uint8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint8]uint8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint8]uint8`) // 1
}

func TestDecodeUint8Uint16Map(t *testing.T) {
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
	var m map[uint8]uint16
	expected := map[uint8]uint16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint8]uint16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint8]uint16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint8]uint16`) // 1
}

func TestDecodeUint8Uint32Map(t *testing.T) {
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
	var m map[uint8]uint32
	expected := map[uint8]uint32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint8]uint32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint8]uint32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint8]uint32`) // 1
}

func TestDecodeUint8Uint64Map(t *testing.T) {
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
	var m map[uint8]uint64
	expected := map[uint8]uint64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint8]uint64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint8]uint64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint8]uint64`) // 1
}

func TestDecodeUint8Float32Map(t *testing.T) {
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
	var m map[uint8]float32
	expected := map[uint8]float32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint8]float32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint8]float32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint8]float32`) // 1
}

func TestDecodeUint8Float64Map(t *testing.T) {
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
	var m map[uint8]float64
	expected := map[uint8]float64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint8]float64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint8]float64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint8]float64`) // 1
}

func TestDecodeUint8BoolMap(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{0, 1, 2, 3, 4})
	enc.Encode([]float32{0, 1, 2, 3, 4})
	enc.Encode(map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var m map[uint8]bool
	expected := map[uint8]bool{0: false, 1: true, 2: true, 3: true, 4: true}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, map[uint8]bool{1: true, 2: true, 3: true, 4: true, 5: true}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint8]bool{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint8]bool`) // 1
}

func TestDecodeUint8StringMap(t *testing.T) {
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
	var m map[uint8]string
	expected := map[uint8]string{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint8]string{1: "1", 2: "4", 3: "8", 4: "16", 5: "32"}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint8]string{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint8]string`) // 1
}

func TestDecodeUint8InterfaceMap(t *testing.T) {
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
	dec.RealType = RealTypeFloat32
	var m map[uint8]interface{}
	dec.Decode(&m)
	assert.Equal(t, map[uint8]interface{}{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[uint8]interface{}{0: float32(1), 1: float32(2), 2: float32(3), 3: float32(4), 4: float32(5)}, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[uint8]interface{}{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint8]interface{}{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint8]interface{}{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint8]interface {}`) // 1
}

func TestDecodeUint8CustomIntMap(t *testing.T) {
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
	type Int int
	var m map[uint8]Int
	expected := map[uint8]Int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint8]Int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint8]Int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint8]encoding.Int`) // 1
}

func TestDecodeUint16IntMap(t *testing.T) {
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
	var m map[uint16]int
	expected := map[uint16]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint16]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint16]int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint16]int`) // 1
}

func TestDecodeUint16Int8Map(t *testing.T) {
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
	var m map[uint16]int8
	expected := map[uint16]int8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint16]int8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint16]int8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint16]int8`) // 1
}

func TestDecodeUint16Int16Map(t *testing.T) {
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
	var m map[uint16]int16
	expected := map[uint16]int16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint16]int16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint16]int16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint16]int16`) // 1
}

func TestDecodeUint16Int32Map(t *testing.T) {
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
	var m map[uint16]int32
	expected := map[uint16]int32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint16]int32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint16]int32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint16]int32`) // 1
}

func TestDecodeUint16Int64Map(t *testing.T) {
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
	var m map[uint16]int64
	expected := map[uint16]int64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint16]int64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint16]int64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint16]int64`) // 1
}

func TestDecodeUint16UintMap(t *testing.T) {
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
	var m map[uint16]uint
	expected := map[uint16]uint{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint16]uint{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint16]uint{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint16]uint`) // 1
}

func TestDecodeUint16Uint8Map(t *testing.T) {
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
	var m map[uint16]uint8
	expected := map[uint16]uint8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint16]uint8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint16]uint8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint16]uint8`) // 1
}

func TestDecodeUint16Uint16Map(t *testing.T) {
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
	var m map[uint16]uint16
	expected := map[uint16]uint16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint16]uint16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint16]uint16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint16]uint16`) // 1
}

func TestDecodeUint16Uint32Map(t *testing.T) {
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
	var m map[uint16]uint32
	expected := map[uint16]uint32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint16]uint32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint16]uint32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint16]uint32`) // 1
}

func TestDecodeUint16Uint64Map(t *testing.T) {
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
	var m map[uint16]uint64
	expected := map[uint16]uint64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint16]uint64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint16]uint64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint16]uint64`) // 1
}

func TestDecodeUint16Float32Map(t *testing.T) {
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
	var m map[uint16]float32
	expected := map[uint16]float32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint16]float32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint16]float32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint16]float32`) // 1
}

func TestDecodeUint16Float64Map(t *testing.T) {
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
	var m map[uint16]float64
	expected := map[uint16]float64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint16]float64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint16]float64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint16]float64`) // 1
}

func TestDecodeUint16BoolMap(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{0, 1, 2, 3, 4})
	enc.Encode([]float32{0, 1, 2, 3, 4})
	enc.Encode(map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var m map[uint16]bool
	expected := map[uint16]bool{0: false, 1: true, 2: true, 3: true, 4: true}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, map[uint16]bool{1: true, 2: true, 3: true, 4: true, 5: true}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint16]bool{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint16]bool`) // 1
}

func TestDecodeUint16StringMap(t *testing.T) {
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
	var m map[uint16]string
	expected := map[uint16]string{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint16]string{1: "1", 2: "4", 3: "8", 4: "16", 5: "32"}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint16]string{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint16]string`) // 1
}

func TestDecodeUint16InterfaceMap(t *testing.T) {
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
	dec.RealType = RealTypeFloat32
	var m map[uint16]interface{}
	dec.Decode(&m)
	assert.Equal(t, map[uint16]interface{}{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[uint16]interface{}{0: float32(1), 1: float32(2), 2: float32(3), 3: float32(4), 4: float32(5)}, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[uint16]interface{}{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint16]interface{}{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint16]interface{}{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint16]interface {}`) // 1
}

func TestDecodeUint16CustomIntMap(t *testing.T) {
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
	type Int int
	var m map[uint16]Int
	expected := map[uint16]Int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint16]Int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint16]Int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint16]encoding.Int`) // 1
}

func TestDecodeUint32IntMap(t *testing.T) {
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
	var m map[uint32]int
	expected := map[uint32]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint32]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint32]int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint32]int`) // 1
}

func TestDecodeUint32Int8Map(t *testing.T) {
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
	var m map[uint32]int8
	expected := map[uint32]int8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint32]int8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint32]int8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint32]int8`) // 1
}

func TestDecodeUint32Int16Map(t *testing.T) {
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
	var m map[uint32]int16
	expected := map[uint32]int16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint32]int16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint32]int16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint32]int16`) // 1
}

func TestDecodeUint32Int32Map(t *testing.T) {
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
	var m map[uint32]int32
	expected := map[uint32]int32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint32]int32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint32]int32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint32]int32`) // 1
}

func TestDecodeUint32Int64Map(t *testing.T) {
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
	var m map[uint32]int64
	expected := map[uint32]int64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint32]int64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint32]int64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint32]int64`) // 1
}

func TestDecodeUint32UintMap(t *testing.T) {
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
	var m map[uint32]uint
	expected := map[uint32]uint{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint32]uint{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint32]uint{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint32]uint`) // 1
}

func TestDecodeUint32Uint8Map(t *testing.T) {
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
	var m map[uint32]uint8
	expected := map[uint32]uint8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint32]uint8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint32]uint8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint32]uint8`) // 1
}

func TestDecodeUint32Uint16Map(t *testing.T) {
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
	var m map[uint32]uint16
	expected := map[uint32]uint16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint32]uint16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint32]uint16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint32]uint16`) // 1
}

func TestDecodeUint32Uint32Map(t *testing.T) {
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
	var m map[uint32]uint32
	expected := map[uint32]uint32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint32]uint32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint32]uint32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint32]uint32`) // 1
}

func TestDecodeUint32Uint64Map(t *testing.T) {
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
	var m map[uint32]uint64
	expected := map[uint32]uint64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint32]uint64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint32]uint64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint32]uint64`) // 1
}

func TestDecodeUint32Float32Map(t *testing.T) {
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
	var m map[uint32]float32
	expected := map[uint32]float32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint32]float32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint32]float32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint32]float32`) // 1
}

func TestDecodeUint32Float64Map(t *testing.T) {
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
	var m map[uint32]float64
	expected := map[uint32]float64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint32]float64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint32]float64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint32]float64`) // 1
}

func TestDecodeUint32BoolMap(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{0, 1, 2, 3, 4})
	enc.Encode([]float32{0, 1, 2, 3, 4})
	enc.Encode(map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var m map[uint32]bool
	expected := map[uint32]bool{0: false, 1: true, 2: true, 3: true, 4: true}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, map[uint32]bool{1: true, 2: true, 3: true, 4: true, 5: true}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint32]bool{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint32]bool`) // 1
}

func TestDecodeUint32StringMap(t *testing.T) {
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
	var m map[uint32]string
	expected := map[uint32]string{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint32]string{1: "1", 2: "4", 3: "8", 4: "16", 5: "32"}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint32]string{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint32]string`) // 1
}

func TestDecodeUint32InterfaceMap(t *testing.T) {
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
	dec.RealType = RealTypeFloat32
	var m map[uint32]interface{}
	dec.Decode(&m)
	assert.Equal(t, map[uint32]interface{}{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[uint32]interface{}{0: float32(1), 1: float32(2), 2: float32(3), 3: float32(4), 4: float32(5)}, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[uint32]interface{}{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint32]interface{}{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint32]interface{}{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint32]interface {}`) // 1
}

func TestDecodeUint32CustomIntMap(t *testing.T) {
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
	type Int int
	var m map[uint32]Int
	expected := map[uint32]Int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint32]Int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint32]Int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint32]encoding.Int`) // 1
}

func TestDecodeUint64IntMap(t *testing.T) {
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
	var m map[uint64]int
	expected := map[uint64]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint64]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint64]int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint64]int`) // 1
}

func TestDecodeUint64Int8Map(t *testing.T) {
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
	var m map[uint64]int8
	expected := map[uint64]int8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint64]int8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint64]int8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint64]int8`) // 1
}

func TestDecodeUint64Int16Map(t *testing.T) {
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
	var m map[uint64]int16
	expected := map[uint64]int16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint64]int16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint64]int16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint64]int16`) // 1
}

func TestDecodeUint64Int32Map(t *testing.T) {
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
	var m map[uint64]int32
	expected := map[uint64]int32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint64]int32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint64]int32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint64]int32`) // 1
}

func TestDecodeUint64Int64Map(t *testing.T) {
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
	var m map[uint64]int64
	expected := map[uint64]int64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint64]int64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint64]int64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint64]int64`) // 1
}

func TestDecodeUint64UintMap(t *testing.T) {
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
	var m map[uint64]uint
	expected := map[uint64]uint{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint64]uint{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint64]uint{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint64]uint`) // 1
}

func TestDecodeUint64Uint8Map(t *testing.T) {
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
	var m map[uint64]uint8
	expected := map[uint64]uint8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint64]uint8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint64]uint8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint64]uint8`) // 1
}

func TestDecodeUint64Uint16Map(t *testing.T) {
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
	var m map[uint64]uint16
	expected := map[uint64]uint16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint64]uint16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint64]uint16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint64]uint16`) // 1
}

func TestDecodeUint64Uint32Map(t *testing.T) {
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
	var m map[uint64]uint32
	expected := map[uint64]uint32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint64]uint32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint64]uint32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint64]uint32`) // 1
}

func TestDecodeUint64Uint64Map(t *testing.T) {
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
	var m map[uint64]uint64
	expected := map[uint64]uint64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint64]uint64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint64]uint64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint64]uint64`) // 1
}

func TestDecodeUint64Float32Map(t *testing.T) {
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
	var m map[uint64]float32
	expected := map[uint64]float32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint64]float32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint64]float32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint64]float32`) // 1
}

func TestDecodeUint64Float64Map(t *testing.T) {
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
	var m map[uint64]float64
	expected := map[uint64]float64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint64]float64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint64]float64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint64]float64`) // 1
}

func TestDecodeUint64BoolMap(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{0, 1, 2, 3, 4})
	enc.Encode([]float32{0, 1, 2, 3, 4})
	enc.Encode(map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var m map[uint64]bool
	expected := map[uint64]bool{0: false, 1: true, 2: true, 3: true, 4: true}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, map[uint64]bool{1: true, 2: true, 3: true, 4: true, 5: true}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint64]bool{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint64]bool`) // 1
}

func TestDecodeUint64StringMap(t *testing.T) {
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
	var m map[uint64]string
	expected := map[uint64]string{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint64]string{1: "1", 2: "4", 3: "8", 4: "16", 5: "32"}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint64]string{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint64]string`) // 1
}

func TestDecodeUint64InterfaceMap(t *testing.T) {
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
	dec.RealType = RealTypeFloat32
	var m map[uint64]interface{}
	dec.Decode(&m)
	assert.Equal(t, map[uint64]interface{}{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[uint64]interface{}{0: float32(1), 1: float32(2), 2: float32(3), 3: float32(4), 4: float32(5)}, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[uint64]interface{}{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint64]interface{}{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint64]interface{}{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint64]interface {}`) // 1
}

func TestDecodeUint64CustomIntMap(t *testing.T) {
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
	type Int int
	var m map[uint64]Int
	expected := map[uint64]Int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uint64]Int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uint64]Int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uint64]encoding.Int`) // 1
}

func TestDecodeFloat32IntMap(t *testing.T) {
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
	var m map[float32]int
	expected := map[float32]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float32]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float32]int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float32]int`) // 1
}

func TestDecodeFloat32Int8Map(t *testing.T) {
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
	var m map[float32]int8
	expected := map[float32]int8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float32]int8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float32]int8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float32]int8`) // 1
}

func TestDecodeFloat32Int16Map(t *testing.T) {
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
	var m map[float32]int16
	expected := map[float32]int16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float32]int16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float32]int16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float32]int16`) // 1
}

func TestDecodeFloat32Int32Map(t *testing.T) {
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
	var m map[float32]int32
	expected := map[float32]int32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float32]int32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float32]int32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float32]int32`) // 1
}

func TestDecodeFloat32Int64Map(t *testing.T) {
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
	var m map[float32]int64
	expected := map[float32]int64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float32]int64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float32]int64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float32]int64`) // 1
}

func TestDecodeFloat32UintMap(t *testing.T) {
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
	var m map[float32]uint
	expected := map[float32]uint{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float32]uint{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float32]uint{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float32]uint`) // 1
}

func TestDecodeFloat32Uint8Map(t *testing.T) {
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
	var m map[float32]uint8
	expected := map[float32]uint8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float32]uint8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float32]uint8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float32]uint8`) // 1
}

func TestDecodeFloat32Uint16Map(t *testing.T) {
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
	var m map[float32]uint16
	expected := map[float32]uint16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float32]uint16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float32]uint16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float32]uint16`) // 1
}

func TestDecodeFloat32Uint32Map(t *testing.T) {
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
	var m map[float32]uint32
	expected := map[float32]uint32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float32]uint32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float32]uint32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float32]uint32`) // 1
}

func TestDecodeFloat32Uint64Map(t *testing.T) {
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
	var m map[float32]uint64
	expected := map[float32]uint64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float32]uint64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float32]uint64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float32]uint64`) // 1
}

func TestDecodeFloat32Float32Map(t *testing.T) {
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
	var m map[float32]float32
	expected := map[float32]float32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float32]float32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float32]float32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float32]float32`) // 1
}

func TestDecodeFloat32Float64Map(t *testing.T) {
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
	var m map[float32]float64
	expected := map[float32]float64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float32]float64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float32]float64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float32]float64`) // 1
}

func TestDecodeFloat32BoolMap(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{0, 1, 2, 3, 4})
	enc.Encode([]float32{0, 1, 2, 3, 4})
	enc.Encode(map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var m map[float32]bool
	expected := map[float32]bool{0: false, 1: true, 2: true, 3: true, 4: true}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, map[float32]bool{1: true, 2: true, 3: true, 4: true, 5: true}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float32]bool{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float32]bool`) // 1
}

func TestDecodeFloat32StringMap(t *testing.T) {
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
	var m map[float32]string
	expected := map[float32]string{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float32]string{1: "1", 2: "4", 3: "8", 4: "16", 5: "32"}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float32]string{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float32]string`) // 1
}

func TestDecodeFloat32InterfaceMap(t *testing.T) {
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
	dec.RealType = RealTypeFloat32
	var m map[float32]interface{}
	dec.Decode(&m)
	assert.Equal(t, map[float32]interface{}{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[float32]interface{}{0: float32(1), 1: float32(2), 2: float32(3), 3: float32(4), 4: float32(5)}, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[float32]interface{}{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float32]interface{}{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float32]interface{}{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float32]interface {}`) // 1
}

func TestDecodeFloat32CustomIntMap(t *testing.T) {
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
	type Int int
	var m map[float32]Int
	expected := map[float32]Int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float32]Int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float32]Int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float32]encoding.Int`) // 1
}

func TestDecodeFloat64IntMap(t *testing.T) {
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
	var m map[float64]int
	expected := map[float64]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float64]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float64]int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float64]int`) // 1
}

func TestDecodeFloat64Int8Map(t *testing.T) {
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
	var m map[float64]int8
	expected := map[float64]int8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float64]int8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float64]int8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float64]int8`) // 1
}

func TestDecodeFloat64Int16Map(t *testing.T) {
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
	var m map[float64]int16
	expected := map[float64]int16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float64]int16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float64]int16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float64]int16`) // 1
}

func TestDecodeFloat64Int32Map(t *testing.T) {
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
	var m map[float64]int32
	expected := map[float64]int32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float64]int32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float64]int32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float64]int32`) // 1
}

func TestDecodeFloat64Int64Map(t *testing.T) {
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
	var m map[float64]int64
	expected := map[float64]int64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float64]int64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float64]int64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float64]int64`) // 1
}

func TestDecodeFloat64UintMap(t *testing.T) {
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
	var m map[float64]uint
	expected := map[float64]uint{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float64]uint{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float64]uint{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float64]uint`) // 1
}

func TestDecodeFloat64Uint8Map(t *testing.T) {
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
	var m map[float64]uint8
	expected := map[float64]uint8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float64]uint8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float64]uint8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float64]uint8`) // 1
}

func TestDecodeFloat64Uint16Map(t *testing.T) {
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
	var m map[float64]uint16
	expected := map[float64]uint16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float64]uint16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float64]uint16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float64]uint16`) // 1
}

func TestDecodeFloat64Uint32Map(t *testing.T) {
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
	var m map[float64]uint32
	expected := map[float64]uint32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float64]uint32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float64]uint32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float64]uint32`) // 1
}

func TestDecodeFloat64Uint64Map(t *testing.T) {
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
	var m map[float64]uint64
	expected := map[float64]uint64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float64]uint64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float64]uint64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float64]uint64`) // 1
}

func TestDecodeFloat64Float32Map(t *testing.T) {
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
	var m map[float64]float32
	expected := map[float64]float32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float64]float32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float64]float32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float64]float32`) // 1
}

func TestDecodeFloat64Float64Map(t *testing.T) {
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
	var m map[float64]float64
	expected := map[float64]float64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float64]float64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float64]float64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float64]float64`) // 1
}

func TestDecodeFloat64BoolMap(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{0, 1, 2, 3, 4})
	enc.Encode([]float32{0, 1, 2, 3, 4})
	enc.Encode(map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var m map[float64]bool
	expected := map[float64]bool{0: false, 1: true, 2: true, 3: true, 4: true}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, map[float64]bool{1: true, 2: true, 3: true, 4: true, 5: true}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float64]bool{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float64]bool`) // 1
}

func TestDecodeFloat64StringMap(t *testing.T) {
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
	var m map[float64]string
	expected := map[float64]string{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float64]string{1: "1", 2: "4", 3: "8", 4: "16", 5: "32"}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float64]string{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float64]string`) // 1
}

func TestDecodeFloat64InterfaceMap(t *testing.T) {
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
	dec.RealType = RealTypeFloat32
	var m map[float64]interface{}
	dec.Decode(&m)
	assert.Equal(t, map[float64]interface{}{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[float64]interface{}{0: float32(1), 1: float32(2), 2: float32(3), 3: float32(4), 4: float32(5)}, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[float64]interface{}{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float64]interface{}{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float64]interface{}{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float64]interface {}`) // 1
}

func TestDecodeFloat64CustomIntMap(t *testing.T) {
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
	type Int int
	var m map[float64]Int
	expected := map[float64]Int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[float64]Int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[float64]Int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[float64]encoding.Int`) // 1
}

func TestDecodeStringIntMap(t *testing.T) {
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
	var m map[string]int
	expected := map[string]int{"0": 1, "1": 2, "2": 3, "3": 4, "4": 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[string]int{"1": 1, "2": 4, "3": 8, "4": 16, "5": 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[string]int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[string]int`) // 1
}

func TestDecodeStringInt8Map(t *testing.T) {
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
	var m map[string]int8
	expected := map[string]int8{"0": 1, "1": 2, "2": 3, "3": 4, "4": 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[string]int8{"1": 1, "2": 4, "3": 8, "4": 16, "5": 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[string]int8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[string]int8`) // 1
}

func TestDecodeStringInt16Map(t *testing.T) {
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
	var m map[string]int16
	expected := map[string]int16{"0": 1, "1": 2, "2": 3, "3": 4, "4": 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[string]int16{"1": 1, "2": 4, "3": 8, "4": 16, "5": 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[string]int16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[string]int16`) // 1
}

func TestDecodeStringInt32Map(t *testing.T) {
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
	var m map[string]int32
	expected := map[string]int32{"0": 1, "1": 2, "2": 3, "3": 4, "4": 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[string]int32{"1": 1, "2": 4, "3": 8, "4": 16, "5": 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[string]int32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[string]int32`) // 1
}

func TestDecodeStringInt64Map(t *testing.T) {
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
	var m map[string]int64
	expected := map[string]int64{"0": 1, "1": 2, "2": 3, "3": 4, "4": 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[string]int64{"1": 1, "2": 4, "3": 8, "4": 16, "5": 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[string]int64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[string]int64`) // 1
}

func TestDecodeStringUintMap(t *testing.T) {
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
	var m map[string]uint
	expected := map[string]uint{"0": 1, "1": 2, "2": 3, "3": 4, "4": 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[string]uint{"1": 1, "2": 4, "3": 8, "4": 16, "5": 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[string]uint{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[string]uint`) // 1
}

func TestDecodeStringUint8Map(t *testing.T) {
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
	var m map[string]uint8
	expected := map[string]uint8{"0": 1, "1": 2, "2": 3, "3": 4, "4": 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[string]uint8{"1": 1, "2": 4, "3": 8, "4": 16, "5": 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[string]uint8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[string]uint8`) // 1
}

func TestDecodeStringUint16Map(t *testing.T) {
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
	var m map[string]uint16
	expected := map[string]uint16{"0": 1, "1": 2, "2": 3, "3": 4, "4": 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[string]uint16{"1": 1, "2": 4, "3": 8, "4": 16, "5": 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[string]uint16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[string]uint16`) // 1
}

func TestDecodeStringUint32Map(t *testing.T) {
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
	var m map[string]uint32
	expected := map[string]uint32{"0": 1, "1": 2, "2": 3, "3": 4, "4": 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[string]uint32{"1": 1, "2": 4, "3": 8, "4": 16, "5": 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[string]uint32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[string]uint32`) // 1
}

func TestDecodeStringUint64Map(t *testing.T) {
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
	var m map[string]uint64
	expected := map[string]uint64{"0": 1, "1": 2, "2": 3, "3": 4, "4": 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[string]uint64{"1": 1, "2": 4, "3": 8, "4": 16, "5": 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[string]uint64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[string]uint64`) // 1
}

func TestDecodeStringFloat32Map(t *testing.T) {
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
	var m map[string]float32
	expected := map[string]float32{"0": 1, "1": 2, "2": 3, "3": 4, "4": 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[string]float32{"1": 1, "2": 4, "3": 8, "4": 16, "5": 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[string]float32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[string]float32`) // 1
}

func TestDecodeStringFloat64Map(t *testing.T) {
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
	var m map[string]float64
	expected := map[string]float64{"0": 1, "1": 2, "2": 3, "3": 4, "4": 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[string]float64{"1": 1, "2": 4, "3": 8, "4": 16, "5": 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[string]float64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[string]float64`) // 1
}

func TestDecodeStringBoolMap(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{0, 1, 2, 3, 4})
	enc.Encode([]float32{0, 1, 2, 3, 4})
	enc.Encode(map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var m map[string]bool
	expected := map[string]bool{"0": false, "1": true, "2": true, "3": true, "4": true}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, map[string]bool{"1": true, "2": true, "3": true, "4": true, "5": true}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[string]bool{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[string]bool`) // 1
}

func TestDecodeStringStringMap(t *testing.T) {
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
	var m map[string]string
	expected := map[string]string{"0": "1", "1": "2", "2": "3", "3": "4", "4": "5"}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[string]string{"1": "1", "2": "4", "3": "8", "4": "16", "5": "32"}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[string]string{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[string]string`) // 1
}

func TestDecodeStringInterfaceMap(t *testing.T) {
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
	dec.RealType = RealTypeFloat32
	var m map[string]interface{}
	dec.Decode(&m)
	assert.Equal(t, map[string]interface{}{"0": 1, "1": 2, "2": 3, "3": 4, "4": 5}, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[string]interface{}{"0": float32(1), "1": float32(2), "2": float32(3), "3": float32(4), "4": float32(5)}, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[string]interface{}{"0": "1", "1": "2", "2": "3", "3": "4", "4": "5"}, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[string]interface{}{"1": 1, "2": 4, "3": 8, "4": 16, "5": 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[string]interface{}{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[string]interface {}`) // 1
}

func TestDecodeStringCustomIntMap(t *testing.T) {
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
	type Int int
	var m map[string]Int
	expected := map[string]Int{"0": 1, "1": 2, "2": 3, "3": 4, "4": 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[string]Int{"1": 1, "2": 4, "3": 8, "4": 16, "5": 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[string]Int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[string]encoding.Int`) // 1
}

func TestDecodeInterfaceIntMap(t *testing.T) {
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
	var m map[interface{}]int
	expected := map[interface{}]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[interface {}]int`) // 1
}

func TestDecodeInterfaceInt8Map(t *testing.T) {
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
	var m map[interface{}]int8
	expected := map[interface{}]int8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]int8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]int8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[interface {}]int8`) // 1
}

func TestDecodeInterfaceInt16Map(t *testing.T) {
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
	var m map[interface{}]int16
	expected := map[interface{}]int16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]int16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]int16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[interface {}]int16`) // 1
}

func TestDecodeInterfaceInt32Map(t *testing.T) {
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
	var m map[interface{}]int32
	expected := map[interface{}]int32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]int32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]int32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[interface {}]int32`) // 1
}

func TestDecodeInterfaceInt64Map(t *testing.T) {
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
	var m map[interface{}]int64
	expected := map[interface{}]int64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]int64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]int64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[interface {}]int64`) // 1
}

func TestDecodeInterfaceUintMap(t *testing.T) {
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
	var m map[interface{}]uint
	expected := map[interface{}]uint{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]uint{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]uint{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[interface {}]uint`) // 1
}

func TestDecodeInterfaceUint8Map(t *testing.T) {
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
	var m map[interface{}]uint8
	expected := map[interface{}]uint8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]uint8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]uint8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[interface {}]uint8`) // 1
}

func TestDecodeInterfaceUint16Map(t *testing.T) {
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
	var m map[interface{}]uint16
	expected := map[interface{}]uint16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]uint16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]uint16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[interface {}]uint16`) // 1
}

func TestDecodeInterfaceUint32Map(t *testing.T) {
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
	var m map[interface{}]uint32
	expected := map[interface{}]uint32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]uint32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]uint32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[interface {}]uint32`) // 1
}

func TestDecodeInterfaceUint64Map(t *testing.T) {
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
	var m map[interface{}]uint64
	expected := map[interface{}]uint64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]uint64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]uint64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[interface {}]uint64`) // 1
}

func TestDecodeInterfaceFloat32Map(t *testing.T) {
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
	var m map[interface{}]float32
	expected := map[interface{}]float32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]float32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]float32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[interface {}]float32`) // 1
}

func TestDecodeInterfaceFloat64Map(t *testing.T) {
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
	var m map[interface{}]float64
	expected := map[interface{}]float64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]float64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]float64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[interface {}]float64`) // 1
}

func TestDecodeInterfaceBoolMap(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{0, 1, 2, 3, 4})
	enc.Encode([]float32{0, 1, 2, 3, 4})
	enc.Encode(map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var m map[interface{}]bool
	expected := map[interface{}]bool{0: false, 1: true, 2: true, 3: true, 4: true}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]bool{1: true, 2: true, 3: true, 4: true, 5: true}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]bool{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[interface {}]bool`) // 1
}

func TestDecodeInterfaceStringMap(t *testing.T) {
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
	var m map[interface{}]string
	expected := map[interface{}]string{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]string{1: "1", 2: "4", 3: "8", 4: "16", 5: "32"}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]string{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[interface {}]string`) // 1
}

func TestDecodeInterfaceInterfaceMap(t *testing.T) {
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
	dec.RealType = RealTypeFloat32
	var m map[interface{}]interface{}
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]interface{}{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]interface{}{0: float32(1), 1: float32(2), 2: float32(3), 3: float32(4), 4: float32(5)}, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]interface{}{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]interface{}{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]interface{}{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[interface {}]interface {}`) // 1
}

func TestDecodeInterfaceCustomIntMap(t *testing.T) {
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
	type Int int
	var m map[interface{}]Int
	expected := map[interface{}]Int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]Int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[interface{}]Int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[interface {}]encoding.Int`) // 1
}

func TestDecodeUintptrIntMap(t *testing.T) {
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
	var m map[uintptr]int
	expected := map[uintptr]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uintptr]int`) // 1
}

func TestDecodeUintptrInt8Map(t *testing.T) {
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
	var m map[uintptr]int8
	expected := map[uintptr]int8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]int8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]int8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uintptr]int8`) // 1
}

func TestDecodeUintptrInt16Map(t *testing.T) {
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
	var m map[uintptr]int16
	expected := map[uintptr]int16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]int16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]int16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uintptr]int16`) // 1
}

func TestDecodeUintptrInt32Map(t *testing.T) {
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
	var m map[uintptr]int32
	expected := map[uintptr]int32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]int32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]int32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uintptr]int32`) // 1
}

func TestDecodeUintptrInt64Map(t *testing.T) {
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
	var m map[uintptr]int64
	expected := map[uintptr]int64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]int64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]int64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uintptr]int64`) // 1
}

func TestDecodeUintptrUintMap(t *testing.T) {
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
	var m map[uintptr]uint
	expected := map[uintptr]uint{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]uint{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]uint{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uintptr]uint`) // 1
}

func TestDecodeUintptrUint8Map(t *testing.T) {
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
	var m map[uintptr]uint8
	expected := map[uintptr]uint8{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]uint8{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]uint8{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uintptr]uint8`) // 1
}

func TestDecodeUintptrUint16Map(t *testing.T) {
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
	var m map[uintptr]uint16
	expected := map[uintptr]uint16{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]uint16{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]uint16{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uintptr]uint16`) // 1
}

func TestDecodeUintptrUint32Map(t *testing.T) {
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
	var m map[uintptr]uint32
	expected := map[uintptr]uint32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]uint32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]uint32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uintptr]uint32`) // 1
}

func TestDecodeUintptrUint64Map(t *testing.T) {
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
	var m map[uintptr]uint64
	expected := map[uintptr]uint64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]uint64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]uint64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uintptr]uint64`) // 1
}

func TestDecodeUintptrFloat32Map(t *testing.T) {
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
	var m map[uintptr]float32
	expected := map[uintptr]float32{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]float32{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]float32{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uintptr]float32`) // 1
}

func TestDecodeUintptrFloat64Map(t *testing.T) {
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
	var m map[uintptr]float64
	expected := map[uintptr]float64{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]float64{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]float64{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uintptr]float64`) // 1
}

func TestDecodeUintptrBoolMap(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{0, 1, 2, 3, 4})
	enc.Encode([]float32{0, 1, 2, 3, 4})
	enc.Encode(map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32})
	enc.Encode(nil)
	enc.Encode("")
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var m map[uintptr]bool
	expected := map[uintptr]bool{0: false, 1: true, 2: true, 3: true, 4: true}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{0, 1, 2, 3, 4}
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]bool{1: true, 2: true, 3: true, 4: true, 5: true}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]bool{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uintptr]bool`) // 1
}

func TestDecodeUintptrStringMap(t *testing.T) {
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
	var m map[uintptr]string
	expected := map[uintptr]string{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]string{1: "1", 2: "4", 3: "8", 4: "16", 5: "32"}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]string{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uintptr]string`) // 1
}

func TestDecodeUintptrInterfaceMap(t *testing.T) {
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
	dec.RealType = RealTypeFloat32
	var m map[uintptr]interface{}
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]interface{}{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]interface{}{0: float32(1), 1: float32(2), 2: float32(3), 3: float32(4), 4: float32(5)}, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]interface{}{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]interface{}{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]interface{}{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uintptr]interface {}`) // 1
}

func TestDecodeUintptrCustomIntMap(t *testing.T) {
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
	type Int int
	var m map[uintptr]Int
	expected := map[uintptr]Int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]Int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[uintptr]Int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[uintptr]encoding.Int`) // 1
}

func TestDecodeComplex64CustomIntMap(t *testing.T) {
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
	type Int int
	var m map[complex64]Int
	expected := map[complex64]Int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[complex64]Int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[complex64]Int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[complex64]encoding.Int`) // 1
}

func TestDecodeComplex128CustomIntMap(t *testing.T) {
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
	type Int int
	var m map[complex128]Int
	expected := map[complex128]Int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[complex128]Int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[complex128]Int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[complex128]encoding.Int`) // 1
}

func TestDecodeCustomIntCustomIntMap(t *testing.T) {
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
	type Int int
	var m map[Int]Int
	expected := map[Int]Int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []int{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []float32{1, 2, 3, 4, 5}
	dec.Decode(&m)
	assert.Equal(t, expected, m) // []string{"1", "2", "3", "4", "5"}
	dec.Decode(&m)
	assert.Equal(t, map[Int]Int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}, m) // map[int]int{1: 1, 2: 4, 3: 8, 4: 16, 5: 32}
	dec.Decode(&m)
	assert.Nil(t, m) // nil
	dec.Decode(&m)
	assert.Equal(t, map[Int]Int{}, m) // ""
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast int to map[encoding.Int]encoding.Int`) // 1
}

func TestDecodeMapError(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{1, 2, 3, 4, 5})
	enc.Encode(map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5})
	dec := NewDecoder(([]byte)(sb.String()))
	var m map[*int]int
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast []interface {} to map[*int]int`)
	dec.Error = nil
	var slice []int
	dec.Decode(&slice)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast map[interface {}]interface {} to []int`)
}

func TestHproseDecodeObjectAsMap(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb)
	type TestStruct1 struct {
		Name     string
		Age      int
		Birthday time.Time
		Male     bool
	}
	birthday := time.Date(2002, 1, 2, 3, 4, 5, 6, time.Local)
	ts := &TestStruct1{
		Name:     "Tom",
		Age:      18,
		Birthday: birthday,
		Male:     true,
	}
	enc.Encode(ts)
	dec := NewDecoder(([]byte)(sb.String()))
	var m map[string]interface{}
	dec.Decode(&m)
	assert.Equal(t, map[string]interface{}{"name": "Tom", "age": 18, "birthday": birthday, "male": true}, m)
}

func TestHproseDecodeObjectAsMap2(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb)
	type TestStruct2 struct {
		Name     string
		Age      int
		Birthday *time.Time
		Male     bool
	}
	Register((*TestStruct2)(nil), "TestStruct2")
	birthday := time.Date(2002, 1, 2, 3, 4, 5, 6, time.Local)
	ts := &TestStruct2{
		Name:     "Tom",
		Age:      18,
		Birthday: &birthday,
		Male:     true,
	}
	enc.Encode(ts)
	dec := NewDecoder(([]byte)(sb.String()))
	var m map[string]interface{}
	dec.Decode(&m)
	assert.Equal(t, map[string]interface{}{"name": "Tom", "age": 18, "birthday": &birthday, "male": true}, m)
}

func TestHproseDecodeObjectAsMapError(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb)
	type TestStruct3 struct {
		Name     string
		Age      int
		Birthday time.Time
		Male     bool
	}
	birthday := time.Date(2002, 1, 2, 3, 4, 5, 6, time.Local)
	ts := &TestStruct3{
		Name:     "Tom",
		Age:      18,
		Birthday: birthday,
		Male:     true,
	}
	enc.Encode(ts)
	dec := NewDecoder(([]byte)(sb.String()))
	var m map[string]string
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast map[string]interface {} to map[string]string`)
}

func TestHproseDecodeObjectAsMapError2(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb)
	type TestStruct4 struct {
		Name     string
		Age      int
		Birthday *time.Time
		Male     bool
	}
	Register((*TestStruct4)(nil), "TestStruct4")
	birthday := time.Date(2002, 1, 2, 3, 4, 5, 6, time.Local)
	ts := &TestStruct4{
		Name:     "Tom",
		Age:      18,
		Birthday: &birthday,
		Male:     true,
	}
	enc.Encode(ts)
	dec := NewDecoder(([]byte)(sb.String()))
	var m map[string]string
	dec.Decode(&m)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not cast *encoding.TestStruct4 to map[string]string`)
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
