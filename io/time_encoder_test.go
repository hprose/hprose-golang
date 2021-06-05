/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/time_encoder_test.go                                  |
|                                                          |
| LastModified: Jun 05, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io_test

import (
	"strings"
	"testing"
	"time"

	. "github.com/hprose/hprose-golang/v3/io"
	"github.com/stretchr/testify/assert"
)

type Time time.Time

func (t Time) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte("\"" + t.String() + "\""), nil
}

func init() {
	RegisterValueEncoder(Time{}, GetValueEncoder(time.Time{}))
}

func TestEncodeTime(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	t1, _ := time.Parse("2006-01-02 15:04:05", "2021-06-05 18:57:32")
	assert.NoError(t, enc.Encode(&t1))
	assert.NoError(t, enc.Encode(t1))
	assert.NoError(t, enc.Encode(&t1))
	assert.NoError(t, enc.Encode(t1))

	assert.Equal(t, `D20210605T185732ZD20210605T185732Zr0;D20210605T185732Z`, sb.String())

	enc.Reset()
	sb.Reset()

	t2 := Time(t1)
	assert.NoError(t, enc.Encode(&t2))
	assert.NoError(t, enc.Encode(t2))
	assert.NoError(t, enc.Encode(&t2))
	assert.NoError(t, enc.Encode(t2))
	assert.Equal(t, `D20210605T185732ZD20210605T185732Zr0;D20210605T185732Z`, sb.String())
}
