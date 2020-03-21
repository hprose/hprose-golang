/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/list_encoder.go                                 |
|                                                          |
| LastModified: Mar 21, 2020                               |
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

func (valenc listEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	return EncodeReference(valenc, enc, v)
}

func (listEncoder) Write(enc *Encoder, v interface{}) (err error) {
	SetReference(enc, v)
	return writeList(enc, (*list.List)(reflect2.PtrOf(v)))
}

func writeList(enc *Encoder, lst *list.List) (err error) {
	count := lst.Len()
	if count == 0 {
		_, err = enc.writer.Write(emptySlice)
		return
	}
	err = WriteHead(enc, count, TagList)
	for e := lst.Front(); e != nil && err == nil; e = e.Next() {
		err = enc.Encode(e.Value)
	}
	if err == nil {
		err = WriteFoot(enc)
	}
	return
}

// elementEncoder is the implementation of ValueEncoder for list.Element/*list.Element.
type elementEncoder struct{}

func (valenc elementEncoder) Encode(enc *Encoder, v interface{}) (err error) {
	e := (*list.Element)(reflect2.PtrOf(v))
	if e == nil {
		return WriteNil(enc)
	}
	return enc.Encode(e.Value)
}

func (elementEncoder) Write(enc *Encoder, v interface{}) (err error) {
	return enc.Write((*list.Element)(reflect2.PtrOf(v)).Value)
}

func init() {
	RegisterValueEncoder((*list.List)(nil), listEncoder{})
	RegisterValueEncoder((*list.Element)(nil), elementEncoder{})
}
