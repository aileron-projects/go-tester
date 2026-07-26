package tester_test

import (
	"strings"
	"testing"

	"github.com/aileron-projects/go-tester"
)

func TestMaxSilentReader(t *testing.T) {
	t.Parallel()
	t.Run("n=-5", func(t *testing.T) {
		r := tester.MaxSilentReader(strings.NewReader("12345"), -5)
		buf := make([]byte, 3)
		n, err := r.Read(buf)
		tester.AssertEqual(t, 0, n)
		tester.AssertEqual(t, nil, err)
		tester.AssertEqual(t, "", r.String())
		tester.AssertEqual(t, "", string(r.Bytes()))
		tester.AssertEqual(t, "", string(buf[:n]))
	})
	t.Run("n=0", func(t *testing.T) {
		r := tester.MaxSilentReader(strings.NewReader("12345"), 0)
		buf := make([]byte, 3)
		n, err := r.Read(buf)
		tester.AssertEqual(t, 0, n)
		tester.AssertEqual(t, nil, err)
		tester.AssertEqual(t, "", r.String())
		tester.AssertEqual(t, "", string(r.Bytes()))
		tester.AssertEqual(t, "", string(buf[:n]))
	})
	t.Run("n=5", func(t *testing.T) {
		r := tester.MaxSilentReader(strings.NewReader("12345"), 5)
		buf := make([]byte, 3)
		n, err := r.Read(buf)
		tester.AssertEqual(t, 3, n)
		tester.AssertEqual(t, nil, err)
		tester.AssertEqual(t, "123", r.String())
		tester.AssertEqual(t, "123", string(r.Bytes()))
		tester.AssertEqual(t, "123", string(buf[:n]))
	})
	t.Run("read full", func(t *testing.T) {
		r := tester.MaxSilentReader(strings.NewReader("12345"), 5)
		buf := make([]byte, 5)
		n, err := r.Read(buf)
		tester.AssertEqual(t, 5, n)
		tester.AssertEqual(t, nil, err)
		tester.AssertEqual(t, "12345", r.String())
		tester.AssertEqual(t, "12345", string(r.Bytes()))
		tester.AssertEqual(t, "12345", string(buf[:n]))
	})
	t.Run("read more than n", func(t *testing.T) {
		r := tester.MaxSilentReader(strings.NewReader("123456"), 5)
		var n int
		var err error
		// First read.
		buf1 := make([]byte, 3)
		n, err = r.Read(buf1)
		tester.AssertEqual(t, 3, n)
		tester.AssertEqual(t, nil, err)
		tester.AssertEqual(t, "123", string(buf1[:n]))
		// Second read.
		buf2 := make([]byte, 3)
		n, err = r.Read(buf2)
		tester.AssertEqual(t, 2, n)
		tester.AssertEqual(t, nil, err)
		tester.AssertEqual(t, "45", string(buf2[:n]))
		// Third read.
		buf3 := make([]byte, 3)
		n, err = r.Read(buf3)
		tester.AssertEqual(t, 0, n)
		tester.AssertEqual(t, nil, err)
		tester.AssertEqual(t, "", string(buf3[:n]))
		// Read empty.
		buf4 := make([]byte, 0)
		n, err = r.Read(buf4)
		tester.AssertEqual(t, 0, n)
		tester.AssertEqual(t, nil, err)
		tester.AssertEqual(t, "", string(buf4[:n]))
	})
}

func TestMaxErrorReader(t *testing.T) {
	t.Parallel()
	t.Run("n=-5", func(t *testing.T) {
		r := tester.MaxErrorReader(strings.NewReader("12345"), -5)
		buf := make([]byte, 3)
		n, err := r.Read(buf)
		tester.AssertEqual(t, 0, n)
		tester.AssertEqual(t, tester.ErrMaxRead, err)
		tester.AssertEqual(t, "", r.String())
		tester.AssertEqual(t, "", string(r.Bytes()))
		tester.AssertEqual(t, "", string(buf[:n]))
	})
	t.Run("n=0", func(t *testing.T) {
		r := tester.MaxErrorReader(strings.NewReader("12345"), 0)
		buf := make([]byte, 3)
		n, err := r.Read(buf)
		tester.AssertEqual(t, 0, n)
		tester.AssertEqual(t, tester.ErrMaxRead, err)
		tester.AssertEqual(t, "", r.String())
		tester.AssertEqual(t, "", string(r.Bytes()))
		tester.AssertEqual(t, "", string(buf[:n]))
	})
	t.Run("n=5", func(t *testing.T) {
		r := tester.MaxErrorReader(strings.NewReader("12345"), 5)
		buf := make([]byte, 3)
		n, err := r.Read(buf)
		tester.AssertEqual(t, 3, n)
		tester.AssertEqual(t, nil, err)
		tester.AssertEqual(t, "123", r.String())
		tester.AssertEqual(t, "123", string(r.Bytes()))
		tester.AssertEqual(t, "123", string(buf[:n]))
	})
	t.Run("read full", func(t *testing.T) {
		r := tester.MaxErrorReader(strings.NewReader("12345"), 5)
		buf := make([]byte, 5)
		n, err := r.Read(buf)
		tester.AssertEqual(t, 5, n)
		tester.AssertEqual(t, tester.ErrMaxRead, err)
		tester.AssertEqual(t, "12345", r.String())
		tester.AssertEqual(t, "12345", string(r.Bytes()))
		tester.AssertEqual(t, "12345", string(buf[:n]))
	})
	t.Run("read more than n", func(t *testing.T) {
		r := tester.MaxErrorReader(strings.NewReader("123456"), 5)
		var n int
		var err error
		// First read.
		buf1 := make([]byte, 3)
		n, err = r.Read(buf1)
		tester.AssertEqual(t, 3, n)
		tester.AssertEqual(t, nil, err)
		tester.AssertEqual(t, "123", string(buf1[:n]))
		// Second read.
		buf2 := make([]byte, 3)
		n, err = r.Read(buf2)
		tester.AssertEqual(t, 2, n)
		tester.AssertEqual(t, tester.ErrMaxRead, err)
		tester.AssertEqual(t, "45", string(buf2[:n]))
		// Third read.
		buf3 := make([]byte, 3)
		n, err = r.Read(buf3)
		tester.AssertEqual(t, 0, n)
		tester.AssertEqual(t, tester.ErrMaxRead, err)
		tester.AssertEqual(t, "", string(buf3[:n]))
		// Read empty.
		buf4 := make([]byte, 0)
		n, err = r.Read(buf4)
		tester.AssertEqual(t, 0, n)
		tester.AssertEqual(t, nil, err)
		tester.AssertEqual(t, "", string(buf4[:n]))
	})
}
