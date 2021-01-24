/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/method.go                                       |
|                                                          |
| LastModified: Jan 25, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"context"
	"reflect"
	"runtime"
	"strings"
)

// MissingMethod is missing method.
type MissingMethod func(context context.Context, name string, args []interface{}) (result []interface{}, err error)

var missingMethodType = reflect.TypeOf((*MissingMethod)(nil)).Elem()
var contextType = reflect.TypeOf((*context.Context)(nil)).Elem()

// Method for RPC.
type Method interface {
	Func() reflect.Value
	Parameters() []reflect.Type
	Name() string
	Missing() bool
	PassContext() bool
	Options() Items
}

type method struct {
	f           reflect.Value
	params      []reflect.Type
	name        string
	missing     bool
	passContext bool
	options     Items
}

func (m method) Func() reflect.Value {
	return m.f
}

func (m method) Parameters() []reflect.Type {
	return m.params
}

func (m method) Name() string {
	return m.name
}

func (m method) Missing() bool {
	return m.missing
}

func (m method) PassContext() bool {
	return m.passContext
}

func (m method) Options() Items {
	return m.options
}

func makeMethod(f reflect.Value, name string, missing bool) method {
	if f.Kind() != reflect.Func {
		panic("f " + name + " is not a function.")
	}
	m := method{f: f, name: name, missing: missing, options: NewSafeHeaders()}
	if name == "" {
		m.name = runtime.FuncForPC(f.Pointer()).Name()
		if i := strings.LastIndexByte(m.name, '.'); i > -1 {
			m.name = m.name[i+1:]
		}
	}
	t := f.Type()
	n := t.NumIn()
	offset := 0
	if n > 0 && t.In(0) == contextType {
		m.passContext = true
		n--
		offset = 1
	}
	m.params = make([]reflect.Type, n)
	for i := 0; i < n; i++ {
		m.params[i] = t.In(i + offset)
	}
	return m
}

// NewMethod returns a Method object.
func NewMethod(f reflect.Value, name string) Method {
	return makeMethod(f, name, false)
}

// NewMissingMethod returns a missing Method object.
func NewMissingMethod(f MissingMethod) Method {
	return makeMethod(reflect.ValueOf(f), "*", true)
}
