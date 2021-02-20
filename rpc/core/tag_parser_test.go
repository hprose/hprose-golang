/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/proxy_test.go                                   |
|                                                          |
| LastModified: Feb 20, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTagParser(t *testing.T) {
	type testStruct struct {
		Test func() string `name:"test" timeout:"1000" context:"str1:Hello World!, str2:'12345', str3:\"\", str4:\"null\", str5:\"1,2,3\", int:12345, boolTrue, boolFalse:false, null:nil, null2:null, float:3.14" header:"oneway,id:123,  "`
	}
	f, ok := reflect.TypeOf(testStruct{}).FieldByName("Test")
	assert.True(t, ok)
	parser := parseTag(NewClientContext(), f.Tag)
	assert.Equal(t, "test", parser.Name)
	assert.Equal(t, time.Second, parser.Context.Timeout)
	items := parser.Context.Items()
	assert.Equal(t, "Hello World!", items.GetString("str1"))
	assert.Equal(t, "12345", items.GetString("str2"))
	assert.Equal(t, "12345", items.GetString("str2"))
	assert.Equal(t, "", items.GetString("str3"))
	assert.Equal(t, "null", items.GetString("str4"))
	assert.Equal(t, "1,2,3", items.GetString("str5"))
	assert.Equal(t, 12345, items.GetInt("int"))
	assert.True(t, items.GetBool("boolTrue"))
	assert.Equal(t, false, items.GetBool("boolFalse"))
	null, ok := items.Get("null")
	assert.True(t, ok)
	assert.Equal(t, nil, null)
	null, ok = items.Get("null2")
	assert.True(t, ok)
	assert.Equal(t, nil, null)
	null, ok = items.Get("null3")
	assert.False(t, ok)
	assert.Equal(t, nil, null)
	assert.Equal(t, 3.14, items.GetFloat("float"))
	header := parser.Context.RequestHeaders()
	assert.True(t, header.GetBool("oneway"))
	assert.False(t, header.GetBool("noway"))
	assert.Equal(t, 123, header.GetInt("id"))
}
