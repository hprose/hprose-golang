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
 * io/reader_test.go                                      *
 *                                                        *
 * hprose Reader Test for Go.                             *
 *                                                        *
 * LastModified: Oct 13, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import (
	"bytes"
	"container/list"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestReadBool(t *testing.T) {
	trueValue := "true"
	data := map[interface{}]bool{
		true:            true,
		false:           false,
		nil:             false,
		"":              false,
		0:               false,
		1:               true,
		9:               true,
		100:             true,
		100000000000000: true,
		0.0:             false,
		"t":             true,
		"f":             false,
		&trueValue:      true,
		&trueValue:      true,
		"false":         false,
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&trueValue)
	reader := NewReader(w.Bytes(), false)
	for _, k := range keys {
		b := reader.ReadBool()
		if b != data[k] {
			t.Error(k, data[k], b)
		}
	}
	b := reader.ReadBool()
	if b != true {
		t.Error(trueValue, true, b)
	}
}

func TestUnserializeBool(t *testing.T) {
	trueValue := "true"
	data := map[interface{}]bool{
		true:            true,
		false:           false,
		nil:             false,
		"":              false,
		0:               false,
		1:               true,
		9:               true,
		100:             true,
		100000000000000: true,
		0.0:             false,
		"t":             true,
		"f":             false,
		&trueValue:      true,
		"false":         false,
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&trueValue)
	reader := NewReader(w.Bytes(), false)
	var p bool
	for _, k := range keys {
		reader.Unserialize(&p)
		if p != data[k] {
			t.Error(k, data[k], p)
		}
	}
	reader.Unserialize(&p)
	if p != true {
		t.Error(trueValue, true, p)
	}
}

func BenchmarkReadBool(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(true)
	bytes := w.Bytes()
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.ReadBool()
	}
}

func BenchmarkUnserializeBool(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(true)
	bytes := w.Bytes()
	var p bool
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func TestReadInt(t *testing.T) {
	intValue := "1234567"
	u := uint(math.MaxUint64)
	data := map[interface{}]int64{
		true:                             1,
		false:                            0,
		nil:                              0,
		"":                               0,
		0:                                0,
		1:                                1,
		9:                                9,
		100:                              100,
		-100:                             -100,
		math.MinInt32:                    int64(math.MinInt32),
		math.MaxInt64:                    int64(math.MaxInt64),
		math.MinInt64:                    int64(math.MinInt64),
		u:                                int64(u),
		0.0:                              0,
		"1":                              1,
		"9":                              9,
		&intValue:                        1234567,
		time.Unix(123, 456):              123000000456,
		time.Unix(1234567890, 123456789): 1234567890123456789,
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&intValue)
	reader := NewReader(w.Bytes(), false)
	for _, k := range keys {
		i := reader.ReadInt()
		if i != data[k] {
			t.Error(k, data[k], i)
		}
	}
	i := reader.ReadInt()
	if i != 1234567 {
		t.Error(intValue, 1234567, i)
	}
}

func TestUnserializeInt(t *testing.T) {
	intValue := "1234567"
	u := uint(math.MaxUint64)
	data := map[interface{}]int{
		true:                             1,
		false:                            0,
		nil:                              0,
		"":                               0,
		0:                                0,
		1:                                1,
		9:                                9,
		100:                              100,
		-100:                             -100,
		math.MinInt32:                    int(math.MinInt32),
		math.MaxInt64:                    int(math.MaxInt64),
		math.MinInt64:                    int(math.MinInt64),
		u:                                int(u),
		0.0:                              0,
		"1":                              1,
		"9":                              9,
		&intValue:                        1234567,
		time.Unix(123, 456):              123000000456,
		time.Unix(1234567890, 123456789): 1234567890123456789,
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&intValue)
	reader := NewReader(w.Bytes(), false)
	var p int
	for _, k := range keys {
		reader.Unserialize(&p)
		if p != data[k] {
			t.Error(k, data[k], p)
		}
	}
	reader.Unserialize(&p)
	if p != 1234567 {
		t.Error(intValue, 1234567, p)
	}
}

func BenchmarkReadInt(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(12345)
	bytes := w.Bytes()
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.ReadInt()
	}
}

func BenchmarkUnserializeInt(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(12345)
	bytes := w.Bytes()
	var p int
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func TestReadUint(t *testing.T) {
	intValue := "1234567"
	u := uint(math.MaxUint64)
	data := map[interface{}]uint64{
		true:                             1,
		false:                            0,
		nil:                              0,
		"":                               0,
		0:                                0,
		1:                                1,
		9:                                9,
		100:                              100,
		math.MaxInt64:                    uint64(math.MaxInt64),
		u:                                uint64(u),
		0.0:                              0,
		"1":                              1,
		"9":                              9,
		&intValue:                        1234567,
		time.Unix(123, 456):              123000000456,
		time.Unix(1234567890, 123456789): 1234567890123456789,
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&intValue)
	reader := NewReader(w.Bytes(), false)
	for _, k := range keys {
		i := reader.ReadUint()
		if i != data[k] {
			t.Error(k, data[k], i)
		}
	}
	i := reader.ReadUint()
	if i != 1234567 {
		t.Error(intValue, 1234567, i)
	}
}

func TestUnserializeUint(t *testing.T) {
	intValue := "1234567"
	u := uint(math.MaxUint64)
	data := map[interface{}]uint{
		true:                             1,
		false:                            0,
		nil:                              0,
		"":                               0,
		0:                                0,
		1:                                1,
		9:                                9,
		100:                              100,
		math.MaxInt64:                    uint(math.MaxInt64),
		u:                                uint(u),
		0.0:                              0,
		"1":                              1,
		"9":                              9,
		&intValue:                        1234567,
		time.Unix(123, 456):              123000000456,
		time.Unix(1234567890, 123456789): 1234567890123456789,
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&intValue)
	reader := NewReader(w.Bytes(), false)
	var p uint
	for _, k := range keys {
		reader.Unserialize(&p)
		if p != data[k] {
			t.Error(k, data[k], p)
		}
	}
	reader.Unserialize(&p)
	if p != 1234567 {
		t.Error(intValue, 1234567, p)
	}
}

func BenchmarkReadUint(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(12345)
	bytes := w.Bytes()
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.ReadUint()
	}
}

