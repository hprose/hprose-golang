/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/encoder/value_encode_test.go                 |
|                                                          |
| LastModified: Feb 22, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoder

import (
	"math"
	"math/big"
	"strings"
	"testing"
	"time"
)

func TestWriteInt(t *testing.T) {
	sb := new(strings.Builder)
	if err := WriteInt(sb, -1); err != nil {
		t.Error(err)
	}
	if err := WriteInt(sb, 0); err != nil {
		t.Error(err)
	}
	if err := WriteInt(sb, 1); err != nil {
		t.Error(err)
	}
	if err := WriteInt(sb, 123); err != nil {
		t.Error(err)
	}
	if err := WriteInt(sb, math.MinInt64); err != nil {
		t.Error(err)
	}
	if err := WriteInt(sb, -math.MaxInt64); err != nil {
		t.Error(err)
	}
	if err := WriteInt(sb, math.MaxInt64); err != nil {
		t.Error(err)
	}
	if sb.String() != "i-1;01i123;l-9223372036854775808;l-9223372036854775807;l9223372036854775807;" {
		t.Error(sb)
	}
}

func TestWriteUint(t *testing.T) {
	sb := new(strings.Builder)
	if err := WriteUint(sb, 0); err != nil {
		t.Error(err)
	}
	if err := WriteUint(sb, 1); err != nil {
		t.Error(err)
	}
	if err := WriteUint(sb, 123); err != nil {
		t.Error(err)
	}
	if err := WriteUint(sb, math.MaxUint64); err != nil {
		t.Error(err)
	}
	if err := WriteUint(sb, math.MaxUint32); err != nil {
		t.Error(err)
	}
	if err := WriteUint(sb, math.MaxInt32); err != nil {
		t.Error(err)
	}
	if sb.String() != "01i123;l18446744073709551615;l4294967295;i2147483647;" {
		t.Error(sb)
	}
}

func TestWriteInt32(t *testing.T) {
	sb := &strings.Builder{}
	if err := WriteInt32(sb, -1); err != nil {
		t.Error(err)
	}
	if err := WriteInt32(sb, 0); err != nil {
		t.Error(err)
	}
	if err := WriteInt32(sb, 1); err != nil {
		t.Error(err)
	}
	if err := WriteInt32(sb, 123); err != nil {
		t.Error(err)
	}
	if err := WriteInt32(sb, math.MinInt32); err != nil {
		t.Error(err)
	}
	if err := WriteInt32(sb, math.MaxInt32); err != nil {
		t.Error(err)
	}
	if sb.String() != "i-1;01i123;i-2147483648;i2147483647;" {
		t.Error(sb)
	}
}

func TestWriteUint32(t *testing.T) {
	sb := &strings.Builder{}
	if err := WriteUint32(sb, 0); err != nil {
		t.Error(err)
	}
	if err := WriteUint32(sb, 1); err != nil {
		t.Error(err)
	}
	if err := WriteUint32(sb, 123); err != nil {
		t.Error(err)
	}
	if err := WriteUint32(sb, math.MaxUint32); err != nil {
		t.Error(err)
	}
	if err := WriteUint32(sb, math.MaxInt32); err != nil {
		t.Error(err)
	}
	if sb.String() != "01i123;l4294967295;i2147483647;" {
		t.Error(sb)
	}
}

func TestWriteUint16(t *testing.T) {
	sb := &strings.Builder{}
	if err := WriteUint16(sb, 0); err != nil {
		t.Error(err)
	}
	if err := WriteUint16(sb, 1); err != nil {
		t.Error(err)
	}
	if err := WriteUint16(sb, 123); err != nil {
		t.Error(err)
	}
	if err := WriteUint16(sb, math.MaxUint16); err != nil {
		t.Error(err)
	}
	if err := WriteUint16(sb, math.MaxInt16); err != nil {
		t.Error(err)
	}
	if sb.String() != "01i123;i65535;i32767;" {
		t.Error(sb)
	}
}

func TestWriteInt16(t *testing.T) {
	sb := &strings.Builder{}
	if err := WriteInt16(sb, 0); err != nil {
		t.Error(err)
	}
	if err := WriteInt16(sb, 1); err != nil {
		t.Error(err)
	}
	if err := WriteInt16(sb, 123); err != nil {
		t.Error(err)
	}
	if err := WriteInt16(sb, math.MinInt16); err != nil {
		t.Error(err)
	}
	if err := WriteInt16(sb, math.MaxInt16); err != nil {
		t.Error(err)
	}
	if sb.String() != "01i123;i-32768;i32767;" {
		t.Error(sb)
	}
}

