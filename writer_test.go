package tester_test

import (
	"testing"

	"github.com/aileron-projects/go-tester"
)

func TestMaxSilentWriter(t *testing.T) {
	t.Parallel()
	t.Run("n=-5", func(t *testing.T) {
		w := tester.MaxSilentWriter(-5)
		n, err := w.Write([]byte("12345"))
		tester.AssertEqual(t, 0, n)
		tester.AssertEqual(t, nil, err)
		tester.AssertEqual(t, "", w.String())
		tester.AssertEqual(t, "", string(w.Bytes()))
	})
	t.Run("n=0", func(t *testing.T) {
		w := tester.MaxSilentWriter(0)
		n, err := w.Write([]byte("12345"))
		tester.AssertEqual(t, 0, n)
		tester.AssertEqual(t, nil, err)
		tester.AssertEqual(t, "", w.String())
		tester.AssertEqual(t, "", string(w.Bytes()))
	})
	t.Run("n=5", func(t *testing.T) {
		w := tester.MaxSilentWriter(5)
		n, err := w.Write([]byte("12345"))
		tester.AssertEqual(t, 5, n)
		tester.AssertEqual(t, nil, err)
		tester.AssertEqual(t, "12345", w.String())
		tester.AssertEqual(t, "12345", string(w.Bytes()))
	})
	t.Run("write more than n", func(t *testing.T) {
		w := tester.MaxSilentWriter(5)
		var n int
		var err error
		// First write.
		n, err = w.Write([]byte("1234"))
		tester.AssertEqual(t, 4, n)
		tester.AssertEqual(t, nil, err)
		tester.AssertEqual(t, "1234", w.String())
		// Second write.
		n, err = w.Write([]byte("56"))
		tester.AssertEqual(t, 1, n)
		tester.AssertEqual(t, nil, err)
		tester.AssertEqual(t, "12345", w.String())
		// Third write.
		n, err = w.Write([]byte("78"))
		tester.AssertEqual(t, 0, n)
		tester.AssertEqual(t, nil, err)
		tester.AssertEqual(t, "12345", w.String())
		// Write nil.
		n, err = w.Write(nil)
		tester.AssertEqual(t, 0, n)
		tester.AssertEqual(t, nil, err)
		tester.AssertEqual(t, "12345", w.String())
	})
}

func TestMaxErrorWriter(t *testing.T) {
	t.Parallel()
	t.Run("n=-5", func(t *testing.T) {
		w := tester.MaxErrorWriter(-5)
		n, err := w.Write([]byte("12345"))
		tester.AssertEqual(t, 0, n)
		tester.AssertEqual(t, tester.ErrMaxWritten, err)
		tester.AssertEqual(t, "", w.String())
		tester.AssertEqual(t, "", string(w.Bytes()))
	})
	t.Run("n=0", func(t *testing.T) {
		w := tester.MaxErrorWriter(0)
		n, err := w.Write([]byte("12345"))
		tester.AssertEqual(t, 0, n)
		tester.AssertEqual(t, tester.ErrMaxWritten, err)
		tester.AssertEqual(t, "", w.String())
		tester.AssertEqual(t, "", string(w.Bytes()))
	})
	t.Run("n=5", func(t *testing.T) {
		w := tester.MaxErrorWriter(5)
		n, err := w.Write([]byte("12345"))
		tester.AssertEqual(t, 5, n)
		tester.AssertEqual(t, tester.ErrMaxWritten, err)
		tester.AssertEqual(t, "12345", w.String())
		tester.AssertEqual(t, "12345", string(w.Bytes()))
	})
	t.Run("write more than n", func(t *testing.T) {
		w := tester.MaxErrorWriter(5)
		var n int
		var err error
		// First write.
		n, err = w.Write([]byte("1234"))
		tester.AssertEqual(t, 4, n)
		tester.AssertEqual(t, nil, err)
		tester.AssertEqual(t, "1234", w.String())
		// Second write.
		n, err = w.Write([]byte("56"))
		tester.AssertEqual(t, 1, n)
		tester.AssertEqual(t, tester.ErrMaxWritten, err)
		tester.AssertEqual(t, "12345", w.String())
		// Third write.
		n, err = w.Write([]byte("78"))
		tester.AssertEqual(t, 0, n)
		tester.AssertEqual(t, tester.ErrMaxWritten, err)
		tester.AssertEqual(t, "12345", w.String())
		// Write nil.
		n, err = w.Write(nil)
		tester.AssertEqual(t, 0, n)
		tester.AssertEqual(t, nil, err)
		tester.AssertEqual(t, "12345", w.String())
	})
}
