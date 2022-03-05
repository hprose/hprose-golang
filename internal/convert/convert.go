package convert

import (
	"reflect"
	"unsafe"
)

//go:nosplit
func ToUnsafeString(v []byte) (s string) {
	(*reflect.StringHeader)(unsafe.Pointer(&s)).Len = (*reflect.SliceHeader)(unsafe.Pointer(&v)).Len
	(*reflect.StringHeader)(unsafe.Pointer(&s)).Data = (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data
	return
}

//go:nosplit
func ToUnsafeBytes(s string) (v []byte) {
	(*reflect.SliceHeader)(unsafe.Pointer(&v)).Cap = (*reflect.StringHeader)(unsafe.Pointer(&s)).Len
	(*reflect.SliceHeader)(unsafe.Pointer(&v)).Len = (*reflect.StringHeader)(unsafe.Pointer(&s)).Len
	(*reflect.SliceHeader)(unsafe.Pointer(&v)).Data = (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
	return
}