func BenchmarkUnserializeUint(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(12345)
	bytes := w.Bytes()
	var p uint
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func TestReadFloat32(t *testing.T) {
	floatValue := "3.14159"
	data := map[interface{}]float32{
		true:                             1,
		false:                            0,
		nil:                              0,
		"":                               0,
		0:                                0,
		1:                                1,
		9:                                9,
		100:                              100,
		math.MaxInt64:                    float32(math.MaxInt64),
		math.MaxFloat32:                  math.MaxFloat32,
		0.0:                              0,
		"1":                              1,
		"9":                              9,
		&floatValue:                      3.14159,
		time.Unix(123, 456):              float32(123.000000456),
		time.Unix(1234567890, 123456789): float32(1234567890.123456789),
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&floatValue)
	reader := NewReader(w.Bytes(), false)
	for _, k := range keys {
		x := reader.ReadFloat32()
		if x != data[k] {
			t.Error(k, data[k], x)
		}
	}
	x := reader.ReadFloat32()
	if x != float32(3.14159) {
		t.Error(floatValue, 3.14159, x)
	}
}

func TestUnserializeFloat32(t *testing.T) {
	floatValue := "3.14159"
	data := map[interface{}]float32{
		true:                             1,
		false:                            0,
		nil:                              0,
		"":                               0,
		0:                                0,
		1:                                1,
		9:                                9,
		100:                              100,
		math.MaxInt64:                    float32(math.MaxInt64),
		math.MaxFloat32:                  math.MaxFloat32,
		0.0:                              0,
		"1":                              1,
		"9":                              9,
		&floatValue:                      3.14159,
		time.Unix(123, 456):              float32(123.000000456),
		time.Unix(1234567890, 123456789): float32(1234567890.123456789),
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&floatValue)
	reader := NewReader(w.Bytes(), false)
	var p float32
	for _, k := range keys {
		reader.Unserialize(&p)
		if p != data[k] {
			t.Error(k, data[k], p)
		}
	}
	reader.Unserialize(&p)
	if p != float32(3.14159) {
		t.Error(floatValue, 3.14159, p)
	}
}

func BenchmarkReadFloat32(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(3.14159)
	bytes := w.Bytes()
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.ReadFloat32()
	}
}

func BenchmarkUnserializeFloat32(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(3.14159)
	bytes := w.Bytes()
	var p float32
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func TestReadFloat64(t *testing.T) {
	floatValue := "3.14159"
	data := map[interface{}]float64{
		true:                             1,
		false:                            0,
		nil:                              0,
		"":                               0,
		0:                                0,
		1:                                1,
		9:                                9,
		100:                              100,
		math.MaxFloat32:                  math.MaxFloat32,
		math.MaxFloat64:                  math.MaxFloat64,
		0.0:                              0,
		"1":                              1,
		"9":                              9,
		&floatValue:                      3.14159,
		time.Unix(123, 456):              float64(123.000000456),
		time.Unix(1234567890, 123456789): float64(1234567890.123456789),
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&floatValue)
	reader := NewReader(w.Bytes(), false)
	for _, k := range keys {
		x := reader.ReadFloat64()
		if x != data[k] {
			t.Error(k, data[k], x)
		}
	}
	x := reader.ReadFloat64()
	if x != float64(3.14159) {
		t.Error(floatValue, 3.14159, x)
	}
}

func TestUnserializeFloat64(t *testing.T) {
	floatValue := "3.14159"
	data := map[interface{}]float64{
		true:                             1,
		false:                            0,
		nil:                              0,
		"":                               0,
		0:                                0,
		1:                                1,
		9:                                9,
		100:                              100,
		math.MaxFloat32:                  math.MaxFloat32,
		math.MaxFloat64:                  math.MaxFloat64,
		0.0:                              0,
		"1":                              1,
		"9":                              9,
		&floatValue:                      3.14159,
		time.Unix(123, 456):              float64(123.000000456),
		time.Unix(1234567890, 123456789): float64(1234567890.123456789),
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&floatValue)
	reader := NewReader(w.Bytes(), false)
	var p float64
	for _, k := range keys {
		reader.Unserialize(&p)
		if p != data[k] {
			t.Error(k, data[k], p)
		}
	}
	reader.Unserialize(&p)
	if p != float64(3.14159) {
		t.Error(floatValue, 3.14159, p)
	}
}

func BenchmarkReadFloat64(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(3.14159)
	bytes := w.Bytes()
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.ReadFloat64()
	}
}

