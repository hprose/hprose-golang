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

func TestDecodeInt(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
	enc.Encode(-1)
	enc.Encode(0)
	enc.Encode(1)
	enc.Encode(123)
	enc.Encode(math.MinInt64)
	enc.Encode(-math.MaxInt64)
	enc.Encode(uint64(math.MaxUint64))
	enc.Encode(true)
	enc.Encode(false)
	enc.Encode(nil)
	enc.Encode(3.14)
	enc.Encode("")
	enc.Encode("1")
	enc.Encode("123")
	enc.Encode("N")
	enc.Encode("NaN")
	enc.Encode(nil)
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var i int
	var maxUint64 uint64 = math.MaxUint64
	dec.Decode(&i)
	assert.Equal(t, -1, i)
	dec.Decode(&i)
	assert.Equal(t, 0, i)
	dec.Decode(&i)
	assert.Equal(t, 1, i)
	dec.Decode(&i)
	assert.Equal(t, 123, i)
	dec.Decode(&i)
	assert.Equal(t, math.MinInt64, i)
	dec.Decode(&i)
	assert.Equal(t, -math.MaxInt64, i)
	dec.Decode(&i)
	assert.Equal(t, int(maxUint64), i)
	dec.Decode(&i)
	assert.Equal(t, 1, i)
	dec.Decode(&i)
	assert.Equal(t, 0, i)
	dec.Decode(&i)
	assert.Equal(t, 0, i)
	dec.Decode(&i)
	assert.Equal(t, 3, i)
	dec.Decode(&i)
	assert.Equal(t, 0, i)
	dec.Decode(&i)
	assert.Equal(t, 1, i)
	dec.Decode(&i)
	assert.Equal(t, 123, i)
	assert.NoError(t, dec.Error)
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `strconv.ParseInt: parsing "N": invalid syntax`)
	dec.Error = nil
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `strconv.ParseInt: parsing "NaN": invalid syntax`)
	dec.Error = nil
	var ip *int
	dec.Decode(&ip)
	assert.Nil(t, ip) // nil
	dec.Decode(&ip)
	assert.Equal(t, 1, *ip) // 1
}

func TestDecodeInt8(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
	enc.Encode(-1)
	enc.Encode(0)
	enc.Encode(1)
	enc.Encode(123)
	enc.Encode(math.MinInt64)
	enc.Encode(-math.MaxInt64)
	enc.Encode(uint64(math.MaxUint64))
	enc.Encode(true)
	enc.Encode(false)
	enc.Encode(nil)
	enc.Encode(3.14)
	enc.Encode("")
	enc.Encode("1")
	enc.Encode("123")
	enc.Encode("N")
	enc.Encode("NaN")
	enc.Encode(nil)
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var i int8
	minInt64, maxInt64 := math.MinInt64, math.MaxInt64
	var maxUint64 uint64 = math.MaxUint64
	dec.Decode(&i)
	assert.Equal(t, int8(-1), i)
	dec.Decode(&i)
	assert.Equal(t, int8(0), i)
	dec.Decode(&i)
	assert.Equal(t, int8(1), i)
	dec.Decode(&i)
	assert.Equal(t, int8(123), i)
	dec.Decode(&i)
	assert.Equal(t, int8(minInt64), i)
	dec.Decode(&i)
	assert.Equal(t, int8(-maxInt64), i)
	dec.Decode(&i)
	assert.Equal(t, int8(maxUint64), i)
	dec.Decode(&i)
	assert.Equal(t, int8(1), i)
	dec.Decode(&i)
	assert.Equal(t, int8(0), i)
	dec.Decode(&i)
	assert.Equal(t, int8(0), i)
	dec.Decode(&i)
	assert.Equal(t, int8(3), i)
	dec.Decode(&i)
	assert.Equal(t, int8(0), i)
	dec.Decode(&i)
	assert.Equal(t, int8(1), i)
	dec.Decode(&i)
	assert.Equal(t, int8(123), i)
	assert.NoError(t, dec.Error)
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `strconv.ParseInt: parsing "N": invalid syntax`)
	dec.Error = nil
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `strconv.ParseInt: parsing "NaN": invalid syntax`)
	dec.Error = nil
	var ip *int8
	dec.Decode(&ip)
	assert.Nil(t, ip) // nil
	dec.Decode(&ip)
	assert.Equal(t, int8(1), *ip) // 1
}

