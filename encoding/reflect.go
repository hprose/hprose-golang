/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/relect.go                                       |
|                                                          |
| LastModified: Apr 12, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import "unsafe"

type eface struct {
	typ uintptr
	ptr unsafe.Pointer
}

func unpackEFace(ptr *interface{}) *eface {
	return (*eface)(unsafe.Pointer(ptr))
}

func unsafeString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}
