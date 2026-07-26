package tester

import (
	"crypto/rand"
	"io"
	"os"
)

// ReplaceRandReader replaces the global variable [rand.Reader] with r.
// Call done after tests finished to set the original Reader back.
// Do not run tests parallel until the done was called.
//
// Example:
//
//	done := tester.ReplaceRandReader(YourReader)
//	defer done()
//	// Your test codes here.
func ReplaceRandReader(r io.Reader) (done func()) {
	tmp := rand.Reader
	rand.Reader = r
	return func() {
		rand.Reader = tmp
	}
}

// ReplaceStdout replaces the global variable [os.Stdout] and return reader.
// Call done after tests finished to set the original Stdout back.
// Do not run tests parallel until the done was called.
//
// Example:
//
//	r, done := tester.ReplaceStdout()
//	defer done()
//	// Your test codes here.
func ReplaceStdout() (r *os.File, done func()) {
	tmp := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	return r, func() {
		os.Stdout = tmp
	}
}

// ReplaceStderr replaces the global variable [os.Stderr] and return reader.
// Call done after tests finished to set the original Stderr back.
// Do not run tests parallel until the done was called.
//
// Example:
//
//	r, done := tester.ReplaceStderr()
//	defer done()
//	// Your test codes here.
func ReplaceStderr() (r *os.File, done func()) {
	tmp := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w
	return r, func() {
		os.Stderr = tmp
	}
}

// ReplaceStdin replaces the global variable [os.Stdin] and return writer.
// Call done after tests finished to set the original Stdin back.
// Do not run tests parallel until the done was called.
//
// Example:
//
//	w, done := tester.ReplaceStdin()
//	defer done()
//	// Your test codes here.
func ReplaceStdin() (w *os.File, done func()) {
	tmp := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	return w, func() {
		os.Stdin = tmp
	}
}
