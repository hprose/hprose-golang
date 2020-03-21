/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/error_encoder.go                                |
|                                                          |
| LastModified: Mar 21, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

// ErrorEncoder is the implementation of ValueEncoder for error/*error.
type ErrorEncoder struct{}

// Encode writes the hprose encoding of v to stream
func (valenc ErrorEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	return valenc.Write(enc, v)
}

// Write writes the hprose encoding of v to stream
func (ErrorEncoder) Write(enc *Encoder, v interface{}) (err error) {
	switch v := v.(type) {
	case error:
		err = WriteError(enc, v)
	case *error:
		err = WriteError(enc, *v)
	}
	return
}

func init() {
	RegisterValueEncoder((*error)(nil), ErrorEncoder{})
}
