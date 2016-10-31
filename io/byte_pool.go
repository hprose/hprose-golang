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
 * io/byte_pool.go                                        *
 *                                                        *
 * byte pool for Go.                                      *
 *                                                        *
 * LastModified: Oct 25, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import (
	"sync"
)

const (
	poolNum = 20
	maxSize = 1 << (poolNum + 8)
)

var bytePool [poolNum]*sync.Pool

func pow2roundup(x int64) int64 {
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x |= x >> 32
	return x + 1
}

var debruijn = [...]int{
	0, 1, 28, 2, 29, 14, 24, 3,
	30, 22, 20, 15, 25, 17, 4, 8,
	31, 27, 13, 23, 21, 19, 16, 7,
	26, 12, 18, 6, 11, 5, 10, 9,
}

func log2(x int64) int {
	return debruijn[uint32(x*0x077CB531)>>27]
}

func init() {
	for i := uint(0); i < poolNum; i++ {
		bytePool[i] = &sync.Pool{
			New: func(n int) func() interface{} {
				return func() interface{} {
					return make([]byte, n)
				}
			}(1 << (i + 9)),
		}
	}
}

// AcquireBytes from pool.
func AcquireBytes(size int) []byte {
	if size < 1 {
		return nil
	}
	capacity := pow2roundup(int64(size))
	if capacity > maxSize {
		return make([]byte, size, capacity)
	}
	if capacity < 512 {
		capacity = 512
	}
	return bytePool[log2(capacity)-9].Get().([]byte)[:size]
}

// ReleaseBytes to pool.
func ReleaseBytes(bytes []byte) bool {
	capacity := int64(cap(bytes))
	if capacity < 512 || capacity > maxSize || capacity != pow2roundup(capacity) {
		return false
	}
	bytePool[log2(capacity)-9].Put(bytes)
	return true
}
