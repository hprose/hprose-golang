/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/int_decoder_test.go                             |
|                                                          |
| LastModified: May 23, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeBool(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
	enc.Encode(-1)
	enc.Encode(0)
	enc.Encode(1.0)
	enc.Encode(123)
	enc.Encode(math.MinInt64)
	enc.Encode(-math.MaxInt64)
	enc.Encode(uint64(math.MaxInt64))
	enc.Encode(true)
	enc.Encode(false)
	enc.Encode(nil)
	enc.Encode(3.14)
	enc.Encode(math.Inf(1))
	enc.Encode(math.NaN())
	enc.Encode("")
	enc.Encode("1")
	enc.Encode("123")
	enc.Encode("N")
	enc.Encode("NaN")
	enc.Encode("F")
	enc.Encode("T")
	enc.Encode("f")
	enc.Encode("t")
	enc.Encode("False")
	enc.Encode("True")
	enc.Encode(nil)
	enc.Encode(true)
	dec := NewDecoder(([]byte)(sb.String()))
	var b bool
	dec.Decode(&b)
	assert.Equal(t, true, b) // -1
	dec.Decode(&b)
	assert.Equal(t, false, b) // 0
	dec.Decode(&b)
	assert.Equal(t, true, b) // 1.0
	dec.Decode(&b)
	assert.Equal(t, true, b) // 123
	dec.Decode(&b)
	assert.Equal(t, true, b) // math.MinInt64
	dec.Decode(&b)
	assert.Equal(t, true, b) // -math.MaxInt64
	dec.Decode(&b)
	assert.Equal(t, true, b) // uint64(math.MaxInt64)
	dec.Decode(&b)
	assert.Equal(t, true, b) // true
	dec.Decode(&b)
	assert.Equal(t, false, b) // false
	dec.Decode(&b)
	assert.Equal(t, false, b) // nil
	dec.Decode(&b)
	assert.Equal(t, true, b) // 3.14
	dec.Decode(&b)
	assert.Equal(t, true, b) // math.Inf(1)
	dec.Decode(&b)
	assert.Equal(t, true, b) // math.NaN()
	dec.Decode(&b)
	assert.Equal(t, false, b) // ""
	dec.Decode(&b)
	assert.Equal(t, true, b) // "1"
	dec.Decode(&b)
	assert.EqualError(t, dec.Error, `strconv.ParseBool: parsing "123": invalid syntax`) // "123"
	dec.Error = nil
	dec.Decode(&b)
	assert.EqualError(t, dec.Error, `strconv.ParseBool: parsing "N": invalid syntax`) // "N"
	dec.Error = nil
	dec.Decode(&b)
	assert.EqualError(t, dec.Error, `strconv.ParseBool: parsing "NaN": invalid syntax`) // "NaN"
	dec.Error = nil
	dec.Decode(&b)
	assert.Equal(t, false, b) // "F"
	dec.Decode(&b)
	assert.Equal(t, true, b) // "T"
	dec.Decode(&b)
	assert.Equal(t, false, b) // "f"
	dec.Decode(&b)
	assert.Equal(t, true, b) // "t"
	dec.Decode(&b)
	assert.Equal(t, false, b) // "False"
	dec.Decode(&b)
	assert.Equal(t, true, b) // "True"
	var bp *bool
	dec.Decode(&bp)
	assert.Equal(t, (*bool)(nil), bp) // nil
	dec.Decode(&bp)
	assert.Equal(t, true, *bp) // true
}
