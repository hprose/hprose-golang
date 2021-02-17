/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/service.go                                      |
|                                                          |
| LastModified: Feb 18, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"context"
	"reflect"
	"sync"
)

// Server is a generic interface used to represent any server.
type Server interface{}

// Handler is an interface used to bind service to any server.
type Handler interface {
	Bind(server Server)
}

// HandlerFactory is a constructor for Handler.
type HandlerFactory interface {
	ServerTypes() []reflect.Type
	New(service *Service) Handler
}

var handlerFactories sync.Map
var serverTypes sync.Map

// RegisterHandler for Service.
func RegisterHandler(name string, handlerFactory HandlerFactory) {
	handlerFactories.Store(name, handlerFactory)
	for _, serverType := range handlerFactory.ServerTypes() {
		if value, loaded := serverTypes.LoadOrStore(serverType, []string{name}); loaded {
			names := value.([]string)
			names = append(names, name)
			serverTypes.Store(serverType, names)
		}
	}
}

// Service for RPC.
type Service struct {
	Codec            ServiceCodec
	MaxRequestLength int
	Options          Dict
	invokeManager    PluginManager
	ioManager        PluginManager
	handlers         map[string]Handler
	methodManager
}

// NewService returns an instance of Service.
func NewService() *Service {
	service := &Service{
		Codec:            serviceCodec{},
		MaxRequestLength: 0x7FFFFFFF,
		Options:          NewSafeDict(),
	}
	handlerFactories.Range(func(key, value interface{}) bool {
		handler := value.(HandlerFactory).New(service)
		service.handlers[key.(string)] = handler
		return true
	})
	service.ioManager = NewIOManager(service.Process)
	service.invokeManager = NewInvokeManager(service.Execute)
	service.AddFunction(service.methodManager.Names, "~")
	return service
}

// Bind to server.
func (s *Service) Bind(server Server) error {
	serverType := reflect.TypeOf(server)
	if value, ok := serverTypes.Load(serverType); ok {
		names := value.([]string)
		for _, name := range names {
			s.handlers[name].Bind(server)
		}
	}
	return UnsupportedServerTypeError{serverType}
}

// Handler returns the handler by the specified name.
func (s *Service) Handler(name string) Handler {
	return s.handlers[name]
}

// Handle the reqeust and returns the response.
func (s *Service) Handle(ctx context.Context, request []byte) ([]byte, error) {
	return s.ioManager.Handler().(NextIOHandler)(ctx, request)
}

// Process the reqeust and returns the response.
func (s *Service) Process(ctx context.Context, request []byte) ([]byte, error) {
	serviceContext := GetServiceContext(ctx)
	name, args, err := s.Codec.Decode(request, serviceContext)
	if err != nil {
		return s.Codec.Encode(err, serviceContext)
	}
	var result interface{}
	func() {
		defer func() {
			if p := recover(); p != nil {
				result = NewPanicError(p)
			}
		}()
		results, err := s.invokeManager.Handler().(NextInvokeHandler)(ctx, name, args)
		if err != nil {
			result = err
			return
		}
		switch len(results) {
		case 0:
			result = nil
		case 1:
			result = results[0]
		default:
			result = results
		}
	}()
	return s.Codec.Encode(result, serviceContext)
}

// Execute the method and returns the results.
func (s *Service) Execute(ctx context.Context, name string, args []interface{}) (result []interface{}, err error) {
	serviceContext := GetServiceContext(ctx)
	method := serviceContext.Method
	if method.Missing() {
		if method.PassContext() {
			return method.(contextMissingMethod)(ctx, name, args)
		}
		return method.(missingMethod)(name, args)
	}
	n := len(args)
	var in []reflect.Value
	if method.PassContext() {
		in = make([]reflect.Value, n+1)
		in[0] = reflect.ValueOf(ctx)
		for i := 0; i < n; i++ {
			in[i+1] = reflect.ValueOf(args[i])
		}
	} else {
		in = make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			in[i] = reflect.ValueOf(args[i])
		}
	}
	f := method.Func()
	out := f.Call(in)
	n = len(out)
	if f.Type().Out(n - 1).Implements(errorType) {
		if !out[n-1].IsNil() {
			err = out[n-1].Interface().(error)
		}
		out = out[:n-1]
		n--
	}
	for i := 0; i < n; i++ {
		result = append(result, out[i].Interface())
	}
	return
}

// Use plugin handlers.
func (s *Service) Use(handler ...PluginHandler) *Service {
	invokeHandlers, ioHandler := separatePluginHandlers(handler)
	if len(invokeHandlers) > 0 {
		s.invokeManager.Use(invokeHandlers...)
	}
	if len(ioHandler) > 0 {
		s.ioManager.Use(ioHandler...)
	}
	return s
}

// Unuse plugin handlers.
func (s *Service) Unuse(handler ...PluginHandler) *Service {
	invokeHandlers, ioHandler := separatePluginHandlers(handler)
	if len(invokeHandlers) > 0 {
		s.invokeManager.Unuse(invokeHandlers...)
	}
	if len(ioHandler) > 0 {
		s.ioManager.Unuse(ioHandler...)
	}
	return s
}

// Remove is used for unpublishing method by the specified name.
func (s *Service) Remove(name string) *Service {
	s.methodManager.Remove(name)
	return s
}

// Add is used for publishing the method.
func (s *Service) Add(method Method) *Service {
	s.methodManager.Add(method)
	return s
}

// AddFunction is used for publishing function f with name.
func (s *Service) AddFunction(f interface{}, name string) *Service {
	s.methodManager.AddFunction(f, name)
	return s
}

// AddMethod is used for publishing method named name on target with alias.
func (s *Service) AddMethod(name string, target interface{}, alias ...string) *Service {
	s.methodManager.AddMethod(name, target, alias...)
	return s
}

// AddMethods is used for publishing methods named names on target with namespace.
func (s *Service) AddMethods(names []string, target interface{}, namespace ...string) *Service {
	s.methodManager.AddMethods(names, target, namespace...)
	return s
}

// AddInstanceMethods is used for publishing all the public methods and func fields with namespace.
func (s *Service) AddInstanceMethods(target interface{}, namespace ...string) *Service {
	s.methodManager.AddInstanceMethods(target, namespace...)
	return s
}

// AddAllMethods will publish all methods and non-nil function fields on the
// obj self and on its anonymous or non-anonymous struct fields (or pointer to
// pointer ... to pointer struct fields). This is a recursive operation.
// So it's a pit, if you do not know what you are doing, do not step on.
func (s *Service) AddAllMethods(target interface{}, namespace ...string) *Service {
	s.methodManager.AddAllMethods(target, namespace...)
	return s
}

// AddMissingMethod is used for publishing a method,
// all methods not explicitly published will be redirected to this method.
func (s *Service) AddMissingMethod(f interface{}) *Service {
	s.methodManager.AddMissingMethod(f)
	return s
}

// AddNetRPCMethods is used for publishing methods defined for net/rpc.
func (s *Service) AddNetRPCMethods(rcvr interface{}, namespace ...string) *Service {
	s.methodManager.AddNetRPCMethods(rcvr, namespace...)
	return s
}
