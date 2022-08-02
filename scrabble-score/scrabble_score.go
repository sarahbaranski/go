package scrabble

import "strings"

func Score(word string) int {
	l := map[string]int{"A": 1,
		"B": 3, "C": 3, "D": 2,
		"E": 1, "F": 4, "G": 2,
		"H": 4, "I": 1, "J": 8,
		"K": 5, "L": 1, "M": 3,
		"N": 1, "O": 1, "P": 3,
		"Q": 10, "R": 1, "S": 1,
		"T": 1, "U": 1, "V": 4,
		"W": 4, "X": 8, "Y": 4,
		"Z": 10,
	}
	sum := 0
	for i := 0; i < len(word); i++ {
		key := strings.ToUpper(string(word[i]))
		if l[key] > 0 {
			sum += l[key]
		}
	}
	return sum
}
