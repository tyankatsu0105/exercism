package cars

import "math"

func percentage(number float64) float64 {
	return number / 100
}

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (float64(productionRate) / 100) * successRate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var expression float64 = (float64(productionRate) / 60) * (successRate / 100)

	return int(math.Floor(expression))
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	var tensPlace, onesPlace = carsCount / 10, carsCount % 10

	return uint(tensPlace)*95000 + uint(onesPlace)*10000
}
