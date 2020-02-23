/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/encoder/marshaler.go                         |
|                                                          |
| LastModified: Feb 23, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoder

import (
	"reflect"
	"unsafe"
)

// Marshaler is the interface that groups the basic Write and Encode methods.
type Marshaler interface {
	Encode(enc *Encoder, v interface{}) error
	Write(enc *Encoder, v interface{}) error
}

// ValueMarshaler is a marshal function for value struct
type ValueMarshaler func(enc *Encoder, v interface{}) error

type typeAddr uintptr

func getTypeAddr(v interface{}) typeAddr {
	return *(*typeAddr)(unsafe.Pointer(&v))
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

func getValueMarshaler(t reflect.Type) ValueMarshaler {
	return nil
}
