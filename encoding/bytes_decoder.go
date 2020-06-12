/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/bytes_decoder.go                                |
|                                                          |
| LastModified: Jun 12, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"reflect"

	"github.com/modern-go/reflect2"
)

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

func (dec *Decoder) decodeBytes(t reflect.Type, tag byte) []byte {
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
	default:
		dec.decodeError(t, tag)
	}
	return nil
}

// bytesDecoder is the implementation of ValueDecoder for []byte.
type bytesDecoder struct {
	t reflect.Type
}

func (valdec bytesDecoder) decode(dec *Decoder, pv *[]byte, tag byte) {
	if bytes := dec.decodeBytes(valdec.t, tag); dec.Error == nil {
		*pv = bytes
	}
}

func (valdec bytesDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (*[]byte)(reflect2.PtrOf(p)), tag)
}

func (valdec bytesDecoder) Type() reflect.Type {
	return valdec.t
}

// bytesPtrDecoder is the implementation of ValueDecoder for *[]byte.
type bytesPtrDecoder struct {
	t reflect.Type
}

func (valdec bytesPtrDecoder) decode(dec *Decoder, pv **[]byte, tag byte) {
	if bytes := dec.decodeBytes(valdec.t, tag); dec.Error == nil {
		*pv = &bytes
	}
}

func (valdec bytesPtrDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	valdec.decode(dec, (**[]byte)(reflect2.PtrOf(p)), tag)
}

func (valdec bytesPtrDecoder) Type() reflect.Type {
	return valdec.t
}

var (
	u8sdec  = bytesDecoder{reflect.TypeOf(([]byte)(nil))}
	pu8sdec = bytesPtrDecoder{reflect.TypeOf((*[]byte)(nil))}
)

func init() {
	RegisterValueDecoder(u8sdec)
	RegisterValueDecoder(pu8sdec)
}
