package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var caseSubject = strings.ToLower(subject)
	var caseCandidates []string
	var empty []string
	for _, word := range candidates {
		caseCandidates = append(caseCandidates, strings.ToLower(word))
	}
	for _, word := range caseCandidates {
		if caseSubject == word {
			return empty
		}
	}
	newSubject := strings.Split(caseSubject, "")
	for i := 0; i < len(caseSubject); i++ {
		sort.Strings(newSubject)
	}
	finalSubject := strings.Join(newSubject, "")
	splitCandidates := make([][]string, len(caseCandidates))
	for i, words := range caseCandidates {
		splitCandidates[i] = strings.Split(words, "")
		sort.Strings(splitCandidates[i])
	}
	var sortedCandidate []string
	for _, word := range splitCandidates {
		sortedCandidate = append(sortedCandidate, strings.Join(word, ""))
	}
	var finalSlice []string
	for i, word := range sortedCandidate {
		if finalSubject == word {
			finalSlice = append(finalSlice, candidates[i])
		}
	}
	return finalSlice
}
