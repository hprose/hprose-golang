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
 * hprose/raw_reader_test.go                              *
 *                                                        *
 * hprose RawReader Test for Go.                          *
 *                                                        *
 * LastModified: Jan 31, 2014                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose_test

import (
	"bytes"
	. "hprose"
	"testing"
)

func TestRawReaderTag(t *testing.T) {
	rawReader := NewRawReader(bytes.NewBufferString("e"))
	raw, err := rawReader.ReadRaw()
	if err != nil {
		t.Error(err.Error())
	}
	if string(raw) != "e" {
		t.Error("read tag error")
	}
}

func TestRawReaderNumber(t *testing.T) {
	rawReader := NewRawReader(bytes.NewBufferString("i123;"))
	raw, err := rawReader.ReadRaw()
	if err != nil {
		t.Error(err.Error())
	}
	if string(raw) != "i123;" {
		t.Error("read integer error: ", string(raw))
	}
}

func TestRawReaderChar(t *testing.T) {
	rawReader := NewRawReader(bytes.NewBufferString(`u我`))
	raw, err := rawReader.ReadRaw()
	if err != nil {
		t.Error(err.Error())
	}
	if string(raw) != `u我` {
		t.Error("read char error")
	}
}

func TestRawReaderBytes(t *testing.T) {
	rawReader := NewRawReader(bytes.NewBufferString(`b12"hello world!"`))
	raw, err := rawReader.ReadRaw()
	if err != nil {
		t.Error(err.Error())
	}
	if string(raw) != `b12"hello world!"` {
		t.Error("read bytes error")
	}
}

func TestRawReaderString(t *testing.T) {
	rawReader := NewRawReader(bytes.NewBufferString(`s4"我爱你!"`))
	raw, err := rawReader.ReadRaw()
	if err != nil {
		t.Error(err.Error())
	}
	if string(raw) != `s4"我爱你!"` {
		t.Error("read string error")
	}
}

func TestRawReaderGuid(t *testing.T) {
	rawReader := NewRawReader(bytes.NewBufferString("g{AFA7F4B1-A64D-46FA-886F-ED7FBCE569B6}"))
	raw, err := rawReader.ReadRaw()
	if err != nil {
		t.Error(err.Error())
	}
	if string(raw) != "g{AFA7F4B1-A64D-46FA-886F-ED7FBCE569B6}" {
		t.Error("read guid error")
	}
}

func TestRawReaderDateTime(t *testing.T) {
	rawReader := NewRawReader(bytes.NewBufferString("D20081211T231221.123433453Z"))
	raw, err := rawReader.ReadRaw()
	if err != nil {
		t.Error(err.Error())
	}
	if string(raw) != "D20081211T231221.123433453Z" {
		t.Error("read datetime error")
	}
}

func TestRawReaderComplex(t *testing.T) {
	rawReader := NewRawReader(bytes.NewBufferString("a10{0123456789}"))
	raw, err := rawReader.ReadRaw()
	if err != nil {
		t.Error(err.Error())
	}
	if string(raw) != "a10{0123456789}" {
		t.Error("read complex error")
	}
}

func TestRawReaderObject(t *testing.T) {
	rawReader := NewRawReader(bytes.NewBufferString(`c6"Person"2{s4"name"s3"age"}o0{s3"马秉尧"i33;}`))
	raw, err := rawReader.ReadRaw()
	if err != nil {
		t.Error(err.Error())
	}
	if string(raw) != `c6"Person"2{s4"name"s3"age"}o0{s3"马秉尧"i33;}` {
		t.Error("read object error")
	}
}
