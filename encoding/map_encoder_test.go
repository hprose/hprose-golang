/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/map_encoder_test.go                             |
|                                                          |
| LastModified: May 9, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding_test

import (
	"math/big"
	"strings"
	"testing"

	. "github.com/hprose/hprose-golang/v3/encoding"
	"github.com/stretchr/testify/assert"
)

func TestEncodeMap(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	assert.NoError(t, enc.Encode(map[string]string{"stringstring": "string"}))
	assert.Equal(t, `m1{s12"stringstring"s6"string"}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[string]int{"stringint": 1}))
	assert.Equal(t, `m1{s9"stringint"1}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[string]int8{"stringint8": 2}))
	assert.Equal(t, `m1{s10"stringint8"2}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[string]int16{"stringint16": 3}))
	assert.Equal(t, `m1{s11"stringint16"3}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[string]int32{"stringint32": 4}))
	assert.Equal(t, `m1{s11"stringint32"4}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[string]int64{"stringint64": 5}))
	assert.Equal(t, `m1{s11"stringint64"l5;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[string]uint{"stringuint": 6}))
	assert.Equal(t, `m1{s10"stringuint"6}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[string]uint8{"stringuint8": 7}))
	assert.Equal(t, `m1{s11"stringuint8"7}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[string]uint16{"stringuint16": 8}))
	assert.Equal(t, `m1{s12"stringuint16"8}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[string]uint32{"stringuint32": 9}))
	assert.Equal(t, `m1{s12"stringuint32"9}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[string]uint64{"stringuint64": 10}))
	assert.Equal(t, `m1{s12"stringuint64"l10;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[string]bool{"stringbool": true}))
	assert.Equal(t, `m1{s10"stringbool"t}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[string]float32{"stringfloat32": 3.14159}))
	assert.Equal(t, `m1{s13"stringfloat32"d3.14159;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[string]float64{"stringfloat64": 2.71828}))
	assert.Equal(t, `m1{s13"stringfloat64"d2.71828;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[string]interface{}{"stringinterface": big.NewInt(0)}))
	assert.Equal(t, `m1{s15"stringinterface"l0;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int]string{1: "string"}))
	assert.Equal(t, `m1{1s6"string"}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int]int{2: 1}))
	assert.Equal(t, `m1{21}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int]int8{3: 2}))
	assert.Equal(t, `m1{32}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int]int16{4: 3}))
	assert.Equal(t, `m1{43}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int]int32{5: 4}))
	assert.Equal(t, `m1{54}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int]int64{6: 5}))
	assert.Equal(t, `m1{6l5;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int]uint{7: 6}))
	assert.Equal(t, `m1{76}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int]uint8{8: 7}))
	assert.Equal(t, `m1{87}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int]uint16{9: 8}))
	assert.Equal(t, `m1{98}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int]uint32{10: 9}))
	assert.Equal(t, `m1{i10;9}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int]uint64{11: 10}))
	assert.Equal(t, `m1{i11;l10;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int]bool{12: true}))
	assert.Equal(t, `m1{i12;t}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int]float32{13: 3.14159}))
	assert.Equal(t, `m1{i13;d3.14159;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int]float64{14: 2.71828}))
	assert.Equal(t, `m1{i14;d2.71828;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int]interface{}{15: big.NewInt(0)}))
	assert.Equal(t, `m1{i15;l0;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int8]string{1: "string"}))
	assert.Equal(t, `m1{1s6"string"}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int8]int{2: 1}))
	assert.Equal(t, `m1{21}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int8]int8{3: 2}))
	assert.Equal(t, `m1{32}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int8]int16{4: 3}))
	assert.Equal(t, `m1{43}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int8]int32{5: 4}))
	assert.Equal(t, `m1{54}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int8]int64{6: 5}))
	assert.Equal(t, `m1{6l5;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int8]uint{7: 6}))
	assert.Equal(t, `m1{76}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int8]uint8{8: 7}))
	assert.Equal(t, `m1{87}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int8]uint16{9: 8}))
	assert.Equal(t, `m1{98}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int8]uint32{10: 9}))
	assert.Equal(t, `m1{i10;9}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int8]uint64{11: 10}))
	assert.Equal(t, `m1{i11;l10;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int8]bool{12: true}))
	assert.Equal(t, `m1{i12;t}`, sb.String())
	enc.Reset()
	sb.Reset()
	assert.NoError(t, enc.Encode(map[int8]float32{13: 3.14159}))
	assert.Equal(t, `m1{i13;d3.14159;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int8]float64{14: 2.71828}))
	assert.Equal(t, `m1{i14;d2.71828;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int8]interface{}{15: big.NewInt(0)}))
	assert.Equal(t, `m1{i15;l0;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int16]string{1: "string"}))
	assert.Equal(t, `m1{1s6"string"}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int16]int{2: 1}))
	assert.Equal(t, `m1{21}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int16]int8{3: 2}))
	assert.Equal(t, `m1{32}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int16]int16{4: 3}))
	assert.Equal(t, `m1{43}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int16]int32{5: 4}))
	assert.Equal(t, `m1{54}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int16]int64{6: 5}))
	assert.Equal(t, `m1{6l5;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int16]uint{7: 6}))
	assert.Equal(t, `m1{76}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int16]uint8{8: 7}))
	assert.Equal(t, `m1{87}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int16]uint16{9: 8}))
	assert.Equal(t, `m1{98}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int16]uint32{10: 9}))
	assert.Equal(t, `m1{i10;9}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int16]uint64{11: 10}))
	assert.Equal(t, `m1{i11;l10;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int16]bool{12: true}))
	assert.Equal(t, `m1{i12;t}`, sb.String())
	enc.Reset()
	sb.Reset()
	assert.NoError(t, enc.Encode(map[int16]float32{13: 3.14159}))
	assert.Equal(t, `m1{i13;d3.14159;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int16]float64{14: 2.71828}))
	assert.Equal(t, `m1{i14;d2.71828;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int16]interface{}{15: big.NewInt(0)}))
	assert.Equal(t, `m1{i15;l0;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int32]string{1: "string"}))
	assert.Equal(t, `m1{1s6"string"}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int32]int{2: 1}))
	assert.Equal(t, `m1{21}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int32]int8{3: 2}))
	assert.Equal(t, `m1{32}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int32]int16{4: 3}))
	assert.Equal(t, `m1{43}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int32]int32{5: 4}))
	assert.Equal(t, `m1{54}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int32]int64{6: 5}))
	assert.Equal(t, `m1{6l5;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int32]uint{7: 6}))
	assert.Equal(t, `m1{76}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int32]uint8{8: 7}))
	assert.Equal(t, `m1{87}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int32]uint16{9: 8}))
	assert.Equal(t, `m1{98}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int32]uint32{10: 9}))
	assert.Equal(t, `m1{i10;9}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int32]uint64{11: 10}))
	assert.Equal(t, `m1{i11;l10;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int32]bool{12: true}))
	assert.Equal(t, `m1{i12;t}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int32]float32{13: 3.14159}))
	assert.Equal(t, `m1{i13;d3.14159;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int32]float64{14: 2.71828}))
	assert.Equal(t, `m1{i14;d2.71828;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int32]interface{}{15: big.NewInt(0)}))
	assert.Equal(t, `m1{i15;l0;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int64]string{1: "string"}))
	assert.Equal(t, `m1{l1;s6"string"}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int64]int{2: 1}))
	assert.Equal(t, `m1{l2;1}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int64]int8{3: 2}))
	assert.Equal(t, `m1{l3;2}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int64]int16{4: 3}))
	assert.Equal(t, `m1{l4;3}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int64]int32{5: 4}))
	assert.Equal(t, `m1{l5;4}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int64]int64{6: 5}))
	assert.Equal(t, `m1{l6;l5;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int64]uint{7: 6}))
	assert.Equal(t, `m1{l7;6}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int64]uint8{8: 7}))
	assert.Equal(t, `m1{l8;7}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int64]uint16{9: 8}))
	assert.Equal(t, `m1{l9;8}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int64]uint32{10: 9}))
	assert.Equal(t, `m1{l10;9}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int64]uint64{11: 10}))
	assert.Equal(t, `m1{l11;l10;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int64]bool{12: true}))
	assert.Equal(t, `m1{l12;t}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int64]float32{13: 3.14159}))
	assert.Equal(t, `m1{l13;d3.14159;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int64]float64{14: 2.71828}))
	assert.Equal(t, `m1{l14;d2.71828;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[int64]interface{}{15: big.NewInt(0)}))
	assert.Equal(t, `m1{l15;l0;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint]string{1: "string"}))
	assert.Equal(t, `m1{1s6"string"}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint]int{2: 1}))
	assert.Equal(t, `m1{21}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint]int8{3: 2}))
	assert.Equal(t, `m1{32}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint]int16{4: 3}))
	assert.Equal(t, `m1{43}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint]int32{5: 4}))
	assert.Equal(t, `m1{54}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint]int64{6: 5}))
	assert.Equal(t, `m1{6l5;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint]uint{7: 6}))
	assert.Equal(t, `m1{76}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint]uint8{8: 7}))
	assert.Equal(t, `m1{87}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint]uint16{9: 8}))
	assert.Equal(t, `m1{98}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint]uint32{10: 9}))
	assert.Equal(t, `m1{i10;9}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint]uint64{11: 10}))
	assert.Equal(t, `m1{i11;l10;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint]bool{12: true}))
	assert.Equal(t, `m1{i12;t}`, sb.String())
	enc.Reset()
	sb.Reset()
	assert.NoError(t, enc.Encode(map[uint]float32{13: 3.14159}))
	assert.Equal(t, `m1{i13;d3.14159;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint]float64{14: 2.71828}))
	assert.Equal(t, `m1{i14;d2.71828;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint]interface{}{15: big.NewInt(0)}))
	assert.Equal(t, `m1{i15;l0;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint8]string{1: "string"}))
	assert.Equal(t, `m1{1s6"string"}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint8]int{2: 1}))
	assert.Equal(t, `m1{21}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint8]int8{3: 2}))
	assert.Equal(t, `m1{32}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint8]int16{4: 3}))
	assert.Equal(t, `m1{43}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint8]int32{5: 4}))
	assert.Equal(t, `m1{54}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint8]int64{6: 5}))
	assert.Equal(t, `m1{6l5;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint8]uint{7: 6}))
	assert.Equal(t, `m1{76}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint8]uint8{8: 7}))
	assert.Equal(t, `m1{87}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint8]uint16{9: 8}))
	assert.Equal(t, `m1{98}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint8]uint32{10: 9}))
	assert.Equal(t, `m1{i10;9}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint8]uint64{11: 10}))
	assert.Equal(t, `m1{i11;l10;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint8]bool{12: true}))
	assert.Equal(t, `m1{i12;t}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint8]float32{13: 3.14159}))
	assert.Equal(t, `m1{i13;d3.14159;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint8]float64{14: 2.71828}))
	assert.Equal(t, `m1{i14;d2.71828;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint8]interface{}{15: big.NewInt(0)}))
	assert.Equal(t, `m1{i15;l0;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint16]string{1: "string"}))
	assert.Equal(t, `m1{1s6"string"}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint16]int{2: 1}))
	assert.Equal(t, `m1{21}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint16]int8{3: 2}))
	assert.Equal(t, `m1{32}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint16]int16{4: 3}))
	assert.Equal(t, `m1{43}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint16]int32{5: 4}))
	assert.Equal(t, `m1{54}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint16]int64{6: 5}))
	assert.Equal(t, `m1{6l5;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint16]uint{7: 6}))
	assert.Equal(t, `m1{76}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint16]uint8{8: 7}))
	assert.Equal(t, `m1{87}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint16]uint16{9: 8}))
	assert.Equal(t, `m1{98}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint16]uint32{10: 9}))
	assert.Equal(t, `m1{i10;9}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint16]uint64{11: 10}))
	assert.Equal(t, `m1{i11;l10;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint16]bool{12: true}))
	assert.Equal(t, `m1{i12;t}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint16]float32{13: 3.14159}))
	assert.Equal(t, `m1{i13;d3.14159;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint16]float64{14: 2.71828}))
	assert.Equal(t, `m1{i14;d2.71828;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint16]interface{}{15: big.NewInt(0)}))
	assert.Equal(t, `m1{i15;l0;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint32]string{1: "string"}))
	assert.Equal(t, `m1{1s6"string"}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint32]int{2: 1}))
	assert.Equal(t, `m1{21}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint32]int8{3: 2}))
	assert.Equal(t, `m1{32}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint32]int16{4: 3}))
	assert.Equal(t, `m1{43}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint32]int32{5: 4}))
	assert.Equal(t, `m1{54}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint32]int64{6: 5}))
	assert.Equal(t, `m1{6l5;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint32]uint{7: 6}))
	assert.Equal(t, `m1{76}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint32]uint8{8: 7}))
	assert.Equal(t, `m1{87}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint32]uint16{9: 8}))
	assert.Equal(t, `m1{98}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint32]uint32{10: 9}))
	assert.Equal(t, `m1{i10;9}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint32]uint64{11: 10}))
	assert.Equal(t, `m1{i11;l10;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint32]bool{12: true}))
	assert.Equal(t, `m1{i12;t}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint32]float32{13: 3.14159}))
	assert.Equal(t, `m1{i13;d3.14159;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint32]float64{14: 2.71828}))
	assert.Equal(t, `m1{i14;d2.71828;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint32]interface{}{15: big.NewInt(0)}))
	assert.Equal(t, `m1{i15;l0;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint64]string{1: "string"}))
	assert.Equal(t, `m1{l1;s6"string"}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint64]int{2: 1}))
	assert.Equal(t, `m1{l2;1}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint64]int8{3: 2}))
	assert.Equal(t, `m1{l3;2}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint64]int16{4: 3}))
	assert.Equal(t, `m1{l4;3}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint64]int32{5: 4}))
	assert.Equal(t, `m1{l5;4}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint64]int64{6: 5}))
	assert.Equal(t, `m1{l6;l5;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint64]uint{7: 6}))
	assert.Equal(t, `m1{l7;6}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint64]uint8{8: 7}))
	assert.Equal(t, `m1{l8;7}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint64]uint16{9: 8}))
	assert.Equal(t, `m1{l9;8}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint64]uint32{10: 9}))
	assert.Equal(t, `m1{l10;9}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint64]uint64{11: 10}))
	assert.Equal(t, `m1{l11;l10;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint64]bool{12: true}))
	assert.Equal(t, `m1{l12;t}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint64]float32{13: 3.14159}))
	assert.Equal(t, `m1{l13;d3.14159;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint64]float64{14: 2.71828}))
	assert.Equal(t, `m1{l14;d2.71828;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[uint64]interface{}{15: big.NewInt(0)}))
	assert.Equal(t, `m1{l15;l0;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float32]string{1: "string"}))
	assert.Equal(t, `m1{d1;s6"string"}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float32]int{2: 1}))
	assert.Equal(t, `m1{d2;1}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float32]int8{3: 2}))
	assert.Equal(t, `m1{d3;2}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float32]int16{4: 3}))
	assert.Equal(t, `m1{d4;3}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float32]int32{5: 4}))
	assert.Equal(t, `m1{d5;4}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float32]int64{6: 5}))
	assert.Equal(t, `m1{d6;l5;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float32]uint{7: 6}))
	assert.Equal(t, `m1{d7;6}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float32]uint8{8: 7}))
	assert.Equal(t, `m1{d8;7}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float32]uint16{9: 8}))
	assert.Equal(t, `m1{d9;8}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float32]uint32{10: 9}))
	assert.Equal(t, `m1{d10;9}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float32]uint64{11: 10}))
	assert.Equal(t, `m1{d11;l10;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float32]bool{12: true}))
	assert.Equal(t, `m1{d12;t}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float32]float32{13: 3.14159}))
	assert.Equal(t, `m1{d13;d3.14159;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float32]float64{14: 2.71828}))
	assert.Equal(t, `m1{d14;d2.71828;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float32]interface{}{15: big.NewInt(0)}))
	assert.Equal(t, `m1{d15;l0;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float64]string{1: "string"}))
	assert.Equal(t, `m1{d1;s6"string"}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float64]int{2: 1}))
	assert.Equal(t, `m1{d2;1}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float64]int8{3: 2}))
	assert.Equal(t, `m1{d3;2}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float64]int16{4: 3}))
	assert.Equal(t, `m1{d4;3}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float64]int32{5: 4}))
	assert.Equal(t, `m1{d5;4}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float64]int64{6: 5}))
	assert.Equal(t, `m1{d6;l5;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float64]uint{7: 6}))
	assert.Equal(t, `m1{d7;6}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float64]uint8{8: 7}))
	assert.Equal(t, `m1{d8;7}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float64]uint16{9: 8}))
	assert.Equal(t, `m1{d9;8}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float64]uint32{10: 9}))
	assert.Equal(t, `m1{d10;9}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float64]uint64{11: 10}))
	assert.Equal(t, `m1{d11;l10;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float64]bool{12: true}))
	assert.Equal(t, `m1{d12;t}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float64]float32{13: 3.14159}))
	assert.Equal(t, `m1{d13;d3.14159;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float64]float64{14: 2.71828}))
	assert.Equal(t, `m1{d14;d2.71828;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[float64]interface{}{15: big.NewInt(0)}))
	assert.Equal(t, `m1{d15;l0;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[interface{}]string{1: "string"}))
	assert.Equal(t, `m1{1s6"string"}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[interface{}]int{2: 1}))
	assert.Equal(t, `m1{21}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[interface{}]int8{3: 2}))
	assert.Equal(t, `m1{32}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[interface{}]int16{4: 3}))
	assert.Equal(t, `m1{43}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[interface{}]int32{5: 4}))
	assert.Equal(t, `m1{54}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[interface{}]int64{6: 5}))
	assert.Equal(t, `m1{6l5;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[interface{}]uint{7: 6}))
	assert.Equal(t, `m1{76}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[interface{}]uint8{8: 7}))
	assert.Equal(t, `m1{87}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[interface{}]uint16{9: 8}))
	assert.Equal(t, `m1{98}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[interface{}]uint32{10: 9}))
	assert.Equal(t, `m1{i10;9}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[interface{}]uint64{11: 10}))
	assert.Equal(t, `m1{i11;l10;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[interface{}]bool{12: true}))
	assert.Equal(t, `m1{i12;t}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[interface{}]float32{13: 3.14159}))
	assert.Equal(t, `m1{i13;d3.14159;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[interface{}]float64{14: 2.71828}))
	assert.Equal(t, `m1{i14;d2.71828;}`, sb.String())
	enc.Reset()
	sb.Reset()

	assert.NoError(t, enc.Encode(map[interface{}]interface{}{15: big.NewInt(0)}))
	assert.Equal(t, `m1{i15;l0;}`, sb.String())
	enc.Reset()
	sb.Reset()
}
