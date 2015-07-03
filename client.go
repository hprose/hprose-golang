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
 * hprose/client.go                                       *
 *                                                        *
 * hprose client for Go.                                  *
 *                                                        *
 * LastModified: Jul 3, 2015                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

/*
Package hprose client example:

		package main

		import (
			"fmt"
			"hprose"
		)

		type testUser struct {
			Name     string
			Sex      int
			Birthday time.Time
			Age      int
			Married  bool
		}

		type testRemoteObject struct {
			Hello               func(string) string
			HelloWithError      func(string) (string, error)               `name:"hello"`
			AsyncHello          func(string) <-chan string                 `name:"hello"`
			AsyncHelloWithError func(string) (<-chan string, <-chan error) `name:"hello"`
			Sum                 func(...int) int
			SwapKeyAndValue     func(*map[string]string) map[string]string `byref:"true"`
			SwapInt             func(int, int) (int, int)                  `name:"swap"`
			SwapFloat           func(float64, float64) (float64, float64)  `name:"swap"`
			Swap                func(interface{}, interface{}) (interface{}, interface{})
			GetUserList         func() []testUser
		}

		func main() {
			client := hprose.NewClient("http://www.hprose.com/example/")
			var ro *RemoteObject
			client.UseService(&ro)

			// If an error occurs, it will panic
			fmt.Println(ro.Hello("World"))

			// If an error occurs, an error value will be returned
			if result, err := ro.HelloWithError("World"); err == nil {
				fmt.Println(result)
			} else {
				fmt.Println(err.Error())
			}

			// If an error occurs, it will be ignored
			result := ro.AsyncHello("World")
			fmt.Println(<-result)

			// If an error occurs, an error chan will be returned
			result, err := ro.AsyncHelloWithError("World")
			if e := <-err; e == nil {
				fmt.Println(<-result)
			} else {
				fmt.Println(e.Error())
			}
			fmt.Println(ro.Sum(1, 2, 3, 4, 5))

			m := make(map[string]string)
			m["Jan"] = "January"
			m["Feb"] = "February"
			m["Mar"] = "March"
			m["Apr"] = "April"
			m["May"] = "May"
			m["Jun"] = "June"
			m["Jul"] = "July"
			m["Aug"] = "August"
			m["Sep"] = "September"
			m["Oct"] = "October"
			m["Nov"] = "November"
			m["Dec"] = "December"

			fmt.Println(m)
			mm := ro.SwapKeyAndValue(&m)
			fmt.Println(m)
			fmt.Println(mm)

			fmt.Println(ro.GetUserList())
			fmt.Println(ro.SwapInt(1, 2))
			fmt.Println(ro.SwapFloat(1.2, 3.4))
			fmt.Println(ro.Swap("Hello", "World"))
		}
*/
package hprose

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"runtime/debug"
	"strings"
)

// InvokeOptions is the invoke options of hprose client
type InvokeOptions struct {
	ByRef      interface{} // true, false, nil
	SimpleMode interface{} // true, false, nil
	ResultMode ResultMode
}

// Client is hprose client
type Client interface {
	UseService(...interface{})
	Invoke(string, []interface{}, *InvokeOptions, interface{}) <-chan error
	Uri() string
	SetUri(string)
	GetFilter() Filter
	SetFilter(filter Filter)
	AddFilter(filter Filter)
	RemoveFilter(filter Filter)
	TLSClientConfig() *tls.Config
	SetTLSClientConfig(config *tls.Config)
	SetKeepAlive(enable bool)
	Close()
}

// ClientContext is the hprose client context
type ClientContext struct {
	*BaseContext
	Client
}

// Transporter is the hprose client transporter
type Transporter interface {
	SendAndReceive(uri string, data []byte) ([]byte, error)
}

// BaseClient is the hprose base client
type BaseClient struct {
	Transporter
	Client
	ByRef        bool
	SimpleMode   bool
	DebugEnabled bool
	uri          *url.URL
	filters      []Filter
}

var clientFactories = make(map[string]func(string) Client)

// NewBaseClient is the constructor of BaseClient
func NewBaseClient(trans Transporter) *BaseClient {
	client := new(BaseClient)
	client.Transporter = trans
	client.filters = make([]Filter, 0)
	return client
}

// NewClient is the constructor of Client
func NewClient(uri string) Client {
	if u, err := url.Parse(uri); err == nil {
		if newClient, ok := clientFactories[u.Scheme]; ok {
			return newClient(uri)
		}
		panic("The " + u.Scheme + "client isn't implemented.")
	} else {
		panic("The uri can't be parsed.")
	}
}

