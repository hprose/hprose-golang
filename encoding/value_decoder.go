/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/value_decoder.go                                |
|                                                          |
| LastModified: Apr 18, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"reflect"
	"sync"
)

// ValueDecoder is the interface that groups the basic Decode methods.
type ValueDecoder interface {
	Decode(dec *Decoder, p interface{}, tag byte)
}

var structDecoderMap = sync.Map{}
var otherDecoderMap = sync.Map{}

func registerValueDecoder(t reflect.Type, valdec ValueDecoder) {
	if t.Kind() == reflect.Struct {
		structDecoderMap.Store(t, valdec)
	} else {
		otherDecoderMap.Store(t, valdec)
	}
}

func getStructDecoder(t reflect.Type) ValueDecoder {
	if valdec, ok := structDecoderMap.Load(t); ok {
		return valdec.(ValueDecoder)
	}
	return nil
	// name := t.Name()
	// if name == "" {
	// 	return newAnonymousStructDecoder(t)
	// }
	// return newStructDecoder(t, name, []string{"json"})
}

func getOtherDecoder(t reflect.Type) ValueDecoder {
	if valdec, ok := otherDecoderMap.Load(t); ok {
		return valdec.(ValueDecoder)
	}
	return nil
}

// RegisterValueDecoder of type(v)
func RegisterValueDecoder(v interface{}, valdec ValueDecoder) {
	registerValueDecoder(checkType(v), valdec)
}

// GetValueDecoder of type(v)
func GetValueDecoder(v interface{}) ValueDecoder {
	t := checkType(v)
	if t.Kind() == reflect.Struct {
		return getStructDecoder(t)
	}
	return getOtherDecoder(t)
}
