/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/ptr_decoder.go                                  |
|                                                          |
| LastModified: Jun 21, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

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
	switch tag {
	case TagNull:
		ep := reflect2.PtrOf(p)
		if !valdec.t.UnsafeIsNil(ep) {
			*(*unsafe.Pointer)(ep) = nil
		}
	default:
		e := valdec.et.New()
		valdec.elemDecoder.Decode(dec, e, tag)
		valdec.t.UnsafeSet(reflect2.PtrOf(p), reflect2.PtrOf(e))
	}
}

func (valdec ptrDecoder) Type() reflect.Type {
	return valdec.t.Type1()
}

func getPtrDecoder(t reflect.Type) ValueDecoder {
	et := t.Elem()
	if elemDecoder := getValueDecoder(et); elemDecoder == nil {
		return ptrDecoder{
			reflect2.Type2(t).(*reflect2.UnsafePtrType),
			reflect2.Type2(et),
			elemDecoder,
		}
	}
	return ptrDecoderFactories[et.Kind()](t)
}

func getArrayPtrDecoder(t reflect.Type) ValueDecoder {
	t2 := reflect2.Type2(t).(*reflect2.UnsafePtrType)
	return ptrDecoder{
		t2,
		t2.Elem(),
		getArrayDecoder(t.Elem()),
	}
}

func getMapPtrDecoder(t reflect.Type) ValueDecoder {
	t2 := reflect2.Type2(t).(*reflect2.UnsafePtrType)
	return ptrDecoder{
		t2,
		t2.Elem(),
		getMapDecoder(t.Elem()),
	}
}

func getPtrPtrDecoder(t reflect.Type) ValueDecoder {
	t2 := reflect2.Type2(t).(*reflect2.UnsafePtrType)
	return ptrDecoder{
		t2,
		t2.Elem(),
		getPtrDecoder(t.Elem()),
	}
}

func getSlicePtrDecoder(t reflect.Type) ValueDecoder {
	t2 := reflect2.Type2(t).(*reflect2.UnsafePtrType)
	return ptrDecoder{
		t2,
		t2.Elem(),
		getSliceDecoder(t.Elem()),
	}
}

func getStructPtrDecoder(t reflect.Type) ValueDecoder {
	t2 := reflect2.Type2(t).(*reflect2.UnsafePtrType)
	return ptrDecoder{
		t2,
		t2.Elem(),
		getStructDecoder(t.Elem()),
	}
}

var ptrDecoderFactories []func(t reflect.Type) ValueDecoder

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
		reflect.Interface:     func(t reflect.Type) ValueDecoder { return interfacePtrDecoder{t} },
		reflect.Map:           getMapPtrDecoder,
		reflect.Ptr:           getPtrPtrDecoder,
		reflect.Slice:         getSlicePtrDecoder,
		reflect.String:        func(t reflect.Type) ValueDecoder { return stringPtrDecoder{t} },
		reflect.Struct:        getStructPtrDecoder,
		reflect.UnsafePointer: invalidDecoder,
	}
}
