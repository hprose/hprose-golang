/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/interface_decoder_test.go                             |
|                                                          |
| LastModified: Feb 7, 2024                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io_test

import (
	"strings"
	"testing"

	. "github.com/hprose/hprose-golang/v3/io"
	"github.com/stretchr/testify/assert"
)

func TestDecodeIntSliceToInterface(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int{0, 1, 2, 3})

	dec := NewDecoder(([]byte)(sb.String()))
	var v interface{}
	dec.Decode(&v)
	assert.Equal(t, []interface{}{0, 1, 2, 3}, v)

	dec = NewDecoder(([]byte)(sb.String()))
	dec.ListType = ListTypeSlice
	dec.Decode(&v)
	assert.Equal(t, []int{0, 1, 2, 3}, v)
}

func TestDecodeInt64SliceToInterface(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]int64{0, 1, 2, 3})

	dec := NewDecoder(([]byte)(sb.String()))
	var v interface{}
	dec.Decode(&v)
	assert.Equal(t, []interface{}{0, 1, 2, 3}, v)

	dec = NewDecoder(([]byte)(sb.String()))
	dec.ListType = ListTypeSlice
	dec.Decode(&v)
	assert.Equal(t, []int{0, 1, 2, 3}, v)

	dec = NewDecoder(([]byte)(sb.String()))
	dec.LongType = LongTypeInt64
	dec.ListType = ListTypeSlice
	dec.Decode(&v)
	assert.Equal(t, []int64{0, 1, 2, 3}, v)
}

func TestDecodeFloat64SliceToInterface(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]float64{0, 1, 2, 3})

	dec := NewDecoder(([]byte)(sb.String()))
	var v interface{}
	dec.Decode(&v)
	assert.Equal(t, []interface{}{float64(0), float64(1), float64(2), float64(3)}, v)

	dec = NewDecoder(([]byte)(sb.String()))
	dec.ListType = ListTypeSlice
	dec.Decode(&v)
	assert.Equal(t, []float64{0, 1, 2, 3}, v)

	dec = NewDecoder(([]byte)(sb.String()))
	dec.RealType = RealTypeFloat32
	dec.ListType = ListTypeSlice
	dec.Decode(&v)
	assert.Equal(t, []float32{0, 1, 2, 3}, v)
}

func TestDecodeStringSliceToInterface(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]string{"", "1", "2", "3"})

	dec := NewDecoder(([]byte)(sb.String()))
	var v interface{}
	dec.Decode(&v)
	assert.Equal(t, []interface{}{"", "1", "2", "3"}, v)

	dec = NewDecoder(([]byte)(sb.String()))
	dec.ListType = ListTypeSlice
	dec.Decode(&v)
	assert.Equal(t, []string{"", "1", "2", "3"}, v)
}

func TestDecodeStructSliceToInterface(t *testing.T) {
	type TestStruct struct {
		A int
		B bool
		C string
		D float32
	}
	Register((*TestStruct)(nil))
	data := []TestStruct{
		{1, true, "1", 1},
		{2, true, "2", 2},
		{3, true, "3", 3},
	}
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode(data)

	expectedData1 := []interface{}{
		&TestStruct{1, true, "1", 1},
		&TestStruct{2, true, "2", 2},
		&TestStruct{3, true, "3", 3},
	}
	dec := NewDecoder(([]byte)(sb.String()))
	var v interface{}
	dec.Decode(&v)
	assert.Equal(t, expectedData1, v)

	expectedData2 := []interface{}{
		TestStruct{1, true, "1", 1},
		TestStruct{2, true, "2", 2},
		TestStruct{3, true, "3", 3},
	}

	dec = NewDecoder(([]byte)(sb.String()))
	dec.StructType = StructTypeStructObject
	dec.Decode(&v)
	assert.Equal(t, expectedData2, v)

	expectedData3 := []*TestStruct{
		{1, true, "1", 1},
		{2, true, "2", 2},
		{3, true, "3", 3},
	}
	dec = NewDecoder(([]byte)(sb.String()))
	dec.ListType = ListTypeSlice
	dec.Decode(&v)
	assert.Equal(t, expectedData3, v)

	dec = NewDecoder(([]byte)(sb.String()))
	dec.StructType = StructTypeStructObject
	dec.ListType = ListTypeSlice
	dec.Decode(&v)
	assert.Equal(t, data, v)

}

