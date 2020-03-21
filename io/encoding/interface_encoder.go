/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/interface_encoder.go                         |
|                                                          |
| LastModified: Mar 21, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

// interfaceEncoder is the implementation of ValueEncoder for interface.
type interfaceEncoder struct{}

var intfenc interfaceEncoder

func (interfaceEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	return enc.Encode(v)
}

func (interfaceEncoder) Write(enc *Encoder, v interface{}) (err error) {
	return enc.Write(v)
}
