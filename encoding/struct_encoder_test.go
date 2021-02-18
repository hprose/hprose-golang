/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/struct_encoder_test.go                          |
|                                                          |
| LastModified: Feb 18, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"errors"
	"math/big"
	"reflect"
	"strings"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestEncodeStruct(t *testing.T) {
	type TestEmbedStruct struct {
		A int
		B bool    `hprose:"-"`
		C string  `json:"json,omitempty"`
		D float32 `json:",omitempty"`
		e float64
	}
	type TestStruct struct {
		TestEmbedStruct
		Test                     TestEmbedStruct
		FuncValue                func()
		ChanValue                chan int
		UnsafePointerValue       unsafe.Pointer
		FuncPtrValue             *func()
		ChanPtrValue             *chan int
		UnsafePointerPtrValue    *unsafe.Pointer
		FuncPtrPtrValue          **func()
		ChanPtrPtrValue          **chan int
		UnsafePointerPtrPtrValue **unsafe.Pointer
		IntValue                 int          `hprose:"i"`
		Int8Value                int8         `hprose:"i8"`
		Int16Value               int16        `hprose:"i16"`
		Int32Value               int32        `hprose:"i32"`
		Int64Value               int64        `hprose:"i64"`
		UintValue                uint         `hprose:"u"`
		Uint8Value               uint8        `hprose:"u8"`
		Uint16Value              uint16       `hprose:"u16"`
		Uint32Value              uint32       `hprose:"u32"`
		Uint64Value              uint64       `hprose:"u64"`
		UintptrValue             uintptr      `hprose:"up"`
		BoolValue                bool         `hprose:"b"`
		Float32Value             float32      `hprose:"f32"`
		Float64Value             float64      `hprose:"f64"`
		Complex64Value           complex64    `hprose:"c64"`
		Complex128Value          complex128   `hprose:"c128"`
		ArrayValue               [3]int       `hprose:"iarr"`
		SliceValue               []int        `hprose:"islice"`
		EmptySliceValue          []int        `hprose:"eslice"`
		NilSliceValue            []int        `hprose:"nslice"`
		MapValue                 map[int]int  `hprose:"imap"`
		EmptyMapValue            map[int]int  `hprose:"emap"`
		NilMapValue              map[int]int  `hprose:"nmap"`
		StringValue              string       `hprose:"s"`
		EmptyStringValue         string       `hprose:"es"`
		InterfaceValue           interface{}  `hprose:"iface"`
		NilInterfaceValue        interface{}  `hprose:"niliface"`
		StructValue              *TestStruct  `hprose:"st"`
		IntPtrValue              *int         `hprose:"iptr"`
		Int8PtrValue             *int8        `hprose:"i8ptr"`
		Int16PtrValue            *int16       `hprose:"i16ptr"`
		Int32PtrValue            *int32       `hprose:"i32ptr"`
		Int64PtrValue            *int64       `hprose:"i64ptr"`
		UintPtrValue             *uint        `hprose:"uptr"`
		Uint8PtrValue            *uint8       `hprose:"u8ptr"`
		Uint16PtrValue           *uint16      `hprose:"u16ptr"`
		Uint32PtrValue           *uint32      `hprose:"u32ptr"`
		Uint64PtrValue           *uint64      `hprose:"u64ptr"`
		UintptrPtrValue          *uintptr     `hprose:"upptr"`
		BoolPtrValue             *bool        `hprose:"bptr"`
		Float32PtrValue          *float32     `hprose:"f32ptr"`
		Float64PtrValue          *float64     `hprose:"f64ptr"`
		Complex64PtrValue        *complex64   `hprose:"c64ptr"`
		Complex128PtrValue       *complex128  `hprose:"c128ptr"`
		ArrayPtrValue            *[3]int      `hprose:"iarrptr"`
		SlicePtrValue            *[]int       `hprose:"isliceptr"`
		EmptySlicePtrValue       *[]int       `hprose:"esliceptr"`
		NilSlicePtrValue         *[]int       `hprose:"nsliceptr"`
		MapPtrValue              *map[int]int `hprose:"imapptr"`
		EmptyMapPtrValue         *map[int]int `hprose:"emapptr"`
		NilMapPtrValue           *map[int]int `hprose:"nmapptr"`
		StringPtrValue           *string      `hprose:"sptr"`
		EmptyStringPtrValue      *string      `hprose:"esptr"`
		InterfacePtrValue        *interface{} `hprose:"ifaceptr"`
		NilInterfacePtrValue     *interface{} `hprose:"nilifaceptr"`
		StructPtrValue           **TestStruct `hprose:"stptr"`
		NilIntPtrValue           *int         `hprose:"niptr"`
		NilInt8PtrValue          *int8        `hprose:"ni8ptr"`
		NilInt16PtrValue         *int16       `hprose:"ni16ptr"`
		NilInt32PtrValue         *int32       `hprose:"ni32ptr"`
		NilInt64PtrValue         *int64       `hprose:"ni64ptr"`
		NilUintPtrValue          *uint        `hprose:"nuptr"`
		NilUint8PtrValue         *uint8       `hprose:"nu8ptr"`
		NilUint16PtrValue        *uint16      `hprose:"nu16ptr"`
		NilUint32PtrValue        *uint32      `hprose:"nu32ptr"`
		NilUint64PtrValue        *uint64      `hprose:"nu64ptr"`
		NilUintptrPtrValue       *uintptr     `hprose:"nupptr"`
		NilBoolPtrValue          *bool        `hprose:"nbptr"`
		NilFloat32PtrValue       *float32     `hprose:"nf32ptr"`
		NilFloat64PtrValue       *float64     `hprose:"nf64ptr"`
		NilComplex64PtrValue     *complex64   `hprose:"nc64ptr"`
		NilComplex128PtrValue    *complex128  `hprose:"nc128ptr"`
		NilArrayPtrValue         *[3]int      `hprose:"niarrptr"`
		NilEmptySlicePtrValue    *[]int       `hprose:"nesliceptr"`
		NilNilSlicePtrValue      *[]int       `hprose:"nnsliceptr"`
		NilEmptyMapPtrValue      *map[int]int `hprose:"nemapptr"`
		NilNilMapPtrValue        *map[int]int `hprose:"nnmapptr"`
		NilStringPtrValue        *string      `hprose:"nsptr"`
		NilEmptyStringPtrValue   *string      `hprose:"nesptr"`
		NilNilInterfacePtrValue  *interface{} `hprose:"nnilifaceptr"`
		NilStructPtrValue        **TestStruct `hprose:"nstptr"`
	}

	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	var s TestStruct
	s.IntValue = 1
	s.Int8Value = 2
	s.Int16Value = 3
	s.Int32Value = 4
	s.Int64Value = 5
	s.UintValue = 6
	s.Uint8Value = 7
	s.Uint16Value = 8
	s.Uint32Value = 9
	s.Uint64Value = 10
	s.UintptrValue = 11
	s.BoolValue = true
	s.Float32Value = 12
	s.Float64Value = 13
	s.Complex64Value = 14
	s.Complex128Value = 15
	s.ArrayValue = [3]int{1, 2, 3}
	s.SliceValue = []int{4, 5, 6}
	s.EmptySliceValue = []int{}
	s.NilSliceValue = nil
	s.MapValue = map[int]int{1: 1}
	s.EmptyMapValue = map[int]int{}
	s.NilMapValue = nil
	s.StringValue = "hello"
	s.EmptyStringValue = ""
	s.InterfaceValue = s
	s.NilInterfaceValue = nil
	s.StructValue = &s
	s.IntPtrValue = &s.IntValue
	s.Int8PtrValue = &s.Int8Value
	s.Int16PtrValue = &s.Int16Value
	s.Int32PtrValue = &s.Int32Value
	s.Int64PtrValue = &s.Int64Value
	s.UintPtrValue = &s.UintValue
	s.Uint8PtrValue = &s.Uint8Value
	s.Uint16PtrValue = &s.Uint16Value
	s.Uint32PtrValue = &s.Uint32Value
	s.Uint64PtrValue = &s.Uint64Value
	s.UintptrPtrValue = &s.UintptrValue
	s.BoolPtrValue = &s.BoolValue
	s.Float32PtrValue = &s.Float32Value
	s.Float64PtrValue = &s.Float64Value
	s.Complex64PtrValue = &s.Complex64Value
	s.Complex128PtrValue = &s.Complex128Value
	s.ArrayPtrValue = &s.ArrayValue
	s.SlicePtrValue = &s.SliceValue
	s.EmptySlicePtrValue = &s.EmptySliceValue
	s.NilSlicePtrValue = &s.NilSliceValue
	s.MapPtrValue = &s.MapValue
	s.EmptyMapPtrValue = &s.EmptyMapValue
	s.NilMapPtrValue = &s.NilMapValue
	s.StringPtrValue = &s.StringValue
	s.EmptyStringPtrValue = &s.EmptyStringValue
	s.InterfacePtrValue = &s.InterfaceValue
	s.NilInterfacePtrValue = &s.NilInterfaceValue
	s.StructPtrValue = &s.StructValue

	assert.NoError(t, enc.Encode(s))
	assert.NoError(t, enc.Encode(&s))
	assert.NoError(t, enc.Encode(&s))
	assert.NoError(t, enc.Encode((*TestStruct)(nil)))
	assert.Equal(t, `c10"TestStruct"`+
		`85{s1"a"s4"json"s1"d"s4"test"s1"i"s2"i8"s3"i16"s3"i32"s3"i64"s1"u"s2"u8"s3"u16"s3"u32"s3"u64"s2"up"s1"`+
		`b"s3"f32"s3"f64"s3"c64"s4"c128"s4"iarr"s6"islice"s6"eslice"s6"nslice"s4"imap"s4"emap"s4"nmap"s1"s"s2"es"`+
		`s5"iface"s8"niliface"s2"st"s4"iptr"s5"i8ptr"s6"i16ptr"s6"i32ptr"s6"i64ptr"s4"uptr"s5"u8ptr"s6"u16ptr"`+
		`s6"u32ptr"s6"u64ptr"s5"upptr"s4"bptr"s6"f32ptr"s6"f64ptr"s6"c64ptr"s7"c128ptr"s7"iarrptr"s9"isliceptr"`+
		`s9"esliceptr"s9"nsliceptr"s7"imapptr"s7"emapptr"s7"nmapptr"s4"sptr"s5"esptr"s8"ifaceptr"s11"nilifaceptr"`+
		`s5"stptr"s5"niptr"s6"ni8ptr"s7"ni16ptr"s7"ni32ptr"s7"ni64ptr"s5"nuptr"s6"nu8ptr"s7"nu16ptr"s7"nu32ptr"`+
		`s7"nu64ptr"s6"nupptr"s5"nbptr"s7"nf32ptr"s7"nf64ptr"s7"nc64ptr"s8"nc128ptr"s8"niarrptr"s10"nesliceptr"`+
		`s10"nnsliceptr"s8"nemapptr"s8"nnmapptr"s5"nsptr"s6"nesptr"s12"nnilifaceptr"s6"nstptr"}`+
		`o0{0ed0;c15"TestEmbedStruct"3{s1"a"s4"json"s1"d"}o1{0ed0;}`+
		`123456789i10;i11;td12;d13;d14;d15;a3{123}a3{456}a{}nm1{11}m{}ns5"hello"eo0{0ed0;o1{0ed0;}`+
		`123456789i10;i11;td12;d13;d14;d15;a3{123}a3{456}a{}nm1{11}m{}nr95;`+
		`ennnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn}no0{0ed0;o1{0ed0;}`+
		`123456789i10;i11;td12;d13;d14;d15;a3{123}a3{456}a{}nm1{11}m{}nr95;eo0{0ed0;o1{0ed0;}`+
		`123456789i10;i11;td12;d13;d14;d15;a3{123}a3{456}a{}nm1{11}m{}nr95;`+
		`ennnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn}nr103;`+
		`123456789i10;i11;td12;d13;d14;d15;a3{123}a3{456}a{}nm1{11}m{}nr95;eo0{0ed0;o1{0ed0;}`+
		`123456789i10;i11;td12;d13;d14;d15;a3{123}a3{456}a{}nm1{11}m{}nr95;`+
		`ennnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn}nr103;`+
		`nnnnnnnnnnnnnnnnnnnnnnnnn}`+
		`123456789i10;i11;td12;d13;d14;d15;r117;r118;r119;nr120;r121;nr95;eo0{0ed0;o1{0ed0;}`+
		`123456789i10;i11;td12;d13;d14;d15;a3{123}a3{456}a{}nm1{11}m{}nr95;`+
		`ennnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn}nr103;`+
		`nnnnnnnnnnnnnnnnnnnnnnnnn}r103;r103;n`, sb.String())
	enc.Reset()
	sb.Reset()
}