// Uri return the uri of hprose client
func (client *BaseClient) Uri() string {
	if client.uri != nil {
		return client.uri.String()
	}
	return ""
}

// SetUri set the uri of hprose client
func (client *BaseClient) SetUri(uri string) {
	if u, err := url.Parse(uri); err == nil {
		client.uri = u
	} else {
		panic("The uri can't be parsed.")
	}
}

// GetFilter return the first filter
func (client *BaseClient) GetFilter() Filter {
	if len(client.filters) == 0 {
		return nil
	}
	return client.filters[0]
}

// SetFilter set the only filter
func (client *BaseClient) SetFilter(filter Filter) {
	client.filters = make([]Filter, 0)
	if filter != nil {
		client.filters = append(client.filters, filter)
	}
}

// AddFilter add a filter
func (client *BaseClient) AddFilter(filter Filter) {
	client.filters = append(client.filters, filter)
}

// RemoveFilter remove a filter
func (client *BaseClient) RemoveFilter(filter Filter) {
	n := len(client.filters)
	for i := 0; i < n; i++ {
		if client.filters[i] == filter {
			if i == n-1 {
				client.filters = client.filters[:i]
			} else {
				client.filters = append(client.filters[:i], client.filters[i+1:]...)
			}
			return
		}
	}
}

// UseService (uri string)
// UseService (remoteObject interface{})
// UseService (uri string, remoteObject interface{})
// UseService (remoteObject interface{}, namespace string)
// UseService (uri string, remoteObject interface{}, namespace string)
func (client *BaseClient) UseService(args ...interface{}) {
	switch len(args) {
	case 1:
		switch arg0 := args[0].(type) {
		case nil:
			panic("The arguments can't be nil.")
		case string:
			client.SetUri(arg0)
			return
		case *string:
			client.SetUri(*arg0)
			return
		default:
			if isStructPointer(arg0) {
				client.createStub(arg0, "")
				return
			}
		}
	case 2:
		switch arg0 := args[0].(type) {
		case nil:
			panic("The arguments can't be nil.")
		case string:
			client.SetUri(arg0)
		case *string:
			client.SetUri(*arg0)
		default:
			switch arg1 := args[1].(type) {
			case nil:
				if isStructPointer(arg0) {
					client.createStub(arg0, "")
					return
				}
			case string:
				if isStructPointer(arg0) {
					client.createStub(arg0, arg1)
					return
				}
			case *string:
				if isStructPointer(arg0) {
					client.createStub(arg0, *arg1)
					return
				}
			}
			panic("Wrong arguments.")
		}
		if args[1] == nil {
			panic("The arguments can't be nil.")
		}
		if isStructPointer(args[1]) {
			client.createStub(args[1], "")
			return
		}
	case 3:
		switch arg0 := args[0].(type) {
		case nil:
			panic("The arguments can't be nil.")
		case string:
			client.SetUri(arg0)
		case *string:
			client.SetUri(*arg0)
		default:
			panic("Wrong arguments.")
		}
		if args[1] == nil {
			panic("The arguments can't be nil.")
		}
		if isStructPointer(args[1]) {
			switch arg2 := args[2].(type) {
			case nil:
				client.createStub(args[1], "")
				return
			case string:
				client.createStub(args[1], arg2)
				return
			case *string:
				client.createStub(args[1], *arg2)
				return
			}
		}
	}
	panic("Wrong arguments.")
}

// Invoke the remote method
func (client *BaseClient) Invoke(name string, args []interface{}, options *InvokeOptions, result interface{}) <-chan error {
	if result == nil {
		panic("The argument result can't be nil")
	}
	v := reflect.ValueOf(result)
	t := v.Type()
	if t.Kind() != reflect.Ptr {
		panic("The argument result must be pointer type")
	}
	r := []reflect.Value{v.Elem()}
	count := len(args)
	a := make([]reflect.Value, count)
	v = reflect.ValueOf(args)
	for i := 0; i < count; i++ {
		a[i] = v.Index(i).Elem()
	}
	return client.invoke(name, a, options, r)
}

// private methods

func (client *BaseClient) invoke(name string, args []reflect.Value, options *InvokeOptions, result []reflect.Value) <-chan error {
	if options == nil {
		options = new(InvokeOptions)
	}
	length := len(result)
	async := false
	for i := 0; i < length; i++ {
		if result[i].Kind() == reflect.Chan {
			async = true
		} else if async {
			panic("The out parameters must be all chan or all non-chan type.")
		}
	}
	byref := client.ByRef
	if br, ok := options.ByRef.(bool); ok {
		byref = br
	}
	if byref && !checkRefArgs(args) {
		panic("The elements in args must be pointer when options.ByRef is true.")
	}
	context := new(ClientContext)
	context.BaseContext = NewBaseContext()
	context.Client = client.Client
	if async {
		return client.asyncInvoke(name, args, options, result, context)
	}
	err := make(chan error, 1)
	err <- client.syncInvoke(name, args, options, result, context)
	return err
}

