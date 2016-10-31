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
 * io/byte_pool_test.go                                   *
 *                                                        *
 * byte pool test for Go.                                 *
 *                                                        *
 * LastModified: Oct 26, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import (
	"testing"
)

func TestBytesPool(t *testing.T) {
	bytes := AcquireBytes(0)
	if bytes != nil {
		t.Error("AcquireBytes(0) must return 0")
	}
	if ReleaseBytes(bytes) {
		t.Error("ReleaseBytes(nil) must return false")
	}
	for i := uint(0); i < 10; i++ {
		bytes := AcquireBytes(1 << i)
		if len(bytes) != 1<<i || cap(bytes) != 512 {
			t.Error(len(bytes), cap(bytes), 1<<i)
		}
		if !ReleaseBytes(bytes) {
			t.Error(len(bytes), cap(bytes), 1<<i)
		}
	}
	for i := uint(10); i < 29; i++ {
		bytes := AcquireBytes((1 << i) - 500)
		if len(bytes) != ((1<<i)-500) || cap(bytes) != 1<<i {
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
