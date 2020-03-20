/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/error_encoder.go                             |
|                                                          |
| LastModified: Mar 20, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"reflect"

	"github.com/hprose/hprose-golang/v3/io"
)

// WriteError to encoder
func WriteError(enc *Encoder, e error) (err error) {
	if err = enc.WriteByte(io.TagError); err == nil {
		enc.AddReferenceCount(1)
		s := e.Error()
		err = writeString(enc.Writer, s, utf16Length(s))
	}
	return
}

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
	RegisterEncoder((*error)(nil), ErrorEncoder{})
}
