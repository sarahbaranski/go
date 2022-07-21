package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var finalCount int

	for i := 0; i < len(birdsPerDay); i++ {
		finalCount = finalCount + birdsPerDay[i]
	}

	return finalCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var weeklyCount int

	if week == 1 {
		for i := 0; i <= 6; i++ {
			weeklyCount = weeklyCount + birdsPerDay[i]
		}
	} else {
		for i := 7; i > 6 && i < 14; i++ {
			weeklyCount = weeklyCount + birdsPerDay[i]
		}
	}

	return weeklyCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i] = birdsPerDay[i] + 1
		}
	}

	return birdsPerDay
}
