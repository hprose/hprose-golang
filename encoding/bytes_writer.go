/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/bytes_writer.go                                 |
|                                                          |
| LastModified: Mar 21, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import "io"

// bytesWriter is the interface that groups the basic Write and WriteByte methods.
type bytesWriter interface {
	io.Writer
	io.ByteWriter
}
