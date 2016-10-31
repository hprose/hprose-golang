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
 * rpc/base_service.go                                    *
 *                                                        *
 * hprose base service for Go.                            *
 *                                                        *
 * LastModified: Oct 29, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/hprose/hprose-golang/io"
	"github.com/hprose/hprose-golang/util"
)

type baseService struct {
	methodManager
	handlerManager
	filterManager
	FixArguments func(args []reflect.Value, context ServiceContext)
	Event        ServiceEvent
	Debug        bool
	Timeout      time.Duration
	Heartbeat    time.Duration
	ErrorDelay   time.Duration
	UserData     map[string]interface{}
	topics       map[string]*topic
	topicLock    sync.RWMutex
}

func defaultFixArguments(args []reflect.Value, context ServiceContext) {
	i := len(args) - 1
	typ := args[i].Type()
	if typ == interfaceType || typ == contextType || typ == serviceContextType {
		args[i] = reflect.ValueOf(context)
	}
}

func (service *baseService) initBaseService() {
	service.initMethodManager()
	service.initHandlerManager()
	service.Timeout = 120 * time.Second
	service.Heartbeat = 3 * time.Second
	service.ErrorDelay = 10 * time.Second
	service.topics = make(map[string]*topic)
	service.AddFunction("#", util.UUIDv4, Options{Simple: true})
	service.override.invokeHandler = func(
		name string, args []reflect.Value,
		context Context) (results []reflect.Value, err error) {
		return invoke(name, args, context.(ServiceContext))
	}
	service.override.beforeFilterHandler = func(
		request []byte, context Context) (response []byte, err error) {
		return service.beforeFilter(request, context.(ServiceContext))
	}
	service.override.afterFilterHandler = func(
		request []byte, context Context) (response []byte, err error) {
		return service.afterFilter(request, context.(ServiceContext))
	}
}

// AddFunction publish a func or bound method
// name is the method name
// function is a func or bound method
// options includes Mode, Simple, Oneway and NameSpace
func (service *baseService) AddFunction(name string, function interface{}, option ...Options) Service {
	service.methodManager.AddFunction(name, function, option...)
	return service
}

// AddFunctions is used for batch publishing service method
func (service *baseService) AddFunctions(names []string, functions []interface{}, option ...Options) Service {
	service.methodManager.AddFunctions(names, functions, option...)
	return service
}

// AddMethod is used for publishing a method on the obj with an alias
func (service *baseService) AddMethod(name string, obj interface{}, alias string, option ...Options) Service {
	service.methodManager.AddMethod(name, obj, alias, option...)
	return service
}

// AddMethods is used for batch publishing methods on the obj with aliases
func (service *baseService) AddMethods(names []string, obj interface{}, aliases []string, option ...Options) Service {
	service.methodManager.AddMethods(names, obj, aliases, option...)
	return service
}

// AddInstanceMethods is used for publishing all the public methods and func fields with options.
func (service *baseService) AddInstanceMethods(obj interface{}, option ...Options) Service {
	service.methodManager.AddInstanceMethods(obj, option...)
	return service
}

// AddAllMethods will publish all methods and non-nil function fields on the
// obj self and on its anonymous or non-anonymous struct fields (or pointer to
// pointer ... to pointer struct fields). This is a recursive operation.
// So it's a pit, if you do not know what you are doing, do not step on.
func (service *baseService) AddAllMethods(obj interface{}, option ...Options) Service {
	service.methodManager.AddAllMethods(obj, option...)
	return service
}

// AddMissingMethod is used for publishing a method,
// all methods not explicitly published will be redirected to this method.
func (service *baseService) AddMissingMethod(method MissingMethod, option ...Options) Service {
	service.methodManager.AddMissingMethod(method, option...)
	return service
}

// AddNetRPCMethods is used for publishing methods defined for net/rpc.
func (service *baseService) AddNetRPCMethods(rcvr interface{}, option ...Options) Service {
	service.methodManager.AddNetRPCMethods(rcvr, option...)
	return service
}

// Remove the published func or method by name
func (service *baseService) Remove(name string) Service {
	service.methodManager.Remove(name)
	return service
}

// SetFilter will replace the current filter settings
func (service *baseService) SetFilter(filter ...Filter) Service {
	service.filterManager.SetFilter(filter...)
	return service
}

// AddFilter add the filter to this Service
func (service *baseService) AddFilter(filter ...Filter) Service {
	service.filterManager.AddFilter(filter...)
	return service
}

// RemoveFilterByIndex remove the filter by the index
func (service *baseService) RemoveFilterByIndex(index int) Service {
	service.filterManager.RemoveFilterByIndex(index)
	return service
}

