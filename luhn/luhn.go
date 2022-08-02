package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(id string) bool {
	idTrim := strings.ReplaceAll(id, " ", "")
	if len(idTrim) <= 1 {
		return false
	}

	if ContainsNonDigit(idTrim) {
		return false
	}

	slice := ConvertToSlice(idTrim)
	for i := len(slice) - 2; i >= 0; i -= 2 {
		indexSlice := slice[i]
		indexSlice *= 2
		if indexSlice > 9 {
			indexSlice -= 9
		}
		slice[i] = indexSlice

	}
	sum := 0
	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}
	return sum%10 == 0
}

func ContainsNonDigit(s string) bool {
	for _, char := range s {
		if !unicode.IsNumber(char) {
			return true
		}
	}
	return false
}
func ConvertToSlice(s string) []int {
	var slice []int
	for i := range s {
		convert := string(s[i])
		v, err := strconv.Atoi(convert)
		if err == nil {
			slice = append(slice, v)
		}
	}
	return slice
}
