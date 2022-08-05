package pangram

import (
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	splits := strings.FieldsFunc(input, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	inputJoin := strings.Join(splits, "")
	inputJoin = strings.ToLower(inputJoin)
	if len(inputJoin) < 26 {
		return false
	}
	for i := 0; i < len(alpha)-1; i++ {
		r := rune(alpha[i])
		if !strings.ContainsRune(inputJoin, r) {
			return false
		}
	}
	return true
}
