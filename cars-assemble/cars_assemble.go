package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var workingCarPerHour float64 = CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(workingCarPerHour / 60.0)
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	const tenCarsCost int = 95000
	const carCost int = 10000
	return uint(((carsCount / 10) * tenCarsCost) + ((carsCount % 10) * carCost))
}