func BenchmarkUnserializeFloat64(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(3.14159)
	bytes := w.Bytes()
	var p float64
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func TestUnserializeArray(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]byte{'h', 'e', 'l', 'l', 'o'}
	w := NewWriter(false)
	w.Serialize(&a)
	w.Serialize(&b)
	w.Serialize(&a)
	w.Serialize(&b)
	w.Serialize(&a)
	w.Serialize(&b)
	w.Serialize(nil)
	w.Serialize("")
	reader := NewReader(w.Bytes(), false)
	var a1 [5]int
	reader.Unserialize(&a1)
	if !reflect.DeepEqual(a1, a) {
		t.Error(a1, a)
	}
	var b1 [5]byte
	reader.Unserialize(&b1)
	if !reflect.DeepEqual(b1, b) {
		t.Error(b1, b)
	}
	var a2 [4]int
	reader.Unserialize(&a2)
	if !reflect.DeepEqual(a2[:4], a[:4]) {
		t.Error(a2[:4], a[:4])
	}
	var b2 [4]byte
	reader.Unserialize(&b2)
	if !reflect.DeepEqual(b2[:4], b[:4]) {
		t.Error(b2[:4], b[:4])
	}
	var a3 [6]int
	reader.Unserialize(&a3)
	if !reflect.DeepEqual(a3, [6]int{1, 2, 3, 4, 5, 0}) {
		t.Error(a3)
	}
	var b3 [6]byte
	reader.Unserialize(&b3)
	if !reflect.DeepEqual(b3, [6]byte{'h', 'e', 'l', 'l', 'o', 0}) {
		t.Error(b3)
	}
	reader.Unserialize(&a3)
	if !reflect.DeepEqual(a3, [6]int{0, 0, 0, 0, 0, 0}) {
		t.Error(a3)
	}
	reader.Unserialize(&b3)
	if !reflect.DeepEqual(b3, [6]byte{0, 0, 0, 0, 0, 0}) {
		t.Error(b3)
	}
}

