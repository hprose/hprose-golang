/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/decoder_refer_test.go                           |
|                                                          |
| LastModified: Apr 27, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding_test

import (
	"container/list"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	. "github.com/hprose/hprose-golang/v3/encoding"
	"github.com/stretchr/testify/assert"
)

func TestDecodeRefer(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Simple(false)
	for i := 0; i < 50; i++ {
		enc.Encode("123")
	}
	dec := NewDecoder(([]byte)(sb.String()))
	dec.Simple(false)
	var s string
	dec.Decode(&s)
	assert.Equal(t, "123", s)
	dec.Decode(&s)
	assert.Equal(t, "123", s)
	var b bool
	dec.Decode(&b)
	assert.Equal(t, false, b)
	assert.Error(t, dec.Error)
	dec.Error = nil
	var i int
	dec.Decode(&i)
	assert.Equal(t, 123, i)
	var i8 int8
	dec.Decode(&i8)
	assert.Equal(t, (int8)(123), i8)
	var i16 int16
	dec.Decode(&i16)
	assert.Equal(t, (int16)(123), i16)
	var i32 int32
	dec.Decode(&i32)
	assert.Equal(t, (int32)(123), i32)
	var i64 int64
	dec.Decode(&i64)
	assert.Equal(t, (int64)(123), i64)
	var u uint
	dec.Decode(&u)
	assert.Equal(t, (uint)(123), u)
	var u8 uint8
	dec.Decode(&u8)
	assert.Equal(t, (uint8)(123), u8)
	var u16 uint16
	dec.Decode(&u16)
	assert.Equal(t, (uint16)(123), u16)
	var u32 uint32
	dec.Decode(&u32)
	assert.Equal(t, (uint32)(123), u32)
	var u64 uint64
	dec.Decode(&u64)
	assert.Equal(t, (uint64)(123), u64)
	var uptr uintptr
	dec.Decode(&uptr)
	assert.Equal(t, (uintptr)(123), uptr)
	var f32 float32
	dec.Decode(&f32)
	assert.Equal(t, (float32)(123), f32)
	var f64 float64
	dec.Decode(&f64)
	assert.Equal(t, (float64)(123), f64)
	var c64 complex64
	dec.Decode(&c64)
	assert.Equal(t, (complex64)(123), c64)
	var c128 complex128
	dec.Decode(&c128)
	assert.Equal(t, (complex128)(123), c128)
	var sp *string
	dec.Decode(&sp)
	assert.Equal(t, "123", *sp)
	var bp *bool
	dec.Decode(&bp)
	assert.Equal(t, false, *bp)
	assert.Error(t, dec.Error)
	dec.Error = nil
	var ip *int
	dec.Decode(&ip)
	assert.Equal(t, 123, *ip)
	var i8p *int8
	dec.Decode(&i8p)
	assert.Equal(t, (int8)(123), *i8p)
	var i16p *int16
	dec.Decode(&i16p)
	assert.Equal(t, (int16)(123), *i16p)
	var i32p *int32
	dec.Decode(&i32p)
	assert.Equal(t, (int32)(123), *i32p)
	var i64p *int64
	dec.Decode(&i64p)
	assert.Equal(t, (int64)(123), *i64p)
	var up *uint
	dec.Decode(&up)
	assert.Equal(t, (uint)(123), *up)
	var u8p *uint8
	dec.Decode(&u8p)
	assert.Equal(t, (uint8)(123), *u8p)
	var u16p *uint16
	dec.Decode(&u16p)
	assert.Equal(t, (uint16)(123), *u16p)
	var u32p *uint32
	dec.Decode(&u32p)
	assert.Equal(t, (uint32)(123), *u32p)
	var u64p *uint64
	dec.Decode(&u64p)
	assert.Equal(t, (uint64)(123), *u64p)
	var uptrp *uintptr
	dec.Decode(&uptrp)
	assert.Equal(t, (uintptr)(123), *uptrp)
	var f32p *float32
	dec.Decode(&f32p)
	assert.Equal(t, (float32)(123), *f32p)
	var f64p *float64
	dec.Decode(&f64p)
	assert.Equal(t, (float64)(123), *f64p)
	var c64p *complex64
	dec.Decode(&c64p)
	assert.Equal(t, (complex64)(123), *c64p)
	var c128p *complex128
	dec.Decode(&c128p)
	assert.Equal(t, (complex128)(123), *c128p)
	var biv big.Int
	dec.Decode(&biv)
	assert.Equal(t, *big.NewInt(123), biv)
	var bi *big.Int
	dec.Decode(&bi)
	assert.Equal(t, big.NewInt(123), bi)
	var bip **big.Int
	dec.Decode(&bip)
	assert.Equal(t, big.NewInt(123), *bip)
	var bfv big.Float
	dec.Decode(&bfv)
	assert.Equal(t, big.NewFloat(123).String(), bfv.String())
	var bf *big.Float
	dec.Decode(&bf)
	assert.Equal(t, big.NewFloat(123).String(), bf.String())
	var brv big.Rat
	dec.Decode(&brv)
	assert.Equal(t, big.NewRat(123, 1).RatString(), brv.RatString())
	var br *big.Rat
	dec.Decode(&br)
	assert.Equal(t, big.NewRat(123, 1).RatString(), br.RatString())
	var spp **string
	dec.Decode(&spp)
	assert.Equal(t, "123", **spp)
	var sppp ***string
	dec.Decode(&sppp)
	assert.Equal(t, "123", ***sppp)
	var spppp ****string
	dec.Decode(&spppp)
	assert.Equal(t, "123", ****spppp)
	var bytes []byte
	dec.Decode(&bytes)
	assert.Equal(t, []byte("123"), bytes)
	var bytesp *[]byte
	dec.Decode(&bytesp)
	assert.Equal(t, []byte("123"), *bytesp)
	var ints []int
	dec.Decode(&ints)
	assert.Error(t, dec.Error)
}

