/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| internal/convert/convert.go                              |
|                                                          |
| LastModified: Mar 18, 2022                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package convert

import (
	"unsafe"
)

// stringHeader is a safe version of reflect.StringHeader used within this package.
type stringHeader struct {
	Data unsafe.Pointer
	Len  int
}

// sliceHeader is a safe version of reflect.SliceHeader used within this package.
type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

//go:nosplit
func ToUnsafeString(v []byte) (s string) {
	(*stringHeader)(unsafe.Pointer(&s)).Len = (*sliceHeader)(unsafe.Pointer(&v)).Len
	(*stringHeader)(unsafe.Pointer(&s)).Data = (*sliceHeader)(unsafe.Pointer(&v)).Data
	return
}

//go:nosplit
func ToUnsafeBytes(s string) (v []byte) {
	(*sliceHeader)(unsafe.Pointer(&v)).Cap = (*stringHeader)(unsafe.Pointer(&s)).Len
	(*sliceHeader)(unsafe.Pointer(&v)).Len = (*stringHeader)(unsafe.Pointer(&s)).Len
	(*sliceHeader)(unsafe.Pointer(&v)).Data = (*stringHeader)(unsafe.Pointer(&s)).Data
	return
}
