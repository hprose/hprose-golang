/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/client_codec_test.go                            |
|                                                          |
| LastModified: Apr 27, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core_test

import (
	"reflect"
	"testing"

	. "github.com/hprose/hprose-golang/v3/rpc/core"
	"github.com/stretchr/testify/assert"
)

func TestClientCodecEncode(t *testing.T) {
	context := NewClientContext()
	context.RequestHeaders().Set("id", "test_id")
	result, err := NewClientCodec().Encode("hello", []interface{}{"World"}, context)
	assert.Equal(t, `Hm1{s2"id"s7"test_id"}Cs5"hello"a1{s5"World"}z`, string(result))
	assert.NoError(t, err)
}

func TestClientCodecDecode(t *testing.T) {
	context := NewClientContext()
	response := ([]byte)(`Hm1{s2"id"s7"test_id"}Rs12"hello World!"z`)
	result, err := NewClientCodec().Decode(response, context)
	id, ok := context.ResponseHeaders().Get("id")
	assert.Nil(t, result)
	assert.Equal(t, id, "test_id")
	assert.Equal(t, ok, true)
	assert.NoError(t, err)

	context.ReturnType = []reflect.Type{reflect.TypeOf("")}
	response = ([]byte)(`Hm1{s2"id"s7"test_id"}Rs12"hello World!"z`)
	result, err = NewClientCodec().Decode(response, context)
	id, ok = context.ResponseHeaders().Get("id")
	assert.Equal(t, `hello World!`, result[0])
	assert.Equal(t, id, "test_id")
	assert.Equal(t, ok, true)
	assert.NoError(t, err)

	context.ReturnType = []reflect.Type{reflect.TypeOf(""), reflect.TypeOf(true)}
	response = ([]byte)(`Hm1{s2"id"s7"test_id"}Rs12"hello World!"z`)
	result, err = NewClientCodec().Decode(response, context)
	id, ok = context.ResponseHeaders().Get("id")
	assert.Equal(t, `hello World!`, result[0])
	assert.Equal(t, false, result[1])
	assert.Equal(t, id, "test_id")
	assert.Equal(t, ok, true)
	assert.NoError(t, err)

	context.ReturnType = []reflect.Type{reflect.TypeOf(""), reflect.TypeOf(true)}
	response = ([]byte)(`Hm1{s2"id"s7"test_id"}Ra2{s12"hello World!"t}z`)
	result, err = NewClientCodec().Decode(response, context)
	id, ok = context.ResponseHeaders().Get("id")
	assert.Equal(t, `hello World!`, result[0])
	assert.Equal(t, true, result[1])
	assert.Equal(t, id, "test_id")
	assert.Equal(t, ok, true)
	assert.NoError(t, err)

	context.ReturnType = []reflect.Type{reflect.TypeOf("")}
	response = ([]byte)(`Hm1{s2"id"s7"test_id"}Es12"hello World!"z`)
	result, err = NewClientCodec().Decode(response, context)
	id, ok = context.ResponseHeaders().Get("id")
	assert.Nil(t, result)
	assert.Equal(t, id, "test_id")
	assert.Equal(t, ok, true)
	assert.EqualError(t, err, "hello World!")

	context.ReturnType = []reflect.Type{reflect.TypeOf(""), reflect.TypeOf(true), reflect.TypeOf(0)}
	response = ([]byte)(`Hm1{s2"id"s7"test_id"}Ra2{s12"hello World!"t}z`)
	result, err = NewClientCodec().Decode(response, context)
	id, ok = context.ResponseHeaders().Get("id")
	assert.Equal(t, `hello World!`, result[0])
	assert.Equal(t, true, result[1])
	assert.Equal(t, 0, result[2])
	assert.Equal(t, id, "test_id")
	assert.Equal(t, ok, true)
	assert.NoError(t, err)

	context.ReturnType = []reflect.Type{reflect.TypeOf(""), reflect.TypeOf(true), reflect.TypeOf(0)}
	response = ([]byte)(`{code:200,msg:"ok"}`)
	result, err = NewClientCodec().Decode(response, context)
	assert.Nil(t, result)
	assert.EqualError(t, err, "hprose/rpc/core: invalid response:\r\n"+`{code:200,msg:"ok"}`)
}
