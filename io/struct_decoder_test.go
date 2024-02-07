/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/struct_encoder_test.go                                |
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

func TestDecodeEmptyToStruct(t *testing.T) {
	type TestStruct struct {
		A int
		B bool    `hprose:"-"`
		C string  `json:"json,omitempty"`
		D float32 `json:",omitempty"`
	}

	sb := &strings.Builder{}
	enc := NewEncoder(sb)
	enc.Encode(nil)
	enc.Encode("")
	dec := NewDecoder(([]byte)(sb.String()))
	var ts *TestStruct
	dec.Decode(&ts)
	assert.Nil(t, ts)
	dec.Decode(&ts)
	assert.Equal(t, &TestStruct{}, ts)
}

func TestDecodeStruct(t *testing.T) {
	type TestStruct struct {
		A int
		B bool    `hprose:"-"`
		C string  `json:"json,omitempty"`
		D float32 `json:",omitempty"`
		e float64
	}

	src := TestStruct{1, true, "hello", 3.14, 2.718}
	sb := &strings.Builder{}
	enc := NewEncoder(sb)
	enc.Encode(src)
	dec := NewDecoder(([]byte)(sb.String()))
	var ts TestStruct
	dec.Decode(&ts)
	assert.Equal(t, TestStruct{1, false, "hello", 3.14, 0}, ts)
}

func TestDecodeStructPtr(t *testing.T) {
	type TestStruct struct {
		A int
		B bool    `hprose:"-"`
		C *string `json:"json,omitempty"`
		D float32 `json:",omitempty"`
		e float64
	}

	hello := "hello"
	src := TestStruct{1, true, &hello, 3.14, 2.718}
	sb := &strings.Builder{}
	enc := NewEncoder(sb)
	enc.Encode(src)
	dec := NewDecoder(([]byte)(sb.String()))
	var ts *TestStruct
	dec.Decode(&ts)
	assert.Equal(t, &TestStruct{1, false, &hello, 3.14, 0}, ts)
}

func TestDecodeMapAsObject(t *testing.T) {
	type TestStruct struct {
		A int
		B bool    `hprose:"-"`
		C *string `json:"json,omitempty"`
		D float32 `json:",omitempty"`
		e float64
	}

	hello := "hello"
	src := make(map[string]interface{})
	src["a"] = 1
	src["b"] = true
	src["c"] = "c"
	src["json"] = "hello"
	src["d"] = 3.14
	src["e"] = 2.178
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	enc.Encode(src)
	dec := NewDecoder(([]byte)(sb.String())).Simple(false)
	var ts *TestStruct
	dec.Decode(&ts)
	assert.Equal(t, &TestStruct{1, false, &hello, 3.14, 0}, ts)
}

func TestDecodeAnonymousStruct(t *testing.T) {
	src := struct {
		A int
		B bool    `hprose:"-"`
		C string  `json:"json,omitempty"`
		D float32 `json:",omitempty"`
		e float64
	}{1, true, "hello", 3.14, 2.718}
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	enc.Encode(src)
	dec := NewDecoder(([]byte)(sb.String())).Simple(false)
	var ts struct {
		A int
		B bool    `hprose:"-"`
		C string  `json:"json,omitempty"`
		D float32 `json:",omitempty"`
		e float64
	}
	dec.Decode(&ts)
	assert.Equal(t, struct {
		A int
		B bool    `hprose:"-"`
		C string  `json:"json,omitempty"`
		D float32 `json:",omitempty"`
		e float64
	}{1, false, "hello", 3.14, 0}, ts)
}

func TestDecodeStructWithDBTag(t *testing.T) {
	type User struct {
		UserID   int    `db:"id"`
		UserName string `db:"name"`
	}
	Register(User{}, "db")
	src := User{1, "张三"}
	sb := &strings.Builder{}
	enc := NewEncoder(sb)
	enc.Encode(src)
	dec := NewDecoder(([]byte)(sb.String()))
	var ts User
	dec.Decode(&ts)
	assert.Equal(t, src, ts)
}

func TestDecodeStructAsInterface(t *testing.T) {
	type TestStruct struct {
		A int
		B bool    `hprose:"-"`
		C string  `json:"json,omitempty"`
		D float32 `json:",omitempty"`
		e float64
	}
	Register((*TestStruct)(nil))
	src := TestStruct{1, true, "hello", 3.14, 2.718}
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	enc.Encode(src)
	dec := NewDecoder(([]byte)(sb.String())).Simple(false)
	var ts interface{}
	dec.Decode(&ts)
	assert.Equal(t, &TestStruct{1, false, "hello", 3.14, 0}, ts)
	dec = NewDecoder(([]byte)(sb.String())).Simple(false)
	dec.StructType = StructTypeStructObject
	dec.Decode(&ts)
	assert.Equal(t, TestStruct{1, false, "hello", 3.14, 0}, ts)
}

func TestDecodeSelfReferenceStructAsInterface(t *testing.T) {
	type TestStruct struct {
		A *TestStruct
	}
	Register((*TestStruct)(nil))
	src := TestStruct{}
	src.A = &src
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	enc.Encode(src)
	dec := NewDecoder(([]byte)(sb.String())).Simple(false)
	var ts interface{}
	dec.Decode(&ts)
	assert.Equal(t, &src, ts)
	dec = NewDecoder(([]byte)(sb.String())).Simple(false)
	dec.StructType = StructTypeStructObject
	dec.Decode(&ts)
	assert.Equal(t, src, ts)
}
