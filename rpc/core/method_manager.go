/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/method_manager.go                               |
|                                                          |
| LastModified: Feb 16, 2021                               |
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
	Names() (names []string)
	Methods() (methods []Method)
	Remove(name string)
	Add(method Method)
	AddFunction(f interface{}, name string)
	AddMethod(name string, target interface{}, alias ...string)
	AddMethods(names []string, target interface{}, namespace ...string)
	AddInstanceMethods(target interface{}, namespace ...string)
	AddAllMethods(target interface{}, namespace ...string)
	AddMissingMethod(f MissingMethod)
	AddNetRPCMethods(rcvr interface{}, namespace ...string)
}

type methodManager struct {
	methods sync.Map
}

// NewMethodManager returns a MethodManager
func NewMethodManager() MethodManager {
	return &methodManager{}
}

func (mm *methodManager) Get(name string) Method {
	if method, ok := mm.methods.Load(strings.ToLower(name)); ok {
		return method.(Method)
	}
	if method, ok := mm.methods.Load("*"); ok {
		return method.(Method)
	}
	return nil
}

func (mm *methodManager) Names() (names []string) {
	mm.methods.Range(func(key, value interface{}) bool {
		names = append(names, value.(Method).Name())
		return true
	})
	return
}

func (mm *methodManager) Methods() (methods []Method) {
	mm.methods.Range(func(key, value interface{}) bool {
		methods = append(methods, value.(Method))
		return true
	})
	return
}

func (mm *methodManager) Remove(name string) {
	mm.methods.Delete(strings.ToLower(name))
}

func (mm *methodManager) Add(method Method) {
	mm.methods.Store(strings.ToLower(method.Name()), method)
}

