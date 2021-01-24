/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/method_manager.go                               |
|                                                          |
| LastModified: Jan 25, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"reflect"
	"strings"
	"sync"
)

// MethodManager for RPC.
type MethodManager interface {
	Get(name string) Method
	GetNames() (names []string)
	Remove(name string)
	Add(method Method)
	AddFunction(f interface{}, name string)
	AddMethod(name string, target interface{}, alias ...string)
	AddMethods(names []string, target interface{}, namespace ...string)
}

type methodManager struct {
	methods *sync.Map
}

// NewMethodManager returns a MethodManager
func NewMethodManager() MethodManager {
	return methodManager{&sync.Map{}}
}

func (mm methodManager) Get(name string) Method {
	if method, ok := mm.methods.Load(name); ok {
		return method.(Method)
	}
	return nil
}

func (mm methodManager) GetNames() (names []string) {
	mm.methods.Range(func(key, value interface{}) bool {
		names = append(names, value.(Method).Name())
		return true
	})
	return
}

func (mm methodManager) Remove(name string) {
	mm.methods.Delete(name)
}

func (mm methodManager) Add(method Method) {
	mm.methods.Store(strings.ToLower(method.Name()), method)
}

func (mm methodManager) AddFunction(f interface{}, name string) {
	var method Method
	switch f := f.(type) {
	case reflect.Value:
		method = NewMethod(f, name)
	case *reflect.Value:
		method = NewMethod(*f, name)
	case reflect.Method:
		if name == "" {
			name = f.Name
		}
		method = NewMethod(f.Func, name)
	case *reflect.Method:
		if name == "" {
			name = f.Name
		}
		method = NewMethod(f.Func, name)
	default:
		method = NewMethod(reflect.ValueOf(f), name)
	}
	mm.Add(method)
}

func (mm methodManager) AddMethod(name string, target interface{}, alias ...string) {
	obj := reflect.ValueOf(target)
	f := obj.MethodByName(name)
	if f.Kind() != reflect.Func && obj.Kind() == reflect.Struct {
		f = obj.FieldByName(name)
	}
	if f.Kind() != reflect.Func {
		if t, ok := target.(reflect.Type); ok {
			if m, ok := t.MethodByName(name); ok {
				f = m.Func
			}
		}
	}
	if len(alias) > 0 && alias[0] != "" {
		name = alias[0]
	}
	mm.Add(NewMethod(f, name))
}

func (mm methodManager) AddMethods(names []string, target interface{}, namespace ...string) {
	for _, name := range names {
		alias := ""
		if len(namespace) > 0 && namespace[0] != "" {
			alias = namespace[0] + "_" + name
		}
		mm.AddMethod(name, target, alias)
	}
}

// AddInstanceMethods is used for publishing all the public methods and func fields with options.
func (mm methodManager) AddInstanceMethods(target interface{}, namespace ...string) {
	// TODO
}

// AddAllMethods will publish all methods and non-nil function fields on the
// obj self and on its anonymous or non-anonymous struct fields (or pointer to
// pointer ... to pointer struct fields). This is a recursive operation.
// So it's a pit, if you do not know what you are doing, do not step on.
func (mm *methodManager) AddAllMethods(target interface{}, namespace ...string) {
	// TODO
}

// AddMissingMethod is used for publishing a method,
// all methods not explicitly published will be redirected to this method.
func (mm *methodManager) AddMissingMethod(f MissingMethod) {
	mm.Add(NewMissingMethod(f))
}

// TODO
/*
func (mm *methodManager) AddNetRPCMethods(rcvr interface{}, namespace ...string) {
	if rcvr == nil {
		panic("rcvr can't be nil")
	}
	v := reflect.ValueOf(rcvr)
	t := v.Type()
	n := t.NumMethod()
	for i := 0; i < n; i++ {
		name := t.Method(i).Name
		method := v.Method(i)
		if method.CanInterface() {
			mm.addNetRPCMethod(name, method, namespace...)
		}
	}
}

func (mm *methodManager) addNetRPCMethod(name string, method reflect.Value, namespace ...string) {
	ft := method.Type()
	if ft.NumIn() != 2 || ft.IsVariadic() {
		// panic("the method " + name + " must has two arguments")
		return
	}
	if ft.In(1).Kind() != reflect.Ptr {
		// panic("the second argument of method " + name + " must be a pointer")
		return
	}
	if ft.NumOut() != 1 || ft.Out(0) != errorType {
		// panic("the result type of method " + name + " must be error")
		return
	}
	argsType := ft.In(0)
	resultType := ft.In(1).Elem()
	in := []reflect.Type{argsType}
	out := []reflect.Type{resultType, errorType}
	newft := reflect.FuncOf(in, out, false)
	newMethod := reflect.MakeFunc(newft, func(
		args []reflect.Value) (results []reflect.Value) {
		result := reflect.New(resultType)
		in := []reflect.Value{args[0], result}
		err := method.Call(in)[0]
		results = []reflect.Value{result.Elem(), err}
		return
	})
	mm.AddFunction(newMethod, name)
}
*/
