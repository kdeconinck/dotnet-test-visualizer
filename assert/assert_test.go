// =====================================================================================================================
// = LICENSE:       Copyright (c) 2023 Kevin De Coninck
// =
// =                Permission is hereby granted, free of charge, to any person
// =                obtaining a copy of this software and associated documentation
// =                files (the "Software"), to deal in the Software without
// =                restriction, including without limitation the rights to use,
// =                copy, modify, merge, publish, distribute, sublicense, and/or sell
// =                copies of the Software, and to permit persons to whom the
// =                Software is furnished to do so, subject to the following
// =                conditions:
// =
// =                The above copyright notice and this permission notice shall be
// =                included in all copies or substantial portions of the Software.
// =
// =                THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// =                EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// =                OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// =                NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// =                HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// =                WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// =                FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// =                OTHER DEALINGS IN THE SOFTWARE.
// =====================================================================================================================

// Quality assurance: Verify (and measure the performance) of the public API of the "assert" package.
package assert_test

import (
	"fmt"
	"testing"

	"github.com/kdeconinck/assert"
)

// The testableT wraps the testing.T struct and adds a field for storing the failure message.
type testableT struct {
	testing.TB
	failureMsg string
}

// Fatal formats args using fmt.Sprintf and stores the result in t.
func (t *testableT) Fatalf(format string, args ...any) {
	t.failureMsg = fmt.Sprintf(format, args...)
}

// UT: Compare a value against nil.
func TestNil(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	for _, tc := range []struct {
		input any
		name  string
		want  string
	}{
		{
			input: true,
			name:  "ValueOf(true)",
			want:  "ValueOf(true) = true, want <nil>",
		},
		{
			name: "ValueOf(nil)",
		},
	} {
		// ARRANGE.
		testingT := &testableT{TB: t}

		// ACT.
		assert.Nil(testingT, tc.input, tc.name)

		// ASSERT.
		if testingT.failureMsg != tc.want {
			t.Fatalf("Failure message = \"%s\", want \"%s\"", testingT.failureMsg, tc.want)
		}
	}
}

// UT: Compare a value against nil (with a custom message).
func TestNilWithCustomMessage(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// ARRANGE.
	testingT := &testableT{TB: t}

	// ACT.
	assert.Nil(testingT, true, "", "UT Failed: `ValueOf(true)` - got %t, want <nil>.", true)

	// ASSERT.
	if testingT.failureMsg != "UT Failed: `ValueOf(true)` - got true, want <nil>." {
		t.Fatalf("Failure message = \"%s\", want \"%s\"", testingT.failureMsg, "UT Failed: `ValueOf(true)` - got true, want <nil>.")
	}
}

// UT: Compare a value against NOT nil.
func TestNotNil(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	for _, tc := range []struct {
		input any
		name  string
		want  string
	}{
		{
			name: "ValueOf(nil)",
			want: "ValueOf(nil) = <nil>, want NOT <nil>",
		},
		{
			input: true,
			name:  "ValueOf(true)",
		},
	} {
		// ARRANGE.
		testingT := &testableT{TB: t}

		// ACT.
		assert.NotNil(testingT, tc.input, tc.name)

		// ASSERT.
		if testingT.failureMsg != tc.want {
			t.Fatalf("Failure message = \"%s\", want \"%s\"", testingT.failureMsg, tc.want)
		}
	}
}

// UT: Compare a value against NOT nil (with a custom message).
func TestNotNilWithCustomMessage(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// ARRANGE.
	testingT := &testableT{TB: t}

	// ACT.
	assert.NotNil(testingT, nil, "", "UT Failed: `ValueOf(nil)` - got <nil>, want NOT <nil>.")

	// ASSERT.
	if testingT.failureMsg != "UT Failed: `ValueOf(nil)` - got <nil>, want NOT <nil>." {
		t.Fatalf("Failure message = \"%s\", want \"%s\"", testingT.failureMsg, "UT Failed: `ValueOf(nil)` - got <nil>, want NOT <nil>.")
	}
}

