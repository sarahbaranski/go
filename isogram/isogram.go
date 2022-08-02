package isogram

import (
	"log"
	"regexp"
	"strings"
)

func IsIsogram(word string) bool {
	word2 := strings.ToLower(word)
	if word2 == "" {
		return true
	}
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	pString := reg.ReplaceAllString(word2, "")
	for i := 0; i < len(pString); i++ {
		char := string(pString[i])
		if strings.Count(pString, char) > 1 {
			return false
		}
	}
	return true
}
