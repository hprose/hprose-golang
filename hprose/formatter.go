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
 * hprose/formatter.go                                    *
 *                                                        *
 * hprose Formatter for Go.                               *
 *                                                        *
 * LastModified: May 22, 2015                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"bytes"
	"io"
	"unicode/utf8"
)

// BytesReader is a bytes reader
type BytesReader struct {
	s []byte
	i int
}

// NewBytesReader is the constructor of BytesReader
func NewBytesReader(b []byte) *BytesReader {
	return &BytesReader{b, 0}
}

// Read bytes from BytesReader
func (r *BytesReader) Read(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	if r.i >= len(r.s) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:])
	r.i += n
	return n, nil
}

// ReadByte from BytesReader
func (r *BytesReader) ReadByte() (b byte, err error) {
	if r.i >= len(r.s) {
		return 0, io.EOF
	}
	b = r.s[r.i]
	r.i++
	return
}

// ReadRune from BytesReader
func (r *BytesReader) ReadRune() (ch rune, size int, err error) {
	if r.i >= len(r.s) {
		return 0, 0, io.EOF
	}
	if c := r.s[r.i]; c < utf8.RuneSelf {
		r.i++
		return rune(c), 1, nil
	}
	ch, size = utf8.DecodeRune(r.s[r.i:])
	r.i += size
	return
}

// ReadString from BytesReader
func (r *BytesReader) ReadString(delim byte) (line string, err error) {
	i := bytes.IndexByte(r.s[r.i:], delim)
	end := r.i + i + 1
	if i < 0 {
		end = len(r.s)
		err = io.EOF
	}
	line = string(r.s[r.i:end])
	r.i = end
	return line, err
}

// Serialize data
func Serialize(v interface{}, simple bool) ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := NewWriter(buf, simple)
	err := writer.Serialize(v)
	return buf.Bytes(), err
}

// Marshal data
func Marshal(v interface{}) ([]byte, error) {
	return Serialize(v, true)
}

// Unserialize data
func Unserialize(b []byte, p interface{}, simple bool) error {
	buf := NewBytesReader(b)
	reader := NewReader(buf, simple)
	return reader.Unserialize(p)
}

// Unmarshal data
func Unmarshal(b []byte, p interface{}) error {
	return Unserialize(b, p, true)
}
