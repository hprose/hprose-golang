/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/formatter.go                                          |
|                                                          |
| LastModified: May 16, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

func Marshal(v interface{}) ([]byte, error) {
	encoder := new(Encoder)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}
	return encoder.Bytes(), nil
}

func Unmarshal(data []byte, v interface{}) error {
	decoder := NewDecoder(data)
	decoder.Decode(v)
	return decoder.Error
}