func BenchmarkUnserializeByteArray(b *testing.B) {
	w := NewWriter(true)
	w.Serialize([5]byte{'h', 'e', 'l', 'l', 'o'})
	bytes := w.Bytes()
	var p [5]byte
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func BenchmarkUnserializeIntArray(b *testing.B) {
	w := NewWriter(true)
	w.Serialize([5]int{1, 2, 3, 4, 5})
	bytes := w.Bytes()
	var p [5]int
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func TestUnserializeSlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []byte{'h', 'e', 'l', 'l', 'o'}
	w := NewWriter(false)
	w.Serialize(a)
	w.Serialize(b)
	w.Serialize(a)
	w.Serialize(b)
	w.Serialize(a)
	w.Serialize(b)
	w.Serialize(nil)
	w.Serialize("")
	reader := NewReader(w.Bytes(), false)
	var a1 []int
	reader.Unserialize(&a1)
	if !reflect.DeepEqual(a1, a) {
		t.Error(a1, a)
	}
	var b1 []byte
	reader.Unserialize(&b1)
	if !reflect.DeepEqual(b1, b) {
		t.Error(b1, b)
	}
	a2 := []int{}
	reader.Unserialize(&a2)
	if !reflect.DeepEqual(a2, a) {
		t.Error(a2, a)
	}
	b2 := []byte{}
	reader.Unserialize(&b2)
	if !reflect.DeepEqual(b2, b) {
		t.Error(b2, b)
	}
	a2 = make([]int, 10)
	reader.Unserialize(&a2)
	if !reflect.DeepEqual(a2, a) {
		t.Error(a2, a)
	}
	b2 = make([]byte, 10)
	reader.Unserialize(&b2)
	if !reflect.DeepEqual(b2, b) {
		t.Error(b2, b)
	}
	reader.Unserialize(&a2)
	if a2 != nil {
		t.Error(a2, nil)
	}
	reader.Unserialize(&b2)
	if b2 != nil {
		t.Error(b2, nil)
	}
}

func TestUnserializeSliceRef(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []byte{'h', 'e', 'l', 'l', 'o'}
	w := NewWriter(false)
	w.Serialize(&a)
	w.Serialize(&b)
	w.Serialize(&a)
	w.Serialize(&b)
	w.Serialize(&a)
	w.Serialize(&b)
	reader := NewReader(w.Bytes(), false)
	var a1 []int
	reader.Unserialize(&a1)
	if !reflect.DeepEqual(a1, a) {
		t.Error(a1, a)
	}
	var b1 []byte
	reader.Unserialize(&b1)
	if !reflect.DeepEqual(b1, b) {
		t.Error(b1, b)
	}
	a2 := []int{}
	reader.Unserialize(&a2)
	if !reflect.DeepEqual(a2, a) {
		t.Error(a2, a)
	}
	b2 := []byte{}
	reader.Unserialize(&b2)
	if !reflect.DeepEqual(b2, b) {
		t.Error(b2, b)
	}
	a2 = make([]int, 10)
	reader.Unserialize(&a2)
	if !reflect.DeepEqual(a2, a) {
		t.Error(a2, a)
	}
	b2 = make([]byte, 10)
	reader.Unserialize(&b2)
	if !reflect.DeepEqual(b2, b) {
		t.Error(b2, b)
	}
}

func BenchmarkUnserializeByteSlice(b *testing.B) {
	w := NewWriter(true)
	w.Serialize([]byte{'h', 'e', 'l', 'l', 'o'})
	bytes := w.Bytes()
	var p []byte
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func BenchmarkUnserializeIntSlice(b *testing.B) {
	w := NewWriter(true)
	w.Serialize([]int{1, 2, 3, 4, 5})
	bytes := w.Bytes()
	var p []int
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func TestUnserializeComplex64(t *testing.T) {
	complex64Value := "3.14159"
	complex64Slice := []float32{math.MaxFloat32, math.MaxFloat32}
	data := map[interface{}]complex64{
		true:                                      1,
		false:                                     0,
		nil:                                       0,
		"":                                        0,
		0:                                         0,
		1:                                         1,
		9:                                         9,
		100:                                       100,
		math.MaxInt64:                             complex(float32(math.MaxInt64), 0),
		math.MaxFloat32:                           complex(math.MaxFloat32, 0),
		0.0:                                       0,
		"1":                                       1,
		"9":                                       9,
		&complex64Value:                           complex(float32(3.14159), 0),
		complex(math.MaxFloat32, math.MaxFloat32): complex(math.MaxFloat32, math.MaxFloat32),
		&complex64Slice:                           complex(math.MaxFloat32, math.MaxFloat32),
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&complex64Value)
	w.Serialize(&complex64Slice)
	reader := NewReader(w.Bytes(), false)
	var p complex64
	for _, k := range keys {
		reader.Unserialize(&p)
		if p != data[k] {
			t.Error(k, data[k], p)
		}
	}
	reader.Unserialize(&p)
	if p != complex64(3.14159) {
		t.Error(complex64Value, 3.14159, p)
	}
	reader.Unserialize(&p)
	if p != complex(math.MaxFloat32, math.MaxFloat32) {
		t.Error(complex64Value, complex(math.MaxFloat32, math.MaxFloat32), p)
	}
}

func BenchmarkReadComplex64(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(complex(math.MaxFloat32, math.MaxFloat32))
	bytes := w.Bytes()
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.ReadComplex64()
	}
}

func BenchmarkUnserializeComplex64(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(complex(math.MaxFloat32, math.MaxFloat32))
	bytes := w.Bytes()
	var p complex64
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func TestUnserializeComplex128(t *testing.T) {
	complex128Value := "3.14159"
	complex128Slice := []float64{math.MaxFloat64, math.MaxFloat64}
	data := map[interface{}]complex128{
		true:                                      1,
		false:                                     0,
		nil:                                       0,
		"":                                        0,
		0:                                         0,
		1:                                         1,
		9:                                         9,
		100:                                       100,
		math.MaxFloat32:                           complex(math.MaxFloat32, 0),
		0.0:                                       0,
		"1":                                       1,
		"9":                                       9,
		&complex128Value:                          complex(float64(3.14159), 0),
		complex(math.MaxFloat32, math.MaxFloat32): complex(math.MaxFloat32, math.MaxFloat32),
		&complex128Slice:                          complex(math.MaxFloat64, math.MaxFloat64),
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&complex128Value)
	w.Serialize(&complex128Slice)
	reader := NewReader(w.Bytes(), false)
	var p complex128
	for _, k := range keys {
		reader.Unserialize(&p)
		if p != data[k] {
			t.Error(k, data[k], p)
		}
	}
	reader.Unserialize(&p)
	if p != complex128(3.14159) {
		t.Error(complex128Value, 3.14159, p)
	}
	reader.Unserialize(&p)
	if p != complex(math.MaxFloat64, math.MaxFloat64) {
		t.Error(complex128Value, complex(math.MaxFloat64, math.MaxFloat64), p)
	}
}

func BenchmarkReadComplex128(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(complex(math.MaxFloat64, math.MaxFloat64))
	bytes := w.Bytes()
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.ReadComplex128()
	}
}

func BenchmarkUnserializeComplex128(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(complex(math.MaxFloat64, math.MaxFloat64))
	bytes := w.Bytes()
	var p complex128
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func TestUnserializeListAsMap(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	m := map[int]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	w := NewWriter(false)
	w.Serialize(a)
	w.Serialize(a)
	w.Serialize(a)
	reader := NewReader(w.Bytes(), false)
	var m1 map[int]int
	reader.Unserialize(&m1)
	if !reflect.DeepEqual(m1, m) {
		t.Error(m1, m, a)
	}
	m2 := make(map[int]int)
	reader.Unserialize(&m2)
	if !reflect.DeepEqual(m2, m) {
		t.Error(m2, m, a)
	}
	m3 := make(map[int]int, 10)
	reader.Unserialize(&m3)
	if !reflect.DeepEqual(m3, m) {
		t.Error(m3, m, a)
	}
}

func TestUnserializeListRefAsMapRef(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	m := map[int]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	w := NewWriter(false)
	w.Serialize(&a)
	w.Serialize(&a)
	w.Serialize(&a)
	w.Serialize(&a)
	w.Serialize(&a)
	w.Serialize(&a)
	reader := NewReader(w.Bytes(), false)
	var m1 map[int]int
	reader.Unserialize(&m1)
	if !reflect.DeepEqual(m1, m) {
		t.Error(m1, m, a)
	}
	m2 := make(map[int]int)
	reader.Unserialize(&m2)
	if !reflect.DeepEqual(m2, m) {
		t.Error(m2, m, a)
	}
	m3 := make(map[int]int, 10)
	reader.Unserialize(&m3)
	if !reflect.DeepEqual(m3, m) {
		t.Error(m3, m, a)
	}
	var m4 map[int]int
	reader.Unserialize(&m4)
	if !reflect.DeepEqual(m4, m) {
		t.Error(m4, m, a)
	}
	m5 := make(map[int]int)
	reader.Unserialize(&m5)
	if !reflect.DeepEqual(m5, m) {
		t.Error(m5, m, a)
	}
	m6 := make(map[int]int, 10)
	reader.Unserialize(&m6)
	if !reflect.DeepEqual(m6, m) {
		t.Error(m6, m, a)
	}
}

func TestUnserializeMap(t *testing.T) {
	m := map[string]string{
		"name": "Tom",
		"国家":   "🇨🇳",
	}
	w := NewWriter(false)
	w.Serialize(m)
	w.Serialize(m)
	w.Serialize(m)
	w.Serialize(nil)
	w.Serialize("")
	reader := NewReader(w.Bytes(), false)
	var m1 map[string]string
	reader.Unserialize(&m1)
	if !reflect.DeepEqual(m1, m) {
		t.Error(m1, m)
	}
	m2 := make(map[string]string)
	reader.Unserialize(&m2)
	if !reflect.DeepEqual(m2, m) {
		t.Error(m2, m)
	}
	m3 := make(map[string]string, 10)
	reader.Unserialize(&m3)
	if !reflect.DeepEqual(m3, m) {
		t.Error(m3, m)
	}
	reader.Unserialize(&m3)
	if m3 != nil {
		t.Error(m3, nil)
	}
	reader.Unserialize(&m3)
	if m3 != nil {
		t.Error(m3, nil)
	}
}

func BenchmarkUnserializeSliceAsMap(b *testing.B) {
	w := NewWriter(true)
	w.Serialize([5]int{1, 2, 3, 4, 5})
	bytes := w.Bytes()
	var p map[int]int
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func BenchmarkUnserializeMap(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(map[string]string{
		"name": "Tom",
		"国家":   "🇨🇳",
	})
	bytes := w.Bytes()
	var p map[string]string
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func TestUnserializeString(t *testing.T) {
	strValue := "你好"
	d := time.Date(1980, 12, 1, 0, 0, 0, 0, time.UTC)
	ti := time.Date(1970, 1, 1, 12, 23, 45, 123456789, time.Local)
	dt := time.Date(2006, 9, 9, 12, 23, 45, 123456789, time.UTC)
	data := map[interface{}]string{
		true:            "true",
		false:           "false",
		nil:             "",
		"":              "",
		0:               "0",
		1:               "1",
		9:               "9",
		100:             "100",
		math.MaxFloat32: strconv.FormatFloat(math.MaxFloat32, 'g', -1, 64),
		math.MaxFloat64: strconv.FormatFloat(math.MaxFloat64, 'g', -1, 64),
		0.0:             "0",
		"1":             "1",
		"9":             "9",
		&strValue:       "你好",
		d:               d.String(),
		ti:              ti.String(),
		dt:              dt.String(),
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&strValue)
	reader := NewReader(w.Bytes(), false)
	var p string
	for _, k := range keys {
		reader.Unserialize(&p)
		if p != data[k] {
			t.Error(k, data[k], p)
		}
	}
	reader.Unserialize(&p)
	if p != strValue {
		t.Error(strValue, p)
	}
}

func BenchmarkReadString(b *testing.B) {
	w := NewWriter(true)
	w.Serialize("你好，🇨🇳")
	bytes := w.Bytes()
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.ReadString()
	}
}

func BenchmarkUnserializeString(b *testing.B) {
	w := NewWriter(true)
	w.Serialize("你好，🇨🇳")
	bytes := w.Bytes()
	var p string
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func TestUnserializeInterface(t *testing.T) {
	strValue := "你好"
	d := time.Date(1980, 12, 1, 0, 0, 0, 0, time.UTC)
	ti := time.Date(1970, 1, 1, 12, 23, 45, 123456789, time.Local)
	dt := time.Date(2006, 9, 9, 12, 23, 45, 123456789, time.UTC)
	a := []int{1, 2, 3, 4, 5}
	m := map[int]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	data := map[interface{}]interface{}{
		true:            true,
		false:           false,
		nil:             nil,
		"":              "",
		0:               0,
		1:               1,
		9:               9,
		100:             100,
		math.MaxInt32:   math.MaxInt32,
		math.MaxFloat32: math.MaxFloat32,
		math.MaxFloat64: math.MaxFloat64,
		0.0:             0.0,
		"1":             "1",
		"9":             "9",
		&strValue:       "你好",
		d:               d,
		ti:              ti,
		dt:              dt,
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&strValue)
	w.Serialize(math.MaxInt64)
	w.Serialize(a)
	w.Serialize(m)
	reader := NewReader(w.Bytes(), false)
	var p interface{}
	for _, k := range keys {
		reader.Unserialize(&p)
		if p != data[k] {
			t.Error(k, data[k], p)
		}
	}
	reader.Unserialize(&p)
	if p != strValue {
		t.Error(strValue, p)
	}
	reader.Unserialize(&p)
	if p.(*big.Int).Cmp(big.NewInt(math.MaxInt64)) != 0 {
		t.Error(p)
	}
	reader.Unserialize(&p)
	if !reflect.DeepEqual(p.([]interface{}), []interface{}{1, 2, 3, 4, 5}) {
		t.Error(p)
	}
	reader.Unserialize(&p)
	if !reflect.DeepEqual(p.(map[interface{}]interface{}), map[interface{}]interface{}{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}) {
		t.Error(p)
	}
}

func TestUnserializeIntPtr(t *testing.T) {
	i := 123
	w := NewWriter(false)
	w.Serialize(&i)
	w.Serialize(&i)
	reader := NewReader(w.Bytes(), false)
	var p *int
	reader.Unserialize(&p)
	if *p != i {
		t.Error(*p, i)
	}
	reader.Unserialize(&p)
	if *p != i {
		t.Error(*p, i)
	}
}

func TestUnserializeFloatPtr(t *testing.T) {
	f := 3.14159
	w := NewWriter(false)
	w.Serialize(&f)
	w.Serialize(&f)
	reader := NewReader(w.Bytes(), false)
	var p *float64
	reader.Unserialize(&p)
	if *p != f {
		t.Error(*p, f)
	}
	reader.Unserialize(&p)
	if *p != f {
		t.Error(*p, f)
	}
}

func TestUnserializeStringPtr(t *testing.T) {
	str := "你好，🇨🇳"
	w := NewWriter(false)
	w.Serialize(&str)
	w.Serialize(&str)
	reader := NewReader(w.Bytes(), false)
	var p *string
	reader.Unserialize(&p)
	if *p != str {
		t.Error(*p, str)
	}
	reader.Unserialize(&p)
	if *p != str {
		t.Error(*p, str)
	}
}

func TestUnserializeBoolPtr(t *testing.T) {
	b := true
	w := NewWriter(false)
	w.Serialize(&b)
	w.Serialize(&b)
	reader := NewReader(w.Bytes(), false)
	var p *bool
	reader.Unserialize(&p)
	if *p != b {
		t.Error(*p, b)
	}
	reader.Unserialize(&p)
	if *p != b {
		t.Error(*p, b)
	}
}

func TestUnserializeArrayPtr(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	w := NewWriter(false)
	w.Serialize(&a)
	w.Serialize(&a)
	reader := NewReader(w.Bytes(), false)
	var p *[5]int
	reader.Unserialize(&p)
	if !reflect.DeepEqual(*p, a) {
		t.Error(*p, a)
	}
	reader.Unserialize(&p)
	if !reflect.DeepEqual(*p, a) {
		t.Error(*p, a)
	}
}

func TestUnserializeSlicePtr(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	w := NewWriter(false)
	w.Serialize(&slice)
	w.Serialize(&slice)
	reader := NewReader(w.Bytes(), false)
	var p *[]int
	reader.Unserialize(&p)
	if !reflect.DeepEqual(*p, slice) {
		t.Error(*p, slice)
	}
	reader.Unserialize(&p)
	if !reflect.DeepEqual(*p, slice) {
		t.Error(*p, slice)
	}
}

func TestUnserializeMapPtr(t *testing.T) {
	m := map[string]interface{}{
		"name": "小明",
		"age":  14,
	}
	w := NewWriter(false)
	w.Serialize(&m)
	w.Serialize(&m)
	reader := NewReader(w.Bytes(), false)
	var p *map[string]interface{}
	reader.Unserialize(&p)
	if !reflect.DeepEqual(*p, m) {
		t.Error(*p, m)
	}
	reader.Unserialize(&p)
	if !reflect.DeepEqual(*p, m) {
		t.Error(*p, m)
	}
}

func TestUnserializeIntPtrPtr(t *testing.T) {
	i := 123
	ip := &i
	w := NewWriter(false)
	w.Serialize(&ip)
	w.Serialize(&ip)
	reader := NewReader(w.Bytes(), false)
	var p **int
	reader.Unserialize(&p)
	if **p != i {
		t.Error(**p, i)
	}
	reader.Unserialize(&p)
	if **p != i {
		t.Error(**p, i)
	}
}

func TestUnserializeNilPtr(t *testing.T) {
	w := NewWriter(false)
	w.Serialize(nil)
	w.Serialize(nil)
	reader := NewReader(w.Bytes(), false)
	var p *int
	reader.Unserialize(&p)
	if p != nil {
		t.Error(p, nil)
	}
	var pb *[]byte
	reader.Unserialize(&pb)
	if pb != nil {
		t.Error(pb, nil)
	}
}

func TestUnserializeBigInt(t *testing.T) {
	w := NewWriter(false)
	bi, _ := new(big.Int).SetString("1234567890987654321234567890987654321", 10)
	data := []*big.Int{
		big.NewInt(0),
		big.NewInt(1),
		big.NewInt(10),
		big.NewInt(100),
		big.NewInt(math.MaxInt64),
		big.NewInt(math.MinInt64),
		bi,
	}
	w.Serialize(nil)
	w.Serialize(1)
	w.Serialize(10)
	w.Serialize("1234567890987654321234567890987654321")
	for _, v := range data {
		w.Serialize(v)
	}
	reader := NewReader(w.Bytes(), false)
	var p *big.Int
	reader.Unserialize(&p)
	if p != nil {
		t.Error(p, nil)
	}
	reader.Unserialize(&p)
	if p.Cmp(big.NewInt(1)) != 0 {
		t.Error(p, big.NewInt(1))
	}
	reader.Unserialize(&p)
	if p.Cmp(big.NewInt(10)) != 0 {
		t.Error(p, big.NewInt(10))
	}
	reader.Unserialize(&p)
	if p.Cmp(bi) != 0 {
		t.Error(p, bi)
	}
	for _, v := range data {
		reader.Unserialize(&p)
		if p.Cmp(v) != 0 {
			t.Error(p, v)
		}
	}
}

func BenchmarkUnserializeNilInterface(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(nil)
	bytes := w.Bytes()
	var p interface{}
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func BenchmarkUnserializeNilPtr(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(nil)
	bytes := w.Bytes()
	var p *interface{}
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func BenchmarkUnserializeIntInterface(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(123)
	bytes := w.Bytes()
	var p interface{}
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func BenchmarkUnserializeIntPtr(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(123)
	bytes := w.Bytes()
	var p *int
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func BenchmarkUnserializeBigInt(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(123)
	bytes := w.Bytes()
	var p *big.Int
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func TestUnserializeBigRat(t *testing.T) {
	w := NewWriter(false)
	br, _ := new(big.Rat).SetString("1234567890987654321234567/890987654321")
	data := []*big.Rat{
		big.NewRat(0, 1),
		big.NewRat(1, 2),
		big.NewRat(10, 3),
		big.NewRat(100, 4),
		big.NewRat(math.MaxInt64, 1),
		big.NewRat(math.MinInt64, 4),
		br,
	}
	w.Serialize(nil)
	w.Serialize(1)
	w.Serialize(10)
	w.Serialize("1234567890987654321234567/890987654321")
	for _, v := range data {
		w.Serialize(v)
	}
	reader := NewReader(w.Bytes(), false)
	var p *big.Rat
	reader.Unserialize(&p)
	if p != nil {
		t.Error(p, nil)
	}
	reader.Unserialize(&p)
	if p.Cmp(big.NewRat(1, 1)) != 0 {
		t.Error(p, big.NewRat(1, 1))
	}
	reader.Unserialize(&p)
	if p.Cmp(big.NewRat(10, 1)) != 0 {
		t.Error(p, big.NewRat(10, 1))
	}
	reader.Unserialize(&p)
	if p.Cmp(br) != 0 {
		t.Error(p, br)
	}
	for _, v := range data {
		reader.Unserialize(&p)
		if p.Cmp(v) != 0 {
			t.Error(p, v)
		}
	}
}

func TestUnserializeBigFloat(t *testing.T) {
	w := NewWriter(false)
	bl, _, _ := new(big.Float).Parse(strconv.Itoa(math.MaxInt64), 10)
	bf, _, _ := new(big.Float).Parse("1234567890987654321234567.890987654321", 10)
	w.Serialize(nil)
	w.Serialize(1)
	w.Serialize(10)
	w.Serialize(math.MaxInt64)
	w.Serialize(bf)
	w.Serialize("1234567890987654321234567.890987654321")
	reader := NewReader(w.Bytes(), false)
	var p *big.Float
	reader.Unserialize(&p)
	if p != nil {
		t.Error(p, nil)
	}
	reader.Unserialize(&p)
	if p.Cmp(big.NewFloat(1)) != 0 {
		t.Error(p, big.NewFloat(1))
	}
	reader.Unserialize(&p)
	if p.Cmp(big.NewFloat(10)) != 0 {
		t.Error(p, big.NewFloat(10))
	}
	reader.Unserialize(&p)
	if p.Cmp(bl) != 0 {
		t.Error(p, bl)
	}
	reader.Unserialize(&p)
	if p.Cmp(bf) != 0 {
		t.Error(p, bf)
	}
	reader.Unserialize(&p)
	if p.Cmp(bf) != 0 {
		t.Error(p, bf)
	}
}

func TestUnserializeTime(t *testing.T) {
	t1 := time.Date(1980, 12, 1, 12, 34, 56, 123456789, time.UTC)
	t2 := time.Date(1970, 1, 1, 12, 34, 56, 123456789, time.UTC)
	w := NewWriter(false)
	w.Serialize(nil)
	w.Serialize(1)
	w.Serialize(10)
	w.Serialize(math.MaxInt64)
	w.Serialize(t1)
	w.Serialize(t2)
	reader := NewReader(w.Bytes(), false)
	var p *time.Time
	reader.Unserialize(&p)
	if p != nil {
		t.Error(p, nil)
	}
	reader.Unserialize(&p)
	if !reflect.DeepEqual(*p, time.Unix(1, 0)) {
		t.Error(p, time.Unix(1, 0))
	}
	reader.Unserialize(&p)
	if !reflect.DeepEqual(*p, time.Unix(10, 0)) {
		t.Error(p, time.Unix(10, 0))
	}
	reader.Unserialize(&p)
	if !reflect.DeepEqual(*p, time.Unix(math.MaxInt64, 0)) {
		t.Error(p, time.Unix(math.MaxInt64, 0))
	}
	reader.Unserialize(&p)
	if !reflect.DeepEqual(*p, t1) {
		t.Error(p, t1)
	}
	reader.Unserialize(&p)
	if !reflect.DeepEqual(*p, t2) {
		t.Error(p, t2)
	}
}

func BenchmarkUnserializeTime(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(123)
	bytes := w.Bytes()
	var p time.Time
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func TestUnserializeList(t *testing.T) {
	a := list.New()
	a.PushBack(1)
	a.PushBack("hello")
	a.PushBack(true)
	w := NewWriter(false)
	w.Serialize(a)
	w.Serialize(nil)
	reader := NewReader(w.Bytes(), false)
	var p *list.List
	reader.Unserialize(&p)
	if !reflect.DeepEqual(*p, *a) {
		t.Error(p, a)
	}
	reader.Unserialize(&p)
	if p != nil {
		t.Error(p, nil)
	}
}

func BenchmarkUnserializeList(b *testing.B) {
	a := list.New()
	a.PushBack(1)
	a.PushBack("hello")
	a.PushBack(true)
	w := NewWriter(true)
	w.Serialize(a)
	bytes := w.Bytes()
	var p *list.List
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func TestUnserializeMapAsStruct(t *testing.T) {
	w := NewWriter(true)
	m := make(map[interface{}]interface{})
	m["name"] = "Tom"
	m["age"] = 36
	m["male"] = true
	w.Serialize(m)
	type Test struct {
		Name string
		Age  int
		Male bool
	}
	reader := NewReader(w.Bytes(), false)
	var p *Test
	reader.Unserialize(&p)
	if p.Name != "Tom" || p.Age != 36 || p.Male != true {
		t.Error(p, m)
	}
}

func BenchmarkUnserializeMapAsStruct(b *testing.B) {
	w := NewWriter(true)
	m := make(map[interface{}]interface{})
	m["name"] = "Tom"
	m["age"] = 36
	m["male"] = true
	w.Serialize(m)
	bytes := w.Bytes()
	type Test struct {
		Name string
		Age  int
		Male bool
	}
	var p Test
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func TestUnserializeStructAsMap(t *testing.T) {
	type Test struct {
		Name string
		Age  int
		Male bool
	}
	test := Test{"Tom", 36, true}
	w := NewWriter(true)
	w.Serialize(test)
	reader := NewReader(w.Bytes(), false)
	m := make(map[string]interface{})
	m["name"] = "Tom"
	m["age"] = 36
	m["male"] = true
	var p map[string]interface{}
	reader.Unserialize(&p)
	if !reflect.DeepEqual(p, m) {
		t.Error(p, m)
	}
}

func TestUnserializeStructAsMapUnregisted(t *testing.T) {
	type TestStructAsMapUnregisted_3 struct {
		Id    string
		Child *TestStructAsMapUnregisted_3
	}
	type TestStructAsMapUnregisted_1 struct {
		Name  string
		Age   int
		Male  bool
		Child *TestStructAsMapUnregisted_3
	}
	test := TestStructAsMapUnregisted_1{"Tom", 36, true,
		&TestStructAsMapUnregisted_3{
			Id: "ok",
			Child: &TestStructAsMapUnregisted_3{
				Id: "yes",
			},
		},
	}
	w := NewWriter(true)
	w.Serialize(test)

	buf := bytes.Replace(w.Bytes(),
		[]byte("TestStructAsMapUnregisted_1"), []byte("TestStructAsMapUnregisted_2"), -1)
	buf = bytes.Replace(buf,
		[]byte("TestStructAsMapUnregisted_3"), []byte("TestStructAsMapUnregisted_4"), -1)

	reader := NewReader(buf, false)

	m := make(map[string]interface{})
	m["name"] = "Tom"
	m["age"] = 36
	m["male"] = true
	m["child"] = map[string]interface{}{
		"id": "ok",
		"child": map[string]interface{}{
			"id":    "yes",
			"child": nil,
		},
	}

	var p map[string]interface{}
	reader.Unserialize(&p)
	if !reflect.DeepEqual(p, m) {
		t.Error(p, m)
	}
}

func BenchmarkUnserializeStructAsMap(b *testing.B) {
	type Test struct {
		Name string
		Age  int
		Male bool
	}
	test := Test{"Tom", 36, true}
	w := NewWriter(true)
	w.Serialize(test)
	bytes := w.Bytes()
	var p map[string]interface{}
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}

func TestUnserializeStruct(t *testing.T) {
	type Test struct {
		Name string
		Age  int
		Male bool
	}
	test := Test{"Tom", 36, true}
	w := NewWriter(true)
	w.Serialize(test)
	reader := NewReader(w.Bytes(), false)
	var p Test
	reader.Unserialize(&p)
	if !reflect.DeepEqual(p, test) {
		t.Error(p, test)
	}
}

func TestUnserializeStructAsInterface(t *testing.T) {
	type Test struct {
		Name string
		Age  int
		Male bool
	}
	test := Test{"Tom", 36, true}
	w := NewWriter(true)
	w.Serialize(test)
	reader := NewReader(w.Bytes(), false)
	var p interface{}
	reader.Unserialize(&p)
	if !reflect.DeepEqual(p, &test) {
		t.Error(p, &test)
	}
}

func BenchmarkUnserializeStruct(b *testing.B) {
	type Test struct {
		Name string
		Age  int
		Male bool
	}
	test := Test{"Tom", 36, true}
	w := NewWriter(true)
	w.Serialize(test)
	bytes := w.Bytes()
	var p Test
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
}
