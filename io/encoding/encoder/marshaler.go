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
		return getArrayMarshaler(v)
	case reflect.Map:
		return getMapMarshaler(v)
	case reflect.Struct:
		return getStructMarshaler(v)
	}
	return nil
}

func getArrayMarshaler(v interface{}) Marshaler {
	return nil
}

func getMapMarshaler(v interface{}) Marshaler {
	return nil
}

func getStructMarshaler(v interface{}) Marshaler {
	return nil
}

// ValueMarshaler is a marshal function for value struct
type ValueMarshaler func(enc *Encoder, v interface{}) error

var valueMarshalerMap = map[reflect.Type]ValueMarshaler{}
var valueMarshalerLocker = sync.RWMutex{}

// RegisterValueMarshaler ...
func RegisterValueMarshaler(t reflect.Type, marshaler ValueMarshaler) {
	valueMarshalerLocker.Lock()
	defer valueMarshalerLocker.Unlock()
	valueMarshalerMap[t] = marshaler
}

// GetValueMarshaler ...
func GetValueMarshaler(t reflect.Type) ValueMarshaler {
	valueMarshalerLocker.RLock()
	defer valueMarshalerLocker.RUnlock()
	return valueMarshalerMap[t]
}
