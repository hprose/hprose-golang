/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/service_codec.go                                |
|                                                          |
| LastModified: Feb 18, 2024                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"errors"
	"reflect"

	"github.com/hprose/hprose-golang/v3/io"
)

// ServiceCodec for RPC.
type ServiceCodec interface {
	Encode(result interface{}, context *ServiceContext) (response []byte, err error)
	Decode(request []byte, context *ServiceContext) (name string, args []interface{}, err error)
}

type serviceCodec struct {
	Debug  bool
	Simple bool
	io.LongType
	io.RealType
	io.MapType
	io.StructType
	io.ListType
}

// Encode response.
func (c serviceCodec) Encode(result interface{}, context *ServiceContext) ([]byte, error) {
	encoder := io.GetEncoder().Simple(c.Simple)
	defer io.FreeEncoder(encoder)
	if c.Simple {
		context.ResponseHeaders().Set("simple", true)
	}
	if context.HasResponseHeaders() {
		encoder.WriteTag(io.TagHeader)
		_ = encoder.Write(context.ResponseHeaders().ToMap())
		encoder.Reset()
	}
	if e, ok := result.(error); ok {
		encoder.WriteTag(io.TagError)
		var msg string
		if pe, ok := e.(*PanicError); ok && c.Debug {
			msg = pe.String()
		} else {
			msg = e.Error()
		}
		encoder.WriteString(msg)
	} else {
		encoder.WriteTag(io.TagResult)
		_ = encoder.Write(result)
	}
	encoder.WriteTag(io.TagEnd)
	return encoder.Bytes(), encoder.Error
}

// Decode request.
func (c serviceCodec) Decode(request []byte, context *ServiceContext) (name string, args []interface{}, err error) {
	if len(request) == 0 {
		name = "~"
		err = c.decodeMethod(name, context)
		return
	}
	decoder := io.GetDecoder().ResetBytes(request)
	defer io.FreeDecoder(decoder)
	decoder.LongType = c.LongType
	decoder.RealType = c.RealType
	decoder.MapType = c.MapType
	decoder.StructType = c.StructType
	decoder.ListType = c.ListType
	tag := decoder.NextByte()
	if tag == io.TagHeader {
		var h map[string]interface{}
		decoder.Decode(&h)
		NewDict(h).CopyTo(context.RequestHeaders())
		decoder.Reset()
		tag = decoder.NextByte()
	}
	switch tag {
	case io.TagCall:
		if context.RequestHeaders().GetBool("simple") {
			decoder.Simple(true)
		}
		decoder.Decode(&name)
		if err = c.decodeMethod(name, context); err == nil {
			args, err = c.decodeArguments(context.Method, decoder)
		}
	case io.TagEnd:
		name = "~"
		err = c.decodeMethod("~", context)
	default:
		err = InvalidRequestError{request}
	}
	return
}

func (c serviceCodec) decodeMethod(name string, context *ServiceContext) (err error) {
	if context.Method = context.Service().Get(name); context.Method == nil {
		err = errors.New("Can't find this method " + name + "().")
	}
	return err
}

func (c serviceCodec) decodeArguments(method Method, decoder *io.Decoder) (args []interface{}, err error) {
	tag := decoder.NextByte()
	if tag != io.TagList {
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
		for i := n - 1; i < count; i++ {
			paramTypes[i] = parameters[n-1].Elem()
		}
	} else {
		copy(paramTypes, parameters)
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
func NewServiceCodec(options ...CodecOption) ServiceCodec {
	c := serviceCodec{}
	for _, option := range options {
		option(&c)
	}
	return c
}