func (client *BaseClient) syncInvoke(name string, args []reflect.Value, options *InvokeOptions, result []reflect.Value, context *ClientContext) (err error) {
	defer func() {
		if e := recover(); e != nil && err == nil {
			if client.DebugEnabled {
				err = fmt.Errorf("%v\r\n%s", e, debug.Stack())
			} else {
				err = fmt.Errorf("%v", e)
			}
		}
	}()
	if odata, e := client.doOutput(name, args, options, context); e != nil {
		err = e
	} else if idata, e := client.SendAndReceive(client.Uri(), odata); e != nil {
		err = e
	} else if e := client.doIntput(idata, args, options, result, context); e != nil {
		err = e
	}
	return err
}

func (client *BaseClient) asyncInvoke(name string, args []reflect.Value, options *InvokeOptions, result []reflect.Value, context *ClientContext) <-chan error {
	length := len(result)
	sender := make([]reflect.Value, length)
	out := make([]reflect.Value, length)
	for i := 0; i < length; i++ {
		t := result[i].Type().Elem()
		out[i] = reflect.New(t).Elem()
		t = reflect.ChanOf(reflect.BothDir, t)
		sender[i] = reflect.MakeChan(t, 1)
		result[i].Set(sender[i])
	}
	errChan := make(chan error, 1)
	go func() {
		err := client.syncInvoke(name, args, options, out, context)
		for i := 0; i < length; i++ {
			sender[i].Send(out[i])
		}
		errChan <- err
	}()
	return errChan
}

func (client *BaseClient) doOutput(name string, args []reflect.Value, options *InvokeOptions, context *ClientContext) (data []byte, err error) {
	buf := new(bytes.Buffer)
	simple := client.SimpleMode
	if s, ok := options.SimpleMode.(bool); ok {
		simple = s
	}
	byref := client.ByRef
	if br, ok := options.ByRef.(bool); ok {
		byref = br
	}
	writer := NewWriter(buf, simple)
	if err = writer.Stream.WriteByte(TagCall); err != nil {
		return nil, err
	}
	if err = writer.WriteString(name); err != nil {
		return nil, err
	}
	if args != nil && (len(args) > 0 || byref) {
		writer.Reset()
		if err = writer.WriteArray(args); err != nil {
			return nil, err
		}
		if byref {
			if err = writer.WriteBool(true); err != nil {
				return nil, err
			}
		}
	}
	if err = writer.Stream.WriteByte(TagEnd); err != nil {
		return nil, err
	}
	data = buf.Bytes()
	n := len(client.filters)
	for i := 0; i < n; i++ {
		data = client.filters[i].OutputFilter(data, context)
	}
	return data, nil
}

func (client *BaseClient) doIntput(data []byte, args []reflect.Value, options *InvokeOptions, result []reflect.Value, context *ClientContext) (err error) {
	for i := len(client.filters) - 1; i >= 0; i-- {
		data = client.filters[i].InputFilter(data, context)
	}
	resultMode := options.ResultMode
	if last := len(data) - 1; data[last] == TagEnd {
		if resultMode == Raw {
			data = data[:last]
		}
	} else {
		return errors.New("Wrong Response: \r\n" + string(data))
	}
	if resultMode == RawWithEndTag || resultMode == Raw {
		if err = setResult(result[0], data); err != nil {
			return err
		}
		return nil
	}
	istream := NewBytesReader(data)
	reader := NewReader(istream, false)
	var tag byte
	for tag, err = istream.ReadByte(); err == nil && tag != TagEnd; tag, err = istream.ReadByte() {
		switch tag {
		case TagResult:
			switch resultMode {
			case Normal:
				reader.Reset()
				length := len(result)
				if length == 1 {
					err = reader.ReadValue(result[0])
				} else if err = reader.CheckTag(TagList); err == nil {
					var count int
					if count, err = reader.ReadInteger(TagOpenbrace); err == nil {
						r := make([]reflect.Value, count)
						if count <= length {
							for i := 0; i < count; i++ {
								r[i] = result[i]
							}
						} else {
							for i := 0; i < length; i++ {
								r[i] = result[i]
							}
							for i := length; i < count; i++ {
								var e interface{}
								r[i] = reflect.ValueOf(&e).Elem()
							}
						}
						err = reader.ReadArray(r)
					}
				}
				if err != nil {
					return err
				}
			case Serialized:
				var buf []byte
				if buf, err = reader.ReadRaw(); err != nil {
					return err
				}
				if err = setResult(result[0], buf); err != nil {
					return err
				}
			}
		case TagArgument:
			reader.Reset()
			if err = reader.CheckTag(TagList); err == nil {
				length := len(args)
				var count int
				if count, err = reader.ReadInteger(TagOpenbrace); err == nil {
					a := make([]reflect.Value, count)
					if count <= length {
						for i := 0; i < count; i++ {
							a[i] = args[i].Elem()
						}
					} else {
						for i := 0; i < length; i++ {
							a[i] = args[i].Elem()
						}
						for i := length; i < count; i++ {
							var e interface{}
							a[i] = reflect.ValueOf(&e).Elem()
						}
					}
					err = reader.ReadArray(a)
				}
			}
			if err != nil {
				return err
			}
		case TagError:
			reader.Reset()
			var e string
			if e, err = reader.ReadString(); err == nil {
				err = errors.New(e)
			}
			return err
		default:
			return errors.New("Wrong Response: \r\n" + string(data))
		}
	}
	return err
}

