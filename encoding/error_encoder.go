/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/error_encoder.go                                |
|                                                          |
| LastModified: Apr 12, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

// errorEncoder is the implementation of ValueEncoder for error/*error.
type errorEncoder struct{}

func (valenc errorEncoder) Encode(enc *Encoder, v interface{}) {
	valenc.Write(enc, v)
}

func (errorEncoder) Write(enc *Encoder, v interface{}) {
	switch v := v.(type) {
	case error:
		enc.WriteError(v)
	case *error:
		enc.WriteError(*v)
	}
}

func init() {
	RegisterValueEncoder((*error)(nil), errorEncoder{})
}
