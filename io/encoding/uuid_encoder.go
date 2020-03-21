/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/uuid_encoder.go                              |
|                                                          |
| LastModified: Mar 21, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"encoding/hex"

	"github.com/google/uuid"
	"github.com/hprose/hprose-golang/v3/io"
	"github.com/modern-go/reflect2"
)

// UUIDEncoder is the implementation of ValueEncoder for uuid.UUID/*uuid.UUID.
type UUIDEncoder struct{}

// Encode writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as reference
func (valenc UUIDEncoder) Encode(enc *Encoder, v interface{}) error {
	return ReferenceEncode(valenc, enc, v)
}

// Write writes the hprose encoding of v to stream
// if v is already written to stream, it will writes it as value
func (UUIDEncoder) Write(enc *Encoder, v interface{}) error {
	SetReference(enc, v)
	return WriteUUID(enc.Writer, *(*[16]byte)(reflect2.PtrOf(v)))
}

// WriteUUID to writer
func WriteUUID(writer io.BytesWriter, id [16]byte) (err error) {
	var buf [36]byte
	encodeHex(buf[:], id)
	if err = writer.WriteByte(io.TagGUID); err == nil {
		if err = writer.WriteByte(io.TagOpenbrace); err == nil {
			if _, err = writer.Write(buf[:]); err == nil {
				err = writer.WriteByte(io.TagClosebrace)
			}
		}
	}
	return
}

func encodeHex(dst []byte, uuid [16]byte) {
	hex.Encode(dst, uuid[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], uuid[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], uuid[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], uuid[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], uuid[10:])
}

func init() {
	RegisterEncoder((*uuid.UUID)(nil), UUIDEncoder{})
}
