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

// Package camelcase defines functions for working with "CamelCase" strings.
package camelcase

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/kdeconinck/slices"
)

// NoSplit is a slice of words that should NOT be split by the Split function.
var NoSplit []string

// A reader designed for reading "CamelCase" strings.
type rdr struct {
	input  string // The data this reader operates on.
	pos    int    // The position of this reader.
	rdRune rune   // The last rune that was read.
}

// Read the next rune from r.
func (r *rdr) readRune() {
	if r.pos > len(r.input) {
		r.rdRune = 0
	} else {
		r.rdRune = rune(r.input[r.pos])
	}

	r.pos = r.pos + 1
}

// Undo the last rune from r.
func (r *rdr) unreadRune() {
	r.pos = r.pos - 1
}

// Peek at the next rune from r without advancing the reader.
func (r *rdr) peekRune() rune {
	return rune(r.input[r.pos])
}

// Verify if the word that's currently read by r is a word that should NOT be split.
func (r *rdr) isNoSplitWord(sIdx int) bool {
	return slices.ContainsFn(NoSplit, r.input[sIdx:r.pos+1], func(got, want string) bool { return strings.HasPrefix(got, want) })
}

// Read the current word from r.
// The word is considered terminated as soon as the reader encounters a new uppercase character.
func (r *rdr) readWord() string {
	sIdx := r.pos

	r.readRune()

	if r.pos < len(r.input) && unicode.IsUpper(r.peekRune()) {
		for r.pos < len(r.input) && (unicode.IsUpper(r.peekRune()) || r.isNoSplitWord(sIdx)) {
			r.readRune()
		}

		if r.pos < len(r.input) {
			r.unreadRune()
		}

		return r.input[sIdx:r.pos]
	}

	for r.pos < len(r.input) && (!unicode.IsUpper(r.peekRune()) || r.isNoSplitWord(sIdx)) {
		r.readRune()
	}

	return r.input[sIdx:r.pos]
}

// Split returns the different words from v.
// If v isn't a valid UTF-8 string, or when v is an empty string, a slice with one element (v) is returned.
func Split(v string) []string {
	if !utf8.ValidString(v) || len(v) == 0 {
		return []string{v}
	}

	vRdr := &rdr{input: v}
	retVal := make([]string, 0)

	for vRdr.pos < len(v) {
		retVal = append(retVal, vRdr.readWord())
	}

	return retVal
}
