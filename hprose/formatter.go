/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.net/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * hprose/reader.go                                       *
 *                                                        *
 * hprose Reader for Go.                                  *
 *                                                        *
 * LastModified: Feb 8, 2014                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package hprose

import (
	"bytes"
	"io"
	"unicode/utf8"
)

type bufReader struct {
	s []byte
	i int
}

func NewBufReader(b []byte) BufReader {
	return &bufReader{b, 0}
}

func (r *bufReader) Read(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	if r.i >= len(r.s) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:])
	r.i += n
	return
}

func (r *bufReader) ReadByte() (b byte, err error) {
	if r.i >= len(r.s) {
		return 0, io.EOF
	}
	b = r.s[r.i]
	r.i++
	return
}

func (r *bufReader) ReadRune() (ch rune, size int, err error) {
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

func (r *bufReader) ReadString(delim byte) (line string, err error) {
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

func Serialize(v interface{}, simple bool) ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := NewWriter(buf, simple)
	err := writer.Serialize(v)
	return buf.Bytes(), err
}

func Marshal(v interface{}) ([]byte, error) {
	return Serialize(v, true)
}

func Unserialize(b []byte, p interface{}, simple bool) error {
	buf := NewBufReader(b)
	reader := NewReader(buf, simple)
	return reader.Unserialize(p)
}

func Unmarshal(b []byte, p interface{}) error {
	return Unserialize(b, p, true)
}
