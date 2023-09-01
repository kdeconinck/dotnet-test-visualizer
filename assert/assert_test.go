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

// UT: Compare a value against nil.
func BenchmarkNil(b *testing.B) {
	// ARRANGE.
	testingT := &testableT{TB: b}

	// RESET.
	b.ResetTimer()

	// EXECUTION.
	for i := 0; i < b.N; i++ {
		// ACT.
		assert.Nil(testingT, true, "BenchmarkNil")
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

// UT: Compare a value against nil (with a custom message).
func BenchmarkNilWithCustomMessage(b *testing.B) {
	// ARRANGE.
	testingT := &testableT{TB: b}

	// RESET.
	b.ResetTimer()

	// EXECUTION.
	for i := 0; i < b.N; i++ {
		// ACT.
		assert.Nil(testingT, true, "", "UT Failed: `ValueOf(true)` - got %t, want <nil>.", true)
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

// UT: Compare a value against NOT nil.
func BenchmarkNotNil(b *testing.B) {
	// ARRANGE.
	testingT := &testableT{TB: b}

	// RESET.
	b.ResetTimer()

	// EXECUTION.
	for i := 0; i < b.N; i++ {
		// ACT.
		assert.NotNil(testingT, nil, "BenchmarkNotNil")
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

// UT: Compare a value against NOT nil (with a custom message).
func BenchmarkNotNilWithCustomMessage(b *testing.B) {
	// ARRANGE.
	testingT := &testableT{TB: b}

	// RESET.
	b.ResetTimer()

	// EXECUTION.
	for i := 0; i < b.N; i++ {
		// ACT.
		assert.NotNil(testingT, nil, "", "UT Failed: `ValueOf(nil)` - got <nil>, want NOT <nil>.")
	}
}
