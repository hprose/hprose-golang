/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/string_encoder.go                               |
|                                                          |
| LastModified: Mar 19, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"github.com/modern-go/reflect2"
)

// stringEncoder is the implementation of ValueEncoder for string.
type stringEncoder struct{}

var strenc stringEncoder

func (stringEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	return EncodeString(enc, *(*string)(reflect2.PtrOf(v)))
}

func (stringEncoder) Write(enc *Encoder, v interface{}) (err error) {
	return WriteString(enc, *(*string)(reflect2.PtrOf(v)))
}

// EncodeString to encoder
func EncodeString(enc *Encoder, s string) (err error) {
	length := utf16Length(s)
	switch length {
	case 0:
		err = enc.Writer.WriteByte(TagEmpty)
	case 1:
		if err = enc.Writer.WriteByte(TagUTF8Char); err == nil {
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
