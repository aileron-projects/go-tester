package tester

import (
	"cmp"
	"errors"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

// TestingT is an interface wrapper around *[testing.T].
type testingT interface {
	Error(...any)
}

// mocT replaces *testing.T when testing.
var mocT testingT

// AssertEqual asserts that the given values are equal.
// `==` operator is used for comparison.
// AssertEqual returns true if two value are equal.
func AssertEqual[T comparable](t *testing.T, want, got T) bool {
	t.Helper()
	if want == got {
		return true
	}
	msg := "\n"
	msg += "- want: " + spew.Sdump(want)
	msg += "+ got: " + spew.Sdump(got)
	cmp.Or(mocT, testingT(t)).Error(msg)
	return false
}

// AssertDeepEqual asserts that the given values are deeply equal.
// [reflect.DeepEqual] is used for comparison.
// AssertDeepEqual returns true if two value are deeply equal.
func AssertDeepEqual[T any](t *testing.T, want, got T) bool {
	t.Helper()
	if reflect.DeepEqual(want, got) {
		return true
	}
	msg := "\n"
	msg += "- want: " + spew.Sdump(want)
	msg += "+ got: " + spew.Sdump(got)
	cmp.Or(mocT, testingT(t)).Error(msg)
	return false
}

// AssertEqualErr asserts that the given errors are equal.
// [errors.Is] is used for comparison.
// Given got will be unwrapped and compare with want in the [errors.Is].
// AssertEqualErr returns true if two errors are equal.
func AssertEqualErr(t *testing.T, want, got error) bool {
	t.Helper()
	if errors.Is(got, want) {
		return true
	}
	msg := "\n"
	msg += "- want: " + spew.Sdump(want)
	msg += "+ got: " + spew.Sdump(got)
	cmp.Or(mocT, testingT(t)).Error(msg)
	return false
}

// AssertPanic asserts that the test function panics.
// It returns recovered value and true when the test panicked.
// Note tha
func AssertPanic(t *testing.T, test func()) (r any, panicked bool) {
	t.Helper()
	defer func() {
		r = recover()
		if r == nil {
			cmp.Or(mocT, testingT(t)).Error("expect panick but test was not panicked")
			panicked = false
			return
		}
		panicked = true
	}()
	test()
	return nil, false
}
