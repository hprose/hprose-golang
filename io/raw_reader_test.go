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
 * io/raw_reader_test.go                                  *
 *                                                        *
 * hprose RawReader Test for Go.                          *
 *                                                        *
 * LastModified: Oct 29, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import (
	"math"
	"testing"
	"time"
)

func TestRawReaderTag(t *testing.T) {
	w := NewWriter(true)
	w.Serialize(nil)
	rawReader := NewRawReader(w.Bytes())
	raw := rawReader.ReadRaw()
	if string(raw) != "n" {
		t.Error("read tag error: ", string(raw))
	}
}

func TestRawReaderNumber(t *testing.T) {
	w := NewWriter(true)
	w.Serialize(123)
	rawReader := NewRawReader(w.Bytes())
	raw := rawReader.ReadRaw()
	if string(raw) != "i123;" {
		t.Error("read integer error: ", string(raw))
	}
}

func TestRawReaderInf(t *testing.T) {
	w := NewWriter(true)
	w.Serialize(math.Inf(1))
	rawReader := NewRawReader(w.Bytes())
	raw := rawReader.ReadRaw()
	if string(raw) != "I+" {
		t.Error("read inf error: ", string(raw))
	}
}

func TestRawReaderChar(t *testing.T) {
	w := NewWriter(true)
	w.Serialize("我")
	rawReader := NewRawReader(w.Bytes())
	raw := rawReader.ReadRaw()
	if string(raw) != `u我` {
		t.Error("read char error: ", string(raw))
	}
}

func TestRawReaderBytes(t *testing.T) {
	w := NewWriter(true)
	w.Serialize([]byte("hello world!"))
	rawReader := NewRawReader(w.Bytes())
	raw := rawReader.ReadRaw()
	if string(raw) != `b12"hello world!"` {
		t.Error("read bytes error: ", string(raw))
	}
}

func TestRawReaderString(t *testing.T) {
	w := NewWriter(true)
	w.Serialize("我爱你!")
	rawReader := NewRawReader(w.Bytes())
	raw := rawReader.ReadRaw()
	if string(raw) != `s4"我爱你!"` {
		t.Error("read string error: ", string(raw))
	}
}

func TestRawReaderGuid(t *testing.T) {
	rawReader := NewRawReader([]byte("g{AFA7F4B1-A64D-46FA-886F-ED7FBCE569B6}"))
	raw := rawReader.ReadRaw()
	if string(raw) != "g{AFA7F4B1-A64D-46FA-886F-ED7FBCE569B6}" {
		t.Error("read guid error: ", string(raw))
	}
}

func TestRawReaderDateTime(t *testing.T) {
	w := NewWriter(true)
	w.Serialize(time.Date(2008, 12, 11, 23, 12, 21, 123433453, time.UTC))
	rawReader := NewRawReader(w.Bytes())
	raw := rawReader.ReadRaw()
	if string(raw) != "D20081211T231221.123433453Z" {
		t.Error("read datetime error: ", string(raw))
	}
}

func TestRawReaderComplex(t *testing.T) {
	w := NewWriter(true)
	w.Serialize([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	rawReader := NewRawReader(w.Bytes())
	raw := rawReader.ReadRaw()
	if string(raw) != "a10{0123456789}" {
		t.Error("read complex error: ", string(raw))
	}
}

func TestRawReaderStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	w := NewWriter(true)
	w.Serialize(Person{"Tom", 18})
	rawReader := NewRawReader(w.Bytes())
	raw := rawReader.ReadRaw()
	if string(raw) != `c6"Person"2{s4"name"s3"age"}o0{s3"Tom"i18;}` {
		t.Error("read object error: ", string(raw))
	}
}

func TestRawReaderError(t *testing.T) {
	w := NewWriter(true)
	w.writeByte(TagError)
	w.Serialize("test error")
	rawReader := NewRawReader(w.Bytes())
	raw := rawReader.ReadRaw()
	if string(raw) != `Es10"test error"` {
		t.Error("read error: ", string(raw))
	}
}

func TestRawReaderUnexpectedTag(t *testing.T) {
	defer func() {
		if e := recover(); e == nil {
			t.Error("read unexpected tag error")
		}
	}()
	w := NewWriter(true)
	w.writeByte('*')
	rawReader := NewRawReader(w.Bytes())
	rawReader.ReadRaw()
}

func BenchmarkRawReaderReadUTF8StringEmpty(b *testing.B) {
	w := NewWriter(true)
	s := "我爱你🇨🇳"
	for i := 0; i < 20; i++ {
		s += s
	}
	w.Serialize(s)
	data := w.Bytes()
	test := func(_ []byte) {}
	for i := 0; i < b.N; i++ {
		test(data)
	}
}

func BenchmarkRawReaderReadUTF8String(b *testing.B) {
	w := NewWriter(true)
	s := "我爱你🇨🇳"
	for i := 0; i < 20; i++ {
		s += s
	}
	w.Serialize(s)
	data := w.Bytes()
	for i := 0; i < b.N; i++ {
		rawReader := NewRawReader(data)
		rawReader.ReadRaw()
	}
}