func TestDecodeInt16(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
	enc.Encode(-1)
	enc.Encode(0)
	enc.Encode(1)
	enc.Encode(123)
	enc.Encode(math.MinInt64)
	enc.Encode(-math.MaxInt64)
	enc.Encode(uint64(math.MaxUint64))
	enc.Encode(true)
	enc.Encode(false)
	enc.Encode(nil)
	enc.Encode(3.14)
	enc.Encode("")
	enc.Encode("1")
	enc.Encode("123")
	enc.Encode("N")
	enc.Encode("NaN")
	enc.Encode(nil)
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var i int16
	minInt64, maxInt64 := math.MinInt64, math.MaxInt64
	var maxUint64 uint64 = math.MaxUint64
	dec.Decode(&i)
	assert.Equal(t, int16(-1), i)
	dec.Decode(&i)
	assert.Equal(t, int16(0), i)
	dec.Decode(&i)
	assert.Equal(t, int16(1), i)
	dec.Decode(&i)
	assert.Equal(t, int16(123), i)
	dec.Decode(&i)
	assert.Equal(t, int16(minInt64), i)
	dec.Decode(&i)
	assert.Equal(t, int16(-maxInt64), i)
	dec.Decode(&i)
	assert.Equal(t, int16(maxUint64), i)
	dec.Decode(&i)
	assert.Equal(t, int16(1), i)
	dec.Decode(&i)
	assert.Equal(t, int16(0), i)
	dec.Decode(&i)
	assert.Equal(t, int16(0), i)
	dec.Decode(&i)
	assert.Equal(t, int16(3), i)
	dec.Decode(&i)
	assert.Equal(t, int16(0), i)
	dec.Decode(&i)
	assert.Equal(t, int16(1), i)
	dec.Decode(&i)
	assert.Equal(t, int16(123), i)
	assert.NoError(t, dec.Error)
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `strconv.ParseInt: parsing "N": invalid syntax`)
	dec.Error = nil
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `strconv.ParseInt: parsing "NaN": invalid syntax`)
	dec.Error = nil
	var ip *int16
	dec.Decode(&ip)
	assert.Nil(t, ip) // nil
	dec.Decode(&ip)
	assert.Equal(t, int16(1), *ip) // 1
}

func TestDecodeInt32(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
	enc.Encode(-1)
	enc.Encode(0)
	enc.Encode(1)
	enc.Encode(123)
	enc.Encode(math.MinInt64)
	enc.Encode(-math.MaxInt64)
	enc.Encode(uint64(math.MaxUint64))
	enc.Encode(true)
	enc.Encode(false)
	enc.Encode(nil)
	enc.Encode(3.14)
	enc.Encode("")
	enc.Encode("1")
	enc.Encode("123")
	enc.Encode("N")
	enc.Encode("NaN")
	enc.Encode(nil)
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var i int32
	minInt64, maxInt64 := math.MinInt64, math.MaxInt64
	var maxUint64 uint64 = math.MaxUint64
	dec.Decode(&i)
	assert.Equal(t, int32(-1), i)
	dec.Decode(&i)
	assert.Equal(t, int32(0), i)
	dec.Decode(&i)
	assert.Equal(t, int32(1), i)
	dec.Decode(&i)
	assert.Equal(t, int32(123), i)
	dec.Decode(&i)
	assert.Equal(t, int32(minInt64), i)
	dec.Decode(&i)
	assert.Equal(t, int32(-maxInt64), i)
	dec.Decode(&i)
	assert.Equal(t, int32(maxUint64), i)
	dec.Decode(&i)
	assert.Equal(t, int32(1), i)
	dec.Decode(&i)
	assert.Equal(t, int32(0), i)
	dec.Decode(&i)
	assert.Equal(t, int32(0), i)
	dec.Decode(&i)
	assert.Equal(t, int32(3), i)
	dec.Decode(&i)
	assert.Equal(t, int32(0), i)
	dec.Decode(&i)
	assert.Equal(t, int32(1), i)
	dec.Decode(&i)
	assert.Equal(t, int32(123), i)
	assert.NoError(t, dec.Error)
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `strconv.ParseInt: parsing "N": invalid syntax`)
	dec.Error = nil
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `strconv.ParseInt: parsing "NaN": invalid syntax`)
	dec.Error = nil
	var ip *int32
	dec.Decode(&ip)
	assert.Nil(t, ip) // nil
	dec.Decode(&ip)
	assert.Equal(t, int32(1), *ip) // 1
}

