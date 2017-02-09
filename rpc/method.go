/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * rpc/method.go                                          *
 *                                                        *
 * hprose method manager for Go.                          *
 *                                                        *
 * LastModified: Oct 31, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import (
	"reflect"
	"strings"
	"sync"
)

// Options is the options of the published service method
type Options struct {
	Mode           ResultMode
	Simple         bool
	Oneway         bool
	NameSpace      string
	JSONCompatible bool
}

// Method is the published service method
type Method struct {
	Function reflect.Value
	Options
}

// methodManager manages published service methods
type methodManager struct {
	MethodNames   []string
	RemoteMethods map[string]*Method
	mmLocker      sync.Mutex
}

func (mm *methodManager) initMethodManager() {
	mm.MethodNames = make([]string, 0, 64)
	mm.RemoteMethods = make(map[string]*Method)
}

// AddFunction publish a func or bound method
// name is the method name
// function is a func or bound method
// option includes Mode, Simple, Oneway and NameSpace
func (mm *methodManager) AddFunction(
	name string, function interface{}, option ...Options) {
	if name == "" {
		panic("name can't be empty")
	}
	if function == nil {
		panic("function can't be nil")
	}
	f, ok := function.(reflect.Value)
	if !ok {
		f = reflect.ValueOf(function)
	}
	if f.Kind() != reflect.Func {
		panic("function must be func or bound method")
	}
	var options Options
	if len(option) > 0 {
		options = option[0]
	}
	if options.NameSpace != "" && name != "*" {
		name = options.NameSpace + "_" + name
	}
	mm.mmLocker.Lock()
	if mm.RemoteMethods[strings.ToLower(name)] == nil {
		mm.MethodNames = append(mm.MethodNames, name)
	}
	mm.RemoteMethods[strings.ToLower(name)] = &Method{f, options}
	mm.mmLocker.Unlock()
}

// AddFunctions is used for batch publishing service method
func (mm *methodManager) AddFunctions(
	names []string, functions []interface{}, option ...Options) {
	count := len(names)
	if count != len(functions) {
		panic("names and functions must have the same length")
	}
	for i := 0; i < count; i++ {
		mm.AddFunction(names[i], functions[i], option...)
	}
}

// AddMethod is used for publishing a method on the obj with an alias
func (mm *methodManager) AddMethod(
	name string, obj interface{}, alias string, option ...Options) {
	if obj == nil {
		panic("obj can't be nil")
	}
	f := reflect.ValueOf(obj).MethodByName(name)
	if alias != "" {
		name = alias
	}
	if f.CanInterface() {
		mm.AddFunction(name, f, option...)
	}
}

// AddMethods is used for batch publishing methods on the obj with aliases
func (mm *methodManager) AddMethods(
	names []string, obj interface{}, aliases []string, option ...Options) {
	if obj == nil {
		panic("obj can't be nil")
	}
	count := len(names)
	if aliases == nil {
		for i := 0; i < count; i++ {
			mm.AddMethod(names[i], obj, "", option...)
		}
		return
	}
	if len(aliases) != count {
		panic("names and aliases must have the same length")
	}
	for i := 0; i < count; i++ {
		mm.AddMethod(names[i], obj, aliases[i], option...)
	}
}

