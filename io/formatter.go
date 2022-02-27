/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/formatter.go                                          |
|                                                          |
| LastModified: Feb 27, 2022                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"io"
	"sync"
)

type Formatter struct {
	Simple bool
	LongType
	RealType
	MapType
}

var encoderPool = sync.Pool{New: func() interface{} { return new(Encoder) }}

func (f Formatter) Marshal(v interface{}) ([]byte, error) {
	encoder := encoderPool.Get().(*Encoder).Simple(f.Simple).ResetBuffer()
	defer encoderPool.Put(encoder)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}
	return encoder.Bytes(), nil
}

var decoderPool = sync.Pool{New: func() interface{} { return new(Decoder) }}

func (f Formatter) Unmarshal(data []byte, v interface{}) error {
	var decoder *Decoder
	if f.Simple {
		decoder = NewDecoder(data)
	} else {
		decoder = decoderPool.Get().(*Decoder).Simple(f.Simple).ResetBytes(data)
		defer func() {
			decoder.ResetBytes(nil)
			decoderPool.Put(decoder)
		}()
	}
	decoder.LongType = f.LongType
	decoder.RealType = f.RealType
	decoder.MapType = f.MapType
	decoder.Decode(v)
	return decoder.Error
}

func (f Formatter) UnmarshalFromReader(reader io.Reader, v interface{}) error {
	decoder := decoderPool.Get().(*Decoder).Simple(f.Simple).ResetReader(reader)
	defer func() {
		decoder.ResetReader(nil)
		decoderPool.Put(decoder)
	}()
	decoder.LongType = f.LongType
	decoder.RealType = f.RealType
	decoder.MapType = f.MapType
	decoder.Decode(v)
	return decoder.Error
}

var defaultFormatter = Formatter{Simple: true}

func Marshal(v interface{}) ([]byte, error) {
	return defaultFormatter.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return defaultFormatter.Unmarshal(data, v)
}

func UnmarshalFromReader(reader io.Reader, v interface{}) error {
	return defaultFormatter.UnmarshalFromReader(reader, v)
}
