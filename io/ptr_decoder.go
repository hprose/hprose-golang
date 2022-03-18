/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/ptr_decoder.go                                        |
|                                                          |
| LastModified: Mar 18, 2022                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

// ptrDecoder is the implementation of ValueDecoder for *T.
type ptrDecoder struct {
	t           *reflect2.UnsafePtrType
	et          reflect2.Type
	elemDecoder ValueDecoder
}

func (valdec ptrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	ptr := (*unsafe.Pointer)(reflect2.PtrOf(p))
	switch tag {
	case TagNull:
		if *ptr != nil {
			*ptr = nil
		}
	default:
		if *ptr == nil {
			*ptr = valdec.et.UnsafeNew()
		}
		valdec.elemDecoder.Decode(dec, valdec.et.PackEFace(*ptr), tag)
	}
}

func getPtrDecoder(t reflect.Type) ValueDecoder {
	et := t.Elem()
	elemDecoder := getRegisteredValueDecoder(et)
	if elemDecoder == nil {
		elemDecoder = getNamedStructDecoder(et)
	}
	if elemDecoder != nil {
		return ptrDecoder{
			reflect2.Type2(t).(*reflect2.UnsafePtrType),
			reflect2.Type2(et),
			elemDecoder,
		}
	}
	return ptrDecoderFactories[et.Kind()](t)
}

func _getPtrDecoder(t reflect.Type, getElemDecoder func(t reflect.Type) ValueDecoder) ValueDecoder {
	et := t.Elem()
	elemDecoder := getElemDecoder(et)
	registerValueDecoder(et, elemDecoder)
	return ptrDecoder{
		reflect2.Type2(t).(*reflect2.UnsafePtrType),
		reflect2.Type2(et),
		elemDecoder,
	}
}

func getArrayPtrDecoder(t reflect.Type) ValueDecoder {
	return _getPtrDecoder(t, getArrayDecoder)
}

func getMapPtrDecoder(t reflect.Type) ValueDecoder {
	return _getPtrDecoder(t, getMapDecoder)
}

func getPtrPtrDecoder(t reflect.Type) ValueDecoder {
	return _getPtrDecoder(t, getPtrDecoder)
}

func getSlicePtrDecoder(t reflect.Type) ValueDecoder {
	return _getPtrDecoder(t, getSliceDecoder)
}

func getStructPtrDecoder(t reflect.Type) ValueDecoder {
	return _getPtrDecoder(t, getStructDecoder)
}

var ptrDecoderFactories []func(t reflect.Type) ValueDecoder

//nolint
func init() {
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
		reflect.Interface:     func(t reflect.Type) ValueDecoder { return interfacePtrDecoder{} },
		reflect.Map:           getMapPtrDecoder,
		reflect.Ptr:           getPtrPtrDecoder,
		reflect.Slice:         getSlicePtrDecoder,
		reflect.String:        func(t reflect.Type) ValueDecoder { return stringPtrDecoder{t} },
		reflect.Struct:        getStructPtrDecoder,
		reflect.UnsafePointer: invalidDecoder,
	}
}
