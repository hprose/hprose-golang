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
 * rpc/reader_pool.go                                     *
 *                                                        *
 * hprose reader pool for Go.                             *
 *                                                        *
 * LastModified: Dec 3, 2016                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package rpc

import (
	"sync"

	"github.com/hprose/hprose-golang/io"
)

// ReaderPool is a reader pool for hprose client & service
type ReaderPool struct {
	sync.Pool
}

// AcquireReader from pool.
func (pool *ReaderPool) AcquireReader(buf []byte) (reader *io.Reader) {
	reader = pool.Get().(*io.Reader)
	reader.Init(buf)
	return
}

// ReleaseReader to pool.
func (pool *ReaderPool) ReleaseReader(reader *io.Reader) {
	reader.Init(nil)
	reader.Reset()
	pool.Put(reader)
}

var defaultReaderPool = &ReaderPool{
	Pool: sync.Pool{
		New: func() interface{} { return new(io.Reader) },
	},
}

// AcquireReader from pool.
func AcquireReader(buf []byte) *io.Reader {
	return defaultReaderPool.AcquireReader(buf)
}

// ReleaseReader to pool.
func ReleaseReader(reader *io.Reader) {
	defaultReaderPool.ReleaseReader(reader)
}
