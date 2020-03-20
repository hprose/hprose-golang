/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/time_encoder.go                              |
|                                                          |
| LastModified: Mar 21, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"time"

	"github.com/modern-go/reflect2"
)

// TimeEncoder is the implementation of ValueEncoder for time.Time/*time.Time.
type TimeEncoder struct{}

// Encode writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as reference
func (valenc TimeEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	return ReferenceEncode(valenc, enc, v)
}

// Write writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as value
func (TimeEncoder) Write(enc *Encoder, v interface{}) (err error) {
	SetReference(enc, v)
	return WriteTime(enc.Writer, *(*time.Time)(reflect2.PtrOf(v)))
}

func init() {
	RegisterEncoder((*time.Time)(nil), TimeEncoder{})
}