func (client *BaseClient) createStub(stub interface{}, ns string) {
	v := reflect.ValueOf(stub).Elem()
	t := v.Type()
	et := t
	if et.Kind() == reflect.Ptr {
		et = et.Elem()
	}
	objPointer := reflect.New(et)
	obj := objPointer.Elem()
	count := obj.NumField()
	for i := 0; i < count; i++ {
		f := obj.Field(i)
		ft := f.Type()
		if ft.Kind() == reflect.Func {
			f.Set(reflect.MakeFunc(ft, client.remoteMethod(ft, et.Field(i), ns)))
		} else if ft.Kind() == reflect.Ptr {
			ft = ft.Elem()
		}
		if ft.Kind() == reflect.Struct && f.CanSet() {
			fp := reflect.New(ft)
			sf := et.Field(i)
			if sf.Anonymous {
				client.createStub(fp.Interface(), ns)
			} else if ns == "" {
				client.createStub(fp.Interface(), sf.Name)
			} else {
				client.createStub(fp.Interface(), (ns + "_" + sf.Name))
			}
			if f.Kind() == reflect.Ptr {
				f.Set(fp)
			} else {
				f.Set(fp.Elem())
			}
		}
	}
	if t.Kind() == reflect.Ptr {
		v.Set(objPointer)
	} else {
		v.Set(obj)
	}
}

func (client *BaseClient) remoteMethod(t reflect.Type, sf reflect.StructField, ns string) func(in []reflect.Value) (out []reflect.Value) {
	name := getFuncName(&sf)
	if ns != "" {
		name = ns + "_" + name
	}
	options := &InvokeOptions{ByRef: getByRef(&sf), SimpleMode: getSimpleMode(&sf), ResultMode: getResultMode(&sf)}
	return func(in []reflect.Value) (out []reflect.Value) {
		inlen := len(in)
		varlen := 0
		argc := inlen
		if t.IsVariadic() {
			argc--
			varlen = in[argc].Len()
			argc += varlen
		}
		args := make([]reflect.Value, argc)
		if argc > 0 {
			for i := 0; i < inlen-1; i++ {
				args[i] = in[i]
			}
			if t.IsVariadic() {
				v := in[inlen-1]
				for i := 0; i < varlen; i++ {
					args[inlen-1+i] = v.Index(i)
				}
			} else {
				args[inlen-1] = in[inlen-1]
			}
		}
		numout := t.NumOut()
		out = make([]reflect.Value, numout)
		switch numout {
		case 0:
			var result interface{}
			err := <-client.invoke(name, args, options, []reflect.Value{reflect.ValueOf(&result).Elem()})
			if err == nil {
				return out
			}
			panic(err.Error())
		case 1:
			rt0 := t.Out(0)
			if rt0.Kind() == reflect.Chan {
				if rt0.Elem().Kind() == reflect.Interface && rt0.Elem().Name() == "error" {
					var result chan interface{}
					err := client.invoke(name, args, options, []reflect.Value{reflect.ValueOf(&result).Elem()})
					out[0] = reflect.ValueOf(&err).Elem()
					return out
				}
				out[0] = reflect.New(rt0).Elem()
				client.invoke(name, args, options, out)
				return out
			}
			if rt0.Kind() == reflect.Interface && rt0.Name() == "error" {
				var result interface{}
				err := <-client.invoke(name, args, options, []reflect.Value{reflect.ValueOf(&result).Elem()})
				out[0] = reflect.ValueOf(&err).Elem()
				return out
			}
			out[0] = reflect.New(rt0).Elem()
			err := <-client.invoke(name, args, options, out)
			if err == nil {
				return out
			}
			panic(err.Error())
		default:
			last := numout - 1
			rtlast := t.Out(last)
			for i := 0; i < last; i++ {
				out[i] = reflect.New(t.Out(i)).Elem()
			}
			if rtlast.Kind() == reflect.Chan &&
				rtlast.Elem().Kind() == reflect.Interface &&
				rtlast.Elem().Name() == "error" {
				err := client.invoke(name, args, options, out[:last])
				out[last] = reflect.ValueOf(&err).Elem()
				return out
			}
			if rtlast.Kind() == reflect.Interface &&
				rtlast.Name() == "error" {
				err := <-client.invoke(name, args, options, out[:last])
				out[last] = reflect.ValueOf(&err).Elem()
				return out
			}
			out[last] = reflect.New(t.Out(last)).Elem()
			if t.Out(0).Kind() == reflect.Chan {
				client.invoke(name, args, options, out)
				return out
			}
			err := <-client.invoke(name, args, options, out)
			if err == nil {
				return out
			}
			panic(err.Error())
		}
	}
}

