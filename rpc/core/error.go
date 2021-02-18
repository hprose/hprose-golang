/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/error.go                                        |
|                                                          |
| LastModified: Feb 18, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
)

var errorType = reflect.TypeOf((*error)(nil)).Elem()

// ErrTimeout represents a timeout error.
var ErrTimeout = errors.New("timeout")

// // ErrServerIsAlreadyStarted represents a error.
// var ErrServerIsAlreadyStarted = errors.New("The server is already started")

// // ErrServerIsNotStarted represents a error.
// var ErrServerIsNotStarted = errors.New("The server is not started")

// // ErrClientIsAlreadyClosed represents a error.
// var ErrClientIsAlreadyClosed = errors.New("The Client is already closed")

// // ErrURIListEmpty represents a error.
// var ErrURIListEmpty = errors.New("uriList must contain at least one uri")

// UnsupportedProtocolError represents a error.
type UnsupportedProtocolError struct {
	Scheme string
}

func (e UnsupportedProtocolError) Error() string {
	return "hprose/rpc/core: unsupported protocol: " + e.Scheme
}

// UnsupportedServerTypeError represents a error.
type UnsupportedServerTypeError struct {
	ServerType reflect.Type
}

func (e UnsupportedServerTypeError) Error() string {
	return "hprose/rpc/core: unsupported server type: " + e.ServerType.String()
}

// PanicError represents a panic error.
type PanicError struct {
	Panic interface{}
	Stack []byte
}

func stack() []byte {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			return buf[:n]
		}
		buf = make([]byte, 2*len(buf))
	}
}

// NewPanicError return a panic error.
func NewPanicError(v interface{}) *PanicError {
	return &PanicError{v, stack()}
}

// Error implements the PanicError Error method.
func (pe *PanicError) Error() string {
	return fmt.Sprintf("%v", pe.Panic)
}

// String returns the panic error message and stack.
func (pe *PanicError) String() string {
	return fmt.Sprintf("%v\r\n%s", pe.Panic, pe.Stack)
}
