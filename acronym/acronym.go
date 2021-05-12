// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

func Abbreviate(s string) string {

	s = strings.Replace(s, "-", " ", -1)
	s = strings.Replace(s, "_", " ", -1)

	text := ""
	// Fields fonksiyonunun slice'ında bulunan cümlelerin 0. indexini text string'ine ekliyoruz.
	for _, v := range strings.Fields(s) {
		text += string(v[0])
	}
	return strings.ToUpper(text)
}
