/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/bytes_reader.go                                 |
|                                                          |
| LastModified: Mar 21, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import "io"

// BytesReader is the interface that groups the basic Read and ReadByte methods.
type BytesReader interface {
	io.Reader
	io.ByteReader
}
