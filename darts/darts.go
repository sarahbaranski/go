package darts

import "math"

func Score(x, y float64) int {
	target := math.Pow(x, 2) + math.Pow(y, 2)

	if target <= 100 && target > 25 {
		return 1
	} else if target <= 25 && target > 1 {
		return 5
	} else if target <= 1 && target >= 0 {
		return 10
	} else {
		return 0
	}

}
