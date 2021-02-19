/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| rpc/core/reflect.go                                      |
|                                                          |
| LastModified: Feb 20, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package core

import (
	"reflect"
	"unsafe"
)

func setAccessible(f reflect.Value) reflect.Value {
	if f.CanInterface() {
		return f
	}
	return reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem()
}
