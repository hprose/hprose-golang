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
 * hprose/writer_test.go                                  *
 *                                                        *
 * hprose Writer Test for Go.                             *
 *                                                        *
 * LastModified: Feb 14, 2014                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose_test

import (
	"bytes"
	"container/list"
	. "github.com/hprose/hprose-go/hprose"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestReaderTime(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, false)
	writer.Serialize(time.Date(2014, 1, 19, 20, 25, 33, 12345678, time.UTC))
	writer.Serialize(time.Date(2014, 1, 19, 20, 25, 33, 12345678, time.UTC))
	writer.Serialize(time.Date(2014, 1, 19, 0, 0, 0, 0, time.Local))
	writer.Serialize(time.Date(1970, 1, 1, 1, 1, 1, 0, time.Local))
	if b.String() != "D20140119T202533.012345678Zr0;D20140119;T010101;" {
		t.Error(b.String())
	}
	reader := NewReader(b, false)
	var t1, t2, t3, t4 time.Time
	if err := reader.Unserialize(&t1); err != nil {
		t.Error(err.Error())
	}
	if err := reader.Unserialize(&t2); err != nil {
		t.Error(err.Error())
	}
	if err := reader.Unserialize(&t3); err != nil {
		t.Error(err.Error())
	}
	if err := reader.Unserialize(&t4); err != nil {
		t.Error(err.Error())
	}
	if t1 != t2 {
		t.Error("t1 != t2")
	}
}

func TestReaderString(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, false)
	writer.Serialize("")
	writer.Serialize("我爱你")
	writer.Serialize("我爱你")

	if b.String() != `es3"我爱你"r0;` {
		t.Error(b.String())
	}
	reader := NewReader(b, false)
	var s1, s2, s3 string
	if err := reader.Unserialize(&s1); err != nil {
		t.Error(err.Error())
	}
	if err := reader.Unserialize(&s2); err != nil {
		t.Error(err.Error())
	}
	if err := reader.Unserialize(&s3); err != nil {
		t.Error(err.Error())
	}
	if s2 != "我爱你" {
		t.Error(s2)
	}
	if s2 != s3 {
		t.Error("s2 != s3")
	}
}

func TestReaderBytes(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, false)
	writer.Serialize([]byte(""))
	bb := []byte("我爱你")
	writer.Serialize(&bb)
	writer.Serialize(&bb)
	if b.String() != `b""b9"我爱你"r1;` {
		t.Error(b.String())
	}
	reader := NewReader(b, false)
	var x1, x2, x3 *[]byte
	if err := reader.Unserialize(&x1); err != nil {
		t.Error(err.Error())
	}
	if err := reader.Unserialize(&x2); err != nil {
		t.Error(err.Error())
	}
	if err := reader.Unserialize(&x3); err != nil {
		t.Error(err.Error())
	}
	if x2 != x3 {
		t.Error("x2 != x3")
	}
}

func TestReaderUUID(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, false)
	u := ToUUID("3f257da1-0b85-48d6-8f5c-6cd13d2d60c9")
	writer.Serialize(&u)
	writer.Serialize(&u)
	writer.Serialize(&u)
	if b.String() != "g{3f257da1-0b85-48d6-8f5c-6cd13d2d60c9}r0;r0;" {
		t.Error(b.String())
	}
	var u2, u3 *UUID
	reader := NewReader(b, false)
	if err := reader.Unserialize(&u); err != nil {
		t.Error(err.Error())
	}
	if err := reader.Unserialize(&u2); err != nil {
		t.Error(err.Error())
	}
	if err := reader.Unserialize(&u3); err != nil {
		t.Error(err.Error())
	}
	if u2 != u3 {
		t.Error(u, u2, u3)
	}
}

func TestReaderList(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, false)
	a := list.New()
	a.PushBack("hello")
	a.PushBack("hprose")
	a.PushBack("world")
	writer.Serialize(a)
	writer.Serialize(*a)
	var aa interface{} = a
	writer.Serialize(aa)
	if b.String() != `a3{s5"hello"s6"hprose"s5"world"}a3{r1;r2;r3;}r0;` {
		t.Error(b.String())
	}
	reader := NewReader(b, false)
	var x1, x2, x3 *list.List
	if err := reader.Unserialize(&x1); err != nil {
		t.Error(err.Error())
	}
	if err := reader.Unserialize(&x2); err != nil {
		t.Error(err.Error())
	}
	if err := reader.Unserialize(&x3); err != nil {
		t.Error(err.Error())
	}
	if x1 == x2 {
		t.Error("x1 == x2")
	}
	if x1 != x3 {
		t.Error("x1 != x3")
	}
}

