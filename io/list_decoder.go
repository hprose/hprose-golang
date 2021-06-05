/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/list_decoder.go                                       |
|                                                          |
| LastModified: Jun 5, 2021                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io

import (
	"container/list"

	"github.com/modern-go/reflect2"
)

// listDecoder is the implementation of ValueDecoder for *list.List.
type listDecoder struct{}

func (valdec listDecoder) Decode(dec *Decoder, p interface{}, tag byte) {
	plist := (**list.List)(reflect2.PtrOf(p))
	switch tag {
	case TagNull:
		*plist = nil
	case TagEmpty:
		*plist = list.New()
	case TagList:
		count := dec.ReadInt()
		l := list.New()
		*plist = l
		dec.AddReference(l)
		for i := 0; i < count; i++ {
			l.PushBack(dec.decodeInterface(dec.NextByte()))
		}
		dec.Skip()
	default:
		dec.defaultDecode(listType, p, tag)
	}
}

func init() {
	registerValueDecoder(listType, listDecoder{})
}
