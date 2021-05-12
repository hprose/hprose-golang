/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/common.go                                            |
|                                                          |
| LastModified: May 12, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package rpc

import "github.com/hprose/hprose-golang/v3/rpc/core"

type (
	Context                    = core.Context
	ClientContext              = core.ClientContext
	ServiceContext             = core.ServiceContext
	InvalidRequestError        = core.InvalidRequestError
	InvalidResponseError       = core.InvalidResponseError
	PanicError                 = core.PanicError
	UnsupportedProtocolError   = core.UnsupportedProtocolError
	UnsupportedServerTypeError = core.UnsupportedServerTypeError
	NextIOHandler              = core.NextIOHandler
	NextInvokeHandler          = core.NextInvokeHandler
	Method                     = core.Method
	ClientCodec                = core.ClientCodec
	ServiceCodec               = core.ServiceCodec
	CodecOption                = core.CodecOption
)

var (
	ErrClosed                = core.ErrClosed
	ErrRequestEntityTooLarge = core.ErrRequestEntityTooLarge
	ErrTimeout               = core.ErrTimeout
	IsTemporaryError         = core.IsTemporaryError
	IsTimeoutError           = core.IsTimeoutError
	RegisterHandler          = core.RegisterHandler
	RegisterTransport        = core.RegisterTransport
	WithContext              = core.WithContext
	GetClientContext         = core.GetClientContext
	NewClientContext         = core.NewClientContext
	GetServiceContext        = core.GetServiceContext
	NewServiceContext        = core.NewServiceContext
	FromContext              = core.FromContext
	NewContext               = core.NewContext
	NewPanicError            = core.NewPanicError
	MissingMethod            = core.MissingMethod
	NewMethod                = core.NewMethod
	NewClientCodec           = core.NewClientCodec
	NewServiceCodec          = core.NewServiceCodec
	WithDebug                = core.WithDebug
	WithLongType             = core.WithLongType
	WithMapType              = core.WithMapType
	WithRealType             = core.WithRealType
	WithSimple               = core.WithSimple
)
