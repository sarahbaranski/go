package resistorcolor

func getMap() map[string]int {
	return map[string]int{
		"black":  0,
		"brown":  1,
		"red":    2,
		"orange": 3,
		"yellow": 4,
		"green":  5,
		"blue":   6,
		"violet": 7,
		"grey":   8,
		"white":  9,
	}
}

// Colors returns the list of all colors.
func Colors() []string {
	var list []string
	for k, _ := range getMap() {
		list = append(list, k)
	}

	return list
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	colorValue := 0
	for k, v := range getMap() {
		if k == color{
			colorValue = v
		}
	}
	return colorValue
}
