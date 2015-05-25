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
 * hprose/jsonrpc_service_filter.go                       *
 *                                                        *
 * jsonrpc service filter for Go.                         *
 *                                                        *
 * LastModified: May 25, 2015                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"bytes"
	"encoding/json"
)

// JSONRPCServiceFilter is a JSONRPC Service Filter
type JSONRPCServiceFilter struct{}

// InputFilter for JSONRPC Service
func (filter JSONRPCServiceFilter) InputFilter(data []byte, context Context) []byte {
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
		buf := new(bytes.Buffer)
		writer := NewWriter(buf, true)
		n := len(requests)
		jsonrpc := make([]map[string]interface{}, n)
		for i, request := range requests {
			j := make(map[string]interface{})
			if id, ok := request["id"]; ok {
				j["id"] = id
			} else {
				j["id"] = nil
			}
			if version, ok := request["version"]; ok {
				j["version"] = version
			} else if jsonrpc, ok := request["jsonrpc"]; ok {
				j["version"] = jsonrpc
			} else {
				j["version"] = "1.0"
			}
			jsonrpc[i] = j
			if method, ok := request["method"].(string); ok && method != "" {
				buf.WriteByte(TagCall)
				writer.WriteString(method)
				if params, ok := request["params"].([]interface{}); ok && params != nil && len(params) > 0 {
					writer.Serialize(params)
				}
			}
		}
		buf.WriteByte(TagEnd)
		data = buf.Bytes()
		context.SetInterface("jsonrpc", jsonrpc)
	}
	return data
}

func initResponse(j map[string]interface{}, err interface{}) map[string]interface{} {
	response := make(map[string]interface{})
	if version, ok := j["version"].(string); ok && version != "2.0" {
		if version == "1.1" {
			response["version"] = "1.1"
		}
		response["result"] = nil
		response["error"] = err
	} else {
		response["jsonrpc"] = "2.0"
	}
	response["id"] = j["id"]
	return response
}

func jsonrpcError(jsonrpc []map[string]interface{}, err error) []byte {
	n := len(jsonrpc)
	responses := make([]map[string]interface{}, n)
	for i, j := range jsonrpc {
		responses[i] = initResponse(j, err.Error())
	}
	if n == 1 {
		data, _ := json.Marshal(responses[0])
		return data
	}
	data, _ := json.Marshal(responses)
	return data
}

// OutputFilter for JSONRPC Service
func (filter JSONRPCServiceFilter) OutputFilter(data []byte, context Context) []byte {
	if jsonrpc, ok := context.GetInterface("jsonrpc"); ok {
		jsonrpc := jsonrpc.([]map[string]interface{})
		n := len(jsonrpc)
		responses := make([]map[string]interface{}, n)
		istream := NewBytesReader(data)
		reader := NewReader(istream, false)
		reader.JSONCompatible = true
		tag, err := istream.ReadByte()
		if err != nil {
			return jsonrpcError(jsonrpc, err)
		}
		i := 0
		for {
			response := initResponse(jsonrpc[i], nil)
			if tag == TagResult {
				reader.Reset()
				var result interface{}
				reader.Unserialize(&result)
				response["result"] = result
				tag, _ = istream.ReadByte()
			} else if tag == TagError {
				reader.Reset()
				e := make(map[string]interface{})
				e["code"] = -1
				message, _ := reader.ReadString()
				e["message"] = message
				tag, _ = istream.ReadByte()
			}
			responses[i] = response
			i++
			if tag == TagEnd {
				break
			}
		}
		if n == 1 {
			data, _ := json.Marshal(responses[0])
			return data
		}
		data, _ := json.Marshal(responses)
		return data
	}
	return data
}
