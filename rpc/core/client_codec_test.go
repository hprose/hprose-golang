/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/client_codec_test.go                            |
|                                                          |
| LastModified: Jan 24, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientCodecEncode(t *testing.T) {
	context := NewClientContext()
	context.RequestHeaders().Set("id", "test_id")
	result := string(DefaultClientCodec.Encode("hello", []interface{}{"World"}, context))
	assert.Equal(t, `Hm1{s2"id"s7"test_id"}Cs5"hello"a1{s5"World"}z`, result)
}

func TestClientCodecDecode(t *testing.T) {
	context := NewClientContext()
	response := ([]byte)(`Hm1{s2"id"s7"test_id"}Rs12"hello World!"z`)
	result, err := DefaultClientCodec.Decode(response, context)
	id, ok := context.ResponseHeaders().Get("id")
	assert.Equal(t, `hello World!`, result)
	assert.Equal(t, id, "test_id")
	assert.Equal(t, ok, true)
	assert.NoError(t, err)

	context.SetReturnType([]reflect.Type{reflect.TypeOf("")})
	response = ([]byte)(`Hm1{s2"id"s7"test_id"}Rs12"hello World!"z`)
	result, err = DefaultClientCodec.Decode(response, context)
	id, ok = context.ResponseHeaders().Get("id")
	assert.Equal(t, `hello World!`, result)
	assert.Equal(t, id, "test_id")
	assert.Equal(t, ok, true)
	assert.NoError(t, err)

	context.SetReturnType([]reflect.Type{reflect.TypeOf(""), reflect.TypeOf(true)})
	response = ([]byte)(`Hm1{s2"id"s7"test_id"}Ra2{s12"hello World!"t}z`)
	result, err = DefaultClientCodec.Decode(response, context)
	results := result.([]interface{})
	id, ok = context.ResponseHeaders().Get("id")
	assert.Equal(t, `hello World!`, results[0])
	assert.Equal(t, true, results[1])
	assert.Equal(t, id, "test_id")
	assert.Equal(t, ok, true)
	assert.NoError(t, err)

	context.SetReturnType([]reflect.Type{reflect.TypeOf("")})
	response = ([]byte)(`Hm1{s2"id"s7"test_id"}Es12"hello World!"z`)
	result, err = DefaultClientCodec.Decode(response, context)
	id, ok = context.ResponseHeaders().Get("id")
	assert.Equal(t, nil, result)
	assert.Equal(t, id, "test_id")
	assert.Equal(t, ok, true)
	assert.EqualError(t, err, "hello World!")

	context.SetReturnType([]reflect.Type{reflect.TypeOf(""), reflect.TypeOf(true), reflect.TypeOf(0)})
	response = ([]byte)(`Hm1{s2"id"s7"test_id"}Ra2{s12"hello World!"t}z`)
	result, err = DefaultClientCodec.Decode(response, context)
	results = result.([]interface{})
	id, ok = context.ResponseHeaders().Get("id")
	assert.Equal(t, `hello World!`, results[0])
	assert.Equal(t, true, results[1])
	assert.Equal(t, 0, results[2])
	assert.Equal(t, id, "test_id")
	assert.Equal(t, ok, true)
	assert.NoError(t, err)

	context.SetReturnType([]reflect.Type{reflect.TypeOf(""), reflect.TypeOf(true), reflect.TypeOf(0)})
	response = ([]byte)(`{code:200,msg:"ok"}`)
	result, err = DefaultClientCodec.Decode(response, context)
	assert.Equal(t, nil, result)
	assert.EqualError(t, err, "Invalid response\r\n"+`{code:200,msg:"ok"}`)
}
