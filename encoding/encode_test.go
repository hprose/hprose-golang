/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/value_encode_test.go                            |
|                                                          |
| LastModified: Apr 12, 2020                               |
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
	enc.WriteInt(-1)
	enc.WriteInt(0)
	enc.WriteInt(1)
	enc.WriteInt(123)
	enc.WriteInt(math.MinInt64)
	enc.WriteInt(-math.MaxInt64)
	enc.WriteInt(math.MaxInt64)
	assert.NoError(t, enc.Flush())
	assert.Equal(t, "i-1;01i123;l-9223372036854775808;l-9223372036854775807;l9223372036854775807;", sb.String())
}

func TestWriteUint(t *testing.T) {
	enc := new(Encoder)
	enc.WriteUint(0)
	enc.WriteUint(1)
	enc.WriteUint(123)
	enc.WriteUint(math.MaxUint64)
	enc.WriteUint(math.MaxUint32)
	enc.WriteUint(math.MaxInt32)
	assert.Equal(t, "01i123;l18446744073709551615;l4294967295;i2147483647;", enc.String())
}

func TestWriteInt32(t *testing.T) {
	enc := new(Encoder)
	enc.WriteInt32(-1)
	enc.WriteInt32(0)
	enc.WriteInt32(1)
	enc.WriteInt32(123)
	enc.WriteInt32(math.MinInt32)
	enc.WriteInt32(math.MaxInt32)
	assert.Equal(t, "i-1;01i123;i-2147483648;i2147483647;", enc.String())
}

func TestWriteUint32(t *testing.T) {
	enc := new(Encoder)
	enc.WriteUint32(0)
	enc.WriteUint32(1)
	enc.WriteUint32(123)
	enc.WriteUint32(math.MaxUint32)
	enc.WriteUint32(math.MaxInt32)
	assert.Equal(t, "01i123;l4294967295;i2147483647;", enc.String())
}

func TestWriteUint16(t *testing.T) {
	enc := new(Encoder)
	enc.WriteUint16(0)
	enc.WriteUint16(1)
	enc.WriteUint16(123)
	enc.WriteUint16(math.MaxUint16)
	enc.WriteUint16(math.MaxInt16)
	assert.Equal(t, "01i123;i65535;i32767;", enc.String())
}

func TestWriteInt16(t *testing.T) {
	enc := new(Encoder)
	enc.WriteInt16(0)
	enc.WriteInt16(1)
	enc.WriteInt16(123)
	enc.WriteInt16(math.MinInt16)
	enc.WriteInt16(math.MaxInt16)
	assert.Equal(t, "01i123;i-32768;i32767;", enc.String())
}

func TestWriteUint8(t *testing.T) {
	enc := new(Encoder)
	enc.WriteUint8(0)
	enc.WriteUint8(1)
	enc.WriteUint8(123)
	enc.WriteUint8(math.MaxUint8)
	enc.WriteUint8(math.MaxInt8)
	assert.Equal(t, "01i123;i255;i127;", enc.String())
}

func TestWriteInt8(t *testing.T) {
	enc := new(Encoder)
	enc.WriteInt8(0)
	enc.WriteInt8(1)
	enc.WriteInt8(123)
	enc.WriteInt8(math.MinInt8)
	enc.WriteInt8(math.MaxInt8)
	assert.Equal(t, "01i123;i-128;i127;", enc.String())
}

func TestWriteBool(t *testing.T) {
	enc := new(Encoder)
	enc.WriteBool(true)
	enc.WriteBool(false)
	assert.Equal(t, "tf", enc.String())
}

func TestWriteFloat(t *testing.T) {
	enc := new(Encoder)
	enc.WriteFloat32(math.E)
	enc.WriteFloat32(math.Pi)
	enc.WriteFloat64(math.E)
	enc.WriteFloat64(math.Pi)
	enc.WriteFloat64(math.Log(1))
	enc.WriteFloat64(math.Log(0))
	enc.WriteFloat64(-math.Log(0))
	enc.WriteFloat64(math.Log(-1))
	assert.Equal(t, "d2.7182817;d3.1415927;d2.718281828459045;d3.141592653589793;d0;I-I+N", enc.String())
}

func TestWriteBigInt(t *testing.T) {
	enc := new(Encoder)
	enc.WriteBigInt(big.NewInt(math.MaxInt64))
	assert.Equal(t, "l9223372036854775807;", enc.String())
}

func TestWriteBigFloat(t *testing.T) {
	enc := new(Encoder)
	enc.WriteBigFloat(big.NewFloat(math.MaxFloat64))
	assert.Equal(t, "d1.7976931348623157e+308;", enc.String())
}
