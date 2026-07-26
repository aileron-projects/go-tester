package tester

import (
	"bytes"
	"errors"
	"io"
)

var (
	// ErrMaxRead tell that the read bytes reached maximum.
	ErrMaxRead = errors.New("go-tester/tester: max readable byte reached")
)

// MaxSilentReader is the alias for NewMaxReader(r, n, nil).
func MaxSilentReader(r io.Reader, n int64) *MaxReader {
	return NewMaxReader(r, n, nil)
}

// MaxErrorReader is the alias for NewMaxReader(r, n, ErrMaxRead).
func MaxErrorReader(r io.Reader, n int64) *MaxReader {
	return NewMaxReader(r, n, ErrMaxRead)
}

// NewMaxReader returns a new [MaxReader] that read n bytes at maximum.
// When the read bytes reached n, the reader returns n and err.
func NewMaxReader(r io.Reader, n int64, err error) *MaxReader {
	return &MaxReader{
		reader:   r,
		buf:      bytes.NewBuffer(nil),
		err:      err,
		errAfter: max(0, n),
	}
}

// MaxReader is a [io.Reader] that can read n bytes at maximum.
// Use [NewMaxReader] to create a new reader.
type MaxReader struct {
	reader   io.Reader
	buf      *bytes.Buffer
	err      error
	errAfter int64
	read     int64
}

func (r *MaxReader) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	left := r.errAfter - r.read
	readable := min(len(p), int(left))
	n, err = r.reader.Read(p[:readable])
	r.read += int64(n)
	r.buf.Write(p[:n])
	if r.read >= r.errAfter {
		return n, r.err
	}
	return n, err
}

// Bytes returns read bytes.
func (w *MaxReader) Bytes() []byte {
	return w.buf.Bytes()
}

// String returns read bytes in string.
func (w *MaxReader) String() string {
	return w.buf.String()
}
