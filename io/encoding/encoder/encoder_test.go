/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/encoding/encoder/encoder_test.go                      |
|                                                          |
| LastModified: Feb 25, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoder

import (
	"math"
	"math/big"
	"strings"
	"testing"
)

func TestEncoderEncode(t *testing.T) {
	sb := new(strings.Builder)
	enc := &Encoder{Writer: sb}
	i := 0
	i8 := int8(1)
	i16 := int16(2)
	i32 := int32(3)
	i64 := int64(4)
	u := uint(5)
	u8 := uint8(6)
	u16 := uint16(7)
	u32 := uint32(8)
	u64 := uint64(9)
	b := true
	f32 := float32(math.Pi)
	f64 := float64(math.Pi)
	e := ""
	c := "我"
	s := "Hello"
	c64 := complex(float32(1), float32(2))
	c128 := complex(float64(3), float64(4))
	r64 := complex(float32(5), float32(0))
	r128 := complex(float64(6), float64(0))
	bi := big.NewInt(0)
	bf := big.NewFloat(1)
	br := big.NewRat(2, 3)
	bri := big.NewRat(4, 1)
	if err := enc.Encode(nil); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(i); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(i8); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(i16); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(i32); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(i64); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(u); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(u8); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(u16); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(u32); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(u64); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(b); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(f32); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(f64); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(e); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(c); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(s); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(c64); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(c128); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(r64); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(r128); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(*bi); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(*bf); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(*br); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(*bri); err != nil {
		t.Error(err)
	}
	if err := enc.Encode((*int)(nil)); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&i); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&i8); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&i16); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&i32); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&i64); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&u); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&u8); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&u16); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&u32); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&u64); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&b); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&f32); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&f64); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&e); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&c); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&s); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&c64); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&c128); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&r64); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(&r128); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(bi); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(bf); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(br); err != nil {
		t.Error(err)
	}
	if err := enc.Encode(bri); err != nil {
		t.Error(err)
	}
	if sb.String() != ``+
		`n0123456789td3.1415927;d3.141592653589793;`+
		`eu我s5"Hello"a2{d1;d2;}a2{d3;d4;}d5;d6;l0;d1;s3"2/3"l4;`+
		`n0123456789td3.1415927;d3.141592653589793;`+
		`eu我s5"Hello"a2{d1;d2;}a2{d3;d4;}d5;d6;l0;d1;s3"2/3"l4;` {
		t.Error(sb)
	}
}
