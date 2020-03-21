/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/bytes_writer.go                                       |
|                                                          |
| LastModified: Mar 21, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import "io"

// BytesWriter is the interface that groups the basic Write and WriteByte methods.
type BytesWriter interface {
	io.Writer
	io.ByteWriter
}
