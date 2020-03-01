/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/encoder/interface_marshaler.go               |
|                                                          |
| LastModified: Mar 1, 2020                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoder

// InterfaceMarshaler is the implementation of Marshaler for interface.
type InterfaceMarshaler struct{}

var interfaceMarshaler InterfaceMarshaler

// Encode writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as reference
func (m InterfaceMarshaler) Encode(enc *Encoder, v interface{}) (err error) {
	return enc.Encode(v)
}

// Write writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as value
func (m InterfaceMarshaler) Write(enc *Encoder, v interface{}) (err error) {
	return enc.Write(v)
}
