/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/codec/jsonrpc/service_codec.go                       |
|                                                          |
| LastModified: May 10, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package jsonrpc

import (
	"reflect"

	"github.com/hprose/hprose-golang/v3/rpc/core"
)

type ServiceCodec struct {
	Codec
	defaultCodec core.ServiceCodec
}

func (c *ServiceCodec) Encode(result interface{}, context *core.ServiceContext) ([]byte, error) {
	if !context.Items().GetBool("jsonrpc") {
		return c.defaultCodec.Encode(result, context)
	}
	response := Response{
		JSONRPC: "2.0",
		ID:      context.Items().GetInt64("jsonrpc.id"),
	}
	if context.HasResponseHeaders() {
		response.Headers = context.ResponseHeaders().ToMap()
	}
	if e, ok := result.(error); ok {
		switch e := e.(type) {
		case *jsonrpcError:
			response.Error = &Error{
				Code:    e.Code,
				Message: e.Message,
			}
		case *core.PanicError:
			response.Error = &Error{
				Message: e.Error(),
				Data:    e.Stack,
			}
		case *core.InvalidRequestError:
			response.Error = &Error{
				Code:    codeInvalidRequest,
				Message: messageInvalidRequest,
			}
		default:
			response.Error = &Error{
				Message: e.Error(),
			}
		}
	} else {
		response.Result = result
	}
	return c.Codec.Marshal(response)
}

func (c *ServiceCodec) Decode(request []byte, context *core.ServiceContext) (name string, args []interface{}, err error) {
	if len(request) == 0 {
		name = "~"
		err = c.decodeMethod(name, context)
		return
	}
	if request[0] != '{' {
		return c.defaultCodec.Decode(request, context)
	}
	context.Items().Set("jsonrpc", true)
	var req Request
	if err = c.Codec.Unmarshal(request, &req); err != nil {
		err = &jsonrpcError{codeParseError, messageParseError}
		return
	}
	if req.JSONRPC != "2.0" || req.Method == "" {
		err = &jsonrpcError{codeInvalidRequest, messageInvalidRequest}
		return
	}
	context.Items().Set("jsonrpc.id", req.ID)
	if req.Headers != nil {
		core.NewDict(req.Headers).CopyTo(context.RequestHeaders())
	}
	name = req.Method
	if err = c.decodeMethod(name, context); err != nil {
		return
	}
	method := context.Method
	if method.Missing() {
		return name, req.Params, nil
	}
	count := len(req.Params)
	parameters := method.Parameters()
	paramTypes := make([]reflect.Type, count)
	if method.Func().Type().IsVariadic() {
		n := len(parameters)
		copy(paramTypes, parameters[:n-1])
		for i := n - 1; i < count; i++ {
			paramTypes[i] = parameters[n-1].Elem()
		}
	} else {
		copy(paramTypes, parameters)
	}
	args = make([]interface{}, count)
	for i, t := range paramTypes {
		data, _ := c.Codec.Marshal(req.Params[i])
		a := reflect.New(t)
		if err = c.Codec.Unmarshal(data, a.Interface()); err != nil {
			err = &jsonrpcError{codeInvalidParams, messageInvalidParams}
			return
		}
		args[i] = a.Elem().Interface()
	}
	return
}

func (c *ServiceCodec) decodeMethod(name string, context *core.ServiceContext) (err error) {
	if context.Method = context.Service().Get(name); context.Method == nil {
		err = &jsonrpcError{codeMethodNotFound, messageMethodNotFound}
	}
	return
}

// NewServiceCodec returns the ServiceCodec.
func NewServiceCodec(codec Codec, options ...core.CodecOption) core.ServiceCodec {
	if codec == nil {
		codec = jsonCodec{}
	}
	return &ServiceCodec{
		Codec:        codec,
		defaultCodec: core.NewServiceCodec(options...),
	}
}
