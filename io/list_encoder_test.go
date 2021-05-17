/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| io/list_encoder_test.go                                  |
|                                                          |
| LastModified: Apr 27, 2021                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package io_test

import (
	"container/list"
	"strings"
	"testing"

	. "github.com/hprose/hprose-golang/v3/io"
	"github.com/stretchr/testify/assert"
)

func TestEncodeList(t *testing.T) {
	sb := &strings.Builder{}
	enc := NewEncoder(sb).Simple(false)
	lst := list.New()
	lst.PushBack(1)
	lst.PushBack(2)
	e := lst.PushBack(3)
	assert.NoError(t, enc.Encode(*e))
	assert.NoError(t, enc.Encode(*lst))
	assert.NoError(t, enc.Encode(e))
	assert.NoError(t, enc.Encode(lst))
	assert.NoError(t, enc.Encode(&e))
	assert.NoError(t, enc.Encode(&lst))
	assert.Equal(t, `3a3{123}3a3{123}3r1;`, sb.String())

	enc.Reset()
	sb.Reset()

	var ts struct {
		Plst *list.List
		Pe   *list.Element
		Nlst *list.List
		Ne   *list.Element
		Lst  list.List
		E    list.Element
		Elst *list.List
	}
	ts.Plst = lst
	ts.Pe = e
	ts.Lst = *lst
	ts.E = *e
	ts.Elst = list.New()
	assert.NoError(t, enc.Encode(ts))
	assert.NoError(t, enc.Encode(&ts))
	assert.NoError(t, enc.Encode(&ts))
	assert.Equal(t, `m7{s4"plst"a3{123}s2"pe"3s4"nlst"ns2"ne"ns3"lst"a3{123}ue3s4"elst"a{}}`+
		`m7{r1;r2;r3;3r4;nr5;nr6;a3{123}ue3r8;r9;}r10;`, sb.String())
}
