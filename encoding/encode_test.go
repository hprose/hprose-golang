/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/value_encode_test.go                            |
|                                                          |
| LastModified: Mar 22, 2020                               |
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
	WriteInt(enc, -1)
	WriteInt(enc, 0)
	WriteInt(enc, 1)
	WriteInt(enc, 123)
	WriteInt(enc, math.MinInt64)
	WriteInt(enc, -math.MaxInt64)
	WriteInt(enc, math.MaxInt64)
	assert.NoError(t, enc.Flush())
	assert.Equal(t, "i-1;01i123;l-9223372036854775808;l-9223372036854775807;l9223372036854775807;", sb.String())
}

func TestWriteUint(t *testing.T) {
	enc := new(Encoder)
	WriteUint(enc, 0)
	WriteUint(enc, 1)
	WriteUint(enc, 123)
	WriteUint(enc, math.MaxUint64)
	WriteUint(enc, math.MaxUint32)
	WriteUint(enc, math.MaxInt32)
	assert.Equal(t, "01i123;l18446744073709551615;l4294967295;i2147483647;", enc.String())
}

func TestWriteInt32(t *testing.T) {
	enc := new(Encoder)
	WriteInt32(enc, -1)
	WriteInt32(enc, 0)
	WriteInt32(enc, 1)
	WriteInt32(enc, 123)
	WriteInt32(enc, math.MinInt32)
	WriteInt32(enc, math.MaxInt32)
	assert.Equal(t, "i-1;01i123;i-2147483648;i2147483647;", enc.String())
}

func TestWriteUint32(t *testing.T) {
	enc := new(Encoder)
	WriteUint32(enc, 0)
	WriteUint32(enc, 1)
	WriteUint32(enc, 123)
	WriteUint32(enc, math.MaxUint32)
	WriteUint32(enc, math.MaxInt32)
	assert.Equal(t, "01i123;l4294967295;i2147483647;", enc.String())
}

func TestWriteUint16(t *testing.T) {
	enc := new(Encoder)
	WriteUint16(enc, 0)
	WriteUint16(enc, 1)
	WriteUint16(enc, 123)
	WriteUint16(enc, math.MaxUint16)
	WriteUint16(enc, math.MaxInt16)
	assert.Equal(t, "01i123;i65535;i32767;", enc.String())
}

func TestWriteInt16(t *testing.T) {
	enc := new(Encoder)
	WriteInt16(enc, 0)
	WriteInt16(enc, 1)
	WriteInt16(enc, 123)
	WriteInt16(enc, math.MinInt16)
	WriteInt16(enc, math.MaxInt16)
	assert.Equal(t, "01i123;i-32768;i32767;", enc.String())
}

func TestWriteUint8(t *testing.T) {
	enc := new(Encoder)
	WriteUint8(enc, 0)
	WriteUint8(enc, 1)
	WriteUint8(enc, 123)
	WriteUint8(enc, math.MaxUint8)
	WriteUint8(enc, math.MaxInt8)
	assert.Equal(t, "01i123;i255;i127;", enc.String())
}

func TestWriteInt8(t *testing.T) {
	enc := new(Encoder)
	WriteInt8(enc, 0)
	WriteInt8(enc, 1)
	WriteInt8(enc, 123)
	WriteInt8(enc, math.MinInt8)
	WriteInt8(enc, math.MaxInt8)
	assert.Equal(t, "01i123;i-128;i127;", enc.String())
}

func TestWriteBool(t *testing.T) {
	enc := new(Encoder)
	WriteBool(enc, true)
	WriteBool(enc, false)
	assert.Equal(t, "tf", enc.String())
}

func TestWriteFloat(t *testing.T) {
	enc := new(Encoder)
	WriteFloat32(enc, math.E)
	WriteFloat32(enc, math.Pi)
	WriteFloat64(enc, math.E)
	WriteFloat64(enc, math.Pi)
	WriteFloat64(enc, math.Log(1))
	WriteFloat64(enc, math.Log(0))
	WriteFloat64(enc, -math.Log(0))
	WriteFloat64(enc, math.Log(-1))
	assert.Equal(t, "d2.7182817;d3.1415927;d2.718281828459045;d3.141592653589793;d0;I-I+N", enc.String())
}

func TestWriteBigInt(t *testing.T) {
	enc := new(Encoder)
	WriteBigInt(enc, big.NewInt(math.MaxInt64))
	assert.Equal(t, "l9223372036854775807;", enc.String())
}

func TestWriteBigFloat(t *testing.T) {
	enc := new(Encoder)
	enc.WriteBigFloat(big.NewFloat(math.MaxFloat64))
	assert.Equal(t, "d1.7976931348623157e+308;", enc.String())
}