func TestEncodeSomeStructField(t *testing.T) {
	type TestStruct2 struct {
		A *big.Int
		B *big.Float
		C *big.Rat
		D big.Int
		E big.Float
		F big.Rat
	}
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	var s TestStruct2
	s.A = big.NewInt(1)
	s.B = big.NewFloat(2)
	s.C = big.NewRat(3, 4)
	s.D = *s.A
	s.E = *s.B
	s.F = *s.C
	assert.NoError(t, enc.Encode(s))
	assert.NoError(t, enc.Encode(&s))
	assert.NoError(t, enc.Encode(&s))
	assert.Equal(t, `c11"TestStruct2"6{s1"a"s1"b"s1"c"s1"d"s1"e"s1"f"}`+
		`o0{l1;d2;s3"3/4"l1;d2;s3"3/4"}o0{l1;d2;s3"3/4"l1;d2;s3"3/4"}r9;`, sb.String())
	enc.Reset()
	sb.Reset()
}

func TestEncodeTimeStructField(t *testing.T) {
	type TestStruct3 struct {
		A time.Time
		B time.Duration
		C *time.Time
		D *time.Duration
		E *time.Time
		F *time.Duration
	}
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	var s TestStruct3
	s.A = time.Date(1980, 12, 1, 2, 3, 4, 5, time.Local)
	s.B = time.Duration(1000)
	s.C = &s.A
	s.D = &s.B
	assert.NoError(t, enc.Encode(s))
	assert.NoError(t, enc.Encode(&s))
	assert.NoError(t, enc.Encode(&s))
	assert.Equal(t, `c11"TestStruct3"6{s1"a"s1"b"s1"c"s1"d"s1"e"s1"f"}`+
		`o0{D19801201T020304.000000005;i1000;D19801201T020304.000000005;i1000;nn}`+
		`o0{D19801201T020304.000000005;i1000;r8;i1000;nn}r9;`, sb.String())
	enc.Reset()
	sb.Reset()
}

