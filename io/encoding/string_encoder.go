/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/string_encoder.go                            |
|                                                          |
| LastModified: Mar 19, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"github.com/hprose/hprose-golang/v3/io"
	"github.com/modern-go/reflect2"
)

// StringEncoder is the implementation of ValueEncoder for string.
type StringEncoder struct{}

var stringEncoder StringEncoder

// Encode writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as reference
func (StringEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	return EncodeString(enc, *(*string)(reflect2.PtrOf(v)))
}

// Write writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as value
func (StringEncoder) Write(enc *Encoder, v interface{}) (err error) {
	return WriteString(enc, *(*string)(reflect2.PtrOf(v)))
}

// EncodeString to encoder
func EncodeString(enc *Encoder, s string) (err error) {
	length := utf16Length(s)
	switch length {
	case 0:
		err = enc.Writer.WriteByte(io.TagEmpty)
	case 1:
		if err = enc.Writer.WriteByte(io.TagUTF8Char); err == nil {
			_, err = enc.Writer.Write(reflect2.UnsafeCastString(s))
		}
	default:
		var ok bool
		if ok, err = enc.WriteStringReference(s); !ok && err == nil {
			enc.SetStringReference(s)
			err = writeString(enc.Writer, s, length)
		}
	}
	return
}

// WriteString to encoder
func WriteString(enc *Encoder, s string) (err error) {
	enc.SetStringReference(s)
	return writeString(enc.Writer, s, utf16Length(s))
}
