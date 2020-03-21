/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/error.go                                        |
|                                                          |
| LastModified: Mar 19, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"errors"
	"reflect"
)

// An UnsupportedTypeError is returned by Encode/Marshal when attempting
// to encode an unsupported value type.
type UnsupportedTypeError struct {
	Type reflect.Type
}

func (e *UnsupportedTypeError) Error() string {
	return "hprose: unsupported type: " + e.Type.String()
}

// ErrInvalidUTF8 means that the string is invalid UTF-8.
var ErrInvalidUTF8 = errors.New("encoding: invalid UTF-8")
