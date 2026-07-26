package tester

import (
	"fmt"
	"io"
	"testing"
)

// testT implements the [testingT] interface
// and replaces [mocT] for testing.
type testT struct{}

func (t *testT) Error(args ...any) {}

func TestAssertEqual(t *testing.T) {
	t.Cleanup(func() { mocT = nil })
	t.Run("string", func(t *testing.T) {
		t.Run("equal", func(t *testing.T) {
			tt := &testT{}
			mocT = tt
			equal := AssertEqual(t, "foo", "foo")
			if !equal {
				t.Error("values should be equal")
			}
		})
		t.Run("not equal", func(t *testing.T) {
			tt := &testT{}
			mocT = tt
			equal := AssertEqual(t, "foo", "bar")
			if equal {
				t.Error("values should not be equal")
			}
		})
	})

	t.Run("pointer", func(t *testing.T) {
		t.Run("equal", func(t *testing.T) {
			v1 := &struct{ s string }{"foo"}
			tt := &testT{}
			mocT = tt
			equal := AssertEqual(t, v1, v1)
			if !equal {
				t.Error("values should be equal")
			}
		})
		t.Run("not equal", func(t *testing.T) {
			v1 := &struct{ s string }{"foo"}
			v2 := &struct{ s string }{"foo"}
			tt := &testT{}
			mocT = tt
			equal := AssertEqual(t, v1, v2)
			if equal {
				t.Error("values should not be equal")
			}
		})
	})
}

func TestAssertDeepEqual(t *testing.T) {
	t.Cleanup(func() { mocT = nil })
	t.Run("string", func(t *testing.T) {
		t.Run("equal", func(t *testing.T) {
			tt := &testT{}
			mocT = tt
			equal := AssertDeepEqual(t, "foo", "foo")
			if !equal {
				t.Error("values should be equal")
			}
		})
		t.Run("not equal", func(t *testing.T) {
			tt := &testT{}
			mocT = tt
			equal := AssertDeepEqual(t, "foo", "bar")
			if equal {
				t.Error("values should not be equal")
			}
		})
	})

	t.Run("pointer", func(t *testing.T) {
		t.Run("equal", func(t *testing.T) {
			v1 := &struct{ s string }{"foo"}
			v2 := &struct{ s string }{"foo"}
			tt := &testT{}
			mocT = tt
			equal := AssertDeepEqual(t, v1, v2)
			if !equal {
				t.Error("values should be equal")
			}
		})
		t.Run("not equal", func(t *testing.T) {
			v1 := &struct{ s string }{"foo"}
			v2 := &struct{ s string }{"bar"}
			tt := &testT{}
			mocT = tt
			equal := AssertDeepEqual(t, v1, v2)
			if equal {
				t.Error("values should not be equal")
			}
		})
	})
}

func TestAssertEqualErr(t *testing.T) {
	t.Cleanup(func() { mocT = nil })
	t.Run("equal", func(t *testing.T) {
		tt := &testT{}
		mocT = tt
		equal := AssertEqualErr(t, io.EOF, io.EOF)
		if !equal {
			t.Error("values should be equal")
		}
	})

	t.Run("wrapped got", func(t *testing.T) {
		tt := &testT{}
		mocT = tt
		equal := AssertEqualErr(t, io.EOF, fmt.Errorf("wrap [%w]", io.EOF))
		if !equal {
			t.Error("values should be equal")
		}
	})

	t.Run("wrapped want", func(t *testing.T) {
		tt := &testT{}
		mocT = tt
		equal := AssertEqualErr(t, fmt.Errorf("wrap [%w]", io.EOF), io.EOF)
		if equal {
			t.Error("values should not be equal")
		}
	})

	t.Run("not equal", func(t *testing.T) {
		tt := &testT{}
		mocT = tt
		equal := AssertDeepEqual(t, io.EOF, io.ErrUnexpectedEOF)
		if equal {
			t.Error("values should not be equal")
		}
	})
}
