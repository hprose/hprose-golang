/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/pool.go                                               |
|                                                          |
| LastModified: Feb 27, 2022                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import "sync"

var encoderPool = sync.Pool{New: func() interface{} { return new(Encoder) }}
var decoderPool = sync.Pool{New: func() interface{} { return new(Decoder) }}

func GetEncoder() *Encoder {
	return encoderPool.Get().(*Encoder)
}

func FreeEncoder(encoder *Encoder) {
	encoderPool.Put(encoder.Simple(false).ResetBuffer())
}

func GetDecoder() *Decoder {
	return decoderPool.Get().(*Decoder)
}

func FreeDecoder(decoder *Decoder) {
	decoderPool.Put(decoder.Simple(false).ResetBuffer())
}
