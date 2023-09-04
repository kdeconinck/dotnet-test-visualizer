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

// Package words defines functions for working with slices of words.
package words

import (
	"strings"

	"github.com/kdeconinck/slices"
)

// NoTransform is a slice of words that should NOT be transformed by the ToSentence function.
var NoTransform []string

// ToSentence returns a slice of words into a sentence.
// Each word in v, except for the first one is converted to lowercase.
func ToSentence(v []string) string {
	retVal := make([]string, 0, len(v))

	for idx, word := range v {
		if idx == 0 {
			retVal = append(retVal, word)
		} else if slices.Contains(NoTransform, word) {
			retVal = append(retVal, word)
		} else {
			retVal = append(retVal, strings.ToLower(word))
		}
	}

	return strings.Join(retVal, " ")
}