func TestDecodeBytesSliceToInterface(t *testing.T) {
	data := [][]byte{
		{1, 2, 3},
		{4, 5, 6},
		nil,
		{7, 8, 9},
	}
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode(data)
	dec := NewDecoder(([]byte)(sb.String()))
	var v interface{}
	dec.Decode(&v)
	assert.Equal(t, []interface{}{
		[]byte{1, 2, 3},
		[]byte{4, 5, 6},
		nil,
		[]byte{7, 8, 9},
	}, v)
	dec = NewDecoder(([]byte)(sb.String()))
	dec.ListType = ListTypeSlice
	dec.Decode(&v)
	assert.Equal(t, data, v)
}

func TestDecodeIntSliceSliceToInterface(t *testing.T) {
	data := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		nil,
		{7, 8, 9},
	}
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode(data)

	dec := NewDecoder(([]byte)(sb.String()))
	var v interface{}
	dec.Decode(&v)
	assert.Equal(t, []interface{}{
		[]interface{}{1, 2, 3},
		[]interface{}{4, 5, 6},
		[]interface{}(nil),
		[]interface{}{7, 8, 9},
	}, v)

	dec = NewDecoder(([]byte)(sb.String()))
	dec.ListType = ListTypeSlice
	dec.Decode(&v)
	assert.Equal(t, data, v)
}

func TestDecodeMapSliceToInterface(t *testing.T) {
	data := []map[string]interface{}{
		{"1": "1", "2": "2", "3": "3"},
		{"4": "4", "5": "5", "6": "6"},
		nil,
		{"7": "7", "8": "8", "9": "9"},
	}
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode(data)

	dec := NewDecoder(([]byte)(sb.String()))
	var v interface{}
	dec.Decode(&v)
	assert.Equal(t, []interface{}{
		map[interface{}]interface{}{"1": "1", "2": "2", "3": "3"},
		map[interface{}]interface{}{"4": "4", "5": "5", "6": "6"},
		nil,
		map[interface{}]interface{}{"7": "7", "8": "8", "9": "9"},
	}, v)

	dec = NewDecoder(([]byte)(sb.String()))
	dec.MapType = MapTypeSIMap
	dec.Decode(&v)
	assert.Equal(t, []interface{}{
		map[string]interface{}{"1": "1", "2": "2", "3": "3"},
		map[string]interface{}{"4": "4", "5": "5", "6": "6"},
		nil,
		map[string]interface{}{"7": "7", "8": "8", "9": "9"},
	}, v)

	dec = NewDecoder(([]byte)(sb.String()))
	dec.ListType = ListTypeSlice
	dec.Decode(&v)
	assert.Equal(t, []map[interface{}]interface{}{
		{"1": "1", "2": "2", "3": "3"},
		{"4": "4", "5": "5", "6": "6"},
		nil,
		{"7": "7", "8": "8", "9": "9"},
	}, v)

	dec = NewDecoder(([]byte)(sb.String()))
	dec.MapType = MapTypeSIMap
	dec.ListType = ListTypeSlice
	dec.Decode(&v)
	assert.Equal(t, data, v)
}

func TestDecodeInterfaceSliceToInterface(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode([]interface{}{1, 2.5, "3"})

	dec := NewDecoder(([]byte)(sb.String()))
	var v interface{}
	dec.Decode(&v)
	assert.Equal(t, []interface{}{1, 2.5, "3"}, v)

	dec = NewDecoder(([]byte)(sb.String()))
	dec.ListType = ListTypeSlice
	dec.Decode(&v)
	assert.Equal(t, []interface{}{1, 2.5, "3"}, v)

}
