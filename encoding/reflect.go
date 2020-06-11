/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/relect.go                                       |
|                                                          |
| LastModified: Apr 18, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"unsafe"
)

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

// sliceHeader is a safe version of SliceHeader used within this package.
type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

func setSliceHeader(slicePtr unsafe.Pointer, arrayPtr unsafe.Pointer, length int) {
	sliceHeader := (*sliceHeader)(slicePtr)
	sliceHeader.Data = arrayPtr
	sliceHeader.Len = length
	sliceHeader.Cap = length
}
