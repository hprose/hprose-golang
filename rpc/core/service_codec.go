/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/service_codec.go                                |
|                                                          |
| LastModified: Jan 24, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"github.com/hprose/hprose-golang/v3/encoding"
)

// ServiceCodec for RPC.
type ServiceCodec interface {
	Encode(result interface{}, context ServiceContext) (response []byte, err error)
	Decode(request []byte, context ClientContext) (name string, args []interface{}, err error)
}

type serviceCodec struct {
	Debug  bool
	Simple bool
	encoding.LongType
	encoding.RealType
	encoding.MapType
}

func (c serviceCodec) Encode(result interface{}, context ServiceContext) (response []byte, err error) {
	encoder := new(encoding.Encoder)
	encoder.Simple(c.Simple)
	if c.Simple {
		context.RequestHeaders().Set("simple", true)
	}
	if context.HasRequestHeaders() {
		encoder.WriteTag(encoding.TagHeader)
		encoder.Write((map[string]interface{})(context.RequestHeaders().(headers)))
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

// func (c serviceCodec) Decode(request []byte, context ClientContext) (name string, args []interface{}, err error) {
// 	if len(request) == 0 {
// 		c.decodeMethod("~", 0, context)
// 		return "~", []interface{}{}, err
// 	}
// 	return "", nil, nil
// }

// func (c serviceCodec) decodeMethod(string name, int paramCount, ServiceContext context) (method Method, err error) {
// 	service := context.Service()
// 	method := service.Get(name, paramCount)
// 	if method == nil {
// 		err = errors.New("Can't find this method " + name + "().")
// 	} else {
// 		context.SetMethod(method)
// 	}
// 	return method, err
// }
