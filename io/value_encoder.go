/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/value_encoder.go                                      |
|                                                          |
| LastModified: Mar 5, 2022                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"reflect"
	"sync"
)

// ValueEncoder is the interface that groups the basic Write and Encode methods.
type ValueEncoder interface {
	Encode(enc *Encoder, v interface{})
	Write(enc *Encoder, v interface{})
}

var structEncoderMap sync.Map
var otherEncoderMap sync.Map

func registerValueEncoder(t reflect.Type, valenc ValueEncoder) {
	if t.Kind() == reflect.Struct {
		structEncoderMap.Store(t, valenc)
	} else {
		otherEncoderMap.Store(t, valenc)
	}
}

func getStructEncoder(t reflect.Type) ValueEncoder {
	if valenc, ok := structEncoderMap.Load(t); ok {
		return valenc.(ValueEncoder)
	}
	if name := t.Name(); name != "" {
		return newNamedStructEncoder(t, name)
	}
	return newAnonymousStructEncoder(t)
}

func getOtherEncoder(t reflect.Type) ValueEncoder {
	if valenc, ok := otherEncoderMap.Load(t); ok {
		return valenc.(ValueEncoder)
	}
	return nil
}

func checkType(v interface{}) reflect.Type {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

// RegisterValueEncoder of type(v).
func RegisterValueEncoder(v interface{}, valenc ValueEncoder) {
	registerValueEncoder(checkType(v), valenc)
}

// GetValueEncoder of type(v).
func GetValueEncoder(v interface{}) ValueEncoder {
	t := checkType(v)
	if t.Kind() == reflect.Struct {
		return getStructEncoder(t)
	}
	return getOtherEncoder(t)
}
