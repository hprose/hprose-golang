/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * io/formatter.go                                        *
 *                                                        *
 * io Formatter for Go.                                   *
 *                                                        *
 * LastModified: Sep 10, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	var i int
	Unmarshal(Marshal(123), &i)
	if i != 123 {
		t.Error(i)
	}
}

func randString(l int) string {
	buf := make([]byte, l)
	for i := 0; i < (l+1)/2; i++ {
		buf[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("%x", buf)[:l]
}

func BenchmarkHproseMarshal(b *testing.B) {
	b.StopTimer()
	type A struct {
		Name     string
		BirthDay time.Time
		Phone    string
		Siblings int
		Spouse   bool
		Money    float64
	}
	var a interface{} = A{
		Name:     randString(16),
		BirthDay: time.Now(),
		Phone:    randString(10),
		Siblings: rand.Intn(5),
		Spouse:   rand.Intn(2) == 1,
		Money:    rand.Float64(),
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Marshal(a)
	}
}

func BenchmarkHproseUnmarshal(b *testing.B) {
	b.StopTimer()
	type A struct {
		Name     string
		BirthDay time.Time
		Phone    string
		Siblings int
		Spouse   bool
		Money    float64
	}
	var a interface{} = A{
		Name:     randString(16),
		BirthDay: time.Now(),
		Phone:    randString(10),
		Siblings: rand.Intn(5),
		Spouse:   rand.Intn(2) == 1,
		Money:    rand.Float64(),
	}
	buf := Marshal(a)
	var x A
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Unmarshal(buf, &x)
	}
}

func BenchmarkHproseSerialize(b *testing.B) {
	b.StopTimer()
	type A struct {
		Name     string
		BirthDay time.Time
		Phone    string
		Siblings int
		Spouse   bool
		Money    float64
	}
	var a interface{} = A{
		Name:     randString(16),
		BirthDay: time.Now(),
		Phone:    randString(10),
		Siblings: rand.Intn(5),
		Spouse:   rand.Intn(2) == 1,
		Money:    rand.Float64(),
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Serialize(a, false)
	}
}

func BenchmarkHproseUnserialize(b *testing.B) {
	b.StopTimer()
	type A struct {
		Name     string
		BirthDay time.Time
		Phone    string
		Siblings int
		Spouse   bool
		Money    float64
	}
	var a interface{} = A{
		Name:     randString(16),
		BirthDay: time.Now(),
		Phone:    randString(10),
		Siblings: rand.Intn(5),
		Spouse:   rand.Intn(2) == 1,
		Money:    rand.Float64(),
	}
	buf := Marshal(a)
	var x A
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Unserialize(buf, &x, false)
	}
}

func BenchmarkJSONMarshal(b *testing.B) {
	b.StopTimer()
	type A struct {
		Name     string
		BirthDay time.Time
		Phone    string
		Siblings int
		Spouse   bool
		Money    float64
	}
	var a interface{} = A{
		Name:     randString(16),
		BirthDay: time.Now(),
		Phone:    randString(10),
		Siblings: rand.Intn(5),
		Spouse:   rand.Intn(2) == 1,
		Money:    rand.Float64(),
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(a)
	}
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	b.StopTimer()
	type A struct {
		Name     string
		BirthDay time.Time
		Phone    string
		Siblings int
		Spouse   bool
		Money    float64
	}
	var a interface{} = A{
		Name:     randString(16),
		BirthDay: time.Now(),
		Phone:    randString(10),
		Siblings: rand.Intn(5),
		Spouse:   rand.Intn(2) == 1,
		Money:    rand.Float64(),
	}
	buf, _ := json.Marshal(a)
	var x A
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		json.Unmarshal(buf, &x)
	}
}
