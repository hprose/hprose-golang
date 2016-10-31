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
 * io/ptr_decoder.go                                      *
 *                                                        *
 * hprose ptr decoder for Go.                             *
 *                                                        *
 * LastModified: Sep 10, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import "reflect"

func ptrDecoder(r *Reader, v reflect.Value, tag byte) {
	if tag == TagNull {
		if v.IsNil() {
			return
		}
		v.Set(reflect.Zero(v.Type()))
		return
	}
	if v.IsNil() {
		v.Set(reflect.New(v.Type().Elem()))
	}
	e := v.Elem()
	decoder := valueDecoders[e.Kind()]
	decoder(r, e, tag)
}
