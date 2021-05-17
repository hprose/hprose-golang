/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/interface_encoder.go                                  |
|                                                          |
| LastModified: Mar 21, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

// interfaceEncoder is the implementation of ValueEncoder for interface.
type interfaceEncoder struct{}

var intfenc interfaceEncoder

func (interfaceEncoder) Encode(enc *Encoder, v interface{}) {
	enc.encode(v)
}

func (interfaceEncoder) Write(enc *Encoder, v interface{}) {
	enc.write(v)
}
