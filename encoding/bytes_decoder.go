/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/bytes_decoder.go                                |
|                                                          |
| LastModified: Jun 2, 2020                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"reflect"

	"github.com/modern-go/reflect2"
)

// bytesDecoder is the implementation of ValueDecoder for []byte.
type bytesDecoder struct {
	destType reflect.Type
}

var bytesdec = bytesDecoder{reflect.TypeOf(([]byte)(nil))}

func (valdec bytesDecoder) decode(dec *Decoder, tag byte) []byte {
	switch tag {
	case TagNull:
		return nil
	case TagEmpty:
		return []byte{}
	case TagBytes:
		return dec.ReadBytes()
	case TagList:
		return dec.readListAsBytes()
	case TagUTF8Char:
		return dec.readStringAsBytes(1)
	case TagString:
		return reflect2.UnsafeCastString(dec.ReadString())
	default:
		dec.decodeError(valdec.destType, tag)
	}
	return nil
}

func (valdec bytesDecoder) decodeValue(dec *Decoder, pv *[]byte, tag byte) {
	if bytes := valdec.decode(dec, tag); dec.Error == nil {
		*pv = bytes
	}
}

func (valdec bytesDecoder) decodePtr(dec *Decoder, pv **[]byte, tag byte) {
	if bytes := valdec.decode(dec, tag); dec.Error == nil {
		*pv = &bytes
	}
}

func (valdec bytesDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	switch pv := p.(type) {
	case *[]byte:
		valdec.decodeValue(dec, pv, tag)
	case **[]byte:
		valdec.decodePtr(dec, pv, tag)
	}
}

func (dec *Decoder) readBytes() []byte {
	bytes := dec.Next(dec.ReadInt())
	dec.Skip()
	return bytes
}

// ReadBytes reads bytes and add reference
func (dec *Decoder) ReadBytes() []byte {
	bytes := dec.readBytes()
	dec.AddReference(bytes)
	return bytes
}

func (dec *Decoder) readListAsBytes() []byte {
	count := dec.ReadInt()
	slice := make([]byte, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = uint8dec.decode(dec, dec.NextByte())
	}
	dec.Skip()
	return slice
}
