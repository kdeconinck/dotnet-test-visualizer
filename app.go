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

// Package main implements .NET test visualizer, a CLI application for visualizing .NET test result(s).
package main

import (
	"fmt"
	"os"

	"github.com/kdeconinck/camelcase"
	"github.com/kdeconinck/words"
	"github.com/kdeconinck/xunit"
)

// The configuration for the application.
type configuration struct {
	ThresholdFast   float32 `json:"thresholdFast"`
	ThresholdNormal float32 `json:"thresholdNormal"`
}

// The standard configuration for the application.
var stdConfiguration configuration = configuration{
	ThresholdFast:   0.05,
	ThresholdNormal: 0.1,
}

// FindNamed returns the value(s) of a "named" argument if it's found.
// It returns a NON <nil> error if either the "named" argument hasn't been found, or when any of the "named" arguments
// doesn't have a value.
func FindNamed(key string) ([]string, error) {
	var args = os.Args[1:]
	var result = make([]string, 0)

	for idx, arg := range args {
		if arg == key && idx >= len(args)-1 {
			return result, fmt.Errorf("no value found for arg '%s'", key)
		}

		if arg == key && len(args) >= idx {
			valuePos := idx + 1

			result = append(result, args[valuePos])
		}
	}

	if len(result) == 0 {
		return result, fmt.Errorf("arg '%s' not found", key)
	}

	return result, nil
}

