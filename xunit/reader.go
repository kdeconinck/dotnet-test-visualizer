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

// Package xunit contains functions for parsing XML files containing .NET test result(s) in xUnit's v2+ XML format.
// More information regarding this format can be found @ https://xunit.net/docs/format-xml-v2.
package xunit

import (
	"io"
	"strings"

	"github.com/kdeconinck/camelcase"
	"github.com/kdeconinck/maps"
)

// TestRun contains the relevant information stored in xUnit's v2+ XML format.
type TestRun struct {
	Computer     string     // The name of the computer that produced xUnit's v2+ XML format.
	User         string     // The name of the user that produced xUnit's v2+ XML format.
	StartTimeRTF string     // The time the first assembly started running.
	EndTimeRTF   string     // The time the last assembly finished running.
	Timestamp    string     // The time the first assembly started running.
	Assemblies   []Assembly // The assemblies that are part of this test run.
}

// Assembly contains information about the run of a single test assembly.
// This includes environmental information.
type Assembly struct {
	Name        string       // The full name of the assembly.
	ErrorCount  int          // The total number of environmental errors experienced in the assembly.
	PassedCount int          // The total number of test cases in the assembly which passed.
	FailedCount int          // The total number of test cases in the assembly which failed.
	NotRunCount int          // The total number of test cases that weren't run.
	TotalCount  int          // The total number of test cases in the assembly.
	RunDate     string       // The date when the test run started.
	RunTime     string       // The time when the test run started.
	Time        float32      // The number of seconds that the assembly took to run.
	TimeRTF     string       // The time spent running the tests in the assembly.
	TestGroups  []*TestGroup // All the tests of the assembly, grouped by trait.
}

// TestGroup is a group of tests.
type TestGroup struct {
	Name   string       // The name of the group.
	Tests  []TestCase   // The tests that belong to this group.
	Groups []*TestGroup // The subgroups of this group.
}

// TestCase contains information about a single test.
type TestCase struct {
	Name   string  // The name of the test, in human-readable format.
	Result string  // The status of the test.
	Time   float32 // The number of seconds that the test took to run.

	// Internal fields.
	groups []string
}

// Load returns a TestRun constructed from the data in rdr.
func Load(rdr io.Reader) (TestRun, error) {
	data, err := unmarshal(rdr)

	if err != nil {
		return TestRun{}, err
	}

	return readResult(data), nil
}

// Read r into a TestRun.
func readResult(r result) TestRun {
	testRun := TestRun{
		Computer:     r.Computer,
		User:         r.User,
		StartTimeRTF: r.StartRTF,
		EndTimeRTF:   r.FinishRTF,
		Timestamp:    r.Timestamp,
		Assemblies:   make([]Assembly, 0, len(r.Assemblies)),
	}

	// Loop over each assembly.
	for _, assembly := range r.Assemblies {
		testRun.Assemblies = append(testRun.Assemblies, Assembly{
			Name:        assembly.name(),
			ErrorCount:  assembly.ErrorCount,
			PassedCount: assembly.PassedCount,
			FailedCount: assembly.FailedCount,
			NotRunCount: assembly.NotRunCount,
			TotalCount:  assembly.Total,
			RunDate:     assembly.RunDate,
			RunTime:     assembly.RunTime,
			TimeRTF:     assembly.TimeRTF,
			Time:        assembly.Time,
			TestGroups:  assembly.groupTests(),
		})
	}

	return testRun
}

// Returns true if t has a display name, false otherwise.
// When t has any space in its name, it's considered to have display name.
// This ie because by design, C# doesn't allow to have spaces in any identifier and the default name of a test is the
// concatenation (with a `.`) of all identifiers (namespace, class, subclass(es) and methods).
func (t *test) hasDisplayName() bool {
	return strings.Contains(t.Name, " ")
}

// Returns true if t is nested, false otherwise.
// When t has any `+` character in its name and when it's NOT a display name, the test is considered nested.
func (t *test) isNested() bool {
	return !t.hasDisplayName() && strings.Contains(t.Name, "+")
}

// Returns the friendly name of the test.
// If t has a display name, the name is returned as is, if not, the name is split based on the `.` character.
// This gives us a slices of strings where each part contains a valid C# identifier. The last part would be the name of
// the function.
// We feed this name to the "CamelCase" package to turn it into a readable sentence.
func (t *test) friendlyName() string {
	if t.hasDisplayName() {
		return t.Name
	}

	fnName := t.Name[strings.LastIndex(t.Name, ".")+1:]
	fnNameWords := camelcase.Split(fnName)

	retVal := make([]string, 0, len(fnNameWords))

	for idx, word := range fnNameWords {
		if idx == 0 {
			retVal = append(retVal, word)
		} else {
			retVal = append(retVal, strings.ToLower(word))
		}
	}

	return strings.Join(retVal, " ")
}

