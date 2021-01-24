/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/client_codec.go                                 |
|                                                          |
| LastModified: Jan 24, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"unsafe"

	"github.com/hprose/hprose-golang/v3/encoding"
	"github.com/modern-go/reflect2"
)

// ClientCodec for RPC.
type ClientCodec interface {
	Encode(name string, args []interface{}, context ClientContext) (reqeust []byte, err error)
	Decode(response []byte, context ClientContext) (result interface{}, err error)
}

type clientCodec struct {
	Simple bool
	encoding.LongType
	encoding.RealType
	encoding.MapType
}

func (c clientCodec) Encode(name string, args []interface{}, context ClientContext) ([]byte, error) {
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
	encoder.Write(name)
	if len(args) > 0 {
		encoder.Reset()
		encoder.Write(args)
	}
	encoder.WriteTag(encoding.TagEnd)
	return encoder.Bytes(), encoder.Error
}

func (c clientCodec) Decode(response []byte, context ClientContext) (result interface{}, err error) {
	decoder := encoding.NewDecoder(response)
	decoder.Simple(false)
	decoder.LongType = c.LongType
	decoder.RealType = c.RealType
	decoder.MapType = c.MapType
	tag := decoder.NextByte()
	if tag == encoding.TagHeader {
		var h map[string]interface{}
		decoder.Decode(&h)
		((headers)(h)).CopyTo(context.ResponseHeaders())
		decoder.Reset()
		tag = decoder.NextByte()
	}
	switch tag {
	case encoding.TagResult:
		if simple, ok := context.ResponseHeaders().Get("simple"); ok && simple.(bool) {
			decoder.Simple(true)
		}
		returnType := context.ReturnType()
		n := len(returnType)
		switch n {
		case 0:
			decoder.Decode(&result)
		case 1:
			result = decoder.Read(returnType[0])
		default:
			results := make([]interface{}, n)
			tag = decoder.NextByte()
			count := 0
			if tag == encoding.TagList {
				count = decoder.ReadInt()
				decoder.AddReference(nil)
			}
			for i := 0; i < n; i++ {
				if i < count {
					results[i] = decoder.Read(returnType[i])
				} else {
					t := reflect2.Type2(returnType[i])
					results[i] = t.Indirect(t.New())
				}
			}
			result = results
		}
		err = decoder.Error
	case encoding.TagError:
		var errstr string
		decoder.Decode(&errstr)
		if decoder.Error != nil {
			err = decoder.Error
		} else {
			err = encoding.DecodeError(errstr)
		}
	case encoding.TagEnd:
	default:
		err = encoding.DecodeError("Invalid response\r\n" + *(*string)(unsafe.Pointer(&response)))
	}
	return
}

// NewClientCodec returns the ClientCodec.
func NewClientCodec(simple bool, longType encoding.LongType, realType encoding.RealType, mapType encoding.MapType) ClientCodec {
	return clientCodec{simple, longType, realType, mapType}
}

// DefaultClientCodec is the default ClientCodec
var DefaultClientCodec ClientCodec = clientCodec{}
