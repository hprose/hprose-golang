/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/value_decoder.go                                |
|                                                          |
| LastModified: Jun 19, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"reflect"
	"sync"
)

var decoderMap sync.Map

// ValueDecoder is the interface that groups the basic Decode methods.
type ValueDecoder interface {
	Decode(dec *Decoder, p interface{}, tag byte)
	Type() reflect.Type
}

func getValueDecoder(t reflect.Type) (valdec ValueDecoder) {
	if valdec, ok := decoderMap.Load(t); ok {
		return valdec.(ValueDecoder)
	}
	return nil
}

// RegisterValueDecoder valdec
func RegisterValueDecoder(valdec ValueDecoder) {
	decoderMap.Store(valdec.Type(), valdec)
}

// GetValueDecoder of Type t
func GetValueDecoder(t reflect.Type) (valdec ValueDecoder) {
	valdec = getValueDecoder(t)
	if valdec == nil {
		valdec = valueDecoderFactories[t.Kind()](t)
		RegisterValueDecoder(valdec)
	}
	return
}

var valueDecoderFactories []func(t reflect.Type) ValueDecoder
var ptrDecoderFactories []func(t reflect.Type) ValueDecoder

func invalidDecoder(t reflect.Type) ValueDecoder {
	panic(UnsupportedTypeError{t})
}

func getPtrDecoder(t reflect.Type) ValueDecoder {
	return ptrDecoderFactories[t.Elem().Kind()](t)
}

func getStructDecoder(t reflect.Type) ValueDecoder {
	panic(UnsupportedTypeError{t})
}

func getArrayPtrDecoder(t reflect.Type) ValueDecoder {
	panic(UnsupportedTypeError{t})
}

func getMapPtrDecoder(t reflect.Type) ValueDecoder {
	panic(UnsupportedTypeError{t})
}

func getPtrPtrDecoder(t reflect.Type) ValueDecoder {
	panic(UnsupportedTypeError{t})
}

func getSlicePtrDecoder(t reflect.Type) ValueDecoder {
	panic(UnsupportedTypeError{t})
}

func getStructPtrDecoder(t reflect.Type) ValueDecoder {
	panic(UnsupportedTypeError{t})
}

func init() {
	valueDecoderFactories = []func(t reflect.Type) ValueDecoder{
		reflect.Invalid:       invalidDecoder,
		reflect.Bool:          func(t reflect.Type) ValueDecoder { return boolDecoder{t} },
		reflect.Int:           func(t reflect.Type) ValueDecoder { return intDecoder{t} },
		reflect.Int8:          func(t reflect.Type) ValueDecoder { return int8Decoder{t} },
		reflect.Int16:         func(t reflect.Type) ValueDecoder { return int16Decoder{t} },
		reflect.Int32:         func(t reflect.Type) ValueDecoder { return int32Decoder{t} },
		reflect.Int64:         func(t reflect.Type) ValueDecoder { return int64Decoder{t} },
		reflect.Uint:          func(t reflect.Type) ValueDecoder { return uintDecoder{t} },
		reflect.Uint8:         func(t reflect.Type) ValueDecoder { return uint8Decoder{t} },
		reflect.Uint16:        func(t reflect.Type) ValueDecoder { return uint16Decoder{t} },
		reflect.Uint32:        func(t reflect.Type) ValueDecoder { return uint32Decoder{t} },
		reflect.Uint64:        func(t reflect.Type) ValueDecoder { return uint64Decoder{t} },
		reflect.Uintptr:       func(t reflect.Type) ValueDecoder { return uintptrDecoder{t} },
		reflect.Float32:       func(t reflect.Type) ValueDecoder { return float32Decoder{t} },
		reflect.Float64:       func(t reflect.Type) ValueDecoder { return float64Decoder{t} },
		reflect.Complex64:     func(t reflect.Type) ValueDecoder { return complex64Decoder{t} },
		reflect.Complex128:    func(t reflect.Type) ValueDecoder { return complex128Decoder{t} },
		reflect.Array:         getArrayDecoder,
		reflect.Chan:          invalidDecoder,
		reflect.Func:          invalidDecoder,
		reflect.Interface:     func(t reflect.Type) ValueDecoder { return interfaceDecoder{t} },
		reflect.Map:           getMapDecoder,
		reflect.Ptr:           getPtrDecoder,
		reflect.Slice:         getSliceDecoder,
		reflect.String:        func(t reflect.Type) ValueDecoder { return stringDecoder{t} },
		reflect.Struct:        getStructDecoder,
		reflect.UnsafePointer: invalidDecoder,
	}

	ptrDecoderFactories = []func(t reflect.Type) ValueDecoder{
		reflect.Invalid:       invalidDecoder,
		reflect.Bool:          func(t reflect.Type) ValueDecoder { return boolPtrDecoder{t} },
		reflect.Int:           func(t reflect.Type) ValueDecoder { return intPtrDecoder{t} },
		reflect.Int8:          func(t reflect.Type) ValueDecoder { return int8PtrDecoder{t} },
		reflect.Int16:         func(t reflect.Type) ValueDecoder { return int16PtrDecoder{t} },
		reflect.Int32:         func(t reflect.Type) ValueDecoder { return int32PtrDecoder{t} },
		reflect.Int64:         func(t reflect.Type) ValueDecoder { return int64PtrDecoder{t} },
		reflect.Uint:          func(t reflect.Type) ValueDecoder { return uintPtrDecoder{t} },
		reflect.Uint8:         func(t reflect.Type) ValueDecoder { return uint8PtrDecoder{t} },
		reflect.Uint16:        func(t reflect.Type) ValueDecoder { return uint16PtrDecoder{t} },
		reflect.Uint32:        func(t reflect.Type) ValueDecoder { return uint32PtrDecoder{t} },
		reflect.Uint64:        func(t reflect.Type) ValueDecoder { return uint64PtrDecoder{t} },
		reflect.Uintptr:       func(t reflect.Type) ValueDecoder { return uintptrPtrDecoder{t} },
		reflect.Float32:       func(t reflect.Type) ValueDecoder { return float32PtrDecoder{t} },
		reflect.Float64:       func(t reflect.Type) ValueDecoder { return float64PtrDecoder{t} },
		reflect.Complex64:     func(t reflect.Type) ValueDecoder { return complex64PtrDecoder{t} },
		reflect.Complex128:    func(t reflect.Type) ValueDecoder { return complex128PtrDecoder{t} },
		reflect.Array:         getArrayPtrDecoder,
		reflect.Chan:          invalidDecoder,
		reflect.Func:          invalidDecoder,
		reflect.Interface:     func(t reflect.Type) ValueDecoder { return interfacePtrDecoder{t} },
		reflect.Map:           getMapPtrDecoder,
		reflect.Ptr:           getPtrPtrDecoder,
		reflect.Slice:         getSlicePtrDecoder,
		reflect.String:        func(t reflect.Type) ValueDecoder { return stringPtrDecoder{t} },
		reflect.Struct:        getStructPtrDecoder,
		reflect.UnsafePointer: invalidDecoder,
	}
}
