/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/decoder_test.go                                 |
|                                                          |
| LastModified: Apr 12, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadIntFromBytes(t *testing.T) {
	dec := NewDecoder(([]byte)(";1;12;123;1234;12345;123456;1234567;12345678;123456789;1234567890;-2147483648;2147483647;-9223372036854775808;18446744073709551615;-2147483648;2147483647;-9223372036854775808;18446744073709551615;-2147483648;2147483647;-9223372036854775808;18446744073709551615;-2147483648;2147483647;-9223372036854775808;18446744073709551615;;1;12;123;1234;12345;123456;1234567;12345678;123456789;1234567890"))
	i32 := int32(-2147483648)
	u32 := uint32(2147483647)
	i64 := int64(-9223372036854775808)
	u64 := uint64(18446744073709551615)
	assert.Equal(t, 0, dec.ReadInt())
	assert.Equal(t, 1, dec.ReadInt())
	assert.Equal(t, 12, dec.ReadInt())
	assert.Equal(t, 123, dec.ReadInt())
	assert.Equal(t, 1234, dec.ReadInt())
	assert.Equal(t, 12345, dec.ReadInt())
	assert.Equal(t, 123456, dec.ReadInt())
	assert.Equal(t, 1234567, dec.ReadInt())
	assert.Equal(t, 12345678, dec.ReadInt())
	assert.Equal(t, 123456789, dec.ReadInt())
	assert.Equal(t, 1234567890, dec.ReadInt())
	assert.Equal(t, int8(i32), dec.ReadInt8())
	assert.Equal(t, uint8(u32), dec.ReadUint8())
	assert.Equal(t, int16(i64), dec.ReadInt16())
	assert.Equal(t, uint16(u64), dec.ReadUint16())
	assert.Equal(t, i32, dec.ReadInt32())
	assert.Equal(t, u32, dec.ReadUint32())
	assert.Equal(t, i64, dec.ReadInt64())
	assert.Equal(t, u64, dec.ReadUint64())
	assert.Equal(t, int(i32), dec.ReadInt())
	assert.Equal(t, int(u32), dec.ReadInt())
	assert.Equal(t, int(i64), dec.ReadInt())
	assert.Equal(t, int(u64), dec.ReadInt())
	assert.Equal(t, uint(i32), dec.ReadUint())
	assert.Equal(t, uint(u32), dec.ReadUint())
	assert.Equal(t, uint(i64), dec.ReadUint())
	assert.Equal(t, uint(u64), dec.ReadUint())
	assert.Equal(t, int32(0), dec.ReadInt32())
	assert.Equal(t, int32(1), dec.ReadInt32())
	assert.Equal(t, int32(12), dec.ReadInt32())
	assert.Equal(t, int32(123), dec.ReadInt32())
	assert.Equal(t, int32(1234), dec.ReadInt32())
	assert.Equal(t, int32(12345), dec.ReadInt32())
	assert.Equal(t, int32(123456), dec.ReadInt32())
	assert.Equal(t, int32(1234567), dec.ReadInt32())
	assert.Equal(t, int32(12345678), dec.ReadInt32())
	assert.Equal(t, int32(123456789), dec.ReadInt32())
	assert.Equal(t, int32(1234567890), dec.ReadInt32())
}