func (mm *methodManager) addMethods(
	v reflect.Value, t reflect.Type, option ...Options) {
	n := t.NumMethod()
	for i := 0; i < n; i++ {
		name := t.Method(i).Name
		method := v.Method(i)
		if method.CanInterface() {
			mm.AddFunction(name, method, option...)
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

func (mm *methodManager) addFuncField(
	v reflect.Value, t reflect.Type, i int, option ...Options) {
	f := v.Field(i)
	name := t.Field(i).Name
	if !f.CanInterface() || !f.IsValid() {
		return
	}
	f, _ = getPtrTo(f, f.Type(), reflect.Func)
	if f.Kind() == reflect.Func && !f.IsNil() {
		mm.AddFunction(name, f, option...)
	}
}

func (mm *methodManager) recursiveAddFuncFields(
	v reflect.Value, t reflect.Type, i int, option ...Options) {
	f := v.Field(i)
	fs := t.Field(i)
	name := fs.Name
	if !f.CanInterface() || !f.IsValid() {
		return
	}
	f, _ = getPtrTo(f, f.Type(), reflect.Func)
	switch f.Kind() {
	case reflect.Func, reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Chan:
		if f.IsNil() {
			return
		}
	}
	if f.Kind() == reflect.Func {
		mm.AddFunction(name, f, option...)
		return
	}
	if fs.Anonymous {
		mm.AddAllMethods(f.Interface(), option...)
	} else {
		var newOptions Options
		if len(option) > 0 {
			newOptions = option[0]
		}
		if newOptions.NameSpace == "" {
			newOptions.NameSpace = name
		} else {
			newOptions.NameSpace += "_" + name
		}
		mm.AddAllMethods(f.Interface(), newOptions)
	}
}

type addFuncFunc func(
	v reflect.Value,
	t reflect.Type,
	i int,
	option ...Options)

func (mm *methodManager) addInstanceMethods(
	obj interface{}, addFunc addFuncFunc, option ...Options) {
	if obj == nil {
		panic("obj can't be nil")
	}
	v := reflect.ValueOf(obj)
	t := v.Type()
	mm.addMethods(v, t, option...)
	v, t = getPtrTo(v, t, reflect.Struct)
	if t.Kind() == reflect.Struct {
		n := t.NumField()
		for i := 0; i < n; i++ {
			addFunc(v, t, i, option...)
		}
	}
}

// AddInstanceMethods is used for publishing all the public methods and func fields with options.
func (mm *methodManager) AddInstanceMethods(
	obj interface{}, option ...Options) {
	mm.addInstanceMethods(obj, mm.addFuncField, option...)
}

// AddAllMethods will publish all methods and non-nil function fields on the
// obj self and on its anonymous or non-anonymous struct fields (or pointer to
// pointer ... to pointer struct fields). This is a recursive operation.
// So it's a pit, if you do not know what you are doing, do not step on.
func (mm *methodManager) AddAllMethods(
	obj interface{}, option ...Options) {
	mm.addInstanceMethods(obj, mm.recursiveAddFuncFields, option...)
}

// MissingMethod is missing method
type MissingMethod func(name string, args []reflect.Value, context Context) (result []reflect.Value, err error)

// AddMissingMethod is used for publishing a method,
// all methods not explicitly published will be redirected to this method.
func (mm *methodManager) AddMissingMethod(
	method MissingMethod, option ...Options) {
	mm.AddFunction("*", method, option...)
}

// AddNetRPCMethods is used for publishing methods defined for net/rpc.
//
// For example:
//
//		package server
//
//		type Args struct {
//			A, B int
//		}
//
//		type Quotient struct {
//			Quo, Rem int
//		}
//
//		type Arith int
//
//		func (t *Arith) Multiply(args *Args, reply *int) error {
//			*reply = args.A * args.B
//			return nil
//		}
//
//		func (t *Arith) Divide(args *Args, quo *Quotient) error {
//			if args.B == 0 {
//				return errors.New("divide by zero")
//			}
//			quo.Quo = args.A / args.B
//			quo.Rem = args.A % args.B
//			return nil
//		}
//
// You can publish it in hprose like this:
//
//		service := rpc.NewHTTPService()
//		service.AddNetRPCMethods(new(Arith), rpc.Options{})
//		http.ListenAndServe(":8080", service)
//
//
// Then define stub struct in client:
//
//		type Stub struct {
//			// Synchronous call
//			Multiply func(args *Args) int
//			// Asynchronous call
//			Divide func(func(*Quotient, error), *Args)
//		}
//
// Then it can make a remote call:
//
//		client := rpc.NewClient("http://127.0.0.1:8080")
//		var stub *Stub
//		client.UseService(&stub)
//		fmt.Println(stub.Multiply(&Args{8, 7}))
//		stub.Divide(func(result *Quotient, err error) {
//			if err != nil {
//				log.Fatal("arith error:", err)
//			} else {
//				fmt.Println(result.Quo, result.Rem)
//			}
//		}, &Args{8, 7})
//
// You can also call it in other languages.
//
// Here is a simple example in JavaScript:
//
//		// The method name is case-insensitive
//		var methods = ['multiply', 'divide'];
//		var client = hprose.Client.create('http://127.0.0.1:8080/', methods);
//		// The first letter of the field name is convert to lowercase.
//		client.multiply({a:8, b:7}).then(function(product) {
//			return client.divide({a:product, b:6});
//		}).then(function(quotient) {
//			// The first letter of the field name is convert to lowercase.
//			console.log(quotient.quo, quotient.rem);
//		}).catch(function(err) {
//			// The returned err is here
//			console.error(err);
//		});
//
func (mm *methodManager) AddNetRPCMethods(rcvr interface{}, option ...Options) {
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
			mm.addNetRPCMethod(name, method, option...)
		}
	}
}

func (mm *methodManager) addNetRPCMethod(
	name string, method reflect.Value, option ...Options) {
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
	mm.AddFunction(name, newMethod, option...)
}

// Remove the published func or method by name
func (mm *methodManager) Remove(name string) {
	name = strings.ToLower(name)
	mm.mmLocker.Lock()
	n := len(mm.MethodNames)
	for i := 0; i < n; i++ {
		if strings.ToLower(mm.MethodNames[i]) == name {
			mm.MethodNames = append(mm.MethodNames[:i], mm.MethodNames[i+1:]...)
			break
		}
	}
	delete(mm.RemoteMethods, name)
	mm.mmLocker.Unlock()
}
