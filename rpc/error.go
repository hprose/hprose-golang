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
 * rpc/error.go                                           *
 *                                                        *
 * rpc error for Go.                                      *
 *                                                        *
 * LastModified: Nov 1, 2016                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import (
	"errors"
	"fmt"
	"runtime"
)

// ErrTimeout represents a timeout error
var ErrTimeout = errors.New("timeout")

// ErrServerIsAlreadyStarted represents a error
var ErrServerIsAlreadyStarted = errors.New("The server is already started")

// ErrServerIsNotStarted represents a error
var ErrServerIsNotStarted = errors.New("The server is not started")

// ErrClientIsAlreadyClosed represents a error
var ErrClientIsAlreadyClosed = errors.New("The Client is already closed")

var errURIListEmpty = errors.New("uriList must contain at least one uri")
var errNotSupportMultpleProtocol = errors.New("Not support multiple protocol.")

// PanicError represents a panic error
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

// NewPanicError return a panic error
func NewPanicError(v interface{}) *PanicError {
	return &PanicError{v, stack()}
}

// Error implements the PanicError Error method.
func (pe *PanicError) Error() string {
	return fmt.Sprintf("%v", pe.Panic)
}
