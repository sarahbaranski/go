// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// this gives me a panic runtime error in terminal, but works in go playground
	removes := strings.ReplaceAll(s, "-", " ")
	removes2 := strings.ReplaceAll(removes, "_", "")
	// .Fields instead? but after uppercase and removing dashes/underscores
	words := strings.Fields(removes2)
	res := ""
	for _, word := range words {
		res = res + string([]rune(word)[0])
	}

	return strings.ToUpper(res)
}
