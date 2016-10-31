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
 * io/decoder.go                                          *
 *                                                        *
 * hprose decoder for Go.                                 *
 *                                                        *
 * LastModified: Oct 25, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import (
	"errors"
	"reflect"
)

type valueDecoder func(r *Reader, v reflect.Value, tag byte)

var valueDecoders []valueDecoder

func nilDecoder(r *Reader, v reflect.Value) {
	v.Set(reflect.Zero(v.Type()))
}

func invalidDecoder(r *Reader, v reflect.Value, tag byte) {
	panic(errors.New("can't unserialize the type: " + v.Type().String()))
}
func init() {
	valueDecoders = []valueDecoder{
		reflect.Invalid:       invalidDecoder,
		reflect.Bool:          boolDecoder,
		reflect.Int:           intDecoder,
		reflect.Int8:          intDecoder,
		reflect.Int16:         intDecoder,
		reflect.Int32:         intDecoder,
		reflect.Int64:         intDecoder,
		reflect.Uint:          uintDecoder,
		reflect.Uint8:         uintDecoder,
		reflect.Uint16:        uintDecoder,
		reflect.Uint32:        uintDecoder,
		reflect.Uint64:        uintDecoder,
		reflect.Uintptr:       uintDecoder,
		reflect.Float32:       float32Decoder,
		reflect.Float64:       float64Decoder,
		reflect.Complex64:     complex64Decoder,
		reflect.Complex128:    complex128Decoder,
		reflect.Array:         arrayDecoder,
		reflect.Chan:          invalidDecoder,
		reflect.Func:          invalidDecoder,
		reflect.Interface:     interfaceDecoder,
		reflect.Map:           mapDecoder,
		reflect.Ptr:           ptrDecoder,
		reflect.Slice:         sliceDecoder,
		reflect.String:        stringDecoder,
		reflect.Struct:        structDecoder,
		reflect.UnsafePointer: invalidDecoder,
	}
}
