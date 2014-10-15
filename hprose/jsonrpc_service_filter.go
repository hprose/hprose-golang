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
 * LastModified: Oct 15, 2014                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"bytes"
	"encoding/json"
)

type JSONRPCServiceFilter struct{}

func (filter JSONRPCServiceFilter) InputFilter(data []byte, context Context) []byte {
	if len(data) > 0 && data[0] == '{' {
		context.SetString("format", "jsonrpc")
		var request map[string]interface{}
		if err := json.Unmarshal(data, &request); err != nil {
			return data
		}
		if id, ok := request["id"]; ok {
			context.SetInterface("id", id)
		} else {
			context.SetInterface("id", nil)
		}
		if version, ok := request["version"].(string); ok {
			context.SetString("version", version)
		} else if jsonrpc, ok := request["jsonrpc"].(string); ok {
			context.SetString("version", jsonrpc)
		} else {
			context.SetString("version", "1.0")
		}
		buf := new(bytes.Buffer)
		writer := NewWriter(buf, true)
		if method, ok := request["method"].(string); ok && method != "" {
			if err := buf.WriteByte(TagCall); err != nil {
				return data
			}
			if err := writer.WriteString(method); err != nil {
				return data
			}
			if params, ok := request["params"].([]interface{}); ok && params != nil && len(params) > 0 {
				if err := writer.Serialize(params); err != nil {
					return data
				}
			}
		}
		buf.WriteByte(TagEnd)
		data = buf.Bytes()
	}
	return data
}

func (filter JSONRPCServiceFilter) OutputFilter(data []byte, context Context) []byte {
	if format, ok := context.GetString("format"); ok && format == "jsonrpc" {
		response := make(map[string]interface{})
		if version, ok := context.GetString("version"); ok && version != "2.0" {
			if version == "1.1" {
				response["version"] = "1.1"
			}
			response["result"] = nil
			response["error"] = nil
		} else {
			response["jsonrpc"] = "2.0"
		}
		response["id"], _ = context.GetInterface("id")
		if len(data) == 0 {
			data, _ = json.Marshal(response)
			return data
		}
		istream := NewBytesReader(data)
		reader := NewReader(istream, false)
		for tag, err := istream.ReadByte(); err == nil && tag != TagEnd; tag, err = istream.ReadByte() {
			switch tag {
			case TagResult:
				reader.Reset()
				var result interface{}
				reader.Unserialize(&result)
				if err != nil {
					e := make(map[string]interface{})
					e["code"] = -1
					e["message"] = err.Error()
					response["error"] = e
				} else {
					response["result"] = result
				}
			case TagError:
				reader.Reset()
				e := make(map[string]interface{})
				e["code"] = -1
				if message, err := reader.ReadString(); err == nil {
					e["message"] = message
				} else {
					e["message"] = err.Error()
				}
			default:
				data, _ = json.Marshal(response)
				return data
			}
		}
		data, _ = json.Marshal(response)
		return data
	}
	return data
}
