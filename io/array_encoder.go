/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/array_encoder.go                                      |
|                                                          |
| LastModified: Feb 18, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"reflect"
)

// arrayEncoder is the implementation of ValueEncoder for *array.
type arrayEncoder struct{}

var arrayenc arrayEncoder

func (valenc arrayEncoder) Encode(enc *Encoder, v interface{}) {
	if ok := enc.WriteReference(v); !ok {
		valenc.Write(enc, v)
	}
}

func (arrayEncoder) Write(enc *Encoder, v interface{}) {
	enc.setReference(v)
	enc.writeArray(reflect.ValueOf(v).Elem().Interface())
}

// WriteArray to encoder.
func (enc *Encoder) WriteArray(v interface{}) {
	enc.AddReferenceCount(1)
	enc.writeArray(v)
}

func (enc *Encoder) writeArray(array interface{}) {
	enc.writeSlice(toSlice(array))
}
