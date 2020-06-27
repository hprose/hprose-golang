/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/uuid_encoder_test.go                            |
|                                                          |
| LastModified: Mar 21, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDecodeUUID(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	src := uuid.UUID{
		0x7d, 0x44, 0x48, 0x40,
		0x9d, 0xc0,
		0x11, 0xd1,
		0xb2, 0x45,
		0x5f, 0xfd, 0xce, 0x74, 0xfa, 0xd2,
	}
	enc.Encode(uuid.Nil)
	enc.Encode(src)
	enc.Encode("{7d444840-9dc0-11d1-b245-5ffdce74fad2}")
	enc.Encode([]byte("{7d444840-9dc0-11d1-b245-5ffdce74fad2}"))
	enc.Encode(src[:])
	enc.Encode((*uuid.UUID)(nil))
	enc.Encode("")
	dec := NewDecoder(([]byte)(sb.String()))
	var id uuid.UUID
	dec.Decode(&id)
	assert.Equal(t, uuid.Nil, id)
	dec.Decode(&id)
	assert.Equal(t, src, id)
	dec.Decode(&id)
	assert.Equal(t, src, id)
	dec.Decode(&id)
	assert.Equal(t, src, id)
	dec.Decode(&id)
	assert.Equal(t, src, id)
	dec.Decode(&id)
	assert.Equal(t, uuid.Nil, id)
	dec.Decode(&id)
	assert.Equal(t, uuid.Nil, id)
}
