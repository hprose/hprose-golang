/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * io/encoder.go                                          *
 *                                                        *
 * hprose encoder for Go.                                 *
 *                                                        *
 * LastModified: Oct 19, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import (
	"container/list"
	"math/big"
	"reflect"
	"time"
	"unsafe"

	"github.com/hprose/hprose-golang/util"
)

type valueEncoder func(w *Writer, v reflect.Value)

var valueEncoders []valueEncoder

func nilEncoder(w *Writer, v reflect.Value) {
	w.WriteNil()
}

func boolEncoder(w *Writer, v reflect.Value) {
	w.WriteBool(v.Bool())
}

func intEncoder(w *Writer, v reflect.Value) {
	w.WriteInt(v.Int())
}

func uintEncoder(w *Writer, v reflect.Value) {
	w.WriteUint(v.Uint())
}

func float32Encoder(w *Writer, v reflect.Value) {
	w.WriteFloat(v.Float(), 32)
}

func float64Encoder(w *Writer, v reflect.Value) {
	w.WriteFloat(v.Float(), 64)
}

func complex64Encoder(w *Writer, v reflect.Value) {
	w.WriteComplex64(complex64(v.Complex()))
}

func complex128Encoder(w *Writer, v reflect.Value) {
	w.WriteComplex128(v.Complex())
}

func interfaceEncoder(w *Writer, v reflect.Value) {
	if v.IsNil() {
		w.WriteNil()
		return
	}
	e := v.Elem()
	valueEncoders[e.Kind()](w, e)
}

func arrayEncoder(w *Writer, v reflect.Value) {
	setWriterRef(w, nil)
	writeArray(w, v)
}

func sliceEncoder(w *Writer, v reflect.Value) {
	setWriterRef(w, nil)
	writeSlice(w, v)
}

func mapEncoder(w *Writer, v reflect.Value) {
	ptr := (*reflectValue)(unsafe.Pointer(&v)).ptr
	if !writeRef(w, ptr) {
		setWriterRef(w, ptr)
		writeMap(w, v)
	}
}

func stringEncoder(w *Writer, v reflect.Value) {
	w.WriteString(v.String())
}

func structEncoder(w *Writer, v reflect.Value) {
	ptr := (*reflectValue)(unsafe.Pointer(&v)).ptr
	structPtrEncoder(w, v, ptr)
}

func arrayPtrEncoder(w *Writer, v reflect.Value, ptr unsafe.Pointer) {
	if !writeRef(w, ptr) {
		setWriterRef(w, ptr)
		writeArray(w, v)
	}
}

func mapPtrEncoder(w *Writer, v reflect.Value, ptr unsafe.Pointer) {
	if !writeRef(w, ptr) {
		setWriterRef(w, ptr)
		writeMap(w, v)
	}
}

func slicePtrEncoder(w *Writer, v reflect.Value, ptr unsafe.Pointer) {
	if !writeRef(w, ptr) {
		setWriterRef(w, ptr)
		writeSlice(w, v)
	}
}

func stringPtrEncoder(w *Writer, v reflect.Value, ptr unsafe.Pointer) {
	str := v.String()
	length := util.UTF16Length(str)
	switch {
	case length == 0:
		w.writeByte(TagEmpty)
	case length < 0:
		w.WriteBytes(*(*[]byte)(unsafe.Pointer(&str)))
	case length == 1:
		w.writeByte(TagUTF8Char)
		w.writeString(str)
	default:
		if !writeRef(w, ptr) {
			setWriterRef(w, ptr)
			writeString(w, str, length)
		}
	}
}

func structPtrEncoder(w *Writer, v reflect.Value, ptr unsafe.Pointer) {
	switch *(*uintptr)(unsafe.Pointer(&v)) {
	case bigIntType:
		w.WriteBigInt((*big.Int)(ptr))
	case bigRatType:
		w.WriteBigRat((*big.Rat)(ptr))
	case bigFloatType:
		w.WriteBigFloat((*big.Float)(ptr))
	case timeType:
		w.WriteTime((*time.Time)(ptr))
	case listType:
		w.WriteList((*list.List)(ptr))
	case reflectValueType:
		w.WriteValue(*(*reflect.Value)(ptr))
	default:
		if !writeRef(w, ptr) {
			writeStruct(w, v)
		}
	}
}

func ptrEncoder(w *Writer, v reflect.Value) {
	if v.IsNil() {
		w.WriteNil()
		return
	}
	e := v.Elem()
	kind := e.Kind()
	ptr := (*reflectValue)(unsafe.Pointer(&e)).ptr
	switch kind {
	case reflect.Array:
		arrayPtrEncoder(w, e, ptr)
	case reflect.Map:
		mapPtrEncoder(w, e, ptr)
	case reflect.Slice:
		slicePtrEncoder(w, e, ptr)
	case reflect.String:
		stringPtrEncoder(w, e, ptr)
	case reflect.Struct:
		structPtrEncoder(w, e, ptr)
	default:
		valueEncoders[kind](w, e)
	}
}

func init() {
	valueEncoders = []valueEncoder{
		reflect.Invalid:       nilEncoder,
		reflect.Bool:          boolEncoder,
		reflect.Int:           intEncoder,
		reflect.Int8:          intEncoder,
		reflect.Int16:         intEncoder,
		reflect.Int32:         intEncoder,
		reflect.Int64:         intEncoder,
		reflect.Uint:          uintEncoder,
		reflect.Uint8:         uintEncoder,
		reflect.Uint16:        uintEncoder,
		reflect.Uint32:        uintEncoder,
		reflect.Uint64:        uintEncoder,
		reflect.Uintptr:       uintEncoder,
		reflect.Float32:       float32Encoder,
		reflect.Float64:       float64Encoder,
		reflect.Complex64:     complex64Encoder,
		reflect.Complex128:    complex128Encoder,
		reflect.Array:         arrayEncoder,
		reflect.Chan:          nilEncoder,
		reflect.Func:          nilEncoder,
		reflect.Interface:     interfaceEncoder,
		reflect.Map:           mapEncoder,
		reflect.Ptr:           ptrEncoder,
		reflect.Slice:         sliceEncoder,
		reflect.String:        stringEncoder,
		reflect.Struct:        structEncoder,
		reflect.UnsafePointer: nilEncoder,
	}
}