// RemoveFilter remove the filter from this Service
func (service *baseService) RemoveFilter(filter ...Filter) Service {
	service.filterManager.RemoveFilter(filter...)
	return service
}

// AddInvokeHandler add the invoke handler to this Service
func (service *baseService) AddInvokeHandler(handler ...InvokeHandler) Service {
	service.handlerManager.AddInvokeHandler(handler...)
	return service
}

// AddBeforeFilterHandler add the filter handler before filters
func (service *baseService) AddBeforeFilterHandler(handler ...FilterHandler) Service {
	service.handlerManager.AddBeforeFilterHandler(handler...)
	return service
}

// AddAfterFilterHandler add the filter handler after filters
func (service *baseService) AddAfterFilterHandler(handler ...FilterHandler) Service {
	service.handlerManager.AddAfterFilterHandler(handler...)
	return service
}

// SetUserData for service
func (service *baseService) SetUserData(
	userdata map[string]interface{}) Service {
	service.UserData = userdata
	return service
}

func callService(
	name string, args []reflect.Value,
	context ServiceContext) (results []reflect.Value, err error) {
	remoteMethod := context.Method()
	function := remoteMethod.Function
	if context.IsMissingMethod() {
		missingMethod := function.Interface().(MissingMethod)
		return missingMethod(name, args, context), nil
	}
	ft := function.Type()
	if !ft.IsVariadic() {
		count := len(args)
		n := ft.NumIn()
		if n < count {
			args = args[:n]
		}
	}
	results = function.Call(args)
	n := ft.NumOut()
	if n == 0 {
		return results, nil
	}
	if ft.Out(n-1) == errorType {
		err, _ = results[n-1].Interface().(error)
		results = results[:n-1]
	}
	return results, err
}

func doOutput(
	args []reflect.Value,
	results []reflect.Value,
	context ServiceContext) []byte {
	method := context.Method()
	writer := io.NewWriter(method.Simple)
	switch method.Mode {
	case RawWithEndTag:
		return results[0].Bytes()
	case Raw:
		writer.Write(results[0].Bytes())
	default:
		writer.WriteByte(io.TagResult)
		if method.Mode == Serialized {
			writer.Write(results[0].Bytes())
		} else {
			switch len(results) {
			case 0:
				writer.WriteNil()
			case 1:
				writer.WriteValue(results[0])
			default:
				writer.WriteSlice(results)
			}
		}
		if context.ByRef() {
			writer.WriteByte(io.TagArgument)
			writer.Reset()
			writer.WriteSlice(args)
		}
	}
	return writer.Bytes()
}

func getErrorMessage(err error, debug bool) string {
	if panicError, ok := err.(*PanicError); ok {
		if debug {
			return fmt.Sprintf("%v\r\n%s", panicError.Panic, panicError.Stack)
		}
		return panicError.Error()
	}
	return err.Error()
}

func fireErrorEvent(event ServiceEvent, e error, context Context) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = NewPanicError(e)
		}
	}()
	err = e
	switch event := event.(type) {
	case sendErrorEvent:
		event.OnSendError(err, context)
	case sendErrorEvent2:
		err = event.OnSendError(err, context)
	}
	return err
}

func fireBeforeInvokeEvent(
	event ServiceEvent,
	name string,
	args []reflect.Value,
	context ServiceContext) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = NewPanicError(e)
		}
	}()
	switch event := event.(type) {
	case beforeInvokeEvent:
		event.OnBeforeInvoke(name, args, context.ByRef(), context)
	case beforeInvokeEvent2:
		err = event.OnBeforeInvoke(name, args, context.ByRef(), context)
	}
	return err
}

func fireAfterInvokeEvent(
	event ServiceEvent,
	name string,
	args []reflect.Value,
	results []reflect.Value,
	context ServiceContext) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = NewPanicError(e)
		}
	}()
	switch event := event.(type) {
	case afterInvokeEvent:
		event.OnAfterInvoke(name, args, context.ByRef(), results, context)
	case afterInvokeEvent2:
		err = event.OnAfterInvoke(name, args, context.ByRef(), results, context)
	}
	return err
}

func (service *baseService) sendError(err error, context Context) []byte {
	err = fireErrorEvent(service.Event, err, context)
	w := io.NewWriter(true)
	w.WriteByte(io.TagError)
	w.WriteString(getErrorMessage(err, service.Debug))
	return w.Bytes()
}

func (service *baseService) endError(err error, context Context) []byte {
	w := io.NewByteWriter(service.sendError(err, context))
	w.WriteByte(io.TagEnd)
	return w.Bytes()
}

