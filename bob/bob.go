// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

func isAsking(str string) bool {
	nStr := strings.ReplaceAll(str, " ", "")
	mStr := strings.ReplaceAll(nStr, "\t", "")
	xStr := strings.ReplaceAll(mStr, "\r", "")
	yStr := strings.ReplaceAll(xStr, "\n", "")

	return strings.HasSuffix(yStr, "?")
}
func isYelling(str string) bool {
	nStr := strings.ReplaceAll(str, " ", "")
	mStr := strings.ReplaceAll(nStr, "\t", "")
	xStr := strings.ReplaceAll(mStr, "\r", "")
	yStr := strings.ReplaceAll(xStr, "\n", "")

	return strings.IndexFunc(yStr, unicode.IsLetter) >= 0 && yStr == strings.ToUpper(yStr)
}
func isSayingAnything(str string) bool {
	nStr := strings.ReplaceAll(str, " ", "")
	mStr := strings.ReplaceAll(nStr, "\t", "")
	xStr := strings.ReplaceAll(mStr, "\r", "")
	yStr := strings.ReplaceAll(xStr, "\n", "")

	return len(yStr) > 0
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	yelling, asking, saying := isYelling(remark), isAsking(remark), isSayingAnything(remark)
	if yelling && asking {
		return "Calm down, I know what I'm doing!"
	} else if asking {
		return "Sure."
	} else if yelling {
		return "Whoa, chill out!"
	} else if !saying {
		return "Fine. Be that way!"
	}

	return "Whatever."
}
