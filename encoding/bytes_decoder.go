/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/bytes_decoder.go                                |
|                                                          |
| LastModified: Jun 1, 2020                                |
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

var bytesdec = bytesDecoder{reflect.TypeOf((*[]byte)(nil)).Elem()}

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

func (valdec bytesDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	if tag == TagNull {
		switch pv := p.(type) {
		case **[]byte:
			*pv = nil
		case *[]byte:
			*pv = nil
		}
		return
	}
	b := valdec.decode(dec, tag)
	if dec.Error != nil {
		return
	}
	switch pv := p.(type) {
	case **[]byte:
		*pv = &b
	case *[]byte:
		*pv = b
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
