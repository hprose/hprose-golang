/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/codec/jsonrpc/common.go                              |
|                                                          |
| LastModified: May 10, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package jsonrpc

import "encoding/json"

type Codec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

type jsonCodec struct{}

func (jsonCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (jsonCodec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

type Request struct {
	JSONRPC string                 `json:"jsonrpc"`
	ID      int64                  `json:"id"`
	Headers map[string]interface{} `json:"headers,omitempty"`
	Method  string                 `json:"method"`
	Params  []interface{}          `json:"params,omitempty"`
}

type Error struct {
	Code    int64  `json:"code,omitempty"`
	Message string `json:"message"`
	Data    []byte `json:"data,omitempty"`
}

type Response struct {
	JSONRPC string                 `json:"jsonrpc"`
	ID      int64                  `json:"id"`
	Headers map[string]interface{} `json:"headers,omitempty"`
	Result  interface{}            `json:"result,omitempty"`
	Error   *Error                 `json:"error,omitempty"`
}

const (
	messageParseError     = "Parse error"
	messageInvalidRequest = "Invalid Request"
	messageMethodNotFound = "Method not found"
	messageInvalidParams  = "Invalid params"

	codeParseError     = -32700
	codeInvalidRequest = -32600
	codeMethodNotFound = -32601
	codeInvalidParams  = -32602
)

type jsonrpcError struct {
	Code    int64
	Message string
}

func (e jsonrpcError) Error() string {
	return "hprose/rpc/codec/jsonrpc: " + e.Message
}