func invoke(
	name string,
	args []reflect.Value,
	context ServiceContext) ([]reflect.Value, error) {
	if context.Method() == nil {
		return nil, errors.New("Can't find this method " + name)
	}
	if context.Method().Oneway {
		go func() {
			defer func() {
				recover()
			}()
			callService(name, args, context)
		}()
		return nil, nil
	}
	return callService(name, args, context)
}

func readArguments(
	fixArguments func(args []reflect.Value, context ServiceContext),
	reader *io.Reader,
	method *Method,
	context ServiceContext) (args []reflect.Value) {
	if method != nil {
		reader.JSONCompatible = method.JSONCompatible
	}
	if method == nil || context.IsMissingMethod() {
		return reader.ReadSliceWithoutTag()
	}
	count := reader.ReadCount()
	ft := method.Function.Type()
	n := ft.NumIn()
	if ft.IsVariadic() {
		n--
	}
	max := util.Max(n, count)
	args = make([]reflect.Value, max)
	for i := 0; i < n; i++ {
		args[i] = reflect.New(ft.In(i)).Elem()
	}
	if n < count {
		if ft.IsVariadic() {
			for i := n; i < count; i++ {
				args[i] = reflect.New(ft.In(n).Elem()).Elem()
			}
		} else {
			for i := n; i < count; i++ {
				args[i] = reflect.New(interfaceType).Elem()
			}
		}
	}
	reader.ReadSlice(args[:count])
	if !ft.IsVariadic() && n > count {
		fixArguments(args, context)
	}
	return
}

func (service *baseService) beforeInvoke(
	name string,
	args []reflect.Value,
	context ServiceContext) ([]byte, error) {
	err := fireBeforeInvokeEvent(service.Event, name, args, context)
	if err != nil {
		return nil, err
	}
	var results []reflect.Value
	results, err = service.handlerManager.invokeHandler(name, args, context)
	if err != nil {
		return nil, err
	}
	err = fireAfterInvokeEvent(service.Event, name, args, results, context)
	if err != nil {
		return nil, err
	}
	return doOutput(args, results, context), nil
}

func mergeResult(results [][]byte) []byte {
	n := len(results)
	if n == 1 {
		return append(results[0], io.TagEnd)
	}
	writer := io.NewByteWriter(results[0])
	for i := 1; i < n; i++ {
		writer.Write(results[i])
	}
	writer.WriteByte(io.TagEnd)
	return writer.Bytes()
}

func (service *baseService) doSingleInvoke(
	reader *io.Reader, context ServiceContext) (result []byte, tag byte) {
	name := reader.ReadString()
	alias := strings.ToLower(name)
	method := service.RemoteMethods[alias]
	if method == nil {
		method = service.RemoteMethods["*"]
		context.setIsMissingMethod(true)
	}
	tag = reader.CheckTags([]byte{io.TagList, io.TagEnd, io.TagCall})
	var args []reflect.Value
	if tag == io.TagList {
		reader.Reset()
		args = readArguments(service.FixArguments, reader, method, context)
		tag = reader.CheckTags([]byte{io.TagTrue, io.TagEnd, io.TagCall})
		if tag == io.TagTrue {
			context.setByRef(true)
			tag = reader.CheckTags([]byte{io.TagEnd, io.TagCall})
		}
	}
	context.setMethod(method)
	result, err := service.beforeInvoke(name, args, context)
	if err != nil {
		return service.sendError(err, context), tag
	}
	return result, tag
}

func (service *baseService) doInvoke(
	reader *io.Reader,
	context ServiceContext) []byte {
	var results [][]byte
	for {
		result, tag := service.doSingleInvoke(reader, context)
		results = append(results, result)
		if tag != io.TagCall {
			break
		}
		reader.Reset()
	}
	return mergeResult(results)
}

func (service *baseService) doFunctionList(context ServiceContext) []byte {
	writer := io.NewWriter(true)
	writer.WriteByte(io.TagFunctions)
	writer.WriteStringSlice(service.MethodNames)
	writer.WriteByte(io.TagEnd)
	return writer.Bytes()
}

func (service *baseService) afterFilter(
	request []byte, context ServiceContext) ([]byte, error) {
	reader := defaultReaderPool.acquireReader(request)
	defer defaultReaderPool.releaseReader(reader)
	reader.Init(request)
	tag, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}
	switch tag {
	case io.TagCall:
		return service.doInvoke(reader, context), nil
	case io.TagEnd:
		return service.doFunctionList(context), nil
	default:
		return nil, fmt.Errorf("Wrong Request: \r\n%s", request)
	}
}

func (service *baseService) delayError(
	err error, context ServiceContext) []byte {
	response := service.endError(err, context)
	if service.ErrorDelay > 0 {
		time.Sleep(service.ErrorDelay)
	}
	return response
}

