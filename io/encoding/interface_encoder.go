/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/interface_encoder.go                         |
|                                                          |
| LastModified: Mar 15, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

// InterfaceEncoder is the implementation of ValueEncoder for interface.
type InterfaceEncoder struct{}

var interfaceEncoder InterfaceEncoder

// Encode writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as reference
func (InterfaceEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	return enc.Encode(v)
}

// Write writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as value
func (InterfaceEncoder) Write(enc *Encoder, v interface{}) (err error) {
	return enc.Write(v)
}
