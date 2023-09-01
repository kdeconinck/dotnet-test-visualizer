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

// Quality assurance: Verify (and measure the performance) of the public API of the "maps" package.
package maps_test

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/kdeconinck/assert"
	"github.com/kdeconinck/maps"
)

// UT: Get the keys of a map (sorted).
func TestKeys(t *testing.T) {
	for _, tc := range []struct {
		input map[int]bool
		want  []int
	}{
		{
			input: map[int]bool{5: true, 4: false, 3: true, 2: false, 1: true},
			want:  []int{1, 2, 3, 4, 5},
		},
	} {
		// ACT.
		got := maps.Keys(tc.input)

		// ASSERT.
		assert.EqualFn(t, got, tc.want, func(got, want []int) bool { return reflect.DeepEqual(got, want) }, "", "\n\n"+
			"UT Name:    Get the keys of a map (sorted).\n"+
			"Input:      %v\n"+
			"\033[32mExpected:   %v\033[0m\n"+
			"\033[31mActual:     %v\033[0m\n\n", tc.input, tc.want, got)
	}
}

// Benchmark: Get the keys of a map (sorted).
func BenchmarkKeys(b *testing.B) {
	// ARRANGE.
	m := make(map[int]bool, 1_000_000)

	// Fill the map with 1 million random elements.
	for {
		m[rand.Intn(1_000_000_000)] = true

		if len(m) == 1_000_000 {
			break
		}
	}

	// RESET.
	b.ResetTimer()

	// EXECUTION.
	for i := 0; i < b.N; i++ {
		// ACT.
		_ = maps.Keys(m)
	}
}
