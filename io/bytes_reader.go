/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/bytes_reader.go                                       |
|                                                          |
| LastModified: Mar 21, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import "io"

// BytesReader is the interface that groups the basic Read and ReadByte methods.
type BytesReader interface {
	io.Reader
	io.ByteReader
}
