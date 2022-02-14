/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/codec/jsonrpc/client_codec.go                        |
|                                                          |
| LastModified: Feb 14, 2022                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package jsonrpc

import (
	"errors"
	"sync/atomic"

	"github.com/hprose/hprose-golang/v3/rpc/core"
	"github.com/modern-go/reflect2"
)

type ClientCodec struct {
	Codec
	counter int64
}

func (c *ClientCodec) Encode(name string, args []interface{}, context *core.ClientContext) ([]byte, error) {
	id := atomic.AddInt64(&c.counter, 1) & int64(0x7fffffff)
	request := Request{
		JSONRPC: "2.0",
		ID:      id,
		Method:  name,
	}
	if context.HasRequestHeaders() {
		request.Headers = context.RequestHeaders().ToMap()
	}
	if len(args) > 0 {
		request.Params = args
	}
	return c.Codec.Marshal(request)
}

func (c *ClientCodec) Decode(response []byte, context *core.ClientContext) (result []interface{}, err error) {
	var resp Response
	if err = c.Codec.Unmarshal(response, &resp); err != nil {
		return
	}
	if resp.Headers != nil {
		core.NewDict(resp.Headers).CopyTo(context.ResponseHeaders())
	}
	if resp.Result != nil {
		switch n := len(context.ReturnType); n {
		case 0:
		case 1:
			data, _ := c.Codec.Marshal(resp.Result)
			t := reflect2.Type2(context.ReturnType[0])
			p := t.New()
			if err = c.Codec.Unmarshal(data, p); err != nil {
				return
			}
			result = []interface{}{t.Indirect(p)}
		default:
			res := resp.Result.([]interface{})
			result = make([]interface{}, 0, len(res))
			for i, r := range res {
				data, _ := c.Codec.Marshal(r)
				t := reflect2.Type2(context.ReturnType[i])
				p := t.New()
				if err = c.Codec.Unmarshal(data, p); err != nil {
					return
				}
				result = append(result, t.Indirect(p))
			}
		}
	}
	if e := resp.Error; e != nil {
		switch code, message, data := e.Code, e.Message, e.Data; {
		case code != 0:
			err = jsonrpcError{code, message}
		case data != nil:
			err = &core.PanicError{
				Panic: message,
				Stack: data,
			}
		default:
			err = errors.New(message)
		}
	}
	return
}

// NewClientCodec returns the ClientCodec.
func NewClientCodec(codec Codec) core.ClientCodec {
	if codec == nil {
		codec = jsonCodec{}
	}
	return &ClientCodec{Codec: codec}
}
