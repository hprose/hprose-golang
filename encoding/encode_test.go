/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/value_encode_test.go                            |
|                                                          |
| LastModified: Mar 20, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"math"
	"math/big"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteInt(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
	assert.NoError(t, WriteInt(enc, -1))
	assert.NoError(t, WriteInt(enc, 0))
	assert.NoError(t, WriteInt(enc, 1))
	assert.NoError(t, WriteInt(enc, 123))
	assert.NoError(t, WriteInt(enc, math.MinInt64))
	assert.NoError(t, WriteInt(enc, -math.MaxInt64))
	assert.NoError(t, WriteInt(enc, math.MaxInt64))
	assert.Equal(t, "i-1;01i123;l-9223372036854775808;l-9223372036854775807;l9223372036854775807;", sb.String())
}

func TestWriteUint(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
	assert.NoError(t, WriteUint(enc, 0))
	assert.NoError(t, WriteUint(enc, 1))
	assert.NoError(t, WriteUint(enc, 123))
	assert.NoError(t, WriteUint(enc, math.MaxUint64))
	assert.NoError(t, WriteUint(enc, math.MaxUint32))
	assert.NoError(t, WriteUint(enc, math.MaxInt32))
	assert.Equal(t, "01i123;l18446744073709551615;l4294967295;i2147483647;", sb.String())
}

func TestWriteInt32(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, true)
	assert.NoError(t, WriteInt32(enc, -1))
	assert.NoError(t, WriteInt32(enc, 0))
	assert.NoError(t, WriteInt32(enc, 1))
	assert.NoError(t, WriteInt32(enc, 123))
	assert.NoError(t, WriteInt32(enc, math.MinInt32))
	assert.NoError(t, WriteInt32(enc, math.MaxInt32))
	assert.Equal(t, "i-1;01i123;i-2147483648;i2147483647;", sb.String())
}

func TestWriteUint32(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, true)
	assert.NoError(t, WriteUint32(enc, 0))
	assert.NoError(t, WriteUint32(enc, 1))
	assert.NoError(t, WriteUint32(enc, 123))
	assert.NoError(t, WriteUint32(enc, math.MaxUint32))
	assert.NoError(t, WriteUint32(enc, math.MaxInt32))
	assert.Equal(t, "01i123;l4294967295;i2147483647;", sb.String())
}

func TestWriteUint16(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, true)
	assert.NoError(t, WriteUint16(enc, 0))
	assert.NoError(t, WriteUint16(enc, 1))
	assert.NoError(t, WriteUint16(enc, 123))
	assert.NoError(t, WriteUint16(enc, math.MaxUint16))
	assert.NoError(t, WriteUint16(enc, math.MaxInt16))
	assert.Equal(t, "01i123;i65535;i32767;", sb.String())
}

func TestWriteInt16(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, true)
	assert.NoError(t, WriteInt16(enc, 0))
	assert.NoError(t, WriteInt16(enc, 1))
	assert.NoError(t, WriteInt16(enc, 123))
	assert.NoError(t, WriteInt16(enc, math.MinInt16))
	assert.NoError(t, WriteInt16(enc, math.MaxInt16))
	assert.Equal(t, "01i123;i-32768;i32767;", sb.String())
}

func TestWriteUint8(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, true)
	assert.NoError(t, WriteUint8(enc, 0))
	assert.NoError(t, WriteUint8(enc, 1))
	assert.NoError(t, WriteUint8(enc, 123))
	assert.NoError(t, WriteUint8(enc, math.MaxUint8))
	assert.NoError(t, WriteUint8(enc, math.MaxInt8))
	assert.Equal(t, "01i123;i255;i127;", sb.String())
}

func TestWriteInt8(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, true)
	assert.NoError(t, WriteInt8(enc, 0))
	assert.NoError(t, WriteInt8(enc, 1))
	assert.NoError(t, WriteInt8(enc, 123))
	assert.NoError(t, WriteInt8(enc, math.MinInt8))
	assert.NoError(t, WriteInt8(enc, math.MaxInt8))
	assert.Equal(t, "01i123;i-128;i127;", sb.String())
}

func TestWriteBool(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, true)
	assert.NoError(t, WriteBool(enc, true))
	assert.NoError(t, WriteBool(enc, false))
	assert.Equal(t, "tf", sb.String())
}

func TestWriteFloat(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, true)
	assert.NoError(t, WriteFloat32(enc, math.E))
	assert.NoError(t, WriteFloat32(enc, math.Pi))
	assert.NoError(t, WriteFloat64(enc, math.E))
	assert.NoError(t, WriteFloat64(enc, math.Pi))
	assert.NoError(t, WriteFloat64(enc, math.Log(1)))
	assert.NoError(t, WriteFloat64(enc, math.Log(0)))
	assert.NoError(t, WriteFloat64(enc, -math.Log(0)))
	assert.NoError(t, WriteFloat64(enc, math.Log(-1)))
	assert.Equal(t, "d2.7182817;d3.1415927;d2.718281828459045;d3.141592653589793;d0;I-I+N", sb.String())
}

func TestWriteBigInt(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, true)
	assert.NoError(t, WriteBigInt(enc, big.NewInt(math.MaxInt64)))
	assert.Equal(t, "l9223372036854775807;", sb.String())
}

func TestWriteBigFloat(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb, true)
	assert.NoError(t, WriteBigFloat(enc, big.NewFloat(math.MaxFloat64)))
	assert.Equal(t, "d1.7976931348623157e+308;", sb.String())
}
