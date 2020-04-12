/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/list_encoder.go                                 |
|                                                          |
| LastModified: Apr 12, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"container/list"

	"github.com/modern-go/reflect2"
)

// listEncoder is the implementation of ValueEncoder for list.List/*list.List.
type listEncoder struct{}

func (valenc listEncoder) Encode(enc *Encoder, v interface{}) {
	enc.EncodeReference(valenc, v)
}

func (listEncoder) Write(enc *Encoder, v interface{}) {
	enc.SetReference(v)
	writeList(enc, (*list.List)(reflect2.PtrOf(v)))
}

func writeList(enc *Encoder, lst *list.List) {
	count := lst.Len()
	if count == 0 {
		enc.buf = append(enc.buf, TagList, TagOpenbrace, TagClosebrace)
		return
	}
	enc.WriteHead(count, TagList)
	for e := lst.Front(); e != nil; e = e.Next() {
		enc.encode(e.Value)
	}
	enc.WriteFoot()
}

// elementEncoder is the implementation of ValueEncoder for list.Element/*list.Element.
type elementEncoder struct{}

func (valenc elementEncoder) Encode(enc *Encoder, v interface{}) {
	e := (*list.Element)(reflect2.PtrOf(v))
	if e == nil {
		WriteNil(enc)
	} else {
		enc.encode(e.Value)
	}
}

func (elementEncoder) Write(enc *Encoder, v interface{}) {
	enc.write((*list.Element)(reflect2.PtrOf(v)).Value)
}

func init() {
	RegisterValueEncoder((*list.List)(nil), listEncoder{})
	RegisterValueEncoder((*list.Element)(nil), elementEncoder{})
}
