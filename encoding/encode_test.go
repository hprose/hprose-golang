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
	assert.NoError(t, WriteInt(sb, -1))
	assert.NoError(t, WriteInt(sb, 0))
	assert.NoError(t, WriteInt(sb, 1))
	assert.NoError(t, WriteInt(sb, 123))
	assert.NoError(t, WriteInt(sb, math.MinInt64))
	assert.NoError(t, WriteInt(sb, -math.MaxInt64))
	assert.NoError(t, WriteInt(sb, math.MaxInt64))
	assert.Equal(t, "i-1;01i123;l-9223372036854775808;l-9223372036854775807;l9223372036854775807;", sb.String())
}

func TestWriteUint(t *testing.T) {
	sb := new(strings.Builder)
	assert.NoError(t, WriteUint(sb, 0))
	assert.NoError(t, WriteUint(sb, 1))
	assert.NoError(t, WriteUint(sb, 123))
	assert.NoError(t, WriteUint(sb, math.MaxUint64))
	assert.NoError(t, WriteUint(sb, math.MaxUint32))
	assert.NoError(t, WriteUint(sb, math.MaxInt32))
	assert.Equal(t, "01i123;l18446744073709551615;l4294967295;i2147483647;", sb.String())
}

func TestWriteInt32(t *testing.T) {
	sb := &strings.Builder{}
	assert.NoError(t, WriteInt32(sb, -1))
	assert.NoError(t, WriteInt32(sb, 0))
	assert.NoError(t, WriteInt32(sb, 1))
	assert.NoError(t, WriteInt32(sb, 123))
	assert.NoError(t, WriteInt32(sb, math.MinInt32))
	assert.NoError(t, WriteInt32(sb, math.MaxInt32))
	assert.Equal(t, "i-1;01i123;i-2147483648;i2147483647;", sb.String())
}

func TestWriteUint32(t *testing.T) {
	sb := &strings.Builder{}
	assert.NoError(t, WriteUint32(sb, 0))
	assert.NoError(t, WriteUint32(sb, 1))
	assert.NoError(t, WriteUint32(sb, 123))
	assert.NoError(t, WriteUint32(sb, math.MaxUint32))
	assert.NoError(t, WriteUint32(sb, math.MaxInt32))
	assert.Equal(t, "01i123;l4294967295;i2147483647;", sb.String())
}

func TestWriteUint16(t *testing.T) {
	sb := &strings.Builder{}
	assert.NoError(t, WriteUint16(sb, 0))
	assert.NoError(t, WriteUint16(sb, 1))
	assert.NoError(t, WriteUint16(sb, 123))
	assert.NoError(t, WriteUint16(sb, math.MaxUint16))
	assert.NoError(t, WriteUint16(sb, math.MaxInt16))
	assert.Equal(t, "01i123;i65535;i32767;", sb.String())
}

func TestWriteInt16(t *testing.T) {
	sb := &strings.Builder{}
	assert.NoError(t, WriteInt16(sb, 0))
	assert.NoError(t, WriteInt16(sb, 1))
	assert.NoError(t, WriteInt16(sb, 123))
	assert.NoError(t, WriteInt16(sb, math.MinInt16))
	assert.NoError(t, WriteInt16(sb, math.MaxInt16))
	assert.Equal(t, "01i123;i-32768;i32767;", sb.String())
}

func TestWriteUint8(t *testing.T) {
	sb := &strings.Builder{}
	assert.NoError(t, WriteUint8(sb, 0))
	assert.NoError(t, WriteUint8(sb, 1))
	assert.NoError(t, WriteUint8(sb, 123))
	assert.NoError(t, WriteUint8(sb, math.MaxUint8))
	assert.NoError(t, WriteUint8(sb, math.MaxInt8))
	assert.Equal(t, "01i123;i255;i127;", sb.String())
}

func TestWriteInt8(t *testing.T) {
	sb := &strings.Builder{}
	assert.NoError(t, WriteInt8(sb, 0))
	assert.NoError(t, WriteInt8(sb, 1))
	assert.NoError(t, WriteInt8(sb, 123))
	assert.NoError(t, WriteInt8(sb, math.MinInt8))
	assert.NoError(t, WriteInt8(sb, math.MaxInt8))
	assert.Equal(t, "01i123;i-128;i127;", sb.String())
}

func TestWriteBool(t *testing.T) {
	sb := &strings.Builder{}
	assert.NoError(t, WriteBool(sb, true))
	assert.NoError(t, WriteBool(sb, false))
	assert.Equal(t, "tf", sb.String())
}

func TestWriteFloat(t *testing.T) {
	sb := &strings.Builder{}
	assert.NoError(t, WriteFloat32(sb, math.E))
	assert.NoError(t, WriteFloat32(sb, math.Pi))
	assert.NoError(t, WriteFloat64(sb, math.E))
	assert.NoError(t, WriteFloat64(sb, math.Pi))
	assert.NoError(t, WriteFloat64(sb, math.Log(1)))
	assert.NoError(t, WriteFloat64(sb, math.Log(0)))
	assert.NoError(t, WriteFloat64(sb, -math.Log(0)))
	assert.NoError(t, WriteFloat64(sb, math.Log(-1)))
	assert.Equal(t, "d2.7182817;d3.1415927;d2.718281828459045;d3.141592653589793;d0;I-I+N", sb.String())
}

func TestWriteBigInt(t *testing.T) {
	sb := &strings.Builder{}
	assert.NoError(t, WriteBigInt(sb, big.NewInt(math.MaxInt64)))
	assert.Equal(t, "l9223372036854775807;", sb.String())
}

func TestWriteBigFloat(t *testing.T) {
	sb := &strings.Builder{}
	assert.NoError(t, WriteBigFloat(sb, big.NewFloat(math.MaxFloat64)))
	assert.Equal(t, "d1.7976931348623157e+308;", sb.String())
}
