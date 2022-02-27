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
)

type Formatter struct {
	Simple bool
	LongType
	RealType
	MapType
}

func (f Formatter) Marshal(v interface{}) ([]byte, error) {
	encoder := GetEncoder().Simple(f.Simple)
	defer FreeEncoder(encoder)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}
	return encoder.Bytes(), nil
}

func (f Formatter) Unmarshal(data []byte, v interface{}) error {
	var decoder *Decoder
	if f.Simple {
		decoder = NewDecoder(data)
	} else {
		decoder = GetDecoder().Simple(f.Simple).ResetBytes(data)
		defer FreeDecoder(decoder)
	}
	decoder.LongType = f.LongType
	decoder.RealType = f.RealType
	decoder.MapType = f.MapType
	decoder.Decode(v)
	return decoder.Error
}

func (f Formatter) UnmarshalFromReader(reader io.Reader, v interface{}) error {
	decoder := GetDecoder().Simple(f.Simple).ResetReader(reader)
	defer FreeDecoder(decoder)
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
