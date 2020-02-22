/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/byte_pool_test.go                                     |
|                                                          |
| LastModified: Feb 22, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import "testing"

func TestBytesPool(t *testing.T) {
	bytes := AcquireBytes(0)
	if bytes != nil {
		t.Error("AcquireBytes(0) must return 0")
	}
	if ReleaseBytes(bytes) {
		t.Error("ReleaseBytes(nil) must return false")
	}
	for i := uint(0); i < 9; i++ {
		bytes := AcquireBytes(1 << i)
		if len(bytes) != 1<<i || cap(bytes) != 1<<i {
			t.Error(len(bytes), cap(bytes), 1<<i)
		}
		if ReleaseBytes(bytes) {
			t.Error(len(bytes), cap(bytes), 1<<i)
		}
	}
	for i := uint(9); i < 29; i++ {
		bytes := AcquireBytes((1 << i))
		if len(bytes) != 1<<i || cap(bytes) != 1<<i {
			t.Error(len(bytes), cap(bytes), 1<<i)
		}
		if !ReleaseBytes(bytes) {
			t.Error(len(bytes), cap(bytes), 1<<i)
		}
	}
	for i := uint(29); i < 32; i++ {
		bytes := AcquireBytes((1 << i) - 500)
		if len(bytes) != (1<<i-500) || cap(bytes) != (1<<i) {
			t.Error(len(bytes), cap(bytes), 1<<i)
		}
		if ReleaseBytes(bytes) {
			t.Error(len(bytes), cap(bytes), 1<<i)
		}
	}
}

func BenchmarkBytesPool(b *testing.B) {
	n := 1 << 10
	for i := 0; i < b.N; i++ {
		bytes := AcquireBytes(n)
		bytes[n-1] = byte(i)
		ReleaseBytes(bytes)
	}
}

func BenchmarkMakeBytes(b *testing.B) {
	n := 1 << 10
	for i := 0; i < b.N; i++ {
		bytes := make([]byte, n)
		bytes[n-1] = byte(i)
	}
}