// The main entry point for the application.
func main() {
	// Configuration of the application.
	camelcase.NoSplit = []string{"HostBuilder", "DBSyncer", "DbSynchronizer"}
	words.NoTransform = []string{"DbSynchronizer", "DBSyncer"}

	// Prints the ASCII header.
	fmt.Println("    _  _ ___ _____   _____       _    __   ___              _ _            ")
	fmt.Println("   | \\| | __|_   _| |_   _|__ __| |_  \\ \\ / (_)____  _ __ _| (_)______ _ _ ")
	fmt.Println("  _| .` | _|  | |     | |/ -_|_-<  _|  \\ V /| (_-< || / _` | | |_ / -_) '_|")
	fmt.Println(" (_)_|\\_|___| |_|     |_|\\___/__/\\__|   \\_/ |_/__/\\_,_\\__,_|_|_/__\\___|_|  ")
	fmt.Println("")

	// Parse the arguments that are passed to the application.
	lFiles, err := FindNamed("--logFile")

	// If there aren't any LOG files found to process, terminate the application with a failure message.
	if err != nil {
		fmt.Println("\033[1;31mFailed\033[0m: No LOG files found to process.")
		fmt.Println("        Use the `--logFile` argument to pass a file containing logs in xUnit's v2+ XML format.")
		fmt.Println("        If you want to specify multiple files, pass the argument once for each log file.")
		fmt.Println("")

		os.Exit(0)
	}

	// Loop over all the LOG files containing results and parse them.
	for _, logFile := range lFiles {
		rdr, err := os.Open(logFile)

		if err != nil {
			fmt.Printf("\033[1;31mFailed\033[0m - %s\n", err.Error())
		}

		tRun, err := xunit.Load(rdr)

		if err != nil {
			fmt.Printf("\033[1;31mFailed\033[0m - %s\n", err.Error())
		}

		fmt.Printf("Input source:         %s\r\n", logFile)
		fmt.Printf("Amount of assemblies: %v\r\n", len(tRun.Assemblies))

		if tRun.Computer != "" {
			fmt.Printf("Computer:             %s\r\n", tRun.Computer)
		}

		if tRun.User != "" {
			fmt.Printf("User:                 %s\r\n", tRun.User)
		}

		if tRun.StartTimeRTF != "" {
			fmt.Printf("Start time:           %s\r\n", tRun.StartTimeRTF)
		}

		if tRun.EndTimeRTF != "" {
			fmt.Printf("End time:             %s\r\n", tRun.EndTimeRTF)
		} else if tRun.Timestamp != "" {
			fmt.Printf("End time:             %s\r\n", tRun.Timestamp)
		}

		// Loop over the assemblies and print the results accordingly.
		for _, assembly := range tRun.Assemblies {
			fmt.Println("")
			fmt.Printf("  Assembly:         %s", assembly.Name)

			if assembly.FailedCount != 0 {
				fmt.Printf(" - \033[1;31m⛌ Failed (%v of %v failed).\033[0m\r\n", assembly.FailedCount, assembly.TotalCount)
			} else {
				fmt.Printf(" - \033[1;32m✓ Passed (%v of %v passed).\033[0m \r\n", assembly.PassedCount, assembly.TotalCount)
			}

			fmt.Printf("  Date / time:      %s %s\r\n", assembly.RunDate, assembly.RunTime)

			if assembly.TimeRTF != "" {
				fmt.Printf("  Total time:       %v.\r\n", assembly.TimeRTF)
			} else {
				fmt.Printf("  Total time:       %v seconds.\r\n", assembly.Time)
			}

			// Print information about the assembly.
			fmt.Println("")
			fmt.Printf("  # tests:        %v\r\n", assembly.TotalCount)
			fmt.Printf("  # Passed tests: %v\r\n", assembly.PassedCount)
			fmt.Printf("  # Failed tests: %v\r\n", assembly.FailedCount)
			fmt.Printf("  # Errors:       %v\r\n", assembly.ErrorCount)
			fmt.Println("")

			// Loop over all the groups in the assembly.
			for _, tGroup := range assembly.TestGroups {
				if tGroup.Name != "" {
					fmt.Println("")
					fmt.Printf("  Trait: %s\r\n", tGroup.Name)
				}

				// Loop over all the test(s) in this group.
				for _, tc := range tGroup.Tests {
					status := "\033[1;32m✓\033[0m"
					if tc.Result != "Pass" {
						status = "\033[1;31m⛌\033[0m"
					}

					suffix := ""

					if tGroup.Name != "" {
						suffix = "  "
					}

					if tc.Time <= stdConfiguration.ThresholdFast {
						fmt.Printf("%s  🚀 %s %s (%v seconds)\r\n", suffix, status, tc.Name, tc.Time)
					} else if tc.Time <= stdConfiguration.ThresholdNormal {
						fmt.Printf("%s  🕐 %s %s (%v seconds)\r\n", suffix, status, tc.Name, tc.Time)
					} else {
						fmt.Printf("%s  🐌 %s %s (%v seconds)\r\n", suffix, status, tc.Name, tc.Time)
					}
				}

				// Loop over all the groups in this group.
				for _, group := range tGroup.Groups {
					fmt.Println("")
					PrintGroup(group, "")
				}
			}
		}
	}
}

func PrintGroup(group *xunit.TestGroup, indent string) {
	fmt.Printf("%s  %s\r\n", indent, group.Name)

	// Loop over all the test(s) in this group.
	for _, tc := range group.Tests {
		status := "\033[1;32m✓\033[0m"
		if tc.Result != "Pass" {
			status = "\033[1;31m⛌\033[0m"
		}

		if tc.Time <= stdConfiguration.ThresholdFast {
			fmt.Printf("%s     🚀 %s %s (%v seconds)\r\n", indent, status, tc.Name, tc.Time)
		} else if tc.Time <= stdConfiguration.ThresholdNormal {
			fmt.Printf("%s     🕐 %s %s (%v seconds)\r\n", indent, status, tc.Name, tc.Time)
		} else {
			fmt.Printf("%s     🐌 %s %s (%v seconds)\r\n", indent, status, tc.Name, tc.Time)
		}
	}

	if len(group.Tests) > 0 {
		fmt.Println("")
	}

	for _, group = range group.Groups {
		PrintGroup(group, indent+"  ")
	}
}
