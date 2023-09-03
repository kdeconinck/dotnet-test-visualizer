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

// Quality assurance: Verify (and measure the performance) of the public API of the "words" package.
package words_test

import (
	"strconv"
	"testing"

	"github.com/kdeconinck/assert"
	"github.com/kdeconinck/words"
)

// UT: Convert a slice of words into a sentence.
func TestToSentence(t *testing.T) {
	for _, tc := range []struct {
		words []string
		want  string
	}{
		{
			words: []string{},
			want:  "",
		},
		{
			words: []string{"A", "collection", "of", "words"},
			want:  "A collection of words",
		},
		{
			words: []string{"A", "Collection", "Of", "Words", "Starting", "With", "An", "Uppercase", "Character"},
			want:  "A collection of words starting with an uppercase character",
		},
	} {
		// ACT.
		got := words.ToSentence(tc.words)

		// ASSERT.
		assert.Equal(t, got, tc.want, "", "\n\n"+
			"UT Name: Convert a slice of words into a sentence.\n"+
			"Input:   %v\n"+
			"\033[32mExpected:   %v\033[0m\n"+
			"\033[31mActual:     %v\033[0m\n\n", tc.words, got, tc.want)
	}
}

// Benchmark: Convert a slice of words into a sentence.
func BenchmarkToSentence(b *testing.B) {
	// ARRANGE.
	m := make([]string, 0, 1_000_000)

	// Fill the map with 1 million random elements.
	for i := 0; i < 1_000_000; i++ {
		m = append(m, strconv.Itoa(i))
	}

	// RESET.
	b.ResetTimer()

	// EXECUTION.
	for i := 0; i < b.N; i++ {
		// ACT.
		_ = words.ToSentence(m)
	}
}
