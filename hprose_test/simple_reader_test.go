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
 * hprose/simple_reader_test.go                           *
 *                                                        *
 * hprose SimpleReader Test for Go.                       *
 *                                                        *
 * LastModified: Feb 15, 2014                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose_test

import (
	"bytes"
	"container/list"
	. "github.com/hprose/hprose-go/hprose"
	"math"
	"math/big"
	"reflect"
	"testing"
	"time"
)

func TestSimpleReaderInt(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	writer.Serialize(-1)
	writer.Serialize(math.MaxInt16)
	writer.Serialize(math.MinInt64)
	var f32 float32 = math.MaxFloat32
	writer.Serialize(f32)
	writer.Serialize("1234567890")
	writer.Serialize("1")
	writer.Serialize("")
	writer.Serialize(nil)
	writer.Serialize(false)
	writer.Serialize(true)
	writer.Serialize(nil)
	reader := NewReader(b, true)
	var i int
	var err error
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != -1 {
		t.Error(i)
	}
	var p *int16
	if err = reader.Unserialize(&p); err != nil {
		t.Error(err.Error())
	}
	if *p != math.MaxInt16 {
		t.Error(*p)
	}
	var i64 int64
	if err = reader.Unserialize(&i64); err != nil {
		t.Error(err.Error())
	}
	if i64 != math.MinInt64 {
		t.Error(i64)
	}
	if err = reader.Unserialize(&i64); err != nil {
		t.Error(err.Error())
	}
	if i64 != int64(f32) {
		t.Error(i64)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 1234567890 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 49 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 0 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 0 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 0 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 1 {
		t.Error(i)
	}
	if err = reader.Unserialize(&p); err != nil {
		t.Error(err.Error())
	}
	if p != nil {
		t.Error(p)
	}
}

func TestSimpleReaderUint(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	writer.Serialize(-1)
	writer.Serialize(math.MaxUint32)
	writer.Serialize(math.MaxUint32)
	writer.Serialize(math.MaxUint32)
	writer.Serialize(math.MaxUint32)
	var f32 float32 = math.MaxFloat32
	writer.Serialize(f32)
	writer.Serialize("1234567890")
	writer.Serialize("1")
	writer.Serialize("")
	writer.Serialize(nil)
	writer.Serialize(false)
	writer.Serialize(true)
	reader := NewReader(b, true)
	var i uint
	var err error
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != math.MaxUint64 {
		t.Error(i)
	}
	var p1 *int8
	if err = reader.Unserialize(&p1); err != nil {
		t.Error(err.Error())
	}
	if *p1 != -1 {
		t.Error(*p1)
	}
	var p2 *int16
	if err = reader.Unserialize(&p2); err != nil {
		t.Error(err.Error())
	}
	if *p2 != -1 {
		t.Error(*p2)
	}
	var p3 *int32
	if err = reader.Unserialize(&p3); err != nil {
		t.Error(err.Error())
	}
	if *p3 != -1 {
		t.Error(*p3)
	}
	var p4 *int64
	if err = reader.Unserialize(&p4); err != nil {
		t.Error(err.Error())
	}
	if *p4 != math.MaxUint32 {
		t.Error(*p4)
	}
	var f uint64
	if err = reader.Unserialize(&f); err != nil {
		t.Error(err.Error())
	}
	if f != uint64(f32) {
		t.Error(f)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 1234567890 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 49 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 0 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 0 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 0 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 1 {
		t.Error(i)
	}
}

func TestSimpleReaderBool(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	writer.Serialize(-1)
	writer.Serialize(0)
	writer.Serialize(1)
	writer.Serialize(math.MaxUint32)
	writer.Serialize(math.MinInt64)
	var f32 float32 = math.MaxFloat32
	writer.Serialize(f32)
	writer.Serialize("false")
	writer.Serialize("1")
	writer.Serialize("t")
	writer.Serialize("0")
	writer.Serialize("f")
	writer.Serialize("")
	writer.Serialize(false)
	writer.Serialize(true)
	writer.Serialize(nil)
	writer.Serialize(nil)
	reader := NewReader(b, true)
	var i bool
	var p *bool = &i
	var err error
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != true {
		t.Error(i)
	}
	if err = reader.Unserialize(&p); err != nil {
		t.Error(err.Error())
	}
	if *p != false {
		t.Error(*p)
	}
	if i != true {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != true {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != true {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != true {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != true {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != false {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != true {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != true {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != false {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != false {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != false {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != false {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != true {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != false {
		t.Error(i)
	}
	if err = reader.Unserialize(&p); err != nil {
		t.Error(err.Error())
	}
	if p != nil {
		t.Error(p)
	}
}

func TestSimpleReaderBigInt(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	writer.Serialize(-1)
	writer.Serialize(math.MaxInt16)
	var f32 float32 = math.MaxFloat32
	writer.Serialize(f32)
	writer.Serialize("1234567890")
	writer.Serialize("1")
	writer.Serialize("")
	writer.Serialize(nil)
	writer.Serialize(false)
	writer.Serialize(true)
	reader := NewReader(b, true)
	var i big.Int
	var err error
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i.Cmp(big.NewInt(-1)) != 0 {
		t.Error(i)
	}
	var p *big.Int
	if err = reader.Unserialize(&p); err != nil {
		t.Error(err.Error())
	}
	if (*p).Cmp(big.NewInt(math.MaxInt16)) != 0 {
		t.Error(*p)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i.Cmp(big.NewInt(int64(f32))) != 0 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i.Cmp(big.NewInt(1234567890)) != 0 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i.Cmp(big.NewInt(49)) != 0 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i.Cmp(big.NewInt(0)) != 0 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i.Cmp(big.NewInt(0)) != 0 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i.Cmp(big.NewInt(0)) != 0 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i.Cmp(big.NewInt(1)) != 0 {
		t.Error(i)
	}
}

func TestSimpleReaderFloat32(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	writer.Serialize(-1)
	writer.Serialize(math.MaxUint32)
	var f32 float32 = math.MaxFloat32
	writer.Serialize(f32)
	writer.Serialize("1234567890")
	writer.Serialize("1")
	writer.Serialize("")
	writer.Serialize(nil)
	writer.Serialize(false)
	writer.Serialize(true)
	reader := NewReader(b, true)
	var i float32
	var err error
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != float32(-1) {
		t.Error(i)
	}
	var p *float32
	if err = reader.Unserialize(&p); err != nil {
		t.Error(err.Error())
	}
	if *p != float32(math.MaxUint32) {
		t.Error(*p)
		t.Error(float32(math.MaxUint32))
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != f32 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 1234567890 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 49 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 0 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 0 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 0 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 1 {
		t.Error(i)
	}
}

func TestSimpleReaderFloat64(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	writer.Serialize(-1)
	writer.Serialize(math.MaxUint32)
	var f64 float64 = math.MaxFloat64
	writer.Serialize(f64)
	writer.Serialize("1234567890")
	writer.Serialize("1")
	writer.Serialize("")
	writer.Serialize(nil)
	writer.Serialize(false)
	writer.Serialize(true)
	reader := NewReader(b, true)
	var i float64
	var err error
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != float64(-1) {
		t.Error(i)
	}
	var p *float64
	if err = reader.Unserialize(&p); err != nil {
		t.Error(err.Error())
	}
	if *p != float64(math.MaxUint32) {
		t.Error(*p)
		t.Error(float64(math.MaxUint32))
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != f64 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 1234567890 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 49 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 0 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 0 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 0 {
		t.Error(i)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != 1 {
		t.Error(i)
	}
}

func TestSimpleReaderTime(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	now := time.Now()
	writer.Serialize(now)
	writer.Serialize(now)
	datetime := time.Date(2014, 1, 21, 12, 13, 14, 0, time.Local)
	writer.Serialize(datetime)
	date := time.Date(2014, 1, 21, 0, 0, 0, 0, time.Local)
	writer.Serialize(date)
	tim := time.Date(1, 1, 1, 19, 23, 19, 123000, time.UTC)
	writer.Serialize(tim)
	// go 1.1 has a bug, 1.2 fixed
	// writer.Serialize(datetime.String())
	reader := NewReader(b, true)
	var x time.Time
	var err error
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if x != now {
		t.Error(x)
	}
	var p *time.Time
	if err = reader.Unserialize(&p); err != nil {
		t.Error(err.Error())
	}
	if *p != now {
		t.Error(*p)
		t.Error(now)
	}
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if x != datetime {
		t.Error(x)
	}
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if x != date {
		t.Error(x)
	}
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if x != tim {
		t.Error(x)
	}
	// go 1.1 has a bug, 1.2 fixed
	/*
		if err = reader.Unserialize(&x); err != nil {
			t.Error(err.Error())
		}
		if x != datetime {
			t.Error(x)
			t.Error(datetime)
		}
	*/
}

func TestSimpleReaderString(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	str := "hello 你好"
	writer.Serialize(str)
	writer.Serialize([]byte(str))
	datetime := time.Date(2014, 1, 21, 12, 13, 14, 0, time.Local)
	writer.Serialize(datetime)
	writer.Serialize(1234567890)
	writer.Serialize(0)
	writer.Serialize(math.NaN())
	writer.Serialize(math.Inf(1))
	writer.Serialize(true)
	writer.Serialize("")
	writer.Serialize(nil)
	writer.Serialize(nil)

	reader := NewReader(b, true)
	var x string
	var err error
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if x != str {
		t.Error(x)
	}
	var p *string
	if err = reader.Unserialize(&p); err != nil {
		t.Error(err.Error())
	}
	if *p != str {
		t.Error(*p)
		t.Error(str)
	}
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if x != datetime.String() {
		t.Error(x)
	}
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if x != "1234567890" {
		t.Error(x)
	}
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if x != "0" {
		t.Error(x)
	}
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if x != "NaN" {
		t.Error(x)
	}
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if x != "+Inf" {
		t.Error(x)
	}
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if x != "true" {
		t.Error(x)
	}
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if x != "" {
		t.Error(x)
	}
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if x != "" {
		t.Error(x)
	}
	if err = reader.Unserialize(&p); err != nil {
		t.Error(err.Error())
	}
	if p != nil {
		t.Error(*p)
		t.Error(str)
	}
}

func TestSimpleReaderBytes(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	str := "hello 你好"
	writer.Serialize(str)
	writer.Serialize([]byte(str))
	writer.Serialize("")
	writer.Serialize("")
	writer.Serialize(nil)
	writer.Serialize(nil)

	reader := NewReader(b, true)
	var x []byte
	var p *[]byte
	var err error
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if string(x) != str {
		t.Error(x)
	}
	if err = reader.Unserialize(&p); err != nil {
		t.Error(err.Error())
	}
	if string(*p) != str {
		t.Error(*p)
		t.Error(str)
	}
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if string(x) != "" {
		t.Error(x)
	}
	if err = reader.Unserialize(&p); err != nil {
		t.Error(err.Error())
	}
	if string(*p) != "" {
		t.Error(*p)
	}
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if string(x) != "" {
		t.Error(x)
	}
	if err = reader.Unserialize(&p); err != nil {
		t.Error(err.Error())
	}
	if p != nil {
		t.Error(*p)
	}
}

func TestSimpleReaderUUID(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	u := ToUUID("3f257da1-0b85-48d6-8f5c-6cd13d2d60c9")
	writer.Serialize(u)
	writer.Serialize(u)
	writer.Serialize(u.String())
	writer.Serialize([]byte(u))
	writer.Serialize(nil)
	writer.Serialize(nil)

	reader := NewReader(b, true)
	var x UUID
	var p *UUID
	var err error
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if x.String() != u.String() {
		t.Error(x)
	}
	if err = reader.Unserialize(&p); err != nil {
		t.Error(err.Error())
	}
	if (*p).String() != u.String() {
		t.Error(*p)
	}
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if x.String() != u.String() {
		t.Error(x)
	}
	if err = reader.Unserialize(&p); err != nil {
		t.Error(err.Error())
	}
	if (*p).String() != u.String() {
		t.Error(*p)
	}
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if x.String() != "" {
		t.Error(x)
	}
	if err = reader.Unserialize(&p); err != nil {
		t.Error(err.Error())
	}
	if p != nil {
		t.Error(*p)
	}
}

func TestSimpleReaderList(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	l := list.New()
	l.PushBack(1)
	l.PushBack(2.0)
	l.PushBack(true)
	l.PushBack(false)
	l.PushBack(nil)
	l.PushBack(math.Inf(1))
	l.PushBack(math.Inf(-1))
	l.PushBack("")
	l.PushBack("我")
	l.PushBack("Hello World")
	l.PushBack("你好")
	writer.Serialize(l)
	writer.Serialize(l)
	writer.Serialize(nil)
	writer.Serialize(nil)

	reader := NewReader(b, true)
	var x list.List
	var p *list.List
	var err error
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	e2 := l.Front()
	for e := x.Front(); e != nil; e = e.Next() {
		if e.Value != e2.Value {
			t.Error(e.Value, e2.Value)
		}
		e2 = e2.Next()
	}

	if err = reader.Unserialize(&p); err != nil {
		t.Error(err.Error())
	}
	e2 = l.Front()
	for e := p.Front(); e != nil; e = e.Next() {
		if e.Value != e2.Value {
			t.Error(e.Value, e2.Value)
		}
		e2 = e2.Next()
	}

	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if x.Len() != 0 {
		t.Error(x)
	}
	if err = reader.Unserialize(&p); err != nil {
		t.Error(err.Error())
	}
	if p != nil {
		t.Error(*p)
	}
}

func TestSimpleReaderSlice(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	a := []interface{}{1, 2.0, true, false, math.Inf(1), math.Inf(-1), "", "我", "Hello World", "你好"}
	ia := []int{1, 2, 3, 4, 5, 6}
	aa := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	writer.Serialize(a)
	writer.Serialize(a)
	writer.Serialize(a)
	writer.Serialize(a)
	writer.Serialize(ia)
	writer.Serialize(aa)
	writer.Serialize(nil)
	writer.Serialize(nil)
	writer.Serialize(nil)
	writer.Serialize(nil)

	reader := NewReader(b, true)
	var x []interface{}
	var p *[]interface{}
	var i interface{}
	var pi *interface{}
	var pia *[]int
	var paa *[][]int
	var err error
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(x, a) {
		t.Error(x)
		t.Error(a)
	}
	if err = reader.Unserialize(&p); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(*p, a) {
		t.Error(*p)
		t.Error(a)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(*i.(*[]interface{}), a) {
		t.Error(*i.(*[]interface{}))
		t.Error(a)
	}
	if err = reader.Unserialize(&pi); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(*(*pi).(*[]interface{}), a) {
		t.Error(*(*pi).(*[]interface{}))
		t.Error(a)
	}
	if err = reader.Unserialize(&pia); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(*pia, ia) {
		t.Error(*pia)
		t.Error(ia)
	}
	if err = reader.Unserialize(&paa); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(*paa, aa) {
		t.Error(*paa)
		t.Error(aa)
	}
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if len(x) != 0 {
		t.Error(x)
	}
	if err = reader.Unserialize(&p); err != nil {
		t.Error(err.Error())
	}
	if p != nil {
		t.Error(*p)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if i != nil {
		t.Error(i)
	}
	if err = reader.Unserialize(&pi); err != nil {
		t.Error(err.Error())
	}
	if *pi != nil {
		t.Error(pi)
	}
}

func TestSimpleReaderMap(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, true)
	m := make(map[interface{}]interface{})
	m["name"] = "张山"
	m["age"] = 38
	m["male"] = false
	writer.Serialize(m)
	writer.Serialize(m)
	writer.Serialize(m)
	writer.Serialize(m)
	writer.Serialize(m)
	writer.Serialize(m)
	writer.Serialize(nil)
	writer.Serialize(nil)
	writer.Serialize(nil)
	writer.Serialize(nil)

	reader := NewReader(b, true)
	var x map[interface{}]interface{}
	var p *map[interface{}]interface{}
	var i interface{}
	var pi *interface{}
	var tp testPerson
	var ptp *testPerson
	var err error
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(x, m) {
		t.Error(x)
		t.Error(m)
	}
	if err = reader.Unserialize(&p); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(*p, m) {
		t.Error(*p)
		t.Error(m)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(*i.(*map[interface{}]interface{}), m) {
		t.Error(*i.(*map[interface{}]interface{}))
		t.Error(m)
	}
	if err = reader.Unserialize(&pi); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(*(*pi).(*map[interface{}]interface{}), m) {
		t.Error(*(*pi).(*map[interface{}]interface{}))
		t.Error(m)
	}
	if err = reader.Unserialize(&tp); err != nil {
		t.Error(err.Error())
	}
	if m["name"] != tp.Name || m["age"] != tp.Age || m["male"] != tp.Male {
		t.Error(m)
		t.Error(tp)
	}
	if err = reader.Unserialize(&ptp); err != nil {
		t.Error(err.Error())
	}
	if m["name"] != ptp.Name || m["age"] != ptp.Age || m["male"] != ptp.Male {
		t.Error(m)
		t.Error(ptp)
	}
}

type emptyInterface interface{}

func TestSimpleReaderObject(t *testing.T) {
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
	err = writer.Serialize(&pp)
	if err != nil {
		t.Error(err.Error())
	}
	err = writer.Serialize(p)
	if err != nil {
		t.Error(err.Error())
	}
	err = writer.Serialize(p)
	if err != nil {
		t.Error(err.Error())
	}

	reader := NewReader(b, true)
	var x testPerson
	var px *testPerson
	var i interface{}
	var pi *emptyInterface
	var m map[string]interface{}
	var pm *map[string]interface{}
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(x, p) {
		t.Error(x)
		t.Error(p)
	}
	if err = reader.Unserialize(&px); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(*px, p) {
		t.Error(*px)
		t.Error(p)
	}
	if err = reader.Unserialize(&i); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(*i.(*testPerson), p) {
		t.Error(*i.(*testPerson))
		t.Error(p)
	}
	if err = reader.Unserialize(&pi); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(*(*pi).(*testPerson), p) {
		t.Error(*(*pi).(*testPerson))
		t.Error(p)
	}
	if err = reader.Unserialize(&m); err != nil {
		t.Error(err.Error())
	}
	if m["name"] != p.Name || m["age"] != p.Age || m["male"] != p.Male {
		t.Error(m)
		t.Error(p)
	}
	if err = reader.Unserialize(&pm); err != nil {
		t.Error(err.Error())
	}
	if (*pm)["name"] != p.Name || (*pm)["age"] != p.Age || (*pm)["male"] != p.Male {
		t.Error(*pm)
		t.Error(p)
	}

	b = bytes.NewBufferString(`c10"testPerson"4{s4"name"s3"age"s4"male"s2"qq"}o0{s3"马秉尧"i33;ts8"53958317"}`)
	reader = NewReader(b, true)
	if err = reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(x, p) {
		t.Error(x)
		t.Error(p)
	}

}
