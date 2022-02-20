/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/bytes_decoder.go                                      |
|                                                          |
| LastModified: Feb 20, 2022                               |
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
	if !dec.IsSimple() {
		dec.refer.Add(bytes)
	}
	return bytes
}

func (dec *Decoder) readUint8Slice(et reflect.Type) []byte {
	count := dec.ReadInt()
	slice := make([]byte, count)
	dec.AddReference(slice)
	for i := 0; i < count; i++ {
		dec.decodeUint8(et, dec.NextByte(), &slice[i])
	}
	dec.Skip()
	return slice
}

func (dec *Decoder) decodeBytes(t reflect.Type, tag byte, p *[]byte) {
	switch tag {
	case TagNull:
		*p = nil
	case TagEmpty:
		*p = []byte{}
	case TagBytes:
		*p = dec.ReadBytes()
	case TagList:
		*p = dec.readUint8Slice(t.Elem())
	case TagUTF8Char:
		*p = dec.readStringAsSafeBytes(1)
	case TagString:
		if dec.IsSimple() {
			*p = dec.ReadStringAsBytes()
		} else {
			*p = reflect2.UnsafeCastString(dec.ReadString())
		}
	case TagGUID:
		*p, _ = dec.ReadUUID().MarshalBinary()
	default:
		dec.defaultDecode(t, p, tag)
	}
}

func (dec *Decoder) decodeBytesPtr(t reflect.Type, tag byte, p **[]byte) {
	if tag == TagNull {
		*p = nil
		return
	}
	var bytes []byte
	dec.decodeBytes(t, tag, &bytes)
	*p = &bytes
}

// bytesDecoder is the implementation of ValueDecoder for []byte.
type bytesDecoder struct {
	t reflect.Type
}

func (valdec bytesDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeBytes(valdec.t, tag, (*[]byte)(reflect2.PtrOf(p)))
}

// bytesPtrDecoder is the implementation of ValueDecoder for *[]byte.
type bytesPtrDecoder struct {
	t reflect.Type
}

func (valdec bytesPtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeBytesPtr(valdec.t, tag, (**[]byte)(reflect2.PtrOf(p)))
}

func init() {
	registerValueDecoder(bytesType, bytesDecoder{bytesType})
	registerValueDecoder(bytesPtrType, bytesPtrDecoder{bytesPtrType})
}
