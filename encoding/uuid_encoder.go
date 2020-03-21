/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/uuid_encoder.go                                 |
|                                                          |
| LastModified: Mar 21, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"encoding/hex"

	"github.com/google/uuid"
	"github.com/modern-go/reflect2"
)

// uuidEncoder is the implementation of ValueEncoder for uuid.UUID/*uuid.UUID.
type uuidEncoder struct{}

func (valenc uuidEncoder) Encode(enc *Encoder, v interface{}) {
	EncodeReference(valenc, enc, v)
}

func (uuidEncoder) Write(enc *Encoder, v interface{}) {
	SetReference(enc, v)
	writeUUID(enc, *(*[16]byte)(reflect2.PtrOf(v)))
}

func writeUUID(enc *Encoder, id [16]byte) {
	var buf [36]byte
	encodeHex(buf[:], id)
	enc.buf = append(enc.buf, TagGUID, TagOpenbrace)
	enc.buf = append(enc.buf, buf[:]...)
	enc.buf = append(enc.buf, TagClosebrace)
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
	RegisterValueEncoder((*uuid.UUID)(nil), uuidEncoder{})
}
