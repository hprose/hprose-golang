/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/common.go                                            |
|                                                          |
| LastModified: Feb 18, 2024                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package rpc

import "github.com/hprose/hprose-golang/v3/rpc/core"

type (
	// Context for RPC.
	Context = core.Context
	// ClientContext for RPC.
	ClientContext = core.ClientContext
	// ServiceContext for RPC.
	ServiceContext = core.ServiceContext
	// InvalidRequestError represents a error.
	InvalidRequestError = core.InvalidRequestError
	// InvalidResponseError represents a error.
	InvalidResponseError = core.InvalidResponseError
	// PanicError represents a panic error.
	PanicError = core.PanicError
	// UnsupportedProtocolError represents a error.
	UnsupportedProtocolError = core.UnsupportedProtocolError
	// UnsupportedServerTypeError represents a error.
	UnsupportedServerTypeError = core.UnsupportedServerTypeError
	// NextIOHandler for RPC.
	NextIOHandler = core.NextIOHandler
	// NextInvokeHandler for RPC.
	NextInvokeHandler = core.NextInvokeHandler
	// Method for RPC.
	Method = core.Method
	// ClientCodec for RPC.
	ClientCodec = core.ClientCodec
	// ServiceCodec for RPC.
	ServiceCodec = core.ServiceCodec
	// CodecOption for clientCodec & serviceCodec.
	CodecOption = core.CodecOption
	// WorkerPool interface
	WorkerPool = core.WorkerPool
)

var (
	// ErrClosed represents a error.
	ErrClosed = core.ErrClosed
	// ErrRequestEntityTooLarge represents a error.
	ErrRequestEntityTooLarge = core.ErrRequestEntityTooLarge
	// ErrTimeout represents a error.
	ErrTimeout = core.ErrTimeout
	// IsTemporaryError returns true if err is a temporary error.
	IsTemporaryError = core.IsTemporaryError
	// IsTimeoutError returns true if err is a timeout error.
	IsTimeoutError = core.IsTimeoutError
	// RegisterHandler for Service.
	RegisterHandler = core.RegisterHandler
	// RegisterTransport for Client.
	RegisterTransport = core.RegisterTransport
	// WithContext returns a copy of the parent context and associates it with a rpc.Context.
	WithContext = core.WithContext
	// GetClientContext returns the *rpc.ClientContext bound to the context.
	GetClientContext = core.GetClientContext
	// NewClientContext returns a rpc.ClientContext.
	NewClientContext = core.NewClientContext
	// GetServiceContext returns the *rpc.ServiceContext bound to the context.
	GetServiceContext = core.GetServiceContext
	// NewServiceContext returns a rpc.ServiceContext.
	NewServiceContext = core.NewServiceContext
	// FromContext returns the rpc.Context bound to the context.
	FromContext = core.FromContext
	// NewContext returns a rpc.Context.
	NewContext = core.NewContext
	// NewPanicError return a panic error.
	NewPanicError = core.NewPanicError
	// MissingMethod returns a missing Method object.
	MissingMethod = core.MissingMethod
	// NewMethod returns a Method object.
	NewMethod = core.NewMethod
	// NewClientCodec returns the ClientCodec.
	NewClientCodec = core.NewClientCodec
	// NewServiceCodec returns the ServiceCodec.
	NewServiceCodec = core.NewServiceCodec
	// WithDebug returns a debug Option for clientCodec & serviceCodec.
	WithDebug = core.WithDebug
	// WithSimple returns a simple Option for clientCodec & serviceCodec.
	WithSimple = core.WithSimple
	// WithLongType returns a longType Option for clientCodec & serviceCodec.
	WithLongType = core.WithLongType
	// WithRealType returns a realType Option for clientCodec & serviceCodec.
	WithRealType = core.WithRealType
	// WithMapType returns a mapType Option for clientCodec & serviceCodec.
	WithMapType = core.WithMapType
	// WithStructType returns a structType Option for clientCodec & serviceCodec.
	WithStructType = core.WithStructType
	// WithListType returns a listType Option for clientCodec & serviceCodec.
	WithListType = core.WithListType
)
