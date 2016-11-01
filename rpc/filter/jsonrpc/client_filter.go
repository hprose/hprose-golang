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
 * rpc/filter/jsonrpc/client_filter.go                    *
 *                                                        *
 * hprose jsonrpc client filter for Go.                   *
 *                                                        *
 * LastModified: Nov 1, 2016                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package jsonrpc

import (
	"encoding/json"
	"sync/atomic"

	"github.com/hprose/hprose-golang/io"
	"github.com/hprose/hprose-golang/rpc"
)

// ClientFilter is a JSONRPC Client Filter
type ClientFilter struct {
	Version string
	id      int32
}

// NewClientFilter is a constructor for JSONRPCClientFilter
func NewClientFilter(version string) *ClientFilter {
	if version == "1.0" || version == "1.1" || version == "2.0" {
		return &ClientFilter{Version: version}
	}
	panic("version must be 1.0, 1.1 or 2.0 in string format.")
}

// InputFilter for JSONRPC Client
func (filter *ClientFilter) InputFilter(data []byte, context rpc.Context) []byte {
	if context.GetBool("jsonrpc") {
		var response map[string]interface{}
		if err := json.Unmarshal(data, &response); err != nil {
			return data
		}
		err := response["error"]
		writer := io.NewWriter(true)
		if err != nil {
			e := err.(map[string]interface{})
			writer.WriteByte(io.TagError)
			writer.WriteString(e["message"].(string))
		} else {
			writer.WriteByte(io.TagResult)
			writer.Serialize(response["result"])
		}
		writer.WriteByte(io.TagEnd)
		data = writer.Bytes()
	}
	return data
}

// OutputFilter for JSONRPC Client
func (filter *ClientFilter) OutputFilter(data []byte, context rpc.Context) []byte {
	if context.GetBool("jsonrpc") {
		request := make(map[string]interface{})
		if filter.Version == "1.1" {
			request["version"] = "1.1"
		} else if filter.Version == "2.0" {
			request["jsonrpc"] = "2.0"
		}
		reader := io.NewReader(data, false)
		reader.JSONCompatible = true
		tag, _ := reader.ReadByte()
		if tag == io.TagCall {
			request["method"] = reader.ReadString()
			tag, _ = reader.ReadByte()
			if tag == io.TagList {
				reader.Reset()
				count := reader.ReadCount()
				params := make([]interface{}, count)
				for i := 0; i < count; i++ {
					reader.Unserialize(&params[i])
				}
				request["params"] = params
			}
		}
		request["id"] = atomic.AddInt32(&filter.id, 1)
		data, _ = json.Marshal(request)
	}
	return data
}
