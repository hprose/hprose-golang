/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/encoder/marshaler.go                         |
|                                                          |
| LastModified: Feb 25, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoder

import (
	"reflect"
	"sync"
)

// Marshaler is the interface that groups the basic Write and Encode methods.
type Marshaler interface {
	Encode(enc *Encoder, v interface{}) error
	Write(enc *Encoder, v interface{}) error
}

var marshalers = sync.Map{}

// RegisterMarshaler ...
func RegisterMarshaler(t reflect.Type, marshaler Marshaler) {
	marshalers.Store(t, marshaler)
}

// GetMarshaler ...
func GetMarshaler(t reflect.Type) Marshaler {
	if m, ok := marshalers.Load(t); ok {
		return m.(Marshaler)
	}
	return nil
}