func (service *baseService) beforeFilter(
	request []byte, context ServiceContext) (response []byte, err error) {
	request = service.inputFilter(request, context)
	response, err = service.afterFilterHandler(request, context)
	if err != nil {
		response = service.delayError(err, context)
	}
	return service.outputFilter(response, context), nil
}

// Handle the hprose request and return the hprose response
func (service *baseService) Handle(request []byte, context Context) []byte {
	if service.UserData != nil {
		for k, v := range service.UserData {
			context.SetInterface(k, v)
		}
	}
	response, err := service.beforeFilterHandler(request, context)
	if err != nil {
		return service.endError(err, context)
	}
	return response
}

func fireSubscribeEvent(
	topic string,
	id string,
	service *baseService) {
	defer func() {
		recover()
	}()
	if event, ok := service.Event.(subscribeEvent); ok {
		event.OnSubscribe(topic, id, service)
	}
}

func fireUnsubscribeEvent(
	topic string,
	id string,
	service *baseService) {
	defer func() {
		recover()
	}()
	if event, ok := service.Event.(unsubscribeEvent); ok {
		event.OnUnsubscribe(topic, id, service)
	}
}

func (service *baseService) offline(t *topic, topic string, id string) {
	if t.exist(id) {
		t.remove(id)
		fireUnsubscribeEvent(topic, id, service)
	}
}

// Publish the hprose push topic
func (service *baseService) Publish(
	topic string,
	timeout time.Duration,
	heartbeat time.Duration) Service {
	if timeout <= 0 {
		timeout = service.Timeout
	}
	if heartbeat <= 0 {
		heartbeat = service.Heartbeat
	}
	t := newTopic(heartbeat)
	service.topicLock.Lock()
	service.topics[topic] = t
	service.topicLock.Unlock()
	return service.AddFunction(topic, func(id string) interface{} {
		message := t.get(id)
		if message == nil {
			message = make(chan interface{})
			t.put(id, message)
			fireSubscribeEvent(topic, id, service)
		}
		select {
		case result := <-message:
			return result
		case <-time.After(timeout):
			service.offline(t, topic, id)
			return nil
		}
	}, Options{})
}

func (service *baseService) getTopic(topic string) (t *topic) {
	service.topicLock.RLock()
	t = service.topics[topic]
	service.topicLock.RUnlock()
	if t == nil {
		panic("topic \"" + topic + "\" is not published.")
	}
	return
}

func (service *baseService) unicast(
	t *topic, topic string, id string, result interface{}, callback func(bool)) {
	message := t.get(id)
	if message == nil {
		if callback != nil {
			callback(false)
		}
		return
	}
	go func() {
		select {
		case message <- result:
			if callback != nil {
				callback(true)
			}
			break
		case <-time.After(t.heartbeat):
			service.offline(t, topic, id)
			if callback != nil {
				callback(false)
			}
		}
	}()
}

// IDList returns the push client id list
func (service *baseService) IDList(topic string) []string {
	return service.getTopic(topic).idlist()
}

// Exist returns true if the client id exist.
func (service *baseService) Exist(topic string, id string) bool {
	return service.getTopic(topic).exist(id)
}

// Push result to clients
func (service *baseService) Push(topic string, result interface{}, id ...string) {
	t := service.getTopic(topic)
	n := len(id)
	if n == 0 {
		id = t.idlist()
		n = len(id)
		if n == 0 {
			return
		}
	}
	for i := 0; i < n; i++ {
		service.unicast(t, topic, id[i], result, nil)
	}
}

// Broadcast push result to all clients
func (service *baseService) Broadcast(
	topic string, result interface{}, callback func([]string)) {
	service.Multicast(topic, service.IDList(topic), result, callback)
}

// Multicast result to the specified clients
func (service *baseService) Multicast(
	topic string, ids []string, result interface{}, callback func([]string)) {
	t := service.getTopic(topic)
	m := 0
	n := len(ids)
	if n == 0 {
		callback(nil)
		return
	}
	sid := make(chan string)
	go func() {
		sended := make([]string, 0, n)
		for id := range sid {
			sended = append(sended, id)
		}
		callback(sended)
	}()
	for i := 0; i < n; i++ {
		id := ids[i]
		service.unicast(t, topic, id, result, func(ok bool) {
			if ok {
				sid <- id
			}
			m++
			if m == n {
				close(sid)
			}
		})
	}
}

// Unicast result to then specified client
func (service *baseService) Unicast(
	topic string, id string, result interface{}, callback func(bool)) {
	service.unicast(service.getTopic(topic), topic, id, result, callback)
}
