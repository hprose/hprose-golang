/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/encoder/string_marshaler.go                  |
|                                                          |
| LastModified: Feb 24, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoder

import (
	"github.com/hprose/hprose-golang/v3/io"
)

// StringMarshaler is the implementation of Marshaler for string.
type StringMarshaler struct{}

var stringMarshaler StringMarshaler

// Encode writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as reference
func (m StringMarshaler) Encode(enc *Encoder, v interface{}) (err error) {
	s := v.(string)
	length := utf16Length(s)
	switch length {
	case 0:
		err = enc.Writer.WriteByte(io.TagEmpty)
	case 1:
		if err = enc.Writer.WriteByte(io.TagUTF8Char); err == nil {
			_, err = enc.Writer.Write(io.StringToBytes(s))
		}
	default:
		var ok bool
		if ok, err = enc.WriteReference(v); !ok && err == nil {
			enc.SetReference(v)
			err = writeString(enc.Writer, s, length)
		}
	}
	return
}

// Write writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as value
func (m StringMarshaler) Write(enc *Encoder, v interface{}) (err error) {
	enc.SetReference(v)
	s := v.(string)
	return writeString(enc.Writer, s, utf16Length(s))
}
