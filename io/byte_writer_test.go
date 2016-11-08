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
 * io/byte_writer_test.go                                 *
 *                                                        *
 * byte writer test for Go.                               *
 *                                                        *
 * LastModified: Nov 8, 2016                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/
package io

import (
	"reflect"
	"testing"
)

func TestByteWriter_Len(t *testing.T) {
	w := NewByteWriter(nil)
	if w.String() != "" {
		t.Error(w.String())
	}
	if w.Len() != 0 {
		t.Error(w.Len())
	}
	w = nil
	if w.String() != "<nil>" {
		t.Error(w.String())
	}
}

func TestByteWriter_Grow(t *testing.T) {
	w := NewByteWriter(nil)
	w.Grow(10)
	if w.Len() != 0 {
		t.Error(w.Len())
	}
	defer func() {
		if e := recover(); e == nil {
			t.Error("grow error")
		}
	}()
	w.Grow(-1)
}

func TestByteWriter_Write(t *testing.T) {
	w := NewByteWriter(nil)
	w.Write([]byte{1, 2, 3})
	if w.Len() != 3 {
		t.Error(w.Len())
	}
}

func TestByteWriter_WriteByte(t *testing.T) {
	w := NewByteWriter(nil)
	w.WriteByte(1)
	w.WriteByte(2)
	w.WriteByte(3)
	if w.Len() != 3 {
		t.Error(w.Len())
	}
}

func TestNewByteWriter(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name  string
		args  args
		wantW *ByteWriter
	}{}
	for _, tt := range tests {
		if gotW := NewByteWriter(tt.args.buf); !reflect.DeepEqual(gotW, tt.wantW) {
			t.Errorf("%q. NewByteWriter() = %v, want %v", tt.name, gotW, tt.wantW)
		}
	}
}
