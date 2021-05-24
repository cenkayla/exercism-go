// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

func IsLeapYear(input int) bool {

	return input % 4 == 0 && input % 100 != 0 || input % 400 == 0
}
