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
 * LastModified: Oct 13, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import "sync"

var readerPool = sync.Pool{
	New: func() interface{} { return new(Reader) },
}

func acquireReader(buf []byte, simple bool) (reader *Reader) {
	reader = readerPool.Get().(*Reader)
	reader.Init(buf)
	reader.Simple = simple
	return
}

func releaseReader(reader *Reader) {
	reader.Init(nil)
	reader.Reset()
	readerPool.Put(reader)
}

// Serialize data
func Serialize(v interface{}, simple bool) []byte {
	return NewWriter(simple).Serialize(v).Bytes()
}

// Marshal data
func Marshal(v interface{}) []byte {
	return Serialize(v, true)
}

// Unserialize data
func Unserialize(b []byte, p interface{}, simple bool) {
	reader := acquireReader(b, simple)
	defer releaseReader(reader)
	reader.Unserialize(p)
}

// Unmarshal data
func Unmarshal(b []byte, p interface{}) {
	Unserialize(b, p, true)
}
