/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/service_codec.go                                |
|                                                          |
| LastModified: Feb 17, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"errors"
	"reflect"

	"github.com/hprose/hprose-golang/v3/encoding"
)

// ServiceCodec for RPC.
type ServiceCodec interface {
	Encode(result interface{}, context *ServiceContext) (response []byte, err error)
	Decode(request []byte, context *ServiceContext) (name string, args []interface{}, err error)
}

type serviceCodec struct {
	Debug  bool
	Simple bool
	encoding.LongType
	encoding.RealType
	encoding.MapType
}

func (c serviceCodec) Encode(result interface{}, context *ServiceContext) (response []byte, err error) {
	encoder := new(encoding.Encoder).Simple(c.Simple)
	if c.Simple {
		context.RequestHeaders().Set("simple", true)
	}
	if context.HasRequestHeaders() {
		encoder.WriteTag(encoding.TagHeader)
		encoder.Write((map[string]interface{})(context.RequestHeaders().(dict)))
		encoder.Reset()
	}
	encoder.WriteTag(encoding.TagCall)
	if e, ok := result.(error); ok {
		encoder.WriteTag(encoding.TagError)
		var msg string
		if pe, ok := e.(*PanicError); ok && c.Debug {
			msg = pe.String()
		} else {
			msg = err.Error()
		}
		encoder.WriteString(msg)
	} else {
		encoder.WriteTag(encoding.TagResult)
		encoder.Write(result)
	}
	encoder.WriteTag(encoding.TagEnd)
	return encoder.Bytes(), encoder.Error
}

func (c serviceCodec) Decode(request []byte, context *ServiceContext) (name string, args []interface{}, err error) {
	if len(request) == 0 {
		name = "~"
		_, err = c.decodeMethod(name, context)
		return
	}
	decoder := encoding.NewDecoder(request).Simple(false)
	decoder.LongType = c.LongType
	decoder.RealType = c.RealType
	decoder.MapType = c.MapType
	tag := decoder.NextByte()
	if tag == encoding.TagHeader {
		var h map[string]interface{}
		decoder.Decode(&h)
		((dict)(h)).CopyTo(context.ResponseHeaders())
		decoder.Reset()
		tag = decoder.NextByte()
	}
	switch tag {
	case encoding.TagCall:
		if context.RequestHeaders().GetBool("simple") {
			decoder.Simple(true)
		}
		decoder.Decode(&name)
		var method Method
		if method, err = c.decodeMethod(name, context); err == nil {
			args, err = c.decodeArguments(method, decoder, context)
		}
	case encoding.TagEnd:
		name = "~"
		_, err = c.decodeMethod("~", context)
	default:
		err = errors.New("Invalid request:\r\n" + string(request))
	}
	return
}

func (c serviceCodec) decodeMethod(name string, context *ServiceContext) (method Method, err error) {
	service := context.Service()
	method = service.Get(name)
	if method == nil {
		err = errors.New("Can't find this method " + name + "().")
	} else {
		context.Method = method
	}
	return method, err
}

func (c serviceCodec) decodeArguments(method Method, decoder *encoding.Decoder, context *ServiceContext) (args []interface{}, err error) {
	tag := decoder.NextByte()
	if tag != encoding.TagList {
		return
	}
	decoder.Reset()
	if method.Missing() {
		decoder.Decode(&args, tag)
		return args, decoder.Error
	}
	count := decoder.ReadInt()
	parameters := method.Parameters()
	paramTypes := make([]reflect.Type, count)
	if method.Func().Type().IsVariadic() {
		n := len(parameters)
		copy(paramTypes, parameters[:n-1])
		for i := n; i < count; i++ {
			paramTypes[i] = parameters[n-1].Elem()
		}
	} else {
		copy(paramTypes, parameters[:])
	}
	args = make([]interface{}, count)
	decoder.AddReference(&args)
	for i := 0; i < count; i++ {
		args[i] = decoder.Read(paramTypes[i])
	}
	decoder.Skip()
	return args, decoder.Error
}

// NewServiceCodec returns the ServiceCodec.
func NewServiceCodec(debug bool, simple bool, longType encoding.LongType, realType encoding.RealType, mapType encoding.MapType) ServiceCodec {
	return serviceCodec{debug, simple, longType, realType, mapType}
}