func TestWriteUint8(t *testing.T) {
	sb := &strings.Builder{}
	if err := WriteUint8(sb, 0); err != nil {
		t.Error(err)
	}
	if err := WriteUint8(sb, 1); err != nil {
		t.Error(err)
	}
	if err := WriteUint8(sb, 123); err != nil {
		t.Error(err)
	}
	if err := WriteUint8(sb, math.MaxUint8); err != nil {
		t.Error(err)
	}
	if err := WriteUint8(sb, math.MaxInt8); err != nil {
		t.Error(err)
	}
	if sb.String() != "01i123;i255;i127;" {
		t.Error(sb)
	}
}

func TestWriteInt8(t *testing.T) {
	sb := &strings.Builder{}
	if err := WriteInt8(sb, 0); err != nil {
		t.Error(err)
	}
	if err := WriteInt8(sb, 1); err != nil {
		t.Error(err)
	}
	if err := WriteInt8(sb, 123); err != nil {
		t.Error(err)
	}
	if err := WriteInt8(sb, math.MinInt8); err != nil {
		t.Error(err)
	}
	if err := WriteInt8(sb, math.MaxInt8); err != nil {
		t.Error(err)
	}
	if sb.String() != "01i123;i-128;i127;" {
		t.Error(sb)
	}
}

func TestWriteNil(t *testing.T) {
	sb := &strings.Builder{}
	if err := WriteNil(sb); err != nil {
		t.Error(err)
	}
	if sb.String() != "n" {
		t.Error(sb)
	}
}

func TestWriteBool(t *testing.T) {
	sb := &strings.Builder{}
	if err := WriteBool(sb, true); err != nil {
		t.Error(err)
	}
	if err := WriteBool(sb, false); err != nil {
		t.Error(err)
	}
	if sb.String() != "tf" {
		t.Error(sb)
	}
}

func TestWriteFloat(t *testing.T) {
	sb := &strings.Builder{}
	if err := WriteFloat32(sb, math.E); err != nil {
		t.Error(err)
	}
	if err := WriteFloat32(sb, math.Pi); err != nil {
		t.Error(err)
	}
	if err := WriteFloat64(sb, math.E); err != nil {
		t.Error(err)
	}
	if err := WriteFloat64(sb, math.Pi); err != nil {
		t.Error(err)
	}
	if err := WriteFloat64(sb, math.Log(1)); err != nil {
		t.Error(err)
	}
	if err := WriteFloat64(sb, math.Log(0)); err != nil {
		t.Error(err)
	}
	if err := WriteFloat64(sb, -math.Log(0)); err != nil {
		t.Error(err)
	}
	if err := WriteFloat64(sb, math.Log(-1)); err != nil {
		t.Error(err)
	}
	if sb.String() != "d2.7182817;d3.1415927;d2.718281828459045;d3.141592653589793;d0;I-I+N" {
		t.Error(sb)
	}
}

func TestWriteBigInt(t *testing.T) {
	sb := &strings.Builder{}
	if err := WriteBigInt(sb, big.NewInt(math.MaxInt64)); err != nil {
		t.Error(err)
	}
	if sb.String() != "l9223372036854775807;" {
		t.Error(sb)
	}
}

func TestWriteBigFloat(t *testing.T) {
	sb := &strings.Builder{}
	if err := WriteBigFloat(sb, big.NewFloat(math.MaxFloat64)); err != nil {
		t.Error(err)
	}
	if sb.String() != "d1.7976931348623157e+308;" {
		t.Error(sb)
	}
}

func TestWriteTime(t *testing.T) {
	sb := &strings.Builder{}
	if err := WriteTime(sb, time.Date(2020, 2, 22, 0, 0, 0, 0, time.UTC)); err != nil {
		t.Error(err)
	}
	if err := WriteTime(sb, time.Date(1970, 1, 1, 12, 12, 12, 0, time.UTC)); err != nil {
		t.Error(err)
	}
	if err := WriteTime(sb, time.Date(1970, 1, 1, 12, 12, 12, 123456789, time.Local)); err != nil {
		t.Error(err)
	}
	if err := WriteTime(sb, time.Date(2020, 2, 22, 12, 12, 12, 123456000, time.Local)); err != nil {
		t.Error(err)
	}
	if err := WriteTime(sb, time.Date(2020, 2, 22, 12, 12, 12, 123000000, time.UTC)); err != nil {
		t.Error(err)
	}
	if sb.String() != "D20200222ZT121212ZT121212.123456789;D20200222T121212.123456;D20200222T121212.123Z" {
		t.Error(sb)
	}
}
