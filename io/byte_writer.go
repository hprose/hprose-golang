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
 * io/byte_writer.go                                      *
 *                                                        *
 * byte writer for Go.                                    *
 *                                                        *
 * LastModified: Oct 24, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

// ByteWriter implements the io.Writer and io.ByteWriter interfaces by writing
// to a byte slice
type ByteWriter struct {
	buf       []byte
	bootstrap [64]byte
}

// NewByteWriter create a ByteWriter in append mode
func NewByteWriter(buf []byte) (w *ByteWriter) {
	w = new(ByteWriter)
	w.buf = buf
	return w
}

// Len return the number of byte of this writer.
func (w *ByteWriter) Len() int {
	return len(w.buf)
}

// Bytes returns the byte slice of this writer.
func (w *ByteWriter) Bytes() []byte {
	return w.buf
}

// String returns the contents of this writer as a string.
// If the ByteWriter is a nil pointer, it returns "<nil>".
func (w *ByteWriter) String() string {
	if w == nil {
		// Special case, useful in debugging.
		return "<nil>"
	}
	return string(w.buf)
}

// Clear the byte slice of this writer.
func (w *ByteWriter) Clear() {
	w.buf = w.buf[:0]
}

func (w *ByteWriter) grow(n int) int {
	p := len(w.buf)
	c := cap(w.buf)
	l := p + n
	if l > c {
		var buf []byte
		if w.buf == nil && n <= len(w.bootstrap) {
			buf = w.bootstrap[0:]
		} else {
			buf = AcquireBytes(c + n)
			copy(buf, w.buf[0:])
			ReleaseBytes(w.buf)
		}
		w.buf = buf
	}
	w.buf = w.buf[:l]
	return p
}

// Grow the the byte slice capacity of this writer.
func (w *ByteWriter) Grow(n int) {
	if n < 0 {
		panic("BytesWriter: negative count")
	}
	p := w.grow(n)
	w.buf = w.buf[0:p]
}

// WriteByte c to the byte slice of this writer.
func (w *ByteWriter) WriteByte(c byte) error {
	p := w.grow(1)
	w.buf[p] = c
	return nil
}

// Write the contents of b to the byte slice of this writer.
func (w *ByteWriter) Write(b []byte) (int, error) {
	p := w.grow(len(b))
	return copy(w.buf[p:], b), nil
}

func (w *ByteWriter) writeByte(c byte) {
	p := w.grow(1)
	w.buf[p] = c
}

func (w *ByteWriter) write(b []byte) int {
	p := w.grow(len(b))
	return copy(w.buf[p:], b)
}

func (w *ByteWriter) writeString(s string) int {
	p := w.grow(len(s))
	return copy(w.buf[p:], s)
}
