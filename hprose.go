/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| hprose.go                                                |
|                                                          |
| LastModified: May 24, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package hprose

import (
	"github.com/hprose/hprose-golang/v3/io"
	"github.com/hprose/hprose-golang/v3/rpc"
	"github.com/hprose/hprose-golang/v3/rpc/codec/jsonrpc"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/circuitbreaker"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/cluster"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/forward"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/limiter"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/loadbalance"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/log"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/oneway"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/push"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/reverse"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/timeout"
)

type (
	Decoder              = io.Decoder
	Encoder              = io.Encoder
	Formatter            = io.Formatter
	LongType             = io.LongType
	MapType              = io.MapType
	RealType             = io.RealType
	CastError            = io.CastError
	DecodeError          = io.DecodeError
	UnsupportedTypeError = io.UnsupportedTypeError
	ValueDecoder         = io.ValueDecoder
	ValueEncoder         = io.ValueEncoder

	Context                    = rpc.Context
	ClientContext              = rpc.ClientContext
	ServiceContext             = rpc.ServiceContext
	InvalidRequestError        = rpc.InvalidRequestError
	InvalidResponseError       = rpc.InvalidResponseError
	PanicError                 = rpc.PanicError
	UnsupportedProtocolError   = rpc.UnsupportedProtocolError
	UnsupportedServerTypeError = rpc.UnsupportedServerTypeError
	NextIOHandler              = rpc.NextIOHandler
	NextInvokeHandler          = rpc.NextInvokeHandler
	Method                     = rpc.Method
	ClientCodec                = rpc.ClientCodec
	ServiceCodec               = rpc.ServiceCodec
	CodecOption                = rpc.CodecOption
	Client                     = rpc.Client
	TransportGetter            = rpc.TransportGetter
	Service                    = rpc.Service
	HandlerGetter              = rpc.HandlerGetter

	JSONRPCClientCodec  = jsonrpc.ClientCodec
	JSONRPCServiceCodec = jsonrpc.ServiceCodec

	CircuitBreaker = circuitbreaker.CircuitBreaker

	Cluster = cluster.Cluster

	Forward = forward.Forward

	ConcurrentLimiter = limiter.ConcurrentLimiter
	RateLimiter       = limiter.RateLimiter

	LeastActiveLoadBalance         = loadbalance.LeastActiveLoadBalance
	NginxRoundRobinLoadBalance     = loadbalance.NginxRoundRobinLoadBalance
	RandomLoadBalance              = loadbalance.RandomLoadBalance
	RoundRobinLoadBalance          = loadbalance.RoundRobinLoadBalance
	WeightedLeastActiveLoadBalance = loadbalance.WeightedLeastActiveLoadBalance
	WeightedRandomLoadBalance      = loadbalance.WeightedRandomLoadBalance
	WeightedRoundRobinLoadBalance  = loadbalance.WeightedRoundRobinLoadBalance

	Log = log.Log

	Oneway = oneway.Oneway

	Broker   = push.Broker
	Message  = push.Message
	Producer = push.Producer
	Prosumer = push.Prosumer

	Caller          = reverse.Caller
	ProviderContext = reverse.ProviderContext
	Provider        = reverse.Provider

	ExecuteTimeout = timeout.ExecuteTimeout
)

const (
	// Serialize Type.
	TagInteger  byte = io.TagInteger
	TagLong     byte = io.TagLong
	TagDouble   byte = io.TagDouble
	TagNull     byte = io.TagNull
	TagEmpty    byte = io.TagEmpty
	TagTrue     byte = io.TagTrue
	TagFalse    byte = io.TagFalse
	TagNaN      byte = io.TagNaN
	TagInfinity byte = io.TagInfinity
	TagDate     byte = io.TagDate
	TagTime     byte = io.TagTime
	TagUTC      byte = io.TagUTC
	TagBytes    byte = io.TagBytes
	TagUTF8Char byte = io.TagUTF8Char
	TagString   byte = io.TagString
	TagGUID     byte = io.TagGUID
	TagList     byte = io.TagList
	TagMap      byte = io.TagMap
	TagClass    byte = io.TagClass
	TagObject   byte = io.TagObject
	TagRef      byte = io.TagRef

	// Serialize Marks.
	TagPos        byte = io.TagPos
	TagNeg        byte = io.TagNeg
	TagSemicolon  byte = io.TagSemicolon
	TagOpenbrace  byte = io.TagOpenbrace
	TagClosebrace byte = io.TagClosebrace
	TagQuote      byte = io.TagQuote
	TagPoint      byte = io.TagPoint

	// Protocol Tags.
	TagHeader byte = io.TagHeader
	TagCall   byte = io.TagCall
	TagResult byte = io.TagResult
	TagError  byte = io.TagError
	TagEnd    byte = io.TagEnd
)