func TestDecodeInt64(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
	enc.Encode(-1)
	enc.Encode(0)
	enc.Encode(1)
	enc.Encode(123)
	enc.Encode(math.MinInt64)
	enc.Encode(-math.MaxInt64)
	enc.Encode(uint64(math.MaxUint64))
	enc.Encode(true)
	enc.Encode(false)
	enc.Encode(nil)
	enc.Encode(3.14)
	enc.Encode("")
	enc.Encode("1")
	enc.Encode("123")
	enc.Encode("N")
	enc.Encode("NaN")
	enc.Encode(nil)
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var i int64
	minInt64, maxInt64 := math.MinInt64, math.MaxInt64
	var maxUint64 uint64 = math.MaxUint64
	dec.Decode(&i)
	assert.Equal(t, int64(-1), i)
	dec.Decode(&i)
	assert.Equal(t, int64(0), i)
	dec.Decode(&i)
	assert.Equal(t, int64(1), i)
	dec.Decode(&i)
	assert.Equal(t, int64(123), i)
	dec.Decode(&i)
	assert.Equal(t, int64(minInt64), i)
	dec.Decode(&i)
	assert.Equal(t, int64(-maxInt64), i)
	dec.Decode(&i)
	assert.Equal(t, int64(maxUint64), i)
	dec.Decode(&i)
	assert.Equal(t, int64(1), i)
	dec.Decode(&i)
	assert.Equal(t, int64(0), i)
	dec.Decode(&i)
	assert.Equal(t, int64(0), i)
	dec.Decode(&i)
	assert.Equal(t, int64(3), i)
	dec.Decode(&i)
	assert.Equal(t, int64(0), i)
	dec.Decode(&i)
	assert.Equal(t, int64(1), i)
	dec.Decode(&i)
	assert.Equal(t, int64(123), i)
	assert.NoError(t, dec.Error)
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `strconv.ParseInt: parsing "N": invalid syntax`)
	dec.Error = nil
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `strconv.ParseInt: parsing "NaN": invalid syntax`)
	dec.Error = nil
	var ip *int64
	dec.Decode(&ip)
	assert.Nil(t, ip) // nil
	dec.Decode(&ip)
	assert.Equal(t, int64(1), *ip) // 1
}

func TestDecodeUint(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
	enc.Encode(-1)
	enc.Encode(0)
	enc.Encode(1)
	enc.Encode(123)
	enc.Encode(math.MinInt64)
	enc.Encode(-math.MaxInt64)
	enc.Encode(uint64(math.MaxUint64))
	enc.Encode(true)
	enc.Encode(false)
	enc.Encode(nil)
	enc.Encode(3.14)
	enc.Encode("")
	enc.Encode("1")
	enc.Encode("123")
	enc.Encode("N")
	enc.Encode("NaN")
	enc.Encode(nil)
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var i uint
	one, minInt64, maxInt64 := 1, math.MinInt64, math.MaxInt64
	var maxUint64 uint64 = math.MaxUint64
	dec.Decode(&i)
	assert.Equal(t, uint(-one), i)
	dec.Decode(&i)
	assert.Equal(t, uint(0), i)
	dec.Decode(&i)
	assert.Equal(t, uint(1), i)
	dec.Decode(&i)
	assert.Equal(t, uint(123), i)
	dec.Decode(&i)
	assert.Equal(t, uint(minInt64), i)
	dec.Decode(&i)
	assert.Equal(t, uint(-maxInt64), i)
	dec.Decode(&i)
	assert.Equal(t, uint(maxUint64), i)
	dec.Decode(&i)
	assert.Equal(t, uint(1), i)
	dec.Decode(&i)
	assert.Equal(t, uint(0), i)
	dec.Decode(&i)
	assert.Equal(t, uint(0), i)
	dec.Decode(&i)
	assert.Equal(t, uint(3), i)
	dec.Decode(&i)
	assert.Equal(t, uint(0), i)
	dec.Decode(&i)
	assert.Equal(t, uint(1), i)
	dec.Decode(&i)
	assert.Equal(t, uint(123), i)
	assert.NoError(t, dec.Error)
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `strconv.ParseUint: parsing "N": invalid syntax`)
	dec.Error = nil
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `strconv.ParseUint: parsing "NaN": invalid syntax`)
	dec.Error = nil
	var ip *uint
	dec.Decode(&ip)
	assert.Nil(t, ip) // nil
	dec.Decode(&ip)
	assert.Equal(t, uint(1), *ip) // 1
}