func TestEncodeErrorStructField(t *testing.T) {
	type TestStruct4 struct {
		A error
		B *error
		C **error
	}
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	var s TestStruct4
	s.A = errors.New("test error")
	s.B = &s.A
	s.C = nil
	assert.NoError(t, enc.Encode(s))
	assert.NoError(t, enc.Encode(&s))
	assert.NoError(t, enc.Encode(&s))
	assert.Equal(t, `c11"TestStruct4"3{s1"a"s1"b"s1"c"}`+
		`o0{Es10"test error"Es10"test error"n}`+
		`o0{Es10"test error"Es10"test error"n}r6;`, sb.String())
	enc.Reset()
	sb.Reset()
}

func TestAmbiguousFields(t *testing.T) {
	assert.PanicsWithValue(t, "hprose/encoding: ambiguous fields with the same name or alias: a", func() {
		type TestStruct struct {
			A int `hprose:"a"`
			B int `json:"a"`
		}
		sb := &strings.Builder{}
		enc := NewEncoder(sb).Simple(false)
		enc.Encode(TestStruct{})
	})
}

func TestInvalidStructName(t *testing.T) {
	assert.PanicsWithValue(t, "hprose/encoding: invalid UTF-8 in struct name", func() {
		type TestStruct struct {
			A int
			B int
		}
		newStructEncoder(reflect.TypeOf((*TestStruct)(nil)).Elem(), "\xFE")
		sb := &strings.Builder{}
		enc := NewEncoder(sb).Simple(false)
		enc.Encode(TestStruct{})
	})
}

