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
 * hprose/service.go                                      *
 *                                                        *
 * hprose service for Go.                                 *
 *                                                        *
 * LastModified: Jul 4, 2015                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"runtime/debug"
	"strings"
)

// MissingMethod is missing method
type MissingMethod func(name string, args []reflect.Value) (result []reflect.Value)

// ServiceEvent is the service event
type ServiceEvent interface{}

type beforeInvokeEvent interface {
	OnBeforeInvoke(name string, args []reflect.Value, byref bool, context Context)
}

type beforeInvoke2Event interface {
	OnBeforeInvoke(name string, args []reflect.Value, byref bool, context Context) error
}

type afterInvokeEvent interface {
	OnAfterInvoke(name string, args []reflect.Value, byref bool, result []reflect.Value, context Context)
}

type afterInvoke2Event interface {
	OnAfterInvoke(name string, args []reflect.Value, byref bool, result []reflect.Value, context Context) error
}

type sendErrorEvent interface {
	OnSendError(err error, context Context)
}

type sendError2Event interface {
	OnSendError(err error, context Context) error
}

// Method is the publish service method
type Method struct {
	Function   reflect.Value
	ResultMode ResultMode
	SimpleMode bool
}

// NewMethod is the constructor for Method
func NewMethod(f reflect.Value, mode ResultMode, simple bool) (method *Method) {
	method = new(Method)
	method.Function = f
	method.ResultMode = mode
	method.SimpleMode = simple
	return
}

// Methods is the publish service methods
type Methods struct {
	MethodNames   []string
	RemoteMethods map[string]*Method
}

// NewMethods is the constructor for Methods
func NewMethods() (methods *Methods) {
	methods = new(Methods)
	methods.MethodNames = make([]string, 0, 64)
	methods.RemoteMethods = make(map[string]*Method)
	return
}

// AddFunction publish a func or bound method
// name is the method name
// function is a func or bound method
// options is ResultMode, SimpleMode and prefix
func (methods *Methods) AddFunction(name string, function interface{}, options ...interface{}) {
	if name == "" {
		panic("name can't be empty")
	}
	if function == nil {
		panic("function can't be nil")
	}
	f := reflect.ValueOf(function)
	if f.Kind() != reflect.Func {
		panic("function must be func or bound method")
	}
	count := len(options)
	resultMode := Normal
	simpleMode := false
	prefix := ""
	for i := 0; i < count; i++ {
		switch opt := options[i].(type) {
		case ResultMode:
			resultMode = opt
		case bool:
			simpleMode = opt
		case string:
			prefix = opt
		default:
			panic("unknown options")
		}
	}
	if prefix != "" && name != "*" {
		name = prefix + "_" + name
	}
	methods.MethodNames = append(methods.MethodNames, name)
	methods.RemoteMethods[strings.ToLower(name)] = NewMethod(f, resultMode, simpleMode)
}

// AddFunctions ...
// names are method names
// functions are funcs or bound methods
// options is the same as AddFuntion
func (methods *Methods) AddFunctions(names []string, functions []interface{}, options ...interface{}) {
	if len(names) != len(functions) {
		panic("names and functions must have the same length")
	}
	count := len(names)
	for i := 0; i < count; i++ {
		methods.AddFunction(names[i], functions[i], options...)
	}
}

// AddMethods ...
// obj is service object. all the public method and func field will be published
// options is the same as AddFuntion
func (methods *Methods) AddMethods(obj interface{}, options ...interface{}) {
	if obj == nil {
		panic("obj can't be nil")
	}
	v := reflect.ValueOf(obj)
	t := v.Type()
	n := t.NumMethod()
	for i := 0; i < n; i++ {
		name := t.Method(i).Name
		method := v.Method(i)
		if method.CanInterface() && 'A' <= name[0] && name[0] <= 'Z' {
			methods.AddFunction(name, method.Interface(), options...)
		}
	}
	for ; t.Kind() == reflect.Ptr && !v.IsNil(); v = v.Elem() {
		t = t.Elem()
	}
	if t.Kind() == reflect.Struct {
		n = t.NumField()
		for i := 0; i < n; i++ {
			f := v.Field(i)
			name := t.Field(i).Name
			if name != "" && f.CanInterface() && 'A' <= name[0] && name[0] <= 'Z' && f.IsValid() {
				for ; f.Kind() == reflect.Ptr && !f.IsNil(); f = f.Elem() {
				}
				if f.Kind() == reflect.Func && !f.IsNil() {
					methods.AddFunction(name, f.Interface(), options...)
				}
			}
		}
	}
}

