package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime int = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	actual := OvenTime - actualMinutesInOven
	return actual
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	prepTime := numberOfLayers * 2
	return prepTime
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	totalMinutes := (numberOfLayers * 2) + actualMinutesInOven
	return totalMinutes
}