// UT: Compare 2 values for equality.
func TestEqual(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	for _, tc := range []struct {
		gInput, wInput bool
		name           string
		want           string
	}{
		{
			gInput: false, wInput: true,
			name: "IsDigit(\"0\")",
			want: "IsDigit(\"0\") = false, want true",
		},
		{
			gInput: true, wInput: true,
			name: "IsDigit(\"0\")",
		},
	} {
		// ARRANGE.
		testingT := &testableT{TB: t}

		// ACT.
		assert.Equal(testingT, tc.gInput, tc.wInput, tc.name)

		// ASSERT.
		if testingT.failureMsg != tc.want {
			t.Fatalf("Failure message = \"%s\", want \"%s\"", testingT.failureMsg, tc.want)
		}
	}
}

// UT: Compare 2 values for equality (with a custom message).
func TestEqualWithCustomMessage(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// ARRANGE.
	testingT := &testableT{TB: t}

	// ACT.
	assert.Equal(testingT, false, true, "", "UT Failed: `IsDigit(\"0\")` - got %t, want %t.", false, true)

	// ASSERT.
	if testingT.failureMsg != "UT Failed: `IsDigit(\"0\")` - got false, want true." {
		t.Fatalf("Failure message = \"%s\", want \"%s\"", testingT.failureMsg, "UT Failed: `ValueOf(true)` - got true, want <nil>.")
	}
}

// UT: Compare 2 slices for equality.
func TestEqualS(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	for _, tc := range []struct {
		gInput, wInput []int
		name           string
		want           string
	}{
		{
			gInput: []int{1, 2, 3},
			wInput: []int{3, 2, 1},
			name:   "Equal([1 2 3], [3 2 1])",
			want:   "Equal([1 2 3], [3 2 1]) = [1 2 3] != [3 2 1]",
		},
		{
			gInput: []int{1, 2, 3},
			wInput: []int{1, 2, 3},
			name:   "Equal([1 2 3], [1 2 3])",
		},
	} {
		// ARRANGE.
		testingT := &testableT{TB: t}

		// ACT.
		assert.EqualS(testingT, tc.gInput, tc.wInput, tc.name)

		// ASSERT.
		if testingT.failureMsg != tc.want {
			t.Fatalf("Failure message = \"%s\", want \"%s\"", testingT.failureMsg, tc.want)
		}
	}
}

// UT: Compare 2 slices for equality (with a custom message).
func TestEqualSWithCustomMessage(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// ARRANGE.
	testingT := &testableT{TB: t}

	// ACT.
	assert.EqualS(testingT, []int{1, 2, 3}, []int{3, 2, 1}, "", "UT Failed: `Equal([1 2 3], [3 2 1])` - %v != %v.", []int{1, 2, 3}, []int{3, 2, 1})

	// ASSERT.
	if testingT.failureMsg != "UT Failed: `Equal([1 2 3], [3 2 1])` - [1 2 3] != [3 2 1]." {
		t.Fatalf("Failure message = \"%s\", want \"%s\"", testingT.failureMsg, "UT Failed: `Equal([1 2 3], [3 2 1])` - [1 2 3] != [3 2 1].")
	}
}

// UT: Compare 2 values for equality using a custom comparison function.
func TestEqualFn(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	for _, tc := range []struct {
		gInput, wInput bool
		name           string
		want           string
	}{
		{
			gInput: false, wInput: true,
			name: "IsDigit(\"0\")",
			want: "IsDigit(\"0\") = false, want true",
		},
		{
			gInput: true, wInput: true,
			name: "IsDigit(\"0\")",
		},
	} {
		// ARRANGE.
		testingT := &testableT{TB: t}

		// ACT.
		assert.EqualFn(testingT, tc.gInput, tc.wInput, func(got, want bool) bool { return got == want }, tc.name)

		// ASSERT.
		if testingT.failureMsg != tc.want {
			t.Fatalf("Failure message = \"%s\", want \"%s\"", testingT.failureMsg, tc.want)
		}
	}
}

// UT: Compare 2 values for equality (with a custom message) using a custom comparison function.
func TestEqualFnWithCustomMessage(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// ARRANGE.
	testingT := &testableT{TB: t}

	// ACT.
	assert.EqualFn(testingT, false, true, func(got, want bool) bool { return false }, "", "UT Failed: `IsDigit(\"0\")` - got %t, want %t.", false, true)

	// ASSERT.
	if testingT.failureMsg != "UT Failed: `IsDigit(\"0\")` - got false, want true." {
		t.Fatalf("Failure message = \"%s\", want \"%s\"", testingT.failureMsg, "UT Failed: `ValueOf(true)` - got true, want <nil>.")
	}
}