// AddAllMethods will publish all methods and non-nil function fields on the obj self
// and on its anonymous or non-anonymous struct fields (or pointer to pointer ...
// to pointer struct fields). This is a recursive operation.
// So it's a pit, if you do not know what you are doing, do not step on.
func (methods *Methods) AddAllMethods(obj interface{}, options ...interface{}) {
	if obj == nil {
		panic("obj can't be nil")
	}
	v := reflect.ValueOf(obj)
	t := v.Type()
	n := t.NumMethod()
	for i := 0; i < n; i++ {
		name := t.Method(i).Name
		method := v.Method(i)
		if method.CanInterface() && 'A' <= name[0] && name[0] <= 'Z' {
			methods.AddFunction(name, method.Interface(), options...)
		}
	}
	for ; t.Kind() == reflect.Ptr && !v.IsNil(); v = v.Elem() {
		t = t.Elem()
	}
	if t.Kind() == reflect.Struct {
		n = t.NumField()
		for i := 0; i < n; i++ {
			f := v.Field(i)
			fs := t.Field(i)
			name := fs.Name
			if f.CanInterface() && f.IsValid() {
				for ; f.Kind() == reflect.Ptr && !f.IsNil(); f = f.Elem() {
				}
				if f.Kind() == reflect.Func && 'A' <= name[0] && name[0] <= 'Z' && !f.IsNil() {
					methods.AddFunction(name, f.Interface(), options...)
				} else if f.Kind() == reflect.Struct {
					if fs.Anonymous {
						methods.AddAllMethods(f.Interface(), options...)
					} else if 'A' <= name[0] && name[0] <= 'Z' {
						prefix := ""
						k := -1
						count := len(options)
						for j := 0; j < count; j++ {
							switch opt := options[i].(type) {
							case string:
								prefix = opt
								k = j
							}
						}
						if prefix == "" {
							prefix = name
						} else {
							prefix = prefix + "_" + name
						}
						if k >= 0 {
							options[k] = prefix
						} else {
							options = append(options, prefix)
						}
						methods.AddAllMethods(f.Interface(), options...)
					}
				}
			}
		}
	}
}

// AddMissingMethod ...
// All methods not explicitly published will be redirected to the method.
func (methods *Methods) AddMissingMethod(method MissingMethod, options ...interface{}) {
	methods.AddFunction("*", method, options...)
}

// ArgsFixer ...
type ArgsFixer interface {
	FixArgs(args []reflect.Value, lastParamType reflect.Type, context Context) []reflect.Value
}

func fixArgs(args []reflect.Value, lastParamType reflect.Type, context Context) []reflect.Value {
	if lastParamType.String() == "interface {}" ||
		lastParamType.String() == "hprose.Context" {
		args = append(args, reflect.ValueOf(context))
	}
	return args
}

// BaseService is the hprose base service
type BaseService struct {
	*Methods
	ServiceEvent
	DebugEnabled bool
	filters      []Filter
	argsfixer    ArgsFixer
}

// NewBaseService is the constructor for BaseService
func NewBaseService() (service *BaseService) {
	service = new(BaseService)
	service.Methods = NewMethods()
	service.filters = make([]Filter, 0)
	return
}

// GetFilter return the first filter
func (service *BaseService) GetFilter() Filter {
	if len(service.filters) == 0 {
		return nil
	}
	return service.filters[0]
}

// SetFilter set the only filter
func (service *BaseService) SetFilter(filter Filter) {
	service.filters = make([]Filter, 0)
	if filter != nil {
		service.filters = append(service.filters, filter)
	}
}

// AddFilter add a filter
func (service *BaseService) AddFilter(filter Filter) {
	service.filters = append(service.filters, filter)
}

// RemoveFilter remove a filter
func (service *BaseService) RemoveFilter(filter Filter) {
	n := len(service.filters)
	for i := 0; i < n; i++ {
		if service.filters[i] == filter {
			if i == n-1 {
				service.filters = service.filters[:i]
			} else {
				service.filters = append(service.filters[:i], service.filters[i+1:]...)
			}
			return
		}
	}
}

