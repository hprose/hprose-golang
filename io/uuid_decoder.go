/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/uuid_decoder.go                                       |
|                                                          |
| LastModified: Feb 7, 2024                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

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

func (dec *Decoder) decodeUUID(t reflect.Type, tag byte, p *uuid.UUID) {
	switch tag {
	case TagEmpty, TagNull:
		*p = uuid.Nil
	case TagGUID:
		*p = dec.ReadUUID()
	case TagBytes:
		if dec.IsSimple() {
			*p = dec.bytesToUUID(dec.readUnsafeBytes())
		} else {
			*p = dec.bytesToUUID(dec.ReadBytes())
		}
	case TagString:
		if dec.IsSimple() {
			*p = dec.stringToUUID(dec.ReadUnsafeString())
		} else {
			*p = dec.stringToUUID(dec.ReadString())
		}
	default:
		dec.defaultDecode(t, p, tag)
	}
}

func (dec *Decoder) decodeUUIDPtr(t reflect.Type, tag byte, p **uuid.UUID) {
	if tag == TagNull {
		*p = nil
		return
	}
	var u uuid.UUID
	dec.decodeUUID(t, tag, &u)
	*p = &u
}

// uuidDecoder is the implementation of ValueDecoder for uuid.UUID.
type uuidDecoder struct{}

func (uuidDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	dec.decodeUUID(reflect.TypeOf(p).Elem(), tag, (*uuid.UUID)(reflect2.PtrOf(p)))
}

func init() {
	registerValueDecoder(uuidType, uuidDecoder{})
}
