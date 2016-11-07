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
 * util/util_test.go                                      *
 *                                                        *
 * util test for Go.                                      *
 *                                                        *
 * LastModified: Nov 7, 2016                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package util

import (
	"math"
	"reflect"
	"strconv"
	"testing"
)

func BenchmarkGetIntBytes(b *testing.B) {
	buf := make([]byte, 20)
	for i := 0; i < b.N; i++ {
		GetIntBytes(buf, int64(i))
		GetIntBytes(buf, int64(-i))
		GetIntBytes(buf, math.MaxInt32-int64(i))
		GetIntBytes(buf, math.MinInt32+int64(i))
		GetIntBytes(buf, math.MaxInt64-int64(i))
		GetIntBytes(buf, math.MinInt64+int64(i))
	}
}

func BenchmarkGetUintBytes(b *testing.B) {
	buf := make([]byte, 20)
	for i := 0; i < b.N; i++ {
		GetUintBytes(buf, uint64(i))
		GetUintBytes(buf, uint64(-i))
		GetUintBytes(buf, math.MaxUint32-uint64(i))
		GetUintBytes(buf, math.MaxUint32+uint64(i))
		GetUintBytes(buf, math.MaxUint64-uint64(i))
		GetUintBytes(buf, math.MaxUint64+uint64(i))
	}
}

func BenchmarkFormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(int64(i), 10)
		strconv.FormatInt(int64(-i), 10)
		strconv.FormatInt(math.MaxInt32-int64(i), 10)
		strconv.FormatInt(math.MinInt32+int64(i), 10)
		strconv.FormatInt(math.MaxInt64-int64(i), 10)
		strconv.FormatInt(math.MinInt64+int64(i), 10)
	}
}

func BenchmarkFormatUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.FormatUint(uint64(i), 10)
		strconv.FormatUint(uint64(-i), 10)
		strconv.FormatUint(math.MaxUint32-uint64(i), 10)
		strconv.FormatUint(math.MaxUint32+uint64(i), 10)
		strconv.FormatUint(math.MaxUint64-uint64(i), 10)
		strconv.FormatUint(math.MaxUint64+uint64(i), 10)
	}
}

func BenchmarkGetIntBytesParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var i int64
		buf := make([]byte, 20)
		for pb.Next() {
			GetIntBytes(buf, i)
			GetIntBytes(buf, -i)
			GetIntBytes(buf, math.MaxInt32-i)
			GetIntBytes(buf, math.MinInt32+i)
			GetIntBytes(buf, math.MaxInt64-i)
			GetIntBytes(buf, math.MinInt64+i)
			i++
		}
	})
}

func BenchmarkGetUintBytesParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var i uint64
		buf := make([]byte, 20)
		for pb.Next() {
			GetUintBytes(buf, i)
			GetUintBytes(buf, -i)
			GetUintBytes(buf, math.MaxUint32-i)
			GetUintBytes(buf, math.MaxUint32+i)
			GetUintBytes(buf, math.MaxUint64-i)
			GetUintBytes(buf, math.MaxUint64+i)
			i++
		}
	})
}

func BenchmarkFormatIntParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var i int64
		for pb.Next() {
			strconv.FormatInt(i, 10)
			strconv.FormatInt(-i, 10)
			strconv.FormatInt(math.MaxInt32-i, 10)
			strconv.FormatInt(math.MinInt32+i, 10)
			strconv.FormatInt(math.MaxInt64-i, 10)
			strconv.FormatInt(math.MinInt64+i, 10)
			i++
		}
	})
}

func BenchmarkFormatUintParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var i uint64
		for pb.Next() {
			strconv.FormatUint(i, 10)
			strconv.FormatUint(-i, 10)
			strconv.FormatUint(math.MaxUint32-i, 10)
			strconv.FormatUint(math.MaxUint32+i, 10)
			strconv.FormatUint(math.MaxUint64-i, 10)
			strconv.FormatUint(math.MaxUint64+i, 10)
			i++
		}
	})
}

func TestGetIntBytes(t *testing.T) {
	data := []int64{
		0, 9, 10, 99, 100, 999, 1000, -1000, 10000, -10000,
		123456789, -123456789, math.MaxInt32, math.MinInt32,
		math.MaxInt64, math.MinInt64}
	buf := make([]byte, 20)
	for _, i := range data {
		b := GetIntBytes(buf, i)
		if !reflect.DeepEqual(b, []byte(strconv.FormatInt(i, 10))) {
			t.Error("b must be []byte(\"" + strconv.FormatInt(i, 10) + "\")")
		}
	}
}

func TestGetUintBytes(t *testing.T) {
	data := []uint64{
		0, 9, 10, 99, 100, 999, 1000, 10000, 123456789,
		math.MaxInt32, math.MaxUint32, math.MaxInt64, math.MaxUint64}
	buf := make([]byte, 20)
	for _, i := range data {
		b := GetUintBytes(buf, i)
		if !reflect.DeepEqual(b, []byte(strconv.FormatUint(i, 10))) {
			t.Error("b must be []byte(\"" + strconv.FormatUint(i, 10) + "\")")
		}
	}
}

func TestUTF16Length(t *testing.T) {
	data := map[string]int{
		"":                            0,
		"π":                           1,
		"你":                           1,
		"你好":                          2,
		"你好啊,hello!":                  10,
		"🇨🇳":                          4,
		string([]byte{128, 129, 130}): -1,
	}
	for k, v := range data {
		if UTF16Length(k) != v {
			t.Error("The UTF16Length of \"" + k + "\" must be " + strconv.Itoa(v))
		}
	}
}

func TestByteString(t *testing.T) {
	s := ([]byte)("你好")
	if ByteString(s) != "你好" {
		t.Error(s)
	}
}

func TestStringByte(t *testing.T) {
	s := "你好"
	if ByteString(StringByte(s)) != "你好" {
		t.Error(s)
	}
}

func TestItoa(t *testing.T) {
	if Itoa(1234567) != "1234567" {
		t.Error(Itoa(1234567))
	}
}

func BenchmarkUtilItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Itoa(i)
	}
}

func BenchmarkStrconvItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.Itoa(i)
	}
}

func TestMin(t *testing.T) {
	if Min(1, 2) != 1 {
		t.Error(Min(1, 2))
	}
}

func TestMax(t *testing.T) {
	if Max(1, 2) != 2 {
		t.Error(Min(1, 2))
	}
}

func TestToUint32(t *testing.T) {
	x := ToUint32([]byte{0x12, 0x34, 0x56, 0x78})
	if x != 0x12345678 {
		t.Error(x)
	}
}

func TestFromUint32(t *testing.T) {
	var b [4]byte
	FromUint32(b[:], 0x12345678)
	if b[0] != 0x12 && b[1] != 0x34 && b[2] != 0x56 && b[3] != 0x78 {
		t.Error(b)
	}
}

type local struct{}

func (local) F1() string {
	return "OK"
}
func (local) F2(a, b int) int {
	return a + b
}

func TestLocalProxy(t *testing.T) {
	type Proxy struct {
		F1 func() string
		F2 func(a, b int) int
	}
	var p *Proxy
	LocalProxy(&p, local{})
	if p.F1() != "OK" {
		t.Error(p.F1())
	}
	if p.F2(1, 2) != 3 {
		t.Error(p.F2(1, 2))
	}
}
