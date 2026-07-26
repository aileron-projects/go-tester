package tester_test

import (
	"crypto/rand"
	"os"
	"strings"
	"testing"

	"github.com/aileron-projects/go-tester"
)

func TestReplaceRandReader(t *testing.T) {
	done := tester.ReplaceRandReader(strings.NewReader("12345"))
	defer done()
	b := make([]byte, 5)
	n, err := rand.Read(b)
	tester.AssertEqual(t, "12345", string(b))
	tester.AssertEqual(t, 5, n)
	tester.AssertEqual(t, nil, err)
}

func TestReplaceStdout(t *testing.T) {
	r, done := tester.ReplaceStdout()
	defer done()
	os.Stdout.Write([]byte("12345"))
	b := make([]byte, 5)
	n, err := r.Read(b)
	tester.AssertEqual(t, "12345", string(b))
	tester.AssertEqual(t, 5, n)
	tester.AssertEqual(t, nil, err)
}

func TestReplaceStderr(t *testing.T) {
	r, done := tester.ReplaceStderr()
	defer done()
	os.Stderr.Write([]byte("12345"))
	b := make([]byte, 5)
	n, err := r.Read(b)
	tester.AssertEqual(t, "12345", string(b))
	tester.AssertEqual(t, 5, n)
	tester.AssertEqual(t, nil, err)
}

func TestReplaceStdin(t *testing.T) {
	w, done := tester.ReplaceStdin()
	defer done()
	w.Write([]byte("12345"))
	b := make([]byte, 5)
	n, err := os.Stdin.Read(b)
	tester.AssertEqual(t, "12345", string(b))
	tester.AssertEqual(t, 5, n)
	tester.AssertEqual(t, nil, err)
}
