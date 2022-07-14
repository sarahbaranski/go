package wordcount

import (
	"strings"
)

type Frequency map[string]int

func WordCount(fullInput string) Frequency {
	freqMap := Frequency{}
	lower := strings.ToLower(fullInput)
	noColons := strings.Replace(lower, ":", " ", -1)
	noPeriods := strings.Replace(noColons, ".", " ", -1)
	noSpecChar := strings.Replace(noPeriods, "!!&@$%^&", " ", -1)
	wordSlice := strings.Fields(strings.Replace(noSpecChar, ",", " ", -1))
	for _, currentWord := range wordSlice {
		freqMap[currentWord]++
	}
	return freqMap
}
