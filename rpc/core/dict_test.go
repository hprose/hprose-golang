/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/dict_test.go                                    |
|                                                          |
| LastModified: May 9, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core_test

import (
	"testing"

	. "github.com/hprose/hprose-golang/v3/rpc/core"
	"github.com/stretchr/testify/assert"
)

func TestSafeDict(t *testing.T) {
	dict := NewSafeDict()
	dict.Set("int", 1)
	dict.Set("uint", uint(2))
	dict.Set("int64", int64(3))
	dict.Set("uint64", uint64(4))
	dict.Set("float", 3.14)
	dict.Set("bool", true)
	dict.Set("str", "hello")

	assert.Equal(t, 1, dict.GetInt("int"))
	assert.Equal(t, uint(2), dict.GetUInt("uint"))
	assert.Equal(t, int64(3), dict.GetInt64("int64"))
	assert.Equal(t, uint64(4), dict.GetUInt64("uint64"))
	assert.Equal(t, 3.14, dict.GetFloat("float"))
	assert.Equal(t, true, dict.GetBool("bool"))
	assert.Equal(t, "hello", dict.GetString("str"))

	assert.Equal(t, 1, dict.GetInt("x", 1))
	assert.Equal(t, uint(2), dict.GetUInt("x", uint(2)))
	assert.Equal(t, int64(3), dict.GetInt64("x", int64(3)))
	assert.Equal(t, uint64(4), dict.GetUInt64("x", uint64(4)))
	assert.Equal(t, 3.14, dict.GetFloat("x", 3.14))
	assert.Equal(t, true, dict.GetBool("x", true))
	assert.Equal(t, "hello", dict.GetString("x", "hello"))

	assert.Equal(t, 0, dict.GetInt("x"))
	assert.Equal(t, uint(0), dict.GetUInt("x"))
	assert.Equal(t, int64(0), dict.GetInt64("x"))
	assert.Equal(t, uint64(0), dict.GetUInt64("x"))
	assert.Equal(t, 0.0, dict.GetFloat("x"))
	assert.Equal(t, false, dict.GetBool("x"))
	assert.Equal(t, "", dict.GetString("x"))

	assert.False(t, dict.Empty())

	dict2 := NewDict()
	dict.CopyTo(dict2)

	assert.False(t, dict2.Empty())

	assert.Equal(t, 1, dict2.GetInt("int"))
	assert.Equal(t, uint(2), dict2.GetUInt("uint"))
	assert.Equal(t, int64(3), dict2.GetInt64("int64"))
	assert.Equal(t, uint64(4), dict2.GetUInt64("uint64"))
	assert.Equal(t, 3.14, dict2.GetFloat("float"))
	assert.Equal(t, true, dict2.GetBool("bool"))
	assert.Equal(t, "hello", dict2.GetString("str"))

	assert.Equal(t, 1, dict2.GetInt("x", 1))
	assert.Equal(t, uint(2), dict2.GetUInt("x", uint(2)))
	assert.Equal(t, int64(3), dict2.GetInt64("x", int64(3)))
	assert.Equal(t, uint64(4), dict2.GetUInt64("x", uint64(4)))
	assert.Equal(t, 3.14, dict2.GetFloat("x", 3.14))
	assert.Equal(t, true, dict2.GetBool("x", true))
	assert.Equal(t, "hello", dict2.GetString("x", "hello"))

	assert.Equal(t, 0, dict2.GetInt("x"))
	assert.Equal(t, uint(0), dict2.GetUInt("x"))
	assert.Equal(t, int64(0), dict2.GetInt64("x"))
	assert.Equal(t, uint64(0), dict2.GetUInt64("x"))
	assert.Equal(t, 0.0, dict2.GetFloat("x"))
	assert.Equal(t, false, dict2.GetBool("x"))
	assert.Equal(t, "", dict2.GetString("x"))

	dict.Del("str")
	assert.Equal(t, "", dict.GetString("str"))
	assert.Equal(t, "hello", dict2.GetString("str"))

	dict2.Del("str")
	assert.Equal(t, "", dict2.GetString("str"))

	dict.Range(func(key string, value interface{}) bool {
		v, ok := dict2.Get(key)
		assert.True(t, ok)
		assert.Equal(t, value, v)
		return true
	})

	dict2.Range(func(key string, value interface{}) bool {
		v, ok := dict.Get(key)
		assert.True(t, ok)
		assert.Equal(t, value, v)
		return true
	})

}
