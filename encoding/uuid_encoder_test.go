/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/uuid_encoder_test.go                            |
|                                                          |
| LastModified: Apr 27, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding_test

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	. "github.com/hprose/hprose-golang/v3/encoding"
	"github.com/stretchr/testify/assert"
)

func TestEncodeUUID(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	id := uuid.UUID{
		0x7d, 0x44, 0x48, 0x40,
		0x9d, 0xc0,
		0x11, 0xd1,
		0xb2, 0x45,
		0x5f, 0xfd, 0xce, 0x74, 0xfa, 0xd2,
	}
	assert.NoError(t, enc.Encode(uuid.Nil))
	assert.NoError(t, enc.Encode(id))
	assert.NoError(t, enc.Encode(&id))
	assert.NoError(t, enc.Encode(id))
	assert.NoError(t, enc.Encode(&id))
	assert.NoError(t, enc.Encode((*uuid.UUID)(nil)))
	assert.Equal(t, `g{00000000-0000-0000-0000-000000000000}`+
		`g{7d444840-9dc0-11d1-b245-5ffdce74fad2}`+
		`g{7d444840-9dc0-11d1-b245-5ffdce74fad2}`+
		`g{7d444840-9dc0-11d1-b245-5ffdce74fad2}r2;n`, sb.String())
}
