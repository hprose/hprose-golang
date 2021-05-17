/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/codec_option.go                                 |
|                                                          |
| LastModified: Mar 25, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import "github.com/hprose/hprose-golang/v3/io"

// CodecOption for clientCodec & serviceCodec.
type CodecOption func(interface{})

// WithDebug returns a debug Option for clientCodec & serviceCodec.
func WithDebug(debug bool) CodecOption {
	return func(c interface{}) {
		if c, ok := c.(*serviceCodec); ok {
			c.Debug = debug
		}
	}
}

// WithSimple returns a simple Option for clientCodec & serviceCodec.
func WithSimple(simple bool) CodecOption {
	return func(c interface{}) {
		switch c := c.(type) {
		case *serviceCodec:
			c.Simple = simple
		case *clientCodec:
			c.Simple = simple
		}
	}
}

// WithDebug returns a longType Option for clientCodec & serviceCodec.
func WithLongType(longType io.LongType) CodecOption {
	return func(c interface{}) {
		switch c := c.(type) {
		case *serviceCodec:
			c.LongType = longType
		case *clientCodec:
			c.LongType = longType
		}
	}
}

// WithRealType returns a realType Option for clientCodec & serviceCodec.
func WithRealType(realType io.RealType) CodecOption {
	return func(c interface{}) {
		switch c := c.(type) {
		case *serviceCodec:
			c.RealType = realType
		case *clientCodec:
			c.RealType = realType
		}
	}
}

// WithMapType returns a mapType Option for clientCodec & serviceCodec.
func WithMapType(mapType io.MapType) CodecOption {
	return func(c interface{}) {
		switch c := c.(type) {
		case *serviceCodec:
			c.MapType = mapType
		case *clientCodec:
			c.MapType = mapType
		}
	}
}
