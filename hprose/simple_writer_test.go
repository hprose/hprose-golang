/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.net/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * hprose/simple_writer_test.go                           *
 *                                                        *
 * hprose SimpleWriter Test for Go.                       *
 *                                                        *
 * LastModified: Feb 11, 2014                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose_test

import (
	"bytes"
	"container/list"
	. "hprose"
	"math"
	"math/big"
	"strings"
	"testing"
	"time"
	"uuid"
)

func TestSimpleWriterNil(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	err := writer.Serialize(nil)
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != "n" {
		t.Error(b.String())
	}
}

func TestSimpleWriterByte(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	err := writer.Serialize(byte(13))
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != "i13;" {
		t.Error(b.String())
	}
}

func TestSimpleWriterUint8(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	err := writer.Serialize(uint8(0))
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != "0" {
		t.Error(b.String())
	}
}

func TestSimpleWriterUint16(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	err := writer.Serialize(uint16(12345))
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != "i12345;" {
		t.Error(b.String())
	}
}

func TestSimpleWriterUint64(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	err := writer.Serialize(uint64(12345))
	if err != nil {
		t.Error(err.Error())
	}
	err = writer.Serialize(uint64(math.MaxUint64))
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != "i12345;l18446744073709551615;" {
		t.Error(b.String())
	}
}

func TestSimpleWriterBigInt(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	var bb big.Int
	err := writer.Serialize(bb)
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != "l0;" {
		t.Error(b.String())
	}
}

func TestSimpleWriterBigIntPointer(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	bb := big.NewInt(1234567890)
	err := writer.Serialize(bb)
	if err != nil {
		t.Error(err.Error())
	}
	var bbb interface{} = *bb
	err = writer.Serialize(&bbb)
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != "l1234567890;l1234567890;" {
		t.Error(b.String())
	}
}

func TestSimpleWriterFloat64(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	err := writer.Serialize(3.1415926)
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != "d3.1415926;" {
		t.Error(b.String())
	}
}

func TestSimpleWriterNaN(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	err := writer.Serialize(math.NaN())
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != "N" {
		t.Error(b.String())
	}
}

func TestSimpleWriterInf(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	err := writer.Serialize(math.Inf(1))
	if err != nil {
		t.Error(err.Error())
	}
	err = writer.Serialize(math.Inf(-1))
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != "I+I-" {
		t.Error(b.String())
	}
}

func TestSimpleWriterBool(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	err := writer.Serialize(true)
	if err != nil {
		t.Error(err.Error())
	}
	err = writer.Serialize(false)
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != "tf" {
		t.Error(b.String())
	}
}

func TestSimpleWriterTime(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	err := writer.Serialize(time.Date(2014, 1, 19, 20, 25, 33, 12345678, time.UTC))
	if err != nil {
		t.Error(err.Error())
	}
	err = writer.Serialize(time.Date(2014, 1, 19, 0, 0, 0, 0, time.Local))
	if err != nil {
		t.Error(err.Error())
	}
	err = writer.Serialize(time.Date(1970, 1, 1, 1, 1, 1, 0, time.Local))
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != "D20140119T202533.012345678ZD20140119;T010101;" {
		t.Error(b.String())
	}
}

func TestSimpleWriterString(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	err := writer.Serialize("")
	if err != nil {
		t.Error(err.Error())
	}
	err = writer.Serialize("我")
	if err != nil {
		t.Error(err.Error())
	}
	err = writer.Serialize("我爱你")
	if err != nil {
		t.Error(err.Error())
	}
	var s interface{} = "字符串"
	err = writer.Serialize(&s)
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != `eu我s3"我爱你"s3"字符串"` {
		t.Error(b.String())
	}
}

func TestSimpleWriterBytes(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	err := writer.Serialize([]byte(""))
	if err != nil {
		t.Error(err.Error())
	}
	err = writer.Serialize([]byte("我"))
	if err != nil {
		t.Error(err.Error())
	}
	err = writer.Serialize([]byte("我爱你"))
	if err != nil {
		t.Error(err.Error())
	}
	s := []byte("字符串")
	err = writer.Serialize(&s)
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != `b""b3"我"b9"我爱你"b9"字符串"` {
		t.Error(b.String())
	}
}

func TestSimpleWriterUUID(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	err := writer.Serialize(uuid.Parse("3f257da1-0b85-48d6-8f5c-6cd13d2d60c9"))
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != "g{3f257da1-0b85-48d6-8f5c-6cd13d2d60c9}" {
		t.Error(b.String())
	}
}

func TestSimpleWriterList(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	a := list.New()
	a.PushBack(1)
	a.PushBack(2)
	a.PushBack(3)
	err := writer.Serialize(a)
	if err != nil {
		t.Error(err.Error())
	}
	err = writer.Serialize(*a)
	if err != nil {
		t.Error(err.Error())
	}
	var aa interface{} = a
	err = writer.Serialize(aa)
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != "a3{123}a3{123}a3{123}" {
		t.Error(b.String())
	}
}

func TestSimpleWriterArray(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	err := writer.Serialize(a)
	if err != nil {
		t.Error(err.Error())
	}
	err = writer.Serialize(&a)
	if err != nil {
		t.Error(err.Error())
	}
	var aa interface{} = a
	err = writer.Serialize(&aa)
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != "a10{0123456789}a10{0123456789}a10{0123456789}" {
		t.Error(b.String())
	}
}

func TestSimpleWriterSlice(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	a := []int{0, 1, 2}
	err := writer.Serialize(a)
	if err != nil {
		t.Error(err.Error())
	}
	err = writer.Serialize(&a)
	if err != nil {
		t.Error(err.Error())
	}
	var aa interface{} = a
	err = writer.Serialize(&aa)
	if err != nil {
		t.Error(err.Error())
	}
	if b.String() != "a3{012}a3{012}a3{012}" {
		t.Error(b.String())
	}
}

func TestSimpleWriterMap(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	m := make(map[string]interface{})
	m["name"] = "马秉尧"
	m["age"] = 33
	m["male"] = true
	err := writer.Serialize(m)
	if err != nil {
		t.Error(err.Error())
	}
	err = writer.Serialize(&m)
	if err != nil {
		t.Error(err.Error())
	}
	var mm interface{} = m
	err = writer.Serialize(&mm)
	if err != nil {
		t.Error(err.Error())
	}
	s := `m3{s4"name"s3"马秉尧"s3"age"i33;s4"male"t}`
	if b.String() != strings.Repeat(s, 3) {
		t.Error(b.String())
	}
}

type testPerson struct {
	Name string
	Age  int
	Male bool
}

func TestSimpleWriterObject(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	p := testPerson{"马秉尧", 33, true}
	err := writer.Serialize(p)
	if err != nil {
		t.Error(err.Error())
	}
	err = writer.Serialize(&p)
	if err != nil {
		t.Error(err.Error())
	}
	var pp interface{} = p
	err = writer.Serialize(&pp)
	if err != nil {
		t.Error(err.Error())
	}
	s := `c10"testPerson"3{s4"name"s3"age"s4"male"}`
	s2 := `o0{s3"马秉尧"i33;t}`
	if b.String() != s+strings.Repeat(s2, 3) {
		t.Error(b.String())
	}
}
