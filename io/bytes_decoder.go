/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/bytes_decoder.go                                      |
|                                                          |
| LastModified: Jun 5, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"reflect"

	"github.com/modern-go/reflect2"
)

func (dec *Decoder) readUnsafeBytes() []byte {
	bytes := dec.UnsafeNext(dec.ReadInt())
	dec.Skip()
	return bytes
}

func (dec *Decoder) readBytes() []byte {
	bytes := dec.Next(dec.ReadInt())
	dec.Skip()
	return bytes
}

// ReadBytes reads bytes and add reference.
func (dec *Decoder) ReadBytes() []byte {
	bytes := dec.readBytes()
	dec.AddReference(bytes)
	return bytes
}

func (dec *Decoder) readUint8Slice(et reflect.Type) []byte {
	count := dec.ReadInt()
	slice := make([]byte, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		slice[i] = dec.decodeUint8(et, dec.NextByte())
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) decodeBytes(t reflect.Type, tag byte) (result []byte) {
	switch tag {
	case TagNull:
		return nil
	case TagEmpty:
		return []byte{}
	case TagBytes:
		return dec.ReadBytes()
	case TagList:
		return dec.readUint8Slice(t.Elem())
	case TagUTF8Char:
		return dec.readStringAsSafeBytes(1)
	case TagString:
		if dec.IsSimple() {
			return dec.ReadStringAsBytes()
		}
		return reflect2.UnsafeCastString(dec.ReadString())
	case TagGUID:
		bytes, _ := dec.ReadUUID().MarshalBinary()
		return bytes
	default:
		dec.defaultDecode(t, &result, tag)
	}
	return
}

func (dec *Decoder) decodeBytesPtr(t reflect.Type, tag byte) *[]byte {
	if tag == TagNull {
		return nil
	}
	bytes := dec.decodeBytes(t, tag)
	return &bytes
}

// bytesDecoder is the implementation of ValueDecoder for []byte.
type bytesDecoder struct {
	t reflect.Type
}

func (valdec bytesDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*[]byte)(reflect2.PtrOf(p)) = dec.decodeBytes(valdec.t, tag)
}

// bytesPtrDecoder is the implementation of ValueDecoder for *[]byte.
type bytesPtrDecoder struct {
	t reflect.Type
}

func (valdec bytesPtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(**[]byte)(reflect2.PtrOf(p)) = dec.decodeBytesPtr(valdec.t, tag)
}

func init() {
	registerValueDecoder(bytesType, bytesDecoder{bytesType})
	registerValueDecoder(bytesPtrType, bytesPtrDecoder{bytesPtrType})
}