func TestDecodeUint8(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
	enc.Encode(-1)
	enc.Encode(0)
	enc.Encode(1)
	enc.Encode(123)
	enc.Encode(math.MinInt64)
	enc.Encode(-math.MaxInt64)
	enc.Encode(uint64(math.MaxUint64))
	enc.Encode(true)
	enc.Encode(false)
	enc.Encode(nil)
	enc.Encode(3.14)
	enc.Encode("")
	enc.Encode("1")
	enc.Encode("123")
	enc.Encode("N")
	enc.Encode("NaN")
	enc.Encode(nil)
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var i uint8
	one, minInt64, maxInt64 := 1, math.MinInt64, math.MaxInt64
	var maxUint64 uint64 = math.MaxUint64
	dec.Decode(&i)
	assert.Equal(t, uint8(-one), i)
	dec.Decode(&i)
	assert.Equal(t, uint8(0), i)
	dec.Decode(&i)
	assert.Equal(t, uint8(1), i)
	dec.Decode(&i)
	assert.Equal(t, uint8(123), i)
	dec.Decode(&i)
	assert.Equal(t, uint8(minInt64), i)
	dec.Decode(&i)
	assert.Equal(t, uint8(-maxInt64), i)
	dec.Decode(&i)
	assert.Equal(t, uint8(maxUint64), i)
	dec.Decode(&i)
	assert.Equal(t, uint8(1), i)
	dec.Decode(&i)
	assert.Equal(t, uint8(0), i)
	dec.Decode(&i)
	assert.Equal(t, uint8(0), i)
	dec.Decode(&i)
	assert.Equal(t, uint8(3), i)
	dec.Decode(&i)
	assert.Equal(t, uint8(0), i)
	dec.Decode(&i)
	assert.Equal(t, uint8(1), i)
	dec.Decode(&i)
	assert.Equal(t, uint8(123), i)
	assert.NoError(t, dec.Error)
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `strconv.ParseUint: parsing "N": invalid syntax`)
	dec.Error = nil
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `strconv.ParseUint: parsing "NaN": invalid syntax`)
	dec.Error = nil
	var ip *uint8
	dec.Decode(&ip)
	assert.Nil(t, ip) // nil
	dec.Decode(&ip)
	assert.Equal(t, uint8(1), *ip) // 1
}

