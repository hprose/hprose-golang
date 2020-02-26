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

func getMarshaler(v interface{}) Marshaler {
	switch reflect.TypeOf(v).Kind() {
	case reflect.String:
		return stringMarshaler
	case reflect.Ptr:
		return ptrMarshaler
	case reflect.Array:
		// return arrayMarshaler
	case reflect.Slice:
		return getSliceMarshaler(v)
	case reflect.Map:
		return getMapMarshaler(v)
	case reflect.Struct:
		return getStructMarshaler(v)
	}
	return nil
}

func getSliceMarshaler(v interface{}) Marshaler {
	return nil
}

func getMapMarshaler(v interface{}) Marshaler {
	return nil
}

func getStructMarshaler(v interface{}) Marshaler {
	return nil
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
