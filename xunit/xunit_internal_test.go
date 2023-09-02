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

// Quality assurance: Verify (and measure the performance) of the public API of the "xunit" package.
package xunit

import (
	"strconv"
	"strings"
	"testing"
)

// Benchmark: Load an XML file containing a .NET test result.
func BenchmarkLoad_MultipleAssemblies(b *testing.B) {
	xmlData := "<assemblies>\n"

	for i := 0; i < 1000; i++ {
		xmlData += "  <assembly />\n"
	}

	xmlData += "</assemblies>"

	benchmarkLoad(xmlData, b)
}

// Benchmark: Load an XML file containing a .NET test result.
func BenchmarkLoad_MultipleTraits(b *testing.B) {
	xmlData := "<assemblies>\n" +
		"  <assembly>\n" +
		"    <collection>\n" +
		"      <test>\n" +
		"        <traits>\n"

	for i := 0; i < 1000; i++ {
		xmlData += "		   <trait name=\"Idx\" value=\"" + strconv.Itoa(i) + "\" />\n"
	}

	xmlData += "        </traits>\n" +
		"      </test>\n" +
		"    </collection>\n" +
		"  </assembly>\n" +
		"</assemblies>"

	benchmarkLoad(xmlData, b)
}

// Benchmark: Load an XML file containing a .NET test result.
func BenchmarkLoad_MultipleTests_MultipleTraits(b *testing.B) {
	xmlData := "<assemblies>\n" +
		"  <assembly>\n" +
		"    <collection>\n"

	for tcIdx := 0; tcIdx < 100; tcIdx++ {
		xmlData += "      <test>\n" +
			"        <traits>\n"

		for i := 0; i < 10; i++ {
			xmlData += "		   <trait name=\"Idx\" value=\"" + strconv.Itoa(i) + "\" />\n"
		}

		xmlData += "        </traits>\n" +
			"      </test>\n"
	}

	xmlData += "    </collection>\n" +
		"  </assembly>\n" +
		"</assemblies>"

	benchmarkLoad(xmlData, b)
}

// Benchmark: Load an XML file containing a .NET test result.
func benchmarkLoad(xmlData string, b *testing.B) {
	// ARRANGE.
	rdr := strings.NewReader(xmlData)
	data, err := unmarshal(rdr)

	if err == nil {
		// RESET.
		b.ResetTimer()

		// EXECUTION.
		for i := 0; i < b.N; i++ {

			_ = readResult(data)
		}
	}
}