func TestDecodeUint16(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
	enc.Encode(-1)
	enc.Encode(0)
	enc.Encode(1)
	enc.Encode(123)
	enc.Encode(math.MinInt64)
	enc.Encode(-math.MaxInt64)
	enc.Encode(uint64(math.MaxUint64))
	enc.Encode(true)
	enc.Encode(false)
	enc.Encode(nil)
	enc.Encode(3.14)
	enc.Encode("")
	enc.Encode("1")
	enc.Encode("123")
	enc.Encode("N")
	enc.Encode("NaN")
	enc.Encode(nil)
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var i uint16
	one, minInt64, maxInt64 := 1, math.MinInt64, math.MaxInt64
	var maxUint64 uint64 = math.MaxUint64
	dec.Decode(&i)
	assert.Equal(t, uint16(-one), i)
	dec.Decode(&i)
	assert.Equal(t, uint16(0), i)
	dec.Decode(&i)
	assert.Equal(t, uint16(1), i)
	dec.Decode(&i)
	assert.Equal(t, uint16(123), i)
	dec.Decode(&i)
	assert.Equal(t, uint16(minInt64), i)
	dec.Decode(&i)
	assert.Equal(t, uint16(-maxInt64), i)
	dec.Decode(&i)
	assert.Equal(t, uint16(maxUint64), i)
	dec.Decode(&i)
	assert.Equal(t, uint16(1), i)
	dec.Decode(&i)
	assert.Equal(t, uint16(0), i)
	dec.Decode(&i)
	assert.Equal(t, uint16(0), i)
	dec.Decode(&i)
	assert.Equal(t, uint16(3), i)
	dec.Decode(&i)
	assert.Equal(t, uint16(0), i)
	dec.Decode(&i)
	assert.Equal(t, uint16(1), i)
	dec.Decode(&i)
	assert.Equal(t, uint16(123), i)
	assert.NoError(t, dec.Error)
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `strconv.ParseUint: parsing "N": invalid syntax`)
	dec.Error = nil
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `strconv.ParseUint: parsing "NaN": invalid syntax`)
	dec.Error = nil
	var ip *uint16
	dec.Decode(&ip)
	assert.Nil(t, ip) // nil
	dec.Decode(&ip)
	assert.Equal(t, uint16(1), *ip) // 1
}

func TestDecodeUint32(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
	enc.Encode(-1)
	enc.Encode(0)
	enc.Encode(1)
	enc.Encode(123)
	enc.Encode(math.MinInt64)
	enc.Encode(-math.MaxInt64)
	enc.Encode(uint64(math.MaxUint64))
	enc.Encode(true)
	enc.Encode(false)
	enc.Encode(nil)
	enc.Encode(3.14)
	enc.Encode("")
	enc.Encode("1")
	enc.Encode("123")
	enc.Encode("N")
	enc.Encode("NaN")
	enc.Encode(nil)
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var i uint32
	one, minInt64, maxInt64 := 1, math.MinInt64, math.MaxInt64
	var maxUint64 uint64 = math.MaxUint64
	dec.Decode(&i)
	assert.Equal(t, uint32(-one), i)
	dec.Decode(&i)
	assert.Equal(t, uint32(0), i)
	dec.Decode(&i)
	assert.Equal(t, uint32(1), i)
	dec.Decode(&i)
	assert.Equal(t, uint32(123), i)
	dec.Decode(&i)
	assert.Equal(t, uint32(minInt64), i)
	dec.Decode(&i)
	assert.Equal(t, uint32(-maxInt64), i)
	dec.Decode(&i)
	assert.Equal(t, uint32(maxUint64), i)
	dec.Decode(&i)
	assert.Equal(t, uint32(1), i)
	dec.Decode(&i)
	assert.Equal(t, uint32(0), i)
	dec.Decode(&i)
	assert.Equal(t, uint32(0), i)
	dec.Decode(&i)
	assert.Equal(t, uint32(3), i)
	dec.Decode(&i)
	assert.Equal(t, uint32(0), i)
	dec.Decode(&i)
	assert.Equal(t, uint32(1), i)
	dec.Decode(&i)
	assert.Equal(t, uint32(123), i)
	assert.NoError(t, dec.Error)
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `strconv.ParseUint: parsing "N": invalid syntax`)
	dec.Error = nil
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `strconv.ParseUint: parsing "NaN": invalid syntax`)
	dec.Error = nil
	var ip *uint32
	dec.Decode(&ip)
	assert.Nil(t, ip) // nil
	dec.Decode(&ip)
	assert.Equal(t, uint32(1), *ip) // 1
}

