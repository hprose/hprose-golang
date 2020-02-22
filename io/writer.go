/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/writer.go                                             |
|                                                          |
| LastModified: Feb 22, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import "io"

// Writer is the interface that groups the basic Write and WriteByte methods.
type Writer interface {
	io.Writer
	io.ByteWriter
}