func TestReaderSlice(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, false)
	a := []string{"hello", "hprose", "world"}
	writer.Serialize(&a)
	writer.Serialize(a)
	var aa interface{} = &a
	writer.Serialize(aa)
	if b.String() != `a3{s5"hello"s6"hprose"s5"world"}a3{r1;r2;r3;}r0;` {
		t.Error(b.String())
	}
	var x []string
	reader := NewReader(b, false)
	if err := reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(x, a) {
		t.Error(x, a)
	}
	if err := reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(x, a) {
		t.Error(x, a)
	}
	if err := reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(x, a) {
		t.Error(x, a)
	}
}

func TestReaderMap(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, false)
	m := make(map[string]interface{})
	m["name"] = "马秉尧"
	m["age"] = 33
	m["male"] = true
	writer.Serialize(m)
	writer.Serialize(&m)
	var mm interface{} = &m
	writer.Serialize(mm)
	s := `m3{s4"name"s3"马秉尧"s3"age"i33;s4"male"t}m3{r1;r2;r3;i33;r4;t}r5;`
	if b.String() != s {
		t.Error(b.String())
	}
	reader := NewReader(b, false)
	var x map[string]interface{}
	if err := reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(x, m) {
		t.Error(x, m)
	}
	if err := reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(x, m) {
		t.Error(x, m)
	}
	if err := reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(x, m) {
		t.Error(x, m)
	}
}

func TestReaderObject(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, false)
	p := testPerson{"马秉尧", 33, true}
	writer.Serialize(p)
	writer.Serialize(&p)
	writer.Serialize(&p)
	var pp interface{} = &p
	writer.Serialize(pp)
	s := `c10"testPerson"3{s4"name"s3"age"s4"male"}o0{s3"马秉尧"i33;t}o0{r4;i33;t}r5;r5;`
	if b.String() != s {
		t.Error(b.String())
	}
	reader := NewReader(b, false)
	var x testPerson
	if err := reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(x, p) {
		t.Error(x, p)
	}
	if err := reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(x, p) {
		t.Error(x, p)
	}
	if err := reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(x, p) {
		t.Error(x, p)
	}
	if err := reader.Unserialize(&x); err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(x, p) {
		t.Error(x, p)
	}
}

func TestReaderReset(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, false)
	p := testPerson{"马秉尧", 33, true}
	writer.Serialize(p)
	writer.Reset()
	writer.Serialize(&p)
	writer.Reset()
	writer.Serialize(&p)
	writer.Reset()
	var pp interface{} = &p
	writer.Serialize(pp)
	s := strings.Repeat(`c10"testPerson"3{s4"name"s3"age"s4"male"}o0{s3"马秉尧"i33;t}`, 4)
	if b.String() != s {
		t.Error(b.String())
	}
	reader := NewReader(b, false)
	reader.Unserialize(&p)
	reader.Reset()
	reader.Unserialize(&pp)
	reader.Reset()
	reader.Unserialize(&pp)
	reader.Reset()
	reader.Unserialize(&p)
}

func TestReaderArray(t *testing.T) {
	b := new(bytes.Buffer)
	writer := NewWriter(b, false)
	p := testPerson{"马秉尧", 33, true}
	a := []interface{}{123, "hello world!", p}
	writer.Serialize(a)
	reader := NewReader(b, false)
	err := reader.CheckTag(TagList)
	if err != nil {
		t.Error(err.Error())
	}
	count, err := reader.ReadInteger(TagOpenbrace)
	if err != nil {
		t.Error(err.Error())
	}
	if count != 3 {
		t.Error("count != 3")
	}
	result := [3]reflect.Value{
		reflect.New(reflect.TypeOf(0)).Elem(),
		reflect.New(reflect.TypeOf("")).Elem(),
		reflect.New(reflect.TypeOf(testPerson{})).Elem(),
	}
	err = reader.ReadArray(result[:])
	if err != nil {
		t.Error(err.Error())
	}
	if result[0].Int() != 123 {
		t.Error(result)
	}
	if result[1].String() != "hello world!" {
		t.Error(result)
	}
	if result[2].Interface().(testPerson) != p {
		t.Error(result)
	}
}