func (mm *methodManager) AddFunction(f interface{}, name string) {
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

func (mm *methodManager) addFunction(f interface{}, name string, namespace ...string) {
	if len(namespace) > 0 && namespace[0] != "" {
		name = namespace[0] + "_" + name
	}
	mm.AddFunction(f, name)
}

func (mm *methodManager) AddMethod(name string, target interface{}, alias ...string) {
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
	if f.CanInterface() {
		if len(alias) > 0 && alias[0] != "" {
			name = alias[0]
		}
		mm.Add(NewMethod(f, name))
	}
}

func (mm *methodManager) addMethod(name string, target interface{}, namespace ...string) {
	alias := ""
	if len(namespace) > 0 && namespace[0] != "" {
		alias = namespace[0] + "_" + name
	}
	mm.AddMethod(name, target, alias)
}

func (mm *methodManager) AddMethods(names []string, target interface{}, namespace ...string) {
	for _, name := range names {
		mm.addMethod(name, target, namespace...)
	}
}

func (mm *methodManager) addMethods(v reflect.Value, t reflect.Type, namespace ...string) {
	n := t.NumMethod()
	for i := 0; i < n; i++ {
		name := t.Method(i).Name
		method := v.Method(i)
		if method.CanInterface() {
			mm.addFunction(method, name, namespace...)
		}
	}
}

func getPtrTo(v reflect.Value, t reflect.Type, kind reflect.Kind) (reflect.Value, reflect.Type) {
	for t.Kind() == reflect.Ptr && !v.IsNil() && t.Elem().Kind() == kind {
		v = v.Elem()
		t = t.Elem()
	}
	return v, t
}

func (mm *methodManager) addFuncField(v reflect.Value, t reflect.Type, i int, namespace ...string) {
	f := v.Field(i)
	name := t.Field(i).Name
	if f.IsValid() {
		f, _ = getPtrTo(f, f.Type(), reflect.Func)
		if f.Kind() == reflect.Func && !f.IsNil() && f.CanInterface() {
			mm.addFunction(f, name, namespace...)
		}
	}
}

func (mm *methodManager) recursiveAddFuncFields(v reflect.Value, t reflect.Type, i int, namespace ...string) {
	f := v.Field(i)
	fs := t.Field(i)
	name := fs.Name
	if f.IsValid() {
		return
	}
	f, _ = getPtrTo(f, f.Type(), reflect.Func)
	switch f.Kind() {
	case reflect.Func, reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Chan:
		if f.IsNil() {
			return
		}
	}
	if !f.CanInterface() {
		return
	}
	if f.Kind() == reflect.Func {
		mm.addFunction(f, name, namespace...)
		return
	}
	if !fs.Anonymous {
		if len(namespace) > 0 {
			if namespace[0] == "" {
				namespace[0] = name
			} else {
				namespace[0] += "_" + name
			}
		} else {
			namespace = append(namespace, name)
		}
	}
	mm.AddAllMethods(f.Interface(), namespace...)
}

type addFuncFunc func(v reflect.Value, t reflect.Type, i int, namespace ...string)

func (mm *methodManager) addInstanceMethods(target interface{}, addFunc addFuncFunc, namespace ...string) {
	if target == nil {
		panic("target can't be nil")
	}
	v := reflect.ValueOf(target)
	t := v.Type()
	mm.addMethods(v, t, namespace...)
	v, t = getPtrTo(v, t, reflect.Struct)
	if t.Kind() == reflect.Struct {
		n := t.NumField()
		for i := 0; i < n; i++ {
			addFunc(v, t, i, namespace...)
		}
	}
}

// AddInstanceMethods is used for publishing all the public methods and func fields with namespace.
func (mm *methodManager) AddInstanceMethods(target interface{}, namespace ...string) {
	mm.addInstanceMethods(target, mm.addFuncField, namespace...)
}

// AddAllMethods will publish all methods and non-nil function fields on the
// obj self and on its anonymous or non-anonymous struct fields (or pointer to
// pointer ... to pointer struct fields). This is a recursive operation.
// So it's a pit, if you do not know what you are doing, do not step on.
func (mm *methodManager) AddAllMethods(target interface{}, namespace ...string) {
	mm.addInstanceMethods(target, mm.recursiveAddFuncFields, namespace...)
}

// AddMissingMethod is used for publishing a method,
// all methods not explicitly published will be redirected to this method.
func (mm *methodManager) AddMissingMethod(f MissingMethod) {
	mm.Add(NewMissingMethod(f))
}

// AddNetRPCMethods is used for publishing methods defined for net/rpc.
func (mm *methodManager) AddNetRPCMethods(rcvr interface{}, namespace ...string) {
	if rcvr == nil {
		panic("rcvr can't be nil")
	}
	v := reflect.ValueOf(rcvr)
	t := v.Type()
	n := t.NumMethod()
	for i := 0; i < n; i++ {
		if method := v.Method(i); method.CanInterface() {
			name := t.Method(i).Name
			if len(namespace) > 0 && namespace[0] != "" {
				name = namespace[0] + "_" + name
			}
			mm.addNetRPCMethod(name, method)
		}
	}
}

func (mm *methodManager) addNetRPCMethod(name string, method reflect.Value) {
	ft := method.Type()
	if ft.NumIn() != 2 || ft.IsVariadic() {
		return
	}
	if ft.In(1).Kind() != reflect.Ptr {
		return
	}
	if ft.NumOut() != 1 || ft.Out(0) != errorType {
		return
	}
	argsType := ft.In(0)
	resultType := ft.In(1)
	in := []reflect.Type{argsType}
	out := []reflect.Type{resultType, errorType}
	newft := reflect.FuncOf(in, out, false)
	newMethod := reflect.MakeFunc(newft, func(
		args []reflect.Value) (results []reflect.Value) {
		result := reflect.New(resultType.Elem())
		in := []reflect.Value{args[0], result}
		err := method.Call(in)[0]
		results = []reflect.Value{result, err}
		return
	})
	mm.AddFunction(newMethod, name)
}