// Returns the groups that t belongs to.
// If t has a display name, an empty string is returned, if not, the name is split based on the `.` character.
// This gives us a slices of strings where each part contains a valid C# identifier. The last part would be the name of
// the function.
// We feed this name to the "CamelCase" package to turn it into a slice of readable words.
func (t *test) groups() []string {
	if !t.isNested() {
		return make([]string, 0)
	}

	groupName := strings.Split(t.Name, "+")

	groupNameParts := make([]string, 0, len(groupName[1:len(groupName)-1]))
	groupNameParts = append(groupNameParts, groupName[0][strings.LastIndex(groupName[0], ".")+1:])
	groupNameParts = append(groupNameParts, groupName[1:len(groupName)-1]...)
	groupNameParts = append(groupNameParts, strings.Split(groupName[len(groupName)-1], ".")[0])

	groupName = make([]string, 0, len(groupNameParts))

	for _, p := range groupNameParts {
		ccSplit := camelcase.Split(p)

		retVal := make([]string, 0, len(ccSplit))

		for idx, word := range ccSplit {
			if idx == 0 {
				retVal = append(retVal, word)
			} else {
				retVal = append(retVal, strings.ToLower(word))
			}
		}

		groupName = append(groupName, strings.Join(retVal, " "))
	}

	return groupName
}

// Returns the name of the assembly.
func (assembly *assembly) name() string {
	if strings.Contains(assembly.FullName, "/") {
		return assembly.FullName[strings.LastIndex(assembly.FullName, "/")+1:]
	}

	return assembly.FullName[strings.LastIndex(assembly.FullName, "\\")+1:]
}

// Returns a map of tests, grouped per trait.
func (assembly *assembly) groupTests() []*TestGroup {
	if !assembly.hasTests() {
		return make([]*TestGroup, 0)
	}

	uniqueTraits := assembly.uniqueTraits()
	resultSet := make([]*TestGroup, 0, len(uniqueTraits))

	for idx, trait := range uniqueTraits {
		cGroup := &TestGroup{Name: trait, Tests: make([]TestCase, 0, len(assembly.testMap[trait]))}
		resultSet = append(resultSet, cGroup)

		for _, tc := range assembly.testMap[trait] {
			if len(tc.groups) == 0 {
				cGroup.Tests = append(cGroup.Tests, TestCase{Name: tc.Name, Result: tc.Result, Time: tc.Time})
			} else {
				for idx, nn := range tc.groups {
					var sGroup *TestGroup

					for _, group := range cGroup.Groups {
						if group.Name == tc.groups[idx] {
							sGroup = group

							break
						}
					}

					if sGroup == nil {
						sGroup = &TestGroup{Name: nn}
						cGroup.Groups = append(cGroup.Groups, sGroup)
					}

					if idx == len(tc.groups)-1 {
						sGroup.Tests = append(sGroup.Tests, TestCase{Name: tc.Name, Result: tc.Result, Time: tc.Time})
					}

					cGroup = sGroup
				}

				cGroup = resultSet[idx]
			}
		}
	}

	return resultSet
}

// Returns true if the assembly has tests, false otherwise.
func (assembly *assembly) hasTests() bool {
	for _, collection := range assembly.Collections {
		if len(collection.Tests) > 0 {
			return true
		}
	}

	return false
}

// Returns all all the unique trait(s).
func (assembly *assembly) uniqueTraits() []string {
	assembly.testMap = make(map[string][]TestCase)

	for _, collection := range assembly.Collections {
		for _, t := range collection.Tests {
			tCase := TestCase{Name: t.friendlyName(), groups: t.groups(), Result: t.Result, Time: t.Time}

			if len(t.TraitSet.Traits) == 0 {
				assembly.testMap[""] = append(assembly.testMap[""], tCase)
			}

			for _, tTrait := range t.TraitSet.Traits {
				traitName := tTrait.friendlyName()

				assembly.testMap[traitName] = append(assembly.testMap[traitName], tCase)
			}
		}
	}

	return maps.Keys(assembly.testMap)
}

// Returns the friendly name of the trait.
func (t *trait) friendlyName() string {
	var b strings.Builder

	b.WriteString(t.Name)
	b.WriteString(" - ")
	b.WriteString(t.Value)

	return b.String()
}