func (service *BaseService) responseEnd(buf []byte, context Context) []byte {
	n := len(service.filters)
	for i := 0; i < n; i++ {
		buf = service.filters[i].OutputFilter(buf, context)
	}
	return buf
}

func (service *BaseService) fireErrorEvent(err error, context Context) error {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	if service.ServiceEvent != nil {
		if event, ok := service.ServiceEvent.(sendErrorEvent); ok {
			event.OnSendError(err, context)
		} else if event, ok := service.ServiceEvent.(sendError2Event); ok {
			return event.OnSendError(err, context)
		}
	}
	return err
}

func (service *BaseService) sendError(err error, context Context) []byte {
	err = service.fireErrorEvent(err, context)
	buf := new(bytes.Buffer)
	writer := NewWriter(buf, true)
	writer.Stream.WriteByte(TagError)
	writer.WriteString(err.Error())
	writer.Stream.WriteByte(TagEnd)
	return service.responseEnd(buf.Bytes(), context)
}

func (service *BaseService) doInvoke(data []byte, context Context) []byte {
	istream := NewBytesReader(data)
	reader := NewReader(istream, false)
	buf := new(bytes.Buffer)
	for {
		reader.Reset()
		name, err := reader.ReadString()
		if err != nil {
			return service.sendError(err, context)
		}
		alias := strings.ToLower(name)
		remoteMethod := service.RemoteMethods[alias]
		count := 0
		var args []reflect.Value
		byref := false
		var tag byte
		if tag, err = reader.CheckTags([]byte{TagList, TagEnd, TagCall}); err != nil {
			return service.sendError(err, context)
		}
		if tag == TagList {
			reader.Reset()
			if count, err = reader.ReadInteger(TagOpenbrace); err != nil {
				return service.sendError(err, context)
			}
			args = make([]reflect.Value, count)
			if remoteMethod == nil {
				for i := 0; i < count; i++ {
					var e interface{}
					args[i] = reflect.ValueOf(&e).Elem()
				}
				if err = reader.ReadArray(args); err != nil {
					return service.sendError(err, context)
				}
			} else {
				ft := remoteMethod.Function.Type()
				n := ft.NumIn()
				if ft.IsVariadic() {
					n--
				}
				if n < count {
					for i := 0; i < n; i++ {
						args[i] = reflect.New(ft.In(i)).Elem()
					}
					if ft.IsVariadic() {
						t := ft.In(n).Elem()
						for i := n; i < count; i++ {
							args[i] = reflect.New(t).Elem()
						}
						if err = reader.ReadArray(args); err != nil {
							return service.sendError(err, context)
						}
					} else {
						for i := n; i < count; i++ {
							var e interface{}
							args[i] = reflect.ValueOf(&e).Elem()
						}
						if err = reader.ReadArray(args); err != nil {
							return service.sendError(err, context)
						}
						args = args[:n]
					}
				} else {
					for i := 0; i < count; i++ {
						args[i] = reflect.New(ft.In(i)).Elem()
					}
					if err = reader.ReadArray(args[0:count]); err != nil {
						return service.sendError(err, context)
					}
					if count+1 == n {
						args = service.argsfixer.FixArgs(args, ft.In(count), context)
					}
				}
			}
			if tag, err = reader.CheckTags([]byte{TagTrue, TagEnd, TagCall}); err != nil {
				return service.sendError(err, context)
			}
			if tag == TagTrue {
				byref = true
				if tag, err = reader.CheckTags([]byte{TagEnd, TagCall}); err != nil {
					return service.sendError(err, context)
				}
			}
		} else {
			args = make([]reflect.Value, 0)
			if remoteMethod != nil {
				ft := remoteMethod.Function.Type()
				if ft.NumIn() == 1 && !ft.IsVariadic() {
					args = service.argsfixer.FixArgs(args, ft.In(0), context)
				}
			}
		}
		if service.ServiceEvent != nil {
			if event, ok := service.ServiceEvent.(beforeInvokeEvent); ok {
				event.OnBeforeInvoke(name, args, byref, context)
			} else if event, ok := service.ServiceEvent.(beforeInvoke2Event); ok {
				err = event.OnBeforeInvoke(name, args, byref, context)
				if err != nil {
					return service.sendError(err, context)
				}
			}
		}
		var result []reflect.Value
		if result, err = func() (out []reflect.Value, err error) {
			defer func() {
				if e := recover(); e != nil && err == nil {
					if service.DebugEnabled {
						err = fmt.Errorf("%v\r\n%s", e, debug.Stack())
					} else {
						err = fmt.Errorf("%v", e)
					}
				}
			}()
			if remoteMethod == nil {
				remoteMethod = service.RemoteMethods["*"]
				if remoteMethod == nil {
					return nil, errors.New("Can't find this method " + name)
				}
				if missingMethod, ok := remoteMethod.Function.Interface().(MissingMethod); ok {
					return missingMethod(name, args), nil
				}
				return nil, errors.New("Can't find this method " + name)
			}
			return remoteMethod.Function.Call(args), nil
		}(); err != nil {
			return service.sendError(err, context)
		}
		if service.ServiceEvent != nil {
			if event, ok := service.ServiceEvent.(afterInvokeEvent); ok {
				event.OnAfterInvoke(name, args, byref, result, context)
			} else if event, ok := service.ServiceEvent.(afterInvoke2Event); ok {
				err = event.OnAfterInvoke(name, args, byref, result, context)
				if err != nil {
					return service.sendError(err, context)
				}
			}
		}
		resultLength := len(result)
		if resultLength > 0 {
			t := remoteMethod.Function.Type().Out(resultLength - 1)
			if t.Implements(reflect.TypeOf(&err).Elem()) {
				if err, ok := result[resultLength-1].Interface().(error); ok {
					return service.sendError(err, context)
				}
				resultLength--
				result = result[:resultLength]
			}
		}
		if remoteMethod.ResultMode != Normal {
			if resultLength == 0 {
				return service.sendError(errors.New("can't find the result value"), context)
			}
			switch r := result[0].Interface().(type) {
			case []byte:
				data = r
			case *[]byte:
				data = *r
			case bytes.Buffer:
				data = r.Bytes()
			case *bytes.Buffer:
				data = r.Bytes()
			case string:
				data = []byte(r)
			case *string:
				data = []byte(*r)
			default:
				return service.sendError(errors.New("the result type is wrong"), context)
			}
			if remoteMethod.ResultMode == RawWithEndTag {
				return service.responseEnd(data, context)
			}
		}
		if remoteMethod.ResultMode == Raw {
			buf.Write(data)
		} else {
			writer := NewWriter(buf, remoteMethod.SimpleMode)
			writer.Stream.WriteByte(TagResult)
			if remoteMethod.ResultMode == Serialized {
				if _, err = writer.Stream.Write(data); err != nil {
					return service.sendError(err, context)
				}
			} else {
				switch resultLength {
				case 0:
					err = writer.Serialize(nil)
				case 1:
					err = writer.WriteValue(result[0])
				default:
					err = writer.WriteArray(result)
				}
				if err != nil {
					return service.sendError(err, context)
				}
			}
			if byref {
				writer.Stream.WriteByte(TagArgument)
				writer.Reset()
				if err = writer.WriteArray(args); err != nil {
					return service.sendError(err, context)
				}
			}
		}
		if tag != TagCall {
			break
		}
	}
	buf.WriteByte(TagEnd)
	return service.responseEnd(buf.Bytes(), context)
}