func TestEmptyAnonymousStruct(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	var s struct{}
	assert.NoError(t, enc.Encode(s))
	assert.NoError(t, enc.Encode((*struct{})(nil)))
	assert.Equal(t, `m{}n`, sb.String())
}

func TestEncodeOnePtrFieldStruct(t *testing.T) {
	type TestEmbedStruct struct {
		A int
	}
	type TestStruct struct {
		B *TestEmbedStruct
	}
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	var s TestStruct
	i := 1
	s.B = &TestEmbedStruct{i}
	assert.NoError(t, enc.Encode(s))
	assert.NoError(t, enc.Encode(s))
	assert.NoError(t, enc.Encode(&s))
	assert.NoError(t, enc.Encode(&s))
	assert.Equal(t, `c10"TestStruct"1{s1"b"}o0{c15"TestEmbedStruct"1{s1"a"}o1{1}}o0{r3;}o0{r3;}r5;`, sb.String())

	enc.Reset()
	sb.Reset()

	var s2 struct {
		B *TestEmbedStruct
	}
	s2.B = &TestEmbedStruct{i}
	assert.NoError(t, enc.Encode(s2))
	assert.NoError(t, enc.Encode(s2))
	assert.NoError(t, enc.Encode(&s2))
	assert.NoError(t, enc.Encode(&s2))
	assert.Equal(t, `m1{ubc15"TestEmbedStruct"1{s1"a"}o0{1}}m1{ubr2;}m1{ubr2;}r4;`, sb.String())
}
