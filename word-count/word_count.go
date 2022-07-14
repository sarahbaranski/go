package wordcount

import (
	"strings"
)

type Frequency map[string]int

const singleQuote = rune('\'')

func WordCount(fullInput string) Frequency {
	freqMap := Frequency{}
	lower := strings.ToLower(fullInput)
	noColons := strings.Replace(lower, ":", " ", -1)
	noPeriods := strings.Replace(noColons, ".", " ", -1)
	noSpecChar := strings.Replace(noPeriods, "!!&@$%^&", " ", -1)
	wordSlice := strings.Fields(strings.Replace(noSpecChar, ",", " ", -1))
	for _, w := range wordSlice {
		if strings.ContainsRune(w, singleQuote) && !IsContraction(w) {
			w = strings.ReplaceAll(w, string(singleQuote), "")
		}
		freqMap[w]++
	}
	return freqMap
}
func IsContraction(fullInput string) bool {
	index := strings.IndexRune(fullInput, singleQuote)
	return index+1 < len(fullInput) && index-1 > 0
}
