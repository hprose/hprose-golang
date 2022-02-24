/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/formatter_test.go                                     |
|                                                          |
| LastModified: Feb 24, 2022                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io_test

import (
	"bytes"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/hprose/hprose-golang/v3/io"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Name     string
	Age      int
	Birthday time.Time
	Male     bool
}

func makeTestStructs() (testStructs []testStruct) {
	for i := 0; i < 1000; i++ {
		testStructs = append(testStructs, testStruct{
			Name:     "Name" + strconv.Itoa(i),
			Age:      i,
			Birthday: time.Now(),
			Male:     i%2 == 1,
		})

	}
	return
}

func TestHproseSimpleFormatter(t *testing.T) {
	wg := new(sync.WaitGroup)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s := makeTestStructs()
			data, err := io.Marshal(s)
			if assert.NoError(t, err) {
				io.Unmarshal(data, &s)
				data2, err := io.Marshal(s)
				if assert.NoError(t, err) {
					assert.Equal(t, string(data), string(data2))
				}
				io.UnmarshalFromReader(bytes.NewReader(data2), &s)
				data3, err := io.Marshal(s)
				if assert.NoError(t, err) {
					assert.Equal(t, string(data2), string(data3))
				}
			}
		}()
	}
	wg.Wait()
}

func TestHproseFormatter(t *testing.T) {
	formatter := io.Formatter{}
	wg := new(sync.WaitGroup)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s := makeTestStructs()
			data, err := formatter.Marshal(s)
			if assert.NoError(t, err) {
				formatter.Unmarshal(data, &s)
				data2, err := formatter.Marshal(s)
				if assert.NoError(t, err) {
					assert.Equal(t, string(data), string(data2))
				}
				formatter.UnmarshalFromReader(bytes.NewReader(data2), &s)
				data3, err := formatter.Marshal(s)
				if assert.NoError(t, err) {
					assert.Equal(t, string(data2), string(data3))
				}
			}
		}()
	}
	wg.Wait()
}

func BenchmarkJsoniterMarshal(b *testing.B) {
	s := makeTestStructs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jsoniter.Marshal(s)
	}
}

func BenchmarkHproseSimpleMarshal(b *testing.B) {
	s := makeTestStructs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		io.Marshal(s)
	}
}

func BenchmarkHproseMarshal(b *testing.B) {
	s := makeTestStructs()
	formatter := io.Formatter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		formatter.Marshal(s)
	}
}

func BenchmarkJsoniterUnmarshal(b *testing.B) {
	s := makeTestStructs()
	data, _ := jsoniter.Marshal(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jsoniter.Unmarshal(data, &s)
	}
}

func BenchmarkHproseSimpleUnmarshal(b *testing.B) {
	s := makeTestStructs()
	data, _ := io.Marshal(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		io.Unmarshal(data, &s)
	}
}

func BenchmarkHproseSimpleUnmarshalFromReader(b *testing.B) {
	s := makeTestStructs()
	data, _ := io.Marshal(s)
	reader := bytes.NewReader(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		io.UnmarshalFromReader(reader, &s)
		reader.Reset(data)
	}
}

func BenchmarkHproseUnmarshal(b *testing.B) {
	s := makeTestStructs()
	formatter := io.Formatter{}
	data, _ := formatter.Marshal(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		formatter.Unmarshal(data, &s)
	}
}

func BenchmarkHproseUnmarshalFromReader(b *testing.B) {
	s := makeTestStructs()
	formatter := io.Formatter{}
	data, _ := formatter.Marshal(s)
	reader := bytes.NewReader(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		formatter.UnmarshalFromReader(reader, &s)
		reader.Reset(data)
	}
}
