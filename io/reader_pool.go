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
 * io/reader_pool.go                                      *
 *                                                        *
 * hprose reader pool for Go.                             *
 *                                                        *
 * LastModified: Dec 3, 2016                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import (
	"sync"
)

// ReaderPool is a reader pool for hprose client & service
type ReaderPool struct {
	sync.Pool
}

// AcquireReader from pool.
func (pool *ReaderPool) AcquireReader(buf []byte, simple bool) (reader *Reader) {
	reader = pool.Get().(*Reader)
	reader.Init(buf)
	reader.Simple = simple
	return
}

// ReleaseReader to pool.
func (pool *ReaderPool) ReleaseReader(reader *Reader) {
	reader.Init(nil)
	reader.Reset()
	pool.Put(reader)
}

var defaultReaderPool = &ReaderPool{
	Pool: sync.Pool{
		New: func() interface{} { return new(Reader) },
	},
}

// AcquireReader from pool.
func AcquireReader(buf []byte, simple bool) *Reader {
	return defaultReaderPool.AcquireReader(buf, simple)
}

// ReleaseReader to pool.
func ReleaseReader(reader *Reader) {
	defaultReaderPool.ReleaseReader(reader)
}