func TestDecodeTimeRefer(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Simple(false)
	for i := 0; i < 3; i++ {
		enc.Encode("2021-02-16 17:26:54")
	}
	dec := NewDecoder(([]byte)(sb.String()))
	dec.Simple(false)
	var s string
	dec.Decode(&s)
	assert.Equal(t, "2021-02-16 17:26:54", s)
	et, _ := time.Parse("2006-01-02 15:04:05", "2021-02-16 17:26:54")
	var at time.Time
	dec.Decode(&at)
	assert.Equal(t, et, at)
	var atp *time.Time
	dec.Decode(&atp)
	assert.Equal(t, et, *atp)

	sb = new(strings.Builder)
	enc = NewEncoder(sb)
	enc.Simple(false)
	for i := 0; i < 4; i++ {
		enc.Encode(&et)
	}
	dec = NewDecoder(([]byte)(sb.String()))
	dec.Simple(false)
	dec.Decode(&s)
	assert.Equal(t, "2021-02-16 17:26:54 +0000 UTC", s)
	dec.Decode(&s)
	assert.Equal(t, "2021-02-16 17:26:54 +0000 UTC", s)
	assert.NoError(t, dec.Error)
	dec.Decode(&at)
	assert.Equal(t, et, at)
	dec.Decode(&atp)
	assert.Equal(t, et, *atp)
}

func TestDecodeUUIDRefer(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Simple(false)
	src := uuid.UUID{
		0x7d, 0x44, 0x48, 0x40,
		0x9d, 0xc0,
		0x11, 0xd1,
		0xb2, 0x45,
		0x5f, 0xfd, 0xce, 0x74, 0xfa, 0xd2,
	}
	for i := 0; i < 3; i++ {
		enc.Encode(&src)
	}
	dec := NewDecoder(([]byte)(sb.String()))
	dec.Simple(false)
	var s string
	dec.Decode(&s)
	assert.Equal(t, "7d444840-9dc0-11d1-b245-5ffdce74fad2", s)
	var u uuid.UUID
	dec.Decode(&u)
	assert.Equal(t, src, u)
	var up *uuid.UUID
	dec.Decode(&up)
	assert.Equal(t, src, *up)

	sb = new(strings.Builder)
	enc = NewEncoder(sb)
	enc.Simple(false)
	for i := 0; i < 4; i++ {
		enc.Encode("f028e399-3807-47e9-bcc6-a07ec2efb06a")
	}
	assert.Equal(t, `s36"f028e399-3807-47e9-bcc6-a07ec2efb06a"r0;r0;r0;`, sb.String())

	src = uuid.MustParse("{f028e399-3807-47e9-bcc6-a07ec2efb06a}")
	dec = NewDecoder(([]byte)(sb.String()))
	dec.Simple(false)
	dec.Decode(&s)
	assert.Equal(t, "f028e399-3807-47e9-bcc6-a07ec2efb06a", s)
	dec.Decode(&s)
	assert.Equal(t, "f028e399-3807-47e9-bcc6-a07ec2efb06a", s)
	dec.Decode(&u)
	assert.Equal(t, src, u)
	dec.Decode(&up)
	assert.Equal(t, src, *up)
}

