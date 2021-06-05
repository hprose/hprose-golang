/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/slice_decoder.go                                      |
|                                                          |
| LastModified: Jun 5, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

// sliceDecoder is the implementation of ValueDecoder for []T.
type sliceDecoder struct {
	t          *reflect2.UnsafeSliceType
	et         reflect.Type
	empty      unsafe.Pointer
	decodeElem DecodeHandler
}

func (valdec sliceDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch tag {
	case TagNull:
		valdec.t.UnsafeSetNil(reflect2.PtrOf(p))
	case TagEmpty:
		setSliceHeader(reflect2.PtrOf(p), valdec.empty, 0)
	case TagList:
		count := dec.ReadInt()
		slice := reflect2.PtrOf(p)
		valdec.t.UnsafeGrow(slice, count)
		dec.AddReference(p)
		for i := 0; i < count; i++ {
			valdec.decodeElem(dec, valdec.et, valdec.t.UnsafeGetIndex(slice, i))
		}
		dec.Skip()
	default:
		dec.defaultDecode(valdec.t.Type1(), p, tag)
	}
}

// makeSliceDecoder returns a ValueDecoder for []T.
func makeSliceDecoder(t reflect.Type, decodeElem DecodeHandler) sliceDecoder {
	return sliceDecoder{
		reflect2.Type2(t).(*reflect2.UnsafeSliceType),
		t.Elem(),
		reflect2.Type2(reflect.ArrayOf(0, t.Elem())).UnsafeNew(),
		decodeElem,
	}
}

func getSliceDecoder(t reflect.Type) ValueDecoder {
	et := t.Elem()
	if et.Kind() == reflect.Uint8 {
		return bytesDecoder{t}
	}
	return makeSliceDecoder(t, GetDecodeHandler(et))
}

var (
	ifsdec sliceDecoder
)

func init() {
	ifsdec = makeSliceDecoder(reflect.TypeOf(([]interface{})(nil)), interfaceDecode)
	registerValueDecoder(ifsdec.t.Type1(), ifsdec)
}
