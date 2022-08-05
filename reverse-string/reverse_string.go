package reverse

func Reverse(input string) string {
	newString := ""
	input2 := []rune(input)
	for i := len(input2) - 1; i >= 0; i-- {
		newString += string(input2[i])
	}
	return newString
}
