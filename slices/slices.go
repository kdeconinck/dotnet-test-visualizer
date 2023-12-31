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

// Package slices defines functions useful with slices of any type.
package slices

// Equal reports whether 2 slices are equal (the same length and all elements equal).
// If the lengths are different, Equal returns false.
// Otherwise, the elements are compared in increasing index order, and the comparison stops at the first unequal pair.
func Equal[S ~[]E, E comparable](s1, s2 S) bool {
	if len(s1) != len(s2) {
		return false
	}

	for idx, el := range s1 {
		if el != s2[idx] {
			return false
		}
	}

	return true
}

// Contains reports whether S contains want.
func Contains[S ~[]E, E comparable](s S, want E) bool {
	for _, el := range s {
		if el == want {
			return true
		}
	}

	return false
}

// ContainsFn reports whether el exists in s using a custom comparison function.
// For each element in s, cmpFn is invoked and when it returns true, ContainsFn returns true.
func ContainsFn[S ~[]E, E any](s S, want E, cmpFn func(E, E) bool) bool {
	for _, el := range s {
		if cmpFn(el, want) {
			return true
		}
	}

	return false
}