func TestDecodeSliceRefer(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Simple(false)
	src := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < 3; i++ {
		enc.Encode(&src)
	}
	assert.Equal(t, `a6{123456}r0;r0;`, sb.String())

	dec := NewDecoder(([]byte)(sb.String()))
	dec.Simple(false)
	var ints []int
	dec.Decode(&ints)
	assert.Equal(t, src, ints)
	var intsp *[]int
	dec.Decode(&intsp)
	assert.Equal(t, src, *intsp)
	var intspp **[]int
	dec.Decode(&intspp)
	assert.Equal(t, src, **intspp)
}

func TestDecodeBytesRefer(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Simple(false)
	src := []byte("hello world")
	for i := 0; i < 4; i++ {
		enc.Encode(&src)
	}
	assert.Equal(t, `b11"hello world"r0;r0;r0;`, sb.String())

	dec := NewDecoder(([]byte)(sb.String()))
	dec.Simple(false)
	var bytes []byte
	dec.Decode(&bytes)
	assert.Equal(t, src, bytes)
	var bytesp *[]byte
	dec.Decode(&bytesp)
	assert.Equal(t, src, *bytesp)
	var bytespp **[]byte
	dec.Decode(&bytespp)
	assert.Equal(t, src, **bytespp)
	var s string
	dec.Decode(&s)
	assert.Equal(t, "hello world", s)
}

func TestDecodeListRefer(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Simple(false)
	src := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < 3; i++ {
		enc.Encode(&src)
	}
	assert.Equal(t, `a6{123456}r0;r0;`, sb.String())
	dec := NewDecoder(([]byte)(sb.String()))
	dec.Simple(false)
	var list1 *list.List
	dec.Decode(&list1)
	var list2 *list.List
	dec.Decode(&list2)
	assert.Equal(t, list1, list2)
	var list3 **list.List
	dec.Decode(&list3)
	assert.Equal(t, list2, *list3)
	println(list2, *list3)
}

func TestDecodeArrayRefer(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Simple(false)
	src := [6]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < 3; i++ {
		enc.Encode(&src)
	}
	assert.Equal(t, `a6{123456}r0;r0;`, sb.String())

	dec := NewDecoder(([]byte)(sb.String()))
	dec.Simple(false)
	var ints [6]int
	dec.Decode(&ints)
	assert.Equal(t, src, ints)
	var intsp *[6]int
	dec.Decode(&intsp)
	assert.Equal(t, src, *intsp)
	var intspp **[6]int
	dec.Decode(&intspp)
	assert.Equal(t, src, **intspp)
}

func TestDecodeMapRefer(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Simple(false)
	src := map[string]interface{}{
		"name": "张三",
		"age":  18,
	}
	for i := 0; i < 3; i++ {
		enc.Encode(&src)
	}
	dec := NewDecoder(([]byte)(sb.String()))
	dec.Simple(false)
	var m map[string]interface{}
	dec.Decode(&m)
	assert.Equal(t, src, m)
	var mp *map[string]interface{}
	dec.Decode(&mp)
	assert.Equal(t, src, *mp)
	var mpp **map[string]interface{}
	dec.Decode(&mpp)
	assert.Equal(t, src, **mpp)
}

func TestDecodeStructRefer(t *testing.T) {
	type TestStruct struct {
		A int
		B bool    `hprose:"-"`
		C string  `json:"json,omitempty"`
		D float32 `json:",omitempty"`
		e float64
	}
	src := TestStruct{1, true, "hello", 3.14, 2.718}
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	for i := 0; i < 4; i++ {
		enc.Encode(&src)
	}
	dec := NewDecoder(([]byte)(sb.String()))
	dec.Simple(false)
	var ts TestStruct
	dec.Decode(&ts)
	assert.Equal(t, TestStruct{1, false, "hello", 3.14, 0}, ts)
	var ts2 TestStruct
	dec.Decode(&ts2)
	assert.Equal(t, ts, ts2)
	var tsp *TestStruct
	dec.Decode(&tsp)
	assert.Equal(t, ts, *tsp)
	var tspp **TestStruct
	dec.Decode(&tspp)
	assert.Equal(t, tsp, *tspp)
}
