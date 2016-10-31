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
 * LastModified: Oct 24, 2016                             *
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

func (pool *ReaderPool) acquireReader(buf []byte) (reader *io.Reader) {
	reader = pool.Get().(*io.Reader)
	reader.Init(buf)
	return
}

func (pool *ReaderPool) releaseReader(reader *io.Reader) {
	reader.Init(nil)
	reader.Reset()
	pool.Put(reader)
}

var defaultReaderPool = &ReaderPool{
	Pool: sync.Pool{
		New: func() interface{} { return new(io.Reader) },
	},
}
