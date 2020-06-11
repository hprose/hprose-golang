/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/error.go                                        |
|                                                          |
| LastModified: Apr 25, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"errors"
	"reflect"
)

// An UnsupportedTypeError is returned by Encoder when attempting
// to encode an unsupported value type.
type UnsupportedTypeError struct {
	Type reflect.Type
}

func (e UnsupportedTypeError) Error() string {
	return "hprose/encoding: unsupported type: " + e.Type.String()
}

// ErrInvalidUTF8 means that a decoder encountered invalid UTF-8.
var ErrInvalidUTF8 = errors.New("hprose/encoding: invalid UTF-8")

// A CastError is returned by Decoder when can not cast source type to destination type.
type CastError struct {
	Source      reflect.Type
	Destination reflect.Type
}

func (e CastError) Error() string {
	if e.Source == nil {
		return "hprose/encoding: can not cast nil to " + e.Destination.String()
	}
	if e.Destination == nil {
		return "hprose/encoding: can not cast " + e.Source.String() + " to nil"
	}
	return "hprose/encoding: can not cast " + e.Source.String() + " to " + e.Destination.String()
}

// DecodeError is returned by Decoder when the data is wrong.
type DecodeError string

func (e DecodeError) Error() string {
	return string(e)
}
