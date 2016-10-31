package bytewriter

import (
	"bytes"
	"strings"
	"testing"

	"github.com/hprose/hprose-golang/io"
)

func BenchmarkByteWriter(b *testing.B) {
	var buffer io.ByteWriter
	for n := 0; n < b.N; n++ {
		buffer.WriteByte('x')
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); buffer.String() != s {
		b.Errorf("unexpected result; got=%s, want=%s", buffer.String(), s)
	}
}

func BenchmarkBuffer(b *testing.B) {
	var buffer bytes.Buffer
	for n := 0; n < b.N; n++ {
		buffer.WriteByte('x')
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); buffer.String() != s {
		b.Errorf("unexpected result; got=%s, want=%s", buffer.String(), s)
	}
}
