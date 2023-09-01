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

// Quality assurance: Verify (and measure the performance) of the public API of the "slices" package.
package slices_test

import (
	"testing"

	"github.com/kdeconinck/assert"
	"github.com/kdeconinck/slices"
)

// UT: Compare 2 slices for equality.
func TestEqual(t *testing.T) {
	for _, tc := range []struct {
		s1Input, s2Input []int
		want             bool
	}{
		{
			s1Input: []int{1},
			s2Input: nil,
			want:    false,
		},
		{
			s1Input: []int{1, 2, 3},
			s2Input: []int{1, 2, 3, 4},
			want:    false,
		},
		{
			s1Input: []int{1, 2, 3},
			s2Input: []int{3, 2, 1},
			want:    false,
		},
		{
			s1Input: []int{},
			s2Input: nil,
			want:    true,
		},
		{
			s1Input: []int{1, 2, 3},
			s2Input: []int{1, 2, 3},
			want:    true,
		},
	} {
		// ACT.
		got := slices.Equal(tc.s1Input, tc.s2Input)

		// ASSERT.
		assert.Equal(t, got, tc.want, "", "\n\n"+
			"UT Name:    Compare 2 slices for equality.\n"+
			"Input (s1): %v\n"+
			"Input (s2): %v\n"+
			"\033[32mExpected:   %v\033[0m\n"+
			"\033[31mActual:     %v\033[0m\n\n", tc.s1Input, tc.s2Input, tc.want, got)
	}
}