func (service *BaseService) doFunctionList(context Context) []byte {
	buf := new(bytes.Buffer)
	writer := NewWriter(buf, true)
	writer.Stream.WriteByte(TagFunctions)
	if err := writer.Serialize(service.MethodNames); err != nil {
		return service.sendError(err, context)
	}
	writer.Stream.WriteByte(TagEnd)
	return service.responseEnd(buf.Bytes(), context)
}

// Handle the hprose request and return the hprose response
func (service *BaseService) Handle(data []byte, context Context) (output []byte) {
	defer func() {
		if e := recover(); e != nil {
			var err error
			if service.DebugEnabled {
				err = fmt.Errorf("%v\r\n%s", e, debug.Stack())
			} else {
				err = fmt.Errorf("%v", e)
			}
			output = service.sendError(err, context)
		}
	}()
	for i := len(service.filters) - 1; i >= 0; i-- {
		data = service.filters[i].InputFilter(data, context)
	}
	if len(data) == 0 {
		return service.sendError(errors.New("no Hprose RPC request"), context)
	}
	tag := data[0]
	switch tag {
	case TagCall:
		return service.doInvoke(data[1:], context)
	case TagEnd:
		return service.doFunctionList(context)
	default:
		return service.sendError(errors.New("Wrong Reqeust: \r\n"+string(data)), context)
	}
}
