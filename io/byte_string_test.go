/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/byte_string_test.go                                   |
|                                                          |
| LastModified: Feb 22, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"testing"
)

func TestBytesToString(t *testing.T) {
	bytes := []byte{'H', 'e', 'l', 'l', 'o'}
	if BytesToString(bytes) != "Hello" {
		t.Error(`BytesToString(bytes) must return "Hello"`)
	}
}

func TestStringToBytes(t *testing.T) {
	s := "Hello"
	if string(StringToBytes(s)) != "Hello" {
		t.Error(`string(StringToBytes(s)) must return "Hello"`)
	}
}
