/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/time_decoder_test.go                                  |
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

func init() {
	RegisterValueDecoder(Time{}, GetValueDecoder(time.Time{}))
}

func TestDecodeTime(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	t1, _ := time.Parse("2006-01-02 15:04:05", "2021-06-05 18:57:32")
	assert.NoError(t, enc.Encode(&t1))
	assert.NoError(t, enc.Encode(t1))
	assert.NoError(t, enc.Encode(&t1))
	assert.NoError(t, enc.Encode(t1))

	t2 := Time(t1)
	assert.NoError(t, enc.Encode(&t2))
	assert.NoError(t, enc.Encode(t2))
	assert.NoError(t, enc.Encode(&t2))
	assert.NoError(t, enc.Encode(t2))

	dec := NewDecoder(([]byte)(sb.String())).Simple(false)
	var t3 *time.Time
	var t4 time.Time
	var t5 *Time
	var t6 Time
	dec.Decode(&t3)
	assert.Equal(t, *t3, t1)
	dec.Decode(&t4)
	assert.Equal(t, t4, t1)
	dec.Decode(&t3)
	assert.Equal(t, *t3, t1)
	dec.Decode(&t4)
	assert.Equal(t, t4, t1)
	dec.Decode(&t5)
	assert.Equal(t, *t5, t2)
	dec.Decode(&t6)
	assert.Equal(t, t6, t2)
	dec.Decode(&t5)
	assert.Equal(t, *t5, t2)
	dec.Decode(&t6)
	assert.Equal(t, t6, t2)
}
