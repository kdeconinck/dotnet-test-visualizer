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

// Quality assurance: Verify (and measure the performance) of the public API of the "camelcase" package.
package camelcase_test

import (
	"strings"
	"testing"

	"github.com/kdeconinck/assert"
	"github.com/kdeconinck/camelcase"
)

// UT: Split a "CamelCase" string.
func TestSplit(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  []string
	}{
		{
			input: "Bad UTF-8 (\xe2\xe2\xa1)",
			want:  []string{"Bad UTF-8 (\xe2\xe2\xa1)"},
		},
		{
			input: "",
			want:  []string{""},
		},
		{
			input: "FullyQualifiedName",
			want:  []string{"Fully", "Qualified", "Name"},
		},
		{
			input: "PDFLoader",
			want:  []string{"PDF", "Loader"},
		},
	} {
		// ACT.
		got := camelcase.Split(tc.input)

		// ASSERT.
		assert.EqualS(t, got, tc.want, "", "\n\n"+
			"UT Name:    Split a \"CamelCase\" string.\n"+
			"Input:      %v\n"+
			"\033[32mExpected:   %v\033[0m\n"+
			"\033[31mActual:     %v\033[0m\n\n", tc.input, tc.want, got)
	}
}

// Benchmark: Split a "CamelCase" string.
func BenchmarkSplit(b *testing.B) {
	// ARRANGE.
	var s strings.Builder

	for i := 0; i < 1_000_000; i++ {
		s.WriteString("HelloWorld")
	}

	input := s.String()

	// RESET.
	b.ResetTimer()

	// EXECUTION.
	for i := 0; i < b.N; i++ {
		_ = camelcase.Split(input)
	}
}
