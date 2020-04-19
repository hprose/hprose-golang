/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/decoder_refer.go                                |
|                                                          |
| LastModified: Apr 19, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

type decoderRefer struct {
	ref []interface{}
}

func (r *decoderRefer) Add(o interface{}) {
	r.ref = append(r.ref, o)
}

func (r *decoderRefer) Last() int {
	return len(r.ref) - 1
}

func (r *decoderRefer) Set(i int, o interface{}) {
	r.ref[i] = o
}

func (r *decoderRefer) Read(i int) interface{} {
	return r.ref[i]
}

func (r *decoderRefer) Reset() {
	r.ref = r.ref[:0]
}
