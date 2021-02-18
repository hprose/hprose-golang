/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/string_encoder.go                               |
|                                                          |
| LastModified: Feb 18, 2021                               |
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

func (stringEncoder) Encode(enc *Encoder, v interface{}) {
	enc.EncodeString(*(*string)(reflect2.PtrOf(v)))
}

func (stringEncoder) Write(enc *Encoder, v interface{}) {
	enc.WriteString(*(*string)(reflect2.PtrOf(v)))
}

// EncodeString to encoder.
func (enc *Encoder) EncodeString(s string) {
	length := utf16Length(s)
	switch length {
	case 0:
		enc.buf = append(enc.buf, TagEmpty)
	case 1:
		enc.buf = append(enc.buf, TagUTF8Char)
		enc.buf = append(enc.buf, s...)
	default:
		if ok := enc.WriteStringReference(s); !ok {
			enc.SetStringReference(s)
			enc.buf = appendString(enc.buf, s, length)
		}
	}
}

// WriteString to encoder.
func (enc *Encoder) WriteString(s string) {
	enc.SetStringReference(s)
	enc.buf = appendString(enc.buf, s, utf16Length(s))
}
