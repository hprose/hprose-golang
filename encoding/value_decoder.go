/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/value_decoder.go                                |
|                                                          |
| LastModified: Jun 7, 2020                                |
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
	valdec = valueDecoderFactories[t.Kind()](t)
	RegisterValueDecoder(valdec)
	return
}

// RegisterValueDecoder valdec
func RegisterValueDecoder(valdec ValueDecoder) {
	decoderMap.Store(valdec.Type(), valdec)
}

// GetValueDecoder of type(p)
func GetValueDecoder(p interface{}) ValueDecoder {
	return getValueDecoder(reflect.TypeOf(p).Elem())
}

var valueDecoderFactories []func(t reflect.Type) ValueDecoder
var arrayDecoderFactories []func(t reflect.Type) ValueDecoder
var sliceDecoderFactories []func(t reflect.Type) ValueDecoder
var ptrDecoderFactories []func(t reflect.Type) ValueDecoder

func invalidDecoder(t reflect.Type) ValueDecoder {
	panic(UnsupportedTypeError{t})
}

func arrayDecoderFactory(t reflect.Type) ValueDecoder {
	return arrayDecoderFactories[t.Elem().Kind()](t)
}

func mapDecoderFactory(t reflect.Type) ValueDecoder {
	panic(UnsupportedTypeError{t})
}

func ptrDecoderFactory(t reflect.Type) ValueDecoder {
	return ptrDecoderFactories[t.Elem().Kind()](t)
}

func sliceDecoderFactory(t reflect.Type) ValueDecoder {
	return sliceDecoderFactories[t.Elem().Kind()](t)
}

func structDecoderFactory(t reflect.Type) ValueDecoder {
	panic(UnsupportedTypeError{t})
}

func arrayPtrDecoderFactory(t reflect.Type) ValueDecoder {
	panic(UnsupportedTypeError{t})
}

func mapPtrDecoderFactory(t reflect.Type) ValueDecoder {
	panic(UnsupportedTypeError{t})
}

func ptrPtrDecoderFactory(t reflect.Type) ValueDecoder {
	panic(UnsupportedTypeError{t})
}

func slicePtrDecoderFactory(t reflect.Type) ValueDecoder {
	panic(UnsupportedTypeError{t})
}

func structPtrDecoderFactory(t reflect.Type) ValueDecoder {
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
		reflect.Array:         arrayDecoderFactory,
		reflect.Chan:          invalidDecoder,
		reflect.Func:          invalidDecoder,
		reflect.Interface:     invalidDecoder,
		reflect.Map:           mapDecoderFactory,
		reflect.Ptr:           ptrDecoderFactory,
		reflect.Slice:         sliceDecoderFactory,
		reflect.String:        func(t reflect.Type) ValueDecoder { return stringDecoder{t} },
		reflect.Struct:        structDecoderFactory,
		reflect.UnsafePointer: invalidDecoder,
	}

	arrayDecoderFactories = []func(t reflect.Type) ValueDecoder{
		reflect.Invalid:       invalidDecoder,
		reflect.Bool:          boolArrayDecoder,
		reflect.Int:           intArrayDecoder,
		reflect.Int8:          int8ArrayDecoder,
		reflect.Int16:         int16ArrayDecoder,
		reflect.Int32:         int32ArrayDecoder,
		reflect.Int64:         int64ArrayDecoder,
		reflect.Uint:          uintArrayDecoder,
		reflect.Uint8:         uint8ArrayDecoder,
		reflect.Uint16:        uint16ArrayDecoder,
		reflect.Uint32:        uint32ArrayDecoder,
		reflect.Uint64:        uint64ArrayDecoder,
		reflect.Uintptr:       uintptrArrayDecoder,
		reflect.Float32:       float32ArrayDecoder,
		reflect.Float64:       float64ArrayDecoder,
		reflect.Complex64:     complex64ArrayDecoder,
		reflect.Complex128:    complex128ArrayDecoder,
		reflect.Array:         otherArrayDecoder,
		reflect.Chan:          invalidDecoder,
		reflect.Func:          invalidDecoder,
		reflect.Interface:     interfaceArrayDecoder,
		reflect.Map:           otherArrayDecoder,
		reflect.Ptr:           otherArrayDecoder,
		reflect.Slice:         otherArrayDecoder,
		reflect.String:        stringArrayDecoder,
		reflect.Struct:        otherArrayDecoder,
		reflect.UnsafePointer: invalidDecoder,
	}

	sliceDecoderFactories = []func(t reflect.Type) ValueDecoder{
		reflect.Invalid:       invalidDecoder,
		reflect.Bool:          boolSliceDecoder,
		reflect.Int:           intSliceDecoder,
		reflect.Int8:          int8SliceDecoder,
		reflect.Int16:         int16SliceDecoder,
		reflect.Int32:         int32SliceDecoder,
		reflect.Int64:         int64SliceDecoder,
		reflect.Uint:          uintSliceDecoder,
		reflect.Uint8:         func(t reflect.Type) ValueDecoder { return bytesDecoder{t} },
		reflect.Uint16:        uint16SliceDecoder,
		reflect.Uint32:        uint32SliceDecoder,
		reflect.Uint64:        uint64SliceDecoder,
		reflect.Uintptr:       uintptrSliceDecoder,
		reflect.Float32:       float32SliceDecoder,
		reflect.Float64:       float64SliceDecoder,
		reflect.Complex64:     complex64SliceDecoder,
		reflect.Complex128:    complex128SliceDecoder,
		reflect.Array:         otherSliceDecoder,
		reflect.Chan:          invalidDecoder,
		reflect.Func:          invalidDecoder,
		reflect.Interface:     interfaceSliceDecoder,
		reflect.Map:           otherSliceDecoder,
		reflect.Ptr:           otherSliceDecoder,
		reflect.Slice:         otherSliceDecoder,
		reflect.String:        stringSliceDecoder,
		reflect.Struct:        otherSliceDecoder,
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
		reflect.Array:         arrayPtrDecoderFactory,
		reflect.Chan:          invalidDecoder,
		reflect.Func:          invalidDecoder,
		reflect.Interface:     invalidDecoder,
		reflect.Map:           mapPtrDecoderFactory,
		reflect.Ptr:           ptrPtrDecoderFactory,
		reflect.Slice:         slicePtrDecoderFactory,
		reflect.String:        func(t reflect.Type) ValueDecoder { return stringPtrDecoder{t} },
		reflect.Struct:        structPtrDecoderFactory,
		reflect.UnsafePointer: invalidDecoder,
	}
}
