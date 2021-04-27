/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/proxy_test.go                                   |
|                                                          |
| LastModified: Apr 27, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core_test

import (
	"errors"
	"reflect"
	"testing"

	. "github.com/hprose/hprose-golang/v3/rpc/core"
	"github.com/stretchr/testify/assert"
)

func TestProxy(t *testing.T) {
	type testEmbedStruct struct {
		Test func() string
	}
	type testStruct struct {
		Test  func() string
		Test2 func(a, b int) (int, int)
		Test3 func(operate string, x ...int) (int, error)
		Embed struct {
			Test  func() string
			embed ***testEmbedStruct
		}
		*testEmbedStruct
		*int
		b bool
	}
	testInvocationHandler := func(proxy interface{}, method reflect.StructField, name string, args []interface{}) (results []interface{}, err error) {
		println(name)
		switch method.Name {
		case "Test":
			return []interface{}{"Hello World!"}, nil
		case "Test2":
			return []interface{}{args[1], args[0]}, nil
		case "Test3":
			n := 0
			for i := 1; i < len(args); i++ {
				switch args[0] {
				case "sum":
					n += args[i].(int)
				case "mul":
					if i == 1 {
						n = 1
					}
					n *= args[i].(int)
				default:
					return nil, errors.New("unknown operate")
				}
			}
			return []interface{}{n}, nil
		}
		return nil, nil
	}
	var testp **testStruct
	Proxy.Build(&testp, testInvocationHandler)
	test := *testp
	assert.Equal(t, `Hello World!`, test.Test())
	a, b := test.Test2(123, 456)
	assert.Equal(t, a, 456)
	assert.Equal(t, b, 123)
	n, err := test.Test3("sum", 1, 2, 3, 4)
	assert.Equal(t, 10, n)
	assert.NoError(t, err)
	n, err = test.Test3("mul", 1, 2, 3, 4)
	assert.Equal(t, 24, n)
	assert.NoError(t, err)
	n, err = test.Test3("div", 1, 2, 3, 4)
	assert.Equal(t, 0, n)
	assert.Error(t, err)
	assert.Equal(t, `Hello World!`, test.Embed.Test())
	assert.Equal(t, `Hello World!`, (**test.Embed.embed).Test())
	assert.Equal(t, `Hello World!`, test.testEmbedStruct.Test())
}
