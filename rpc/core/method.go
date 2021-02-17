/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/method.go                                       |
|                                                          |
| LastModified: Feb 8, 2021                                |
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

// Method for RPC.
type Method interface {
	Func() reflect.Value
	Parameters() []reflect.Type
	Name() string
	Missing() bool
	PassContext() bool
	Options() Dict
}

var contextType = reflect.TypeOf((context.Context)(nil))
var nameType = reflect.TypeOf("")
var argsType = reflect.TypeOf([]interface{}{})

type contextMissingMethod func(ctx context.Context, name string, args []interface{}) (result []interface{}, err error)

func (m contextMissingMethod) Func() reflect.Value {
	return reflect.ValueOf(m)
}

func (m contextMissingMethod) Parameters() []reflect.Type {
	return []reflect.Type{nameType, argsType}
}

func (m contextMissingMethod) Name() string {
	return "*"
}

func (m contextMissingMethod) Missing() bool {
	return true
}

func (m contextMissingMethod) PassContext() bool {
	return true
}

func (m contextMissingMethod) Options() Dict {
	return nil
}

type missingMethod func(name string, args []interface{}) (result []interface{}, err error)

func (m missingMethod) Func() reflect.Value {
	return reflect.ValueOf(m)
}

func (m missingMethod) Parameters() []reflect.Type {
	return []reflect.Type{nameType, argsType}
}

func (m missingMethod) Name() string {
	return "*"
}

func (m missingMethod) Missing() bool {
	return true
}

func (m missingMethod) PassContext() bool {
	return false
}

func (m missingMethod) Options() Dict {
	return nil
}

// MissingMethod returns a missing Method object.
func MissingMethod(f interface{}) Method {
	switch m := f.(type) {
	case func(ctx context.Context, name string, args []interface{}) (result []interface{}, err error):
		return contextMissingMethod(m)
	case func(name string, args []interface{}) (result []interface{}, err error):
		return missingMethod(m)
	}
	return nil
}

type method struct {
	f           reflect.Value
	params      []reflect.Type
	name        string
	passContext bool
	options     Dict
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
	return false
}

func (m method) PassContext() bool {
	return m.passContext
}

func (m method) Options() Dict {
	return m.options
}

func makeMethod(f reflect.Value, name string) method {
	if f.Kind() != reflect.Func {
		panic("f " + name + " is not a function.")
	}
	m := method{f: f, name: name, options: NewSafeDict()}
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
	return makeMethod(f, name)
}
