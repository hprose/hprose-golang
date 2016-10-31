/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * io/types.go                                            *
 *                                                        *
 * reflect types for Go.                                  *
 *                                                        *
 * LastModified: Oct 19, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import (
	"container/list"
	"math/big"
	"reflect"
	"time"
	"unsafe"
)

type emptyInterface struct {
	typ uintptr
	ptr uintptr
}

type reflectValue struct {
	typ  uintptr
	ptr  unsafe.Pointer
	flag uintptr
}

func getType(v interface{}) uintptr {
	return *(*uintptr)(unsafe.Pointer(&v))
}

var bytesType = getType(([]byte)(nil))
var bigIntType = getType(big.Int{})
var bigRatType = getType(big.Rat{})
var bigFloatType = getType(big.Float{})
var timeType = getType(time.Time{})
var listType = getType(list.List{})
var reflectValueType = getType(reflect.Value{})

var interfaceType = reflect.TypeOf((*interface{})(nil)).Elem()
