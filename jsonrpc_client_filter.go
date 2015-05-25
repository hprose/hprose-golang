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
 * hprose/jsonrpc_client_filter.go                        *
 *                                                        *
 * jsonrpc client filter for Go.                          *
 *                                                        *
 * LastModified: May 22, 2015                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"bytes"
	"encoding/json"
)

var id = 1

// JSONRPCClientFilter is a JSONRPC Client Filter
type JSONRPCClientFilter struct {
	Version string
}

// NewJSONRPCClientFilter is a constructor for JSONRPCClientFilter
func NewJSONRPCClientFilter(version string) JSONRPCClientFilter {
	if version == "1.0" || version == "1.1" || version == "2.0" {
		return JSONRPCClientFilter{Version: version}
	}
	panic("version must be 1.0, 1.1 or 2.0 in string format.")
}

// InputFilter for JSONRPC Client
func (filter JSONRPCClientFilter) InputFilter(data []byte, context Context) []byte {
	var response map[string]interface{}
	if err := json.Unmarshal(data, &response); err != nil {
		return data
	}
	result := response["result"]
	err := response["error"]
	buf := new(bytes.Buffer)
	writer := NewWriter(buf, true)
	if err != nil {
		e := err.(map[string]interface{})
		buf.WriteByte(TagError)
		writer.WriteString(e["message"].(string))
	} else {
		buf.WriteByte(TagResult)
		writer.Serialize(result)
	}
	buf.WriteByte(TagEnd)
	data = buf.Bytes()
	return data
}

// OutputFilter for JSONRPC Client
func (filter JSONRPCClientFilter) OutputFilter(data []byte, context Context) []byte {
	request := make(map[string]interface{})
	if filter.Version == "1.1" {
		request["version"] = "1.1"
	} else if filter.Version == "2.0" {
		request["jsonrpc"] = "2.0"
	}
	istream := NewBytesReader(data)
	reader := NewReader(istream, false)
	reader.JSONCompatible = true
	tag, _ := istream.ReadByte()
	if tag == TagCall {
		request["method"], _ = reader.ReadString()
		tag, _ = istream.ReadByte()
		if tag == TagList {
			reader.Reset()
			if count, err := reader.ReadInteger(TagOpenbrace); err == nil {
				params := make([]interface{}, count)
				for i := 0; i < count; i++ {
					reader.Unserialize(&params[i])
				}
				request["params"] = params
			}
		}
	}
	request["id"] = id
	id++
	data, _ = json.Marshal(request)
	return data
}
