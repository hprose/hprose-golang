/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/struct_encoder_test.go                          |
|                                                          |
| LastModified: Jun 25, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeEmptyToStruct(t *testing.T) {
	type TestStruct struct {
		A int
		B bool    `hprose:"-"`
		C string  `json:"json,omitempty"`
		D float32 `json:",omitempty"`
		e float64
	}

	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
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
	enc := NewEncoder(sb).Simple(false)
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
	enc := NewEncoder(sb).Simple(false)
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
	dec := NewDecoder(([]byte)(sb.String()))
	var ts *TestStruct
	dec.Decode(&ts)
	assert.Equal(t, &TestStruct{1, false, &hello, 3.14, 0}, ts)
}
