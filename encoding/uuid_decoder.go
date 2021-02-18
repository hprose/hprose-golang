/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/uuid_decoder.go                                 |
|                                                          |
| LastModified: Feb 18, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"reflect"

	"github.com/google/uuid"
	"github.com/modern-go/reflect2"
)

func (dec *Decoder) stringToUUID(value string) uuid.UUID {
	uuid, err := uuid.Parse(value)
	if err != nil {
		dec.Error = err
	}
	return uuid
}

func (dec *Decoder) bytesToUUID(value []byte) (id uuid.UUID) {
	if len(value) == 16 {
		copy(id[:], value)
		return
	}
	var err error
	id, err = uuid.ParseBytes(value)
	if err != nil {
		dec.Error = err
	}
	return
}

// ReadUUID reads uuid.UUID and add reference.
func (dec *Decoder) ReadUUID() uuid.UUID {
	uuid, err := uuid.ParseBytes(dec.UnsafeNext(38))
	dec.AddReference(uuid)
	if err != nil {
		dec.Error = err
	}
	return uuid
}

func (dec *Decoder) decodeUUID(t reflect.Type, tag byte) (id uuid.UUID) {
	switch tag {
	case TagEmpty, TagNull:
		return uuid.Nil
	case TagGUID:
		return dec.ReadUUID()
	case TagBytes:
		if dec.IsSimple() {
			return dec.bytesToUUID(dec.readUnsafeBytes())
		}
		return dec.bytesToUUID(dec.ReadBytes())
	case TagString:
		if dec.IsSimple() {
			return dec.stringToUUID(dec.ReadUnsafeString())
		}
		return dec.stringToUUID(dec.ReadString())
	default:
		dec.decodeError(t, tag)
	}
	return
}

func (dec *Decoder) decodeUUIDPtr(t reflect.Type, tag byte) *uuid.UUID {
	if tag == TagNull {
		return nil
	}
	uuid := dec.decodeUUID(t, tag)
	return &uuid
}

// uuidDecoder is the implementation of ValueDecoder for uuid.UUID.
type uuidDecoder struct{}

func (uuidDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	*(*uuid.UUID)(reflect2.PtrOf(p)) = dec.decodeUUID(uuidType, tag)
}

func (uuidDecoder) Type() reflect.Type {
	return uuidType
}

func init() {
	RegisterValueDecoder(uuidDecoder{})
}
