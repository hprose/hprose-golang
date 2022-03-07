/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| hprose.go                                                |
|                                                          |
| LastModified: Mar 7, 2022                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package hprose

import (
	"github.com/hprose/hprose-golang/v3/io"
	"github.com/hprose/hprose-golang/v3/rpc"
	"github.com/hprose/hprose-golang/v3/rpc/codec/jsonrpc"
	"github.com/hprose/hprose-golang/v3/rpc/http/cookie"
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
	Service                    = rpc.Service

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

	ExecuteTimeout      = timeout.ExecuteTimeout
	CookieManagerOption = cookie.CookieManagerOption
)

const (
	// Serialize Type.
	TagInteger  = io.TagInteger
	TagLong     = io.TagLong
	TagDouble   = io.TagDouble
	TagNull     = io.TagNull
	TagEmpty    = io.TagEmpty
	TagTrue     = io.TagTrue
	TagFalse    = io.TagFalse
	TagNaN      = io.TagNaN
	TagInfinity = io.TagInfinity
	TagDate     = io.TagDate
	TagTime     = io.TagTime
	TagUTC      = io.TagUTC
	TagBytes    = io.TagBytes
	TagUTF8Char = io.TagUTF8Char
	TagString   = io.TagString
	TagGUID     = io.TagGUID
	TagList     = io.TagList
	TagMap      = io.TagMap
	TagClass    = io.TagClass
	TagObject   = io.TagObject
	TagRef      = io.TagRef

	// Serialize Marks.
	TagPos        = io.TagPos
	TagNeg        = io.TagNeg
	TagSemicolon  = io.TagSemicolon
	TagOpenbrace  = io.TagOpenbrace
	TagClosebrace = io.TagClosebrace
	TagQuote      = io.TagQuote
	TagPoint      = io.TagPoint

	// Protocol Tags.
	TagHeader = io.TagHeader
	TagCall   = io.TagCall
	TagResult = io.TagResult
	TagError  = io.TagError
	TagEnd    = io.TagEnd

	LongTypeInt      = io.LongTypeInt
	LongTypeUint     = io.LongTypeUint
	LongTypeInt64    = io.LongTypeInt64
	LongTypeUint64   = io.LongTypeUint64
	LongTypeBigInt   = io.LongTypeBigInt
	RealTypeFloat64  = io.RealTypeFloat64
	RealTypeFloat32  = io.RealTypeFloat32
	RealTypeBigFloat = io.RealTypeBigFloat
	MapTypeIIMap     = io.MapTypeIIMap
	MapTypeSIMap     = io.MapTypeSIMap

	NoCookieManager     = cookie.NoCookieManager
	GlobalCookieManager = cookie.GlobalCookieManager
	ClientCookieManager = cookie.ClientCookieManager
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
	GetEncoder           = io.GetEncoder
	FreeEncoder          = io.FreeEncoder
	GetDecoder           = io.GetDecoder
	FreeDecoder          = io.FreeDecoder

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
	FastHTTPTransport        = rpc.FastHTTPTransport
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