func TestDecodeUint64(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
	enc.Encode(-1)
	enc.Encode(0)
	enc.Encode(1)
	enc.Encode(123)
	enc.Encode(math.MinInt64)
	enc.Encode(-math.MaxInt64)
	enc.Encode(uint64(math.MaxUint64))
	enc.Encode(true)
	enc.Encode(false)
	enc.Encode(nil)
	enc.Encode(3.14)
	enc.Encode("")
	enc.Encode("1")
	enc.Encode("123")
	enc.Encode("N")
	enc.Encode("NaN")
	enc.Encode(nil)
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var i uint64
	one, minInt64, maxInt64 := 1, math.MinInt64, math.MaxInt64
	var maxUint64 uint64 = math.MaxUint64
	dec.Decode(&i)
	assert.Equal(t, uint64(-one), i)
	dec.Decode(&i)
	assert.Equal(t, uint64(0), i)
	dec.Decode(&i)
	assert.Equal(t, uint64(1), i)
	dec.Decode(&i)
	assert.Equal(t, uint64(123), i)
	dec.Decode(&i)
	assert.Equal(t, uint64(minInt64), i)
	dec.Decode(&i)
	assert.Equal(t, uint64(-maxInt64), i)
	dec.Decode(&i)
	assert.Equal(t, uint64(maxUint64), i)
	dec.Decode(&i)
	assert.Equal(t, uint64(1), i)
	dec.Decode(&i)
	assert.Equal(t, uint64(0), i)
	dec.Decode(&i)
	assert.Equal(t, uint64(0), i)
	dec.Decode(&i)
	assert.Equal(t, uint64(3), i)
	dec.Decode(&i)
	assert.Equal(t, uint64(0), i)
	dec.Decode(&i)
	assert.Equal(t, uint64(1), i)
	dec.Decode(&i)
	assert.Equal(t, uint64(123), i)
	assert.NoError(t, dec.Error)
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `strconv.ParseUint: parsing "N": invalid syntax`)
	dec.Error = nil
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `strconv.ParseUint: parsing "NaN": invalid syntax`)
	dec.Error = nil
	var ip *uint64
	dec.Decode(&ip)
	assert.Nil(t, ip) // nil
	dec.Decode(&ip)
	assert.Equal(t, uint64(1), *ip) // 1
}

func TestDecodeUintptr(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb, true)
	enc.Encode(-1)
	enc.Encode(0)
	enc.Encode(1)
	enc.Encode(123)
	enc.Encode(math.MinInt64)
	enc.Encode(-math.MaxInt64)
	enc.Encode(uint64(math.MaxUint64))
	enc.Encode(true)
	enc.Encode(false)
	enc.Encode(nil)
	enc.Encode(3.14)
	enc.Encode("")
	enc.Encode("1")
	enc.Encode("123")
	enc.Encode("N")
	enc.Encode("NaN")
	enc.Encode(nil)
	enc.Encode(1)
	dec := NewDecoder(([]byte)(sb.String()))
	var i uintptr
	one, minInt64, maxInt64 := 1, math.MinInt64, math.MaxInt64
	dec.Decode(&i)
	assert.Equal(t, uintptr(-one), i)
	dec.Decode(&i)
	assert.Equal(t, uintptr(0), i)
	dec.Decode(&i)
	assert.Equal(t, uintptr(1), i)
	dec.Decode(&i)
	assert.Equal(t, uintptr(123), i)
	dec.Decode(&i)
	assert.Equal(t, uintptr(minInt64), i)
	dec.Decode(&i)
	assert.Equal(t, uintptr(-maxInt64), i)
	dec.Decode(&i)
	assert.Equal(t, uintptr(math.MaxUint64), i)
	dec.Decode(&i)
	assert.Equal(t, uintptr(1), i)
	dec.Decode(&i)
	assert.Equal(t, uintptr(0), i)
	dec.Decode(&i)
	assert.Equal(t, uintptr(0), i)
	dec.Decode(&i)
	assert.Equal(t, uintptr(3), i)
	dec.Decode(&i)
	assert.Equal(t, uintptr(0), i)
	dec.Decode(&i)
	assert.Equal(t, uintptr(1), i)
	dec.Decode(&i)
	assert.Equal(t, uintptr(123), i)
	assert.NoError(t, dec.Error)
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `strconv.ParseUint: parsing "N": invalid syntax`)
	dec.Error = nil
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `strconv.ParseUint: parsing "NaN": invalid syntax`)
	dec.Error = nil
	var ip *uintptr
	dec.Decode(&ip)
	assert.Nil(t, ip) // nil
	dec.Decode(&ip)
	assert.Equal(t, uintptr(1), *ip) // 1
}
