/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/encoder/big_marshaler.go                     |
|                                                          |
| LastModified: Feb 25, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoder

import (
	"math/big"
	"reflect"
)

func writeBigInt(enc *Encoder, i interface{}) (err error) {
	switch i := i.(type) {
	case *big.Int:
		return WriteBigInt(enc.Writer, i)
	case big.Int:
		return WriteBigInt(enc.Writer, &i)
	}
	return &UnsupportedTypeError{Type: reflect.TypeOf(i)}
}

func writeBigFloat(enc *Encoder, i interface{}) (err error) {
	switch f := i.(type) {
	case *big.Float:
		return WriteBigFloat(enc.Writer, f)
	case big.Float:
		return WriteBigFloat(enc.Writer, &f)
	}
	return &UnsupportedTypeError{Type: reflect.TypeOf(i)}
}

func writeBigRat(enc *Encoder, i interface{}) (err error) {
	switch r := i.(type) {
	case *big.Rat:
		return WriteBigRat(enc, r)
	case big.Rat:
		return WriteBigRat(enc, &r)
	}
	return &UnsupportedTypeError{Type: reflect.TypeOf(i)}
}

func init() {
	RegisterValueMarshaler(reflect.TypeOf((*big.Int)(nil)).Elem(), writeBigInt)
	RegisterValueMarshaler(reflect.TypeOf((*big.Float)(nil)).Elem(), writeBigFloat)
	RegisterValueMarshaler(reflect.TypeOf((*big.Rat)(nil)).Elem(), writeBigRat)
}