func TestReadIntFromReader(t *testing.T) {
	reader := bytes.NewBufferString(";1;12;123;1234;12345;123456;1234567;12345678;123456789;1234567890;-2147483648;2147483647;-9223372036854775808;18446744073709551615;-2147483648;2147483647;-9223372036854775808;18446744073709551615;-2147483648;2147483647;-9223372036854775808;18446744073709551615;-2147483648;2147483647;-9223372036854775808;18446744073709551615;;1;12;123;1234;12345;123456;1234567;12345678;123456789;1234567890")
	dec := NewDecoderFromReader(reader, 0)
	i32 := int32(-2147483648)
	u32 := uint32(2147483647)
	i64 := int64(-9223372036854775808)
	u64 := uint64(18446744073709551615)
	assert.Equal(t, 0, dec.ReadInt())
	assert.Equal(t, 1, dec.ReadInt())
	assert.Equal(t, 12, dec.ReadInt())
	assert.Equal(t, 123, dec.ReadInt())
	assert.Equal(t, 1234, dec.ReadInt())
	assert.Equal(t, 12345, dec.ReadInt())
	assert.Equal(t, 123456, dec.ReadInt())
	assert.Equal(t, 1234567, dec.ReadInt())
	assert.Equal(t, 12345678, dec.ReadInt())
	assert.Equal(t, 123456789, dec.ReadInt())
	assert.Equal(t, 1234567890, dec.ReadInt())
	assert.Equal(t, int8(i32), dec.ReadInt8())
	assert.Equal(t, uint8(u32), dec.ReadUint8())
	assert.Equal(t, int16(i64), dec.ReadInt16())
	assert.Equal(t, uint16(u64), dec.ReadUint16())
	assert.Equal(t, i32, dec.ReadInt32())
	assert.Equal(t, u32, dec.ReadUint32())
	assert.Equal(t, i64, dec.ReadInt64())
	assert.Equal(t, u64, dec.ReadUint64())
	assert.Equal(t, int(i32), dec.ReadInt())
	assert.Equal(t, int(u32), dec.ReadInt())
	assert.Equal(t, int(i64), dec.ReadInt())
	assert.Equal(t, int(u64), dec.ReadInt())
	assert.Equal(t, uint(i32), dec.ReadUint())
	assert.Equal(t, uint(u32), dec.ReadUint())
	assert.Equal(t, uint(i64), dec.ReadUint())
	assert.Equal(t, uint(u64), dec.ReadUint())
}

func TestResetReader(t *testing.T) {
	reader := bytes.NewBufferString(";1;12;123;1234;12345;123456;1234567;12345678;123456789;1234567890;-2147483648;2147483647;-9223372036854775808;18446744073709551615;-2147483648;2147483647;-9223372036854775808;18446744073709551615;-2147483648;2147483647;-9223372036854775808;18446744073709551615;-2147483648;2147483647;-9223372036854775808;18446744073709551615;;1;12;123;1234;12345;123456;1234567;12345678;123456789;1234567890")
	dec := &Decoder{}
	dec.ResetReader(reader)
	i32 := int32(-2147483648)
	u32 := uint32(2147483647)
	i64 := int64(-9223372036854775808)
	u64 := uint64(18446744073709551615)
	assert.Equal(t, 0, dec.ReadInt())
	assert.Equal(t, 1, dec.ReadInt())
	assert.Equal(t, 12, dec.ReadInt())
	assert.Equal(t, 123, dec.ReadInt())
	assert.Equal(t, 1234, dec.ReadInt())
	assert.Equal(t, 12345, dec.ReadInt())
	assert.Equal(t, 123456, dec.ReadInt())
	assert.Equal(t, 1234567, dec.ReadInt())
	assert.Equal(t, 12345678, dec.ReadInt())
	assert.Equal(t, 123456789, dec.ReadInt())
	assert.Equal(t, 1234567890, dec.ReadInt())
	assert.Equal(t, int8(i32), dec.ReadInt8())
	assert.Equal(t, uint8(u32), dec.ReadUint8())
	assert.Equal(t, int16(i64), dec.ReadInt16())
	assert.Equal(t, uint16(u64), dec.ReadUint16())
	assert.Equal(t, i32, dec.ReadInt32())
	assert.Equal(t, u32, dec.ReadUint32())
	assert.Equal(t, i64, dec.ReadInt64())
	assert.Equal(t, u64, dec.ReadUint64())
	assert.Equal(t, int(i32), dec.ReadInt())
	assert.Equal(t, int(u32), dec.ReadInt())
	assert.Equal(t, int(i64), dec.ReadInt())
	assert.Equal(t, int(u64), dec.ReadInt())
	assert.Equal(t, uint(i32), dec.ReadUint())
	assert.Equal(t, uint(u32), dec.ReadUint())
	assert.Equal(t, uint(i64), dec.ReadUint())
	assert.Equal(t, uint(u64), dec.ReadUint())
}