var (
	ErrInvalidUTF8       = io.ErrInvalidUTF8
	Marshal              = io.Marshal
	Unmarshal            = io.Unmarshal
	Register             = io.Register
	RegisterName         = io.RegisterName
	RegisterValueDecoder = io.RegisterValueDecoder
	RegisterValueEncoder = io.RegisterValueEncoder
	RegisterConverter    = io.RegisterConverter
	GetConverter         = io.GetConverter
	Convert              = io.Convert
	GetStructType        = io.GetStructType
	NewDecoder           = io.NewDecoder
	NewDecoderFromReader = io.NewDecoderFromReader
	NewEncoder           = io.NewEncoder
	GetValueDecoder      = io.GetValueDecoder
	GetValueEncoder      = io.GetValueEncoder

	ErrClosed                = rpc.ErrClosed
	ErrRequestEntityTooLarge = rpc.ErrRequestEntityTooLarge
	ErrTimeout               = rpc.ErrTimeout
	IsTemporaryError         = rpc.IsTemporaryError
	IsTimeoutError           = rpc.IsTimeoutError
	RegisterHandler          = rpc.RegisterHandler
	RegisterTransport        = rpc.RegisterTransport
	WithContext              = rpc.WithContext
	GetClientContext         = rpc.GetClientContext
	NewClientContext         = rpc.NewClientContext
	GetServiceContext        = rpc.GetServiceContext
	NewServiceContext        = rpc.NewServiceContext
	FromContext              = rpc.FromContext
	NewContext               = rpc.NewContext
	NewPanicError            = rpc.NewPanicError
	MissingMethod            = rpc.MissingMethod
	NewMethod                = rpc.NewMethod
	NewClientCodec           = rpc.NewClientCodec
	NewServiceCodec          = rpc.NewServiceCodec
	WithDebug                = rpc.WithDebug
	WithLongType             = rpc.WithLongType
	WithMapType              = rpc.WithMapType
	WithRealType             = rpc.WithRealType
	WithSimple               = rpc.WithSimple
	NewClient                = rpc.NewClient
	NewService               = rpc.NewService
	HTTPTransport            = rpc.HTTPTransport
	SocketTransport          = rpc.SocketTransport
	UDPTransport             = rpc.UDPTransport
	WebSocketTransport       = rpc.WebSocketTransport
	HTTPHandler              = rpc.HTTPHandler
	SocketHandler            = rpc.SocketHandler
	UDPHandler               = rpc.UDPHandler
	WebSocketHandler         = rpc.WebSocketHandler

	NewJSONRPCClientCodec  = jsonrpc.NewClientCodec
	NewJSONRPCServiceCodec = jsonrpc.NewServiceCodec

	WithThreshold     = circuitbreaker.WithThreshold
	WithRecoverTime   = circuitbreaker.WithRecoverTime
	WithMockService   = circuitbreaker.WithMockService
	NewCircuitBreaker = circuitbreaker.New

	WithRetry       = cluster.WithRetry
	WithIdempotent  = cluster.WithIdempotent
	WithMinInterval = cluster.WithMinInterval
	WithMaxInterval = cluster.WithMaxInterval
	FailoverConfig  = cluster.FailoverConfig
	FailtryConfig   = cluster.FailtryConfig
	FailfastConfig  = cluster.FailfastConfig
	NewCluster      = cluster.New
	ForkingPlugin   = cluster.Forking
	BroadcastPlugin = cluster.Broadcast

	NewForward = forward.New

	WithMaxPermits       = limiter.WithMaxPermits
	WithTimeout          = limiter.WithTimeout
	NewRateLimiter       = limiter.NewRateLimiter
	NewConcurrentLimiter = limiter.NewConcurrentLimiter

	NewLeastActiveLoadBalance         = loadbalance.NewLeastActiveLoadBalance
	NewRandomLoadBalance              = loadbalance.NewRandomLoadBalance
	NewRoundRobinLoadBalance          = loadbalance.NewRoundRobinLoadBalance
	NewNginxRoundRobinLoadBalance     = loadbalance.NewNginxRoundRobinLoadBalance
	NewWeightedLeastActiveLoadBalance = loadbalance.NewWeightedLeastActiveLoadBalance
	NewWeightedRandomLoadBalance      = loadbalance.NewWeightedRandomLoadBalance
	NewWeightedRoundRobinLoadBalance  = loadbalance.NewWeightedRoundRobinLoadBalance

	NewLog    = log.New
	LogPlugin = log.Plugin

	NewBroker   = push.NewBroker
	GetProducer = push.GetProducer
	NewProsumer = push.NewProsumer

	NewCaller          = reverse.NewCaller
	UseService         = reverse.UseService
	GetProviderContext = reverse.GetProviderContext
	NewProvider        = reverse.NewProvider

	NewExecuteTimeout = timeout.New
)
