package isbn

import (
	"strconv"
	"strings"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func IsValidISBN(isbn string) bool {
	newIsbn := strings.Replace(isbn, "-", "", -1)
	if len(newIsbn) != 10 {
		return false
	}
	for _, v := range newIsbn {
		if strings.ContainsRune(letters, v) {
			if v != 'X' {
				return false
			}
		}
	}
	splitIsbn := strings.Split(newIsbn, "")
	if splitIsbn[len(splitIsbn)-1] == "X" {
		splitIsbn[len(splitIsbn)-1] = "10"
	}
	sum := 0
	count := 10
	for i := 0; i < len(splitIsbn); i++ {
		num, _ := strconv.Atoi(splitIsbn[i])
		sum += int(num) * count
		count--
	}
	return sum%11 == 0
}