// public functions

// RegisterClientFactory register client factory
func RegisterClientFactory(scheme string, newClient func(string) Client) {
	clientFactories[strings.ToLower(scheme)] = newClient
}

// private functions

func isStructPointer(p interface{}) bool {
	v := reflect.ValueOf(p)
	if !v.IsValid() || v.IsNil() {
		return false
	}
	t := v.Type()
	return t.Kind() == reflect.Ptr && (t.Elem().Kind() == reflect.Struct ||
		(t.Elem().Kind() == reflect.Ptr && t.Elem().Elem().Kind() == reflect.Struct))
}

func checkRefArgs(args []reflect.Value) bool {
	count := len(args)
	for i := 0; i < count; i++ {
		if args[i].Kind() != reflect.Ptr {
			return false
		}
	}
	return true
}

func setResult(result reflect.Value, buf []byte) error {
	switch result.Interface().(type) {
	case []byte, interface{}:
		result.Set(reflect.ValueOf(buf))
	case *bytes.Buffer:
		result.Set(reflect.ValueOf(bytes.NewBuffer(buf)))
	default:
		return errors.New("The argument result must be a *[]byte or **bytes.Buffer if the ResultMode is different from Normal.")
	}
	return nil
}

func getFuncName(sf *reflect.StructField) string {
	keys := []string{"name", "Name", "funcname", "funcName", "FuncName"}
	for i := range keys {
		if name := sf.Tag.Get(keys[i]); name != "" {
			return name
		}
	}
	return sf.Name
}

func getByRef(sf *reflect.StructField) interface{} {
	keys := []string{"byref", "byRef", "Byref", "ByRef"}
	for i := range keys {
		switch strings.ToLower(sf.Tag.Get(keys[i])) {
		case "true", "t", "1":
			return true
		case "false", "f", "0":
			return false
		}
	}
	return nil
}

func getSimpleMode(sf *reflect.StructField) interface{} {
	keys := []string{"simple", "Simple", "simpleMode", "SimpleMode"}
	for i := range keys {
		switch strings.ToLower(sf.Tag.Get(keys[i])) {
		case "true", "t", "1":
			return true
		case "false", "f", "0":
			return false
		}
	}
	return nil
}

func getResultMode(sf *reflect.StructField) ResultMode {
	keys := []string{"result", "Result", "resultMode", "ResultMode"}
	for i := range keys {
		switch strings.ToLower(sf.Tag.Get(keys[i])) {
		case "normal":
			return Normal
		case "serialized":
			return Serialized
		case "raw":
			return Raw
		case "rawwithendtag":
			return RawWithEndTag
		}
	}
	return Normal
}

func init() {
	RegisterClientFactory("http", newHttpClient)
	RegisterClientFactory("https", newHttpClient)
	RegisterClientFactory("tcp", newTcpClient)
	RegisterClientFactory("tcp4", newTcpClient)
	RegisterClientFactory("tcp6", newTcpClient)
	RegisterClientFactory("unix", newUnixClient)
	RegisterClientFactory("ws", newWebSocketClient)
	RegisterClientFactory("wss", newWebSocketClient)
}
