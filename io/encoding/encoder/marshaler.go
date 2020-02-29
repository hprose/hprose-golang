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

var marshalerMap = map[reflect.Type]Marshaler{}
var marshalerLocker = sync.RWMutex{}

// RegisterMarshaler ...
func RegisterMarshaler(t reflect.Type, marshaler Marshaler) {
	marshalerLocker.Lock()
	defer marshalerLocker.Unlock()
	marshalerMap[t] = marshaler
}

// GetMarshaler ...
func GetMarshaler(t reflect.Type) Marshaler {
	marshalerLocker.RLock()
	defer marshalerLocker.RUnlock()
	return marshalerMap[t]
}
