/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/formatter.go                                          |
|                                                          |
| LastModified: May 17, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

type Formatter struct {
	Simple bool
	LongType
	RealType
	MapType
}

func (f Formatter) Marshal(v interface{}) ([]byte, error) {
	encoder := new(Encoder).Simple(f.Simple)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}
	return encoder.Bytes(), nil
}

func (f Formatter) Unmarshal(data []byte, v interface{}) error {
	decoder := NewDecoder(data).Simple(f.Simple)
	decoder.LongType = f.LongType
	decoder.RealType = f.RealType
	decoder.MapType = f.MapType
	decoder.Decode(v)
	return decoder.Error
}

var defaultFormatter = Formatter{}

func Marshal(v interface{}) ([]byte, error) {
	return defaultFormatter.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return defaultFormatter.Unmarshal(data, v)
}