func TestResetBytes(t *testing.T) {
	data := ([]byte)(";1;12;123;1234;12345;123456;1234567;12345678;123456789;1234567890;-2147483648;2147483647;-9223372036854775808;18446744073709551615;-2147483648;2147483647;-9223372036854775808;18446744073709551615;-2147483648;2147483647;-9223372036854775808;18446744073709551615;-2147483648;2147483647;-9223372036854775808;18446744073709551615;;1;12;123;1234;12345;123456;1234567;12345678;123456789;1234567890")
	dec := &Decoder{}
	dec.ResetBytes(data)
	i32 := int32(-2147483648)
	u32 := uint32(2147483647)
	i64 := int64(-9223372036854775808)
	u64 := uint64(18446744073709551615)
	assert.Equal(t, 0, dec.ReadInt())
	assert.Equal(t, 1, dec.ReadInt())
	assert.Equal(t, 12, dec.ReadInt())
	assert.Equal(t, 123, dec.ReadInt())
	assert.Equal(t, 1234, dec.ReadInt())
	assert.Equal(t, 12345, dec.ReadInt())
	assert.Equal(t, 123456, dec.ReadInt())
	assert.Equal(t, 1234567, dec.ReadInt())
	assert.Equal(t, 12345678, dec.ReadInt())
	assert.Equal(t, 123456789, dec.ReadInt())
	assert.Equal(t, 1234567890, dec.ReadInt())
	assert.Equal(t, int8(i32), dec.ReadInt8())
	assert.Equal(t, uint8(u32), dec.ReadUint8())
	assert.Equal(t, int16(i64), dec.ReadInt16())
	assert.Equal(t, uint16(u64), dec.ReadUint16())
	assert.Equal(t, i32, dec.ReadInt32())
	assert.Equal(t, u32, dec.ReadUint32())
	assert.Equal(t, i64, dec.ReadInt64())
	assert.Equal(t, u64, dec.ReadUint64())
	assert.Equal(t, int(i32), dec.ReadInt())
	assert.Equal(t, int(u32), dec.ReadInt())
	assert.Equal(t, int(i64), dec.ReadInt())
	assert.Equal(t, int(u64), dec.ReadInt())
	assert.Equal(t, uint(i32), dec.ReadUint())
	assert.Equal(t, uint(u32), dec.ReadUint())
	assert.Equal(t, uint(i64), dec.ReadUint())
	assert.Equal(t, uint(u64), dec.ReadUint())
}

func TestNext(t *testing.T) {
	data := ([]byte)(";1;12;123;1234;12345;123456;1234567;12345678;123456789;1234567890")
	dec := NewDecoderFromReader(bytes.NewBuffer(data), 32)
	assert.Equal(t, ";1;12;123;1234;12345", string(dec.Next(20)))
	assert.Equal(t, ";123456;1234567;1234", string(dec.Next(20)))
	assert.Equal(t, "5678;123456789;12345", string(dec.Next(20)))
	assert.Equal(t, "67890", string(dec.Next(20)))
	assert.Equal(t, "", string(dec.Next(20)))
	assert.Equal(t, byte(0), dec.NextByte())
}

func TestRemains(t *testing.T) {
	data := ([]byte)(";1;12;123;1234;12345;123456;1234567;12345678;123456789;1234567890")
	dec := NewDecoderFromReader(bytes.NewBuffer(data), 32)
	assert.Equal(t, ";1;12;123;1234;12345", string(dec.Next(20)))
	assert.Equal(t, ";123456;1234567;1234", string(dec.Next(20)))
	assert.Equal(t, "5678;123456789;1234567890", string(dec.Remains()))
	assert.Equal(t, "", string(dec.Remains()))
	assert.Equal(t, byte(0), dec.NextByte())
	assert.EqualError(t, io.EOF, dec.Error.Error())
}

func BenchmarkReadIntFromReader(b *testing.B) {
	data := ([]byte)(";1;12;123;1234;12345;123456;1234567;12345678;123456789;1234567890;-2147483648;2147483647;-9223372036854775808;18446744073709551615;-2147483648;2147483647;-9223372036854775808;18446744073709551615;-2147483648;2147483647;-9223372036854775808;18446744073709551615;-2147483648;2147483647;-9223372036854775808;18446744073709551615;;1;12;123;1234;12345;123456;1234567;12345678;123456789;1234567890")
	dec := &Decoder{}
	for i := 0; i < b.N; i++ {
		reader := bytes.NewBuffer(data)
		dec.ResetReader(reader)
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
	}
}

func BenchmarkReadIntFromBytes(b *testing.B) {
	data := ([]byte)(";1;12;123;1234;12345;123456;1234567;12345678;123456789;1234567890;-2147483648;2147483647;-9223372036854775808;18446744073709551615;-2147483648;2147483647;-9223372036854775808;18446744073709551615;-2147483648;2147483647;-9223372036854775808;18446744073709551615;-2147483648;2147483647;-9223372036854775808;18446744073709551615;;1;12;123;1234;12345;123456;1234567;12345678;123456789;1234567890")
	dec := &Decoder{}
	for i := 0; i < b.N; i++ {
		dec.ResetBytes(data)
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
		dec.ReadInt()
	}
}
