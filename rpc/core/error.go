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

type timeout interface {
	Timeout() bool
}

// IsTimeoutError returns true if err is a timeout error.
func IsTimeoutError(err error) bool {
	t, ok := err.(timeout)
	return ok && t.Timeout()
}

type temporary interface {
	Temporary() bool
}

// IsTemporaryError returns true if err is a temporary error.
func IsTemporaryError(err error) bool {
	t, ok := err.(temporary)
	return ok && t.Temporary()
}

type timeoutError struct{}

func (e timeoutError) Error() string {
	return "timeout"
}

func (e timeoutError) Timeout() bool {
	return true
}

func (e timeoutError) Temporary() bool {
	return true
}

// ErrTimeout represents a error.
var ErrTimeout = timeoutError{}

// // ErrServerIsAlreadyStarted represents a error.
// var ErrServerIsAlreadyStarted = errors.New("The server is already started")

// ErrServerIsStoped represents a error.
var ErrServerIsStoped = errors.New("hprose/rpc/core: server is stoped")

// // ErrClientIsAlreadyClosed represents a error.
// var ErrClientIsAlreadyClosed = errors.New("The Client is already closed")

// // ErrURIListEmpty represents a error.
// var ErrURIListEmpty = errors.New("uriList must contain at least one uri")

// ErrRequestEntityTooLarge represents a error.
var ErrRequestEntityTooLarge = errors.New("Request entity too large")

// InvalidRequestError represents a error.
type InvalidRequestError struct {
	Request []byte
}

func (e InvalidRequestError) Error() string {
	return "hprose/rpc/core: invalid request:\r\n" + string(e.Request)
}

// InvalidResponseError represents a error.
type InvalidResponseError struct {
	Response []byte
}

func (e InvalidResponseError) Error() string {
	return "hprose/rpc/core: invalid response:\r\n" + string(e.Response)
}

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
