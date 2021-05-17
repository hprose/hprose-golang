/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/client_codec.go                                 |
|                                                          |
| LastModified: May 10, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"github.com/hprose/hprose-golang/v3/io"
	"github.com/modern-go/reflect2"
)

// ClientCodec for RPC.
type ClientCodec interface {
	Encode(name string, args []interface{}, context *ClientContext) (reqeust []byte, err error)
	Decode(response []byte, context *ClientContext) (result []interface{}, err error)
}

type clientCodec struct {
	Simple bool
	io.LongType
	io.RealType
	io.MapType
}

func (c clientCodec) Encode(name string, args []interface{}, context *ClientContext) ([]byte, error) {
	encoder := new(io.Encoder).Simple(c.Simple)
	if c.Simple {
		context.RequestHeaders().Set("simple", true)
	}
	if context.HasRequestHeaders() {
		encoder.WriteTag(io.TagHeader)
		_ = encoder.Write(context.RequestHeaders().ToMap())
		encoder.Reset()
	}
	encoder.WriteTag(io.TagCall)
	_ = encoder.Write(name)
	if len(args) > 0 {
		encoder.Reset()
		_ = encoder.Write(args)
	}
	encoder.WriteTag(io.TagEnd)
	return encoder.Bytes(), encoder.Error
}

func (c clientCodec) Decode(response []byte, context *ClientContext) (result []interface{}, err error) {
	decoder := io.NewDecoder(response).Simple(false)
	decoder.LongType = c.LongType
	decoder.RealType = c.RealType
	decoder.MapType = c.MapType
	tag := decoder.NextByte()
	if tag == io.TagHeader {
		var h map[string]interface{}
		decoder.Decode(&h)
		NewDict(h).CopyTo(context.ResponseHeaders())
		decoder.Reset()
		tag = decoder.NextByte()
	}
	switch tag {
	case io.TagResult:
		if context.ResponseHeaders().GetBool("simple") {
			decoder.Simple(true)
		}
		returnType := context.ReturnType
		n := len(returnType)
		switch n {
		case 0:
			// Ignore the result to speed up.
		case 1:
			result = []interface{}{decoder.Read(returnType[0])}
		default:
			results := make([]interface{}, n)
			tag = decoder.NextByte()
			count := 1
			if tag == io.TagList {
				count = decoder.ReadInt()
				decoder.AddReference(nil)
				for i := 0; i < n && i < count; i++ {
					results[i] = decoder.Read(returnType[i])
				}
			} else {
				results[0] = decoder.Read(returnType[0], tag)
			}
			for i := count; i < n; i++ {
				t := reflect2.Type2(returnType[i])
				results[i] = t.Indirect(t.New())
			}
			result = results
		}
		err = decoder.Error
	case io.TagError:
		var errstr string
		decoder.Decode(&errstr)
		switch {
		case decoder.Error != nil:
			err = decoder.Error
		case errstr == "timeout":
			err = ErrTimeout
		default:
			err = io.DecodeError(errstr)
		}
	case io.TagEnd:
	default:
		err = InvalidResponseError{response}
	}
	return
}

// NewClientCodec returns the ClientCodec.
func NewClientCodec(options ...CodecOption) ClientCodec {
	c := clientCodec{}
	for _, option := range options {
		option(&c)
	}
	return c
}
