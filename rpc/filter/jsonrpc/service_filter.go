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
 * rpc/filter/jsonrpc/service_filter.go                   *
 *                                                        *
 * hprose jsonrpc service filter for Go.                  *
 *                                                        *
 * LastModified: Nov 1, 2016                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package jsonrpc

import (
	"encoding/json"

	"github.com/hprose/hprose-golang/io"
	"github.com/hprose/hprose-golang/rpc"
)

// ServiceFilter is a JSONRPC Service Filter
type ServiceFilter struct{}

func createResponse(request map[string]interface{}) (response map[string]interface{}) {
	response = make(map[string]interface{})
	if id, ok := request["id"]; ok {
		response["id"] = id
	} else {
		response["id"] = nil
	}
	if version, ok := request["jsonrpc"]; ok {
		response["jsonrpc"] = version
	} else {
		if version, ok := request["version"]; ok {
			response["version"] = version
		}
		response["result"] = nil
		response["error"] = nil
	}
	return
}

// InputFilter for JSONRPC Service
func (filter ServiceFilter) InputFilter(data []byte, context rpc.Context) []byte {
	if (len(data) > 0) && (data[0] == '[' || data[0] == '{') {
		var requests []map[string]interface{}
		if data[0] == '[' {
			if err := json.Unmarshal(data, &requests); err != nil {
				return data
			}
		} else {
			requests = make([]map[string]interface{}, 1)
			if err := json.Unmarshal(data, &requests[0]); err != nil {
				return data
			}
		}
		writer := io.NewWriter(true)
		n := len(requests)
		responses := make([]map[string]interface{}, n)
		for i, request := range requests {
			responses[i] = createResponse(request)
			if method, ok := request["method"].(string); ok && method != "" {
				writer.WriteByte(io.TagCall)
				writer.WriteString(method)
				if params, ok := request["params"].([]interface{}); ok && params != nil && len(params) > 0 {
					writer.Serialize(params)
				}
			}
		}
		writer.WriteByte(io.TagEnd)
		data = writer.Bytes()
		context.SetInterface("jsonrpc", responses)
	}
	return data
}

// OutputFilter for JSONRPC Service
func (filter ServiceFilter) OutputFilter(data []byte, context rpc.Context) []byte {
	responses, ok := context.GetInterface("jsonrpc").([]map[string]interface{})
	if ok && responses != nil {
		reader := io.NewReader(data, false)
		reader.JSONCompatible = true
		tag, _ := reader.ReadByte()
		for _, response := range responses {
			if tag == io.TagResult {
				reader.Reset()
				var result interface{}
				reader.Unserialize(&result)
				response["result"] = result
				tag, _ = reader.ReadByte()
			} else if tag == io.TagError {
				reader.Reset()
				err := make(map[string]interface{})
				err["code"] = -1
				message := reader.ReadString()
				err["message"] = message
				tag, _ = reader.ReadByte()
				response["error"] = err
			}
			if tag == io.TagEnd {
				break
			}
		}
		if len(responses) == 1 {
			data, _ = json.Marshal(responses[0])
		} else {
			data, _ = json.Marshal(responses)
		}
	}
	return data
}
