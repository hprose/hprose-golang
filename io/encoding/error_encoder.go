/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/error_encoder.go                             |
|                                                          |
| LastModified: Mar 21, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"reflect"
)

// ErrorEncoder is the implementation of ValueEncoder for error.
type ErrorEncoder struct{}

// Encode writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as reference
func (valenc ErrorEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	return valenc.Write(enc, v)
}

// Write writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as value
func (ErrorEncoder) Write(enc *Encoder, v interface{}) (err error) {
	switch v := v.(type) {
	case error:
		return WriteError(enc, v)
	case *error:
		return WriteError(enc, *v)
	}
	return &UnsupportedTypeError{Type: reflect.TypeOf(v)}
}

func init() {
	RegisterValueEncoder((*error)(nil), ErrorEncoder{})
}
