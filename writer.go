package tester

import (
	"bytes"
	"errors"
)

var (
	// ErrMaxWritten tell that the written bytes reached maximum.
	ErrMaxWritten = errors.New("go-tester/tester: max writable byte reached")
)

// MaxSilentWriter is the alias for NewMaxWriter(n, nil).
func MaxSilentWriter(n int64) *MaxWriter {
	return NewMaxWriter(n, nil)
}

// MaxErrorWriter is the alias for NewMaxWriter(n, ErrMaxWritten).
func MaxErrorWriter(n int64) *MaxWriter {
	return NewMaxWriter(n, ErrMaxWritten)
}

// NewMaxWriter returns a new [MaxWriter] that accepts n bytes at maximum.
// When the written bytes reached n, the writer returns n and err.
func NewMaxWriter(n int64, err error) *MaxWriter {
	return &MaxWriter{
		buf:      bytes.NewBuffer(nil),
		err:      err,
		errAfter: max(0, n),
	}
}

// MaxWriter is a [io.Writer] that accepts n bytes at maximum.
// Use [NewMaxWriter] to create a new writer.
type MaxWriter struct {
	buf      *bytes.Buffer
	err      error
	errAfter int64
	written  int64
}

func (w *MaxWriter) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	left := w.errAfter - w.written
	read := min(int(left), len(p))
	w.written += int64(read)
	_, _ = w.buf.Write(p[:read])
	if w.written >= w.errAfter {
		return read, w.err
	}
	return read, nil
}

// Bytes returns written bytes.
func (w *MaxWriter) Bytes() []byte {
	return w.buf.Bytes()
}

// String returns written bytes in string.
func (w *MaxWriter) String() string {
	return w.buf.String()
}
